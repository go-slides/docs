---
title: "Línea de comandos"
weight: 50
description: "Todos los comandos y opciones de slides, y el archivo de configuración slides.yaml."
tags: [cli]
---

{{< callout type="warning" >}}
Esta página describe v0.13.0.
{{< /callout >}}

```text
slides [deck.md] [flags]
slides [command] [flags] deck.md
```

## Comandos {#commands}

| Comando | Qué hace |
| --- | --- |
| `build` | escribe la presentación como HTML, PDF o imágenes (véase más abajo). **Es el comando por defecto**: `slides deck.md` equivale a `slides build deck.md` |
| `expand` | imprime la presentación como Markdown de Marp plano: plantillas expandidas, kinds sustituidos por las directivas que representan |
| `notes` | imprime las notas del orador |
| `kinds` | lista los kinds que la presentación puede usar y lo que fijan; `--json` (en v0.6.0) los escribe como `[{"name", "extends", "directives": [{"key", "value"}]}]`, con `extends` ya resuelto |
| `themes` | lista los temas |
| `version` | imprime la versión (también lo hace `slides --version`) |
| `completion` | imprime un script de autocompletado para el shell: `bash`, `zsh`, `fish` o `powershell` |

`slides <command> --help` imprime las opciones de un comando. Las opciones
pueden ir antes o después de la presentación: `slides deck.md -o out.html` y
`slides -o out.html
deck.md` son equivalentes. Las opciones largas llevan dos guiones (`--html`);
`-o` es la única corta. `build` escribe `deck.html` junto a `deck.md`, salvo que
`-o` indique otra cosa.

## Opciones {#flags}

