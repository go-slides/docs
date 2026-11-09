---
title: "ビデオ"
weight: 39
description: "プレゼンテーション内のビデオ: HTML ページではブラウザーが再生し、PDF ではリンク付きのポスターフレーム、PowerPoint と OpenDocument には埋め込まれます。"
tags: [video]
---

ビデオは、ファイル名が `.mp4`、`.m4v`、`.webm` で終わる画像です:

```markdown
![video](clip.mp4)
![The demo](clip.mp4#t=12.5)
![demo loop muted](clip.mp4)
```

## 構文 {#syntax}

ファイルはプレゼンテーションのフォルダー、または `http` か `https` の URL から読み込まれます。絶対パス、フォルダーの外に出るパス、その他のスキームは拒否されます。`#t=12.5`、または分と秒で書く `#t=1:23` は、ポスターフレームを選びます。再生前と、ビデオを再生できない出力で表示される画像です。メディアフラグメント形式の `#t=10,20` も受け付け、その開始値が使われます。指定がなければ、ポスターは最初のフレームです。

代替テキスト内のキーワードは、Marp の画像キーワードのように書きますが go-slides 独自のもので、再生方法を決めます: `autoplay`、`loop`、`muted`。`autoplay` は `muted` と併用したときだけ有効で、そうでなければ警告とともに無視されます。代替テキストのそれ以外の語は、アクセシブルなラベルです。

## 出力ごとの表示 {#outputs}

- **HTML** (ページと playground のプレビュー): `<video controls playsinline preload="metadata" poster=...>` 要素。ポスターはデコードしたフレームで、JPEG の data URI として埋め込まれます。ビデオが自分から音を出して再生されることはありません。ブラウザーが動きの軽減を求めると (`prefers-reduced-motion`)、`autoplay` と `loop` は取り除かれ、ビデオは一時停止されます。自動再生される (ミュートの) ビデオは、半分以上が見えている間だけ再生され、そのスライドが表示されていないときは一時停止します。
- **PDF** (`--pdf`): ポスターフレームを画像として入れ、ビデオファイルへのリンクを付けます。リンクは、プレゼンテーションに書いた相対パス、または `http` か `https` のソースなら URL です。再生はリンクを開いたときだけ行われます。PDF には `RichMedia` 注釈を入れません。それは Acrobat でしか再生できないため、あえて使いません。
- **PowerPoint** (`--pptx`): ビデオはメディアパート (`ppt/media/slidesvideoN.mp4|webm`) としてポスターとともに埋め込まれ、`a:videoFile` に `p14:media` 拡張を加えた形になります。Keynote: ムービーとして読み込み、再生します (確認済み)。PowerPoint: まだ確認していません。`autoplay`、`loop`、`muted` キーワードは PowerPoint と OpenDocument には引き継がれません。
- **OpenDocument** (`--odp`): ビデオは `Video/` の下に `draw:plugin` (ODF 1.3 パート 3、10.4.8) として、ポスターの `draw:image` の前に埋め込まれ、マニフェストに列挙されます。メディアバックエンドがそのコーデックに対応している Impress で再生されます (未確認)。
- **URL ソース**: HTML ページは URL からビデオをストリーミングします。エクスポートでは URL へのリンク付きの中立的なカードが表示されます。ビルド時には何も取得せず、PowerPoint や OpenDocument にも何も埋め込まれません。

## コンテナーとポスター {#codecs}

MP4 (ISO BMFF) と WebM (Matroska) のファイルは Go で読み込まれ、ffmpeg は実行されません。ポスターは H.264 (`avc1` または `avc3`) についてのみデコードされます。これは go-avkit が持つ唯一のビデオデコーダーです。ほかのコーデック (H.265、VP8、VP9、AV1) も HTML では再生され、PowerPoint と OpenDocument に埋め込まれますが、ポスターは再生三角形の付いた中立的な暗いカードになり、警告 `video.codec` が出力されます。

## 上限 {#limits}

ビデオは信頼できない入力なので、上限があります: ファイル (64 MiB)、ボックスまたは要素 (262,144 個) とその入れ子 (16 段)、フレーム (4096 × 4096)、ポスターのデコード (直前のキーフレームから最大 600 サンプル、1 サンプル 8 MiB)。キーフレームから遠すぎるポスターは、カードと警告 `video.limit` になります。

拒否されたビデオは、スライド内で `[video not shown: ...]` というメモに置き換えられます。ビルドは続行され、問題はコード付きの警告になります:

- `video.syntax`: MP4 や WebM のファイルではない、ビデオトラックがない、または時刻が無効;
- `video.limit`: 上記の上限のいずれかを超えた;
- `video.codec`: このコーデックのポスターをデコードできない;
- `video.path`: 絶対パス、プレゼンテーションのフォルダーの外に出るパス、または `http` と `https` 以外のスキーム;
- `file.missing`: ファイルが存在しない。

## playground では {#playground}

リリース後は、ほかのファイルと同じようにプロジェクトにビデオを追加できます。そのバイトは blob URL としてブラウザー内にとどまり、アップロードされません。[playground](https://go-slides.github.io/playground/) で試してください。
