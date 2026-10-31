---
title: "Matemática e código"
weight: 35
description: "Fórmulas entre cifrões são compostas e o código delimitado é destacado, na página HTML, em imagens e no PDF."
tags: [math, code]
---

{{< callout type="warning" >}}
Matemática e código estão na v0.2.0.
{{< /callout >}}

## Matemática {#math}

`$…$` é uma fórmula em linha e `$$…$$` uma fórmula em destaque:

```markdown
The famous one is $E=mc^2$, and this is a display formula:

$$
\int_0^\infty e^{-x^2}\,dx = \frac{\sqrt{\pi}}{2}
$$
```

- As regras de delimitadores são as do Marp Core, lidas do marp-core 4.4.0: um
  `$` seguido de espaço não abre uma fórmula, e um seguido de dígito não a
  fecha, de modo que preços como `$5 and $10` continuam texto.
- O [go-tex/math](https://github.com/go-tex/math) compõe cada fórmula em um SVG
  em linha dimensionado em `em`, de modo que ela acompanha o tamanho do texto e
  se apoia na sua linha de base. O mesmo SVG é usado na página HTML, nas imagens
  e no PDF.
- Uma fórmula que não pode ser analisada continua visível como foi escrita; o
  erro é guardado no seu `title`.

## Código {#code}

Os blocos de código delimitados são coloridos pelo
[go-rouge](https://github.com/go-rouge/rouge), que escreve nomes de classe do
highlight.js, os que os temas do Marp estilizam:

````markdown
```go
fmt.Println("hello")
```
````

Uma linguagem que o go-rouge não conhece é escrita como texto simples, como o
Marp Core faz.

### Números de linha {#line-numbers}

Na v0.8.0, um bloco pode numerar suas linhas. O Marp Core não tem números de
linha e ignora o que vem depois da linguagem de uma cerca, então as opções vão
entre chaves ali, e uma apresentação que as usa é renderizada no Marp como seria
sem elas:

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

| Escreva | Significa |
| --- | --- |
| `{lineNumbers}` | numerar a partir de 1 (o `{lines:true}` do Slidev e o `{.numberLines}` do Pandoc também são lidos) |
| `{lineNumbers=10}` | numerar a partir de 10 (também `startLine=10`, o `startFrom="10"` do Pandoc) |
| `{lineNumbers=false}` | sem números, o que quer que o slide diga |

Para todos os blocos de um slide e dos slides seguintes, defina a diretiva:
`<!-- lineNumbers: true -->` (`_lineNumbers` para um único slide). Um kind (o tipo de um slide) pode
defini-la uma vez para uma família de slides:

```yaml
kinds:
  listing: {lineNumbers: true}
```

Cada número inicia a sua linha dentro da caixa de código do tema, separado por
um filete. Ele permanece na linha de base da sua linha mesmo quando a linha é
mais alta, por exemplo com uma chamada (callout) ou um emoji. Os números são
atenuados por uma cor misturada a partir da cor do código e do seu fundo, com
contraste de 4,5:1 ou mais nos temas default, gaia e uncover, claros e escuros.
Eles **não são selecionáveis** na página e não são lidos por leitores de tela.
Em um PDF, eles são desenhados mas ficam fora do texto, de modo que um
extrator de texto lê apenas o código. Em um arquivo PowerPoint ou OpenDocument,
são texto editável no início de cada linha. Uma apresentação que não usa a opção
é renderizada como antes, byte a byte. Destacar linhas escolhidas ainda não é
suportado.

### Chamadas no código {#code-callouts}

Na v0.8.0, uma linha de código pode terminar com um disco numerado que remete a
uma nota abaixo do bloco. A sintaxe é a do
[Quarto](https://quarto.org/docs/authoring/code-annotation.html), que também é a
forma de comentário do Asciidoctor. O marcador é um comentário na linguagem do
próprio código, de modo que o código continua válido, e as notas são a lista
ordenada logo após o bloco:

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

| Escreva | Significa |
| --- | --- |
| `// <1>`, `# <1>`, `-- <1>`, `; <1>`, `% <1>`, `' <1>` | um comentário de linha que contém apenas o marcador: ambos são substituídos pelo disco |
| `<!-- <1> -->`, `/* <1> */` | o mesmo, para linguagens com comentários delimitados |
| `// <1> <2>` | dois discos em uma linha |
| `x = 1 # count <1>` | o texto do comentário permanece e o disco o segue |

Os marcadores só viram discos quando a lista logo após o bloco começa em 1 e tem
um item para cada número. Caso contrário, o bloco continua como foi escrito. Os
discos são anéis na cor do código, nos temas default, gaia e uncover, claros ou
escuros. Os itens da lista são numerados com os mesmos discos.

Os discos **não são selecionáveis**. Copiar o código da página dá o código sem
os marcadores. Em um PDF, eles são vetores deixados fora do texto, e em um
arquivo PowerPoint ou OpenDocument são formas editáveis. As chamadas se combinam
com os números de linha. O Marp Core não tem chamadas (marp-core 4.4.0), então
no Marp os comentários e a lista numerada aparecem como foram escritos. Uma
apresentação que não as usa é renderizada como antes, byte a byte.

Na v0.9.0, um disco e a sua nota se encontram na página HTML. Passar o mouse
sobre um disco, ou chegar a ele com Tab, realça a sua nota, e uma nota faz o
mesmo com o seu disco. É o que o `code-annotations: hover` do Quarto faz, mas
aqui a folha de estilo faz isso sozinha, sem script. Por isso funciona do mesmo
jeito na pré-visualização em sandbox do Playground e na pré-visualização do VS
Code, que restringem scripts. Cada disco é descrito pela sua nota, então um
leitor de tela lê a nota quando chega ao disco. O realce surge gradualmente
somente quando o leitor não pediu movimento reduzido. Imagens, PDF e arquivos
PowerPoint e OpenDocument permanecem inalterados, byte a byte.

## Fontes {#fonts}

Lato Regular, Bold, Black e itálicos e Roboto Mono Regular, Bold e itálicos
estão incorporadas, de modo que os títulos do gaia ganham o seu peso e o texto
em itálico é renderizado em itálico.
