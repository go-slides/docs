---
title: "HCL"
weight: 45
description: "Du HCL2 partout où une présentation accepte du YAML : front matter HCL, slides.hcl, fichiers de kinds et de données, avec expressions et fonctions."
tags: [hcl, configuration]
---

{{< callout type="warning" >}}
HCL est arrivé en v0.3.0 ([go-slides/slides#13](https://github.com/go-slides/slides/pull/13)).
{{< /callout >}}

Partout où `slides` lit du YAML, il lit aussi la syntaxe native de
[HCL2](https://github.com/hashicorp/hcl/blob/main/hclsyntax/spec.md) — pour
ses expressions et ses fonctions.

| Où | YAML | HCL |
| --- | --- | --- |
| front matter | `---` … `---` | `---hcl` … `---` |
| configuration | `slides.yaml` | `slides.hcl`, cherché en premier |
| fichiers de *kinds* | `--kinds kinds.yaml`, `kinds: kinds.yaml` | tout fichier nommé `*.hcl` |
| fichiers de données | `--data data.yaml` ou `.json` | tout fichier nommé `*.hcl` |

Les commentaires de directive (`<!-- class: invert -->`) restent en YAML, comme dans Marp, et le
corps de la présentation reste un [gabarit Go]({{< relref "templates" >}}).

## Un front matter en HCL {#a-front-matter-in-hcl}

```hcl
---hcl
theme    = "gaia"
paginate = true
title    = "Quarterly review"
data = {
  company = "ACME"
  quarter = 3
}
header = format("%s — Q%d", data.company, data.quarter)
footer = "${upper(data.company)} · ${page()}/${total()}"
date   = formatdate("D MMMM YYYY", now())

kind "chapter" {
  extends = "section"
  class   = ["lead", "invert"]
}
---

# {{ .title }}
```

- Les **attributs** sont les clés YAML, sous les mêmes noms : les directives de Marp
  (`theme`, `paginate`, `header`, `footer`, `class`, `backgroundColor`…),
  les métadonnées (`title`, `author`…), `template`, `delims`, `kinds` (chemins de
  fichiers) et `data`, un objet.
- Les blocs **`kind "<name>" { … }`** définissent des [types de diapositives (kinds)]({{< relref "kinds" >}}),
  avec les directives locales et `extends`.
- Les **valeurs** parviennent aux directives sous les formes textuelles que leur donne YAML —
  `true` devient `"true"`, une liste de classes est une liste — si bien qu'une directive se comporte
  de la même façon quelle que soit la syntaxe qui l'a définie. `null` laisse une directive non définie.
- **`data`** conserve ses types : dans les gabarits, un nombre entier est un `int`,
  les autres nombres des `float64`, et listes et objets restent des listes et des maps : ainsi,
  `{{ add .quarter 1 }}` fonctionne.
- Une valeur HCL est définitive : les gabarits Go ne repassent pas dessus.

## Variables {#variables}

`data` contient les données du gabarit : le `data` du front matter lui-même, surchargé par
celui de `slides.hcl`, par les fichiers `--data` et par `--set`. `"${data.company}"` fonctionne
partout où `{{ .company }}` fonctionnerait.

Les attributs du front matter peuvent faire référence à `data.*`, mais **pas les uns aux autres** :
`footer = title` est une erreur. Les directives ponctuelles (`_class`…) ne sont pas des
identifiants HCL ; écrivez-les dans un commentaire.

## Fonctions {#functions}

| | |
| --- | --- |
| `page()`, `total()` | le numéro de page et celui de la dernière page, comme `{{ page }}` et `{{ total }}` : un pied de page peut ainsi paginer |
| `now()`, `timestamp()` | l'heure, au format RFC 3339, locale et UTC ; `formatdate` la met en forme |
| `env(name)` | une variable d'environnement, `""` si elle n'est pas définie |
| `file(path)` | un fichier texte UTF-8 d'au plus 1 Mio, relatif à la présentation ou au fichier de configuration |
| nombres | `abs ceil floor log max min parseint pow signum` |
| chaînes | `chomp format formatlist join lower regex regexall regexreplace replace split strlen strrev substr title trim trimprefix trimspace trimsuffix upper` |
| collections | `chunklist coalesce coalescelist compact concat contains distinct element flatten keys length lookup merge range reverse slice sort values zipmap` |
| encodages | `csvdecode jsondecode jsonencode` |
| dates | `formatdate timeadd` |

Toutes, sauf celles des quatre premières lignes, viennent de la bibliothèque standard de
[go-cty](https://github.com/zclconf/go-cty). Aucune n'accède au réseau.

## slides.hcl {#slideshcl}

```hcl
kinds  = "kinds.hcl"          # or a list of files
themes = ["themes/acme.css"]
data   = { company = "ACME" } # or file paths
html   = false
lang   = "en"

kind "quote" {
  extends = "blank"
  footer  = "${data.company}"
}
```

Un attribut inconnu est une erreur, et non une opération silencieusement ignorée. Un fichier de *kinds* contient
des blocs `kind` ; ses autres attributs sont ignorés, si bien que `slides.hcl` peut en tenir
lieu. Les attributs d'un fichier de données sont les clés des données :

```hcl
# data.hcl
company = "ACME"
year    = 2026
motto   = trimspace(file("motto.txt"))
```

## Erreurs {#errors}

Les erreurs sont les diagnostics de HCL, positionnés dans le fichier de la présentation lui-même :
`footer = data.compnay` sur la quatrième ligne de la présentation donne

```
deck.md:4:14: Unsupported attribute: This object does not have an attribute named "compnay".
```

## Marp et `slides expand` {#marp-and-slides-expand}

Marp lit lui aussi `---hcl` comme le début d'un front matter — sa règle de
front matter s'ouvre sur toute première ligne commençant par trois tirets — et, le bloc
n'étant pas un mapping YAML, n'en applique rien : la présentation s'affiche quand même dans Marp,
sans ses directives. Un test le vérifie par rapport à la sortie de Marp Core.

`slides expand` réécrit le front matter en YAML : la présentation développée est donc
du Marp ordinaire et s'affiche comme `slides` affiche l'original :

```yaml
---
theme: "gaia"
paginate: true
title: "Quarterly review"
header: "ACME — Q3"
footer: "ACME · 1/2"
date: "19 October 2026"
---
```

## Licences {#licences}

La prise en charge de HCL lie [hashicorp/hcl](https://github.com/hashicorp/hcl)
(MPL-2.0, utilisé sans modification comme bibliothèque) et
[zclconf/go-cty](https://github.com/zclconf/go-cty) (MIT).
