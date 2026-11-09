---
title: "Paquets de thème"
weight: 42
description: "Un thème en un seul fichier : une archive .gsta.xz avec un theme.hcl, son CSS, ses polices, ses ressources et ses kinds, vérifiée comme une entrée non fiable."
tags: [themes, hcl]
---

Un **paquet de thème** est un seul fichier, `<name>.gsta.xz` (Go-Slides Theme
Archive) : une archive tar (ustar ou pax) compressée avec xz, avec `theme.hcl`
à sa racine et les fichiers qu'il déclare. Il porte le CSS d'un thème, ses
polices, ses logos et ses arrière-plans et, si on le souhaite, des *kinds* de
diapositives, avec un nom, une version et une licence, de sorte qu'un thème
puisse circuler et être installé. Les paquets de thème sont dans la v0.18.0
et les suivantes.

## Le fichier theme.hcl {#the-themehcl-file}

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

`theme.hcl` utilise la syntaxe native de HCL, avec des valeurs littérales
seulement : pas de fonction, pas de variable, pas de `file()` ni de `env()`. Un
nom se compose de lettres minuscules, de chiffres et de tirets, et n'est ni
`default`, ni `gaia`, ni `uncover`. `extends` désigne le thème de base ; chaque
bloc `css` ajoute une feuille de style, dans l'ordre ; chaque bloc `font` et
`asset` désigne un fichier du paquet ; chaque bloc `kind` est un [kind]({{< relref "kinds" >}}).
Un `sha256` sur un bloc `css`, `font` ou `asset` est vérifié à la lecture du
paquet.

Le CSS désigne les polices et les ressources par leur chemin dans le paquet,
relatif au fichier CSS : `background: url(assets/logo.svg)`. Au chargement du
paquet, chacune de ces références devient une URL `data:`, de sorte que la page
est autonome, et un bloc `font` devient une `@font-face`. Un paquet peut aussi
contenir des fichiers `LICENSE`, `NOTICE` et `README` ; tout autre fichier que
`theme.hcl` ne déclare pas est refusé.

## Empaqueter, installer et utiliser {#packing-installing-and-using}

```
slides theme pack dir       write dir/theme.hcl and its files as <name>.gsta.xz
slides theme verify file    run every check on a package
slides theme install file   check it, then install it for the current user
slides theme list           the installed packages
slides theme remove name
```

Les paquets installés se trouvent dans `go-slides/themes`, sous le répertoire de
configuration de l'utilisateur (`os.UserConfigDir()` : `~/Library/Application Support`,
`$XDG_CONFIG_HOME` ou `~/.config`, `%AppData%`). Dans une présentation :

```markdown
---
theme: acme
---
```

trouve un paquet installé, après les thèmes intégrés et `--theme-set`.
`--theme acme.gsta.xz` et `--theme-set acme.gsta.xz` utilisent
directement un fichier, sans l'installer ; voir [Ligne de commande]({{< relref "cli" >}}).

Les *kinds* d'un paquet s'appliquent avec `--theme file.gsta.xz`, avec
`--theme name` pour un paquet installé, ou avec `theme: name` dans le front
matter d'une présentation ; les *kinds* propres à la présentation l'emportent
sur eux.

## Un paquet est une entrée non fiable {#a-package-is-untrusted-input}

Un paquet de thème peut venir de n'importe où ; chaque lecteur (`pack`, `verify`,
`install`, `--theme`, une génération) le vérifie donc de la même façon, et le
refuse avec une erreur qui porte un code (`theme.archive`, `theme.limit`,
`theme.meta`, `theme.css`, `theme.svg`, `theme.font`, `theme.checksum`,
`theme.version`) :

- **Archive.** Un seul flux xz, dont le dictionnaire est borné et dont la
  fenêtre grandit avec les données. Dans le tar : ni `..`, ni chemin absolu, ni
  lettre de lecteur, ni barre oblique inverse dans un nom ; seulement des
  fichiers ordinaires et des répertoires (ni lien symbolique ou physique, ni
  périphérique, ni FIFO) ; pas de nom en double (comparé sans tenir compte de la
  casse) ; seulement des en-têtes ustar et pax (ni extension GNU, ni fichier
  creux), un en-tête pax d'au plus 8 Kio. Les tailles sont celles qui sont
  lues, pas celles qu'un en-tête déclare.
- **Bornes.** Au plus 16 Mio d'archive, 256 entrées, 8 Mio par entrée, 24 Mio
  décompressés, un dictionnaire xz de 16 Mio, un flux décompressé d'au plus 100
  fois la taille de l'archive (de 64 Kio ou plus), `theme.hcl` d'au plus 64 Kio
  et de 32 niveaux d'imbrication, 1 Mio de CSS par fichier, 4 Mio par police ou
  par image.
- **CSS.** Pas de `@import`, pas de `@import-theme` (utilisez `extends`), pas
  de `image-set()`, de `src()` ni d'autre fonction qui charge une ressource,
  pas d'échappement par barre oblique inverse hors d'une chaîne, pas de
  `</style`. Un `url()` est un `#fragment`, une image ou une police `data:`, ou
  un fichier déclaré du paquet, jamais une adresse qui en sort.
- **SVG.** Pas de `script`, de `foreignObject`, de cadre, d'embed, de `set`,
  de DOCTYPE ni de gestionnaire d'événement ; un `href` est un `#fragment` ; pas
  de `url()` qui sort du fichier.
- **Polices et images** sont analysées : une police doit être une police, une
  image doit correspondre à son extension, et chacune a une borne de taille.

Le dossier `testdata/themes/acme` du dépôt est un petit exemple.
