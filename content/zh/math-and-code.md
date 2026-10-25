---
title: "数学公式与代码"
weight: 35
description: "美元符号之间的公式会被排版，围栏代码会被高亮，适用于 HTML 页面、图片和 PDF。"
tags: [math, code]
---

{{< callout type="warning" >}}
数学公式与代码自 v0.2.0 起提供。
{{< /callout >}}

## 数学公式 {#math}

`$…$` 是行内公式，`$$…$$` 是独立公式：

```markdown
The famous one is $E=mc^2$, and this is a display formula:

$$
\int_0^\infty e^{-x^2}\,dx = \frac{\sqrt{\pi}}{2}
$$
```

- 定界符规则与 Marp Core 相同（依据 marp-core 4.4.0）：后面跟空格的 `$`
  不会开启公式，后面跟数字的美元符号不会结束公式，因此像 `$5 and $10` 这样的价格会保留为文本。
- [go-tex/math](https://github.com/go-tex/math) 将每个公式排版为以 `em` 为单位的行内 SVG，
  因此它会随文字大小变化，并位于文字的基线上。HTML 页面、图片和 PDF 使用的是同一份 SVG。
- 无法解析的公式会按原样显示；错误信息保存在其 `title` 中。

## 代码 {#code}

围栏代码块由 [go-rouge](https://github.com/go-rouge/rouge) 着色，
它输出 highlight.js 的类名，也就是 Marp 主题所设置样式的那些类名：

````markdown
```go
fmt.Println("hello")
```
````

go-rouge 不认识的语言会按纯文本输出，与 Marp Core 相同。

### 行号 {#line-numbers}

从 v0.8.0 起，代码块可以为各行编号。Marp Core 没有行号功能，
并且会忽略围栏中语言名之后的内容，所以这些选项写在那里的花括号中，
使用了它们的演示文稿在 Marp 中的渲染效果与不使用时相同：

````markdown
```go {lineNumbers}
func main() {
	fmt.Println("hello")
}
```

```go {lineNumbers=10}
// starts at 10
```
````

| 写法 | 含义 |
| --- | --- |
| `{lineNumbers}` | 从 1 开始编号（也支持 Slidev 的 `{lines:true}` 和 Pandoc 的 `{.numberLines}`） |
| `{lineNumbers=10}` | 从 10 开始编号（也可写作 `startLine=10`，或 Pandoc 的 `startFrom="10"`） |
| `{lineNumbers=false}` | 不编号，无论幻灯片如何设置 |

要对一张幻灯片及其后幻灯片中的所有代码块生效，请设置指令：
`<!-- lineNumbers: true -->`（只对一张幻灯片生效则用 `_lineNumbers`）。
kind（幻灯片类型）可以为一类幻灯片统一设置：

```yaml
kinds:
  listing: {lineNumbers: true}
```

每个行号位于主题代码框内对应行的开头，并以一条分隔线隔开。
即使该行较高（例如带有代码标注或 emoji），行号也始终位于该行的基线上。
行号使用由代码颜色和背景色混合而成的柔和颜色，在 default、
gaia 和 uncover 主题的浅色和深色模式下，对比度均不低于 4.5:1。行号在页面中**不可选中**，
屏幕阅读器也不会朗读。在 PDF 中，行号会被绘制，但不计入文本，因此文本提取工具只会读到代码本身。
在 PowerPoint 或 OpenDocument 文件中，行号是每行开头的可编辑文本。
不使用该选项的演示文稿，渲染结果与以前逐字节相同。目前尚不支持高亮指定的行。

### 代码标注 {#code-callouts}

从 v0.8.0 起，代码行末尾可以带一个编号圆标，指向代码块下方的一条说明。语法采用
[Quarto](https://quarto.org/docs/authoring/code-annotation.html) 的写法，
这也是 Asciidoctor 的注释形式。标记是代码自身语言中的注释，因此代码仍然有效；
说明则是紧跟在代码块之后的有序列表：

````markdown
```go
func main() { // <1>
	x := load() // <2>
	fmt.Println(x)
}
```

1. The entry point.
2. Reads the input.
````

| 写法 | 含义 |
| --- | --- |
| `// <1>`, `# <1>`, `-- <1>`, `; <1>`, `% <1>`, `' <1>` | 只包含标记的行注释：注释和标记一起被替换为圆标 |
| `<!-- <1> -->`, `/* <1> */` | 同上，适用于使用包围式注释的语言 |
| `// <1> <2>` | 同一行上的两个圆标 |
| `x = 1 # count <1>` | 注释文字保留，圆标跟在其后 |

只有当紧跟在代码块之后的列表从 1 开始、并且每个编号都有对应的列表项时，标记才会变成圆标；
否则代码块按原样保留。在 default、gaia 和 uncover 主题的浅色或深色模式下，
圆标都是与代码同色的圆环。列表项也使用同样的圆标编号。

圆标**不可选中**。从页面复制代码时，得到的是不含标记的代码。在 PDF 中，圆标是不计入文本的矢量图形；
在 PowerPoint 或 OpenDocument 文件中，圆标是可编辑的形状。代码标注可以与行号同时使用。
Marp Core 没有代码标注功能（marp-core 4.4.0），因此在 Marp 中，注释和编号列表会按原样显示。
不使用代码标注的演示文稿，渲染结果与以前逐字节相同。

从 v0.9.0 起，在 HTML 页面中，圆标与其说明可以相互关联。将指针悬停在圆标上，
或用 Tab 键移到圆标上，其说明会被着色；说明对其圆标也是如此。这与 Quarto 的
`code-annotations: hover` 效果相同，但这里完全由样式表实现，不需要脚本。
因此，它在 Playground 的沙盒预览和 VS Code 的预览中（两者都限制脚本）效果相同。
每个圆标都以其说明作为描述，因此屏幕阅读器读到圆标时会朗读对应的说明。
只有在读者未要求减弱动态效果时，着色才会以渐显方式出现。图片、PDF、PowerPoint 和
OpenDocument 文件保持不变，逐字节相同。

## 字体 {#fonts}

内嵌了 Lato Regular、Bold、Black 及其斜体，以及 Roboto Mono Regular、Bold 及其斜体，
因此 gaia 的标题能获得应有的字重，斜体文本也能正确显示为斜体。
