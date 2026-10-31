---
title: "Linha de comando"
weight: 50
description: "Todos os comandos e flags do slides, e o arquivo de configuração slides.yaml."
tags: [cli]
---

{{< callout type="warning" >}}
Esta página descreve a v0.14.0.
{{< /callout >}}

```text
slides [deck.md] [flags]
slides [command] [flags] deck.md
```

## Comandos {#commands}

| Comando | O que faz |
| --- | --- |
| `build` | escreve a apresentação como HTML, ou como PDF ou imagens (veja abaixo). **O padrão**: `slides deck.md` é `slides build deck.md` |
| `expand` | imprime a apresentação como Markdown Marp simples: templates expandidos, kinds (o tipo de um slide) substituídos pelas diretivas que eles representam |
| `notes` | imprime as notas do apresentador |
| `kinds` | lista os kinds que a apresentação pode usar e o que eles definem; `--json` (na v0.6.0) os escreve como `[{"name", "extends", "directives": [{"key", "value"}]}]`, com `extends` resolvido |
| `themes` | lista os temas |
| `version` | imprime a versão (assim como `slides --version`) |
| `completion` | imprime um script de autocompletar do shell: `bash`, `zsh`, `fish` ou `powershell` |

`slides <command> --help` imprime as flags de um comando. As flags podem vir
antes ou depois da apresentação: `slides deck.md -o out.html` e `slides -o
out.html deck.md` são a mesma coisa. As flags longas levam dois hifens
(`--html`); `-o` é a única curta. O `build` escreve `deck.html` ao lado de
`deck.md`, a menos que `-o` diga outra coisa.

## Flags {#flags}

