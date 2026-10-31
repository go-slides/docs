---
title: "Primeiros passos"
weight: 10
description: "Uma primeira apresentação, construída em HTML e expandida de volta para Markdown Marp."
tags: [start]
---

## Instalar {#install}

Go 1.27.2 ou posterior:

```sh
go install github.com/go-slides/slides/cmd/slides@latest
```

`slides` é um único binário Go estático, sem cgo: nada mais a instalar, sem
Node, sem Chrome. Esta página descreve a v0.14.0.

## Uma primeira apresentação {#a-first-deck}

Qualquer apresentação Marp serve. Salve isto como `talk.md`:

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

O front matter é o do Marp, mais `data:`. O corpo é um template Go: o `range`
escreve um slide por tópico, cada um começando com `---`, e o primeiro slide é
do kind (o tipo de slide) embutido `title` (sem cabeçalho, sem rodapé, sem número de página).

## Construir {#build-it}

```sh
slides talk.md                 # writes talk.html beside the deck
slides -o slides.html talk.md  # or somewhere else; -o - for standard output
```

`slides` sem comando é `slides build`. As flags podem vir antes ou depois da
apresentação. Abra o HTML em um navegador; o teclado navega entre os slides.

## Exportar imagens, PDF, PPTX e ODP {#export-images-pdf-pptx-and-odp}

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

Nenhum navegador é usado: o go-webengine renderiza os slides e o go-pdfkit
escreve o PDF. Imagens locais são recusadas, a menos que `--allow-local-files`
seja informado. Desde a v0.5.0, os slides de um PPTX ou ODP são editáveis
— caixas de texto, formas e imagens, com fontes incorporadas — enquanto os do
marp-cli são imagens; `--pptx-raster` mantém imagens
([PPTX e ODP editáveis]({{< relref "cli#editable-pptx-and-odp" >}})). Desde a
v0.4.0, o PDF é vetorial, com o texto selecionável (a v0.3.0 escrevia uma
imagem por página). Veja [Linha de comando]({{< relref "cli#vector-pdf" >}}).

## Expandir {#expand-it}

```sh
slides expand talk.md
```

imprime a mesma apresentação como **Markdown Marp simples** (saída de `main` em
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

Os quatro slides que o `range` gerou, o rodapé com o seu valor e o kind `title`
substituído pelas diretivas locais que ele representa. Qualquer coisa que leia
Marp — marp-cli, a extensão do VS Code — pode continuar a partir daí, então
nada escrito com os acréscimos do go-slides fica preso a ele.

## Próximos passos {#next}

- [Compatibilidade com o Marp]({{< relref "marp" >}}): o que é lido do mesmo jeito que o Marp lê.
- [Templates]({{< relref "templates" >}}) e [Kinds]({{< relref "kinds" >}}): os dois acréscimos.
- [Linha de comando]({{< relref "cli" >}}): todos os comandos e flags.
