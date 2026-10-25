---
title: "HCL"
weight: 45
description: "HCL2 allí donde una presentación admite YAML: un front matter en HCL, slides.hcl, archivos de kinds y de datos, con expresiones y funciones."
tags: [hcl, configuration]
---

{{< callout type="warning" >}}
HCL llegó en v0.3.0 ([go-slides/slides#13](https://github.com/go-slides/slides/pull/13)).
{{< /callout >}}

Allí donde `slides` lee YAML, también lee la sintaxis nativa de
[HCL2](https://github.com/hashicorp/hcl/blob/main/hclsyntax/spec.md), por sus
expresiones y sus funciones.

| Dónde | YAML | HCL |
| --- | --- | --- |
| front matter | `---` … `---` | `---hcl` … `---` |
| configuración | `slides.yaml` | `slides.hcl`, que se busca primero |
| archivos de kinds | `--kinds kinds.yaml`, `kinds: kinds.yaml` | cualquier archivo llamado `*.hcl` |
| archivos de datos | `--data data.yaml` o `.json` | cualquier archivo llamado `*.hcl` |

Los comentarios de directivas (`<!-- class: invert -->`) siguen siendo YAML, como
en Marp, y el cuerpo de la presentación sigue siendo una
[plantilla de Go]({{< relref "templates" >}}).

## Un front matter en HCL {#a-front-matter-in-hcl}

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

- Los **atributos** son las claves de YAML con los mismos nombres: las directivas
  de Marp (`theme`, `paginate`, `header`, `footer`, `class`, `backgroundColor`…),
  los metadatos (`title`, `author`…), `template`, `delims`, `kinds` (rutas de
  archivos) y `data`, un objeto.
- Los bloques **`kind "<name>" { … }`** definen [kinds]({{< relref "kinds" >}}) (tipos
  de diapositiva), con las directivas locales y `extends`.
- Los **valores** llegan a las directivas en las formas de cadena que les da
  YAML (`true` es `"true"`, una lista de clases es una lista), de modo que una
  directiva se comporta igual sea cual sea la sintaxis que la fije. `null` deja
  una directiva sin definir.
- **`data`** conserva sus tipos: en las plantillas un número entero es un `int`,
  los demás números son `float64`, y las listas y los objetos siguen siendo
  listas y mapas, de modo que `{{ add .quarter 1 }}` funciona.
- Un valor HCL es definitivo: las plantillas de Go no se ejecutan sobre él por
  segunda vez.

## Variables {#variables}

`data` son los datos de la plantilla: el propio `data` del front matter,
superpuesto al de `slides.hcl`, a los archivos de `--data` y a `--set`.
`"${data.company}"` funciona allí donde funcionaría `{{ .company }}`.

Los atributos del front matter pueden hacer referencia a `data.*`, pero **no
entre sí**: `footer = title` es un error. Las directivas locales de una sola
diapositiva (`_class`…) no son identificadores de HCL; escríbalas en un
comentario.

## Funciones {#functions}

| | |
| --- | --- |
| `page()`, `total()` | el número de página y el número de la última página, como `{{ page }}` y `{{ total }}`: un pie de página pagina |
| `now()`, `timestamp()` | la hora, en RFC 3339, local y UTC; `formatdate` le da formato |
| `env(name)` | una variable de entorno, `""` si no está definida |
| `file(path)` | un archivo de texto UTF-8 de como máximo 1 MiB, relativo a la presentación o al archivo de configuración |
| números | `abs ceil floor log max min parseint pow signum` |
| cadenas | `chomp format formatlist join lower regex regexall regexreplace replace split strlen strrev substr title trim trimprefix trimspace trimsuffix upper` |
| colecciones | `chunklist coalesce coalescelist compact concat contains distinct element flatten keys length lookup merge range reverse slice sort values zipmap` |
| codificaciones | `csvdecode jsondecode jsonencode` |
| fechas | `formatdate timeadd` |

Todas salvo las cuatro primeras filas son la biblioteca estándar de
[go-cty](https://github.com/zclconf/go-cty). Ninguna accede a la red.

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

Un atributo desconocido es un error, no una operación nula silenciosa. Un archivo
de kinds contiene bloques `kind`; sus demás atributos se ignoran, así que
`slides.hcl` puede servir como uno. Los atributos de un archivo de datos son las
claves de los datos:

```hcl
# data.hcl
company = "ACME"
year    = 2026
motto   = trimspace(file("motto.txt"))
```

## Errores {#errors}

Los errores son los diagnósticos de HCL, situados en el propio archivo de la
presentación: `footer = data.compnay` en la cuarta línea de la presentación da

```
deck.md:4:14: Unsupported attribute: This object does not have an attribute named "compnay".
```

## Marp y `slides expand` {#marp-and-slides-expand}

Marp también lee `---hcl` como el inicio de un front matter (su regla de front
matter se abre con cualquier primera línea que empiece por tres guiones) y, al no
ser el bloque un mapa YAML, no aplica nada de él: la presentación se sigue
renderizando en Marp, pero sin sus directivas. Una prueba lo fija contra la
propia salida de Marp Core.

`slides expand` vuelve a escribir el front matter como YAML, de modo que la
presentación expandida es Marp plano y se renderiza como `slides` renderiza la
original:

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

## Licencias {#licences}

La compatibilidad con HCL enlaza [hashicorp/hcl](https://github.com/hashicorp/hcl)
(MPL-2.0, usado sin modificar como biblioteca) y
[zclconf/go-cty](https://github.com/zclconf/go-cty) (MIT).