| Flag | Comandos | O que faz |
| --- | --- | --- |
| `-o`, `--output file` | `build`, `expand`, `notes` | arquivo de saída; `-` para a saída padrão |
| `-I`, `--input-dir dir` | `build` | converte os arquivos Markdown sob `dir`, recursivamente (na v0.6.0); veja abaixo |
| `-P`, `--parallel n` | `build` | apresentações construídas ao mesmo tempo (na v0.6.0): 5 por padrão, no máximo o número de núcleos; 1 as constrói uma após a outra |
| `-w`, `--watch` | `build` | constrói de novo quando uma apresentação, um tema, um arquivo de kinds ou de dados ou a configuração muda, e recarrega as páginas abertas (na v0.6.0); veja abaixo |
| `-s`, `--server dir` | `build` | serve um diretório, convertendo as suas apresentações sob demanda (na v0.6.0); veja abaixo |
| `-p`, `--preview` | `build` | abre o resultado no navegador padrão (na v0.6.0): as páginas HTML com `--watch` (implícito), o servidor com `--server` |
| `--host addr` | `build` | o endereço em que `--server` escuta, `127.0.0.1` por padrão (na v0.6.0) |
| `--allowed-host name` | `build` | o nome de host público de um proxy reverso ao qual `--server` responde, repetível (na v0.14.0); veja [Auto-hospedagem]({{< relref "self-hosting" >}}) |
| `--template bespoke\|bare` | `build` | o template da página HTML (na v0.6.0): `bespoke`, o padrão, com o script de apresentação, o controle na tela e as transições, ou `bare`, os slides empilhados (`--bare`) |
| `--bespoke.osc`, `--bespoke.progress`, `--bespoke.transition` | `build` | as opções do template bespoke, com os padrões do marp-cli (na v0.6.0): controle ligado, **barra de progresso desligada** (antes ela era sempre desenhada), transições ligadas; `--bespoke.osc=false` desliga uma delas |
| `-c`, `--config-file file`, `--no-config-file` | `build` | o arquivo de configuração do marp-cli, JSON ou YAML (na v0.6.0); veja abaixo. `--config` é o arquivo próprio do `slides` |
| `--browser`, `--browser-path`, `--browser-protocol`, `--browser-timeout` | `build` | aceitas e ignoradas com um aviso (na v0.6.0): não há navegador |
| `--pptx-editable`, `--no-parallel` | `build` | as grafias do marp-cli para `--pptx` e `-P 1` (na v0.6.0) |
| `--config file` | todos | arquivo de configuração; por padrão `slides.yaml` ao lado da apresentação, se houver |
| `--no-config` | todos | não lê um arquivo de configuração |
| `--theme-set path` | todos | um arquivo CSS de tema, ou um diretório deles (repetível) |
| `--kinds file` | todos | um arquivo YAML de kinds (repetível) |
| `--data file` | todos | um arquivo YAML ou JSON de dados de template (repetível) |
| `--set key=value` | todos | um valor de template (repetível) |
| `--no-template` | todos | não expande os templates Go |
| `--strict-yaml` | todos | analisa as diretivas como YAML estrito, como o Marpit puro faz (o Marp Core as analisa de forma flexível) |
| `--no-emoji` | todos | deixa `:shortcodes:` e emoji Unicode como texto em vez de desenhá-los como imagens (desde a v0.4.0: [Emoji]({{< relref "emoji" >}})) |
| `--html` | `build` | permite todo o HTML bruto; por padrão, a lista de permissões do Marp Core |
| `--lang code` | `build` | o `lang` da página quando a apresentação não define nenhum |
| `--bare` | `build` | empilha os slides, sem o script de apresentação |
| `--pdf` | `build` | escreve um PDF, uma página por slide, com sumário — texto e formas como vetores (também escolhido por `-o file.pdf`) |
| `--pdf-raster` | `build` | escreve o PDF com cada página sendo uma imagem do slide, como a v0.3.0 fazia |
| `--notes` | `build` | escreve as notas do apresentador como um arquivo de texto, com os slides separados por `---` como o marp-cli faz (na v0.6.0; também escolhido por `-o file.txt`). O comando `notes` as imprime em outro formato |
| `--pdf-notes` | `build` | acrescenta as notas do apresentador de cada slide à sua página do PDF como uma anotação de nota, como o marp-cli faz (na v0.6.0). Pede um PDF quando nada mais diz o que escrever |
| `--pdf-outlines`, `--pdf-outlines.pages`, `--pdf-outlines.headings` | `build` | os marcadores do PDF (na v0.6.0): "Page N" por slide e os títulos aninhados sob ele. Ligado por padrão aqui, diferente do marp-cli; `=false` desliga uma parte. Veja abaixo |
| `--images png\|jpeg` | `build` | escreve uma imagem por slide: `deck.001.png`, … |
| `--image png\|jpeg` | `build` | escreve o primeiro slide como imagem (também escolhido por `-o file.png` ou `-o file.jpg`) |
| `--pptx` | `build` | escreve um arquivo PowerPoint com as notas do apresentador (também escolhido por `-o file.pptx`): editável desde a v0.5.0, uma imagem por slide na v0.4.0 |
| `--pptx-raster` | `build` | escreve um arquivo PowerPoint com uma imagem de slide inteiro por slide, como o marp-cli faz (desde a v0.5.0); `--image-scale` tem padrão 2 |
| `--odp` | `build` | escreve uma apresentação OpenDocument com as notas do apresentador (também escolhido por `-o file.odp`): editável desde a v0.5.0 |
| `--odp-raster` | `build` | escreve uma apresentação OpenDocument com uma imagem por slide (desde a v0.5.0) |
| `--jpeg-quality n` | `build` | qualidade do JPEG, de 1 a 100 (padrão 85) |
| `--image-scale n` | `build` | pixels do dispositivo por pixel CSS para imagens e PDF (padrão 1); 2 dá 2560×1440 para 16:9, com texto e vetores desenhados nessa densidade, não ampliados |
| `--allow-local-files` | `build` | deixa imagens e PDF lerem arquivos locais, somente de dentro do diretório da apresentação; recusado por padrão, como no marp-cli |
| `--theme name\|file.css` | `build` | substitui o tema da apresentação (na v0.6.0): um nome de tema, ou um arquivo CSS, que se torna o tema qualquer que seja o `@theme` que ele declare; um diretório é recusado, use `--theme-set` |
| `--title`, `--description`, `--author`, `--keywords`, `--url`, `--og-image` | `build` | os metadados da página (na v0.6.0), sobre as diretivas `title`, `description`, `author`, `keywords`, `url` e `image` da apresentação; veja abaixo |

## Várias apresentações, e um diretório {#several-decks-and-a-directory}

Na v0.6.0, o `build` aceita várias apresentações, um diretório (as
apresentações sob ele, encontradas recursivamente) ou um padrão, e constrói até
`--parallel` delas ao mesmo tempo. `-` é a entrada padrão; diferente do
marp-cli, um pipe não é lido a menos que `-` o diga.

