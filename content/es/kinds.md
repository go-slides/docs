---
title: "Kinds"
weight: 40
description: "Una diapositiva declara lo que es: kind nombra un conjunto predefinido de directivas locales."
tags: [kinds]
---

Un ***kind*** (tipo de diapositiva) es un conjunto predefinido y con nombre de
directivas locales: encabezado, pie de página, paginación, clase, fondo. Una
diapositiva declara lo que es (`title`, `section`) en lugar de repetir cómo se
ve, y el aspecto de una presentación se cambia en un único lugar.

## Usar un kind {#using-a-kind}

```markdown
<!-- kind: section -->
# Part two

---

<!-- _kind: blank -->
![bg](photo.jpg)
```

- `kind:` se aplica como una directiva local: a esta diapositiva y a las
  siguientes.
- `_kind:` se aplica como una directiva de una sola diapositiva: solo a esta.
- Cambiar de kind **borra** primero lo que fijó el kind anterior, de modo que una
  clase o un pie de página nunca pasa de un kind al siguiente.
- Las directivas escritas en la propia diapositiva **prevalecen** sobre el kind.

## Incorporados {#built-in}

| Kind | Fija |
| --- | --- |
| `title` | `class: lead`, sin número de página, sin encabezado, sin pie de página |
| `section` | `class: lead` |
| `content` | nada: los valores por defecto de la presentación |
| `blank` | sin número de página, sin encabezado, sin pie de página |

Cualquiera de ellos se puede redefinir.

## Definir kinds {#defining-kinds}

En el front matter:

```yaml
---
marp: true
kinds:
  chapter:
    class: invert
    backgroundColor: "#1e1b4b"
    paginate: false
  quote:
    extends: blank
    class: quote
---
```

`extends` parte de las directivas de otro kind. Un kind acepta las directivas
locales y `extends`.

## Compartir kinds {#sharing-kinds}

Los kinds de un equipo van en un único archivo que nombran todas las
presentaciones:

```yaml
# kinds.yaml
chapter:
  class: invert
  paginate: false
```

```yaml
---
marp: true
kinds: kinds.yaml        # or a list of files
---
```

o bien `--kinds kinds.yaml` en la línea de comandos, o `kinds:` en `slides.yaml`.
El orden es: los kinds incorporados, luego la línea de comandos y el archivo de
configuración y, por último, los `kinds:` de la propia presentación.

`slides kinds talk.md` lista los kinds que puede usar una presentación y lo que
fija cada uno.

## Salir de go-slides {#leaving-go-slides}

`slides expand` escribe cada kind como las directivas locales de una sola
diapositiva que representa, diapositiva a diapositiva, de modo que el resultado
es Markdown de Marp plano que lee cualquier herramienta de Marp.
