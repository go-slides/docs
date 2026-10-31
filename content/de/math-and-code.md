---
title: "Mathematik und Code"
weight: 35
description: "Formeln zwischen Dollarzeichen werden gesetzt, Codeblöcke hervorgehoben, auf der HTML-Seite, in Bildern und im PDF."
tags: [math, code]
---

{{< callout type="warning" >}}
Mathematik und Code gibt es ab v0.2.0.
{{< /callout >}}

## Mathematik {#math}

`$…$` ist eine Inline-Formel und `$$…$$` eine abgesetzte Formel:

```markdown
The famous one is $E=mc^2$, and this is a display formula:

$$
\int_0^\infty e^{-x^2}\,dx = \frac{\sqrt{\pi}}{2}
$$
```

- Die Begrenzerregeln sind die von Marp Core, gelesen aus marp-core 4.4.0: Ein `$`
  mit folgendem Leerzeichen öffnet keine Formel, und eines mit folgender Ziffer
  schließt keine, sodass Preise wie `$5 and $10` Text bleiben.
- [go-tex/math](https://github.com/go-tex/math) setzt jede Formel in ein
  Inline-SVG in der Einheit `em`, das sich der Textgröße anpasst und auf der
  Grundlinie sitzt. Dasselbe SVG dient auf der HTML-Seite, in Bildern und im PDF.
- Eine Formel, die sich nicht parsen lässt, bleibt wie geschrieben sichtbar; der Fehler steht in
  ihrem `title`.

## Code {#code}

Codeblöcke werden von [go-rouge](https://github.com/go-rouge/rouge) eingefärbt,
das die Klassennamen von highlight.js schreibt, die die Themes von Marp formatieren:

````markdown
```go
fmt.Println("hello")
```
````

Eine Sprache, die go-rouge nicht kennt, wird als reiner Text geschrieben, wie es Marp Core tut.

### Zeilennummern {#line-numbers}

Ab v0.8.0 kann ein Block seine Zeilen nummerieren. Marp Core kennt keine Zeilennummern und
ignoriert, was hinter der Sprache eines Codeblocks steht; die Optionen stehen deshalb in geschweiften Klammern,
und ein Foliensatz, der sie verwendet, wird in Marp so gerendert, als gäbe es sie nicht:

````markdown
```go {lineNumbers}
func main() {
	fmt.Println("hello")
}
```

```go {lineNumbers=10}
// starts at 10
```
````

| Schreibweise | Bedeutung |
| --- | --- |
| `{lineNumbers}` | ab 1 nummerieren (Slidevs `{lines:true}` und Pandocs `{.numberLines}` werden ebenfalls gelesen) |
| `{lineNumbers=10}` | ab 10 nummerieren (auch `startLine=10`, Pandocs `startFrom="10"`) |
| `{lineNumbers=false}` | keine Nummern, was die Folie auch sagt |

Für alle Blöcke einer Folie und der folgenden Folien setzen Sie die Direktive:
`<!-- lineNumbers: true -->` (`_lineNumbers` für eine einzelne Folie). Ein Kind (die Art einer Folie) kann sie
einmal für eine ganze Gruppe von Folien setzen:

```yaml
kinds:
  listing: {lineNumbers: true}
```

Jede Nummer beginnt ihre Zeile innerhalb des Codekastens des Themes, durch eine Linie abgesetzt.
Sie bleibt auf der Grundlinie ihrer Zeile, auch wenn die Zeile höher ist, etwa
durch eine Markierung oder ein Emoji. Die Nummern sind in einer Farbe gedämpft, die aus der
Farbe des Codes und seinem Hintergrund gemischt ist, mit einem Kontrast von 4,5:1 oder mehr in den Themes
default, gaia und uncover, hell und dunkel. Sie sind auf der Seite **nicht auswählbar**
und werden von Screenreadern nicht vorgelesen. Im PDF werden sie gezeichnet, aber aus dem Text
ausgelassen, sodass ein Textextraktor nur den Code liest. In einer PowerPoint- oder
OpenDocument-Datei sind sie editierbarer Text am Anfang jeder Zeile. Ein Foliensatz,
der die Option nicht verwendet, wird wie zuvor gerendert, bis auf das Byte genau. Das Hervorheben
ausgewählter Zeilen wird noch nicht unterstützt.

### Markierungen im Code {#code-callouts}

Ab v0.8.0 kann eine Codezeile mit einer nummerierten Scheibe enden, die auf eine Notiz
unter dem Block verweist. Die Syntax ist die von
[Quarto](https://quarto.org/docs/authoring/code-annotation.html), die auch
die Kommentarform von Asciidoctor ist. Die Markierung ist ein Kommentar in der Sprache des Codes,
sodass der Code gültig bleibt, und die Notizen sind die geordnete Liste direkt
hinter dem Block:

````markdown
```go
func main() { // <1>
	x := load() // <2>
	fmt.Println(x)
}
```

1. The entry point.
2. Reads the input.
````

| Schreibweise | Bedeutung |
| --- | --- |
| `// <1>`, `# <1>`, `-- <1>`, `; <1>`, `% <1>`, `' <1>` | ein Zeilenkommentar, der nur die Markierung enthält: beide werden durch die Scheibe ersetzt |
| `<!-- <1> -->`, `/* <1> */` | dasselbe für Sprachen mit umschließenden Kommentaren |
| `// <1> <2>` | zwei Scheiben in einer Zeile |
| `x = 1 # count <1>` | der Text des Kommentars bleibt, und die Scheibe folgt ihm |

Die Markierungen werden nur dann zu Scheiben, wenn die Liste direkt hinter dem Block bei 1 beginnt
und für jede Nummer einen Eintrag hat. Andernfalls bleibt der Block wie geschrieben. Die
Scheiben sind Ringe in der Farbe des Codes, in den Themes default, gaia und uncover,
hell oder dunkel. Die Einträge der Liste werden mit denselben Scheiben nummeriert.

Die Scheiben sind **nicht auswählbar**. Kopiert man den Code von der Seite, erhält man den Code
ohne seine Markierungen. Im PDF sind sie Vektoren, die aus dem Text ausgelassen werden, und in einer
PowerPoint- oder OpenDocument-Datei sind sie editierbare Formen. Markierungen lassen sich
mit Zeilennummern kombinieren. Marp Core kennt keine Markierungen (marp-core 4.4.0); in Marp erscheinen daher die
Kommentare und die nummerierte Liste wie geschrieben. Ein Foliensatz, der sie nicht verwendet,
wird wie zuvor gerendert, bis auf das Byte genau.

Ab v0.9.0 finden sich eine Scheibe und ihre Notiz auf der HTML-Seite gegenseitig. Fährt man mit der Maus
über eine Scheibe oder erreicht sie mit der Tabulatortaste, färbt sich ihre Notiz ein, und eine Notiz bewirkt dasselbe für ihre
Scheibe. Das entspricht `code-annotations: hover` von Quarto, nur erledigt es hier allein das Stylesheet,
ohne Skript. Es funktioniert daher gleich in der Sandbox-Vorschau des
Playgrounds und in der Vorschau von VS Code, die beide Skripte einschränken. Jede Scheibe wird durch ihre Notiz beschrieben,
sodass ein Screenreader die Notiz vorliest, wenn er die Scheibe erreicht. Die Färbung blendet nur dann ein, wenn der Leser
keine reduzierte Bewegung verlangt hat. Bilder, PDF-, PowerPoint- und OpenDocument-Dateien
bleiben Byte für Byte unverändert.

## Schriften {#fonts}

Lato Regular, Bold, Black und Kursiv sowie Roboto Mono Regular, Bold und
Kursiv sind eingebettet, sodass die Überschriften von gaia ihre Strichstärke erhalten und kursiver Text
kursiv gerendert wird.
