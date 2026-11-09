---
title: "Grabaciones de terminal"
weight: 38
description: "Una grabación de asciinema en una presentación: se reproduce en un pequeño reproductor de la página HTML y se dibuja como texto vectorial en el PDF, PowerPoint y OpenDocument."
tags: [asciinema, terminal]
---

{{< callout type="info" >}}
Las grabaciones de terminal están en v0.18.0 y posteriores.
{{< /callout >}}

Una grabación de [asciinema](https://docs.asciinema.org/) ([asciicast v2](https://docs.asciinema.org/manual/asciicast/v2/) o [v3](https://docs.asciinema.org/manual/asciicast/v3/)) es una imagen cuyo archivo termina en `.cast`:

```markdown
![cast](demo.cast)
![The failing build](demo.cast#t=12.5)
```

## Sintaxis {#syntax}

El archivo se lee de la carpeta de la presentación: se rechaza una URL, una ruta absoluta o una ruta que salga de la carpeta. El texto alternativo nombra la grabación para los lectores de pantalla; `cast` solo significa el título propio de la grabación. `#t=12.5`, o `#t=1:23` en minutos y segundos, elige el momento que muestran las salidas que no pueden moverse; sin él, es el final de la grabación.

## Qué muestra cada salida {#outputs}

- **HTML** (la página y la vista previa del playground): la grabación se reproduce en un pequeño reproductor de go-slides incrustado en la página: reproducir y pausar, un botón de velocidad y una barra de posición. No carga ningún script de ningún sitio y no usa la red.
- **PDF, PowerPoint y OpenDocument** (`--pdf`, `--pptx`, `--odp`): el terminal en el momento que indica `#t=`, como texto vectorial en la fuente monoespaciada del tema, con los colores de 16, 256 y directos, negrita, cursiva, subrayado y texto invertido. El texto se puede seleccionar y buscar. Nunca es una imagen. La negrita y la cursiva se dibujan cuando la fuente monoespaciada del tema las tiene (Roboto Mono en el tema gaia).

## El reproductor {#player}

Se reproduce una vez, la primera vez que la grabación está en pantalla. Cuando el navegador pide menos movimiento (`prefers-reduced-motion`), empieza en pausa, mostrando el momento que indica `#t=`, y se reproduce cuando pulsa reproducir. Sus botones y teclas no cambian de diapositiva.

## Límites {#limits}

Una grabación es una entrada no confiable, así que está acotada: el archivo (16 MiB), los eventos (500 000), la duración (3 horas), el terminal (400 por 200 celdas) y el trabajo que puede hacer el emulador. Una grabación que supera un límite, o que no es un archivo asciicast bien formado, se rechaza: la compilación continúa, la diapositiva muestra una nota en su lugar y el problema es una advertencia con un código: `cast.syntax`, `cast.limit`, `cast.version` o `cast.path`.

## En el playground {#playground}

Añada un archivo `.cast` al proyecto desde el panel Archivos, soltándolo o con el selector de archivos. Va a `media/` y se inserta como una imagen. Es texto que la presentación lee, así que cuenta con los demás archivos de texto del proyecto, limitados a 2 MiB. Pruébelo en el [playground](https://go-slides.github.io/playground/).
