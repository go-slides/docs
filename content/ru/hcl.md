---
title: "HCL"
weight: 45
description: "HCL2 везде, где презентация принимает YAML: front matter на HCL, slides.hcl, файлы типов и данных, с выражениями и функциями."
tags: [hcl, configuration]
---

{{< callout type="warning" >}}
HCL появился в v0.3.0 ([go-slides/slides#13](https://github.com/go-slides/slides/pull/13)).
{{< /callout >}}

Везде, где `slides` читает YAML, он читает и собственный синтаксис
[HCL2](https://github.com/hashicorp/hcl/blob/main/hclsyntax/spec.md) — ради
его выражений и функций.

| Где | YAML | HCL |
| --- | --- | --- |
| front matter | `---` … `---` | `---hcl` … `---` |
| конфигурация | `slides.yaml` | `slides.hcl`, ищется первым |
| файлы типов | `--kinds kinds.yaml`, `kinds: kinds.yaml` | любой файл с именем `*.hcl` |
| файлы данных | `--data data.yaml` или `.json` | любой файл с именем `*.hcl` |

Комментарии-директивы (`<!-- class: invert -->`) остаются в YAML, как в Marp, а
тело презентации остаётся [шаблоном Go]({{< relref "templates" >}}).

## Front matter на HCL {#a-front-matter-in-hcl}

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

- **Атрибуты** — это ключи YAML под теми же именами: директивы Marp
  (`theme`, `paginate`, `header`, `footer`, `class`, `backgroundColor`…),
  метаданные (`title`, `author`…), `template`, `delims`, `kinds` (пути к
  файлам) и `data` — объект.
- **Блоки `kind "<name>" { … }`** определяют [типы слайдов]({{< relref "kinds" >}})
  с локальными директивами и `extends`.
- **Значения** попадают в директивы в тех строковых формах, которые им даёт
  YAML, — `true` становится `"true"`, список классов остаётся списком, — так что
  директива ведёт себя одинаково, каким бы синтаксисом она ни была задана.
  `null` оставляет директиву незаданной.
- **`data`** сохраняет типы: в шаблонах целое число — это `int`, остальные
  числа — `float64`, а списки и объекты остаются списками и отображениями, так
  что `{{ add .quarter 1 }}` работает.
- Значение HCL окончательное: шаблоны Go не запускаются над ним второй раз.

## Переменные {#variables}

`data` — это данные шаблона: собственный `data` из front matter, поверх
которого накладываются данные из `slides.hcl`, из файлов `--data` и из `--set`.
`"${data.company}"` работает везде, где работал бы `{{ .company }}`.

Атрибуты front matter могут ссылаться на `data.*`, но **не друг на друга**:
`footer = title` — ошибка. Точечные директивы (`_class`…) не являются
идентификаторами HCL; пишите их в комментарии.

## Функции {#functions}

| | |
| --- | --- |
| `page()`, `total()` | номер страницы и номер последней страницы, как `{{ page }}` и `{{ total }}`: колонтитул нумеруется |
| `now()`, `timestamp()` | время в формате RFC 3339, местное и UTC; `formatdate` форматирует его |
| `env(name)` | переменная окружения, `""`, если не задана |
| `file(path)` | текстовый файл UTF-8 размером не более 1 МиБ, путь отсчитывается от презентации или файла конфигурации |
| числа | `abs ceil floor log max min parseint pow signum` |
| строки | `chomp format formatlist join lower regex regexall regexreplace replace split strlen strrev substr title trim trimprefix trimspace trimsuffix upper` |
| коллекции | `chunklist coalesce coalescelist compact concat contains distinct element flatten keys length lookup merge range reverse slice sort values zipmap` |
| кодировки | `csvdecode jsondecode jsonencode` |
| даты | `formatdate timeadd` |

Все строки, кроме первых четырёх, — стандартная библиотека
[go-cty](https://github.com/zclconf/go-cty). Ни одна не обращается к сети.

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

Неизвестный атрибут — это ошибка, а не молчаливое отсутствие действия. Файл
типов содержит блоки `kind`; остальные его атрибуты игнорируются, поэтому
`slides.hcl` может служить и таким файлом. Атрибуты файла данных — это ключи
данных:

```hcl
# data.hcl
company = "ACME"
year    = 2026
motto   = trimspace(file("motto.txt"))
```

## Ошибки {#errors}

Ошибки — это диагностика HCL с указанием места в самом файле презентации:
`footer = data.compnay` в четвёртой строке презентации даёт

```
deck.md:4:14: Unsupported attribute: This object does not have an attribute named "compnay".
```

## Marp и `slides expand` {#marp-and-slides-expand}

Marp тоже читает `---hcl` как начало front matter — его правило front matter
открывается на любой первой строке, начинающейся с трёх дефисов, — и, раз блок
не является отображением YAML, не применяет из него ничего: презентация всё
равно отрисовывается в Marp, но без своих директив. Тест сверяет это с
собственным выводом Marp Core.

`slides expand` записывает front matter обратно в YAML, так что развёрнутая
презентация — это обычный Marp, и отрисовывается она так же, как `slides`
отрисовывает оригинал:

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

## Лицензии {#licences}

Поддержка HCL подключает [hashicorp/hcl](https://github.com/hashicorp/hcl)
(MPL-2.0, используется без изменений как библиотека) и
[zclconf/go-cty](https://github.com/zclconf/go-cty) (MIT).
