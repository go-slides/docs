---
title: "Kommandozeile"
weight: 50
description: "Jeder Befehl und jedes Flag von slides sowie die Konfigurationsdatei slides.yaml."
tags: [cli]
---

{{< callout type="warning" >}}
Diese Seite beschreibt v0.14.0.
{{< /callout >}}

```text
slides [deck.md] [flags]
slides [command] [flags] deck.md
```

## Befehle {#commands}

| Befehl | Was er tut |
| --- | --- |
| `build` | schreibt die Präsentation als HTML oder als PDF oder Bilder (siehe unten). **Der Standard**: `slides deck.md` ist `slides build deck.md` |
| `expand` | gibt den Foliensatz als reines Marp-Markdown aus: Vorlagen expandiert, Kinds (Folientypen) durch die Direktiven ersetzt, für die sie stehen |
| `notes` | gibt die Sprechernotizen aus |
| `kinds` | listet die Kinds auf, die der Foliensatz verwenden kann, und was sie setzen; `--json` (in v0.6.0) schreibt sie als `[{"name", "extends", "directives": [{"key", "value"}]}]`, `extends` aufgelöst |
| `themes` | listet die Themes auf |
| `version` | gibt die Version aus (ebenso `slides --version`) |
| `completion` | gibt ein Shell-Vervollständigungsskript aus: `bash`, `zsh`, `fish` oder `powershell` |

`slides <command> --help` gibt die Flags eines Befehls aus. Die Flags dürfen vor
oder nach dem Foliensatz stehen: `slides deck.md -o out.html` und `slides -o out.html
deck.md` sind dasselbe. Lange Flags haben zwei Striche (`--html`); `-o` ist das einzige
kurze. `build` schreibt `deck.html` neben `deck.md`, sofern `-o` nichts anderes
angibt.

## Flags {#flags}

