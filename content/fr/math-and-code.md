---
title: "Mathématiques et code"
weight: 35
description: "Les formules entre dollars sont composées et le code délimité est coloré, dans la page HTML, les images et le PDF."
tags: [math, code]
---

{{< callout type="warning" >}}
Les mathématiques et le code sont arrivés en v0.2.0.
{{< /callout >}}

## Mathématiques {#math}

`$…$` est une formule en ligne et `$$…$$` une formule centrée :

```markdown
The famous one is $E=mc^2$, and this is a display formula:

$$
\int_0^\infty e^{-x^2}\,dx = \frac{\sqrt{\pi}}{2}
$$
```

- Les règles de délimitation sont celles de Marp Core, lues dans marp-core 4.4.0 : un `$`
  suivi d'une espace n'ouvre pas de formule, et un dollar suivi d'un chiffre
  ne la ferme pas, si bien que des prix comme `$5 and $10` restent du texte.
- [go-tex/math](https://github.com/go-tex/math) compose chaque formule en
  SVG en ligne dimensionné en `em` : elle suit la taille du texte et repose sur sa
  ligne de base. Le même SVG sert dans la page HTML, dans les images et dans le PDF.
- Une formule qui ne s'analyse pas reste visible telle qu'elle est écrite ; l'erreur est conservée dans
  son `title`.

## Code {#code}

Les blocs de code délimités sont colorés par [go-rouge](https://github.com/go-rouge/rouge),
qui écrit les noms de classes de highlight.js, ceux que stylent les thèmes de Marp :

````markdown
```go
fmt.Println("hello")
```
````

Un langage que go-rouge ne connaît pas est écrit en texte brut, comme le fait Marp Core.

### Numéros de ligne {#line-numbers}

Depuis la v0.8.0, un bloc peut numéroter ses lignes. Marp Core n'a pas de numéros de ligne et
ignore ce qui suit le langage d'une clôture : les options se placent donc là, entre accolades,
et une présentation qui les utilise s'affiche dans Marp comme elle s'afficherait sans elles :

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

| Écriture | Signification |
| --- | --- |
| `{lineNumbers}` | numérotation à partir de 1 (le `{lines:true}` de Slidev et le `{.numberLines}` de Pandoc sont aussi reconnus) |
| `{lineNumbers=10}` | numérotation à partir de 10 (également `startLine=10`, ou le `startFrom="10"` de Pandoc) |
| `{lineNumbers=false}` | aucun numéro, quoi qu'indique la diapositive |

Pour tous les blocs d'une diapositive et des suivantes, définissez la directive
`<!-- lineNumbers: true -->` (`_lineNumbers` pour une seule diapositive). Un *kind* (le type d'une diapositive) peut la définir
une fois pour toute une famille de diapositives :

```yaml
kinds:
  listing: {lineNumbers: true}
```

Chaque numéro ouvre sa ligne à l'intérieur du cadre de code du thème, séparé par un filet.
Il reste sur la ligne de base de sa ligne même quand celle-ci est plus haute, par exemple
à cause d'une annotation ou d'un emoji. Les numéros sont atténués par une couleur mélangeant la
couleur du code et celle de son fond, avec un contraste d'au moins 4,5:1 dans les thèmes default,
gaia et uncover, clairs comme sombres. Ils ne sont **pas sélectionnables** dans la
page et ne sont pas lus par les lecteurs d'écran. Dans un PDF, ils sont dessinés mais exclus
du texte : un extracteur de texte ne lit que le code. Dans un fichier PowerPoint ou
OpenDocument, ce sont des textes modifiables en début de chaque ligne. Une présentation
qui n'utilise pas l'option s'affiche comme avant, à l'octet près. La mise en évidence
de lignes choisies n'est pas encore prise en charge.

### Annotations de code {#code-callouts}

Depuis la v0.8.0, une ligne de code peut se terminer par une pastille numérotée qui renvoie à une note
sous le bloc. La syntaxe est
[celle de Quarto](https://quarto.org/docs/authoring/code-annotation.html), qui est
aussi la forme en commentaire d'Asciidoctor. Le marqueur est un commentaire dans le langage même
du code, qui reste donc valide, et les notes forment la liste numérotée qui suit
immédiatement le bloc :

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

| Écriture | Signification |
| --- | --- |
| `// <1>`, `# <1>`, `-- <1>`, `; <1>`, `% <1>`, `' <1>` | un commentaire de ligne qui ne contient que le marqueur : les deux sont remplacés par la pastille |
| `<!-- <1> -->`, `/* <1> */` | la même chose, pour les langages à commentaires délimités |
| `// <1> <2>` | deux pastilles sur une ligne |
| `x = 1 # count <1>` | le texte du commentaire reste et la pastille le suit |

Les marqueurs ne deviennent des pastilles que si la liste qui suit immédiatement le bloc commence à 1
et comporte un élément pour chaque numéro. Sinon, le bloc reste tel qu'il est écrit. Les
pastilles sont des anneaux de la couleur du code, dans les thèmes default, gaia et uncover,
clairs ou sombres. Les éléments de la liste sont numérotés avec les mêmes pastilles.

Les pastilles ne sont **pas sélectionnables**. Copier le code depuis la page donne le code
sans ses marqueurs. Dans un PDF, ce sont des vecteurs exclus du texte, et dans un fichier
PowerPoint ou OpenDocument, des formes modifiables. Les annotations se combinent
avec les numéros de ligne. Marp Core n'a pas d'annotations (marp-core 4.4.0) : dans Marp, les
commentaires et la liste numérotée s'affichent donc tels qu'ils sont écrits. Une présentation qui ne les utilise pas
s'affiche comme avant, à l'octet près.

Depuis la v0.9.0, une pastille et sa note se répondent dans la page HTML. Survoler une
pastille, ou l'atteindre avec Tab, teinte sa note, et une note fait de même pour sa
pastille. C'est ce que fait `code-annotations: hover` dans Quarto, mais ici la feuille de style
s'en charge seule, sans script. Cela fonctionne donc de la même façon dans l'aperçu
isolé du playground et dans l'aperçu de VS Code, qui restreignent tous deux les
scripts. Chaque pastille est décrite par sa note : un lecteur d'écran lit la note
quand il atteint la pastille. La teinte n'apparaît en fondu que si le lecteur n'a pas demandé
de mouvements réduits. Les images, les PDF et les fichiers PowerPoint et OpenDocument sont
inchangés, à l'octet près.

## Polices {#fonts}

Lato Regular, Bold, Black et italiques ainsi que Roboto Mono Regular, Bold et
italiques sont intégrées : les titres de gaia ont leur graisse et le texte en italique
s'affiche bien en italique.
