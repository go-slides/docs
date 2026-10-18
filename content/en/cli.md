---
title: "Command line"
weight: 50
description: "Every command and flag of slides, and the slides.yaml configuration file."
tags: [cli]
---

{{< callout type="warning" >}}
No version is tagged yet. This page is checked against `slides --help` built
from `main` after go-slides/slides#4 (2026-10-18); anything here may change
before the first release.
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
| `--html` | `build` | allow all raw HTML; by default, Marp Core's allowlist |
| `--lang code` | `build` | the page's `lang` when the deck sets none |
| `--bare` | `build` | stack the slides, without the presenting script |
| `--pdf` | `build` | write a PDF, one page per slide, with an outline (also chosen by `-o file.pdf`) |
| `--images png\|jpeg` | `build` | write one image per slide: `deck.001.png`, … |
| `--image png\|jpeg` | `build` | write the first slide as an image (also chosen by `-o file.png` or `-o file.jpg`) |
| `--jpeg-quality n` | `build` | JPEG quality, 1 to 100 (default 85) |
| `--image-scale n` | `build` | device pixels per CSS pixel for images and PDF (default 1); 2 gives 2560×1440 for 16:9, with text and vectors drawn at that density, not enlarged |
| `--allow-local-files` | `build` | let images and PDF read local files, from under the deck's directory only; refused by default, as with marp-cli |

## Images and PDF

`-o` with a `.pdf`, `.png` or `.jpg` extension picks the format, as marp-cli
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

{{< callout type="warning" >}}
Known limit, tracked in the library it belongs to: headings are lighter than
with Lato Black. go-opentype/fonts carries one regular instance of Lato and
Roboto Mono, so gaia and uncover headings are not as heavy.
{{< /callout >}}

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

Not built yet: PPTX output, math, code highlighting and emoji. When they arrive, they are commands or flags here.
