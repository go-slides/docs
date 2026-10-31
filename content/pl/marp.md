---
title: "Zgodność z Marp"
linkTitle: "Marp"
weight: 20
description: "Markdown, który czyta slides, to Markdown Marpit i Marp Core, a nie jego dialekt."
tags: [marp]
---

Prezentacja napisana dla [Marp](https://marp.app) jest prezentacją go-slides. Poniższe
reguły pochodzą z [Marpit](https://marpit.marp.app) i
[Marp Core](https://github.com/marp-team/marp-core); tam, gdzie go-slides
się różni, ta strona o tym mówi.

## Slajdy {#slides}

Linia pozioma — `---` w osobnej linii, jak dzieli Marpit — zaczyna
nowy slajd, podobnie jak nagłówek poziomu wskazanego przez dyrektywę
`headingDivider`. Front matter, między pierwszymi dwoma `---`, zawiera dyrektywy globalne
i dane prezentacji.

## Dyrektywy {#directives}

| Rodzaj | Gdzie | Dotyczy |
| --- | --- | --- |
| globalne | front matter lub `<!-- theme: gaia -->` w dowolnym miejscu | całej prezentacji: `theme`, `style`, `headingDivider`, `lang`… |
| lokalne | `<!-- paginate: true -->` | tego slajdu i następnych |
| punktowe | `<!-- _class: lead -->` (podkreślnik) | tylko tego slajdu |

Dyrektywy lokalne pochodzą z Marpit: `paginate`, `header`, `footer`, `class`,
`backgroundColor`, `backgroundImage`, `backgroundPosition`,
`backgroundRepeat`, `backgroundSize`, `color`. Dyrektywy są parsowane jako
luźny YAML, tak jak robi to Marp Core; `--strict-yaml` parsuje je tak, jak zwykły
Marpit.

## Obrazy {#images}

`![bg](image.jpg)` czyni obraz tłem slajdu, ze słowami kluczowymi Marp:
`bg contain`, `bg cover`, `bg 50%`, `bg left`/`bg right` dla tła dzielonego,
kilka obrazów `bg` obok siebie oraz filtry obrazów
(`blur`, `sepia`, `grayscale`…). Rozmiary obrazów w tekście (`w:200px`, `h:`) są
czytane tak, jak czyta je Marp.

## Wokół slajdu {#around-the-slide}

- **Nagłówek i stopka**: dyrektywy `header:` i `footer:`, w Markdownie.
- **Numeracja stron**: `paginate: true`.
- **Notatki prelegenta**: komentarz HTML, który nie jest dyrektywą, jest notatką.
  `slides notes` je wypisuje.
- **Listy stopniowane**: lista zapisana z `*` (lub `1)`) pojawia się element po elemencie,
  jak w Marp.
- **`<style scoped>`**: CSS tylko dla jednego slajdu.
- **Surowy HTML**: domyślnie lista dozwolonych z Marp Core; `--html` dopuszcza całość.
- **Emoji**: `:shortcodes:` i emoji Unicode stają się obrazami Twemoji, jak w
  Marp Core, ale osadzonymi zamiast ładowanymi z CDN: zobacz
  [Emoji]({{< relref "emoji" >}}).

## Motywy {#themes}

`default`, `gaia` i `uncover` są wbudowane, wybierane przez `theme:`. Własny
motyw to plik CSS z komentarzem Marp `/* @theme name */`,
podany przez `--theme-set` lub wymieniony w `themes:` w `slides.yaml`.

## Czym go-slides się różni {#where-go-slides-differs}

- **Szablony.** `{{ }}` w ciele to akcja szablonu Go (zobacz
  [Szablony]({{< relref "templates" >}})). Prezentacja, która pokazuje `{{` jako tekst
  poza blokiem kodu, potrzebuje `--no-template` albo tekstu z ucieczką tak, jak
  robią to szablony Go. Bloki kodu i fragmenty kodu pozostają nietknięte.
- **Kinds.** `kind:` i `_kind:` to dyrektywy, których Marp nie ma (zobacz
  [Kinds]({{< relref "kinds" >}})). `slides expand` zastępuje je
  dyrektywami, które reprezentują, więc rozwinięta prezentacja to zwykły Marp.
- **Formaty wyjściowe.** HTML, PNG, JPEG, PDF, PPTX i ODP, przez biblioteki Go, a nie
  przez przeglądarkę headless (znane ograniczenie opisuje [Wiersz poleceń]({{< relref "cli#images-pdf-pptx-and-odp" >}})).
  W v0.4.0, jak w marp-cli, PPTX zawiera obrazy; od
  v0.5.0 jego slajdy to edytowalne pola tekstowe i kształty, a `--pptx-raster` zachowuje
  obrazy.
- **Dopasowane nagłówki.** `# <!-- fit --> Title` jest obsługiwane: nagłówek wypełnia
  szerokość slajdu w jednej linii, na stronie i w każdym eksporcie.
- **Matematyka i kod.** `$…$` i `$$…$$` podlegają regułom ograniczników Marp Core i
  są składane przez go-tex/math; kod w blokach jest podświetlany przez go-rouge
  z nazwami klas highlight.js (zobacz [Matematyka i kod]({{< relref "math-and-code" >}})).
