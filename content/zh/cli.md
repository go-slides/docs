---
title: "命令行"
weight: 50
description: "slides 的所有命令和参数，以及 slides.yaml 配置文件。"
tags: [cli]
---

{{< callout type="warning" >}}
本页描述的是 v0.15.0。
{{< /callout >}}

```text
slides [deck.md] [flags]
slides [command] [flags] deck.md
```

## 命令 {#commands}

| 命令 | 作用 |
| --- | --- |
| `build` | 将演示文稿写出为 HTML，或 PDF、图片（见下文）。**默认命令**：`slides deck.md` 即 `slides build deck.md` |
| `expand` | 将演示文稿输出为纯 Marp Markdown：模板已展开，kind（幻灯片类型）已替换为其所代表的指令 |
| `notes` | 输出演讲者备注 |
| `kinds` | 列出演示文稿可用的 kind 及其设置的内容；`--json`（v0.6.0 起）将其写为 `[{"name", "extends", "directives": [{"key", "value"}]}]`，其中 `extends` 已解析 |
| `themes` | 列出主题 |
| `version` | 输出版本号（`slides --version` 亦可） |
| `completion` | 输出 shell 自动补全脚本：`bash`、`zsh`、`fish` 或 `powershell` |

`slides <command> --help` 会列出某个命令的参数。参数可以写在演示文稿之前或之后：
`slides deck.md -o out.html` 与 `slides -o out.html
deck.md` 等价。长参数使用两个短横线（`--html`）；`-o` 是唯一的短参数。
除非 `-o` 另行指定，`build` 会在 `deck.md` 旁边写出 `deck.html`。

## 参数 {#flags}

