---
title: "Autoalojamiento"
weight: 58
description: "Sirva el Playground y slides --server desde su propio equipo con Caddy: HTTPS automático, las cabeceras que necesita el playground y un inicio de sesión OpenID Connect por delante."
tags: [playground, server, security]
---

El [Playground]({{< relref "playground" >}}) son archivos estáticos, y
[`slides --server`]({{< relref "cli" >}}) es un solo proceso en el bucle local. Ambos
pueden ejecutarse en su propio equipo, con su propio nombre, para personas que inician
sesión con la cuenta de su organización. Esta página lo hace con
[Caddy](https://caddyserver.com/docs/) (probado con v2.11.7), que obtiene
y renueva por sí mismo los certificados HTTPS, y
[oauth2-proxy](https://oauth2-proxy.github.io/oauth2-proxy/) para el
inicio de sesión. Todos los archivos que se muestran aquí están en
[`selfhost/`](https://github.com/go-slides/go-slides.github.io/tree/main/selfhost),
y su CI comprueba cada Caddyfile con `caddy validate`.

## El Playground {#the-playground}

Cada versión del sitio adjunta el sitio completo, Playground incluido, como
un solo tarball construido para la raíz de un dominio, con su SHA-256:

```sh
curl -fLO https://github.com/go-slides/go-slides.github.io/releases/download/v0.1.0/go-slides-site-v0.1.0.tar.gz
curl -fLO https://github.com/go-slides/go-slides.github.io/releases/download/v0.1.0/go-slides-site-v0.1.0.tar.gz.sha256
sha256sum -c go-slides-site-v0.1.0.tar.gz.sha256
tar -xzf go-slides-site-v0.1.0.tar.gz -C /srv
mv /srv/go-slides-site-v0.1.0 /srv/go-slides-site
```

El tarball es reproducible: construir de nuevo la misma etiqueta da los mismos
bytes. Junto a cada archivo que se comprime (los módulos WebAssembly, las
fuentes, las páginas, los scripts), lleva el mismo archivo comprimido con
Brotli (`.br`) y gzip (`.gz`), de modo que el servidor los envía tal cual y
no comprime nada al vuelo.

### Caddyfile {#caddyfile}

`slides.example.org` es su nombre. Sus registros A y AAAA apuntan al
equipo, y los puertos 80 y 443 llegan a Caddy:

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

El Playground queda entonces en `https://slides.example.org/playground/`.

### Lo que el servidor debe hacer bien {#what-the-server-must-get-right}

- **HTTPS.** Caddy obtiene el certificado de Let's Encrypt con ACME
  ([RFC 8555](https://www.rfc-editor.org/rfc/rfc8555)), demostrando que posee
  el nombre al responder en el puerto 80 (HTTP-01) o 443 (TLS-ALPN-01), y
  lo renueva. Un equipo al que la CA no puede llegar puede usar en su lugar el desafío
  DNS-01 (un registro TXT, que requiere un Caddy compilado con el módulo de su
  proveedor de DNS), o su propio certificado con `tls cert.pem key.pem`.
  En un equipo de pruebas, `tls internal` usa la CA local de Caddy.
- **`application/wasm`.** Caddy envía este tipo para `.wasm`; en otro servidor
  puede que haya que indicarlo. El Playground compila sus módulos a partir de los bytes que
  ha leído (el motor de exportación mientras se descarga, en una respuesta propia
  con ese tipo), así que arranca incluso desde un servidor que envía otro tipo;
  con `nosniff`, el tipo correcto sigue siendo el que hay que enviar.
- **Compresión.** `precompressed br gzip` envía `playground.wasm.br` a un
  navegador que acepta Brotli, `.gz` a uno que acepta gzip, y el archivo
  mismo en caso contrario, con `Vary: Accept-Encoding`. El módulo del editor pesa
  unos 7 MiB en Brotli, frente a 33 MiB sin comprimir.
- **Caché.** La página pide cada módulo, worker, fuente e imagen
  con la versión en la consulta (`playground.wasm?v=v0.1.0`), de modo que se pueden
  conservar durante un año (`immutable`). Las páginas en sí se revalidan
  (`no-cache`): una versión nueva se ve en la siguiente visita.
- **Las cabeceras de `headers.caddy`:**

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

La política es la que necesita el Playground: sus scripts (el cargador está
en línea), su WebAssembly (`'wasm-unsafe-eval'`) y su worker, todos desde
su servidor. Las imágenes, hojas de estilo y fuentes de una presentación pueden venir de cualquier
dirección `https:`, como en la salida del comando; quite `https:` de
`img-src`, `style-src`, `font-src`, `media-src` y `connect-src` para mantener
todas las peticiones en su servidor, y las presentaciones que apunten a otra parte se mostrarán sin
esos recursos. La presentación en sí se ejecuta en iframes aislados (sandbox) sin su
origen, bajo una política propia más estricta que el Playground escribe
en cada marco: no puede llegar al Playground, a su almacenamiento ni a sus
cookies, diga lo que diga esta cabecera. `frame-ancestors 'self'` impide que otros
sitios incrusten el Playground en un marco.

- **Ni COOP ni COEP.** El Playground no usa `SharedArrayBuffer` ni
  hilos, así que no necesita estar aislado entre orígenes. Enviar
  `Cross-Origin-Embedder-Policy: require-corp` bloquearía las imágenes de una presentación
  procedentes de otros sitios.

### Sin red {#without-the-network}

Una vez cargada la página, todo ocurre en el navegador: el editor,
la vista previa, la vista del presentador y las exportaciones a PDF, PPTX y ODP, cuyas
fuentes están en el tarball. La presentación no se envía ni a su servidor ni
a ningún otro sitio. Lo que necesita la red es lo que la presentación nombra: imágenes u
hojas de estilo en una dirección `https:`, y la fuente web del tema `gaia`, que
pide a `fonts.bunny.net` (sin ella, las diapositivas recurren a otra
fuente). Los enlaces a esta documentación salen de su servidor.

## `slides --server` detrás del proxy {#slides---server-behind-the-proxy}

`slides --server` se queda en `127.0.0.1` y Caddy le reenvía las peticiones. El servidor
comprueba la cabecera `Host` de cada petición y el `Origin` del WebSocket de cada
recarga en vivo, contra el DNS rebinding y las páginas de otros sitios, y el
proxy conserva el nombre público en ambos. `--allowed-host` añade ese nombre a
lo que aceptan las comprobaciones; estas siguen activas:

```sh
slides -s --allowed-host decks.example.org /srv/decks
```

```caddy
decks.example.org {
	reverse_proxy 127.0.0.1:8080
}
```

La recarga en vivo pasa tal cual (Caddy hace de proxy de los WebSockets sin
configuración). No use `--host 0.0.0.0` para esto: serviría el
directorio a toda la red, saltándose el proxy y su inicio de sesión.

{{< callout type="info" >}}
`--allowed-host` está en v0.14.0 y posteriores:
`go install github.com/go-slides/slides/cmd/slides@v0.20.0`. Requiere Go 1.27.3
o posterior.
{{< /callout >}}

## Inicio de sesión con OpenID Connect {#sign-in-with-openid-connect}

Caddy consulta a oauth2-proxy en cada petición (`forward_auth`). oauth2-proxy
responde `202` a un navegador que ha iniciado sesión, y Caddy entonces lo sirve,
y `401` en caso contrario, que Caddy convierte en una redirección al inicio de sesión. El
inicio de sesión es el flujo de código de autorización de
[OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html):
el navegador va a su proveedor (Keycloak, Authentik, Entra ID, Google,
Okta, Dex...), vuelve a `/oauth2/callback` con un código, oauth2-proxy
intercambia el código por un ID Token, lo comprueba y establece su propia cookie.

1. **Registre un cliente** en su proveedor, confidencial, con el flujo de
   código de autorización. Sus URI de redirección, que deben coincidir exactamente:
   `https://slides.example.org/oauth2/callback` y
   `https://decks.example.org/oauth2/callback`.
2. **Ponga los secretos en archivos** que solo oauth2-proxy pueda leer: el
   secreto del cliente, sin salto de línea al final, y un secreto de cookie de
   exactamente 32 bytes aleatorios (`head -c 32 /dev/urandom`). Su
   almacén de secretos los escribe (credenciales de systemd, secretos de Docker o Kubernetes,
   Vault...); ninguno va en un archivo de configuración, una línea
   de comandos ni una variable de entorno, donde los registros y `ps` los muestran.
3. **Configure oauth2-proxy** (`oauth2-proxy.cfg`):

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

`redirect_url` sin host devuelve el inicio de sesión de cada sitio a su propio
callback, de modo que cada sitio conserva su propia cookie. `whitelist_domains` es adonde puede volver un
inicio de sesión: si se le pide ir a cualquier otro sitio, se queda en el suyo.
`trusted_proxy_ips` solo cree las cabeceras `X-Forwarded-*` que vienen de Caddy.
`email_domains` indica quién puede entrar; oauth2-proxy también tiene opciones para
grupos y listas de direcciones.

4. **Configure Caddy** (`Caddyfile.oidc`):

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

oauth2-proxy y `slides` escuchan solo en el bucle local: la única entrada pasa
por Caddy y, por tanto, por el inicio de sesión. Las cabeceras del Playground van
dentro del bloque con sesión iniciada, solo en el Playground: la página de inicio
de sesión de oauth2-proxy es un formulario, que su `form-action 'none'` bloquearía.

La alternativa es [caddy-security](https://github.com/greenpau/caddy-security),
un módulo de Caddy que hace el inicio de sesión dentro de Caddy, sin un segundo proceso.
Necesita un Caddy compilado con el módulo (`xcaddy`) y su propio lenguaje
de configuración.

## slides-gitd: las credenciales se quedan en el servidor {#slides-gitd}

[slides-gitd](https://github.com/go-slides/gitd) (v0.1.0) es el backend de git
de un Playground autoalojado. Carga la carpeta de una presentación desde una
forja, propone los cambios en una rama nueva con una pull (merge) request y
busca cambios en origen, mientras las credenciales de la forja **se quedan en
el servidor**: el navegador solo tiene una cookie de sesión HttpOnly. Va detrás
de Caddy y oauth2-proxy (OIDC), escucha en `127.0.0.1:8091`, toma la identidad
del usuario solo del proxy y solo escribe en ramas llamadas
`slides/<usuario>/<título>-<hora>` que el propio servidor nombra.

El modo servidor del Playground, que habla con slides-gitd en lugar de que el
navegador guarde un token, se describe en
[El Playground detrás de slides-gitd](#playground-server-mode).

Configuración (los secretos son rutas de archivo, nunca valores; cada archivo
solo debe ser legible por su propietario):

```json
{
  "listen": "127.0.0.1:8091",
  "origin": "https://slides.example.org",
  "proxy_secret_file": "/etc/slides-gitd/proxy.secret",
  "master_key_file": "/etc/slides-gitd/master.key",
  "data_dir": "/var/lib/slides-gitd",
  "audit_file": "/var/log/slides-gitd/audit.log",
  "forges": {
    "github": {"kind": "github", "api": "https://api.github.com", "auth": "github-app",
               "app_client_id": "Iv1.…", "installation_id": 12345678,
               "app_key_file": "/etc/slides-gitd/github-app.pem"},
    "gitlab": {"kind": "gitlab", "api": "https://gitlab.example.org/api/v4", "auth": "user-token"}
  },
  "users": {
    "3f2b…-oidc-sub": {"slug": "ada", "repos": ["github:lab/talks", "gitlab:team/slides"]}
  }
}
```

Un usuario solo llega a las entradas `forja:propietario/nombre` listadas para su
sujeto OIDC (denegado por defecto). Cree la clave maestra con
`head -c 32 /dev/urandom > master.key && chmod 600 master.key`, luego ejecute
`slides-gitd check -c config.json`, que carga todos los archivos de secretos y
termina, y `slides-gitd serve -c config.json`.

El lado de Caddy, con oauth2-proxy iniciado con `--set-xauthrequest`:

```caddyfile
slides.example.org {
    handle /oauth2/* {
        reverse_proxy 127.0.0.1:4180
    }
    handle /api/git/* {
        # never let a client choose its identity, whatever Caddy's version does
        request_header -X-Auth-Request-User
        request_header -X-Auth-Request-Email
        request_header -X-Gitd-Proxy-Secret
        forward_auth 127.0.0.1:4180 {
            uri /oauth2/auth
            copy_headers X-Auth-Request-User X-Auth-Request-Email
        }
        request_body { max_size 64MB }
        reverse_proxy 127.0.0.1:8091 {
            header_up X-Gitd-Proxy-Secret {file./etc/caddy/gitd-proxy.secret}
        }
    }
    handle { root * /srv/slides; file_server }
}
```

Las líneas `request_header -…` borran cualquier cabecera de identidad que envíe el
cliente, así que un cliente no puede elegir quién es; el Caddy reciente lo hace
solo antes de `copy_headers`, las versiones antiguas no. `{file.*}` lee el
secreto del proxy desde un archivo, así que no está en una variable de entorno.

**GitHub** usa una App. Créela en Settings ▸ Developer settings ▸ GitHub Apps
con los permisos Contents de lectura y escritura, Pull requests de lectura y
escritura y Metadata de lectura, sin webhook; instálela solo en los
repositorios de las presentaciones. slides-gitd emite un token de instalación
para el único repositorio de cada llamada, válido una hora.

**GitLab:** cada usuario guarda un token de acceso personal o de proyecto con el
alcance `api`. Por qué `api` y no `write_repository`, y qué puede hacer un token
robado, se explica [en la página del Playground]({{< relref "playground#gitlab" >}});
prefiera un token de acceso de proyecto o de grupo limitado a los repositorios
de las presentaciones, o una cuenta dedicada.

### El Playground detrás de slides-gitd (modo servidor) {#playground-server-mode}

Cuando el Playground se sirve detrás de slides-gitd en su propio origen, lo
detecta preguntando `GET /api/git/session` en ese origen. `401` significa que
hay un servidor y nadie ha iniciado sesión: el panel muestra **Iniciar
sesión**. Una sesión JSON significa que la sesión está iniciada. Cualquier otra
cosa, como el `404` de un alojamiento estático, significa que no hay servidor, y
el panel funciona exactamente como en go-slides.github.io: un token guardado en
la pestaña o el inicio de sesión de GitLab.

En modo servidor el panel oculta los campos del token y de OAuth. Muestra
`Signed in as <identity>` (sesión iniciada como) y la lista de repositorios que
el servidor concede. **Cargar**, **Proponer cambios…** (el servidor nombra la
rama `slides/<readable>-<hash>/…`; la página no puede elegir una),
**Buscar cambios** y **Tomar los suyos** van todos a `/api/git/*` en el mismo
origen. El navegador no contacta con ningún host de la forja.

Qué queda dónde:

- Los tokens viven en el servidor, en una bóveda AES-GCM con una clave por
  usuario.
- El navegador solo guarda la cookie `__Host-gitd` (`HttpOnly`, `Secure`,
  `SameSite=Strict`), que los scripts no pueden leer.
- El token CSRF se guarda solo en la memoria de la página y nunca se almacena.
- Toda escritura lleva el `Origin` exacto, `X-Requested-With: slides-gitd`, un
  cuerpo JSON y `X-CSRF-Token`.

Cuando el inicio de sesión ha caducado o falta (HTTP `401`), el panel muestra
«Su sesión ha terminado» con un botón **Iniciar sesión**
(`/oauth2/sign_in?rd=<this page>` de oauth2-proxy). Nunca recurre a un token
escrito en el navegador.

Para servirlo, use `selfhost/Caddyfile.gitd` de go-slides/go-slides.github.io.
Su Content-Security-Policy es `connect-src 'self'`; el sitio sigue siendo
público y solo `/api/git/*` queda detrás de oauth2-proxy. El ajuste `origin` de
slides-gitd debe ser igual al origen del Playground. Una forja cuyo `auth` es
`user-token` (GitLab, por ejemplo) necesita el token propio de cada usuario
guardado en el servidor mediante `POST /api/git/credential`. El panel no ofrece
un campo para ello, así que lo hace el administrador o el usuario fuera de la
página.

Un límite: el Playground guarda los borradores y el último nombre de repositorio
en el perfil del navegador (IndexedDB y localStorage), como antes. En un perfil
de navegador compartido, use **Draft ▸ Forget my drafts** o perfiles separados.
En modo servidor, un repositorio recordado al que el usuario no tiene acceso se
descarta.

## Compruébelo {#check-it}

```sh
curl -sI -H 'Accept-Encoding: br' 'https://slides.example.org/playground/playground.wasm?v=v0.1.0'
```

La respuesta indica `content-type: application/wasm`, `content-encoding: br`
y `cache-control: public, max-age=31536000, immutable`, con el
`content-security-policy` de arriba. Con el inicio de sesión por delante, la misma
petición sin cookie responde `302` a `/oauth2/sign_in`.

## Referencias {#references}

- Caddy: [HTTPS automático](https://caddyserver.com/docs/automatic-https),
  [`file_server`](https://caddyserver.com/docs/caddyfile/directives/file_server),
  [`header`](https://caddyserver.com/docs/caddyfile/directives/header),
  [`reverse_proxy`](https://caddyserver.com/docs/caddyfile/directives/reverse_proxy),
  [`forward_auth`](https://caddyserver.com/docs/caddyfile/directives/forward_auth).
- [RFC 8555](https://www.rfc-editor.org/rfc/rfc8555), ACME: HTTP-01 (8.3),
  DNS-01 (8.4).
- [OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html):
  el flujo de código de autorización (3.1), la URI de redirección (3.1.2.1), el ID
  Token (2).
- oauth2-proxy: [configuración](https://oauth2-proxy.github.io/oauth2-proxy/configuration/overview),
  [proveedor OpenID Connect](https://oauth2-proxy.github.io/oauth2-proxy/configuration/providers/openid_connect).
