---
title: "Per iniziare"
weight: 10
description: "Una prima presentazione, compilata in HTML ed espansa di nuovo in Markdown Marp."
tags: [start]
---

## Installazione {#install}

Go 1.27.3 o successivo:

```sh
go install github.com/go-slides/slides/cmd/slides@latest
```

`slides` è un unico binario Go statico senza cgo: nient'altro da installare, niente
Node, niente Chrome. Questa pagina descrive la v0.16.0.

## Una prima presentazione {#a-first-deck}

Qualsiasi presentazione Marp lo è. Salva questo come `talk.md`:

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

Il front matter è quello di Marp, più `data:`. Il corpo è un template Go: il
`range` scrive una slide per argomento, ciascuna che inizia con `---`, e la prima
slide è del kind (il tipo di una slide) predefinito `title` (senza intestazione, senza piè di pagina, senza numero di pagina).

## Compilarla {#build-it}

```sh
slides talk.md                 # writes talk.html beside the deck
slides -o slides.html talk.md  # or somewhere else; -o - for standard output
```

`slides` senza comando equivale a `slides build`. I flag possono stare prima o dopo la
presentazione. Apri l'HTML in un browser; la tastiera permette di spostarsi tra le slide.

## Esportare immagini, PDF, PPTX e ODP {#export-images-pdf-pptx-and-odp}

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

Non interviene alcun browser: go-webengine rende le slide e go-pdfkit scrive
il PDF. Le immagini locali vengono rifiutate se non si passa `--allow-local-files`.
Dalla v0.5.0 le slide di un PPTX o ODP sono modificabili — caselle di testo, forme
e immagini, font incorporati — mentre quelle di marp-cli sono immagini; `--pptx-raster`
mantiene le immagini ([PPTX e ODP modificabili]({{< relref "cli#editable-pptx-and-odp" >}})). Dalla v0.4.0 il PDF è vettoriale, con testo
selezionabile (la v0.3.0 scriveva un'immagine per pagina). Vedi [Riga di comando]({{< relref "cli#vector-pdf" >}}).

## Espanderla {#expand-it}

```sh
slides expand talk.md
```

stampa la stessa presentazione come **semplice Markdown Marp** (output di `main` a
`4ae5c49`, 2026-10-18):

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

Le quattro slide generate dal `range`, il piè di pagina con il suo valore e il
kind `title` sostituito dalle direttive spot che rappresenta. Tutto ciò che legge Marp — marp-cli, l'estensione di VS
Code — può proseguire da qui, quindi nulla di ciò che è scritto con le aggiunte di go-slides
resta bloccato.

## Prossimi passi {#next}

- [Compatibilità con Marp]({{< relref "marp" >}}): ciò che viene letto come lo legge Marp.
- [Template]({{< relref "templates" >}}) e [Kinds]({{< relref "kinds" >}}): le due aggiunte.
- [Riga di comando]({{< relref "cli" >}}): ogni comando e flag.
