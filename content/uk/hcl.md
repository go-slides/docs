---
title: "HCL"
weight: 45
description: "HCL2 скрізь, де презентація приймає YAML: front matter у HCL, slides.hcl, файли kinds і даних, з виразами та функціями."
tags: [hcl, configuration]
---

{{< callout type="warning" >}}
HCL з'явився у v0.3.0 ([go-slides/slides#13](https://github.com/go-slides/slides/pull/13)).
{{< /callout >}}

Скрізь, де `slides` читає YAML, він читає й рідний синтаксис
[HCL2](https://github.com/hashicorp/hcl/blob/main/hclsyntax/spec.md) — заради
його виразів і функцій.

| Де | YAML | HCL |
| --- | --- | --- |
| front matter | `---` … `---` | `---hcl` … `---` |
| конфігурація | `slides.yaml` | `slides.hcl`, шукається першим |
| файли kinds | `--kinds kinds.yaml`, `kinds: kinds.yaml` | будь-який файл з іменем `*.hcl` |
| файли даних | `--data data.yaml` або `.json` | будь-який файл з іменем `*.hcl` |

Коментарі-директиви (`<!-- class: invert -->`) залишаються YAML, як у Marp, а
тіло презентації залишається [шаблоном Go]({{< relref "templates" >}}).

## Front matter у HCL {#a-front-matter-in-hcl}

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

- **Атрибути** — це ключі YAML під тими самими іменами: директиви Marp
  (`theme`, `paginate`, `header`, `footer`, `class`, `backgroundColor`…),
  метадані (`title`, `author`…), `template`, `delims`, `kinds` (шляхи до
  файлів) і `data`, об'єкт.
- **Блоки `kind "<name>" { … }`** визначають [kinds]({{< relref "kinds" >}}),
  з локальними директивами та `extends`.
- **Значення** потрапляють до директив у рядкових формах, які їм дає YAML, —
  `true` стає `"true"`, список класів лишається списком, — тож директива
  поводиться однаково, яким би синтаксисом її не задано. `null` залишає
  директиву незаданою.
- **`data`** зберігає свої типи: у шаблонах ціле число — це `int`, інші числа —
  `float64`, а списки й об'єкти лишаються списками й мапами, тож
  `{{ add .quarter 1 }}` працює.
- Значення HCL остаточне: шаблони Go вдруге по ньому не проходять.

## Змінні {#variables}

`data` — це дані шаблону: власний `data` front matter, поверх якого накладено
дані зі `slides.hcl`, файли `--data` і `--set`. `"${data.company}"` працює
всюди, де працювало б `{{ .company }}`.

Атрибути front matter можуть посилатися на `data.*`, але **не одні на одних**:
`footer = title` — помилка. Spot-директиви (`_class`…) не є ідентифікаторами
HCL; пишіть їх у коментарі.

## Функції {#functions}

| | |
| --- | --- |
| `page()`, `total()` | номер сторінки й номер останньої сторінки, як `{{ page }}` та `{{ total }}`: нижній колонтитул нумерує сторінки |
| `now()`, `timestamp()` | час у форматі RFC 3339, локальний і UTC; `formatdate` форматує його |
| `env(name)` | змінна середовища, `""`, якщо не задана |
| `file(path)` | текстовий файл UTF-8 не більше 1 МіБ, відносно презентації або файлу конфігурації |
| числа | `abs ceil floor log max min parseint pow signum` |
| рядки | `chomp format formatlist join lower regex regexall regexreplace replace split strlen strrev substr title trim trimprefix trimspace trimsuffix upper` |
| колекції | `chunklist coalesce coalescelist compact concat contains distinct element flatten keys length lookup merge range reverse slice sort values zipmap` |
| кодування | `csvdecode jsondecode jsonencode` |
| дати | `formatdate timeadd` |

Усі рядки, крім перших чотирьох, — стандартна бібліотека
[go-cty](https://github.com/zclconf/go-cty). Жодна не звертається до мережі.

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

Невідомий атрибут — це помилка, а не мовчазна відсутність дії. Файл kinds
містить блоки `kind`; його інші атрибути ігноруються, тож `slides.hcl` може
слугувати таким файлом. Атрибути файлу даних — це ключі даних:

```hcl
# data.hcl
company = "ACME"
year    = 2026
motto   = trimspace(file("motto.txt"))
```

## Помилки {#errors}

Помилки — це діагностика HCL, з позицією в самому файлі презентації:
`footer = data.compnay` у четвертому рядку презентації дає

```
deck.md:4:14: Unsupported attribute: This object does not have an attribute named "compnay".
```

## Marp і `slides expand` {#marp-and-slides-expand}

Marp також читає `---hcl` як початок front matter — його правило front matter
відкривається на будь-якому першому рядку, що починається з трьох дефісів, — і,
оскільки блок не є відображенням YAML, не застосовує нічого з нього:
презентація все ще рендериться в Marp, але без своїх директив. Тест фіксує це
щодо власного виводу Marp Core.

`slides expand` записує front matter назад як YAML, тож розгорнута презентація
— звичайний Marp і рендериться так, як `slides` рендерить оригінал:

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

## Ліцензії {#licences}

Підтримка HCL лінкує [hashicorp/hcl](https://github.com/hashicorp/hcl)
(MPL-2.0, використовується без змін як бібліотека) і
[zclconf/go-cty](https://github.com/zclconf/go-cty) (MIT).
