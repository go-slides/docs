---
title: "Registrazioni del terminale"
weight: 38
description: "Una registrazione asciinema in una presentazione: riprodotta in un piccolo lettore della pagina HTML, disegnata come testo vettoriale nel PDF, in PowerPoint e in OpenDocument."
tags: [asciinema, terminal]
---

{{< callout type="warning" >}}
Le registrazioni del terminale sono nel ramo principale dopo la v0.17.1 e in nessuna versione pubblicata.
{{< /callout >}}

Una registrazione [asciinema](https://docs.asciinema.org/) ([asciicast v2](https://docs.asciinema.org/manual/asciicast/v2/) o [v3](https://docs.asciinema.org/manual/asciicast/v3/)) è un'immagine il cui file termina con `.cast`:

```markdown
![cast](demo.cast)
![The failing build](demo.cast#t=12.5)
```

## Sintassi {#syntax}

Il file viene letto dalla cartella della presentazione: un URL, un percorso assoluto o un percorso che esce dalla cartella viene rifiutato. Il testo alternativo nomina la registrazione per gli screen reader; `cast` da solo indica il titolo della registrazione. `#t=12.5`, oppure `#t=1:23` in minuti e secondi, sceglie il momento che mostrano le uscite che non possono muoversi; senza, è la fine della registrazione.

## Cosa mostra ogni uscita {#outputs}

- **HTML** (la pagina e l'anteprima del playground): la registrazione viene riprodotta in un piccolo lettore di go-slides incorporato nella pagina: riproduzione e pausa, un pulsante di velocità e una barra di posizione. Non carica script da nessuna parte e non usa la rete.
- **PDF, PowerPoint e OpenDocument** (`--pdf`, `--pptx`, `--odp`): il terminale nel momento indicato da `#t=`, come testo vettoriale nel carattere a spaziatura fissa del tema, con i colori a 16, 256 e diretti, grassetto, corsivo, sottolineato e testo invertito. Il testo si può selezionare e cercare. Non è mai un'immagine. Grassetto e corsivo sono disegnati dove il carattere a spaziatura fissa del tema li ha (Roboto Mono nel tema gaia).

## Il lettore {#player}

Viene riprodotto una volta, la prima volta che la registrazione è sullo schermo. Quando il browser chiede meno movimento (`prefers-reduced-motion`), parte in pausa, mostrando il momento indicato da `#t=`, e si avvia quando premi riproduci. I suoi pulsanti e tasti non cambiano diapositiva.

## Limiti {#limits}

Una registrazione è un input non attendibile, perciò è limitata: il file (16 MiB), gli eventi (500.000), la durata (3 ore), il terminale (400 per 200 celle) e il lavoro che l'emulatore può fare. Una registrazione oltre un limite, o che non è un file asciicast ben formato, viene rifiutata: la compilazione prosegue, la diapositiva mostra una nota al suo posto e il problema è un avviso con un codice: `cast.syntax`, `cast.limit`, `cast.version` o `cast.path`.

## Nel playground {#playground}

Aggiungi un file `.cast` al progetto dal pannello File, trascinandolo o dal selettore di file. Va in `media/` e viene inserito come immagine. È testo che la presentazione legge, quindi conta con gli altri file di testo del progetto, limitati a 2 MiB. Provalo nel [playground](https://go-slides.github.io/playground/).
