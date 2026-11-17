---
title: "Prestazioni"
weight: 60
description: "Quanto velocemente go-slides compila una presentazione, con quanta memoria, misurato contro marp-cli sulla stessa macchina."
tags: [performance]
---

{{< callout type="info" >}}
Le cifre PPTX qui sotto sono per la v0.5.0 e quelle del PDF vettoriale per la v0.4.0.
La prima tabella è per la v0.3.0,
la versione precedente, che rende le slide in modo concorrente e analizza i propri font
una sola volta ([go-slides/slides#11](https://github.com/go-slides/slides/pull/11)).
{{< /callout >}}

Le righe della PDF e del PPTX da 100 slide delle tre tabelle seguenti sono state misurate prima della v0.18.0. Sono state misurate di nuovo con la v0.26.1, una volta corretta la regressione tracciata in [go-slides/slides#117](https://github.com/go-slides/slides/issues/117): vedi [Più core](#several-cores).

go-slides e [marp-cli](https://github.com/marp-team/marp-cli) 4.5.1 sono stati
eseguiti come sottoprocessi sulle stesse presentazioni, su un Apple M4 Max (16 CPU) con un
carico medio a 1 minuto inferiore a 8, cinque round alternati dopo un riscaldamento.
Il tempo è la mediana del tempo reale; la memoria è il picco dell'intero albero di processi
(marp-cli pilota un Chrome, i cui processi vengono contati).

*Misurato con la v0.3.0 su tutti i core (16 CPU).*

| Presentazione | Output | go-slides | marp-cli | memoria go-slides | memoria marp-cli |
| --- | --- | ---: | ---: | ---: | ---: |
| 100 slide | HTML | 36 ms | 315 ms | 47 MiB | 140 MiB |
| 100 slide | PDF | 1,14 s | 1,88 s | 242 MiB | 1.509 MiB |
| 100 slide | PNG, ogni slide | 0,97 s | 5,44 s | 225 MiB | 1.354 MiB |
| 100 slide | PPTX | 3,39 s | 8,03 s | 530 MiB | 1.373 MiB |
| gaia, 20 slide | PDF | 283 ms | 1,57 s | 154 MiB | 1.315 MiB |
| matematica e codice, 20 slide | PNG, ogni slide | 250 ms | 2,32 s | 168 MiB | 1.286 MiB |

## PDF vettoriale {#vector-pdf}

Nella v0.4.0 il PDF è disegnato in vettoriale (vedi
[Riga di comando]({{< relref "cli#vector-pdf" >}})). Stessa macchina e stesso metodo,
19-10-2026:

*Misurato con la v0.4.0 su tutti i core (16 CPU), il 2026-10-19.*

| Presentazione | go-slides | marp-cli | v0.3.0 | dimensione go-slides | dimensione marp-cli | dimensione v0.3.0 |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| 100 slide | 306 ms | 1.763 ms | 1.147 ms | 758 KiB | 837 KiB | 10,0 MiB |
| gaia, 20 slide | 93 ms | 1.546 ms | 285 ms | 32 KiB | 66 KiB | 582 KiB |
| matematica e codice, 20 slide | 140 ms | 1.643 ms | 294 ms | 124 KiB | 335 KiB | 638 KiB |

Memoria per il PDF da 100 slide: 158 MiB, contro 1.512 MiB di marp-cli; tempo
CPU 0,9 s contro 2,7 s.

Le cifre della v0.4.0 per il PDF da 100 slide sono state misurate il 19-10-2026 alle 21:29,
con un carico medio a 1 minuto di 7,7, su tre round: go-slides 306 ms, 758
KiB e 158 MiB; marp-cli 1.838 ms, 837 KiB e 1.512 MiB.

## PPTX modificabile {#editable-pptx}

Dalla v0.5.0 il PPTX è fatto di caselle di testo, forme e immagini (vedi
[Riga di comando]({{< relref "cli#editable-pptx-and-odp" >}})). Misurato il
19-10-2026 alle 22:51, con un carico medio a 1 minuto di 7,1, su tre round:

*Misurato con la v0.5.0 su tutti i core (16 CPU), il 2026-10-19.*

| Presentazione | go-slides | marp-cli | dimensione go-slides | dimensione marp-cli | memoria go-slides | memoria marp-cli |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| 100 slide | 348 ms | 7.795 ms | 760 KiB | 7,2 MiB | 204 MiB | 1.370 MiB |
| gaia, 20 slide | 112 ms | 4.608 ms | 893 KiB | 4,8 MiB | | |
| matematica e codice, 20 slide | 207 ms | 2.780 ms | 842 KiB | 1,5 MiB | | |

Per la presentazione da 100 slide, la v0.4.0 impiegava 4,7 s, 4,8 MiB e 540 MiB. Il PPTX
di marp-cli contiene immagini. Gran parte della dimensione di una presentazione modificabile sono i font incorporati.

## Più core {#several-cores}

Misurato con la v0.26.1 (go-webengine/engine v0.48.0): tempo reale in ms con 1, 2, 4 e 16 core, contro marp-cli. Il carico medio a 1 minuto era 5,7 all'inizio. Il tempo di marp-cli non cambia con il numero di core, perché il suo `-P` si limita a distribuire file separati tra i browser. `slidesbench -cores` produce questa tabella. go-slides è più veloce di marp-cli in ogni riga, con ogni numero di core.

| Presentazione, modalità | 1 core | 2 | 4 | 16 | marp-cli |
| --- | ---: | ---: | ---: | ---: | ---: |
| 100 slide, PNG | 2628 | 1397 | 899 | 766 | 5197 |
| gaia, 20 slide, PNG | 731 | 394 | 227 | 222 | 2331 |
| matematica e codice, 20 slide, PNG | 600 | 333 | 199 | 192 | 2140 |
| gaia, 20 slide, PDF | 85 | 64 | 57 | 58 | 1411 |
| matematica e codice, 20 slide, PDF | 278 | 183 | 139 | 135 | 1457 |
| gaia, 20 slide, PPTX | 99 | 77 | 69 | 70 | 3110 |
| matematica e codice, 20 slide, PPTX | 395 | 268 | 189 | 179 | 2527 |
| 100 slide, PDF | 544 | 331 | 247 | 256 | 1534 |
| 100 slide, PPTX | 692 | 435 | 317 | 276 | 7248 |

**Corretto nella v0.26.1.** Dalla v0.18.0 alla v0.26.0 la PDF da 100 slide era più lenta di marp-cli, e la PDF e il PPTX da 100 slide non scalavano con i core: il motore ricostruiva le istanze del font variabile per ogni pagina. go-webengine/engine v0.48.0 le condivide tra le pagine ([go-webengine/engine#443](https://github.com/go-webengine/engine/issues/443)) e la v0.26.1 lo include: la PDF da 100 slide richiede 544 ms su un core invece di 3166 ms, con un picco di 169 MiB invece di 2182 MiB, e pesa 792,4 KiB invece di 1005,5 KiB, perché il font è incorporato 3 volte invece di 102. Le pagine sono identiche al pixel. La segnalazione è [go-slides/slides#117](https://github.com/go-slides/slides/issues/117) (chiusa).

## Perché è veloce {#why-it-is-fast}

- **Nessun browser da avviare.** marp-cli lancia Chrome per ogni esportazione in PDF, immagini e
  PPTX; go-slides esegue il rendering con
  [go-webengine](https://github.com/go-webengine/engine) all'interno del proprio
  processo.
- **Le slide vengono rese in parallelo**, quattro alla volta per impostazione predefinita, e ogni immagine viene
  scritta non appena è il suo turno: una presentazione lunga non è mai tenuta in memoria come
  immagini.
- **I font vengono analizzati una sola volta** per processo, non una volta per slide.

## Dove non vince {#what-it-does-not-win}

- **Dimensione del PDF, nella v0.3.0.** Ogni slide era un'immagine nel PDF: 10 MiB contro
  837 KiB per la presentazione da 100 slide. Il PDF vettoriale qui sopra è più piccolo di quello di
  marp-cli.
- **Tempo CPU sulle presentazioni grandi.** Quattro worker spendono più CPU di quanto il tempo reale
  mostri: nella v0.3.0, 3,2 s contro i 2,7 s di marp-cli per il PDF da 100 slide
  (0,9 s per il PDF vettoriale).
- **La memoria è un limite superiore per marp-cli**: i processi di Chrome condividono pagine
  che la somma conta in ciascuno. Il suo singolo processo più grande è di 220–320 MiB.

Il metodo, ogni presentazione e modalità, e i comandi per riprodurre le cifre
sono in
[BENCHMARKS.md](https://github.com/go-slides/slides/blob/main/BENCHMARKS.md).
