---
title: "Paquetes de tema"
weight: 42
description: "Un tema en un solo archivo: un archivo .gsta.xz con un theme.hcl, su CSS, fuentes, recursos y kinds, comprobado como entrada no fiable."
tags: [themes, hcl]
---

Un **paquete de tema** es un solo archivo, `<name>.gsta.xz` (Go-Slides Theme
Archive): un archivo tar (ustar o pax) comprimido con xz, con `theme.hcl` en su
raíz y los archivos que declara. Lleva el CSS de un tema, sus fuentes,
logotipos y fondos y, si se quiere, kinds de diapositivas, con un nombre, una
versión y una licencia, de modo que un tema pueda circular e instalarse. Los
paquetes de tema están en v0.18.0 y posteriores.

## El archivo theme.hcl {#the-themehcl-file}

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

`theme.hcl` usa la sintaxis nativa de HCL con valores literales solamente: sin
funciones, sin variables, sin `file()` ni `env()`. Un nombre se compone de
letras minúsculas, dígitos y guiones, y no es `default`, `gaia` ni `uncover`.
`extends` nombra el tema base; cada bloque `css` añade una hoja de estilo, en
orden; cada bloque `font` y `asset` nombra un archivo del paquete; cada bloque
`kind` es un [kind]({{< relref "kinds" >}}). Un `sha256` en un bloque `css`,
`font` o `asset` se comprueba cuando se lee el paquete.

El CSS se refiere a las fuentes y a los recursos por su ruta en el paquete,
relativa al archivo CSS: `background: url(assets/logo.svg)`. Cuando se carga el
paquete, cada una de esas referencias se convierte en una URL `data:`, de modo
que la página es autónoma, y un bloque `font` se convierte en una `@font-face`.
Un paquete también puede llevar archivos `LICENSE`, `NOTICE` y `README`; se
rechaza cualquier otro archivo que `theme.hcl` no declare.

## Empaquetar, instalar y usar {#packing-installing-and-using}

```
slides theme pack dir       write dir/theme.hcl and its files as <name>.gsta.xz
slides theme verify file    run every check on a package
slides theme install file   check it, then install it for the current user
slides theme list           the installed packages
slides theme remove name
```

Los paquetes instalados están en `go-slides/themes`, dentro del directorio de
configuración del usuario (`os.UserConfigDir()`: `~/Library/Application Support`,
`$XDG_CONFIG_HOME` o `~/.config`, `%AppData%`). En una presentación:

```markdown
---
theme: acme
---
```

encuentra un paquete instalado, después de los temas integrados y de
`--theme-set`. `--theme acme.gsta.xz` y `--theme-set acme.gsta.xz`
usan un archivo directamente, sin instalarlo; véase [Línea de comandos]({{< relref "cli" >}}).

Los kinds de un paquete se aplican con `--theme file.gsta.xz`, con
`--theme name` para un paquete instalado, o con `theme: name` en el front
matter de una presentación; los kinds propios de la presentación prevalecen
sobre ellos.

## Un paquete es una entrada no fiable {#a-package-is-untrusted-input}

Un paquete de tema puede venir de cualquier parte, así que cada lector (`pack`,
`verify`, `install`, `--theme`, una construcción) lo comprueba de la misma
manera y lo rechaza con un error que lleva un código (`theme.archive`,
`theme.limit`, `theme.meta`, `theme.css`, `theme.svg`, `theme.font`,
`theme.checksum`, `theme.version`):

- **Archivo.** Un solo flujo xz, cuyo diccionario está acotado y cuya ventana
  crece con los datos. Dentro del tar: ni `..`, ni ruta absoluta, ni letra de
  unidad, ni barra invertida en un nombre; solo archivos regulares y
  directorios (ni enlace simbólico o duro, ni dispositivo, ni FIFO); ni nombre
  duplicado (comparado sin distinguir mayúsculas); solo cabeceras ustar y pax
  (ni extensión de GNU, ni archivo disperso), una cabecera pax de como máximo
  8 KiB. Los tamaños son los que se leen, no los que declara una cabecera.
- **Límites.** Como máximo 16 MiB de archivo, 256 entradas, 8 MiB por entrada,
  24 MiB descomprimidos, un diccionario xz de 16 MiB, un flujo descomprimido de
  como máximo 100 veces el tamaño del archivo (de 64 KiB o más), `theme.hcl` de
  como máximo 64 KiB y 32 niveles de anidamiento, 1 MiB de CSS por archivo,
  4 MiB por fuente o imagen.
- **CSS.** Ni `@import`, ni `@import-theme` (use `extends`), ni `image-set()`,
  `src()` u otra función que cargue un recurso, ni escapes con barra invertida
  fuera de una cadena, ni `</style`. Un `url()` es un `#fragment`, una imagen o
  fuente `data:`, o un archivo declarado del paquete, nunca uno que salga de él.
- **SVG.** Ni `script`, `foreignObject`, marcos, embed, `set`, DOCTYPE ni
  manejadores de eventos; un `href` es un `#fragment`; ni `url()` que salga del
  archivo.
- **Fuentes e imágenes** se analizan: una fuente debe ser una fuente, una
  imagen debe coincidir con su extensión, y cada una tiene un límite de tamaño.

El directorio `testdata/themes/acme` del repositorio es un pequeño ejemplo.
