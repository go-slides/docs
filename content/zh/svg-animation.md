---
title: "动画 SVG"
weight: 40
description: "带有 SMIL 或 CSS 动画的 SVG 图片：在 HTML 中由浏览器播放，在 PowerPoint 和 OpenDocument 中是带动画的形状，原生无法表达时用动画 GIF，在 PDF 和图片中显示初始状态。"
tags: [svg, animation]
---

会自己动起来的 SVG 图片，写法和任何图片一样：

```markdown
![](spinner.svg)
![w:400](chart.svg)
```

## 读取什么 {#read}

写在 SVG 文件本身中的动画：SMIL（`animate`、`animateTransform`、`animateMotion`、`animateColor`、`set`，以及 `begin`、`dur`、`repeatCount`、`fill`、`calcMode`、`keyTimes` 和 `keySplines`）和 CSS（带有 `animation` 属性的 `@keyframes`、`transform-origin` 和 `transform-box`）。不能理解的内容，例如 `begin="click"` 这样的事件，会被略去并以警告报告：不会有任何内容被悄悄丢弃。

## 各种输出显示什么 {#outputs}

- **HTML**：由浏览器自行播放动画。
- **PowerPoint**（`--pptx`）和 **OpenDocument**（`--odp`）：每个有动画的元素都是独立的形状，其动画随幻灯片一起开始：移动、旋转和缩放、不透明度、两种颜色之间的变化。元素的运动按动画自身的时刻采样，因此缓动、多个关键帧以及绕另一点的旋转都是精确的。重复（包括 `indefinite`）、延迟和反向都会保留。
- **逐帧**：当动画无法这样表达时（会变形的路径、自己绘制的描边、倾斜），整张图片按动画自身的时刻渲染，并作为动画 GIF 嵌入，PowerPoint 和 Impress 都能播放。这是有界的：每秒 12 帧，最多 240 帧，所有帧合计 1600 万像素。
- **PDF**（`--pdf`）和图片（`--images`）：它们不能动画。图形以初始状态显示，不含 JavaScript，也不含 `RichMedia`，并用一条警告列出这些动画图片。

每张动画图片会得到一行代码为 `svg-animation` 的警告，说明它是如何处理的：原生，还是逐帧以及原因。

## 各输出的支持情况 {#support}

| 特性 | HTML | PowerPoint | OpenDocument | PDF 和图片 |
|---|---|---|---|---|
| 平移、沿路径移动 | 播放 | 原生 | 原生 | 初始状态 |
| 绕任意点的旋转和缩放 | 播放 | 原生 | 原生 | 初始状态 |
| 圆的半径、矩形的大小（无描边） | 播放 | 原生 | 原生 | 初始状态 |
| 单个形状元素的不透明度 | 播放 | 原生 | 原生 | 初始状态 |
| 填充或描边的颜色，从一种颜色到另一种 | 播放 | 原生 | 原生 | 初始状态 |
| 重复、延迟、反向、结束时保持 | 播放 | 原生 | 原生 | 初始状态 |
| 缓动、多个关键帧 | 播放 | 原生（采样） | 原生（采样） | 初始状态 |
| 阶梯、离散变化 | 播放 | 逐帧 | 逐帧 | 初始状态 |
| 路径变形、虚线偏移、位置和大小属性、渐变 | 播放 | 逐帧 | 逐帧 | 初始状态 |
| 倾斜、翻转、一个元素上的多个时钟、嵌套的动画元素 | 播放 | 逐帧 | 逐帧 | 初始状态 |
| 事件 | 播放 | 不保留（警告） | 不保留（警告） | 初始状态 |

## 本地文件 {#local-files}

本地图片只有在使用 `--allow-local-files` 时才会读取，且只读取演示文稿所在文件夹中的文件，与 marp-cli 相同。没有该选项时，图片不会出现在导出结果中，构建会用 `local-file-blocked` 说明，并指出文件、所在幻灯片和该选项。不存在的文件会报告为 `local-file-missing`。

## 限制 {#limits}

SVG 是不受信任的输入，因此有边界：8 MiB 的 SVG、50,000 个元素、2,000 个动画、单个动画 256 个关键帧、60 秒时间线和 240 帧。超出边界的图形会被拒绝并给出警告，构建继续进行。

## 已验证的内容 {#checked}

LibreOffice Impress（macOS）能打开 PowerPoint 和 OpenDocument 文件，并在以两种格式重新保存时保留每个动画。Keynote（macOS）导入形状和动画 GIF；其动画未经检查。PowerPoint：尚未验证。
