---
title: "Playground"
weight: 15
description: "Escreva uma apresentação no navegador e veja os slides acompanharem: o go-slides compilado para WebAssembly, com autocompletar, ajuda e exemplos."
tags: [playground, browser]
---

**[Abrir o Playground →](https://go-slides.github.io/pt/playground/)**

O Playground é o próprio go-slides, v0.20.1, compilado para WebAssembly e
executado no seu navegador. Digite uma apresentação à esquerda; os slides
acompanham à direita. Nada é enviado a um servidor: nem a apresentação, nem um
link de compartilhamento.

A interface do Playground (faixa de opções, ajuda, mensagens) fala o idioma da
página: abra `/pt/playground/` para o português, e assim para os doze idiomas
deste site. O seletor de idioma da página o altera e a sua escolha fica
guardada neste navegador; um link que nomeia um idioma sempre prevalece. As
apresentações de exemplo continuam em inglês. Ucraniano, russo, chinês,
japonês e coreano são desenhados com um subconjunto do Noto Sans (de 0,06 a
2,3 MiB) que a página baixa só para esses idiomas.

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
  páginas. Desde a v0.16.0, a ajuda das diretivas segue o idioma da interface,
  em 11 idiomas.
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
- **Fechar um painel**: cada painel tem um pequeno botão × no cabeçalho, com uma
  dica. Os outros dividem o seu espaço e o painel volta com a largura que tinha.
  Para reabri-lo, use o seu interruptor na aba View (cada painel tem um).
  O editor e a pré-visualização podem ser fechados cada um, nunca os dois: o
  último se recusa, o botão fica esmaecido e a dica diz por quê. Pelo teclado,
  F6 e depois Tab ou Shift+Tab levam a um botão ×, e Enter ou Espaço fecha o
  painel; o teclado volta ao editor. Os painéis abertos são guardados neste
  navegador com os tamanhos (Versions, GitHub e o editor de temas, abertos para
  uma tarefa, começam fechados).
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

## Editor de temas {#theme-editor}

**Página Inicial ▸ Aparência ▸ Editor de temas** abre um painel que edita um
dos temas do projeto com controles, e a pré-visualização acompanha. O tema é
guardado como uma pasta de pacote, `theme-src/<name>/`; a folha que a
pré-visualização e as exportações leem, `themes/<name>.css`, é compilada a
partir dela: o que você vê é o que exporta. Um pacote de tema que você importa
(**Importar…** no painel, **Adicionar mídia…** ou um arrastar e soltar) mantém
ali os seus arquivos, então você pode editá-lo.

- **Paleta**: as cores de fundo, texto, títulos, destaque e links. Escolha uma
  cor ou digite-a (`#1a2b3c`, `rgb()`, `hsl()` ou um nome).
- **Tipografia**: as fontes do texto, dos títulos e do código, com os seus
  tamanhos e pesos. As fontes são as genéricas e as comuns, e as que você
  adicionar (**Adicionar fonte…**: `.woff2`, `.woff`, `.ttf` ou `.otf`).
- **Logotipo**: uma imagem (`.svg`, `.png`), onde ela vai (um canto ou todo o
  fundo) e o seu tamanho. **Número de página e rodapé**: onde vai o número de
  página, e os tamanhos e o alinhamento.
- **Detalhes**: o que diz o `theme.hcl`: versão, descrição, autores, licença,
  página inicial, o tema base (`extends`) e a versão do go-slides de que ele
  precisa.
- **Contraste (WCAG)**: para cada texto sobre o seu fundo, a razão e o seu
  nível, `AA` ou `AAA`, ou “Reprovado”, ao vivo. Os títulos contam como texto
  grande (3:1 basta para AA; o texto normal exige 4,5:1).

**Editar CSS** mostra a folha do tema no editor. Os controles escrevem só
dentro de um bloco delimitado por dois comentários `gsta:visual`, e o releem
quando você edita os seus valores à mão. O que você escreve fora do bloco nunca
é tocado. Se você acrescentar no bloco algo que os controles não conseguem
mostrar, eles param de escrever e o painel avisa; **Restaurar o bloco** o
reescreve a partir dos valores que ele contém e nada mais.

**Importar…** abre um `.gsta.xz` ou um `.css` avulso. Um pacote é verificado
como o `slides theme verify` o verifica, com os mesmos limites: um pacote que
não é seguro (um caminho que sai dele, um link, um script num SVG, uma `url()`
remota, uma soma de verificação que não confere) é recusado com o seu código
`theme.*`, em Problemas, e nada muda. Numa folha avulsa, um `@import "gaia";`
passa a ser o tema base (`extends`), porque um pacote nomeia ali a sua base.
Uma fonte que você adiciona tem a licença `LicenseRef-unspecified` até você
indicá-la em Detalhes.

**Exportar .gsta.xz** monta o pacote no seu navegador e o relê como o faz o
`slides theme verify`, de modo que o arquivo passa nele; **Exportar CSS**
baixa a folha compilada. Nada é enviado. Cada edição é um passo de
**Desfazer** (as edições de um mesmo controle em sequência são um passo) e
guarda o que substitui como versão; **Salvar versão** guarda o tema e a sua
versão, pelo nome, no painel Versões.

## Desfazer, refazer e versões {#undo-redo-and-versions}

**Undo** e **Redo** (desfazer e refazer) ficam na aba Home e em Ctrl+Z,
Ctrl+Shift+Z ou Ctrl+Y (Cmd no Mac). Eles desfazem a última edição do arquivo
exibido (cada arquivo tem seu próprio histórico) e depois a última operação
sobre o projeto inteiro: um exemplo escolhido, Reset, um .zip ou uma pasta
aberta, um rascunho restaurado, um arquivo ou uma pasta excluído ou renomeado,
uma versão restaurada. Redo traz tudo de volta como estava quando Undo foi
pressionado.

As **versões** guardam o projeto inteiro neste navegador, sem enviar nada para
lugar nenhum. Uma é guardada pouco depois de você parar de digitar (no máximo a
cada 2 minutos), antes de cada operação sobre o projeto inteiro e a cada
exportação e link de compartilhamento. **Save version…** guarda uma com um nome.
**Home ▸ Versions** abre o painel Versions:

- **Preview** mostra uma versão na visualização.
- **Compare with now** mostra as mudanças de cada arquivo linha a linha, em
  linha ou lado a lado; F7 e Shift+F7 passam de uma mudança à outra.
- **Restore** coloca uma versão de volta; Undo desfaz a restauração.
- **Delete** remove uma versão.

O que é guardado: as 20 versões automáticas mais recentes, depois uma por hora
durante um dia e uma por dia durante um mês. As versões com nome ficam até você
excluí-las (no máximo 50). Todas as versões juntas ocupam no máximo 96 MiB, e
os arquivos que não mudaram são armazenados uma só vez. Quando o armazenamento
do navegador está cheio, as versões automáticas mais antigas saem primeiro.

## GitHub {#github}

**Home ▸ History ▸ GitHub** abre um painel que carrega a pasta de uma
apresentação de um repositório do GitHub e propõe suas mudanças de volta como
um pull request. Preencha:

- **Repositório (dono/nome)**;
- **Branch (vazio: o padrão)**;
- **Pasta (vazia: a raiz)**: a pasta deve conter um `deck.md`.

Acima desses campos o painel mostra o **quadro git**: um pequeno painel em que
você escolhe a forja, digita o **token** (mascarado), marca **Lembrar neste
dispositivo** e clica em **Conectar** (ou **Entrar com o GitLab**) ou
**Desconectar**. O quadro é uma página de outro site,
`go-slides-usercontent.github.io`, incorporada ao playground: **a página do
playground nunca vê o seu token**.

O token é um token de acesso pessoal de granularidade fina
([criar um](https://github.com/settings/personal-access-tokens/new)), limitado
aos repositórios escolhidos, com as permissões Contents de leitura e escrita e
Pull requests de leitura e escrita (Metadata de leitura está implícito). Depois:

- **Carregar** substitui o projeto pela pasta do GitHub. Undo desfaz, e o
  projeto como estava fica guardado como versão «Do GitHub».
- **Propor mudanças…** pede um título, cria um novo branch
  `slides/<título>-<data-hora>` e abre um pull request. Nunca escreve em um
  branch existente nem no branch padrão. **Abrir o pull request** leva você até
  ele.
- **Procurar mudanças** compara a versão mais recente no GitHub com a sua na
  vista Compare. **Ficar com as deles** a carrega.

Nenhum proxy é usado para o GitHub: o quadro fala diretamente com ele.

O token é digitado no quadro e permanece no site dele. A página do playground
não consegue lê-lo (o quadro é outra origem, com armazenamento próprio),
nenhum comando que a página possa enviar o devolve, e a Content Security Policy
da página não lhe permite nenhuma conexão com o GitHub nem com o GitLab. O
quadro envia o token somente a `api.github.com` (ou `gitlab.com`). Ele nunca é
salvo no rascunho, nas versões, nos links de compartilhamento nem nas
exportações, e nunca chega à pré-visualização dos slides nem ao worker de
renderização. **Lembrar neste dispositivo** o guarda no localStorage próprio do
quadro, em `go-slides-usercontent.github.io`, e não no armazenamento do
playground, até **Desconectar** (ou Forget my drafts).

**Se você usou uma versão anterior:** um token que ela guardava no próprio
armazenamento da página do playground é removido na primeira vez que você abre
esta versão, e o painel avisa. Digite-o de novo no quadro (e marque Lembrar se
quiser que ele seja mantido).

**Propor mudanças…** é uma escrita, então o quadro pergunta antes: ele mostra o
título, o repositório e o número de arquivos, e espera seu clique no próprio
botão **Permitir** (no máximo dois minutos, uma proposta por vez). Nada é
escrito antes do seu clique.

Restam dois riscos, porque são inerentes a uma página que incorpora um quadro.
Um script injetado na página do playground pode ler o nonce do quadro e enviar
os comandos do quadro: ler a pasta que você carregou ou pedir uma proposta (que
você ainda precisa permitir no quadro). E pode desenhar um painel falso sobre o
quadro e pedir o seu token por conta própria: se o quadro parecer errado ou não
estiver onde deveria, não digite o token. Nenhum dos dois permite ler o token
que o quadro guarda. Prefira um token limitado aos repositórios de que você
precisa. Para equipes, o modo servidor auto-hospedado
(em breve) mantém os tokens no servidor.

### GitLab.com {#gitlab}

O painel tem uma lista de forjas: **GitHub** e **GitLab.com**. Para o GitLab,
informe um token de acesso pessoal com o escopo `api`, ou clique em **Entrar com
o GitLab** (código de autorização OAuth com PKCE, sem segredo). O login abre uma
pequena janela do site do quadro e volta a ele; a página do playground em que
você trabalha não é abandonada. Seus tokens ficam só na memória do quadro
(nunca lembrados no dispositivo, mesmo com **Lembrar neste dispositivo**), se
renovam sozinhos antes de expirar, e **Desconectar** os revoga no GitLab. Um
login não iniciado pelo quadro, ou um link reutilizado, é recusado. O login fica disponível quando o aplicativo GitLab do
site estiver registrado; até lá, use um token de acesso pessoal.

As requisições vão só para `api.github.com` e `gitlab.com`, e só a partir do
quadro: a página do playground não faz nenhuma e não envia referrer.

Por que `api`? A
[tabela de escopos](https://docs.gitlab.com/security/tokens/access_token_scopes/)
do GitLab diz de `write_repository` «Does not support API authentication»: serve
só para git sobre HTTP, que o GitLab serve sem CORS, então um navegador não pode
usá-lo. As propostas usam a API REST, que exige `api`, «complete read and write
access to the API». Um token `api` roubado pode fazer tudo o que seu usuário pode
fazer pela API, em todos os projetos que esse usuário alcança. Prefira um token
de acesso de projeto ou de grupo limitado aos repositórios das apresentações, ou
uma conta do GitLab dedicada que seja apenas membro dos projetos das
apresentações.

## Segurança {#security}

A apresentação é exibida em um iframe em sandbox, sem a origem da página: ele
não alcança nem o Playground nem o armazenamento dele, não envia nenhum
formulário e não navega para lugar algum. A página dele carrega uma Content
Security Policy que executa somente os scripts dela. O HTML bruto da
apresentação segue a lista de permissões do Marp, como o `slides` faz por
padrão. Uma renderização que dura mais de 5 segundos (um template em laço) é
interrompida.

A janela do apresentador mantém essa sandbox. Ela é aberta pelo próprio
Playground e mostra a apresentação em um iframe em sandbox do mesmo modo. Os dois
iframes não podem compartilhar armazenamento, então o Playground repassa o
número do slide entre eles em mensagens. Nenhum frame que executa a
apresentação recebe a origem do Playground.

A mídia do projeto chega aos frames em sandbox como bytes em uma mensagem, e
cada frame faz com eles as suas próprias URLs `blob:`, na sua própria origem. Uma
URL `blob:` feita pelo Playground não carrega ali, e a política dos frames ganha
apenas `blob:` para mídia e fontes.

O Playground pode ser servido a partir da sua própria máquina, atrás de um
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
