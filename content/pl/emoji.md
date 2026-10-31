---
title: "Emoji"
weight: 36
description: "Kody emoji i emoji Unicode stają się obrazami Twemoji, bez sieci: na stronie HTML, w obrazach, PDF, PowerPoint i OpenDocument."
tags: [emoji]
---

{{< callout type="warning" >}}
Emoji są w v0.4.0; v0.3.0 ich nie miała.
{{< /callout >}}

Zapisz emoji jako kod albo jako sam znak:

```markdown
# Launch :rocket: 🎉

Shortcodes :smile: :+1: :heart_eyes: and Unicode 😄 👍🏽 👩‍💻 🇫🇷 1️⃣.
```

Każde staje się obrazem [Twemoji](https://github.com/jdecked/twemoji), który
rysuje Marp Core, ze znacznikami, które zapisuje Marp Core:

```html
<img class="emoji" draggable="false" alt="😄" src="data:image/svg+xml;base64,…" data-marp-twemoji=""/>
```

## Bez sieci {#no-network}

`src` w Marp Core to adres URL na CDN, więc prezentacja pokazywana offline,
drukowana albo zamieniana na PDF na maszynie bez połączenia pokazuje zepsute
obrazy. go-slides umieszcza sam SVG w URI `data:`. Grafiki (4009 sztuk,
Twemoji 17.0.3, Emoji 17.0) są osadzone w programie z modułu w czystym Go
[go-icons/twemoji](https://github.com/go-icons/twemoji), 2 MB po kompresji,
czytane tylko wtedy, gdy prezentacja zawiera emoji.

Obrazy rysuje tak samo strona HTML, `--images`/`--image`, `--pdf`,
`--pptx` i `--odp`: go-webengine rasteryzuje SVG. Wszystkie 4009
grafik daje się sparsować i narysować jego rasteryzatorem SVG; sprawdza to test. Od
v0.5.0 tym rasteryzatorem jest [go-gfx/oksvg](https://github.com/go-gfx/oksvg)
v0.2.0: `clipPath`, `fill-rule: evenodd` i polecenia łuków rysują się teraz tak,
jak rysuje je Chrome, co zmieniło jedno Twemoji i 211 innych, a logo SVG nie są
już zniekształcane.

Reguła dopasowująca rozmiar emoji do tekstu (`img[data-marp-twemoji]`, kwadrat
o boku jednego `em`, wyrównany do linii bazowej) jest dodawana do motywu tak, jak
dodaje ją Marp Core, więc własny motyw też ją dostaje.

## Co jest konwertowane {#what-is-converted}

Reguły są z Marp Core; test porównuje je z HTML, który Marp Core 4.4.0
zapisuje dla prezentacji ze wszystkimi przypadkami.

| | |
| --- | --- |
| `:smile:`, `:+1:`, `:heart_eyes:` | 1903 nazwy z markdown-it-emoji, w tekście, nagłówkach, odnośnikach, tabelach, `header` i `footer`; nie w kodzie, adresach URL ani tekście alternatywnym obrazu |
| `😄`, `👩‍💻`, `🇫🇷`, `1️⃣`, `👍🏽` | emoji Unicode, w tekście **i w kodzie**: Marp Core przepisuje wyrenderowany HTML kodu |
| `✔`, `©`, `™`, `▶` | zostają tekstem: znaki, których domyślną prezentacją jest tekst, są emoji tylko z U+FE0F (`✔️`); wyjątkiem jest `❤` |
| `☺︎` | zostaje tekstem: U+FE0E po emoji prosi o zwykły glif |
| `12:30:45`, `:not_a_name:` | zostają tak, jak napisano |

`id` nagłówka i spis treści używają emoji, a nie nazwy kodu:
`# Launch :rocket:` ma id `launch-%F0%9F%9A%80`, jak w Marp.

Tam, gdzie go-slides różni się od Marp Core, konwertuje więcej: `✌️ ☝️ ✍️ 🏋️` i
kilka innych znaków domyślnie tekstowych, które emoji-regex w Marp Core zostawia
jako tekst, oraz garść sekwencji, które rysuje Twemoji, a Unicode ich nie wymienia.

## Wyłączanie {#turning-it-off}

`--no-emoji` (w kodzie `slides.Config{NoEmoji: true}`) zostawia kody i
emoji Unicode takim tekstem, jakim je napisano, co odpowiada `emoji: false`
w Marp Core.

## Autorstwo {#attribution}

Grafiki to Twemoji, Copyright 2020 Twitter, Inc i inni współtwórcy,
na licencji [CC-BY 4.0](https://creativecommons.org/licenses/by/4.0/),
kontynuowane przez projekt [jdecked/twemoji](https://github.com/jdecked/twemoji).
Prezentacja, która je pokazuje, powinna wskazać autorstwo, na przykład w stopce lub na
slajdzie końcowym. Tabela kodów pochodzi z markdown-it-emoji (MIT).
