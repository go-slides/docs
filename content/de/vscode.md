---
title: "VS Code"
weight: 55
description: "Die go-slides-Erweiterung für VS Code: Vorschau, Präsentieren, Probleme, Vervollständigung und Export, so wie Marp for VS Code es für Marp bietet."
tags: [vscode, editor]
---

{{< callout type="info" >}}
Die Erweiterung ist [go-slides/vscode](https://github.com/go-slides/vscode),
mit eigener Versionierung (v0.5.0). Sie rendert mit go-slides v0.14.0 (dem Paket
`editor`, den Quellzeilen für die Scroll-Synchronisierung sowie der Übersichts- und
Referentenansicht). Installiert wird sie über die an das GitHub-Release angehängte Datei
`go-slides-0.5.0.vsix`: Die Veröffentlichung im VS Code Marketplace und auf Open VSX ist
vorbereitet, aber noch nicht erfolgt.
{{< /callout >}}

Die Erweiterung ist das Gegenstück von go-slides zu
[Marp for VS Code](https://github.com/marp-team/marp-vscode): Der Foliensatz erscheint in
der Markdown-Vorschau von VS Code als Folien, seine Fehler im Problems-Panel, und
Direktiven, Kinds (die Art einer Folie) und Vorlagenfunktionen werden beim Tippen vervollständigt.

## Installation {#install}

Laden Sie `go-slides-<version>.vsix` von den
[Releases](https://github.com/go-slides/vscode/releases) herunter und führen Sie dann
*Extensions: Install from VSIX...* aus, oder:

```sh
code --install-extension go-slides-0.5.0.vsix
```

Die Vorschau braucht nichts weiter. Für den Export benötigen Sie den Befehl `slides`:

```sh
go install github.com/go-slides/slides/cmd/slides@latest
```

## Vorschau {#preview}

Öffnen Sie einen Foliensatz und wählen Sie *Markdown: Open Preview to the Side*. Ein Foliensatz ist eine Markdown-Datei,
deren Front-Matter `slides: true` enthält:

```markdown
---
slides: true
paginate: true
data:
  who: World
---

# Hello {{ .who }}

---

<!-- _kind: title -->
# Page {{ page }} of {{ total }}
```

Dateien mit `marp: true` werden ebenfalls in der Vorschau gezeigt, es sei denn, Marp for VS Code ist installiert:
Beide Erweiterungen würden dasselbe Dokument rendern. Die Einstellung
`goSlides.preview.marpDocuments` entscheidet: `auto` (der Standard), `always` oder
`never`.

Die Vorschau ist go-slides selbst, nach WebAssembly kompiliert und von der
Erweiterung geladen: kein Binary nötig, und sie funktioniert auf [vscode.dev](https://vscode.dev). Sie liest
aus dem Workspace, was der Foliensatz verwendet – `slides.yaml` oder `slides.hcl` neben
dem Foliensatz, die [Kinds]({{< relref "kinds" >}})-Dateien, die diese benennen, Dateien, die eine
[Vorlage]({{< relref "templates" >}}) einbindet, Theme-CSS – und das nur aus dem
Workspace-Ordner des Foliensatzes.

## Scroll-Synchronisierung {#scroll-sync}

Seit v0.2.0 der Erweiterung (go-slides v0.7.0) scrollen Editor und
Vorschau gemeinsam, in beide Richtungen, wie bei reinem Markdown: Die
Vorschau folgt der Folie und dem Block unter dem Cursor des Editors, und
Scrollen in der Vorschau bewegt den Editor. Die Einstellungen sind die von VS Code selbst,
`markdown.preview.scrollPreviewWithEditor` und
`markdown.preview.scrollEditorWithPreview`.

Zeilen werden durch Vorlagen hindurch verfolgt. Eine Folie, die ein `{{ range }}` erzeugt,
folgt der Zeile ihrer Vorlage, und was ein `{{ include }}` hereinholt,
folgt der Zeile des Includes. Ein durch `headingDivider` aufgeteilter Foliensatz synchronisiert sich an seinen
Überschriften.

## Präsentieren {#presenting}

Seit v0.4.0 der Erweiterung (go-slides v0.9.0) lässt sich ein Foliensatz präsentieren,
ohne VS Code zu verlassen. Mit dem Foliensatz im Editor, über die Befehlspalette
oder das `…`-Menü des Editors:

| Befehl | Öffnet |
|---|---|
| *go-slides: Open Overview* | ein Panel mit allen Folien in einem Raster. Eine gewählte Folie zeigt ihre Quelle im Editor; `Esc` schließt das Panel |
| *go-slides: Start Presenter View* | die Folien und daneben die Referentenansicht: den nächsten Schritt, die Notizen des Vortragenden und einen Timer |

Die Panels liegen neben dem Editor und folgen einander: Wechseln Sie die Folie in einem,
wechseln die anderen mit. Sie folgen auch Ihren Änderungen am Foliensatz. Die Taste `p`
in den Folien öffnet die Referentenansicht.

Um auf einem anderen Bildschirm zu präsentieren, verschieben Sie das Folien-Panel mit *View: Move
Editor into New Window* dorthin.

Jedes Panel führt den Foliensatz in einem Sandbox-Frame aus (ohne `allow-same-origin`) unter einer
Richtlinie mit `default-src 'none'`: Der Frame lädt Bilder, Medien, Stile und Schriften
und sonst nichts und führt nur die Skripte aus, die go-slides in die Seite eingebaut hat,
benannt per Hash. Ein Skript, das der Foliensatz selbst mitbringt, läuft nicht.

## Probleme, Vervollständigung, Hover {#problems-completion-hover}

- **Probleme**: Fehler in Vorlagen, YAML und [HCL]({{< relref "hcl" >}}), in der
  Zeile, in der sie auftreten, im Foliensatz oder in der Kinds-Datei, aus der sie stammen.
- **Vervollständigung**: Direktivenschlüssel im Front-Matter und in `<!-- -->`-Kommentaren
  (mit `_` für eine Spot-Direktive), die Werte von `kind` und `_kind` (mit dem, was
  das jeweilige Kind setzt), `theme`, `size`, `paginate`, `math`, `headingDivider`; innerhalb von
  `{{ }}` die Vorlagenfunktionen und die Datenfelder.
- **Hover**: was eine Direktive bewirkt, was ein Kind setzt, die Argumente einer Vorlagenfunktion. Seit v0.16.0 folgt die Hilfe zu den Direktiven
  der Sprache der Editor-Oberfläche (der Anzeigesprache von VS Code), in 11 Sprachen.

## Export {#export}

*go-slides: Export Slide Deck...* (auch in der Titelleiste des Editors) schreibt HTML,
Vektor-PDF, editierbares PPTX und ODP, PNG oder JPEG mit dem Befehl `slides`, genau wie
`slides -o deck.pdf deck.md` es täte; siehe [Kommandozeile]({{< relref "cli" >}}).
Dafür ist ein vertrauenswürdiger Workspace nötig. Wird `slides` nicht gefunden, bietet die Erweiterung an,
das obige `go install` in einem Terminal auszuführen.

Seit v0.5.0 der Erweiterung (go-slides v0.14.0) ist das geschriebene PDF
**barrierefrei**: Mit `slides` ab v0.12.0 ist das Vektor-PDF für Screenreader
getaggt und deklariert PDF/UA-1, mit `title` und `lang` des Foliensatzes und
einem Abschnitt pro Folie. `slides --pdf-tags=false` schreibt die Datei ohne Tags.
Aktualisieren Sie den Befehl mit dem obigen `go install`.

## Einstellungen {#settings}

| Einstellung | Standard | |
|---|---|---|
| `goSlides.cliPath` | `""` | der Befehl `slides`; leer: `slides` im `PATH`. Eine Maschineneinstellung: Ein Workspace kann sie nicht setzen |
| `goSlides.exportType` | `pdf` | das Format, das der Exportdialog zuerst vorschlägt |
| `goSlides.exportAutoOpen` | `true` | die exportierte Datei öffnen |
| `goSlides.html` | `false` | alles rohe HTML statt der Allowlist von Marp erlauben; wird nur in einem vertrauenswürdigen Workspace beachtet |
| `goSlides.preview.marpDocuments` | `auto` | `marp: true`-Foliensätze in der Vorschau zeigen: `auto`, `always`, `never` |

## Sicherheit {#security}

Foliensatz und Workspace gelten als nicht vertrauenswürdig. Rohes HTML folgt der Allowlist von Marp,
es sei denn, `goSlides.html` ist in einem vertrauenswürdigen Workspace eingeschaltet (eine `slides.yaml` mit
`html: true` schaltet es in der Vorschau nicht ein), und die Vorschau von VS Code führt kein
Skript aus dem Foliensatz aus. Die Vorschau liest nur Dateien aus dem Workspace-Ordner des Foliensatzes,
folgt keinem symbolischen Link, und Vorlagen sehen keine Umgebungsvariable.
Der Export führt ein Programm aus, wartet daher auf die Workspace-Vertrauensstellung, und das Programm
wird nie vom Workspace bestimmt.

## Performance {#performance}

Mit dem 100-Folien-Benchmark-Foliensatz von go-slides, in VS Code 1.142 auf einem Apple M4 Max,
mit Marp for VS Code 3.6.1 in derselben Sitzung gemessen:

| | go-slides | Marp for VS Code |
|---|---:|---:|
| Aktivierung (Laden des WebAssembly-Moduls) | 189 ms | — |
| erste Vorschau | 52 ms | 78 ms |
| Vorschau nach einer Änderung (Median) | **19 ms** | 23–24 ms |
| `.vsix` | 7,5 MB | 36,3 MB |

Die Werte gelten für Erweiterung v0.3.0 (go-slides v0.8.0) mit einem Foliensatz von 100 Folien,
bei Eingabe eines Zeichens in Folie 50, aus abwechselnd ausgeführten Durchläufen bei einem Load Average über eine Minute unter 8. Die erste Vorschau ist schneller als die von Marp, und die Vorschau nach einer
Änderung ist es nun ebenfalls.

### In v0.8.0 und Erweiterung v0.3.0 {#in-v080-and-extension-v030}

go-slides v0.8.0 behält die gerenderten Folien jedes geöffneten Foliensatzes und rendert nach einer Änderung
nur die neu, deren Quelle oder geerbter Zustand sich geändert hat (Direktiven, Seiten-
nummer, Gesamtzahl, Überschrifts-IDs, Linkreferenzen, Quellzeilen); die Ausgabe ist
dieselbe wie bei einem vollständigen Rendern, Byte für Byte. Tippen in Folie 50 von 100 rendert eine
Folie neu: 2,5 ms nativ, 11,5 ms in WebAssembly. Die Vorschau von VS Code überführt dann
nur das DOM der geänderten Folie, und der Smoke-Test der Erweiterung tippt ein
Zeichen in Folie 50 zwischen zeitgemessenen Renderings, für beide Erweiterungen. Die
Referentenansicht der Erweiterung synchronisiert sich über das übergeordnete Fenster (`postMessage`).

### In v0.7.1 und Erweiterung v0.2.2 {#in-v071-and-extension-v022}

go-slides v0.7.1 hält geparste Themes, gepackte Stylesheets und hervorgehobenen Code
zwischen Renderings im Cache, markiert Scroll-Sync-Zeilen, ohne Byte für Byte vorzugehen, und verwendet
seine Writer wieder: Ein Rendern des Foliensatzes mit 100 Folien belegt 5,7 MB, gegenüber 52,7 MB
bei Erweiterung v0.2.0. Damit verzichtet Erweiterung v0.2.2 auf die Überschreibung `GOGC=400`.
Zum Vergleich: Erweiterung v0.2.1 brauchte 72 ms für die erste Vorschau und 44 ms
nach einer Änderung bei Standard-`GOGC` (36 ms mit `GOGC=400`), und v0.2.0 brauchte
47–51 ms.

Erweiterung v0.2.2 maß 52 ms für die erste Vorschau und 24 ms nach einer Änderung,
auf Augenhöhe mit den 23 ms von Marp.
