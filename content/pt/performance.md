---
title: "Desempenho"
weight: 60
description: "Com que rapidez o go-slides constrói uma apresentação, em quanta memória, medido contra o marp-cli na mesma máquina."
tags: [performance]
---

{{< callout type="info" >}}
Os números do PPTX abaixo são da v0.5.0 e os do PDF vetorial, da v0.4.0. A
primeira tabela é da v0.3.0, a versão anterior, que renderiza os slides
concorrentemente e analisa as suas fontes uma única vez
([go-slides/slides#11](https://github.com/go-slides/slides/pull/11)).
{{< /callout >}}

O go-slides e o [marp-cli](https://github.com/marp-team/marp-cli) 4.5.1 foram
executados como subprocessos nas mesmas apresentações, em um Apple M4 Max (16
CPUs) com carga média de 1 minuto abaixo de 8, em cinco rodadas intercaladas
após um aquecimento. O tempo é a mediana do tempo real; a memória é o pico de
toda a árvore de processos (o marp-cli controla um Chrome, cujos processos
contam).

| Apresentação | Saída | go-slides | marp-cli | memória do go-slides | memória do marp-cli |
| --- | --- | ---: | ---: | ---: | ---: |
| 100 slides | HTML | 36 ms | 315 ms | 47 MiB | 140 MiB |
| 100 slides | PDF | 1,14 s | 1,88 s | 242 MiB | 1.509 MiB |
| 100 slides | PNG, todos os slides | 0,97 s | 5,44 s | 225 MiB | 1.354 MiB |
| 100 slides | PPTX | 3,39 s | 8,03 s | 530 MiB | 1.373 MiB |
| gaia, 20 slides | PDF | 283 ms | 1,57 s | 154 MiB | 1.315 MiB |
| matemática e código, 20 slides | PNG, todos os slides | 250 ms | 2,32 s | 168 MiB | 1.286 MiB |

## PDF vetorial {#vector-pdf}

Na v0.4.0, o PDF é desenhado como vetores (veja
[Linha de comando]({{< relref "cli#vector-pdf" >}})). Mesma máquina e mesmo
método, 2026-10-19:

| Apresentação | go-slides | marp-cli | v0.3.0 | tamanho do go-slides | tamanho do marp-cli | tamanho da v0.3.0 |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| 100 slides | 306 ms | 1.763 ms | 1.147 ms | 758 KiB | 837 KiB | 10,0 MiB |
| gaia, 20 slides | 93 ms | 1.546 ms | 285 ms | 32 KiB | 66 KiB | 582 KiB |
| matemática e código, 20 slides | 140 ms | 1.643 ms | 294 ms | 124 KiB | 335 KiB | 638 KiB |

Memória do PDF de 100 slides: 158 MiB, contra 1.512 MiB do marp-cli; tempo de
CPU de 0,9 s contra 2,7 s.

Os números da v0.4.0 para o PDF de 100 slides foram medidos em 2026-10-19 às
21:29, com carga média de 1 minuto de 7,7, em três rodadas: go-slides 306 ms,
758 KiB e 158 MiB; marp-cli 1.838 ms, 837 KiB e 1.512 MiB.

## PPTX editável {#editable-pptx}

Desde a v0.5.0, o PPTX é feito de caixas de texto, formas e imagens (veja
[Linha de comando]({{< relref "cli#editable-pptx-and-odp" >}})). Medido em
2026-10-19 às 22:51, com carga média de 1 minuto de 7,1, em três rodadas:

| Apresentação | go-slides | marp-cli | tamanho do go-slides | tamanho do marp-cli | memória do go-slides | memória do marp-cli |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| 100 slides | 348 ms | 7.795 ms | 760 KiB | 7,2 MiB | 204 MiB | 1.370 MiB |
| gaia, 20 slides | 112 ms | 4.608 ms | 893 KiB | 4,8 MiB | | |
| matemática e código, 20 slides | 207 ms | 2.780 ms | 842 KiB | 1,5 MiB | | |

Para a apresentação de 100 slides, a v0.4.0 levou 4,7 s, 4,8 MiB e 540 MiB. O
PPTX do marp-cli contém imagens. A maior parte do tamanho de uma apresentação
editável são as fontes incorporadas.

## Por que é rápido {#why-it-is-fast}

- **Nenhum navegador para iniciar.** O marp-cli inicia um Chrome para cada
  exportação de PDF, imagem e PPTX; o go-slides renderiza com o
  [go-webengine](https://github.com/go-webengine/engine) dentro do próprio
  processo.
- **Os slides são renderizados em paralelo**, quatro por vez por padrão, e cada
  imagem é escrita assim que chega a sua vez: uma apresentação longa nunca é
  mantida na memória como imagens.
- **As fontes são analisadas uma vez** por processo, não uma vez por slide.

## Onde ele não vence {#what-it-does-not-win}

- **Tamanho do PDF, na v0.3.0.** Cada slide era uma imagem no PDF: 10 MiB contra
  837 KiB para a apresentação de 100 slides. O PDF vetorial acima é menor que o
  do marp-cli.
- **Tempo de CPU em apresentações grandes.** Quatro workers gastam mais CPU do
  que o tempo real mostra: na v0.3.0, 3,2 s contra 2,7 s do marp-cli para o PDF
  de 100 slides (0,9 s para o PDF vetorial).
- **A memória é um limite superior para o marp-cli**: os processos do Chrome
  compartilham páginas que a soma conta em cada um. Seu maior processo
  individual tem 220–320 MiB.

O método, cada apresentação e modo, e os comandos para reproduzir os números
estão em
[BENCHMARKS.md](https://github.com/go-slides/slides/blob/main/BENCHMARKS.md).
