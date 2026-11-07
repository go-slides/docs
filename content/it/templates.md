---
title: "Template"
weight: 30
description: "Interpolazione con text/template di Go in una presentazione: valori, slide generate, numeri di pagina, inclusioni."
tags: [templates]
---

La presentazione è un [`text/template`](https://pkg.go.dev/text/template) di Go,
valutato **prima** che il Markdown venga analizzato. Tutto ciò che il linguaggio
dei template Go offre — azioni, pipeline, `if`, `range`, `with`, variabili — funziona
in una presentazione.

## Dati {#data}

Il punto del template contiene le chiavi del front matter, poi le chiavi della mappa
`data:`, poi i dati forniti sulla riga di comando o in `slides.yaml`
— ciascuno sovrascrive il precedente.

```markdown
---
marp: true
title: Quarterly review
data:
  team: Platform
---

# {{ .title }}

Presented by the {{ .team }} team.
```

Dalla riga di comando:

| Flag | Fornisce |
| --- | --- |
| `--data file.yaml` | un file di dati YAML o JSON (ripetibile) |
| `--set key=value` | un valore (ripetibile) |

## Generare slide {#generating-slides}

L'intero corpo è un unico template, quindi un `range` che scrive `---` scrive slide:

```markdown
{{ range .releases }}
---

# {{ .version }}

{{ range .changes }}- {{ . }}
{{ end }}
{{ end }}
```

## Pagina e totale {#page-and-total}

`{{ page }}` e `{{ total }}` sono il numero della slide e il numero di
slide. Sono noti solo dopo che la presentazione è stata divisa, quindi vengono compilati
in quel momento — una presentazione generata conta le slide che ha generato.

```markdown
<!-- footer: "{{ page }} / {{ total }}" -->
```

## Include {#include}

`{{ include "parts/agenda.md" }}` inserisce un altro file, relativo alla
presentazione.

## Il codice non viene toccato {#code-is-left-alone}

I blocchi di codice delimitati e le porzioni di codice in linea vengono messi da parte prima che il template
venga eseguito e ripristinati alla lettera, quindi una slide sui template Go può mostrare
`{{ .Name }}` senza fare l'escape.

## Limiti {#limits}

L'esecuzione dei template ha un budget di passi, in modo che una presentazione i cui template annidano
`range` su grandi liste di dati non possa tenere occupato il renderer per minuti. Oltre
1.048.576 passi (iterazioni di ciclo e chiamate di template, condivisi con gli `include`
di una stessa espansione) o 10 s, la presentazione viene rifiutata con la diagnostica
`template.limit` invece di bloccarsi. Un passo costa circa 260 ns, quindi le presentazioni
ordinarie restano ben al di sotto del limite. Dalla v0.17.1.

## Disattivarlo {#turning-it-off}

`--no-template` legge la presentazione come Marp semplice: `{{` è testo.
