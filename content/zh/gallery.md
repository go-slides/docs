---
title: "示例库"
weight: 12
description: "十二个小型演示文稿，各自展示一项功能：主题、kind、模板、HCL、数学公式、图片、代码、emoji、过渡效果，以及一场完整的演讲。每一个都可以在 Playground 中打开。"
tags: [gallery, examples]
---

十二个小型演示文稿，来自 go-slides 的
[`examples/`](https://github.com/go-slides/slides/tree/main/examples) 文件夹，
每个都着重展示一项功能。缩略图是每个演示文稿的第一张幻灯片，在构建本站点时由 go-slides 自己渲染；
CI 会将每个演示文稿构建为 HTML 和 PDF，确保它们都不会损坏。

**在 Playground 中打开** 会在你的浏览器中加载演示文稿，所用的链接不会向任何服务器发送数据。
除 `deck.md` 外还有其他文件（kind 文件、引入的片段、图片）的演示文稿无法通过链接传递：
请下载它的 **.zip**，再通过 Playground 的文件面板导入；或者解压后用
`slides build --allow-local-files deck.md` 构建。

图片是 go-slides 作者原创的 SVG，以
[CC0 1.0](https://creativecommons.org/publicdomain/zero/1.0/) 发布，
使用这些图片的演示文稿中也有说明。

{{< gallery >}}
