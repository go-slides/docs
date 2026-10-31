---
title: "Емодзі"
weight: 36
description: "Шорткоди емодзі й емодзі Unicode стають зображеннями Twemoji без мережі: у HTML-сторінці, зображеннях, PDF, PowerPoint та OpenDocument."
tags: [emoji]
---

{{< callout type="warning" >}}
Емодзі з'явилися у v0.4.0; у v0.3.0 їх не було.
{{< /callout >}}

Пишіть емодзі як шорткод або як сам символ:

```markdown
# Launch :rocket: 🎉

Shortcodes :smile: :+1: :heart_eyes: and Unicode 😄 👍🏽 👩‍💻 🇫🇷 1️⃣.
```

Кожне стає зображенням [Twemoji](https://github.com/jdecked/twemoji), яке малює
Marp Core, з розміткою, яку пише Marp Core:

```html
<img class="emoji" draggable="false" alt="😄" src="data:image/svg+xml;base64,…" data-marp-twemoji=""/>
```

## Без мережі {#no-network}

`src` у Marp Core — це URL на CDN, тож презентація, показана офлайн,
надрукована або перетворена на PDF на машині без з'єднання, показує зламані
зображення. go-slides кладе сам SVG у URI `data:`. Графіка (4 009 зображень,
Twemoji 17.0.3, Emoji 17.0) вбудована в програму з модуля на чистому Go
[go-icons/twemoji](https://github.com/go-icons/twemoji), 2 МБ у стиснутому
вигляді, і читається лише тоді, коли в презентації є емодзі.

Зображення малюються однаково HTML-сторінкою, `--images`/`--image`, `--pdf`,
`--pptx` та `--odp`: SVG растеризує go-webengine. Усі 4 009 зображень
розбираються й малюються його растеризатором SVG; це перевіряє тест. Починаючи
з v0.5.0 цей растеризатор — [go-gfx/oksvg](https://github.com/go-gfx/oksvg)
v0.2.0: `clipPath`, `fill-rule: evenodd` і команди дуг тепер малюються так, як
їх малює Chrome, що змінило один Twemoji і 211 інших, а логотипи SVG більше не
спотворюються.

Правило, яке підганяє розмір емодзі під текст (`img[data-marp-twemoji]`,
квадрат в один `em`, вирівняний по базовій лінії), додається до теми так, як
його додає Marp Core, тож власна тема теж його отримує.

## Що перетворюється {#what-is-converted}

Правила — це правила Marp Core; тест порівнює їх з HTML, який Marp Core 4.4.0
пише для презентації з усіма випадками.

| | |
| --- | --- |
| `:smile:`, `:+1:`, `:heart_eyes:` | 1 903 назви markdown-it-emoji, у тексті, заголовках, посиланнях, таблицях, `header` і `footer`; не в коді, URL чи альтернативному тексті зображення |
| `😄`, `👩‍💻`, `🇫🇷`, `1️⃣`, `👍🏽` | емодзі Unicode, у тексті **і в коді**: Marp Core переписує відрендерений HTML коду |
| `✔`, `©`, `™`, `▶` | залишаються текстом: символи, типове представлення яких — текст, є емодзі лише з U+FE0F (`✔️`); `❤` — виняток |
| `☺︎` | залишається текстом: U+FE0E після емодзі просить звичайний гліф |
| `12:30:45`, `:not_a_name:` | залишаються як написано |

`id` заголовка та зміст використовують емодзі, а не назву шорткоду:
`# Launch :rocket:` має id `launch-%F0%9F%9A%80`, як у Marp.

Там, де go-slides відрізняється від Marp Core, він перетворює більше: `✌️ ☝️ ✍️ 🏋️`
та кілька інших символів із текстовим типовим представленням, які emoji-regex
Marp Core залишає текстом, а також пару послідовностей, які малює Twemoji, але
яких Unicode не перелічує.

## Вимкнення {#turning-it-off}

`--no-emoji` (у коді — `slides.Config{NoEmoji: true}`) залишає шорткоди й емодзі
Unicode таким текстом, яким їх написано; це `emoji: false` у Marp Core.

## Авторство {#attribution}

Графіка — це Twemoji, Copyright 2020 Twitter, Inc and other contributors,
ліцензія [CC-BY 4.0](https://creativecommons.org/licenses/by/4.0/); проєкт
продовжує [jdecked/twemoji](https://github.com/jdecked/twemoji). Презентація,
що їх показує, має вказати авторство, наприклад у нижньому колонтитулі або на
завершальному слайді. Таблиця шорткодів — із markdown-it-emoji (MIT).
