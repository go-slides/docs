---
title: "Emoji"
weight: 36
description: "Les shortcodes d'emoji et les emoji Unicode deviennent des images Twemoji, sans réseau : dans la page HTML, les images, le PDF, PowerPoint et OpenDocument."
tags: [emoji]
---

{{< callout type="warning" >}}
Les emoji sont arrivés en v0.4.0 ; la v0.3.0 n'en avait pas.
{{< /callout >}}

Écrivez un emoji sous forme de shortcode ou avec le caractère lui-même :

```markdown
# Launch :rocket: 🎉

Shortcodes :smile: :+1: :heart_eyes: and Unicode 😄 👍🏽 👩‍💻 🇫🇷 1️⃣.
```

Chacun devient l'image [Twemoji](https://github.com/jdecked/twemoji) que dessine Marp Core,
avec le balisage qu'écrit Marp Core :

```html
<img class="emoji" draggable="false" alt="😄" src="data:image/svg+xml;base64,…" data-marp-twemoji=""/>
```

## Sans réseau {#no-network}

Dans Marp Core, `src` est une URL sur un CDN : une présentation affichée hors ligne, imprimée ou
convertie en PDF sur une machine sans connexion montre des images cassées.
go-slides place le SVG lui-même dans une URI `data:`. Les graphismes (4 009, issus de
Twemoji 17.0.3, Emoji 17.0) sont intégrés au programme depuis le module en pur Go
[go-icons/twemoji](https://github.com/go-icons/twemoji) : 2 Mo compressés, lus
seulement quand une présentation contient un emoji.

Les images sont dessinées de la même façon par la page HTML, par `--images`/`--image`, par `--pdf`, par
`--pptx` et par `--odp` : go-webengine rastérise le SVG. Les 4 009
graphismes s'analysent et se dessinent tous avec son rastériseur SVG ; un test le vérifie. Depuis
la v0.5.0, ce rastériseur est [go-gfx/oksvg](https://github.com/go-gfx/oksvg)
v0.2.0 : `clipPath`, `fill-rule: evenodd` et les commandes d'arc se dessinent désormais comme dans Chrome,
ce qui a modifié un Twemoji et 211 autres, et les logos SVG ne sont plus
déformés.

La règle qui dimensionne un emoji selon le texte (`img[data-marp-twemoji]`, un carré d'un `em`,
aligné sur la ligne de base) est ajoutée au thème comme le fait Marp Core :
un thème personnel en bénéficie aussi.

## Ce qui est converti {#what-is-converted}

Les règles sont celles de Marp Core ; un test les compare au HTML qu'écrit Marp Core 4.4.0
pour une présentation couvrant tous les cas.

| | |
| --- | --- |
| `:smile:`, `:+1:`, `:heart_eyes:` | les 1 903 noms de markdown-it-emoji, dans le texte, les titres, les liens, les tableaux, `header` et `footer` ; pas dans le code, les URL ni le texte alternatif d'une image |
| `😄`, `👩‍💻`, `🇫🇷`, `1️⃣`, `👍🏽` | les emoji Unicode, dans le texte **et dans le code** : Marp Core réécrit le HTML du code rendu |
| `✔`, `©`, `™`, `▶` | restent du texte : les caractères dont la présentation par défaut est textuelle ne sont des emoji qu'avec U+FE0F (`✔️`) ; `❤` fait exception |
| `☺︎` | reste du texte : un U+FE0E après un emoji demande le glyphe simple |
| `12:30:45`, `:not_a_name:` | restent tels qu'ils sont écrits |

L'`id` d'un titre et le plan utilisent l'emoji, pas le nom du shortcode :
`# Launch :rocket:` a pour id `launch-%F0%9F%9A%80`, comme dans Marp.

Là où go-slides s'écarte de Marp Core, il convertit davantage : `✌️ ☝️ ✍️ 🏋️` et
quelques autres caractères textuels par défaut que l'emoji-regex de Marp Core laisse en texte,
ainsi que la poignée de séquences que dessine Twemoji sans qu'Unicode les recense.

## Désactivation {#turning-it-off}

`--no-emoji` (dans le code, `slides.Config{NoEmoji: true}`) laisse les shortcodes et les
emoji Unicode sous la forme textuelle où ils ont été écrits, ce qui correspond à
`emoji: false` dans Marp Core.

## Attribution {#attribution}

Les graphismes sont des Twemoji, Copyright 2020 Twitter, Inc and other contributors,
sous licence [CC-BY 4.0](https://creativecommons.org/licenses/by/4.0/) et
poursuivis par le projet [jdecked/twemoji](https://github.com/jdecked/twemoji).
Une présentation qui les affiche doit les créditer, par exemple dans un pied de page ou sur une
diapositive de clôture. La table des shortcodes est celle de markdown-it-emoji (MIT).
