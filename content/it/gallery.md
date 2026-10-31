---
title: "Galleria"
weight: 12
description: "Dodici piccole presentazioni, ciascuna con una sola cosa ben fatta: temi, kinds, template, HCL, matematica, immagini, codice, emoji, transizioni e un intervento completo. Aprile qualsiasi nel Playground."
tags: [gallery, examples]
---

Dodici piccole presentazioni dalla cartella
[`examples/`](https://github.com/go-slides/slides/tree/main/examples) di
go-slides, ciascuna con una sola cosa ben fatta. Le miniature sono la prima slide di
ogni presentazione, resa da go-slides stesso quando questo sito viene compilato; la CI compila ogni
presentazione in HTML e PDF, quindi nessuna si rompe.

**Apri nel Playground** (l'editor di go-slides nel browser) carica la presentazione nel tuo browser, tramite un link che non invia
nulla a nessun server. Una presentazione che ha file oltre a `deck.md` (un file di kinds, cioè di tipi di slide,
parti incluse, immagini) non può viaggiare in un link: scarica il suo **.zip**
e importalo con il pannello dei file del Playground, oppure decomprimilo e compilalo con
`slides build --allow-local-files deck.md`.

Le immagini sono SVG originali degli autori di go-slides, rilasciati con
[CC0 1.0](https://creativecommons.org/publicdomain/zero/1.0/), e le presentazioni
che le usano lo dichiarano.

{{< gallery >}}