| Opción | Comandos | Qué hace |
| --- | --- | --- |
| `-o`, `--output file` | `build`, `expand`, `notes` | archivo de salida; `-` para la salida estándar |
| `-I`, `--input-dir dir` | `build` | convierte los archivos Markdown de `dir`, de forma recursiva (en v0.6.0); véase más abajo |
| `-P`, `--parallel n` | `build` | presentaciones construidas a la vez (en v0.6.0): 5 por defecto, como máximo el número de núcleos; con 1 se construyen una tras otra |
| `-w`, `--watch` | `build` | vuelve a construir cuando cambia una presentación, un tema, un archivo de kinds o de datos, o la configuración, y recarga las páginas abiertas (en v0.6.0); véase más abajo |
| `-s`, `--server dir` | `build` | sirve un directorio y convierte sus presentaciones bajo demanda (en v0.6.0); véase más abajo |
| `-p`, `--preview` | `build` | abre el resultado en el navegador por defecto (en v0.6.0): las páginas HTML con `--watch` (implícito), el servidor con `--server` |
| `--host addr` | `build` | la dirección en la que escucha `--server`, `127.0.0.1` por defecto (en v0.6.0) |
| `--allowed-host name` | `build` | el nombre de host público de un proxy inverso al que responde `--server`, repetible (después de v0.13.0); consulte [Autoalojamiento]({{< relref "self-hosting" >}}) |
| `--template bespoke\|bare` | `build` | la plantilla de la página HTML (en v0.6.0): `bespoke`, la de por defecto, con el script de presentación, el controlador en pantalla y las transiciones, o `bare`, con las diapositivas apiladas (`--bare`) |
| `--bespoke.osc`, `--bespoke.progress`, `--bespoke.transition` | `build` | las opciones de la plantilla bespoke, con los valores por defecto de marp-cli (en v0.6.0): controlador activado, **barra de progreso desactivada** (antes se dibujaba siempre), transiciones activadas; `--bespoke.osc=false` desactiva una |
| `-c`, `--config-file file`, `--no-config-file` | `build` | el archivo de configuración de marp-cli, JSON o YAML (en v0.6.0); véase más abajo. `--config` es el archivo propio de `slides` |
| `--browser`, `--browser-path`, `--browser-protocol`, `--browser-timeout` | `build` | se aceptan y se ignoran con un aviso (en v0.6.0): no hay navegador |
| `--pptx-editable`, `--no-parallel` | `build` | las grafías de marp-cli para `--pptx` y `-P 1` (en v0.6.0) |
| `--config file` | todos | archivo de configuración; por defecto `slides.yaml` junto a la presentación, si existe |
| `--no-config` | todos | no lee ningún archivo de configuración |
| `--theme-set path` | todos | un archivo CSS de tema, o un directorio de ellos (repetible) |
| `--kinds file` | todos | un archivo YAML de kinds (repetible) |
| `--data file` | todos | un archivo YAML o JSON con datos para las plantillas (repetible) |
| `--set key=value` | todos | un valor para las plantillas (repetible) |
| `--no-template` | todos | no expande las plantillas de Go |
| `--strict-yaml` | todos | analiza las directivas como YAML estricto, como hace Marpit (Marp Core las analiza con más tolerancia) |
| `--no-emoji` | todos | deja los `:shortcodes:` y los emoji Unicode como texto en lugar de dibujarlos como imágenes (desde v0.4.0: [Emoji]({{< relref "emoji" >}})) |
| `--html` | `build` | permite todo el HTML sin procesar; por defecto, la lista de permitidos de Marp Core |
| `--lang code` | `build` | el `lang` de la página cuando la presentación no define ninguno |
| `--bare` | `build` | apila las diapositivas, sin el script de presentación |
| `--pdf` | `build` | escribe un PDF, una página por diapositiva, con índice: texto y formas en vectores (también se elige con `-o file.pdf`) |
| `--pdf-raster` | `build` | escribe el PDF con cada página como una imagen de la diapositiva, como hacía v0.3.0 |
| `--notes` | `build` | escribe las notas del orador en un archivo de texto, con las diapositivas separadas por `---` como hace marp-cli (en v0.6.0; también se elige con `-o file.txt`). El comando `notes` las imprime con otra disposición |
| `--pdf-notes` | `build` | añade las notas del orador de cada diapositiva a su página del PDF como anotación de nota, como hace marp-cli (en v0.6.0). Pide un PDF cuando nada más indica qué escribir |
| `--pdf-outlines`, `--pdf-outlines.pages`, `--pdf-outlines.headings` | `build` | los marcadores del PDF (en v0.6.0): «Page N» por diapositiva y los títulos anidados debajo. Activados por defecto aquí, a diferencia de marp-cli; `=false` desactiva uno. Véase más abajo |
| `--images png\|jpeg` | `build` | escribe una imagen por diapositiva: `deck.001.png`, … |
| `--image png\|jpeg` | `build` | escribe la primera diapositiva como imagen (también se elige con `-o file.png` o `-o file.jpg`) |
| `--pptx` | `build` | escribe un archivo de PowerPoint con sus notas del orador (también se elige con `-o file.pptx`): editable desde v0.5.0, una imagen por diapositiva en v0.4.0 |
| `--pptx-raster` | `build` | escribe un archivo de PowerPoint con una imagen a diapositiva completa por diapositiva, como hace marp-cli (desde v0.5.0); `--image-scale` vale 2 por defecto |
| `--odp` | `build` | escribe una presentación OpenDocument con sus notas del orador (también se elige con `-o file.odp`): editable desde v0.5.0 |
| `--odp-raster` | `build` | escribe una presentación OpenDocument con una imagen por diapositiva (desde v0.5.0) |
| `--jpeg-quality n` | `build` | calidad JPEG, de 1 a 100 (85 por defecto) |
| `--image-scale n` | `build` | píxeles del dispositivo por píxel CSS para las imágenes y el PDF (1 por defecto); con 2 se obtienen 2560×1440 en 16:9, con el texto y los vectores dibujados a esa densidad, no ampliados |
| `--allow-local-files` | `build` | permite que las imágenes y el PDF lean archivos locales, solo desde el directorio de la presentación; se rechaza por defecto, como en marp-cli |
| `--theme name\|file.css` | `build` | sustituye el tema de la presentación (en v0.6.0): un nombre de tema o un archivo CSS, que pasa a ser el tema sea cual sea el `@theme` que declare; un directorio se rechaza, use `--theme-set` |
| `--title`, `--description`, `--author`, `--keywords`, `--url`, `--og-image` | `build` | los metadatos de la página (en v0.6.0), por encima de las directivas `title`, `description`, `author`, `keywords`, `url` e `image` de la presentación; véase más abajo |

## Varias presentaciones y un directorio {#several-decks-and-a-directory}

