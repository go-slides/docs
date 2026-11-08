---
title: "主题包"
weight: 42
description: "把主题装进一个文件：包含 theme.hcl、CSS、字体、资源和 kind 的 .slidestheme 归档，按不可信输入进行检查。"
tags: [themes, hcl]
---

**主题包**是一个文件 `<name>.slidestheme`：一个 ZIP 归档（deflate 压缩），
根目录下有 `theme.hcl`，以及它声明的各个文件。它携带主题的 CSS、字体、徽标和背景，
还可以按需带上幻灯片 kind，并附有名称、版本和许可证，因此主题可以被传递和安装。
主题包在 v0.17.1 之后提供。

## theme.hcl 文件 {#the-themehcl-file}

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
  min_slides  = "0.17.0"                # oldest go-slides it works with
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

`theme.hcl` 使用 HCL 的原生语法，且只允许字面量：没有函数，没有变量，
没有 `file()` 或 `env()`。名称由小写字母、数字和连字符组成，且不能是 `default`、
`gaia` 或 `uncover`。`extends` 指明基础主题；每个 `css` 块按顺序添加一个样式表；
每个 `font` 和 `asset` 块都指向包中的一个文件；每个 `kind` 块就是一个 [kind]({{< relref "kinds" >}})。
`css`、`font` 或 `asset` 块上的 `sha256` 会在读取主题包时校验。

CSS 通过文件在包中的路径引用字体和资源，路径相对于该 CSS 文件：
`background: url(assets/logo.svg)`。加载主题包时，每个这样的引用都会变成 `data:` URL，
使页面自成一体，而 `font` 块会变成 `@font-face`。主题包还可以带有 `LICENSE`、
`NOTICE` 和 `README` 文件；`theme.hcl` 没有声明的其他任何文件都会被拒绝。

## 打包、安装和使用 {#packing-installing-and-using}

```
slides theme pack dir       write dir/theme.hcl and its files as <name>.slidestheme
slides theme verify file    run every check on a package
slides theme install file   check it, then install it for the current user
slides theme list           the installed packages
slides theme remove name
```

已安装的主题包位于用户配置目录下的 `go-slides/themes`
（`os.UserConfigDir()`：`~/Library/Application Support`、`$XDG_CONFIG_HOME` 或
`~/.config`、`%AppData%`）。在演示文稿中：

```markdown
---
theme: acme
---
```

会查找已安装的主题包，顺序在内置主题和 `--theme-set` 之后。
`--theme acme.slidestheme` 和 `--theme-set acme.slidestheme` 直接使用文件，无需安装；
见[命令行]({{< relref "cli" >}})。

主题包中的 kind 通过 `--theme file.slidestheme` 生效，对已安装的主题包则用
`--theme name`。只在 front matter 中指定主题的演示文稿会得到它的外观，而不会得到它的 kind，
可以用 `--kinds` 来指定这些 kind。

## 主题包是不可信输入 {#a-package-is-untrusted-input}

主题包可能来自任何地方，因此每一个读取者（`pack`、`verify`、`install`、`--theme`、
构建）都以同样的方式检查它，并以带有错误码的错误拒绝
（`theme.archive`、`theme.limit`、`theme.meta`、`theme.css`、`theme.svg`、`theme.font`、
`theme.checksum`、`theme.version`）：

- **归档。** 名称中不得有 `..`、绝对路径、盘符或反斜杠；不得有符号链接或特殊文件；
  不得有重复名称（比较时不区分大小写）；不得加密，除存储和 deflate 之外不得使用其他压缩方法。
  大小以实际读取的为准，而不是归档所声明的。
- **上限。** 归档最大 16 MiB，最多 256 个条目，每个条目最大 8 MiB，解压后共 24 MiB，
  64 KiB 及以上的条目压缩比不得超过 100 比 1，`theme.hcl` 最大 64 KiB、嵌套括号最深 32 层，
  每个 CSS 文件最大 1 MiB，每个字体或图像最大 4 MiB。
- **CSS。** 不得有 `@import`，不得有 `@import-theme`（请用 `extends`），不得有 `image-set()`、
  `src()` 或其他加载资源的函数，字符串之外不得有反斜杠转义，不得有 `</style`。
  `url()` 只能是 `#fragment`、`data:` 图像或字体，或包中已声明的文件，绝不能指向包外。
- **SVG。** 不得有 `script`、`foreignObject`、frame、embed、`set`、DOCTYPE 或事件处理器；
  `href` 只能是 `#fragment`；不得有指向文件之外的 `url()`。
- **字体和图像**会被解析：字体必须是字体，图像必须与其扩展名相符，且各有大小上限。

仓库中的 `testdata/themes/acme` 是一个小示例。
