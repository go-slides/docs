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
**On `main`, not released yet.** The engine and the command are merged in
[go-slides/slides](https://github.com/go-slides/slides), but no version is
tagged. These pages describe `main` and are published under `/docs/main/`;
they will describe the first release, under `/docs/0.1/`, once it is tagged,
and anything here may change before then.
{{< /callout >}}

## What it does today

| | |
| --- | --- |
| `slides build` (the default) | the deck as one HTML page: each slide an SVG-framed `<section>` that scales to the window, the theme's CSS, a keyboard viewer |
| `slides build --pdf`, `--images png`, `--image jpeg` | the deck as a PDF (one page per slide, with an outline) or as PNG/JPEG images, with no browser: [go-webengine](https://github.com/go-webengine) renders, [go-pdfkit](https://github.com/go-pdfkit) writes the PDF. [Known limits]({{< relref "cli#images-and-pdf" >}}) apply |
| `slides expand` | the deck as plain Marp Markdown: templates evaluated, kinds replaced by the spot directives they stand for |
| `slides notes` | the speaker notes |
| `slides kinds`, `slides themes` | what a deck can name |
| `slides.yaml` | shared kinds, themes, data, HTML policy and language for a folder of decks |

## Planned

Not built yet, and listed so nobody has to guess: **PPTX** through
[go-ooxml](https://github.com/go-ooxml), **math** through
[go-tex](https://github.com/go-tex), **code highlighting** through
[go-rouge](https://github.com/go-rouge), and emoji. The difference with
marp-cli is the route: it renders images and PDF by driving a headless Chrome;
here that is done by Go libraries.

## Pages

- [Getting started]({{< relref "getting-started" >}})
- [Marp compatibility]({{< relref "marp" >}})
- [Templates]({{< relref "templates" >}})
- [Kinds]({{< relref "kinds" >}})
- [Command line]({{< relref "cli" >}})
