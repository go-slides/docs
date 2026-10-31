---
title: "Pierwsze kroki"
weight: 10
description: "Pierwsza prezentacja, zbudowana do HTML i rozwinięta z powrotem do Markdownu Marp."
tags: [start]
---

## Instalacja {#install}

Go 1.27.2 lub nowsze:

```sh
go install github.com/go-slides/slides/cmd/slides@latest
```

`slides` to jeden statyczny plik wykonywalny Go bez cgo: nie trzeba instalować
niczego więcej, ani Node, ani Chrome. Ta strona opisuje v0.14.0.

## Pierwsza prezentacja {#a-first-deck}

Każda prezentacja Marp nią jest. Zapisz to jako `talk.md`:

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

Front matter jest z Marp, plus `data:`. Ciało to szablon Go: `range`
zapisuje jeden slajd na temat, każdy zaczynający się od `---`, a pierwszy
slajd ma wbudowany kind `title` (bez nagłówka, bez stopki, bez numeru strony).

## Zbuduj ją {#build-it}

```sh
slides talk.md                 # writes talk.html beside the deck
slides -o slides.html talk.md  # or somewhere else; -o - for standard output
```

`slides` bez polecenia to `slides build`. Flagi mogą stać przed prezentacją lub
po niej. Otwórz HTML w przeglądarce; klawiatura przełącza slajdy.

## Eksport obrazów, PDF, PPTX i ODP {#export-images-pdf-pptx-and-odp}

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

Nie bierze w tym udziału żadna przeglądarka: slajdy renderuje go-webengine, a PDF
zapisuje go-pdfkit. Obrazy lokalne są odrzucane, chyba że podano `--allow-local-files`.
Od v0.5.0 slajdy w PPTX lub ODP są edytowalne — pola tekstowe, kształty
i obrazy, osadzone czcionki — podczas gdy w marp-cli to obrazy; `--pptx-raster`
zachowuje obrazy ([Edytowalne PPTX i ODP]({{< relref "cli#editable-pptx-and-odp" >}})). Od v0.4.0 PDF jest wektorowy, z zaznaczalnym
tekstem (v0.3.0 zapisywała jeden obraz na stronę). Zobacz [Wiersz poleceń]({{< relref "cli#vector-pdf" >}}).

## Rozwiń ją {#expand-it}

```sh
slides expand talk.md
```

wypisuje tę samą prezentację jako **zwykły Markdown Marp** (wynik `main` w
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

Cztery slajdy wygenerowane przez `range`, stopka z jej wartością i kind
`title` zastąpiony dyrektywami punktowymi, które reprezentuje. Wszystko, co czyta Marp — marp-cli, rozszerzenie VS
Code — może stąd przejąć pracę, więc nic napisanego z użyciem dodatków go-slides
nie jest uwięzione.

## Dalej {#next}

- [Zgodność z Marp]({{< relref "marp" >}}): co jest czytane tak, jak czyta Marp.
- [Szablony]({{< relref "templates" >}}) i [Kinds]({{< relref "kinds" >}}): dwa dodatki.
- [Wiersz poleceń]({{< relref "cli" >}}): każde polecenie i flaga.
