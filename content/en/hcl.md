---
title: "HCL"
weight: 45
description: "HCL2 wherever a deck takes YAML: an HCL front matter, slides.hcl, kinds and data files, with expressions and functions."
tags: [hcl, configuration]
---

{{< callout type="warning" >}}
HCL is on `main` ([go-slides/slides#13](https://github.com/go-slides/slides/pull/13)),
not in v0.2.0: it arrives with the next release.
{{< /callout >}}

Everywhere `slides` reads YAML, it also reads the native syntax of
[HCL2](https://github.com/hashicorp/hcl/blob/main/hclsyntax/spec.md) — for
its expressions and its functions.

| Where | YAML | HCL |
| --- | --- | --- |
| front matter | `---` … `---` | `---hcl` … `---` |
| configuration | `slides.yaml` | `slides.hcl`, looked up first |
| kinds files | `--kinds kinds.yaml`, `kinds: kinds.yaml` | any file named `*.hcl` |
| data files | `--data data.yaml` or `.json` | any file named `*.hcl` |

Directive comments (`<!-- class: invert -->`) stay YAML, as in Marp, and the
deck body stays a [Go template]({{< relref "templates" >}}).

## A front matter in HCL

```hcl
---hcl
theme    = "gaia"
paginate = true
title    = "Quarterly review"
data = {
  company = "ACME"
  quarter = 3
}
header = format("%s — Q%d", data.company, data.quarter)
footer = "${upper(data.company)} · ${page()}/${total()}"
date   = formatdate("D MMMM YYYY", now())

kind "chapter" {
  extends = "section"
  class   = ["lead", "invert"]
}
---

# {{ .title }}
```

- **Attributes** are the YAML keys under the same names: Marp's directives
  (`theme`, `paginate`, `header`, `footer`, `class`, `backgroundColor`…),
  the metadata (`title`, `author`…), `template`, `delims`, `kinds` (file
  paths) and `data`, an object.
- **`kind "<name>" { … }`** blocks define [kinds]({{< relref "kinds" >}}),
  with the local directives and `extends`.
- **Values** reach the directives in the string forms YAML gives them —
  `true` is `"true"`, a list of classes is a list — so a directive behaves
  the same whichever syntax set it. `null` leaves a directive unset.
- **`data`** keeps its types: in templates an integral number is an `int`,
  other numbers a `float64`, and lists and objects stay lists and maps, so
  `{{ add .quarter 1 }}` works.
- An HCL value is final: Go templates do not run over it a second time.

## Variables

`data` is the template data: the front matter's own `data`, overlaid by
`slides.hcl`'s, by `--data` files and by `--set`. `"${data.company}"` works
wherever `{{ .company }}` would.

Front-matter attributes may reference `data.*` but **not each other**:
`footer = title` is an error. Spot directives (`_class`…) are not HCL
identifiers; write them in a comment.

## Functions

| | |
| --- | --- |
| `page()`, `total()` | the page number and the last page number, as `{{ page }}` and `{{ total }}`: a footer paginates |
| `now()`, `timestamp()` | the time, RFC 3339, local and UTC; `formatdate` formats it |
| `env(name)` | an environment variable, `""` when unset |
| `file(path)` | a UTF-8 text file of at most 1 MiB, relative to the deck or the configuration file |
| numbers | `abs ceil floor log max min parseint pow signum` |
| strings | `chomp format formatlist join lower regex regexall regexreplace replace split strlen strrev substr title trim trimprefix trimspace trimsuffix upper` |
| collections | `chunklist coalesce coalescelist compact concat contains distinct element flatten keys length lookup merge range reverse slice sort values zipmap` |
| encodings | `csvdecode jsondecode jsonencode` |
| dates | `formatdate timeadd` |

All but the first four rows are [go-cty](https://github.com/zclconf/go-cty)'s
standard library. None reaches the network.

## slides.hcl

```hcl
kinds  = "kinds.hcl"          # or a list of files
themes = ["themes/acme.css"]
data   = { company = "ACME" } # or file paths
html   = false
lang   = "en"

kind "quote" {
  extends = "blank"
  footer  = "${data.company}"
}
```

An unknown attribute is an error, not a silent no-op. A kinds file holds
`kind` blocks; its other attributes are ignored, so `slides.hcl` can serve as
one. A data file's attributes are the data's keys:

```hcl
# data.hcl
company = "ACME"
year    = 2026
motto   = trimspace(file("motto.txt"))
```

## Errors

Errors are HCL's diagnostics, positioned in the deck file itself:
`footer = data.compnay` on the deck's fourth line gives

```
deck.md:4:14: Unsupported attribute: This object does not have an attribute named "compnay".
```

## Marp and `slides expand`

Marp reads `---hcl` as the start of a front matter too — its front-matter
rule opens on any first line that begins with three dashes — and, the block
not being a YAML mapping, applies none of it: the deck still renders in Marp,
without its directives. A test pins this against Marp Core's own output.

`slides expand` writes the front matter back as YAML, so the expanded deck is
plain Marp and renders as `slides` renders the original:

```yaml
---
theme: "gaia"
paginate: true
title: "Quarterly review"
header: "ACME — Q3"
footer: "ACME · 1/2"
date: "19 October 2026"
---
```

## Licences

HCL support links [hashicorp/hcl](https://github.com/hashicorp/hcl)
(MPL-2.0, used unmodified as a library) and
[zclconf/go-cty](https://github.com/zclconf/go-cty) (MIT).
