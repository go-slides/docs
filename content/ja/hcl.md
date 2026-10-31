---
title: "HCL"
weight: 45
description: "スライドデッキで YAML を使う場所なら、どこでも HCL2 が使えます。HCL のフロントマター、slides.hcl、kind とデータのファイルで、式と関数を利用できます。"
tags: [hcl, configuration]
---

{{< callout type="warning" >}}
HCL は v0.3.0 で追加されました（[go-slides/slides#13](https://github.com/go-slides/slides/pull/13)）。
{{< /callout >}}

`slides` は YAML を読むすべての場所で、
[HCL2](https://github.com/hashicorp/hcl/blob/main/hclsyntax/spec.md) のネイティブ構文も読み込みます。
HCL2 の式と関数を使うためです。

| 場所 | YAML | HCL |
| --- | --- | --- |
| フロントマター | `---` … `---` | `---hcl` … `---` |
| 設定 | `slides.yaml` | `slides.hcl`（こちらが先に検索されます） |
| kind（スライドの種類）のファイル | `--kinds kinds.yaml`、`kinds: kinds.yaml` | `*.hcl` という名前の任意のファイル |
| データファイル | `--data data.yaml` または `.json` | `*.hcl` という名前の任意のファイル |

ディレクティブのコメント（`<!-- class: invert -->`）は Marp と同じく YAML のままで、
スライドデッキの本文も [Go テンプレート]({{< relref "templates" >}}) のままです。

## HCL で書くフロントマター {#a-front-matter-in-hcl}

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

- **属性** は、同じ名前の YAML のキーに対応します。Marp のディレクティブ
  （`theme`、`paginate`、`header`、`footer`、`class`、`backgroundColor`…）、
  メタデータ（`title`、`author`…）、`template`、`delims`、`kinds`（ファイルの
  パス）、そしてオブジェクトである `data` です。
- **`kind "<name>" { … }`** ブロックは [kind]({{< relref "kinds" >}}) を定義します。
  中にはローカルディレクティブと `extends` を書けます。
- **値** は、YAML の場合と同じ文字列の形でディレクティブに渡されます。
  `true` は `"true"` になり、クラスのリストはリストのままです。そのため、どちらの構文で設定しても
  ディレクティブは同じように動作します。`null` を指定すると、そのディレクティブは未設定のままになります。
- **`data`** は型を保持します。テンプレート内では整数は `int`、
  それ以外の数値は `float64` になり、リストとオブジェクトはリストとマップのままです。そのため
  `{{ add .quarter 1 }}` が動作します。
- HCL の値は確定値です。Go テンプレートがその値に対してもう一度実行されることはありません。

## 変数 {#variables}

`data` はテンプレートのデータです。フロントマター自体の `data` に、
`slides.hcl` のもの、`--data` のファイル、`--set` の順に重ねられます。`{{ .company }}` が使える場所なら、
どこでも `"${data.company}"` が使えます。

フロントマターの属性は `data.*` を参照できますが、**属性どうしは参照できません**。
`footer = title` はエラーになります。スポットディレクティブ（`_class`…）は HCL の
識別子ではないので、コメントに書いてください。

## 関数 {#functions}

| | |
| --- | --- |
| `page()`、`total()` | `{{ page }}` や `{{ total }}` と同じく、ページ番号と最終ページの番号です。フッターにページ番号を入れられます |
| `now()`、`timestamp()` | 現在時刻（RFC 3339、ローカル時刻と UTC）です。`formatdate` で書式化できます |
| `env(name)` | 環境変数です。未設定の場合は `""` になります |
| `file(path)` | 最大 1 MiB の UTF-8 テキストファイルです。パスはスライドデッキまたは設定ファイルからの相対パスです |
| 数値 | `abs ceil floor log max min parseint pow signum` |
| 文字列 | `chomp format formatlist join lower regex regexall regexreplace replace split strlen strrev substr title trim trimprefix trimspace trimsuffix upper` |
| コレクション | `chunklist coalesce coalescelist compact concat contains distinct element flatten keys length lookup merge range reverse slice sort values zipmap` |
| エンコーディング | `csvdecode jsondecode jsonencode` |
| 日付 | `formatdate timeadd` |

最初の 4 行以外は、[go-cty](https://github.com/zclconf/go-cty) の
標準ライブラリです。ネットワークにアクセスする関数はありません。

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

未知の属性は、黙って無視されるのではなくエラーになります。kind のファイルには
`kind` ブロックを書きます。それ以外の属性は無視されるので、`slides.hcl` を kind のファイルとしても
使えます。データファイルの属性は、そのままデータのキーになります。

```hcl
# data.hcl
company = "ACME"
year    = 2026
motto   = trimspace(file("motto.txt"))
```

## エラー {#errors}

エラーは HCL の診断メッセージで、スライドデッキのファイル内の位置が示されます。
スライドデッキの 4 行目に `footer = data.compnay` と書くと、次のようになります。

```
deck.md:4:14: Unsupported attribute: This object does not have an attribute named "compnay".
```

## Marp と `slides expand` {#marp-and-slides-expand}

Marp も `---hcl` をフロントマターの開始として読みます（Marp のフロントマターの
ルールは、ダッシュ 3 つで始まる最初の行であれば何でも開始とみなすためです）。ただし、そのブロックは
YAML のマッピングではないので、内容は何も適用されません。スライドデッキは Marp でもレンダリングされますが、
ディレクティブは反映されません。この挙動は、Marp Core 自身の出力と照らし合わせるテストで固定しています。

`slides expand` はフロントマターを YAML に書き戻すので、展開後のスライドデッキは
プレーンな Marp になり、`slides` が元のスライドデッキをレンダリングしたときと同じようにレンダリングされます。

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

## ライセンス {#licences}

HCL のサポートでは、[hashicorp/hcl](https://github.com/hashicorp/hcl)
（MPL-2.0、ライブラリとして無改変で使用）と
[zclconf/go-cty](https://github.com/zclconf/go-cty)（MIT）をリンクしています。
