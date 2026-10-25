---
title: "Emoji"
weight: 36
description: "Emoji-Shortcodes und Unicode-Emoji werden zu Twemoji-Bildern, ohne Netzwerk: auf der HTML-Seite, in Bildern, PDF, PowerPoint und OpenDocument."
tags: [emoji]
---

{{< callout type="warning" >}}
Emoji gibt es ab v0.4.0; v0.3.0 hatte keine.
{{< /callout >}}

Schreiben Sie ein Emoji als Shortcode oder als das Zeichen selbst:

```markdown
# Launch :rocket: 🎉

Shortcodes :smile: :+1: :heart_eyes: and Unicode 😄 👍🏽 👩‍💻 🇫🇷 1️⃣.
```

Jedes wird zu dem [Twemoji](https://github.com/jdecked/twemoji)-Bild, das Marp Core
zeichnet, mit dem Markup, das Marp Core schreibt:

```html
<img class="emoji" draggable="false" alt="😄" src="data:image/svg+xml;base64,…" data-marp-twemoji=""/>
```

## Kein Netzwerk {#no-network}

Das `src` von Marp Core ist eine URL auf einem CDN. Ein Foliensatz, der offline gezeigt, gedruckt
oder auf einem Rechner ohne Verbindung in ein PDF umgewandelt wird, zeigt deshalb kaputte Bilder.
go-slides legt das SVG selbst in einen `data:`-URI. Die Grafiken (4.009 Stück,
Twemoji 17.0.3, Emoji 17.0) sind aus dem reinen Go-Modul
[go-icons/twemoji](https://github.com/go-icons/twemoji) in das Programm eingebettet, 2 MB komprimiert,
und werden nur gelesen, wenn ein Foliensatz ein Emoji enthält.

Die Bilder zeichnen die HTML-Seite, `--images`/`--image`, `--pdf`, `--pptx`
und `--odp` gleichermaßen: go-webengine rastert das SVG. Alle 4.009
Grafiken lassen sich mit seinem SVG-Rasterizer parsen und zeichnen; ein Test prüft das. Seit
v0.5.0 ist dieser Rasterizer [go-gfx/oksvg](https://github.com/go-gfx/oksvg)
v0.2.0: `clipPath`, `fill-rule: evenodd` und Bogenbefehle werden jetzt so gezeichnet, wie
Chrome sie zeichnet, was ein Twemoji und 211 weitere verändert hat; SVG-Logos
sind nicht mehr verzerrt.

Die Regel, die ein Emoji an den Text anpasst (`img[data-marp-twemoji]`, ein `em`
im Quadrat, an der Grundlinie ausgerichtet), wird dem Theme so hinzugefügt, wie Marp Core sie hinzufügt,
sodass auch ein eigenes Theme sie erhält.

## Was umgewandelt wird {#what-is-converted}

Die Regeln sind die von Marp Core; ein Test vergleicht sie mit dem HTML, das Marp Core 4.4.0
für einen Foliensatz mit allen Fällen schreibt.

| | |
| --- | --- |
| `:smile:`, `:+1:`, `:heart_eyes:` | die 1.903 Namen von markdown-it-emoji, in Text, Überschriften, Links, Tabellen, `header` und `footer`; nicht in Code, URLs oder im Alt-Text eines Bildes |
| `😄`, `👩‍💻`, `🇫🇷`, `1️⃣`, `👍🏽` | Unicode-Emoji, im Text **und im Code**: Marp Core schreibt das gerenderte Code-HTML um |
| `✔`, `©`, `™`, `▶` | bleiben Text: Zeichen, deren Standarddarstellung Text ist, sind nur mit U+FE0F ein Emoji (`✔️`); `❤` ist die Ausnahme |
| `☺︎` | bleibt Text: Ein U+FE0E hinter einem Emoji verlangt die schlichte Glyphe |
| `12:30:45`, `:not_a_name:` | bleiben unverändert |

Die `id` einer Überschrift und die Gliederung verwenden das Emoji, nicht den Namen des Shortcodes:
`# Launch :rocket:` hat die ID `launch-%F0%9F%9A%80`, wie in Marp.

Wo go-slides von Marp Core abweicht, wandelt es mehr um: `✌️ ☝️ ✍️ 🏋️` und einige
weitere Zeichen mit Text als Standarddarstellung, die der emoji-regex von Marp Core als Text belässt,
sowie die wenigen Sequenzen, die Twemoji zeichnet, Unicode aber nicht auflistet.

## Abschalten {#turning-it-off}

`--no-emoji` (im Code `slides.Config{NoEmoji: true}`) belässt Shortcodes und
Unicode-Emoji als den Text, als der sie geschrieben wurden; das entspricht
`emoji: false` von Marp Core.

## Quellenangabe {#attribution}

Die Grafiken stammen von Twemoji, Copyright 2020 Twitter, Inc. und weitere Mitwirkende,
lizenziert unter [CC-BY 4.0](https://creativecommons.org/licenses/by/4.0/) und
weitergeführt vom Projekt [jdecked/twemoji](https://github.com/jdecked/twemoji).
Ein Foliensatz, der sie zeigt, sollte sie nennen, etwa in einer Fußzeile oder auf einer
Schlussfolie. Die Shortcode-Tabelle stammt von markdown-it-emoji (MIT).
