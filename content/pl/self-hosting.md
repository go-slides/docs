---
title: "Self-hosting"
weight: 58
description: "Serwuj Playground i slides --server z własnej maszyny za pomocą Caddy: automatyczny HTTPS, nagłówki potrzebne dla Playground i logowanie OpenID Connect z przodu."
tags: [playground, server, security]
---

[Playground]({{< relref "playground" >}}) to pliki statyczne, a
[`slides --server`]({{< relref "cli" >}}) to jeden proces na pętli zwrotnej. Oba
mogą działać na twojej maszynie, pod twoją nazwą, dla osób, które logują się
kontem twojej organizacji. Ta strona robi to za pomocą
[Caddy](https://caddyserver.com/docs/) (testowane z v2.11.7), który sam uzyskuje
i odnawia certyfikaty HTTPS, oraz
[oauth2-proxy](https://oauth2-proxy.github.io/oauth2-proxy/) do
logowania. Każdy pokazany tu plik jest w
[`selfhost/`](https://github.com/go-slides/go-slides.github.io/tree/main/selfhost),
a jego CI sprawdza każdy Caddyfile poleceniem `caddy validate`.

## Playground {#the-playground}

Każde wydanie witryny dołącza całą witrynę, wraz z Playground, jako
jeden tarball zbudowany dla korzenia domeny, z jego SHA-256:

```sh
curl -fLO https://github.com/go-slides/go-slides.github.io/releases/download/v0.1.0/go-slides-site-v0.1.0.tar.gz
curl -fLO https://github.com/go-slides/go-slides.github.io/releases/download/v0.1.0/go-slides-site-v0.1.0.tar.gz.sha256
sha256sum -c go-slides-site-v0.1.0.tar.gz.sha256
tar -xzf go-slides-site-v0.1.0.tar.gz -C /srv
mv /srv/go-slides-site-v0.1.0 /srv/go-slides-site
```

Tarball jest odtwarzalny: ponowne zbudowanie tego samego tagu daje te same
bajty. Obok każdego pliku, który się kompresuje (moduły WebAssembly,
czcionki, strony, skrypty), zawiera ten sam plik skompresowany
Brotli (`.br`) i gzip (`.gz`), więc serwer wysyła je tak, jak są, i
niczego nie kompresuje w locie.

### Caddyfile {#caddyfile}

`slides.example.org` to twoja nazwa. Jej rekordy A i AAAA wskazują
maszynę, a porty 80 i 443 docierają do Caddy:

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

Playground jest wtedy pod adresem `https://slides.example.org/playground/`.

### Co serwer musi zrobić dobrze {#what-the-server-must-get-right}

- **HTTPS.** Caddy dostaje certyfikat od Let's Encrypt przez ACME
  ([RFC 8555](https://www.rfc-editor.org/rfc/rfc8555)), dowodząc, że posiada
  nazwę, odpowiadając na porcie 80 (HTTP-01) lub 443 (TLS-ALPN-01), i
  go odnawia. Maszyna, do której CA nie ma dostępu, może użyć zamiast tego
  wyzwania DNS-01 (rekord TXT, wymagający Caddy zbudowanego z modułem twojego
  dostawcy DNS) albo własnego certyfikatu z `tls cert.pem key.pem`.
  Na maszynie testowej `tls internal` używa lokalnego CA Caddy.
- **`application/wasm`.** Caddy wysyła ten typ dla `.wasm`; inny serwer
  może wymagać, by mu to powiedzieć. Playground kompiluje swoje moduły z odczytanych
  bajtów (silnik eksportu w trakcie pobierania, w osobnej odpowiedzi
  z tym typem), więc startuje nawet z serwera wysyłającego inny typ;
  przy `nosniff` właściwy typ nadal jest tym, który należy wysłać.
- **Kompresja.** `precompressed br gzip` wysyła `playground.wasm.br` do
  przeglądarki akceptującej Brotli, `.gz` do akceptującej gzip, a w pozostałych przypadkach sam
  plik, z `Vary: Accept-Encoding`. Moduł edytora ma
  około 7 MiB w Brotli wobec 33 MiB bez kompresji.
- **Cache.** Strona żąda każdego modułu, workera, czcionki i obrazu
  z wydaniem w zapytaniu (`playground.wasm?v=v0.1.0`), więc można je
  trzymać przez rok (`immutable`). Same strony są rewalidowane
  (`no-cache`): nowe wydanie jest widoczne przy następnej wizycie.
- **Nagłówki w `headers.caddy`:**

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

Polityka jest taka, jakiej potrzebuje Playground: jego skrypty (loader jest
wbudowany), jego WebAssembly (`'wasm-unsafe-eval'`) i jego worker, wszystko z
twojego serwera. Obrazy, arkusze stylów i czcionki prezentacji mogą pochodzić z dowolnego
adresu `https:`, tak jak w wyniku polecenia; usuń `https:` z
`img-src`, `style-src`, `font-src`, `media-src` i `connect-src`, by zatrzymać
każde żądanie na twoim serwerze, a prezentacje wskazujące gdzie indziej pokażą się bez
tych zasobów. Sama prezentacja działa w iframe z piaskownicą bez twojego
pochodzenia, pod własną, surowszą polityką, którą Playground zapisuje
w każdej ramce: nie sięgnie do Playground, jego pamięci ani twoich
ciasteczek, cokolwiek mówi ten nagłówek. `frame-ancestors 'self'` powstrzymuje inne
witryny przed osadzaniem Playground w ramce.

- **Bez COOP ani COEP.** Playground nie używa ani `SharedArrayBuffer`, ani
  wątków, więc nie musi być odizolowany między źródłami. Wysłanie
  `Cross-Origin-Embedder-Policy: require-corp` zablokowałoby obrazy prezentacji
  z innych witryn.

### Bez sieci {#without-the-network}

Po załadowaniu strony wszystko dzieje się w przeglądarce: edytor,
podgląd, widok prelegenta i eksporty PDF, PPTX i ODP, których
czcionki są w tarballu. Prezentacja nie jest wysyłana ani na twój serwer, ani
nigdzie indziej. Sieci potrzebuje to, co nazywa prezentacja: obrazy lub
arkusze stylów pod adresem `https:` oraz czcionka internetowa motywu `gaia`, którą
pobiera z `fonts.bunny.net` (bez niej slajdy wracają do innej
czcionki). Odnośniki do tej dokumentacji wychodzą z twojego serwera.

## `slides --server` za proxy {#slides---server-behind-the-proxy}

`slides --server` zostaje na `127.0.0.1`, a Caddy przekazuje do niego żądania. Serwer
sprawdza nagłówek `Host` każdego żądania i `Origin` każdego WebSocketu
przeładowania na żywo, przeciw DNS rebinding i stronom innych witryn, a
proxy zachowuje w obu publiczną nazwę. `--allowed-host` dodaje tę nazwę do
tego, co kontrole akceptują; pozostają włączone:

```sh
slides -s --allowed-host decks.example.org /srv/decks
```

```caddy
decks.example.org {
	reverse_proxy 127.0.0.1:8080
}
```

Przeładowanie na żywo przechodzi bez zmian (Caddy proxuje WebSockety bez
konfiguracji). Nie używaj do tego `--host 0.0.0.0`: serwuje
katalog całej sieci, z pominięciem proxy i jego logowania.

{{< callout type="info" >}}
`--allowed-host` jest w v0.14.0 i nowszych:
`go install github.com/go-slides/slides/cmd/slides@v0.15.0`. Wymaga Go 1.27.3
lub nowszego.
{{< /callout >}}

## Logowanie przez OpenID Connect {#sign-in-with-openid-connect}

Caddy pyta oauth2-proxy o każde żądanie (`forward_auth`). oauth2-proxy
odpowiada `202` dla przeglądarki, która się zalogowała, co Caddy wtedy serwuje,
a w przeciwnym razie `401`, które Caddy zamienia na przekierowanie do logowania.
Logowanie to przepływ kodu autoryzacji
[OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html):
przeglądarka idzie do twojego dostawcy (Keycloak, Authentik, Entra ID, Google,
Okta, Dex...), wraca do `/oauth2/callback` z kodem, oauth2-proxy
wymienia kod na ID Token, sprawdza go i ustawia własne ciasteczko.

1. **Zarejestruj klienta** u swojego dostawcy, poufnego, z przepływem
   kodu autoryzacji. Jego adresy przekierowania, które muszą dokładnie pasować:
   `https://slides.example.org/oauth2/callback` i
   `https://decks.example.org/oauth2/callback`.
2. **Umieść sekrety w plikach** czytelnych tylko dla oauth2-proxy: sekret
   klienta, bez znaku nowej linii na końcu, oraz sekret ciasteczka z
   dokładnie 32 losowych bajtów (`head -c 32 /dev/urandom`). Zapisuje je twój
   magazyn sekretów (poświadczenia systemd, sekrety Dockera lub Kubernetesa,
   Vault...); żaden nie trafia do pliku konfiguracyjnego, wiersza poleceń ani
   zmiennej środowiskowej, gdzie pokazują go logi i `ps`.
3. **Skonfiguruj oauth2-proxy** (`oauth2-proxy.cfg`):

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

`redirect_url` bez hosta sprowadza logowanie każdej witryny do jej własnego
callbacku, więc każda witryna ma własne ciasteczko. `whitelist_domains` to miejsca, do
których logowanie może wrócić: poproszone o przejście gdzie indziej, zostaje na twojej witrynie.
`trusted_proxy_ips` wierzy nagłówkom `X-Forwarded-*` tylko od Caddy.
`email_domains` mówi, kto może wejść; oauth2-proxy ma też opcje dla
grup i list adresów.

4. **Skonfiguruj Caddy** (`Caddyfile.oidc`):

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

oauth2-proxy i `slides` nasłuchują tylko na pętli zwrotnej: jedyne wejście
prowadzi przez Caddy, a więc przez logowanie. Nagłówki Playground idą
do bloku signed-in, tylko dla samego Playground: strona logowania oauth2-proxy
to formularz, który zatrzymałoby ich `form-action 'none'`.

Alternatywą jest [caddy-security](https://github.com/greenpau/caddy-security),
moduł Caddy, który realizuje logowanie wewnątrz Caddy, bez drugiego procesu.
Wymaga Caddy zbudowanego z tym modułem (`xcaddy`) i własnego
języka konfiguracji.

## Sprawdź to {#check-it}

```sh
curl -sI -H 'Accept-Encoding: br' 'https://slides.example.org/playground/playground.wasm?v=v0.1.0'
```

Odpowiedź mówi `content-type: application/wasm`, `content-encoding: br`
i `cache-control: public, max-age=31536000, immutable`, z powyższym
`content-security-policy`. Przy logowaniu z przodu to samo
żądanie bez ciasteczka dostaje odpowiedź `302` do `/oauth2/sign_in`.

## Źródła {#references}

- Caddy: [automatic HTTPS](https://caddyserver.com/docs/automatic-https),
  [`file_server`](https://caddyserver.com/docs/caddyfile/directives/file_server),
  [`header`](https://caddyserver.com/docs/caddyfile/directives/header),
  [`reverse_proxy`](https://caddyserver.com/docs/caddyfile/directives/reverse_proxy),
  [`forward_auth`](https://caddyserver.com/docs/caddyfile/directives/forward_auth).
- [RFC 8555](https://www.rfc-editor.org/rfc/rfc8555), ACME: HTTP-01 (8.3),
  DNS-01 (8.4).
- [OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html):
  przepływ kodu autoryzacji (3.1), adres przekierowania (3.1.2.1), ID
  Token (2).
- oauth2-proxy: [configuration](https://oauth2-proxy.github.io/oauth2-proxy/configuration/overview),
  [OpenID Connect provider](https://oauth2-proxy.github.io/oauth2-proxy/configuration/providers/openid_connect).
