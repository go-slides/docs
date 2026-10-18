---
title: "Templates"
weight: 30
description: "Go text/template interpolation in a deck: values, generated slides, page numbers, includes."
tags: [templates]
---

The deck is a Go [`text/template`](https://pkg.go.dev/text/template),
evaluated **before** the Markdown is parsed. Everything the Go template
language has — actions, pipelines, `if`, `range`, `with`, variables — works
in a deck.

## Data

The template's dot holds the front matter's keys, then the keys of its
`data:` mapping, then the data given on the command line or in `slides.yaml`
— each overriding the one before.

```markdown
---
marp: true
title: Quarterly review
data:
  team: Platform
---

# {{ .title }}

Presented by the {{ .team }} team.
```

From the command line:

| Flag | Gives |
| --- | --- |
| `--data file.yaml` | a YAML or JSON file of data (repeatable) |
| `--set key=value` | one value (repeatable) |

## Generating slides

The whole body is one template, so a `range` that writes `---` writes slides:

```markdown
{{ range .releases }}
---

# {{ .version }}

{{ range .changes }}- {{ . }}
{{ end }}
{{ end }}
```

## Page and total

`{{ page }}` and `{{ total }}` are the slide's number and the number of
slides. They are known only once the deck has been split, so they are filled
in then — a generated deck counts the slides it generated.

```markdown
<!-- footer: "{{ page }} / {{ total }}" -->
```

## Include

`{{ include "parts/agenda.md" }}` inserts another file, relative to the
deck.

## Code is left alone

Fenced code blocks and inline code spans are set aside before the template
runs and restored verbatim, so a slide about Go templates can show
`{{ .Name }}` without escaping it.

## Turning it off

`--no-template` reads the deck as plain Marp: `{{` is text.