| 参数 | 适用命令 | 作用 |
| --- | --- | --- |
| `-o`, `--output file` | `build`、`expand`、`notes` | 输出文件；`-` 表示标准输出 |
| `-I`, `--input-dir dir` | `build` | 递归转换 `dir` 下的 Markdown 文件（v0.6.0 起）；见下文 |
| `-P`, `--parallel n` | `build` | 同时构建的演示文稿数（v0.6.0 起）：默认 5，最多为 CPU 核数；为 1 时逐个构建 |
| `-w`, `--watch` | `build` | 当演示文稿、主题、kind 或数据文件、配置发生变化时重新构建，并刷新已打开的页面（v0.6.0 起）；见下文 |
| `-s`, `--server dir` | `build` | 托管一个目录，按请求转换其中的演示文稿（v0.6.0 起）；见下文 |
| `-p`, `--preview` | `build` | 在默认浏览器中打开结果（v0.6.0 起）：配合 `--watch`（隐含启用）打开 HTML 页面，配合 `--server` 打开服务器 |
| `--host addr` | `build` | `--server` 监听的地址，默认 `127.0.0.1`（v0.6.0 起） |
| `--allowed-host name` | `build` | `--server` 应答的反向代理公开主机名，可重复使用（自 v0.14.0 起）；参见 [自托管]({{< relref "self-hosting" >}}) |
| `--template bespoke\|bare` | `build` | HTML 页面的模板（v0.6.0 起）：默认的 `bespoke` 带有演示脚本、屏幕控制器和过渡效果；`bare` 则将幻灯片依次堆叠（`--bare`） |
| `--bespoke.osc`, `--bespoke.progress`, `--bespoke.transition` | `build` | bespoke 模板的选项，默认值与 marp-cli 相同（v0.6.0 起）：控制器开启，**进度条关闭**（以前总是绘制），过渡效果开启；`--bespoke.osc=false` 可关闭其中一项 |
| `-c`, `--config-file file`, `--no-config-file` | `build` | marp-cli 的配置文件，JSON 或 YAML 格式（v0.6.0 起）；见下文。`--config` 是 `slides` 自己的配置文件 |
| `--browser`, `--browser-path`, `--browser-protocol`, `--browser-timeout` | `build` | 接受但忽略，并给出警告（v0.6.0 起）：这里没有浏览器 |
| `--pptx-editable`, `--no-parallel` | `build` | marp-cli 中 `--pptx` 和 `-P 1` 的写法（v0.6.0 起） |
| `--config file` | 全部 | 配置文件；默认为演示文稿旁边的 `slides.yaml`（如果存在） |
| `--no-config` | 全部 | 不读取配置文件 |
| `--theme-set path` | 全部 | 一个主题 CSS 文件，或包含主题文件的目录（可重复） |
| `--kinds file` | 全部 | 一个定义 kind 的 YAML 文件（可重复） |
| `--data file` | 全部 | 一个提供模板数据的 YAML 或 JSON 文件（可重复） |
| `--set key=value` | 全部 | 一个模板值（可重复） |
| `--no-template` | 全部 | 不展开 Go 模板 |
| `--strict-yaml` | 全部 | 按严格 YAML 解析指令，与原生 Marpit 一致（Marp Core 的解析较宽松） |
| `--no-emoji` | 全部 | 将 `:shortcodes:` 和 Unicode emoji 保留为文本，而不绘制为图片（自 v0.4.0 起：[Emoji]({{< relref "emoji" >}})） |
| `--html` | `build` | 允许所有原始 HTML；默认使用 Marp Core 的白名单 |
| `--lang code` | `build` | 演示文稿未设置时页面使用的 `lang` |
| `--bare` | `build` | 堆叠幻灯片，不带演示脚本 |
| `--pdf` | `build` | 写出 PDF，每张幻灯片一页，带书签大纲；文本和形状为矢量（也可由 `-o file.pdf` 选择） |
| `--pdf-raster` | `build` | 写出每页为一张幻灯片图片的 PDF，与 v0.3.0 相同 |
| `--notes` | `build` | 将演讲者备注写为文本文件，幻灯片之间以 `---` 分隔，与 marp-cli 相同（v0.6.0 起；也可由 `-o file.txt` 选择）。`notes` 命令以另一种格式输出备注 |
| `--pdf-notes` | `build` | 将每张幻灯片的演讲者备注作为注释批注加到对应的 PDF 页面上，与 marp-cli 相同（v0.6.0 起）。未指定其他输出格式时，会写出 PDF |
| `--pdf-outlines`, `--pdf-outlines.pages`, `--pdf-outlines.headings` | `build` | PDF 的书签（v0.6.0 起）：每张幻灯片一个 “Page N”，其下嵌套该幻灯片的标题。与 marp-cli 不同，这里默认开启；`=false` 可关闭其中一项。见下文 |
| `--images png\|jpeg` | `build` | 每张幻灯片写出一张图片：`deck.001.png`、… |
| `--image png\|jpeg` | `build` | 将第一张幻灯片写为图片（也可由 `-o file.png` 或 `-o file.jpg` 选择） |
| `--pptx` | `build` | 写出带演讲者备注的 PowerPoint 文件（也可由 `-o file.pptx` 选择）：自 v0.5.0 起可编辑，v0.4.0 中每张幻灯片是一张图片 |
| `--pptx-raster` | `build` | 写出每张幻灯片为一整张图片的 PowerPoint 文件，与 marp-cli 相同（自 v0.5.0 起）；`--image-scale` 默认为 2 |
| `--odp` | `build` | 写出带演讲者备注的 OpenDocument 演示文稿（也可由 `-o file.odp` 选择）：自 v0.5.0 起可编辑 |
| `--odp-raster` | `build` | 写出每张幻灯片为一张图片的 OpenDocument 演示文稿（自 v0.5.0 起） |
| `--jpeg-quality n` | `build` | JPEG 质量，1 到 100（默认 85） |
| `--image-scale n` | `build` | 图片和 PDF 中每个 CSS 像素对应的设备像素数（默认 1）；为 2 时 16:9 幻灯片的尺寸为 2560×1440，文本和矢量图形按该密度绘制，而不是放大 |
| `--allow-local-files` | `build` | 允许图片和 PDF 读取本地文件，且仅限演示文稿所在目录之下；默认拒绝，与 marp-cli 相同 |
| `--theme name\|file.css` | `build` | 覆盖演示文稿的主题（v0.6.0 起）：主题名，或一个 CSS 文件，无论其声明的 `@theme` 是什么，都会成为当前主题；不接受目录，目录请用 `--theme-set` |
| `--title`, `--description`, `--author`, `--keywords`, `--url`, `--og-image` | `build` | 页面的元数据（v0.6.0 起），优先于演示文稿的 `title`、`description`、`author`、`keywords`、`url` 和 `image` 指令；见下文 |

