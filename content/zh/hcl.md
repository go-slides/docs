---
title: "HCL"
weight: 45
description: "凡是演示文稿接受 YAML 的地方都可以使用 HCL2：HCL front matter、slides.hcl、kind 文件和数据文件，支持表达式和函数。"
tags: [hcl, configuration]
---

{{< callout type="warning" >}}
HCL 自 v0.3.0 起提供（[go-slides/slides#13](https://github.com/go-slides/slides/pull/13)）。
{{< /callout >}}

凡是 `slides` 读取 YAML 的地方，它也能读取
[HCL2](https://github.com/hashicorp/hcl/blob/main/hclsyntax/spec.md) 的原生语法，
以便使用其表达式和函数。

| 位置 | YAML | HCL |
| --- | --- | --- |
| front matter | `---` … `---` | `---hcl` … `---` |
| 配置文件 | `slides.yaml` | `slides.hcl`，优先查找 |
| kind 文件 | `--kinds kinds.yaml`、`kinds: kinds.yaml` | 任何名为 `*.hcl` 的文件 |
| 数据文件 | `--data data.yaml` 或 `.json` | 任何名为 `*.hcl` 的文件 |

与 Marp 一样，指令注释（`<!-- class: invert -->`）仍然使用 YAML，
演示文稿正文仍然是 [Go 模板]({{< relref "templates" >}})。

## 用 HCL 编写 front matter {#a-front-matter-in-hcl}

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

- **属性** 与同名的 YAML 键对应：Marp 的指令
  （`theme`、`paginate`、`header`、`footer`、`class`、`backgroundColor`…）、
  元数据（`title`、`author`…）、`template`、`delims`、`kinds`（文件路径）
  以及 `data`（一个对象）。
- **`kind "<name>" { … }`** 块用于定义 [kind（幻灯片类型）]({{< relref "kinds" >}})，
  可包含局部指令和 `extends`。
- **值** 以 YAML 给出的字符串形式传给指令：
  `true` 即 `"true"`，class 列表仍是列表，因此无论用哪种语法设置，指令的行为都相同。
  `null` 表示不设置该指令。
- **`data`** 保留其类型：在模板中，整数是 `int`，
  其他数字是 `float64`，列表和对象仍是列表和 map，因此
  `{{ add .quarter 1 }}` 可以正常工作。
- HCL 的值是最终值：Go 模板不会再对其执行第二遍。

## 变量 {#variables}

`data` 就是模板数据：front matter 自身的 `data`，依次被
`slides.hcl` 中的、`--data` 文件中的以及 `--set` 中的数据覆盖。凡是能用
`{{ .company }}` 的地方，都可以用 `"${data.company}"`。

front matter 的属性可以引用 `data.*`，但**不能相互引用**：
`footer = title` 会报错。单页指令（`_class`…）不是合法的 HCL
标识符；请把它们写在注释中。

## 函数 {#functions}

| | |
| --- | --- |
| `page()`、`total()` | 页码和最后一页的页码，相当于 `{{ page }}` 和 `{{ total }}`：页脚可以显示页码 |
| `now()`、`timestamp()` | 当前时间，RFC 3339 格式，分别为本地时间和 UTC；可用 `formatdate` 格式化 |
| `env(name)` | 环境变量，未设置时为 `""` |
| `file(path)` | 一个不超过 1 MiB 的 UTF-8 文本文件，路径相对于演示文稿或配置文件 |
| 数字 | `abs ceil floor log max min parseint pow signum` |
| 字符串 | `chomp format formatlist join lower regex regexall regexreplace replace split strlen strrev substr title trim trimprefix trimspace trimsuffix upper` |
| 集合 | `chunklist coalesce coalescelist compact concat contains distinct element flatten keys length lookup merge range reverse slice sort values zipmap` |
| 编码 | `csvdecode jsondecode jsonencode` |
| 日期 | `formatdate timeadd` |

除前四行外，其余均来自 [go-cty](https://github.com/zclconf/go-cty) 的标准库。
所有函数都不会访问网络。

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

未知属性会报错，而不是被静默忽略。kind 文件包含
`kind` 块，其他属性会被忽略，因此 `slides.hcl` 本身也可以充当 kind 文件。
数据文件的属性就是数据的键：

```hcl
# data.hcl
company = "ACME"
year    = 2026
motto   = trimspace(file("motto.txt"))
```

## 错误 {#errors}

错误信息就是 HCL 的诊断信息，定位到演示文稿文件本身：
在演示文稿第四行写 `footer = data.compnay`，会得到

```
deck.md:4:14: Unsupported attribute: This object does not have an attribute named "compnay".
```

## Marp 与 `slides expand` {#marp-and-slides-expand}

Marp 同样会把 `---hcl` 当作 front matter 的开头（它的 front matter
规则会在任何以三个短横线开头的第一行开启），但由于该块不是 YAML 映射，
其中的内容一概不会应用：演示文稿在 Marp 中仍能渲染，只是没有这些指令。
有测试针对 Marp Core 自身的输出固定了这一行为。

`slides expand` 会把 front matter 写回为 YAML，因此展开后的演示文稿是纯 Marp，
渲染结果与 `slides` 渲染原文件时相同：

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

## 许可证 {#licences}

HCL 支持链接了 [hashicorp/hcl](https://github.com/hashicorp/hcl)
（MPL-2.0，作为库原样使用，未作修改）和
[zclconf/go-cty](https://github.com/zclconf/go-cty)（MIT）。
