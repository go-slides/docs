---
title: "Compatibilità con Marp"
linkTitle: "Marp"
weight: 20
description: "Il Markdown che slides legge è quello di Marpit e Marp Core, non un suo dialetto."
tags: [marp]
---

Una presentazione scritta per [Marp](https://marp.app) è una presentazione go-slides. Le regole
qui sotto sono quelle di [Marpit](https://marpit.marp.app) e di
[Marp Core](https://github.com/marp-team/marp-core); dove go-slides
differisce, questa pagina lo dice.

## Slide {#slides}

Una linea orizzontale — `---` su una riga a sé, come la divide Marpit — inizia una
nuova slide, e lo fa anche un titolo del livello indicato dalla direttiva
`headingDivider`. Il front matter, tra i primi due `---`, contiene le direttive globali
e i dati della presentazione.

## Direttive {#directives}

| Tipo | Dove | Si applica a |
| --- | --- | --- |
| globale | front matter, o `<!-- theme: gaia -->` ovunque | l'intera presentazione: `theme`, `style`, `headingDivider`, `lang`… |
| locale | `<!-- paginate: true -->` | questa slide e quelle successive |
| spot | `<!-- _class: lead -->` (trattino basso) | solo questa slide |

Le direttive locali sono quelle di Marpit: `paginate`, `header`, `footer`, `class`,
`backgroundColor`, `backgroundImage`, `backgroundPosition`,
`backgroundRepeat`, `backgroundSize`, `color`. Le direttive vengono analizzate come YAML
permissivo, come fa Marp Core; `--strict-yaml` le analizza come fa Marpit semplice.

## Immagini {#images}

`![bg](image.jpg)` rende un'immagine lo sfondo della slide, con le parole chiave di Marp:
`bg contain`, `bg cover`, `bg 50%`, `bg left`/`bg right` per uno sfondo
diviso, più immagini `bg` affiancate e i filtri per immagini
(`blur`, `sepia`, `grayscale`…). Le dimensioni delle immagini in linea (`w:200px`, `h:`) vengono
lette come le legge Marp.

## Attorno alla slide {#around-the-slide}

- **Intestazione e piè di pagina**: direttive `header:` e `footer:`, in Markdown.
- **Numerazione**: `paginate: true`.
- **Note del relatore**: un commento HTML che non è una direttiva è una nota.
  `slides notes` le stampa.
- **Elenchi a frammenti**: un elenco scritto con `*` (o `1)`) appare voce per voce,
  come in Marp.
- **`<style scoped>`**: CSS per una sola slide.
- **HTML grezzo**: per impostazione predefinita la lista consentita di Marp Core; `--html` lo consente tutto.
- **Emoji**: `:shortcodes:` ed emoji Unicode diventano immagini Twemoji, come in
  Marp Core, ma incorporate invece che caricate da una CDN: vedi
  [Emoji]({{< relref "emoji" >}}).

## Temi {#themes}

`default`, `gaia` e `uncover` sono integrati e si scelgono con `theme:`. Un
tema tutto tuo è un file CSS con il commento `/* @theme name */` di Marp,
passato con `--theme-set` o elencato sotto `themes:` in `slides.yaml`.

## Dove go-slides differisce {#where-go-slides-differs}

- **Template.** `{{ }}` nel corpo è un'azione di template Go (vedi
  [Template]({{< relref "templates" >}})). Una presentazione che mostra `{{` come testo
  fuori da un blocco di codice richiede `--no-template`, oppure il testo con l'escape come lo
  usano i template Go. I blocchi di codice e le porzioni di codice non vengono toccati.
- **Tipi.** `kind:` e `_kind:` sono direttive che Marp non ha (vedi
  [Tipi di slide]({{< relref "kinds" >}})). `slides expand` le sostituisce con le
  direttive che rappresentano, quindi la presentazione espansa è Marp semplice.
- **Output.** HTML, PNG, JPEG, PDF, PPTX e ODP, tramite librerie Go anziché
  un browser headless (vedi [Riga di comando]({{< relref "cli#images-pdf-pptx-and-odp" >}})
  per il limite noto). Nella v0.4.0, come in marp-cli, un PPTX contiene immagini; dalla
  v0.5.0 le sue slide sono caselle di testo e forme modificabili, con `--pptx-raster` che mantiene
  le immagini.
- **Titoli adattati.** `# <!-- fit --> Title` è supportato: il titolo riempie
  la larghezza della slide su una riga, nella pagina e in ogni esportazione.
- **Matematica e codice.** `$…$` e `$$…$$` seguono le regole sui delimitatori di Marp Core e
  sono composti da go-tex/math; il codice delimitato viene evidenziato da go-rouge con
  i nomi di classe di highlight.js (vedi [Matematica e codice]({{< relref "math-and-code" >}})).