| Flag | Befehle | Was es tut |
| --- | --- | --- |
| `-o`, `--output file` | `build`, `expand`, `notes` | Ausgabedatei; `-` für die Standardausgabe |
| `-I`, `--input-dir dir` | `build` | wandelt die Markdown-Dateien unter `dir` rekursiv um (in v0.6.0); siehe unten |
| `-P`, `--parallel n` | `build` | gleichzeitig gebaute Foliensätze (in v0.6.0): standardmäßig 5, höchstens die Anzahl der Kerne; 1 baut sie nacheinander |
| `-w`, `--watch` | `build` | baut erneut, wenn sich ein Foliensatz, ein Theme, eine Kinds- oder Datendatei oder die Konfiguration ändert, und lädt die geöffneten Seiten neu (in v0.6.0); siehe unten |
| `-s`, `--server dir` | `build` | stellt ein Verzeichnis bereit und wandelt seine Foliensätze auf Anfrage um (in v0.6.0); siehe unten |
| `-p`, `--preview` | `build` | öffnet das Ergebnis im Standardbrowser (in v0.6.0): die HTML-Seiten mit `--watch` (impliziert), den Server mit `--server` |
| `--host addr` | `build` | die Adresse, auf der `--server` lauscht, standardmäßig `127.0.0.1` (in v0.6.0) |
| `--allowed-host name` | `build` | der öffentliche Hostname eines Reverse-Proxys, auf den `--server` antwortet, wiederholbar (ab v0.14.0); siehe [Self-Hosting]({{< relref "self-hosting" >}}) |
| `--template bespoke\|bare` | `build` | die Vorlage der HTML-Seite (in v0.6.0): `bespoke`, der Standard, mit Präsentationsskript, Bedienelementen auf dem Bildschirm und Übergängen, oder `bare`, die Folien untereinander (`--bare`) |
| `--bespoke.osc`, `--bespoke.progress`, `--bespoke.transition` | `build` | die Optionen der bespoke-Vorlage, mit den Standardwerten von marp-cli (in v0.6.0): Bedienelemente an, **Fortschrittsbalken aus** (er wurde zuvor immer gezeichnet), Übergänge an; `--bespoke.osc=false` schaltet eines ab |
| `-c`, `--config-file file`, `--no-config-file` | `build` | die Konfigurationsdatei von marp-cli, JSON oder YAML (in v0.6.0); siehe unten. `--config` ist die eigene Datei von `slides` |
| `--browser`, `--browser-path`, `--browser-protocol`, `--browser-timeout` | `build` | werden akzeptiert und mit einer Warnung ignoriert (in v0.6.0): Es gibt keinen Browser |
| `--pptx-editable`, `--no-parallel` | `build` | die Schreibweisen von marp-cli für `--pptx` und `-P 1` (in v0.6.0) |
| `--config file` | alle | Konfigurationsdatei; standardmäßig `slides.yaml` neben dem Foliensatz, falls vorhanden |
| `--no-config` | alle | keine Konfigurationsdatei lesen |
| `--theme-set path` | alle | eine Theme-CSS-Datei oder ein Verzeichnis davon (wiederholbar) |
| `--kinds file` | alle | eine YAML-Datei mit Kinds (wiederholbar) |
| `--data file` | alle | eine YAML- oder JSON-Datei mit Vorlagendaten (wiederholbar) |
| `--set key=value` | alle | ein Vorlagenwert (wiederholbar) |
| `--no-template` | alle | Go-Vorlagen nicht expandieren |
| `--strict-yaml` | alle | Direktiven als striktes YAML parsen, wie es reines Marpit tut (Marp Core parst sie nachsichtig) |
| `--no-emoji` | alle | `:shortcodes:` und Unicode-Emoji als Text belassen, statt sie als Bilder zu zeichnen (seit v0.4.0: [Emoji]({{< relref "emoji" >}})) |
| `--html` | `build` | alles rohe HTML erlauben; standardmäßig die Allowlist von Marp Core |
| `--lang code` | `build` | das `lang` der Seite, wenn der Foliensatz keines setzt |
| `--bare` | `build` | die Folien untereinander stellen, ohne Präsentationsskript |
| `--pdf` | `build` | ein PDF schreiben, eine Seite pro Folie, mit Gliederung – Text und Formen als Vektoren (auch durch `-o file.pdf` gewählt) |
| `--pdf-raster` | `build` | das PDF schreiben, wobei jede Seite ein Bild der Folie ist, wie es v0.3.0 tat |
| `--notes` | `build` | die Sprechernotizen als Textdatei schreiben, Folien durch `---` getrennt wie bei marp-cli (in v0.6.0; auch durch `-o file.txt` gewählt). Der Befehl `notes` gibt sie in einem anderen Layout aus |
| `--pdf-notes` | `build` | die Sprechernotizen jeder Folie als Notizanmerkung auf ihre PDF-Seite setzen, wie es marp-cli tut (in v0.6.0). Verlangt ein PDF, wenn nichts anderes sagt, was geschrieben werden soll |
| `--pdf-outlines`, `--pdf-outlines.pages`, `--pdf-outlines.headings` | `build` | die Lesezeichen des PDFs (in v0.6.0): „Page N“ pro Folie und darunter verschachtelt die Überschriften. Hier standardmäßig an, anders als bei marp-cli; `=false` schaltet eines ab. Siehe unten |
| `--images png\|jpeg` | `build` | ein Bild pro Folie schreiben: `deck.001.png`, … |
| `--image png\|jpeg` | `build` | die erste Folie als Bild schreiben (auch durch `-o file.png` oder `-o file.jpg` gewählt) |
| `--pptx` | `build` | eine PowerPoint-Datei mit ihren Sprechernotizen schreiben (auch durch `-o file.pptx` gewählt): seit v0.5.0 editierbar, in v0.4.0 ein Bild pro Folie |
| `--pptx-raster` | `build` | eine PowerPoint-Datei mit einem ganzseitigen Bild pro Folie schreiben, wie es marp-cli tut (seit v0.5.0); `--image-scale` ist standardmäßig 2 |
| `--odp` | `build` | eine OpenDocument-Präsentation mit ihren Sprechernotizen schreiben (auch durch `-o file.odp` gewählt): seit v0.5.0 editierbar |
| `--odp-raster` | `build` | eine OpenDocument-Präsentation mit einem Bild pro Folie schreiben (seit v0.5.0) |
| `--jpeg-quality n` | `build` | JPEG-Qualität, 1 bis 100 (Standard 85) |
| `--image-scale n` | `build` | Geräte-Pixel pro CSS-Pixel für Bilder und PDF (Standard 1); 2 ergibt 2560×1440 bei 16:9, wobei Text und Vektoren in dieser Dichte gezeichnet und nicht vergrößert werden |
| `--allow-local-files` | `build` | Bildern und PDF das Lesen lokaler Dateien erlauben, nur unterhalb des Verzeichnisses des Foliensatzes; standardmäßig abgelehnt, wie bei marp-cli |
| `--theme name\|file.css` | `build` | das Theme des Foliensatzes überschreiben (in v0.6.0): ein Theme-Name oder eine CSS-Datei, die zum Theme wird, welches `@theme` sie auch deklariert; ein Verzeichnis wird abgelehnt, verwenden Sie `--theme-set` |
| `--title`, `--description`, `--author`, `--keywords`, `--url`, `--og-image` | `build` | die Metadaten der Seite (in v0.6.0), vor den Direktiven `title`, `description`, `author`, `keywords`, `url` und `image` des Foliensatzes; siehe unten |

