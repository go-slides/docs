---
title: "Self-hosting"
weight: 58
description: "Обслуговуйте Playground і slides --server зі своєї машини за допомогою Caddy: автоматичний HTTPS, заголовки, яких потребує Playground, і вхід через OpenID Connect попереду."
tags: [playground, server, security]
---

[Playground]({{< relref "playground" >}}) — це статичні файли, а
[`slides --server`]({{< relref "cli" >}}) — один процес на loopback. Обидва
можна запустити на власній машині, під власним іменем, для людей, які входять
через обліковий запис вашої організації. Ця сторінка робить це за допомогою
[Caddy](https://caddyserver.com/docs/) (перевірено з v2.11.7), який сам
отримує й поновлює сертифікати HTTPS, і
[oauth2-proxy](https://oauth2-proxy.github.io/oauth2-proxy/) для входу. Кожен
показаний тут файл є в
[`selfhost/`](https://github.com/go-slides/go-slides.github.io/tree/main/selfhost),
і його CI перевіряє кожен Caddyfile командою `caddy validate`.

## Playground {#the-playground}

Кожен випуск сайту додає весь сайт, разом із Playground, як один tar-архів,
зібраний для кореня домену, з його SHA-256:

```sh
curl -fLO https://github.com/go-slides/go-slides.github.io/releases/download/v0.1.0/go-slides-site-v0.1.0.tar.gz
curl -fLO https://github.com/go-slides/go-slides.github.io/releases/download/v0.1.0/go-slides-site-v0.1.0.tar.gz.sha256
sha256sum -c go-slides-site-v0.1.0.tar.gz.sha256
tar -xzf go-slides-site-v0.1.0.tar.gz -C /srv
mv /srv/go-slides-site-v0.1.0 /srv/go-slides-site
```

Архів відтворюваний: повторне збирання того самого тега дає ті самі байти.
Поруч із кожним файлом, що стискається (модулі WebAssembly, шрифти, сторінки,
скрипти), він містить той самий файл, стиснутий Brotli (`.br`) і gzip (`.gz`),
тож сервер надсилає їх як є й нічого не стискає на льоту.

### Caddyfile {#caddyfile}

`slides.example.org` — це ваше ім'я. Його записи A і AAAA вказують на машину, а
порти 80 і 443 досягають Caddy:

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

Після цього Playground розташований за адресою
`https://slides.example.org/playground/`.

### Що має зробити правильно сервер {#what-the-server-must-get-right}

- **HTTPS.** Caddy отримує сертифікат від Let's Encrypt за допомогою ACME
  ([RFC 8555](https://www.rfc-editor.org/rfc/rfc8555)), доводячи, що володіє
  іменем, відповіддю на порту 80 (HTTP-01) або 443 (TLS-ALPN-01), і поновлює
  його. Машина, до якої CA не може дістатися, може натомість використати
  виклик DNS-01 (запис TXT, для якого потрібен Caddy зі збіркою модуля вашого
  DNS-провайдера) або власний сертифікат через `tls cert.pem key.pem`. На
  тестовій машині `tls internal` використовує локальний CA Caddy.
- **`application/wasm`.** Caddy надсилає цей тип для `.wasm`; іншому серверу
  може знадобитися це вказати. Playground компілює свої модулі з байтів, які
  прочитав (рушій експорту — у міру завантаження, у власній відповіді з цим
  типом), тож він запускається навіть із сервера, що надсилає інший тип; з
  `nosniff` правильний тип усе ж той, який слід надсилати.
- **Стиснення.** `precompressed br gzip` надсилає `playground.wasm.br`
  браузеру, що приймає Brotli, `.gz` — тому, що приймає gzip, і сам файл в
  іншому разі, з `Vary: Accept-Encoding`. Модуль редактора має близько 7 МіБ у
  Brotli проти 33 МіБ без стиснення.
- **Кешування.** Сторінка запитує кожен модуль, воркер, шрифт і картинку з
  випуском у запиті (`playground.wasm?v=v0.1.0`), тож їх можна зберігати рік
  (`immutable`). Самі сторінки перевіряються повторно (`no-cache`): новий
  випуск видно при наступному візиті.
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

Політика — та, що потрібна Playground: його скрипти (завантажувач
вбудований), його WebAssembly (`'wasm-unsafe-eval'`) і його воркер, усе з вашого
сервера. Зображення, таблиці стилів і шрифти презентації можуть надходити з
будь-якої адреси `https:`, як і у виводі команди; приберіть `https:` з
`img-src`, `style-src`, `font-src`, `media-src` і `connect-src`, щоб тримати
кожен запит на вашому сервері, і презентації, що вказують деінде, показуватимуться
без цих ресурсів. Сама презентація працює в iframe, ізольованих без вашого
походження, під суворішою власною політикою, яку Playground записує в кожен
фрейм: вона не може дістатися ні до Playground, ні до його сховища, ні до ваших
cookie, що б не казав цей заголовок. `frame-ancestors 'self'` не дає іншим
сайтам вбудовувати Playground у фрейми.

- **Ні COOP, ні COEP.** Playground не використовує ні `SharedArrayBuffer`, ні
  потоків, тож йому не потрібна міжджерельна ізоляція. Надсилання
  `Cross-Origin-Embedder-Policy: require-corp` заблокувало б зображення
  презентації з інших сайтів.

### Без мережі {#without-the-network}

Щойно сторінку завантажено, усе відбувається в браузері: редактор, попередній
перегляд, вигляд доповідача й експорти PDF, PPTX та ODP, шрифти яких є в
архіві. Презентація не надсилається ні на ваш сервер, ні кудись іще. Мережі
потребує те, що називає презентація: зображення чи таблиці стилів за адресою
`https:` і веб-шрифт теми `gaia`, який вона запитує в `fonts.bunny.net` (без
нього слайди переходять на інший шрифт). Посилання на цю документацію ведуть за
межі вашого сервера.

## `slides --server` за проксі {#slides---server-behind-the-proxy}

`slides --server` залишається на `127.0.0.1`, а Caddy перенаправляє до нього.
Сервер перевіряє заголовок `Host` кожного запиту й `Origin` WebSocket живого
перезавантаження, проти DNS rebinding і сторінок інших сайтів, а проксі
зберігає публічне ім'я в обох. `--allowed-host` додає це ім'я до того, що
приймають перевірки; вони лишаються ввімкненими:

```sh
slides -s --allowed-host decks.example.org /srv/decks
```

```caddy
decks.example.org {
	reverse_proxy 127.0.0.1:8080
}
```

Живе перезавантаження проходить як є (Caddy проксіює WebSocket без
налаштування). Не використовуйте для цього `--host 0.0.0.0`: він віддає теку
всій мережі, в обхід проксі та його входу.

{{< callout type="info" >}}
`--allowed-host` є у v0.14.0 і пізніших:
`go install github.com/go-slides/slides/cmd/slides@v0.18.0`. Потрібен Go 1.27.3
або новіший.
{{< /callout >}}

## Вхід через OpenID Connect {#sign-in-with-openid-connect}

Caddy питає oauth2-proxy про кожен запит (`forward_auth`). oauth2-proxy
відповідає `202` для браузера, що увійшов, і тоді Caddy його обслуговує, та
`401` в іншому разі, і тоді Caddy перетворює це на перенаправлення на вхід.
Вхід — це потік коду авторизації
[OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html):
браузер іде до вашого провайдера (Keycloak, Authentik, Entra ID, Google, Okta,
Dex...), повертається на `/oauth2/callback` з кодом, oauth2-proxy обмінює код на
ID Token, перевіряє його й ставить власний cookie.

1. **Зареєструйте клієнта** у вашого провайдера, конфіденційного, з потоком
   коду авторизації. Його redirect URI, які мають збігатися точно:
   `https://slides.example.org/oauth2/callback` та
   `https://decks.example.org/oauth2/callback`.
2. **Покладіть секрети у файли**, читані лише oauth2-proxy: секрет клієнта, без
   символу нового рядка наприкінці, і секрет cookie рівно з 32 випадкових байтів
   (`head -c 32 /dev/urandom`). Їх записує ваше сховище секретів (облікові дані
   systemd, секрети Docker чи Kubernetes, Vault...); жоден не потрапляє до файлу
   конфігурації, командного рядка чи змінної середовища, де їх показують логи
   та `ps`.
3. **Налаштуйте oauth2-proxy** (`oauth2-proxy.cfg`):

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

`redirect_url` без хоста повертає вхід кожного сайту на його власний callback,
тож кожен сайт має власний cookie. `whitelist_domains` — це те, куди вхід може
повернутися: якщо його просять піти кудись інде, він залишається на вашому
сайті. `trusted_proxy_ips` вірить заголовкам `X-Forwarded-*` лише від Caddy.
`email_domains` каже, хто може увійти; oauth2-proxy має також параметри для
груп і списків адрес.

4. **Налаштуйте Caddy** (`Caddyfile.oidc`):

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

oauth2-proxy і `slides` слухають лише на loopback: єдиний вхід — через Caddy, а
отже, через авторизацію. Заголовки Playground розміщено всередині блока
signed-in, лише для Playground: сторінка входу oauth2-proxy є формою, яку
зупинив би їхній `form-action 'none'`.

Альтернатива — [caddy-security](https://github.com/greenpau/caddy-security),
модуль Caddy, що робить вхід усередині Caddy, без другого процесу. Йому
потрібен Caddy, зібраний із модулем (`xcaddy`), і власна мова конфігурації.

## slides-gitd: облікові дані лишаються на сервері {#slides-gitd}

[slides-gitd](https://github.com/go-slides/gitd) (v0.1.0) — це git-бекенд
власного Playground. Він завантажує теку презентації з forge, пропонує зміни в
новій гілці з pull (merge) request і перевіряє зміни в першоджерелі, а облікові
дані forge **лишаються на сервері**: браузер має лише сесійний cookie HttpOnly.
Він стоїть за Caddy і oauth2-proxy (OIDC), слухає `127.0.0.1:8091`, особу
користувача бере лише від проксі й пише лише в гілки з назвою
`slides/<користувач>/<заголовок>-<час>`, які називає сам сервер.

Серверний режим Playground, який звертається до slides-gitd замість того, щоб
браузер тримав токен, описано в розділі
[Playground за slides-gitd](#playground-server-mode).

Конфігурація (секрети — це шляхи до файлів, ніколи не значення; кожен файл має
читатися лише власником):

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

Користувач сягає лише записів `forge:власник/назва`, перелічених для його
суб'єкта OIDC (типово заборонено). Створіть головний ключ командою
`head -c 32 /dev/urandom > master.key && chmod 600 master.key`, потім запустіть
`slides-gitd check -c config.json`, що завантажує всі файли секретів і
завершується, та `slides-gitd serve -c config.json`.

Бік Caddy, з oauth2-proxy, запущеним з `--set-xauthrequest`:

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

Рядки `request_header -…` видаляють будь-який заголовок особи, надісланий
клієнтом, тож клієнт не може обрати, хто він; свіжий Caddy робить це сам перед
`copy_headers`, старіші версії — ні. `{file.*}` читає секрет проксі з файлу, тож
його немає в змінній середовища.

**GitHub** використовує App. Створіть її в Settings ▸ Developer settings ▸
GitHub Apps з дозволами Contents читання й запис, Pull requests читання й запис
та Metadata читання, без webhook; встановіть її лише на репозиторії презентацій.
slides-gitd видає токен інсталяції для єдиного репозиторію кожного виклику,
чинний одну годину.

**GitLab:** кожен користувач зберігає персональний токен доступу або токен
проєкту з областю `api`. Чому `api`, а не `write_repository`, і що може вкрадений
токен, пояснено [на сторінці Playground]({{< relref "playground#gitlab" >}});
обирайте токен доступу проєкту або групи, обмежений репозиторіями презентацій,
або окремий обліковий запис.

### Playground за slides-gitd (серверний режим) {#playground-server-mode}

Коли Playground віддається за slides-gitd зі свого origin, він виявляє це,
запитуючи `GET /api/git/session` на цьому origin. `401` означає, що сервер є,
але ніхто не ввійшов: панель показує **Увійти**. JSON-сесія означає, що вхід
виконано. Усе інше, наприклад `404` статичного хоста, означає, що сервера немає,
і панель працює достеменно як на go-slides.github.io: з токеном у вкладці або зі
входом через GitLab.

У серверному режимі панель ховає поля токена й OAuth. Вона показує
`Signed in as <identity>` (ви ввійшли як) і список репозиторіїв, які надає
сервер. **Завантажити**, **Запропонувати зміни…** (гілку
`slides/<readable>-<hash>/…` називає сервер; сторінка не може вибрати свою),
**Перевірити зміни** та **Взяти їхні** звертаються до `/api/git/*` на тому самому
origin. Браузер не звертається до жодного хоста forge.

Що де лишається:

- Токени живуть на сервері, у сховищі AES-GCM з одним ключем на користувача.
- Браузер тримає лише cookie `__Host-gitd` (`HttpOnly`, `Secure`,
  `SameSite=Strict`), яку скрипти прочитати не можуть.
- Токен CSRF зберігається лише в пам’яті сторінки й ніколи не записується.
- Кожен запис несе точний `Origin`, `X-Requested-With: slides-gitd`, тіло JSON і
  `X-CSRF-Token`.

Якщо вхід минув або його немає (HTTP `401`), панель показує «Ваш вхід
завершено» з кнопкою **Увійти** (`/oauth2/sign_in?rd=<this page>` з
oauth2-proxy). Вона ніколи не повертається до токена, введеного в браузері.

Щоб віддавати його, використайте `selfhost/Caddyfile.gitd` з
go-slides/go-slides.github.io. Його Content-Security-Policy — `connect-src
'self'`; сайт лишається публічним, і лише `/api/git/*` стоїть за oauth2-proxy.
Параметр `origin` у slides-gitd має збігатися з origin Playground. Forge, у якого
`auth` дорівнює `user-token` (наприклад, GitLab), вимагає власний токен кожного
користувача, збережений на сервері через `POST /api/git/credential`. Поля для
цього на панелі немає, тож адміністратор або користувач робить це поза
сторінкою.

Одне обмеження: Playground, як і раніше, зберігає чернетки й останню назву
репозиторію в профілі браузера (IndexedDB і localStorage). У спільному профілі
браузера користуйтеся **Draft ▸ Forget my drafts** або окремими профілями. У
серверному режимі запам’ятаний репозиторій, до якого користувач не має доступу,
відкидається.

## Перевірка {#check-it}

```sh
curl -sI -H 'Accept-Encoding: br' 'https://slides.example.org/playground/playground.wasm?v=v0.1.0'
```

Відповідь містить `content-type: application/wasm`, `content-encoding: br` і
`cache-control: public, max-age=31536000, immutable`, а також наведений вище
`content-security-policy`. Якщо попереду стоїть вхід, той самий запит без cookie
відповідає `302` на `/oauth2/sign_in`.

## Джерела {#references}

- Caddy: [automatic HTTPS](https://caddyserver.com/docs/automatic-https),
  [`file_server`](https://caddyserver.com/docs/caddyfile/directives/file_server),
  [`header`](https://caddyserver.com/docs/caddyfile/directives/header),
  [`reverse_proxy`](https://caddyserver.com/docs/caddyfile/directives/reverse_proxy),
  [`forward_auth`](https://caddyserver.com/docs/caddyfile/directives/forward_auth).
- [RFC 8555](https://www.rfc-editor.org/rfc/rfc8555), ACME: HTTP-01 (8.3),
  DNS-01 (8.4).
- [OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html):
  потік коду авторизації (3.1), redirect URI (3.1.2.1), ID Token (2).
- oauth2-proxy: [configuration](https://oauth2-proxy.github.io/oauth2-proxy/configuration/overview),
  [OpenID Connect provider](https://oauth2-proxy.github.io/oauth2-proxy/configuration/providers/openid_connect).
