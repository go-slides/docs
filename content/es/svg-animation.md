---
title: "SVG animado"
weight: 40
description: "Una imagen SVG con animaciones SMIL o CSS: reproducida por el navegador en HTML, formas animadas en PowerPoint y OpenDocument, un GIF animado donde estas no pueden expresarlo, el estado inicial en PDF e imágenes."
tags: [svg, animation]
---

Una imagen SVG que se anima por sí misma se escribe como cualquier imagen:

```markdown
![](spinner.svg)
![w:400](chart.svg)
```

## Qué se lee {#read}

Las animaciones escritas en el propio archivo SVG: SMIL (`animate`, `animateTransform`, `animateMotion`, `animateColor`, `set`, con `begin`, `dur`, `repeatCount`, `fill`, `calcMode`, `keyTimes` y `keySplines`) y CSS (`@keyframes` con las propiedades `animation`, `transform-origin` y `transform-box`). Lo que no se entiende, un evento como `begin="click"` por ejemplo, se deja fuera y se avisa con una advertencia: nada se descarta en silencio.

## Qué muestra cada salida {#outputs}

- **HTML**: el navegador reproduce las animaciones por sí mismo.
- **PowerPoint** (`--pptx`) y **OpenDocument** (`--odp`): cada elemento animado es una forma propia, y sus animaciones empiezan con la diapositiva: movimiento, rotación y escala, opacidad, un cambio entre dos colores. El movimiento de un elemento se muestrea en los instantes propios de la animación, de modo que las aceleraciones, los fotogramas clave múltiples y una rotación alrededor de otro punto son exactos. Se conservan las repeticiones (también `indefinite`), los retrasos y las inversiones.
- **Fotogramas**: cuando una animación no puede expresarse así (un trazado que se deforma, un trazo que se dibuja, una inclinación), la imagen entera se renderiza en los instantes propios de la animación y se incrusta como un GIF animado, que PowerPoint e Impress reproducen. Está acotado: 12 fotogramas por segundo, 240 como máximo, 16 millones de píxeles entre todos.
- **PDF** (`--pdf`) e imágenes (`--images`): no pueden animarse. El dibujo se muestra en su estado inicial, sin JavaScript ni `RichMedia`, y una advertencia nombra las imágenes animadas.

Cada imagen animada recibe una línea de advertencia con el código `svg-animation`, que dice cómo se llevó: de forma nativa, o como fotogramas y por qué.

## Compatibilidad por salida {#support}

| Característica | HTML | PowerPoint | OpenDocument | PDF e imágenes |
|---|---|---|---|---|
| Traslación, movimiento a lo largo de un trazado | se reproduce | nativo | nativo | estado inicial |
| Rotación y escala, alrededor de cualquier punto | se reproduce | nativo | nativo | estado inicial |
| Radio de un círculo, tamaño de un rectángulo (sin contorno) | se reproduce | nativo | nativo | estado inicial |
| Opacidad de un elemento de una sola forma | se reproduce | nativo | nativo | estado inicial |
| Color de un relleno o de un contorno, de un color a otro | se reproduce | nativo | nativo | estado inicial |
| Repeticiones, retrasos, inversiones, mantener al final | se reproduce | nativo | nativo | estado inicial |
| Aceleraciones, fotogramas clave múltiples | se reproduce | nativo (muestreado) | nativo (muestreado) | estado inicial |
| Pasos, cambios discretos | se reproduce | fotogramas | fotogramas | estado inicial |
| Deformación de trazado, desplazamiento de guiones, atributos de posición y tamaño, degradados | se reproduce | fotogramas | fotogramas | estado inicial |
| Inclinación, reflejo, varios relojes en un elemento, elementos animados anidados | se reproduce | fotogramas | fotogramas | estado inicial |
| Eventos | se reproduce | no se lleva (advertencia) | no se lleva (advertencia) | estado inicial |

## Archivos locales {#local-files}

Una imagen local solo se lee con `--allow-local-files`, desde la carpeta de la presentación, como en marp-cli. Sin ella la imagen queda fuera de las exportaciones, y la compilación lo dice con `local-file-blocked`, nombrando el archivo, sus diapositivas y la opción. Un archivo que no existe se avisa como `local-file-missing`.

## Límites {#limits}

Un SVG es una entrada no fiable, por eso está acotado: 8 MiB de SVG, 50 000 elementos, 2 000 animaciones, 256 fotogramas clave por animación, 60 segundos de línea de tiempo y 240 fotogramas. Un dibujo que pasa un límite se rechaza con una advertencia, y la compilación continúa.

## Qué se ha comprobado {#checked}

LibreOffice Impress (macOS) abre los archivos de PowerPoint y de OpenDocument y conserva cada animación al guardarlos de nuevo, en los dos formatos. Keynote (macOS) importa las formas y el GIF animado; sus animaciones no se inspeccionaron. PowerPoint: aún sin verificar.
