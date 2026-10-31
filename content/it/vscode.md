---
title: "VS Code"
weight: 55
description: "L'estensione go-slides per VS Code: anteprima, presentazione, problemi, completamento ed esportazione, come Marp for VS Code fa per Marp."
tags: [vscode, editor]
---

{{< callout type="info" >}}
L'estensione è [go-slides/vscode](https://github.com/go-slides/vscode),
con versioni proprie (v0.5.0). Esegue il rendering con go-slides v0.14.0 (il
pacchetto `editor`, le righe sorgente della sincronizzazione dello scorrimento, e la vista panoramica e
del relatore). Si installa dal file `go-slides-0.5.0.vsix` allegato alla
sua release su GitHub: la pubblicazione sul VS Code Marketplace e su Open VSX è
preparata, ma non è ancora disponibile.
{{< /callout >}}

L'estensione è la controparte di go-slides di
[Marp for VS Code](https://github.com/marp-team/marp-vscode): la presentazione appare
nell'anteprima Markdown di VS Code come slide, i suoi errori nel pannello Problemi, e
direttive, kinds (i tipi di slide) e funzioni di template vengono completati mentre digiti.

## Installazione {#install}

Scarica `go-slides-<version>.vsix` dalle
[release](https://github.com/go-slides/vscode/releases), poi esegui
*Extensions: Install from VSIX...*, oppure:

```sh
code --install-extension go-slides-0.5.0.vsix
```

L'anteprima non richiede altro. L'esportazione richiede il comando `slides`:

```sh
go install github.com/go-slides/slides/cmd/slides@latest
```

## Anteprima {#preview}

Apri una presentazione e *Markdown: Open Preview to the Side*. Una presentazione è un file Markdown
il cui front matter dice `slides: true`:

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

Vengono mostrati in anteprima anche i file con `marp: true`, a meno che non sia installato Marp for VS Code:
entrambe le estensioni renderebbero lo stesso documento. Lo decide l'impostazione
`goSlides.preview.marpDocuments`: `auto` (il valore predefinito), `always` o
`never`.

L'anteprima è go-slides stesso, compilato in WebAssembly e caricato
dall'estensione: nessun binario, e funziona su [vscode.dev](https://vscode.dev). Legge
ciò che la presentazione usa dall'area di lavoro — `slides.yaml` o `slides.hcl` accanto alla
presentazione, i file di [kinds]({{< relref "kinds" >}}) che nominano, i file che un
[template]({{< relref "templates" >}}) include, il CSS dei temi — e solo dalla
cartella di lavoro della presentazione.

## Sincronizzazione dello scorrimento {#scroll-sync}

Dalla v0.2.0 dell'estensione (go-slides v0.7.0), l'editor e l'
anteprima scorrono insieme, in entrambe le direzioni, come per il Markdown semplice: l'
anteprima segue la slide e il blocco sotto il cursore dell'editor, e
scorrere l'anteprima sposta l'editor. Le impostazioni sono quelle di VS Code,
`markdown.preview.scrollPreviewWithEditor` e
`markdown.preview.scrollEditorWithPreview`.

Le righe sono tracciate attraverso i template. Una slide generata da un `{{ range }}`
segue la riga del suo template, e ciò che un `{{ include }}` porta dentro
segue la riga dell'include. Una presentazione divisa da `headingDivider` si sincronizza sui
suoi titoli.

## Presentare {#presenting}

Dalla v0.4.0 dell'estensione (go-slides v0.9.0), una presentazione può essere presentata
senza uscire da VS Code. Con la presentazione nell'editor, dalla Tavolozza dei comandi
o dal menu `…` dell'editor:

| Comando | Apre |
|---|---|
| *go-slides: Open Overview* | un pannello con tutte le slide in una griglia. Una slide scelta mostra il suo sorgente nell'editor; `Esc` chiude il pannello |
| *go-slides: Start Presenter View* | le slide, e accanto la vista del relatore: il passo successivo, le note del relatore e un timer |

I pannelli stanno accanto all'editor e si seguono a vicenda: cambia slide in uno
e gli altri cambiano con esso. Seguono anche le tue modifiche alla presentazione. Il tasto `p`
nelle slide apre la vista del relatore.

Per presentare su un altro schermo, sposta lì il pannello delle slide con *View: Move
Editor into New Window*.

Ogni pannello esegue la presentazione in un frame in sandbox (senza `allow-same-origin`) con una
policy `default-src 'none'`: il frame carica immagini, media, stili e font
e nient'altro, ed esegue solo gli script che go-slides ha incorporato nella pagina,
indicati per hash. Uno script scritto dalla presentazione non viene eseguito.

## Problemi, completamento, hover {#problems-completion-hover}

- **Problemi**: errori di template, YAML e [HCL]({{< relref "hcl" >}}), sulla
  riga in cui si trovano, nella presentazione o nel file di kinds da cui provengono.
- **Completamento**: chiavi delle direttive nel front matter e nei commenti `<!-- -->`
  (con `_` per una direttiva spot), i valori di `kind` e `_kind` (con ciò che
  imposta ciascun kind), `theme`, `size`, `paginate`, `math`, `headingDivider`; dentro
  `{{ }}`, le funzioni dei template e i campi dei dati.
- **Hover**: cosa fa una direttiva, cosa imposta un kind, gli argomenti di una funzione
  di template.

## Esportazione {#export}

*go-slides: Export Slide Deck...* (presente anche nella barra del titolo dell'editor) scrive HTML,
PDF vettoriale, PPTX e ODP modificabili, PNG o JPEG con il comando `slides`, come
farebbe `slides -o deck.pdf deck.md`; vedi [Riga di comando]({{< relref "cli" >}}).
Richiede un'area di lavoro attendibile. Se `slides` non viene trovato, l'estensione propone di
eseguire il `go install` qui sopra in un terminale.

Dalla v0.5.0 dell'estensione (go-slides v0.14.0), il PDF che scrive è
**accessibile**: con `slides` v0.12.0 o successivo, il PDF vettoriale ha i tag per
gli screen reader e dichiara PDF/UA-1, con il `title` e il `lang` della presentazione e una
sezione per slide. `slides --pdf-tags=false` scrive il file senza tag. Aggiorna
il comando con il `go install` qui sopra per ottenerlo.

## Impostazioni {#settings}

| Impostazione | Predefinito | |
|---|---|---|
| `goSlides.cliPath` | `""` | il comando `slides`; vuoto: `slides` nel `PATH`. Un'impostazione della macchina: un'area di lavoro non può definirla |
| `goSlides.exportType` | `pdf` | il formato che la finestra di esportazione propone per primo |
| `goSlides.exportAutoOpen` | `true` | apre il file esportato |
| `goSlides.html` | `false` | consente tutto l'HTML grezzo invece della lista consentita di Marp; rispettata solo in un'area di lavoro attendibile |
| `goSlides.preview.marpDocuments` | `auto` | mostra in anteprima le presentazioni `marp: true`: `auto`, `always`, `never` |

## Sicurezza {#security}

La presentazione e l'area di lavoro non sono attendibili. L'HTML grezzo segue la lista consentita di Marp
a meno che `goSlides.html` sia attivo in un'area di lavoro attendibile (un `slides.yaml` con
`html: true` non lo attiva nell'anteprima), e l'anteprima di VS Code non esegue alcuno
script della presentazione. L'anteprima legge solo i file della cartella di lavoro della presentazione,
non segue alcun collegamento simbolico, e i template non vedono alcuna variabile d'ambiente.
L'esportazione esegue un programma, quindi attende l'attendibilità dell'area di lavoro, e il programma
non viene mai nominato dall'area di lavoro.

## Prestazioni {#performance}

Sulla presentazione di benchmark da 100 slide di go-slides, in VS Code 1.142 su un Apple M4 Max,
con Marp for VS Code 3.6.1 misurato nella stessa sessione:

| | go-slides | Marp for VS Code |
|---|---:|---:|
| attivazione (caricamento del modulo WebAssembly) | 189 ms | — |
| prima anteprima | 52 ms | 78 ms |
| anteprima dopo una modifica (mediana) | **19 ms** | 23–24 ms |
| `.vsix` | 7,5 MB | 36,3 MB |

Le cifre sono per l'estensione v0.3.0 (go-slides v0.8.0) su una presentazione da 100 slide,
digitando un carattere nella slide 50, da round alternati con un carico a 1 minuto
inferiore a 8. La prima anteprima è più veloce di quella di Marp, e l'anteprima dopo una
modifica è ora più veloce anche di quella di Marp.

### Nella v0.8.0 e nell'estensione v0.3.0 {#in-v080-and-extension-v030}

go-slides v0.8.0 conserva le slide rese di ogni presentazione aperta e, dopo una modifica,
rende di nuovo solo quelle il cui sorgente o stato ereditato è cambiato (direttive, numero
di pagina, totale, id dei titoli, riferimenti dei link, righe sorgente); l'output è
lo stesso di un rendering completo, byte per byte. Digitare nella slide 50 di 100 rende di nuovo una
slide: 2,5 ms in nativo, 11,5 ms in WebAssembly. L'anteprima di VS Code poi trasforma
solo il DOM della slide cambiata, e il test di fumo dell'estensione digita un
carattere nella slide 50 tra rendering cronometrati, per entrambe le estensioni. La
vista del relatore dell'estensione si sincronizza tramite la finestra padre (`postMessage`).

### Nella v0.7.1 e nell'estensione v0.2.2 {#in-v071-and-extension-v022}

go-slides v0.7.1 mette in cache i temi analizzati, i fogli di stile impacchettati e il codice evidenziato
tra un rendering e l'altro, marca le righe della sincronizzazione senza procedere byte per byte,
e riutilizza i propri writer: un rendering della presentazione da 100 slide alloca 5,7 MB, contro 52,7 MB
per l'estensione v0.2.0. Con esso, l'estensione v0.2.2 elimina l'override `GOGC=400`.
A titolo di confronto, l'estensione v0.2.1 impiegava 72 ms per la prima anteprima e 44 ms
dopo una modifica con il `GOGC` predefinito (36 ms con `GOGC=400`), e la v0.2.0 impiegava
47–51 ms.

L'estensione v0.2.2 ha misurato 52 ms per la prima anteprima e 24 ms dopo una modifica,
alla pari con i 23 ms di Marp.
