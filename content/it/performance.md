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

go-slides e [marp-cli](https://github.com/marp-team/marp-cli) 4.5.1 sono stati
eseguiti come sottoprocessi sulle stesse presentazioni, su un Apple M4 Max (16 CPU) con un
carico medio a 1 minuto inferiore a 8, cinque round alternati dopo un riscaldamento.
Il tempo è la mediana del tempo reale; la memoria è il picco dell'intero albero di processi
(marp-cli pilota un Chrome, i cui processi vengono contati).

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

| Presentazione | go-slides | marp-cli | dimensione go-slides | dimensione marp-cli | memoria go-slides | memoria marp-cli |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| 100 slide | 348 ms | 7.795 ms | 760 KiB | 7,2 MiB | 204 MiB | 1.370 MiB |
| gaia, 20 slide | 112 ms | 4.608 ms | 893 KiB | 4,8 MiB | | |
| matematica e codice, 20 slide | 207 ms | 2.780 ms | 842 KiB | 1,5 MiB | | |

Per la presentazione da 100 slide, la v0.4.0 impiegava 4,7 s, 4,8 MiB e 540 MiB. Il PPTX
di marp-cli contiene immagini. Gran parte della dimensione di una presentazione modificabile sono i font incorporati.

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
