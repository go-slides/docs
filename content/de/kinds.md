---
title: "Kinds"
weight: 40
description: "Eine Folie sagt, was sie ist: kind benennt eine Vorgabe lokaler Direktiven."
tags: [kinds]
---

Ein **Kind** (die Art einer Folie) ist eine benannte Vorgabe lokaler Direktiven: Kopfzeile, Fußzeile,
Seitennummerierung, Klasse, Hintergrund. Eine Folie gibt an, was sie ist – `title`,
`section` –, statt zu wiederholen, wie sie aussieht, und das Erscheinungsbild eines Foliensatzes ändert sich
an einer einzigen Stelle.

## Ein Kind verwenden {#using-a-kind}

```markdown
<!-- kind: section -->
# Part two

---

<!-- _kind: blank -->
![bg](photo.jpg)
```

- `kind:` wirkt wie eine lokale Direktive: auf diese Folie und die folgenden.
- `_kind:` wirkt wie eine Spot-Direktive: nur auf diese Folie.
- Beim Wechsel eines Kinds wird zuerst **zurückgesetzt**, was das vorherige Kind gesetzt hat, sodass eine Klasse oder eine
  Fußzeile nie von einem Kind ins nächste durchsickert.
- Direktiven, die auf der Folie selbst stehen, **haben Vorrang** vor dem Kind.

## Eingebaut {#built-in}

| Kind | Setzt |
| --- | --- |
| `title` | `class: lead`, keine Seitennummer, keine Kopfzeile, keine Fußzeile |
| `section` | `class: lead` |
| `content` | nichts: die Standardwerte des Foliensatzes |
| `blank` | keine Seitennummer, keine Kopfzeile, keine Fußzeile |

Jedes davon lässt sich neu definieren.

## Kinds definieren {#defining-kinds}

Im Front-Matter:

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

`extends` lässt ein Kind mit den Direktiven eines anderen beginnen. Ein Kind akzeptiert die
lokalen Direktiven und `extends`.

## Kinds teilen {#sharing-kinds}

Die Kinds eines Teams gehören in eine Datei, die jeder Foliensatz benennt:

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

oder `--kinds kinds.yaml` auf der Kommandozeile, oder `kinds:` in `slides.yaml`.
Die Reihenfolge ist: die eingebauten Kinds, dann die Kommandozeile und die
Konfigurationsdatei, dann das eigene `kinds:` des Foliensatzes.

`slides kinds talk.md` listet die Kinds auf, die ein Foliensatz verwenden kann, und was jedes davon setzt.

## go-slides verlassen {#leaving-go-slides}

`slides expand` schreibt jedes Kind als die Spot-Direktiven aus, für die es steht,
Folie für Folie, sodass das Ergebnis reines Marp-Markdown ist, das jedes Marp-Werkzeug
lesen kann.
