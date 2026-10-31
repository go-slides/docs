---
title: "Primeros pasos"
weight: 10
description: "Una primera presentación, construida como HTML y expandida de nuevo a Markdown de Marp."
tags: [start]
---

## Instalación {#install}

Go 1.27.3 o posterior:

```sh
go install github.com/go-slides/slides/cmd/slides@latest
```

`slides` es un único binario estático de Go, sin cgo: no hay nada más que
instalar, ni Node ni Chrome. Esta página describe v0.16.0.

## Una primera presentación {#a-first-deck}

Cualquier presentación de Marp sirve. Guarde esto como `talk.md`:

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

El front matter es el de Marp, más `data:`. El cuerpo es una plantilla de Go:
`range` escribe una diapositiva por tema, cada una empezando por `---`, y la
primera diapositiva es del kind (tipo de diapositiva) incorporado `title` (sin
encabezado, sin pie de página, sin número de página).

## Constrúyala {#build-it}

```sh
slides talk.md                 # writes talk.html beside the deck
slides -o slides.html talk.md  # or somewhere else; -o - for standard output
```

`slides` sin ningún comando equivale a `slides build`. Las opciones pueden ir
antes o después de la presentación. Abra el HTML en un navegador; el teclado
permite moverse entre diapositivas.

## Exportar imágenes, PDF, PPTX y ODP {#export-images-pdf-pptx-and-odp}

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

No interviene ningún navegador: go-webengine renderiza las diapositivas y
go-pdfkit escribe el PDF. Las imágenes locales se rechazan salvo que se indique
`--allow-local-files`. Desde v0.5.0 las diapositivas de un PPTX o un ODP son
editables (cuadros de texto, formas e imágenes, con las fuentes incrustadas),
mientras que las de marp-cli son imágenes; `--pptx-raster` conserva las imágenes
([PPTX y ODP editables]({{< relref "cli#editable-pptx-and-odp" >}})). Desde
v0.4.0 el PDF es vectorial y su texto se puede seleccionar (v0.3.0 escribía una
imagen por página). Véase [Línea de comandos]({{< relref "cli#vector-pdf" >}}).

## Expandirla {#expand-it}

```sh
slides expand talk.md
```

imprime la misma presentación como **Markdown de Marp plano** (salida de `main`
en `4ae5c49`, 2026-10-18):

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

Las cuatro diapositivas que generó `range`, el pie de página con su valor y el
kind `title` sustituido por las directivas locales que representa. Cualquier
herramienta que lea Marp (marp-cli, la extensión de VS Code) puede continuar a
partir de ahí, de modo que nada de lo escrito con las ampliaciones de go-slides
queda atrapado.

## Siguientes pasos {#next}

- [Compatibilidad con Marp]({{< relref "marp" >}}): qué se lee tal como lo lee Marp.
- [Plantillas]({{< relref "templates" >}}) y [Kinds]({{< relref "kinds" >}}): las dos ampliaciones.
- [Línea de comandos]({{< relref "cli" >}}): todos los comandos y opciones.
