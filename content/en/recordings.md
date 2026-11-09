---
title: "Terminal recordings"
weight: 38
description: "An asciinema recording in a deck: replayed in a small player in the HTML page, drawn as vector text in the PDF, PowerPoint and OpenDocument."
tags: [asciinema, terminal]
---

{{< callout type="info" >}}
Terminal recordings are in v0.18.0 and later.
{{< /callout >}}

An [asciinema](https://docs.asciinema.org/) recording ([asciicast v2](https://docs.asciinema.org/manual/asciicast/v2/) or [v3](https://docs.asciinema.org/manual/asciicast/v3/)) is an image whose file ends in `.cast`:

```markdown
![cast](demo.cast)
![The failing build](demo.cast#t=12.5)
```

## Syntax {#syntax}

The file is read from the deck's folder: a URL, an absolute path or a path that leaves the folder is refused. The alternative text names the recording for screen readers; `cast` alone means the recording's own title. `#t=12.5`, or `#t=1:23` for minutes and seconds, picks the moment that the outputs which cannot move show; without it, that is the end of the recording.

## What each output shows {#outputs}

- **HTML** (the page, and the playground's preview): the recording replays in a small player of go-slides, embedded in the page: play and pause, a speed button and a position bar. It loads no script from anywhere and uses no network.
- **PDF, PowerPoint and OpenDocument** (`--pdf`, `--pptx`, `--odp`): the terminal at the moment `#t=` names, as vector text in the theme's monospace font, with the 16, 256 and direct colours, bold, italic, underline and inverse text. The text can be selected and searched. It is never a picture. Bold and italic are drawn where the theme's monospace font has them (Roboto Mono in the gaia theme).

## The player {#player}

It plays once, the first time the recording is on screen. When the browser asks for less motion (`prefers-reduced-motion`), it starts paused, showing the moment `#t=` names, and plays when you press play. Its buttons and keys do not turn the slide.

## Limits {#limits}

A recording is untrusted input, so it is bounded: the file (16 MiB), the events (500,000), the length (3 hours), the terminal (400 by 200 cells) and the work the emulator may do. A recording over a bound, or not a well formed asciicast file, is refused: the build goes on, the slide shows a note in its place, and the problem is a warning with a code: `cast.syntax`, `cast.limit`, `cast.version` or `cast.path`.

## In the playground {#playground}

Add a `.cast` file to the project from the Files panel, by dropping it or from the file picker. It goes into `media/` and is inserted as an image. It is text that the deck reads, so it counts with the project's other text files, which are limited to 2 MiB. Try it in the [playground](https://go-slides.github.io/playground/).
