---
title: "Matematica e codice"
weight: 35
description: "Le formule tra simboli del dollaro vengono composte e il codice delimitato viene evidenziato, nella pagina HTML, nelle immagini e nel PDF."
tags: [math, code]
---

{{< callout type="warning" >}}
Matematica e codice sono nella v0.2.0.
{{< /callout >}}

## Matematica {#math}

`$…$` è una formula in linea e `$$…$$` una formula in evidenza:

```markdown
The famous one is $E=mc^2$, and this is a display formula:

$$
\int_0^\infty e^{-x^2}\,dx = \frac{\sqrt{\pi}}{2}
$$
```

- Le regole sui delimitatori sono quelle di Marp Core, lette da marp-core 4.4.0: un `$`
  seguito da uno spazio non apre una formula, e uno seguito da una cifra
  non la chiude, quindi prezzi come `$5 and $10` restano testo.
- [go-tex/math](https://github.com/go-tex/math) compone ogni formula in
  un SVG in linea dimensionato in `em`, così segue la dimensione del testo e poggia sulla sua
  linea di base. Lo stesso SVG è usato nella pagina HTML, nelle immagini e nel PDF.
- Una formula che non viene analizzata resta visibile com'è scritta; l'errore è conservato
  nel suo `title`.

## Codice {#code}

I blocchi di codice delimitati sono colorati da [go-rouge](https://github.com/go-rouge/rouge),
che scrive i nomi di classe di highlight.js, quelli che i temi di Marp stilizzano:

````markdown
```go
fmt.Println("hello")
```
````

Un linguaggio che go-rouge non conosce viene scritto come testo semplice, come fa Marp Core.

### Numeri di riga {#line-numbers}

Nella v0.8.0 un blocco può numerare le proprie righe. Marp Core non ha numeri di riga e
ignora ciò che segue il linguaggio di un blocco, quindi le opzioni vanno tra parentesi graffe
e una presentazione che le usa viene resa in Marp come lo sarebbe senza:

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

| Scrivi | Significa |
| --- | --- |
| `{lineNumbers}` | numera da 1 (vengono lette anche `{lines:true}` di Slidev e `{.numberLines}` di Pandoc) |
| `{lineNumbers=10}` | numera da 10 (anche `startLine=10`, `startFrom="10"` di Pandoc) |
| `{lineNumbers=false}` | nessun numero, qualunque cosa dica la slide |

Per ogni blocco di una slide e delle slide successive, imposta la direttiva:
`<!-- lineNumbers: true -->` (`_lineNumbers` per una sola slide). Un kind (il tipo di una slide) può impostarla
una volta per tutta una famiglia di slide:

```yaml
kinds:
  listing: {lineNumbers: true}
```

Ogni numero apre la propria riga dentro il riquadro del codice del tema, separato da una linea.
Resta sulla linea di base della sua riga anche quando la riga è più alta, per esempio
con una chiamata o un'emoji. I numeri sono attenuati da un colore ottenuto mescolando
il colore del codice e il suo sfondo, con un contrasto di 4,5:1 o più nei temi default,
gaia e uncover, chiari e scuri. **Non sono selezionabili** nella
pagina e non vengono letti dagli screen reader. In un PDF sono disegnati ma esclusi
dal testo, quindi un estrattore di testo legge solo il codice. In un file PowerPoint o
OpenDocument sono testo modificabile all'inizio di ogni riga. Una presentazione
che non usa l'opzione viene resa come prima, byte per byte. L'evidenziazione
di righe scelte non è ancora supportata.

### Chiamate nel codice {#code-callouts}

Nella v0.8.0 una riga di codice può terminare con un cerchio numerato che rimanda a una nota
sotto il blocco. La sintassi è
[quella di Quarto](https://quarto.org/docs/authoring/code-annotation.html), che è
anche la forma a commento di Asciidoctor. Il marcatore è un commento nel linguaggio stesso del codice,
quindi il codice resta valido, e le note sono l'elenco ordinato subito
dopo il blocco:

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

| Scrivi | Significa |
| --- | --- |
| `// <1>`, `# <1>`, `-- <1>`, `; <1>`, `% <1>`, `' <1>` | un commento di riga che contiene solo il marcatore: entrambi vengono sostituiti dal cerchio |
| `<!-- <1> -->`, `/* <1> */` | lo stesso, per i linguaggi con commenti delimitati |
| `// <1> <2>` | due cerchi su una riga |
| `x = 1 # count <1>` | il testo del commento resta e il cerchio lo segue |

I marcatori diventano cerchi solo se l'elenco subito dopo il blocco inizia da 1
e ha una voce per ogni numero. Altrimenti il blocco resta com'è scritto. I
cerchi sono anelli nel colore del codice, nei temi default, gaia e uncover,
chiari o scuri. Le voci dell'elenco sono numerate con gli stessi cerchi.

I cerchi **non sono selezionabili**. Copiare il codice dalla pagina dà il codice
senza i suoi marcatori. In un PDF sono vettori esclusi dal testo, e in un
file PowerPoint o OpenDocument sono forme modificabili. Le chiamate si combinano
con i numeri di riga. Marp Core non ha chiamate (marp-core 4.4.0), quindi in Marp i
commenti e l'elenco numerato appaiono come scritti. Una presentazione che non le usa
viene resa come prima, byte per byte.

Nella v0.9.0, un cerchio e la sua nota si ritrovano a vicenda nella pagina HTML. Passare sopra
un cerchio, o raggiungerlo con Tab, colora la sua nota, e una nota fa lo stesso con il suo
cerchio. È ciò che fa `code-annotations: hover` di Quarto, ma qui il foglio di
stile lo fa da solo, senza script. Funziona quindi allo stesso modo nell'anteprima
isolata del Playground e nell'anteprima di VS Code, che limitano entrambe gli script. Ogni cerchio è descritto dalla sua nota, quindi uno screen reader legge la nota
quando raggiunge il cerchio. La colorazione compare con dissolvenza solo se il lettore non ha chiesto
di ridurre il movimento. Immagini, PDF, file PowerPoint e OpenDocument
restano invariati, byte per byte.

## Font {#fonts}

Lato Regular, Bold, Black e corsivi e Roboto Mono Regular, Bold e
corsivi sono incorporati, così i titoli di gaia ottengono il loro peso e il testo
in corsivo viene reso in corsivo.
