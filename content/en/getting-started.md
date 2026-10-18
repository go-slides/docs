---
title: "Getting started"
weight: 10
description: "A first deck, built to HTML and expanded back to Marp Markdown."
tags: [start]
---

## Install

**Not yet.** Nothing is released: the install line (`go install`) and the
binaries will be given here with the first release. `slides` is one static
Go binary with no cgo, so it will need nothing else installed — no Node, no
Chrome.

## A first deck

Any Marp deck is one. Save this as `talk.md`:

```markdown
---
marp: true
theme: default
paginate: true
footer: "{{ .talk }}"
data:
  talk: Pure Go slides
  topics: [Why, How, Next]
---

<!-- _kind: title -->
# {{ .talk }}

{{ range .topics }}
---

# {{ . }}
{{ end }}
```

The front matter is Marp's, plus `data:`. The body is a Go template: the
`range` writes one slide per topic, each starting with `---`, and the first
slide is of the built-in kind `title` (no header, no footer, no page number).

## Build it

```sh
slides talk.md            # writes the HTML presentation
slides talk.md -o -       # or to standard output
```

`slides` with no command is `slides build`. Open the HTML in a browser; the
keyboard moves between slides.

## Expand it

```sh
slides expand talk.md
```

prints the same deck as **plain Marp Markdown**: the four slides the `range`
generated, the footer with its value, and the `title` kind replaced by the
spot directives it stands for. Anything that reads Marp — marp-cli, the VS
Code extension — can take it from there, so nothing written with go-slides'
additions is locked in.

## Next

- [Marp compatibility]({{< relref "marp" >}}): what is read the way Marp reads it.
- [Templates]({{< relref "templates" >}}) and [Kinds]({{< relref "kinds" >}}): the two additions.
- [Command line]({{< relref "cli" >}}): every command and flag.
