---
title: "Galería"
weight: 12
description: "Doce presentaciones pequeñas, cada una con un solo tema bien mostrado: temas, kinds, plantillas, HCL, matemáticas, imágenes, código, emoji, transiciones y una charla completa. Ábralas en el Playground."
tags: [gallery, examples]
---

Doce presentaciones pequeñas de la carpeta
[`examples/`](https://github.com/go-slides/slides/tree/main/examples) de
go-slides, cada una con un solo tema bien mostrado. Las miniaturas son la primera
diapositiva de cada presentación, renderizada por el propio go-slides al
construir este sitio; la CI construye cada presentación en HTML y PDF, así que
ninguna se rompe.

**Open in the Playground** carga la presentación en su navegador, mediante un
enlace que no envía nada a ningún servidor. Una presentación con más archivos
además de `deck.md` (un archivo de kinds, partes incluidas, imágenes) no puede
viajar en un enlace: descargue su **.zip** e impórtelo con el panel de archivos
del Playground, o descomprímalo y constrúyalo con
`slides build --allow-local-files deck.md`.

Las imágenes son SVG originales de los autores de go-slides, publicados bajo
[CC0 1.0](https://creativecommons.org/publicdomain/zero/1.0/), y las
presentaciones que las usan lo indican.

{{< gallery >}}
