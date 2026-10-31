---
title: "Galeria"
weight: 12
description: "Doze pequenas apresentações, cada uma mostrando bem uma coisa: temas, kinds, templates, HCL, matemática, imagens, código, emoji, transições e uma palestra completa. Abra qualquer uma delas no playground."
tags: [gallery, examples]
---

Doze pequenas apresentações da pasta
[`examples/`](https://github.com/go-slides/slides/tree/main/examples) do
go-slides, cada uma mostrando bem uma coisa. As miniaturas são o primeiro slide
de cada apresentação, renderizado pelo próprio go-slides quando este site é
construído; a CI constrói cada apresentação em HTML e PDF, então nenhuma delas
quebra.

**Abrir no playground** carrega a apresentação no seu navegador, em um link que
não envia nada a nenhum servidor. Uma apresentação que tem arquivos além de
`deck.md` (um arquivo de kinds, partes incluídas, imagens) não cabe em um link:
baixe o **.zip** e importe-o com o painel de arquivos do playground, ou
descompacte-o e construa-o com `slides build --allow-local-files deck.md`.

As imagens são SVGs originais dos autores do go-slides, liberados sob
[CC0 1.0](https://creativecommons.org/publicdomain/zero/1.0/), e as
apresentações que as usam dizem isso.

{{< gallery >}}