## 多个演示文稿与目录 {#several-decks-and-a-directory}

从 v0.6.0 起，`build` 可以接受多个演示文稿、一个目录（递归查找其中的演示文稿）
或一个通配模式，并同时构建最多 `--parallel` 个。`-` 表示标准输入；与 marp-cli
不同，除非用 `-` 明确指定，否则不会读取管道。

通配模式的解释方式与 marp-cli 的 glob 相同（`**` 和 `{a,b}` 自 v0.7.0 起支持；
请加引号，让 shell 把它原样交给 `slides`）：名称中可使用 `*`、`?`、`[a-z]` 和
`{md,txt}`，`**` 匹配任意层目录，因此 `'docs/**/*.md'` 会选取整棵目录树中的演示文稿。
通配符不会匹配以点开头的名称，不会在 `node_modules` 中查找，区分大小写
（`*.md` 找不到 `UPPER.MD`），并且模式匹配到的任何文件都会被转换，无论其扩展名是什么。
没有匹配到任何文件的模式会报错。

- **`-I dir`** 与 marp-cli 一样，查找 `dir` 下的 Markdown 文件（`.md`、`.mdown`、`.markdown`、
  `.markdn`，小写）：不会进入 `node_modules` 和以点开头的目录。配合 `-o out` 时，
  每个演示文稿按其在 `dir` 下的路径，镜像写到 `out` 下：`dir/sub/b.md` 生成
  `out/sub/b.html`（或 `b.pdf`、`b.001.png`、`b.txt`、…，取决于参数）。
  不带 `-o` 时，输出文件写在各自的演示文稿旁边。同时指定演示文稿和
  `-I` 会报错，对多个演示文稿使用 `-o file` 也会报错（允许 `-o -`）。
- **`-P n`** 同时构建 `n` 个演示文稿。单个演示文稿的渲染本身已经会利用多核，
  所以 `n` 的上限是 CPU 核数；使用 `-o -` 时逐个构建。同一个演示文稿的消息会集中在一起输出，
  一个演示文稿失败不会中断其他演示文稿：退出状态为 1，消息中会列出失败项。

## 过渡效果与演示页面 {#transitions-and-the-presenting-page}

从 v0.6.0 起，marp-cli 的 `transition` 指令可用：`<!-- transition:
fade 1s -->`（名称加可选的时长）作用于当前幻灯片及其后的幻灯片，
`<!-- _transition: cover -->` 则只作用于一张。内置名称与 marp-cli 相同：
`clockwise`、`counterclockwise`、`cover`、`coverflow`、`cube`、`cylinder`、
`diamond`、`drop`、`explode`、`fade`、`fade-out`、`fall`、`flip`、`glow`、
`implode`、`in-out`、`iris-in`、`iris-out`、`melt`、`overlap`、`pivot`、`pull`、
`push`、`reveal`、`rotate`、`slide`、`star`、`swap`、`swipe`、`swoosh`、`wipe`、
`wiper` 和 `zoom`；`none` 表示关闭。当观众向前离开带有过渡效果的幻灯片，
以及返回该幻灯片时，页面会使用浏览器的 **View Transition API** 播放过渡；
不支持该 API 的浏览器照旧切换幻灯片，而要求减弱动态效果的观众会看到简单的淡入淡出。
主题可以自带关键帧，命名为 `marp-transition-NAME`（或 `marp-outgoing-` 和
`marp-incoming-`，以及 `-backward`），它们优先于内置关键帧。

页面带有**屏幕控制器**（上一张、页码、下一张、全屏），在指针移动或按键时显示，
片刻后隐藏（`--bespoke.osc=false` 可将其移除）；使用 `--bespoke.progress` 时还会显示进度条。
**进度条现在默认关闭**，与 marp-cli 相同；以前总是绘制。

### 演讲者视图 {#presenter-view}

从 v0.7.0 起，页面提供 marp-cli 的演讲者视图。按 `p` 或控制器的最后一个按钮，
会打开一个窗口，显示当前幻灯片、下一步、你的演讲者备注、页码、时间和计时器
（点击计时器可重置；`+` 和 `-` 调整备注字号；幻灯片与备注之间的分隔条可以拖动）。
两个窗口双向同步：在任一窗口中按方向键、点击下一步或点击 Previous/Next，
另一个窗口都会跟着移动，每次一个片段。

