---
title: "Playground"
weight: 15
description: "Scrivi una presentazione nel browser e guarda le slide seguirti: go-slides compilato in WebAssembly, con completamento, guida ed esempi."
tags: [playground, browser]
---

**[Apri il playground →](https://go-slides.github.io/playground/)**

Il playground è go-slides stesso, v0.14.0, compilato in WebAssembly ed
eseguito nel tuo browser. Scrivi una presentazione a sinistra; le slide seguono a
destra. Nulla viene inviato a un server: né la presentazione, né un link di condivisione.

## Cosa offre {#what-it-has}

- **Anteprima in tempo reale** di ogni slide, resa fuori dal thread della pagina (un Web
  Worker), quindi la digitazione non deve mai aspettarla. Fai clic su una slide per portare l'editor
  alla sua riga; sposta il cursore e l'anteprima scorre fino alla sua slide.
- **Slideshow**: la presentazione vera e propria, con la tastiera, il
  controller a schermo e le transizioni.
- **Vista del relatore**: **Present** (o `p` nello slideshow) la apre in una
  seconda finestra: la slide corrente, il passo successivo, le note del relatore (i
  `<!-- comments -->` della presentazione), un orologio e un timer. I tasti freccia, Spazio,
  Home e Fine funzionano in entrambe le finestre, e ciascuna segue l'altra, anche
  quando modifichi la presentazione. Se la finestra non si apre, consenti i pop-up per
  il sito. La vista panoramica (`o`) esiste solo nella pagina scaricata.
- **Completamento** durante la digitazione: chiavi delle direttive (quelle spot con `_` nei commenti),
  i valori di `kind`, `theme`, `size`, `paginate`, e le funzioni dei template e i campi dei dati
  dentro `{{ }}`.
- **Guida** per ciò che si trova sotto il cursore: una direttiva, un kind (il tipo di una slide: cosa imposta),
  una funzione di template. Il riquadro della guida rimanda anche a queste pagine.
- **Problemi** sulla loro riga: gli errori di template, YAML e HCL sono evidenziati
  nell'editor ed elencati sotto di esso, e l'anteprima conserva l'ultima presentazione
  che è stata resa.
- **Esempi**: una presentazione di benvenuto, template che generano slide, kinds,
  un front matter HCL, temi e direttive spot, e matematica, immagini e
  sfondi (`![bg]`, un `![bg left:40%]` diviso, filtri).
- **File**: il pannello a sinistra dell'editor è il progetto della presentazione, organizzato
  come lo leggono il [comando]({{< relref "cli" >}}) e
  l'[estensione di VS Code]({{< relref "vscode" >}}): `deck.md`,
  `slides.yaml`, `themes/`, `kinds/`, `media/`.
  - Fai clic su un file per modificarlo. L'editor lo colora secondo il suo linguaggio
    (Markdown, YAML, HCL, JSON, CSS) e ne segnala i problemi.
  - Clic destro (o Maiusc+F10) per Nuovo file, Nuova cartella, Rinomina (F2),
    Elimina (chiede conferma) e Nuovo tema / Nuovo kind. Trascina un file su una cartella
    per spostarlo. La tastiera segue il pattern ad albero WAI-ARIA, compresa
    la ricerca digitando.
  - **Project ▸ New theme** crea `themes/my-gaia.css` (oppure da default,
    uncover): un `@import` del tema integrato e le regole che aggiungi. Entra
    nell'elenco dei temi. Anche **New kind** e **New data file** partono da
    modelli. Ognuno è elencato in `slides.yaml`, così `slides build deck.md`
    lo trova.
  - **Project ▸ Download project (.zip)** salva tutto. Decompresso,
    `slides build deck.md` lo compila. **Open project** riprende uno `.zip` o una
    cartella. Un archivio che esce dalla propria cartella (`..`, percorsi assoluti),
    contiene più di 256 file o 8 MiB in un file, o si espande più di 100
    volte viene rifiutato.
  - Tutto viene conservato in questo browser insieme alla bozza. Un progetto contiene fino
    a 64 MiB, e la riga sotto i file dice quanto ne è usato.
- **Media**: trascina o incolla un'immagine, un video o un suono sull'editor e finisce
  in `media/`, e la presentazione lo nomina: `![](media/photo.png)`. Trascina
  file su una cartella del pannello per metterli lì. La presentazione resta piccola e
  leggibile, e nulla viene caricato. La sandbox dell'anteprima li carica
  senza essere aperta (la pagina le passa i byte). **Save as ▸ HTML
  page** li inserisce nella pagina. Funzionano anche le immagini `https:`, ma il browser
  le preleva dal loro server, che vede così il tuo indirizzo. L'anteprima
  non invia alcun referrer.
- **Riquadri ridimensionabili**: il pannello dei file, l'editor, le miniature delle slide, l'
  anteprima e la guida sono separati da barre che puoi trascinare (il puntatore
  diventa un cursore di ridimensionamento sopra una di esse). Il riquadro accanto alla barra prende lo spazio,
  gli altri mantengono la loro dimensione. Fai doppio clic su una barra, o premi Invio su di essa,
  per chiudere il riquadro e riaprirlo. Da tastiera, F6 passa tra
  le barre, i tasti freccia ne spostano una di 16 px, Home e Fine portano il riquadro
  che la precede alla dimensione minima e massima, ed Esc torna all'editor.
  La scheda View continua a disattivare i riquadri. Le dimensioni impostate sono conservate in questo
  browser, con la bozza; le dimensioni predefinite (190, 168 e 280 px, con l'editor
  e l'anteprima che si dividono il resto 6 a 5) non vengono mai memorizzate.
- **Formato**: la scheda Format modifica il Markdown al posto tuo, e ogni pressione è
  un passo di Annulla.
  - **Testo**: Grassetto, Corsivo, Barrato e Codice in linea agiscono sulla selezione
    o sulla parola sotto il cursore e si alternano (Ctrl+B, Ctrl+I, Ctrl+Maiusc+X e
    Ctrl+E; Cmd su Mac). Un pulsante appare premuto quando il cursore è dentro
    i suoi segni. Markdown non ha il sottolineato: Underline (Ctrl+U) racchiude il testo in
    `<u>…</u>`, un tag che la lista consentita HTML predefinita di Marp mantiene, quindi funziona con
    l'HTML grezzo disattivato.
  - **Paragrafo**: Titolo da 1 a 3 (Ctrl+Alt+1..3), elenchi puntati (Ctrl+Maiusc+8),
    numerati (Ctrl+Maiusc+7) e una citazione agiscono sulle righe selezionate e si alternano.
    La linea orizzontale è il tag `<hr>`: `---`, `***` e `___` inizierebbero ciascuno
    una nuova slide in Marp.
  - **Tabella**: una griglia sceglie colonne e righe. Con il cursore in una tabella
    aggiungi o rimuovi una riga o una colonna e allinei la colonna (`:--`, `:-:`,
    `--:`); la tabella viene riscritta allineata dopo ogni modifica.
  - **Testo della slide** a sinistra, al centro o a destra scrive un blocco `<style scoped>` nella
    slide corrente. Marp non ha una sintassi di allineamento, e un attributo `style` non
    supera la lista consentita HTML predefinita, mentre il blocco scoped sì, quindi
    funziona con l'HTML grezzo disattivato e compare nell'anteprima, nella pagina HTML e nel PDF.
  - **Link** (Ctrl+K) scrive `[text](url)` con l'url selezionato.
- **Evidenziazione** delle chiavi del front matter a qualsiasi profondità, in YAML (anche
  nelle mappe `{a: 1}`) e in HCL.
- Selettore di **tema**: default, gaia, uncover, i temi del progetto
  (`themes/*.css`) e i temi definiti dalla presentazione (`/* @theme name */` in un
  `<style>`). Scrive `theme:` nel front
  matter, YAML o HCL, così un link condiviso lo conserva, e segue il front
  matter quando lo modifichi.
- L'anteprima **mantiene la posizione** mentre digiti: vengono sostituite solo le slide che
  sono cambiate, e un cambio di tema riapre l'anteprima sulla
  stessa slide.
- **Link di condivisione**: la presentazione viaggia nel link stesso (nella sua parte `#`,
  compressa), quindi aprire il link apre la presentazione. I link trasportano fino a
  64 KiB compressi, e solo la presentazione: quando il progetto contiene di più
  (temi, kinds, media), Share lo dice e propone lo `.zip`.
- **Save as**:
  - la pagina HTML, la stessa che scrive `slides deck.md`;
  - un **PDF** (vettoriale, con testo selezionabile, con indice);
  - una presentazione **PowerPoint** o **OpenDocument** modificabile.

  Sono prodotti nel tuo browser, dallo stesso codice di `slides --pdf`,
  `--pptx` e `--odp`: il PDF della presentazione di benchmark da 100 slide è
  identico byte per byte a quello del comando. La prima esportazione scarica il motore
  di esportazione (circa 5 MiB, una volta; il browser lo conserva fino al prossimo aggiornamento),
  con una riga di avanzamento, e lo compila man mano che arriva. Poi scarica solo
  i font che le tue slide disegnano. Con un collegamento a 3 MB/s il primo PDF della
  presentazione di benvenuto è pronto in meno di 3 secondi. La presentazione da 100 slide richiede poi circa
  3 secondi. Un'immagine
  `https:` è inclusa solo se il suo server lo consente (CORS); quelle `data:`
  lo sono sempre.
- **La tua bozza viene conservata**:
  - La presentazione è salvata in questo browser mentre digiti (IndexedDB, solo per questo sito)
    e ritorna quando ricarichi. Nulla viene inviato altrove. I
    file di testo del progetto (temi, kinds, dati, `slides.yaml`) sono conservati
    allo stesso modo, anche se chiudi la scheda subito dopo aver digitato; i media sono
    memorizzati nel momento in cui vengono trascinati.
  - Un link condiviso apre la sua presentazione e conserva la tua bozza: **Draft ▸ Restore my
    draft** la riporta indietro.
  - Un esempio, **Draft ▸ Reset to the example** (che chiede conferma) e
    un'altra scheda non la cancellano mai: **Draft ▸ Restore replaced** la
    riporta indietro.
  - **Draft ▸ Forget my drafts** cancella tutto ciò che è memorizzato.

Ogni esportazione rende il progetto che mostra l'anteprima: il suo `slides.yaml`,
temi, kinds, file di dati e media, con i media dentro il documento.

Le immagini come file (slides `--images`) non vengono prodotte nel browser: usa il
[comando]({{< relref "cli" >}}) o l'
[estensione di VS Code]({{< relref "vscode" >}}).

## Sicurezza {#security}

La presentazione è mostrata in un iframe in sandbox senza l'origine della pagina: non può
raggiungere né il playground né la sua memoria, non può inviare moduli e non naviga
da nessuna parte. La sua pagina ha una Content Security Policy che esegue solo i propri
script. L'HTML grezzo nella presentazione segue la lista consentita di Marp, come fa `slides` per
impostazione predefinita. Un rendering che dura più di 5 secondi (un template che va in ciclo)
viene interrotto.

La finestra del relatore mantiene quella sandbox. È aperta dal playground
stesso e mostra la presentazione in un iframe in sandbox allo stesso modo. I due
iframe non possono condividere la memoria, quindi il playground inoltra il numero della slide
tra loro tramite messaggi. Nessun frame che esegue la presentazione ottiene l'origine
del playground.

I media del progetto raggiungono i frame in sandbox come byte in un messaggio, e
ogni frame ne crea i propri URL `blob:`, nella propria origine. Un URL `blob:`
creato dal playground non vi si carica, e la policy dei frame
guadagna solo `blob:` per media e font.

Il playground può essere servito dalla tua macchina, dietro un accesso:
vedi [Self-hosting]({{< relref "self-hosting" >}}).

## Dimensione e velocità {#size-and-speed}

Il modulo è di 8,8 MiB come lo serve GitHub Pages (gzip; Pages non
serve brotli). Viene scaricato una volta con una barra di avanzamento, poi messo in cache dal
browser fino al prossimo deploy. Su un Apple M4 Max con Chrome 155, misurato
sulla pagina pubblicata, l'editor è pronto 2,1 s dopo l'apertura della pagina, di cui 1,5 s
sono il download, e le prime slide compaiono a 2,4 s. Una modifica alla
presentazione di benchmark da 100 slide raggiunge l'anteprima in circa 30 ms: go-slides
rende solo le slide cambiate, e solo quelle vengono sostituite
nell'anteprima.
