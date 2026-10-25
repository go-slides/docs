---
title: "VS Code"
weight: 55
description: "The go-slides extension for VS Code: preview, presenting, problems, completion and export, as Marp for VS Code does for Marp."
tags: [vscode, editor]
---

{{< callout type="info" >}}
The extension is [go-slides/vscode](https://github.com/go-slides/vscode),
versioned on its own (v0.4.0). It renders with go-slides v0.9.0 (the
`editor` package, the source lines of scroll sync, and the overview and
presenter view). It is installed from the `go-slides-0.4.0.vsix` attached to
its GitHub release: its publication on the VS Code Marketplace and Open VSX is
prepared, but it is not there yet.
{{< /callout >}}

The extension is go-slides' counterpart of
[Marp for VS Code](https://github.com/marp-team/marp-vscode): the deck shows in
VS Code's Markdown preview as slides, its errors in the Problems panel, and
directives, kinds and template functions complete as you type.

## Install

Download `go-slides-<version>.vsix` from the
[releases](https://github.com/go-slides/vscode/releases), then run
*Extensions: Install from VSIX...*, or:

```sh
code --install-extension go-slides-0.4.0.vsix
```

The preview needs nothing else. Exporting needs the `slides` command:

```sh
go install github.com/go-slides/slides/cmd/slides@latest
```

## Preview

Open a deck and *Markdown: Open Preview to the Side*. A deck is a Markdown file
whose front matter says `slides: true`:

```markdown
---
slides: true
paginate: true
data:
  who: World
---

# Hello {{ .who }}

---

<!-- _kind: title -->
# Page {{ page }} of {{ total }}
```

Files with `marp: true` are previewed too, unless Marp for VS Code is installed:
both extensions would render the same document. The setting
`goSlides.preview.marpDocuments` decides: `auto` (the default), `always` or
`never`.

The preview is go-slides itself, compiled to WebAssembly and loaded by the
extension: no binary, and it works on [vscode.dev](https://vscode.dev). It reads
what the deck uses from the workspace — `slides.yaml` or `slides.hcl` next to
the deck, the [kinds]({{< relref "kinds" >}}) files they name, files a
[template]({{< relref "templates" >}}) includes, theme CSS — and only from the
deck's workspace folder.

## Scroll sync

Since v0.2.0 of the extension (go-slides v0.7.0), the editor and the
preview scroll together, both ways, as they do for plain Markdown: the
preview follows the slide and the block under the editor's cursor, and
scrolling the preview moves the editor. The settings are VS Code's own,
`markdown.preview.scrollPreviewWithEditor` and
`markdown.preview.scrollEditorWithPreview`.

Lines are traced through templates. A slide that a `{{ range }}` generates
follows the line of its template, and what an `{{ include }}` brings in
follows the include's line. A deck split by `headingDivider` syncs on its
headings.

## Presenting

Since v0.4.0 of the extension (go-slides v0.9.0), a deck can be presented
without leaving VS Code. With the deck in the editor, from the Command Palette
or the editor's `…` menu:

| Command | Opens |
|---|---|
| *go-slides: Open Overview* | a panel with every slide in a grid. A chosen slide shows its source in the editor; `Esc` closes the panel |
| *go-slides: Start Presenter View* | the slides, and beside them the presenter view: the next step, the speaker's notes and a timer |

The panels sit beside the editor and follow one another: change slide in one
and the others change with it. They also follow your edits to the deck. The `p`
key in the slides opens the presenter view.

To present on another screen, move the slides panel there with *View: Move
Editor into New Window*.

Each panel runs the deck in a sandboxed frame (no `allow-same-origin`) under a
policy of `default-src 'none'`: the frame loads images, media, styles and fonts
and nothing else, and runs only the scripts go-slides built into the page,
named by hash. A script the deck wrote does not run.

## Problems, completion, hover

- **Problems**: template, YAML and [HCL]({{< relref "hcl" >}}) errors, on the
  line they are on, in the deck or in the kinds file they come from.
- **Completion**: directive keys in the front matter and in `<!-- -->` comments
  (with `_` for a spot directive), the values of `kind` and `_kind` (with what
  each kind sets), `theme`, `size`, `paginate`, `math`, `headingDivider`; inside
  `{{ }}`, the template functions and the data fields.
- **Hover**: what a directive does, what a kind sets, a template function's
  arguments.

## Export

*go-slides: Export Slide Deck...* (also in the editor's title bar) writes HTML,
vector PDF, editable PPTX and ODP, PNG or JPEG with the `slides` command, as
`slides -o deck.pdf deck.md` would; see [Command line]({{< relref "cli" >}}).
It needs a trusted workspace. If `slides` is not found, the extension offers to
run the `go install` above in a terminal.

## Settings

| Setting | Default | |
|---|---|---|
| `goSlides.cliPath` | `""` | the `slides` command; empty: `slides` on `PATH`. A machine setting: a workspace cannot set it |
| `goSlides.exportType` | `pdf` | the format the export dialog proposes first |
| `goSlides.exportAutoOpen` | `true` | open the exported file |
| `goSlides.html` | `false` | allow all raw HTML instead of Marp's allowlist; honoured only in a trusted workspace |
| `goSlides.preview.marpDocuments` | `auto` | preview `marp: true` decks: `auto`, `always`, `never` |

## Security

The deck and the workspace are untrusted. Raw HTML follows Marp's allowlist
unless `goSlides.html` is on in a trusted workspace (a `slides.yaml` with
`html: true` does not turn it on in the preview), and VS Code's preview runs no
script from the deck. The preview reads files of the deck's workspace folder
only, follows no symbolic link, and templates see no environment variable.
Exporting runs a program, so it waits for workspace trust, and the program is
never named by the workspace.

## Performance

On go-slides' 100-slide benchmark deck, in VS Code 1.142 on an Apple M4 Max,
with Marp for VS Code 3.6.1 measured in the same session:

| | go-slides | Marp for VS Code |
|---|---:|---:|
| activation (loading the WebAssembly module) | 189 ms | — |
| first preview | 52 ms | 78 ms |
| preview after an edit (median) | **19 ms** | 23–24 ms |
| `.vsix` | 7.5 MB | 36.3 MB |

The figures are for extension v0.3.0 (go-slides v0.8.0) on a 100-slide deck,
typing one character into slide 50, from interleaved rounds under a 1-minute
load below 8. The first preview is faster than Marp's, and the preview after an
edit is now faster than Marp's too.

### In v0.8.0 and extension v0.3.0

go-slides v0.8.0 keeps each open deck's rendered slides and, after an edit,
re-renders only those whose source or inherited state changed (directives, page
number, total, heading ids, link references, source lines); the output is the
same as a full render, byte for byte. Typing in slide 50 of 100 re-renders one
slide: 2.5 ms natively, 11.5 ms in WebAssembly. VS Code's preview then morphs
only the changed slide's DOM, and the extension's smoke test types one
character into slide 50 between timed renders, for both extensions. The
extension's presenter view syncs through the parent window (`postMessage`).

### In v0.7.1 and extension v0.2.2

go-slides v0.7.1 caches parsed themes, packed style sheets and highlighted code
between renders, marks scroll-sync lines without going byte by byte, and reuses
its writers: a render of the 100-slide deck allocates 5.7 MB, against 52.7 MB
for extension v0.2.0. With it, extension v0.2.2 drops the `GOGC=400` override.
For comparison, extension v0.2.1 took 72 ms for the first preview and 44 ms
after an edit at the default `GOGC` (36 ms with `GOGC=400`), and v0.2.0 took
47–51 ms.

Extension v0.2.2 measured 52 ms for the first preview and 24 ms after an edit,
level with Marp's 23 ms.
