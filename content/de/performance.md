---
title: "Performance"
weight: 60
description: "Wie schnell go-slides einen Foliensatz baut und wie viel Speicher es dafür braucht, gemessen gegen marp-cli auf demselben Rechner."
tags: [performance]
---

{{< callout type="info" >}}
Die PPTX-Werte unten gelten für v0.5.0, die Werte für das Vektor-PDF für v0.4.0.
Die erste Tabelle gilt für v0.3.0,
die vorherige Version, die Folien nebenläufig rendert und ihre Schriften
nur einmal parst ([go-slides/slides#11](https://github.com/go-slides/slides/pull/11)).
{{< /callout >}}

Die Zeilen für das PDF und das PPTX mit 100 Folien in den drei Tabellen unten gelten seit v0.18.0 nicht mehr: siehe [Mehrere Kerne](#several-cores) und [go-slides/slides#117](https://github.com/go-slides/slides/issues/117).

go-slides und [marp-cli](https://github.com/marp-team/marp-cli) 4.5.1 wurden
als Subprozesse mit denselben Foliensätzen ausgeführt, auf einem Apple M4 Max (16 CPUs) mit
einem Load Average über eine Minute unter 8, fünf abwechselnd ausgeführte Durchläufe nach einem Warm-up.
Die Zeit ist der Median der Wanduhrzeit; der Speicher ist der Spitzenwert des gesamten Prozessbaums
(marp-cli steuert einen Chrome, dessen Prozesse mitzählen).

*Gemessen mit v0.3.0 auf allen Kernen (16 CPUs).*

| Foliensatz | Ausgabe | go-slides | marp-cli | Speicher go-slides | Speicher marp-cli |
| --- | --- | ---: | ---: | ---: | ---: |
| 100 Folien | HTML | 36 ms | 315 ms | 47 MiB | 140 MiB |
| 100 Folien | PDF | 1,14 s | 1,88 s | 242 MiB | 1.509 MiB |
| 100 Folien | PNG, jede Folie | 0,97 s | 5,44 s | 225 MiB | 1.354 MiB |
| 100 Folien | PPTX | 3,39 s | 8,03 s | 530 MiB | 1.373 MiB |
| gaia, 20 Folien | PDF | 283 ms | 1,57 s | 154 MiB | 1.315 MiB |
| Mathematik und Code, 20 Folien | PNG, jede Folie | 250 ms | 2,32 s | 168 MiB | 1.286 MiB |

## Vektor-PDF {#vector-pdf}

Ab v0.4.0 wird das PDF als Vektoren gezeichnet (siehe
[Kommandozeile]({{< relref "cli#vector-pdf" >}})). Gleicher Rechner, gleiche Methode,
2026-10-19:

*Gemessen mit v0.4.0 auf allen Kernen (16 CPUs), am 2026-10-19.*

| Foliensatz | go-slides | marp-cli | v0.3.0 | Größe go-slides | Größe marp-cli | Größe v0.3.0 |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| 100 Folien | 306 ms | 1.763 ms | 1.147 ms | 758 KiB | 837 KiB | 10,0 MiB |
| gaia, 20 Folien | 93 ms | 1.546 ms | 285 ms | 32 KiB | 66 KiB | 582 KiB |
| Mathematik und Code, 20 Folien | 140 ms | 1.643 ms | 294 ms | 124 KiB | 335 KiB | 638 KiB |

Speicher für das PDF mit 100 Folien: 158 MiB gegenüber 1.512 MiB bei marp-cli; CPU-Zeit
0,9 s gegenüber 2,7 s.

Die Werte von v0.4.0 für das PDF mit 100 Folien wurden am 2026-10-19 um 21:29 gemessen,
bei einem Load Average über eine Minute von 7,7, in drei Durchläufen: go-slides 306 ms, 758
KiB und 158 MiB; marp-cli 1.838 ms, 837 KiB und 1.512 MiB.

## Editierbares PPTX {#editable-pptx}

Seit v0.5.0 besteht das PPTX aus Textfeldern, Formen und Bildern (siehe
[Kommandozeile]({{< relref "cli#editable-pptx-and-odp" >}})). Gemessen am
2026-10-19 um 22:51, bei einem Load Average über eine Minute von 7,1, in drei Durchläufen:

*Gemessen mit v0.5.0 auf allen Kernen (16 CPUs), am 2026-10-19.*

| Foliensatz | go-slides | marp-cli | Größe go-slides | Größe marp-cli | Speicher go-slides | Speicher marp-cli |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| 100 Folien | 348 ms | 7.795 ms | 760 KiB | 7,2 MiB | 204 MiB | 1.370 MiB |
| gaia, 20 Folien | 112 ms | 4.608 ms | 893 KiB | 4,8 MiB | | |
| Mathematik und Code, 20 Folien | 207 ms | 2.780 ms | 842 KiB | 1,5 MiB | | |

Für den Foliensatz mit 100 Folien brauchte v0.4.0 4,7 s, 4,8 MiB und 540 MiB. Das PPTX
von marp-cli enthält Bilder. Den größten Teil der Größe eines editierbaren Foliensatzes machen die eingebetteten Schriften aus.

## Mehrere Kerne {#several-cores}

Gemessen mit v0.26.0: Wanduhrzeit in ms mit 1, 2, 4 und 16 Kernen, gegen marp-cli. Die Lastzahl über 1 Minute lag bei 3 bis 4. Die Zeit von marp-cli ändert sich nicht mit der Kernzahl, weil sein `-P` nur getrennte Dateien auf Browser verteilt. `slidesbench -cores` erzeugt diese Tabelle. Die Ausgabedateien sind byteidentisch mit denen von v0.25.0.

| Foliensatz, Modus | 1 Kern | 2 | 4 | 16 | marp-cli |
| --- | ---: | ---: | ---: | ---: | ---: |
| 100 Folien, PNG | 2629 | 1387 | 906 | 765 | 5181 |
| gaia, 20 Folien, PNG | 745 | 395 | 225 | 222 | 2316 |
| Mathematik und Code, 20 Folien, PNG | 590 | 331 | 200 | 192 | 2122 |
| gaia, 20 Folien, PDF | 87 | 65 | 57 | 59 | 1382 |
| Mathematik und Code, 20 Folien, PDF | 801 | 588 | 576 | 582 | 1447 |
| gaia, 20 Folien, PPTX | 98 | 77 | 70 | 70 | 3036 |
| Mathematik und Code, 20 Folien, PPTX | 911 | 625 | 593 | 605 | 2502 |
| 100 Folien, PDF | 3166 | 2573 | 2545 | 2568 | 1513 |
| 100 Folien, PPTX | 3299 | 2560 | 2547 | 2586 | 7238 |

**Bekannter Fehler.** Das PDF mit 100 Folien ist langsamer als marp-cli, und das PDF und das PPTX mit 100 Folien skalieren nicht mit den Kernen. Das ist seit v0.18.0 so: Die Engine baut die Instanzen der variablen Schrift für jede Seite neu auf, was auf einem Kern das 16-Fache an Speicher und das 3,5- bis 4-Fache an Zeit kostet. Der Fehler wird in [go-slides/slides#117](https://github.com/go-slides/slides/issues/117) und [go-webengine/engine#443](https://github.com/go-webengine/engine/issues/443) verfolgt und ist in v0.26.0 nicht behoben.

## Warum es schnell ist {#why-it-is-fast}

- **Kein Browser, der gestartet werden muss.** marp-cli startet für jeden Export nach PDF, Bild und
  PPTX einen Chrome; go-slides rendert mit
  [go-webengine](https://github.com/go-webengine/engine) im eigenen
  Prozess.
- **Folien werden parallel gerendert**, standardmäßig vier gleichzeitig, und jedes Bild wird
  geschrieben, sobald es an der Reihe ist: Ein langer Foliensatz liegt nie als
  Bilder im Speicher.
- **Schriften werden nur einmal pro Prozess geparst**, nicht einmal pro Folie.

## Was es nicht gewinnt {#what-it-does-not-win}

- **PDF-Größe in v0.3.0.** Jede Folie war ein Bild im PDF: 10 MiB gegenüber
  837 KiB beim Foliensatz mit 100 Folien. Das Vektor-PDF oben ist kleiner als das von
  marp-cli.
- **CPU-Zeit bei großen Foliensätzen.** Vier Worker verbrauchen mehr CPU, als die Wanduhrzeit
  zeigt: in v0.3.0 3,2 s gegenüber 2,7 s bei marp-cli für das PDF mit 100 Folien
  (0,9 s beim Vektor-PDF).
- **Der Speicher ist bei marp-cli eine Obergrenze**: Die Prozesse von Chrome teilen sich Seiten,
  die die Summe in jedem davon mitzählt. Sein größter einzelner Prozess hat 220–320 MiB.

Die Methode, jeder Foliensatz und jeder Modus sowie die Befehle zum Reproduzieren der Werte
stehen in
[BENCHMARKS.md](https://github.com/go-slides/slides/blob/main/BENCHMARKS.md).
