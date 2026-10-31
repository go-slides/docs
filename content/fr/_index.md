---
title: "Documentation de go-slides"
linkTitle: "Accueil"
type: docs
cascade:
  type: docs
description: "Des présentations compatibles Marp en pur Go (CGO=0), avec interpolation de gabarits Go et kinds."
---

**`slides` fait d'une présentation Marp un diaporama, en pur Go.** Il lit le
Markdown que lit [Marp](https://marp.app) — les directives de Marpit, `---` entre
les diapositives, les arrière-plans `![bg]`, les notes de l'orateur, les thèmes `default`, `gaia` et
`uncover` — et y ajoute deux choses que Marp n'a pas :

- des **[gabarits Go]({{< relref "templates" >}})** dans la présentation, évalués
  avant l'analyse du Markdown, si bien qu'une présentation peut générer ses propres diapositives ;
- des **[kinds]({{< relref "kinds" >}})** (le type d'une diapositive) : une diapositive dit ce qu'elle est
  (`title`, `section`…) au lieu de répéter à quoi elle ressemble.

{{< callout type="warning" >}}
**v0.14.0.** Ces pages décrivent
[go-slides/slides v0.14.0](https://github.com/go-slides/slides/releases/tag/v0.14.0)
et sont publiées sous `/docs/0.14/`.
{{< /callout >}}

## Ce qu'il fait aujourd'hui {#what-it-does-today}

| | |
| --- | --- |
| `slides build` (par défaut) | la présentation en une seule page HTML : chaque diapositive est une `<section>` dans un cadre SVG qui s'adapte à la fenêtre, avec le CSS du thème et une visionneuse pilotée au clavier |
| `slides build --pdf`, `--images png`, `--image jpeg` | la présentation en PDF (une page par diapositive, avec un plan) ou en images PNG/JPEG, sans navigateur : [go-webengine](https://github.com/go-webengine) fait le rendu, [go-pdfkit](https://github.com/go-pdfkit) écrit le PDF — en vectoriel depuis la v0.4.0, avec un texte sélectionnable : [PDF vectoriel]({{< relref "cli#vector-pdf" >}}) |
| mathématiques et code | `$…$` et `$$…$$` composés par go-tex/math, code délimité coloré par go-rouge, en HTML, en images et en PDF : [Mathématiques et code]({{< relref "math-and-code" >}}) |
| emoji | `:smile:` et 😄 deviennent des images Twemoji intégrées : elles se passent de réseau et s'affichent dans toutes les sorties : [Emoji]({{< relref "emoji" >}}) (depuis la v0.4.0) |
| `slides build --pptx` | la présentation en fichier PowerPoint, avec ses notes de l'orateur et ses métadonnées : une image par diapositive en v0.4.0, comme marp-cli ; **modifiable** depuis la v0.5.0 — zones de texte, formes, images, polices intégrées ([PPTX et ODP modifiables]({{< relref "cli#editable-pptx-and-odp" >}})) |
| `slides build --odp` | la même chose en présentation OpenDocument (`.odp`) |
| titres ajustés | `# <!-- fit --> Title` occupe toute la largeur de la diapositive sur une ligne, dans la page comme dans chaque export |
| `slides expand` | la présentation en Markdown Marp ordinaire : gabarits évalués, *kinds* remplacés par les directives ponctuelles qu'ils représentent |
| `slides notes` | les notes de l'orateur |
| `slides kinds`, `slides themes` | ce qu'une présentation peut nommer |
| `slides.yaml` | *kinds*, thèmes, données, politique HTML et langue communs à un dossier de présentations |
| HCL | front matter `---hcl`, `slides.hcl`, fichiers de *kinds* et de données en HCL2, avec expressions et fonctions : [HCL]({{< relref "hcl" >}}) |

## Prévu {#planned}

Pas encore réalisé, et indiqué ici pour que personne n'ait à deviner : les formules en vectoriel dans le
PDF (elles y sont en bitmap). La différence avec
marp-cli tient à la méthode : il produit images et PDF en pilotant un Chrome headless ;
ici, ce sont des bibliothèques Go qui s'en chargent.

## Pages {#pages}

- [Premiers pas]({{< relref "getting-started" >}})
- [Galerie]({{< relref "gallery" >}}) : douze présentations d'exemple, à ouvrir dans le playground
- [Compatibilité Marp]({{< relref "marp" >}})
- [Gabarits]({{< relref "templates" >}})
- [Kinds]({{< relref "kinds" >}})
- [HCL]({{< relref "hcl" >}})
- [Mathématiques et code]({{< relref "math-and-code" >}})
- [Emoji]({{< relref "emoji" >}})
- [Ligne de commande]({{< relref "cli" >}})
