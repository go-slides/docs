# docs

[![docs](https://github.com/go-slides/docs/actions/workflows/docs.yml/badge.svg)](https://github.com/go-slides/docs/actions/workflows/docs.yml)
[![License](https://img.shields.io/badge/license-BSD--3--Clause-4F46E5?style=flat-square)](LICENSE)

**The go-slides documentation site**, published at
<https://go-slides.github.io/docs/>.

A [Hugo](https://gohugo.io/) site with the [tannevaled/hextra](https://github.com/tannevaled/hextra)
fork of the [Hextra](https://github.com/imfing/hextra) theme, pinned at its
**v0.16.0** tag, built the same way as the
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

`https://go-slides.github.io/docs/` redirects to `latest/`. Nothing is
released yet, so the pages describe the branch being written and are
published under `main/`. At the first release: set `params.slides.version` to
its tag, check every page against that release (flags from its `slides -h`),
and remove the "nothing is released" callouts.

## Layout

| Path | What |
| --- | --- |
| `content/en/` | the pages; `_index.md` is the home, `weight` orders the sidebar |
| `i18n/en.yaml` | the site's strings, beside the theme's |
| `hugo.yaml` | `baseURL`, the slides version described, the navbar, the brand mounts |
| `layouts/_partials/custom/version-select.html` | the version selector (reads `versions.json`), from go-plot/docs |
| `layouts/_partials/custom/content-begin.html` | the outdated-version banner |
| `layouts/_partials/favicons.html` | the brand's favicons |
| `assets/css/custom.css` | the brand indigo as Hextra's primary colour, with its contrast measured |
| `themes/hextra` | submodule: tannevaled/hextra at tag v0.16.0 |
| `branding` | submodule: [go-slides/brand](https://github.com/go-slides/brand) |
| `scripts/publish-version.sh` | puts one build into a `gh-pages` checkout and rewrites `versions.json`, `latest` and the root redirect |

## Working locally

```sh
git clone --recurse-submodules https://github.com/go-slides/docs.git
cd docs
mkdir -p data && sh themes/hextra/scripts/page-history.sh > data/pagehistory.json
hugo server --baseURL http://localhost:1313/docs/main/
```

Hugo 0.146 or later; no Python, no Node.

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
