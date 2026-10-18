---
title: "Kinds"
weight: 40
description: "A slide says what it is: kind names a preset of local directives."
tags: [kinds]
---

A **kind** is a named preset of local directives: header, footer,
pagination, class, background. A slide says what it is — `title`,
`section` — instead of repeating how it looks, and a deck's look is changed
in one place.

## Using a kind

```markdown
<!-- kind: section -->
# Part two

---

<!-- _kind: blank -->
![bg](photo.jpg)
```

- `kind:` applies like a local directive: to this slide and the ones after it.
- `_kind:` applies like a spot directive: to this slide only.
- Switching kinds first **clears** what the previous kind set, so a class or a
  footer never leaks from one kind into the next.
- Directives written on the slide itself **win** over the kind.

## Built in

| Kind | Sets |
| --- | --- |
| `title` | `class: lead`, no page number, no header, no footer |
| `section` | `class: lead` |
| `content` | nothing: the deck's own defaults |
| `blank` | no page number, no header, no footer |

Any of them can be redefined.

## Defining kinds

In the front matter:

```yaml
---
marp: true
kinds:
  chapter:
    class: invert
    backgroundColor: "#1e1b4b"
    paginate: false
  quote:
    extends: blank
    class: quote
---
```

`extends` starts a kind from another one's directives. A kind accepts the
local directives and `extends`.

## Sharing kinds

A team's kinds belong in one file every deck names:

```yaml
# kinds.yaml
chapter:
  class: invert
  paginate: false
```

```yaml
---
marp: true
kinds: kinds.yaml        # or a list of files
---
```

or `--kinds kinds.yaml` on the command line, or `kinds:` in `slides.yaml`.
The order is: the built-in kinds, then the command line and the
configuration file, then the deck's own `kinds:`.

`slides kinds talk.md` lists the kinds a deck can use and what each one sets.

## Leaving go-slides

`slides expand` writes every kind out as the spot directives it stands for,
slide by slide, so the result is plain Marp Markdown that any Marp tool
reads.
