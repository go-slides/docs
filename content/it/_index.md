---
title: "Documentazione di go-slides"
linkTitle: "Home"
type: docs
cascade:
  type: docs
description: "Presentazioni compatibili con Marp in puro Go (CGO=0), con interpolazione di template Go e slide kinds."
---

**`slides` trasforma una presentazione Marp in una presentazione vera e propria, in puro Go.** Legge il
Markdown che legge [Marp](https://marp.app) — le direttive di Marpit, `---` tra
le slide, gli sfondi `![bg]`, le note del relatore, i temi `default`, `gaia` e
`uncover` — e aggiunge due cose che Marp non ha:

- **[template Go]({{< relref "templates" >}})** dentro la presentazione, valutati
  prima che il Markdown venga analizzato, così una presentazione può generare le proprie slide;
- **[slide kinds]({{< relref "kinds" >}})** (il tipo di una slide): una slide dichiara cos'è
  (`title`, `section`…) invece di ripetere che aspetto ha.

{{< callout type="warning" >}}
**v0.15.0.** Queste pagine descrivono
[go-slides/slides v0.15.0](https://github.com/go-slides/slides/releases/tag/v0.15.0)
e sono pubblicate in `/docs/0.15/`.
{{< /callout >}}

## Cosa fa oggi {#what-it-does-today}

| | |
| --- | --- |
| `slides build` (il comando predefinito) | la presentazione come un'unica pagina HTML: ogni slide è una `<section>` inquadrata in un SVG che si adatta alla finestra, il CSS del tema, un visualizzatore da tastiera |
| `slides build --pdf`, `--images png`, `--image jpeg` | la presentazione come PDF (una pagina per slide, con un indice) o come immagini PNG/JPEG, senza browser: [go-webengine](https://github.com/go-webengine) esegue il rendering, [go-pdfkit](https://github.com/go-pdfkit) scrive il PDF — in vettoriale dalla v0.4.0, con testo selezionabile: [PDF vettoriale]({{< relref "cli#vector-pdf" >}}) |
| matematica e codice | `$…$` e `$$…$$` composti da go-tex/math, blocchi di codice evidenziati da go-rouge, in HTML, immagini e PDF: [Matematica e codice]({{< relref "math-and-code" >}}) |
| emoji | `:smile:` e 😄 diventano immagini Twemoji incorporate, quindi non richiedono rete e vengono disegnate in ogni output: [Emoji]({{< relref "emoji" >}}) (dalla v0.4.0) |
| `slides build --pptx` | la presentazione come file PowerPoint con le note del relatore e i metadati della presentazione: un'immagine per slide nella v0.4.0, come in marp-cli; **modificabile** dalla v0.5.0 — caselle di testo, forme, immagini, font incorporati ([PPTX e ODP modificabili]({{< relref "cli#editable-pptx-and-odp" >}})) |
| `slides build --odp` | lo stesso come presentazione OpenDocument (`.odp`) |
| titoli adattati | `# <!-- fit --> Title` riempie la larghezza della slide su una sola riga, nella pagina e in ogni esportazione |
| `slides expand` | la presentazione come semplice Markdown Marp: template valutati, kinds sostituiti dalle direttive spot che rappresentano |
| `slides notes` | le note del relatore |
| `slides kinds`, `slides themes` | ciò che una presentazione può nominare |
| `slides.yaml` | kinds, temi, dati, criterio HTML e lingua condivisi per una cartella di presentazioni |
| HCL | front matter `---hcl`, `slides.hcl`, kinds e file di dati in HCL2, con espressioni e funzioni: [HCL]({{< relref "hcl" >}}) |

La differenza con
marp-cli è il percorso: esso produce immagini e PDF pilotando un Chrome headless;
qui lo fanno librerie Go.

## Pagine {#pages}

- [Per iniziare]({{< relref "getting-started" >}})
- [Galleria]({{< relref "gallery" >}}): dodici presentazioni di esempio, da aprire nel Playground
- [Compatibilità con Marp]({{< relref "marp" >}})
- [Template]({{< relref "templates" >}})
- [Kinds]({{< relref "kinds" >}})
- [HCL]({{< relref "hcl" >}})
- [Matematica e codice]({{< relref "math-and-code" >}})
- [Emoji]({{< relref "emoji" >}})
- [Riga di comando]({{< relref "cli" >}})
