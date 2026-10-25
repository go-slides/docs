---
title: "快速上手"
weight: 10
description: "第一个演示文稿：构建为 HTML，再展开回 Marp Markdown。"
tags: [start]
---

## 安装 {#install}

需要 Go 1.27.2 或更高版本：

```sh
go install github.com/go-slides/slides/cmd/slides@latest
```

`slides` 是一个不使用 cgo 的静态 Go 二进制文件：无需安装其他任何东西，
不需要 Node，也不需要 Chrome。本页描述的是 v0.9.0。

## 第一个演示文稿 {#a-first-deck}

任何 Marp 演示文稿都可以直接使用。将以下内容保存为 `talk.md`：

```markdown
---
marp: true
theme: default
paginate: true
footer: "{{ .talk }}"
data:
  talk: Pure Go slides
  topics: [Why, How, Next]
---

<!-- _kind: title -->
# {{ .talk }}

{{ range .topics }}
---

# {{ . }}
{{ end }}
```

front matter 与 Marp 相同，只是多了 `data:`。正文是一个 Go 模板：
`range` 为每个主题生成一张幻灯片，每张都以 `---` 开头；第一张幻灯片使用内置的 kind（幻灯片类型）
`title`（没有页眉、页脚和页码）。

## 构建 {#build-it}

```sh
slides talk.md                 # writes talk.html beside the deck
slides -o slides.html talk.md  # or somewhere else; -o - for standard output
```

不带命令的 `slides` 即 `slides build`。参数可以写在演示文稿之前或之后。
用浏览器打开生成的 HTML；用键盘即可在幻灯片之间切换。

## 导出图片、PDF、PPTX 和 ODP {#export-images-pdf-pptx-and-odp}

```sh
slides talk.md --pdf                  # one page per slide, with an outline
slides talk.md --images png           # talk.001.png, talk.002.png, …
slides talk.md --image jpeg           # the first slide only
slides talk.md -o out/slides.pdf      # the extension picks the format; out/ is created
slides talk.md --images jpeg --jpeg-quality 90
slides talk.md --images png --image-scale 2   # 2560×1440 for 16:9
slides talk.md --pdf --allow-local-files      # images from under the deck's directory
slides talk.md --pptx                         # or -o talk.pptx, with notes
slides talk.md --odp                          # or -o talk.odp: the same, as OpenDocument
```

整个过程不涉及浏览器：go-webengine 渲染幻灯片，go-pdfkit 写出 PDF。
除非指定 `--allow-local-files`，否则会拒绝本地图片。
自 v0.5.0 起，PPTX 或 ODP 中的幻灯片是可编辑的（文本框、形状和图片，并内嵌字体），
而 marp-cli 生成的是图片；`--pptx-raster`
仍生成图片（[可编辑的 PPTX 和 ODP]({{< relref "cli#editable-pptx-and-odp" >}})）。自 v0.4.0 起 PDF 为矢量格式，文本可选中
（v0.3.0 每页写出一张图片）。参见[命令行]({{< relref "cli#vector-pdf" >}})。

## 展开 {#expand-it}

```sh
slides expand talk.md
```

会将同一个演示文稿输出为**纯 Marp Markdown**（`main` 在
`4ae5c49`，2026-10-18 的输出）：

```markdown
---
marp: true
theme: default
paginate: true
footer: "Pure Go slides"
---
<!--
_class: "lead"
_footer: ""
_paginate: "false"
-->


# Pure Go slides


---

# Why

---

# How

---

# Next
```

可以看到 `range` 生成的四张幻灯片、填入了值的页脚，以及
`title` 这个 kind 已被替换为它所代表的单页指令。任何能读取 Marp 的工具（marp-cli、VS
Code 扩展）都可以接着处理它，因此用 go-slides 的扩展功能编写的内容不会被锁定。

## 下一步 {#next}

- [Marp 兼容性]({{< relref "marp" >}})：哪些内容的读取方式与 Marp 相同。
- [模板]({{< relref "templates" >}})与 [Kind]({{< relref "kinds" >}})：两项扩展功能。
- [命令行]({{< relref "cli" >}})：所有命令和参数。