Um padrão é lido como o glob do marp-cli o lê (`**` e `{a,b}` na v0.7.0; coloque
entre aspas para que o shell o deixe para o `slides`): `*`, `?`, `[a-z]` e
`{md,txt}` dentro de um nome, `**` para qualquer número de diretórios, de modo
que `'docs/**/*.md'` pega as apresentações da árvore. Um curinga não alcança um
nome que começa com ponto, nada é procurado em `node_modules`, maiúsculas e
minúsculas importam (`*.md` não encontra `UPPER.MD`), e tudo o que um padrão
casa é convertido, qualquer que seja a extensão. Um padrão que não casa com nada
é um erro.

- **`-I dir`** encontra os arquivos Markdown (`.md`, `.mdown`, `.markdown`,
  `.markdn`, em minúsculas) sob `dir`, como o marp-cli faz: `node_modules` e
  diretórios com ponto não são vasculhados. Com `-o out`, cada apresentação é
  escrita no seu caminho sob `dir`, espelhado sob `out`: `dir/sub/b.md` dá
  `out/sub/b.html` (`b.pdf`, `b.001.png`, `b.txt`, … o que as flags disserem).
  Sem `-o`, as saídas vão para ao lado das suas apresentações. Nomear
  apresentações além de `-I` é um erro, e `-o file` para várias apresentações
  também (`-o -` é permitido).
- **`-P n`** constrói `n` apresentações ao mesmo tempo. A renderização de uma
  apresentação já usa os núcleos, então `n` é limitado pelo número deles; com
  `-o -`, as apresentações são construídas uma após a outra. As mensagens de uma
  apresentação ficam juntas, e uma apresentação que falha não interrompe as
  outras: o status de saída é 1 e a mensagem lista as falhas.

## Transições e a página de apresentação {#transitions-and-the-presenting-page}

Na v0.6.0, a diretiva `transition` do marp-cli funciona: `<!-- transition:
fade 1s -->`, um nome e uma duração opcional, para o slide e os seguintes, ou
`<!-- _transition: cover -->` para um só. Os nomes embutidos são os do marp-cli:
`clockwise`, `counterclockwise`, `cover`, `coverflow`, `cube`, `cylinder`,
`diamond`, `drop`, `explode`, `fade`, `fade-out`, `fall`, `flip`, `glow`,
`implode`, `in-out`, `iris-in`, `iris-out`, `melt`, `overlap`, `pivot`, `pull`,
`push`, `reveal`, `rotate`, `slide`, `star`, `swap`, `swipe`, `swoosh`, `wipe`,
`wiper` e `zoom`; `none` a desliga. A página a executa com a **View Transition
API** do navegador quando o espectador sai do slide que a tem, avançando, e
quando volta a ele; os navegadores sem a API trocam de slide como antes, e um
espectador que pede movimento reduzido recebe um fade simples. Um tema pode
trazer keyframes próprios, chamados `marp-transition-NAME` (ou `marp-outgoing-` e
`marp-incoming-`, e `-backward`), que são usados antes dos embutidos.

A página tem o **controle na tela** — anterior, página, próximo, tela cheia —
que aparece quando o ponteiro ou uma tecla se move e some depois de um instante
(`--bespoke.osc=false` o remove) e, com `--bespoke.progress`, uma barra de
progresso. **A barra de progresso agora vem desligada por padrão**, como no
marp-cli; antes ela era sempre desenhada.

### Visão do apresentador {#presenter-view}

Na v0.7.0, a página tem a visão do apresentador do marp-cli. Pressione `p`, ou o
último botão do controle, e abre-se uma janela com o slide atual, o passo
seguinte, as suas notas do apresentador, a página, a hora e um cronômetro
(clique no cronômetro para zerá-lo; `+` e `-` redimensionam as notas; a barra
entre os slides e as notas pode ser arrastada). As duas janelas se acompanham nos
dois sentidos: as setas, um clique no próximo passo ou em Previous/Next em
qualquer uma move a outra, um fragmento de cada vez.

