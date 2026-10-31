---
title: "Первые шаги"
weight: 10
description: "Первая презентация, собранная в HTML и развёрнутая обратно в Markdown для Marp."
tags: [start]
---

## Установка {#install}

Go 1.27.3 или новее:

```sh
go install github.com/go-slides/slides/cmd/slides@latest
```

`slides` — это один статический бинарный файл на Go без cgo: больше ничего
устанавливать не нужно, ни Node, ни Chrome. Эта страница описывает v0.16.0.

## Первая презентация {#a-first-deck}

Подойдёт любая презентация Marp. Сохраните это как `talk.md`:

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

Front matter — из Marp, плюс `data:`. Тело — шаблон Go: `range` пишет по слайду
на каждую тему, и каждый начинается с `---`, а первый слайд относится к
встроенному kind (типу слайда) `title` (без верхнего и нижнего колонтитулов, без номера
страницы).

## Сборка {#build-it}

```sh
slides talk.md                 # writes talk.html beside the deck
slides -o slides.html talk.md  # or somewhere else; -o - for standard output
```

`slides` без команды — это `slides build`. Флаги можно ставить до или после
презентации. Откройте HTML в браузере; между слайдами можно переходить с
клавиатуры.

## Экспорт изображений, PDF, PPTX и ODP {#export-images-pdf-pptx-and-odp}

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

Браузер не участвует: слайды рисует go-webengine, а PDF пишет go-pdfkit.
Локальные изображения отклоняются, пока не указан `--allow-local-files`.
Начиная с v0.5.0 слайды PPTX и ODP редактируемые — текстовые блоки, фигуры и
картинки, встроенные шрифты, — тогда как у marp-cli это картинки; `--pptx-raster`
сохраняет картинки ([Редактируемые PPTX и ODP]({{< relref "cli#editable-pptx-and-odp" >}})). Начиная с v0.4.0 PDF векторный, его текст
можно выделять (v0.3.0 писал по одному изображению на страницу). См. [Командная строка]({{< relref "cli#vector-pdf" >}}).

## Развёртывание {#expand-it}

```sh
slides expand talk.md
```

выводит ту же презентацию как **обычный Markdown для Marp** (вывод `main` на
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

Четыре слайда, которые породил `range`, нижний колонтитул с его значением и kind
`title`, замещённый точечными директивами, которые он обозначает. Всё, что
читает Marp, — marp-cli, расширение VS Code — может продолжить отсюда, так что
ничто из написанного с дополнениями go-slides не оказывается заперто в нём.

## Дальше {#next}

- [Совместимость с Marp]({{< relref "marp" >}}): что читается так же, как читает Marp.
- [Шаблоны]({{< relref "templates" >}}) и [Kinds]({{< relref "kinds" >}}): два дополнения.
- [Командная строка]({{< relref "cli" >}}): каждая команда и флаг.
