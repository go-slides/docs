---
title: "Galerie"
weight: 12
description: "Douze petites présentations, chacune montrant bien une chose : thèmes, kinds, gabarits, HCL, mathématiques, images, code, emoji, transitions et un exposé complet. Ouvrez-les dans le Playground."
tags: [gallery, examples]
---

Douze petites présentations tirées du dossier
[`examples/`](https://github.com/go-slides/slides/tree/main/examples) de
go-slides, chacune montrant bien une chose. Les vignettes sont la première diapositive de
chaque présentation, rendue par go-slides lui-même lors de la génération de ce site ; la CI génère chaque
présentation en HTML et en PDF, pour qu'aucune ne se casse.

**Ouvrir dans le Playground** (bac à sable en ligne) charge la présentation dans votre navigateur, par un lien qui n'envoie
rien à aucun serveur. Une présentation qui comporte d'autres fichiers que `deck.md` (un fichier de *kinds*,
des parties incluses, des images) ne peut pas tenir dans un lien : téléchargez son **.zip**, puis
importez-le avec le panneau de fichiers du Playground, ou décompressez-le et générez la présentation avec
`slides build --allow-local-files deck.md`.

Les illustrations sont des SVG originaux des auteurs de go-slides, placés sous
[CC0 1.0](https://creativecommons.org/publicdomain/zero/1.0/), et les présentations
qui les utilisent le mentionnent.

{{< gallery >}}
