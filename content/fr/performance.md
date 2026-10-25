---
title: "Performances"
weight: 60
description: "La vitesse à laquelle go-slides génère une présentation, et avec quelle mémoire, mesurées face à marp-cli sur la même machine."
tags: [performance]
---

{{< callout type="info" >}}
Les chiffres PPTX ci-dessous concernent la v0.5.0 et ceux du PDF vectoriel la v0.4.0.
Le premier tableau concerne la v0.3.0,
la version précédente, qui fait le rendu des diapositives en parallèle et n'analyse ses polices
qu'une fois ([go-slides/slides#11](https://github.com/go-slides/slides/pull/11)).
{{< /callout >}}

go-slides et [marp-cli](https://github.com/marp-team/marp-cli) 4.5.1 ont été
lancés en sous-processus sur les mêmes présentations, sur un Apple M4 Max (16 CPU) avec une
charge moyenne sur 1 minute inférieure à 8, en cinq tours entrelacés après un échauffement.
Le temps est la médiane du temps écoulé ; la mémoire est le pic de toute l'arborescence de processus
(marp-cli pilote un Chrome, dont les processus sont comptés).

| Présentation | Sortie | go-slides | marp-cli | mémoire go-slides | mémoire marp-cli |
| --- | --- | ---: | ---: | ---: | ---: |
| 100 diapositives | HTML | 36 ms | 315 ms | 47 Mio | 140 Mio |
| 100 diapositives | PDF | 1,14 s | 1,88 s | 242 Mio | 1 509 Mio |
| 100 diapositives | PNG, toutes les diapositives | 0,97 s | 5,44 s | 225 Mio | 1 354 Mio |
| 100 diapositives | PPTX | 3,39 s | 8,03 s | 530 Mio | 1 373 Mio |
| gaia, 20 diapositives | PDF | 283 ms | 1,57 s | 154 Mio | 1 315 Mio |
| mathématiques et code, 20 diapositives | PNG, toutes les diapositives | 250 ms | 2,32 s | 168 Mio | 1 286 Mio |

## PDF vectoriel {#vector-pdf}

Depuis la v0.4.0, le PDF est dessiné en vectoriel (voir
[Ligne de commande]({{< relref "cli#vector-pdf" >}})). Même machine et même méthode,
le 2026-10-19 :

| Présentation | go-slides | marp-cli | v0.3.0 | taille go-slides | taille marp-cli | taille v0.3.0 |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| 100 diapositives | 306 ms | 1 763 ms | 1 147 ms | 758 Kio | 837 Kio | 10,0 Mio |
| gaia, 20 diapositives | 93 ms | 1 546 ms | 285 ms | 32 Kio | 66 Kio | 582 Kio |
| mathématiques et code, 20 diapositives | 140 ms | 1 643 ms | 294 ms | 124 Kio | 335 Kio | 638 Kio |

Mémoire pour le PDF de 100 diapositives : 158 Mio, contre 1 512 Mio pour marp-cli ; temps
CPU : 0,9 s contre 2,7 s.

Les chiffres de la v0.4.0 pour le PDF de 100 diapositives ont été mesurés le 2026-10-19 à 21 h 29,
avec une charge moyenne sur 1 minute de 7,7, sur trois tours : go-slides 306 ms, 758
Kio et 158 Mio ; marp-cli 1 838 ms, 837 Kio et 1 512 Mio.

## PPTX modifiable {#editable-pptx}

Depuis la v0.5.0, le PPTX est composé de zones de texte, de formes et d'images (voir
[Ligne de commande]({{< relref "cli#editable-pptx-and-odp" >}})). Mesuré le
2026-10-19 à 22 h 51, avec une charge moyenne sur 1 minute de 7,1, sur trois tours :

| Présentation | go-slides | marp-cli | taille go-slides | taille marp-cli | mémoire go-slides | mémoire marp-cli |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| 100 diapositives | 348 ms | 7 795 ms | 760 Kio | 7,2 Mio | 204 Mio | 1 370 Mio |
| gaia, 20 diapositives | 112 ms | 4 608 ms | 893 Kio | 4,8 Mio | | |
| mathématiques et code, 20 diapositives | 207 ms | 2 780 ms | 842 Kio | 1,5 Mio | | |

Pour la présentation de 100 diapositives, la v0.4.0 demandait 4,7 s, 4,8 Mio et 540 Mio. Le
PPTX de marp-cli contient des images. L'essentiel de la taille d'une présentation modifiable tient aux polices intégrées.

## Pourquoi c'est rapide {#why-it-is-fast}

- **Aucun navigateur à démarrer.** marp-cli lance Chrome pour chaque export PDF, image et
  PPTX ; go-slides fait le rendu avec
  [go-webengine](https://github.com/go-webengine/engine) dans son propre
  processus.
- **Les diapositives sont rendues en parallèle**, quatre à la fois par défaut, et chaque image est
  écrite dès que son tour vient : une longue présentation n'est jamais conservée en mémoire sous forme
  d'images.
- **Les polices sont analysées une seule fois** par processus, et non une fois par diapositive.

## Là où il ne gagne pas {#what-it-does-not-win}

- **La taille du PDF, en v0.3.0.** Chaque diapositive était une image dans le PDF : 10 Mio contre
  837 Kio pour la présentation de 100 diapositives. Le PDF vectoriel ci-dessus est plus petit que celui de
  marp-cli.
- **Le temps CPU sur les grosses présentations.** Quatre workers consomment plus de CPU que ne le laisse voir
  le temps écoulé : en v0.3.0, 3,2 s contre 2,7 s pour marp-cli sur le PDF de 100 diapositives
  (0,9 s pour le PDF vectoriel).
- **La mémoire est une borne supérieure pour marp-cli** : les processus de Chrome partagent des pages
  que la somme compte dans chacun d'eux. Son plus gros processus pris isolément occupe 220 à 320 Mio.

La méthode, chaque présentation et chaque mode, ainsi que les commandes permettant de reproduire les chiffres,
se trouvent dans
[BENCHMARKS.md](https://github.com/go-slides/slides/blob/main/BENCHMARKS.md).
