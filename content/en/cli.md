---
title: "Command line"
weight: 50
description: "Every command and flag of slides, and the slides.yaml configuration file."
tags: [cli]
---

{{< callout type="warning" >}}
This page describes v0.4.0.
{{< /callout >}}

```text
slides [deck.md] [flags]
slides [command] [flags] deck.md
```

## Commands

| Command | What it does |
| --- | --- |
| `build` | writes the presentation as HTML, or as a PDF or images (see below). **The default**: `slides deck.md` is `slides build deck.md` |
| `expand` | prints the deck as plain Marp Markdown: templates expanded, kinds replaced by the directives they stand for |
| `notes` | prints the speaker notes |
| `kinds` | lists the kinds the deck can use and what they set |
| `themes` | lists the themes |
| `version` | prints the version (so does `slides --version`) |
| `completion` | prints a shell completion script: `bash`, `zsh`, `fish` or `powershell` |

`slides <command> --help` prints the flags of a command. Flags may come before
or after the deck: `slides deck.md -o out.html` and `slides -o out.html
deck.md` are the same. Long flags take two dashes (`--html`); `-o` is the only
short one. `build` writes `deck.html` beside `deck.md` unless `-o` says
otherwise.

## Flags

| Flag | Commands | What it does |
| --- | --- | --- |
| `-o`, `--output file` | `build`, `expand`, `notes` | output file; `-` for standard output |
| `--config file` | all | configuration file; by default `slides.yaml` next to the deck, if there is one |
| `--no-config` | all | do not read a configuration file |
| `--theme-set path` | all | a theme CSS file, or a directory of them (repeatable) |
| `--kinds file` | all | a YAML file of kinds (repeatable) |
| `--data file` | all | a YAML or JSON file of template data (repeatable) |
| `--set key=value` | all | one template value (repeatable) |
| `--no-template` | all | do not expand Go templates |
| `--strict-yaml` | all | parse directives as strict YAML, as plain Marpit does (Marp Core parses them loosely) |
| `--no-emoji` | all | leave `:shortcodes:` and Unicode emoji as text instead of drawing them as images (since v0.4.0: [Emoji]({{< relref "emoji" >}})) |
| `--html` | `build` | allow all raw HTML; by default, Marp Core's allowlist |
| `--lang code` | `build` | the page's `lang` when the deck sets none |
| `--bare` | `build` | stack the slides, without the presenting script |
| `--pdf` | `build` | write a PDF, one page per slide, with an outline — text and shapes as vectors (also chosen by `-o file.pdf`) |
| `--pdf-raster` | `build` | write the PDF with each page one image of the slide, as v0.3.0 did |
| `--images png\|jpeg` | `build` | write one image per slide: `deck.001.png`, … |
| `--image png\|jpeg` | `build` | write the first slide as an image (also chosen by `-o file.png` or `-o file.jpg`) |
| `--pptx` | `build` | write a PowerPoint file with its speaker notes (also chosen by `-o file.pptx`): editable after v0.4.0, one picture per slide in v0.4.0 |
| `--pptx-raster` | `build` | write a PowerPoint file of one full-slide picture per slide, as marp-cli does (after v0.4.0); `--image-scale` defaults to 2 |
| `--odp` | `build` | write an OpenDocument presentation with its speaker notes (also chosen by `-o file.odp`): editable after v0.4.0 |
| `--odp-raster` | `build` | write an OpenDocument presentation of one picture per slide (after v0.4.0) |
| `--jpeg-quality n` | `build` | JPEG quality, 1 to 100 (default 85) |
| `--image-scale n` | `build` | device pixels per CSS pixel for images and PDF (default 1); 2 gives 2560×1440 for 16:9, with text and vectors drawn at that density, not enlarged |
| `--allow-local-files` | `build` | let images and PDF read local files, from under the deck's directory only; refused by default, as with marp-cli |

## Images, PDF, PPTX and ODP

