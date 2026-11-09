---
title: "Пакеты тем"
weight: 42
description: "Тема в одном файле: архив .gsta.xz с theme.hcl, её CSS, шрифтами, ресурсами и kinds, проверяемый как недоверенный ввод."
tags: [themes, hcl]
---

**Пакет темы** — это один файл, `<name>.gsta.xz` (Go-Slides Theme Archive):
tar-архив (ustar или pax), сжатый xz, с файлом `theme.hcl` в корне и файлами, которые тот объявляет. Он несёт CSS
темы, шрифты, логотипы и фоны и, если нужно, типы слайдов (kinds), а также имя,
версию и лицензию, чтобы тему можно было передавать и устанавливать. Пакеты тем
появились после v0.17.1.

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

`theme.hcl` использует родной синтаксис HCL только с литеральными значениями:
никаких функций, переменных, `file()` и `env()`. Имя состоит из строчных букв,
цифр и дефисов и не может быть `default`, `gaia` или `uncover`. `extends`
называет базовую тему; каждый блок `css` добавляет таблицу стилей по порядку;
каждый блок `font` и `asset` называет файл пакета; каждый блок `kind` — это
[kind]({{< relref "kinds" >}}). `sha256` в блоке `css`, `font` или `asset`
проверяется при чтении пакета.

CSS ссылается на шрифты и ресурсы по их пути в пакете, относительно файла CSS:
`background: url(assets/logo.svg)`. При загрузке пакета каждая такая ссылка
становится URL `data:`, чтобы страница была самодостаточной, а блок `font`
становится `@font-face`. Пакет может также нести файлы `LICENSE`, `NOTICE`
и `README`; любой другой файл, который `theme.hcl` не объявляет, отклоняется.

## Packing, installing and using {#packing-installing-and-using}

```
slides theme pack dir       write dir/theme.hcl and its files as <name>.gsta.xz
slides theme verify file    run every check on a package
slides theme install file   check it, then install it for the current user
slides theme list           the installed packages
slides theme remove name
```

Установленные пакеты лежат в `go-slides/themes` в каталоге конфигурации
пользователя (`os.UserConfigDir()`: `~/Library/Application Support`,
`$XDG_CONFIG_HOME` или `~/.config`, `%AppData%`). В презентации:

```markdown
---
theme: acme
---
```

находит установленный пакет, после встроенных тем и `--theme-set`.
`--theme acme.gsta.xz` и `--theme-set acme.gsta.xz` используют файл
напрямую, без установки; см. [Командная строка]({{< relref "cli" >}}).

Kinds пакета применяются с `--theme file.gsta.xz`, с `--theme name`
для установленного пакета или с `theme: name` в front matter презентации;
собственные kinds презентации имеют над ними приоритет.

## A package is untrusted input {#a-package-is-untrusted-input}

Пакет темы может прийти откуда угодно, поэтому каждый читатель (`pack`,
`verify`, `install`, `--theme`, сборка) проверяет его одинаково и отклоняет
с ошибкой, у которой есть код (`theme.archive`, `theme.limit`, `theme.meta`,
`theme.css`, `theme.svg`, `theme.font`, `theme.checksum`, `theme.version`):

- **Архив.** Один поток xz, словарь которого ограничен, а окно растёт вместе
  с данными. В tar: никаких `..`, абсолютного пути, буквы диска или обратной
  косой черты в имени; только обычные файлы и каталоги (никаких символических
  или жёстких ссылок, устройств и FIFO); никаких повторяющихся имён (сравнение
  без учёта регистра); только заголовки ustar и pax (никаких расширений GNU,
  никаких разреженных файлов), заголовок pax не более 8 КиБ. Размеры — те, что
  прочитаны, а не те, что заявляет заголовок.
- **Границы.** Не более 16 МиБ архива, 256 записей, 8 МиБ на запись, 24 МиБ
  после распаковки, словарь xz 16 МиБ, распакованный поток не более чем в 100
  раз больше архива (от 64 КиБ), `theme.hcl` не более 64 КиБ и 32 уровней
  вложенности скобок, 1 МиБ CSS на файл, 4 МиБ на шрифт или изображение.
- **CSS.** Никаких `@import`, никаких `@import-theme` (используйте `extends`),
  никаких `image-set()`, `src()` и других функций, загружающих ресурс, никаких
  экранирований обратной косой чертой вне строки, никаких `</style`. `url()` —
  это `#fragment`, изображение или шрифт `data:` либо объявленный файл пакета,
  но никогда не тот, что выходит за его пределы.
- **SVG.** Никаких `script`, `foreignObject`, фреймов, embed, `set`, DOCTYPE
  и обработчиков событий; `href` — это `#fragment`; никаких `url()`, выходящих
  за пределы файла.
- **Шрифты и изображения** разбираются: шрифт должен быть шрифтом, изображение
  должно соответствовать своему расширению, и у каждого есть предел размера.

Каталог `testdata/themes/acme` в репозитории — небольшой пример.
