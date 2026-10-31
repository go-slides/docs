---
title: "HCL"
weight: 45
description: "HCL2 wszędzie tam, gdzie prezentacja przyjmuje YAML: nagłówek w HCL, slides.hcl, pliki rodzajów i danych, z wyrażeniami i funkcjami."
tags: [hcl, configuration]
---

{{< callout type="warning" >}}
HCL jest w v0.3.0 ([go-slides/slides#13](https://github.com/go-slides/slides/pull/13)).
{{< /callout >}}

Wszędzie tam, gdzie `slides` czyta YAML, czyta też natywną składnię
[HCL2](https://github.com/hashicorp/hcl/blob/main/hclsyntax/spec.md) — dla
jej wyrażeń i funkcji.

| Gdzie | YAML | HCL |
| --- | --- | --- |
| front matter | `---` … `---` | `---hcl` … `---` |
| konfiguracja | `slides.yaml` | `slides.hcl`, szukany w pierwszej kolejności |
| pliki rodzajów | `--kinds kinds.yaml`, `kinds: kinds.yaml` | dowolny plik o nazwie `*.hcl` |
| pliki danych | `--data data.yaml` lub `.json` | dowolny plik o nazwie `*.hcl` |

Komentarze z dyrektywami (`<!-- class: invert -->`) pozostają w YAML, jak w Marp, a
ciało prezentacji pozostaje [szablonem Go]({{< relref "templates" >}}).

## Front matter w HCL {#a-front-matter-in-hcl}

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

- **Atrybuty** to klucze YAML pod tymi samymi nazwami: dyrektywy Marp
  (`theme`, `paginate`, `header`, `footer`, `class`, `backgroundColor`…),
  metadane (`title`, `author`…), `template`, `delims`, `kinds` (ścieżki
  plików) i `data`, obiekt.
- **Bloki `kind "<name>" { … }`** definiują [rodzaje]({{< relref "kinds" >}}),
  z dyrektywami lokalnymi i `extends`.
- **Wartości** trafiają do dyrektyw w postaci tekstowej, jaką nadaje im YAML —
  `true` to `"true"`, lista klas to lista — więc dyrektywa zachowuje się
  tak samo, którąkolwiek składnią ją ustawiono. `null` zostawia dyrektywę nieustawioną.
- **`data`** zachowuje typy: w szablonach liczba całkowita to `int`,
  inne liczby to `float64`, a listy i obiekty pozostają listami i mapami, więc
  `{{ add .quarter 1 }}` działa.
- Wartość HCL jest ostateczna: szablony Go nie przechodzą po niej drugi raz.

## Zmienne {#variables}

`data` to dane szablonu: własne `data` front matter, nałożone przez to z
`slides.hcl`, przez pliki `--data` i przez `--set`. `"${data.company}"` działa
wszędzie tam, gdzie zadziałałoby `{{ .company }}`.

Atrybuty front matter mogą odwoływać się do `data.*`, ale **nie do siebie
nawzajem**: `footer = title` jest błędem. Dyrektywy punktowe (`_class`…) nie są
identyfikatorami HCL; zapisz je w komentarzu.

## Funkcje {#functions}

| | |
| --- | --- |
| `page()`, `total()` | numer strony i numer ostatniej strony, jak `{{ page }}` i `{{ total }}`: stopka numeruje strony |
| `now()`, `timestamp()` | czas, RFC 3339, lokalny i UTC; `formatdate` go formatuje |
| `env(name)` | zmienna środowiskowa, `""`, gdy nieustawiona |
| `file(path)` | plik tekstowy UTF-8 o rozmiarze do 1 MiB, względem prezentacji lub pliku konfiguracyjnego |
| liczby | `abs ceil floor log max min parseint pow signum` |
| napisy | `chomp format formatlist join lower regex regexall regexreplace replace split strlen strrev substr title trim trimprefix trimspace trimsuffix upper` |
| kolekcje | `chunklist coalesce coalescelist compact concat contains distinct element flatten keys length lookup merge range reverse slice sort values zipmap` |
| kodowania | `csvdecode jsondecode jsonencode` |
| daty | `formatdate timeadd` |

Wszystkie wiersze poza pierwszymi czterema to biblioteka standardowa
[go-cty](https://github.com/zclconf/go-cty). Żadna nie sięga do sieci.

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

Nieznany atrybut jest błędem, a nie cichym zignorowaniem. Plik rodzajów zawiera
bloki `kind`; jego pozostałe atrybuty są ignorowane, więc `slides.hcl` może służyć
jako taki plik. Atrybuty pliku danych to klucze danych:

```hcl
# data.hcl
company = "ACME"
year    = 2026
motto   = trimspace(file("motto.txt"))
```

## Błędy {#errors}

Błędy to diagnostyki HCL, wskazujące miejsce w samym pliku prezentacji:
`footer = data.compnay` w czwartej linii prezentacji daje

```
deck.md:4:14: Unsupported attribute: This object does not have an attribute named "compnay".
```

## Marp i `slides expand` {#marp-and-slides-expand}

Marp również czyta `---hcl` jako początek front matter — jego reguła
front matter otwiera się na każdej pierwszej linii zaczynającej się od trzech myślników — i,
skoro blok nie jest mapą YAML, nie stosuje z niego niczego: prezentacja nadal
renderuje się w Marp, bez swoich dyrektyw. Test przypina to do wyniku samego Marp Core.

`slides expand` zapisuje front matter z powrotem jako YAML, więc rozwinięta prezentacja to
zwykły Marp i renderuje się tak, jak `slides` renderuje oryginał:

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

## Licencje {#licences}

Obsługa HCL linkuje [hashicorp/hcl](https://github.com/hashicorp/hcl)
(MPL-2.0, używane bez modyfikacji jako biblioteka) i
[zclconf/go-cty](https://github.com/zclconf/go-cty) (MIT).
