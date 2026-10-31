---
title: "Szablony"
weight: 30
description: "Interpolacja text/template z Go w prezentacji: wartości, generowane slajdy, numery stron, dołączanie plików."
tags: [templates]
---

Prezentacja jest szablonem Go [`text/template`](https://pkg.go.dev/text/template),
wyliczanym **przed** parsowaniem Markdownu. Wszystko, co ma język szablonów
Go — akcje, potoki, `if`, `range`, `with`, zmienne — działa
w prezentacji.

## Dane {#data}

Kropka szablonu przechowuje klucze front matter, potem klucze jego mapy
`data:`, potem dane podane w wierszu poleceń lub w `slides.yaml`
— każde nadpisuje poprzednie.

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

Z wiersza poleceń:

| Flaga | Daje |
| --- | --- |
| `--data file.yaml` | plik danych YAML lub JSON (można powtarzać) |
| `--set key=value` | jedną wartość (można powtarzać) |

## Generowanie slajdów {#generating-slides}

Całe ciało jest jednym szablonem, więc `range`, który zapisuje `---`, zapisuje slajdy:

```markdown
{{ range .releases }}
---

# {{ .version }}

{{ range .changes }}- {{ . }}
{{ end }}
{{ end }}
```

## Strona i suma {#page-and-total}

`{{ page }}` i `{{ total }}` to numer slajdu i liczba slajdów.
Są znane dopiero po podziale prezentacji, więc są wtedy uzupełniane
— wygenerowana prezentacja liczy wygenerowane przez siebie slajdy.

```markdown
<!-- footer: "{{ page }} / {{ total }}" -->
```

## Include {#include}

`{{ include "parts/agenda.md" }}` wstawia inny plik, względem
prezentacji.

## Kod zostaje nietknięty {#code-is-left-alone}

Bloki kodu i fragmenty kodu w tekście są odkładane na bok przed uruchomieniem
szablonu i przywracane bez zmian, więc slajd o szablonach Go może pokazać
`{{ .Name }}` bez ucieczki znaków.

## Wyłączanie {#turning-it-off}

`--no-template` czyta prezentację jako zwykły Marp: `{{` jest tekstem.
