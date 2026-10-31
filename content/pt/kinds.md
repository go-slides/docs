---
title: "Kinds"
weight: 40
description: "Um slide diz o que ele é: kind nomeia um conjunto predefinido de diretivas locais."
tags: [kinds]
---

Um **kind** é um conjunto predefinido e nomeado de diretivas locais:
cabeçalho, rodapé, paginação, classe, fundo. Um slide diz o que ele é —
`title`, `section` — em vez de repetir a sua aparência, e o visual de uma
apresentação muda em um único lugar.

## Usar um kind {#using-a-kind}

```markdown
<!-- kind: section -->
# Part two

---

<!-- _kind: blank -->
![bg](photo.jpg)
```

- `kind:` se aplica como uma diretiva local: a este slide e aos seguintes.
- `_kind:` se aplica como uma diretiva pontual (spot): somente a este slide.
- Trocar de kind primeiro **limpa** o que o kind anterior definiu, de modo que
  uma classe ou um rodapé nunca vaza de um kind para o seguinte.
- As diretivas escritas no próprio slide **prevalecem** sobre o kind.

## Embutidos {#built-in}

| Kind | Define |
| --- | --- |
| `title` | `class: lead`, sem número de página, sem cabeçalho, sem rodapé |
| `section` | `class: lead` |
| `content` | nada: os padrões da própria apresentação |
| `blank` | sem número de página, sem cabeçalho, sem rodapé |

Qualquer um deles pode ser redefinido.

## Definir kinds {#defining-kinds}

No front matter:

```yaml
---
marp: true
kinds:
  chapter:
    class: invert
    backgroundColor: "#1e1b4b"
    paginate: false
  quote:
    extends: blank
    class: quote
---
```

`extends` inicia um kind a partir das diretivas de outro. Um kind aceita as
diretivas locais e `extends`.

## Compartilhar kinds {#sharing-kinds}

Os kinds de uma equipe pertencem a um único arquivo que todas as apresentações
nomeiam:

```yaml
# kinds.yaml
chapter:
  class: invert
  paginate: false
```

```yaml
---
marp: true
kinds: kinds.yaml        # or a list of files
---
```

ou `--kinds kinds.yaml` na linha de comando, ou `kinds:` em `slides.yaml`. A
ordem é: os kinds embutidos, depois a linha de comando e o arquivo de
configuração, depois o `kinds:` da própria apresentação.

`slides kinds talk.md` lista os kinds que uma apresentação pode usar e o que
cada um define.

## Sair do go-slides {#leaving-go-slides}

`slides expand` escreve cada kind como as diretivas pontuais que ele
representa, slide por slide, de modo que o resultado é Markdown Marp simples,
que qualquer ferramenta Marp lê.
