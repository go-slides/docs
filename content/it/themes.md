---
title: "Pacchetti di tema"
weight: 42
description: "Un tema in un solo file: un archivio .slidestheme con un theme.hcl, il suo CSS, font, risorse e kinds, controllato come input non attendibile."
tags: [themes, hcl]
---

Un **pacchetto di tema** è un solo file, `<name>.slidestheme`: un archivio ZIP
(deflate) con `theme.hcl` alla radice e i file che dichiara. Porta con sé il
CSS di un tema, i font, i loghi e gli sfondi e, se si vuole, kinds di slide, con
un nome, una versione e una licenza, in modo che un tema possa essere
distribuito e installato. I pacchetti di tema esistono dopo la v0.17.1.

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
  min_slides  = "0.17.0"                # oldest go-slides it works with
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
slides theme pack dir       write dir/theme.hcl and its files as <name>.slidestheme
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
`--theme acme.slidestheme` e `--theme-set acme.slidestheme` usano direttamente
un file, senza installarlo; vedi [Riga di comando]({{< relref "cli" >}}).

I kinds di un pacchetto si applicano con `--theme file.slidestheme`, o con
`--theme name` per un pacchetto installato. Una presentazione che nomina il tema
solo nel suo front matter ne ottiene l'aspetto, non i kinds, e può nominarli con
`--kinds`.

## Un pacchetto è un input non attendibile {#a-package-is-untrusted-input}

Un pacchetto di tema può arrivare da qualsiasi parte, quindi ogni lettore
(`pack`, `verify`, `install`, `--theme`, una compilazione) lo controlla allo
stesso modo e lo rifiuta con un errore che ha un codice (`theme.archive`,
`theme.limit`, `theme.meta`, `theme.css`, `theme.svg`, `theme.font`,
`theme.checksum`, `theme.version`):

- **Archivio.** Niente `..`, percorso assoluto, lettera di unità o barra
  rovesciata in un nome; niente collegamento simbolico né file speciale; niente
  nome duplicato (confrontato senza distinguere maiuscole e minuscole); niente
  cifratura e nessun metodo se non stored e deflate. Le dimensioni sono quelle
  lette, non quelle dichiarate dall'archivio.
- **Limiti.** Al massimo 16 MiB di archivio, 256 voci, 8 MiB per voce, 24 MiB
  decompressi, un rapporto di compressione di 100 a 1 per una voce di 64 KiB o
  più, `theme.hcl` di al massimo 64 KiB e 32 livelli di annidamento, 1 MiB di
  CSS per file, 4 MiB per font o immagine.
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
