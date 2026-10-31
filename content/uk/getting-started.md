---
title: "Початок роботи"
weight: 10
description: "Перша презентація, зібрана в HTML і розгорнута назад у Markdown Marp."
tags: [start]
---

## Встановлення {#install}

Go 1.27.2 або новіший:

```sh
go install github.com/go-slides/slides/cmd/slides@latest
```

`slides` — це один статичний бінарний файл Go без cgo: нічого більше
встановлювати не треба, ні Node, ні Chrome. Ця сторінка описує v0.14.0.

## Перша презентація {#a-first-deck}

Будь-яка презентація Marp підходить. Збережіть це як `talk.md`:

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

Front matter — з Marp, плюс `data:`. Тіло — шаблон Go: `range` пише по одному
слайду на тему, кожен починається з `---`, а перший слайд має вбудований kind
`title` (без верхнього колонтитула, без нижнього, без номера сторінки).

## Збирання {#build-it}

```sh
slides talk.md                 # writes talk.html beside the deck
slides -o slides.html talk.md  # or somewhere else; -o - for standard output
```

`slides` без команди — це `slides build`. Прапорці можуть стояти перед
презентацією або після неї. Відкрийте HTML у браузері; клавіатура перемикає
слайди.

## Експорт зображень, PDF, PPTX і ODP {#export-images-pdf-pptx-and-odp}

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

Браузер не задіяно: go-webengine рендерить слайди, а go-pdfkit пише PDF.
Локальні зображення відхиляються, доки не вказано `--allow-local-files`.
Починаючи з v0.5.0 слайди PPTX або ODP редагуються — текстові блоки, фігури й
зображення, вбудовані шрифти — тоді як у marp-cli це картинки; `--pptx-raster`
залишає картинки ([Редагований PPTX і ODP]({{< relref "cli#editable-pptx-and-odp" >}})). Починаючи з v0.4.0 PDF векторний, його текст
можна виділяти (v0.3.0 писала одне зображення на сторінку). Див. [Командний рядок]({{< relref "cli#vector-pdf" >}}).

## Розгортання {#expand-it}

```sh
slides expand talk.md
```

виводить ту саму презентацію як **звичайний Markdown Marp** (вивід `main` на
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

Чотири слайди, згенеровані `range`, нижній колонтитул з його значенням і kind
`title`, замінений на spot-директиви, які він означає. Усе, що читає Marp —
marp-cli, розширення VS Code — може братися за це далі, тож нічого, написаного
з доповненнями go-slides, не заблоковано.

## Далі {#next}

- [Сумісність з Marp]({{< relref "marp" >}}): що читається так, як читає Marp.
- [Шаблони]({{< relref "templates" >}}) і [Kinds]({{< relref "kinds" >}}): два доповнення.
- [Командний рядок]({{< relref "cli" >}}): усі команди й прапорці.
