---
title: "go-slides-Dokumentation"
linkTitle: "Start"
type: docs
cascade:
  type: docs
description: "Marp-kompatible Foliensätze in reinem Go (CGO=0), mit Go-Vorlagen zur Interpolation und Kinds."
---

**`slides` macht aus einem Marp-Foliensatz eine Präsentation, in reinem Go.** Es liest
das Markdown, das auch [Marp](https://marp.app) liest – die Direktiven von Marpit, `---` zwischen
den Folien, `![bg]`-Hintergründe, Sprechernotizen, die Themes `default`, `gaia` und
`uncover` – und ergänzt zwei Dinge, die Marp nicht kennt:

- **[Go-Vorlagen]({{< relref "templates" >}})** im Foliensatz, die ausgewertet werden,
  bevor das Markdown geparst wird, sodass ein Foliensatz seine eigenen Folien erzeugen kann;
- **[Kinds (die Art einer Folie)]({{< relref "kinds" >}})**: Eine Folie gibt an, was sie ist
  (`title`, `section` …), statt immer wieder zu wiederholen, wie sie aussieht.

{{< callout type="warning" >}}
**v0.15.0.** Diese Seiten beschreiben
[go-slides/slides v0.15.0](https://github.com/go-slides/slides/releases/tag/v0.15.0)
und sind unter `/docs/0.15/` veröffentlicht.
{{< /callout >}}

## Was es heute kann {#what-it-does-today}

| | |
| --- | --- |
| `slides build` (der Standard) | der Foliensatz als eine HTML-Seite: jede Folie ein `<section>` im SVG-Rahmen, der sich dem Fenster anpasst, dazu das CSS des Themes und ein Tastatur-Viewer |
| `slides build --pdf`, `--images png`, `--image jpeg` | der Foliensatz als PDF (eine Seite pro Folie, mit Gliederung) oder als PNG-/JPEG-Bilder, ohne Browser: [go-webengine](https://github.com/go-webengine) rendert, [go-pdfkit](https://github.com/go-pdfkit) schreibt das PDF – seit v0.4.0 als Vektoren, mit auswählbarem Text: [Vektor-PDF]({{< relref "cli#vector-pdf" >}}) |
| Mathematik und Code | `$…$` und `$$…$$` setzt go-tex/math, Codeblöcke hebt go-rouge hervor, in HTML, Bildern und PDF: [Mathematik und Code]({{< relref "math-and-code" >}}) |
| Emoji | `:smile:` und 😄 werden zu eingebetteten Twemoji-Bildern, brauchen also kein Netzwerk und erscheinen in jeder Ausgabe: [Emoji]({{< relref "emoji" >}}) (seit v0.4.0) |
| `slides build --pptx` | der Foliensatz als PowerPoint-Datei mit Sprechernotizen und den Metadaten des Foliensatzes: in v0.4.0 ein Bild pro Folie wie bei marp-cli; seit v0.5.0 **editierbar** – Textfelder, Formen, Bilder, eingebettete Schriften ([Editierbares PPTX und ODP]({{< relref "cli#editable-pptx-and-odp" >}})) |
| `slides build --odp` | dasselbe als OpenDocument-Präsentation (`.odp`) |
| einpassende Überschriften | `# <!-- fit --> Title` füllt die Folienbreite in einer Zeile, auf der Seite und in jedem Export |
| `slides expand` | der Foliensatz als reines Marp-Markdown: Vorlagen ausgewertet, Kinds durch die Spot-Direktiven ersetzt, für die sie stehen |
| `slides notes` | die Sprechernotizen |
| `slides kinds`, `slides themes` | was ein Foliensatz benennen kann |
| `slides.yaml` | gemeinsame Kinds, Themes, Daten, HTML-Richtlinie und Sprache für einen Ordner voller Foliensätze |
| HCL | `---hcl`-Front-Matter, `slides.hcl`, Kinds- und Datendateien in HCL2, mit Ausdrücken und Funktionen: [HCL]({{< relref "hcl" >}}) |

Der Unterschied zu
marp-cli liegt im Weg: Es rendert Bilder und PDF, indem es einen Headless-Chrome steuert;
hier erledigen das Go-Bibliotheken.

## Seiten {#pages}

- [Erste Schritte]({{< relref "getting-started" >}})
- [Galerie]({{< relref "gallery" >}}): zwölf Beispiel-Foliensätze, zum Öffnen im Playground
- [Marp-Kompatibilität]({{< relref "marp" >}})
- [Vorlagen]({{< relref "templates" >}})
- [Kinds]({{< relref "kinds" >}})
- [HCL]({{< relref "hcl" >}})
- [Mathematik und Code]({{< relref "math-and-code" >}})
- [Emoji]({{< relref "emoji" >}})
- [Kommandozeile]({{< relref "cli" >}})