它的工作方式与 marp-cli 相同（依据其 4.5.1 源码）：该窗口是同一文件的
`?view=presenter&sync=KEY`，下一步是在框架中以 `?view=next` 打开的同一页面，
窗口之间通过 `localStorage`（`bespoke-marp-sync-KEY`）共享位置，因此不需要服务器，
也不需要网络；在没有 `localStorage` 的浏览器中无法使用（按钮会给出提示）。
备注是纯文本：`&`、`<` 和 `>` 会被转义，所以备注永远不会成为标记。
备注隐藏在页面中，供演讲者视图读取，这样屏幕阅读器就不会把它们读给观众听。
`--bare` 页面不包含这些功能。该视图会使页面增加约 10 KB。

### 概览视图 {#overview-view}

从 v0.8.0 起，按 `o`（或 `Escape`：marp-cli 两者都可切换）、控制器的网格按钮，
或演讲者视图的页码计数器，所有幻灯片会以网格形式完整显示在页面上方。
方向键选择幻灯片（上下键按行移动），`Home` 和 `End` 跳到首尾，其他窗口随之同步；
`Enter`、`Space` 或点击会选中幻灯片，在主页面上还会关闭概览。
在演讲者视图上选择后概览保持打开（演讲者挑选要展示的幻灯片，并继续浏览）。
“Close overview” 按钮、`o` 和 `Escape` 可关闭概览。它与 marp-cli 的实现相同：
同一页面再次在框架中以 `?view=overview&sync=KEY` 打开，通过 `localStorage`
共享位置，无需网络。网格中的幻灯片可获得焦点，标签为 “Slide N”；
概览关闭时处于 inert 状态，因此键盘用户和屏幕阅读器用户不会误入其中。
两个视图合计使页面增加约 15 KB。

### 配置文件 {#configuration-files}

marp-cli 从工作目录中的文件读取选项，`slides` 也一样，前提是该文件是数据文件：
`.marprc`（YAML；JSON 也是 YAML）、`.marprc.json`、`.marprc.yaml`、`.marprc.yml`、
`.config/marprc*`，或 `package.json` 中的 `marp` 键；`--config-file`（`-c`）
可指定任意 JSON 或 YAML 文件，例如 `marp.config.json`。与 marp-cli 的查找方式一样，
只查找工作目录，不查找其上级目录。JavaScript 或 TypeScript 文件
（`marp.config.js` ...）是程序，`slides` 不会运行它：会给出提示并继续。
支持的选项是 marp-cli 文档中在这里有意义的那些：`title`、`description`、`author`、
`keywords`、`url`、`ogImage`、`lang`、`theme`、`themeSet`、`html`、`allowLocalFiles`、
`template`、`bespoke`、`inputDir`、`output`、`parallel`、`pdf`、`pptx`、`notes`、
`pdfNotes`、`pdfOutlines`、`images`、`image`、`imageScale`、`jpegQuality`、`watch`、
`server`、`preview`；路径相对于该文件。命令行优先。
`engine`、`options` 和 `browser*` 需要 Marp 的引擎和浏览器，会被忽略并给出警告，
`--browser*` 参数也是如此：为 marp-cli 编写的脚本可以继续使用。
（`slides.yaml` 和 `slides.hcl`，连同 kind 和模板数据，是独立的，读取方式不变。）

## 监视、托管与预览 {#watch-serve-and-preview}

自 v0.6.0 起，采用 marp-cli 的路由和实时刷新协议（依据其 4.5.1 源码，
并与运行中的 `marp -s` 对照过）。

**`slides --server dir`**（`-s`）在 `http://127.0.0.1:8080/` 上托管该目录
（环境变量 `PORT` 可修改端口；端口已被占用时会报错并说明原因）：

| 请求 | 响应 |
| --- | --- |
| `/deck.md` | 按请求转换为 HTML 的演示文稿 |
| `/deck.md?pdf`、`?pptx`、`?png`、`?jpg`（或 `?jpeg`）、`?notes`（或 `?txt`） | 按查询参数转换后的演示文稿；PPTX 以附件形式返回 |
| 目录 | 其中的 `index.md` 或 `PITCHME.md` 作为页面返回，否则返回一个索引，为每个演示文稿提供 PDF 和 PPTX 链接 |
| 其他文件 | 文件本身 |
| 无法转换的演示文稿 | 503 及错误信息 |

