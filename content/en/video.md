---
title: "Video"
weight: 39
description: "A video in a deck: played by the browser in the HTML page, a poster frame with a link in the PDF, embedded in PowerPoint and OpenDocument."
tags: [video]
---

A video is an image whose file ends in `.mp4`, `.m4v` or `.webm`:

```markdown
![video](clip.mp4)
![The demo](clip.mp4#t=12.5)
![demo loop muted](clip.mp4)
```

## Syntax {#syntax}

The file is read from the deck's folder, or from an `http` or `https` URL. An absolute path, a path that leaves the folder or any other scheme is refused. `#t=12.5`, or `#t=1:23` for minutes and seconds, picks the poster frame: the picture shown before playback and by the outputs that cannot play the video. The media-fragment form `#t=10,20` is accepted too, and its start is used. Without it, the poster is the first frame.

Keywords in the alternative text, written like Marp image keywords but go-slides' own, set how the video plays: `autoplay`, `loop` and `muted`. `autoplay` is honoured only together with `muted`; otherwise it is ignored, with a warning. The other words of the alternative text are the accessible label.

## What each output shows {#outputs}

- **HTML** (the page, and the playground's preview): a `<video controls playsinline preload="metadata" poster=...>` element. The poster is the decoded frame, embedded as a JPEG data URI. A video never plays with sound by itself. When the browser asks for less motion (`prefers-reduced-motion`), `autoplay` and `loop` are removed and the video is paused. A video that autoplays (muted) plays only while at least half of it is visible, and pauses when its slide is not shown.
- **PDF** (`--pdf`): the poster frame as a picture, with a link to the video file: the relative path as written in the deck, or the URL for an `http` or `https` source. Playback happens only by opening the link. The PDF carries no `RichMedia` annotation: it would play only in Acrobat, and is deliberately not used.
- **PowerPoint** (`--pptx`): the video is embedded as a media part (`ppt/media/slidesvideoN.mp4|webm`) with its poster, as `a:videoFile` plus the `p14:media` extension. Keynote (macOS): imports it as a movie and plays it (verified). LibreOffice Impress (macOS): imports PowerPoint and OpenDocument files as a video and decodes it (verified). PowerPoint: not yet verified. The `autoplay`, `loop` and `muted` keywords are not carried into PowerPoint or OpenDocument.
- **OpenDocument** (`--odp`): the video is embedded under `Video/` as a `draw:plugin` (ODF 1.3 Part 3, 10.4.8), before the poster `draw:image`, and listed in the manifest. LibreOffice Impress imports it as a video and decodes it (verified); live slideshow playback depends on its media backend and the codec.
- **URL sources**: the HTML page streams the video from the URL. The exports show a neutral card linked to the URL. Nothing is fetched at build time, and nothing is embedded in PowerPoint or OpenDocument.

## Containers and posters {#codecs}

MP4 (ISO BMFF) and WebM (Matroska) files are read, in Go; ffmpeg is never run. The poster is decoded for H.264 only (`avc1` or `avc3`), the one video decoder of go-avkit. Other codecs (H.265, VP8, VP9, AV1) still play in HTML and are embedded in PowerPoint and OpenDocument, but their poster is a neutral dark card with a play triangle, and the warning `video.codec` is printed.

## Limits {#limits}

A video is untrusted input, so it is bounded: the file (64 MiB), the boxes or elements (262,144) and their nesting (16 levels), the frame (4096 by 4096) and the poster decoding (at most 600 samples from the preceding keyframe, 8 MiB per sample). A poster too far from a keyframe gets the card and the warning `video.limit`.

A refused video is replaced in the slide by a note, `[video not shown: ...]`; the build goes on, and the problem is a warning with a code:

- `video.syntax`: not an MP4 or WebM file, no video track, or an invalid time;
- `video.limit`: a bound above is exceeded;
- `video.codec`: the poster cannot be decoded for this codec;
- `video.path`: an absolute path, a path that leaves the deck's folder, or a scheme other than `http` and `https`;
- `file.missing`: the file does not exist.

## In the playground {#playground}

Once released, a video can be added to the project like the other files. Its bytes stay in the browser, as blob URLs, and are never uploaded. Try it in the [playground](https://go-slides.github.io/playground/).

Video support loads on demand. The reader that decodes the poster frame (about 3 MiB compressed) is a module of its own, downloaded the first time a deck needs a poster. A deck without video never fetches it, and the playground's first load is that much smaller. It runs in a Web Worker, so typing is never blocked, and it reads a file with the same limits as `slides` (64 MiB, boxes and frame size) and the same `video.limit` and `video.syntax` errors.