## Mehrere Foliensätze und ein Verzeichnis {#several-decks-and-a-directory}

Ab v0.6.0 nimmt `build` mehrere Foliensätze, ein Verzeichnis (die darunter
rekursiv gefundenen Foliensätze) oder ein Muster entgegen und baut bis zu `--parallel` davon gleichzeitig.
`-` ist die Standardeingabe; anders als bei marp-cli wird eine Pipe nur gelesen, wenn `-` es verlangt.

Ein Muster wird so gelesen, wie es der Glob von marp-cli liest (`**` und `{a,b}` in v0.7.0;
setzen Sie es in Anführungszeichen, damit die Shell es `slides` überlässt): `*`, `?`, `[a-z]` und
`{md,txt}` innerhalb eines Namens, `**` für beliebig viele Verzeichnisse, sodass
`'docs/**/*.md'` die Foliensätze des Baums erfasst. Ein Platzhalter erreicht keinen Namen,
der mit einem Punkt beginnt, in `node_modules` wird nicht gesucht, die Groß- und Kleinschreibung zählt
(`*.md` findet `UPPER.MD` nicht), und was immer ein Muster trifft, wird umgewandelt,
gleich welche Endung es hat. Ein Muster ohne Treffer ist ein Fehler.

- **`-I dir`** findet die Markdown-Dateien (`.md`, `.mdown`, `.markdown`,
  `.markdn`, in Kleinbuchstaben) unter `dir`, wie marp-cli es tut: `node_modules` und
  Punkt-Verzeichnisse werden nicht durchsucht. Mit `-o out` wird jeder Foliensatz mit seinem
  Pfad unter `dir` gespiegelt unter `out` geschrieben: `dir/sub/b.md` ergibt
  `out/sub/b.html` (`b.pdf`, `b.001.png`, `b.txt`, … je nachdem, was die Flags sagen).
  Ohne `-o` landen die Ausgaben neben ihren Foliensätzen. Foliensätze zusätzlich zu
  `-I` zu nennen ist ein Fehler, ebenso `-o file` bei mehreren Foliensätzen (`-o -` ist erlaubt).
- **`-P n`** baut `n` Foliensätze gleichzeitig. Das Rendern eines einzelnen Foliensatzes nutzt schon die
  Kerne, deshalb ist `n` durch deren Anzahl begrenzt; mit `-o -` werden Foliensätze nacheinander
  gebaut. Die Meldungen eines Foliensatzes bleiben beisammen, und ein Foliensatz, der scheitert, hält die
  anderen nicht auf: Der Exit-Status ist 1, und die Meldung listet die Fehlschläge auf.

## Übergänge und die Präsentationsseite {#transitions-and-the-presenting-page}

Ab v0.6.0 funktioniert die Direktive `transition` von marp-cli: `<!-- transition:
fade 1s -->`, ein Name und eine optionale Dauer, für die Folie und die folgenden,
oder `<!-- _transition: cover -->` für eine einzelne. Die eingebauten Namen sind die von marp-cli:
`clockwise`, `counterclockwise`, `cover`, `coverflow`, `cube`, `cylinder`,
`diamond`, `drop`, `explode`, `fade`, `fade-out`, `fall`, `flip`, `glow`,
`implode`, `in-out`, `iris-in`, `iris-out`, `melt`, `overlap`, `pivot`, `pull`,
`push`, `reveal`, `rotate`, `slide`, `star`, `swap`, `swipe`, `swoosh`, `wipe`,
`wiper` und `zoom`; `none` schaltet ihn ab. Die Seite spielt ihn mit der
**View Transition API** des Browsers ab, wenn der Betrachter die Folie, die ihn hat, vorwärts verlässt,
und wenn er zu ihr zurückkehrt; Browser ohne die API wechseln die Folien wie zuvor,
und ein Betrachter, der reduzierte Bewegung verlangt, erhält eine einfache Überblendung. Ein Theme kann
eigene Keyframes mitbringen, benannt `marp-transition-NAME` (oder `marp-outgoing-` und
`marp-incoming-` sowie `-backward`), die vor den eingebauten verwendet werden.

