---
title: "Templates"
weight: 30
description: "Interpolação com text/template do Go em uma apresentação: valores, slides gerados, números de página, includes."
tags: [templates]
---

A apresentação é um [`text/template`](https://pkg.go.dev/text/template) do Go,
avaliado **antes** de o Markdown ser analisado. Tudo o que a linguagem de
templates do Go tem — ações, pipelines, `if`, `range`, `with`, variáveis —
funciona em uma apresentação.

## Dados {#data}

O ponto (dot) do template contém as chaves do front matter, depois as chaves do
mapeamento `data:`, depois os dados informados na linha de comando ou em
`slides.yaml` — cada um sobrescrevendo o anterior.

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

Pela linha de comando:

| Flag | Fornece |
| --- | --- |
| `--data file.yaml` | um arquivo YAML ou JSON de dados (repetível) |
| `--set key=value` | um valor (repetível) |

## Gerar slides {#generating-slides}

O corpo inteiro é um só template, então um `range` que escreve `---` escreve
slides:

```markdown
{{ range .releases }}
---

# {{ .version }}

{{ range .changes }}- {{ . }}
{{ end }}
{{ end }}
```

## Página e total {#page-and-total}

`{{ page }}` e `{{ total }}` são o número do slide e o número de slides. Eles só
são conhecidos depois que a apresentação foi dividida, então são preenchidos
nesse momento — uma apresentação gerada conta os slides que gerou.

```markdown
<!-- footer: "{{ page }} / {{ total }}" -->
```

## Include {#include}

`{{ include "parts/agenda.md" }}` insere outro arquivo, relativo à
apresentação.

## O código é deixado em paz {#code-is-left-alone}

Blocos de código delimitados e trechos de código em linha são postos de lado
antes de o template rodar e restaurados literalmente, então um slide sobre
templates Go pode mostrar `{{ .Name }}` sem escapá-lo.

## Desativar {#turning-it-off}

`--no-template` lê a apresentação como Marp simples: `{{` é texto.
