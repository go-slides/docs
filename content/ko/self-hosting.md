---
title: "셀프 호스팅"
weight: 58
description: "Caddy로 Playground와 slides --server를 자신의 컴퓨터에서 서비스합니다. 자동 HTTPS, Playground에 필요한 헤더, 앞단의 OpenID Connect 로그인까지 다룹니다."
tags: [playground, server, security]
---

[Playground]({{< relref "playground" >}})는 정적 파일이고,
[`slides --server`]({{< relref "cli" >}})는 루프백에서 동작하는 하나의 프로세스입니다. 둘 다
자신의 컴퓨터에서, 자신의 이름으로 실행해 조직의 계정으로 로그인하는 사람들에게
제공할 수 있습니다. 이 페이지에서는 HTTPS 인증서를 직접 받고 갱신하는
[Caddy](https://caddyserver.com/docs/)(v2.11.7로 테스트)와, 로그인을 맡는
[oauth2-proxy](https://oauth2-proxy.github.io/oauth2-proxy/)를 사용합니다.
여기에 나오는 모든 파일은
[`selfhost/`](https://github.com/go-slides/go-slides.github.io/tree/main/selfhost)에
있으며, 그곳의 CI가 각 Caddyfile을 `caddy validate`로 검사합니다.

## Playground {#the-playground}

사이트의 각 릴리스에는 Playground를 포함한 사이트 전체가
도메인 루트용으로 빌드된 tarball 하나로, SHA-256과 함께 첨부됩니다.

```sh
curl -fLO https://github.com/go-slides/go-slides.github.io/releases/download/v0.1.0/go-slides-site-v0.1.0.tar.gz
curl -fLO https://github.com/go-slides/go-slides.github.io/releases/download/v0.1.0/go-slides-site-v0.1.0.tar.gz.sha256
sha256sum -c go-slides-site-v0.1.0.tar.gz.sha256
tar -xzf go-slides-site-v0.1.0.tar.gz -C /srv
mv /srv/go-slides-site-v0.1.0 /srv/go-slides-site
```

이 tarball은 재현 가능합니다. 같은 태그를 다시 빌드하면 같은
바이트가 나옵니다. 압축되는 모든 파일(WebAssembly 모듈, 글꼴, 페이지, 스크립트) 옆에는
같은 파일을 Brotli(`.br`)와 gzip(`.gz`)으로 압축한 것도 들어 있어서, 서버는 이를
그대로 보내고 즉석에서는 아무것도 압축하지 않습니다.

### Caddyfile {#caddyfile}

`slides.example.org`는 여러분의 이름입니다. 이 이름의 A 레코드와 AAAA 레코드가 해당
컴퓨터를 가리키고, 80번과 443번 포트가 Caddy에 도달해야 합니다.

```caddy
import headers.caddy

slides.example.org {
	root * /srv/go-slides-site
	file_server {
		precompressed br gzip
	}
	import playground_headers
}
```

```sh
caddy run --config Caddyfile.playground
```

그러면 Playground는 `https://slides.example.org/playground/`에서 열립니다.

### 서버가 올바르게 해야 하는 일 {#what-the-server-must-get-right}

- **HTTPS.** Caddy는 ACME([RFC 8555](https://www.rfc-editor.org/rfc/rfc8555))로
  Let's Encrypt에서 인증서를 받습니다. 80번 포트(HTTP-01) 또는 443번 포트(TLS-ALPN-01)에서
  응답해 이름을 가지고 있음을 증명하고, 인증서를 갱신합니다. CA가 접근할 수 없는 컴퓨터는 대신
  DNS-01 챌린지(TXT 레코드이며, DNS 제공업체의 모듈이 들어간 Caddy가 필요합니다)나
  `tls cert.pem key.pem`으로 지정한 자체 인증서를 쓸 수 있습니다.
  테스트용 컴퓨터에서는 `tls internal`이 Caddy의 로컬 CA를 사용합니다.
- **`application/wasm`.** Caddy는 `.wasm`에 이 타입을 보냅니다. 다른 서버는
  따로 지정해야 할 수 있습니다. Playground는 읽어 들인 바이트에서 모듈을 컴파일하므로
  (내보내기 엔진은 내려받는 동안, 이 타입의 별도 응답으로), 다른 타입을 보내는 서버에서도
  시작되지만, `nosniff`가 있더라도 올바른 타입을 보내는 것이 맞습니다.
- **압축.** `precompressed br gzip`은 Brotli를 받는 브라우저에는 `playground.wasm.br`을,
  gzip을 받는 브라우저에는 `.gz`를, 둘 다 아니면 파일 자체를 `Vary: Accept-Encoding`과
  함께 보냅니다. 편집기 모듈은 Brotli로 약 7 MiB이고, 압축하지 않으면
  33 MiB입니다.
- **캐싱.** 페이지는 모든 모듈, 워커, 글꼴, 그림을 쿼리에 릴리스를 붙여
  (`playground.wasm?v=v0.1.0`) 요청하므로, 이들은 1년 동안
  보관할 수 있습니다(`immutable`). 페이지 자체는 매번 재검증됩니다
  (`no-cache`). 새 릴리스는 다음 방문 때 보입니다.
- **`headers.caddy`의 헤더:**

```caddy
(playground_headers) {
	header {
		Content-Security-Policy "default-src 'self'; script-src 'self' 'unsafe-inline' 'wasm-unsafe-eval'; worker-src 'self'; connect-src 'self' https:; img-src 'self' data: blob: https:; media-src 'self' data: blob: https:; style-src 'self' 'unsafe-inline' https:; font-src 'self' data: blob: https:; frame-src 'self'; object-src 'none'; base-uri 'none'; form-action 'none'; frame-ancestors 'self'"
		X-Content-Type-Options nosniff
		Referrer-Policy no-referrer
		Strict-Transport-Security "max-age=31536000"
	}
	@versioned query v=*
	header @versioned Cache-Control "public, max-age=31536000, immutable"
	@unversioned not query v=*
	header @unversioned Cache-Control "no-cache"
}
```

이 정책은 Playground가 필요로 하는 그대로입니다. 스크립트(로더는
인라인), WebAssembly(`'wasm-unsafe-eval'`), 워커가 모두
여러분의 서버에서 옵니다. 덱의 그림, 스타일시트, 글꼴은 명령의 출력과 마찬가지로
어떤 `https:` 주소에서든 올 수 있습니다. 모든 요청을 여러분의 서버에 두려면
`img-src`, `style-src`, `font-src`, `media-src`, `connect-src`에서 `https:`를 빼십시오.
그러면 다른 곳을 가리키는 덱은 그 리소스 없이 표시됩니다. 덱 자체는 여러분의 출처가 없는
샌드박스 iframe 안에서, Playground가 각 프레임에 써 넣는 더 엄격한 자체 정책 아래 실행됩니다.
이 헤더가 무엇이라 하든, 덱은 Playground, 그 저장소, 여러분의 쿠키에 접근할 수
없습니다. `frame-ancestors 'self'`는 다른 사이트가 Playground를 프레임에 넣지
못하게 합니다.

- **COOP도 COEP도 없음.** Playground는 `SharedArrayBuffer`도 스레드도 쓰지 않으므로
  교차 출처 격리가 필요하지 않습니다.
  `Cross-Origin-Embedder-Policy: require-corp`를 보내면 덱이 다른 사이트에서 가져오는 그림이
  차단됩니다.

### 네트워크 없이 {#without-the-network}

페이지를 불러온 뒤에는 편집기, 미리보기, 발표자 보기, 그리고 PDF·PPTX·ODP 내보내기
(글꼴은 tarball에 들어 있습니다)까지 모든 것이 브라우저 안에서 일어납니다. 덱은 여러분의
서버로도, 다른 어디로도 전송되지 않습니다. 네트워크가 필요한 것은 덱이 직접 가리키는 것들, 즉
`https:` 주소의 그림이나 스타일시트와, `gaia` 테마의 웹 글꼴입니다. 이 글꼴은
`fonts.bunny.net`에 요청합니다(없으면 슬라이드는 다른 글꼴로 대체됩니다).
이 문서로 향하는 링크는 여러분의 서버를 벗어납니다.

## 프록시 뒤의 `slides --server` {#slides---server-behind-the-proxy}

`slides --server`는 `127.0.0.1`에 머물고 Caddy가 그쪽으로 전달합니다. 서버는 DNS 리바인딩과
다른 사이트의 페이지를 막기 위해 각 요청의 `Host` 헤더와 각 라이브 리로드
WebSocket의 `Origin`을 검사하며, 프록시는 둘 다에 공개 이름을 유지합니다.
`--allowed-host`는 그 이름을 검사가 받아들이는 대상에 추가합니다. 검사는 켜진 채로 있습니다.

```sh
slides -s --allowed-host decks.example.org /srv/decks
```

```caddy
decks.example.org {
	reverse_proxy 127.0.0.1:8080
}
```

라이브 리로드는 그대로 통과합니다(Caddy는 설정 없이 WebSocket을 프록시합니다).
이를 위해 `--host 0.0.0.0`을 쓰지 마십시오. 디렉터리를 네트워크 전체에 제공하게 되어
프록시와 로그인을 우회하게 됩니다.

{{< callout type="info" >}}
`--allowed-host`는 v0.13.0 이후에 들어왔습니다. 다음 릴리스가 나올 때까지는
이 옵션이 포함된 커밋을 설치하십시오.
`go install github.com/go-slides/slides/cmd/slides@v0.13.1-0.20261026090714-2e1da79f663e`.
{{< /callout >}}

## OpenID Connect로 로그인 {#sign-in-with-openid-connect}

Caddy는 모든 요청에 대해 oauth2-proxy에 물어봅니다(`forward_auth`). oauth2-proxy는
로그인한 브라우저에는 `202`로 답하고 Caddy가 그 요청을 처리하며,
그렇지 않으면 `401`로 답하고 Caddy가 이를 로그인으로의 리디렉션으로 바꿉니다.
로그인은
[OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html)의
인증 코드 흐름입니다. 브라우저는 제공자(Keycloak, Authentik, Entra ID, Google,
Okta, Dex 등)에게 갔다가 코드를 가지고 `/oauth2/callback`으로 돌아오고, oauth2-proxy가
그 코드를 ID Token으로 교환해 검증한 뒤 자체 쿠키를 설정합니다.

1. 제공자에 **클라이언트를 등록**합니다. 기밀 클라이언트이며 인증 코드 흐름을 사용합니다.
   리디렉션 URI는 정확히 일치해야 합니다.
   `https://slides.example.org/oauth2/callback`과
   `https://decks.example.org/oauth2/callback`입니다.
2. **비밀 값은 파일에 두고** oauth2-proxy만 읽을 수 있게 합니다. 클라이언트의
   비밀 값(끝에 줄바꿈 없음)과, 정확히 32바이트의 무작위 쿠키
   비밀 값(`head -c 32 /dev/urandom`)입니다. 비밀 값 저장소(systemd credentials,
   Docker나 Kubernetes secret, Vault 등)가 이를 기록합니다.
   둘 다 설정 파일, 명령줄, 환경 변수에는 넣지 마십시오. 로그와 `ps`에 노출됩니다.
3. **oauth2-proxy를 설정**합니다(`oauth2-proxy.cfg`).

```toml
http_address = "127.0.0.1:4180"
reverse_proxy = true
trusted_proxy_ips = ["127.0.0.1/32", "::1/128"]
upstreams = ["static://202"]
set_xauthrequest = true

provider = "oidc"
oidc_issuer_url = "https://id.example.org/realms/example"
client_id = "go-slides"
client_secret_file = "/run/secrets/oauth2-proxy-client-secret"
code_challenge_method = "S256"

redirect_url = "/oauth2/callback"
whitelist_domains = ["slides.example.org", "decks.example.org"]
email_domains = ["example.org"]

cookie_secret_file = "/run/secrets/oauth2-proxy-cookie-secret"
cookie_secure = true
cookie_samesite = "lax"
```

호스트가 없는 `redirect_url`은 각 사이트의 로그인을 그 사이트 자신의 콜백으로 되돌려 보내므로,
각 사이트가 자기 쿠키를 가집니다. `whitelist_domains`는 로그인 후 돌아갈 수 있는 곳입니다.
다른 곳으로 가라고 요청받아도 여러분의 사이트에 머뭅니다.
`trusted_proxy_ips`는 Caddy가 보낸 `X-Forwarded-*` 헤더만 믿습니다.
`email_domains`는 누가 들어올 수 있는지를 정합니다. oauth2-proxy에는 그룹과
주소 목록을 위한 옵션도 있습니다.

4. **Caddy를 설정**합니다(`Caddyfile.oidc`).

```caddy
import headers.caddy

(signed_in) {
	handle /oauth2/* {
		reverse_proxy 127.0.0.1:4180
	}
	handle {
		forward_auth 127.0.0.1:4180 {
			uri /oauth2/auth
			copy_headers X-Auth-Request-User X-Auth-Request-Email
			@unauthorized status 401
			handle_response @unauthorized {
				redir * /oauth2/sign_in?rd={scheme}://{hostport}{uri} 302
			}
		}
		{block}
	}
}

slides.example.org {
	import signed_in {
		root * /srv/go-slides-site
		file_server {
			precompressed br gzip
		}
		import playground_headers
	}
}

decks.example.org {
	import signed_in {
		reverse_proxy 127.0.0.1:8080
	}
}
```

```sh
slides -s --allowed-host decks.example.org /srv/decks
oauth2-proxy --config /etc/oauth2-proxy/oauth2-proxy.cfg
caddy run --config Caddyfile.oidc
```

oauth2-proxy와 `slides`는 루프백에서만 수신합니다. 들어오는 길은 Caddy뿐이며,
따라서 로그인을 반드시 거칩니다. Playground의 헤더는 로그인된 블록 안쪽에,
Playground에만 둡니다. oauth2-proxy의 로그인 페이지는 폼인데,
이 헤더의 `form-action 'none'`이 그것을 막기 때문입니다.

대안은 [caddy-security](https://github.com/greenpau/caddy-security)입니다.
Caddy 안에서 로그인을 처리하는 Caddy 모듈로, 두 번째 프로세스가 필요 없습니다.
다만 이 모듈을 넣어 빌드한 Caddy(`xcaddy`)와 자체 설정 언어가 필요합니다.

## 확인하기 {#check-it}

```sh
curl -sI -H 'Accept-Encoding: br' 'https://slides.example.org/playground/playground.wasm?v=v0.1.0'
```

응답에는 `content-type: application/wasm`, `content-encoding: br`,
`cache-control: public, max-age=31536000, immutable`이 나오며, 위에서 본
`content-security-policy`도 함께 옵니다. 앞단에 로그인이 있으면, 같은
요청을 쿠키 없이 보낼 때 `302`로 `/oauth2/sign_in`에 리디렉션됩니다.

## 참고 자료 {#references}

- Caddy: [automatic HTTPS](https://caddyserver.com/docs/automatic-https),
  [`file_server`](https://caddyserver.com/docs/caddyfile/directives/file_server),
  [`header`](https://caddyserver.com/docs/caddyfile/directives/header),
  [`reverse_proxy`](https://caddyserver.com/docs/caddyfile/directives/reverse_proxy),
  [`forward_auth`](https://caddyserver.com/docs/caddyfile/directives/forward_auth)
- [RFC 8555](https://www.rfc-editor.org/rfc/rfc8555), ACME: HTTP-01(8.3),
  DNS-01(8.4)
- [OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html):
  인증 코드 흐름(3.1), 리디렉션 URI(3.1.2.1), ID
  Token(2)
- oauth2-proxy: [configuration](https://oauth2-proxy.github.io/oauth2-proxy/configuration/overview),
  [OpenID Connect provider](https://oauth2-proxy.github.io/oauth2-proxy/configuration/providers/openid_connect)