Ela funciona como a do marp-cli, lida a partir dos fontes da 4.5.1: a janela é
`?view=presenter&sync=KEY` do mesmo arquivo, o próximo passo é a mesma página que
`?view=next` em um frame, e as janelas compartilham a sua posição por meio do
`localStorage` (`bespoke-marp-sync-KEY`), então não precisa de servidor nem de
rede, e não funciona em um navegador sem `localStorage` (o botão avisa). As notas
são texto: `&`, `<` e `>` são escapados, de modo que uma nota nunca pode ser
marcação. Elas ficam na página, ocultas, para a visão do apresentador pegá-las, o
que as mantém fora do que um leitor de tela lê para o público. As páginas
`--bare` não têm nada disso. A visão acrescenta cerca de 10 KB a uma página.

### Visão geral {#overview-view}

Na v0.8.0, pressione `o` (ou `Escape`: o marp-cli alterna com os dois), o botão
de grade do controle, ou o contador de páginas da visão do apresentador, e cada
slide é exibido inteiro em uma grade sobre a página. As setas escolhem um slide
(para cima e para baixo, por linha), `Home` e `End` vão às extremidades, e as
outras janelas acompanham; `Enter`, `Space` ou um clique escolhe um slide e,
sobre a página, fecha a visão geral. Sobre a visão do apresentador, uma escolha a
deixa aberta (o apresentador escolhe o slide a mostrar e continua olhando). O
botão "Close overview", `o` e `Escape` a fecham. É a do marp-cli: a página de
novo em um frame em `?view=overview&sync=KEY`, compartilhando a posição pelo
`localStorage`, sem rede. Os slides da grade são focalizáveis e rotulados "Slide
N", e a visão geral fechada é inerte, de modo que os usuários de teclado e de
leitor de tela nunca caem nela. As duas visões juntas acrescentam cerca de 15 KB
a uma página.

### Arquivos de configuração {#configuration-files}

O marp-cli lê suas opções de um arquivo no diretório de trabalho, e o `slides`
também, quando o arquivo é de dados: `.marprc` (YAML; JSON também é YAML),
`.marprc.json`, `.marprc.yaml`, `.marprc.yml`, `.config/marprc*`, ou a chave
`marp` do `package.json`; `--config-file` (`-c`) nomeia qualquer arquivo JSON ou
YAML, `marp.config.json` para um deles. Somente o diretório de trabalho é
procurado, não os seus pais, como a busca do marp-cli. Um arquivo de JavaScript
ou TypeScript (`marp.config.js` ...) é um programa, que o `slides` não executa:
ele avisa e segue em frente. As opções são as da documentação do marp-cli que
fazem sentido aqui: `title`, `description`, `author`, `keywords`, `url`,
`ogImage`, `lang`, `theme`, `themeSet`, `html`, `allowLocalFiles`, `template`,
`bespoke`, `inputDir`, `output`, `parallel`, `pdf`, `pptx`, `notes`, `pdfNotes`,
`pdfOutlines`, `images`, `image`, `imageScale`, `jpegQuality`, `watch`, `server`,
`preview`; os caminhos são relativos ao arquivo. A linha de comando prevalece.
`engine`, `options` e `browser*` precisam do motor do Marp e de um navegador, e
são ignoradas com um aviso, assim como as flags `--browser*`: os scripts escritos
para o marp-cli continuam funcionando. (`slides.yaml` e `slides.hcl`, com os
kinds e os dados de template, são separados e lidos como antes.)

## Observar, servir e pré-visualizar {#watch-serve-and-preview}

Na v0.6.0, com as rotas e o protocolo de recarga ao vivo do marp-cli (lidos dos
fontes da 4.5.1 e comparados com um `marp -s` em execução).

**`slides --server dir`** (`-s`) serve o diretório em `http://127.0.0.1:8080/` (a
variável de ambiente `PORT` muda a porta; uma porta já em uso é um erro que diz
isso):

| Requisição | Resposta |
| --- | --- |
| `/deck.md` | a apresentação como HTML, convertida sob demanda |
| `/deck.md?pdf`, `?pptx`, `?png`, `?jpg` (ou `?jpeg`), `?notes` (ou `?txt`) | a apresentação convertida como a query pede; o PPTX é um anexo |
| um diretório | o seu `index.md` ou `PITCHME.md` como página, senão um índice com um link de PDF e um de PPTX para cada apresentação |
| qualquer outro arquivo | o arquivo |
| uma apresentação que não converte | 503 com o erro |

