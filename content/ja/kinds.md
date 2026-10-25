---
title: "kind"
weight: 40
description: "スライドは自分が何であるかを宣言します。kind はローカルディレクティブのプリセットに名前を付けたものです。"
tags: [kinds]
---

**kind**（スライドの種類）は、ローカルディレクティブ（ヘッダー、フッター、
ページ番号、クラス、背景）のプリセットに名前を付けたものです。スライドは見た目を毎回繰り返し書く代わりに、
自分が何であるか（`title`、`section`）を宣言します。こうすれば、スライドデッキの見た目を
1 か所で変更できます。

## kind を使う {#using-a-kind}

```markdown
<!-- kind: section -->
# Part two

---

<!-- _kind: blank -->
![bg](photo.jpg)
```

- `kind:` はローカルディレクティブと同じように、そのスライドと以降のスライドに適用されます。
- `_kind:` はスポットディレクティブと同じように、そのスライドだけに適用されます。
- kind を切り替えると、まず直前の kind が設定した内容が **クリア** されます。そのため、クラスや
  フッターがある kind から次の kind に漏れ出すことはありません。
- スライド自体に書いたディレクティブは、kind より **優先** されます。

## 組み込みの kind {#built-in}

| kind | 設定内容 |
| --- | --- |
| `title` | `class: lead`、ページ番号なし、ヘッダーなし、フッターなし |
| `section` | `class: lead` |
| `content` | なし（スライドデッキ自体のデフォルト） |
| `blank` | ページ番号なし、ヘッダーなし、フッターなし |

どれも再定義できます。

## kind を定義する {#defining-kinds}

フロントマターに記述します。

```yaml
---
marp: true
kinds:
  chapter:
    class: invert
    backgroundColor: "#1e1b4b"
    paginate: false
  quote:
    extends: blank
    class: quote
---
```

`extends` を使うと、別の kind のディレクティブを土台にして kind を定義できます。kind に書けるのは
ローカルディレクティブと `extends` です。

## kind を共有する {#sharing-kinds}

チームで使う kind は、すべてのスライドデッキが参照する 1 つのファイルにまとめます。

```yaml
# kinds.yaml
chapter:
  class: invert
  paginate: false
```

```yaml
---
marp: true
kinds: kinds.yaml        # or a list of files
---
```

または、コマンドラインで `--kinds kinds.yaml` を指定するか、`slides.yaml` に `kinds:` を書きます。
適用順は、組み込みの kind、次にコマンドラインと設定ファイル、最後にスライドデッキ自体の `kinds:` です。

`slides kinds talk.md` は、スライドデッキで使える kind と、それぞれが設定する内容を一覧表示します。

## go-slides から離れる {#leaving-go-slides}

`slides expand` は、すべての kind をそれが表すスポットディレクティブとして
スライドごとに書き出します。その結果はプレーンな Marp Markdown になり、どの Marp ツールでも
読み込めます。
