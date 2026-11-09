---
title: "主题包"
weight: 42
description: "把主题装进一个文件：包含 theme.hcl、CSS、字体、资源和 kind 的 .gsta.xz 归档，按不可信输入进行检查。"
tags: [themes, hcl]
---

**主题包**是一个文件 `<name>.gsta.xz`（Go-Slides Theme Archive）：一个用 xz 压缩的
tar 归档（ustar 或 pax），根目录下有 `theme.hcl`，以及它声明的各个文件。它携带主题的 CSS、字体、徽标和背景，
还可以按需带上幻灯片 kind，并附有名称、版本和许可证，因此主题可以被传递和安装。
主题包自 v0.18.0 起提供。

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
slides theme pack dir       write dir/theme.hcl and its files as <name>.gsta.xz
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
`--theme acme.gsta.xz` 和 `--theme-set acme.gsta.xz` 直接使用文件，无需安装；
见[命令行]({{< relref "cli" >}})。

主题包中的 kind 通过 `--theme file.gsta.xz` 生效，对已安装的主题包则用
`--theme name`，也可以在演示文稿的 front matter 中写 `theme: name`；演示文稿自己的 kind 优先于它们。

## 主题包是不可信输入 {#a-package-is-untrusted-input}

主题包可能来自任何地方，因此每一个读取者（`pack`、`verify`、`install`、`--theme`、
构建）都以同样的方式检查它，并以带有错误码的错误拒绝
（`theme.archive`、`theme.limit`、`theme.meta`、`theme.css`、`theme.svg`、`theme.font`、
`theme.checksum`、`theme.version`）：

- **归档。** 只有一个 xz 流，其字典有上限，窗口随数据增长。tar 中：名称里不得有 `..`、
  绝对路径、盘符或反斜杠；只允许普通文件和目录（不得有符号链接、硬链接、设备或 FIFO）；
  不得有重复名称（比较时不区分大小写）；只允许 ustar 和 pax 头（不得有 GNU 扩展，不得有稀疏文件），
  pax 头最大 8 KiB。大小以实际读取的为准，而不是头部所声明的。
- **上限。** 归档最大 16 MiB，最多 256 个条目，每个条目最大 8 MiB，解压后共 24 MiB，
  xz 字典 16 MiB，解压后的数据流最大为归档大小的 100 倍（64 KiB 及以上），
  `theme.hcl` 最大 64 KiB、嵌套括号最深 32 层，
  每个 CSS 文件最大 1 MiB，每个字体或图像最大 4 MiB。
- **CSS。** 不得有 `@import`，不得有 `@import-theme`（请用 `extends`），不得有 `image-set()`、
  `src()` 或其他加载资源的函数，字符串之外不得有反斜杠转义，不得有 `</style`。
  `url()` 只能是 `#fragment`、`data:` 图像或字体，或包中已声明的文件，绝不能指向包外。
- **SVG。** 不得有 `script`、`foreignObject`、frame、embed、`set`、DOCTYPE 或事件处理器；
  `href` 只能是 `#fragment`；不得有指向文件之外的 `url()`。
- **字体和图像**会被解析：字体必须是字体，图像必须与其扩展名相符，且各有大小上限。

仓库中的 `testdata/themes/acme` 是一个小示例。