Uma página abre um WebSocket em `/.__marp-cli-watch-notifier__/<id>`, em que o id
é o SHA-256 do caminho da apresentação; o servidor envia `ready` e depois
`reload` quando a apresentação, ou qualquer coisa sob o diretório, muda. Um
diretório requisitado sem a barra final é redirecionado para ela (o marp-cli
serve o índice ali, o que quebra os links relativos da página).

**`--watch`** (`-w`) constrói as apresentações e depois constrói de novo a que
muda e, quando um arquivo de `--theme-set`, um arquivo de kinds ou de dados ou a
configuração muda, todas elas. As páginas que ele escreve abrem um socket no
loopback (porta 37717, ou a próxima livre) e recarregam quando a apresentação
delas é construída de novo. Uma apresentação que falha ao ser construída é
reportada e continua sendo observada. **`--preview`** (`-p`) abre o navegador
padrão no servidor, ou nas páginas HTML recém-construídas; não há janela
própria, já que o `slides` não tem um navegador para controlar.

O servidor é uma porta de entrada para a máquina, então é mais estrito que o do
marp-cli, que escuta em todas as interfaces e não verifica nada:

- ele escuta em **127.0.0.1**; `--host` muda isso, com um aviso, e um endereço
  curinga não consegue mais verificar o cabeçalho `Host`;
- o cabeçalho **`Host`** deve ser um nome de loopback (`localhost`, `127.0.0.1`,
  `[::1]`, `*.localhost`) ou o host em que ele escuta, o que impede o DNS
  rebinding;
- a **`Origin`** de um WebSocket deve ser a do próprio servidor; os sockets do
  modo watch também aceitam a origem `null` de uma página aberta a partir de um
  arquivo; uma requisição de outro site que não seja uma navegação (um script,
  uma imagem, um `fetch`) é recusada;
- um caminho é limpo, seus links simbólicos são resolvidos, e o resultado deve
  estar dentro do diretório: `..` e um link que sai dele são recusados, e o
  `?pdf` de uma apresentação de fora também; nomes que começam com ponto
  (`.git`, `.env`) não são servidos;
- somente `GET` e `HEAD`; poucas conversões ao mesmo tempo; no máximo 4096
  diretórios observados, e nenhum em `node_modules` ou sob um ponto.

Atrás de um proxy reverso, `--allowed-host name` (na v0.14.0, repetível)
acrescenta o nome público do proxy aos `Host` e `Origin` aceitos; as verificações
continuam ativas e o servidor continua no loopback.
[Auto-hospedagem]({{< relref "self-hosting" >}}) o coloca atrás do Caddy, com
HTTPS e um login.

## Metadados da página {#page-metadata}

A linha de comando prevalece sobre a apresentação, como no marp-cli, onde as
flags são anexadas à apresentação como um último comentário de diretiva. Regras
lidas do marp-cli 4.5.1 e fixadas por testes contra a saída dele (na v0.6.0):

- Uma apresentação que não define `title` usa o texto do seu primeiro título do
  nível mais baixo, como escrito (`# **Bold** title` dá `**Bold** title`). Sem
  título, não há `<title>`.
- `keywords` são divididas nas vírgulas, aparadas e sem duplicatas; uma lista do
  front matter é tomada como está. Como o marp-cli, a página traz a tag
  `keywords` a partir da segunda palavra-chave, não para uma só.
- Uma `url` inválida (não absoluta) é ignorada com um aviso, e o valor anterior
  é mantido; uma válida também escreve `<link rel="canonical">` e `og:url`.
- `author` e `description` vêm com seus gêmeos do Open Graph; `og:image:alt` e o
  cartão grande do Twitter precisam de um título e de uma imagem.

## Imagens, PDF, PPTX e ODP {#images-pdf-pptx-and-odp}

`-o` com a extensão `.pdf`, `.pptx`, `.odp`, `.png` ou `.jpg` escolhe o formato,
como o marp-cli faz. Os diretórios de saída são criados. Um PDF tem uma página
por slide, no tamanho do slide, e um sumário dos slides e dos seus títulos.

### Notas e sumário do PDF {#pdf-notes-and-outline}

