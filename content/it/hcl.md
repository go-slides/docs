---
title: "HCL"
weight: 45
description: "HCL2 ovunque una presentazione accetti YAML: un front matter HCL, slides.hcl, file di kinds e di dati, con espressioni e funzioni."
tags: [hcl, configuration]
---

{{< callout type="warning" >}}
HCL è nella v0.3.0 ([go-slides/slides#13](https://github.com/go-slides/slides/pull/13)).
{{< /callout >}}

Ovunque `slides` legga YAML, legge anche la sintassi nativa di
[HCL2](https://github.com/hashicorp/hcl/blob/main/hclsyntax/spec.md) — per
le sue espressioni e le sue funzioni.

| Dove | YAML | HCL |
| --- | --- | --- |
| front matter | `---` … `---` | `---hcl` … `---` |
| configurazione | `slides.yaml` | `slides.hcl`, cercato per primo |
| file di kinds (il tipo di una slide) | `--kinds kinds.yaml`, `kinds: kinds.yaml` | qualsiasi file chiamato `*.hcl` |
| file di dati | `--data data.yaml` o `.json` | qualsiasi file chiamato `*.hcl` |

I commenti-direttiva (`<!-- class: invert -->`) restano YAML, come in Marp, e il
corpo della presentazione resta un [template Go]({{< relref "templates" >}}).

## Un front matter in HCL {#a-front-matter-in-hcl}

```hcl
---hcl
theme    = "gaia"
paginate = true
title    = "Quarterly review"
data = {
  company = "ACME"
  quarter = 3
}
header = format("%s — Q%d", data.company, data.quarter)
footer = "${upper(data.company)} · ${page()}/${total()}"
date   = formatdate("D MMMM YYYY", now())

kind "chapter" {
  extends = "section"
  class   = ["lead", "invert"]
}
---

# {{ .title }}
```

- Gli **attributi** sono le chiavi YAML con gli stessi nomi: le direttive di Marp
  (`theme`, `paginate`, `header`, `footer`, `class`, `backgroundColor`…),
  i metadati (`title`, `author`…), `template`, `delims`, `kinds` (percorsi di
  file) e `data`, un oggetto.
- I blocchi **`kind "<name>" { … }`** definiscono i [kinds]({{< relref "kinds" >}}),
  con le direttive locali e `extends`.
- I **valori** arrivano alle direttive nelle forme stringa che dà loro YAML —
  `true` diventa `"true"`, un elenco di classi resta un elenco — così una direttiva si comporta
  allo stesso modo con qualunque sintassi sia stata impostata. `null` lascia una direttiva non impostata.
- **`data`** mantiene i suoi tipi: nei template un numero intero è un `int`,
  gli altri numeri un `float64`, mentre elenchi e oggetti restano liste e mappe, quindi
  `{{ add .quarter 1 }}` funziona.
- Un valore HCL è definitivo: i template Go non vengono eseguiti su di esso una seconda volta.

## Variabili {#variables}

`data` sono i dati del template: il `data` del front matter, sovrapposto da quello di
`slides.hcl`, dai file `--data` e da `--set`. `"${data.company}"` funziona
ovunque funzionerebbe `{{ .company }}`.

Gli attributi del front matter possono fare riferimento a `data.*` ma **non l'uno all'altro**:
`footer = title` è un errore. Le direttive spot (`_class`…) non sono identificatori
HCL; scrivile in un commento.

## Funzioni {#functions}

| | |
| --- | --- |
| `page()`, `total()` | il numero di pagina e quello dell'ultima pagina, come `{{ page }}` e `{{ total }}`: un piè di pagina viene numerato |
| `now()`, `timestamp()` | l'ora, RFC 3339, locale e UTC; `formatdate` la formatta |
| `env(name)` | una variabile d'ambiente, `""` se non impostata |
| `file(path)` | un file di testo UTF-8 di al massimo 1 MiB, relativo alla presentazione o al file di configurazione |
| numeri | `abs ceil floor log max min parseint pow signum` |
| stringhe | `chomp format formatlist join lower regex regexall regexreplace replace split strlen strrev substr title trim trimprefix trimspace trimsuffix upper` |
| collezioni | `chunklist coalesce coalescelist compact concat contains distinct element flatten keys length lookup merge range reverse slice sort values zipmap` |
| codifiche | `csvdecode jsondecode jsonencode` |
| date | `formatdate timeadd` |

Tutte tranne le prime quattro righe fanno parte della libreria standard di
[go-cty](https://github.com/zclconf/go-cty). Nessuna accede alla rete.

## slides.hcl {#slideshcl}

```hcl
kinds  = "kinds.hcl"          # or a list of files
themes = ["themes/acme.css"]
data   = { company = "ACME" } # or file paths
html   = false
lang   = "en"

kind "quote" {
  extends = "blank"
  footer  = "${data.company}"
}
```

Un attributo sconosciuto è un errore, non un'operazione ignorata in silenzio. Un file di kinds contiene
blocchi `kind`; i suoi altri attributi vengono ignorati, quindi `slides.hcl` può servire
da file di kinds. Gli attributi di un file di dati sono le chiavi dei dati:

```hcl
# data.hcl
company = "ACME"
year    = 2026
motto   = trimspace(file("motto.txt"))
```

## Errori {#errors}

Gli errori sono le diagnostiche di HCL, posizionate nel file stesso della presentazione:
`footer = data.compnay` alla quarta riga della presentazione produce

```
deck.md:4:14: Unsupported attribute: This object does not have an attribute named "compnay".
```

## Marp e `slides expand` {#marp-and-slides-expand}

Anche Marp legge `---hcl` come inizio di un front matter — la sua regola
sul front matter si apre su qualsiasi prima riga che inizi con tre trattini — e, dato che il blocco
non è una mappa YAML, non ne applica nulla: la presentazione viene comunque resa in Marp,
ma senza le sue direttive. Un test lo verifica rispetto all'output di Marp Core.

`slides expand` riscrive il front matter come YAML, quindi la presentazione espansa è
Marp semplice e viene resa come `slides` rende l'originale:

```yaml
---
theme: "gaia"
paginate: true
title: "Quarterly review"
header: "ACME — Q3"
footer: "ACME · 1/2"
date: "19 October 2026"
---
```

## Licenze {#licences}

Il supporto HCL collega [hashicorp/hcl](https://github.com/hashicorp/hcl)
(MPL-2.0, usato senza modifiche come libreria) e
[zclconf/go-cty](https://github.com/zclconf/go-cty) (MIT).
