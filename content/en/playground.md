---
title: "Playground"
weight: 15
description: "Write a deck in your browser and watch the slides follow: go-slides compiled to WebAssembly, with completion, help and examples."
tags: [playground, browser]
---

**[Open the playground →](https://go-slides.github.io/playground/)**

The playground is go-slides itself, v0.7.0, compiled to WebAssembly and
running in your browser. Type a deck on the left; the slides follow on the
right. Nothing is sent to a server: not the deck, not a share link.

## What it has

- **Live preview** of every slide, rendered off the page's thread (a Web
  Worker), so typing never waits for it. Click a slide to move the editor
  to its line; move the caret and the preview scrolls to its slide.
- **Slideshow**: the presentation itself, with the keyboard, the
  on-screen controller and the transitions. The presenter view needs a
  second window, so it works in the downloaded page rather than in the
  preview.
- **Completion** as you type: directive keys (spot `_` ones in comments),
  the values of `kind`, `theme`, `size`, `paginate`, and template functions
  and data fields inside `{{ }}`.
- **Help** for what is under the caret: a directive, a kind (what it sets),
  a template function. The help pane also links to these pages.
- **Problems** on their line: template, YAML and HCL errors are highlighted
  in the editor and listed under it, and the preview keeps the last deck
  that rendered.
- **Examples**: a welcome deck, templates that generate slides, slide kinds,
  an HCL front matter, themes and spot directives.
- **Share link**: the deck travels in the link itself (in its `#` part,
  compressed), so opening the link opens the deck. Decks up to 256 KiB.
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

## Size and speed

The module is 8.7 MiB as GitHub Pages serves it (gzip), downloaded once with
a progress bar and kept in the browser's cache until the next release. On
an Apple M4 Max with Chrome 155, it starts in a quarter of a second once
downloaded. A change to the 100-slide benchmark deck reaches the preview in
70–100 ms, 35 ms of which is go-slides rendering it.