Die Seite hat die **Bedienelemente auf dem Bildschirm** – zurück, Seite, weiter, Vollbild –,
die erscheinen, wenn sich der Mauszeiger oder eine Taste bewegt, und nach einem Moment verschwinden
(`--bespoke.osc=false` entfernt sie), und mit `--bespoke.progress` einen Fortschrittsbalken.
**Der Fortschrittsbalken ist jetzt standardmäßig aus**, wie bei marp-cli; zuvor wurde er immer
gezeichnet.

### Referentenansicht {#presenter-view}

Ab v0.7.0 hat die Seite die Referentenansicht von marp-cli. Drücken Sie `p` oder die letzte
Schaltfläche der Bedienelemente, und ein Fenster öffnet sich mit der aktuellen Folie, dem Schritt
danach, Ihren Sprechernotizen, der Seite, der Uhrzeit und einem Timer (ein Klick auf den Timer
setzt ihn zurück; `+` und `-` ändern die Größe der Notizen; die Leiste zwischen den Folien und den
Notizen lässt sich ziehen). Die beiden Fenster folgen einander in beide Richtungen: die
Pfeiltasten, ein Klick auf den nächsten Schritt oder auf Previous/Next in einem von beiden bewegt das
andere, jeweils ein Fragment auf einmal.

Sie funktioniert wie die von marp-cli, gelesen aus dessen Quellen der Version 4.5.1: Das Fenster ist
`?view=presenter&sync=KEY` derselben Datei, der nächste Schritt ist dieselbe Seite als
`?view=next` in einem Frame, und die Fenster teilen ihre Position über `localStorage` (`bespoke-marp-sync-KEY`),
sodass sie weder Server noch Netzwerk braucht und in einem Browser ohne
`localStorage` nicht funktioniert (die Schaltfläche sagt es). Die Notizen sind Text: `&`, `<` und `>` werden maskiert, sodass eine Notiz
nie Markup sein kann. Sie stehen versteckt in der Seite, damit die Referentenansicht sie übernehmen kann, was sie
aus dem heraushält, was ein Screenreader dem Publikum vorliest. `--bare`-Seiten haben nichts davon.
Die Ansicht fügt einer Seite etwa 10 KB hinzu.

### Übersichtsansicht {#overview-view}

Ab v0.8.0 zeigt sich jede Folie ganz in einem Raster über der Seite, wenn Sie `o` drücken (oder `Escape`: marp-cli schaltet mit
beiden um), die Rasterschaltfläche der Bedienelemente oder den Seitenzähler der Referentenansicht.
Die Pfeiltasten wählen eine Folie (aufwärts und abwärts zeilenweise), Pos1 und Ende (`Home` und `End`) springen zu den Enden,
und die anderen Fenster folgen; `Enter`, `Space` oder ein Klick wählt eine Folie und schließt, über der Seite, die
Übersicht. Über der Referentenansicht lässt eine Wahl sie geöffnet (der Referent wählt
die zu zeigende Folie und schaut weiter). Die Schaltfläche „Close overview“, `o` und
`Escape` schließen sie. Sie ist die von marp-cli: die Seite noch einmal in einem Frame unter
`?view=overview&sync=KEY`, die Position über `localStorage` teilend, ohne
Netzwerk. Die Folien im Raster sind fokussierbar und mit „Slide N“ beschriftet, und die
geschlossene Übersicht ist inert, sodass Tastatur- und Screenreader-Nutzer nie darin landen.
Die beiden Ansichten zusammen fügen einer Seite etwa 15 KB hinzu.

### Konfigurationsdateien {#configuration-files}

