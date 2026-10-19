---
title: "go-slides documentation"
linkTitle: "Home"
type: docs
cascade:
  type: docs
description: "Marp-compatible slide decks in pure Go (CGO=0), with Go template interpolation and slide kinds."
---

**`slides` turns a Marp deck into a presentation, in pure Go.** It reads the
Markdown [Marp](https://marp.app) reads — Marpit's directives, `---` between
slides, `![bg]` backgrounds, speaker notes, the `default`, `gaia` and
`uncover` themes — and adds two things Marp does not have:

- **[Go templates]({{< relref "templates" >}})** inside the deck, evaluated
  before the Markdown is parsed, so a deck can generate its own slides;
- **[slide kinds]({{< relref "kinds" >}})**: a slide says what it is
  (`title`, `section`…) instead of repeating how it looks.

{{< callout type="warning" >}}
**v0.3.0.** These pages describe
[go-slides/slides v0.3.0](https://github.com/go-slides/slides/releases/tag/v0.3.0)
and are published under `/docs/0.3/`.
{{< /callout >}}

## What it does today

| | |
| --- | --- |
| `slides build` (the default) | the deck as one HTML page: each slide an SVG-framed `<section>` that scales to the window, the theme's CSS, a keyboard viewer |
| `slides build --pdf`, `--images png`, `--image jpeg` | the deck as a PDF (one page per slide, with an outline) or as PNG/JPEG images, with no browser: [go-webengine](https://github.com/go-webengine) renders, [go-pdfkit](https://github.com/go-pdfkit) writes the PDF — as vectors after v0.3.0, its text selectable: [Vector PDF]({{< relref "cli#vector-pdf" >}}) |
| math and code | `$…$` and `$$…$$` typeset by go-tex/math, fenced code highlighted by go-rouge, in HTML, images and PDF: [Math and code]({{< relref "math-and-code" >}}) |
| emoji | `:smile:` and 😄 become Twemoji images, embedded, so they need no network and draw in every output: [Emoji]({{< relref "emoji" >}}) (after v0.3.0) |
| `slides build --pptx` | the deck as a PowerPoint file: one full-slide picture per slide, with its speaker notes and the deck's metadata (not editable text, as in marp-cli) |
| `slides build --odp` | the same as an OpenDocument presentation (`.odp`) |
| fitting headings | `# <!-- fit --> Title` fills the slide width on one line, in the page and in every export |
| `slides expand` | the deck as plain Marp Markdown: templates evaluated, kinds replaced by the spot directives they stand for |
| `slides notes` | the speaker notes |
| `slides kinds`, `slides themes` | what a deck can name |
| `slides.yaml` | shared kinds, themes, data, HTML policy and language for a folder of decks |
| HCL | `---hcl` front matter, `slides.hcl`, kinds and data files in HCL2, with expressions and functions: [HCL]({{< relref "hcl" >}}) |

## Planned

Not built yet, and listed so nobody has to guess: formulas as vectors in the
PDF (they are bitmaps there). The difference with
marp-cli is the route: it renders images and PDF by driving a headless Chrome;
here that is done by Go libraries.

## Pages

- [Getting started]({{< relref "getting-started" >}})
- [Marp compatibility]({{< relref "marp" >}})
- [Templates]({{< relref "templates" >}})
- [Kinds]({{< relref "kinds" >}})
- [HCL]({{< relref "hcl" >}})
- [Math and code]({{< relref "math-and-code" >}})
- [Emoji]({{< relref "emoji" >}})
- [Command line]({{< relref "cli" >}})
