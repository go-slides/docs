---
title: "ギャラリー"
weight: 12
description: "テーマ、kind、テンプレート、HCL、数式、画像、コード、絵文字、トランジション、そして完全な講演まで、それぞれ 1 つのことをうまく示す 12 個の小さなスライドデッキ。どれも Playground で開けます。"
tags: [gallery, examples]
---

go-slides の
[`examples/`](https://github.com/go-slides/slides/tree/main/examples) フォルダーにある
12 個の小さなスライドデッキで、それぞれが 1 つのことをうまく示しています。サムネイルは
各スライドデッキの最初のスライドで、このサイトのビルド時に go-slides 自身がレンダリングしています。CI はすべての
スライドデッキを HTML と PDF にビルドしているので、どれも壊れることはありません。

**Playground で開く** を押すと、サーバーに何も送信しないリンクで、スライドデッキがブラウザーに読み込まれます。
`deck.md` 以外のファイル（kind（スライドの種類）のファイル、インクルードする部品、画像）を含むスライドデッキはリンクでは渡せません。
その場合は **.zip** をダウンロードして Playground のファイルパネルからインポートするか、展開して
`slides build --allow-local-files deck.md` でビルドしてください。

画像は go-slides の作者によるオリジナルの SVG で、
[CC0 1.0](https://creativecommons.org/publicdomain/zero/1.0/) で公開されており、
それを使うスライドデッキにはその旨が記載されています。

{{< gallery >}}
