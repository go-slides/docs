---
title: "Riga di comando"
weight: 50
description: "Ogni comando e flag di slides, e il file di configurazione slides.yaml."
tags: [cli]
---

{{< callout type="warning" >}}
Questa pagina descrive la v0.15.0.
{{< /callout >}}

```text
slides [deck.md] [flags]
slides [command] [flags] deck.md
```

## Comandi {#commands}

| Comando | Cosa fa |
| --- | --- |
| `build` | scrive la presentazione come HTML, o come PDF o immagini (vedi sotto). **Il predefinito**: `slides deck.md` equivale a `slides build deck.md` |
| `expand` | stampa la presentazione come semplice Markdown Marp: template espansi, kinds sostituiti dalle direttive che rappresentano |
| `notes` | stampa le note del relatore |
| `kinds` | elenca i kinds (il tipo di una slide) che la presentazione può usare e cosa impostano; `--json` (nella v0.6.0) li scrive come `[{"name", "extends", "directives": [{"key", "value"}]}]`, con `extends` risolto |
| `themes` | elenca i temi |
| `version` | stampa la versione (lo fa anche `slides --version`) |
| `completion` | stampa uno script di completamento per la shell: `bash`, `zsh`, `fish` o `powershell` |

`slides <command> --help` stampa i flag di un comando. I flag possono stare prima
o dopo la presentazione: `slides deck.md -o out.html` e `slides -o out.html
deck.md` sono equivalenti. I flag lunghi prendono due trattini (`--html`); `-o` è l'unico
corto. `build` scrive `deck.html` accanto a `deck.md` salvo diversa indicazione di
`-o`.

## Flag {#flags}

