---
title: "HCL"
weight: 45
description: "HCL2 überall dort, wo ein Foliensatz YAML erwartet: ein HCL-Front-Matter, slides.hcl, Kinds- und Datendateien, mit Ausdrücken und Funktionen."
tags: [hcl, configuration]
---

{{< callout type="warning" >}}
HCL gibt es ab v0.3.0 ([go-slides/slides#13](https://github.com/go-slides/slides/pull/13)).
{{< /callout >}}

Überall, wo `slides` YAML liest, liest es auch die native Syntax von
[HCL2](https://github.com/hashicorp/hcl/blob/main/hclsyntax/spec.md) – wegen
ihrer Ausdrücke und Funktionen.

| Wo | YAML | HCL |
| --- | --- | --- |
| Front-Matter | `---` … `---` | `---hcl` … `---` |
| Konfiguration | `slides.yaml` | `slides.hcl`, wird zuerst gesucht |
| Kinds-Dateien | `--kinds kinds.yaml`, `kinds: kinds.yaml` | jede Datei namens `*.hcl` |
| Datendateien | `--data data.yaml` oder `.json` | jede Datei namens `*.hcl` |

Direktivenkommentare (`<!-- class: invert -->`) bleiben wie in Marp YAML, und der
Rumpf des Foliensatzes bleibt eine [Go-Vorlage]({{< relref "templates" >}}).

## Ein Front-Matter in HCL {#a-front-matter-in-hcl}

```hcl
---hcl
theme    = "gaia"
paginate = true
title    = "Quarterly review"
data = {
  company = "ACME"
  quarter = 3
}
header = format("%s — Q%d", data.company, data.quarter)
footer = "${upper(data.company)} · ${page()}/${total()}"
date   = formatdate("D MMMM YYYY", now())

kind "chapter" {
  extends = "section"
  class   = ["lead", "invert"]
}
---

# {{ .title }}
```

- **Attribute** sind die YAML-Schlüssel unter denselben Namen: die Direktiven von Marp
  (`theme`, `paginate`, `header`, `footer`, `class`, `backgroundColor` …),
  die Metadaten (`title`, `author` …), `template`, `delims`, `kinds` (Dateipfade)
  und `data`, ein Objekt.
- **`kind "<name>" { … }`**-Blöcke definieren [Kinds (Folientypen)]({{< relref "kinds" >}})
  mit den lokalen Direktiven und `extends`.
- **Werte** erreichen die Direktiven in den Zeichenkettenformen, die YAML ihnen gibt –
  `true` wird zu `"true"`, eine Liste von Klassen bleibt eine Liste –, sodass sich eine Direktive
  gleich verhält, welche Syntax sie auch gesetzt hat. `null` lässt eine Direktive ungesetzt.
- **`data`** behält seine Typen: In Vorlagen ist eine ganze Zahl ein `int`,
  andere Zahlen sind `float64`, und Listen und Objekte bleiben Listen und Maps, sodass
  `{{ add .quarter 1 }}` funktioniert.
- Ein HCL-Wert ist endgültig: Go-Vorlagen laufen kein zweites Mal darüber.

## Variablen {#variables}

`data` sind die Vorlagendaten: das `data` des Front-Matters, überlagert von dem
aus `slides.hcl`, von `--data`-Dateien und von `--set`. `"${data.company}"` funktioniert
überall dort, wo auch `{{ .company }}` funktionieren würde.

Attribute des Front-Matters dürfen auf `data.*` verweisen, aber **nicht aufeinander**:
`footer = title` ist ein Fehler. Spot-Direktiven (`_class` …) sind keine HCL-
Bezeichner; schreiben Sie sie in einen Kommentar.

## Funktionen {#functions}

| | |
| --- | --- |
| `page()`, `total()` | die Seitennummer und die Nummer der letzten Seite, wie `{{ page }}` und `{{ total }}`: eine Fußzeile paginiert |
| `now()`, `timestamp()` | die Zeit, RFC 3339, lokal und UTC; `formatdate` formatiert sie |
| `env(name)` | eine Umgebungsvariable, `""` wenn nicht gesetzt |
| `file(path)` | eine UTF-8-Textdatei von höchstens 1 MiB, relativ zum Foliensatz oder zur Konfigurationsdatei |
| Zahlen | `abs ceil floor log max min parseint pow signum` |
| Zeichenketten | `chomp format formatlist join lower regex regexall regexreplace replace split strlen strrev substr title trim trimprefix trimspace trimsuffix upper` |
| Collections | `chunklist coalesce coalescelist compact concat contains distinct element flatten keys length lookup merge range reverse slice sort values zipmap` |
| Kodierungen | `csvdecode jsondecode jsonencode` |
| Datumsangaben | `formatdate timeadd` |

Alle Zeilen außer den ersten vier stammen aus der Standardbibliothek von
[go-cty](https://github.com/zclconf/go-cty). Keine davon greift auf das Netzwerk zu.

## slides.hcl {#slideshcl}

```hcl
kinds  = "kinds.hcl"          # or a list of files
themes = ["themes/acme.css"]
data   = { company = "ACME" } # or file paths
html   = false
lang   = "en"

kind "quote" {
  extends = "blank"
  footer  = "${data.company}"
}
```

Ein unbekanntes Attribut ist ein Fehler, kein stilles Nichts. Eine Kinds-Datei enthält
`kind`-Blöcke; ihre übrigen Attribute werden ignoriert, sodass `slides.hcl` als
solche dienen kann. Die Attribute einer Datendatei sind die Schlüssel der Daten:

```hcl
# data.hcl
company = "ACME"
year    = 2026
motto   = trimspace(file("motto.txt"))
```

## Fehler {#errors}

Fehler sind die Diagnosen von HCL, an der Position in der Foliensatz-Datei selbst:
`footer = data.compnay` in der vierten Zeile des Foliensatzes ergibt

```
deck.md:4:14: Unsupported attribute: This object does not have an attribute named "compnay".
```

## Marp und `slides expand` {#marp-and-slides-expand}

Auch Marp liest `---hcl` als Beginn eines Front-Matters – seine Front-Matter-
Regel greift bei jeder ersten Zeile, die mit drei Strichen beginnt – und wendet davon nichts an,
da der Block keine YAML-Map ist: Der Foliensatz wird in Marp trotzdem gerendert,
nur ohne seine Direktiven. Ein Test legt dies gegen die eigene Ausgabe von Marp Core fest.

`slides expand` schreibt das Front-Matter als YAML zurück, sodass der expandierte Foliensatz
reines Marp ist und so gerendert wird, wie `slides` das Original rendert:

```yaml
---
theme: "gaia"
paginate: true
title: "Quarterly review"
header: "ACME — Q3"
footer: "ACME · 1/2"
date: "19 October 2026"
---
```

## Lizenzen {#licences}

Die HCL-Unterstützung bindet [hashicorp/hcl](https://github.com/hashicorp/hcl)
(MPL-2.0, unverändert als Bibliothek verwendet) und
[zclconf/go-cty](https://github.com/zclconf/go-cty) (MIT) ein.
