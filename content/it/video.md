---
title: "Video"
weight: 39
description: "Un video in una presentazione: riprodotto dal browser nella pagina HTML, un fotogramma di anteprima con collegamento nel PDF, incorporato in PowerPoint e OpenDocument."
tags: [video]
---

Un video è un'immagine il cui file termina con `.mp4`, `.m4v` o `.webm`:

```markdown
![video](clip.mp4)
![The demo](clip.mp4#t=12.5)
![demo loop muted](clip.mp4)
```

## Sintassi {#syntax}

Il file viene letto dalla cartella della presentazione, o da un URL `http` o `https`. Un percorso assoluto, un percorso che esce dalla cartella o qualsiasi altro schema viene rifiutato. `#t=12.5`, o `#t=1:23` in minuti e secondi, sceglie il fotogramma di anteprima: l'immagine mostrata prima della riproduzione e dalle uscite che non possono riprodurre il video. Anche la forma di frammento multimediale `#t=10,20` è accettata, e se ne usa l'inizio. Senza, l'anteprima è il primo fotogramma.

Le parole chiave nel testo alternativo, scritte come le parole chiave di immagine di Marp ma proprie di go-slides, stabiliscono come il video viene riprodotto: `autoplay`, `loop` e `muted`. `autoplay` è rispettato solo insieme a `muted`; altrimenti è ignorato, con un avviso. Le altre parole del testo alternativo sono l'etichetta accessibile.

## Cosa mostra ogni uscita {#outputs}

- **HTML** (la pagina e l'anteprima del playground): un elemento `<video controls playsinline preload="metadata" poster=...>`. L'anteprima è il fotogramma decodificato, incorporato come URI di dati JPEG. Un video non suona mai da solo. Quando il browser chiede meno movimento (`prefers-reduced-motion`), `autoplay` e `loop` vengono rimossi e il video è messo in pausa. Un video che parte da solo (senza audio) viene riprodotto solo finché almeno metà è visibile, e va in pausa quando la sua diapositiva non è mostrata.
- **PDF** (`--pdf`): il fotogramma di anteprima come immagine, con un collegamento al file video: il percorso relativo come scritto nella presentazione, o l'URL per una sorgente `http` o `https`. La riproduzione avviene solo aprendo il collegamento. Il PDF non contiene alcuna annotazione `RichMedia`: si riprodurrebbe solo in Acrobat, e non viene usata di proposito.
- **PowerPoint** (`--pptx`): il video è incorporato come parte multimediale (`ppt/media/slidesvideoN.mp4|webm`) con la sua anteprima, come `a:videoFile` più l'estensione `p14:media`. Keynote: lo importa come filmato e lo riproduce (verificato). LibreOffice Impress: importa i file PowerPoint e OpenDocument come video e lo decodifica (verificato). PowerPoint: non ancora verificato. Le parole chiave `autoplay`, `loop` e `muted` non passano in PowerPoint né in OpenDocument.
- **OpenDocument** (`--odp`): il video è incorporato sotto `Video/` come `draw:plugin` (ODF 1.3 parte 3, 10.4.8), prima del `draw:image` dell'anteprima, ed elencato nel manifesto. LibreOffice Impress lo importa come video e lo decodifica (verificato); la riproduzione nella presentazione dipende dal suo backend multimediale e dal codec.
- **Sorgenti URL**: la pagina HTML trasmette il video dall'URL. Le esportazioni mostrano una scheda neutra con collegamento all'URL. Niente viene scaricato durante la costruzione, e niente viene incorporato in PowerPoint o OpenDocument.

## Contenitori e anteprime {#codecs}

I file MP4 (ISO BMFF) e WebM (Matroska) vengono letti, in Go; ffmpeg non viene mai eseguito. L'anteprima è decodificata solo per H.264 (`avc1` o `avc3`), l'unico decodificatore video di go-avkit. Gli altri codec (H.265, VP8, VP9, AV1) si riproducono comunque in HTML e sono incorporati in PowerPoint e OpenDocument, ma la loro anteprima è una scheda scura neutra con un triangolo di riproduzione, e viene stampato l'avviso `video.codec`.

## Limiti {#limits}

Un video è un input non attendibile, quindi è limitato: il file (64 MiB), i box o elementi (262.144) e il loro annidamento (16 livelli), il fotogramma (4096 per 4096) e la decodifica dell'anteprima (al massimo 600 campioni dal fotogramma chiave precedente, 8 MiB per campione). Un'anteprima troppo lontana da un fotogramma chiave riceve la scheda e l'avviso `video.limit`.

Un video rifiutato è sostituito nella diapositiva da una nota, `[video not shown: ...]`; la costruzione prosegue, e il problema è un avviso con un codice:

- `video.syntax`: non è un file MP4 o WebM, nessuna traccia video, o un tempo non valido;
- `video.limit`: un limite qui sopra è superato;
- `video.codec`: l'anteprima non può essere decodificata per questo codec;
- `video.path`: un percorso assoluto, un percorso che esce dalla cartella della presentazione, o uno schema diverso da `http` e `https`;
- `file.missing`: il file non esiste.

## Nel playground {#playground}

Una volta rilasciato, un video potrà essere aggiunto al progetto come gli altri file. I suoi byte restano nel browser, come URL blob, e non vengono mai caricati. Provalo nel [playground](https://go-slides.github.io/playground/).

Il supporto video si carica su richiesta. Il lettore che decodifica il fotogramma del poster (circa 3 MiB compressi) è un modulo a parte, scaricato la prima volta che una presentazione ha bisogno di un poster. Una presentazione senza video non lo scarica mai, e il primo caricamento del playground è più leggero di altrettanto. Gira in un Web Worker, quindi la digitazione non viene mai bloccata, e legge un file con gli stessi limiti di `slides` (64 MiB, box e dimensione del fotogramma) e gli stessi errori `video.limit` e `video.syntax`.
