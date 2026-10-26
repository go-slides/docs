---
title: "Documentación de go-slides"
linkTitle: "Inicio"
type: docs
cascade:
  type: docs
description: "Presentaciones compatibles con Marp en Go puro (CGO=0), con interpolación de plantillas de Go y kinds de diapositiva."
---

**`slides` convierte una presentación de Marp en una presentación lista para
proyectar, en Go puro.** Lee el Markdown que lee [Marp](https://marp.app): las
directivas de Marpit, `---` entre diapositivas, los fondos `![bg]`, las notas del
orador y los temas `default`, `gaia` y `uncover`. Y añade dos cosas que Marp no
tiene:

- **[Plantillas de Go]({{< relref "templates" >}})** dentro de la presentación,
  evaluadas antes de analizar el Markdown, de modo que una presentación puede
  generar sus propias diapositivas;
- **[kinds de diapositiva]({{< relref "kinds" >}})**: una diapositiva declara lo
  que es (`title`, `section`…) en lugar de repetir cómo se ve.

{{< callout type="warning" >}}
**v0.14.0.** Estas páginas describen
[go-slides/slides v0.14.0](https://github.com/go-slides/slides/releases/tag/v0.14.0)
y se publican en `/docs/0.14/`.
{{< /callout >}}

## Qué hace hoy {#what-it-does-today}

| | |
| --- | --- |
| `slides build` (el comando por defecto) | la presentación como una sola página HTML: cada diapositiva es una `<section>` enmarcada en SVG que se ajusta a la ventana, el CSS del tema y un visor con teclado |
| `slides build --pdf`, `--images png`, `--image jpeg` | la presentación como PDF (una página por diapositiva, con índice) o como imágenes PNG/JPEG, sin navegador: [go-webengine](https://github.com/go-webengine) renderiza y [go-pdfkit](https://github.com/go-pdfkit) escribe el PDF, en vectores desde v0.4.0 y con el texto seleccionable: [PDF vectorial]({{< relref "cli#vector-pdf" >}}) |
| matemáticas y código | `$…$` y `$$…$$` compuestos por go-tex/math, bloques de código resaltados por go-rouge, en HTML, imágenes y PDF: [Matemáticas y código]({{< relref "math-and-code" >}}) |
| emoji | `:smile:` y 😄 se convierten en imágenes Twemoji incrustadas, de modo que no necesitan red y se dibujan en todas las salidas: [Emoji]({{< relref "emoji" >}}) (desde v0.4.0) |
| `slides build --pptx` | la presentación como archivo de PowerPoint con sus notas del orador y los metadatos de la presentación: una imagen por diapositiva en v0.4.0, como en marp-cli; **editable** desde v0.5.0 con cuadros de texto, formas, imágenes y fuentes incrustadas ([PPTX y ODP editables]({{< relref "cli#editable-pptx-and-odp" >}})) |
| `slides build --odp` | lo mismo como presentación OpenDocument (`.odp`) |
| títulos ajustados | `# <!-- fit --> Title` ocupa todo el ancho de la diapositiva en una línea, en la página y en todas las exportaciones |
| `slides expand` | la presentación como Markdown de Marp plano: plantillas evaluadas, kinds sustituidos por las directivas locales que representan |
| `slides notes` | las notas del orador |
| `slides kinds`, `slides themes` | lo que una presentación puede nombrar |
| `slides.yaml` | kinds, temas, datos, política de HTML e idioma compartidos por una carpeta de presentaciones |
| HCL | front matter `---hcl`, `slides.hcl`, archivos de kinds y de datos en HCL2, con expresiones y funciones: [HCL]({{< relref "hcl" >}}) |

## Previsto {#planned}

Todavía no está hecho, y se indica para que nadie tenga que adivinarlo: las
fórmulas como vectores en el PDF (ahora son mapas de bits). La diferencia con
marp-cli está en el camino: él renderiza imágenes y PDF controlando un Chrome
sin interfaz; aquí lo hacen bibliotecas de Go.

## Páginas {#pages}

- [Primeros pasos]({{< relref "getting-started" >}})
- [Galería]({{< relref "gallery" >}}): doce presentaciones de ejemplo, para abrir en el playground
- [Compatibilidad con Marp]({{< relref "marp" >}})
- [Plantillas]({{< relref "templates" >}})
- [Kinds]({{< relref "kinds" >}})
- [HCL]({{< relref "hcl" >}})
- [Matemáticas y código]({{< relref "math-and-code" >}})
- [Emoji]({{< relref "emoji" >}})
- [Línea de comandos]({{< relref "cli" >}})
