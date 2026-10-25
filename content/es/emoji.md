---
title: "Emoji"
weight: 36
description: "Los códigos cortos y los emoji Unicode se convierten en imágenes Twemoji, sin red: en la página HTML, las imágenes, el PDF, PowerPoint y OpenDocument."
tags: [emoji]
---

{{< callout type="warning" >}}
Los emoji llegaron en v0.4.0; v0.3.0 no tenía ninguno.
{{< /callout >}}

Escriba un emoji como código corto o como el propio carácter:

```markdown
# Launch :rocket: 🎉

Shortcodes :smile: :+1: :heart_eyes: and Unicode 😄 👍🏽 👩‍💻 🇫🇷 1️⃣.
```

Cada uno se convierte en la imagen [Twemoji](https://github.com/jdecked/twemoji)
que dibuja Marp Core, con el marcado que escribe Marp Core:

```html
<img class="emoji" draggable="false" alt="😄" src="data:image/svg+xml;base64,…" data-marp-twemoji=""/>
```

## Sin red {#no-network}

El `src` de Marp Core es una URL de una CDN, así que una presentación mostrada
sin conexión, impresa o convertida a PDF en una máquina sin conexión muestra
imágenes rotas. go-slides coloca el propio SVG en un URI `data:`. Los gráficos
(4009, Twemoji 17.0.3, Emoji 17.0) van incrustados en el programa desde el
módulo en Go puro [go-icons/twemoji](https://github.com/go-icons/twemoji): 2 MB
comprimidos, que solo se leen cuando una presentación contiene algún emoji.

Las imágenes las dibuja por igual la página HTML, `--images`/`--image`, `--pdf`,
`--pptx` y `--odp`: go-webengine rasteriza el SVG. Los 4009 gráficos se analizan
y se dibujan con su rasterizador de SVG; una prueba lo comprueba. Desde v0.5.0
ese rasterizador es [go-gfx/oksvg](https://github.com/go-gfx/oksvg) v0.2.0:
`clipPath`, `fill-rule: evenodd` y los comandos de arco se dibujan ahora como los
dibuja Chrome, lo que cambió un Twemoji y otros 211, y los logotipos SVG ya no
salen deformados.

La regla que ajusta el tamaño de un emoji al del texto (`img[data-marp-twemoji]`,
un cuadrado de un `em`, alineado con la línea base) se añade al tema igual que
lo hace Marp Core, de modo que un tema propio también la recibe.

## Qué se convierte {#what-is-converted}

Las reglas son las de Marp Core; una prueba las compara con el HTML que escribe
Marp Core 4.4.0 para una presentación con todos los casos.

| | |
|---|---|
| `:smile:`, `:+1:`, `:heart_eyes:` | los 1903 nombres de markdown-it-emoji, en texto, títulos, enlaces, tablas, `header` y `footer`; no en código, URL ni en el texto alternativo de una imagen |
| `😄`, `👩‍💻`, `🇫🇷`, `1️⃣`, `👍🏽` | emoji Unicode, en el texto **y en el código**: Marp Core reescribe el HTML del código ya renderizado |
| `✔`, `©`, `™`, `▶` | se quedan como texto: los caracteres cuya presentación por defecto es de texto solo son emoji con U+FE0F (`✔️`); `❤` es la excepción |
| `☺︎` | se queda como texto: un U+FE0E tras un emoji pide el glifo plano |
| `12:30:45`, `:not_a_name:` | se quedan tal como están escritos |

El `id` de un título y el índice usan el emoji, no el nombre del código corto:
`# Launch :rocket:` tiene el id `launch-%F0%9F%9A%80`, como en Marp.

Donde go-slides difiere de Marp Core, convierte más: `✌️ ☝️ ✍️ 🏋️` y algunos
otros caracteres de presentación textual por defecto que el emoji-regex de Marp
Core deja como texto, además de las pocas secuencias que Twemoji dibuja y que
Unicode no lista.

## Desactivarlo {#turning-it-off}

`--no-emoji` (en código, `slides.Config{NoEmoji: true}`) deja los códigos cortos
y los emoji Unicode como el texto con el que se escribieron, que es el
`emoji: false` de Marp Core.

## Atribución {#attribution}

Los gráficos son Twemoji, Copyright 2020 Twitter, Inc y otros colaboradores,
con licencia [CC-BY 4.0](https://creativecommons.org/licenses/by/4.0/) y
mantenidos por el proyecto [jdecked/twemoji](https://github.com/jdecked/twemoji).
Una presentación que los muestre debe acreditarlos, por ejemplo en un pie de
página o en una diapositiva final. La tabla de códigos cortos es la de
markdown-it-emoji (MIT).
