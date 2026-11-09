---
title: "Animated SVG"
weight: 40
description: "An SVG image with SMIL or CSS animations: played by the browser in HTML, animated shapes in PowerPoint and OpenDocument, an animated GIF where those cannot say it, the initial state in PDF and images."
tags: [svg, animation]
---

An SVG image that animates itself is written like any image:

```markdown
![](spinner.svg)
![w:400](chart.svg)
```

## What is read {#read}

The animations written in the SVG file itself: SMIL (`animate`, `animateTransform`, `animateMotion`, `animateColor`, `set`, with `begin`, `dur`, `repeatCount`, `fill`, `calcMode`, `keyTimes` and `keySplines`) and CSS (`@keyframes` with the `animation` properties, `transform-origin` and `transform-box`). What is not understood, an event such as `begin="click"` for instance, is left out and reported as a warning: nothing is dropped silently.

## What each output shows {#outputs}

- **HTML**: the browser plays the animations by itself.
- **PowerPoint** (`--pptx`) and **OpenDocument** (`--odp`): each animated element is a shape of its own, and its animations start with the slide: movement, rotation and scale, opacity, a change between two colours. The movement of an element is sampled at the animation's own times, so easing, several keyframes and a rotation about another point are exact. Repeats (`indefinite` too), delays and reversals are carried.
- **Frames**: when an animation cannot be said that way (a path that morphs, a stroke that draws itself, a skew), the whole image is rendered at the animation's own times and embedded as an animated GIF, which PowerPoint and Impress play. It is bounded: 12 frames a second, at most 240 frames, 16 million pixels over all the frames.
- **PDF** (`--pdf`) and images (`--images`): they cannot animate. The drawing is shown in its initial state, with no JavaScript and no `RichMedia`, and one warning names the animated images.

Each animated image gets one warning line with the code `svg-animation`, saying which way it was carried: natively, or as frames and why.

## Support per output {#support}

| Feature | HTML | PowerPoint | OpenDocument | PDF and images |
|---|---|---|---|---|
| Translation, movement along a path | plays | native | native | initial state |
| Rotation and scale, about any point | plays | native | native | initial state |
| Radius of a circle, size of a rectangle (no stroke) | plays | native | native | initial state |
| Opacity of a one-shape element | plays | native | native | initial state |
| Colour of a fill or a stroke, from one colour to another | plays | native | native | initial state |
| Repeats, delays, reversals, freezing at the end | plays | native | native | initial state |
| Easing, several keyframes | plays | native (sampled) | native (sampled) | initial state |
| Steps, discrete changes | plays | frames | frames | initial state |
| Path morphing, dash offset, position and size attributes, gradients | plays | frames | frames | initial state |
| Skew, reflection, several clocks on one element, nested animated elements | plays | frames | frames | initial state |
| Events | plays | not carried (warning) | not carried (warning) | initial state |

## Local files {#local-files}

A local image is read only with `--allow-local-files`, from the deck's folder, as in marp-cli. Without it the image is left out of the exports, and the build says so with `local-file-blocked`, naming the file, its slides and the option. A file that does not exist is reported as `local-file-missing`.

## Limits {#limits}

An SVG is untrusted input, so it is bounded: 8 MiB of SVG, 50,000 elements, 2,000 animations, 256 keyframes for one animation, 60 seconds of timeline and 240 frames. A drawing past a bound is refused with a warning, and the build goes on.

## What was checked {#checked}

LibreOffice Impress (macOS) opens the PowerPoint and OpenDocument files and keeps every animation when it saves them again, in both formats. Keynote (macOS) imports the shapes and the animated GIF; its animations were not inspected. PowerPoint: not yet verified.
