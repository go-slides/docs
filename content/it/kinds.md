---
title: "Tipi di slide"
weight: 40
description: "Una slide dichiara cos'è: kind indica un insieme predefinito di direttive locali."
tags: [kinds]
---

Un **tipo** (*kind*) è un insieme predefinito e con nome di direttive locali: intestazione, piè di pagina,
numerazione, classe, sfondo. Una slide dichiara cos'è — `title`,
`section` — invece di ripetere che aspetto ha, e l'aspetto di una presentazione si cambia
in un unico punto.

## Usare un tipo {#using-a-kind}

```markdown
<!-- kind: section -->
# Part two

---

<!-- _kind: blank -->
![bg](photo.jpg)
```

- `kind:` si applica come una direttiva locale: a questa slide e a quelle successive.
- `_kind:` si applica come una direttiva spot: solo a questa slide.
- Cambiare tipo **azzera** prima ciò che il tipo precedente aveva impostato, quindi una classe o un
  piè di pagina non passa mai da un tipo al successivo.
- Le direttive scritte sulla slide stessa **prevalgono** sul tipo.

## Predefiniti {#built-in}

| Tipo | Imposta |
| --- | --- |
| `title` | `class: lead`, nessun numero di pagina, nessuna intestazione, nessun piè di pagina |
| `section` | `class: lead` |
| `content` | nulla: i valori predefiniti della presentazione |
| `blank` | nessun numero di pagina, nessuna intestazione, nessun piè di pagina |

Qualsiasi tipo può essere ridefinito.

## Definire i tipi {#defining-kinds}

Nel front matter:

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

`extends` fa partire un tipo dalle direttive di un altro. Un tipo accetta le
direttive locali e `extends`.

## Condividere i tipi {#sharing-kinds}

I tipi di un team vanno in un unico file che ogni presentazione nomina:

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

oppure `--kinds kinds.yaml` sulla riga di comando, oppure `kinds:` in `slides.yaml`.
L'ordine è: i tipi predefiniti, poi la riga di comando e il
file di configurazione, poi i `kinds:` della presentazione stessa.

`slides kinds talk.md` elenca i tipi che una presentazione può usare e cosa imposta ciascuno.

## Lasciare go-slides {#leaving-go-slides}

`slides expand` scrive ogni tipo come le direttive spot che rappresenta,
slide per slide, così il risultato è semplice Markdown Marp che qualsiasi strumento Marp
legge.
