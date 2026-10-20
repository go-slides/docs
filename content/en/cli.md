---
title: "Command line"
weight: 50
description: "Every command and flag of slides, and the slides.yaml configuration file."
tags: [cli]
---

{{< callout type="warning" >}}
This page describes v0.7.0.
{{< /callout >}}

```text
slides [deck.md] [flags]
slides [command] [flags] deck.md
```

## Commands

| Command | What it does |
| --- | --- |
| `build` | writes the presentation as HTML, or as a PDF or images (see below). **The default**: `slides deck.md` is `slides build deck.md` |
| `expand` | prints the deck as plain Marp Markdown: templates expanded, kinds replaced by the directives they stand for |
| `notes` | prints the speaker notes |
| `kinds` | lists the kinds the deck can use and what they set; `--json` (in v0.6.0) writes them as `[{"name", "extends", "directives": [{"key", "value"}]}]`, `extends` resolved |
| `themes` | lists the themes |
| `version` | prints the version (so does `slides --version`) |
| `completion` | prints a shell completion script: `bash`, `zsh`, `fish` or `powershell` |

`slides <command> --help` prints the flags of a command. Flags may come before
or after the deck: `slides deck.md -o out.html` and `slides -o out.html
deck.md` are the same. Long flags take two dashes (`--html`); `-o` is the only
short one. `build` writes `deck.html` beside `deck.md` unless `-o` says
otherwise.

## Flags

| Flag | Commands | What it does |
| --- | --- | --- |
| `-o`, `--output file` | `build`, `expand`, `notes` | output file; `-` for standard output |
| `-I`, `--input-dir dir` | `build` | convert the Markdown files under `dir`, recursively (in v0.6.0); see below |
| `-P`, `--parallel n` | `build` | decks built at once (in v0.6.0): 5 by default, at most the number of cores; 1 builds them one after the other |
| `-w`, `--watch` | `build` | build again when a deck, a theme, kinds or data file or the configuration changes, and reload the open pages (in v0.6.0); see below |
| `-s`, `--server dir` | `build` | serve a directory, converting its decks on request (in v0.6.0); see below |
| `-p`, `--preview` | `build` | open the result in the default browser (in v0.6.0): the HTML pages with `--watch` (implied), the server with `--server` |
| `--host addr` | `build` | the address `--server` listens on, `127.0.0.1` by default (in v0.6.0) |
| `--template bespoke\|bare` | `build` | the HTML page's template (in v0.6.0): `bespoke`, the default, with the presenting script, on-screen controller and transitions, or `bare`, the slides stacked (`--bare`) |
| `--bespoke.osc`, `--bespoke.progress`, `--bespoke.transition` | `build` | the bespoke template's options, with marp-cli's defaults (in v0.6.0): controller on, **progress bar off** (it was always drawn before), transitions on; `--bespoke.osc=false` turns one off |
| `-c`, `--config-file file`, `--no-config-file` | `build` | marp-cli's configuration file, JSON or YAML (in v0.6.0); see below. `--config` is `slides`' own file |
| `--browser`, `--browser-path`, `--browser-protocol`, `--browser-timeout` | `build` | accepted and ignored with a warning (in v0.6.0): there is no browser |
| `--pptx-editable`, `--no-parallel` | `build` | marp-cli's spellings of `--pptx` and `-P 1` (in v0.6.0) |
| `--config file` | all | configuration file; by default `slides.yaml` next to the deck, if there is one |
| `--no-config` | all | do not read a configuration file |
| `--theme-set path` | all | a theme CSS file, or a directory of them (repeatable) |
| `--kinds file` | all | a YAML file of kinds (repeatable) |
| `--data file` | all | a YAML or JSON file of template data (repeatable) |
| `--set key=value` | all | one template value (repeatable) |
| `--no-template` | all | do not expand Go templates |
| `--strict-yaml` | all | parse directives as strict YAML, as plain Marpit does (Marp Core parses them loosely) |
| `--no-emoji` | all | leave `:shortcodes:` and Unicode emoji as text instead of drawing them as images (since v0.4.0: [Emoji]({{< relref "emoji" >}})) |
| `--html` | `build` | allow all raw HTML; by default, Marp Core's allowlist |
| `--lang code` | `build` | the page's `lang` when the deck sets none |
| `--bare` | `build` | stack the slides, without the presenting script |
| `--pdf` | `build` | write a PDF, one page per slide, with an outline — text and shapes as vectors (also chosen by `-o file.pdf`) |
| `--pdf-raster` | `build` | write the PDF with each page one image of the slide, as v0.3.0 did |
| `--notes` | `build` | write the speaker notes as a text file, slides separated by `---` as marp-cli does (in v0.6.0; also chosen by `-o file.txt`). The `notes` command prints them in another layout |
| `--pdf-notes` | `build` | add each slide's speaker notes to its PDF page as a note annotation, as marp-cli does (in v0.6.0). Asks for a PDF when nothing else says what to write |
| `--pdf-outlines`, `--pdf-outlines.pages`, `--pdf-outlines.headings` | `build` | the PDF's bookmarks (in v0.6.0): "Page N" per slide and the headings nested under it. On by default here, unlike marp-cli; `=false` turns one off. See below |
| `--images png\|jpeg` | `build` | write one image per slide: `deck.001.png`, … |
| `--image png\|jpeg` | `build` | write the first slide as an image (also chosen by `-o file.png` or `-o file.jpg`) |
| `--pptx` | `build` | write a PowerPoint file with its speaker notes (also chosen by `-o file.pptx`): editable since v0.5.0, one picture per slide in v0.4.0 |
| `--pptx-raster` | `build` | write a PowerPoint file of one full-slide picture per slide, as marp-cli does (since v0.5.0); `--image-scale` defaults to 2 |
| `--odp` | `build` | write an OpenDocument presentation with its speaker notes (also chosen by `-o file.odp`): editable since v0.5.0 |
| `--odp-raster` | `build` | write an OpenDocument presentation of one picture per slide (since v0.5.0) |
| `--jpeg-quality n` | `build` | JPEG quality, 1 to 100 (default 85) |
| `--image-scale n` | `build` | device pixels per CSS pixel for images and PDF (default 1); 2 gives 2560×1440 for 16:9, with text and vectors drawn at that density, not enlarged |
| `--allow-local-files` | `build` | let images and PDF read local files, from under the deck's directory only; refused by default, as with marp-cli |
| `--theme name\|file.css` | `build` | override the deck's theme (in v0.6.0): a theme name, or a CSS file, which becomes the theme whatever `@theme` it declares; a directory is refused, use `--theme-set` |
| `--title`, `--description`, `--author`, `--keywords`, `--url`, `--og-image` | `build` | the page's metadata (in v0.6.0), over the deck's `title`, `description`, `author`, `keywords`, `url` and `image` directives; see below |

