---
title: "Gabarits"
weight: 30
description: "Interpolation de gabarits Go text/template dans une présentation : valeurs, diapositives générées, numéros de page, inclusions."
tags: [templates]
---

La présentation est un gabarit Go [`text/template`](https://pkg.go.dev/text/template),
évalué **avant** l'analyse du Markdown. Tout ce qu'offre le langage de gabarits
de Go — actions, pipelines, `if`, `range`, `with`, variables — fonctionne
dans une présentation.

## Données {#data}

Le point du gabarit contient les clés du front matter, puis les clés de son
mapping `data:`, puis les données fournies en ligne de commande ou dans `slides.yaml`
— chacune remplaçant la précédente.

```markdown
---
marp: true
title: Quarterly review
data:
  team: Platform
---

# {{ .title }}

Presented by the {{ .team }} team.
```

En ligne de commande :

| Option | Fournit |
| --- | --- |
| `--data file.yaml` | un fichier de données YAML ou JSON (répétable) |
| `--set key=value` | une valeur (répétable) |

## Générer des diapositives {#generating-slides}

Le corps entier est un seul gabarit : un `range` qui écrit `---` écrit donc des diapositives :

```markdown
{{ range .releases }}
---

# {{ .version }}

{{ range .changes }}- {{ . }}
{{ end }}
{{ end }}
```

## Numéro de page et total {#page-and-total}

`{{ page }}` et `{{ total }}` sont le numéro de la diapositive et le nombre de
diapositives. Ils ne sont connus qu'une fois la présentation découpée : ils sont donc renseignés
à ce moment-là — une présentation générée compte les diapositives qu'elle a générées.

```markdown
<!-- footer: "{{ page }} / {{ total }}" -->
```

## Inclusion {#include}

`{{ include "parts/agenda.md" }}` insère un autre fichier, relatif à la
présentation.

## Le code n'est pas touché {#code-is-left-alone}

Les blocs de code délimités et les extraits de code en ligne sont mis de côté avant l'exécution du gabarit
puis restitués tels quels : une diapositive consacrée aux gabarits Go peut ainsi montrer
`{{ .Name }}` sans l'échapper.

## Désactivation {#turning-it-off}

`--no-template` lit la présentation comme du Marp ordinaire : `{{` est du texte.