| Flag | Comandi | Cosa fa |
| --- | --- | --- |
| `-o`, `--output file` | `build`, `expand`, `notes` | file di output; `-` per lo standard output |
| `-I`, `--input-dir dir` | `build` | converte i file Markdown sotto `dir`, ricorsivamente (nella v0.6.0); vedi sotto |
| `-P`, `--parallel n` | `build` | presentazioni compilate contemporaneamente (nella v0.6.0): 5 per impostazione predefinita, al massimo il numero di core; 1 le compila una dopo l'altra |
| `-w`, `--watch` | `build` | compila di nuovo quando cambia una presentazione, un tema, un file di kinds o di dati o la configurazione, e ricarica le pagine aperte (nella v0.6.0); vedi sotto |
| `-s`, `--server dir` | `build` | serve una cartella, convertendo le sue presentazioni su richiesta (nella v0.6.0); vedi sotto |
| `-p`, `--preview` | `build` | apre il risultato nel browser predefinito (nella v0.6.0): le pagine HTML con `--watch` (implicito), il server con `--server` |
| `--host addr` | `build` | l'indirizzo su cui ascolta `--server`, `127.0.0.1` per impostazione predefinita (nella v0.6.0) |
| `--allowed-host name` | `build` | il nome host pubblico di un reverse proxy a cui `--server` risponde, ripetibile (nella v0.14.0); vedi [Self-hosting]({{< relref "self-hosting" >}}) |
| `--template bespoke\|bare` | `build` | il template della pagina HTML (nella v0.6.0): `bespoke`, il predefinito, con lo script di presentazione, il controller a schermo e le transizioni, oppure `bare`, le slide impilate (`--bare`) |
| `--bespoke.osc`, `--bespoke.progress`, `--bespoke.transition` | `build` | le opzioni del template bespoke, con i valori predefiniti di marp-cli (nella v0.6.0): controller attivo, **barra di avanzamento disattivata** (prima era sempre disegnata), transizioni attive; `--bespoke.osc=false` ne disattiva una |
| `-c`, `--config-file file`, `--no-config-file` | `build` | il file di configurazione di marp-cli, JSON o YAML (nella v0.6.0); vedi sotto. `--config` è il file proprio di `slides` |
| `--browser`, `--browser-path`, `--browser-protocol`, `--browser-timeout` | `build` | accettati e ignorati con un avviso (nella v0.6.0): non c'è nessun browser |
| `--pptx-editable`, `--no-parallel` | `build` | le grafie di marp-cli per `--pptx` e `-P 1` (nella v0.6.0) |
| `--config file` | tutti | file di configurazione; per impostazione predefinita `slides.yaml` accanto alla presentazione, se esiste |
| `--no-config` | tutti | non legge alcun file di configurazione |
| `--theme-set path` | tutti | un file CSS di tema, o una cartella di essi (ripetibile) |
| `--kinds file` | tutti | un file YAML di kinds (ripetibile) |
| `--data file` | tutti | un file YAML o JSON di dati per i template (ripetibile) |
| `--set key=value` | tutti | un valore per i template (ripetibile) |
| `--no-template` | tutti | non espande i template Go |
| `--strict-yaml` | tutti | analizza le direttive come YAML rigoroso, come fa Marpit semplice (Marp Core le analizza in modo permissivo) |
| `--no-emoji` | tutti | lascia `:shortcodes:` ed emoji Unicode come testo invece di disegnarli come immagini (dalla v0.4.0: [Emoji]({{< relref "emoji" >}})) |
| `--html` | `build` | consente tutto l'HTML grezzo; per impostazione predefinita, la lista consentita di Marp Core |
| `--lang code` | `build` | il `lang` della pagina quando la presentazione non ne imposta uno |
| `--bare` | `build` | impila le slide, senza lo script di presentazione |
| `--pdf` | `build` | scrive un PDF, una pagina per slide, con un indice — testo e forme in vettoriale (scelto anche da `-o file.pdf`) |
| `--pdf-raster` | `build` | scrive il PDF con ogni pagina come un'unica immagine della slide, come faceva la v0.3.0 |
| `--notes` | `build` | scrive le note del relatore come file di testo, con le slide separate da `---` come fa marp-cli (nella v0.6.0; scelto anche da `-o file.txt`). Il comando `notes` le stampa con un altro layout |
| `--pdf-notes` | `build` | aggiunge le note del relatore di ogni slide alla sua pagina PDF come annotazione di nota, come fa marp-cli (nella v0.6.0). Richiede un PDF quando nient'altro dice cosa scrivere |
| `--pdf-outlines`, `--pdf-outlines.pages`, `--pdf-outlines.headings` | `build` | i segnalibri del PDF (nella v0.6.0): "Page N" per slide e i titoli annidati sotto di essa. Attivi per impostazione predefinita qui, a differenza di marp-cli; `=false` ne disattiva uno. Vedi sotto |
| `--images png\|jpeg` | `build` | scrive un'immagine per slide: `deck.001.png`, … |
| `--image png\|jpeg` | `build` | scrive la prima slide come immagine (scelto anche da `-o file.png` o `-o file.jpg`) |
| `--pptx` | `build` | scrive un file PowerPoint con le note del relatore (scelto anche da `-o file.pptx`): modificabile dalla v0.5.0, un'immagine per slide nella v0.4.0 |
| `--pptx-raster` | `build` | scrive un file PowerPoint con una sola immagine a tutta slide per slide, come fa marp-cli (dalla v0.5.0); `--image-scale` vale 2 per impostazione predefinita |
| `--odp` | `build` | scrive una presentazione OpenDocument con le note del relatore (scelto anche da `-o file.odp`): modificabile dalla v0.5.0 |
| `--odp-raster` | `build` | scrive una presentazione OpenDocument con un'immagine per slide (dalla v0.5.0) |
| `--jpeg-quality n` | `build` | qualità JPEG, da 1 a 100 (predefinito 85) |
| `--image-scale n` | `build` | pixel del dispositivo per pixel CSS per immagini e PDF (predefinito 1); 2 dà 2560×1440 per il 16:9, con testo e vettori disegnati a quella densità, non ingranditi |
| `--allow-local-files` | `build` | permette a immagini e PDF di leggere file locali, solo da sotto la cartella della presentazione; rifiutato per impostazione predefinita, come con marp-cli |
| `--theme name\|file.css` | `build` | sostituisce il tema della presentazione (nella v0.6.0): un nome di tema, o un file CSS, che diventa il tema qualunque `@theme` dichiari; una cartella viene rifiutata, usa `--theme-set` |
| `--title`, `--description`, `--author`, `--keywords`, `--url`, `--og-image` | `build` | i metadati della pagina (nella v0.6.0), al posto delle direttive `title`, `description`, `author`, `keywords`, `url` e `image` della presentazione; vedi sotto |

## Più presentazioni, e una cartella {#several-decks-and-a-directory}

