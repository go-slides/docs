---
title: "Premiers pas"
weight: 10
description: "Une première présentation, générée en HTML puis développée en Markdown Marp."
tags: [start]
---

## Installation {#install}

Go 1.27.3 ou ultérieur :

```sh
go install github.com/go-slides/slides/cmd/slides@latest
```

`slides` est un binaire Go statique, sans cgo : rien d'autre à installer, ni
Node ni Chrome. Cette page décrit la v0.16.0.

## Une première présentation {#a-first-deck}

Toute présentation Marp en est une. Enregistrez ceci sous `talk.md` :

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

Le front matter est celui de Marp, plus `data:`. Le corps est un gabarit Go : le
`range` écrit une diapositive par sujet, chacune commençant par `---`, et la première
diapositive est du *kind* (le type d'une diapositive) intégré `title` (ni en-tête, ni pied de page, ni numéro de page).

## Génération {#build-it}

```sh
slides talk.md                 # writes talk.html beside the deck
slides -o slides.html talk.md  # or somewhere else; -o - for standard output
```

`slides` sans commande équivaut à `slides build`. Les options peuvent précéder ou suivre la
présentation. Ouvrez le HTML dans un navigateur ; le clavier permet de passer d'une diapositive à l'autre.

## Export en images, PDF, PPTX et ODP {#export-images-pdf-pptx-and-odp}

```sh
slides talk.md --pdf                  # one page per slide, with an outline
slides talk.md --images png           # talk.001.png, talk.002.png, …
slides talk.md --image jpeg           # the first slide only
slides talk.md -o out/slides.pdf      # the extension picks the format; out/ is created
slides talk.md --images jpeg --jpeg-quality 90
slides talk.md --images png --image-scale 2   # 2560×1440 for 16:9
slides talk.md --pdf --allow-local-files      # images from under the deck's directory
slides talk.md --pptx                         # or -o talk.pptx, with notes
slides talk.md --odp                          # or -o talk.odp: the same, as OpenDocument
```

Aucun navigateur n'intervient : go-webengine fait le rendu des diapositives et go-pdfkit écrit
le PDF. Les images locales sont refusées, sauf si `--allow-local-files` est indiqué.
Depuis la v0.5.0, les diapositives d'un PPTX ou d'un ODP sont modifiables — zones de texte, formes
et images, polices intégrées —, là où celles de marp-cli sont des images ; `--pptx-raster`
conserve les images ([PPTX et ODP modifiables]({{< relref "cli#editable-pptx-and-odp" >}})). Depuis la v0.4.0, le PDF est vectoriel et son texte
sélectionnable (la v0.3.0 écrivait une image par page). Voir [Ligne de commande]({{< relref "cli#vector-pdf" >}}).

## Développer la présentation {#expand-it}

```sh
slides expand talk.md
```

affiche la même présentation en **Markdown Marp ordinaire** (sortie de `main` au commit
`4ae5c49`, le 2026-10-18) :

```markdown
---
marp: true
theme: default
paginate: true
footer: "Pure Go slides"
---
<!--
_class: "lead"
_footer: ""
_paginate: "false"
-->


# Pure Go slides


---

# Why

---

# How

---

# Next
```

On y retrouve les quatre diapositives générées par le `range`, le pied de page avec sa valeur, et le
*kind* `title` remplacé par les directives ponctuelles qu'il représente. Tout outil qui lit Marp — marp-cli, l'extension VS
Code — peut prendre le relais : rien de ce qui est écrit avec les ajouts de go-slides
ne vous enferme.

## Pour aller plus loin {#next}

- [Compatibilité Marp]({{< relref "marp" >}}) : ce qui est lu comme Marp le lit.
- [Gabarits]({{< relref "templates" >}}) et [Kinds]({{< relref "kinds" >}}) : les deux ajouts.
- [Ligne de commande]({{< relref "cli" >}}) : toutes les commandes et toutes les options.
