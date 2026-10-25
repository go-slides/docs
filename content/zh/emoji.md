---
title: "Emoji"
weight: 36
description: "Emoji 短代码和 Unicode emoji 会变成 Twemoji 图片，无需网络：适用于 HTML 页面、图片、PDF、PowerPoint 和 OpenDocument。"
tags: [emoji]
---

{{< callout type="warning" >}}
Emoji 自 v0.4.0 起提供；v0.3.0 不支持。
{{< /callout >}}

emoji 可以写成短代码，也可以直接写字符本身：

```markdown
# Launch :rocket: 🎉

Shortcodes :smile: :+1: :heart_eyes: and Unicode 😄 👍🏽 👩‍💻 🇫🇷 1️⃣.
```

每个 emoji 都会变成 Marp Core 所绘制的 [Twemoji](https://github.com/jdecked/twemoji) 图片，
标记与 Marp Core 写出的相同：

```html
<img class="emoji" draggable="false" alt="😄" src="data:image/svg+xml;base64,…" data-marp-twemoji=""/>
```

## 无需网络 {#no-network}

Marp Core 的 `src` 是 CDN 上的一个 URL，因此离线展示、打印，或在没有网络连接的机器上转换为 PDF 的演示文稿，
都会显示破损的图片。go-slides 则把 SVG 本身放进 `data:` URI。这些图形（共 4,009 个，
Twemoji 17.0.3，Emoji 17.0）来自纯 Go 模块
[go-icons/twemoji](https://github.com/go-icons/twemoji)，内嵌在程序中，压缩后 2 MB，
只有在演示文稿包含 emoji 时才会读取。

HTML 页面、`--images`/`--image`、`--pdf`、`--pptx` 和 `--odp` 都同样会绘制这些图片：
由 go-webengine 对 SVG 进行栅格化。全部 4,009 个图形都能被其 SVG 栅格化器解析和绘制；
有测试对此进行检查。自 v0.5.0 起，该栅格化器为 [go-gfx/oksvg](https://github.com/go-gfx/oksvg)
v0.2.0：`clipPath`、`fill-rule: evenodd` 和圆弧命令现在的绘制效果与 Chrome 一致，
这改变了一个 Twemoji 和另外 211 个图形，SVG 徽标也不再错乱。

让 emoji 与文字等大的规则（`img[data-marp-twemoji]`，一个 `em` 见方，与基线对齐）
会像 Marp Core 那样添加到主题中，因此你自己的主题也能用上。

## 转换范围 {#what-is-converted}

规则与 Marp Core 相同；有测试将其与 Marp Core 4.4.0 为一个涵盖所有情况的演示文稿写出的 HTML 进行比较。

| | |
| --- | --- |
| `:smile:`、`:+1:`、`:heart_eyes:` | markdown-it-emoji 的 1,903 个名称，适用于正文、标题、链接、表格、`header` 和 `footer`；不适用于代码、URL 或图片的替代文本 |
| `😄`、`👩‍💻`、`🇫🇷`、`1️⃣`、`👍🏽` | Unicode emoji，适用于正文**以及代码**：Marp Core 会改写渲染后的代码 HTML |
| `✔`、`©`、`™`、`▶` | 保留为文本：默认以文本形式呈现的字符，只有带 U+FE0F（`✔️`）时才是 emoji；`❤` 是例外 |
| `☺︎` | 保留为文本：emoji 后的 U+FE0E 表示要求使用普通字形 |
| `12:30:45`、`:not_a_name:` | 保持原样 |

标题的 `id` 和书签大纲使用的是 emoji 本身，而不是短代码的名称：
`# Launch :rocket:` 的 id 为 `launch-%F0%9F%9A%80`，与 Marp 相同。

go-slides 与 Marp Core 的不同之处在于它转换得更多：`✌️ ☝️ ✍️ 🏋️`，以及其他一些
Marp Core 的 emoji-regex 保留为文本的默认文本字符，还有少数 Twemoji 能绘制但 Unicode 未列出的序列。

## 关闭 {#turning-it-off}

`--no-emoji`（在代码中为 `slides.Config{NoEmoji: true}`）会将短代码和
Unicode emoji 保留为书写时的文本，相当于 Marp Core 的
`emoji: false`。

## 署名 {#attribution}

这些图形来自 Twemoji，Copyright 2020 Twitter, Inc and other contributors，
采用 [CC-BY 4.0](https://creativecommons.org/licenses/by/4.0/) 许可，
由 [jdecked/twemoji](https://github.com/jdecked/twemoji) 项目继续维护。
展示这些图形的演示文稿应注明出处，例如在页脚或结尾幻灯片中。短代码表来自 markdown-it-emoji（MIT）。
