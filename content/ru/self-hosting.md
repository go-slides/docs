---
title: "Самостоятельный хостинг"
weight: 58
description: "Раздавайте песочницу и slides --server со своей машины через Caddy: автоматический HTTPS, нужные песочнице заголовки и вход через OpenID Connect перед ними."
tags: [playground, server, security]
---

[Песочница]({{< relref "playground" >}}) — это статические файлы, а
[`slides --server`]({{< relref "cli" >}}) — один процесс на loopback. Оба можно
запустить на своей машине, под своим именем, для людей, входящих с учётной
записью вашей организации. Эта страница делает это через
[Caddy](https://caddyserver.com/docs/) (проверено с v2.11.7), который сам
получает и продлевает сертификаты HTTPS, и
[oauth2-proxy](https://oauth2-proxy.github.io/oauth2-proxy/) для входа. Все
показанные здесь файлы лежат в
[`selfhost/`](https://github.com/go-slides/go-slides.github.io/tree/main/selfhost),
а CI проверяет каждый Caddyfile командой `caddy validate`.

## Песочница {#the-playground}

Каждый выпуск сайта прикладывает весь сайт, включая песочницу, одним архивом
tar, собранным для корня домена, вместе с его SHA-256:

```sh
curl -fLO https://github.com/go-slides/go-slides.github.io/releases/download/v0.1.0/go-slides-site-v0.1.0.tar.gz
curl -fLO https://github.com/go-slides/go-slides.github.io/releases/download/v0.1.0/go-slides-site-v0.1.0.tar.gz.sha256
sha256sum -c go-slides-site-v0.1.0.tar.gz.sha256
tar -xzf go-slides-site-v0.1.0.tar.gz -C /srv
mv /srv/go-slides-site-v0.1.0 /srv/go-slides-site
```

Архив воспроизводим: повторная сборка того же тега даёт те же байты. Рядом с
каждым файлом, который сжимается (модули WebAssembly, шрифты, страницы,
скрипты), он содержит тот же файл, сжатый Brotli (`.br`) и gzip (`.gz`), так
что сервер отдаёт их как есть и ничего не сжимает на лету.

### Caddyfile {#caddyfile}

`slides.example.org` — это ваше имя. Его записи A и AAAA указывают на машину, а
порты 80 и 443 доходят до Caddy:

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

Песочница после этого доступна по адресу `https://slides.example.org/playground/`.

### Что сервер должен делать правильно {#what-the-server-must-get-right}

- **HTTPS.** Caddy получает сертификат от Let's Encrypt по ACME
  ([RFC 8555](https://www.rfc-editor.org/rfc/rfc8555)), доказывая, что владеет
  именем, ответом на порту 80 (HTTP-01) или 443 (TLS-ALPN-01), и продлевает его.
  Машина, до которой центр сертификации не может добраться, может использовать
  вызов DNS-01 (запись TXT, для чего нужен Caddy, собранный с модулем вашего
  DNS-провайдера) или собственный сертификат через `tls cert.pem key.pem`. На
  тестовой машине `tls internal` использует локальный центр сертификации Caddy.
- **`application/wasm`.** Caddy отправляет этот тип для `.wasm`; другому
  серверу, возможно, придётся об этом сказать. Песочница компилирует свои
  модули из прочитанных байтов (движок экспорта — по мере загрузки, в
  собственном ответе с этим типом), поэтому запускается даже с сервера,
  отправляющего другой тип; при `nosniff` правильный тип всё же стоит
  отправлять.
- **Сжатие.** `precompressed br gzip` отправляет `playground.wasm.br`
  браузеру, принимающему Brotli, `.gz` — принимающему gzip, а иначе сам файл, с
  `Vary: Accept-Encoding`. Модуль редактора занимает около 7 МиБ в Brotli
  против 33 МиБ без сжатия.
- **Кэширование.** Страница запрашивает каждый модуль, worker, шрифт и картинку
  с выпуском в запросе (`playground.wasm?v=v0.1.0`), поэтому их можно хранить
  год (`immutable`). Сами страницы перепроверяются (`no-cache`): новый выпуск
  виден при следующем посещении.
- **Заголовки в `headers.caddy`:**

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

Политика — та, что нужна песочнице: её скрипты (загрузчик встроенный), её
WebAssembly (`'wasm-unsafe-eval'`) и её worker, всё с вашего сервера.
Картинки, таблицы стилей и шрифты презентации могут приходить с любого адреса
`https:`, как и в выводе команды; уберите `https:` из `img-src`, `style-src`,
`font-src`, `media-src` и `connect-src`, чтобы все запросы шли на ваш сервер, —
тогда презентации, ссылающиеся на сторонние ресурсы, будут показываться без
них. Сама презентация работает в iframe, изолированных без вашего источника, под
собственной, более строгой политикой, которую песочница записывает в каждый
фрейм: он не может обратиться ни к песочнице, ни к её хранилищу, ни к вашим
cookie, что бы ни говорил этот заголовок. `frame-ancestors 'self'` не даёт
другим сайтам встраивать песочницу во фреймы.

- **Без COOP и COEP.** Песочница не использует ни `SharedArrayBuffer`, ни
  потоки, поэтому ей не нужна кросс-источниковая изоляция. Отправка
  `Cross-Origin-Embedder-Policy: require-corp` заблокировала бы картинки
  презентации с других сайтов.

### Без сети {#without-the-network}

Когда страница загружена, всё происходит в браузере: редактор, предпросмотр,
режим докладчика и экспорт в PDF, PPTX и ODP, шрифты для которых лежат в
архиве. Презентация не отправляется ни на ваш сервер, ни куда-либо ещё. Сети
требует то, что называет сама презентация: картинки или таблицы стилей по
адресу `https:` и веб-шрифт темы `gaia`, который она запрашивает у
`fonts.bunny.net` (без него слайды переходят на другой шрифт). Ссылки на эту
документацию уходят с вашего сервера.

## `slides --server` за прокси {#slides---server-behind-the-proxy}

`slides --server` остаётся на `127.0.0.1`, а Caddy пересылает ему запросы.
Сервер проверяет заголовок `Host` каждого запроса и `Origin` WebSocket каждой
живой перезагрузки — против DNS rebinding и страниц других сайтов, — а прокси
сохраняет публичное имя в обоих. `--allowed-host` добавляет это имя к тому, что
проверки принимают; сами они остаются включёнными:

```sh
slides -s --allowed-host decks.example.org /srv/decks
```

```caddy
decks.example.org {
	reverse_proxy 127.0.0.1:8080
}
```

Живая перезагрузка проходит как есть (Caddy проксирует WebSocket без настройки).
Не используйте для этого `--host 0.0.0.0`: он отдаёт каталог всей сети в обход
прокси и его входа.

{{< callout type="info" >}}
`--allowed-host` есть в v0.14.0 и позже:
`go install github.com/go-slides/slides/cmd/slides@v0.14.0`.
{{< /callout >}}

## Вход через OpenID Connect {#sign-in-with-openid-connect}

Caddy спрашивает oauth2-proxy о каждом запросе (`forward_auth`). oauth2-proxy
отвечает `202` браузеру, который вошёл, — и Caddy его обслуживает, — и `401` в
остальных случаях, который Caddy превращает в перенаправление на вход. Вход —
это поток кода авторизации
[OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html):
браузер идёт к вашему провайдеру (Keycloak, Authentik, Entra ID, Google, Okta,
Dex...), возвращается на `/oauth2/callback` с кодом, oauth2-proxy обменивает код
на ID Token, проверяет его и ставит собственный cookie.

1. **Зарегистрируйте клиента** у вашего провайдера, конфиденциального, с потоком
   кода авторизации. Его URI перенаправления, которые должны совпадать точно:
   `https://slides.example.org/oauth2/callback` и
   `https://decks.example.org/oauth2/callback`.
2. **Положите секреты в файлы**, читаемые только oauth2-proxy: секрет клиента без
   перевода строки в конце и секрет cookie ровно из 32 случайных байт
   (`head -c 32 /dev/urandom`). Их записывает ваше хранилище секретов
   (учётные данные systemd, секреты Docker или Kubernetes, Vault...); ни один
   не попадает в файл конфигурации, командную строку или переменную окружения,
   где его покажут журналы и `ps`.
3. **Настройте oauth2-proxy** (`oauth2-proxy.cfg`):

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

`redirect_url` без хоста возвращает вход каждого сайта на его собственный
callback, так что у каждого сайта свой cookie. `whitelist_domains` — это то,
куда вход может вернуть: если попросить перейти куда-то ещё, он остаётся на
вашем сайте. `trusted_proxy_ips` доверяет заголовкам `X-Forwarded-*` только от
Caddy. `email_domains` определяет, кто может войти; у oauth2-proxy есть и
параметры для групп и списков адресов.

4. **Настройте Caddy** (`Caddyfile.oidc`):

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

oauth2-proxy и `slides` слушают только loopback: единственный путь внутрь
идёт через Caddy, а значит, через вход. Заголовки песочницы находятся внутри
блока с входом и относятся только к песочнице: страница входа oauth2-proxy — это
форма, которую остановил бы их `form-action 'none'`.

Альтернатива — [caddy-security](https://github.com/greenpau/caddy-security),
модуль Caddy, выполняющий вход внутри Caddy, без второго процесса. Ему нужен
Caddy, собранный с модулем (`xcaddy`), и собственный язык конфигурации.

## Проверка {#check-it}

```sh
curl -sI -H 'Accept-Encoding: br' 'https://slides.example.org/playground/playground.wasm?v=v0.1.0'
```

В ответе будут `content-type: application/wasm`, `content-encoding: br` и
`cache-control: public, max-age=31536000, immutable`, а также
`content-security-policy`, приведённая выше. При входе перед сайтом тот же
запрос без cookie получает ответ `302` на `/oauth2/sign_in`.

## Источники {#references}

- Caddy: [автоматический HTTPS](https://caddyserver.com/docs/automatic-https),
  [`file_server`](https://caddyserver.com/docs/caddyfile/directives/file_server),
  [`header`](https://caddyserver.com/docs/caddyfile/directives/header),
  [`reverse_proxy`](https://caddyserver.com/docs/caddyfile/directives/reverse_proxy),
  [`forward_auth`](https://caddyserver.com/docs/caddyfile/directives/forward_auth).
- [RFC 8555](https://www.rfc-editor.org/rfc/rfc8555), ACME: HTTP-01 (8.3),
  DNS-01 (8.4).
- [OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html):
  поток кода авторизации (3.1), URI перенаправления (3.1.2.1), ID Token (2).
- oauth2-proxy: [конфигурация](https://oauth2-proxy.github.io/oauth2-proxy/configuration/overview),
  [провайдер OpenID Connect](https://oauth2-proxy.github.io/oauth2-proxy/configuration/providers/openid_connect).
