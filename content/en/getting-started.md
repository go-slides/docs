---
title: "Getting started"
weight: 10
description: "A first deck, built to HTML and expanded back to Marp Markdown."
tags: [start]
---

## Install

Go 1.27.2 or later:

```sh
go install github.com/go-slides/slides/cmd/slides@latest
```

`slides` is one static Go binary with no cgo: nothing else to install, no
Node, no Chrome. This page describes v0.7.1.

## A first deck

Any Marp deck is one. Save this as `talk.md`:

```markdown
---
marp: true
theme: default
paginate: true
footer: "{{ .talk }}"
data:
  talk: Pure Go slides
  topics: [Why, How, Next]
---

<!-- _kind: title -->
# {{ .talk }}

{{ range .topics }}
---

# {{ . }}
{{ end }}
```

The front matter is Marp's, plus `data:`. The body is a Go template: the
`range` writes one slide per topic, each starting with `---`, and the first
slide is of the built-in kind `title` (no header, no footer, no page number).

## Build it

```sh
slides talk.md                 # writes talk.html beside the deck
slides -o slides.html talk.md  # or somewhere else; -o - for standard output
```

`slides` with no command is `slides build`. Flags may come before or after the
deck. Open the HTML in a browser; the keyboard moves between slides.

## Export images, PDF, PPTX and ODP

```sh
slides talk.md --pdf                  # one page per slide, with an outline
slides talk.md --images png           # talk.001.png, talk.002.png, …
slides talk.md --image jpeg           # the first slide only
slides talk.md -o out/slides.pdf      # the extension picks the format; out/ is created
slides talk.md --images jpeg --jpeg-quality 90
slides talk.md --images png --image-scale 2   # 2560×1440 for 16:9
slides talk.md --pdf --allow-local-files      # images from under the deck's directory
slides talk.md --pptx                         # or -o talk.pptx, with notes
slides talk.md --odp                          # or -o talk.odp: the same, as OpenDocument
```

No browser is involved: go-webengine renders the slides and go-pdfkit writes
the PDF. Local images are refused unless `--allow-local-files` is given.
Since v0.5.0 the slides of a PPTX or ODP are editable — text boxes, shapes
and pictures, fonts embedded — where marp-cli's are pictures; `--pptx-raster`
keeps pictures ([Editable PPTX and ODP]({{< relref "cli#editable-pptx-and-odp" >}})). Since v0.4.0 the PDF is vectors, its text
selectable (v0.3.0 wrote one image per page). See [Command line]({{< relref "cli#vector-pdf" >}}).

## Expand it

```sh
slides expand talk.md
```

prints the same deck as **plain Marp Markdown** (output of `main` at
`4ae5c49`, 2026-10-18):

```markdown
---
marp: true
theme: default
paginate: true
footer: "Pure Go slides"
---
<!--
_class: "lead"
_footer: ""
_paginate: "false"
-->


# Pure Go slides


---

# Why

---

# How

---

# Next
```

The four slides the `range` generated, the footer with its value, and the
`title` kind replaced by the spot directives it stands for. Anything that reads Marp — marp-cli, the VS
Code extension — can take it from there, so nothing written with go-slides'
additions is locked in.

## Next

- [Marp compatibility]({{< relref "marp" >}}): what is read the way Marp reads it.
- [Templates]({{< relref "templates" >}}) and [Kinds]({{< relref "kinds" >}}): the two additions.
- [Command line]({{< relref "cli" >}}): every command and flag.
