---
title: "Playground"
weight: 15
description: "Write a deck in your browser and watch the slides follow: go-slides compiled to WebAssembly, with completion, help and examples."
tags: [playground, browser]
---

**[Open the playground →](https://go-slides.github.io/playground/)**

The playground is go-slides itself, v0.7.1, compiled to WebAssembly and
running in your browser. Type a deck on the left; the slides follow on the
right. Nothing is sent to a server: not the deck, not a share link.

## What it has

- **Live preview** of every slide, rendered off the page's thread (a Web
  Worker), so typing never waits for it. Click a slide to move the editor
  to its line; move the caret and the preview scrolls to its slide.
- **Slideshow**: the presentation itself, with the keyboard, the
  on-screen controller and the transitions.
- **Presenter view**: **Present** (or `p` in the slideshow) opens it in a
  second window: the current slide, the next step, the speaker notes (the
  deck's `<!-- comments -->`), a clock and a timer. The arrow keys, Space,
  Home and End work in either window, and each follows the other, also
  when you edit the deck. If the window does not open, allow pop-ups for
  the site. The overview view (`o`) is in the downloaded page only.
- **Completion** as you type: directive keys (spot `_` ones in comments),
  the values of `kind`, `theme`, `size`, `paginate`, and template functions
  and data fields inside `{{ }}`.
- **Help** for what is under the caret: a directive, a kind (what it sets),
  a template function. The help pane also links to these pages.
- **Problems** on their line: template, YAML and HCL errors are highlighted
  in the editor and listed under it, and the preview keeps the last deck
  that rendered.
- **Examples**: a welcome deck, templates that generate slides, slide kinds,
  an HCL front matter, themes and spot directives, and math, images and
  backgrounds (`![bg]`, a split `![bg left:40%]`, filters).
- **Images**: drop or paste an image file (PNG, JPEG, GIF, WebP or SVG, up
  to 256 KiB) into the editor and it goes into the deck as a `data:` URI. It
  travels with the deck and nothing is uploaded. `https:` images work too,
  but the browser fetches them from their server, which then sees your
  address. The preview sends no referrer.
- **Highlighting** of the front matter's keys at any depth, in YAML (also
  in `{a: 1}` mappings) and in HCL.
- **Share link**: the deck travels in the link itself (in its `#` part,
  compressed), so opening the link opens the deck. Links carry up to
  64 KiB compressed, so a deck with images is shared as a download
  instead.
- **Download HTML**: the presentation as one page, the same as
  `slides deck.md` writes.

PDF, PPTX, ODP and images are not made in the browser. go-slides' export
engines compiled to WebAssembly would weigh 21 MiB more on the wire. Use
the [command]({{< relref "cli" >}}) or the
[VS Code extension]({{< relref "vscode" >}}) for them.

## Security

The deck is shown in an iframe sandboxed without the page's origin: it can
reach neither the playground nor its storage, submit no form and navigate
nothing. Its page carries a Content Security Policy that runs only its own
scripts. Raw HTML in the deck follows Marp's allowlist, as `slides` does by
default. A render that runs for more than 5 seconds (a template that loops)
is stopped.

The presenter window keeps that sandbox. It is opened by the playground
itself and shows the deck in an iframe sandboxed the same way. The two
iframes cannot share storage, so the playground relays the slide number
between them in messages. No frame that runs the deck gets the
playground's origin.

## Size and speed

The module is 8.8 MiB as GitHub Pages serves it (gzip; Pages does not
serve brotli). It is downloaded once with a progress bar, then cached by the
browser until the next deploy. On an Apple M4 Max with Chrome 155, measured
against the live page, the editor is ready 2.1 s after the page opens, 1.5 s
of which is the download, and the first slides appear at 2.4 s. A change to
the 100-slide benchmark deck reaches the preview in about 70 ms, 37 ms of
which is go-slides rendering it.
