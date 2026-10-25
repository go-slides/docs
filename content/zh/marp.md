---
title: "Marp 兼容性"
linkTitle: "Marp"
weight: 20
description: "slides 读取的 Markdown 就是 Marpit 和 Marp Core 的 Markdown，而非其方言。"
tags: [marp]
---

为 [Marp](https://marp.app) 编写的演示文稿就是 go-slides 演示文稿。下面的规则来自
[Marpit](https://marpit.marp.app) 和
[Marp Core](https://github.com/marp-team/marp-core)；凡是 go-slides
有所不同的地方，本页都会注明。

## 幻灯片 {#slides}

水平分隔线（单独一行的 `---`，与 Marpit 的分隔方式相同）开始一张新幻灯片，
`headingDivider` 指令所指定级别的标题也会开始新幻灯片。front matter 位于最前面的两个
`---` 之间，包含全局指令和演示文稿的数据。

## 指令 {#directives}

| 类型 | 位置 | 作用范围 |
| --- | --- | --- |
| 全局 | front matter，或任意位置的 `<!-- theme: gaia -->` | 整个演示文稿：`theme`、`style`、`headingDivider`、`lang`… |
| 局部 | `<!-- paginate: true -->` | 当前幻灯片及其后的幻灯片 |
| 单页 | `<!-- _class: lead -->`（下划线开头） | 仅当前幻灯片 |

局部指令与 Marpit 相同：`paginate`、`header`、`footer`、`class`、
`backgroundColor`、`backgroundImage`、`backgroundPosition`、
`backgroundRepeat`、`backgroundSize`、`color`。与 Marp Core 一样，
指令按宽松 YAML 解析；`--strict-yaml` 则与原生 Marpit 一样严格解析。

## 图片 {#images}

`![bg](image.jpg)` 将图片设为幻灯片背景，支持 Marp 的关键字：
`bg contain`、`bg cover`、`bg 50%`、用于分栏背景的 `bg left`/`bg right`、
多张并排的 `bg` 图片，以及图片滤镜（`blur`、`sepia`、`grayscale`…）。
行内图片尺寸（`w:200px`、`h:`）的解释方式与 Marp 相同。

## 幻灯片周边 {#around-the-slide}

- **页眉和页脚**：`header:` 和 `footer:` 指令，内容为 Markdown。
- **页码**：`paginate: true`。
- **演讲者备注**：不是指令的 HTML 注释就是备注。
  `slides notes` 可输出备注。
- **逐项显示列表**：用 `*`（或 `1)`）编写的列表会逐项出现，
  与 Marp 相同。
- **`<style scoped>`**：只作用于一张幻灯片的 CSS。
- **原始 HTML**：默认使用 Marp Core 的白名单；`--html` 允许所有 HTML。
- **Emoji**：`:shortcodes:` 和 Unicode emoji 会变成 Twemoji 图片，与
  Marp Core 相同，但图片是内嵌的，而不是从 CDN 加载：参见
  [Emoji]({{< relref "emoji" >}})。

## 主题 {#themes}

内置 `default`、`gaia` 和 `uncover` 主题，通过 `theme:` 选择。
自定义主题是一个带有 Marp 的 `/* @theme name */` 注释的 CSS 文件，
通过 `--theme-set` 指定，或列在 `slides.yaml` 的 `themes:` 下。

## go-slides 的不同之处 {#where-go-slides-differs}

- **模板。** 正文中的 `{{ }}` 是 Go 模板动作（参见
  [模板]({{< relref "templates" >}})）。如果演示文稿要在代码块之外把 `{{` 作为文本显示，
  需要使用 `--no-template`，或按 Go 模板的方式转义该文本。代码块和行内代码不受影响。
- **Kind。** `kind:` 和 `_kind:` 是 Marp 没有的指令（参见
  [Kind]({{< relref "kinds" >}})）。`slides expand` 会将它们替换为其所代表的指令，
  因此展开后的演示文稿是纯 Marp。
- **输出格式。** HTML、PNG、JPEG、PDF、PPTX 和 ODP，由 Go 库生成，而不是无头浏览器
  （已知限制参见[命令行]({{< relref "cli#images-pdf-pptx-and-odp" >}})）。
  在 v0.4.0 中，PPTX 与 marp-cli 一样由图片组成；自
  v0.5.0 起，其幻灯片由可编辑的文本框和形状组成，`--pptx-raster` 则保留图片形式。
- **自适应标题。** 支持 `# <!-- fit --> Title`：标题在一行内撑满幻灯片宽度，
  在页面和所有导出格式中均有效。
- **数学公式与代码。** `$…$` 和 `$$…$$` 遵循 Marp Core 的定界符规则，
  由 go-tex/math 排版；围栏代码由 go-rouge 高亮，使用
  highlight.js 的类名（参见[数学公式与代码]({{< relref "math-and-code" >}})）。
