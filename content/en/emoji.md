---
title: "Emoji"
weight: 36
description: "Emoji shortcodes and Unicode emoji become Twemoji images, with no network: in the HTML page, images, PDF, PowerPoint and OpenDocument."
tags: [emoji]
---

{{< callout type="warning" >}}
Emoji are in v0.4.0; v0.3.0 had none.
{{< /callout >}}

Write an emoji as a shortcode or as the character itself:

```markdown
# Launch :rocket: 🎉

Shortcodes :smile: :+1: :heart_eyes: and Unicode 😄 👍🏽 👩‍💻 🇫🇷 1️⃣.
```

Each becomes the [Twemoji](https://github.com/jdecked/twemoji) image Marp Core
draws, with the markup Marp Core writes:

```html
<img class="emoji" draggable="false" alt="😄" src="data:image/svg+xml;base64,…" data-marp-twemoji=""/>
```

## No network

Marp Core's `src` is a URL on a CDN, so a deck shown offline, printed, or
turned into a PDF on a machine without a connection shows broken images.
go-slides puts the SVG itself in a `data:` URI. The graphics (4,009 of them,
Twemoji 17.0.3, Emoji 17.0) are embedded in the program from the pure-Go module
[go-icons/twemoji](https://github.com/go-icons/twemoji), 2 MB compressed and
read only when a deck has an emoji.

The images are drawn by the HTML page, by `--images`/`--image`, by `--pdf`, by
`--pptx` and by `--odp` alike: go-webengine rasterizes the SVG. All 4,009
graphics parse and draw under its SVG rasterizer; a test checks it. Since
v0.5.0 that rasterizer is [go-gfx/oksvg](https://github.com/go-gfx/oksvg)
v0.2.0: `clipPath`, `fill-rule: evenodd` and arc commands now draw as Chrome
draws them, which changed one Twemoji and 211 others, and SVG logos are no
longer garbled.

The rule that sizes an emoji to the text (`img[data-marp-twemoji]`, one `em`
square, aligned with the baseline) is added to the theme as Marp Core adds it,
so a theme of your own gets it too.

## What is converted

The rules are Marp Core's; a test compares them with the HTML Marp Core 4.4.0
writes for a deck of every case.

| | |
| --- | --- |
| `:smile:`, `:+1:`, `:heart_eyes:` | the 1,903 names of markdown-it-emoji, in text, headings, links, tables, `header` and `footer`; not in code, URLs or an image's alt text |
| `😄`, `👩‍💻`, `🇫🇷`, `1️⃣`, `👍🏽` | Unicode emoji, in text **and in code**: Marp Core rewrites the rendered code HTML |
| `✔`, `©`, `™`, `▶` | stay text: characters whose default presentation is text are emoji only with U+FE0F (`✔️`); `❤` is the exception |
| `☺︎` | stays text: a U+FE0E after an emoji asks for the plain glyph |
| `12:30:45`, `:not_a_name:` | stay as written |

A heading's `id` and the outline use the emoji, not the shortcode's name:
`# Launch :rocket:` has the id `launch-%F0%9F%9A%80`, as in Marp.

Where go-slides differs from Marp Core, it converts more: `✌️ ☝️ ✍️ 🏋️` and a
few other text-default characters that Marp Core's emoji-regex leaves as text,
and the handful of sequences Twemoji draws that Unicode does not list.

## Turning it off

`--no-emoji` (in code, `slides.Config{NoEmoji: true}`) leaves shortcodes and
Unicode emoji as the text they were written as, which is Marp Core's
`emoji: false`.

## Attribution

The graphics are Twemoji, Copyright 2020 Twitter, Inc and other contributors,
licensed [CC-BY 4.0](https://creativecommons.org/licenses/by/4.0/) and
continued by the [jdecked/twemoji](https://github.com/jdecked/twemoji) project.
A deck that shows them should credit them, for instance in a footer or on a
closing slide. The shortcode table is markdown-it-emoji's (MIT).