En v0.6.0, `build` admite varias presentaciones, un directorio (las
presentaciones que contiene, buscadas de forma recursiva) o un patrón, y
construye hasta `--parallel` de ellas a la vez. `-` es la entrada estándar; a
diferencia de marp-cli, no se lee una tubería salvo que `-` lo indique.

Un patrón se interpreta como lo hace el glob de marp-cli (`**` y `{a,b}` en
v0.7.0; escríbalo entre comillas para que el shell se lo deje a `slides`): `*`,
`?`, `[a-z]` y `{md,txt}` dentro de un nombre, `**` para cualquier número de
directorios, de modo que `'docs/**/*.md'` toma las presentaciones de todo el
árbol. Un comodín no alcanza los nombres que empiezan por un punto, no se busca
nada en `node_modules`, se distinguen mayúsculas y minúsculas (`*.md` no
encuentra `UPPER.MD`), y se convierte todo lo que el patrón encuentre, sea cual
sea su extensión. Un patrón que no encuentra nada es un error.

- **`-I dir`** busca los archivos Markdown (`.md`, `.mdown`, `.markdown`,
  `.markdn`, en minúsculas) dentro de `dir`, como hace marp-cli: no se mira en
  `node_modules` ni en los directorios que empiezan por un punto. Con `-o out`,
  cada presentación se escribe en su ruta bajo `dir`, replicada bajo `out`:
  `dir/sub/b.md` da `out/sub/b.html` (`b.pdf`, `b.001.png`, `b.txt`, … según las
  opciones). Sin `-o`, las salidas van junto a sus presentaciones. Indicar
  presentaciones además de `-I` es un error, y también lo es `-o file` con varias
  presentaciones (`-o -` está permitido).
- **`-P n`** construye `n` presentaciones a la vez. El renderizado de una
  presentación ya usa los núcleos, así que `n` está limitado por su número; con
  `-o -` las presentaciones se construyen una tras otra. Los mensajes de cada
  presentación se mantienen juntos, y una presentación que falla no detiene a las
  demás: el estado de salida es 1 y el mensaje lista los fallos.

## Transiciones y la página de presentación {#transitions-and-the-presenting-page}

En v0.6.0 funciona la directiva `transition` de marp-cli: `<!-- transition:
fade 1s -->`, un nombre y una duración opcional, para la diapositiva y las
siguientes, o `<!-- _transition: cover -->` para una sola. Los nombres
incorporados son los de marp-cli: `clockwise`, `counterclockwise`, `cover`,
`coverflow`, `cube`, `cylinder`, `diamond`, `drop`, `explode`, `fade`,
`fade-out`, `fall`, `flip`, `glow`, `implode`, `in-out`, `iris-in`, `iris-out`,
`melt`, `overlap`, `pivot`, `pull`, `push`, `reveal`, `rotate`, `slide`, `star`,
`swap`, `swipe`, `swoosh`, `wipe`, `wiper` y `zoom`; `none` la desactiva. La
página la reproduce con la **View Transition API** del navegador cuando el
espectador sale de la diapositiva que la tiene, hacia delante, y cuando vuelve a
ella; los navegadores sin esa API cambian de diapositiva como antes, y un
espectador que pide movimiento reducido obtiene un simple fundido. Un tema puede
aportar sus propios fotogramas clave, con el nombre `marp-transition-NAME` (o
`marp-outgoing-` y `marp-incoming-`, y `-backward`), que se usan antes que los
incorporados.

La página incluye el **controlador en pantalla** (anterior, página, siguiente,
pantalla completa), que aparece cuando se mueve el puntero o una tecla y se
oculta pasado un momento (`--bespoke.osc=false` lo elimina) y, con
`--bespoke.progress`, una barra de progreso. **La barra de progreso está ahora
desactivada por defecto**, como en marp-cli; antes se dibujaba siempre.

### Vista del presentador {#presenter-view}

En v0.7.0 la página incluye la vista del presentador de marp-cli. Pulse `p`, o
el último botón del controlador, y se abre una ventana con la diapositiva
actual, el paso siguiente, sus notas del orador, la página, la hora y un
cronómetro (haga clic en el cronómetro para reiniciarlo; `+` y `-` cambian el
tamaño de las notas; la barra entre las diapositivas y las notas se puede
arrastrar). Las dos ventanas se siguen en ambos sentidos: las teclas de flecha,
un clic en el paso siguiente o en Anterior/Siguiente en cualquiera de las dos
mueve la otra, un fragmento cada vez.

