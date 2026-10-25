---
title: "テンプレート"
weight: 30
description: "スライドデッキでの Go text/template による埋め込み：値、スライドの生成、ページ番号、インクルード。"
tags: [templates]
---

スライドデッキは Go の [`text/template`](https://pkg.go.dev/text/template) であり、
Markdown の解析 **前** に評価されます。Go のテンプレート言語が備えるもの
（アクション、パイプライン、`if`、`range`、`with`、変数）は、すべて
スライドデッキの中で使えます。

## データ {#data}

テンプレートのドットには、フロントマターのキー、次にその
`data:` マッピングのキー、さらにコマンドラインまたは `slides.yaml` で与えたデータが、
この順に入ります。後のものが前のものを上書きします。

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

コマンドラインからは次のように指定します。

| フラグ | 与えるもの |
| --- | --- |
| `--data file.yaml` | データを記述した YAML または JSON ファイル（複数回指定可） |
| `--set key=value` | 1 つの値（複数回指定可） |

## スライドを生成する {#generating-slides}

本文全体が 1 つのテンプレートなので、`---` を書き出す `range` はスライドを書き出すことになります。

```markdown
{{ range .releases }}
---

# {{ .version }}

{{ range .changes }}- {{ . }}
{{ end }}
{{ end }}
```

## ページ番号と総数 {#page-and-total}

`{{ page }}` と `{{ total }}` は、そのスライドの番号とスライドの総数です。
これらはスライドデッキが分割されて初めて決まるため、その時点で埋め込まれます。
生成されたスライドデッキでは、生成されたスライドも数えられます。

```markdown
<!-- footer: "{{ page }} / {{ total }}" -->
```

## インクルード {#include}

`{{ include "parts/agenda.md" }}` は、スライドデッキからの相対パスで指定した別のファイルを
挿入します。

## コードはそのまま {#code-is-left-alone}

フェンスで囲んだコードブロックとインラインコードは、テンプレートの実行前に取り分けられ、
そのまま元に戻されます。そのため、Go テンプレートについてのスライドでも
`{{ .Name }}` をエスケープせずに表示できます。

## 無効にする {#turning-it-off}

`--no-template` を指定すると、スライドデッキをプレーンな Marp として読み込みます。`{{` はただのテキストになります。
