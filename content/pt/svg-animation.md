---
title: "SVG animado"
weight: 40
description: "Uma imagem SVG com animações SMIL ou CSS: reproduzida pelo navegador em HTML, formas animadas no PowerPoint e no OpenDocument, um GIF animado onde estes não conseguem expressá-la, o estado inicial no PDF e nas imagens."
tags: [svg, animation]
---

Uma imagem SVG que se anima sozinha escreve-se como qualquer imagem:

```markdown
![](spinner.svg)
![w:400](chart.svg)
```

## O que é lido {#read}

As animações escritas no próprio arquivo SVG: SMIL (`animate`, `animateTransform`, `animateMotion`, `animateColor`, `set`, com `begin`, `dur`, `repeatCount`, `fill`, `calcMode`, `keyTimes` e `keySplines`) e CSS (`@keyframes` com as propriedades `animation`, `transform-origin` e `transform-box`). O que não é entendido, um evento como `begin="click"` por exemplo, fica de fora e é indicado por um aviso: nada é descartado em silêncio.

## O que cada saída mostra {#outputs}

- **HTML**: o navegador reproduz as animações por conta própria.
- **PowerPoint** (`--pptx`) e **OpenDocument** (`--odp`): cada elemento animado é uma forma própria, e suas animações começam com o slide: movimento, rotação e escala, opacidade, a mudança entre duas cores. O movimento de um elemento é amostrado nos instantes próprios da animação, de modo que as acelerações, os quadros-chave múltiplos e uma rotação em torno de outro ponto são exatos. Repetições (também `indefinite`), atrasos e inversões são mantidos.
- **Quadros**: quando uma animação não pode ser expressa assim (um caminho que se deforma, um traço que se desenha, uma inclinação), a imagem inteira é renderizada nos instantes próprios da animação e incorporada como um GIF animado, que o PowerPoint e o Impress reproduzem. É limitado: 12 quadros por segundo, no máximo 240 quadros, 16 milhões de pixels no conjunto.
- **PDF** (`--pdf`) e imagens (`--images`): não podem se animar. O desenho aparece em seu estado inicial, sem JavaScript nem `RichMedia`, e um aviso nomeia as imagens animadas.

Cada imagem animada recebe uma linha de aviso com o código `svg-animation`, dizendo como ela foi levada: de forma nativa, ou como quadros e por quê.

## Suporte por saída {#support}

| Recurso | HTML | PowerPoint | OpenDocument | PDF e imagens |
|---|---|---|---|---|
| Translação, movimento ao longo de um caminho | reproduzido | nativo | nativo | estado inicial |
| Rotação e escala, em torno de qualquer ponto | reproduzido | nativo | nativo | estado inicial |
| Raio de um círculo, tamanho de um retângulo (sem contorno) | reproduzido | nativo | nativo | estado inicial |
| Opacidade de um elemento de uma só forma | reproduzido | nativo | nativo | estado inicial |
| Cor de um preenchimento ou de um contorno, de uma cor a outra | reproduzido | nativo | nativo | estado inicial |
| Repetições, atrasos, inversões, manter no final | reproduzido | nativo | nativo | estado inicial |
| Acelerações, quadros-chave múltiplos | reproduzido | nativo (amostrado) | nativo (amostrado) | estado inicial |
| Degraus, mudanças discretas | reproduzido | quadros | quadros | estado inicial |
| Deformação de caminho, deslocamento de tracejado, atributos de posição e tamanho, gradientes | reproduzido | quadros | quadros | estado inicial |
| Inclinação, reflexão, vários relógios em um elemento, elementos animados aninhados | reproduzido | quadros | quadros | estado inicial |
| Eventos | reproduzido | não mantido (aviso) | não mantido (aviso) | estado inicial |

## Arquivos locais {#local-files}

Uma imagem local só é lida com `--allow-local-files`, a partir da pasta da apresentação, como no marp-cli. Sem a opção, a imagem fica fora das exportações, e a compilação avisa com `local-file-blocked`, nomeando o arquivo, seus slides e a opção. Um arquivo que não existe é indicado como `local-file-missing`.

## Limites {#limits}

Um SVG é uma entrada não confiável, por isso é limitado: 8 MiB de SVG, 50 000 elementos, 2 000 animações, 256 quadros-chave por animação, 60 segundos de linha do tempo e 240 quadros. Um desenho que passa de um limite é recusado com um aviso, e a compilação continua.

## O que foi verificado {#checked}

O LibreOffice Impress (macOS) abre os arquivos do PowerPoint e do OpenDocument e mantém cada animação ao salvá-los de novo, nos dois formatos. O Keynote (macOS) importa as formas e o GIF animado; suas animações não foram inspecionadas. PowerPoint: ainda não verificado.
