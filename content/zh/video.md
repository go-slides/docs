---
title: "视频"
weight: 39
description: "演示文稿中的视频：在 HTML 页面中由浏览器播放，在 PDF 中是带链接的封面帧，在 PowerPoint 和 OpenDocument 中嵌入。"
tags: [video]
---

{{< callout type="warning" >}}
视频尚未包含在任何发行版中。
{{< /callout >}}

视频是文件以 `.mp4`、`.m4v` 或 `.webm` 结尾的图片：

```markdown
![video](clip.mp4)
![The demo](clip.mp4#t=12.5)
![demo loop muted](clip.mp4)
```

## 语法 {#syntax}

文件从演示文稿所在的文件夹读取，或从 `http` 或 `https` URL 读取。绝对路径、离开该文件夹的路径或任何其他协议都会被拒绝。`#t=12.5`（或用分钟和秒的 `#t=1:23`）选择封面帧：在播放前以及无法播放视频的输出中显示的画面。也接受媒体片段形式 `#t=10,20`，并使用其起点。没有它时，封面是第一帧。

替代文字中的关键字决定视频如何播放，写法像 Marp 的图片关键字，但是 go-slides 自己的：`autoplay`、`loop` 和 `muted`。`autoplay` 仅在与 `muted` 同时出现时生效；否则会被忽略，并给出警告。替代文字的其余词语是无障碍标签。

## 各输出显示什么 {#outputs}

- **HTML**（页面和 playground 的预览）：一个 `<video controls playsinline preload="metadata" poster=...>` 元素。封面是解码后的帧，以 JPEG data URI 嵌入。视频从不自行带声音播放。当浏览器请求减少动态效果（`prefers-reduced-motion`）时，会移除 `autoplay` 和 `loop`，并暂停视频。自动播放（静音）的视频只在至少一半可见时播放，其幻灯片未显示时暂停。
- **PDF**（`--pdf`）：封面帧作为图片，并带有指向视频文件的链接：演示文稿中所写的相对路径，或 `http`、`https` 来源的 URL。只有打开链接才会播放。PDF 不含 `RichMedia` 注释：它只能在 Acrobat 中播放，因此特意不使用。
- **PowerPoint**（`--pptx`）：视频作为媒体部件（`ppt/media/slidesvideoN.mp4|webm`）连同封面嵌入，形式为 `a:videoFile` 加 `p14:media` 扩展。它可在 PowerPoint 中播放；在 PowerPoint 和 Keynote 中的播放尚未验证。`autoplay`、`loop` 和 `muted` 关键字不会带入 PowerPoint 或 OpenDocument。
- **OpenDocument**（`--odp`）：视频作为 `draw:plugin`（ODF 1.3 第 3 部分，10.4.8）嵌入 `Video/` 下，位于封面 `draw:image` 之前，并列入清单。在其媒体后端支持该编解码器的 Impress 中可播放（未验证）。
- **URL 来源**：HTML 页面从该 URL 流式播放视频。导出显示一张链接到该 URL 的中性卡片。构建时不会获取任何内容，也不会在 PowerPoint 或 OpenDocument 中嵌入任何内容。

## 容器与封面 {#codecs}

MP4（ISO BMFF）和 WebM（Matroska）文件由 Go 读取，从不运行 ffmpeg。只有 H.264（`avc1` 或 `avc3`）会解码封面，这是 go-avkit 唯一的视频解码器。其他编解码器（H.265、VP8、VP9、AV1）仍可在 HTML 中播放，并嵌入 PowerPoint 和 OpenDocument，但封面是带播放三角形的中性深色卡片，并输出警告 `video.codec`。

## 限制 {#limits}

视频是不受信任的输入，因此有界限：文件（64 MiB）、box 或元素（262,144）及其嵌套（16 层）、画面（4096 乘 4096）以及封面解码（从前一个关键帧起最多 600 个样本，每个样本 8 MiB）。离关键帧太远的封面会得到卡片和警告 `video.limit`。

被拒绝的视频在幻灯片中由一条注释 `[video not shown: ...]` 代替；构建继续，问题是一条带代码的警告：

- `video.syntax`：不是 MP4 或 WebM 文件、没有视频轨，或时间无效；
- `video.limit`：超出了上述某个界限；
- `video.codec`：无法为该编解码器解码封面；
- `video.path`：绝对路径、离开演示文稿文件夹的路径，或 `http` 和 `https` 以外的协议；
- `file.missing`：文件不存在。

## 在 playground 中 {#playground}

发布后，可以像其他文件一样把视频添加到项目中。它的字节留在浏览器中，作为 blob URL，绝不会上传。请在 [playground](https://go-slides.github.io/playground/) 中试用。
