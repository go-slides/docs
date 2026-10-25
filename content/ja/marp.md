---
title: "Marp との互換性"
linkTitle: "Marp"
weight: 20
description: "slides が読む Markdown は Marpit と Marp Core のものそのもので、その方言ではありません。"
tags: [marp]
---

[Marp](https://marp.app) 向けに書かれたスライドデッキは、そのまま go-slides のスライドデッキです。以下の
ルールは [Marpit](https://marpit.marp.app) と
[Marp Core](https://github.com/marp-team/marp-core) のものです。go-slides が
異なる点は、このページで明記します。

## スライド {#slides}

水平線（Marpit が分割するのと同じく、1 行に単独で書いた `---`）で
新しいスライドが始まります。`headingDivider` ディレクティブで指定したレベルの見出しでも
新しいスライドが始まります。最初の 2 つの `---` の間にあるフロントマターには、グローバルディレクティブと
スライドデッキのデータを書きます。

## ディレクティブ {#directives}

| 種類 | 書く場所 | 適用範囲 |
| --- | --- | --- |
| グローバル | フロントマター、または任意の場所の `<!-- theme: gaia -->` | スライドデッキ全体：`theme`、`style`、`headingDivider`、`lang`… |
| ローカル | `<!-- paginate: true -->` | そのスライドと以降のスライド |
| スポット | `<!-- _class: lead -->`（アンダースコア付き） | そのスライドのみ |

ローカルディレクティブは Marpit のものです：`paginate`、`header`、`footer`、`class`、
`backgroundColor`、`backgroundImage`、`backgroundPosition`、
`backgroundRepeat`、`backgroundSize`、`color`。ディレクティブは Marp Core と同じく
緩い YAML として解析されます。`--strict-yaml` を指定すると、素の Marpit と同じように
解析されます。

## 画像 {#images}

`![bg](image.jpg)` は画像をスライドの背景にします。Marp のキーワードが使えます：
`bg contain`、`bg cover`、`bg 50%`、分割背景の `bg left`/`bg right`、複数の `bg` 画像の横並び、
画像フィルター（`blur`、`sepia`、`grayscale`…）。インライン画像のサイズ指定（`w:200px`、`h:`）も
Marp と同じように解釈されます。

## スライドの周辺要素 {#around-the-slide}

- **ヘッダーとフッター**：Markdown で書く `header:` と `footer:` のディレクティブ。
- **ページ番号**：`paginate: true`。
- **スピーカーノート**：ディレクティブではない HTML コメントはノートになります。
  `slides notes` で出力できます。
- **段階的に表示するリスト**：`*`（または `1)`）で書いたリストは、Marp と同じく
  項目が 1 つずつ表示されます。
- **`<style scoped>`**：1 枚のスライドだけに適用される CSS。
- **生の HTML**：デフォルトでは Marp Core の許可リストに従います。`--html` を指定するとすべて許可されます。
- **絵文字**：`:shortcodes:` と Unicode の絵文字は、Marp Core と同じく Twemoji の画像になります。
  ただし、CDN から読み込むのではなく埋め込まれます。詳しくは
  [絵文字]({{< relref "emoji" >}}) を参照してください。

## テーマ {#themes}

`default`、`gaia`、`uncover` が組み込まれており、`theme:` で選択します。
独自のテーマは、Marp の `/* @theme name */` コメントを含む CSS ファイルです。
`--theme-set` で指定するか、`slides.yaml` の `themes:` に列挙します。

## go-slides が異なる点 {#where-go-slides-differs}

- **テンプレート。** 本文中の `{{ }}` は Go テンプレートのアクションです（
  [テンプレート]({{< relref "templates" >}}) を参照）。コードブロックの外で `{{` をテキストとして表示する
  スライドデッキでは、`--no-template` を指定するか、Go テンプレートの方法でそのテキストをエスケープする必要があります。
  コードブロックとインラインコードはそのまま残ります。
- **kind。** `kind:` と `_kind:` は Marp にはないディレクティブです（
  [kind]({{< relref "kinds" >}})（スライドの種類）を参照）。`slides expand` はこれらを、それが表す
  ディレクティブに置き換えるので、展開後のスライドデッキはプレーンな Marp になります。
- **出力。** HTML、PNG、JPEG、PDF、PPTX、ODP を、ヘッドレスブラウザーではなく Go のライブラリで
  出力します（既知の制限については [コマンドライン]({{< relref "cli#images-pdf-pptx-and-odp" >}})
  を参照）。v0.4.0 では marp-cli と同じく PPTX の中身は画像でしたが、
  v0.5.0 以降のスライドは編集可能なテキストボックスと図形で構成されます。`--pptx-raster` を指定すると
  画像のままにできます。
- **見出しのフィット。** `# <!-- fit --> Title` に対応しています。ページでもすべてのエクスポートでも、
  見出しが 1 行でスライドの幅いっぱいに表示されます。
- **数式とコード。** `$…$` と `$$…$$` は Marp Core の区切り文字のルールに従い、
  go-tex/math で組版されます。フェンスで囲んだコードは go-rouge によって
  highlight.js のクラス名でハイライトされます（[数式とコード]({{< relref "math-and-code" >}}) を参照）。
