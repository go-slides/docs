---
title: "Playground"
weight: 15
description: "Escriba una presentación en el navegador y vea cómo las diapositivas la siguen: go-slides compilado a WebAssembly, con autocompletado, ayuda y ejemplos."
tags: [playground, browser]
---

**[Abrir el playground →](https://go-slides.github.io/playground/)**

El playground es el propio go-slides, v0.9.0, compilado a WebAssembly y
ejecutándose en su navegador. Escriba una presentación a la izquierda; las
diapositivas la siguen a la derecha. No se envía nada a ningún servidor: ni la
presentación, ni un enlace para compartir.

## Qué ofrece {#what-it-has}

- **Vista previa en vivo** de cada diapositiva, renderizada fuera del hilo de la
  página (un Web Worker), de modo que escribir nunca espera por ella. Haga clic en
  una diapositiva para llevar el editor a su línea; mueva el cursor y la vista
  previa se desplaza a su diapositiva.
- **Pase de diapositivas**: la propia presentación, con el teclado, el
  controlador en pantalla y las transiciones.
- **Vista del presentador**: **Present** (o `p` en el pase de diapositivas) la
  abre en una segunda ventana: la diapositiva actual, el paso siguiente, las
  notas del orador (los `<!-- comments -->` de la presentación), un reloj y un
  cronómetro. Las teclas de flecha, Espacio, Inicio y Fin funcionan en cualquiera
  de las dos ventanas, y cada una sigue a la otra, también cuando se edita la
  presentación. Si la ventana no se abre, permita las ventanas emergentes para el
  sitio. La vista general (`o`) solo está en la página descargada.
- **Autocompletado** mientras escribe: claves de directivas (las de una sola
  diapositiva, con `_`, en los comentarios), los valores de `kind`, `theme`,
  `size`, `paginate`, y las funciones de plantilla y los campos de datos dentro
  de `{{ }}`.
- **Ayuda** sobre lo que hay bajo el cursor: una directiva, un kind (tipo de
  diapositiva; qué fija), una función de plantilla. El panel de ayuda enlaza
  también con estas páginas.
- **Problemas** en su línea: los errores de plantilla, de YAML y de HCL se
  resaltan en el editor y se listan debajo, y la vista previa conserva la última
  presentación que se renderizó.
- **Ejemplos**: una presentación de bienvenida, plantillas que generan
  diapositivas, kinds de diapositiva, un front matter en HCL, temas y directivas
  de una sola diapositiva, y matemáticas, imágenes y fondos (`![bg]`, un fondo
  dividido `![bg left:40%]`, filtros).
- **Archivos**: el panel a la izquierda del editor es el proyecto de la
  presentación, organizado como lo leen el [comando]({{< relref "cli" >}}) y la
  [extensión de VS Code]({{< relref "vscode" >}}): `deck.md`, `slides.yaml`,
  `themes/`, `kinds/`, `media/`.
  - Haga clic en un archivo para editarlo. El editor lo colorea según su lenguaje
    (Markdown, YAML, HCL, JSON, CSS) y marca sus problemas.
  - Clic derecho (o Mayús+F10) para New File, New Folder, Rename (F2), Delete
    (pregunta antes) y New theme / New kind. Arrastre un archivo a una carpeta
    para moverlo. El teclado sigue el patrón de árbol de WAI-ARIA, incluida la
    búsqueda al escribir.
  - **Project ▸ New theme** crea `themes/my-gaia.css` (o a partir de default,
    uncover): un `@import` del tema incorporado y las reglas que usted añada. Se
    incorpora a la lista de temas. **New kind** y **New data file** también
    parten de plantillas. Cada uno se lista en `slides.yaml`, de modo que
    `slides build deck.md` lo encuentra.
  - **Project ▸ Download project (.zip)** lo guarda todo. Una vez descomprimido,
    `slides build deck.md` lo construye. **Open project** recupera un `.zip` o
    una carpeta. Se rechaza un archivo comprimido que salga de su carpeta (`..`,
    rutas absolutas), que contenga más de 256 archivos o más de 8 MiB en un
    archivo, o que se descomprima a más de 100 veces su tamaño.
  - Todo se conserva en este navegador junto con el borrador. Un proyecto admite
    hasta 64 MiB, y la línea bajo los archivos indica cuánto se usa.
- **Medios**: suelte o pegue una imagen, un vídeo o un sonido en el editor y pasa
  a `media/`, y la presentación lo nombra: `![](media/photo.png)`. Suelte
  archivos sobre una carpeta del panel para ponerlos allí. La presentación sigue
  siendo pequeña y legible, y no se sube nada. El entorno aislado de la vista
  previa los carga sin abrirse (la página le entrega los bytes). **Save as ▸ HTML
  page** los incluye en la página. Las imágenes `https:` también funcionan, pero
  el navegador las obtiene de su servidor, que entonces ve su dirección. La vista
  previa no envía referrer.
- **Paneles redimensionables**: el panel de archivos, el editor, las miniaturas
  de las diapositivas, la vista previa y la ayuda están separados por barras que
  se pueden arrastrar (el puntero se convierte en un cursor de redimensionado
  sobre una de ellas). El panel junto a la barra se queda con el espacio y los
  demás conservan su tamaño. Haga doble clic en una barra, o pulse Intro sobre
  ella, para cerrar el panel y volver a abrirlo. Con el teclado, F6 pasa de una
  barra a otra, las teclas de flecha mueven una barra 16 px, Inicio y Fin llevan
  el panel anterior a su tamaño mínimo y máximo, y Escape vuelve al editor. La
  pestaña View sigue pudiendo desactivar paneles. Los tamaños que fije se
  conservan en este navegador, con el borrador; los tamaños por defecto (190, 168
  y 280 px, y el editor y la vista previa se reparten el resto en proporción 6 a
  5) nunca se almacenan.
- **Formato**: la pestaña Format edita el Markdown por usted, y cada pulsación es
  un paso de Deshacer.
  - **Texto**: Bold, Italic, Strikethrough e Inline code actúan sobre la
    selección o la palabra bajo el cursor y se alternan (Ctrl+B, Ctrl+I,
    Ctrl+Shift+X y Ctrl+E; Cmd en un Mac). Un botón aparece pulsado cuando el
    cursor está dentro de sus marcas. Markdown no tiene subrayado: Underline
    (Ctrl+U) envuelve el texto en `<u>…</u>`, una etiqueta que la lista de
    permitidos de HTML por defecto de Marp conserva, así que funciona con el HTML
    sin procesar desactivado.
  - **Párrafo**: Heading 1 a 3 (Ctrl+Alt+1..3), viñetas (Ctrl+Shift+8),
    numeración (Ctrl+Shift+7) y una cita actúan sobre las líneas seleccionadas y
    se alternan. La regla horizontal es la etiqueta `<hr>`: `---`, `***` y `___`
    iniciarían cada una una nueva diapositiva en Marp.
  - **Tabla**: una cuadrícula elige las columnas y las filas. Con el cursor en una
    tabla se añade o se elimina una fila o una columna y se alinea la columna
    (`:--`, `:-:`, `--:`); la tabla se reescribe alineada tras cada edición.
  - **Slide text** a la izquierda, al centro o a la derecha escribe un bloque
    `<style scoped>` en la diapositiva actual. Marp no tiene sintaxis de
    alineación, y un atributo `style` no pasa la lista de permitidos de HTML por
    defecto, pero el bloque con ámbito sí, de modo que funciona con el HTML sin
    procesar desactivado y se ve en la vista previa, la página HTML y el PDF.
  - **Link** (Ctrl+K) escribe `[text](url)` con la url seleccionada.
- **Resaltado** de las claves del front matter a cualquier profundidad, en YAML
  (también en mapas `{a: 1}`) y en HCL.
- Selector de **tema**: default, gaia, uncover, los temas del proyecto
  (`themes/*.css`) y los temas que define la presentación (`/* @theme name */` en
  un `<style>`). Escribe `theme:` en el front matter, YAML o HCL, de modo que un
  enlace compartido lo conserva, y sigue al front matter cuando lo edita.
- La vista previa **mantiene su posición** mientras escribe: solo se sustituyen
  las diapositivas que cambiaron, y un cambio de tema reabre la vista previa en
  la misma diapositiva.
- **Enlace para compartir**: la presentación viaja en el propio enlace (en su
  parte `#`, comprimida), así que abrir el enlace abre la presentación. Los
  enlaces admiten hasta 64 KiB comprimidos, y solo la presentación: cuando el
  proyecto contiene más (temas, kinds, medios), Share lo indica y ofrece el
  `.zip`.
- **Save as**:
  - la página HTML, la misma que escribe `slides deck.md`;
  - un **PDF** (vectorial, con el texto seleccionable, con índice);
  - una presentación **PowerPoint** u **OpenDocument** editable.

  Se generan en su navegador, con el mismo código que `slides --pdf`, `--pptx` y
  `--odp`: el PDF de la presentación de prueba de 100 diapositivas es idéntico
  byte a byte al del comando. La primera exportación descarga el motor de
  exportación (unos 5 MiB, una sola vez; su navegador lo conserva hasta la
  siguiente actualización), con una línea de progreso, y lo compila a medida que
  llega. Después solo descarga las fuentes que dibujan sus diapositivas. Con un
  enlace de 3 MB/s, el primer PDF de la presentación de bienvenida está listo en
  menos de 3 segundos. La presentación de 100
  diapositivas tarda entonces unos 3 segundos. Una imagen `https:` solo se incluye
  si su servidor lo permite (CORS); las imágenes `data:` siempre se incluyen.
- **Su borrador se conserva**:
  - La presentación se guarda en este navegador mientras escribe (IndexedDB,
    solo para este sitio) y reaparece al recargar. No se envía nada a ninguna
    parte. Los archivos de texto del proyecto (temas, kinds, datos,
    `slides.yaml`) se conservan del mismo modo, incluso si cierra la pestaña justo
    después de escribir; los medios se almacenan en cuanto se sueltan.
  - Un enlace compartido abre su presentación y conserva su borrador: **Draft ▸
    Restore my draft** lo recupera.
  - Un ejemplo, **Draft ▸ Reset to the example** (que pregunta antes) y otra
    pestaña tampoco lo borran nunca: **Draft ▸ Restore replaced** lo recupera.
  - **Draft ▸ Forget my drafts** borra todo lo almacenado.

Toda exportación renderiza el proyecto que muestra la vista previa: su
`slides.yaml`, temas, kinds, archivos de datos y medios, con los medios dentro
del documento.

Las imágenes como archivos (`--images` de slides) no se generan en el navegador:
use el [comando]({{< relref "cli" >}}) o la
[extensión de VS Code]({{< relref "vscode" >}}).

## Seguridad {#security}

La presentación se muestra en un iframe aislado, sin el origen de la página: no
puede alcanzar ni el playground ni su almacenamiento, no puede enviar ningún
formulario y no navega a ninguna parte. Su página lleva una Content Security
Policy que ejecuta solo sus propios scripts. El HTML sin procesar de la
presentación sigue la lista de permitidos de Marp, como hace `slides` por
defecto. Un renderizado que dure más de 5 segundos (una plantilla que entra en
bucle) se detiene.

La ventana del presentador mantiene ese aislamiento. La abre el propio
playground y muestra la presentación en un iframe aislado de la misma manera. Los
dos iframes no pueden compartir almacenamiento, así que el playground retransmite
el número de diapositiva entre ellos mediante mensajes. Ningún marco que ejecute
la presentación obtiene el origen del playground.

Los medios del proyecto llegan a los marcos aislados como bytes en un mensaje, y
cada marco crea con ellos sus propias URL `blob:`, en su propio origen. Una URL
`blob:` creada por el playground no se carga allí, y la política de los marcos
solo añade `blob:` para medios y fuentes.

## Tamaño y velocidad {#size-and-speed}

El módulo pesa 8,8 MiB tal como lo sirve GitHub Pages (gzip; Pages no sirve
brotli). Se descarga una sola vez con una barra de progreso y luego el navegador
lo conserva en caché hasta el siguiente despliegue. En un Apple M4 Max con Chrome
155, medido sobre la página en línea, el editor está listo 2,1 s después de abrir
la página, 1,5 s de los cuales son la descarga, y las primeras diapositivas
aparecen a los 2,4 s. Un cambio en la presentación de prueba de 100 diapositivas
llega a la vista previa en unos 30 ms: go-slides renderiza solo las diapositivas
que cambiaron, y solo esas se sustituyen en la vista previa.
