---
title: "Marp-Kompatibilität"
linkTitle: "Marp"
weight: 20
description: "Das Markdown, das slides liest, ist das von Marpit und Marp Core, kein Dialekt davon."
tags: [marp]
---

Ein für [Marp](https://marp.app) geschriebener Foliensatz ist ein go-slides-Foliensatz. Die folgenden
Regeln stammen von [Marpit](https://marpit.marp.app) und
[Marp Core](https://github.com/marp-team/marp-core); wo go-slides
abweicht, sagt diese Seite es.

## Folien {#slides}

Eine horizontale Linie – `---` allein auf einer Zeile, so wie Marpit trennt – beginnt eine
neue Folie, ebenso eine Überschrift der Ebene, die die Direktive `headingDivider`
nennt. Das Front-Matter zwischen den ersten beiden `---` enthält globale Direktiven
und die Daten des Foliensatzes.

## Direktiven {#directives}

| Art | Wo | Gilt für |
| --- | --- | --- |
| global | Front-Matter oder `<!-- theme: gaia -->` an beliebiger Stelle | den ganzen Foliensatz: `theme`, `style`, `headingDivider`, `lang` … |
| lokal | `<!-- paginate: true -->` | diese Folie und die folgenden |
| Spot | `<!-- _class: lead -->` (Unterstrich) | nur diese Folie |

Die lokalen Direktiven sind die von Marpit: `paginate`, `header`, `footer`, `class`,
`backgroundColor`, `backgroundImage`, `backgroundPosition`,
`backgroundRepeat`, `backgroundSize`, `color`. Direktiven werden als
nachsichtiges YAML geparst, wie es Marp Core tut; `--strict-yaml` parst sie so, wie es
reines Marpit tut.

## Bilder {#images}

`![bg](image.jpg)` macht ein Bild zum Hintergrund der Folie, mit den Schlüsselwörtern
von Marp: `bg contain`, `bg cover`, `bg 50%`, `bg left`/`bg right` für einen geteilten
Hintergrund, mehrere `bg`-Bilder nebeneinander und die Bildfilter
(`blur`, `sepia`, `grayscale` …). Größenangaben für Inline-Bilder (`w:200px`, `h:`) werden
so gelesen, wie Marp sie liest.

## Rund um die Folie {#around-the-slide}

- **Kopf- und Fußzeile**: die Direktiven `header:` und `footer:`, in Markdown.
- **Seitennummerierung**: `paginate: true`.
- **Sprechernotizen**: Ein HTML-Kommentar, der keine Direktive ist, ist eine Notiz.
  `slides notes` gibt sie aus.
- **Schrittweise Listen**: Eine mit `*` (oder `1)`) geschriebene Liste erscheint Eintrag für Eintrag,
  wie in Marp.
- **`<style scoped>`**: CSS nur für eine Folie.
- **Rohes HTML**: standardmäßig die Allowlist von Marp Core; `--html` erlaubt alles.
- **Emoji**: `:shortcodes:` und Unicode-Emoji werden zu Twemoji-Bildern, wie in
  Marp Core, jedoch eingebettet statt von einem CDN geladen: siehe
  [Emoji]({{< relref "emoji" >}}).

## Themes {#themes}

`default`, `gaia` und `uncover` sind eingebaut und werden mit `theme:` gewählt. Ein
eigenes Theme ist eine CSS-Datei mit dem Marp-Kommentar `/* @theme name */`,
die Sie mit `--theme-set` angeben oder in `slides.yaml` unter `themes:` aufführen.

## Wo go-slides abweicht {#where-go-slides-differs}

- **Vorlagen.** `{{ }}` im Rumpf ist eine Aktion einer Go-Vorlage (siehe
  [Vorlagen]({{< relref "templates" >}})). Ein Foliensatz, der `{{` außerhalb eines Codeblocks als Text
  zeigt, braucht `--no-template` oder den Text maskiert, so wie Go-Vorlagen
  ihn maskieren. Codeblöcke und Code-Spans bleiben unberührt.
- **Kinds (die Art einer Folie).** `kind:` und `_kind:` sind Direktiven, die Marp nicht kennt (siehe
  [Kinds]({{< relref "kinds" >}})). `slides expand` ersetzt sie durch die
  Direktiven, für die sie stehen, sodass der expandierte Foliensatz reines Marp ist.
- **Ausgaben.** HTML, PNG, JPEG, PDF, PPTX und ODP, über Go-Bibliotheken statt über
  einen Headless-Browser (zur bekannten Einschränkung siehe [Kommandozeile]({{< relref "cli#images-pdf-pptx-and-odp" >}})).
  In v0.4.0 enthält ein PPTX wie bei marp-cli Bilder; seit
  v0.5.0 sind seine Folien editierbare Textfelder und Formen, und `--pptx-raster` behält
  die Bilder bei.
- **Einpassende Überschriften.** `# <!-- fit --> Title` wird unterstützt: Die Überschrift füllt
  die Folienbreite in einer Zeile, auf der Seite und in jedem Export.
- **Mathematik und Code.** `$…$` und `$$…$$` folgen den Begrenzerregeln von Marp Core und
  werden von go-tex/math gesetzt; Codeblöcke hebt go-rouge mit den Klassennamen von
  highlight.js hervor (siehe [Mathematik und Code]({{< relref "math-and-code" >}})).