marp-cli liest seine Optionen aus einer Datei im Arbeitsverzeichnis, und `slides` tut es auch,
sofern die Datei aus Daten besteht: `.marprc` (YAML, JSON ist auch YAML),
`.marprc.json`, `.marprc.yaml`, `.marprc.yml`, `.config/marprc*` oder der
Schlüssel `marp` in `package.json`; `--config-file` (`-c`) benennt jede JSON- oder YAML-
Datei, `marp.config.json` eine solche. Durchsucht wird nur das Arbeitsverzeichnis, nicht
seine übergeordneten, wie bei der Suche von marp-cli. Eine JavaScript- oder TypeScript-Datei
(`marp.config.js` ...) ist ein Programm, das `slides` nicht ausführt: Es sagt das und
macht weiter. Die Optionen sind die aus der Dokumentation von marp-cli, die hier etwas
bedeuten: `title`, `description`, `author`, `keywords`, `url`, `ogImage`, `lang`,
`theme`, `themeSet`, `html`, `allowLocalFiles`, `template`, `bespoke`,
`inputDir`, `output`, `parallel`, `pdf`, `pptx`, `notes`, `pdfNotes`,
`pdfOutlines`, `images`, `image`, `imageScale`, `jpegQuality`, `watch`,
`server`, `preview`; Pfade sind relativ zur Datei. Die Kommandozeile hat Vorrang.
`engine`, `options` und `browser*` brauchen die Engine von Marp und einen Browser und werden
mit einer Warnung ignoriert, ebenso wie die Flags `--browser*`: Für marp-cli geschriebene Skripte
funktionieren weiter. (`slides.yaml` und `slides.hcl`, mit den Kinds und den
Vorlagendaten, sind getrennt und werden wie bisher gelesen.)

## Beobachten, bereitstellen und Vorschau {#watch-serve-and-preview}

Ab v0.6.0, mit den Routen und dem Live-Reload-Protokoll von marp-cli (gelesen aus dessen
Quellen der Version 4.5.1 und verglichen mit einem laufenden `marp -s`).

**`slides --server dir`** (`-s`) stellt das Verzeichnis unter `http://127.0.0.1:8080/` bereit
(die Umgebungsvariable `PORT` ändert den Port; ein bereits belegter Port ist ein
Fehler, der das sagt):

| Anfrage | Antwort |
| --- | --- |
| `/deck.md` | der Foliensatz als HTML, auf Anfrage umgewandelt |
| `/deck.md?pdf`, `?pptx`, `?png`, `?jpg` (oder `?jpeg`), `?notes` (oder `?txt`) | der Foliensatz, umgewandelt wie in der Abfrage verlangt; das PPTX ist ein Anhang |
| ein Verzeichnis | sein `index.md` oder `PITCHME.md` als Seite, sonst ein Index mit einem PDF- und einem PPTX-Link für jeden Foliensatz |
| jede andere Datei | die Datei |
| ein Foliensatz, der sich nicht umwandeln lässt | 503 mit dem Fehler |

Eine Seite öffnet einen WebSocket zu `/.__marp-cli-watch-notifier__/<id>`, wobei die ID
der SHA-256 des Pfades des Foliensatzes ist; der Server sendet `ready`, dann `reload`, wenn sich
der Foliensatz oder irgendetwas unter dem Verzeichnis ändert. Ein Verzeichnis, das ohne
den abschließenden Schrägstrich angefordert wird, wird dorthin umgeleitet (marp-cli liefert seinen Index dort aus,
was die relativen Links der Seite zerbricht).

**`--watch`** (`-w`) baut die Foliensätze und baut dann den, der sich ändert, erneut
und, wenn sich eine `--theme-set`-Datei, eine Kinds- oder Datendatei oder die Konfiguration
ändert, alle. Die Seiten, die es schreibt, öffnen einen Socket auf dem Loopback (Port
37717 oder der nächste freie) und laden neu, sobald ihr Foliensatz erneut gebaut wird. Ein Foliensatz,
der nicht gebaut werden kann, wird gemeldet und weiter beobachtet. **`--preview`** (`-p`) öffnet
den Standardbrowser auf dem Server oder auf den soeben gebauten HTML-Seiten; es gibt kein
eigenes Fenster, da `slides` keinen Browser steuern kann.

Der Server ist ein Zugang zur Maschine und deshalb strenger als der von marp-cli, der
auf jeder Schnittstelle lauscht und nichts prüft:

