---
title: "Playground"
weight: 15
description: "Escreva uma apresentação no navegador e veja os slides acompanharem: o go-slides compilado para WebAssembly, com autocompletar, ajuda e exemplos."
tags: [playground, browser]
---

**[Abrir o playground →](https://go-slides.github.io/playground/)**

O playground é o próprio go-slides, v0.14.0, compilado para WebAssembly e
executado no seu navegador. Digite uma apresentação à esquerda; os slides
acompanham à direita. Nada é enviado a um servidor: nem a apresentação, nem um
link de compartilhamento.

## O que ele tem {#what-it-has}

- **Pré-visualização ao vivo** de cada slide, renderizada fora da thread da
  página (um Web Worker), de modo que digitar nunca espera por ela. Clique em um
  slide para levar o editor até a sua linha; mova o cursor e a pré-visualização
  rola até o slide dele.
- **Apresentação de slides**: a apresentação em si, com o teclado, o controle na
  tela e as transições.
- **Visão do apresentador**: **Present** (ou `p` na apresentação) a abre em uma
  segunda janela: o slide atual, o próximo passo, as notas do apresentador (os
  `<!-- comments -->` da apresentação), um relógio e um cronômetro. As setas,
  a barra de espaço, Home e End funcionam em qualquer uma das janelas, e cada uma
  acompanha a outra, inclusive quando você edita a apresentação. Se a janela não
  abrir, permita pop-ups para o site. A visão geral (`o`) existe somente na
  página baixada.
- **Autocompletar** enquanto você digita: chaves de diretiva (as pontuais com
  `_`, em comentários), os valores de `kind`, `theme`, `size`, `paginate`, e as
  funções de template e os campos de dados dentro de `{{ }}`.
- **Ajuda** sobre o que está sob o cursor: uma diretiva, um kind (o tipo de um slide: o que ele
  define), uma função de template. O painel de ajuda também tem links para estas
  páginas.
- **Problemas** na própria linha: os erros de template, YAML e HCL são
  destacados no editor e listados abaixo dele, e a pré-visualização mantém a
  última apresentação que foi renderizada.
- **Exemplos**: uma apresentação de boas-vindas, templates que geram slides,
  kinds de slide, um front matter em HCL, temas e diretivas pontuais, e
  matemática, imagens e fundos (`![bg]`, um `![bg left:40%]` dividido, filtros).
- **Arquivos**: o painel à esquerda do editor é o projeto da apresentação,
  organizado do jeito que o [comando]({{< relref "cli" >}}) e a
  [extensão do VS Code]({{< relref "vscode" >}}) o leem: `deck.md`,
  `slides.yaml`, `themes/`, `kinds/`, `media/`.
  - Clique em um arquivo para editá-lo. O editor o colore pela sua linguagem
    (Markdown, YAML, HCL, JSON, CSS) e marca os seus problemas.
  - Clique com o botão direito (ou Shift+F10) para New File, New Folder, Rename
    (F2), Delete (pergunta antes) e New theme / New kind. Arraste um arquivo
    para uma pasta para movê-lo. O teclado segue o padrão de árvore do
    WAI-ARIA, inclusive a busca por digitação.
  - **Project ▸ New theme** cria `themes/my-gaia.css` (ou a partir de default,
    uncover): um `@import` do tema embutido e as regras que você acrescentar.
    Ele entra na lista de temas. **New kind** e **New data file** também partem
    de modelos. Cada um é listado em `slides.yaml`, de modo que
    `slides build deck.md` o encontra.
  - **Project ▸ Download project (.zip)** salva tudo. Descompactado,
    `slides build deck.md` o constrói. **Open project** recebe de volta um
    `.zip` ou uma pasta. Um arquivo compactado que sai da sua pasta (`..`,
    caminhos absolutos), contém mais de 256 arquivos ou 8 MiB em um arquivo, ou
    se expande mais de 100 vezes é recusado.
  - Tudo é guardado neste navegador junto com o rascunho. Um projeto comporta até
    64 MiB, e a linha abaixo dos arquivos diz quanto está em uso.
- **Mídia**: solte ou cole uma imagem, um vídeo ou um som no editor e ele vai
  para `media/`, e a apresentação o nomeia: `![](media/photo.png)`. Solte
  arquivos em uma pasta do painel para colocá-los ali. A apresentação continua
  pequena e legível, e nada é enviado. A sandbox da pré-visualização os carrega
  sem ser aberta (a página lhe entrega os bytes). **Save as ▸ HTML page** os
  coloca na página. Imagens `https:` também funcionam, mas o navegador as busca
  no servidor delas, que então vê o seu endereço. A pré-visualização não envia
  referrer.
- **Painéis redimensionáveis**: o painel de arquivos, o editor, as miniaturas
  dos slides, a pré-visualização e a ajuda são separados por barras que você
  pode arrastar (o ponteiro vira um cursor de redimensionamento sobre uma
  delas). O painel ao lado da barra fica com o espaço, os outros mantêm o
  tamanho. Dê um duplo clique em uma barra, ou pressione Enter nela, para fechar
  o painel e trazê-lo de volta. Pelo teclado, F6 passa de barra em barra, as
  setas movem uma delas em 16 px, Home e End levam o painel anterior ao menor e
  ao maior tamanho, e Escape volta ao editor. A aba View ainda desliga painéis.
  Os tamanhos que você define são guardados neste navegador, com o rascunho; os
  tamanhos padrão (190, 168 e 280 px, com o editor e a pré-visualização
  dividindo o restante na proporção de 6 para 5) nunca são armazenados.
- **Formatar**: a aba Format edita o Markdown para você, e cada clique é um
  passo de Undo.
  - **Texto**: Bold, Italic, Strikethrough e Inline code funcionam sobre a
    seleção ou a palavra sob o cursor e alternam (Ctrl+B, Ctrl+I, Ctrl+Shift+X e
    Ctrl+E; Cmd no Mac). Um botão aparece pressionado quando o cursor está
    dentro das suas marcas. O Markdown não tem sublinhado: Underline (Ctrl+U)
    envolve o texto em `<u>…</u>`, uma tag que a lista de permissões de HTML
    padrão do Marp mantém, então funciona com o HTML bruto desativado.
  - **Parágrafo**: Heading 1 a 3 (Ctrl+Alt+1..3), marcadores (Ctrl+Shift+8),
    numeração (Ctrl+Shift+7) e uma citação funcionam sobre as linhas
    selecionadas e alternam. A régua horizontal é a tag `<hr>`: `---`, `***` e
    `___` iniciariam, cada um, um novo slide no Marp.
  - **Tabela**: uma grade escolhe as colunas e as linhas. Com o cursor em uma
    tabela, você adiciona ou remove uma linha ou uma coluna e alinha a coluna
    (`:--`, `:-:`, `--:`); a tabela é reescrita alinhada após cada edição.
  - **Slide text** à esquerda, ao centro ou à direita escreve um bloco
    `<style scoped>` no slide atual. O Marp não tem sintaxe de alinhamento, e um
    atributo `style` não passa pela lista de permissões de HTML padrão, mas o
    bloco com escopo passa, então funciona com o HTML bruto desativado e aparece
    na pré-visualização, na página HTML e no PDF.
  - **Link** (Ctrl+K) escreve `[text](url)` com a url selecionada.
- **Realce** das chaves do front matter em qualquer profundidade, em YAML
  (também em mapeamentos `{a: 1}`) e em HCL.
- Seletor de **tema**: default, gaia, uncover, os temas do projeto
  (`themes/*.css`) e os temas que a apresentação define (`/* @theme name */` em
  um `<style>`). Ele escreve `theme:` no front matter, YAML ou HCL, de modo que
  um link compartilhado o mantém, e acompanha o front matter quando você o edita.
- A pré-visualização **mantém a sua posição** enquanto você digita: somente os
  slides que mudaram são substituídos, e uma troca de tema reabre a
  pré-visualização no mesmo slide.
- **Link de compartilhamento**: a apresentação viaja no próprio link (na parte
  `#`, compactada), então abrir o link abre a apresentação. Os links carregam até
  64 KiB compactados, e somente a apresentação: quando o projeto contém mais
  (temas, kinds, mídia), Share avisa e oferece o `.zip`.
- **Save as**:
  - a página HTML, a mesma que `slides deck.md` escreve;
  - um **PDF** (vetorial, com o texto selecionável, com o sumário);
  - uma apresentação **PowerPoint** ou **OpenDocument** editável.

  Eles são feitos no seu navegador, pelo mesmo código de `slides --pdf`,
  `--pptx` e `--odp`: o PDF da apresentação de referência de 100 slides é
  idêntico, byte a byte, ao do comando. A primeira exportação baixa o motor de
  exportação (cerca de 5 MiB, uma única vez; o navegador o mantém até a próxima
  atualização), com uma linha de progresso, e o compila à medida que chega.
  Depois, baixa somente as fontes que os seus slides desenham. Em um link de
  3 MB/s, o primeiro PDF da apresentação de boas-vindas fica pronto em menos de
  3 segundos. A apresentação de 100 slides leva então cerca de 3 segundos. Uma
  imagem `https:` só é incluída se o seu servidor permitir (CORS); imagens
  `data:` sempre são.
- **O seu rascunho é mantido**:
  - A apresentação é salva neste navegador enquanto você digita (IndexedDB,
    somente para este site) e volta quando você recarrega. Nada é enviado a lugar
    algum. Os arquivos de texto do projeto (temas, kinds, dados, `slides.yaml`)
    são mantidos do mesmo modo, mesmo quando você fecha a aba logo depois de
    digitar; a mídia é armazenada à medida que é solta.
  - Um link compartilhado abre a sua apresentação e mantém o seu rascunho:
    **Draft ▸ Restore my draft** o traz de volta.
  - Um exemplo, **Draft ▸ Reset to the example** (que pergunta antes) e outra
    aba também nunca o apagam: **Draft ▸ Restore replaced** o recupera.
  - **Draft ▸ Forget my drafts** apaga tudo o que está armazenado.

Toda exportação renderiza o projeto que a pré-visualização mostra: o seu
`slides.yaml`, temas, kinds, arquivos de dados e mídia, com a mídia dentro do
documento.

Imagens como arquivos (slides `--images`) não são feitas no navegador: use o
[comando]({{< relref "cli" >}}) ou a
[extensão do VS Code]({{< relref "vscode" >}}).

## Segurança {#security}

A apresentação é exibida em um iframe em sandbox, sem a origem da página: ele
não alcança nem o playground nem o armazenamento dele, não envia nenhum
formulário e não navega para lugar algum. A página dele carrega uma Content
Security Policy que executa somente os scripts dela. O HTML bruto da
apresentação segue a lista de permissões do Marp, como o `slides` faz por
padrão. Uma renderização que dura mais de 5 segundos (um template em laço) é
interrompida.

A janela do apresentador mantém essa sandbox. Ela é aberta pelo próprio
playground e mostra a apresentação em um iframe em sandbox do mesmo modo. Os dois
iframes não podem compartilhar armazenamento, então o playground repassa o
número do slide entre eles em mensagens. Nenhum frame que executa a
apresentação recebe a origem do playground.

A mídia do projeto chega aos frames em sandbox como bytes em uma mensagem, e
cada frame faz com eles as suas próprias URLs `blob:`, na sua própria origem. Uma
URL `blob:` feita pelo playground não carrega ali, e a política dos frames ganha
apenas `blob:` para mídia e fontes.

O playground pode ser servido a partir da sua própria máquina, atrás de um
login: veja [Auto-hospedagem]({{< relref "self-hosting" >}}).

## Tamanho e velocidade {#size-and-speed}

O módulo tem 8,8 MiB como o GitHub Pages o serve (gzip; o Pages não serve
brotli). Ele é baixado uma vez com uma barra de progresso e depois fica em cache
no navegador até o próximo deploy. Em um Apple M4 Max com Chrome 155, medido
contra a página no ar, o editor fica pronto 2,1 s depois de a página abrir, dos
quais 1,5 s são o download, e os primeiros slides aparecem aos 2,4 s. Uma
alteração na apresentação de referência de 100 slides chega à pré-visualização
em cerca de 30 ms: o go-slides renderiza somente os slides que mudaram, e
somente esses são substituídos na pré-visualização.
