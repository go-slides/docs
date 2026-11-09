---
title: "セルフホスティング"
weight: 58
description: "Caddy を使って、Playground と slides --server を自分のマシンから配信します。自動 HTTPS、Playground に必要なヘッダー、前段の OpenID Connect サインインまで説明します。"
tags: [playground, server, security]
---

[Playground]({{< relref "playground" >}})は静的ファイルの集まりで、
[`slides --server`]({{< relref "cli" >}})はループバック上で動く 1 つのプロセスです。どちらも
自分のマシン上で、自分のドメイン名のもとに動かし、組織のアカウントでサインインする人に
提供できます。このページでは、HTTPS 証明書を自分で取得・更新する
[Caddy](https://caddyserver.com/docs/)（v2.11.7 で動作確認）と、
サインインのための
[oauth2-proxy](https://oauth2-proxy.github.io/oauth2-proxy/) を使います。ここに示すファイルはすべて
[`selfhost/`](https://github.com/go-slides/go-slides.github.io/tree/main/selfhost)
にあり、その CI が各 Caddyfile を `caddy validate` で検証しています。

## Playground {#the-playground}

サイトの各リリースには、Playground を含むサイト全体が、
ドメインのルート向けにビルドされた 1 つの tarball として、SHA-256 とともに添付されています。

```sh
curl -fLO https://github.com/go-slides/go-slides.github.io/releases/download/v0.1.0/go-slides-site-v0.1.0.tar.gz
curl -fLO https://github.com/go-slides/go-slides.github.io/releases/download/v0.1.0/go-slides-site-v0.1.0.tar.gz.sha256
sha256sum -c go-slides-site-v0.1.0.tar.gz.sha256
tar -xzf go-slides-site-v0.1.0.tar.gz -C /srv
mv /srv/go-slides-site-v0.1.0 /srv/go-slides-site
```

この tarball は再現可能です。同じタグをもう一度ビルドすると、同じバイト列になります。
圧縮できるファイル（WebAssembly モジュール、フォント、ページ、スクリプト）のそれぞれの隣には、
同じファイルを Brotli（`.br`）と gzip（`.gz`）で圧縮したものも入っているので、サーバーはそれをそのまま送り、
その場で圧縮することはありません。

### Caddyfile {#caddyfile}

`slides.example.org` はあなたのドメイン名です。その A レコードと AAAA レコードがこのマシンを指し、
ポート 80 と 443 が Caddy に届くようにします。

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

これで Playground は `https://slides.example.org/playground/` で使えます。

### サーバーが正しく行うべきこと {#what-the-server-must-get-right}

- **HTTPS。** Caddy は ACME（[RFC 8555](https://www.rfc-editor.org/rfc/rfc8555)）で
  Let's Encrypt から証明書を取得します。80 番ポート（HTTP-01）または 443 番ポート（TLS-ALPN-01）で
  応答して名前の所有を証明し、証明書を更新します。CA から到達できないマシンでは、代わりに
  DNS-01 チャレンジ（TXT レコード。DNS プロバイダーのモジュールを組み込んだ Caddy が必要）か、
  `tls cert.pem key.pem` による自前の証明書を使えます。
  テスト用のマシンでは、`tls internal` が Caddy のローカル CA を使います。
- **`application/wasm`。** Caddy は `.wasm` にこの型を付けて送ります。他のサーバーでは
  設定が必要かもしれません。Playground は読み込んだバイト列からモジュールをコンパイルする
  （エクスポートエンジンはダウンロードしながら、この型の専用レスポンスで）ので、別の型を送るサーバーからでも
  起動しますが、`nosniff` がある場合でも、正しい型を送るのが望ましい形です。
- **圧縮。** `precompressed br gzip` は、Brotli を受け付けるブラウザーには `playground.wasm.br` を、
  gzip を受け付けるブラウザーには `.gz` を、どちらも受け付けなければファイルそのものを、
  `Vary: Accept-Encoding` を付けて送ります。エディターのモジュールは Brotli で約 7 MiB、
  無圧縮では 33 MiB です。
- **キャッシュ。** ページは、モジュール、ワーカー、フォント、画像のすべてを、
  クエリにリリースを付けて（`playground.wasm?v=v0.1.0`）要求するので、これらは
  1 年間保持できます（`immutable`）。ページ自体は毎回再検証されます
  （`no-cache`）。新しいリリースは次の訪問で反映されます。
- **`headers.caddy` のヘッダー：**

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

このポリシーは Playground に必要なものです。スクリプト（ローダーはインライン）、
WebAssembly（`'wasm-unsafe-eval'`）、ワーカーは、すべてあなたのサーバーから読み込まれます。
デッキの画像、スタイルシート、フォントは、コマンドの出力と同様に、任意の `https:` アドレスから
読み込めます。すべてのリクエストを自分のサーバーにとどめたい場合は、`img-src`、`style-src`、
`font-src`、`media-src`、`connect-src` から `https:` を外してください。
他所を指しているデッキは、それらのリソースなしで表示されます。デッキ自体は、あなたのオリジンを持たない
サンドボックス化された iframe の中で、Playground が各フレームに書き込む、より厳しい独自のポリシーのもとで動きます。
このヘッダーの内容にかかわらず、デッキは Playground にも、そのストレージにも、あなたの
Cookie にも届きません。`frame-ancestors 'self'` は、他のサイトが Playground を
フレームに埋め込むのを防ぎます。

- **COOP も COEP もなし。** Playground は `SharedArrayBuffer` もスレッドも使わないので、
  クロスオリジン分離は不要です。
  `Cross-Origin-Embedder-Policy: require-corp` を送ると、デッキが他のサイトから読み込む画像が
  ブロックされてしまいます。

### ネットワークなしで {#without-the-network}

ページを読み込んだ後は、エディター、プレビュー、発表者ビュー、そして PDF・PPTX・ODP へのエクスポート
（フォントは tarball に含まれています）まで、すべてブラウザー内で完結します。デッキはあなたの
サーバーにも他のどこにも送られません。ネットワークが必要なのは、デッキが名指ししているもの、
つまり `https:` アドレスの画像やスタイルシート、そして `gaia` テーマの Web フォントです。
このフォントは `fonts.bunny.net` に要求します（取得できない場合、スライドは別のフォントで表示されます）。
このドキュメントへのリンクは、あなたのサーバーの外に出ます。

## プロキシの背後で `slides --server` を動かす {#slides---server-behind-the-proxy}

`slides --server` は `127.0.0.1` にとどまり、Caddy がそこへ転送します。サーバーは、DNS リバインディングや
他のサイトのページからの攻撃を防ぐため、各リクエストの `Host` ヘッダーと、各ライブリロードの
WebSocket の `Origin` を確認します。プロキシはどちらにも公開名を保ちます。
`--allowed-host` は、その名前をチェックが受け付ける対象に加えます。チェックは有効なままです。

```sh
slides -s --allowed-host decks.example.org /srv/decks
```

```caddy
decks.example.org {
	reverse_proxy 127.0.0.1:8080
}
```

ライブリロードはそのまま通ります（Caddy は設定なしで WebSocket をプロキシします）。
このために `--host 0.0.0.0` を使ってはいけません。ディレクトリがネットワーク全体に公開され、
プロキシとそのサインインを素通りしてしまいます。

{{< callout type="info" >}}
`--allowed-host` は v0.14.0 以降にあります。
`go install github.com/go-slides/slides/cmd/slides@v0.20.0`
Go 1.27.3 以降が必要です。
{{< /callout >}}

## OpenID Connect によるサインイン {#sign-in-with-openid-connect}

Caddy はすべてのリクエストについて oauth2-proxy に問い合わせます（`forward_auth`）。oauth2-proxy は、
サインイン済みのブラウザーには `202` を返し、Caddy はそのリクエストを配信します。
それ以外には `401` を返し、Caddy はこれをサインインへのリダイレクトに変えます。
サインインは
[OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html)
の認可コードフローです。ブラウザーはプロバイダー（Keycloak、Authentik、Entra ID、Google、
Okta、Dex など）に行き、コードを持って `/oauth2/callback` に戻り、oauth2-proxy が
そのコードを ID Token に交換して検証し、独自の Cookie を設定します。

1. プロバイダーに**クライアントを登録**します。機密クライアントで、認可コードフローを使います。
   リダイレクト URI は完全に一致する必要があります。
   `https://slides.example.org/oauth2/callback` と
   `https://decks.example.org/oauth2/callback` です。
2. **シークレットはファイルに置き**、oauth2-proxy だけが読めるようにします。クライアントの
   シークレット（末尾に改行を付けない）と、ちょうど 32 バイトのランダムな
   Cookie シークレット（`head -c 32 /dev/urandom`）です。
   シークレットストア（systemd の資格情報、Docker や Kubernetes の secret、Vault など）が
   書き出します。どちらも、設定ファイル、コマンドライン、環境変数には入れないでください。
   ログや `ps` に表示されてしまいます。
3. **oauth2-proxy を設定**します（`oauth2-proxy.cfg`）。

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

ホストを含まない `redirect_url` は、各サイトのサインインをそのサイト自身のコールバックに戻すので、
各サイトは独自の Cookie を持ちます。`whitelist_domains` はサインイン後に戻ってよい先です。
それ以外の場所へ行くよう求められても、サインインはあなたのサイトにとどまります。
`trusted_proxy_ips` は、Caddy からの `X-Forwarded-*` ヘッダーだけを信頼します。
`email_domains` は入れる人を決めます。oauth2-proxy には、グループやアドレス一覧のための
オプションもあります。

4. **Caddy を設定**します（`Caddyfile.oidc`）。

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

oauth2-proxy と `slides` はループバックだけで待ち受けます。入口は Caddy だけで、
したがってサインインを必ず通ります。Playground のヘッダーはサインイン済みブロックの内側に置き、
Playground にだけ適用します。oauth2-proxy のサインインページはフォームであり、
`form-action 'none'` がそれを止めてしまうからです。

別の方法として、[caddy-security](https://github.com/greenpau/caddy-security) があります。
これは Caddy の中でサインインを行う Caddy モジュールで、2 つ目のプロセスは要りません。
ただし、このモジュールを組み込んだ Caddy（`xcaddy`）と、独自の設定言語が必要です。

## slides-gitd：認証情報はサーバーに残る {#slides-gitd}

[slides-gitd](https://github.com/go-slides/gitd)（v0.1.0）は、セルフホストした
Playground の git バックエンドです。forge からスライドのフォルダーを読み込み、変更を
新しいブランチとプル（マージ）リクエストで提案し、上流の変更を確認します。forge の認証情報は
**サーバーに残り**、ブラウザーが持つのは HttpOnly のセッション cookie だけです。Caddy と
oauth2-proxy（OIDC）の背後に置き、`127.0.0.1:8091` で待ち受け、ユーザーの身元はプロキシからだけ
受け取り、サーバー自身が名付けた `slides/<ユーザー>/<タイトル>-<時刻>` というブランチにしか
書き込みません。

ブラウザーがトークンを持つ代わりに slides-gitd と通信する Playground のサーバーモードは、
[slides-gitd の背後の Playground](#playground-server-mode)で説明します。

設定（秘密はファイルパスで、値そのものは書きません。各ファイルは所有者だけが読める必要があります）：

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

ユーザーが届くのは、その OIDC サブジェクトに列挙された `forge:所有者/名前` だけです
（既定で拒否）。マスターキーは
`head -c 32 /dev/urandom > master.key && chmod 600 master.key` で作り、
`slides-gitd check -c config.json`（すべての秘密ファイルを読み込んで終了します）を実行してから
`slides-gitd serve -c config.json` を起動します。

Caddy 側です。oauth2-proxy は `--set-xauthrequest` で起動します。

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

`request_header -…` の行は、クライアントが送った身元ヘッダーをすべて削除するので、
クライアントは自分が誰かを選べません。新しい Caddy は `copy_headers` の前に自分で削除しますが、
古いバージョンはしません。`{file.*}` はプロキシの秘密をファイルから読むので、環境変数には
置かれません。

**GitHub** は App を使います。Settings ▸ Developer settings ▸ GitHub Apps で、
Contents の読み取りと書き込み、Pull requests の読み取りと書き込み、Metadata の読み取りの
権限で作成し（webhook なし）、スライドのリポジトリにだけインストールします。slides-gitd は
呼び出しごとにその 1 つのリポジトリ向けのインストールトークンを発行し、有効期間は 1 時間です。

**GitLab：** 各ユーザーはスコープ `api` の個人またはプロジェクトのアクセストークンを保存します。
`write_repository` ではなく `api` が必要な理由と、盗まれたトークンでできることは
[Playground のページ]({{< relref "playground#gitlab" >}})で説明しています。スライドの
リポジトリに限定したプロジェクトまたはグループのアクセストークン、または専用アカウントを使ってください。

### slides-gitd の背後の Playground（サーバーモード） {#playground-server-mode}

Playground を slides-gitd の背後で独自のオリジンから配信すると、Playground はそのオリジンに
`GET /api/git/session` を問い合わせて、そのことを検出します。`401` はサーバーがあってまだ誰も
サインインしていないことを意味し、ペインには **サインイン** が表示されます。JSON のセッションが
返ればサインイン済みです。それ以外、たとえば静的ホストの `404` は、サーバーがないことを意味し、
ペインは go-slides.github.io とまったく同じように動作します。つまり、タブに保持するトークン、
または GitLab サインインを使います。

サーバーモードでは、ペインはトークンと OAuth のフィールドを隠します。代わりに
`Signed in as <identity>`（サインイン中のユーザー）と、サーバーが許可したリポジトリの一覧を表示します。
**読み込む**、**変更を提案…**（ブランチ `slides/<readable>-<hash>/…` はサーバーが名付け、ページからは選べません）、
**変更を確認**、**向こうを採用** は、すべて同じオリジンの `/api/git/*` に送られます。
ブラウザーが forge のホストに接続することはありません。

何がどこに残るか：

- トークンはサーバー上の、ユーザーごとに 1 つの鍵を持つ AES-GCM の保管庫にあります。
- ブラウザーが持つのは `__Host-gitd` cookie（`HttpOnly`、`Secure`、`SameSite=Strict`）だけで、
  スクリプトからは読めません。
- CSRF トークンはページのメモリ内にのみ保持され、保存されません。
- すべての書き込みは、正確な `Origin`、`X-Requested-With: slides-gitd`、JSON の本文、
  `X-CSRF-Token` を伴います。

サインインの期限が切れている、またはない場合（HTTP `401`）、ペインは「サインインが終了しました」と
**サインイン** ボタン（oauth2-proxy の `/oauth2/sign_in?rd=<this page>`）を表示します。
ブラウザーに入力したトークンにフォールバックすることは決してありません。

配信には go-slides/go-slides.github.io の `selfhost/Caddyfile.gitd` を使います。その
Content-Security-Policy は `connect-src 'self'` で、サイトは公開のままにし、oauth2-proxy の背後に置くのは
`/api/git/*` だけです。slides-gitd の `origin` 設定は Playground のオリジンと一致させる必要があります。
`auth` が `user-token` の forge（たとえば GitLab）では、各ユーザー自身のトークンを
`POST /api/git/credential` でサーバーに保存する必要があります。ペインにはそのためのフィールドがないので、
管理者またはユーザーがページの外で行います。

制限が 1 つあります。Playground は、これまでどおり下書きと最後のリポジトリ名をブラウザープロファイル
（IndexedDB と localStorage）に保持します。共有のブラウザープロファイルでは、**Draft ▸ Forget my drafts**
を使うか、プロファイルを分けてください。サーバーモードでは、記憶されているリポジトリのうち
ユーザーに許可されていないものは破棄されます。

## 確認する {#check-it}

```sh
curl -sI -H 'Accept-Encoding: br' 'https://slides.example.org/playground/playground.wasm?v=v0.1.0'
```

応答には `content-type: application/wasm`、`content-encoding: br`、
`cache-control: public, max-age=31536000, immutable` が含まれ、上記の
`content-security-policy` も付いています。サインインを前段に置いている場合、同じリクエストを
Cookie なしで送ると、`302` で `/oauth2/sign_in` に転送されます。

## 参考資料 {#references}

- Caddy：[automatic HTTPS](https://caddyserver.com/docs/automatic-https)、
  [`file_server`](https://caddyserver.com/docs/caddyfile/directives/file_server)、
  [`header`](https://caddyserver.com/docs/caddyfile/directives/header)、
  [`reverse_proxy`](https://caddyserver.com/docs/caddyfile/directives/reverse_proxy)、
  [`forward_auth`](https://caddyserver.com/docs/caddyfile/directives/forward_auth)
- [RFC 8555](https://www.rfc-editor.org/rfc/rfc8555)、ACME：HTTP-01（8.3）、
  DNS-01（8.4）
- [OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html)：
  認可コードフロー（3.1）、リダイレクト URI（3.1.2.1）、ID
  Token（2）
- oauth2-proxy：[configuration](https://oauth2-proxy.github.io/oauth2-proxy/configuration/overview)、
  [OpenID Connect provider](https://oauth2-proxy.github.io/oauth2-proxy/configuration/providers/openid_connect)
