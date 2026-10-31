---
title: "go-slides ドキュメント"
linkTitle: "ホーム"
type: docs
cascade:
  type: docs
description: "純粋な Go（CGO=0）で書かれた Marp 互換のスライドデッキ。Go テンプレートによる埋め込みとスライドの kind に対応しています。"
---

**`slides` は Marp のスライドデッキを、純粋な Go だけでプレゼンテーションに変換します。**
[Marp](https://marp.app) が読むのと同じ Markdown（Marpit のディレクティブ、スライド間の `---`、`![bg]` による背景、スピーカーノート、`default`・`gaia`・
`uncover` の各テーマ）を読み込み、さらに Marp にはない機能を 2 つ追加します。

- スライドデッキ内の **[Go テンプレート]({{< relref "templates" >}})**。Markdown の解析前に評価されるため、スライドデッキが自分自身のスライドを生成できます。
- **[スライドの kind]({{< relref "kinds" >}})**（スライドの種類）。スライドは見た目を毎回繰り返し書く代わりに、自分が何であるか
  （`title`、`section`…）を宣言します。

{{< callout type="warning" >}}
**v0.14.0。** このページ群は
[go-slides/slides v0.14.0](https://github.com/go-slides/slides/releases/tag/v0.14.0)
について説明しており、`/docs/0.14/` で公開されています。
{{< /callout >}}

## 現在できること {#what-it-does-today}

| | |
| --- | --- |
| `slides build`（デフォルト） | スライドデッキを 1 つの HTML ページとして出力します。各スライドはウィンドウに合わせて拡大縮小される SVG 枠付きの `<section>` で、テーマの CSS とキーボード操作のビューアーが付きます |
| `slides build --pdf`、`--images png`、`--image jpeg` | スライドデッキを PDF（スライドごとに 1 ページ、アウトライン付き）または PNG/JPEG 画像として出力します。ブラウザーは不要です。[go-webengine](https://github.com/go-webengine) がレンダリングし、[go-pdfkit](https://github.com/go-pdfkit) が PDF を書き出します。v0.4.0 以降はベクター形式で、テキストを選択できます：[ベクター PDF]({{< relref "cli#vector-pdf" >}}) |
| 数式とコード | `$…$` と `$$…$$` は go-tex/math で組版され、フェンスで囲んだコードは go-rouge でハイライトされます。HTML、画像、PDF のいずれでも同様です：[数式とコード]({{< relref "math-and-code" >}}) |
| 絵文字 | `:smile:` や 😄 は Twemoji の画像に変換されて埋め込まれるため、ネットワークを必要とせず、すべての出力で描画されます：[絵文字]({{< relref "emoji" >}})（v0.4.0 以降） |
| `slides build --pptx` | スライドデッキを、スピーカーノートとメタデータ付きの PowerPoint ファイルとして出力します。v0.4.0 では marp-cli と同じくスライドごとに 1 枚の画像でしたが、v0.5.0 以降は **編集可能** です。テキストボックス、図形、画像を使い、フォントも埋め込まれます（[編集可能な PPTX と ODP]({{< relref "cli#editable-pptx-and-odp" >}})） |
| `slides build --odp` | 同じ内容を OpenDocument プレゼンテーション（`.odp`）として出力します |
| 見出しのフィット | `# <!-- fit --> Title` は、ページでもすべてのエクスポートでも、見出しを 1 行でスライドの幅いっぱいに表示します |
| `slides expand` | スライドデッキをプレーンな Marp Markdown として出力します。テンプレートは評価され、kind はそれが表すスポットディレクティブに置き換えられます |
| `slides notes` | スピーカーノートを出力します |
| `slides kinds`、`slides themes` | スライドデッキで指定できる名前を一覧表示します |
| `slides.yaml` | フォルダー内のスライドデッキで共有する kind、テーマ、データ、HTML ポリシー、言語を定義します |
| HCL | `---hcl` のフロントマター、`slides.hcl`、kind とデータファイルを、式と関数が使える HCL2 で記述できます：[HCL]({{< relref "hcl" >}}) |

marp-cli との違いは処理の経路です。marp-cli はヘッドレス Chrome を操作して画像や PDF をレンダリングしますが、
go-slides ではそれを Go のライブラリで行います。

## ページ一覧 {#pages}

- [はじめに]({{< relref "getting-started" >}})
- [ギャラリー]({{< relref "gallery" >}})：Playground で開ける 12 個のサンプルスライドデッキ
- [Marp との互換性]({{< relref "marp" >}})
- [テンプレート]({{< relref "templates" >}})
- [kind]({{< relref "kinds" >}})
- [HCL]({{< relref "hcl" >}})
- [数式とコード]({{< relref "math-and-code" >}})
- [絵文字]({{< relref "emoji" >}})
- [コマンドライン]({{< relref "cli" >}})