Nella v0.6.0 `build` accetta più presentazioni, una cartella (le presentazioni al suo interno,
trovate ricorsivamente) o un pattern, e ne compila fino a `--parallel` contemporaneamente.
`-` è lo standard input; a differenza di marp-cli, una pipe non viene letta se non lo dice `-`.

Un pattern è letto come lo legge il glob di marp-cli (`**` e `{a,b}` nella v0.7.0;
mettilo tra virgolette perché la shell lo lasci a `slides`): `*`, `?`, `[a-z]` e
`{md,txt}` all'interno di un nome, `**` per un numero qualsiasi di cartelle, così
`'docs/**/*.md'` prende le presentazioni dell'albero. Un carattere jolly non raggiunge un nome
che inizia con un punto, nulla viene cercato in `node_modules`, le maiuscole contano
(`*.md` non trova `UPPER.MD`), e tutto ciò che un pattern individua viene convertito,
qualunque sia la sua estensione. Un pattern che non individua nulla è un errore.

- **`-I dir`** trova i file Markdown (`.md`, `.mdown`, `.markdown`,
  `.markdn`, in minuscolo) sotto `dir`, come fa marp-cli: `node_modules` e
  le cartelle che iniziano con un punto non vengono esplorate. Con `-o out`, ogni presentazione viene scritta al suo
  percorso sotto `dir`, riprodotto sotto `out`: `dir/sub/b.md` dà
  `out/sub/b.html` (`b.pdf`, `b.001.png`, `b.txt`, … qualunque cosa dicano i flag).
  Senza `-o` gli output vanno accanto alle loro presentazioni. Nominare presentazioni oltre a
  `-I` è un errore, e lo è anche `-o file` per più presentazioni (`-o -` è consentito).
- **`-P n`** compila `n` presentazioni contemporaneamente. Il rendering di una presentazione usa già i
  core, quindi `n` è limitato al loro numero; con `-o -` le presentazioni sono compilate una dopo
  l'altra. I messaggi di una presentazione restano insieme, e una presentazione che fallisce non
  ferma le altre: lo stato di uscita è 1 e il messaggio elenca i fallimenti.

## Transizioni e pagina di presentazione {#transitions-and-the-presenting-page}

Nella v0.6.0 la direttiva `transition` di marp-cli funziona: `<!-- transition:
fade 1s -->`, un nome e una durata facoltativa, per la slide e per quelle successive,
oppure `<!-- _transition: cover -->` per una sola. I nomi integrati sono quelli di marp-cli:
`clockwise`, `counterclockwise`, `cover`, `coverflow`, `cube`, `cylinder`,
`diamond`, `drop`, `explode`, `fade`, `fade-out`, `fall`, `flip`, `glow`,
`implode`, `in-out`, `iris-in`, `iris-out`, `melt`, `overlap`, `pivot`, `pull`,
`push`, `reveal`, `rotate`, `slide`, `star`, `swap`, `swipe`, `swoosh`, `wipe`,
`wiper` e `zoom`; `none` la disattiva. La pagina la esegue con la
**View Transition API** del browser quando lo spettatore lascia la slide che la possiede, avanzando,
e quando vi ritorna; i browser senza l'API cambiano slide come prima,
e uno spettatore che chiede di ridurre il movimento ottiene una semplice dissolvenza. Un tema può portare
keyframe propri, chiamati `marp-transition-NAME` (o `marp-outgoing-` e
`marp-incoming-`, e `-backward`), che vengono usati prima di quelli integrati.

La pagina ha il **controller a schermo** — precedente, pagina, successiva, schermo intero —
che compare quando si muove il puntatore o un tasto e si nasconde dopo un momento
(`--bespoke.osc=false` lo rimuove), e, con `--bespoke.progress`, una barra di
avanzamento. **La barra di avanzamento è ora disattivata per impostazione predefinita**, come in marp-cli; prima
era sempre disegnata.

### Vista del relatore {#presenter-view}

Nella v0.7.0 la pagina ha la vista del relatore di marp-cli. Premi `p`, o l'ultimo
pulsante del controller, e si apre una finestra con la slide corrente, il passo
successivo, le tue note del relatore, la pagina, l'ora e un timer (fai clic sul timer
per azzerarlo; `+` e `-` ridimensionano le note; la barra tra le slide e le
note si può trascinare). Le due finestre si seguono a vicenda in entrambe le direzioni: i
tasti freccia, un clic sul passo successivo o su Previous/Next in una delle due spostano
l'altra, un frammento alla volta.