Funciona como la de marp-cli, según se lee en sus fuentes de la versión 4.5.1: la
ventana es `?view=presenter&sync=KEY` del mismo archivo, el paso siguiente es la
misma página como `?view=next` dentro de un marco, y las ventanas comparten su
posición mediante `localStorage` (`bespoke-marp-sync-KEY`), de modo que no
necesita servidor ni red, y no funciona en un navegador sin `localStorage` (el
botón lo indica). Las notas son texto: `&`, `<` y `>` se escapan, por lo que una
nota nunca puede ser marcado. Están en la página, ocultas, para que la vista del
presentador las tome, lo que las mantiene fuera de lo que un lector de pantalla
lee al público. Las páginas `--bare` no tienen nada de esto. La vista añade
unos 10 KB a una página.

### Vista general {#overview-view}

En v0.8.0, pulse `o` (o `Escape`: marp-cli alterna con ambas), el botón de
cuadrícula del controlador o el contador de páginas de la vista del presentador,
y todas las diapositivas se muestran enteras en una cuadrícula sobre la página.
Las teclas de flecha eligen una diapositiva (arriba y abajo, por filas), `Home`
y `End` van a los extremos, y las demás ventanas la siguen; `Enter`, `Space` o
un clic eligen una diapositiva y, sobre la página, cierran la vista general.
Sobre la vista del presentador, elegir deja la vista abierta (el presentador
escoge la diapositiva que se muestra y sigue mirando). El botón «Close
overview», `o` y `Escape` la cierran. Es la de marp-cli: la misma página otra
vez dentro de un marco en `?view=overview&sync=KEY`, que comparte la posición
mediante `localStorage`, sin red. Las diapositivas de la cuadrícula admiten el
foco y llevan la etiqueta «Slide N», y la vista general cerrada es inerte, de
modo que quien use el teclado o un lector de pantalla nunca cae en ella. Las dos
vistas juntas añaden unos 15 KB a una página.

### Archivos de configuración {#configuration-files}

marp-cli lee sus opciones de un archivo en el directorio de trabajo, y `slides`
también, cuando el archivo es de datos: `.marprc` (YAML; JSON también es YAML),
`.marprc.json`, `.marprc.yaml`, `.marprc.yml`, `.config/marprc*`, o la clave
`marp` de `package.json`; `--config-file` (`-c`) nombra cualquier archivo JSON o
YAML, `marp.config.json` para uno de ellos. Solo se busca en el directorio de
trabajo, no en sus padres, como hace la búsqueda de marp-cli. Un archivo de
JavaScript o TypeScript (`marp.config.js` ...) es un programa, que `slides` no
ejecuta: lo avisa y continúa. Las opciones son las de la documentación de
marp-cli que tienen sentido aquí: `title`, `description`, `author`, `keywords`,
`url`, `ogImage`, `lang`, `theme`, `themeSet`, `html`, `allowLocalFiles`,
`template`, `bespoke`, `inputDir`, `output`, `parallel`, `pdf`, `pptx`, `notes`,
`pdfNotes`, `pdfOutlines`, `images`, `image`, `imageScale`, `jpegQuality`,
`watch`, `server`, `preview`; las rutas son relativas al archivo. La línea de
comandos prevalece. `engine`, `options` y `browser*` necesitan el motor de Marp
y un navegador, y se ignoran con un aviso, igual que las opciones `--browser*`:
los scripts escritos para marp-cli siguen funcionando. (`slides.yaml` y
`slides.hcl`, con los kinds y los datos de las plantillas, son aparte y se leen
como antes.)

## Observar, servir y previsualizar {#watch-serve-and-preview}

En v0.6.0, con las rutas y el protocolo de recarga en vivo de marp-cli (leídos
de sus fuentes de la versión 4.5.1 y comparados con un `marp -s` en ejecución).

**`slides --server dir`** (`-s`) sirve el directorio en
`http://127.0.0.1:8080/` (la variable de entorno `PORT` cambia el puerto; un
puerto ya en uso es un error que lo dice):

