---
title: "Command line"
weight: 50
description: "Every command and flag of slides, and the slides.yaml configuration file."
tags: [cli]
---

{{< callout type="warning" >}}
No version is tagged yet. This page is checked against `slides -h` built from
`main` at `4ae5c49` (2026-10-18); anything here may change before the first
release.
{{< /callout >}}

```text
slides [command] [flags] deck.md
```

## Commands

| Command | What it does |
| --- | --- |
| `build` | writes the presentation as HTML. **The default**: `slides deck.md` is `slides build deck.md` |
| `expand` | prints the deck as plain Marp Markdown: templates expanded, kinds replaced by the directives they stand for |
| `notes` | prints the speaker notes |
| `kinds` | lists the kinds the deck can use and what they set |
| `themes` | lists the themes |
| `version` | prints the version |

`slides <command> -h` prints the flags of a command. **Flags go before the
deck** (`slides -o out.html deck.md`): Go's `flag` package stops at the first
argument that is not a flag. `-flag` and `--flag` are the same. `build`
writes `deck.html` beside `deck.md` unless `-o` says otherwise.

## Flags

| Flag | Commands | What it does |
| --- | --- | --- |
| `-o file` | all | output file; `-` for standard output |
| `--config file` | all | configuration file; by default `slides.yaml` next to the deck, if there is one |
| `--no-config` | all | do not read a configuration file |
| `--theme-set path` | all | a theme CSS file, or a directory of them (repeatable) |
| `--kinds file` | `build`, `expand`, `notes`, `kinds` | a YAML file of kinds (repeatable) |
| `--data file` | `build`, `expand`, `notes`, `kinds` | a YAML or JSON file of template data (repeatable) |
| `--set key=value` | `build`, `expand`, `notes`, `kinds` | one template value (repeatable) |
| `--no-template` | `build`, `expand`, `notes`, `kinds` | do not expand Go templates |
| `--strict-yaml` | `build`, `expand`, `notes`, `kinds` | parse directives as strict YAML, as plain Marpit does (Marp Core parses them loosely) |
| `--html` | `build` | allow all raw HTML; by default, Marp Core's allowlist |
| `--lang code` | `build` | the page's `lang` when the deck sets none |
| `--bare` | `build` | stack the slides, without the presenting script |

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

Not in the first release: image (PNG/JPEG), PDF and PPTX output, math and
code highlighting. When they arrive, they are commands or flags here.
