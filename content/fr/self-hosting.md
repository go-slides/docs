---
title: "Auto-hébergement"
weight: 58
description: "Servir le Playground et slides --server depuis votre propre machine avec Caddy : HTTPS automatique, les en-têtes dont le Playground a besoin, et une authentification OpenID Connect devant."
tags: [playground, server, security]
---

Le [Playground]({{< relref "playground" >}}) (bac à sable en ligne) est un ensemble de fichiers statiques, et
[`slides --server`]({{< relref "cli" >}}) est un seul processus sur la boucle locale. Les deux
peuvent tourner sur votre propre machine, sous votre propre nom, pour des personnes qui se
connectent avec le compte de votre organisation. Cette page le fait avec
[Caddy](https://caddyserver.com/docs/) (testé avec la v2.11.7), qui obtient
et renouvelle lui-même les certificats HTTPS, et
[oauth2-proxy](https://oauth2-proxy.github.io/oauth2-proxy/) pour
l'authentification. Tous les fichiers montrés ici se trouvent dans
[`selfhost/`](https://github.com/go-slides/go-slides.github.io/tree/main/selfhost),
et son CI vérifie chaque Caddyfile avec `caddy validate`.

## Le Playground {#the-playground}

Chaque version du site joint le site entier, Playground compris, sous la forme
d'une seule archive tarball construite pour la racine d'un domaine, avec son SHA-256 :

```sh
curl -fLO https://github.com/go-slides/go-slides.github.io/releases/download/v0.1.0/go-slides-site-v0.1.0.tar.gz
curl -fLO https://github.com/go-slides/go-slides.github.io/releases/download/v0.1.0/go-slides-site-v0.1.0.tar.gz.sha256
sha256sum -c go-slides-site-v0.1.0.tar.gz.sha256
tar -xzf go-slides-site-v0.1.0.tar.gz -C /srv
mv /srv/go-slides-site-v0.1.0 /srv/go-slides-site
```

L'archive est reproductible : reconstruire le même tag donne les mêmes
octets. À côté de chaque fichier qui se compresse (les modules WebAssembly, les
polices, les pages, les scripts), elle contient le même fichier compressé en
Brotli (`.br`) et en gzip (`.gz`), de sorte que le serveur les envoie tels quels et
ne compresse rien à la volée.

### Caddyfile {#caddyfile}

`slides.example.org` est votre nom. Ses enregistrements A et AAAA pointent vers la
machine, et les ports 80 et 443 atteignent Caddy :

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

Le Playground est alors à l'adresse `https://slides.example.org/playground/`.

### Ce que le serveur doit faire correctement {#what-the-server-must-get-right}

- **HTTPS.** Caddy obtient le certificat auprès de Let's Encrypt avec ACME
  ([RFC 8555](https://www.rfc-editor.org/rfc/rfc8555)), en prouvant qu'il détient
  le nom par une réponse sur le port 80 (HTTP-01) ou 443 (TLS-ALPN-01), puis
  le renouvelle. Une machine que l'autorité de certification ne peut pas joindre peut utiliser le défi
  DNS-01 à la place (un enregistrement TXT, qui demande un Caddy construit avec le module de
  votre fournisseur DNS), ou son propre certificat avec `tls cert.pem key.pem`.
  Sur une machine de test, `tls internal` utilise l'autorité locale de Caddy.
- **`application/wasm`.** Caddy envoie ce type pour `.wasm` ; un autre serveur
  peut demander à être configuré. Le Playground compile ses modules à partir des octets qu'il
  a lus (le moteur d'export au fil du téléchargement, dans une réponse à part
  avec ce type), si bien qu'il démarre même depuis un serveur qui envoie un autre type ;
  avec `nosniff`, le bon type reste celui qu'il faut envoyer.
- **Compression.** `precompressed br gzip` envoie `playground.wasm.br` à un
  navigateur qui accepte Brotli, `.gz` à un navigateur qui accepte gzip, et le fichier
  lui-même sinon, avec `Vary: Accept-Encoding`. Le module de l'éditeur fait
  environ 7 Mio en Brotli, contre 33 Mio non compressé.
- **Cache.** La page demande chaque module, worker, police et image
  avec la version dans la requête (`playground.wasm?v=v0.1.0`), ils peuvent donc
  être conservés un an (`immutable`). Les pages elles-mêmes sont revalidées
  (`no-cache`) : une nouvelle version est vue à la visite suivante.
- **Les en-têtes de `headers.caddy` :**

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

La politique est celle dont le Playground a besoin : ses scripts (le chargeur est
en ligne), son WebAssembly (`'wasm-unsafe-eval'`) et son worker, tous depuis
votre serveur. Les images, feuilles de style et polices d'une présentation peuvent venir de toute
adresse `https:`, comme dans la sortie de la commande ; retirez `https:` de
`img-src`, `style-src`, `font-src`, `media-src` et `connect-src` pour garder
toutes les requêtes sur votre serveur, et les présentations qui pointent ailleurs s'afficheront sans
ces ressources. La présentation elle-même s'exécute dans des iframes isolées (sandbox) sans votre
origine, sous une politique plus stricte qui leur est propre et que le Playground écrit
dans chaque cadre : elle ne peut atteindre ni le Playground, ni son stockage, ni vos
cookies, quoi que dise cet en-tête. `frame-ancestors 'self'` empêche les autres
sites d'inclure le Playground dans un cadre.

- **Ni COOP ni COEP.** Le Playground n'utilise ni `SharedArrayBuffer` ni
  threads, il n'a donc pas besoin d'être isolé en cross-origin. Envoyer
  `Cross-Origin-Embedder-Policy: require-corp` bloquerait les images d'une présentation
  venant d'autres sites.

### Sans le réseau {#without-the-network}

Une fois la page chargée, tout se passe dans le navigateur : l'éditeur,
l'aperçu, la vue présentateur, et les exports PDF, PPTX et ODP, dont les
polices sont dans l'archive. La présentation n'est envoyée ni à votre serveur ni
ailleurs. Ce qui demande le réseau, c'est ce que la présentation nomme : des images ou
des feuilles de style à une adresse `https:`, et la police web du thème `gaia`, qu'il
demande à `fonts.bunny.net` (sans elle, les diapositives se rabattent sur une autre
police). Les liens vers cette documentation quittent votre serveur.

## `slides --server` derrière le proxy {#slides---server-behind-the-proxy}

`slides --server` reste sur `127.0.0.1` et Caddy lui transmet les requêtes. Le serveur
vérifie l'en-tête `Host` de chaque requête et l'`Origin` du WebSocket de chaque
rechargement à chaud, contre le DNS rebinding et les pages d'autres sites, et le
proxy conserve le nom public dans les deux. `--allowed-host` ajoute ce nom à
ce que les contrôles acceptent ; ils restent actifs :

```sh
slides -s --allowed-host decks.example.org /srv/decks
```

```caddy
decks.example.org {
	reverse_proxy 127.0.0.1:8080
}
```

Le rechargement à chaud passe tel quel (Caddy relaie les WebSockets sans
configuration). N'utilisez pas `--host 0.0.0.0` pour cela : il servirait le
répertoire à tout le réseau, en contournant le proxy et son authentification.

{{< callout type="info" >}}
`--allowed-host` est dans la v0.14.0 et les suivantes :
`go install github.com/go-slides/slides/cmd/slides@v0.16.0`. Il demande Go 1.27.3
ou plus récent.
{{< /callout >}}

## Authentification avec OpenID Connect {#sign-in-with-openid-connect}

Caddy interroge oauth2-proxy pour chaque requête (`forward_auth`). oauth2-proxy
répond `202` pour un navigateur authentifié, que Caddy sert alors,
et `401` sinon, que Caddy transforme en redirection vers la page de connexion. La
connexion suit le flux par code d'autorisation d'
[OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html) :
le navigateur va chez votre fournisseur (Keycloak, Authentik, Entra ID, Google,
Okta, Dex...), revient sur `/oauth2/callback` avec un code, oauth2-proxy
échange le code contre un ID Token, le vérifie et pose son propre cookie.

1. **Enregistrez un client** auprès de votre fournisseur, confidentiel, avec le
   flux par code d'autorisation. Ses URI de redirection, qui doivent correspondre exactement :
   `https://slides.example.org/oauth2/callback` et
   `https://decks.example.org/oauth2/callback`.
2. **Mettez les secrets dans des fichiers** lisibles par oauth2-proxy seul : le
   secret du client, sans saut de ligne final, et un secret de cookie de
   exactement 32 octets aléatoires (`head -c 32 /dev/urandom`). Votre
   gestionnaire de secrets les écrit (credentials systemd, secrets Docker ou Kubernetes,
   Vault...) ; aucun ne va dans un fichier de configuration, une ligne
   de commande ou une variable d'environnement, où les journaux et `ps` les montrent.
3. **Configurez oauth2-proxy** (`oauth2-proxy.cfg`) :

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

`redirect_url` sans hôte ramène la connexion de chaque site vers son propre
callback, de sorte que chaque site garde son propre cookie. `whitelist_domains` est l'endroit où une
connexion peut revenir : si on lui demande d'aller ailleurs, elle reste sur votre site.
`trusted_proxy_ips` ne croit les en-têtes `X-Forwarded-*` que s'ils viennent de Caddy.
`email_domains` dit qui peut entrer ; oauth2-proxy a aussi des options pour
les groupes et les listes d'adresses.

4. **Configurez Caddy** (`Caddyfile.oidc`) :

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

oauth2-proxy et `slides` n'écoutent que sur la boucle locale : la seule entrée passe
par Caddy, donc par l'authentification. Les en-têtes du Playground vont
à l'intérieur du bloc authentifié, pour le Playground seul : la page de connexion
d'oauth2-proxy est un formulaire, que leur `form-action 'none'` bloquerait.

L'alternative est [caddy-security](https://github.com/greenpau/caddy-security),
un module Caddy qui fait l'authentification dans Caddy, sans second processus.
Il demande un Caddy construit avec le module (`xcaddy`), et son propre langage
de configuration.

## Vérifier {#check-it}

```sh
curl -sI -H 'Accept-Encoding: br' 'https://slides.example.org/playground/playground.wasm?v=v0.1.0'
```

La réponse indique `content-type: application/wasm`, `content-encoding: br`
et `cache-control: public, max-age=31536000, immutable`, avec le
`content-security-policy` ci-dessus. Avec l'authentification devant, la même
requête sans cookie répond `302` vers `/oauth2/sign_in`.

## Références {#references}

- Caddy : [HTTPS automatique](https://caddyserver.com/docs/automatic-https),
  [`file_server`](https://caddyserver.com/docs/caddyfile/directives/file_server),
  [`header`](https://caddyserver.com/docs/caddyfile/directives/header),
  [`reverse_proxy`](https://caddyserver.com/docs/caddyfile/directives/reverse_proxy),
  [`forward_auth`](https://caddyserver.com/docs/caddyfile/directives/forward_auth).
- [RFC 8555](https://www.rfc-editor.org/rfc/rfc8555), ACME : HTTP-01 (8.3),
  DNS-01 (8.4).
- [OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html) :
  le flux par code d'autorisation (3.1), l'URI de redirection (3.1.2.1), l'ID
  Token (2).
- oauth2-proxy : [configuration](https://oauth2-proxy.github.io/oauth2-proxy/configuration/overview),
  [fournisseur OpenID Connect](https://oauth2-proxy.github.io/oauth2-proxy/configuration/providers/openid_connect).
