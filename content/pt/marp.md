---
title: "Compatibilidade com o Marp"
linkTitle: "Marp"
weight: 20
description: "O Markdown que o slides lê é o do Marpit e do Marp Core, não um dialeto dele."
tags: [marp]
---

Uma apresentação escrita para o [Marp](https://marp.app) é uma apresentação do
go-slides. As regras abaixo são as do [Marpit](https://marpit.marp.app) e do
[Marp Core](https://github.com/marp-team/marp-core); onde o go-slides difere,
esta página diz.

## Slides {#slides}

Uma régua horizontal — `---` em uma linha só dela, como o Marpit divide — inicia
um novo slide, e o mesmo vale para um título do nível que a diretiva
`headingDivider` indica. O front matter, entre os dois primeiros `---`, contém
as diretivas globais e os dados da apresentação.

## Diretivas {#directives}

| Tipo | Onde | Aplica-se a |
| --- | --- | --- |
| global | front matter, ou `<!-- theme: gaia -->` em qualquer lugar | a apresentação inteira: `theme`, `style`, `headingDivider`, `lang`… |
| local | `<!-- paginate: true -->` | este slide e os seguintes |
| pontual (spot) | `<!-- _class: lead -->` (sublinhado) | somente este slide |

As diretivas locais são as do Marpit: `paginate`, `header`, `footer`, `class`,
`backgroundColor`, `backgroundImage`, `backgroundPosition`,
`backgroundRepeat`, `backgroundSize`, `color`. As diretivas são analisadas como
YAML flexível, como o Marp Core faz; `--strict-yaml` as analisa como o Marpit
puro.

## Imagens {#images}

`![bg](image.jpg)` faz de uma imagem o fundo do slide, com as palavras-chave do
Marp: `bg contain`, `bg cover`, `bg 50%`, `bg left`/`bg right` para um fundo
dividido, várias imagens `bg` lado a lado, e os filtros de imagem (`blur`,
`sepia`, `grayscale`…). Os tamanhos de imagem em linha (`w:200px`, `h:`) são
lidos como o Marp os lê.

## Ao redor do slide {#around-the-slide}

- **Cabeçalho e rodapé**: as diretivas `header:` e `footer:`, em Markdown.
- **Paginação**: `paginate: true`.
- **Notas do apresentador**: um comentário HTML que não é uma diretiva é uma
  nota. `slides notes` as imprime.
- **Listas fragmentadas**: uma lista escrita com `*` (ou `1)`) aparece item por
  item, como no Marp.
- **`<style scoped>`**: CSS para um único slide.
- **HTML bruto**: a lista de permissões do Marp Core por padrão; `--html`
  permite todo ele.
- **Emoji**: `:shortcodes:` e emoji Unicode viram imagens Twemoji, como no Marp
  Core, mas incorporadas em vez de carregadas de um CDN: veja
  [Emoji]({{< relref "emoji" >}}).

## Temas {#themes}

`default`, `gaia` e `uncover` são embutidos, selecionados com `theme:`. Um tema
seu é um arquivo CSS com o comentário `/* @theme name */` do Marp, informado com
`--theme-set` ou listado em `themes:` no `slides.yaml`.

## Onde o go-slides difere {#where-go-slides-differs}

- **Templates.** `{{ }}` no corpo é uma ação de template Go (veja
  [Templates]({{< relref "templates" >}})). Uma apresentação que mostra `{{`
  como texto fora de um bloco de código precisa de `--no-template`, ou do texto
  escapado como os templates Go escapam. Blocos e trechos de código são
  deixados em paz.
- **Kinds.** `kind:` e `_kind:` são diretivas que o Marp não tem (veja
  [Kinds]({{< relref "kinds" >}})). `slides expand` as substitui pelas
  diretivas que elas representam, de modo que a apresentação expandida é Marp
  simples.
- **Saídas.** HTML, PNG, JPEG, PDF, PPTX e ODP, por meio de bibliotecas Go em
  vez de um navegador sem interface (veja
  [Linha de comando]({{< relref "cli#images-pdf-pptx-and-odp" >}}) para o
  limite conhecido). Na v0.4.0, como no marp-cli, um PPTX contém imagens; desde
  a v0.5.0 seus slides são caixas de texto e formas editáveis, e `--pptx-raster`
  mantém as imagens.
- **Títulos ajustados.** `# <!-- fit --> Title` é suportado: o título ocupa
  toda a largura do slide em uma linha, na página e em todas as exportações.
- **Matemática e código.** `$…$` e `$$…$$` seguem as regras de delimitadores do
  Marp Core e são compostos pelo go-tex/math; o código delimitado é destacado
  pelo go-rouge com nomes de classe do highlight.js (veja
  [Matemática e código]({{< relref "math-and-code" >}})).
