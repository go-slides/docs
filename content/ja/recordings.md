---
title: "ターミナルの録画"
weight: 38
description: "デッキの中の asciinema 録画: HTML ページの小さなプレーヤーで再生し、PDF、PowerPoint、OpenDocument ではベクターテキストとして描きます。"
tags: [asciinema, terminal]
---

{{< callout type="warning" >}}
ターミナルの録画は v0.17.1 より後のメインブランチにあり、まだどのリリースにも入っていません。
{{< /callout >}}

[asciinema](https://asciinema.org) の録画（[asciicast v2](https://docs.asciinema.org/manual/asciicast/v2/) または [v3](https://docs.asciinema.org/manual/asciicast/v3/)）は、ファイル名が `.cast` で終わる画像です。

```markdown
![cast](demo.cast)
![The failing build](demo.cast#t=12.5)
```

## 構文 {#syntax}

ファイルはデッキのフォルダーから読み込まれます。URL、絶対パス、フォルダーの外に出るパスは拒否されます。代替テキストはスクリーンリーダー向けに録画の名前になります。`cast` だけの場合は録画自身のタイトルを意味します。`#t=12.5`（分と秒なら `#t=1:23`）は、動かせない出力が示す瞬間を選びます。省略すると録画の最後です。

## 出力ごとに表示されるもの {#outputs}

- **HTML**（ページと playground のプレビュー）: 録画は、ページに埋め込まれた go-slides の小さなプレーヤーで再生されます。再生と一時停止、速度ボタン、位置バーがあります。どこからもスクリプトを読み込まず、ネットワークも使いません。
- **PDF、PowerPoint、OpenDocument**（`--pdf`、`--pptx`、`--odp`）: `#t=` が指す瞬間のターミナルを、テーマの等幅フォントのベクターテキストとして描きます。16 色、256 色、直接指定の色、太字、斜体、下線、反転テキストに対応します。テキストは選択して検索でき、画像になることはありません。太字と斜体は、テーマの等幅フォントにある場合に描かれます（gaia テーマの Roboto Mono）。

## プレーヤー {#player}

録画が初めて画面に現れたときに一度だけ再生します。ブラウザーが動きの軽減を求めている場合（`prefers-reduced-motion`）は、一時停止の状態で始まり、`#t=` が指す瞬間を表示し、再生を押すと再生します。ボタンやキーでスライドが切り替わることはありません。

## 制限 {#limits}

録画は信頼できない入力なので、上限があります。ファイル（16 MiB）、イベント（500,000 件）、長さ（3 時間）、ターミナル（400 × 200 セル）、エミュレーターが行える処理量です。上限を超える録画や、正しい形式の asciicast ファイルではない録画は拒否されます。ビルドは続行し、スライドにはその場所に注記が表示され、問題はコード付きの警告になります: `cast.syntax`、`cast.limit`、`cast.version`、`cast.path`。

## playground では {#playground}

ファイルパネルからドロップするか、ファイル選択で、プロジェクトに `.cast` ファイルを追加します。ファイルは `media/` に入り、画像として挿入されます。デッキが読むテキストなので、プロジェクトのほかのテキストファイルと合わせて数えられ、上限は 2 MiB です。[playground](https://go-slides.github.io/playground/) で試してみてください。
