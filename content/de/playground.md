---
title: "Playground"
weight: 15
description: "Schreiben Sie einen Foliensatz im Browser und sehen Sie, wie die Folien folgen: go-slides, nach WebAssembly kompiliert, mit Vervollständigung, Hilfe und Beispielen."
tags: [playground, browser]
---

**[Playground öffnen →](https://go-slides.github.io/de/playground/)**

Der Playground ist go-slides selbst, v0.15.0, nach WebAssembly kompiliert und
in Ihrem Browser ausgeführt. Tippen Sie links einen Foliensatz; rechts
folgen die Folien. Nichts wird an einen Server gesendet: weder der Foliensatz noch ein Link zum Teilen.

Die Oberfläche des Playgrounds (Menüband, Hilfe, Meldungen) spricht die Sprache
der Seite: Öffne `/de/playground/` für Deutsch, und ebenso für die zwölf
Sprachen dieser Website. Die Sprachauswahl der Seite ändert sie, und deine Wahl
bleibt in diesem Browser gespeichert; ein Link, der eine Sprache nennt,
gewinnt immer. Die Beispielpräsentationen bleiben englisch. Ukrainisch,
Russisch, Chinesisch, Japanisch und Koreanisch werden mit einer Teilmenge von
Noto Sans (0,06 bis 2,3 MiB) dargestellt, die die Seite nur für diese Sprachen
herunterlädt.

## Was er bietet {#what-it-has}

- **Live-Vorschau** jeder Folie, außerhalb des Threads der Seite gerendert (ein Web
  Worker), sodass das Tippen nie darauf wartet. Klicken Sie auf eine Folie, um den Editor
  zu ihrer Zeile springen zu lassen; bewegen Sie den Cursor, und die Vorschau scrollt zu ihrer Folie.
- **Bildschirmpräsentation**: die Präsentation selbst, mit Tastatur,
  Bedienelementen auf dem Bildschirm und den Übergängen.
- **Referentenansicht**: **Present** (oder `p` in der Bildschirmpräsentation) öffnet sie in einem
  zweiten Fenster: die aktuelle Folie, der nächste Schritt, die Sprechernotizen (die
  `<!-- comments -->` des Foliensatzes), eine Uhr und ein Timer. Die Pfeiltasten, die Leertaste,
  Pos1 und Ende funktionieren in beiden Fenstern, und jedes folgt dem anderen, auch
  wenn Sie den Foliensatz bearbeiten. Öffnet sich das Fenster nicht, erlauben Sie Pop-ups für
  die Website. Die Übersichtsansicht (`o`) gibt es nur auf der heruntergeladenen Seite.
- **Vervollständigung** beim Tippen: Direktivenschlüssel (Spot-Direktiven mit `_` in Kommentaren),
  die Werte von `kind`, `theme`, `size`, `paginate` sowie Vorlagenfunktionen
  und Datenfelder innerhalb von `{{ }}`.
- **Hilfe** zu dem, was unter dem Cursor steht: eine Direktive, ein Kind (die Art einer Folie; was es setzt),
  eine Vorlagenfunktion. Der Hilfebereich verlinkt auch auf diese Seiten.
- **Probleme** in ihrer Zeile: Fehler in Vorlagen, YAML und HCL werden im Editor hervorgehoben
  und darunter aufgelistet, und die Vorschau behält den letzten Foliensatz,
  der sich rendern ließ.
- **Beispiele**: ein Willkommens-Foliensatz, Vorlagen, die Folien erzeugen, Kinds,
  ein HCL-Front-Matter, Themes und Spot-Direktiven sowie Mathematik, Bilder und
  Hintergründe (`![bg]`, ein geteiltes `![bg left:40%]`, Filter).
- **Dateien**: Der Bereich links vom Editor ist das Projekt des Foliensatzes, so angeordnet,
  wie der [Befehl]({{< relref "cli" >}}) und die
  [VS-Code-Erweiterung]({{< relref "vscode" >}}) es lesen: `deck.md`,
  `slides.yaml`, `themes/`, `kinds/`, `media/`.
  - Klicken Sie auf eine Datei, um sie zu bearbeiten. Der Editor färbt sie nach ihrer Sprache ein
    (Markdown, YAML, HCL, JSON, CSS) und markiert ihre Probleme.
  - Rechtsklick (oder Umschalt+F10) für Neue Datei, Neuer Ordner, Umbenennen (F2),
    Löschen (fragt vorher nach) und Neues Theme / Neues Kind. Ziehen Sie eine Datei auf einen Ordner,
    um sie zu verschieben. Die Tastaturbedienung folgt dem WAI-ARIA-Tree-Muster, einschließlich
    Type-ahead.
  - **Project ▸ New theme** erzeugt `themes/my-gaia.css` (oder auf Basis von default,
    uncover): ein `@import` des eingebauten Themes und die Regeln, die Sie ergänzen. Es
    erscheint in der Theme-Liste. **New kind** und **New data file** beginnen ebenfalls mit
    Vorlagen. Jedes wird in `slides.yaml` aufgeführt, sodass `slides build deck.md`
    es findet.
  - **Project ▸ Download project (.zip)** sichert alles. Entpackt, baut
    `slides build deck.md` es. **Open project** nimmt eine `.zip`-Datei oder einen
    Ordner wieder entgegen. Ein Archiv, das seinen Ordner verlässt (`..`, absolute Pfade),
    mehr als 256 Dateien oder 8 MiB in einer Datei enthält oder sich auf mehr als das 100-Fache
    entpackt, wird abgelehnt.
  - Alles wird zusammen mit dem Entwurf in diesem Browser gespeichert. Ein Projekt fasst
    bis zu 64 MiB, und die Zeile unter den Dateien zeigt, wie viel belegt ist.
- **Medien**: Ziehen Sie ein Bild, ein Video oder einen Ton auf den Editor oder fügen Sie es ein, und es
  landet in `media/`, und der Foliensatz benennt es: `![](media/photo.png)`. Ziehen Sie
  Dateien auf einen Ordner des Bereichs, um sie dort abzulegen. Der Foliensatz bleibt klein und
  lesbar, und nichts wird hochgeladen. Die Sandbox der Vorschau lädt sie,
  ohne geöffnet zu werden (die Seite reicht ihr die Bytes). **Save as ▸ HTML
  page** legt sie in die Seite. Auch `https:`-Bilder funktionieren, aber der Browser
  holt sie von ihrem Server, der dann Ihre Adresse sieht. Die
  Vorschau sendet keinen Referrer.
- **Größenveränderbare Bereiche**: Der Dateibereich, der Editor, die Folienminiaturen, die
  Vorschau und die Hilfe sind durch Leisten getrennt, die Sie ziehen können (der Mauszeiger
  wird darüber zum Größenänderungs-Cursor). Der Bereich neben der Leiste erhält den Platz,
  die anderen behalten ihre Größe. Doppelklicken Sie auf eine Leiste oder drücken Sie die Eingabetaste darauf, um den Bereich
  zu schließen und wieder zu öffnen. Per Tastatur wechselt F6 zwischen
  den Leisten, die Pfeiltasten verschieben eine um 16 px, Pos1 und Ende bringen den davorliegenden Bereich
  auf seine kleinste und größte Größe, und Escape führt zurück in den Editor.
  Der Reiter View schaltet Bereiche weiterhin ab. Die von Ihnen eingestellten Größen werden in diesem
  Browser mit dem Entwurf gespeichert; die Standardgrößen (190, 168 und 280 px, wobei sich Editor
  und Vorschau den Rest im Verhältnis 6 zu 5 teilen) werden nie gespeichert.
- **Format**: Der Reiter Format bearbeitet das Markdown für Sie, und jeder Druck
  ist ein Schritt für Rückgängig.
  - **Text**: Bold, Italic, Strikethrough und Inline code wirken auf die Auswahl
    oder das Wort am Cursor und schalten um (Strg+B, Strg+I, Strg+Umschalt+X und
    Strg+E; Cmd auf dem Mac). Eine Schaltfläche erscheint gedrückt, wenn der Cursor innerhalb
    ihrer Auszeichnung steht. Markdown kennt kein Unterstreichen: Underline (Strg+U) umschließt den Text mit
    `<u>…</u>`, einem Tag, das die Standard-HTML-Allowlist von Marp beibehält, sodass es auch bei ausgeschaltetem
    rohem HTML funktioniert.
  - **Paragraph**: Heading 1 bis 3 (Strg+Alt+1..3), Aufzählung (Strg+Umschalt+8),
    Nummerierung (Strg+Umschalt+7) und ein Zitat wirken auf die ausgewählten Zeilen und schalten um.
    Die horizontale Linie ist das Tag `<hr>`: `---`, `***` und `___` würden in Marp jeweils
    eine neue Folie beginnen.
  - **Table**: Ein Raster wählt Spalten und Zeilen. Am Cursor in einer Tabelle
    fügen Sie eine Zeile oder Spalte hinzu oder entfernen sie und richten die Spalte aus (`:--`, `:-:`,
    `--:`); die Tabelle wird nach jeder Änderung ausgerichtet neu geschrieben.
  - **Slide text** links, zentriert oder rechts schreibt einen `<style scoped>`-Block in
    die aktuelle Folie. Marp kennt keine Ausrichtungssyntax, und ein `style`-Attribut
    passiert die Standard-HTML-Allowlist nicht, der Scoped-Block aber schon, sodass er bei ausgeschaltetem
    rohem HTML funktioniert und in der Vorschau, auf der HTML-Seite und im PDF erscheint.
  - **Link** (Strg+K) schreibt `[text](url)` mit markierter URL.
- **Hervorhebung** der Schlüssel des Front-Matters in beliebiger Tiefe, in YAML (auch
  in `{a: 1}`-Maps) und in HCL.
- **Theme**-Auswahl: default, gaia, uncover, die Themes des Projekts
  (`themes/*.css`) und die Themes, die der Foliensatz definiert (`/* @theme name */` in einem
  `<style>`). Sie schreibt `theme:` in das Front-Matter, YAML oder HCL, sodass ein geteilter Link es beibehält,
  und folgt dem Front-Matter, wenn Sie es bearbeiten.
- Die Vorschau **behält ihre Position** beim Tippen: Nur die Folien, die sich
  geändert haben, werden ersetzt, und ein Themenwechsel öffnet die Vorschau
  wieder an derselben Folie.
- **Link zum Teilen**: Der Foliensatz reist im Link selbst (in dessen `#`-Teil,
  komprimiert), sodass das Öffnen des Links den Foliensatz öffnet. Links fassen bis zu
  64 KiB komprimiert, und nur den Foliensatz: Enthält das Projekt mehr
  (Themes, Kinds, Medien), weist Share darauf hin und bietet die `.zip`-Datei an.
- **Save as**:
  - die HTML-Seite, dieselbe wie von `slides deck.md` geschrieben;
  - ein **PDF** (Vektor, Text auswählbar, mit Gliederung);
  - ein editierbarer **PowerPoint**- oder **OpenDocument**-Foliensatz.

  Sie entstehen in Ihrem Browser, mit demselben Code wie `slides --pdf`,
  `--pptx` und `--odp`: Das PDF des Benchmark-Foliensatzes mit 100 Folien ist
  byteidentisch mit dem des Befehls. Der erste Export lädt die Export-Engine herunter
  (etwa 5 MiB, einmalig; Ihr Browser behält sie bis zum nächsten Update),
  mit einer Fortschrittszeile, und kompiliert sie, während sie eintrifft. Danach lädt er nur die
  Schriften, die Ihre Folien zeichnen. Bei 3 MB/s ist das erste PDF des Willkommens-Foliensatzes
  in unter 3 Sekunden fertig. Der Foliensatz mit 100 Folien braucht dann etwa 3 Sekunden. Ein
  `https:`-Bild wird nur eingebunden, wenn sein Server es erlaubt (CORS); `data:`-Bilder
  immer.
- **Ihr Entwurf bleibt erhalten**:
  - Der Foliensatz wird beim Tippen in diesem Browser gespeichert (IndexedDB, nur für diese
    Website) und kommt beim Neuladen zurück. Nichts wird irgendwohin gesendet. Die
    Textdateien des Projekts (Themes, Kinds, Daten, `slides.yaml`) werden auf dieselbe Weise
    gespeichert, auch wenn Sie den Tab direkt nach dem Tippen schließen; Medien werden
    gespeichert, sobald sie abgelegt werden.
  - Ein geteilter Link öffnet seinen Foliensatz und behält Ihren Entwurf: **Draft ▸ Restore my
    draft** holt ihn zurück.
  - Ein Beispiel, **Draft ▸ Reset to the example** (fragt vorher nach) und
    ein anderer Tab löschen ihn ebenfalls nie: **Draft ▸ Restore replaced** holt ihn
    zurück.
  - **Draft ▸ Forget my drafts** löscht alles Gespeicherte.

Jeder Export rendert das Projekt, das die Vorschau zeigt: seine `slides.yaml`,
Themes, Kinds, Datendateien und Medien, die Medien im Dokument.

Bilder als Dateien (`--images` von slides) entstehen nicht im Browser: Verwenden Sie den
[Befehl]({{< relref "cli" >}}) oder die
[VS-Code-Erweiterung]({{< relref "vscode" >}}).

## Rückgängig, Wiederholen und Versionen {#undo-redo-and-versions}

**Undo** und **Redo** (Rückgängig und Wiederholen) stehen im Reiter Home und
auf Strg+Z, Strg+Umschalt+Z oder Strg+Y (Cmd auf dem Mac). Sie nehmen die
letzte Änderung der angezeigten Datei zurück (jede Datei hat ihren eigenen
Verlauf) und danach die letzte Operation am ganzen Projekt: ein gewähltes
Beispiel, Reset, ein geöffnetes .zip oder ein geöffneter Ordner, ein
wiederhergestellter Entwurf, eine gelöschte oder umbenannte Datei oder ein
solcher Ordner, eine wiederhergestellte Version. Redo bringt alles so zurück,
wie es war, als Undo gedrückt wurde.

**Versionen** halten das ganze Projekt in diesem Browser fest, ohne dass etwas
irgendwohin gesendet wird. Eine Version entsteht kurz nachdem Sie aufhören zu
tippen (höchstens alle 2 Minuten), vor jeder Operation am ganzen Projekt und bei
jedem Export und Teilen-Link. **Save version…** hält eine unter einem Namen
fest. **Home ▸ Versions** öffnet den Bereich Versions:

- **Preview** zeigt eine Version in der Vorschau.
- **Compare with now** zeigt die Änderungen jeder Datei Zeile für Zeile, inline
  oder nebeneinander; F7 und Umschalt+F7 springen zwischen den Änderungen.
- **Restore** stellt eine Version wieder her; Undo nimmt die Wiederherstellung
  zurück.
- **Delete** entfernt eine Version.

Was behalten wird: die 20 neuesten automatischen Versionen, dann eine pro Stunde
für einen Tag und eine pro Tag für einen Monat. Benannte Versionen bleiben, bis
Sie sie löschen (höchstens 50). Alle Versionen zusammen belegen höchstens
96 MiB, und unveränderte Dateien werden nur einmal gespeichert. Ist der
Speicher des Browsers voll, gehen die ältesten automatischen Versionen zuerst.

## Sicherheit {#security}

Der Foliensatz wird in einem iframe gezeigt, das ohne den Ursprung der Seite in einer Sandbox läuft: Er kann
weder den Playground noch dessen Speicher erreichen, kein Formular absenden und
nirgendwohin navigieren. Seine Seite trägt eine Content Security Policy, die nur ihre eigenen
Skripte ausführt. Rohes HTML im Foliensatz folgt der Allowlist von Marp, wie es `slides`
standardmäßig tut. Ein Rendern, das länger als 5 Sekunden läuft (eine Vorlage mit Endlosschleife),
wird abgebrochen.

Das Referentenfenster behält diese Sandbox. Es wird vom Playground
selbst geöffnet und zeigt den Foliensatz in einem ebenso abgeschotteten iframe. Die beiden
iframes können keinen Speicher teilen, deshalb reicht der Playground die Foliennummer
in Nachrichten zwischen ihnen weiter. Kein Frame, der den Foliensatz ausführt, erhält den
Ursprung des Playgrounds.

Die Medien des Projekts erreichen die Sandbox-Frames als Bytes in einer Nachricht, und
jeder Frame erzeugt daraus seine eigenen `blob:`-URLs, in seinem eigenen Ursprung. Eine `blob:`-URL, die vom Playground
erzeugt wurde, lädt dort nicht, und die Richtlinie der Frames
erhält lediglich `blob:` für Medien und Schriften hinzu.

Der Playground kann von Ihrem eigenen Rechner aus bereitgestellt werden, hinter
einer Anmeldung: siehe [Self-Hosting]({{< relref "self-hosting" >}}).

## Größe und Geschwindigkeit {#size-and-speed}

Das Modul umfasst 8,8 MiB, wie GitHub Pages es ausliefert (gzip; Pages liefert
kein Brotli aus). Es wird einmal mit Fortschrittsbalken heruntergeladen und dann vom
Browser bis zum nächsten Deploy zwischengespeichert. Auf einem Apple M4 Max mit Chrome 155, gemessen
an der Live-Seite, ist der Editor 2,1 s nach dem Öffnen der Seite bereit, davon 1,5 s
für den Download, und die ersten Folien erscheinen nach 2,4 s. Eine Änderung am
Benchmark-Foliensatz mit 100 Folien erreicht die Vorschau in etwa 30 ms: go-slides
rendert nur die Folien neu, die sich geändert haben, und nur diese werden in der
Vorschau ersetzt.