Funziona come quella di marp-cli, letta dai suoi sorgenti 4.5.1: la finestra è
`?view=presenter&sync=KEY` dello stesso file, il passo successivo è la stessa pagina di
`?view=next` in un frame, e le finestre condividono la loro posizione tramite
`localStorage` (`bespoke-marp-sync-KEY`), quindi non richiede server né
rete, e non funziona in un browser privo di `localStorage` (il
pulsante lo dice). Le note sono testo: `&`, `<` e `>` hanno l'escape, quindi una nota
non può mai essere markup. Sono nella pagina, nascoste, perché la vista del relatore le
prelevi, il che le tiene fuori da ciò che uno screen reader legge al pubblico.
Le pagine `--bare` non hanno nulla di tutto ciò. La vista aggiunge circa 10 KB a una pagina.

### Vista panoramica {#overview-view}

Nella v0.8.0, premi `o` (o `Escape`: marp-cli alterna con entrambi), il pulsante
a griglia del controller, o il contatore di pagina della vista del relatore, e ogni slide
appare intera in una griglia sopra la pagina. I tasti freccia scelgono una slide (su e giù
per riga), `Home` e `End` vanno alle estremità, e le altre finestre seguono;
`Enter`, `Space` o un clic scelgono una slide e, sopra la pagina, chiudono la
panoramica. Sopra la vista del relatore una scelta la lascia aperta (il relatore sceglie
la slide da mostrare, e continua a guardare). Il pulsante "Close overview", `o` e
`Escape` la chiudono. È quella di marp-cli: la pagina di nuovo in un frame a
`?view=overview&sync=KEY`, che condivide la posizione tramite `localStorage`, senza
rete. Le slide nella griglia sono attivabili con il focus ed etichettate "Slide N", e la
panoramica chiusa è inerte, così gli utenti di tastiera e screen reader non vi finiscono mai
dentro. Le due viste insieme aggiungono circa 15 KB a una pagina.

### File di configurazione {#configuration-files}

marp-cli legge le proprie opzioni da un file nella cartella di lavoro, e lo fa anche
`slides`, quando il file è dati: `.marprc` (YAML, anche JSON è YAML),
`.marprc.json`, `.marprc.yaml`, `.marprc.yml`, `.config/marprc*`, o la chiave
`marp` di `package.json`; `--config-file` (`-c`) nomina qualsiasi file JSON o YAML,
`marp.config.json` per uno. Si cerca solo nella cartella di lavoro, non
nelle sue superiori, come fa la ricerca di marp-cli. Un file JavaScript o TypeScript
(`marp.config.js` ...) è un programma, che `slides` non esegue: lo dice e
prosegue. Le opzioni sono quelle della documentazione di marp-cli che hanno senso
qui: `title`, `description`, `author`, `keywords`, `url`, `ogImage`, `lang`,
`theme`, `themeSet`, `html`, `allowLocalFiles`, `template`, `bespoke`,
`inputDir`, `output`, `parallel`, `pdf`, `pptx`, `notes`, `pdfNotes`,
`pdfOutlines`, `images`, `image`, `imageScale`, `jpegQuality`, `watch`,
`server`, `preview`; i percorsi sono relativi al file. La riga di comando prevale.
`engine`, `options` e `browser*` richiedono il motore di Marp e un browser, e vengono
ignorate con un avviso, come i flag `--browser*`: gli script scritti per
marp-cli continuano a funzionare. (`slides.yaml` e `slides.hcl`, con i kinds e i
dati dei template, sono separati e letti come prima.)

## Watch, server e anteprima {#watch-serve-and-preview}

Nella v0.6.0, con le rotte e il protocollo di ricaricamento in tempo reale di marp-cli (letti dai suoi
sorgenti 4.5.1 e confrontati con un `marp -s` in esecuzione).

**`slides --server dir`** (`-s`) serve la cartella su `http://127.0.0.1:8080/`
(la variabile d'ambiente `PORT` cambia la porta; una porta già in uso è un
errore che lo dice):

| Richiesta | Risposta |
| --- | --- |
| `/deck.md` | la presentazione come HTML, convertita su richiesta |
| `/deck.md?pdf`, `?pptx`, `?png`, `?jpg` (o `?jpeg`), `?notes` (o `?txt`) | la presentazione convertita come chiede la query; il PPTX è un allegato |
| una cartella | il suo `index.md` o `PITCHME.md` come pagina, altrimenti un indice con un link PDF e uno PPTX per ogni presentazione |
| qualsiasi altro file | il file |
| una presentazione che non si converte | 503 con l'errore |