## Several decks, and a directory

In v0.6.0 `build` takes several decks, a directory (the decks under it,
found recursively) or a pattern, and builds up to `--parallel` of them at once.
`-` is standard input; unlike marp-cli, a pipe is not read unless `-` says so.

A pattern is read as marp-cli's glob reads it (`**` and `{a,b}` in v0.7.0;
quote it so that the shell leaves it to `slides`): `*`, `?`, `[a-z]` and
`{md,txt}` within a name, `**` for any number of directories, so that
`'docs/**/*.md'` takes the decks of the tree. A wildcard does not reach a name
that starts with a dot, nothing is looked for in `node_modules`, case matters
(`*.md` does not find `UPPER.MD`), and whatever a pattern matches is converted,
whatever its extension. A pattern that matches nothing is an error.

- **`-I dir`** finds the Markdown files (`.md`, `.mdown`, `.markdown`,
  `.markdn`, in lower case) under `dir`, as marp-cli does: `node_modules` and
  dot directories are not looked in. With `-o out`, each deck is written at its
  path under `dir`, mirrored under `out`: `dir/sub/b.md` gives
  `out/sub/b.html` (`b.pdf`, `b.001.png`, `b.txt`, … whatever the flags say).
  Without `-o` the outputs go next to their decks. Naming decks as well as
  `-I` is an error, and so is `-o file` for several decks (`-o -` is allowed).
