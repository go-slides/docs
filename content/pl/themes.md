---
title: "Pakiety motywów"
weight: 42
description: "Motyw w jednym pliku: archiwum .gsta.xz z theme.hcl, jego CSS, czcionkami, zasobami i kinds, sprawdzane jak niezaufane dane wejściowe."
tags: [themes, hcl]
---

**Pakiet motywu** to jeden plik, `<name>.gsta.xz` (Go-Slides Theme Archive):
archiwum tar (ustar lub pax) skompresowane xz, z plikiem `theme.hcl` w katalogu głównym i plikami, które ten deklaruje.
Zawiera CSS motywu, czcionki, logotypy i tła oraz, jeśli trzeba, rodzaje
slajdów (kinds), a także nazwę, wersję i licencję, dzięki czemu motyw można
przekazywać dalej i instalować. Pakiety motywów pojawiają się po v0.17.1.

## The theme.hcl file {#the-themehcl-file}

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

`theme.hcl` używa natywnej składni HCL wyłącznie z wartościami literalnymi:
bez funkcji, bez zmiennych, bez `file()` i `env()`. Nazwa składa się z małych
liter, cyfr i myślników i nie jest `default`, `gaia` ani `uncover`. `extends`
wskazuje motyw bazowy; każdy blok `css` dodaje arkusz, po kolei; każdy blok
`font` i `asset` wskazuje plik pakietu; każdy blok `kind` to [kind]({{< relref "kinds" >}}).
Wartość `sha256` w bloku `css`, `font` lub `asset` jest sprawdzana przy
odczycie pakietu.

CSS odwołuje się do czcionek i zasobów przez ich ścieżkę w pakiecie,
względem pliku CSS: `background: url(assets/logo.svg)`. Przy ładowaniu pakietu
każde takie odwołanie staje się adresem `data:`, tak aby strona była
samodzielna, a blok `font` staje się `@font-face`. Pakiet może też zawierać
pliki `LICENSE`, `NOTICE` i `README`; każdy inny plik, którego `theme.hcl` nie
deklaruje, jest odrzucany.

## Packing, installing and using {#packing-installing-and-using}

```
slides theme pack dir       write dir/theme.hcl and its files as <name>.gsta.xz
slides theme verify file    run every check on a package
slides theme install file   check it, then install it for the current user
slides theme list           the installed packages
slides theme remove name
```

Zainstalowane pakiety znajdują się w `go-slides/themes` w katalogu
konfiguracyjnym użytkownika (`os.UserConfigDir()`: `~/Library/Application Support`,
`$XDG_CONFIG_HOME` lub `~/.config`, `%AppData%`). W prezentacji:

```markdown
---
theme: acme
---
```

znajduje zainstalowany pakiet, po motywach wbudowanych i `--theme-set`.
`--theme acme.gsta.xz` i `--theme-set acme.gsta.xz` używają pliku
bezpośrednio, bez instalowania; zobacz [Wiersz poleceń]({{< relref "cli" >}}).

Kinds pakietu stosują się z `--theme file.gsta.xz`, z `--theme name`
dla zainstalowanego pakietu lub z `theme: name` w front matter prezentacji;
własne kinds prezentacji mają pierwszeństwo przed nimi.

## A package is untrusted input {#a-package-is-untrusted-input}

Pakiet motywu może pochodzić skądkolwiek, więc każdy czytnik (`pack`, `verify`,
`install`, `--theme`, budowanie) sprawdza go tak samo i odrzuca go błędem
z kodem (`theme.archive`, `theme.limit`, `theme.meta`, `theme.css`,
`theme.svg`, `theme.font`, `theme.checksum`, `theme.version`):

- **Archiwum.** Jeden strumień xz, którego słownik jest ograniczony, a okno
  rośnie wraz z danymi. W tar: bez `..`, ścieżki bezwzględnej, litery dysku ani
  odwrotnego ukośnika w nazwie; tylko zwykłe pliki i katalogi (bez dowiązania
  symbolicznego ani twardego, urządzenia ani FIFO); bez zduplikowanej nazwy
  (porównywanej bez rozróżniania wielkości liter); tylko nagłówki ustar i pax
  (bez rozszerzeń GNU, bez plików rzadkich), nagłówek pax co najwyżej 8 KiB.
  Rozmiary to to, co faktycznie odczytano, a nie to, co deklaruje nagłówek.
- **Limity.** Co najwyżej 16 MiB archiwum, 256 wpisów, 8 MiB na wpis, 24 MiB
  po rozpakowaniu, słownik xz 16 MiB, strumień po rozpakowaniu co najwyżej 100
  razy większy od archiwum (od 64 KiB), `theme.hcl` co najwyżej 64 KiB i 32
  poziomy zagnieżdżenia nawiasów, 1 MiB CSS na plik, 4 MiB na czcionkę lub
  obraz.
- **CSS.** Bez `@import`, bez `@import-theme` (użyj `extends`), bez
  `image-set()`, `src()` ani innej funkcji ładującej zasób, bez sekwencji
  z odwrotnym ukośnikiem poza łańcuchem znaków, bez `</style`. `url()` to
  `#fragment`, obraz lub czcionka `data:` albo zadeklarowany plik pakietu,
  nigdy taki, który z niego wychodzi.
- **SVG.** Bez `script`, `foreignObject`, ramki, embed, `set`, DOCTYPE ani
  procedury obsługi zdarzenia; `href` to `#fragment`; bez `url()`, który
  wychodzi poza plik.
- **Czcionki i obrazy** są parsowane: czcionka musi być czcionką, obraz musi
  zgadzać się ze swoim rozszerzeniem, a każde z nich ma limit rozmiaru.

Katalog `testdata/themes/acme` w repozytorium to mały przykład.
