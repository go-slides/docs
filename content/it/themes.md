---
title: "Pacchetti di tema"
weight: 42
description: "Un tema in un solo file: un archivio .gsta.xz con un theme.hcl, il suo CSS, font, risorse e kinds, controllato come input non attendibile."
tags: [themes, hcl]
---

Un **pacchetto di tema** è un solo file, `<name>.gsta.xz` (Go-Slides Theme
Archive): un archivio tar (ustar o pax) compresso con xz, con `theme.hcl` alla
radice e i file che dichiara. Porta con sé il CSS di un tema, i font, i loghi e
gli sfondi e, se si vuole, kinds di slide, con un nome, una versione e una
licenza, in modo che un tema possa essere distribuito e installato. I pacchetti
di tema esistono dopo la v0.17.1.

## Il file theme.hcl {#the-themehcl-file}

```hcl
# theme.hcl
theme {
  name        = "acme"
  version     = "1.0.0"                 # semantic version
  description = "Example theme"
  authors     = ["go-slides contributors"]
  license     = "MIT"                   # SPDX expression
  homepage    = "https://go-slides.github.io"
  extends     = "default"               # the base theme, built in or installed
  min_slides  = "0.18.0"                # oldest go-slides it works with
}

css "theme.css" {}                      # in order; sha256 = "…" is optional

font "Roboto Mono" {
  weight  = 400                         # or "100 900"
  style   = "normal"
  file    = "fonts/RobotoMono-Regular.ttf"
  license = "OFL-1.1"                   # required for a font
}

asset "logo" { file = "assets/logo.svg" }

kind "acme-title" {                     # optional slide kinds
  class    = "lead"
  paginate = "false"
}
```

`theme.hcl` usa la sintassi nativa di HCL, solo con valori letterali: nessuna
funzione, nessuna variabile, niente `file()` né `env()`. Un nome è fatto di
lettere minuscole, cifre e trattini, e non è `default`, `gaia` né `uncover`.
`extends` indica il tema di base; ogni blocco `css` aggiunge un foglio di stile,
in ordine; ogni blocco `font` e `asset` indica un file del pacchetto; ogni
blocco `kind` è un [kind]({{< relref "kinds" >}}). Un `sha256` su un blocco
`css`, `font` o `asset` viene controllato quando il pacchetto viene letto.

Il CSS fa riferimento a font e risorse con il loro percorso nel pacchetto,
relativo al file CSS: `background: url(assets/logo.svg)`. Quando il pacchetto
viene caricato, ognuno di questi riferimenti diventa un URL `data:`, così la
pagina è autosufficiente, e un blocco `font` diventa una `@font-face`. Un
pacchetto può contenere anche i file `LICENSE`, `NOTICE` e `README`; qualsiasi
altro file che `theme.hcl` non dichiara viene rifiutato.

## Impacchettare, installare e usare {#packing-installing-and-using}

```
slides theme pack dir       write dir/theme.hcl and its files as <name>.gsta.xz
slides theme verify file    run every check on a package
slides theme install file   check it, then install it for the current user
slides theme list           the installed packages
slides theme remove name
```

I pacchetti installati si trovano in `go-slides/themes`, nella cartella di
configurazione dell'utente (`os.UserConfigDir()`: `~/Library/Application Support`,
`$XDG_CONFIG_HOME` o `~/.config`, `%AppData%`). In una presentazione:

```markdown
---
theme: acme
---
```

trova un pacchetto installato, dopo i temi integrati e `--theme-set`.
`--theme acme.gsta.xz` e `--theme-set acme.gsta.xz` usano direttamente
un file, senza installarlo; vedi [Riga di comando]({{< relref "cli" >}}).

I kinds di un pacchetto si applicano con `--theme file.gsta.xz`, con
`--theme name` per un pacchetto installato, o con `theme: name` nel front matter
di una presentazione; i kinds propri della presentazione prevalgono su di essi.

L'`extends` di un pacchetto nomina il suo tema di base, cercato in quest'ordine: integrato (o indicato con `--theme-set`), poi installato, poi un `<base>.gsta.xz` fratello nella stessa directory del file del pacchetto, in modo che una famiglia di temi (una base, diversi piccoli figli) possa essere provata direttamente dai suoi file. L'errore nomina i tre luoghi cercati. Una catena di pacchetti è lunga al massimo 8 e un ciclo viene rifiutato (`theme.meta`).

## Un pacchetto è un input non attendibile {#a-package-is-untrusted-input}

Un pacchetto di tema può arrivare da qualsiasi parte, quindi ogni lettore
(`pack`, `verify`, `install`, `--theme`, una compilazione) lo controlla allo
stesso modo e lo rifiuta con un errore che ha un codice (`theme.archive`,
`theme.limit`, `theme.meta`, `theme.css`, `theme.svg`, `theme.font`,
`theme.checksum`, `theme.version`):

- **Archivio.** Un solo flusso xz, il cui dizionario è limitato e la cui
  finestra cresce con i dati. Nel tar: niente `..`, percorso assoluto, lettera
  di unità o barra rovesciata in un nome; solo file regolari e cartelle (niente
  collegamento simbolico o fisico, dispositivo o FIFO); niente nome duplicato
  (confrontato senza distinguere maiuscole e minuscole); solo intestazioni ustar
  e pax (niente estensione GNU, niente file sparso), un'intestazione pax di al
  massimo 8 KiB. Le dimensioni sono quelle lette, non quelle dichiarate da
  un'intestazione.
- **Limiti.** Al massimo 16 MiB di archivio, 256 voci, 8 MiB per voce, 24 MiB
  decompressi, un dizionario xz di 16 MiB, un flusso decompresso di al massimo
  100 volte la dimensione dell'archivio (di 64 KiB o più), `theme.hcl` di al
  massimo 64 KiB e 32 livelli di annidamento, 1 MiB di CSS per file, 4 MiB per
  font o immagine.
- **CSS.** Niente `@import`, niente `@import-theme` (usa `extends`), niente
  `image-set()`, `src()` o altra funzione che carichi una risorsa, niente
  escape con barra rovesciata fuori da una stringa, niente `</style`. Un
  `url()` è un `#fragment`, un'immagine o un font `data:`, o un file dichiarato
  del pacchetto, mai uno che ne esca.
- **SVG.** Niente `script`, `foreignObject`, frame, embed, `set`, DOCTYPE né
  gestori di eventi; un `href` è un `#fragment`; niente `url()` che esca dal
  file.
- **Font e immagini** vengono analizzati: un font deve essere un font, un'immagine
  deve corrispondere alla sua estensione, e ciascuno ha un limite di dimensione.

La cartella `testdata/themes/acme` del repository è un piccolo esempio.