- **`-P n`** builds `n` decks at once. A deck's own rendering already uses the
  cores, so `n` is capped by their number; with `-o -` decks are built one after
  the other. The messages of a deck stay together, and a deck that fails does not
  stop the others: the exit status is 1 and the message lists the failures.

## Transitions and the presenting page

In v0.6.0 the `transition` directive of marp-cli works: `<!-- transition:
fade 1s -->`, a name and an optional duration, for the slide and the ones after
it, or `<!-- _transition: cover -->` for one. The built-in names are marp-cli's:
`clockwise`, `counterclockwise`, `cover`, `coverflow`, `cube`, `cylinder`,
`diamond`, `drop`, `explode`, `fade`, `fade-out`, `fall`, `flip`, `glow`,
`implode`, `in-out`, `iris-in`, `iris-out`, `melt`, `overlap`, `pivot`, `pull`,
`push`, `reveal`, `rotate`, `slide`, `star`, `swap`, `swipe`, `swoosh`, `wipe`,
`wiper` and `zoom`; `none` turns it off. The page plays it with the browser's
**View Transition API** when the viewer leaves the slide that has it, forwards,
and when it comes back to it; browsers without the API change slides as before,
and a viewer who asks for reduced motion gets a plain fade. A theme can bring
keyframes of its own, named `marp-transition-NAME` (or `marp-outgoing-` and
`marp-incoming-`, and `-backward`), which are used before the built-in ones.

The page has the **on-screen controller** — previous, page, next, full screen —
which shows when the pointer or a key moves and hides after a moment
(`--bespoke.osc=false` removes it), and, with `--bespoke.progress`, a progress
bar. **The progress bar is off by default now**, as in marp-cli; it was always
drawn before.

### Presenter view

In v0.7.0 the page has marp-cli's presenter view. Press `p`, or the last
button of the controller, and a window opens with the current slide, the step
after it, your speaker notes, the page, the time and a timer (click the timer
to reset it; `+` and `-` resize the notes; the bar between the slides and the
notes can be dragged). The two windows follow each other both ways: the
arrow keys, a click on the next step or on Previous/Next in either one moves the
other, a fragment at a time.

It works as marp-cli's does, read from its 4.5.1 sources: the window is
`?view=presenter&sync=KEY` of the same file, the next step is the same page as
`?view=next` in a frame, and the windows share their position through
`localStorage` (`bespoke-marp-sync-KEY`), so it needs no server and no
network, and does not work in a browser that has no `localStorage` (the
button says so). The notes are text: `&`, `<` and `>` are escaped, so a note
can never be markup. They are in the page, hidden, for the presenter view to
take, which keeps them out of what a screen reader reads to the audience.
`--bare` pages have none of it. The view adds about 10 KB to a page. There is
no overview view yet.

### Configuration files

marp-cli reads its options from a file in the working directory, and so does
`slides`, when the file is data: `.marprc` (YAML, JSON is YAML too),
`.marprc.json`, `.marprc.yaml`, `.marprc.yml`, `.config/marprc*`, or the
`marp` key of `package.json`; `--config-file` (`-c`) names any JSON or YAML
file, `marp.config.json` for one. Only the working directory is searched, not
its parents, as marp-cli's search does. A file of JavaScript or TypeScript
(`marp.config.js` ...) is a program, which `slides` does not run: it says so and
goes on. The options are those of marp-cli's documentation that mean something
here: `title`, `description`, `author`, `keywords`, `url`, `ogImage`, `lang`,
`theme`, `themeSet`, `html`, `allowLocalFiles`, `template`, `bespoke`,
`inputDir`, `output`, `parallel`, `pdf`, `pptx`, `notes`, `pdfNotes`,
`pdfOutlines`, `images`, `image`, `imageScale`, `jpegQuality`, `watch`,
`server`, `preview`; paths are relative to the file. The command line wins.
`engine`, `options` and `browser*` need Marp's engine and a browser, and are
ignored with a warning, as are the `--browser*` flags: scripts written for
marp-cli keep working. (`slides.yaml` and `slides.hcl`, with the kinds and the
template data, are separate and read as before.)

## Watch, serve and preview

In v0.6.0, with marp-cli's routes and live reload protocol (read from its
4.5.1 sources and compared with a running `marp -s`).

