---
title: "自托管"
weight: 58
description: "用 Caddy 在你自己的机器上提供 Playground 和 slides --server：自动 HTTPS、Playground 所需的响应头，以及前置的 OpenID Connect 登录。"
tags: [playground, server, security]
---

[Playground]({{< relref "playground" >}}) 是一组静态文件，
[`slides --server`]({{< relref "cli" >}}) 是运行在回环地址上的单个进程。两者都
可以运行在你自己的机器上，使用你自己的域名，供用你所在组织的账号登录的人使用。
本页用 [Caddy](https://caddyserver.com/docs/)（已用 v2.11.7 测试）来完成，它会自行获取并
续期 HTTPS 证书；登录则由
[oauth2-proxy](https://oauth2-proxy.github.io/oauth2-proxy/) 负责。这里展示的每个文件都在
[`selfhost/`](https://github.com/go-slides/go-slides.github.io/tree/main/selfhost)
中，其 CI 会用 `caddy validate` 检查每个 Caddyfile。

## Playground {#the-playground}

站点的每个发布版本都会把整个站点（包括 Playground）打包成
一个为域名根路径构建的 tarball，并附上它的 SHA-256：

```sh
curl -fLO https://github.com/go-slides/go-slides.github.io/releases/download/v0.1.0/go-slides-site-v0.1.0.tar.gz
curl -fLO https://github.com/go-slides/go-slides.github.io/releases/download/v0.1.0/go-slides-site-v0.1.0.tar.gz.sha256
sha256sum -c go-slides-site-v0.1.0.tar.gz.sha256
tar -xzf go-slides-site-v0.1.0.tar.gz -C /srv
mv /srv/go-slides-site-v0.1.0 /srv/go-slides-site
```

这个 tarball 是可复现的：再次构建同一个标签会得到相同的
字节。每个可压缩的文件（WebAssembly 模块、字体、页面、脚本）旁边，
都带有用 Brotli（`.br`）和 gzip（`.gz`）压缩的同一文件，所以服务器会原样发送它们，
而不必即时压缩任何内容。

### Caddyfile {#caddyfile}

`slides.example.org` 是你的域名。它的 A 记录和 AAAA 记录指向这台
机器，80 和 443 端口可以到达 Caddy：

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

之后 Playground 就在 `https://slides.example.org/playground/`。

### 服务器必须做对的事 {#what-the-server-must-get-right}

- **HTTPS。** Caddy 通过 ACME（[RFC 8555](https://www.rfc-editor.org/rfc/rfc8555)）
  从 Let's Encrypt 获取证书，通过在 80 端口（HTTP-01）或 443 端口（TLS-ALPN-01）上应答
  来证明自己拥有该域名，并负责续期。CA 无法访问的机器可以改用
  DNS-01 质询（一条 TXT 记录，需要内置了你的 DNS 提供商模块的 Caddy），
  或者用 `tls cert.pem key.pem` 使用自己的证书。
  在测试机器上，`tls internal` 使用 Caddy 的本地 CA。
- **`application/wasm`。** Caddy 会为 `.wasm` 发送这个类型；换成别的服务器时
  可能需要手动指定。Playground 是根据读到的字节来编译模块的
  （导出引擎边下载边编译，用的是一个带此类型的独立响应），所以即使服务器发送了别的类型它也能启动；
  不过有 `nosniff` 时，发送正确的类型仍然是应有的做法。
- **压缩。** `precompressed br gzip` 会向支持 Brotli 的浏览器发送 `playground.wasm.br`，
  向支持 gzip 的浏览器发送 `.gz`，否则发送文件本身，并带上
  `Vary: Accept-Encoding`。编辑器的模块用 Brotli 约 7 MiB，
  未压缩为 33 MiB。
- **缓存。** 页面在查询中带上发布版本来请求每个模块、worker、字体和图片
  （`playground.wasm?v=v0.1.0`），所以这些文件可以
  缓存一年（`immutable`）。页面本身每次都会重新验证
  （`no-cache`）：新版本在下次访问时就能看到。
- **`headers.caddy` 中的响应头：**

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

这条策略正是 Playground 所需要的：它的脚本（加载器是
内联的）、它的 WebAssembly（`'wasm-unsafe-eval'`）和它的 worker，全部来自
你的服务器。演示文稿的图片、样式表和字体可以来自任何
`https:` 地址，与命令的输出一致；如果想让所有请求都留在你的服务器上，
请从 `img-src`、`style-src`、`font-src`、`media-src` 和 `connect-src` 中去掉 `https:`，
这样指向别处的演示文稿将不显示
这些资源。演示文稿本身运行在不带你的来源的沙盒 iframe 中，
受 Playground 写入每个框架的、更严格的专属策略约束：
无论这个响应头怎么写，它都无法访问 Playground、它的存储或你的
Cookie。`frame-ancestors 'self'` 可防止其他网站
把 Playground 嵌入框架。

- **没有 COOP 和 COEP。** Playground 既不使用 `SharedArrayBuffer`，也不使用
  线程，所以不需要跨源隔离。发送
  `Cross-Origin-Embedder-Policy: require-corp` 会阻止演示文稿加载其他网站的
  图片。

### 没有网络时 {#without-the-network}

页面加载完成后，一切都在浏览器中进行：编辑器、
预览、演讲者视图，以及 PDF、PPTX 和 ODP 导出（所需
字体都在 tarball 中）。演示文稿既不会发送到你的服务器，
也不会发送到别处。需要网络的是演示文稿自己引用的内容：`https:` 地址上的图片或
样式表，以及 `gaia` 主题的网络字体，它
向 `fonts.bunny.net` 请求这个字体（没有它，幻灯片会退回到其他
字体）。指向本文档的链接会离开你的服务器。

## 代理之后的 `slides --server` {#slides---server-behind-the-proxy}

`slides --server` 保持在 `127.0.0.1` 上，由 Caddy 转发给它。服务器会检查每个请求的 `Host` 头，
以及每个实时重载 WebSocket 的 `Origin`，以防范 DNS 重绑定和其他网站的页面，而代理在两者中
都保留公开域名。`--allowed-host` 把这个域名加入
检查所接受的范围；检查仍然开启：

```sh
slides -s --allowed-host decks.example.org /srv/decks
```

```caddy
decks.example.org {
	reverse_proxy 127.0.0.1:8080
}
```

实时重载会原样通过（Caddy 无需配置就能代理 WebSocket）。
不要为此使用 `--host 0.0.0.0`：它会把目录提供给整个网络，
绕过代理及其登录。

{{< callout type="info" >}}
`--allowed-host` 自 v0.14.0 起提供：
`go install github.com/go-slides/slides/cmd/slides@v0.14.0`。
{{< /callout >}}

## 使用 OpenID Connect 登录 {#sign-in-with-openid-connect}

Caddy 会就每个请求询问 oauth2-proxy（`forward_auth`）。对已登录的浏览器，oauth2-proxy
回答 `202`，Caddy 随后提供页面；
否则回答 `401`，Caddy 把它变成跳转到登录页的重定向。
登录采用
[OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html)
的授权码流程：浏览器前往你的提供商（Keycloak、Authentik、Entra ID、Google、
Okta、Dex 等），带着授权码回到 `/oauth2/callback`，oauth2-proxy
把授权码换成 ID Token，对其进行校验，并设置自己的 Cookie。

1. **在提供商处注册一个客户端**，机密客户端，使用授权码流程。
   它的重定向 URI 必须完全匹配：
   `https://slides.example.org/oauth2/callback` 和
   `https://decks.example.org/oauth2/callback`。
2. **把密钥放进文件**，且只有 oauth2-proxy 可读：客户端的
   密钥（末尾不带换行），以及恰好 32 个随机字节的 Cookie 密钥
   （`head -c 32 /dev/urandom`）。由你的密钥存储来写入它们
   （systemd 凭据、Docker 或 Kubernetes 的 secret、Vault 等）；
   两者都不要放进配置文件、命令行或环境变量，日志和 `ps` 会把它们显示出来。
3. **配置 oauth2-proxy**（`oauth2-proxy.cfg`）：

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

不带主机的 `redirect_url` 会让每个站点的登录回到它自己的
回调，因此每个站点都有自己的 Cookie。`whitelist_domains` 是登录后允许返回的地方：
如果要求跳转到其他任何地方，登录会留在你的站点上。
`trusted_proxy_ips` 只信任来自 Caddy 的 `X-Forwarded-*` 头。
`email_domains` 规定谁可以进入；oauth2-proxy 还提供针对
群组和地址列表的选项。

4. **配置 Caddy**（`Caddyfile.oidc`）：

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

oauth2-proxy 和 `slides` 只监听回环地址：唯一的入口是
Caddy，因而必须经过登录。Playground 的响应头放在
已登录块内部，且只用于 Playground：oauth2-proxy 的登录
页面是一个表单，而这些响应头中的 `form-action 'none'` 会阻止它。

另一种方案是 [caddy-security](https://github.com/greenpau/caddy-security)，
这是一个在 Caddy 内部完成登录的 Caddy 模块，不需要第二个进程。
它需要内置了该模块的 Caddy（`xcaddy`），以及它自己的
配置语言。

## 检查 {#check-it}

```sh
curl -sI -H 'Accept-Encoding: br' 'https://slides.example.org/playground/playground.wasm?v=v0.1.0'
```

响应会显示 `content-type: application/wasm`、`content-encoding: br`
和 `cache-control: public, max-age=31536000, immutable`，并带有上面的
`content-security-policy`。前面有登录时，同样的
请求在不带 Cookie 的情况下会返回 `302`，跳转到 `/oauth2/sign_in`。

## 参考资料 {#references}

- Caddy：[automatic HTTPS](https://caddyserver.com/docs/automatic-https)、
  [`file_server`](https://caddyserver.com/docs/caddyfile/directives/file_server)、
  [`header`](https://caddyserver.com/docs/caddyfile/directives/header)、
  [`reverse_proxy`](https://caddyserver.com/docs/caddyfile/directives/reverse_proxy)、
  [`forward_auth`](https://caddyserver.com/docs/caddyfile/directives/forward_auth)。
- [RFC 8555](https://www.rfc-editor.org/rfc/rfc8555)，ACME：HTTP-01（8.3）、
  DNS-01（8.4）。
- [OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html)：
  授权码流程（3.1）、重定向 URI（3.1.2.1）、ID
  Token（2）。
- oauth2-proxy：[configuration](https://oauth2-proxy.github.io/oauth2-proxy/configuration/overview)、
  [OpenID Connect provider](https://oauth2-proxy.github.io/oauth2-proxy/configuration/providers/openid_connect)。