| Petición | Respuesta |
| --- | --- |
| `/deck.md` | la presentación como HTML, convertida bajo demanda |
| `/deck.md?pdf`, `?pptx`, `?png`, `?jpg` (o `?jpeg`), `?notes` (o `?txt`) | la presentación convertida como pide la consulta; el PPTX es un adjunto |
| un directorio | su `index.md` o `PITCHME.md` como página; si no, un índice con un enlace al PDF y otro al PPTX de cada presentación |
| cualquier otro archivo | el archivo |
| una presentación que no se convierte | 503 con el error |

Una página abre un WebSocket en `/.__marp-cli-watch-notifier__/<id>`, donde el id
es el SHA-256 de la ruta de la presentación; el servidor envía `ready` y luego
`reload` cuando cambia la presentación o cualquier cosa bajo el directorio. Un
directorio solicitado sin la barra final se redirige a la versión con barra
(marp-cli sirve su índice ahí, lo que rompe los enlaces relativos de la página).

**`--watch`** (`-w`) construye las presentaciones y luego vuelve a construir la
que cambia y, cuando cambia un archivo de `--theme-set`, un archivo de kinds o de
datos o la configuración, todas. Las páginas que escribe abren un socket en el
loopback (puerto 37717, o el siguiente libre) y se recargan cuando se vuelve a
construir su presentación. Una presentación que no se construye se notifica y
se sigue observando. **`--preview`** (`-p`) abre el navegador por defecto en el
servidor, o en las páginas HTML recién construidas; no hay ventana propia, ya que
`slides` no tiene un navegador que controlar.

El servidor es una puerta de entrada a la máquina, así que es más estricto que el
de marp-cli, que escucha en todas las interfaces y no comprueba nada:

- escucha en **127.0.0.1**; `--host` lo cambia, con un aviso, y una dirección
  comodín ya no permite comprobar la cabecera `Host`;
- la cabecera **`Host`** debe ser un nombre de loopback (`localhost`,
  `127.0.0.1`, `[::1]`, `*.localhost`) o el host en el que escucha, lo que impide
  el DNS rebinding;
- el **`Origin`** de un WebSocket debe ser el del propio servidor; los sockets del
  modo watch aceptan además el origen `null` de una página abierta desde un
  archivo; se rechaza una petición de otro sitio que no sea una navegación (un
  script, una imagen, un `fetch`);
- una ruta se limpia, se resuelven sus enlaces simbólicos y el resultado debe
  quedar dentro del directorio: `..` y un enlace que salga de él se rechazan, y
  también `?pdf` de una presentación fuera de él; los nombres que empiezan por un
  punto (`.git`, `.env`) no se sirven;
- solo `GET` y `HEAD`; pocas conversiones a la vez; como máximo 4096 directorios
  observados, y ninguno en `node_modules` ni bajo un punto.

Detrás de un proxy inverso, `--allowed-host name` (después de v0.13.0,
repetible) añade el nombre público del proxy a los `Host` y `Origin` aceptados;
las comprobaciones siguen activas y el servidor sigue en el bucle local.
[Autoalojamiento]({{< relref "self-hosting" >}}) lo coloca detrás de Caddy, con
HTTPS y un inicio de sesión.

## Metadatos de la página {#page-metadata}

La línea de comandos prevalece sobre la presentación, como en marp-cli, donde las
opciones se añaden a la presentación como un último comentario de directivas.
Reglas leídas de marp-cli 4.5.1 y fijadas con pruebas contra su salida (en
v0.6.0):

- Una presentación que no define `title` toma el texto de su primer título del
  nivel más bajo, tal como está escrito (`# **Bold** title` da `**Bold** title`).
  Sin ningún título no hay `<title>`.
- Las `keywords` se separan por comas, se recortan y se deduplican; una lista del
  front matter se toma tal cual. Como marp-cli, la página incluye la etiqueta
  `keywords` a partir de la segunda palabra clave, no para una sola.
- Una `url` no válida (no absoluta) se ignora con un aviso y se conserva el valor
  anterior; una válida escribe además `<link rel="canonical">` y `og:url`.
- `author` y `description` vienen con sus gemelos de Open Graph; `og:image:alt` y
  la tarjeta grande de Twitter necesitan un título y una imagen.

## Imágenes, PDF, PPTX y ODP {#images-pdf-pptx-and-odp}

