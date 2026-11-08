---
title: "Gravações de terminal"
weight: 38
description: "Uma gravação do asciinema em uma apresentação: reproduzida em um pequeno player da página HTML e desenhada como texto vetorial no PDF, no PowerPoint e no OpenDocument."
tags: [asciinema, terminal]
---

{{< callout type="warning" >}}
As gravações de terminal estão no ramo principal depois da v0.17.1 e em nenhuma versão publicada.
{{< /callout >}}

Uma gravação do [asciinema](https://docs.asciinema.org/) ([asciicast v2](https://docs.asciinema.org/manual/asciicast/v2/) ou [v3](https://docs.asciinema.org/manual/asciicast/v3/)) é uma imagem cujo arquivo termina em `.cast`:

```markdown
![cast](demo.cast)
![The failing build](demo.cast#t=12.5)
```

## Sintaxe {#syntax}

O arquivo é lido da pasta da apresentação: uma URL, um caminho absoluto ou um caminho que sai da pasta é recusado. O texto alternativo nomeia a gravação para leitores de tela; `cast` sozinho significa o título da própria gravação. `#t=12.5`, ou `#t=1:23` em minutos e segundos, escolhe o momento que as saídas que não podem se mover mostram; sem ele, é o fim da gravação.

## O que cada saída mostra {#outputs}

- **HTML** (a página e a prévia do playground): a gravação é reproduzida em um pequeno player do go-slides embutido na página: reproduzir e pausar, um botão de velocidade e uma barra de posição. Ele não carrega script de lugar nenhum e não usa a rede.
- **PDF, PowerPoint e OpenDocument** (`--pdf`, `--pptx`, `--odp`): o terminal no momento que `#t=` indica, como texto vetorial na fonte monoespaçada do tema, com as cores de 16, 256 e diretas, negrito, itálico, sublinhado e texto invertido. O texto pode ser selecionado e pesquisado. Nunca é uma imagem. Negrito e itálico são desenhados quando a fonte monoespaçada do tema os tem (Roboto Mono no tema gaia).

## O player {#player}

Ele reproduz uma vez, na primeira vez que a gravação aparece na tela. Quando o navegador pede menos movimento (`prefers-reduced-motion`), ele começa pausado, mostrando o momento que `#t=` indica, e reproduz quando você aperta reproduzir. Seus botões e teclas não mudam o slide.

## Limites {#limits}

Uma gravação é uma entrada não confiável, por isso é limitada: o arquivo (16 MiB), os eventos (500.000), a duração (3 horas), o terminal (400 por 200 células) e o trabalho que o emulador pode fazer. Uma gravação acima de um limite, ou que não seja um arquivo asciicast bem formado, é recusada: a compilação continua, o slide mostra uma nota no lugar e o problema é um aviso com um código: `cast.syntax`, `cast.limit`, `cast.version` ou `cast.path`.

## No playground {#playground}

Adicione um arquivo `.cast` ao projeto pelo painel Arquivos, soltando-o ou pelo seletor de arquivos. Ele vai para `media/` e é inserido como imagem. É texto que a apresentação lê, então conta junto com os outros arquivos de texto do projeto, limitados a 2 MiB. Experimente no [playground](https://go-slides.github.io/playground/).
