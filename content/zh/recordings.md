---
title: "终端录制"
weight: 38
description: "演示文稿中的 asciinema 录制：在 HTML 页面的小型播放器中回放，在 PDF、PowerPoint 和 OpenDocument 中绘制为矢量文本。"
tags: [asciinema, terminal]
---

{{< callout type="warning" >}}
终端录制位于 v0.17.1 之后的主分支，尚未包含在任何发布版本中。
{{< /callout >}}

[asciinema](https://docs.asciinema.org/) 录制（[asciicast v2](https://docs.asciinema.org/manual/asciicast/v2/) 或 [v3](https://docs.asciinema.org/manual/asciicast/v3/)）是文件以 `.cast` 结尾的图像：

```markdown
![cast](demo.cast)
![The failing build](demo.cast#t=12.5)
```

## 语法 {#syntax}

文件从演示文稿所在的文件夹读取：URL、绝对路径或离开该文件夹的路径都会被拒绝。替代文本为屏幕阅读器命名这段录制；单独的 `cast` 表示录制本身的标题。`#t=12.5`（或以分秒表示的 `#t=1:23`）选择无法动起来的输出所显示的时刻；不写时，就是录制的结尾。

## 各输出显示什么 {#outputs}

- **HTML**（页面和 playground 的预览）：录制在嵌入页面的 go-slides 小型播放器中回放：播放与暂停、速度按钮和位置条。它不会从任何地方加载脚本，也不使用网络。
- **PDF、PowerPoint 和 OpenDocument**（`--pdf`、`--pptx`、`--odp`）：`#t=` 所指时刻的终端，以主题等宽字体的矢量文本绘制，含 16 色、256 色和直接色、粗体、斜体、下划线和反色文本。文本可以选择和搜索，绝不是图片。主题的等宽字体有粗体和斜体时才会绘制它们（gaia 主题的 Roboto Mono）。

## 播放器 {#player}

录制第一次出现在屏幕上时，它播放一次。当浏览器要求减少动态效果（`prefers-reduced-motion`）时，它以暂停状态开始，显示 `#t=` 所指的时刻，按下播放后才播放。它的按钮和按键不会翻动幻灯片。

## 限制 {#limits}

录制是不受信任的输入，因此有上限：文件（16 MiB）、事件（500,000 个）、时长（3 小时）、终端（400 × 200 个单元格）以及模拟器可做的工作量。超出上限或不是格式正确的 asciicast 文件的录制会被拒绝：构建继续，幻灯片在原处显示一条说明，问题则成为带代码的警告：`cast.syntax`、`cast.limit`、`cast.version` 或 `cast.path`。

## 在 playground 中 {#playground}

从“文件”面板把 `.cast` 文件添加到项目，可拖放或使用文件选择器。它会进入 `media/` 并作为图像插入。它是演示文稿读取的文本，因此与项目的其他文本文件一同计算，上限为 2 MiB。请在 [playground](https://go-slides.github.io/playground/) 中试一试。
