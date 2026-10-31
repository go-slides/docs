---
title: "VS Code"
weight: 55
description: "VS Code 的 go-slides 扩展：预览、演示、问题提示、自动补全和导出，就像 Marp for VS Code 之于 Marp。"
tags: [vscode, editor]
---

{{< callout type="info" >}}
该扩展是 [go-slides/vscode](https://github.com/go-slides/vscode)，
独立进行版本管理（v0.5.0）。它使用 go-slides v0.14.0 渲染（包括
`editor` 包、滚动同步所用的源码行，以及概览视图和演讲者视图）。
请从其 GitHub release 附带的 `go-slides-0.5.0.vsix` 安装：
在 VS Code Marketplace 和 Open VSX 上的发布已在准备中，但目前尚未上架。
{{< /callout >}}

该扩展是 go-slides 版的
[Marp for VS Code](https://github.com/marp-team/marp-vscode)：演示文稿会在
VS Code 的 Markdown 预览中以幻灯片形式显示，错误会显示在“问题”面板中，
指令、kind（幻灯片类型）和模板函数会在输入时自动补全。

## 安装 {#install}

从 [releases](https://github.com/go-slides/vscode/releases) 页面下载
`go-slides-<version>.vsix`，然后运行
*Extensions: Install from VSIX...*，或者：

```sh
code --install-extension go-slides-0.5.0.vsix
```

预览不需要任何其他依赖。导出则需要 `slides` 命令：

```sh
go install github.com/go-slides/slides/cmd/slides@latest
```

## 预览 {#preview}

打开一个演示文稿，然后执行 *Markdown: Open Preview to the Side*。演示文稿就是 front matter 中写有
`slides: true` 的 Markdown 文件：

```markdown
---
slides: true
paginate: true
data:
  who: World
---

# Hello {{ .who }}

---

<!-- _kind: title -->
# Page {{ page }} of {{ total }}
```

带有 `marp: true` 的文件也会被预览，除非已安装 Marp for VS Code：
因为那样两个扩展会渲染同一个文档。由设置项
`goSlides.preview.marpDocuments` 决定：`auto`（默认）、`always` 或
`never`。

预览就是 go-slides 本身，编译为 WebAssembly 并由扩展加载：不需要二进制文件，
在 [vscode.dev](https://vscode.dev) 上也能使用。它从工作区中读取演示文稿用到的内容
（演示文稿旁边的 `slides.yaml` 或 `slides.hcl`、其中引用的 [kind]({{< relref "kinds" >}}) 文件、
[模板]({{< relref "templates" >}})引入的文件、主题 CSS），且只从演示文稿所在的工作区文件夹中读取。

## 滚动同步 {#scroll-sync}

从扩展 v0.2.0（go-slides v0.7.0）起，编辑器与预览会双向同步滚动，
与普通 Markdown 一样：预览跟随编辑器光标所在的幻灯片和块，
滚动预览也会移动编辑器。相关设置是 VS Code 自带的
`markdown.preview.scrollPreviewWithEditor` 和
`markdown.preview.scrollEditorWithPreview`。

行号映射会穿过模板。由 `{{ range }}` 生成的幻灯片跟随其模板所在的行，
由 `{{ include }}` 引入的内容跟随 include 所在的行。用 `headingDivider` 拆分的演示文稿
按其标题同步。

## 演示 {#presenting}

从扩展 v0.4.0（go-slides v0.9.0）起，无需离开 VS Code 即可进行演示。
在编辑器中打开演示文稿后，从命令面板或编辑器的 `…` 菜单中选择：

| 命令 | 打开的内容 |
|---|---|
| *go-slides: Open Overview* | 一个以网格显示所有幻灯片的面板。选中某张幻灯片后，编辑器会显示其源码；`Esc` 关闭面板 |
| *go-slides: Start Presenter View* | 幻灯片，以及旁边的演讲者视图：下一步、演讲者备注和计时器 |

这些面板位于编辑器旁边，并相互同步：在其中一个面板中切换幻灯片，其他面板也会随之切换。
它们也会跟随你对演示文稿的编辑。在幻灯片中按 `p`
键会打开演讲者视图。

要在另一块屏幕上演示，请用 *View: Move
Editor into New Window* 把幻灯片面板移过去。

每个面板都在一个沙盒框架（不带 `allow-same-origin`）中运行演示文稿，并采用
`default-src 'none'` 策略：框架只加载图片、媒体、样式和字体，
且只运行 go-slides 内置于页面、按哈希值指定的脚本。演示文稿自己写的脚本不会运行。

## 问题提示、自动补全与悬停提示 {#problems-completion-hover}

- **问题**：模板、YAML 和 [HCL]({{< relref "hcl" >}}) 错误，显示在出错的行上，
  位于演示文稿中或其来源的 kind 文件中。
- **自动补全**：front matter 中和 `<!-- -->` 注释中的指令键
  （单页指令带 `_`），`kind` 和 `_kind` 的取值（附带每个 kind 设置的内容），
  `theme`、`size`、`paginate`、`math`、`headingDivider`；在
  `{{ }}` 内则补全模板函数和数据字段。
- **悬停提示**：指令的作用、kind 设置的内容、模板函数的参数。自 v0.16.0 起，
  指令帮助会跟随编辑器的界面语言（VS Code 的显示语言），支持 11 种语言。

## 导出 {#export}

*go-slides: Export Slide Deck...*（也位于编辑器标题栏中）会调用 `slides` 命令写出 HTML、
矢量 PDF、可编辑的 PPTX 和 ODP、PNG 或 JPEG，效果与
`slides -o deck.pdf deck.md` 相同；参见[命令行]({{< relref "cli" >}})。
导出需要受信任的工作区。如果找不到 `slides`，扩展会提议在终端中运行上面的
`go install`。

从扩展 v0.5.0（go-slides v0.14.0）起，导出的 PDF 具备**无障碍**能力：
使用 `slides` v0.12.0 或更高版本时，矢量 PDF 带有供屏幕阅读器使用的标签，
声明 PDF/UA-1，包含演示文稿的 `title` 和 `lang`，并为每张幻灯片提供一个章节。
`slides --pdf-tags=false` 会写出不带标签的文件。请用上面的 `go install` 更新该命令。

## 设置 {#settings}

| 设置项 | 默认值 | |
|---|---|---|
| `goSlides.cliPath` | `""` | `slides` 命令的路径；为空时使用 `PATH` 中的 `slides`。这是机器级设置：工作区无法设置它 |
| `goSlides.exportType` | `pdf` | 导出对话框默认推荐的格式 |
| `goSlides.exportAutoOpen` | `true` | 导出后打开文件 |
| `goSlides.html` | `false` | 允许所有原始 HTML，而不使用 Marp 的白名单；仅在受信任的工作区中生效 |
| `goSlides.preview.marpDocuments` | `auto` | 是否预览 `marp: true` 演示文稿：`auto`、`always`、`never` |

## 安全 {#security}

演示文稿和工作区都被视为不可信。原始 HTML 遵循 Marp 的白名单，
除非在受信任的工作区中开启了 `goSlides.html`（在 `slides.yaml` 中设置
`html: true` 不会在预览中开启它），而且 VS Code 的预览不会运行演示文稿中的任何脚本。
预览只读取演示文稿所在工作区文件夹中的文件，不跟随符号链接，模板也看不到任何环境变量。
导出会运行一个程序，因此要等工作区被信任后才会进行，而且该程序永远不会由工作区指定。

## 性能 {#performance}

在 go-slides 的 100 张幻灯片基准演示文稿上，使用 Apple M4 Max 上的 VS Code 1.142，
并在同一会话中测量 Marp for VS Code 3.6.1：

| | go-slides | Marp for VS Code |
|---|---:|---:|
| 激活（加载 WebAssembly 模块） | 189 ms | — |
| 首次预览 | 52 ms | 78 ms |
| 编辑后的预览（中位数） | **19 ms** | 23–24 ms |
| `.vsix` | 7.5 MB | 36.3 MB |

以上数据针对扩展 v0.3.0（go-slides v0.8.0），使用 100 张幻灯片的演示文稿，
在第 50 张幻灯片中输入一个字符，在 1 分钟负载低于 8 的条件下交错多轮测得。
首次预览比 Marp 快，现在编辑后的预览也比 Marp 快。

### v0.8.0 与扩展 v0.3.0 {#in-v080-and-extension-v030}

go-slides v0.8.0 会保留每个已打开演示文稿的渲染结果，编辑后只重新渲染
源码或继承状态（指令、页码、总页数、标题 id、链接引用、源码行）发生变化的幻灯片；
输出与完整渲染逐字节相同。在 100 张幻灯片中的第 50 张中输入，只会重新渲染一张幻灯片：
原生 2.5 ms，WebAssembly 中 11.5 ms。随后 VS Code 的预览只对变化幻灯片的 DOM 做差量更新；
扩展的冒烟测试会对两个扩展都在两次计时渲染之间向第 50 张幻灯片输入一个字符。
扩展的演讲者视图通过父窗口（`postMessage`）同步。

### v0.7.1 与扩展 v0.2.2 {#in-v071-and-extension-v022}

go-slides v0.7.1 会在多次渲染之间缓存已解析的主题、打包好的样式表和已高亮的代码，
标记滚动同步行时不再逐字节处理，并复用其写入器：渲染 100 张幻灯片的演示文稿需分配 5.7 MB 内存，
扩展 v0.2.0 时为 52.7 MB。借此，扩展 v0.2.2 去掉了 `GOGC=400` 覆盖设置。
作为对比，扩展 v0.2.1 在默认 `GOGC` 下首次预览耗时 72 ms，编辑后 44 ms
（使用 `GOGC=400` 时为 36 ms），而 v0.2.0 为
47–51 ms。

扩展 v0.2.2 测得首次预览 52 ms、编辑后 24 ms，
与 Marp 的 23 ms 持平。
