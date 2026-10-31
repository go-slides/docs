---
title: "Documentação do go-slides"
linkTitle: "Início"
type: docs
cascade:
  type: docs
description: "Apresentações compatíveis com Marp em Go puro (CGO=0), com interpolação de templates Go e kinds de slide."
---

**`slides` transforma uma apresentação Marp em uma apresentação pronta para
exibir, em Go puro.** Ele lê o Markdown que o [Marp](https://marp.app) lê: as
diretivas do Marpit, `---` entre os slides, os fundos `![bg]`, as notas do
apresentador e os temas `default`, `gaia` e `uncover`. E acrescenta duas coisas
que o Marp não tem:

- **[Templates Go]({{< relref "templates" >}})** dentro da apresentação,
  avaliados antes de o Markdown ser analisado, de modo que uma apresentação
  pode gerar os próprios slides;
- **[kinds de slide]({{< relref "kinds" >}})** (o tipo de um slide): um slide declara o que ele é
  (`title`, `section`…) em vez de repetir a sua aparência.

{{< callout type="warning" >}}
**v0.14.0.** Estas páginas descrevem
[go-slides/slides v0.14.0](https://github.com/go-slides/slides/releases/tag/v0.14.0)
e são publicadas em `/docs/0.14/`.
{{< /callout >}}

## O que ele faz hoje {#what-it-does-today}

| | |
| --- | --- |
| `slides build` (o padrão) | a apresentação como uma única página HTML: cada slide é uma `<section>` emoldurada em SVG que se ajusta à janela, o CSS do tema e um visualizador com teclado |
| `slides build --pdf`, `--images png`, `--image jpeg` | a apresentação como PDF (uma página por slide, com sumário) ou como imagens PNG/JPEG, sem navegador: o [go-webengine](https://github.com/go-webengine) renderiza e o [go-pdfkit](https://github.com/go-pdfkit) escreve o PDF, em vetores desde a v0.4.0 e com o texto selecionável: [PDF vetorial]({{< relref "cli#vector-pdf" >}}) |
| matemática e código | `$…$` e `$$…$$` compostos pelo go-tex/math, blocos de código destacados pelo go-rouge, em HTML, imagens e PDF: [Matemática e código]({{< relref "math-and-code" >}}) |
| emoji | `:smile:` e 😄 viram imagens Twemoji incorporadas, de modo que não precisam de rede e são desenhadas em todas as saídas: [Emoji]({{< relref "emoji" >}}) (desde a v0.4.0) |
| `slides build --pptx` | a apresentação como arquivo PowerPoint com as notas do apresentador e os metadados da apresentação: uma imagem por slide na v0.4.0, como no marp-cli; **editável** desde a v0.5.0 com caixas de texto, formas, imagens e fontes incorporadas ([PPTX e ODP editáveis]({{< relref "cli#editable-pptx-and-odp" >}})) |
| `slides build --odp` | o mesmo como apresentação OpenDocument (`.odp`) |
| títulos ajustados | `# <!-- fit --> Title` ocupa toda a largura do slide em uma linha, na página e em todas as exportações |
| `slides expand` | a apresentação como Markdown Marp simples: templates avaliados, kinds substituídos pelas diretivas locais que eles representam |
| `slides notes` | as notas do apresentador |
| `slides kinds`, `slides themes` | o que uma apresentação pode nomear |
| `slides.yaml` | kinds, temas, dados, política de HTML e idioma compartilhados por uma pasta de apresentações |
| HCL | front matter `---hcl`, `slides.hcl`, arquivos de kinds e de dados em HCL2, com expressões e funções: [HCL]({{< relref "hcl" >}}) |

A diferença em relação ao
marp-cli está no caminho: ele renderiza imagens e PDF controlando um Chrome
sem interface; aqui isso é feito por bibliotecas Go.

## Páginas {#pages}

- [Primeiros passos]({{< relref "getting-started" >}})
- [Galeria]({{< relref "gallery" >}}): doze apresentações de exemplo, para abrir no Playground
- [Compatibilidade com o Marp]({{< relref "marp" >}})
- [Templates]({{< relref "templates" >}})
- [Kinds]({{< relref "kinds" >}})
- [HCL]({{< relref "hcl" >}})
- [Matemática e código]({{< relref "math-and-code" >}})
- [Emoji]({{< relref "emoji" >}})
- [Linha de comando]({{< relref "cli" >}})