- er lauscht auf **127.0.0.1**; `--host` ändert das mit einer Warnung, und eine
  Wildcard-Adresse kann den `Host`-Header nicht mehr prüfen;
- der **`Host`**-Header muss ein Loopback-Name sein (`localhost`, `127.0.0.1`,
  `[::1]`, `*.localhost`) oder der Host, auf dem er lauscht, was DNS-Rebinding unterbindet;
- der **`Origin`** eines WebSockets muss der des Servers selbst sein; die Sockets des Watch-Modus
  akzeptieren auch den `null`-Origin einer aus einer Datei geöffneten Seite; eine Anfrage von
  einer anderen Website, die keine Navigation ist (ein Skript, ein Bild, ein `fetch`), wird
  abgelehnt;
- ein Pfad wird bereinigt, seine symbolischen Links werden aufgelöst, und das Ergebnis muss
  im Verzeichnis liegen: `..` und ein Link hinaus werden abgelehnt, ebenso
  `?pdf` eines Foliensatzes außerhalb; Namen, die mit einem Punkt beginnen (`.git`, `.env`), werden nicht
  ausgeliefert;
- nur `GET` und `HEAD`; wenige Umwandlungen gleichzeitig; höchstens 4096 Verzeichnisse
  beobachtet, und keines in `node_modules` oder unter einem Punkt.

Hinter einem Reverse-Proxy fügt `--allowed-host name` (ab v0.14.0,
wiederholbar) den öffentlichen Namen des Proxys zu den akzeptierten `Host` und
`Origin` hinzu; die Prüfungen bleiben aktiv, und der Server bleibt auf dem
Loopback. [Self-Hosting]({{< relref "self-hosting" >}}) stellt ihn hinter Caddy,
mit HTTPS und einer Anmeldung.

## Seitenmetadaten {#page-metadata}

Die Kommandozeile hat Vorrang vor dem Foliensatz, wie bei marp-cli, wo die Flags
als letzter Direktivenkommentar an den Foliensatz angehängt werden. Die Regeln sind aus marp-cli
4.5.1 gelesen und durch Tests gegen dessen Ausgabe festgelegt (in v0.6.0):

- Ein Foliensatz ohne `title` übernimmt den Text seiner ersten Überschrift der niedrigsten
  Ebene, wie geschrieben (`# **Bold** title` ergibt `**Bold** title`). Ohne
  Überschrift gibt es kein `<title>`.
- `keywords` werden an Kommas getrennt, getrimmt und dedupliziert; eine Liste im Front-Matter
  wird so übernommen, wie sie ist. Wie bei marp-cli trägt die Seite das Tag `keywords` ab
  dem zweiten Schlüsselwort, nicht bei nur einem.
- Eine ungültige `url` (nicht absolut) wird mit einer Warnung ignoriert, und der vorherige
  Wert bleibt; eine gültige schreibt außerdem `<link rel="canonical">` und `og:url`.
- `author` und `description` kommen mit ihren Open-Graph-Gegenstücken; `og:image:alt`
  und die große Twitter-Karte brauchen sowohl einen Titel als auch ein Bild.

## Bilder, PDF, PPTX und ODP {#images-pdf-pptx-and-odp}

`-o` mit der Endung `.pdf`, `.pptx`, `.odp`, `.png` oder `.jpg` wählt das Format, wie bei marp-cli.
Ausgabeverzeichnisse werden angelegt. Ein PDF hat eine Seite pro Folie, in der
Größe der Folie, und eine Gliederung der Folien und ihrer Überschriften.

### PDF-Notizen und Gliederung {#pdf-notes-and-outline}

Ab v0.6.0 übernimmt das PDF die beiden Optionen von marp-cli, und sie werden mit
den PDFs verglichen, die marp-cli 4.5.1 für denselben Foliensatz schreibt:

- `--pdf-notes` setzt die Sprechernotizen jeder Folie als Notizanmerkung auf ihre Seite:
  ein gelbes Notiz-Symbol unten links, dessen Inhalt die
  Notizen sind (mehrere Kommentare durch eine Leerzeile verbunden) und dessen Titel der `author`
  des Foliensatzes ist. `--pdf-notes` allein verlangt ein PDF, wie bei marp-cli.
