---
title: "Vídeo"
weight: 39
description: "Um vídeo em uma apresentação: reproduzido pelo navegador na página HTML, um quadro de pôster com link no PDF, incorporado no PowerPoint e no OpenDocument."
tags: [video]
---

Um vídeo é uma imagem cujo arquivo termina em `.mp4`, `.m4v` ou `.webm`:

```markdown
![video](clip.mp4)
![The demo](clip.mp4#t=12.5)
![demo loop muted](clip.mp4)
```

## Sintaxe {#syntax}

O arquivo é lido da pasta da apresentação, ou de uma URL `http` ou `https`. Um caminho absoluto, um caminho que sai da pasta ou qualquer outro esquema é recusado. `#t=12.5`, ou `#t=1:23` em minutos e segundos, escolhe o quadro de pôster: a imagem mostrada antes da reprodução e pelas saídas que não podem reproduzir o vídeo. A forma de fragmento de mídia `#t=10,20` também é aceita, e seu início é usado. Sem ela, o pôster é o primeiro quadro.

Palavras-chave no texto alternativo, escritas como as palavras-chave de imagem do Marp, mas próprias do go-slides, definem como o vídeo é reproduzido: `autoplay`, `loop` e `muted`. `autoplay` só vale junto com `muted`; caso contrário é ignorado, com um aviso. As demais palavras do texto alternativo são o rótulo acessível.

## O que cada saída mostra {#outputs}

- **HTML** (a página e a prévia do playground): um elemento `<video controls playsinline preload="metadata" poster=...>`. O pôster é o quadro decodificado, incorporado como URI de dados JPEG. Um vídeo nunca toca com som sozinho. Quando o navegador pede menos movimento (`prefers-reduced-motion`), `autoplay` e `loop` são removidos e o vídeo é pausado. Um vídeo que começa sozinho (sem som) só toca enquanto pelo menos metade dele está visível, e pausa quando seu slide não está sendo mostrado.
- **PDF** (`--pdf`): o quadro de pôster como imagem, com um link para o arquivo de vídeo: o caminho relativo como escrito na apresentação, ou a URL para uma fonte `http` ou `https`. A reprodução só acontece ao abrir o link. O PDF não leva anotação `RichMedia`: ela só tocaria no Acrobat, e deliberadamente não é usada.
- **PowerPoint** (`--pptx`): o vídeo é incorporado como uma parte de mídia (`ppt/media/slidesvideoN.mp4|webm`) com seu pôster, como `a:videoFile` mais a extensão `p14:media`. Keynote: importa como filme e reproduz (verificado). LibreOffice Impress: importa arquivos PowerPoint e OpenDocument como vídeo e o decodifica (verificado). PowerPoint: ainda não verificado. As palavras-chave `autoplay`, `loop` e `muted` não passam para o PowerPoint nem para o OpenDocument.
- **OpenDocument** (`--odp`): o vídeo é incorporado em `Video/` como um `draw:plugin` (ODF 1.3 parte 3, 10.4.8), antes do `draw:image` do pôster, e listado no manifesto. O LibreOffice Impress o importa como vídeo e o decodifica (verificado); a reprodução na apresentação depende do seu backend de mídia e do codec.
- **Fontes URL**: a página HTML transmite o vídeo a partir da URL. As exportações mostram um cartão neutro com link para a URL. Nada é baixado na construção, e nada é incorporado no PowerPoint nem no OpenDocument.

## Contêineres e pôsteres {#codecs}

Arquivos MP4 (ISO BMFF) e WebM (Matroska) são lidos, em Go; o ffmpeg nunca é executado. O pôster só é decodificado para H.264 (`avc1` ou `avc3`), o único decodificador de vídeo do go-avkit. Os outros codecs (H.265, VP8, VP9, AV1) ainda tocam em HTML e são incorporados no PowerPoint e no OpenDocument, mas o pôster deles é um cartão escuro neutro com um triângulo de reprodução, e o aviso `video.codec` é impresso.

## Limites {#limits}

Um vídeo é uma entrada não confiável, por isso é limitado: o arquivo (64 MiB), as caixas ou elementos (262.144) e seu aninhamento (16 níveis), o quadro (4096 por 4096) e a decodificação do pôster (no máximo 600 amostras a partir do quadro-chave anterior, 8 MiB por amostra). Um pôster longe demais de um quadro-chave recebe o cartão e o aviso `video.limit`.

Um vídeo recusado é substituído no slide por uma nota, `[video not shown: ...]`; a construção continua, e o problema é um aviso com um código:

- `video.syntax`: não é um arquivo MP4 ou WebM, não há faixa de vídeo, ou um tempo inválido;
- `video.limit`: um dos limites acima foi excedido;
- `video.codec`: o pôster não pode ser decodificado para este codec;
- `video.path`: um caminho absoluto, um caminho que sai da pasta da apresentação, ou um esquema diferente de `http` e `https`;
- `file.missing`: o arquivo não existe.

## No playground {#playground}

Depois de publicado, um vídeo poderá ser adicionado ao projeto como os outros arquivos. Seus bytes ficam no navegador, como URLs blob, e nunca são enviados. Experimente no [playground](https://go-slides.github.io/playground/).

O suporte a vídeo é carregado sob demanda. O leitor que decodifica o quadro do pôster (cerca de 3 MiB compactados) é um módulo à parte, baixado na primeira vez que uma apresentação precisa de um pôster. Uma apresentação sem vídeo nunca o baixa, e o primeiro carregamento do playground fica mais leve na mesma medida. Ele roda em um Web Worker, então a digitação nunca é bloqueada, e lê um arquivo com os mesmos limites do `slides` (64 MiB, caixas e tamanho do quadro) e os mesmos erros `video.limit` e `video.syntax`.
