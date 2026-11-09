---
title: "Пакети тем"
weight: 42
description: "Тема в одному файлі: архів .gsta.xz з theme.hcl, її CSS, шрифтами, ресурсами та kinds, який перевіряють як недовірений вхід."
tags: [themes, hcl]
---

**Пакет теми** — це один файл, `<name>.gsta.xz` (Go-Slides Theme Archive):
tar-архів (ustar або pax), стиснений xz, з файлом `theme.hcl` у корені та файлами, які той оголошує. Він містить CSS
теми, шрифти, логотипи та тла і, якщо потрібно, типи слайдів (kinds), а також
назву, версію та ліцензію, щоб тему можна було передавати й встановлювати.
Пакети тем є у v0.18.0 і пізніших.

## The theme.hcl file {#the-themehcl-file}

```hcl
# theme.hcl
theme {
  name        = "acme"
  version     = "1.0.0"                 # semantic version
  description = "Example theme"
  authors     = ["go-slides contributors"]
  license     = "MIT"                   # SPDX expression
  homepage    = "https://go-slides.github.io"
  extends     = "default"               # the base theme, built in or installed
  min_slides  = "0.18.0"                # oldest go-slides it works with
}

css "theme.css" {}                      # in order; sha256 = "…" is optional

font "Roboto Mono" {
  weight  = 400                         # or "100 900"
  style   = "normal"
  file    = "fonts/RobotoMono-Regular.ttf"
  license = "OFL-1.1"                   # required for a font
}

asset "logo" { file = "assets/logo.svg" }

kind "acme-title" {                     # optional slide kinds
  class    = "lead"
  paginate = "false"
}
```

`theme.hcl` використовує рідний синтаксис HCL лише з літеральними значеннями:
жодних функцій, змінних, `file()` чи `env()`. Назва складається з малих літер,
цифр і дефісів і не може бути `default`, `gaia` чи `uncover`. `extends`
називає базову тему; кожен блок `css` додає таблицю стилів за порядком;
кожен блок `font` і `asset` називає файл пакета; кожен блок `kind` — це
[kind]({{< relref "kinds" >}}). `sha256` у блоці `css`, `font` чи `asset`
перевіряють під час читання пакета.

CSS посилається на шрифти й ресурси за їхнім шляхом у пакеті, відносно файлу
CSS: `background: url(assets/logo.svg)`. Під час завантаження пакета кожне таке
посилання стає URL `data:`, щоб сторінка була самодостатньою, а блок `font`
стає `@font-face`. Пакет може також містити файли `LICENSE`, `NOTICE`
і `README`; будь-який інший файл, якого `theme.hcl` не оголошує, відхиляють.

## Packing, installing and using {#packing-installing-and-using}

```
slides theme pack dir       write dir/theme.hcl and its files as <name>.gsta.xz
slides theme verify file    run every check on a package
slides theme install file   check it, then install it for the current user
slides theme list           the installed packages
slides theme remove name
```

Встановлені пакети лежать у `go-slides/themes` у теці конфігурації користувача
(`os.UserConfigDir()`: `~/Library/Application Support`,
`$XDG_CONFIG_HOME` або `~/.config`, `%AppData%`). У презентації:

```markdown
---
theme: acme
---
```

знаходить встановлений пакет, після вбудованих тем і `--theme-set`.
`--theme acme.gsta.xz` і `--theme-set acme.gsta.xz` використовують файл
безпосередньо, без встановлення; див. [Командний рядок]({{< relref "cli" >}}).

Kinds пакета застосовуються з `--theme file.gsta.xz`, з `--theme name`
для встановленого пакета або з `theme: name` у front matter презентації;
власні kinds презентації мають над ними пріоритет.

`extends` пакета називає його базову тему, яку шукають у такому порядку: вбудовану (або задану через `--theme-set`), потім встановлену, потім сусідній `<base>.gsta.xz` у тому самому каталозі, що й файл пакета, щоб родину тем (одна базова, кілька малих дочірніх) можна було спробувати просто з її файлів. Помилка називає три місця пошуку. Ланцюг пакетів налічує щонайбільше 8 ланок, а цикл відхиляється (`theme.meta`).

## A package is untrusted input {#a-package-is-untrusted-input}

Пакет теми може прийти звідки завгодно, тож кожен читач (`pack`, `verify`,
`install`, `--theme`, збірка) перевіряє його однаково і відхиляє з помилкою,
що має код (`theme.archive`, `theme.limit`, `theme.meta`, `theme.css`,
`theme.svg`, `theme.font`, `theme.checksum`, `theme.version`):

- **Архів.** Один потік xz, словник якого обмежений, а вікно росте разом
  з даними. У tar: жодних `..`, абсолютного шляху, літери диска чи зворотної
  косої риски в назві; лише звичайні файли й каталоги (жодних символічних чи
  жорстких посилань, пристроїв і FIFO); жодних повторюваних назв (порівняння
  без урахування регістру); лише заголовки ustar і pax (жодних розширень GNU,
  жодних розріджених файлів), заголовок pax не більше 8 КіБ. Розміри — це те,
  що прочитано, а не те, що заявляє заголовок.
- **Межі.** Щонайбільше 16 МіБ архіву, 256 записів, 8 МіБ на запис, 24 МіБ
  після розпакування, словник xz 16 МіБ, розпакований потік щонайбільше в 100
  разів більший за архів (від 64 КіБ), `theme.hcl` не більше 64 КіБ і 32 рівнів
  вкладеності дужок, 1 МіБ CSS на файл, 4 МіБ на шрифт чи зображення.
- **CSS.** Жодних `@import`, жодних `@import-theme` (використовуйте `extends`),
  жодних `image-set()`, `src()` чи інших функцій, що завантажують ресурс,
  жодних екранувань зворотною косою рискою поза рядком, жодних `</style`.
  `url()` — це `#fragment`, зображення чи шрифт `data:` або оголошений файл
  пакета, але ніколи не той, що виходить за його межі.
- **SVG.** Жодних `script`, `foreignObject`, фреймів, embed, `set`, DOCTYPE
  і обробників подій; `href` — це `#fragment`; жодних `url()`, що виходять
  за межі файлу.
- **Шрифти та зображення** розбирають: шрифт має бути шрифтом, зображення має
  відповідати своєму розширенню, і кожне має межу розміру.

Тека `testdata/themes/acme` у репозиторії — невеликий приклад.