页面会打开一个连接到 `/.__marp-cli-watch-notifier__/<id>` 的 WebSocket，其中 id 是
演示文稿路径的 SHA-256；服务器先发送 `ready`，在演示文稿或目录下任何内容变化时发送 `reload`。
请求目录时如果缺少结尾的斜杠，会重定向到带斜杠的地址（marp-cli 会在不带斜杠的地址上返回索引，
这会破坏页面的相对链接）。

**`--watch`**（`-w`）先构建演示文稿，之后重新构建发生变化的那一个；当 `--theme-set`
文件、kind 或数据文件、配置发生变化时，则全部重新构建。它写出的页面会在回环地址上打开一个套接字
（端口 37717，或下一个空闲端口），并在对应演示文稿重新构建时刷新。构建失败的演示文稿会报告错误，
但仍继续监视。**`--preview`**（`-p`）在默认浏览器中打开服务器，或刚刚构建的 HTML 页面；
它没有自己的窗口，因为 `slides` 没有可以驱动的浏览器。

服务器是进入本机的一个入口，所以它比 marp-cli 的更严格（marp-cli 的服务器监听所有网络接口，
且不做任何检查）：

- 它监听 **127.0.0.1**；`--host` 可以更改，但会给出警告，而且通配地址将无法再检查 `Host` 请求头；
- **`Host`** 请求头必须是回环名称（`localhost`、`127.0.0.1`、
  `[::1]`、`*.localhost`）或其监听的主机名，以防范 DNS 重绑定；
- WebSocket 的 **`Origin`** 必须是服务器自身；监视模式的套接字还接受从文件打开的页面的
  `null` 来源；来自其他站点、且不是页面导航的请求（脚本、图片、`fetch`）会被拒绝；
- 路径会被规范化，符号链接会被解析，结果必须位于该目录内：`..` 和指向目录外的链接会被拒绝，
  对目录外演示文稿的 `?pdf` 请求也是如此；以点开头的名称（`.git`、`.env`）不会被提供；
- 只接受 `GET` 和 `HEAD`；同时进行的转换数量有限；最多监视 4096 个目录，
  且不监视 `node_modules` 中或以点开头的目录。

在反向代理之后，`--allowed-host name`（自 v0.14.0 起，可重复使用）会把代理的公开名称
加入可接受的 `Host` 和 `Origin`；检查保持开启，服务器仍停留在回环地址上。
[自托管]({{< relref "self-hosting" >}})把它放在 Caddy 之后，并配有 HTTPS 和登录。

## 页面元数据 {#page-metadata}

与 marp-cli 一样，命令行优先于演示文稿；在 marp-cli 中，这些参数会作为最后一条指令注释追加到演示文稿末尾。
以下规则依据 marp-cli 4.5.1 整理，并由针对其输出的测试固定下来（v0.6.0 起）：

- 未设置 `title` 的演示文稿，取其最低层级的第一个标题的原文
  （`# **Bold** title` 得到 `**Bold** title`）。没有标题时就没有 `<title>`。
- `keywords` 按逗号拆分、去除首尾空白并去重；front matter 中的列表按原样使用。
  与 marp-cli 一样，只有从第二个关键词起，页面才会带 `keywords` 标签，只有一个关键词时不带。
- 无效的 `url`（非绝对地址）会被忽略并给出警告，保留之前的值；
  有效值还会写出 `<link rel="canonical">` 和 `og:url`。
- `author` 和 `description` 会附带对应的 Open Graph 标签；`og:image:alt`
  和 Twitter 大图卡片需要同时有标题和图片。

## 图片、PDF、PPTX 和 ODP {#images-pdf-pptx-and-odp}

与 marp-cli 一样，`-o` 配合 `.pdf`、`.pptx`、`.odp`、`.png` 或 `.jpg` 扩展名即可选择格式。
输出目录会自动创建。PDF 每张幻灯片一页，页面尺寸与幻灯片相同，并带有由幻灯片及其标题组成的书签大纲。