**`slides --server dir`** (`-s`) serves the directory on `http://127.0.0.1:8080/`
(the `PORT` environment variable changes the port; a port already in use is an
error that says so):

| Request | Answer |
| --- | --- |
| `/deck.md` | the deck as HTML, converted on request |
| `/deck.md?pdf`, `?pptx`, `?png`, `?jpg` (or `?jpeg`), `?notes` (or `?txt`) | the deck converted as the query asks; the PPTX is an attachment |
| a directory | its `index.md` or `PITCHME.md` as a page, else an index with a PDF and a PPTX link for each deck |
| any other file | the file |
| a deck that does not convert | 503 with the error |

A page opens a WebSocket to `/.__marp-cli-watch-notifier__/<id>`, where the id is
the SHA-256 of the deck's path; the server sends `ready`, then `reload` when the
deck, or anything under the directory, changes. A directory requested without
its trailing slash is redirected to it (marp-cli serves its index there, which
breaks the page's relative links).

**`--watch`** (`-w`) builds the decks, then builds again the one that changes
and, when a `--theme-set` file, a kinds or data file or the configuration
changes, all of them. The pages it writes open a socket on the loopback (port
37717, or the next free one) and reload when their deck is built again. A deck
that fails to build is reported and watched still. **`--preview`** (`-p`) opens
the default browser on the server, or on the HTML pages just built; there is no
window of its own, since `slides` has no browser to drive.

The server is a way into the machine, so it is stricter than marp-cli's, which
listens on every interface and checks nothing:

- it listens on **127.0.0.1**; `--host` changes that, with a warning, and a
  wildcard address cannot check the `Host` header any more;
- the **`Host`** header must be a loopback name (`localhost`, `127.0.0.1`,
  `[::1]`, `*.localhost`) or the host it listens on, which stops DNS rebinding;
- a WebSocket's **`Origin`** must be the server's own; the watch mode's sockets
  also take the `null` origin of a page opened from a file; a request from
  another site that is not a navigation (a script, an image, a `fetch`) is
  refused;
- a path is cleaned, its symbolic links are resolved, and the result must be
  inside the directory: `..` and a link out of it are refused, and so is
  `?pdf` of a deck outside; names that start with a dot (`.git`, `.env`) are not
  served;
- only `GET` and `HEAD`; a few conversions at once; at most 4096 directories
  watched, and none in `node_modules` or under a dot.

## Page metadata

The command line wins over the deck, as in marp-cli, where the flags are
appended to the deck as one last directive comment. Rules read from marp-cli
4.5.1 and pinned by tests against its output (in v0.6.0):

- A deck that sets no `title` takes the text of its first heading of the lowest
  level, as written (`# **Bold** title` gives `**Bold** title`). With no
  heading there is no `<title>`.
- `keywords` are split at commas, trimmed and deduplicated; a front-matter list
  is taken as it is. Like marp-cli, the page carries the `keywords` tag from
  the second keyword on, not for a single one.
- An invalid `url` (not absolute) is ignored with a warning, and the previous
  value kept; a valid one also writes `<link rel="canonical">` and `og:url`.
- `author` and `description` come with their Open Graph twins; `og:image:alt`
  and the large Twitter card need both a title and an image.

## Images, PDF, PPTX and ODP

`-o` with a `.pdf`, `.pptx`, `.odp`, `.png` or `.jpg` extension picks the format, as marp-cli
does. Output directories are created. A PDF has one page per slide, at the
slide's size, and an outline of the slides and their headings.

### PDF notes and outline

In v0.6.0 the PDF takes marp-cli's two options, and they are compared with
the PDFs marp-cli 4.5.1 writes for the same deck:

- `--pdf-notes` puts each slide's speaker notes on its page as a note
  annotation: a yellow Note icon at the bottom left whose contents are the
  notes (several comments joined by a blank line) and whose title is the deck's
  `author`. `--pdf-notes` alone asks for a PDF, as in marp-cli.
- The outline holds a bold bookmark "Page N" per slide and, under it, the
  slide's headings nested by level (`--pdf-outlines.pages` and
  `--pdf-outlines.headings` choose the two). **The outline is on by default
  here and off in marp-cli.** `slides` has always written one; it costs a few
  hundred bytes and makes a long deck navigable, so the default stays.
  `--pdf-outlines=false`, or marp-cli's `--no-pdf-outlines`, turns it off, and
  `=false` on `.pages` or `.headings` drops that part. A heading bookmark jumps
  to its slide; marp-cli also scrolls to the heading.

