---
title: "Matemáticas y código"
weight: 35
description: "Las fórmulas entre signos de dólar se componen y el código en bloques se resalta, en la página HTML, las imágenes y el PDF."
tags: [math, code]
---

{{< callout type="warning" >}}
Las matemáticas y el código llegaron en v0.2.0.
{{< /callout >}}

## Matemáticas {#math}

`$…$` es una fórmula en línea y `$$…$$` una fórmula en bloque:

```markdown
The famous one is $E=mc^2$, and this is a display formula:

$$
\int_0^\infty e^{-x^2}\,dx = \frac{\sqrt{\pi}}{2}
$$
```

- Las reglas de delimitadores son las de Marp Core, leídas de marp-core 4.4.0: un
  `$` seguido de un espacio no abre una fórmula, y uno seguido de un dígito no la
  cierra, de modo que precios como `$5 and $10` se quedan como texto.
- [go-tex/math](https://github.com/go-tex/math) compone cada fórmula en un SVG en
  línea con tamaño en `em`, de modo que sigue el tamaño del texto y se apoya en
  su línea base. El mismo SVG se usa en la página HTML, en las imágenes y en el
  PDF.
- Una fórmula que no se analiza correctamente queda visible tal como se escribió;
  el error se conserva en su `title`.

## Código {#code}

Los bloques de código se colorean con [go-rouge](https://github.com/go-rouge/rouge),
que escribe los nombres de clase de highlight.js, los que usan los temas de Marp:

````markdown
```go
fmt.Println("hello")
```
````

Un lenguaje que go-rouge no conoce se escribe como texto plano, igual que hace
Marp Core.

### Números de línea {#line-numbers}

En v0.8.0 un bloque puede numerar sus líneas. Marp Core no tiene números de línea
e ignora lo que sigue al lenguaje de una valla, así que las opciones van entre
llaves y una presentación que las usa se renderiza en Marp como lo haría sin
ellas:

````markdown
```go {lineNumbers}
func main() {
	fmt.Println("hello")
}
```

```go {lineNumbers=10}
// starts at 10
```
````

| Se escribe | Significa |
| --- | --- |
| `{lineNumbers}` | numera desde 1 (también se leen el `{lines:true}` de Slidev y el `{.numberLines}` de Pandoc) |
| `{lineNumbers=10}` | numera desde 10 (también `startLine=10`, el `startFrom="10"` de Pandoc) |
| `{lineNumbers=false}` | sin números, diga lo que diga la diapositiva |

Para todos los bloques de una diapositiva y de las siguientes, fije la directiva:
`<!-- lineNumbers: true -->` (`_lineNumbers` para una sola diapositiva). Un kind
(tipo de diapositiva) puede fijarla una vez para toda una familia de
diapositivas:

```yaml
kinds:
  listing: {lineNumbers: true}
```

Cada número empieza su línea dentro de la caja de código del tema, separado por
una línea. Se mantiene en la línea base de su línea aunque esta sea más alta, por
ejemplo con una llamada o un emoji. Los números se atenúan con un color mezcla
del color del código y el de su fondo, con un contraste de 4,5:1 o más en los
temas default, gaia y uncover, claros y oscuros. **No se pueden seleccionar** en
la página y los lectores de pantalla no los leen. En un PDF se dibujan pero se
dejan fuera del texto, de modo que un extractor de texto lee solo el código. En
un archivo de PowerPoint o de OpenDocument son texto editable al comienzo de cada
línea. Una presentación que no usa la opción se renderiza como antes, byte a
byte. Resaltar líneas escogidas todavía no es compatible.

### Llamadas en el código {#code-callouts}

En v0.8.0 una línea de código puede terminar con un disco numerado que remite a
una nota situada debajo del bloque. La sintaxis es la de
[Quarto](https://quarto.org/docs/authoring/code-annotation.html), que es también
la forma de comentario de Asciidoctor. El marcador es un comentario en el propio
lenguaje del código, así que el código sigue siendo válido, y las notas son la
lista ordenada que sigue justo después del bloque:

````markdown
```go
func main() { // <1>
	x := load() // <2>
	fmt.Println(x)
}
```

1. The entry point.
2. Reads the input.
````

| Se escribe | Significa |
| --- | --- |
| `// <1>`, `# <1>`, `-- <1>`, `; <1>`, `% <1>`, `' <1>` | un comentario de línea que contiene solo el marcador: ambos se sustituyen por el disco |
| `<!-- <1> -->`, `/* <1> */` | lo mismo, para los lenguajes con comentarios delimitados |
| `// <1> <2>` | dos discos en una línea |
| `x = 1 # count <1>` | el texto del comentario se conserva y el disco lo sigue |

Los marcadores se convierten en discos solo cuando la lista que sigue justo
después del bloque empieza en 1 y tiene un elemento por cada número. En caso
contrario, el bloque se queda como se escribió. Los discos son anillos del color
del código, en los temas default, gaia y uncover, claros u oscuros. Los
elementos de la lista se numeran con los mismos discos.

Los discos **no se pueden seleccionar**. Al copiar el código de la página se
obtiene el código sin sus marcadores. En un PDF son vectores que quedan fuera del
texto, y en un archivo de PowerPoint o de OpenDocument son formas editables. Las
llamadas se combinan con los números de línea. Marp Core no tiene llamadas
(marp-core 4.4.0), así que en Marp los comentarios y la lista numerada se ven
tal como se escribieron. Una presentación que no las usa se renderiza como antes,
byte a byte.

En v0.9.0, un disco y su nota se encuentran mutuamente en la página HTML. Pasar
el puntero por un disco, o llegar a él con Tab, tiñe su nota, y una nota hace lo
mismo con su disco. Es lo que hace el `code-annotations: hover` de Quarto, pero
aquí lo hace solo la hoja de estilos, sin ningún script. Por eso funciona igual
en la vista previa aislada del Playground y en la vista previa de VS Code, que
restringen los scripts. Cada disco se describe con su nota, de modo que un lector
de pantalla lee la nota al llegar al disco. El tinte aparece gradualmente solo
cuando el lector no ha pedido movimiento reducido. Las imágenes, el PDF y los
archivos de PowerPoint y de OpenDocument no cambian, byte a byte.

## Fuentes {#fonts}

Lato Regular, Bold, Black e itálicas, y Roboto Mono Regular, Bold e itálicas van
incrustadas, de modo que los títulos de gaia tienen su grosor y el texto en
cursiva se renderiza en cursiva.
