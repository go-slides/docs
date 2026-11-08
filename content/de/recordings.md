---
title: "Terminal-Aufnahmen"
weight: 38
description: "Eine asciinema-Aufnahme in einer Präsentation: in einem kleinen Player der HTML-Seite abgespielt, im PDF, in PowerPoint und OpenDocument als Vektortext gezeichnet."
tags: [asciinema, terminal]
---

{{< callout type="warning" >}}
Terminal-Aufnahmen sind im Hauptzweig nach v0.17.1 enthalten, aber noch in keinem Release.
{{< /callout >}}

Eine [asciinema](https://asciinema.org)-Aufnahme ([asciicast v2](https://docs.asciinema.org/manual/asciicast/v2/) oder [v3](https://docs.asciinema.org/manual/asciicast/v3/)) ist ein Bild, dessen Datei auf `.cast` endet:

```markdown
![cast](demo.cast)
![The failing build](demo.cast#t=12.5)
```

## Syntax {#syntax}

Die Datei wird aus dem Ordner der Präsentation gelesen: eine URL, ein absoluter Pfad oder ein Pfad, der den Ordner verlässt, wird abgelehnt. Der Alternativtext benennt die Aufnahme für Screenreader; `cast` allein steht für den eigenen Titel der Aufnahme. `#t=12.5`, oder `#t=1:23` für Minuten und Sekunden, wählt den Moment, den die Ausgaben zeigen, die sich nicht bewegen können; ohne diese Angabe ist es das Ende der Aufnahme.

## Was jede Ausgabe zeigt {#outputs}

- **HTML** (die Seite und die Vorschau des Playgrounds): Die Aufnahme wird in einem kleinen Player von go-slides abgespielt, der in die Seite eingebettet ist: Wiedergabe und Pause, eine Geschwindigkeitstaste und eine Positionsleiste. Er lädt von nirgendwoher ein Skript und braucht kein Netzwerk.
- **PDF, PowerPoint und OpenDocument** (`--pdf`, `--pptx`, `--odp`): das Terminal in dem Moment, den `#t=` nennt, als Vektortext in der Festbreitenschrift des Themes, mit den 16, 256 und direkten Farben, Fett, Kursiv, Unterstreichung und invertiertem Text. Der Text lässt sich markieren und durchsuchen. Es ist nie ein Bild. Fett und Kursiv werden gezeichnet, wo die Festbreitenschrift des Themes sie hat (Roboto Mono im Theme gaia).

## Der Player {#player}

Er spielt einmal, wenn die Aufnahme zum ersten Mal auf dem Bildschirm ist. Wenn der Browser weniger Bewegung wünscht (`prefers-reduced-motion`), startet er pausiert, zeigt den Moment, den `#t=` nennt, und spielt, wenn Sie auf Wiedergabe drücken. Seine Tasten und Tastenkürzel blättern die Folie nicht weiter.

## Grenzen {#limits}

Eine Aufnahme ist nicht vertrauenswürdige Eingabe und daher begrenzt: die Datei (16 MiB), die Ereignisse (500.000), die Länge (3 Stunden), das Terminal (400 mal 200 Zellen) und die Arbeit, die der Emulator leisten darf. Eine Aufnahme über einer Grenze oder keine wohlgeformte asciicast-Datei wird abgelehnt: Der Build läuft weiter, die Folie zeigt an ihrer Stelle einen Hinweis, und das Problem ist eine Warnung mit einem Code: `cast.syntax`, `cast.limit`, `cast.version` oder `cast.path`.

## Im Playground {#playground}

Fügen Sie dem Projekt im Dateien-Bereich eine `.cast`-Datei hinzu, per Ablegen oder über die Dateiauswahl. Sie kommt nach `media/` und wird als Bild eingefügt. Sie ist Text, den die Präsentation liest, und zählt daher mit den übrigen Textdateien des Projekts, die auf 2 MiB begrenzt sind. Probieren Sie es im [Playground](https://go-slides.github.io/playground/).
