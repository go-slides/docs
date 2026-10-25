---
title: "Erste Schritte"
weight: 10
description: "Ein erster Foliensatz, zu HTML gebaut und zurück zu Marp-Markdown expandiert."
tags: [start]
---

## Installation {#install}

Go 1.27.2 oder neuer:

```sh
go install github.com/go-slides/slides/cmd/slides@latest
```

`slides` ist ein einzelnes statisches Go-Binary ohne cgo: Sie müssen nichts weiter installieren, kein
Node, kein Chrome. Diese Seite beschreibt v0.13.0.

## Ein erster Foliensatz {#a-first-deck}

Jeder Marp-Foliensatz ist einer. Speichern Sie Folgendes als `talk.md`:

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

Das Front-Matter ist das von Marp, ergänzt um `data:`. Der Rumpf ist eine Go-Vorlage: Das
`range` schreibt pro Thema eine Folie, jeweils mit `---` eingeleitet, und die erste
Folie hat das eingebaute Kind (Folientyp) `title` (ohne Kopfzeile, ohne Fußzeile, ohne Seitenzahl).

## Bauen {#build-it}

```sh
slides talk.md                 # writes talk.html beside the deck
slides -o slides.html talk.md  # or somewhere else; -o - for standard output
```

`slides` ohne Befehl ist `slides build`. Die Flags dürfen vor oder nach dem
Foliensatz stehen. Öffnen Sie das HTML in einem Browser; mit der Tastatur wechseln Sie zwischen den Folien.

## Bilder, PDF, PPTX und ODP exportieren {#export-images-pdf-pptx-and-odp}

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

Es ist kein Browser beteiligt: go-webengine rendert die Folien, und go-pdfkit schreibt
das PDF. Lokale Bilder werden abgelehnt, solange `--allow-local-files` nicht angegeben ist.
Seit v0.5.0 sind die Folien eines PPTX oder ODP editierbar – Textfelder, Formen
und Bilder, eingebettete Schriften –, während sie bei marp-cli Bilder sind; `--pptx-raster`
behält Bilder bei ([Editierbares PPTX und ODP]({{< relref "cli#editable-pptx-and-odp" >}})). Seit v0.4.0 besteht das PDF aus Vektoren, sein Text ist
auswählbar (v0.3.0 schrieb ein Bild pro Seite). Siehe [Kommandozeile]({{< relref "cli#vector-pdf" >}}).

## Expandieren {#expand-it}

```sh
slides expand talk.md
```

gibt denselben Foliensatz als **reines Marp-Markdown** aus (Ausgabe von `main` bei
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

Die vier Folien, die das `range` erzeugt hat, die Fußzeile mit ihrem Wert und das
Kind `title`, ersetzt durch die Spot-Direktiven, für die es steht. Alles, was Marp liest – marp-cli, die VS-
Code-Erweiterung –, kann damit weiterarbeiten; nichts, was Sie mit den Erweiterungen von go-slides geschrieben haben, sperrt Sie ein.

## Weiter {#next}

- [Marp-Kompatibilität]({{< relref "marp" >}}): was so gelesen wird, wie Marp es liest.
- [Vorlagen]({{< relref "templates" >}}) und [Kinds]({{< relref "kinds" >}}): die beiden Erweiterungen.
- [Kommandozeile]({{< relref "cli" >}}): jeder Befehl und jedes Flag.