### PDF 备注与书签大纲 {#pdf-notes-and-outline}

从 v0.6.0 起，PDF 支持 marp-cli 的两个选项，并与 marp-cli 4.5.1 为同一演示文稿写出的 PDF 做过对比：

- `--pdf-notes` 将每张幻灯片的演讲者备注作为注释批注放在对应页面上：
  左下角一个黄色的 Note 图标，内容为备注（多条注释之间以空行连接），标题为演示文稿的
  `author`。与 marp-cli 一样，单独使用 `--pdf-notes` 时会写出 PDF。
- 书签大纲为每张幻灯片包含一个粗体书签 “Page N”，其下按层级嵌套该幻灯片的标题
  （`--pdf-outlines.pages` 和 `--pdf-outlines.headings` 分别控制这两部分）。
  **书签大纲在这里默认开启，在 marp-cli 中默认关闭。** `slides` 一直都会写出书签大纲；
  它只占几百字节，却能让长篇演示文稿便于导航，所以保留这一默认值。
  `--pdf-outlines=false` 或 marp-cli 的 `--no-pdf-outlines` 可将其关闭，
  在 `.pages` 或 `.headings` 上使用 `=false` 则去掉对应部分。标题书签跳转到其所在的幻灯片；
  marp-cli 还会滚动到该标题。

