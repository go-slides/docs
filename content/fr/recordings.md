---
title: "Enregistrements de terminal"
weight: 38
description: "Un enregistrement asciinema dans une présentation : rejoué dans un petit lecteur de la page HTML, dessiné en texte vectoriel dans le PDF, PowerPoint et OpenDocument."
tags: [asciinema, terminal]
---

{{< callout type="info" >}}
Les enregistrements de terminal sont dans la v0.18.0 et les suivantes.
{{< /callout >}}

Un enregistrement [asciinema](https://docs.asciinema.org/) ([asciicast v2](https://docs.asciinema.org/manual/asciicast/v2/) ou [v3](https://docs.asciinema.org/manual/asciicast/v3/)) est une image dont le fichier se termine par `.cast` :

```markdown
![cast](demo.cast)
![The failing build](demo.cast#t=12.5)
```

## Syntaxe {#syntax}

Le fichier est lu dans le dossier de la présentation : une URL, un chemin absolu ou un chemin qui sort du dossier est refusé. Le texte alternatif nomme l'enregistrement pour les lecteurs d'écran ; `cast` seul désigne le titre de l'enregistrement. `#t=12.5`, ou `#t=1:23` en minutes et secondes, choisit l'instant que montrent les sorties qui ne peuvent pas bouger ; sans lui, c'est la fin de l'enregistrement.

## Ce que montre chaque sortie {#outputs}

- **HTML** (la page, et l'aperçu du bac à sable) : l'enregistrement est rejoué dans un petit lecteur de go-slides intégré à la page : lecture et pause, un bouton de vitesse et une barre de position. Il ne charge aucun script d'ailleurs et n'utilise pas le réseau.
- **PDF, PowerPoint et OpenDocument** (`--pdf`, `--pptx`, `--odp`) : le terminal à l'instant que désigne `#t=`, en texte vectoriel dans la police à chasse fixe du thème, avec les couleurs 16, 256 et directes, le gras, l'italique, le soulignement et le texte inversé. Le texte peut être sélectionné et recherché. Ce n'est jamais une image. Le gras et l'italique sont dessinés quand la police à chasse fixe du thème les possède (Roboto Mono dans le thème gaia).

## Le lecteur {#player}

Il joue une fois, la première fois que l'enregistrement est à l'écran. Quand le navigateur demande moins de mouvement (`prefers-reduced-motion`), il démarre en pause, en montrant l'instant que désigne `#t=`, et joue quand vous appuyez sur lecture. Ses boutons et ses touches ne font pas changer de diapositive.

## Limites {#limits}

Un enregistrement est une entrée non fiable ; il est donc borné : le fichier (16 MiB), les événements (500 000), la durée (3 heures), le terminal (400 par 200 cellules) et le travail que l'émulateur peut faire. Un enregistrement qui dépasse une borne, ou qui n'est pas un fichier asciicast bien formé, est refusé : la construction continue, la diapositive montre une note à sa place, et le problème est un avertissement portant un code : `cast.syntax`, `cast.limit`, `cast.version` ou `cast.path`.

## Dans le bac à sable {#playground}

Ajoutez un fichier `.cast` au projet depuis le panneau Fichiers, en le déposant ou par le sélecteur de fichiers. Il va dans `media/` et est inséré comme une image. C'est du texte que la présentation lit : il compte avec les autres fichiers texte du projet, limités à 2 MiB. Essayez-le dans le [bac à sable](https://go-slides.github.io/playground/).
