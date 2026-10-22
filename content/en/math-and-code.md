---
title: "Math and code"
weight: 35
description: "Formulas between dollar signs are typeset, fenced code is highlighted, in the HTML page, images and PDF."
tags: [math, code]
---

{{< callout type="warning" >}}
Math and code are in v0.2.0.
{{< /callout >}}

## Math

`$…$` is an inline formula and `$$…$$` a display formula:

```markdown
The famous one is $E=mc^2$, and this is a display formula:

$$
\int_0^\infty e^{-x^2}\,dx = \frac{\sqrt{\pi}}{2}
$$
```

- The delimiter rules are Marp Core's, read from marp-core 4.4.0: a `$`
  followed by a space does not open a formula, and one followed by a digit
  does not close it, so prices such as `$5 and $10` stay text.
- [go-tex/math](https://github.com/go-tex/math) typesets each formula into
  inline SVG sized in `em`, so it follows the text's size and sits on its
  baseline. The same SVG is used in the HTML page, in images and in the PDF.
- A formula that does not parse stays visible as written; the error is kept in
  its `title`.

## Code

Fenced code blocks are coloured by [go-rouge](https://github.com/go-rouge/rouge),
which writes highlight.js class names, the ones Marp's themes style:

````markdown
```go
fmt.Println("hello")
```
````

A language go-rouge does not know is written as plain text, as Marp Core does.

### Line numbers

After v0.7.1 a block can number its lines. Marp Core has no line numbers, and
it ignores what follows a fence's language, so the options go in braces there
and a deck that uses them renders in Marp as it would without:

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

| Write | Means |
| --- | --- |
| `{lineNumbers}` | number from 1 (Slidev's `{lines:true}` and Pandoc's `{.numberLines}` are read too) |
| `{lineNumbers=10}` | number from 10 (also `startLine=10`, Pandoc's `startFrom="10"`) |
| `{lineNumbers=false}` | no numbers, whatever the slide says |

For every block of a slide and the slides after it, set the directive:
`<!-- lineNumbers: true -->` (`_lineNumbers` for one slide). A kind can set it
once for a family of slides:

```yaml
kinds:
  listing: {lineNumbers: true}
```

The numbers sit in a gutter inside the theme's code box, muted, and are
**not selectable** in the page nor read by screen readers. The default, gaia
and uncover themes look right. In a PDF they are text in a block of their own,
never interleaved with the code in what a text extractor reads; in a PowerPoint
or OpenDocument file they are a text box of their own. A deck that does not use
the option renders as before, to the byte. Highlighting chosen lines is not
supported yet.

## Fonts

Lato Regular, Bold, Black and italics and Roboto Mono Regular, Bold and
italics are embedded, so gaia's headings get their weight and italic text
renders italic.