Na v0.6.0, o PDF aceita as duas opções do marp-cli, e elas são comparadas com os
PDFs que o marp-cli 4.5.1 escreve para a mesma apresentação:

- `--pdf-notes` coloca as notas do apresentador de cada slide na sua página como
  uma anotação de nota: um ícone Note amarelo no canto inferior esquerdo, cujo
  conteúdo são as notas (vários comentários unidos por uma linha em branco) e
  cujo título é o `author` da apresentação. `--pdf-notes` sozinho pede um PDF,
  como no marp-cli.
- O sumário contém um marcador em negrito "Page N" por slide e, sob ele, os
  títulos do slide aninhados por nível (`--pdf-outlines.pages` e
  `--pdf-outlines.headings` escolhem as duas partes). **O sumário vem ligado por
  padrão aqui e desligado no marp-cli.** O `slides` sempre escreveu um; ele custa
  algumas centenas de bytes e torna navegável uma apresentação longa, então o
  padrão permanece. `--pdf-outlines=false`, ou o `--no-pdf-outlines` do
  marp-cli, o desliga, e `=false` em `.pages` ou `.headings` descarta essa
  parte. Um marcador de título salta para o seu slide; o marp-cli também rola até
  o título.

A renderização é feita pelo [go-webengine](https://github.com/go-webengine/engine)
e o PDF é escrito pelo [go-pdfkit](https://github.com/go-pdfkit/pdfkit); as
fontes que os temas embutidos pedem (Lato, Roboto Mono) vêm do
[go-opentype/fonts](https://github.com/go-opentype/fonts) e são incorporadas,
então nenhuma renderização toca a rede. Não há Chrome.

`--image-scale 2` renderiza a dois pixels do dispositivo por pixel CSS. Imagens
locais são recusadas a menos que `--allow-local-files` seja informado, e mesmo
assim só os arquivos sob o diretório da apresentação são lidos: uma apresentação
de outro lugar não consegue alcançar `~/.ssh`. Fundos múltiplos e divididos
(`![bg left:40%]`) são renderizados como as camadas do Marp.

`--pptx` escreve a apresentação com o `WritePresentation` do
[go-ooxml](https://github.com/go-ooxml), com as notas do apresentador e os
metadados da apresentação. Desde a v0.5.0, os slides são editáveis (abaixo);
`--pptx-raster` escreve uma imagem de slide inteiro por slide, como o marp-cli
faz, em uma escala com padrão 2 que `--image-scale` substitui. O LibreOffice abre
o resultado. Somente o LibreOffice foi testado; o PowerPoint e o Keynote não.

`--odp` escreve uma apresentação OpenDocument, por meio do `WritePresentation` do
go-odf, com as suas notas do apresentador; `--odp-raster` escreve uma imagem por
slide.

### PPTX e ODP editáveis {#editable-pptx-and-odp}

Desde a v0.5.0, os slides de um PPTX ou de um ODP são **editáveis** — o PPTX do
marp-cli é uma imagem por slide, e o seu `--pptx-editable` experimental converte
um PDF por meio do LibreOffice. Aqui, cada slide é desenhado a partir do seu
layout, o mesmo que as exportações para PNG e PDF desenham:

- cada linha de texto é uma caixa de texto na sua posição calculada, nas fontes
  com que o slide foi diagramado — **incorporadas ao arquivo**, de modo que ele
  fica igual em uma máquina sem Lato ou Inter (verificado com o LibreOffice
  25.8) — com negrito, itálico, cores, espaçamento entre letras e links
  clicáveis;
- fundos, bordas e caixas arredondadas são formas; os gradientes são
  preenchimentos em gradiente no PPTX e imagens no ODP (um gradiente ODF tem duas
  cores);
- as imagens são armazenadas uma única vez, por mais slides que as mostrem;
- o que um slide não consegue expressar — um elemento sob `opacity`, um
  `filter`, uma sombra desfocada — é uma imagem desse elemento, com
  o seu texto mantido por cima;
- uma fórmula, um emoji ou uma imagem SVG é um grupo de formas vetoriais (`p:grpSp` no PPTX,
  `draw:g` no ODP) que se move como um só. Um SVG que precisa de um recorte ou de um gradiente radial continua a ser uma imagem.

Renderizado de volta pelo LibreOffice, um slide fica a 0,04–2 % dos pixels da
exportação para PNG. `--pptx-raster` e `--odp-raster` mantêm uma imagem por
slide.

### PDF vetorial {#vector-pdf}

Desde a v0.4.0, o PDF é desenhado como vetores: o go-webengine pinta o slide
diagramado no PDF em vez de em pixels. O texto é texto, nas fontes com que o
slide foi diagramado, incorporadas em subconjunto com um mapa `/ToUnicode` — ele
pode ser selecionado, pesquisado e copiado. Fundos, bordas e gradientes são
caminhos e sombreamentos, uma imagem mostrada em muitos slides é incorporada uma
única vez, e os links são clicáveis. A apresentação de referência de 100 slides
tem 758 KiB, contra 837 KiB do marp-cli e 10 MiB do PDF de imagens da v0.3.0, e
é escrita em 0,3 s.

Um elemento sem forma vetorial aqui — sob `opacity` < 1, com um `filter` ou uma
rotação, um `box-shadow` desfocado, um gradiente cuja opacidade se desvanece — é
desenhado como um bitmap apenas desse elemento, com o seu texto colocado por
cima de forma invisível para que continue selecionável. `--pdf-raster` escreve uma imagem por página, como a v0.3.0 fazia.

As fórmulas, os emoji e as imagens SVG também são vetores (go-webengine v0.22.0): cada desenho
distinto é um XObject de formulário PDF, pintado onde é usado, nítido em qualquer zoom.

### PDF acessível e etiquetado {#accessible-tagged-pdf}

Desde a v0.12.0, o PDF vetorial é etiquetado (tagged). Ele carrega a sua
estrutura lógica, que um leitor de tela lê em ordem, como lê uma página web:

| Na apresentação | No PDF |
| --- | --- |
| um slide | uma seção (`Sect`) |
| `#` … `######` | `H1` … `H6` |
| um parágrafo, uma lista, uma tabela | `P`; `L` com o rótulo e o corpo de cada item; `Table`, células de cabeçalho com o seu escopo |
| um link | um `Link`, com a sua área clicável incluída e o seu texto como descrição |
| uma imagem | uma `Figure` com o seu texto alternativo; `![](…)` sem texto alternativo é decoração |
| `$…$`, `$$…$$` | uma `Formula`, com o seu código TeX como `/Alt` e `/ActualText` (os seus glifos são contornos: pdftotext e mupdf não os extraem, os leitores que seguem a estrutura sim) |
| código | `Code`, sem os seus números de linha nem discos de chamada |
| `<span lang="fr">` | um span nesse idioma |
| fundos, bordas, números de página | artefatos, que um leitor ignora |
| notas do apresentador (`--pdf-notes`) | anotações na estrutura |

O arquivo declara PDF/UA-1, o padrão para PDF acessível. Ele carrega o `title` da
apresentação, que o visualizador mostra, e o seu `lang` (`--lang` quando a
apresentação não define nenhum, `en` caso contrário).

Uma apresentação verificada com o veraPDF passa no PDF/UA-1, a menos que os seus
próprios títulos quebrem a ordem do padrão. O PDF/UA exige que o primeiro título
seja um `H1` e que nenhum nível seja pulado (`#` e depois `###`). Somente o autor
da apresentação pode corrigir isso. O PDF do Chrome da mesma apresentação, que o
marp-cli escreve, também é etiquetado, mas falha em sete regras do PDF/UA, não
tem uma seção por slide e não tem fórmulas.

Etiquetar a apresentação de referência de 100 slides custa 4% em tamanho (788
contra 757 KiB) e nenhum tempo mensurável. `--pdf-tags=false` escreve o PDF sem
etiquetas.

## `slides.yaml` {#slidesyaml}

Um arquivo de configuração ao lado da apresentação se aplica a ela, de modo que
uma pasta de apresentações compartilha um só:

```yaml
kinds: kinds.yaml        # a mapping of kinds, or file(s) holding one
themes:                  # theme CSS files or directories
  - themes/
data:                    # template data
  company: Example Inc.
html: false              # allow all raw HTML
lang: en
```

Kinds: o `kinds:` da própria apresentação sobrescreve o do arquivo e o da linha
de comando. Dados: o `data` do arquivo e o da linha de comando sobrescrevem o
front matter da apresentação, de modo que uma apresentação pode ser construída
com dados diferentes.
