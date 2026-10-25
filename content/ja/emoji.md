---
title: "絵文字"
weight: 36
description: "絵文字のショートコードと Unicode の絵文字は、ネットワークなしで Twemoji の画像になります。HTML ページ、画像、PDF、PowerPoint、OpenDocument のいずれでも同様です。"
tags: [emoji]
---

{{< callout type="warning" >}}
絵文字は v0.4.0 で追加されました。v0.3.0 では使えません。
{{< /callout >}}

絵文字はショートコードでも、文字そのものでも書けます。

```markdown
# Launch :rocket: 🎉

Shortcodes :smile: :+1: :heart_eyes: and Unicode 😄 👍🏽 👩‍💻 🇫🇷 1️⃣.
```

いずれも Marp Core が描画するのと同じ [Twemoji](https://github.com/jdecked/twemoji) の画像になり、
Marp Core と同じマークアップで出力されます。

```html
<img class="emoji" draggable="false" alt="😄" src="data:image/svg+xml;base64,…" data-marp-twemoji=""/>
```

## ネットワーク不要 {#no-network}

Marp Core の `src` は CDN 上の URL なので、オフラインで表示したり、印刷したり、
ネットワークに接続していないマシンで PDF に変換したりすると、画像が壊れて表示されます。
go-slides は SVG そのものを `data:` URI に埋め込みます。画像データ（Twemoji 17.0.3、Emoji 17.0 の 4,009 個）は、
純粋な Go のモジュール [go-icons/twemoji](https://github.com/go-icons/twemoji) からプログラムに組み込まれています。
圧縮時のサイズは 2 MB で、スライドデッキに絵文字が含まれる場合にだけ読み込まれます。

画像は HTML ページでも、`--images`/`--image`、`--pdf`、
`--pptx`、`--odp` でも同じように描画されます。SVG のラスタライズは go-webengine が行います。4,009 個の
画像はすべて、その SVG ラスタライザーで解析・描画でき、テストでそれを確認しています。
v0.5.0 以降、このラスタライザーは [go-gfx/oksvg](https://github.com/go-gfx/oksvg)
v0.2.0 です。`clipPath`、`fill-rule: evenodd`、円弧コマンドが Chrome と同じように描画されるようになり、
その結果 Twemoji の 1 個とそのほか 211 個の描画が変わり、SVG のロゴも崩れなくなりました。

絵文字をテキストに合わせたサイズにするルール（`img[data-marp-twemoji]`、1 `em` 四方、
ベースラインに揃える）は、Marp Core と同様にテーマへ追加されます。
そのため、独自のテーマにも適用されます。

## 変換されるもの {#what-is-converted}

ルールは Marp Core と同じです。あらゆるケースを含むスライドデッキについて、
Marp Core 4.4.0 が出力する HTML と比較するテストがあります。

| | |
| --- | --- |
| `:smile:`、`:+1:`、`:heart_eyes:` | markdown-it-emoji の 1,903 個の名前。本文、見出し、リンク、表、`header`、`footer` で変換されます。コード、URL、画像の代替テキストでは変換されません |
| `😄`、`👩‍💻`、`🇫🇷`、`1️⃣`、`👍🏽` | Unicode の絵文字。本文に加えて **コード内でも** 変換されます。Marp Core はレンダリング後のコードの HTML を書き換えるためです |
| `✔`、`©`、`™`、`▶` | テキストのままです。デフォルトの表示がテキストである文字は、U+FE0F が付いた場合（`✔️`）にだけ絵文字になります。`❤` は例外です |
| `☺︎` | テキストのままです。絵文字の後の U+FE0E は、通常のグリフでの表示を求めます |
| `12:30:45`、`:not_a_name:` | 書いたとおりに残ります |

見出しの `id` とアウトラインでは、ショートコードの名前ではなく絵文字が使われます。
`# Launch :rocket:` の id は、Marp と同じく `launch-%F0%9F%9A%80` になります。

go-slides が Marp Core と異なる点では、go-slides のほうが多くを変換します。`✌️ ☝️ ✍️ 🏋️` など、
Marp Core の emoji-regex がテキストのまま残す、デフォルト表示がテキストの文字をいくつか変換するほか、
Twemoji は描画できるものの Unicode には登録されていない少数のシーケンスも変換します。

## 無効にする {#turning-it-off}

`--no-emoji`（コードでは `slides.Config{NoEmoji: true}`）を指定すると、ショートコードと
Unicode の絵文字は書いたとおりのテキストのまま残ります。これは Marp Core の
`emoji: false` に相当します。

## クレジット表記 {#attribution}

画像は Twemoji（Copyright 2020 Twitter, Inc and other contributors）で、
[CC-BY 4.0](https://creativecommons.org/licenses/by/4.0/) の下でライセンスされており、
[jdecked/twemoji](https://github.com/jdecked/twemoji) プロジェクトによって引き継がれています。
これらを表示するスライドデッキでは、フッターや最後のスライドなどでクレジットを記載してください。
ショートコードの対応表は markdown-it-emoji のもの（MIT）です。
