---
title: "Vídeo"
weight: 39
description: "Un vídeo en una presentación: reproducido por el navegador en la página HTML, un fotograma de póster con enlace en el PDF, incrustado en PowerPoint y OpenDocument."
tags: [video]
---

Un vídeo es una imagen cuyo archivo termina en `.mp4`, `.m4v` o `.webm`:

```markdown
![video](clip.mp4)
![The demo](clip.mp4#t=12.5)
![demo loop muted](clip.mp4)
```

## Sintaxis {#syntax}

El archivo se lee de la carpeta de la presentación, o de una URL `http` o `https`. Se rechaza una ruta absoluta, una ruta que sale de la carpeta o cualquier otro esquema. `#t=12.5`, o `#t=1:23` en minutos y segundos, elige el fotograma de póster: la imagen que se muestra antes de la reproducción y en las salidas que no pueden reproducir el vídeo. También se acepta la forma de fragmento multimedia `#t=10,20`, y se usa su inicio. Sin él, el póster es el primer fotograma.

Las palabras clave del texto alternativo, escritas como las palabras clave de imagen de Marp pero propias de go-slides, fijan cómo se reproduce el vídeo: `autoplay`, `loop` y `muted`. `autoplay` solo se respeta junto con `muted`; si no, se ignora, con una advertencia. Las demás palabras del texto alternativo son la etiqueta accesible.

## Qué muestra cada salida {#outputs}

- **HTML** (la página y la vista previa del playground): un elemento `<video controls playsinline preload="metadata" poster=...>`. El póster es el fotograma decodificado, incrustado como URI de datos JPEG. Un vídeo nunca suena por sí solo. Cuando el navegador pide menos movimiento (`prefers-reduced-motion`), se quitan `autoplay` y `loop` y el vídeo se pausa. Un vídeo que se reproduce solo (silenciado) suena únicamente mientras al menos la mitad es visible, y se pausa cuando su diapositiva no se muestra.
- **PDF** (`--pdf`): el fotograma de póster como imagen, con un enlace al archivo de vídeo: la ruta relativa tal como está escrita en la presentación, o la URL si la fuente es `http` o `https`. La reproducción solo ocurre al abrir el enlace. El PDF no lleva anotación `RichMedia`: solo se reproduciría en Acrobat, y deliberadamente no se usa.
- **PowerPoint** (`--pptx`): el vídeo se incrusta como una parte multimedia (`ppt/media/slidesvideoN.mp4|webm`) con su póster, como `a:videoFile` más la extensión `p14:media`. Keynote: lo importa como película y lo reproduce (verificado). LibreOffice Impress: importa los archivos PowerPoint y OpenDocument como vídeo y lo decodifica (verificado). PowerPoint: aún sin verificar. Las palabras clave `autoplay`, `loop` y `muted` no pasan a PowerPoint ni a OpenDocument.
- **OpenDocument** (`--odp`): el vídeo se incrusta bajo `Video/` como un `draw:plugin` (ODF 1.3 parte 3, 10.4.8), antes del `draw:image` del póster, y se lista en el manifiesto. LibreOffice Impress lo importa como vídeo y lo decodifica (verificado); la reproducción en la presentación depende de su motor multimedia y del códec.
- **Fuentes URL**: la página HTML transmite el vídeo desde la URL. Las exportaciones muestran una tarjeta neutra con enlace a la URL. No se descarga nada al construir, y nada se incrusta en PowerPoint ni en OpenDocument.

## Contenedores y pósteres {#codecs}

Se leen archivos MP4 (ISO BMFF) y WebM (Matroska), en Go; nunca se ejecuta ffmpeg. El póster solo se decodifica para H.264 (`avc1` o `avc3`), el único decodificador de vídeo de go-avkit. Los demás códecs (H.265, VP8, VP9, AV1) se reproducen igualmente en HTML y se incrustan en PowerPoint y OpenDocument, pero su póster es una tarjeta oscura neutra con un triángulo de reproducción, y se imprime la advertencia `video.codec`.

## Límites {#limits}

Un vídeo es una entrada no fiable, así que está acotado: el archivo (64 MiB), las cajas o elementos (262.144) y su anidamiento (16 niveles), el fotograma (4096 por 4096) y la decodificación del póster (como máximo 600 muestras desde el fotograma clave anterior, 8 MiB por muestra). Un póster demasiado lejos de un fotograma clave recibe la tarjeta y la advertencia `video.limit`.

Un vídeo rechazado se sustituye en la diapositiva por una nota, `[video not shown: ...]`; la construcción continúa, y el problema es una advertencia con un código:

- `video.syntax`: no es un archivo MP4 o WebM, no hay pista de vídeo, o un tiempo no válido;
- `video.limit`: se supera uno de los límites anteriores;
- `video.codec`: el póster no se puede decodificar para este códec;
- `video.path`: una ruta absoluta, una ruta que sale de la carpeta de la presentación, o un esquema distinto de `http` y `https`;
- `file.missing`: el archivo no existe.

## En el playground {#playground}

Una vez publicada, se podrá añadir un vídeo al proyecto como los demás archivos. Sus bytes se quedan en el navegador, como URL blob, y nunca se suben. Pruébelo en el [playground](https://go-slides.github.io/playground/).

La compatibilidad con vídeo se carga bajo demanda. El lector que decodifica el fotograma del póster (unos 3 MiB comprimidos) es un módulo aparte, que se descarga la primera vez que una presentación necesita un póster. Una presentación sin vídeo nunca lo descarga, y la primera carga del playground es así más ligera. Se ejecuta en un Web Worker, de modo que escribir nunca se bloquea, y lee un archivo con los mismos límites que `slides` (64 MiB, cajas y tamaño de fotograma) y los mismos errores `video.limit` y `video.syntax`.
