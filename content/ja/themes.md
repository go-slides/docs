---
title: "テーマパッケージ"
weight: 42
description: "テーマを 1 つのファイルに：theme.hcl、CSS、フォント、アセット、kind を収めた .gsta.xz アーカイブ。信頼できない入力として検査されます。"
tags: [themes, hcl]
---

**テーマパッケージ**は `<name>.gsta.xz`（Go-Slides Theme Archive）という 1 つのファイルです。
ルートに `theme.hcl` があり、それが宣言したファイルを収めた、xz で圧縮した tar アーカイブ
（ustar または pax）です。
テーマの CSS、フォント、ロゴ、背景を収め、必要ならスライドの kind も収めることができ、
名前、バージョン、ライセンスが付くので、テーマを受け渡してインストールできます。
テーマパッケージは v0.17.1 より後で使えます。

## theme.hcl ファイル {#the-themehcl-file}

```hcl
# theme.hcl
theme {
  name        = "acme"
  version     = "1.0.0"                 # semantic version
  description = "Example theme"
  authors     = ["go-slides contributors"]
  license     = "MIT"                   # SPDX expression
  homepage    = "https://go-slides.github.io"
  extends     = "default"               # the base theme, built in or installed
  min_slides  = "0.18.0"                # oldest go-slides it works with
}

css "theme.css" {}                      # in order; sha256 = "…" is optional

font "Roboto Mono" {
  weight  = 400                         # or "100 900"
  style   = "normal"
  file    = "fonts/RobotoMono-Regular.ttf"
  license = "OFL-1.1"                   # required for a font
}

asset "logo" { file = "assets/logo.svg" }

kind "acme-title" {                     # optional slide kinds
  class    = "lead"
  paginate = "false"
}
```

`theme.hcl` は、リテラル値だけを使う HCL のネイティブ構文です。関数も変数もなく、
`file()` や `env()` も使えません。名前は小文字、数字、ダッシュで構成し、`default`、
`gaia`、`uncover` は使えません。`extends` はベースとなるテーマを指定し、`css` ブロックは
1 つにつき 1 枚のスタイルシートを順に追加し、`font` と `asset` ブロックはそれぞれ
パッケージ内のファイルを指し、`kind` ブロックはそれぞれ [kind]({{< relref "kinds" >}}) です。
`css`、`font`、`asset` ブロックの `sha256` は、パッケージを読み込むときに検査されます。

CSS は、フォントとアセットをパッケージ内のパスで、CSS ファイルからの相対パスで参照します。
`background: url(assets/logo.svg)` のようにです。パッケージを読み込むとき、こうした参照は
すべて `data:` URL になってページが自己完結し、`font` ブロックは `@font-face` になります。
パッケージには `LICENSE`、`NOTICE`、`README` ファイルを含めることもできます。
`theme.hcl` が宣言していないそれ以外のファイルは拒否されます。

## パッケージ化、インストール、使い方 {#packing-installing-and-using}

```
slides theme pack dir       write dir/theme.hcl and its files as <name>.gsta.xz
slides theme verify file    run every check on a package
slides theme install file   check it, then install it for the current user
slides theme list           the installed packages
slides theme remove name
```

インストールされたパッケージは、ユーザーの設定ディレクトリの下の `go-slides/themes` に
あります（`os.UserConfigDir()`：`~/Library/Application Support`、`$XDG_CONFIG_HOME` または
`~/.config`、`%AppData%`）。スライドデッキで次のように書くと：

```markdown
---
theme: acme
---
```

組み込みテーマと `--theme-set` の次に、インストール済みのパッケージを探します。
`--theme acme.gsta.xz` と `--theme-set acme.gsta.xz` は、インストールせずに
ファイルを直接使います。[コマンドライン]({{< relref "cli" >}})を参照してください。

パッケージの kind は `--theme file.gsta.xz` で適用され、インストール済みのパッケージなら
`--theme name` で、またはスライドデッキのフロントマターに `theme: name` と書いても適用されます。
スライドデッキ自身の kind が優先されます。

パッケージの `extends` はベーステーマを指定し、次の順で探されます。まず組み込み（または `--theme-set` で指定されたもの）、次にインストール済みのもの、最後にパッケージファイルと同じディレクトリにある兄弟の `<base>.gsta.xz` です。これにより、テーマのファミリー（1 つのベースと複数の小さな子）を、ファイルだけでそのまま試せます。エラーには探した 3 か所が示されます。パッケージのチェーンは最長 8 で、循環は拒否されます（`theme.meta`）。

## パッケージは信頼できない入力 {#a-package-is-untrusted-input}

テーマパッケージはどこからでも来る可能性があるため、パッケージを読むすべての経路
（`pack`、`verify`、`install`、`--theme`、ビルド）が同じように検査し、コード付きの
エラーで拒否します（`theme.archive`、`theme.limit`、`theme.meta`、`theme.css`、
`theme.svg`、`theme.font`、`theme.checksum`、`theme.version`）。

- **アーカイブ。** xz ストリームは 1 つだけで、辞書には上限があり、ウィンドウはデータに
  応じて広がる。tar の中では、名前に `..`、絶対パス、ドライブ文字、バックスラッシュを含めない。
  通常ファイルとディレクトリのみ（シンボリックリンク、ハードリンク、デバイス、FIFO は不可）。
  名前の重複も不可（大文字小文字を区別せず比較）。ustar と pax のヘッダーのみ（GNU 拡張も
  スパースファイルも不可）で、pax ヘッダーは最大 8 KiB。サイズはヘッダーが宣言した値ではなく、
  実際に読み取った値です。
- **上限。** アーカイブは最大 16 MiB、エントリは 256 個、1 エントリあたり 8 MiB、展開後 24 MiB、
  xz の辞書は 16 MiB、展開したストリームはアーカイブの大きさの 100 倍まで（64 KiB 以上の場合）、
  `theme.hcl` は最大 64 KiB でかっこの入れ子は
  32 段まで、CSS は 1 ファイルあたり 1 MiB、フォントや画像は 4 MiB までです。
- **CSS。** `@import` も `@import-theme` も不可（`extends` を使ってください）。`image-set()`、
  `src()` などリソースを読み込む関数も、文字列の外のバックスラッシュエスケープも、
  `</style` も不可です。`url()` は `#fragment`、`data:` の画像またはフォント、あるいは
  パッケージ内の宣言済みファイルに限り、パッケージの外に出るものは認められません。
- **SVG。** `script`、`foreignObject`、フレーム、embed、`set`、DOCTYPE、イベントハンドラーは
  不可。`href` は `#fragment` でなければならず、ファイルの外に出る `url()` も不可です。
- **フォントと画像**は解析されます。フォントはフォントでなければならず、画像は拡張子と
  一致しなければならず、それぞれにサイズの上限があります。

リポジトリの `testdata/themes/acme` が小さな例です。
