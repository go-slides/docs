---
title: "Auto-hospedagem"
weight: 58
description: "Sirva o Playground e o slides --server a partir da sua própria máquina com o Caddy: HTTPS automático, os cabeçalhos de que o Playground precisa e um login OpenID Connect na frente."
tags: [playground, server, security]
---

O [Playground]({{< relref "playground" >}}) são arquivos estáticos, e o
[`slides --server`]({{< relref "cli" >}}) é um único processo no loopback. Ambos
podem rodar na sua própria máquina, sob o seu próprio nome, para pessoas que
entram com a conta da sua organização. Esta página faz isso com o
[Caddy](https://caddyserver.com/docs/) (testado com a v2.11.7), que obtém e
renova sozinho os certificados HTTPS, e com o
[oauth2-proxy](https://oauth2-proxy.github.io/oauth2-proxy/) para o login. Todos
os arquivos mostrados aqui estão em
[`selfhost/`](https://github.com/go-slides/go-slides.github.io/tree/main/selfhost),
e a CI dele verifica cada Caddyfile com `caddy validate`.

## O Playground {#the-playground}

Cada release do site anexa o site inteiro, Playground incluído, como um único
tarball construído para a raiz de um domínio, com o seu SHA-256:

```sh
curl -fLO https://github.com/go-slides/go-slides.github.io/releases/download/v0.1.0/go-slides-site-v0.1.0.tar.gz
curl -fLO https://github.com/go-slides/go-slides.github.io/releases/download/v0.1.0/go-slides-site-v0.1.0.tar.gz.sha256
sha256sum -c go-slides-site-v0.1.0.tar.gz.sha256
tar -xzf go-slides-site-v0.1.0.tar.gz -C /srv
mv /srv/go-slides-site-v0.1.0 /srv/go-slides-site
```

O tarball é reproduzível: construir a mesma tag de novo dá os mesmos bytes. Ao
lado de cada arquivo que comprime (os módulos WebAssembly, as fontes, as
páginas, os scripts) ele traz o mesmo arquivo comprimido com Brotli (`.br`) e
gzip (`.gz`), de modo que o servidor os envia como estão e não comprime nada na
hora.

### Caddyfile {#caddyfile}

`slides.example.org` é o seu nome. Os registros A e AAAA dele apontam para a
máquina, e as portas 80 e 443 chegam ao Caddy:

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

O Playground fica então em `https://slides.example.org/playground/`.

### O que o servidor precisa acertar {#what-the-server-must-get-right}

- **HTTPS.** O Caddy obtém o certificado da Let's Encrypt com ACME
  ([RFC 8555](https://www.rfc-editor.org/rfc/rfc8555)), provando que possui o
  nome ao responder na porta 80 (HTTP-01) ou 443 (TLS-ALPN-01), e o renova. Uma
  máquina que a CA não consegue alcançar pode usar o desafio DNS-01 (um registro
  TXT, que exige um Caddy construído com o módulo do seu provedor de DNS), ou o
  seu próprio certificado com `tls cert.pem key.pem`. Em uma máquina de teste,
  `tls internal` usa a CA local do Caddy.
- **`application/wasm`.** O Caddy envia este tipo para `.wasm`; outro servidor
  pode precisar ser instruído. O Playground compila os seus módulos a partir dos
  bytes que leu (o motor de exportação à medida que baixa, em uma resposta
  própria com esse tipo), então ele inicia mesmo a partir de um servidor que
  envia outro tipo; com `nosniff`, o tipo certo continua sendo o que se deve
  enviar.
- **Compressão.** `precompressed br gzip` envia `playground.wasm.br` a um
  navegador que aceita Brotli, `.gz` a um que aceita gzip, e o próprio arquivo
  nos demais casos, com `Vary: Accept-Encoding`. O módulo do editor tem cerca de
  7 MiB em Brotli, contra 33 MiB sem compressão.
- **Cache.** A página pede cada módulo, worker, fonte e imagem com a release na
  query (`playground.wasm?v=v0.1.0`), de modo que eles podem ser mantidos por um
  ano (`immutable`). As próprias páginas são revalidadas (`no-cache`): uma nova
  release é vista na próxima visita.
- **Os cabeçalhos em `headers.caddy`:**

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

A política é a de que o Playground precisa: os seus scripts (o loader é em
linha), o seu WebAssembly (`'wasm-unsafe-eval'`) e o seu worker, todos vindos do
seu servidor. As imagens, folhas de estilo e fontes de uma apresentação podem vir
de qualquer endereço `https:`, como na saída do comando; retire `https:` de
`img-src`, `style-src`, `font-src`, `media-src` e `connect-src` para manter todas
as requisições no seu servidor, e as apresentações que apontam para outro lugar
são exibidas sem esses recursos. A própria apresentação roda em iframes em
sandbox sem a sua origem, sob uma política própria mais estrita que o Playground
escreve em cada frame: ela não alcança o Playground, o armazenamento dele nem os
seus cookies, diga o que disser este cabeçalho. `frame-ancestors 'self'` impede
que outros sites incorporem o Playground em frames.

- **Nem COOP nem COEP.** O Playground não usa `SharedArrayBuffer` nem threads,
  então não precisa de isolamento cross-origin. Enviar
  `Cross-Origin-Embedder-Policy: require-corp` bloquearia as imagens de uma
  apresentação vindas de outros sites.

### Sem a rede {#without-the-network}

Depois que a página é carregada, tudo acontece no navegador: o editor, a
pré-visualização, a visão do apresentador, e as exportações para PDF, PPTX e
ODP, cujas fontes estão no tarball. A apresentação não é enviada nem ao seu
servidor nem a qualquer outro lugar. O que precisa da rede é o que a
apresentação nomeia: imagens ou folhas de estilo em um endereço `https:`, e a
fonte web do tema `gaia`, que ela pede ao `fonts.bunny.net` (sem ela, os slides
recorrem a outra fonte). Os links para esta documentação saem do seu servidor.

## `slides --server` atrás do proxy {#slides---server-behind-the-proxy}

O `slides --server` fica em `127.0.0.1` e o Caddy encaminha para ele. O servidor
verifica o cabeçalho `Host` de cada requisição e o `Origin` do WebSocket de cada
recarga ao vivo, contra DNS rebinding e páginas de outros sites, e o proxy
mantém o nome público nos dois. `--allowed-host` acrescenta esse nome ao que as
verificações aceitam; elas continuam ativas:

```sh
slides -s --allowed-host decks.example.org /srv/decks
```

```caddy
decks.example.org {
	reverse_proxy 127.0.0.1:8080
}
```

A recarga ao vivo passa como está (o Caddy faz proxy de WebSockets sem
configuração). Não use `--host 0.0.0.0` para isso: ele serve o diretório para a
rede inteira, contornando o proxy e o login dele.

{{< callout type="info" >}}
`--allowed-host` está na v0.14.0 e posteriores:
`go install github.com/go-slides/slides/cmd/slides@v0.15.0`. Requer Go 1.27.3
ou posterior.
{{< /callout >}}

## Login com OpenID Connect {#sign-in-with-openid-connect}

O Caddy consulta o oauth2-proxy sobre cada requisição (`forward_auth`). O
oauth2-proxy responde `202` para um navegador que fez login, o que o Caddy então
serve, e `401` caso contrário, o que o Caddy transforma em um redirecionamento
para o login. O login é o fluxo de código de autorização do
[OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html):
o navegador vai ao seu provedor (Keycloak, Authentik, Entra ID, Google, Okta,
Dex...), volta a `/oauth2/callback` com um código, o oauth2-proxy troca o código
por um ID Token, o verifica e define o seu próprio cookie.

1. **Registre um cliente** no seu provedor, confidencial, com o fluxo de código
   de autorização. As URIs de redirecionamento dele, que devem coincidir
   exatamente: `https://slides.example.org/oauth2/callback` e
   `https://decks.example.org/oauth2/callback`.
2. **Coloque os segredos em arquivos** legíveis somente pelo oauth2-proxy: o
   segredo do cliente, sem quebra de linha no final, e um segredo de cookie de
   exatamente 32 bytes aleatórios (`head -c 32 /dev/urandom`). O seu repositório
   de segredos os escreve (credenciais do systemd, segredos do Docker ou do
   Kubernetes, Vault...); nenhum dos dois vai para um arquivo de configuração,
   uma linha de comando ou uma variável de ambiente, onde os logs e o `ps` os
   mostram.
3. **Configure o oauth2-proxy** (`oauth2-proxy.cfg`):

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

`redirect_url` sem host leva o login de cada site de volta ao seu próprio
callback, de modo que cada site mantém o seu próprio cookie. `whitelist_domains`
é para onde um login pode retornar: se pedirem para ir a qualquer outro lugar,
ele permanece no seu site. `trusted_proxy_ips` acredita nos cabeçalhos
`X-Forwarded-*` somente do Caddy. `email_domains` diz quem pode entrar; o
oauth2-proxy também tem opções para grupos e listas de endereços.

4. **Configure o Caddy** (`Caddyfile.oidc`):

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

O oauth2-proxy e o `slides` escutam somente no loopback: o único caminho de
entrada passa pelo Caddy, e portanto pelo login. Os cabeçalhos do Playground
ficam dentro do bloco de login, somente no Playground: a página de login do
oauth2-proxy é um formulário, que o `form-action 'none'` deles bloquearia.

A alternativa é o [caddy-security](https://github.com/greenpau/caddy-security),
um módulo do Caddy que faz o login dentro do próprio Caddy, sem um segundo
processo. Ele exige um Caddy construído com o módulo (`xcaddy`) e tem a sua
própria linguagem de configuração.

## Verificar {#check-it}

```sh
curl -sI -H 'Accept-Encoding: br' 'https://slides.example.org/playground/playground.wasm?v=v0.1.0'
```

A resposta diz `content-type: application/wasm`, `content-encoding: br` e
`cache-control: public, max-age=31536000, immutable`, com o
`content-security-policy` acima. Com o login na frente, a mesma requisição sem
cookie responde `302` para `/oauth2/sign_in`.

## Referências {#references}

- Caddy: [HTTPS automático](https://caddyserver.com/docs/automatic-https),
  [`file_server`](https://caddyserver.com/docs/caddyfile/directives/file_server),
  [`header`](https://caddyserver.com/docs/caddyfile/directives/header),
  [`reverse_proxy`](https://caddyserver.com/docs/caddyfile/directives/reverse_proxy),
  [`forward_auth`](https://caddyserver.com/docs/caddyfile/directives/forward_auth).
- [RFC 8555](https://www.rfc-editor.org/rfc/rfc8555), ACME: HTTP-01 (8.3),
  DNS-01 (8.4).
- [OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html):
  o fluxo de código de autorização (3.1), a URI de redirecionamento (3.1.2.1), o
  ID Token (2).
- oauth2-proxy: [configuração](https://oauth2-proxy.github.io/oauth2-proxy/configuration/overview),
  [provedor OpenID Connect](https://oauth2-proxy.github.io/oauth2-proxy/configuration/providers/openid_connect).
