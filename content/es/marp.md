---
title: "Compatibilidad con Marp"
linkTitle: "Marp"
weight: 20
description: "El Markdown que lee slides es el de Marpit y Marp Core, no un dialecto suyo."
tags: [marp]
---

Una presentación escrita para [Marp](https://marp.app) es una presentación de
go-slides. Las reglas siguientes son las de [Marpit](https://marpit.marp.app) y
de [Marp Core](https://github.com/marp-team/marp-core); donde go-slides difiere,
esta página lo indica.

## Diapositivas {#slides}

Una regla horizontal (`---` en una línea propia, como divide Marpit) inicia una
nueva diapositiva, y también lo hace un título del nivel que indique la directiva
`headingDivider`. El front matter, entre los dos primeros `---`, contiene las
directivas globales y los datos de la presentación.

## Directivas {#directives}

| Tipo | Dónde | Se aplica a |
| --- | --- | --- |
| global | el front matter, o `<!-- theme: gaia -->` en cualquier lugar | toda la presentación: `theme`, `style`, `headingDivider`, `lang`… |
| local | `<!-- paginate: true -->` | esta diapositiva y las siguientes |
| de una sola diapositiva (spot) | `<!-- _class: lead -->` (guion bajo) | solo esta diapositiva |

Las directivas locales son las de Marpit: `paginate`, `header`, `footer`,
`class`, `backgroundColor`, `backgroundImage`, `backgroundPosition`,
`backgroundRepeat`, `backgroundSize`, `color`. Las directivas se analizan como
YAML tolerante, igual que hace Marp Core; `--strict-yaml` las analiza como lo
hace Marpit.

## Imágenes {#images}

`![bg](image.jpg)` convierte una imagen en el fondo de la diapositiva, con las
palabras clave de Marp: `bg contain`, `bg cover`, `bg 50%`, `bg left`/`bg right`
para un fondo dividido, varias imágenes `bg` una junto a otra, y los filtros de
imagen (`blur`, `sepia`, `grayscale`…). Los tamaños de las imágenes en línea
(`w:200px`, `h:`) se leen como los lee Marp.

## Alrededor de la diapositiva {#around-the-slide}

- **Encabezado y pie de página**: las directivas `header:` y `footer:`, en Markdown.
- **Paginación**: `paginate: true`.
- **Notas del orador**: un comentario HTML que no es una directiva es una nota.
  `slides notes` las imprime.
- **Listas fragmentadas**: una lista escrita con `*` (o `1)`) aparece elemento a
  elemento, como en Marp.
- **`<style scoped>`**: CSS para una sola diapositiva.
- **HTML sin procesar**: la lista de permitidos de Marp Core por defecto; `--html`
  lo permite todo.
- **Emoji**: los `:shortcodes:` y los emoji Unicode se convierten en imágenes
  Twemoji, como en Marp Core, pero incrustadas en lugar de cargadas desde una
  CDN: véase [Emoji]({{< relref "emoji" >}}).

## Temas {#themes}

`default`, `gaia` y `uncover` vienen incorporados y se eligen con `theme:`. Un
tema propio es un archivo CSS con el comentario `/* @theme name */` de Marp, que
se indica con `--theme-set` o se lista bajo `themes:` en `slides.yaml`.

## En qué difiere go-slides {#where-go-slides-differs}

- **Plantillas.** `{{ }}` en el cuerpo es una acción de plantilla de Go (véase
  [Plantillas]({{< relref "templates" >}})). Una presentación que muestre `{{`
  como texto fuera de un bloque de código necesita `--no-template`, o el texto
  escapado como lo escapan las plantillas de Go. Los bloques de código y los
  fragmentos de código en línea se dejan tal cual.
- **Kinds** (tipos de diapositiva). `kind:` y `_kind:` son directivas que Marp no tiene (véase
  [Kinds]({{< relref "kinds" >}})). `slides expand` las sustituye por las
  directivas que representan, de modo que la presentación expandida es Marp
  plano.
- **Salidas.** HTML, PNG, JPEG, PDF, PPTX y ODP, mediante bibliotecas de Go en
  lugar de un navegador sin interfaz (véase
  [Línea de comandos]({{< relref "cli#images-pdf-pptx-and-odp" >}}) para el
  límite conocido). En v0.4.0, como en marp-cli, un PPTX contiene imágenes; desde
  v0.5.0 sus diapositivas son cuadros de texto y formas editables, y
  `--pptx-raster` conserva las imágenes.
- **Títulos ajustados.** `# <!-- fit --> Title` es compatible: el título ocupa
  todo el ancho de la diapositiva en una línea, en la página y en todas las
  exportaciones.
- **Matemáticas y código.** `$…$` y `$$…$$` siguen las reglas de delimitadores de
  Marp Core y las compone go-tex/math; el código en bloques lo resalta go-rouge
  con los nombres de clase de highlight.js (véase
  [Matemáticas y código]({{< relref "math-and-code" >}})).
