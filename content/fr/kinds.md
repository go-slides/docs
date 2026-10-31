---
title: "Kinds"
weight: 40
description: "Une diapositive dit ce qu'elle est : kind désigne un préréglage de directives locales."
tags: [kinds]
---

Un **kind** (le type d'une diapositive) est un préréglage nommé de directives locales : en-tête, pied de page,
pagination, classe, arrière-plan. Une diapositive dit ce qu'elle est — `title`,
`section` — au lieu de répéter à quoi elle ressemble, et l'apparence d'une présentation se modifie
à un seul endroit.

## Utiliser un kind {#using-a-kind}

```markdown
<!-- kind: section -->
# Part two

---

<!-- _kind: blank -->
![bg](photo.jpg)
```

- `kind:` s'applique comme une directive locale : à cette diapositive et aux suivantes.
- `_kind:` s'applique comme une directive ponctuelle : à cette diapositive seulement.
- Changer de *kind* **efface** d'abord ce que le précédent avait défini : une classe ou un
  pied de page ne déborde donc jamais d'un *kind* sur le suivant.
- Les directives écrites sur la diapositive elle-même **l'emportent** sur le *kind*.

## Kinds intégrés {#built-in}

| Kind | Définit |
| --- | --- |
| `title` | `class: lead`, ni numéro de page, ni en-tête, ni pied de page |
| `section` | `class: lead` |
| `content` | rien : les valeurs par défaut de la présentation |
| `blank` | ni numéro de page, ni en-tête, ni pied de page |

Chacun d'eux peut être redéfini.

## Définir des kinds {#defining-kinds}

Dans le front matter :

```yaml
---
marp: true
kinds:
  chapter:
    class: invert
    backgroundColor: "#1e1b4b"
    paginate: false
  quote:
    extends: blank
    class: quote
---
```

`extends` fait partir un *kind* des directives d'un autre. Un *kind* accepte les
directives locales et `extends`.

## Partager des kinds {#sharing-kinds}

Les *kinds* d'une équipe ont leur place dans un fichier unique que chaque présentation désigne :

```yaml
# kinds.yaml
chapter:
  class: invert
  paginate: false
```

```yaml
---
marp: true
kinds: kinds.yaml        # or a list of files
---
```

ou `--kinds kinds.yaml` en ligne de commande, ou `kinds:` dans `slides.yaml`.
L'ordre est le suivant : les *kinds* intégrés, puis la ligne de commande et le
fichier de configuration, puis les `kinds:` de la présentation elle-même.

`slides kinds talk.md` liste les *kinds* qu'une présentation peut utiliser et ce que chacun définit.

## Quitter go-slides {#leaving-go-slides}

`slides expand` écrit chaque *kind* sous la forme des directives ponctuelles qu'il représente,
diapositive par diapositive : le résultat est du Markdown Marp ordinaire, que tout outil Marp
sait lire.
