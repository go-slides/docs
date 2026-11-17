---
title: "Rendimiento"
weight: 60
description: "Qué tan rápido construye go-slides una presentación y con cuánta memoria, medido frente a marp-cli en la misma máquina."
tags: [performance]
---

{{< callout type="info" >}}
Las cifras de PPTX de abajo corresponden a v0.5.0 y las del PDF vectorial a
v0.4.0. La primera tabla es de v0.3.0, la versión anterior, que renderiza las
diapositivas de forma concurrente y analiza sus fuentes una sola vez
([go-slides/slides#11](https://github.com/go-slides/slides/pull/11)).
{{< /callout >}}

Las filas del PDF y del PPTX de 100 diapositivas de las tres tablas siguientes se midieron antes de la v0.18.0. Se volvieron a medir con la v0.26.1, una vez corregida la regresión que se seguía en [go-slides/slides#117](https://github.com/go-slides/slides/issues/117): véase [Varios núcleos](#several-cores).

go-slides y [marp-cli](https://github.com/marp-team/marp-cli) 4.5.1 se ejecutaron
como subprocesos sobre las mismas presentaciones, en un Apple M4 Max (16 CPU) con
una carga media de 1 minuto inferior a 8, cinco rondas intercaladas tras un
calentamiento. El tiempo es la mediana del tiempo real; la memoria es el pico de
todo el árbol de procesos (marp-cli controla un Chrome, cuyos procesos cuentan).

*Medido con la v0.3.0 en todos los núcleos (16 CPU).*

| Presentación | Salida | go-slides | marp-cli | memoria de go-slides | memoria de marp-cli |
| --- | --- | ---: | ---: | ---: | ---: |
| 100 diapositivas | HTML | 36 ms | 315 ms | 47 MiB | 140 MiB |
| 100 diapositivas | PDF | 1,14 s | 1,88 s | 242 MiB | 1509 MiB |
| 100 diapositivas | PNG, todas las diapositivas | 0,97 s | 5,44 s | 225 MiB | 1354 MiB |
| 100 diapositivas | PPTX | 3,39 s | 8,03 s | 530 MiB | 1373 MiB |
| gaia, 20 diapositivas | PDF | 283 ms | 1,57 s | 154 MiB | 1315 MiB |
| matemáticas y código, 20 diapositivas | PNG, todas las diapositivas | 250 ms | 2,32 s | 168 MiB | 1286 MiB |

## PDF vectorial {#vector-pdf}

En v0.4.0 el PDF se dibuja en vectores (véase
[Línea de comandos]({{< relref "cli#vector-pdf" >}})). Misma máquina y mismo
método, 2026-10-19:

*Medido con la v0.4.0 en todos los núcleos (16 CPU), el 2026-10-19.*

| Presentación | go-slides | marp-cli | v0.3.0 | tamaño de go-slides | tamaño de marp-cli | tamaño de v0.3.0 |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| 100 diapositivas | 306 ms | 1763 ms | 1147 ms | 758 KiB | 837 KiB | 10,0 MiB |
| gaia, 20 diapositivas | 93 ms | 1546 ms | 285 ms | 32 KiB | 66 KiB | 582 KiB |
| matemáticas y código, 20 diapositivas | 140 ms | 1643 ms | 294 ms | 124 KiB | 335 KiB | 638 KiB |

Memoria del PDF de 100 diapositivas: 158 MiB, frente a 1512 MiB de marp-cli;
tiempo de CPU de 0,9 s frente a 2,7 s.

Las cifras de v0.4.0 para el PDF de 100 diapositivas se midieron el 2026-10-19 a
las 21:29, con una carga media de 1 minuto de 7,7, en tres rondas: go-slides
306 ms, 758 KiB y 158 MiB; marp-cli 1838 ms, 837 KiB y 1512 MiB.

## PPTX editable {#editable-pptx}

Desde v0.5.0 el PPTX se compone de cuadros de texto, formas e imágenes (véase
[Línea de comandos]({{< relref "cli#editable-pptx-and-odp" >}})). Medido el
2026-10-19 a las 22:51, con una carga media de 1 minuto de 7,1, en tres rondas:

*Medido con la v0.5.0 en todos los núcleos (16 CPU), el 2026-10-19.*

| Presentación | go-slides | marp-cli | tamaño de go-slides | tamaño de marp-cli | memoria de go-slides | memoria de marp-cli |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| 100 diapositivas | 348 ms | 7795 ms | 760 KiB | 7,2 MiB | 204 MiB | 1370 MiB |
| gaia, 20 diapositivas | 112 ms | 4608 ms | 893 KiB | 4,8 MiB | | |
| matemáticas y código, 20 diapositivas | 207 ms | 2780 ms | 842 KiB | 1,5 MiB | | |

Para la presentación de 100 diapositivas, v0.4.0 tardaba 4,7 s y ocupaba 4,8 MiB
y 540 MiB. El PPTX de marp-cli contiene imágenes. La mayor parte del tamaño de
una presentación editable son las fuentes incrustadas.

## Varios núcleos {#several-cores}

Medido con la v0.26.1 (go-webengine/engine v0.48.0): tiempo real en ms con 1, 2, 4 y 16 núcleos, frente a marp-cli. La carga media de 1 minuto era de 5,7 al empezar. El tiempo de marp-cli no cambia con el número de núcleos, porque su `-P` solo reparte archivos distintos entre navegadores. `slidesbench -cores` produce esta tabla. go-slides es más rápido que marp-cli en cada fila, con cada número de núcleos.

| Presentación, modo | 1 núcleo | 2 | 4 | 16 | marp-cli |
| --- | ---: | ---: | ---: | ---: | ---: |
| 100 diapositivas, PNG | 2628 | 1397 | 899 | 766 | 5197 |
| gaia, 20 diapositivas, PNG | 731 | 394 | 227 | 222 | 2331 |
| matemáticas y código, 20 diapositivas, PNG | 600 | 333 | 199 | 192 | 2140 |
| gaia, 20 diapositivas, PDF | 85 | 64 | 57 | 58 | 1411 |
| matemáticas y código, 20 diapositivas, PDF | 278 | 183 | 139 | 135 | 1457 |
| gaia, 20 diapositivas, PPTX | 99 | 77 | 69 | 70 | 3110 |
| matemáticas y código, 20 diapositivas, PPTX | 395 | 268 | 189 | 179 | 2527 |
| 100 diapositivas, PDF | 544 | 331 | 247 | 256 | 1534 |
| 100 diapositivas, PPTX | 692 | 435 | 317 | 276 | 7248 |

**Corregido en la v0.26.1.** De la v0.18.0 a la v0.26.0, el PDF de 100 diapositivas era más lento que marp-cli, y el PDF y el PPTX de 100 diapositivas no escalaban con los núcleos: el motor reconstruía las instancias de la fuente variable para cada página. go-webengine/engine v0.48.0 las comparte entre páginas ([go-webengine/engine#443](https://github.com/go-webengine/engine/issues/443)) y la v0.26.1 lo incorpora: el PDF de 100 diapositivas tarda 544 ms en un núcleo en vez de 3166 ms, con un pico de 169 MiB en vez de 2182 MiB, y pesa 792,4 KiB en vez de 1005,5 KiB, porque la fuente se incrusta 3 veces en vez de 102. Las páginas son idénticas píxel a píxel. El informe es [go-slides/slides#117](https://github.com/go-slides/slides/issues/117) (cerrado).

## Por qué es rápido {#why-it-is-fast}

- **No hay que arrancar ningún navegador.** marp-cli lanza Chrome en cada
  exportación a PDF, imagen o PPTX; go-slides renderiza con
  [go-webengine](https://github.com/go-webengine/engine) dentro de su propio
  proceso.
- **Las diapositivas se renderizan en paralelo**, cuatro a la vez por defecto, y
  cada imagen se escribe en cuanto le toca el turno: una presentación larga nunca
  se mantiene en memoria como imágenes.
- **Las fuentes se analizan una sola vez** por proceso, no una vez por
  diapositiva.

## Dónde no gana {#what-it-does-not-win}

- **Tamaño del PDF, en v0.3.0.** Cada diapositiva era una imagen en el PDF: 10 MiB
  frente a 837 KiB en la presentación de 100 diapositivas. El PDF vectorial de
  arriba es más pequeño que el de marp-cli.
- **Tiempo de CPU en presentaciones grandes.** Cuatro trabajadores gastan más CPU
  de lo que muestra el tiempo real: en v0.3.0, 3,2 s frente a los 2,7 s de
  marp-cli en el PDF de 100 diapositivas (0,9 s con el PDF vectorial).
- **La memoria de marp-cli es una cota superior**: los procesos de Chrome
  comparten páginas que la suma cuenta en cada uno. Su mayor proceso individual
  ocupa entre 220 y 320 MiB.

El método, cada presentación y modo, y los comandos para reproducir las cifras
están en
[BENCHMARKS.md](https://github.com/go-slides/slides/blob/main/BENCHMARKS.md).