`-o` con una extensión `.pdf`, `.pptx`, `.odp`, `.png` o `.jpg` elige el formato,
como hace marp-cli. Los directorios de salida se crean. Un PDF tiene una página
por diapositiva, del tamaño de la diapositiva, y un índice de las diapositivas y
sus títulos.

### Notas e índice del PDF {#pdf-notes-and-outline}

En v0.6.0 el PDF admite las dos opciones de marp-cli, y se comparan con los PDF
que escribe marp-cli 4.5.1 para la misma presentación:

- `--pdf-notes` coloca las notas del orador de cada diapositiva en su página como
  anotación de nota: un icono amarillo de nota abajo a la izquierda cuyo
  contenido son las notas (varios comentarios unidos por una línea en blanco) y
  cuyo título es el `author` de la presentación. `--pdf-notes` solo pide un PDF,
  como en marp-cli.
- El índice contiene un marcador en negrita «Page N» por diapositiva y, debajo,
  los títulos de la diapositiva anidados por nivel (`--pdf-outlines.pages` y
  `--pdf-outlines.headings` eligen las dos partes). **El índice está activado por
  defecto aquí y desactivado en marp-cli.** `slides` siempre ha escrito uno;
  cuesta unos cientos de bytes y hace navegable una presentación larga, así que
  se mantiene el valor por defecto. `--pdf-outlines=false`, o el
  `--no-pdf-outlines` de marp-cli, lo desactiva, y `=false` en `.pages` o
  `.headings` suprime esa parte. El marcador de un título salta a su diapositiva;
  marp-cli además se desplaza hasta el título.

