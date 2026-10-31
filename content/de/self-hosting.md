---
title: "Self-Hosting"
weight: 58
description: "Den Playground und slides --server vom eigenen Rechner aus mit Caddy bereitstellen: automatisches HTTPS, die Header, die der Playground braucht, und eine OpenID-Connect-Anmeldung davor."
tags: [playground, server, security]
---

Der [Playground]({{< relref "playground" >}}) besteht aus statischen Dateien, und
[`slides --server`]({{< relref "cli" >}}) ist ein einzelner Prozess auf dem Loopback. Beide
können auf Ihrem eigenen Rechner laufen, unter Ihrem eigenen Namen, für Personen, die sich
mit dem Konto Ihrer Organisation anmelden. Diese Seite richtet das ein mit
[Caddy](https://caddyserver.com/docs/) (getestet mit v2.11.7), der die HTTPS-Zertifikate
selbst bezieht und erneuert, und
[oauth2-proxy](https://oauth2-proxy.github.io/oauth2-proxy/) für die
Anmeldung. Jede hier gezeigte Datei liegt in
[`selfhost/`](https://github.com/go-slides/go-slides.github.io/tree/main/selfhost),
und deren CI prüft jedes Caddyfile mit `caddy validate`.

## Der Playground {#the-playground}

Jedes Release der Website hängt die ganze Website samt Playground als
ein Tarball an, gebaut für die Wurzel einer Domain, mit seiner SHA-256:

```sh
curl -fLO https://github.com/go-slides/go-slides.github.io/releases/download/v0.1.0/go-slides-site-v0.1.0.tar.gz
curl -fLO https://github.com/go-slides/go-slides.github.io/releases/download/v0.1.0/go-slides-site-v0.1.0.tar.gz.sha256
sha256sum -c go-slides-site-v0.1.0.tar.gz.sha256
tar -xzf go-slides-site-v0.1.0.tar.gz -C /srv
mv /srv/go-slides-site-v0.1.0 /srv/go-slides-site
```

Der Tarball ist reproduzierbar: Wer denselben Tag erneut baut, erhält dieselben
Bytes. Neben jeder Datei, die sich komprimieren lässt (die WebAssembly-Module, die
Schriften, die Seiten, die Skripte), liegt dieselbe Datei komprimiert mit
Brotli (`.br`) und gzip (`.gz`), sodass der Server sie unverändert sendet und
nichts im laufenden Betrieb komprimiert.

### Caddyfile {#caddyfile}

`slides.example.org` ist Ihr Name. Seine A- und AAAA-Einträge zeigen auf den
Rechner, und die Ports 80 und 443 erreichen Caddy:

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

Der Playground liegt dann unter `https://slides.example.org/playground/`.

### Was der Server richtig machen muss {#what-the-server-must-get-right}

- **HTTPS.** Caddy bezieht das Zertifikat von Let's Encrypt per ACME
  ([RFC 8555](https://www.rfc-editor.org/rfc/rfc8555)), beweist den Besitz
  des Namens, indem es auf Port 80 (HTTP-01) oder 443 (TLS-ALPN-01) antwortet, und
  erneuert es. Ein Rechner, den die CA nicht erreichen kann, kann stattdessen die
  DNS-01-Challenge nutzen (ein TXT-Eintrag, wofür ein Caddy mit dem Modul Ihres
  DNS-Anbieters gebaut sein muss) oder ein eigenes Zertifikat mit `tls cert.pem key.pem`.
  Auf einem Testrechner verwendet `tls internal` die lokale CA von Caddy.
- **`application/wasm`.** Caddy sendet diesen Typ für `.wasm`; ein anderer Server
  muss es möglicherweise ausdrücklich gesagt bekommen. Der Playground kompiliert seine Module aus den Bytes, die er
  gelesen hat (die Export-Engine schon beim Herunterladen, in einer eigenen Antwort
  mit diesem Typ), er startet also auch von einem Server, der einen anderen Typ sendet;
  mit `nosniff` ist der richtige Typ trotzdem der, den man senden sollte.
- **Kompression.** `precompressed br gzip` sendet `playground.wasm.br` an einen
  Browser, der Brotli akzeptiert, `.gz` an einen, der gzip akzeptiert, und sonst die Datei
  selbst, mit `Vary: Accept-Encoding`. Das Modul des Editors hat
  mit Brotli etwa 7 MiB, unkomprimiert 33 MiB.
- **Caching.** Die Seite fordert jedes Modul, jeden Worker, jede Schrift und jedes Bild
  mit dem Release in der Query an (`playground.wasm?v=v0.1.0`), daher können diese
  ein Jahr lang aufbewahrt werden (`immutable`). Die Seiten selbst werden neu validiert
  (`no-cache`): Ein neues Release ist beim nächsten Besuch zu sehen.
- **Die Header in `headers.caddy`:**

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

Die Richtlinie ist die, die der Playground braucht: seine Skripte (der Loader ist
inline), sein WebAssembly (`'wasm-unsafe-eval'`) und sein Worker, alle von
Ihrem Server. Bilder, Stylesheets und Schriften einer Präsentation dürfen von jeder
`https:`-Adresse kommen, wie in der Ausgabe des Befehls; entfernen Sie `https:` aus
`img-src`, `style-src`, `font-src`, `media-src` und `connect-src`, damit
alle Anfragen auf Ihrem Server bleiben; Präsentationen, die woandershin verweisen, werden dann ohne
diese Ressourcen angezeigt. Die Präsentation selbst läuft in iframes in einer Sandbox ohne Ihren
Ursprung, unter einer strengeren eigenen Richtlinie, die der Playground
in jeden Frame schreibt: Sie kann weder den Playground noch seinen Speicher noch Ihre
Cookies erreichen, was auch immer dieser Header sagt. `frame-ancestors 'self'` verhindert, dass andere
Websites den Playground in einen Frame einbetten.

- **Kein COOP und kein COEP.** Der Playground verwendet weder `SharedArrayBuffer` noch
  Threads, er muss also nicht cross-origin-isoliert sein. Wer
  `Cross-Origin-Embedder-Policy: require-corp` sendet, blockiert die Bilder einer Präsentation
  von anderen Websites.

### Ohne Netzwerk {#without-the-network}

Sobald die Seite geladen ist, passiert alles im Browser: der Editor,
die Vorschau, die Referentenansicht und die PDF-, PPTX- und ODP-Exporte, deren
Schriften im Tarball liegen. Die Präsentation wird weder an Ihren Server noch
sonst irgendwohin gesendet. Das Netzwerk brauchen die Dinge, die die Präsentation nennt: Bilder oder
Stylesheets unter einer `https:`-Adresse und die Webschrift des Themes `gaia`, die
es bei `fonts.bunny.net` anfordert (ohne sie weichen die Folien auf eine andere
Schrift aus). Die Links zu dieser Dokumentation verlassen Ihren Server.

## `slides --server` hinter dem Proxy {#slides---server-behind-the-proxy}

`slides --server` bleibt auf `127.0.0.1`, und Caddy leitet an ihn weiter. Der Server
prüft den `Host`-Header jeder Anfrage und den `Origin` des WebSockets jedes
Live-Reloads, gegen DNS-Rebinding und Seiten anderer Websites, und der
Proxy behält in beiden den öffentlichen Namen. `--allowed-host` fügt diesen Namen dem hinzu,
was die Prüfungen akzeptieren; sie bleiben eingeschaltet:

```sh
slides -s --allowed-host decks.example.org /srv/decks
```

```caddy
decks.example.org {
	reverse_proxy 127.0.0.1:8080
}
```

Der Live-Reload geht unverändert durch (Caddy leitet WebSockets ohne
Konfiguration weiter). Verwenden Sie dafür nicht `--host 0.0.0.0`: Das stellt das
Verzeichnis dem ganzen Netzwerk bereit, am Proxy und seiner Anmeldung vorbei.

{{< callout type="info" >}}
`--allowed-host` gibt es ab v0.14.0:
`go install github.com/go-slides/slides/cmd/slides@v0.15.0`. Dafür ist Go 1.27.3
oder neuer nötig.
{{< /callout >}}

## Anmeldung mit OpenID Connect {#sign-in-with-openid-connect}

Caddy fragt oauth2-proxy bei jeder Anfrage (`forward_auth`). oauth2-proxy
antwortet `202` für einen Browser, der sich angemeldet hat, was Caddy dann ausliefert,
und sonst `401`, was Caddy in eine Weiterleitung zur Anmeldung verwandelt. Die
Anmeldung ist der Authorization-Code-Flow von
[OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html):
Der Browser geht zu Ihrem Anbieter (Keycloak, Authentik, Entra ID, Google,
Okta, Dex...), kommt mit einem Code auf `/oauth2/callback` zurück, oauth2-proxy
tauscht den Code gegen ein ID Token, prüft es und setzt sein eigenes Cookie.

1. **Registrieren Sie einen Client** bei Ihrem Anbieter, vertraulich, mit dem
   Authorization-Code-Flow. Seine Redirect-URIs, die genau übereinstimmen müssen:
   `https://slides.example.org/oauth2/callback` und
   `https://decks.example.org/oauth2/callback`.
2. **Legen Sie die Geheimnisse in Dateien**, die nur oauth2-proxy lesen kann: das
   Geheimnis des Clients, ohne Zeilenumbruch am Ende, und ein Cookie-Geheimnis aus
   genau 32 Zufallsbytes (`head -c 32 /dev/urandom`). Ihr
   Secret-Speicher schreibt sie (systemd-Credentials, Docker- oder Kubernetes-Secrets,
   Vault...); keines gehört in eine Konfigurationsdatei, eine Befehlszeile
   oder eine Umgebungsvariable, wo Logs und `ps` sie zeigen.
3. **Konfigurieren Sie oauth2-proxy** (`oauth2-proxy.cfg`):

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

`redirect_url` ohne Host führt die Anmeldung jeder Website zu ihrem eigenen
Callback zurück, sodass jede Website ihr eigenes Cookie behält. `whitelist_domains` ist, wohin eine
Anmeldung zurückkehren darf: Soll sie anderswohin gehen, bleibt sie auf Ihrer Website.
`trusted_proxy_ips` glaubt den `X-Forwarded-*`-Headern nur von Caddy.
`email_domains` legt fest, wer hinein darf; oauth2-proxy hat außerdem Optionen für
Gruppen und Adresslisten.

4. **Konfigurieren Sie Caddy** (`Caddyfile.oidc`):

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

oauth2-proxy und `slides` lauschen nur auf dem Loopback: Der einzige Zugang führt
über Caddy und damit über die Anmeldung. Die Header des Playgrounds stehen
innerhalb des angemeldeten Blocks, nur beim Playground: Die Anmeldeseite von
oauth2-proxy ist ein Formular, das ihr `form-action 'none'` verhindern würde.

Die Alternative ist [caddy-security](https://github.com/greenpau/caddy-security),
ein Caddy-Modul, das die Anmeldung innerhalb von Caddy erledigt, ohne zweiten Prozess.
Es braucht einen Caddy, der mit dem Modul gebaut wurde (`xcaddy`), und eine eigene
Konfigurationssprache.

## Prüfen {#check-it}

```sh
curl -sI -H 'Accept-Encoding: br' 'https://slides.example.org/playground/playground.wasm?v=v0.1.0'
```

Die Antwort nennt `content-type: application/wasm`, `content-encoding: br`
und `cache-control: public, max-age=31536000, immutable`, mit dem
`content-security-policy` von oben. Mit der Anmeldung davor antwortet dieselbe
Anfrage ohne Cookie mit `302` auf `/oauth2/sign_in`.

## Referenzen {#references}

- Caddy: [automatisches HTTPS](https://caddyserver.com/docs/automatic-https),
  [`file_server`](https://caddyserver.com/docs/caddyfile/directives/file_server),
  [`header`](https://caddyserver.com/docs/caddyfile/directives/header),
  [`reverse_proxy`](https://caddyserver.com/docs/caddyfile/directives/reverse_proxy),
  [`forward_auth`](https://caddyserver.com/docs/caddyfile/directives/forward_auth).
- [RFC 8555](https://www.rfc-editor.org/rfc/rfc8555), ACME: HTTP-01 (8.3),
  DNS-01 (8.4).
- [OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html):
  der Authorization-Code-Flow (3.1), die Redirect-URI (3.1.2.1), das ID
  Token (2).
- oauth2-proxy: [Konfiguration](https://oauth2-proxy.github.io/oauth2-proxy/configuration/overview),
  [OpenID-Connect-Anbieter](https://oauth2-proxy.github.io/oauth2-proxy/configuration/providers/openid_connect).
