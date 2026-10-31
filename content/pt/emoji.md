---
title: "Emoji"
weight: 36
description: "Shortcodes de emoji e emoji Unicode viram imagens Twemoji, sem rede: na página HTML, em imagens, PDF, PowerPoint e OpenDocument."
tags: [emoji]
---

{{< callout type="warning" >}}
Os emoji estão na v0.4.0; a v0.3.0 não tinha nenhum.
{{< /callout >}}

Escreva um emoji como shortcode ou como o próprio caractere:

```markdown
# Launch :rocket: 🎉

Shortcodes :smile: :+1: :heart_eyes: and Unicode 😄 👍🏽 👩‍💻 🇫🇷 1️⃣.
```

Cada um vira a imagem [Twemoji](https://github.com/jdecked/twemoji) que o Marp
Core desenha, com a marcação que o Marp Core escreve:

```html
<img class="emoji" draggable="false" alt="😄" src="data:image/svg+xml;base64,…" data-marp-twemoji=""/>
```

## Sem rede {#no-network}

O `src` do Marp Core é uma URL em um CDN, então uma apresentação exibida
offline, impressa ou transformada em PDF em uma máquina sem conexão mostra
imagens quebradas. O go-slides coloca o próprio SVG em uma URI `data:`. Os
gráficos (4.009 deles, Twemoji 17.0.3, Emoji 17.0) são incorporados ao programa
a partir do módulo em Go puro
[go-icons/twemoji](https://github.com/go-icons/twemoji), com 2 MB compactados e
lidos somente quando uma apresentação tem um emoji.

As imagens são desenhadas igualmente pela página HTML, por
`--images`/`--image`, por `--pdf`, por `--pptx` e por `--odp`: o go-webengine
rasteriza o SVG. Todos os 4.009 gráficos são analisados e desenhados pelo seu
rasterizador de SVG; um teste verifica isso. Desde a v0.5.0, esse rasterizador é
o [go-gfx/oksvg](https://github.com/go-gfx/oksvg) v0.2.0: `clipPath`,
`fill-rule: evenodd` e os comandos de arco agora são desenhados como o Chrome os
desenha, o que mudou um Twemoji e outros 211, e os logotipos SVG não ficam mais
embaralhados.

A regra que dimensiona um emoji ao texto (`img[data-marp-twemoji]`, um quadrado
de um `em`, alinhado à linha de base) é adicionada ao tema como o Marp Core a
adiciona, então um tema seu também a recebe.

## O que é convertido {#what-is-converted}

As regras são as do Marp Core; um teste as compara com o HTML que o Marp Core
4.4.0 escreve para uma apresentação com todos os casos.

| | |
| --- | --- |
| `:smile:`, `:+1:`, `:heart_eyes:` | os 1.903 nomes do markdown-it-emoji, em texto, títulos, links, tabelas, `header` e `footer`; não em código, URLs nem no texto alternativo de uma imagem |
| `😄`, `👩‍💻`, `🇫🇷`, `1️⃣`, `👍🏽` | emoji Unicode, em texto **e em código**: o Marp Core reescreve o HTML do código renderizado |
| `✔`, `©`, `™`, `▶` | continuam texto: caracteres cuja apresentação padrão é texto só são emoji com U+FE0F (`✔️`); `❤` é a exceção |
| `☺︎` | continua texto: um U+FE0E depois de um emoji pede o glifo simples |
| `12:30:45`, `:not_a_name:` | continuam como foram escritos |

O `id` de um título e o sumário usam o emoji, não o nome do shortcode:
`# Launch :rocket:` tem o id `launch-%F0%9F%9A%80`, como no Marp.

Onde o go-slides difere do Marp Core, ele converte mais: `✌️ ☝️ ✍️ 🏋️` e alguns
outros caracteres de apresentação padrão em texto que o emoji-regex do Marp Core
deixa como texto, e as poucas sequências que o Twemoji desenha e que o Unicode
não lista.

## Desativar {#turning-it-off}

`--no-emoji` (no código, `slides.Config{NoEmoji: true}`) deixa shortcodes e
emoji Unicode como o texto em que foram escritos, o que corresponde ao
`emoji: false` do Marp Core.

## Atribuição {#attribution}

Os gráficos são Twemoji, Copyright 2020 Twitter, Inc e outros colaboradores,
licenciados sob [CC-BY 4.0](https://creativecommons.org/licenses/by/4.0/) e
continuados pelo projeto [jdecked/twemoji](https://github.com/jdecked/twemoji).
Uma apresentação que os mostre deve dar o crédito, por exemplo em um rodapé ou
em um slide final. A tabela de shortcodes é a do markdown-it-emoji (MIT).
