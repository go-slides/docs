---
title: "SVG animato"
weight: 40
description: "Un'immagine SVG con animazioni SMIL o CSS: riprodotta dal browser in HTML, forme animate in PowerPoint e OpenDocument, una GIF animata dove queste non possono esprimerla, lo stato iniziale nel PDF e nelle immagini."
tags: [svg, animation]
---

Un'immagine SVG che si anima da sola si scrive come qualsiasi immagine:

```markdown
![](spinner.svg)
![w:400](chart.svg)
```

## Che cosa viene letto {#read}

Le animazioni scritte nel file SVG stesso: SMIL (`animate`, `animateTransform`, `animateMotion`, `animateColor`, `set`, con `begin`, `dur`, `repeatCount`, `fill`, `calcMode`, `keyTimes` e `keySplines`) e CSS (`@keyframes` con le proprietà `animation`, `transform-origin` e `transform-box`). Ciò che non viene capito, per esempio un evento come `begin="click"`, viene lasciato fuori e segnalato con un avviso: nulla viene scartato in silenzio.

## Che cosa mostra ogni uscita {#outputs}

- **HTML**: il browser riproduce da solo le animazioni.
- **PowerPoint** (`--pptx`) e **OpenDocument** (`--odp`): ogni elemento animato è una forma a sé, e le sue animazioni partono con la diapositiva: movimento, rotazione e scala, opacità, un passaggio tra due colori. Il movimento di un elemento è campionato negli istanti propri dell'animazione, così le accelerazioni, i fotogrammi chiave multipli e una rotazione attorno a un altro punto sono esatti. Ripetizioni (anche `indefinite`), ritardi e inversioni vengono riportati.
- **Fotogrammi**: quando un'animazione non si può esprimere così (un tracciato che si deforma, un tratto che si disegna, un'inclinazione), l'intera immagine viene resa negli istanti propri dell'animazione e incorporata come GIF animata, che PowerPoint e Impress riproducono. È limitato: 12 fotogrammi al secondo, al massimo 240 fotogrammi, 16 milioni di pixel in totale.
- **PDF** (`--pdf`) e immagini (`--images`): non possono animarsi. Il disegno è mostrato nello stato iniziale, senza JavaScript né `RichMedia`, e un avviso nomina le immagini animate.

Ogni immagine animata riceve una riga di avviso con il codice `svg-animation`, che dice come è stata portata: in modo nativo, o come fotogrammi e perché.

## Supporto per uscita {#support}

| Funzione | HTML | PowerPoint | OpenDocument | PDF e immagini |
|---|---|---|---|---|
| Traslazione, movimento lungo un tracciato | riprodotto | nativo | nativo | stato iniziale |
| Rotazione e scala, attorno a qualsiasi punto | riprodotto | nativo | nativo | stato iniziale |
| Raggio di un cerchio, dimensione di un rettangolo (senza contorno) | riprodotto | nativo | nativo | stato iniziale |
| Opacità di un elemento di una sola forma | riprodotto | nativo | nativo | stato iniziale |
| Colore di un riempimento o di un contorno, da un colore a un altro | riprodotto | nativo | nativo | stato iniziale |
| Ripetizioni, ritardi, inversioni, mantenimento alla fine | riprodotto | nativo | nativo | stato iniziale |
| Accelerazioni, fotogrammi chiave multipli | riprodotto | nativo (campionato) | nativo (campionato) | stato iniziale |
| Gradini, cambi discreti | riprodotto | fotogrammi | fotogrammi | stato iniziale |
| Deformazione di tracciato, spostamento dei trattini, attributi di posizione e dimensione, sfumature | riprodotto | fotogrammi | fotogrammi | stato iniziale |
| Inclinazione, riflessione, più orologi su un elemento, elementi animati annidati | riprodotto | fotogrammi | fotogrammi | stato iniziale |
| Eventi | riprodotto | non riportato (avviso) | non riportato (avviso) | stato iniziale |

## File locali {#local-files}

Un'immagine locale viene letta solo con `--allow-local-files`, dalla cartella della presentazione, come in marp-cli. Senza l'opzione l'immagine resta fuori dalle esportazioni, e la compilazione lo dice con `local-file-blocked`, nominando il file, le sue diapositive e l'opzione. Un file che non esiste è segnalato come `local-file-missing`.

## Limiti {#limits}

Un SVG è un input non attendibile, perciò è limitato: 8 MiB di SVG, 50 000 elementi, 2 000 animazioni, 256 fotogrammi chiave per animazione, 60 secondi di linea temporale e 240 fotogrammi. Un disegno oltre un limite viene rifiutato con un avviso, e la compilazione prosegue.

## Che cosa è stato verificato {#checked}

LibreOffice Impress (macOS) apre i file PowerPoint e OpenDocument e mantiene ogni animazione quando li salva di nuovo, in entrambi i formati. Keynote (macOS) importa le forme e la GIF animata; le sue animazioni non sono state ispezionate. PowerPoint: non ancora verificato.
