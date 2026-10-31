---
title: "VS Code"
weight: 55
description: "La extensión de go-slides para VS Code: vista previa, presentación, problemas, autocompletado y exportación, como hace Marp for VS Code con Marp."
tags: [vscode, editor]
---

{{< callout type="info" >}}
La extensión es [go-slides/vscode](https://github.com/go-slides/vscode), con
versionado propio (v0.5.0). Renderiza con go-slides v0.14.0 (el paquete `editor`,
las líneas de origen de la sincronización del desplazamiento, y la vista general
y la vista del presentador). Se instala desde el `go-slides-0.5.0.vsix` adjunto a
su versión de GitHub: su publicación en VS Code Marketplace y en Open VSX está
preparada, pero todavía no está allí.
{{< /callout >}}

La extensión es el equivalente de go-slides de
[Marp for VS Code](https://github.com/marp-team/marp-vscode): la presentación se
muestra como diapositivas en la vista previa de Markdown de VS Code, sus errores
en el panel Problemas, y las directivas, los kinds (tipos de diapositiva) y las
funciones de plantilla se autocompletan mientras se escribe.

## Instalación {#install}

Descargue `go-slides-<version>.vsix` de las
[versiones](https://github.com/go-slides/vscode/releases) y ejecute
*Extensions: Install from VSIX...*, o bien:

```sh
code --install-extension go-slides-0.5.0.vsix
```

La vista previa no necesita nada más. Para exportar hace falta el comando
`slides`:

```sh
go install github.com/go-slides/slides/cmd/slides@latest
```

## Vista previa {#preview}

Abra una presentación y ejecute *Markdown: Open Preview to the Side*. Una
presentación es un archivo Markdown cuyo front matter dice `slides: true`:

```markdown
---
slides: true
paginate: true
data:
  who: World
---

# Hello {{ .who }}

---

<!-- _kind: title -->
# Page {{ page }} of {{ total }}
```

Los archivos con `marp: true` también se previsualizan, salvo que esté instalado
Marp for VS Code: ambas extensiones renderizarían el mismo documento. Lo decide
el ajuste `goSlides.preview.marpDocuments`: `auto` (el valor por defecto),
`always` o `never`.

La vista previa es el propio go-slides, compilado a WebAssembly y cargado por la
extensión: sin binario, y funciona en [vscode.dev](https://vscode.dev). Lee del
espacio de trabajo lo que usa la presentación (`slides.yaml` o `slides.hcl` junto
a la presentación, los archivos de [kinds]({{< relref "kinds" >}}) que nombran,
los archivos que incluye una [plantilla]({{< relref "templates" >}}), el CSS de
los temas) y solo de la carpeta del espacio de trabajo de la presentación.

## Sincronización del desplazamiento {#scroll-sync}

Desde la v0.2.0 de la extensión (go-slides v0.7.0), el editor y la vista previa
se desplazan a la vez, en ambos sentidos, como ocurre con el Markdown normal: la
vista previa sigue la diapositiva y el bloque que hay bajo el cursor del editor,
y desplazar la vista previa mueve el editor. Los ajustes son los propios de VS
Code, `markdown.preview.scrollPreviewWithEditor` y
`markdown.preview.scrollEditorWithPreview`.

Las líneas se rastrean a través de las plantillas. Una diapositiva que genera un
`{{ range }}` sigue la línea de su plantilla, y lo que aporta un `{{ include }}`
sigue la línea de la inclusión. Una presentación dividida por `headingDivider` se
sincroniza con sus títulos.

## Presentar {#presenting}

Desde la v0.4.0 de la extensión (go-slides v0.9.0), una presentación se puede
proyectar sin salir de VS Code. Con la presentación en el editor, desde la paleta
de comandos o el menú `…` del editor:

| Comando | Abre |
|---|---|
| *go-slides: Open Overview* | un panel con todas las diapositivas en una cuadrícula. Una diapositiva elegida muestra su código fuente en el editor; `Esc` cierra el panel |
| *go-slides: Start Presenter View* | las diapositivas y, a su lado, la vista del presentador: el paso siguiente, las notas del orador y un cronómetro |

Los paneles se sitúan junto al editor y se siguen entre sí: al cambiar de
diapositiva en uno, cambian también los demás. También siguen sus ediciones de la
presentación. La tecla `p` en las diapositivas abre la vista del presentador.

Para presentar en otra pantalla, mueva allí el panel de las diapositivas con
*View: Move Editor into New Window*.

Cada panel ejecuta la presentación en un marco aislado (sin `allow-same-origin`)
con una política `default-src 'none'`: el marco carga imágenes, medios, estilos y
fuentes y nada más, y ejecuta solo los scripts que go-slides incorporó a la
página, identificados por su hash. Un script escrito por la presentación no se
ejecuta.

## Problemas, autocompletado, información al pasar el puntero {#problems-completion-hover}

- **Problemas**: errores de plantilla, de YAML y de [HCL]({{< relref "hcl" >}}),
  en la línea en que se producen, en la presentación o en el archivo de kinds de
  donde proceden.
- **Autocompletado**: las claves de directivas en el front matter y en los
  comentarios `<!-- -->` (con `_` para una directiva de una sola diapositiva), los
  valores de `kind` y `_kind` (con lo que fija cada kind), `theme`, `size`,
  `paginate`, `math`, `headingDivider`; dentro de `{{ }}`, las funciones de
  plantilla y los campos de datos.
- **Información al pasar el puntero**: qué hace una directiva, qué fija un kind,
  los argumentos de una función de plantilla.

## Exportación {#export}

*go-slides: Export Slide Deck...* (también en la barra de título del editor)
escribe HTML, PDF vectorial, PPTX y ODP editables, PNG o JPEG con el comando
`slides`, como lo haría `slides -o deck.pdf deck.md`; véase
[Línea de comandos]({{< relref "cli" >}}). Necesita un espacio de trabajo de
confianza. Si no se encuentra `slides`, la extensión ofrece ejecutar en un
terminal el `go install` de arriba.

Desde la v0.5.0 de la extensión (go-slides v0.14.0), el PDF que escribe es
**accesible**: con `slides` v0.12.0 o posterior, el PDF vectorial lleva etiquetas
para lectores de pantalla y declara PDF/UA-1, con el `title` y el `lang` de la
presentación y una sección por diapositiva. `slides --pdf-tags=false` escribe el
archivo sin etiquetas. Actualice el comando con el `go install` de arriba.

## Ajustes {#settings}

| Ajuste | Valor por defecto | |
|---|---|---|
| `goSlides.cliPath` | `""` | el comando `slides`; vacío: `slides` en el `PATH`. Es un ajuste de la máquina: un espacio de trabajo no puede fijarlo |
| `goSlides.exportType` | `pdf` | el formato que propone primero el diálogo de exportación |
| `goSlides.exportAutoOpen` | `true` | abre el archivo exportado |
| `goSlides.html` | `false` | permite todo el HTML sin procesar en lugar de la lista de permitidos de Marp; se respeta solo en un espacio de trabajo de confianza |
| `goSlides.preview.marpDocuments` | `auto` | previsualiza las presentaciones con `marp: true`: `auto`, `always`, `never` |

## Seguridad {#security}

La presentación y el espacio de trabajo no son de confianza. El HTML sin procesar
sigue la lista de permitidos de Marp salvo que `goSlides.html` esté activado en un
espacio de trabajo de confianza (un `slides.yaml` con `html: true` no lo activa en
la vista previa), y la vista previa de VS Code no ejecuta ningún script de la
presentación. La vista previa solo lee archivos de la carpeta del espacio de
trabajo de la presentación, no sigue ningún enlace simbólico, y las plantillas no
ven ninguna variable de entorno. Exportar ejecuta un programa, así que espera a
que se confíe en el espacio de trabajo, y el espacio de trabajo nunca nombra el
programa.

## Rendimiento {#performance}

Con la presentación de prueba de 100 diapositivas de go-slides, en VS Code 1.142
sobre un Apple M4 Max, con Marp for VS Code 3.6.1 medido en la misma sesión:

| | go-slides | Marp for VS Code |
|---|---:|---:|
| activación (carga del módulo WebAssembly) | 189 ms | — |
| primera vista previa | 52 ms | 78 ms |
| vista previa tras una edición (mediana) | **19 ms** | 23–24 ms |
| `.vsix` | 7,5 MB | 36,3 MB |

Las cifras corresponden a la extensión v0.3.0 (go-slides v0.8.0) con una
presentación de 100 diapositivas, escribiendo un carácter en la diapositiva 50,
en rondas intercaladas con una carga de 1 minuto inferior a 8. La primera vista
previa es más rápida que la de Marp, y la vista previa tras una edición también
lo es ya.

### En v0.8.0 y la extensión v0.3.0 {#in-v080-and-extension-v030}

go-slides v0.8.0 conserva las diapositivas renderizadas de cada presentación
abierta y, tras una edición, vuelve a renderizar solo aquellas cuyo origen o
estado heredado cambió (directivas, número de página, total, ids de títulos,
referencias de enlaces, líneas de origen); el resultado es idéntico al de un
renderizado completo, byte a byte. Escribir en la diapositiva 50 de 100 vuelve a
renderizar una diapositiva: 2,5 ms de forma nativa, 11,5 ms en WebAssembly. La
vista previa de VS Code transforma entonces solo el DOM de la diapositiva
modificada, y la prueba de humo de la extensión escribe un carácter en la
diapositiva 50 entre renderizados cronometrados, para ambas extensiones. La vista
del presentador de la extensión se sincroniza a través de la ventana padre
(`postMessage`).

### En v0.7.1 y la extensión v0.2.2 {#in-v071-and-extension-v022}

go-slides v0.7.1 almacena en caché entre renderizados los temas analizados, las
hojas de estilo empaquetadas y el código resaltado, marca las líneas de
sincronización del desplazamiento sin recorrer byte a byte, y reutiliza sus
escritores: un renderizado de la presentación de 100 diapositivas reserva 5,7 MB,
frente a los 52,7 MB de la extensión v0.2.0. Con ello, la extensión v0.2.2
elimina la sobrescritura `GOGC=400`. A modo de comparación, la extensión v0.2.1
tardaba 72 ms en la primera vista previa y 44 ms tras una edición con el `GOGC`
por defecto (36 ms con `GOGC=400`), y v0.2.0 tardaba entre 47 y 51 ms.

La extensión v0.2.2 midió 52 ms en la primera vista previa y 24 ms tras una
edición, a la par con los 23 ms de Marp.
