---
title: "はじめに"
weight: 10
description: "最初のスライドデッキを HTML にビルドし、Marp Markdown に展開し直します。"
tags: [start]
---

## インストール {#install}

Go 1.27.3 以降が必要です。

```sh
go install github.com/go-slides/slides/cmd/slides@latest
```

`slides` は cgo を使わない単一の静的 Go バイナリです。ほかにインストールするものはなく、
Node も Chrome も不要です。このページは v0.15.0 について説明しています。

## 最初のスライドデッキ {#a-first-deck}

Marp のスライドデッキであれば、そのまま使えます。次の内容を `talk.md` として保存してください。

```markdown
---
marp: true
theme: default
paginate: true
footer: "{{ .talk }}"
data:
  talk: Pure Go slides
  topics: [Why, How, Next]
---

<!-- _kind: title -->
# {{ .talk }}

{{ range .topics }}
---

# {{ . }}
{{ end }}
```

フロントマターは Marp のものに `data:` を加えたものです。本文は Go テンプレートです。
`range` はトピックごとに `---` で始まるスライドを 1 枚ずつ書き出し、最初の
スライドは組み込みの kind（スライドの種類）`title` です（ヘッダー、フッター、ページ番号なし）。

## ビルドする {#build-it}

```sh
slides talk.md                 # writes talk.html beside the deck
slides -o slides.html talk.md  # or somewhere else; -o - for standard output
```

コマンドを付けずに `slides` を実行すると `slides build` になります。フラグはスライドデッキの前にも後にも置けます。
HTML をブラウザーで開き、キーボードでスライド間を移動します。

## 画像、PDF、PPTX、ODP にエクスポートする {#export-images-pdf-pptx-and-odp}

```sh
slides talk.md --pdf                  # one page per slide, with an outline
slides talk.md --images png           # talk.001.png, talk.002.png, …
slides talk.md --image jpeg           # the first slide only
slides talk.md -o out/slides.pdf      # the extension picks the format; out/ is created
slides talk.md --images jpeg --jpeg-quality 90
slides talk.md --images png --image-scale 2   # 2560×1440 for 16:9
slides talk.md --pdf --allow-local-files      # images from under the deck's directory
slides talk.md --pptx                         # or -o talk.pptx, with notes
slides talk.md --odp                          # or -o talk.odp: the same, as OpenDocument
```

ブラウザーは一切使いません。go-webengine がスライドをレンダリングし、go-pdfkit が
PDF を書き出します。ローカルの画像は、`--allow-local-files` を指定しない限り読み込まれません。
v0.5.0 以降、PPTX や ODP のスライドは編集可能です。marp-cli では画像になるところを、テキストボックス、図形、
画像で構成し、フォントも埋め込みます。`--pptx-raster`
を指定すると画像のままにできます（[編集可能な PPTX と ODP]({{< relref "cli#editable-pptx-and-odp" >}})）。v0.4.0 以降、PDF はベクター形式で、テキストを
選択できます（v0.3.0 ではページごとに 1 枚の画像を書き出していました）。詳しくは [コマンドライン]({{< relref "cli#vector-pdf" >}}) を参照してください。

## 展開する {#expand-it}

```sh
slides expand talk.md
```

は、同じスライドデッキを **プレーンな Marp Markdown** として出力します（`main` の
`4ae5c49`、2026-10-18 時点の出力）。

```markdown
---
marp: true
theme: default
paginate: true
footer: "Pure Go slides"
---
<!--
_class: "lead"
_footer: ""
_paginate: "false"
-->


# Pure Go slides


---

# Why

---

# How

---

# Next
```

`range` が生成した 4 枚のスライド、値が埋め込まれたフッター、そして
`title` kind はそれが表すスポットディレクティブに置き換えられています。marp-cli や VS
Code 拡張機能など、Marp を読めるものならこの先を引き継げるので、go-slides の
追加機能を使って書いたものが囲い込まれることはありません。

## 次のステップ {#next}

- [Marp との互換性]({{< relref "marp" >}})：Marp と同じように読み込まれるもの。
- [テンプレート]({{< relref "templates" >}})と[kind]({{< relref "kinds" >}})：2 つの追加機能。
- [コマンドライン]({{< relref "cli" >}})：すべてのコマンドとフラグ。
