---
title: "SVG animé"
weight: 40
description: "Une image SVG avec des animations SMIL ou CSS : jouée par le navigateur en HTML, formes animées dans PowerPoint et OpenDocument, GIF animé là où elles ne peuvent pas le dire, état initial dans le PDF et les images."
tags: [svg, animation]
---

Une image SVG qui s'anime d'elle-même s'écrit comme n'importe quelle image :

```markdown
![](spinner.svg)
![w:400](chart.svg)
```

## Ce qui est lu {#read}

Les animations écrites dans le fichier SVG lui-même : SMIL (`animate`, `animateTransform`, `animateMotion`, `animateColor`, `set`, avec `begin`, `dur`, `repeatCount`, `fill`, `calcMode`, `keyTimes` et `keySplines`) et CSS (`@keyframes` avec les propriétés `animation`, `transform-origin` et `transform-box`). Ce qui n'est pas compris, un événement comme `begin="click"` par exemple, est laissé de côté et signalé par un avertissement : rien n'est abandonné en silence.

## Ce que montre chaque sortie {#outputs}

- **HTML** : le navigateur joue les animations lui-même.
- **PowerPoint** (`--pptx`) et **OpenDocument** (`--odp`) : chaque élément animé est une forme à part, et ses animations démarrent avec la diapositive : déplacement, rotation et échelle, opacité, passage d'une couleur à une autre. Le mouvement d'un élément est échantillonné aux instants propres de l'animation, si bien que les accélérations, les images clés multiples et une rotation autour d'un autre point sont exactes. Les répétitions (`indefinite` aussi), les délais et les allers-retours sont repris.
- **Images successives** : quand une animation ne peut pas se dire ainsi (un tracé qui se déforme, un trait qui se dessine, une inclinaison), l'image entière est rendue aux instants propres de l'animation et incorporée comme GIF animé, que PowerPoint et Impress lisent. C'est borné : 12 images par seconde, 240 images au plus, 16 millions de pixels pour l'ensemble des images.
- **PDF** (`--pdf`) et images (`--images`) : ils ne peuvent pas s'animer. Le dessin est montré dans son état initial, sans JavaScript ni `RichMedia`, et un avertissement nomme les images animées.

Chaque image animée reçoit une ligne d'avertissement de code `svg-animation`, qui dit comment elle a été reprise : nativement, ou en images successives et pourquoi.

## Prise en charge par sortie {#support}

| Fonctionnalité | HTML | PowerPoint | OpenDocument | PDF et images |
|---|---|---|---|---|
| Translation, déplacement le long d'un tracé | joué | natif | natif | état initial |
| Rotation et échelle, autour de n'importe quel point | joué | natif | natif | état initial |
| Rayon d'un cercle, taille d'un rectangle (sans contour) | joué | natif | natif | état initial |
| Opacité d'un élément d'une seule forme | joué | natif | natif | état initial |
| Couleur d'un remplissage ou d'un contour, d'une couleur à une autre | joué | natif | natif | état initial |
| Répétitions, délais, allers-retours, maintien à la fin | joué | natif | natif | état initial |
| Accélération, images clés multiples | joué | natif (échantillonné) | natif (échantillonné) | état initial |
| Paliers, changements discrets | joué | images successives | images successives | état initial |
| Déformation de tracé, décalage de tirets, attributs de position et de taille, dégradés | joué | images successives | images successives | état initial |
| Inclinaison, symétrie, plusieurs horloges sur un élément, éléments animés imbriqués | joué | images successives | images successives | état initial |
| Événements | joué | non repris (avertissement) | non repris (avertissement) | état initial |

## Fichiers locaux {#local-files}

Une image locale n'est lue qu'avec `--allow-local-files`, dans le dossier de la présentation, comme dans marp-cli. Sans elle, l'image est écartée des exports, et la construction le dit avec `local-file-blocked`, en nommant le fichier, ses diapositives et l'option. Un fichier qui n'existe pas est signalé par `local-file-missing`.

## Limites {#limits}

Un SVG est une entrée non fiable, donc il est borné : 8 Mio de SVG, 50 000 éléments, 2 000 animations, 256 images clés pour une animation, 60 secondes de chronologie et 240 images. Un dessin qui dépasse une borne est refusé avec un avertissement, et la construction continue.

## Ce qui a été vérifié {#checked}

LibreOffice Impress (macOS) ouvre les fichiers PowerPoint et OpenDocument et garde chaque animation quand il les enregistre de nouveau, dans les deux formats. Keynote (macOS) importe les formes et le GIF animé ; ses animations n'ont pas été inspectées. PowerPoint : pas encore vérifié.
