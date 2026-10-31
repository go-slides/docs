# docs

[![docs](https://github.com/go-slides/docs/actions/workflows/docs.yml/badge.svg)](https://github.com/go-slides/docs/actions/workflows/docs.yml)
[![License](https://img.shields.io/badge/license-BSD--3--Clause-4F46E5?style=flat-square)](LICENSE)

**The go-slides documentation site**, published at
<https://go-slides.github.io/docs/>.

A [Hugo](https://gohugo.io/) site with the [tannevaled/hextra](https://github.com/tannevaled/hextra)
fork of the [Hextra](https://github.com/imfing/hextra) theme, pinned at its
**v0.24.1** tag, built the same way as the
[go-plot documentation](https://github.com/go-plot/docs), with go-slides'
branding. It documents [`go-slides/slides`](https://github.com/go-slides/slides).

## Versions

A version of the documentation is **the minor version of
[`go-slides/slides`](https://github.com/go-slides/slides) it describes**, or
`main` while no release is tagged:

| Source | Published under |
| --- | --- |
| branch `main`, with `params.slides.version` in `hugo.yaml` | `https://go-slides.github.io/docs/<MAJOR.MINOR>/`, or `/docs/main/` while it says `main` |
| the newest release (else `main`) | also `https://go-slides.github.io/docs/latest/` |

`https://go-slides.github.io/docs/` redirects to `latest/`. The pages describe
v0.16.0 (`params.slides.version`) and are published under `0.16/`. At the next
minor release: set `params.slides.version` to its tag and check every page
against that release (flags from its `slides -h`).

## Layout

| Path | What |
| --- | --- |
| `content/en/` | the pages; `_index.md` is the home, `weight` orders the sidebar |
| `i18n/en.yaml` | the site's strings, beside the theme's |
| `hugo.yaml` | `baseURL`, the slides version described, the navbar, the brand mounts |
| `layouts/_partials/custom/content-begin.html` | the outdated-version banner |
| `layouts/_partials/favicons.html` | the brand's favicons |
| `assets/css/custom.css` | the brand indigo as Hextra's primary colour, with its contrast measured |
| `themes/hextra` | submodule: tannevaled/hextra at tag v0.24.1; `.hextra-version` pins the oldest tag the site accepts, and CI (`tools/hextrapin`) fails on a submodule older than it, so a stale branch cannot move the theme back |
| `branding` | submodule: [go-slides/brand](https://github.com/go-slides/brand) |
| `content/<lang>/` | the translations (fr, de, es, zh, ko, ja), with the English file names; Chinese is Simplified (`zh-Hans`). Code, flags, directive names and file names stay as in English; each translated heading carries the English heading's `{#anchor}` |
| `tools/i18n` | a Go program (stdlib only, with tests): the CI fails when a page of `content/en` is missing in a language of `hugo.yaml` (or exists only in a translation), and warns when a translation changes code, shortcodes, link targets or anchors (`-strict` fails on those too); `-fix` adds the anchors, `-built public` compares the built site's heading ids with the English ones |
| `static/images/flags` | the language flags, lipis/flag-icons (MIT), for the first seven languages (`languageSwitch.flags` is false: the menu shows none) |
| languages | 12: English, then by script (Latin: de, es, fr, it, pl, pt; Cyrillic: ru, uk; CJK: ja, ko, zh) and by code within a script; each under its own name. Portuguese is Brazilian (`locale: pt-BR`, path `/pt/`), and has its own `i18n/pt.yaml` because Hextra's is European; Polish has one because Hextra has none. Russian and Polish dates are numeric (Hugo prints the nominative month) |
| `layouts/baseof.html` | Hextra's, with `lang` taken from the language's `locale` (`zh-Hans`) |
| `tools/gallery` | a Go program (stdlib only, with tests) that renders the gallery: it reads go-slides/slides' `examples/`, renders each deck's first slide with the slides command, and writes `data/gallery.json` and `static/gallery/` (both ignored by git) |
| `layouts/_shortcodes/gallery.html` | the gallery page's cards, from `data/gallery.json` |
| `scripts/publish-version.sh` | puts one build into a `gh-pages` checkout and rewrites `versions.json`, `latest` and the root redirect |

## Working locally

```sh
git clone --recurse-submodules https://github.com/go-slides/docs.git
cd docs
mkdir -p data && go run github.com/tannevaled/hextra/cmd/pagehistory@v0.21.0 > data/pagehistory.json
git clone https://github.com/go-slides/slides ../slides      # the gallery's decks
(cd tools/gallery && go run . -src ../../../slides -out ../..)   # Go 1.27.2 or later
hugo server --baseURL http://localhost:1313/docs/main/
```

Hugo 0.146 or later; no Python, no Node. The gallery page fails the build when
`data/gallery.json` is missing, which is the `go run` line above. CI does it
before every build, from the ref `params.gallery.ref` names.

## Publication

`.github/workflows/docs.yml`:

| Job | When | What |
| --- | --- | --- |
| `build` | pull requests, `main` | builds with `--panicOnWarning`, checks internal links and anchors (lychee, offline), uploads the site |
| `links` | pull requests, `main` | fetches every external URL the pages name |
| `deploy` | `main` | `scripts/publish-version.sh` into `gh-pages` (started empty on the first deploy), then a commit and a push; write access on this job only |

GitHub Pages serves the `gh-pages` branch.

## Licence

BSD-3-Clause.
