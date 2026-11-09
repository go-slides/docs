---
title: "Theme-Pakete"
weight: 42
description: "Ein Theme als eine Datei: ein .gsta.xz-Archiv mit einer theme.hcl, ihrem CSS, Schriften, Ressourcen und Kinds, geprüft wie nicht vertrauenswürdige Eingaben."
tags: [themes, hcl]
---

Ein **Theme-Paket** ist eine einzige Datei, `<name>.gsta.xz` (Go-Slides Theme
Archive): ein mit xz komprimiertes tar-Archiv (ustar oder pax) mit `theme.hcl`
im Wurzelverzeichnis und den Dateien, die es deklariert. Es enthält das CSS
eines Themes, Schriften, Logos und Hintergründe und, wenn gewünscht,
Folien-Kinds, mit einem Namen, einer Version und einer Lizenz, sodass ein Theme
weitergegeben und installiert werden kann. Theme-Pakete gibt es nach v0.17.1.

## Die Datei theme.hcl {#the-themehcl-file}

```hcl
# theme.hcl
theme {
  name        = "acme"
  version     = "1.0.0"                 # semantic version
  description = "Example theme"
  authors     = ["go-slides contributors"]
  license     = "MIT"                   # SPDX expression
  homepage    = "https://go-slides.github.io"
  extends     = "default"               # the base theme, built in or installed
  min_slides  = "0.18.0"                # oldest go-slides it works with
}

css "theme.css" {}                      # in order; sha256 = "…" is optional

font "Roboto Mono" {
  weight  = 400                         # or "100 900"
  style   = "normal"
  file    = "fonts/RobotoMono-Regular.ttf"
  license = "OFL-1.1"                   # required for a font
}

asset "logo" { file = "assets/logo.svg" }

kind "acme-title" {                     # optional slide kinds
  class    = "lead"
  paginate = "false"
}
```

`theme.hcl` verwendet die native Syntax von HCL, nur mit Literalwerten: keine
Funktion, keine Variable, kein `file()` und kein `env()`. Ein Name besteht aus
Kleinbuchstaben, Ziffern und Bindestrichen und ist weder `default` noch `gaia`
noch `uncover`. `extends` nennt das Basis-Theme; jeder `css`-Block fügt der
Reihe nach ein Stylesheet hinzu; jeder `font`- und `asset`-Block nennt eine
Datei des Pakets; jeder `kind`-Block ist ein [Kind]({{< relref "kinds" >}}).
Ein `sha256` an einem `css`-, `font`- oder `asset`-Block wird beim Lesen des
Pakets geprüft.

Das CSS verweist auf Schriften und Ressourcen über ihren Pfad im Paket,
relativ zur CSS-Datei: `background: url(assets/logo.svg)`. Beim Laden des
Pakets wird jeder solche Verweis zu einer `data:`-URL, sodass die Seite in sich
geschlossen ist, und ein `font`-Block wird zu einer `@font-face`. Ein Paket
kann außerdem die Dateien `LICENSE`, `NOTICE` und `README` enthalten; jede
andere Datei, die `theme.hcl` nicht deklariert, wird abgelehnt.

## Packen, installieren und verwenden {#packing-installing-and-using}

```
slides theme pack dir       write dir/theme.hcl and its files as <name>.gsta.xz
slides theme verify file    run every check on a package
slides theme install file   check it, then install it for the current user
slides theme list           the installed packages
slides theme remove name
```

Installierte Pakete liegen in `go-slides/themes` im Konfigurationsverzeichnis
des Benutzers (`os.UserConfigDir()`: `~/Library/Application Support`,
`$XDG_CONFIG_HOME` oder `~/.config`, `%AppData%`). In einem Foliensatz:

```markdown
---
theme: acme
---
```

findet ein installiertes Paket, nach den eingebauten Themes und `--theme-set`.
`--theme acme.gsta.xz` und `--theme-set acme.gsta.xz` verwenden eine
Datei direkt, ohne sie zu installieren; siehe [Befehlszeile]({{< relref "cli" >}}).

Die Kinds eines Pakets gelten mit `--theme file.gsta.xz`, mit `--theme name`
für ein installiertes Paket oder mit `theme: name` im Front Matter eines
Foliensatzes; die eigenen Kinds des Foliensatzes haben Vorrang.

Das `extends` eines Pakets benennt sein Basis-Theme, das in dieser Reihenfolge
gesucht wird: eingebaut (oder mit `--theme-set` angegeben), dann installiert,
dann ein danebenliegendes `<base>.gsta.xz` im selben Verzeichnis wie die
Paketdatei, sodass eine Familie von Themes (eine Basis, mehrere kleine Kinder)
direkt aus ihren Dateien ausprobiert werden kann. Der Fehler nennt die drei
durchsuchten Orte. Eine Kette von Paketen ist höchstens 8 lang, und ein Zyklus
wird abgelehnt (`theme.meta`).

## Ein Paket ist eine nicht vertrauenswürdige Eingabe {#a-package-is-untrusted-input}

Ein Theme-Paket kann von überall stammen; deshalb prüft jeder Leser (`pack`,
`verify`, `install`, `--theme`, ein Build) es auf dieselbe Weise und lehnt es
mit einem Fehler ab, der einen Code trägt (`theme.archive`, `theme.limit`,
`theme.meta`, `theme.css`, `theme.svg`, `theme.font`, `theme.checksum`,
`theme.version`):

- **Archiv.** Ein xz-Stream, dessen Wörterbuch begrenzt ist und dessen Fenster
  mit den Daten wächst. Im tar: kein `..`, kein absoluter Pfad, kein
  Laufwerksbuchstabe und kein Backslash in einem Namen; nur reguläre Dateien und
  Verzeichnisse (kein symbolischer oder harter Link, kein Gerät, keine FIFO);
  kein doppelter Name (ohne Beachtung der Groß- und Kleinschreibung verglichen);
  nur ustar- und pax-Header (keine GNU-Erweiterung, keine Sparse-Datei), ein
  pax-Header von höchstens 8 KiB. Größen sind das, was gelesen wird, nicht das,
  was ein Header angibt.
- **Grenzen.** Höchstens 16 MiB Archiv, 256 Einträge, 8 MiB je Eintrag, 24 MiB
  entpackt, ein xz-Wörterbuch von 16 MiB, ein entpackter Stream von höchstens
  dem 100-Fachen der Archivgröße (ab 64 KiB), `theme.hcl` höchstens 64 KiB und
  32 Klammern tief, 1 MiB CSS je Datei, 4 MiB je Schrift oder Bild.
- **CSS.** Kein `@import`, kein `@import-theme` (verwenden Sie `extends`), kein
  `image-set()`, `src()` oder eine andere Funktion, die eine Ressource lädt,
  keine Backslash-Escape-Sequenz außerhalb eines Strings, kein `</style`. Ein
  `url()` ist ein `#fragment`, ein Bild oder eine Schrift mit `data:` oder
  eine deklarierte Datei des Pakets, nie eine Adresse, die das Paket verlässt.
- **SVG.** Kein `script`, kein `foreignObject`, kein Frame, kein embed, kein
  `set`, kein DOCTYPE und kein Event-Handler; ein `href` ist ein `#fragment`;
  kein `url()`, das die Datei verlässt.
- **Schriften und Bilder** werden geparst: Eine Schrift muss eine Schrift sein,
  ein Bild muss zu seiner Dateiendung passen, und für beide gilt eine
  Größengrenze.

Das Verzeichnis `testdata/themes/acme` des Repositorys ist ein kleines Beispiel.