Una pagina apre un WebSocket verso `/.__marp-cli-watch-notifier__/<id>`, dove l'id è
lo SHA-256 del percorso della presentazione; il server invia `ready`, poi `reload` quando la
presentazione, o qualsiasi cosa sotto la cartella, cambia. Una cartella richiesta senza
la barra finale viene reindirizzata ad essa (marp-cli serve lì il suo indice, il che
rompe i link relativi della pagina).

**`--watch`** (`-w`) compila le presentazioni, poi ricompila quella che cambia
e, quando cambia un file `--theme-set`, un file di kinds o di dati o la configurazione,
tutte. Le pagine che scrive aprono un socket sul loopback (porta
37717, o la successiva libera) e si ricaricano quando la loro presentazione viene compilata di nuovo. Una presentazione
che non si compila viene segnalata e continua a essere osservata. **`--preview`** (`-p`) apre
il browser predefinito sul server, o sulle pagine HTML appena compilate; non c'è
una finestra propria, dato che `slides` non ha un browser da pilotare.

Il server è un modo di entrare nella macchina, quindi è più rigoroso di quello di marp-cli, che
ascolta su ogni interfaccia e non controlla nulla:

- ascolta su **127.0.0.1**; `--host` lo cambia, con un avviso, e un
  indirizzo jolly non può più controllare l'header `Host`;
- l'header **`Host`** deve essere un nome di loopback (`localhost`, `127.0.0.1`,
  `[::1]`, `*.localhost`) o l'host su cui ascolta, il che impedisce il DNS rebinding;
