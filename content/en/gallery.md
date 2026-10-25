---
title: "Gallery"
weight: 12
description: "Twelve small decks, each showing one thing well: themes, kinds, templates, HCL, math, images, code, emoji, transitions and a complete talk. Open any of them in the playground."
tags: [gallery, examples]
---

Twelve small decks from the
[`examples/`](https://github.com/go-slides/slides/tree/main/examples) folder of
go-slides, each showing one thing well. The thumbnails are the first slide of
each deck, rendered by go-slides itself when this site is built; CI builds every
deck to HTML and PDF, so none of them breaks.

**Open in the playground** loads the deck in your browser, in a link that sends
nothing to any server. A deck that has files besides `deck.md` (a kinds file,
included parts, images) cannot travel in a link: download its **.zip**, and
import it with the playground's file panel, or unzip it and build it with
`slides build --allow-local-files deck.md`.

The pictures are original SVGs of the go-slides authors, released under
[CC0 1.0](https://creativecommons.org/publicdomain/zero/1.0/), and the decks
that use them say so.

{{< gallery >}}
