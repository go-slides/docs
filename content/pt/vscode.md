---
title: "VS Code"
weight: 55
description: "A extensão do go-slides para o VS Code: pré-visualização, apresentação, problemas, autocompletar e exportação, como o Marp for VS Code faz para o Marp."
tags: [vscode, editor]
---

{{< callout type="info" >}}
A extensão é o [go-slides/vscode](https://github.com/go-slides/vscode), com
versionamento próprio (v0.5.0). Ela renderiza com o go-slides v0.14.0 (o pacote
`editor`, as linhas de origem da rolagem sincronizada, e a visão geral e a visão
do apresentador). Ela é instalada a partir do `go-slides-0.5.0.vsix` anexado à
sua release no GitHub: a publicação no VS Code Marketplace e no Open VSX está
preparada, mas ainda não está lá.
{{< /callout >}}

A extensão é a contraparte do go-slides para o
[Marp for VS Code](https://github.com/marp-team/marp-vscode): a apresentação
aparece na pré-visualização de Markdown do VS Code como slides, os erros dela no
painel Problems, e as diretivas, os kinds e as funções de template são
autocompletados enquanto você digita.

## Instalar {#install}

Baixe `go-slides-<version>.vsix` das
[releases](https://github.com/go-slides/vscode/releases) e execute
*Extensions: Install from VSIX...*, ou:

```sh
code --install-extension go-slides-0.5.0.vsix
```

A pré-visualização não precisa de mais nada. Exportar precisa do comando
`slides`:

```sh
go install github.com/go-slides/slides/cmd/slides@latest
```

## Pré-visualização {#preview}

Abra uma apresentação e *Markdown: Open Preview to the Side*. Uma apresentação é
um arquivo Markdown cujo front matter diz `slides: true`:

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

Arquivos com `marp: true` também são pré-visualizados, a menos que o Marp for VS
Code esteja instalado: as duas extensões renderizariam o mesmo documento. A
configuração `goSlides.preview.marpDocuments` decide: `auto` (o padrão),
`always` ou `never`.

A pré-visualização é o próprio go-slides, compilado para WebAssembly e carregado
pela extensão: nenhum binário, e funciona no [vscode.dev](https://vscode.dev).
Ela lê do workspace o que a apresentação usa — `slides.yaml` ou `slides.hcl` ao
lado da apresentação, os arquivos de [kinds]({{< relref "kinds" >}}) que eles
nomeiam, os arquivos que um [template]({{< relref "templates" >}}) inclui, o CSS
dos temas — e somente da pasta de workspace da apresentação.

## Rolagem sincronizada {#scroll-sync}

Desde a v0.2.0 da extensão (go-slides v0.7.0), o editor e a pré-visualização
rolam juntos, nos dois sentidos, como fazem para o Markdown simples: a
pré-visualização acompanha o slide e o bloco sob o cursor do editor, e rolar a
pré-visualização move o editor. As configurações são as do próprio VS Code,
`markdown.preview.scrollPreviewWithEditor` e
`markdown.preview.scrollEditorWithPreview`.

As linhas são rastreadas através dos templates. Um slide que um `{{ range }}`
gera acompanha a linha do seu template, e o que um `{{ include }}` traz acompanha
a linha do include. Uma apresentação dividida por `headingDivider` sincroniza
pelos seus títulos.

## Apresentar {#presenting}

Desde a v0.4.0 da extensão (go-slides v0.9.0), uma apresentação pode ser
apresentada sem sair do VS Code. Com a apresentação no editor, pela Command
Palette ou pelo menu `…` do editor:

| Comando | Abre |
|---|---|
| *go-slides: Open Overview* | um painel com todos os slides em uma grade. Um slide escolhido mostra o seu código-fonte no editor; `Esc` fecha o painel |
| *go-slides: Start Presenter View* | os slides e, ao lado deles, a visão do apresentador: o próximo passo, as notas do apresentador e um cronômetro |

Os painéis ficam ao lado do editor e acompanham uns aos outros: troque de slide
em um e os outros mudam junto. Eles também acompanham as suas edições na
apresentação. A tecla `p` nos slides abre a visão do apresentador.

Para apresentar em outra tela, mova o painel dos slides para lá com *View: Move
Editor into New Window*.

Cada painel executa a apresentação em um frame em sandbox (sem
`allow-same-origin`) sob uma política `default-src 'none'`: o frame carrega
imagens, mídia, estilos e fontes e nada mais, e executa somente os scripts que o
go-slides incorporou à página, nomeados por hash. Um script que a apresentação
escreveu não é executado.

## Problemas, autocompletar, hover {#problems-completion-hover}

- **Problemas**: erros de template, YAML e [HCL]({{< relref "hcl" >}}), na linha
  em que estão, na apresentação ou no arquivo de kinds de onde vêm.
- **Autocompletar**: as chaves de diretiva no front matter e em comentários
  `<!-- -->` (com `_` para uma diretiva pontual), os valores de `kind` e `_kind`
  (com o que cada kind define), `theme`, `size`, `paginate`, `math`,
  `headingDivider`; dentro de `{{ }}`, as funções de template e os campos de
  dados.
- **Hover**: o que uma diretiva faz, o que um kind define, os argumentos de uma
  função de template.

## Exportar {#export}

*go-slides: Export Slide Deck...* (também na barra de título do editor) escreve
HTML, PDF vetorial, PPTX e ODP editáveis, PNG ou JPEG com o comando `slides`,
como `slides -o deck.pdf deck.md` faria; veja
[Linha de comando]({{< relref "cli" >}}). Requer um workspace confiável. Se o
`slides` não for encontrado, a extensão oferece executar o `go install` acima em
um terminal.

Desde a v0.5.0 da extensão (go-slides v0.14.0), o PDF que ela escreve é
**acessível**: com o `slides` v0.12.0 ou posterior, o PDF vetorial é etiquetado
para leitores de tela e declara PDF/UA-1, com o `title` e o `lang` da
apresentação e uma seção por slide. `slides --pdf-tags=false` escreve o arquivo
sem etiquetas. Atualize o comando com o `go install` acima para obtê-lo.

## Configurações {#settings}

| Configuração | Padrão | |
|---|---|---|
| `goSlides.cliPath` | `""` | o comando `slides`; vazio: `slides` no `PATH`. Uma configuração de máquina: um workspace não pode defini-la |
| `goSlides.exportType` | `pdf` | o formato que a caixa de diálogo de exportação propõe primeiro |
| `goSlides.exportAutoOpen` | `true` | abrir o arquivo exportado |
| `goSlides.html` | `false` | permitir todo o HTML bruto em vez da lista de permissões do Marp; respeitado somente em um workspace confiável |
| `goSlides.preview.marpDocuments` | `auto` | pré-visualizar apresentações `marp: true`: `auto`, `always`, `never` |

## Segurança {#security}

A apresentação e o workspace não são confiáveis. O HTML bruto segue a lista de
permissões do Marp, a menos que `goSlides.html` esteja ativado em um workspace
confiável (um `slides.yaml` com `html: true` não o ativa na pré-visualização), e
a pré-visualização do VS Code não executa nenhum script da apresentação. A
pré-visualização lê somente arquivos da pasta de workspace da apresentação, não
segue nenhum link simbólico, e os templates não veem nenhuma variável de
ambiente. Exportar executa um programa, então espera a confiança do workspace, e
o programa nunca é nomeado pelo workspace.

## Desempenho {#performance}

Na apresentação de referência de 100 slides do go-slides, no VS Code 1.142 em um
Apple M4 Max, com o Marp for VS Code 3.6.1 medido na mesma sessão:

| | go-slides | Marp for VS Code |
|---|---:|---:|
| ativação (carregamento do módulo WebAssembly) | 189 ms | — |
| primeira pré-visualização | 52 ms | 78 ms |
| pré-visualização após uma edição (mediana) | **19 ms** | 23–24 ms |
| `.vsix` | 7,5 MB | 36,3 MB |

Os números são da extensão v0.3.0 (go-slides v0.8.0) em uma apresentação de 100
slides, digitando um caractere no slide 50, a partir de rodadas intercaladas sob
uma carga de 1 minuto abaixo de 8. A primeira pré-visualização é mais rápida que
a do Marp, e a pré-visualização após uma edição agora também é mais rápida que a
do Marp.

### Na v0.8.0 e na extensão v0.3.0 {#in-v080-and-extension-v030}

O go-slides v0.8.0 guarda os slides renderizados de cada apresentação aberta e,
após uma edição, renderiza de novo somente aqueles cuja fonte ou estado herdado
mudou (diretivas, número de página, total, ids de títulos, referências de links,
linhas de origem); a saída é a mesma de uma renderização completa, byte a byte.
Digitar no slide 50 de 100 renderiza de novo um slide: 2,5 ms nativamente,
11,5 ms em WebAssembly. A pré-visualização do VS Code então transforma (morph)
somente o DOM do slide alterado, e o teste de fumaça da extensão digita um
caractere no slide 50 entre renderizações cronometradas, para as duas extensões.
A visão do apresentador da extensão sincroniza pela janela pai (`postMessage`).

### Na v0.7.1 e na extensão v0.2.2 {#in-v071-and-extension-v022}

O go-slides v0.7.1 mantém em cache temas analisados, folhas de estilo
empacotadas e código destacado entre renderizações, marca as linhas de rolagem
sincronizada sem ir byte por byte, e reutiliza seus writers: uma renderização da
apresentação de 100 slides aloca 5,7 MB, contra 52,7 MB da extensão v0.2.0. Com
isso, a extensão v0.2.2 abandona o override `GOGC=400`. Para comparação, a
extensão v0.2.1 levou 72 ms para a primeira pré-visualização e 44 ms após uma
edição com o `GOGC` padrão (36 ms com `GOGC=400`), e a v0.2.0 levou 47–51 ms.

A extensão v0.2.2 mediu 52 ms para a primeira pré-visualização e 24 ms após uma
edição, no mesmo nível dos 23 ms do Marp.
