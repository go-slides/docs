---
title: "Compatibilité Marp"
linkTitle: "Marp"
weight: 20
description: "Le Markdown que lit slides est celui de Marpit et de Marp Core, pas un dialecte."
tags: [marp]
---

Une présentation écrite pour [Marp](https://marp.app) est une présentation go-slides. Les règles
ci-dessous sont celles de [Marpit](https://marpit.marp.app) et de
[Marp Core](https://github.com/marp-team/marp-core) ; quand go-slides
s'en écarte, cette page le signale.

## Diapositives {#slides}

Une règle horizontale — `---` seul sur sa ligne, comme le découpage de Marpit — commence une
nouvelle diapositive, de même qu'un titre du niveau indiqué par la directive `headingDivider`.
Le front matter, entre les deux premiers `---`, contient les directives globales
et les données de la présentation.

## Directives {#directives}

| Type | Où | S'applique à |
| --- | --- | --- |
| globale | front matter, ou `<!-- theme: gaia -->` n'importe où | toute la présentation : `theme`, `style`, `headingDivider`, `lang`… |
| locale | `<!-- paginate: true -->` | cette diapositive et les suivantes |
| ponctuelle | `<!-- _class: lead -->` (tiret bas) | cette diapositive seulement |

Les directives locales sont celles de Marpit : `paginate`, `header`, `footer`, `class`,
`backgroundColor`, `backgroundImage`, `backgroundPosition`,
`backgroundRepeat`, `backgroundSize`, `color`. Les directives sont analysées comme du
YAML souple, comme le fait Marp Core ; `--strict-yaml` les analyse comme le fait
Marpit seul.

## Images {#images}

`![bg](image.jpg)` fait d'une image l'arrière-plan de la diapositive, avec les
mots-clés de Marp : `bg contain`, `bg cover`, `bg 50%`, `bg left`/`bg right` pour un arrière-plan
partagé, plusieurs images `bg` côte à côte, et les filtres d'image
(`blur`, `sepia`, `grayscale`…). Les tailles d'image en ligne (`w:200px`, `h:`) sont
lues comme Marp les lit.

## Autour de la diapositive {#around-the-slide}

- **En-tête et pied de page** : directives `header:` et `footer:`, en Markdown.
- **Pagination** : `paginate: true`.
- **Notes de l'orateur** : un commentaire HTML qui n'est pas une directive est une note.
  `slides notes` les affiche.
- **Listes progressives** : une liste écrite avec `*` (ou `1)`) apparaît élément par élément,
  comme dans Marp.
- **`<style scoped>`** : du CSS pour une seule diapositive.
- **HTML brut** : liste d'autorisation de Marp Core par défaut ; `--html` autorise tout.
- **Emoji** : les `:shortcodes:` et les emoji Unicode deviennent des images Twemoji, comme dans
  Marp Core, mais intégrées au lieu d'être chargées depuis un CDN : voir
  [Emoji]({{< relref "emoji" >}}).

## Thèmes {#themes}

`default`, `gaia` et `uncover` sont intégrés et se choisissent avec `theme:`. Un
thème personnel est un fichier CSS portant le commentaire `/* @theme name */` de Marp,
fourni avec `--theme-set` ou listé sous `themes:` dans `slides.yaml`.

## Différences avec Marp {#where-go-slides-differs}

- **Gabarits.** `{{ }}` dans le corps est une action de gabarit Go (voir
  [Gabarits]({{< relref "templates" >}})). Une présentation qui affiche `{{` comme du texte
  hors d'un bloc de code a besoin de `--no-template`, ou d'échapper ce texte comme les gabarits Go
  l'échappent. Les blocs et les extraits de code ne sont pas touchés.
- **Kinds.** `kind:` et `_kind:` sont des directives que Marp n'a pas (voir
  [Kinds]({{< relref "kinds" >}})). `slides expand` les remplace par les
  directives qu'elles représentent : la présentation développée est du Marp ordinaire.
- **Sorties.** HTML, PNG, JPEG, PDF, PPTX et ODP, au moyen de bibliothèques Go plutôt
  que d'un navigateur headless (voir [Ligne de commande]({{< relref "cli#images-pdf-pptx-and-odp" >}})
  pour la limite connue). En v0.4.0, comme avec marp-cli, un PPTX contient des images ; depuis
  la v0.5.0, ses diapositives sont des zones de texte et des formes modifiables, `--pptx-raster`
  conservant les images.
- **Titres ajustés.** `# <!-- fit --> Title` est pris en charge : le titre occupe toute
  la largeur de la diapositive sur une ligne, dans la page comme dans chaque export.
- **Mathématiques et code.** `$…$` et `$$…$$` suivent les règles de délimitation de Marp Core et
  sont composés par go-tex/math ; le code délimité est coloré par go-rouge avec
  les noms de classes de highlight.js (voir [Mathématiques et code]({{< relref "math-and-code" >}})).
