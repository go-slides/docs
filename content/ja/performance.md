---
title: "パフォーマンス"
weight: 60
description: "go-slides がスライドデッキをどれだけ速く、どれだけのメモリでビルドするか。同じマシン上で marp-cli と比較して計測しました。"
tags: [performance]
---

{{< callout type="info" >}}
以下の PPTX の数値は v0.5.0、ベクター PDF の数値は v0.4.0 のものです。
最初の表は v0.3.0（その前のリリース）のものです。v0.3.0 はスライドを並行してレンダリングし、
フォントの解析を 1 回だけ行います（[go-slides/slides#11](https://github.com/go-slides/slides/pull/11)）。
{{< /callout >}}

go-slides と [marp-cli](https://github.com/marp-team/marp-cli) 4.5.1 を、
同じスライドデッキに対してサブプロセスとして実行しました。マシンは Apple M4 Max（16 CPU）、
1 分間のロードアベレージは 8 未満で、ウォームアップの後に交互に 5 ラウンド計測しています。
時間は実時間の中央値、メモリはプロセスツリー全体のピークです
（marp-cli は Chrome を操作するため、Chrome のプロセスも含めています）。

| スライドデッキ | 出力 | go-slides | marp-cli | go-slides のメモリ | marp-cli のメモリ |
| --- | --- | ---: | ---: | ---: | ---: |
| 100 枚 | HTML | 36 ms | 315 ms | 47 MiB | 140 MiB |
| 100 枚 | PDF | 1.14 s | 1.88 s | 242 MiB | 1,509 MiB |
| 100 枚 | PNG（全スライド） | 0.97 s | 5.44 s | 225 MiB | 1,354 MiB |
| 100 枚 | PPTX | 3.39 s | 8.03 s | 530 MiB | 1,373 MiB |
| gaia、20 枚 | PDF | 283 ms | 1.57 s | 154 MiB | 1,315 MiB |
| 数式とコード、20 枚 | PNG（全スライド） | 250 ms | 2.32 s | 168 MiB | 1,286 MiB |

## ベクター PDF {#vector-pdf}

v0.4.0 では、PDF はベクターで描画されます（
[コマンドライン]({{< relref "cli#vector-pdf" >}}) を参照）。同じマシンと方法で、
2026-10-19 に計測しました。

| スライドデッキ | go-slides | marp-cli | v0.3.0 | go-slides のサイズ | marp-cli のサイズ | v0.3.0 のサイズ |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| 100 枚 | 306 ms | 1,763 ms | 1,147 ms | 758 KiB | 837 KiB | 10.0 MiB |
| gaia、20 枚 | 93 ms | 1,546 ms | 285 ms | 32 KiB | 66 KiB | 582 KiB |
| 数式とコード、20 枚 | 140 ms | 1,643 ms | 294 ms | 124 KiB | 335 KiB | 638 KiB |

100 枚の PDF のメモリ使用量は 158 MiB で、marp-cli は 1,512 MiB です。CPU
時間は 0.9 s 対 2.7 s です。

100 枚の PDF についての v0.4.0 の数値は、2026-10-19 の 21:29 に、
1 分間のロードアベレージ 7.7 の状態で 3 ラウンド計測したものです：go-slides は 306 ms、758
KiB、158 MiB、marp-cli は 1,838 ms、837 KiB、1,512 MiB でした。

## 編集可能な PPTX {#editable-pptx}

v0.5.0 以降、PPTX はテキストボックス、図形、画像で構成されます（
[コマンドライン]({{< relref "cli#editable-pptx-and-odp" >}}) を参照）。
2026-10-19 の 22:51 に、1 分間のロードアベレージ 7.1 の状態で 3 ラウンド計測しました。

| スライドデッキ | go-slides | marp-cli | go-slides のサイズ | marp-cli のサイズ | go-slides のメモリ | marp-cli のメモリ |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| 100 枚 | 348 ms | 7,795 ms | 760 KiB | 7.2 MiB | 204 MiB | 1,370 MiB |
| gaia、20 枚 | 112 ms | 4,608 ms | 893 KiB | 4.8 MiB | | |
| 数式とコード、20 枚 | 207 ms | 2,780 ms | 842 KiB | 1.5 MiB | | |

100 枚のスライドデッキでは、v0.4.0 は 4.7 s、4.8 MiB、540 MiB でした。marp-cli の
PPTX の中身は画像です。編集可能なスライドデッキのサイズの大部分は、埋め込まれたフォントが占めています。

## 速い理由 {#why-it-is-fast}

- **起動するブラウザーがない。** marp-cli は PDF、画像、
  PPTX へのエクスポートのたびに Chrome を起動します。go-slides は
  [go-webengine](https://github.com/go-webengine/engine) を使って自身の
  プロセス内でレンダリングします。
- **スライドは並列にレンダリングされます。** デフォルトでは 4 枚ずつで、各画像は
  順番が来しだい書き出されます。そのため、長いスライドデッキが画像としてメモリに保持され続けることはありません。
- **フォントの解析はプロセスごとに 1 回** だけで、スライドごとには行いません。

## 優位でない点 {#what-it-does-not-win}

- **v0.3.0 での PDF のサイズ。** PDF の各スライドが画像だったため、100 枚のスライドデッキでは
  10 MiB 対 837 KiB でした。上記のベクター PDF は marp-cli のものより小さくなっています。
- **大きなスライドデッキでの CPU 時間。** 4 つのワーカーは、実時間から見えるより多くの CPU を消費します。
  v0.3.0 の 100 枚の PDF では 3.2 s で、marp-cli の 2.7 s を上回りました
  （ベクター PDF では 0.9 s）。
- **marp-cli のメモリは上限値です。** Chrome のプロセスはページを共有しており、合計ではそれを
  プロセスごとに重複して数えています。単一プロセスで最大のものは 220–320 MiB です。

計測方法、すべてのスライドデッキとモード、数値を再現するためのコマンドは
[BENCHMARKS.md](https://github.com/go-slides/slides/blob/main/BENCHMARKS.md)
にあります。
