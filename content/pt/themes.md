---
title: "Pacotes de tema"
weight: 42
description: "Um tema em um só arquivo: um arquivo .gsta.xz com um theme.hcl, seu CSS, fontes, recursos e kinds, verificado como entrada não confiável."
tags: [themes, hcl]
---

Um **pacote de tema** é um único arquivo, `<name>.gsta.xz` (Go-Slides Theme
Archive): um arquivo tar (ustar ou pax) compactado com xz, com `theme.hcl` na raiz e os arquivos que ele declara. Ele leva o CSS
de um tema, fontes, logotipos e fundos e, se quiser, kinds de slide, com um
nome, uma versão e uma licença, para que um tema possa ser repassado e
instalado. Os pacotes de tema estão na v0.18.0 e posteriores.

## The theme.hcl file {#the-themehcl-file}

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

`theme.hcl` usa a sintaxe nativa do HCL, apenas com valores literais: sem
função, sem variável, sem `file()` nem `env()`. Um nome é feito de letras
minúsculas, dígitos e hifens, e não é `default`, `gaia` nem `uncover`.
`extends` nomeia o tema base; cada bloco `css` acrescenta uma folha, em ordem;
cada bloco `font` e `asset` nomeia um arquivo do pacote; cada bloco `kind` é um
[kind]({{< relref "kinds" >}}). Um `sha256` em um bloco `css`, `font` ou `asset` é
verificado quando o pacote é lido.

O CSS se refere a fontes e recursos pelo caminho deles no pacote, relativo ao
arquivo CSS: `background: url(assets/logo.svg)`. Quando o pacote é carregado,
cada referência desse tipo vira uma URL `data:`, para que a página seja
autossuficiente, e um bloco `font` vira um `@font-face`. Um pacote também pode
levar arquivos `LICENSE`, `NOTICE` e `README`; qualquer outro arquivo que o
`theme.hcl` não declare é recusado.

## Packing, installing and using {#packing-installing-and-using}

```
slides theme pack dir       write dir/theme.hcl and its files as <name>.gsta.xz
slides theme verify file    run every check on a package
slides theme install file   check it, then install it for the current user
slides theme list           the installed packages
slides theme remove name
```

Os pacotes instalados ficam em `go-slides/themes`, no diretório de configuração
do usuário (`os.UserConfigDir()`: `~/Library/Application Support`,
`$XDG_CONFIG_HOME` ou `~/.config`, `%AppData%`). Em uma apresentação:

```markdown
---
theme: acme
---
```

encontra um pacote instalado, depois dos temas embutidos e do `--theme-set`.
`--theme acme.gsta.xz` e `--theme-set acme.gsta.xz` usam um arquivo
diretamente, sem instalá-lo; veja [Linha de comando]({{< relref "cli" >}}).

Os kinds de um pacote se aplicam com `--theme file.gsta.xz`, com `--theme name`
para um pacote instalado, ou com `theme: name` no front matter de uma
apresentação; os kinds da própria apresentação prevalecem sobre eles.

## A package is untrusted input {#a-package-is-untrusted-input}

Um pacote de tema pode vir de qualquer lugar, então todo leitor (`pack`,
`verify`, `install`, `--theme`, uma construção) o verifica da mesma forma e o
recusa com um erro que tem um código (`theme.archive`, `theme.limit`,
`theme.meta`, `theme.css`, `theme.svg`, `theme.font`, `theme.checksum`,
`theme.version`):

- **Arquivo.** Um único fluxo xz, cujo dicionário é limitado e cuja janela
  cresce com os dados. No tar: sem `..`, caminho absoluto, letra de unidade nem
  barra invertida em um nome; apenas arquivos regulares e diretórios (sem link
  simbólico nem físico, dispositivo ou FIFO); sem nome duplicado (comparado sem
  diferenciar maiúsculas de minúsculas); apenas cabeçalhos ustar e pax (sem
  extensão GNU, sem arquivo esparso), com um cabeçalho pax de no máximo 8 KiB.
  Os tamanhos são os que foram lidos, não os que um cabeçalho declara.
- **Limites.** No máximo 16 MiB de arquivo, 256 entradas, 8 MiB por entrada,
  24 MiB descompactados, um dicionário xz de 16 MiB, um fluxo descompactado de
  no máximo 100 vezes o tamanho do arquivo (a partir de 64 KiB), `theme.hcl` com
  no máximo 64 KiB e 32 níveis de colchetes, 1 MiB de CSS por arquivo, 4 MiB por
  fonte ou imagem.
- **CSS.** Sem `@import`, sem `@import-theme` (use `extends`), sem
  `image-set()`, `src()` nem outra função que carregue um recurso, sem escape
  com barra invertida fora de uma string, sem `</style`. Um `url()` é um
  `#fragment`, uma imagem ou fonte `data:`, ou um arquivo declarado do pacote,
  nunca um que saia dele.
- **SVG.** Sem `script`, `foreignObject`, frame, embed, `set`, DOCTYPE nem
  tratador de evento; um `href` é um `#fragment`; sem `url()` que saia do
  arquivo.
- **Fontes e imagens** são analisadas: uma fonte precisa ser uma fonte, uma
  imagem precisa corresponder à sua extensão, e cada uma tem um limite de
  tamanho.

O diretório `testdata/themes/acme` do repositório é um pequeno exemplo.