- Die Gliederung enthält ein fettes Lesezeichen „Page N“ pro Folie und darunter die
  Überschriften der Folie nach Ebene verschachtelt (`--pdf-outlines.pages` und
  `--pdf-outlines.headings` wählen die beiden). **Die Gliederung ist hier standardmäßig an und
  bei marp-cli aus.** `slides` hat schon immer eine geschrieben; sie kostet ein paar
  hundert Byte und macht einen langen Foliensatz navigierbar, also bleibt der Standard.
  `--pdf-outlines=false` oder `--no-pdf-outlines` von marp-cli schaltet sie ab, und
  `=false` bei `.pages` oder `.headings` lässt diesen Teil weg. Ein Überschriften-Lesezeichen springt zu
  seiner Folie; marp-cli scrollt außerdem zur Überschrift.

Das Rendern übernimmt [go-webengine](https://github.com/go-webengine/engine),
und das PDF schreibt [go-pdfkit](https://github.com/go-pdfkit/pdfkit);
die Schriften, die die eingebauten Themes verlangen (Lato, Roboto Mono), stammen aus
[go-opentype/fonts](https://github.com/go-opentype/fonts) und sind eingebettet, sodass
kein Rendern das Netzwerk berührt. Es gibt keinen Chrome.

`--image-scale 2` rendert mit zwei Geräte-Pixeln pro CSS-Pixel. Lokale Bilder werden
abgelehnt, solange `--allow-local-files` nicht angegeben ist, und selbst dann werden nur Dateien unter
dem Verzeichnis des Foliensatzes gelesen: Ein Foliensatz von anderswo kann `~/.ssh` nicht erreichen.
Mehrere und geteilte Hintergründe (`![bg left:40%]`) werden gerendert, wie Marp ihre Ebenen rendert.

`--pptx` schreibt den Foliensatz mit `WritePresentation` von [go-ooxml](https://github.com/go-ooxml),
mit den Sprechernotizen und den Metadaten des Foliensatzes. Seit
v0.5.0 sind die Folien editierbar (unten); `--pptx-raster` schreibt ein ganzseitiges
Bild pro Folie, wie marp-cli es tut, mit einem Maßstab, der standardmäßig 2 ist und den
`--image-scale` überschreibt. LibreOffice öffnet das Ergebnis. Getestet wurde nur mit LibreOffice;
PowerPoint und Keynote wurden nicht getestet.

`--odp` schreibt eine OpenDocument-Präsentation über `WritePresentation` von go-odf, mit
ihren Sprechernotizen; `--odp-raster` schreibt ein Bild pro Folie.

### Editierbares PPTX und ODP {#editable-pptx-and-odp}

Seit v0.5.0 sind die Folien eines PPTX oder ODP **editierbar** – das PPTX von marp-cli
ist ein Bild pro Folie, und sein experimentelles `--pptx-editable` wandelt ein PDF
über LibreOffice um. Hier wird jede Folie aus ihrem Layout gezeichnet, demselben,
das die PNG- und PDF-Exporte zeichnen:

- jede Textzeile ist ein Textfeld an ihrer berechneten Position, in den Schriften,
  mit denen die Folie gesetzt wurde – **in die Datei eingebettet**, sodass sie auf einem Rechner ohne Lato oder Inter
  gleich aussieht (geprüft mit LibreOffice 25.8) –
  mit Fett, Kursiv, Farben, Buchstabenabstand und anklickbaren Links;
- Hintergründe, Rahmen und abgerundete Kästen sind Formen; Verläufe sind Verlaufs-
  füllungen im PPTX und Bilder im ODP (ein ODF-Verlauf hat zwei
  Farben);
- Bilder werden nur einmal gespeichert, egal wie viele Folien sie zeigen;
- was eine Folie nicht ausdrücken kann – ein Element unter `opacity`, ein `filter`, ein
  weichgezeichneter Schatten, eine Formel – ist ein Bild dieses Elements, mit seinem darüberliegenden Text.

Von LibreOffice zurückgerendert, liegt eine Folie 0,04–2 % von den Pixeln des
PNG-Exports entfernt. `--pptx-raster` und `--odp-raster` behalten ein Bild pro Folie.

### Vektor-PDF {#vector-pdf}

Seit v0.4.0 wird das PDF als Vektoren gezeichnet: go-webengine malt die berechnete
Folie in das PDF statt in Pixel. Der Text ist Text, in den Schriften, mit denen die
Folie gesetzt wurde, als Teilmenge eingebettet mit einer `/ToUnicode`-Zuordnung – er lässt sich
auswählen, durchsuchen und kopieren. Hintergründe, Rahmen und Verläufe sind Pfade
und Schattierungen, ein Bild, das auf vielen Folien erscheint, wird einmal eingebettet, und Links sind
anklickbar. Der Benchmark-Foliensatz mit 100 Folien ist 758 KiB groß, gegenüber 837 KiB bei
marp-cli und 10 MiB beim Bild-PDF von v0.3.0, und er ist in 0,3 s geschrieben.

Ein Element ohne Vektorform – unter `opacity` < 1, mit einem `filter`
oder einer Drehung, ein weichgezeichneter `box-shadow`, ein Verlauf mit ausblendender Deckkraft –
wird als Bitmap nur dieses Elements gezeichnet, mit unsichtbar darübergelegtem Text, damit er
auswählbar bleibt. Formeln sind weiterhin Bitmaps. `--pdf-raster` schreibt ein Bild
pro Seite, wie es v0.3.0 tat.

### Barrierefreies, getaggtes PDF {#accessible-tagged-pdf}

Seit v0.12.0 ist das Vektor-PDF getaggt. Es trägt seine logische Struktur, die ein
Screenreader der Reihe nach liest, wie er eine Webseite liest:

| Im Foliensatz | Im PDF |
| --- | --- |
| eine Folie | ein Abschnitt (`Sect`) |
| `#` … `######` | `H1` … `H6` |
| ein Absatz, eine Liste, eine Tabelle | `P`; `L` mit Marke und Inhalt jedes Eintrags; `Table`, Kopfzellen mit ihrem Bezugsbereich |
| ein Link | ein `Link`, einschließlich seiner anklickbaren Fläche, sein Text als Beschreibung |
| ein Bild | eine `Figure` mit ihrem Alternativtext; `![](…)` ohne Alternativtext ist Dekoration |
| `$…$`, `$$…$$` | eine `Formula`, ihr TeX-Quelltext als Beschreibung |
| Code | `Code`, ohne Zeilennummern und Callout-Scheiben |
| `<span lang="fr">` | ein Span in dieser Sprache |
| Hintergründe, Rahmen, Seitenzahlen | Artefakte, die ein Reader überspringt |
| Sprechernotizen (`--pdf-notes`) | Annotationen in der Struktur |

Die Datei deklariert PDF/UA-1, den Standard für barrierefreies PDF. Sie trägt den
`title` des Foliensatzes, den der Viewer anzeigt, und seine `lang` (`--lang`, wenn der
Foliensatz keine setzt, sonst `en`).

Ein mit veraPDF geprüfter Foliensatz besteht PDF/UA-1, es sei denn, seine eigenen
Überschriften verletzen die Reihenfolge des Standards. PDF/UA verlangt, dass die erste
Überschrift eine `H1` ist und keine Ebene übersprungen wird (`#`, dann `###`). Das kann
nur der Autor des Foliensatzes beheben. Das PDF von Chrome für denselben Foliensatz, das
marp-cli schreibt, ist ebenfalls getaggt, verletzt aber sieben Regeln von PDF/UA, hat
keinen Abschnitt pro Folie und keine Formeln.

Das Taggen des Benchmark-Foliensatzes mit 100 Folien kostet 4 % an Größe (788 gegenüber
757 KiB) und keine messbare Zeit. `--pdf-tags=false` schreibt das PDF ohne Tags.

## `slides.yaml` {#slidesyaml}

Eine Konfigurationsdatei neben dem Foliensatz gilt für ihn, sodass sich ein Ordner voller Foliensätze
eine teilt:

```yaml
kinds: kinds.yaml        # a mapping of kinds, or file(s) holding one
themes:                  # theme CSS files or directories
  - themes/
data:                    # template data
  company: Example Inc.
html: false              # allow all raw HTML
lang: en
```

Kinds: Das eigene `kinds:` des Foliensatzes überschreibt das der Datei und das der Kommandozeile.
Daten: Das `data` der Datei und der Kommandozeile überschreibt das Front-Matter des Foliensatzes,
sodass ein Foliensatz mit verschiedenen Daten gebaut werden kann.

## Geplant {#planned}

Noch nicht gebaut: Formeln als Vektoren im PDF. Sobald sie kommen, sind sie hier ein
Befehl oder Flag.
