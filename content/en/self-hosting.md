---
title: "Self-hosting"
weight: 58
description: "Serve the playground and slides --server from your own machine with Caddy: automatic HTTPS, the headers the playground needs, and an OpenID Connect sign-in in front."
tags: [playground, server, security]
---

The [playground]({{< relref "playground" >}}) is static files, and
[`slides --server`]({{< relref "cli" >}}) is one process on the loopback. Both
can run on your own machine, behind your own name, for people who sign in
with your organisation's account. This page does it with
[Caddy](https://caddyserver.com/docs/) (tested with v2.11.7), which obtains
and renews the HTTPS certificates itself, and
[oauth2-proxy](https://oauth2-proxy.github.io/oauth2-proxy/) for the
sign-in. Every file shown here is in
[`selfhost/`](https://github.com/go-slides/go-slides.github.io/tree/main/selfhost),
and its CI checks each Caddyfile with `caddy validate`.

## The playground

Each release of the site attaches the whole site, playground included, as
one tarball built for the root of a domain, with its SHA-256:

```sh
curl -fLO https://github.com/go-slides/go-slides.github.io/releases/download/v0.1.0/go-slides-site-v0.1.0.tar.gz
curl -fLO https://github.com/go-slides/go-slides.github.io/releases/download/v0.1.0/go-slides-site-v0.1.0.tar.gz.sha256
sha256sum -c go-slides-site-v0.1.0.tar.gz.sha256
tar -xzf go-slides-site-v0.1.0.tar.gz -C /srv
mv /srv/go-slides-site-v0.1.0 /srv/go-slides-site
```

The tarball is reproducible: building the same tag again gives the same
bytes. Beside every file that compresses (the WebAssembly modules, the
fonts, the pages, the scripts) it carries the same file compressed with
Brotli (`.br`) and gzip (`.gz`), so the server sends them as they are and
compresses nothing on the fly.

### Caddyfile

`slides.example.org` is your name. Its A and AAAA records point at the
machine, and ports 80 and 443 reach Caddy:

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

The playground is then at `https://slides.example.org/playground/`.

### What the server must get right

- **HTTPS.** Caddy gets the certificate from Let's Encrypt with ACME
  ([RFC 8555](https://www.rfc-editor.org/rfc/rfc8555)), proving that it holds
  the name by answering on port 80 (HTTP-01) or 443 (TLS-ALPN-01), and
  renews it. A machine that the CA cannot reach can use the DNS-01
  challenge instead (a TXT record, which needs a Caddy built with your DNS
  provider's module), or your own certificate with `tls cert.pem key.pem`.
  On a test machine, `tls internal` uses Caddy's local CA.
- **`application/wasm`.** Caddy sends this type for `.wasm`; another server
  may need to be told. The playground compiles its modules from the bytes it
  has read (the export engine as it downloads, in a response of its own
  with that type), so it starts even from a server that sends another type;
  with `nosniff`, the right type is still the one to send.
- **Compression.** `precompressed br gzip` sends `playground.wasm.br` to a
  browser that accepts Brotli, `.gz` to one that accepts gzip, and the file
  itself otherwise, with `Vary: Accept-Encoding`. The editor's module is
  about 7 MiB in Brotli, against 33 MiB uncompressed.
- **Caching.** The page asks for every module, worker, font and picture
  with the release in the query (`playground.wasm?v=v0.1.0`), so those can
  be kept for a year (`immutable`). The pages themselves are revalidated
  (`no-cache`): a new release is seen at the next visit.
- **The headers in `headers.caddy`:**

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

The policy is the one the playground needs: its scripts (the loader is
inline), its WebAssembly (`'wasm-unsafe-eval'`) and its worker, all from
your server. A deck's pictures, style sheets and fonts may come from any
`https:` address, as they do in the command's output; drop `https:` from
`img-src`, `style-src`, `font-src`, `media-src` and `connect-src` to keep
every request on your server, and decks that point elsewhere show without
those resources. The deck itself runs in iframes sandboxed without your
origin, under a stricter policy of their own that the playground writes
into each frame: it cannot reach the playground, its storage or your
cookies, whatever this header says. `frame-ancestors 'self'` stops other
sites from framing the playground.

- **No COOP or COEP.** The playground uses neither `SharedArrayBuffer` nor
  threads, so it does not need to be cross-origin isolated. Sending
  `Cross-Origin-Embedder-Policy: require-corp` would block a deck's pictures
  from other sites.

### Without the network

Once the page is loaded, everything happens in the browser: the editor,
the preview, the presenter view, and the PDF, PPTX and ODP exports, whose
fonts are in the tarball. The deck is sent neither to your server nor
anywhere else. What needs the network is what the deck names: pictures or
style sheets at an `https:` address, and the `gaia` theme's web font, which
it asks of `fonts.bunny.net` (without it, the slides fall back to another
font). The links to this documentation leave your server.

## `slides --server` behind the proxy

`slides --server` stays on `127.0.0.1` and Caddy forwards to it. The server
checks the `Host` header of each request and the `Origin` of each live
reload's WebSocket, against DNS rebinding and other sites' pages, and the
proxy keeps the public name in both. `--allowed-host` adds that name to
what the checks accept; they stay on:

```sh
slides -s --allowed-host decks.example.org /srv/decks
```

```caddy
decks.example.org {
	reverse_proxy 127.0.0.1:8080
}
```

The live reload goes through as it is (Caddy proxies WebSockets without
configuration). Do not use `--host 0.0.0.0` for this: it serves the
directory to the whole network, around the proxy and its sign-in.

{{< callout type="info" >}}
`--allowed-host` is in v0.14.0 and later:
`go install github.com/go-slides/slides/cmd/slides@v0.15.0`. It needs Go 1.27.3
or later.
{{< /callout >}}

## Sign-in with OpenID Connect

Caddy asks oauth2-proxy about every request (`forward_auth`). oauth2-proxy
answers `202` for a browser that has signed in, which Caddy then serves,
and `401` otherwise, which Caddy turns into a redirect to the sign-in. The
sign-in is the authorization code flow of
[OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html):
the browser goes to your provider (Keycloak, Authentik, Entra ID, Google,
Okta, Dex...), comes back to `/oauth2/callback` with a code, oauth2-proxy
exchanges the code for an ID Token, checks it, and sets its own cookie.

1. **Register a client** with your provider, confidential, with the
   authorization code flow. Its redirect URIs, which must match exactly:
   `https://slides.example.org/oauth2/callback` and
   `https://decks.example.org/oauth2/callback`.
2. **Put the secrets in files** readable by oauth2-proxy alone: the
   client's secret, with no newline at the end, and a cookie secret of
   exactly 32 random bytes (`head -c 32 /dev/urandom`). Your
   secret store writes them (systemd credentials, Docker or Kubernetes
   secrets, Vault...); neither goes into a configuration file, a command
   line or an environment variable, where logs and `ps` show them.
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

`redirect_url` without a host brings each site's sign-in back to its own
callback, so each site keeps its own cookie. `whitelist_domains` is where a
sign-in may return to: asked to go anywhere else, it stays on your site.
`trusted_proxy_ips` believes the `X-Forwarded-*` headers from Caddy alone.
`email_domains` says who may come in; oauth2-proxy also has options for
groups and lists of addresses.

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

oauth2-proxy and `slides` listen on the loopback only: the one way in is
through Caddy, and so through the sign-in. The playground's headers go
inside the signed-in block, on the playground alone: oauth2-proxy's sign-in
page is a form, which their `form-action 'none'` would stop.

The alternative is [caddy-security](https://github.com/greenpau/caddy-security),
a Caddy module that does the sign-in inside Caddy, with no second process.
It needs a Caddy built with the module (`xcaddy`), and its own
configuration language.

## Check it

```sh
curl -sI -H 'Accept-Encoding: br' 'https://slides.example.org/playground/playground.wasm?v=v0.1.0'
```

The answer says `content-type: application/wasm`, `content-encoding: br`
and `cache-control: public, max-age=31536000, immutable`, with the
`content-security-policy` above. With the sign-in in front, the same
request without a cookie answers `302` to `/oauth2/sign_in`.

## References

- Caddy: [automatic HTTPS](https://caddyserver.com/docs/automatic-https),
  [`file_server`](https://caddyserver.com/docs/caddyfile/directives/file_server),
  [`header`](https://caddyserver.com/docs/caddyfile/directives/header),
  [`reverse_proxy`](https://caddyserver.com/docs/caddyfile/directives/reverse_proxy),
  [`forward_auth`](https://caddyserver.com/docs/caddyfile/directives/forward_auth).
- [RFC 8555](https://www.rfc-editor.org/rfc/rfc8555), ACME: HTTP-01 (8.3),
  DNS-01 (8.4).
- [OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html):
  the authorization code flow (3.1), the redirect URI (3.1.2.1), the ID
  Token (2).
- oauth2-proxy: [configuration](https://oauth2-proxy.github.io/oauth2-proxy/configuration/overview),
  [OpenID Connect provider](https://oauth2-proxy.github.io/oauth2-proxy/configuration/providers/openid_connect).
