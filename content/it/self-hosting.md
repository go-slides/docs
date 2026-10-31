---
title: "Self-hosting"
weight: 58
description: "Servire il Playground e slides --server dalla propria macchina con Caddy: HTTPS automatico, gli header necessari al Playground e un accesso OpenID Connect davanti."
tags: [playground, server, security]
---

Il [Playground]({{< relref "playground" >}}) è fatto di file statici, e
[`slides --server`]({{< relref "cli" >}}) è un solo processo sul loopback. Entrambi
possono girare sulla tua macchina, con il tuo nome, per persone che accedono
con l'account della tua organizzazione. Questa pagina lo fa con
[Caddy](https://caddyserver.com/docs/) (provato con la v2.11.7), che ottiene
e rinnova da sé i certificati HTTPS, e con
[oauth2-proxy](https://oauth2-proxy.github.io/oauth2-proxy/) per
l'accesso. Ogni file mostrato qui si trova in
[`selfhost/`](https://github.com/go-slides/go-slides.github.io/tree/main/selfhost),
e la sua CI controlla ogni Caddyfile con `caddy validate`.

## Il Playground {#the-playground}

Ogni release del sito allega l'intero sito, Playground compreso, come
un unico tarball costruito per la radice di un dominio, con il relativo SHA-256:

```sh
curl -fLO https://github.com/go-slides/go-slides.github.io/releases/download/v0.1.0/go-slides-site-v0.1.0.tar.gz
curl -fLO https://github.com/go-slides/go-slides.github.io/releases/download/v0.1.0/go-slides-site-v0.1.0.tar.gz.sha256
sha256sum -c go-slides-site-v0.1.0.tar.gz.sha256
tar -xzf go-slides-site-v0.1.0.tar.gz -C /srv
mv /srv/go-slides-site-v0.1.0 /srv/go-slides-site
```

Il tarball è riproducibile: ricostruire lo stesso tag dà gli stessi
byte. Accanto a ogni file comprimibile (i moduli WebAssembly, i
font, le pagine, gli script) contiene lo stesso file compresso con
Brotli (`.br`) e gzip (`.gz`), così il server li invia così come sono e
non comprime nulla al volo.

### Caddyfile {#caddyfile}

`slides.example.org` è il tuo nome. I suoi record A e AAAA puntano alla
macchina, e le porte 80 e 443 raggiungono Caddy:

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

Il Playground si trova allora in `https://slides.example.org/playground/`.

### Cosa deve fare bene il server {#what-the-server-must-get-right}

- **HTTPS.** Caddy ottiene il certificato da Let's Encrypt con ACME
  ([RFC 8555](https://www.rfc-editor.org/rfc/rfc8555)), dimostrando di possedere
  il nome rispondendo sulla porta 80 (HTTP-01) o 443 (TLS-ALPN-01), e lo
  rinnova. Una macchina che la CA non può raggiungere può usare invece la sfida
  DNS-01 (un record TXT, che richiede un Caddy compilato con il modulo del tuo
  provider DNS), oppure un certificato proprio con `tls cert.pem key.pem`.
  Su una macchina di prova, `tls internal` usa la CA locale di Caddy.
- **`application/wasm`.** Caddy invia questo tipo per `.wasm`; un altro server
  potrebbe doverne essere informato. Il Playground compila i propri moduli dai byte che ha
  letto (il motore di esportazione mentre scarica, in una risposta a sé
  con quel tipo), quindi parte anche da un server che invia un altro tipo;
  con `nosniff`, il tipo giusto resta comunque quello da inviare.
- **Compressione.** `precompressed br gzip` invia `playground.wasm.br` a un
  browser che accetta Brotli, `.gz` a uno che accetta gzip, e il file
  stesso negli altri casi, con `Vary: Accept-Encoding`. Il modulo dell'editor è
  di circa 7 MiB in Brotli, contro 33 MiB non compresso.
- **Cache.** La pagina richiede ogni modulo, worker, font e immagine
  con la release nella query (`playground.wasm?v=v0.1.0`), quindi questi possono
  essere conservati per un anno (`immutable`). Le pagine stesse vengono rivalidate
  (`no-cache`): una nuova release si vede alla visita successiva.
- **Gli header in `headers.caddy`:**

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

La policy è quella di cui il Playground ha bisogno: i suoi script (il loader è
in linea), il suo WebAssembly (`'wasm-unsafe-eval'`) e il suo worker, tutti dal
tuo server. Immagini, fogli di stile e font di una presentazione possono provenire da qualsiasi
indirizzo `https:`, come nell'output del comando; togli `https:` da
`img-src`, `style-src`, `font-src`, `media-src` e `connect-src` per mantenere
ogni richiesta sul tuo server, e le presentazioni che puntano altrove vengono mostrate senza
quelle risorse. La presentazione stessa gira in iframe in sandbox senza la tua
origine, sotto una policy più rigida di loro proprietà che il Playground scrive
in ogni frame: non può raggiungere il Playground, la sua memoria né i tuoi
cookie, qualunque cosa dica questo header. `frame-ancestors 'self'` impedisce ad altri
siti di incorporare il Playground in un frame.

- **Né COOP né COEP.** Il Playground non usa né `SharedArrayBuffer` né
  thread, quindi non ha bisogno di essere isolato cross-origin. Inviare
  `Cross-Origin-Embedder-Policy: require-corp` bloccherebbe le immagini di una presentazione
  provenienti da altri siti.

### Senza rete {#without-the-network}

Una volta caricata la pagina, tutto avviene nel browser: l'editor,
l'anteprima, la vista del relatore, e le esportazioni PDF, PPTX e ODP, i cui
font sono nel tarball. La presentazione non è inviata né al tuo server né
altrove. Ciò che richiede la rete è ciò che la presentazione nomina: immagini o
fogli di stile a un indirizzo `https:`, e il web font del tema `gaia`, che
chiede a `fonts.bunny.net` (senza di esso, le slide ripiegano su un altro
font). I link a questa documentazione escono dal tuo server.

## `slides --server` dietro il proxy {#slides---server-behind-the-proxy}

`slides --server` resta su `127.0.0.1` e Caddy inoltra verso di esso. Il server
controlla l'header `Host` di ogni richiesta e l'`Origin` del WebSocket di ogni
ricaricamento in tempo reale, contro il DNS rebinding e le pagine di altri siti, e il
proxy mantiene in entrambi il nome pubblico. `--allowed-host` aggiunge quel nome a
ciò che i controlli accettano; i controlli restano attivi:

```sh
slides -s --allowed-host decks.example.org /srv/decks
```

```caddy
decks.example.org {
	reverse_proxy 127.0.0.1:8080
}
```

Il ricaricamento in tempo reale passa così com'è (Caddy fa da proxy ai WebSocket senza
configurazione). Non usare `--host 0.0.0.0` per questo: servirebbe la
cartella all'intera rete, aggirando il proxy e il suo accesso.

{{< callout type="info" >}}
`--allowed-host` è nella v0.14.0 e successive:
`go install github.com/go-slides/slides/cmd/slides@v0.14.0`.
{{< /callout >}}

## Accesso con OpenID Connect {#sign-in-with-openid-connect}

Caddy interroga oauth2-proxy per ogni richiesta (`forward_auth`). oauth2-proxy
risponde `202` a un browser che ha effettuato l'accesso, e Caddy allora serve la richiesta,
e `401` altrimenti, che Caddy trasforma in un reindirizzamento all'accesso. L'
accesso è l'authorization code flow di
[OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html):
il browser va dal tuo provider (Keycloak, Authentik, Entra ID, Google,
Okta, Dex...), torna a `/oauth2/callback` con un codice, oauth2-proxy
scambia il codice con un ID Token, lo verifica e imposta un proprio cookie.

1. **Registra un client** presso il tuo provider, confidenziale, con l'
   authorization code flow. I suoi redirect URI, che devono corrispondere esattamente:
   `https://slides.example.org/oauth2/callback` e
   `https://decks.example.org/oauth2/callback`.
2. **Metti i segreti in file** leggibili solo da oauth2-proxy: il segreto del
   client, senza a capo finale, e un segreto per il cookie di
   esattamente 32 byte casuali (`head -c 32 /dev/urandom`). Li scrive il tuo
   gestore di segreti (credenziali systemd, segreti Docker o Kubernetes,
   Vault...); nessuno dei due va in un file di configurazione, in una riga di
   comando o in una variabile d'ambiente, dove log e `ps` li mostrano.
3. **Configura oauth2-proxy** (`oauth2-proxy.cfg`):

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

`redirect_url` senza host riporta l'accesso di ciascun sito al proprio
callback, così ogni sito mantiene il proprio cookie. `whitelist_domains` è dove un
accesso può tornare: se gli si chiede di andare altrove, resta sul tuo sito.
`trusted_proxy_ips` crede agli header `X-Forwarded-*` solo da Caddy.
`email_domains` dice chi può entrare; oauth2-proxy ha anche opzioni per
gruppi ed elenchi di indirizzi.

4. **Configura Caddy** (`Caddyfile.oidc`):

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

oauth2-proxy e `slides` ascoltano solo sul loopback: l'unico ingresso passa
da Caddy, e quindi dall'accesso. Gli header del Playground vanno
dentro il blocco con accesso, solo sul Playground: la pagina di accesso di oauth2-proxy
è un modulo, che il loro `form-action 'none'` bloccherebbe.

L'alternativa è [caddy-security](https://github.com/greenpau/caddy-security),
un modulo di Caddy che gestisce l'accesso dentro Caddy, senza un secondo processo.
Richiede un Caddy compilato con il modulo (`xcaddy`) e un proprio
linguaggio di configurazione.

## Verificarlo {#check-it}

```sh
curl -sI -H 'Accept-Encoding: br' 'https://slides.example.org/playground/playground.wasm?v=v0.1.0'
```

La risposta riporta `content-type: application/wasm`, `content-encoding: br`
e `cache-control: public, max-age=31536000, immutable`, con la
`content-security-policy` vista sopra. Con l'accesso davanti, la stessa
richiesta senza cookie risponde `302` verso `/oauth2/sign_in`.

## Riferimenti {#references}

- Caddy: [HTTPS automatico](https://caddyserver.com/docs/automatic-https),
  [`file_server`](https://caddyserver.com/docs/caddyfile/directives/file_server),
  [`header`](https://caddyserver.com/docs/caddyfile/directives/header),
  [`reverse_proxy`](https://caddyserver.com/docs/caddyfile/directives/reverse_proxy),
  [`forward_auth`](https://caddyserver.com/docs/caddyfile/directives/forward_auth).
- [RFC 8555](https://www.rfc-editor.org/rfc/rfc8555), ACME: HTTP-01 (8.3),
  DNS-01 (8.4).
- [OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html):
  l'authorization code flow (3.1), il redirect URI (3.1.2.1), l'ID
  Token (2).
- oauth2-proxy: [configurazione](https://oauth2-proxy.github.io/oauth2-proxy/configuration/overview),
  [provider OpenID Connect](https://oauth2-proxy.github.io/oauth2-proxy/configuration/providers/openid_connect).
