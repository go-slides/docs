---
title: "Theme packages"
weight: 42
description: "A theme as one file: a .gsta.xz archive with a theme.hcl, its CSS, fonts, assets and kinds, checked as untrusted input."
tags: [themes, hcl]
---

A **theme package** is one file, `<name>.gsta.xz` (Go-Slides Theme Archive): a
tar archive (ustar or pax) compressed with xz, with `theme.hcl` at its root and
the files it declares. It carries a theme's CSS, fonts, logos and backgrounds
and, if it likes, slide kinds, with a name, a version and a licence, so that a
theme can be passed around and installed. Theme packages come after v0.17.1.

## The theme.hcl file

```hcl
# theme.hcl
theme {
  name        = "acme"
  version     = "1.0.0"                 # semantic version
  description = "Example theme"
  authors     = ["go-slides contributors"]
  license     = "MIT"                   # SPDX expression
  homepage    = "https://go-slides.github.io"
  extends     = "default"               # the base theme, built in or installed
  min_slides  = "0.18.0"                # oldest go-slides it works with
}

css "theme.css" {}                      # in order; sha256 = "…" is optional

font "Roboto Mono" {
  weight  = 400                         # or "100 900"
  style   = "normal"
  file    = "fonts/RobotoMono-Regular.ttf"
  license = "OFL-1.1"                   # required for a font
}

asset "logo" { file = "assets/logo.svg" }

kind "acme-title" {                     # optional slide kinds
  class    = "lead"
  paginate = "false"
}
```

`theme.hcl` is HCL's native syntax with literal values only: no function, no
variable, no `file()` or `env()`. A name is lower-case letters, digits and
dashes, and is not `default`, `gaia` or `uncover`. `extends` names the base
theme; every `css` block adds a sheet, in order; every `font` and `asset` block
names a file of the package; every `kind` block is a [kind]({{< relref "kinds" >}}).
A `sha256` on a `css`, `font` or `asset` block is checked when the package is
read.

The CSS refers to fonts and assets by their path in the package, relative to
the CSS file: `background: url(assets/logo.svg)`. When the package is loaded,
each such reference becomes a `data:` URL, so that the page is self-contained,
and a `font` block becomes an `@font-face`. A package may also carry `LICENSE`,
`NOTICE` and `README` files; any other file that `theme.hcl` does not declare is
refused.

## Packing, installing and using

```
slides theme pack dir       write dir/theme.hcl and its files as <name>.gsta.xz
slides theme verify file    run every check on a package
slides theme install file   check it, then install it for the current user
slides theme list           the installed packages
slides theme remove name
```

Installed packages live in `go-slides/themes` under the user's configuration
directory (`os.UserConfigDir()`: `~/Library/Application Support`,
`$XDG_CONFIG_HOME` or `~/.config`, `%AppData%`). In a deck:

```markdown
---
theme: acme
---
```

finds an installed package, after the built-in themes and `--theme-set`.
`--theme acme.gsta.xz` and `--theme-set acme.gsta.xz` use a file
directly, without installing it; see [Command line]({{< relref "cli" >}}).

The kinds of a package apply with `--theme file.gsta.xz`, with `--theme name`
for an installed package, or with `theme: name` in a deck's front matter; the
deck's own kinds win over them.

## A package is untrusted input

A theme package can come from anywhere, so every reader (`pack`, `verify`,
`install`, `--theme`, a build) checks it the same way, and refuses with an error
that has a code (`theme.archive`, `theme.limit`, `theme.meta`, `theme.css`,
`theme.svg`, `theme.font`, `theme.checksum`, `theme.version`):

- **Archive.** One xz stream, whose dictionary is bounded and whose window
  grows with the data. In the tar: no `..`, absolute path, drive letter or
  backslash in a name; only regular files and directories (no symbolic or hard
  link, device or FIFO); no duplicate name (compared without case); only ustar
  and pax headers (no GNU extension, no sparse file), a pax header of at most
  8 KiB. Sizes are what is read, not what a header declares.
- **Bounds.** At most 16 MiB of archive, 256 entries, 8 MiB an entry, 24 MiB
  inflated, an xz dictionary of 16 MiB, an inflated stream of at most 100 times
  the archive's size (64 KiB or more), `theme.hcl` at most 64 KiB and 32
  brackets deep, 1 MiB of CSS a file, 4 MiB a font or an image.
- **CSS.** No `@import`, no `@import-theme` (use `extends`), no `image-set()`,
  `src()` or other function that loads a resource, no backslash escape outside a
  string, no `</style`. A `url()` is a `#fragment`, a `data:` image or font, or
  a declared file of the package, never one that leaves it.
- **SVG.** No `script`, `foreignObject`, frame, embed, `set`, DOCTYPE or event
  handler; an `href` is a `#fragment`; no `url()` that leaves the file.
- **Fonts and images** are parsed: a font must be a font, an image must match
  its extension, and each has a size bound.

The repository's `testdata/themes/acme` is a small example.