渲染由 [go-webengine](https://github.com/go-webengine/engine) 完成，
PDF 由 [go-pdfkit](https://github.com/go-pdfkit/pdfkit) 写出；
内置主题所需的字体（Lato、Roboto Mono）来自
[go-opentype/fonts](https://github.com/go-opentype/fonts) 并被内嵌，
因此任何渲染都不会访问网络。这里没有 Chrome。

`--image-scale 2` 以每个 CSS 像素两个设备像素进行渲染。除非指定 `--allow-local-files`，
否则会拒绝本地图片；即使指定了，也只读取演示文稿所在目录之下的文件：
来自别处的演示文稿无法访问 `~/.ssh`。多重背景和分栏背景（`![bg left:40%]`）
的渲染方式与 Marp 的图层相同。

`--pptx` 使用 [go-ooxml](https://github.com/go-ooxml) 的
`WritePresentation` 写出演示文稿，包含演讲者备注和演示文稿的元数据。自
v0.5.0 起幻灯片可编辑（见下文）；`--pptx-raster` 则与 marp-cli 一样，每张幻灯片写出一整张图片，
缩放比例默认为 2，可由 `--image-scale` 覆盖。生成的文件可以用 LibreOffice 打开。
只在 LibreOffice 中测试过，未在 PowerPoint 和 Keynote 中测试。

`--odp` 通过 go-odf 的 `WritePresentation` 写出 OpenDocument 演示文稿，
包含演讲者备注；`--odp-raster` 每张幻灯片写出一张图片。

### 可编辑的 PPTX 和 ODP {#editable-pptx-and-odp}

自 v0.5.0 起，PPTX 或 ODP 中的幻灯片是**可编辑的**：marp-cli 的 PPTX 每张幻灯片是一张图片，
其实验性的 `--pptx-editable` 则是借助 LibreOffice 转换 PDF。这里每张幻灯片都依据其布局绘制，
与 PNG 和 PDF 导出所用的布局相同：

- 每行文本都是一个位于其布局位置的文本框，使用幻灯片排版时所用的字体，
  且**字体内嵌在文件中**，因此在没有 Lato 或 Inter 的机器上看起来也一样
  （已用 LibreOffice 25.8 验证），并保留粗体、斜体、颜色、字间距和可点击的链接；
- 背景、边框和圆角框是形状；渐变在 PPTX 中是渐变填充，在 ODP 中是图片
  （ODF 的渐变只有两种颜色）；
- 图片只存储一次，无论有多少张幻灯片显示它；
- 幻灯片无法表达的内容（受 `opacity` 影响的元素、`filter`、
  模糊阴影）会成为该元素的一张图片，其上保留文本；
- 自 v0.15.0 起，公式、emoji 或 SVG 图像是一组矢量形状（PPTX 中为 `p:grpSp`，ODP 中为 `draw:g`），作为一个整体移动。
  需要裁剪或径向渐变的 SVG 仍是图片。

经 LibreOffice 重新渲染后，幻灯片与 PNG 导出结果的像素差异在 0.04–2 % 之间。
`--pptx-raster` 和 `--odp-raster` 仍然每张幻灯片一张图片。

### 矢量 PDF {#vector-pdf}

自 v0.4.0 起，PDF 以矢量方式绘制：go-webengine 将排好版的幻灯片绘制到 PDF 中，而不是绘制成像素。
文本就是文本，使用幻灯片排版时所用的字体，以子集方式内嵌并带有 `/ToUnicode` 映射，
因此可以选中、搜索和复制。背景、边框和渐变是路径和着色，在多张幻灯片上出现的图片只内嵌一次，
链接可以点击。100 张幻灯片的基准演示文稿大小为 758 KiB，marp-cli 为 837 KiB，
v0.3.0 的图片 PDF 为 10 MiB，写出耗时 0.3 s。

这里没有矢量形式的元素（`opacity` < 1、带有 `filter` 或旋转、模糊的 `box-shadow`、
透明度渐变的渐变色）只会将该元素绘制为位图，其上覆盖不可见的文本，以保持可选中。
`--pdf-raster` 与 v0.3.0 一样，每页写出一张图片。

自 v0.15.0 起，公式、emoji 和 SVG 图像同样是矢量（go-webengine v0.22.0）：每个不同的图形是一个 PDF 表单
XObject，在使用处绘制，任意缩放都清晰。

### 无障碍的带标签 PDF {#accessible-tagged-pdf}

自 v0.12.0 起，矢量 PDF 带有标签。它包含文档的逻辑结构，屏幕阅读器会按顺序朗读，
就像朗读网页一样：

| 演示文稿中 | PDF 中 |
| --- | --- |
| 一张幻灯片 | 一个节（`Sect`） |
| `#` … `######` | `H1` … `H6` |
| 段落、列表、表格 | `P`；`L`，包含每一项的标号和正文；`Table`，表头单元格带有其作用范围 |
| 链接 | `Link`，包含其可点击区域，以其文本作为描述 |
| 图片 | `Figure`，带有替代文本；没有替代文本的 `![](…)` 视为装饰 |
| `$…$`、`$$…$$` | `Formula`，以其 TeX 源码作为 `/Alt` 和 `/ActualText`（其字形是轮廓，pdftotext 和 mupdf 不会提取，识别结构的阅读器会） |
| 代码 | `Code`，不含行号和标注圆点 |
| `<span lang="fr">` | 使用该语言的文本片段 |
| 背景、边框、页码 | 工件（artifact），阅读器会跳过 |
| 演讲者备注（`--pdf-notes`） | 结构中的批注 |

该文件声明符合 PDF/UA-1，即无障碍 PDF 的标准。它包含演示文稿的 `title`（由阅读器显示）
及其 `lang`（演示文稿未设置时取 `--lang`，否则为 `en`）。

用 veraPDF 检查的演示文稿可以通过 PDF/UA-1，除非其自身的标题违反了标准规定的层级顺序。
PDF/UA 要求第一个标题为 `H1`，且不得跳级（如 `#` 之后直接 `###`）。这只能由演示文稿的作者修正。
同一演示文稿由 Chrome 生成的 PDF（marp-cli 写出的就是它）也带有标签，但违反了 PDF/UA 的
七条规则，没有按幻灯片划分的节，也没有公式。

为 100 张幻灯片的基准演示文稿添加标签，体积增加 4%（788 KiB 对 757 KiB），
耗时没有可测量的增加。`--pdf-tags=false` 写出不带标签的 PDF。

## `slides.yaml` {#slidesyaml}

演示文稿旁边的配置文件会作用于该演示文稿，因此一个文件夹中的演示文稿可以共享同一个配置文件：

```yaml
kinds: kinds.yaml        # a mapping of kinds, or file(s) holding one
themes:                  # theme CSS files or directories
  - themes/
data:                    # template data
  company: Example Inc.
html: false              # allow all raw HTML
lang: en
```

kind：演示文稿自身的 `kinds:` 优先于配置文件和命令行中的定义。
数据：配置文件和命令行中的 `data` 优先于演示文稿的 front matter，
因此同一个演示文稿可以用不同的数据构建。
