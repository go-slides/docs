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

## Fonts

Lato Regular, Bold, Black and italics and Roboto Mono Regular, Bold and
italics are embedded, so gaia's headings get their weight and italic text
renders italic.
