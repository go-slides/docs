---
title: "Emoji"
weight: 36
description: "Gli shortcode emoji e le emoji Unicode diventano immagini Twemoji, senza rete: nella pagina HTML, nelle immagini, nel PDF, in PowerPoint e in OpenDocument."
tags: [emoji]
---

{{< callout type="warning" >}}
Le emoji sono nella v0.4.0; la v0.3.0 non ne aveva.
{{< /callout >}}

Scrivi un'emoji come shortcode o come carattere stesso:

```markdown
# Launch :rocket: 🎉

Shortcodes :smile: :+1: :heart_eyes: and Unicode 😄 👍🏽 👩‍💻 🇫🇷 1️⃣.
```

Ciascuna diventa l'immagine [Twemoji](https://github.com/jdecked/twemoji) che Marp Core
disegna, con il markup che Marp Core scrive:

```html
<img class="emoji" draggable="false" alt="😄" src="data:image/svg+xml;base64,…" data-marp-twemoji=""/>
```

## Senza rete {#no-network}

Il `src` di Marp Core è un URL su una CDN, quindi una presentazione mostrata offline, stampata o
trasformata in PDF su una macchina senza connessione mostra immagini rotte.
go-slides mette l'SVG stesso in un URI `data:`. Le grafiche (4.009, 
Twemoji 17.0.3, Emoji 17.0) sono incorporate nel programma dal modulo in puro Go
[go-icons/twemoji](https://github.com/go-icons/twemoji), 2 MB compressi e
letti solo quando una presentazione contiene un'emoji.

Le immagini sono disegnate allo stesso modo dalla pagina HTML, da `--images`/`--image`, da `--pdf`, da
`--pptx` e da `--odp`: go-webengine rasterizza l'SVG. Tutte le 4.009
grafiche vengono analizzate e disegnate dal suo rasterizzatore SVG; lo verifica un test. Dalla
v0.5.0 quel rasterizzatore è [go-gfx/oksvg](https://github.com/go-gfx/oksvg)
v0.2.0: `clipPath`, `fill-rule: evenodd` e i comandi ad arco ora vengono disegnati come li
disegna Chrome, il che ha cambiato una Twemoji e altre 211, e i loghi SVG non
risultano più alterati.

La regola che dimensiona un'emoji sul testo (`img[data-marp-twemoji]`, un quadrato
di un `em`, allineato alla linea di base) viene aggiunta al tema come la aggiunge Marp Core,
quindi la ottiene anche un tema tutto tuo.

## Cosa viene convertito {#what-is-converted}

Le regole sono quelle di Marp Core; un test le confronta con l'HTML che Marp Core 4.4.0
scrive per una presentazione contenente ogni caso.

| | |
| --- | --- |
| `:smile:`, `:+1:`, `:heart_eyes:` | i 1.903 nomi di markdown-it-emoji, nel testo, nei titoli, nei link, nelle tabelle, in `header` e `footer`; non nel codice, negli URL né nel testo alternativo di un'immagine |
| `😄`, `👩‍💻`, `🇫🇷`, `1️⃣`, `👍🏽` | emoji Unicode, nel testo **e nel codice**: Marp Core riscrive l'HTML del codice già reso |
| `✔`, `©`, `™`, `▶` | restano testo: i caratteri la cui presentazione predefinita è testuale sono emoji solo con U+FE0F (`✔️`); `❤` è l'eccezione |
| `☺︎` | resta testo: un U+FE0E dopo un'emoji richiede il glifo semplice |
| `12:30:45`, `:not_a_name:` | restano come scritti |

L'`id` di un titolo e l'indice usano l'emoji, non il nome dello shortcode:
`# Launch :rocket:` ha l'id `launch-%F0%9F%9A%80`, come in Marp.

Dove go-slides differisce da Marp Core, converte di più: `✌️ ☝️ ✍️ 🏋️` e alcuni
altri caratteri a presentazione testuale che l'emoji-regex di Marp Core lascia come testo,
e le poche sequenze che Twemoji disegna e che Unicode non elenca.

## Disattivarle {#turning-it-off}

`--no-emoji` (nel codice, `slides.Config{NoEmoji: true}`) lascia shortcode ed
emoji Unicode come il testo in cui sono stati scritti, ovvero l'`emoji: false` di
Marp Core.

## Attribuzione {#attribution}

Le grafiche sono Twemoji, Copyright 2020 Twitter, Inc e altri collaboratori,
con licenza [CC-BY 4.0](https://creativecommons.org/licenses/by/4.0/) e
portate avanti dal progetto [jdecked/twemoji](https://github.com/jdecked/twemoji).
Una presentazione che le mostra dovrebbe citarle, per esempio in un piè di pagina o in una
slide finale. La tabella degli shortcode è quella di markdown-it-emoji (MIT).
