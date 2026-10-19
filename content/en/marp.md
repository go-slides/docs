---
title: "Marp compatibility"
linkTitle: "Marp"
weight: 20
description: "The Markdown slides reads is Marpit's and Marp Core's, not a dialect of it."
tags: [marp]
---

A deck written for [Marp](https://marp.app) is a go-slides deck. The rules
below are [Marpit](https://marpit.marp.app)'s and
[Marp Core](https://github.com/marp-team/marp-core)'s; where go-slides
differs, this page says so.

## Slides

A horizontal ruler — `---` on a line of its own, as Marpit splits — starts a
new slide, and so does a heading of the level the `headingDivider` directive
names. The front matter, between the first two `---`, holds global directives
and the deck's data.

## Directives

| Kind | Where | Applies to |
| --- | --- | --- |
| global | front matter, or `<!-- theme: gaia -->` anywhere | the whole deck: `theme`, `style`, `headingDivider`, `lang`… |
| local | `<!-- paginate: true -->` | this slide and the ones after it |
| spot | `<!-- _class: lead -->` (underscore) | this slide only |

The local directives are Marpit's: `paginate`, `header`, `footer`, `class`,
`backgroundColor`, `backgroundImage`, `backgroundPosition`,
`backgroundRepeat`, `backgroundSize`, `color`. Directives are parsed as
loose YAML, as Marp Core does; `--strict-yaml` parses them as plain Marpit
does.

## Images

`![bg](image.jpg)` makes an image the slide's background, with Marp's
keywords: `bg contain`, `bg cover`, `bg 50%`, `bg left`/`bg right` for a split
background, several `bg` images side by side, and the image filters
(`blur`, `sepia`, `grayscale`…). Inline image sizes (`w:200px`, `h:`) are
read as Marp reads them.

## Around the slide

- **Header and footer**: `header:` and `footer:` directives, in Markdown.
- **Pagination**: `paginate: true`.
- **Speaker notes**: an HTML comment that is not a directive is a note.
  `slides notes` prints them.
- **Fragmented lists**: a list written with `*` (or `1)`) appears item by item,
  as in Marp.
- **`<style scoped>`**: CSS for one slide only.
- **Raw HTML**: Marp Core's allowlist by default; `--html` allows all of it.
- **Emoji**: `:shortcodes:` and Unicode emoji become Twemoji images, as in
  Marp Core, but embedded instead of loaded from a CDN: see
  [Emoji]({{< relref "emoji" >}}).

## Themes

`default`, `gaia` and `uncover` are built in, selected with `theme:`. A
theme of your own is a CSS file with Marp's `/* @theme name */` comment,
given with `--theme-set` or listed under `themes:` in `slides.yaml`.

## Where go-slides differs

- **Templates.** `{{ }}` in the body is a Go template action (see
  [Templates]({{< relref "templates" >}})). A deck that shows `{{` as text
  outside a code block needs `--no-template`, or the text escaped as Go
  templates escape it. Code blocks and code spans are left alone.
- **Kinds.** `kind:` and `_kind:` are directives Marp does not have (see
  [Kinds]({{< relref "kinds" >}})). `slides expand` replaces them with the
  directives they stand for, so the expanded deck is plain Marp.
- **Outputs.** HTML, PNG, JPEG, PDF, PPTX and ODP, through Go libraries rather
  than a headless browser (see [Command line]({{< relref "cli#images-pdf-pptx-and-odp" >}})
  for the known limit). In v0.4.0, as in marp-cli, a PPTX holds pictures; since
  v0.5.0 its slides are editable text boxes and shapes, `--pptx-raster` keeping
  the pictures.
- **Fitting headings.** `# <!-- fit --> Title` is supported: the heading fills
  the slide width on one line, in the page and in every export.
- **Math and code.** `$…$` and `$$…$$` follow Marp Core's delimiter rules and
  are typeset by go-tex/math; fenced code is highlighted by go-rouge with
  highlight.js class names (see [Math and code]({{< relref "math-and-code" >}})).
