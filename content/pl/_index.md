---
title: "Dokumentacja go-slides"
linkTitle: "Strona główna"
type: docs
cascade:
  type: docs
description: "Prezentacje zgodne z Marp w czystym Go (CGO=0), z interpolacją szablonów Go i kinds slajdów."
---

**`slides` zamienia prezentację Marp w gotowy pokaz, w czystym Go.** Czyta ten
Markdown, który czyta [Marp](https://marp.app) — dyrektywy Marpit, `---` między
slajdami, tła `![bg]`, notatki prelegenta, motywy `default`, `gaia` i
`uncover` — i dodaje dwie rzeczy, których Marp nie ma:

- **[szablony Go]({{< relref "templates" >}})** wewnątrz prezentacji, wyliczane
  przed parsowaniem Markdownu, dzięki czemu prezentacja może sama generować
  swoje slajdy;
- **[kinds (rodzaj slajdu)]({{< relref "kinds" >}})**: slajd mówi, czym jest
  (`title`, `section`…), zamiast powtarzać, jak wygląda.

{{< callout type="warning" >}}
**v0.16.0.** Te strony opisują
[go-slides/slides v0.16.0](https://github.com/go-slides/slides/releases/tag/v0.16.0)
i są publikowane pod adresem `/docs/0.16/`.
{{< /callout >}}

## Co potrafi dzisiaj {#what-it-does-today}

| | |
| --- | --- |
| `slides build` (domyślne) | prezentacja jako jedna strona HTML: każdy slajd to `<section>` w ramce SVG, skalowana do okna, z CSS motywu i przeglądarką obsługiwaną z klawiatury |
| `slides build --pdf`, `--images png`, `--image jpeg` | prezentacja jako PDF (jedna strona na slajd, ze spisem treści) albo jako obrazy PNG/JPEG, bez przeglądarki: renderuje [go-webengine](https://github.com/go-webengine), PDF zapisuje [go-pdfkit](https://github.com/go-pdfkit) — od v0.4.0 jako wektory, z zaznaczalnym tekstem: [PDF wektorowy]({{< relref "cli#vector-pdf" >}}) |
| matematyka i kod | `$…$` i `$$…$$` składane przez go-tex/math, bloki kodu podświetlane przez go-rouge, w HTML, obrazach i PDF: [Matematyka i kod]({{< relref "math-and-code" >}}) |
| emoji | `:smile:` i 😄 stają się osadzonymi obrazami Twemoji, więc nie potrzebują sieci i rysują się w każdym formacie wyjściowym: [Emoji]({{< relref "emoji" >}}) (od v0.4.0) |
| `slides build --pptx` | prezentacja jako plik PowerPoint z notatkami prelegenta i metadanymi prezentacji: jeden obraz na slajd w v0.4.0, jak w marp-cli; **edytowalny** od v0.5.0 — pola tekstowe, kształty, obrazy, osadzone czcionki ([Edytowalne PPTX i ODP]({{< relref "cli#editable-pptx-and-odp" >}})) |
| `slides build --odp` | to samo jako prezentacja OpenDocument (`.odp`) |
| dopasowane nagłówki | `# <!-- fit --> Title` wypełnia szerokość slajdu w jednej linii, na stronie i w każdym eksporcie |
| `slides expand` | prezentacja jako zwykły Markdown Marp: szablony wyliczone, kinds zastąpione dyrektywami punktowymi, które reprezentują |
| `slides notes` | notatki prelegenta |
| `slides kinds`, `slides themes` | to, co prezentacja może nazwać |
| `slides.yaml` | wspólne kinds, motywy, dane, polityka HTML i język dla katalogu prezentacji |
| HCL | nagłówek `---hcl`, `slides.hcl`, pliki kinds i danych w HCL2, z wyrażeniami i funkcjami: [HCL]({{< relref "hcl" >}}) |

Różnica względem
marp-cli to droga: on renderuje obrazy i PDF, sterując headless Chrome;
tutaj robią to biblioteki Go.

## Strony {#pages}

- [Pierwsze kroki]({{< relref "getting-started" >}})
- [Galeria]({{< relref "gallery" >}}): dwanaście przykładowych prezentacji do otwarcia w Playground
- [Zgodność z Marp]({{< relref "marp" >}})
- [Szablony]({{< relref "templates" >}})
- [Kinds]({{< relref "kinds" >}})
- [HCL]({{< relref "hcl" >}})
- [Matematyka i kod]({{< relref "math-and-code" >}})
- [Emoji]({{< relref "emoji" >}})
- [Wiersz poleceń]({{< relref "cli" >}})
