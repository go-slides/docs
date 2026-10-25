---
title: "性能"
weight: 60
description: "go-slides 构建演示文稿有多快、占用多少内存，在同一台机器上与 marp-cli 对比测量。"
tags: [performance]
---

{{< callout type="info" >}}
下文的 PPTX 数据针对 v0.5.0，矢量 PDF 数据针对 v0.4.0。
第一张表针对 v0.3.0，
即当时的上一个版本，该版本并行渲染幻灯片，并且只解析一次字体
（[go-slides/slides#11](https://github.com/go-slides/slides/pull/11)）。
{{< /callout >}}

go-slides 和 [marp-cli](https://github.com/marp-team/marp-cli) 4.5.1
以子进程方式在相同的演示文稿上运行，机器为 Apple M4 Max（16 个 CPU），
1 分钟平均负载低于 8，预热后交错进行五轮测试。
时间为实际耗时（wall time）的中位数；内存为整个进程树的峰值
（marp-cli 会驱动一个 Chrome，其进程也计算在内）。

| 演示文稿 | 输出 | go-slides | marp-cli | go-slides 内存 | marp-cli 内存 |
| --- | --- | ---: | ---: | ---: | ---: |
| 100 张幻灯片 | HTML | 36 ms | 315 ms | 47 MiB | 140 MiB |
| 100 张幻灯片 | PDF | 1.14 s | 1.88 s | 242 MiB | 1,509 MiB |
| 100 张幻灯片 | PNG，全部幻灯片 | 0.97 s | 5.44 s | 225 MiB | 1,354 MiB |
| 100 张幻灯片 | PPTX | 3.39 s | 8.03 s | 530 MiB | 1,373 MiB |
| gaia，20 张幻灯片 | PDF | 283 ms | 1.57 s | 154 MiB | 1,315 MiB |
| 数学公式与代码，20 张幻灯片 | PNG，全部幻灯片 | 250 ms | 2.32 s | 168 MiB | 1,286 MiB |

## 矢量 PDF {#vector-pdf}

从 v0.4.0 起，PDF 以矢量方式绘制（参见
[命令行]({{< relref "cli#vector-pdf" >}})）。机器和方法相同，
测于 2026-10-19：

| 演示文稿 | go-slides | marp-cli | v0.3.0 | go-slides 大小 | marp-cli 大小 | v0.3.0 大小 |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| 100 张幻灯片 | 306 ms | 1,763 ms | 1,147 ms | 758 KiB | 837 KiB | 10.0 MiB |
| gaia，20 张幻灯片 | 93 ms | 1,546 ms | 285 ms | 32 KiB | 66 KiB | 582 KiB |
| 数学公式与代码，20 张幻灯片 | 140 ms | 1,643 ms | 294 ms | 124 KiB | 335 KiB | 638 KiB |

100 张幻灯片 PDF 的内存占用：158 MiB，marp-cli 为 1,512 MiB；CPU
时间 0.9 s，marp-cli 为 2.7 s。

v0.4.0 的 100 张幻灯片 PDF 数据测于 2026-10-19 21:29，
1 分钟平均负载为 7.7，共三轮：go-slides 306 ms、758
KiB、158 MiB；marp-cli 1,838 ms、837 KiB、1,512 MiB。

## 可编辑 PPTX {#editable-pptx}

自 v0.5.0 起，PPTX 由文本框、形状和图片组成（参见
[命令行]({{< relref "cli#editable-pptx-and-odp" >}})）。测于
2026-10-19 22:51，1 分钟平均负载为 7.1，共三轮：

| 演示文稿 | go-slides | marp-cli | go-slides 大小 | marp-cli 大小 | go-slides 内存 | marp-cli 内存 |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| 100 张幻灯片 | 348 ms | 7,795 ms | 760 KiB | 7.2 MiB | 204 MiB | 1,370 MiB |
| gaia，20 张幻灯片 | 112 ms | 4,608 ms | 893 KiB | 4.8 MiB | | |
| 数学公式与代码，20 张幻灯片 | 207 ms | 2,780 ms | 842 KiB | 1.5 MiB | | |

对于 100 张幻灯片的演示文稿，v0.4.0 耗时 4.7 s，大小 4.8 MiB，内存 540 MiB。marp-cli 的
PPTX 由图片组成。可编辑演示文稿的体积大部分来自内嵌字体。

## 为什么快 {#why-it-is-fast}

- **无需启动浏览器。** marp-cli 每次导出 PDF、图片和
  PPTX 时都要启动 Chrome；go-slides 则在自己的进程内用
  [go-webengine](https://github.com/go-webengine/engine) 渲染。
- **幻灯片并行渲染**，默认同时渲染四张，每张图片一轮到就立即写出：
  长篇演示文稿永远不会以图片形式整体驻留在内存中。
- **字体每个进程只解析一次**，而不是每张幻灯片解析一次。

## 不占优势的方面 {#what-it-does-not-win}

- **v0.3.0 的 PDF 大小。** 当时 PDF 中每张幻灯片都是一张图片：100 张幻灯片的演示文稿为 10 MiB，
  marp-cli 为 837 KiB。上文的矢量 PDF 已经比 marp-cli 的更小。
- **大型演示文稿的 CPU 时间。** 四个工作线程消耗的 CPU 时间多于实际耗时所显示的：
  在 v0.3.0 中，100 张幻灯片的 PDF 为 3.2 s，marp-cli 为 2.7 s
  （矢量 PDF 为 0.9 s）。
- **marp-cli 的内存是一个上限值**：Chrome 的各进程共享内存页，而求和时每个进程都计入了这些页。
  其中最大的单个进程为 220–320 MiB。

测试方法、每个演示文稿和模式，以及复现这些数据的命令，见
[BENCHMARKS.md](https://github.com/go-slides/slides/blob/main/BENCHMARKS.md)。
