---
title: "Performance"
weight: 60
description: "How fast go-slides builds a deck, in how much memory, measured against marp-cli on the same machine."
tags: [performance]
---

{{< callout type="info" >}}
The PPTX figures below are for v0.5.0 and the vector PDF figures for v0.4.0.
The first table is for v0.3.0,
the previous release, which renders slides concurrently and parses its fonts
once ([go-slides/slides#11](https://github.com/go-slides/slides/pull/11)).
{{< /callout >}}

The 100-slide PDF and PPTX rows of the three tables below no longer hold since v0.18.0: see [Several cores](#several-cores) and [go-slides/slides#117](https://github.com/go-slides/slides/issues/117).

go-slides and [marp-cli](https://github.com/marp-team/marp-cli) 4.5.1 were
run as subprocesses on the same decks, on an Apple M4 Max (16 CPUs) with a
1-minute load average under 8, five interleaved rounds after a warm-up.
Time is the median wall time; memory is the peak of the whole process tree
(marp-cli drives a Chrome, whose processes count).

*Measured with v0.3.0 on all cores (16 CPUs).*

| Deck | Output | go-slides | marp-cli | go-slides memory | marp-cli memory |
| --- | --- | ---: | ---: | ---: | ---: |
| 100 slides | HTML | 36 ms | 315 ms | 47 MiB | 140 MiB |
| 100 slides | PDF | 1.14 s | 1.88 s | 242 MiB | 1,509 MiB |
| 100 slides | PNG, every slide | 0.97 s | 5.44 s | 225 MiB | 1,354 MiB |
| 100 slides | PPTX | 3.39 s | 8.03 s | 530 MiB | 1,373 MiB |
| gaia, 20 slides | PDF | 283 ms | 1.57 s | 154 MiB | 1,315 MiB |
| math and code, 20 slides | PNG, every slide | 250 ms | 2.32 s | 168 MiB | 1,286 MiB |

## Vector PDF

In v0.4.0 the PDF is drawn as vectors (see
[Command line]({{< relref "cli#vector-pdf" >}})). Same machine and method,
2026-10-19:

*Measured with v0.4.0 on all cores (16 CPUs), on 2026-10-19.*

| Deck | go-slides | marp-cli | v0.3.0 | go-slides size | marp-cli size | v0.3.0 size |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| 100 slides | 306 ms | 1,763 ms | 1,147 ms | 758 KiB | 837 KiB | 10.0 MiB |
| gaia, 20 slides | 93 ms | 1,546 ms | 285 ms | 32 KiB | 66 KiB | 582 KiB |
| math and code, 20 slides | 140 ms | 1,643 ms | 294 ms | 124 KiB | 335 KiB | 638 KiB |

Memory for the 100-slide PDF: 158 MiB, against 1,512 MiB for marp-cli; CPU
time 0.9 s against 2.7 s.

The v0.4.0 figures for the 100-slide PDF were measured on 2026-10-19 at 21:29,
with a 1-minute load average of 7.7, over three rounds: go-slides 306 ms, 758
KiB and 158 MiB; marp-cli 1,838 ms, 837 KiB and 1,512 MiB.

## Editable PPTX

Since v0.5.0 the PPTX is made of text boxes, shapes and pictures (see
[Command line]({{< relref "cli#editable-pptx-and-odp" >}})). Measured on
2026-10-19 at 22:51, with a 1-minute load average of 7.1, over three rounds:

*Measured with v0.5.0 on all cores (16 CPUs), on 2026-10-19.*

| Deck | go-slides | marp-cli | go-slides size | marp-cli size | go-slides memory | marp-cli memory |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| 100 slides | 348 ms | 7,795 ms | 760 KiB | 7.2 MiB | 204 MiB | 1,370 MiB |
| gaia, 20 slides | 112 ms | 4,608 ms | 893 KiB | 4.8 MiB | | |
| math and code, 20 slides | 207 ms | 2,780 ms | 842 KiB | 1.5 MiB | | |

For the 100-slide deck, v0.4.0 took 4.7 s, 4.8 MiB and 540 MiB. marp-cli's
PPTX holds pictures. Most of an editable deck's size is the embedded fonts.

## Several cores {#several-cores}

Measured with v0.26.0: wall time in ms with 1, 2, 4 and 16 cores, against marp-cli. The 1-minute load average was 3 to 4. marp-cli's time does not change with the number of cores, because its `-P` only spreads separate files across browsers. `slidesbench -cores` produces this table. The output files are byte-identical to those of v0.25.0.

| Deck, mode | 1 core | 2 | 4 | 16 | marp-cli |
| --- | ---: | ---: | ---: | ---: | ---: |
| 100 slides, PNG | 2629 | 1387 | 906 | 765 | 5181 |
| gaia, 20 slides, PNG | 745 | 395 | 225 | 222 | 2316 |
| math and code, 20 slides, PNG | 590 | 331 | 200 | 192 | 2122 |
| gaia, 20 slides, PDF | 87 | 65 | 57 | 59 | 1382 |
| math and code, 20 slides, PDF | 801 | 588 | 576 | 582 | 1447 |
| gaia, 20 slides, PPTX | 98 | 77 | 70 | 70 | 3036 |
| math and code, 20 slides, PPTX | 911 | 625 | 593 | 605 | 2502 |
| 100 slides, PDF | 3166 | 2573 | 2545 | 2568 | 1513 |
| 100 slides, PPTX | 3299 | 2560 | 2547 | 2586 | 7238 |

**Known defect.** The 100-slide PDF is slower than marp-cli, and the 100-slide PDF and PPTX do not scale with cores. This has been so since v0.18.0: the engine rebuilds the variable-font instances for every page, which costs 16× the memory and 3.5 to 4× the time on one core. It is tracked in [go-slides/slides#117](https://github.com/go-slides/slides/issues/117) and [go-webengine/engine#443](https://github.com/go-webengine/engine/issues/443), and is not fixed in v0.26.0.

## Why it is fast

- **No browser to start.** marp-cli launches Chrome for every PDF, image and
  PPTX export; go-slides renders with
  [go-webengine](https://github.com/go-webengine/engine) inside its own
  process.
- **Slides render in parallel**, four at a time by default, and each image is
  written as soon as it is its turn: a long deck is never held in memory as
  images.
- **Fonts are parsed once** per process, not once per slide.

## What it does not win

- **PDF size, in v0.3.0.** Each slide was an image in the PDF: 10 MiB against
  837 KiB for the 100-slide deck. The vector PDF above is smaller than
  marp-cli's.
- **CPU time on big decks.** Four workers spend more CPU than the wall time
  shows: in v0.3.0, 3.2 s against marp-cli's 2.7 s for the 100-slide PDF
  (0.9 s for the vector PDF).
- **Memory is an upper bound for marp-cli**: Chrome's processes share pages
  that the sum counts in each. Its largest single process is 220–320 MiB.

The method, every deck and mode, and the commands to reproduce the figures
are in
[BENCHMARKS.md](https://github.com/go-slides/slides/blob/main/BENCHMARKS.md).