El renderizado lo hace [go-webengine](https://github.com/go-webengine/engine) y
el PDF lo escribe [go-pdfkit](https://github.com/go-pdfkit/pdfkit); las fuentes
que piden los temas incorporados (Lato, Roboto Mono) provienen de
[go-opentype/fonts](https://github.com/go-opentype/fonts) y van incrustadas, de
modo que ningún renderizado toca la red. No hay Chrome.

`--image-scale 2` renderiza a dos píxeles del dispositivo por píxel CSS. Las
imágenes locales se rechazan salvo que se indique `--allow-local-files`, e
incluso entonces solo se leen los archivos bajo el directorio de la
presentación: una presentación ajena no puede alcanzar `~/.ssh`. Los fondos
múltiples y divididos (`![bg left:40%]`) se renderizan como lo hacen las capas de
Marp.

`--pptx` escribe la presentación con el `WritePresentation` de
[go-ooxml](https://github.com/go-ooxml), con las notas del orador y los metadatos
de la presentación. Desde v0.5.0 las diapositivas son editables (véase más
abajo); `--pptx-raster` escribe una imagen a diapositiva completa por diapositiva,
como hace marp-cli, con una escala que vale 2 por defecto y que `--image-scale`
sustituye. LibreOffice abre el resultado. Solo se probó LibreOffice; PowerPoint y
Keynote no.

`--odp` escribe una presentación OpenDocument, mediante el `WritePresentation` de
go-odf, con sus notas del orador; `--odp-raster` escribe una imagen por
diapositiva.

### PPTX y ODP editables {#editable-pptx-and-odp}

Desde v0.5.0 las diapositivas de un PPTX o de un ODP son **editables**: el PPTX de
marp-cli es una imagen por diapositiva, y su `--pptx-editable` experimental
convierte un PDF mediante LibreOffice. Aquí cada diapositiva se dibuja a partir de
su disposición, la misma que dibujan las exportaciones a PNG y PDF:

- cada línea de texto es un cuadro de texto en su posición calculada, con las
  fuentes con las que se maquetó la diapositiva, **incrustadas en el archivo**,
  de modo que se ve igual en una máquina sin Lato o Inter (comprobado con
  LibreOffice 25.8), con negrita, cursiva, colores, espaciado entre letras y
  enlaces en los que se puede hacer clic;
- los fondos, los bordes y los cuadros redondeados son formas; los degradados son
  rellenos degradados en el PPTX e imágenes en el ODP (un degradado de ODF tiene
  dos colores);
- las imágenes se almacenan una sola vez, por muchas diapositivas que las
  muestren;
- lo que una diapositiva no puede expresar (un elemento bajo `opacity`, un
  `filter`, una sombra difuminada, una fórmula) es una imagen de ese elemento, con
  su texto conservado encima.

Al renderizarla de nuevo con LibreOffice, una diapositiva queda entre el 0,04 y el
2 % de los píxeles de la exportación a PNG. `--pptx-raster` y `--odp-raster`
mantienen una imagen por diapositiva.

### PDF vectorial {#vector-pdf}

Desde v0.4.0 el PDF se dibuja en vectores: go-webengine pinta la diapositiva ya
maquetada en el PDF en lugar de en píxeles. El texto es texto, con las fuentes con
las que se maquetó la diapositiva, incrustadas en subconjunto con un mapa
`/ToUnicode`: se puede seleccionar, buscar y copiar. Los fondos, los bordes y los
degradados son trazados y sombreados, una imagen que aparece en muchas
diapositivas se incrusta una sola vez y los enlaces admiten clic. La presentación
de prueba de 100 diapositivas pesa 758 KiB, frente a 837 KiB de marp-cli y 10 MiB
del PDF de imágenes de v0.3.0, y se escribe en 0,3 s.

Un elemento sin forma vectorial aquí (bajo `opacity` < 1, con un `filter` o una
rotación, un `box-shadow` difuminado, un degradado cuya opacidad se atenúa) se
dibuja como un mapa de bits solo de ese elemento, con su texto superpuesto de
forma invisible para que siga siendo seleccionable. Las fórmulas siguen siendo
mapas de bits. `--pdf-raster` escribe una imagen por página, como hacía v0.3.0.

### PDF accesible y etiquetado {#accessible-tagged-pdf}

Desde v0.12.0 el PDF vectorial está etiquetado. Lleva su estructura lógica, que
un lector de pantalla lee en orden, como lee una página web:

| En la presentación | En el PDF |
| --- | --- |
| una diapositiva | una sección (`Sect`) |
| `#` … `######` | `H1` … `H6` |
| un párrafo, una lista, una tabla | `P`; `L` con la etiqueta y el cuerpo de cada elemento; `Table`, con celdas de encabezado y su ámbito |
| un enlace | un `Link`, con su área clicable incluida y su texto como descripción |
| una imagen | una `Figure` con su texto alternativo; `![](…)` sin texto alternativo es decoración |
| `$…$`, `$$…$$` | una `Formula`, con su código fuente TeX como descripción |
| código | `Code`, sin sus números de línea ni sus círculos de llamada |
| `<span lang="fr">` | un tramo en ese idioma |
| fondos, bordes, números de página | artefactos, que un lector omite |
| notas del orador (`--pdf-notes`) | anotaciones en la estructura |

El archivo declara PDF/UA-1, el estándar del PDF accesible. Lleva el `title` de
la presentación, que el visor muestra, y su `lang` (`--lang` cuando la
presentación no define ninguno, `en` en caso contrario).

Una presentación verificada con veraPDF cumple PDF/UA-1, salvo que sus propios
encabezados rompan el orden del estándar. PDF/UA exige que el primer encabezado
sea un `H1` y que no se salte ningún nivel (`#` y luego `###`). Solo el autor de
la presentación puede corregirlo. El PDF de Chrome de la misma presentación, el
que escribe marp-cli, también está etiquetado, pero incumple siete de las reglas
de PDF/UA, no tiene una sección por diapositiva ni fórmulas.

Etiquetar la presentación de prueba de 100 diapositivas cuesta un 4 % de tamaño
(788 frente a 757 KiB) y ningún tiempo medible. `--pdf-tags=false` escribe el PDF
sin etiquetas.

## `slides.yaml` {#slidesyaml}

Un archivo de configuración junto a la presentación se aplica a ella, de modo que
una carpeta de presentaciones comparte uno solo:

```yaml
kinds: kinds.yaml        # a mapping of kinds, or file(s) holding one
themes:                  # theme CSS files or directories
  - themes/
data:                    # template data
  company: Example Inc.
html: false              # allow all raw HTML
lang: en
```

Kinds: los `kinds:` de la propia presentación prevalecen sobre los del archivo y
los de la línea de comandos. Datos: los `data` del archivo y de la línea de
comandos prevalecen sobre el front matter de la presentación, de modo que una
misma presentación se puede construir con datos distintos.

## Previsto {#planned}

Todavía no está hecho: las fórmulas como vectores en el PDF. Cuando lleguen, será
un comando u opción aquí.
