---
title: "Vorlagen"
weight: 30
description: "Interpolation mit Go-text/template in einem Foliensatz: Werte, erzeugte Folien, Seitenzahlen, Includes."
tags: [templates]
---

Der Foliensatz ist ein Go-[`text/template`](https://pkg.go.dev/text/template),
das **ausgewertet wird, bevor** das Markdown geparst wird. Alles, was die Go-Template-
Sprache bietet – Aktionen, Pipelines, `if`, `range`, `with`, Variablen –, funktioniert
in einem Foliensatz.

## Daten {#data}

Der Punkt der Vorlage enthält die Schlüssel des Front-Matters, dann die Schlüssel seiner
`data:`-Map, dann die Daten, die auf der Kommandozeile oder in `slides.yaml` angegeben werden
– jede Stufe überschreibt die vorherige.

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

Von der Kommandozeile:

| Flag | Liefert |
| --- | --- |
| `--data file.yaml` | eine YAML- oder JSON-Datei mit Daten (wiederholbar) |
| `--set key=value` | einen Wert (wiederholbar) |

## Folien erzeugen {#generating-slides}

Der gesamte Rumpf ist eine einzige Vorlage, sodass ein `range`, das `---` schreibt, Folien schreibt:

```markdown
{{ range .releases }}
---

# {{ .version }}

{{ range .changes }}- {{ . }}
{{ end }}
{{ end }}
```

## Seite und Gesamtzahl {#page-and-total}

`{{ page }}` und `{{ total }}` sind die Nummer der Folie und die Anzahl der
Folien. Sie sind erst bekannt, nachdem der Foliensatz aufgeteilt wurde, und werden deshalb
erst dann eingesetzt – ein erzeugter Foliensatz zählt die Folien, die er erzeugt hat.

```markdown
<!-- footer: "{{ page }} / {{ total }}" -->
```

## Include {#include}

`{{ include "parts/agenda.md" }}` fügt eine andere Datei ein, relativ zum
Foliensatz.

## Code bleibt unberührt {#code-is-left-alone}

Codeblöcke und Inline-Code-Spans werden beiseitegelegt, bevor die Vorlage
läuft, und unverändert wiederhergestellt, sodass eine Folie über Go-Vorlagen
`{{ .Name }}` zeigen kann, ohne es zu maskieren.

## Grenzen {#limits}

Die Ausführung einer Vorlage hat ein Schrittbudget, damit ein Foliensatz, dessen
Vorlagen `range` über große Datenlisten verschachteln, den Renderer nicht
minutenlang beschäftigen kann. Jenseits von 1.048.576 Schritten (Schleifendurchläufe
und Vorlagenaufrufe, gemeinsam mit den `include`s einer Expansion) oder 10 s wird der
Foliensatz mit der Diagnose `template.limit` abgelehnt, statt zu hängen. Ein Schritt
kostet etwa 260 ns, sodass gewöhnliche Foliensätze weit unter der Grenze bleiben.
Seit v0.17.1.

## Abschalten {#turning-it-off}

`--no-template` liest den Foliensatz als reines Marp: `{{` ist Text.
