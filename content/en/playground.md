---
title: "Playground"
weight: 15
description: "Write a deck in your browser and watch the slides follow: go-slides compiled to WebAssembly, with completion, help and examples."
tags: [playground, browser]
---

**[Open the playground →](https://go-slides.github.io/playground/)**

The playground is go-slides itself, v0.9.0, compiled to WebAssembly and
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
- **Files**: the panel left of the editor is the deck's project, laid out
  the way the [command]({{< relref "cli" >}}) and the
  [VS Code extension]({{< relref "vscode" >}}) read it: `deck.md`,
  `slides.yaml`, `themes/`, `kinds/`, `media/`.
  - Click a file to edit it. The editor colours it by its language
    (Markdown, YAML, HCL, JSON, CSS) and marks its problems.
  - Right-click (or Shift+F10) for New File, New Folder, Rename (F2),
    Delete (asks first) and New theme / New kind. Drag a file onto a folder
    to move it. The keyboard follows the WAI-ARIA tree pattern, type-ahead
    included.
  - **Project ▸ New theme** makes `themes/my-gaia.css` (or from default,
    uncover): `@import` of the built-in theme and the rules you add. It
    joins the Theme list. **New kind** and **New data file** start from
    templates too. Each is listed in `slides.yaml`, so `slides build deck.md`
    finds it.
  - **Project ▸ Download project (.zip)** saves it all. Unzipped,
    `slides build deck.md` builds it. **Open project** takes a `.zip` or a
    folder back. An archive that leaves its folder (`..`, absolute paths),
    holds more than 256 files or 8 MiB in a file, or inflates more than 100
    times is refused.
  - Everything is kept in this browser with the draft. A project holds up
    to 64 MiB, and the line under the files says how much is used.
- **Media**: drop or paste an image, a video or a sound on the editor and it
  goes into `media/`, and the deck names it: `![](media/photo.png)`. Drop
  files on a folder of the panel to put them there. The deck stays small and
  readable, and nothing is uploaded. The preview's sandbox loads them
  without being opened up (the page hands it the bytes). **Save as ▸ HTML
  page** puts them in the page. `https:` images work too, but the browser
  fetches them from their server, which then sees your address. The
  preview sends no referrer.
- **Resizable panes**: the file panel, the editor, the slide thumbnails, the
  preview and the help are separated by bars you can drag (the pointer
  becomes a resize cursor over one). The pane beside the bar takes the room,
  the others keep their size. Double-click a bar, or press Enter on it, to
  pull the pane shut and bring it back. From the keyboard, F6 moves between
  the bars, the arrow keys move one by 16 px, Home and End take the pane
  before it to its smallest and largest, and Escape returns to the editor.
  The View tab still switches panes off. The sizes you set are kept in this
  browser, with the draft; the default sizes (190, 168 and 280 px, the editor
  and the preview sharing the rest 6 to 5) are never stored.
- **Highlighting** of the front matter's keys at any depth, in YAML (also
  in `{a: 1}` mappings) and in HCL.
- **Theme** picker: default, gaia, uncover, the project's themes
  (`themes/*.css`) and the themes the deck defines (`/* @theme name */` in a
  `<style>`). It writes `theme:` into the front
  matter, YAML or HCL, so a shared link keeps it, and it follows the front
  matter when you edit it.
- The preview **keeps its place** as you type: only the slides that
  changed are replaced, and a change of theme reopens the preview on the
  same slide.
- **Share link**: the deck travels in the link itself (in its `#` part,
  compressed), so opening the link opens the deck. Links carry up to
  64 KiB compressed, and the deck only: when the project holds more
  (themes, kinds, media), Share says so and offers the `.zip`.
- **Save as**:
  - the HTML page, the same as `slides deck.md` writes;
  - a **PDF** (vector, its text selectable, with the outline);
  - an editable **PowerPoint** or **OpenDocument** deck.

  They are made in your browser, by the same code as `slides --pdf`,
  `--pptx` and `--odp`: the 100-slide benchmark deck's PDF is
  byte-identical to the command's. The first export downloads the export
  engine (about 21 MiB, once; your browser keeps it until the next update),
  with a progress line. The 100-slide deck then takes about 3 seconds. An
  `https:` image is included only if its server allows it (CORS); `data:`
  images always are.
- **Your draft is kept**:
  - The deck is saved in this browser as you type (IndexedDB, for this site
    only) and comes back when you reload. Nothing is sent anywhere. The
    project's text files (themes, kinds, data, `slides.yaml`) are kept the
    same way, even when you close the tab right after typing; media are
    stored as they are dropped.
  - A shared link opens its deck and keeps your draft: **Draft ▸ Restore my
    draft** brings it back.
  - An example, **Draft ▸ Reset to the example** (which asks first) and
    another tab never erase it either: **Draft ▸ Restore replaced** swaps it
    back.
  - **Draft ▸ Forget my drafts** erases everything stored.

Every export renders the project the preview shows: its `slides.yaml`,
themes, kinds, data files and media, the media inside the document.

Images as files (slides `--images`) are not made in the browser: use the
[command]({{< relref "cli" >}}) or the
[VS Code extension]({{< relref "vscode" >}}).

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

The project's media reach the sandboxed frames as bytes in a message, and
each frame makes its own `blob:` URLs of them, in its own origin. A `blob:`
URL made by the playground does not load there, and the frames' policy
gains only `blob:` for media and fonts.

## Size and speed

The module is 8.8 MiB as GitHub Pages serves it (gzip; Pages does not
serve brotli). It is downloaded once with a progress bar, then cached by the
browser until the next deploy. On an Apple M4 Max with Chrome 155, measured
against the live page, the editor is ready 2.1 s after the page opens, 1.5 s
of which is the download, and the first slides appear at 2.4 s. A change to
the 100-slide benchmark deck reaches the preview in about 30 ms: go-slides
renders only the slides that changed, and only those are replaced in the
preview.