Rendering is done by [go-webengine](https://github.com/go-webengine/engine)
and the PDF is written by [go-pdfkit](https://github.com/go-pdfkit/pdfkit);
the fonts the built-in themes ask for (Lato, Roboto Mono) come from
[go-opentype/fonts](https://github.com/go-opentype/fonts) and are embedded, so
no render touches the network. There is no Chrome.

`--image-scale 2` renders at two device pixels per CSS pixel. Local images are
refused unless `--allow-local-files` is given, and even then only files under
the deck's directory are read: a deck from elsewhere cannot reach `~/.ssh`.
Several and split backgrounds (`![bg left:40%]`) render as Marp's layers do.

`--pptx` writes the deck with [go-ooxml](https://github.com/go-ooxml)'s
`WritePresentation`, with the speaker notes and the deck's metadata. Since
v0.5.0 the slides are editable (below); `--pptx-raster` writes one full-slide
picture per slide, as marp-cli does, at a scale that defaults to 2 that
`--image-scale` overrides. LibreOffice opens the result. Only LibreOffice was
tested; PowerPoint and Keynote were not.

`--odp` writes an OpenDocument presentation, through go-odf's
`WritePresentation`, with its speaker notes; `--odp-raster` writes one picture
per slide.

### Editable PPTX and ODP

Since v0.5.0 the slides of a PPTX or an ODP are **editable** — marp-cli's
PPTX is one picture a slide, and its experimental `--pptx-editable` converts
a PDF through LibreOffice. Here each slide is drawn from its layout, the one
the PNG and PDF exports draw:

- every line of text is a text box at its laid-out position, in the fonts
  the slide was laid out with — **embedded in the file**, so it looks the
  same on a machine without Lato or Inter (checked with LibreOffice 25.8) —
  with bold, italic, colours, letter spacing and clickable links;
- backgrounds, borders and rounded boxes are shapes; gradients are gradient
  fills in the PPTX and pictures in the ODP (an ODF gradient has two
  colours);
- pictures are stored once, however many slides show them;
- what a slide cannot express — an element under `opacity`, a `filter`, a
  blurred shadow, a formula — is a picture of that element, its text kept
  over it.

Rendered back by LibreOffice, a slide is within 0.04–2 % of the PNG export's
pixels. `--pptx-raster` and `--odp-raster` keep one picture a slide.

### Vector PDF

Since v0.4.0 the PDF is drawn as vectors: go-webengine paints the laid-out
slide into the PDF instead of into pixels. The text is text, in the fonts the
slide was laid out with, subset-embedded with a `/ToUnicode` map — it can be
selected, searched and copied. Backgrounds, borders and gradients are paths
and shadings, an image shown on many slides is embedded once, and links are
clickable. The 100-slide benchmark deck is 758 KiB, against 837 KiB for
marp-cli and 10 MiB for the image PDF of v0.3.0, and it is written in 0.3 s.

An element with no vector form here — under `opacity` < 1, with a `filter`
or a rotation, a blurred `box-shadow`, a gradient whose opacity fades — is
drawn as a bitmap of that element only, its text laid over it invisibly so it
stays selectable. Formulas are still bitmaps. `--pdf-raster` writes one image
per page, as v0.3.0 did.

## `slides.yaml`

A configuration file next to the deck applies to it, so a folder of decks
shares one:

```yaml
kinds: kinds.yaml        # a mapping of kinds, or file(s) holding one
themes:                  # theme CSS files or directories
  - themes/
data:                    # template data
  company: Example Inc.
html: false              # allow all raw HTML
lang: en
```

Kinds: the deck's own `kinds:` override the file's and the command line's.
Data: the file's and the command line's `data` override the deck's front
matter, so one deck can be built with different data.

## Planned

Not built yet: formulas as vectors in the PDF. When they arrive, they are a
command or flag here.
