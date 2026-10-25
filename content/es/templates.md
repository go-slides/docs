---
title: "Plantillas"
weight: 30
description: "Interpolación con text/template de Go en una presentación: valores, diapositivas generadas, números de página, inclusiones."
tags: [templates]
---

La presentación es una [`text/template`](https://pkg.go.dev/text/template) de Go,
evaluada **antes** de analizar el Markdown. Todo lo que ofrece el lenguaje de
plantillas de Go (acciones, tuberías, `if`, `range`, `with`, variables)
funciona en una presentación.

## Datos {#data}

El punto de la plantilla contiene las claves del front matter, luego las claves
de su mapa `data:`, y luego los datos indicados en la línea de comandos o en
`slides.yaml`, cada uno sobrescribiendo al anterior.

```markdown
---
marp: true
title: Quarterly review
data:
  team: Platform
---

# {{ .title }}

Presented by the {{ .team }} team.
```

Desde la línea de comandos:

| Opción | Aporta |
| --- | --- |
| `--data file.yaml` | un archivo YAML o JSON de datos (repetible) |
| `--set key=value` | un valor (repetible) |

## Generar diapositivas {#generating-slides}

Todo el cuerpo es una plantilla, así que un `range` que escribe `---` escribe
diapositivas:

```markdown
{{ range .releases }}
---

# {{ .version }}

{{ range .changes }}- {{ . }}
{{ end }}
{{ end }}
```

## Página y total {#page-and-total}

`{{ page }}` y `{{ total }}` son el número de la diapositiva y el número de
diapositivas. Solo se conocen una vez dividida la presentación, así que se
rellenan entonces: una presentación generada cuenta las diapositivas que generó.

```markdown
<!-- footer: "{{ page }} / {{ total }}" -->
```

## Incluir {#include}

`{{ include "parts/agenda.md" }}` inserta otro archivo, relativo a la
presentación.

## El código se deja tal cual {#code-is-left-alone}

Los bloques de código y los fragmentos de código en línea se apartan antes de
ejecutar la plantilla y se restituyen sin cambios, de modo que una diapositiva
sobre plantillas de Go puede mostrar `{{ .Name }}` sin escaparlo.

## Desactivarlas {#turning-it-off}

`--no-template` lee la presentación como Marp plano: `{{` es texto.
