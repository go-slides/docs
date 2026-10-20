---
title: "VS Code"
weight: 55
description: "The go-slides extension for VS Code: preview, problems, completion and export, as Marp for VS Code does for Marp."
tags: [vscode, editor]
---

{{< callout type="info" >}}
The extension is [go-slides/vscode](https://github.com/go-slides/vscode),
versioned on its own (v0.1.0). It renders with go-slides in v0.6.0 (the
`editor` package). It is installed from the `.vsix` attached to its GitHub
release: it is not on the VS Code Marketplace or Open VSX yet.
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
code --install-extension go-slides-0.1.0.vsix
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
| first preview | 85 ms | 78 ms |
| preview after an edit (median) | 47–51 ms | 24 ms |
| `.vsix` | 7.5 MB | 36.3 MB |

The preview after an edit is twice Marp's: a known defect, tracked in
[go-regexp/engine#10](https://github.com/go-regexp/engine/issues/10) (the
regular expressions behind code highlighting allocate most of what a render
allocates, and WebAssembly collects garbage on the preview's thread).
