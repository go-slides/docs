---
title: "go-slides 文档"
linkTitle: "首页"
type: docs
cascade:
  type: docs
description: "纯 Go（CGO=0）实现、兼容 Marp 的幻灯片演示文稿，支持 Go 模板插值和幻灯片 kind。"
---

**`slides` 用纯 Go 把 Marp 演示文稿变成一场演示。** 它读取的 Markdown 与
[Marp](https://marp.app) 相同（Marpit 的指令、幻灯片之间的 `---`、`![bg]`
背景、演讲者备注，以及 `default`、`gaia` 和 `uncover` 主题），并增加了两项
Marp 没有的功能：

- 演示文稿中的 **[Go 模板]({{< relref "templates" >}})**：在解析 Markdown
  之前求值，因此演示文稿可以自己生成幻灯片；
- **[kind（幻灯片类型）]({{< relref "kinds" >}})**：幻灯片声明自己是什么
  （`title`、`section`…），而不必反复描述自己的外观。

{{< callout type="warning" >}}
**v0.9.0。** 本文档描述的是
[go-slides/slides v0.9.0](https://github.com/go-slides/slides/releases/tag/v0.9.0)，
发布于 `/docs/0.9/` 下。
{{< /callout >}}

## 现有功能 {#what-it-does-today}

| | |
| --- | --- |
| `slides build`（默认命令） | 将演示文稿生成为一个 HTML 页面：每张幻灯片是一个以 SVG 为框架、随窗口缩放的 `<section>`，附带主题 CSS 和键盘浏览器 |
| `slides build --pdf`、`--images png`、`--image jpeg` | 将演示文稿导出为 PDF（每张幻灯片一页，带书签大纲）或 PNG/JPEG 图片，无需浏览器：由 [go-webengine](https://github.com/go-webengine) 渲染，[go-pdfkit](https://github.com/go-pdfkit) 写出 PDF；自 v0.4.0 起为矢量格式，文本可选中：[矢量 PDF]({{< relref "cli#vector-pdf" >}}) |
| 数学公式与代码 | `$…$` 和 `$$…$$` 由 go-tex/math 排版，围栏代码由 go-rouge 高亮，适用于 HTML、图片和 PDF：[数学公式与代码]({{< relref "math-and-code" >}}) |
| emoji | `:smile:` 和 😄 会变成内嵌的 Twemoji 图片，因此无需网络，并能在所有输出格式中绘制：[Emoji]({{< relref "emoji" >}})（自 v0.4.0 起） |
| `slides build --pptx` | 将演示文稿导出为 PowerPoint 文件，包含演讲者备注和演示文稿的元数据：v0.4.0 中每张幻灯片是一张图片，与 marp-cli 相同；自 v0.5.0 起**可编辑**：文本框、形状、图片，并内嵌字体（[可编辑的 PPTX 和 ODP]({{< relref "cli#editable-pptx-and-odp" >}})） |
| `slides build --odp` | 同上，但导出为 OpenDocument 演示文稿（`.odp`） |
| 自适应标题 | `# <!-- fit --> Title` 在一行内撑满幻灯片宽度，在页面和所有导出格式中均有效 |
| `slides expand` | 将演示文稿输出为纯 Marp Markdown：模板已求值，kind 已替换为其所代表的单页指令 |
| `slides notes` | 演讲者备注 |
| `slides kinds`、`slides themes` | 演示文稿可以引用的名称 |
| `slides.yaml` | 为一个文件夹中的演示文稿共享 kind、主题、数据、HTML 策略和语言 |
| HCL | `---hcl` front matter、`slides.hcl`、HCL2 格式的 kind 和数据文件，支持表达式和函数：[HCL]({{< relref "hcl" >}}) |

## 计划中 {#planned}

尚未实现，列在这里以免大家猜测：PDF 中的公式矢量化（目前公式在 PDF 中是位图）。
与 marp-cli 的区别在于实现途径：marp-cli 通过驱动无头 Chrome 来渲染图片和 PDF；
这里则由 Go 库完成。

## 页面 {#pages}

- [快速上手]({{< relref "getting-started" >}})
- [示例库]({{< relref "gallery" >}})：十二个示例演示文稿，可在 Playground 中打开
- [Marp 兼容性]({{< relref "marp" >}})
- [模板]({{< relref "templates" >}})
- [Kind]({{< relref "kinds" >}})
- [HCL]({{< relref "hcl" >}})
- [数学公式与代码]({{< relref "math-and-code" >}})
- [Emoji]({{< relref "emoji" >}})
- [命令行]({{< relref "cli" >}})