- l'**`Origin`** di un WebSocket deve essere quello del server stesso; i socket della modalità watch
  accettano anche l'origine `null` di una pagina aperta da un file; una richiesta da
  un altro sito che non sia una navigazione (uno script, un'immagine, un `fetch`) viene
  rifiutata;
- un percorso viene ripulito, i suoi collegamenti simbolici risolti, e il risultato deve essere
  dentro la cartella: `..` e un collegamento che ne esce vengono rifiutati, così come
  `?pdf` di una presentazione esterna; i nomi che iniziano con un punto (`.git`, `.env`) non vengono
  serviti;
- solo `GET` e `HEAD`; poche conversioni contemporanee; al massimo 4096 cartelle
  osservate, e nessuna in `node_modules` o sotto un punto.

Dietro un reverse proxy, `--allowed-host name` (nella v0.14.0, ripetibile)
aggiunge il nome pubblico del proxy agli `Host` e `Origin` accettati; i controlli
restano attivi e il server resta sul loopback.
[Self-hosting]({{< relref "self-hosting" >}}) lo mette dietro Caddy, con
HTTPS e un accesso.

## Metadati della pagina {#page-metadata}

La riga di comando prevale sulla presentazione, come in marp-cli, dove i flag vengono
aggiunti alla presentazione come un ultimo commento-direttiva. Regole lette da marp-cli
4.5.1 e fissate da test rispetto al suo output (nella v0.6.0):

- Una presentazione che non imposta alcun `title` prende il testo del suo primo titolo di livello
  più basso, com'è scritto (`# **Bold** title` dà `**Bold** title`). Senza
  titoli non c'è `<title>`.
- Le `keywords` sono divise alle virgole, ripulite e deduplicate; un elenco nel front matter
  viene preso così com'è. Come marp-cli, la pagina porta il tag `keywords` dalla
  seconda parola chiave in poi, non per una sola.
- Un `url` non valido (non assoluto) è ignorato con un avviso, e il valore precedente
  viene mantenuto; uno valido scrive anche `<link rel="canonical">` e `og:url`.
- `author` e `description` arrivano con i loro gemelli Open Graph; `og:image:alt`
  e la Twitter card grande richiedono sia un titolo sia un'immagine.

## Immagini, PDF, PPTX e ODP {#images-pdf-pptx-and-odp}

`-o` con un'estensione `.pdf`, `.pptx`, `.odp`, `.png` o `.jpg` sceglie il formato, come fa marp-cli.
Le cartelle di output vengono create. Un PDF ha una pagina per slide, alla
dimensione della slide, e un indice delle slide e dei loro titoli.

### Note e indice del PDF {#pdf-notes-and-outline}

Nella v0.6.0 il PDF accetta le due opzioni di marp-cli, e vengono confrontate con
i PDF che marp-cli 4.5.1 scrive per la stessa presentazione:

- `--pdf-notes` mette le note del relatore di ogni slide sulla sua pagina come annotazione di
  nota: un'icona Note gialla in basso a sinistra il cui contenuto sono le
  note (più commenti uniti da una riga vuota) e il cui titolo è l'`author`
  della presentazione. `--pdf-notes` da solo richiede un PDF, come in marp-cli.
- L'indice contiene un segnalibro in grassetto "Page N" per slide e, sotto di esso, i
  titoli della slide annidati per livello (`--pdf-outlines.pages` e
  `--pdf-outlines.headings` scelgono le due parti). **L'indice è attivo per impostazione predefinita
  qui e disattivato in marp-cli.** `slides` ne ha sempre scritto uno; costa poche
  centinaia di byte e rende navigabile una presentazione lunga, quindi il valore predefinito resta.
  `--pdf-outlines=false`, o `--no-pdf-outlines` di marp-cli, lo disattiva, e
  `=false` su `.pages` o `.headings` elimina quella parte. Un segnalibro di titolo salta
  alla sua slide; marp-cli scorre anche fino al titolo.

Il rendering è fatto da [go-webengine](https://github.com/go-webengine/engine)
e il PDF è scritto da [go-pdfkit](https://github.com/go-pdfkit/pdfkit);
i font richiesti dai temi integrati (Lato, Roboto Mono) provengono da
[go-opentype/fonts](https://github.com/go-opentype/fonts) e sono incorporati, così
nessun rendering tocca la rete. Non c'è Chrome.

`--image-scale 2` esegue il rendering a due pixel del dispositivo per pixel CSS. Le immagini locali vengono
rifiutate se non si passa `--allow-local-files`, e anche allora vengono letti solo i file sotto
la cartella della presentazione: una presentazione proveniente da altrove non può raggiungere `~/.ssh`.
Sfondi multipli e divisi (`![bg left:40%]`) vengono resi come i livelli di Marp.

`--pptx` scrive la presentazione con il `WritePresentation` di [go-ooxml](https://github.com/go-ooxml),
con le note del relatore e i metadati della presentazione. Dalla
v0.5.0 le slide sono modificabili (vedi sotto); `--pptx-raster` scrive una sola immagine a tutta
slide per slide, come fa marp-cli, a una scala che per impostazione predefinita è 2 e che
`--image-scale` sostituisce. LibreOffice apre il risultato. È stato provato solo LibreOffice;
PowerPoint e Keynote no.

`--odp` scrive una presentazione OpenDocument, tramite il `WritePresentation`
di go-odf, con le sue note del relatore; `--odp-raster` scrive un'immagine
per slide.

### PPTX e ODP modificabili {#editable-pptx-and-odp}

Dalla v0.5.0 le slide di un PPTX o di un ODP sono **modificabili** — il PPTX di marp-cli
è un'immagine per slide, e il suo sperimentale `--pptx-editable` converte
un PDF tramite LibreOffice. Qui ogni slide è disegnata a partire dal suo layout, quello
che disegnano le esportazioni PNG e PDF:

- ogni riga di testo è una casella di testo nella posizione calcolata, nei font
  con cui la slide è stata impaginata — **incorporati nel file**, così appare
  identica su una macchina senza Lato o Inter (verificato con LibreOffice 25.8) —
  con grassetto, corsivo, colori, spaziatura tra le lettere e link cliccabili;
- sfondi, bordi e riquadri arrotondati sono forme; i gradienti sono riempimenti
  a gradiente nel PPTX e immagini nell'ODP (un gradiente ODF ha due
  colori);
- le immagini vengono memorizzate una volta sola, per quante slide le mostrino;
- ciò che una slide non può esprimere — un elemento sotto `opacity`, un `filter`, un'
  ombra sfocata — è un'immagine di quell'elemento, con il suo testo mantenuto
  sopra;
- dalla v0.15.0 una formula, un'emoji o un'immagine SVG è un gruppo di forme vettoriali
  (`p:grpSp` nel PPTX, `draw:g` nell'ODP) che si sposta come un tutt'uno. Un SVG che richiede un ritaglio o un gradiente radiale resta un'immagine.

Rese di nuovo da LibreOffice, una slide è entro lo 0,04–2 % dei pixel dell'esportazione
PNG. `--pptx-raster` e `--odp-raster` mantengono un'immagine per slide.

### PDF vettoriale {#vector-pdf}

Dalla v0.4.0 il PDF è disegnato in vettoriale: go-webengine dipinge la slide
impaginata nel PDF invece che in pixel. Il testo è testo, nei font con cui la
slide è stata impaginata, incorporati come sottoinsieme con una mappa `/ToUnicode` — può essere
selezionato, cercato e copiato. Sfondi, bordi e gradienti sono tracciati
e sfumature, un'immagine mostrata in molte slide viene incorporata una volta sola, e i link sono
cliccabili. La presentazione di benchmark da 100 slide pesa 758 KiB, contro 837 KiB di
marp-cli e 10 MiB del PDF a immagini della v0.3.0, e viene scritta in 0,3 s.

Un elemento senza forma vettoriale qui — sotto `opacity` < 1, con un `filter`
o una rotazione, un `box-shadow` sfocato, un gradiente la cui opacità sfuma — è
disegnato come bitmap di quel solo elemento, con il suo testo sovrapposto in modo invisibile perché
resti selezionabile. `--pdf-raster` scrive un'immagine
per pagina, come faceva la v0.3.0.

Dalla v0.15.0, anche formule, emoji e immagini SVG sono vettori (go-webengine v0.22.0): ogni disegno
distinto è un form XObject PDF, dipinto dove viene usato, nitido a qualsiasi zoom.

### PDF accessibile e con tag {#accessible-tagged-pdf}

Dalla v0.12.0 il PDF vettoriale ha i tag. Porta la propria struttura logica, che
uno screen reader legge in ordine, come legge una pagina web:

| Nella presentazione | Nel PDF |
| --- | --- |
| una slide | una sezione (`Sect`) |
| `#` … `######` | `H1` … `H6` |
| un paragrafo, un elenco, una tabella | `P`; `L` con l'etichetta e il corpo di ogni voce; `Table`, celle d'intestazione con il loro scope |
| un link | un `Link`, con la sua area cliccabile, il suo testo come descrizione |
| un'immagine | una `Figure` con il suo testo alternativo; `![](…)` senza testo alternativo è decorazione |
| `$…$`, `$$…$$` | una `Formula`, con il suo sorgente TeX come `/Alt` e `/ActualText` (i suoi glifi sono contorni: pdftotext e mupdf non li estraggono, i lettori che seguono la struttura sì) |
| codice | `Code`, senza i numeri di riga né i cerchi delle chiamate |
| `<span lang="fr">` | uno span in quella lingua |
| sfondi, bordi, numeri di pagina | artefatti, che un lettore salta |
| note del relatore (`--pdf-notes`) | annotazioni nella struttura |

Il file dichiara PDF/UA-1, lo standard per il PDF accessibile. Porta il `title` della
presentazione, che il visualizzatore mostra, e il suo `lang` (`--lang` quando la presentazione
non ne imposta uno, `en` altrimenti).

Una presentazione verificata con veraPDF supera PDF/UA-1, a meno che i suoi stessi titoli violino
l'ordine dello standard. PDF/UA richiede che il primo titolo sia un `H1` e che
nessun livello sia saltato (`#` poi `###`). Solo l'autore della presentazione può correggerlo.
Anche il PDF di Chrome per la stessa presentazione, che scrive marp-cli, ha i tag, ma non supera
sette regole di PDF/UA, non ha una sezione per slide né formule.

Aggiungere i tag alla presentazione di benchmark da 100 slide costa il 4% in dimensione (788 contro 757 KiB)
e nessun tempo misurabile. `--pdf-tags=false` scrive il PDF senza tag.

## `slides.yaml` {#slidesyaml}

Un file di configurazione accanto alla presentazione si applica ad essa, così una cartella di presentazioni
ne condivide uno:

```yaml
kinds: kinds.yaml        # a mapping of kinds, or file(s) holding one
themes:                  # theme CSS files or directories
  - themes/
data:                    # template data
  company: Example Inc.
html: false              # allow all raw HTML
lang: en
```

Kinds: i `kinds:` della presentazione prevalgono su quelli del file e della riga di comando.
Dati: i `data` del file e della riga di comando prevalgono sul front
matter della presentazione, così una presentazione può essere compilata con dati diversi.
