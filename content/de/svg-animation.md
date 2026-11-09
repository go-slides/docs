---
title: "Animiertes SVG"
weight: 40
description: "Ein SVG-Bild mit SMIL- oder CSS-Animationen: vom Browser in HTML abgespielt, animierte Formen in PowerPoint und OpenDocument, ein animiertes GIF, wo diese es nicht ausdrücken können, der Anfangszustand in PDF und Bildern."
tags: [svg, animation]
---

Ein SVG-Bild, das sich selbst animiert, wird wie jedes Bild geschrieben:

```markdown
![](spinner.svg)
![w:400](chart.svg)
```

## Was gelesen wird {#read}

Die Animationen, die in der SVG-Datei selbst stehen: SMIL (`animate`, `animateTransform`, `animateMotion`, `animateColor`, `set`, mit `begin`, `dur`, `repeatCount`, `fill`, `calcMode`, `keyTimes` und `keySplines`) und CSS (`@keyframes` mit den `animation`-Eigenschaften, `transform-origin` und `transform-box`). Was nicht verstanden wird, etwa ein Ereignis wie `begin="click"`, wird ausgelassen und als Warnung gemeldet: nichts wird stillschweigend verworfen.

## Was jede Ausgabe zeigt {#outputs}

- **HTML**: Der Browser spielt die Animationen selbst ab.
- **PowerPoint** (`--pptx`) und **OpenDocument** (`--odp`): Jedes animierte Element ist eine eigene Form, und seine Animationen starten mit der Folie: Bewegung, Drehung und Skalierung, Deckkraft, ein Wechsel zwischen zwei Farben. Die Bewegung eines Elements wird zu den eigenen Zeitpunkten der Animation abgetastet, sodass Beschleunigungskurven, mehrere Schlüsselbilder und eine Drehung um einen anderen Punkt exakt sind. Wiederholungen (auch `indefinite`), Verzögerungen und Umkehrungen werden übernommen.
- **Einzelbilder**: Lässt sich eine Animation so nicht ausdrücken (ein Pfad, der sich verformt, eine Linie, die sich zeichnet, eine Scherung), wird das ganze Bild zu den eigenen Zeitpunkten der Animation gerendert und als animiertes GIF eingebettet, das PowerPoint und Impress abspielen. Das ist begrenzt: 12 Bilder pro Sekunde, höchstens 240 Bilder, 16 Millionen Pixel über alle Bilder.
- **PDF** (`--pdf`) und Bilder (`--images`): Sie können nicht animieren. Die Zeichnung erscheint im Anfangszustand, ohne JavaScript und ohne `RichMedia`, und eine Warnung nennt die animierten Bilder.

Jedes animierte Bild erhält eine Warnzeile mit dem Code `svg-animation`, die sagt, wie es übernommen wurde: nativ oder als Einzelbilder und warum.

## Unterstützung je Ausgabe {#support}

| Merkmal | HTML | PowerPoint | OpenDocument | PDF und Bilder |
|---|---|---|---|---|
| Verschiebung, Bewegung entlang eines Pfads | abgespielt | nativ | nativ | Anfangszustand |
| Drehung und Skalierung um einen beliebigen Punkt | abgespielt | nativ | nativ | Anfangszustand |
| Radius eines Kreises, Größe eines Rechtecks (ohne Kontur) | abgespielt | nativ | nativ | Anfangszustand |
| Deckkraft eines Elements aus einer Form | abgespielt | nativ | nativ | Anfangszustand |
| Farbe einer Füllung oder Kontur, von einer Farbe zur anderen | abgespielt | nativ | nativ | Anfangszustand |
| Wiederholungen, Verzögerungen, Umkehrungen, Halten am Ende | abgespielt | nativ | nativ | Anfangszustand |
| Beschleunigungskurven, mehrere Schlüsselbilder | abgespielt | nativ (abgetastet) | nativ (abgetastet) | Anfangszustand |
| Stufen, diskrete Wechsel | abgespielt | Einzelbilder | Einzelbilder | Anfangszustand |
| Pfadverformung, Strichversatz, Positions- und Größenattribute, Verläufe | abgespielt | Einzelbilder | Einzelbilder | Anfangszustand |
| Scherung, Spiegelung, mehrere Uhren an einem Element, verschachtelte animierte Elemente | abgespielt | Einzelbilder | Einzelbilder | Anfangszustand |
| Ereignisse | abgespielt | nicht übernommen (Warnung) | nicht übernommen (Warnung) | Anfangszustand |

## Lokale Dateien {#local-files}

Ein lokales Bild wird nur mit `--allow-local-files` aus dem Ordner der Präsentation gelesen, wie in marp-cli. Ohne diese Option bleibt das Bild aus den Exporten heraus, und der Build sagt es mit `local-file-blocked`, unter Nennung der Datei, ihrer Folien und der Option. Eine Datei, die nicht existiert, wird als `local-file-missing` gemeldet.

## Grenzen {#limits}

Ein SVG ist nicht vertrauenswürdige Eingabe und daher begrenzt: 8 MiB SVG, 50 000 Elemente, 2 000 Animationen, 256 Schlüsselbilder je Animation, 60 Sekunden Zeitachse und 240 Bilder. Eine Zeichnung über einer Grenze wird mit einer Warnung abgelehnt, und der Build läuft weiter.

## Was geprüft wurde {#checked}

LibreOffice Impress (macOS) öffnet die PowerPoint- und OpenDocument-Dateien und behält jede Animation, wenn es sie erneut speichert, in beiden Formaten. Keynote (macOS) importiert die Formen und das animierte GIF; seine Animationen wurden nicht untersucht. PowerPoint: noch nicht geprüft.
