---
title: "HCL"
weight: 45
description: "HCL2 onde quer que uma apresentação aceite YAML: um front matter em HCL, slides.hcl, arquivos de kinds e de dados, com expressões e funções."
tags: [hcl, configuration]
---

{{< callout type="warning" >}}
O HCL está na v0.3.0 ([go-slides/slides#13](https://github.com/go-slides/slides/pull/13)).
{{< /callout >}}

Em todo lugar onde o `slides` lê YAML, ele também lê a sintaxe nativa do
[HCL2](https://github.com/hashicorp/hcl/blob/main/hclsyntax/spec.md) — por
causa das suas expressões e das suas funções.

| Onde | YAML | HCL |
| --- | --- | --- |
| front matter | `---` … `---` | `---hcl` … `---` |
| configuração | `slides.yaml` | `slides.hcl`, procurado primeiro |
| arquivos de kinds (o tipo de um slide) | `--kinds kinds.yaml`, `kinds: kinds.yaml` | qualquer arquivo chamado `*.hcl` |
| arquivos de dados | `--data data.yaml` ou `.json` | qualquer arquivo chamado `*.hcl` |

Os comentários de diretiva (`<!-- class: invert -->`) continuam em YAML, como no
Marp, e o corpo da apresentação continua sendo um
[template Go]({{< relref "templates" >}}).

## Um front matter em HCL {#a-front-matter-in-hcl}

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

- Os **atributos** são as chaves do YAML com os mesmos nomes: as diretivas do
  Marp (`theme`, `paginate`, `header`, `footer`, `class`, `backgroundColor`…),
  os metadados (`title`, `author`…), `template`, `delims`, `kinds` (caminhos de
  arquivo) e `data`, um objeto.
- Os blocos **`kind "<name>" { … }`** definem [kinds]({{< relref "kinds" >}}),
  com as diretivas locais e `extends`.
- Os **valores** chegam às diretivas nas formas de string que o YAML lhes dá —
  `true` vira `"true"`, uma lista de classes continua uma lista — de modo que uma
  diretiva se comporta do mesmo jeito, qualquer que seja a sintaxe que a
  definiu. `null` deixa a diretiva sem definição.
- **`data`** mantém os seus tipos: nos templates, um número inteiro é um `int`,
  os demais números são `float64`, e listas e objetos continuam listas e mapas,
  então `{{ add .quarter 1 }}` funciona.
- Um valor HCL é final: os templates Go não passam por ele uma segunda vez.

## Variáveis {#variables}

`data` são os dados do template: o `data` do próprio front matter, sobreposto
pelo de `slides.hcl`, pelos arquivos de `--data` e por `--set`.
`"${data.company}"` funciona em qualquer lugar onde `{{ .company }}` funcionaria.

Os atributos do front matter podem referenciar `data.*`, mas **não uns aos
outros**: `footer = title` é um erro. As diretivas pontuais (`_class`…) não são
identificadores HCL; escreva-as em um comentário.

## Funções {#functions}

| | |
| --- | --- |
| `page()`, `total()` | o número da página e o número da última página, como `{{ page }}` e `{{ total }}`: um rodapé pagina |
| `now()`, `timestamp()` | a hora, em RFC 3339, local e UTC; `formatdate` a formata |
| `env(name)` | uma variável de ambiente, `""` quando não definida |
| `file(path)` | um arquivo de texto UTF-8 de no máximo 1 MiB, relativo à apresentação ou ao arquivo de configuração |
| números | `abs ceil floor log max min parseint pow signum` |
| strings | `chomp format formatlist join lower regex regexall regexreplace replace split strlen strrev substr title trim trimprefix trimspace trimsuffix upper` |
| coleções | `chunklist coalesce coalescelist compact concat contains distinct element flatten keys length lookup merge range reverse slice sort values zipmap` |
| codificações | `csvdecode jsondecode jsonencode` |
| datas | `formatdate timeadd` |

Todas as linhas, exceto as quatro primeiras, vêm da biblioteca padrão do
[go-cty](https://github.com/zclconf/go-cty). Nenhuma acessa a rede.

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

Um atributo desconhecido é um erro, não uma operação nula silenciosa. Um arquivo
de kinds contém blocos `kind`; os outros atributos dele são ignorados, então o
`slides.hcl` pode servir como um. Os atributos de um arquivo de dados são as
chaves dos dados:

```hcl
# data.hcl
company = "ACME"
year    = 2026
motto   = trimspace(file("motto.txt"))
```

## Erros {#errors}

Os erros são os diagnósticos do HCL, posicionados no próprio arquivo da
apresentação: `footer = data.compnay` na quarta linha da apresentação produz

```
deck.md:4:14: Unsupported attribute: This object does not have an attribute named "compnay".
```

## Marp e `slides expand` {#marp-and-slides-expand}

O Marp também lê `---hcl` como o início de um front matter — a regra de front
matter dele abre em qualquer primeira linha que comece com três hifens — e, como
o bloco não é um mapeamento YAML, não aplica nada dele: a apresentação ainda é
renderizada no Marp, sem as suas diretivas. Um teste fixa isso em comparação com
a saída do próprio Marp Core.

`slides expand` escreve o front matter de volta como YAML, de modo que a
apresentação expandida é Marp simples e é renderizada como o `slides` renderiza
a original:

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

## Licenças {#licences}

O suporte a HCL liga com [hashicorp/hcl](https://github.com/hashicorp/hcl)
(MPL-2.0, usado sem modificação como biblioteca) e
[zclconf/go-cty](https://github.com/zclconf/go-cty) (MIT).
