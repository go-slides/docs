---
title: "VS Code"
weight: 55
description: "VS Code 向けの go-slides 拡張機能。Marp for VS Code が Marp に対して行うのと同じく、プレビュー、発表、問題の表示、補完、エクスポートを提供します。"
tags: [vscode, editor]
---

{{< callout type="info" >}}
拡張機能は [go-slides/vscode](https://github.com/go-slides/vscode) で、
独自にバージョン管理されています（v0.4.0）。レンダリングには go-slides v0.9.0（
`editor` パッケージ、スクロール同期用のソース行、一覧表示と
発表者ビュー）を使います。インストールには、GitHub リリースに添付された `go-slides-0.4.0.vsix`
を使います。VS Code Marketplace と Open VSX での公開は
準備中で、まだ公開されていません。
{{< /callout >}}

この拡張機能は、go-slides 版の
[Marp for VS Code](https://github.com/marp-team/marp-vscode) です。スライドデッキは
VS Code の Markdown プレビューにスライドとして表示され、エラーは「問題」パネルに表示され、
ディレクティブ、kind（スライドの種類）、テンプレート関数は入力中に補完されます。

## インストール {#install}

[リリース](https://github.com/go-slides/vscode/releases) から
`go-slides-<version>.vsix` をダウンロードし、
*Extensions: Install from VSIX...* を実行するか、次のコマンドを実行します。

```sh
code --install-extension go-slides-0.4.0.vsix
```

プレビューにはほかに何も必要ありません。エクスポートには `slides` コマンドが必要です。

```sh
go install github.com/go-slides/slides/cmd/slides@latest
```

## プレビュー {#preview}

スライドデッキを開き、*Markdown: Open Preview to the Side* を実行します。スライドデッキとは、
フロントマターに `slides: true` と書かれた Markdown ファイルです。

```markdown
---
slides: true
paginate: true
data:
  who: World
---

# Hello {{ .who }}

---

<!-- _kind: title -->
# Page {{ page }} of {{ total }}
```

`marp: true` のファイルもプレビューされます。ただし、Marp for VS Code がインストールされている場合は除きます。
両方の拡張機能が同じドキュメントをレンダリングしてしまうためです。この挙動は設定
`goSlides.preview.marpDocuments` で決まります：`auto`（デフォルト）、`always`、
`never`。

プレビューは、WebAssembly にコンパイルされて拡張機能に読み込まれる go-slides そのものです。
バイナリは不要で、[vscode.dev](https://vscode.dev) でも動作します。スライドデッキが使うもの
（スライドデッキと同じ場所の `slides.yaml` または `slides.hcl`、そこで指定された
[kind]({{< relref "kinds" >}}) のファイル、[テンプレート]({{< relref "templates" >}}) がインクルードするファイル、
テーマの CSS）をワークスペースから読み込みます。読み込むのは、
スライドデッキのワークスペースフォルダー内のものだけです。

## スクロール同期 {#scroll-sync}

拡張機能の v0.2.0（go-slides v0.7.0）以降、エディターとプレビューは、
通常の Markdown と同じく双方向に連動してスクロールします：
プレビューはエディターのカーソル位置にあるスライドとブロックに追従し、
プレビューをスクロールするとエディターが移動します。設定は VS Code 自体のもの、
`markdown.preview.scrollPreviewWithEditor` と
`markdown.preview.scrollEditorWithPreview` です。

行はテンプレートを通して追跡されます。`{{ range }}` が生成したスライドは
そのテンプレートの行に追従し、`{{ include }}` で取り込まれた内容は
インクルードした行に追従します。`headingDivider` で分割されたスライドデッキは、
見出しの位置で同期します。

## 発表 {#presenting}

拡張機能の v0.4.0（go-slides v0.9.0）以降、VS Code から離れずに
スライドデッキで発表できます。スライドデッキをエディターで開いた状態で、コマンドパレット
またはエディターの `…` メニューから実行します。

| コマンド | 開くもの |
|---|---|
| *go-slides: Open Overview* | すべてのスライドをグリッドに並べたパネル。スライドを選ぶとエディターにそのソースが表示されます。`Esc` でパネルを閉じます |
| *go-slides: Start Presenter View* | スライドと、その横の発表者ビュー：次のステップ、スピーカーノート、タイマー |

パネルはエディターの横に配置され、互いに連動します：どれかでスライドを切り替えると、
ほかのパネルも切り替わります。スライドデッキの編集にも追従します。スライドのパネルで `p`
キーを押すと、発表者ビューが開きます。

別の画面で発表するには、*View: Move
Editor into New Window* でスライドのパネルをその画面に移動します。

各パネルは、サンドボックス化されたフレーム（`allow-same-origin` なし）で、
`default-src 'none'` のポリシーの下でスライドデッキを実行します：フレームが読み込むのは画像、メディア、スタイル、フォントだけで、
実行するのは go-slides がページに組み込んだ、ハッシュで指定されたスクリプトだけです。
スライドデッキ内に書かれたスクリプトは実行されません。

## 問題、補完、ホバー {#problems-completion-hover}

- **問題**：テンプレート、YAML、[HCL]({{< relref "hcl" >}}) のエラーを、スライドデッキ内または
  エラー元の kind のファイル内の、該当する行に表示します。
- **補完**：フロントマターと `<!-- -->` コメント内のディレクティブのキー
  （スポットディレクティブには `_` 付き）、`kind` と `_kind` の値（各 kind が設定する内容付き）、
  `theme`、`size`、`paginate`、`math`、`headingDivider`。
  `{{ }}` 内では、テンプレート関数とデータのフィールド。
- **ホバー**：ディレクティブの働き、kind が設定する内容、テンプレート関数の
  引数を表示します。

## エクスポート {#export}

*go-slides: Export Slide Deck...*（エディターのタイトルバーにもあります）は、
`slides -o deck.pdf deck.md` と同じように `slides` コマンドを使って、HTML、
ベクター PDF、編集可能な PPTX と ODP、PNG、JPEG を書き出します。[コマンドライン]({{< relref "cli" >}}) を参照してください。
信頼されたワークスペースが必要です。`slides` が見つからない場合、拡張機能は上記の `go install` を
ターミナルで実行することを提案します。

## 設定 {#settings}

| 設定 | デフォルト | |
|---|---|---|
| `goSlides.cliPath` | `""` | `slides` コマンド。空の場合は `PATH` 上の `slides`。マシン設定なので、ワークスペースからは設定できません |
| `goSlides.exportType` | `pdf` | エクスポートダイアログで最初に提案される形式 |
| `goSlides.exportAutoOpen` | `true` | エクスポートしたファイルを開きます |
| `goSlides.html` | `false` | Marp の許可リストではなく、生の HTML をすべて許可します。信頼されたワークスペースでのみ有効です |
| `goSlides.preview.marpDocuments` | `auto` | `marp: true` のスライドデッキをプレビューするかどうか：`auto`、`always`、`never` |

## セキュリティ {#security}

スライドデッキとワークスペースは信頼されていないものとして扱います。生の HTML は、
信頼されたワークスペースで `goSlides.html` がオンになっていない限り、Marp の許可リストに従います
（`html: true` を含む `slides.yaml` があっても、プレビューでは有効になりません）。また、VS Code のプレビューは
スライドデッキのスクリプトを実行しません。プレビューが読み込むのはスライドデッキのワークスペースフォルダー内の
ファイルだけで、シンボリックリンクはたどらず、テンプレートから環境変数は見えません。
エクスポートはプログラムを実行するので、ワークスペースが信頼されるまで待機します。また、そのプログラムが
ワークスペースによって指定されることはありません。

## パフォーマンス {#performance}

go-slides の 100 枚のベンチマーク用スライドデッキについて、Apple M4 Max 上の VS Code 1.142 で、
同じセッション内で Marp for VS Code 3.6.1 と比較して計測しました。

| | go-slides | Marp for VS Code |
|---|---:|---:|
| アクティベーション（WebAssembly モジュールの読み込み） | 189 ms | — |
| 最初のプレビュー | 52 ms | 78 ms |
| 編集後のプレビュー（中央値） | **19 ms** | 23–24 ms |
| `.vsix` | 7.5 MB | 36.3 MB |

この数値は、拡張機能 v0.3.0（go-slides v0.8.0）で 100 枚のスライドデッキを使い、
50 枚目のスライドに 1 文字入力したときのもので、1 分間のロードアベレージが 8 未満の状態で
交互に行ったラウンドから得ています。最初のプレビューは Marp より速く、編集後のプレビューも
現在は Marp より速くなっています。

### v0.8.0 と拡張機能 v0.3.0 {#in-v080-and-extension-v030}

go-slides v0.8.0 は、開いている各スライドデッキのレンダリング済みスライドを保持し、編集後は、
ソースまたは継承された状態（ディレクティブ、ページ番号、総数、見出しの id、リンク参照、ソース行）が
変わったスライドだけを再レンダリングします。出力は完全なレンダリングとバイト単位で
同じです。100 枚中 50 枚目のスライドで入力すると、再レンダリングされるのは 1
枚だけで、ネイティブでは 2.5 ms、WebAssembly では 11.5 ms です。その後、VS Code のプレビューは
変更されたスライドの DOM だけを差分更新します。拡張機能のスモークテストでは、両方の拡張機能について、
計測するレンダリングの合間に 50 枚目のスライドへ 1 文字を入力します。
拡張機能の発表者ビューは、親ウィンドウ経由（`postMessage`）で同期します。

### v0.7.1 と拡張機能 v0.2.2 {#in-v071-and-extension-v022}

go-slides v0.7.1 は、解析済みのテーマ、パック済みのスタイルシート、ハイライト済みのコードを
レンダリング間でキャッシュし、スクロール同期用の行をバイト単位で走査せずに記録し、
ライターを再利用します：100 枚のスライドデッキの 1 回のレンダリングで割り当てるメモリは 5.7 MB で、
拡張機能 v0.2.0 の 52.7 MB から減りました。これにより、拡張機能 v0.2.2 では `GOGC=400` の上書きをやめています。
比較のために挙げると、拡張機能 v0.2.1 は、デフォルトの `GOGC` で最初のプレビューに 72 ms、
編集後に 44 ms（`GOGC=400` では 36 ms）かかり、v0.2.0 は
47–51 ms でした。

拡張機能 v0.2.2 の計測値は、最初のプレビューが 52 ms、編集後が 24 ms で、
Marp の 23 ms とほぼ同じでした。
