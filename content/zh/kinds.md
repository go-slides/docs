---
title: "Kinds"
weight: 40
description: "幻灯片声明自己是什么：kind 是一组命名的局部指令预设。"
tags: [kinds]
---

**kind**（幻灯片类型）是一组命名的局部指令预设：页眉、页脚、
页码、class、背景。幻灯片声明自己是什么（`title`、
`section`），而不必反复描述自己的外观；要修改整个演示文稿的外观，
只需改一处。

## 使用 kind {#using-a-kind}

```markdown
<!-- kind: section -->
# Part two

---

<!-- _kind: blank -->
![bg](photo.jpg)
```

- `kind:` 的作用方式与局部指令相同：作用于当前幻灯片及其后的幻灯片。
- `_kind:` 的作用方式与单页指令相同：只作用于当前幻灯片。
- 切换 kind 时，会先**清除**上一个 kind 设置的内容，因此 class 或
  页脚永远不会从一个 kind 泄漏到下一个。
- 直接写在幻灯片上的指令**优先于** kind。

## 内置 kind {#built-in}

| Kind | 设置内容 |
| --- | --- |
| `title` | `class: lead`，无页码、无页眉、无页脚 |
| `section` | `class: lead` |
| `content` | 不设置任何内容：使用演示文稿自身的默认值 |
| `blank` | 无页码、无页眉、无页脚 |

它们都可以重新定义。

## 定义 kind {#defining-kinds}

在 front matter 中：

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

`extends` 让一个 kind 以另一个 kind 的指令为起点。kind 可以包含
局部指令和 `extends`。

## 共享 kind {#sharing-kinds}

团队的 kind 应放在一个文件中，供每个演示文稿引用：

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

也可以在命令行中使用 `--kinds kinds.yaml`，或在 `slides.yaml` 中使用 `kinds:`。
优先顺序为：先是内置 kind，然后是命令行和配置文件，最后是演示文稿自身的 `kinds:`。

`slides kinds talk.md` 会列出演示文稿可用的 kind 以及每个 kind 设置的内容。

## 离开 go-slides {#leaving-go-slides}

`slides expand` 会逐张幻灯片地把每个 kind 写成它所代表的单页指令，
因此结果是任何 Marp 工具都能读取的纯 Marp Markdown。
