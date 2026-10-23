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

In v0.8.0 a block can number its lines. Marp Core has no line numbers, and
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

Each number starts its line inside the theme's code box, set apart by a rule.
It stays on its line's baseline even when the line is taller, for example
with a callout or an emoji. The numbers are muted by a colour mixed from the
code's colour and its background, with 4.5:1 contrast or more in the default,
gaia and uncover themes, light and dark. They are **not selectable** in the
page and are not read by screen readers. In a PDF they are drawn but left out
of the text, so a text extractor reads the code alone. In a PowerPoint or
OpenDocument file they are editable text at the start of each line. A deck
that does not use the option renders as before, to the byte. Highlighting
chosen lines is not supported yet.

### Code callouts

In v0.8.0 a code line can end with a numbered disc that refers to a note
below the block. The syntax is
[Quarto's](https://quarto.org/docs/authoring/code-annotation.html), which is
also Asciidoctor's comment form. The marker is a comment in the code's own
language, so the code stays valid, and the notes are the ordered list right
after the block:

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

| Write | Means |
| --- | --- |
| `// <1>`, `# <1>`, `-- <1>`, `; <1>`, `% <1>`, `' <1>` | a line comment that holds only the marker: both are replaced by the disc |
| `<!-- <1> -->`, `/* <1> */` | the same, for languages with wrapped comments |
| `// <1> <2>` | two discs on one line |
| `x = 1 # count <1>` | the comment's text stays and the disc follows it |

The markers become discs only when the list right after the block starts at 1
and has an item for every number. Otherwise the block stays as written. The
discs are rings in the code's colour, in the default, gaia and uncover themes,
light or dark. The list's items are numbered with the same discs.

The discs are **not selectable**. Copying the code from the page gives the code
without its markers. In a PDF they are vectors left out of the text, and in a
PowerPoint or OpenDocument file they are editable shapes. Callouts combine
with line numbers. Marp Core has no callouts (marp-core 4.4.0), so in Marp the
comments and the numbered list show as written. A deck that does not use them
renders as before, to the byte.

In v0.9.0, a disc and its note find each other in the HTML page. Hovering a
disc, or reaching it with Tab, tints its note, and a note does the same for its
disc. This is what Quarto's `code-annotations: hover` does, but here the style
sheet does it alone, with no script. It therefore works the same in the
playground's sandboxed preview and in VS Code's preview, which both restrict
scripts. Each disc is described by its note, so a screen reader reads the note
when it reaches the disc. The tint fades in only when the reader has not asked
for reduced motion. Images, PDF, PowerPoint and OpenDocument files are
unchanged, byte for byte.

## Fonts

Lato Regular, Bold, Black and italics and Roboto Mono Regular, Bold and
italics are embedded, so gaia's headings get their weight and italic text
renders italic.
