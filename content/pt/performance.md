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

As linhas do PDF e do PPTX de 100 slides das três tabelas a seguir foram medidas antes da v0.18.0. Foram medidas de novo com a v0.26.1, depois de corrigida a regressão acompanhada em [go-slides/slides#117](https://github.com/go-slides/slides/issues/117): veja [Vários núcleos](#several-cores).

O go-slides e o [marp-cli](https://github.com/marp-team/marp-cli) 4.5.1 foram
executados como subprocessos nas mesmas apresentações, em um Apple M4 Max (16
CPUs) com carga média de 1 minuto abaixo de 8, em cinco rodadas intercaladas
após um aquecimento. O tempo é a mediana do tempo real; a memória é o pico de
toda a árvore de processos (o marp-cli controla um Chrome, cujos processos
contam).

*Medido com a v0.3.0 em todos os núcleos (16 CPUs).*

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

*Medido com a v0.4.0 em todos os núcleos (16 CPUs), em 2026-10-19.*

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

*Medido com a v0.5.0 em todos os núcleos (16 CPUs), em 2026-10-19.*

| Apresentação | go-slides | marp-cli | tamanho do go-slides | tamanho do marp-cli | memória do go-slides | memória do marp-cli |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| 100 slides | 348 ms | 7.795 ms | 760 KiB | 7,2 MiB | 204 MiB | 1.370 MiB |
| gaia, 20 slides | 112 ms | 4.608 ms | 893 KiB | 4,8 MiB | | |
| matemática e código, 20 slides | 207 ms | 2.780 ms | 842 KiB | 1,5 MiB | | |

Para a apresentação de 100 slides, a v0.4.0 levou 4,7 s, 4,8 MiB e 540 MiB. O
PPTX do marp-cli contém imagens. A maior parte do tamanho de uma apresentação
editável são as fontes incorporadas.

## Vários núcleos {#several-cores}

Medido com a v0.26.1 (go-webengine/engine v0.48.0): tempo real em ms com 1, 2, 4 e 16 núcleos, contra o marp-cli. A carga média de 1 minuto era de 5,7 no início. O tempo do marp-cli não muda com o número de núcleos, porque o seu `-P` só distribui arquivos separados entre navegadores. O `slidesbench -cores` produz esta tabela. O go-slides é mais rápido que o marp-cli em cada linha, com qualquer número de núcleos.

| Apresentação, modo | 1 núcleo | 2 | 4 | 16 | marp-cli |
| --- | ---: | ---: | ---: | ---: | ---: |
| 100 slides, PNG | 2628 | 1397 | 899 | 766 | 5197 |
| gaia, 20 slides, PNG | 731 | 394 | 227 | 222 | 2331 |
| matemática e código, 20 slides, PNG | 600 | 333 | 199 | 192 | 2140 |
| gaia, 20 slides, PDF | 85 | 64 | 57 | 58 | 1411 |
| matemática e código, 20 slides, PDF | 278 | 183 | 139 | 135 | 1457 |
| gaia, 20 slides, PPTX | 99 | 77 | 69 | 70 | 3110 |
| matemática e código, 20 slides, PPTX | 395 | 268 | 189 | 179 | 2527 |
| 100 slides, PDF | 544 | 331 | 247 | 256 | 1534 |
| 100 slides, PPTX | 692 | 435 | 317 | 276 | 7248 |

**Corrigido na v0.26.1.** Da v0.18.0 à v0.26.0, o PDF de 100 slides era mais lento que o marp-cli, e o PDF e o PPTX de 100 slides não escalavam com os núcleos: o motor reconstruía as instâncias da fonte variável para cada página. O go-webengine/engine v0.48.0 as compartilha entre as páginas ([go-webengine/engine#443](https://github.com/go-webengine/engine/issues/443)), e a v0.26.1 o incorpora: o PDF de 100 slides leva 544 ms em um núcleo em vez de 3166 ms, com um pico de 169 MiB em vez de 2182 MiB, e tem 792,4 KiB em vez de 1005,5 KiB, porque a fonte é incorporada 3 vezes em vez de 102. As páginas são idênticas pixel a pixel. O relato é [go-slides/slides#117](https://github.com/go-slides/slides/issues/117) (fechado).

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