`-o` with a `.pdf`, `.pptx`, `.odp`, `.png` or `.jpg` extension picks the format, as marp-cli
does. Output directories are created. A PDF has one page per slide, at the
slide's size, and an outline of the slides and their headings.

Rendering is done by [go-webengine](https://github.com/go-webengine/engine)
and the PDF is written by [go-pdfkit](https://github.com/go-pdfkit/pdfkit);
the fonts the built-in themes ask for (Lato, Roboto Mono) come from
[go-opentype/fonts](https://github.com/go-opentype/fonts) and are embedded, so
no render touches the network. There is no Chrome.

`--image-scale 2` renders at two device pixels per CSS pixel. Local images are
refused unless `--allow-local-files` is given, and even then only files under
the deck's directory are read: a deck from elsewhere cannot reach `~/.ssh`.
Several and split backgrounds (`![bg left:40%]`) render as Marp's layers do.

`--pptx` writes the deck with [go-ooxml](https://github.com/go-ooxml)'s
`WritePresentation`: each slide is one full-slide picture, with its speaker
notes and the deck's metadata. As in marp-cli, the slides are pictures, not
editable text. The scale defaults to 2 and `--image-scale` overrides it;
LibreOffice opens the result.

`--odp` writes the same thing as an OpenDocument presentation, through
go-odf's `WritePresentation`: one picture per slide, with its speaker notes.

### Editable PPTX and ODP

After v0.4.0 the slides of a PPTX or an ODP are **editable** — marp-cli's
PPTX is one picture a slide, and its experimental `--pptx-editable` converts
a PDF through LibreOffice. Here each slide is drawn from its layout, the one
the PNG and PDF exports draw:

- every line of text is a text box at its laid-out position, in the fonts
  the slide was laid out with — **embedded in the file**, so it looks the
  same on a machine without Lato or Inter (checked with LibreOffice 25.8) —
  with bold, italic, colours, letter spacing and clickable links;
- backgrounds, borders and rounded boxes are shapes; gradients are gradient
  fills in the PPTX and pictures in the ODP (an ODF gradient has two
  colours);
- pictures are stored once, however many slides show them;
- what a slide cannot express — an element under `opacity`, a `filter`, a
  blurred shadow, a formula — is a picture of that element, its text kept
  over it.

Rendered back by LibreOffice, a slide is within 0.04–2 % of the PNG export's
pixels. `--pptx-raster` and `--odp-raster` keep one picture a slide.

### Vector PDF

Since v0.4.0 the PDF is drawn as vectors: go-webengine paints the laid-out
slide into the PDF instead of into pixels. The text is text, in the fonts the
slide was laid out with, subset-embedded with a `/ToUnicode` map — it can be
selected, searched and copied. Backgrounds, borders and gradients are paths
and shadings, an image shown on many slides is embedded once, and links are
clickable. The 100-slide benchmark deck is 758 KiB, against 837 KiB for
marp-cli and 10 MiB for the image PDF of v0.3.0, and it is written in 0.3 s.

An element with no vector form here — under `opacity` < 1, with a `filter`
or a rotation, a blurred `box-shadow`, a gradient whose opacity fades — is
drawn as a bitmap of that element only, its text laid over it invisibly so it
stays selectable. Formulas are still bitmaps. `--pdf-raster` writes one image
per page, as v0.3.0 did.

## `slides.yaml`

A configuration file next to the deck applies to it, so a folder of decks
shares one:

```yaml
kinds: kinds.yaml        # a mapping of kinds, or file(s) holding one
themes:                  # theme CSS files or directories
  - themes/
data:                    # template data
  company: Example Inc.
html: false              # allow all raw HTML
lang: en
```

Kinds: the deck's own `kinds:` override the file's and the command line's.
Data: the file's and the command line's `data` override the deck's front
matter, so one deck can be built with different data.

## Planned

Not built yet: formulas as vectors in the PDF. When they arrive, they are a
command or flag here.
