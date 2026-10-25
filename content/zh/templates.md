---
title: "模板"
weight: 30
description: "在演示文稿中使用 Go text/template 插值：变量值、生成幻灯片、页码、文件引入。"
tags: [templates]
---

演示文稿本身就是一个 Go [`text/template`](https://pkg.go.dev/text/template)，
在解析 Markdown **之前**求值。Go 模板语言的所有功能（动作、管道、`if`、`range`、`with`、变量）
都可以在演示文稿中使用。

## 数据 {#data}

模板的点（dot）依次包含 front matter 的键、其 `data:` 映射中的键，
以及通过命令行或 `slides.yaml` 提供的数据，后者覆盖前者。

```markdown
---
marp: true
title: Quarterly review
data:
  team: Platform
---

# {{ .title }}

Presented by the {{ .team }} team.
```

通过命令行提供：

| 参数 | 提供的内容 |
| --- | --- |
| `--data file.yaml` | 一个 YAML 或 JSON 数据文件（可重复） |
| `--set key=value` | 一个值（可重复） |

## 生成幻灯片 {#generating-slides}

整个正文是一个模板，因此写出 `---` 的 `range` 就能生成幻灯片：

```markdown
{{ range .releases }}
---

# {{ .version }}

{{ range .changes }}- {{ . }}
{{ end }}
{{ end }}
```

## 页码与总页数 {#page-and-total}

`{{ page }}` 和 `{{ total }}` 分别是当前幻灯片的序号和幻灯片总数。
只有在演示文稿被拆分之后才能知道它们的值，所以它们在那时才被填入：
生成的演示文稿统计的是它实际生成的幻灯片。

```markdown
<!-- footer: "{{ page }} / {{ total }}" -->
```

## 引入文件 {#include}

`{{ include "parts/agenda.md" }}` 会插入另一个文件，路径相对于演示文稿。

## 代码不受影响 {#code-is-left-alone}

围栏代码块和行内代码在模板执行之前会被暂存，之后原样还原，
因此一张讲解 Go 模板的幻灯片可以直接显示 `{{ .Name }}`，无需转义。

## 关闭 {#turning-it-off}

`--no-template` 会将演示文稿当作纯 Marp 读取：`{{` 就是普通文本。
