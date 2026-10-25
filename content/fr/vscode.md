---
title: "VS Code"
weight: 55
description: "L'extension go-slides pour VS Code : aperçu, présentation, problèmes, complétion et export, comme Marp for VS Code le fait pour Marp."
tags: [vscode, editor]
---

{{< callout type="info" >}}
L'extension est [go-slides/vscode](https://github.com/go-slides/vscode),
versionnée séparément (v0.4.0). Elle fait le rendu avec go-slides v0.9.0 (le paquet
`editor`, les lignes sources de la synchronisation du défilement, la vue d'ensemble et le
mode présentateur). Elle s'installe à partir du fichier `go-slides-0.4.0.vsix` joint à
sa version GitHub : sa publication sur le VS Code Marketplace et sur Open VSX est
prête, mais elle n'y figure pas encore.
{{< /callout >}}

L'extension est l'équivalent, pour go-slides, de
[Marp for VS Code](https://github.com/marp-team/marp-vscode) : la présentation s'affiche dans
l'aperçu Markdown de VS Code sous forme de diapositives, ses erreurs dans le panneau Problèmes, et
les directives, les types de diapositives (kinds) et les fonctions de gabarit se complètent pendant la saisie.

## Installation {#install}

Téléchargez `go-slides-<version>.vsix` depuis les
[versions publiées](https://github.com/go-slides/vscode/releases), puis lancez
*Extensions: Install from VSIX...*, ou :

```sh
code --install-extension go-slides-0.4.0.vsix
```

L'aperçu n'a besoin de rien d'autre. L'export nécessite la commande `slides` :

```sh
go install github.com/go-slides/slides/cmd/slides@latest
```

## Aperçu {#preview}

Ouvrez une présentation, puis *Markdown: Open Preview to the Side*. Une présentation est un fichier Markdown
dont le front matter indique `slides: true` :

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

Les fichiers avec `marp: true` sont eux aussi prévisualisés, sauf si Marp for VS Code est installé :
les deux extensions rendraient le même document. Le paramètre
`goSlides.preview.marpDocuments` en décide : `auto` (par défaut), `always` ou
`never`.

L'aperçu, c'est go-slides lui-même, compilé en WebAssembly et chargé par
l'extension : aucun binaire, et il fonctionne sur [vscode.dev](https://vscode.dev). Il lit
dans l'espace de travail ce qu'utilise la présentation — `slides.yaml` ou `slides.hcl` à côté de
la présentation, les fichiers de [kinds]({{< relref "kinds" >}}) qu'ils désignent, les fichiers qu'un
[gabarit]({{< relref "templates" >}}) inclut, le CSS des thèmes — et uniquement dans le
dossier de l'espace de travail de la présentation.

## Synchronisation du défilement {#scroll-sync}

Depuis la v0.2.0 de l'extension (go-slides v0.7.0), l'éditeur et
l'aperçu défilent ensemble, dans les deux sens, comme pour du Markdown ordinaire : l'aperçu
suit la diapositive et le bloc sous le curseur de l'éditeur, et
faire défiler l'aperçu déplace l'éditeur. Les paramètres sont ceux de VS Code lui-même,
`markdown.preview.scrollPreviewWithEditor` et
`markdown.preview.scrollEditorWithPreview`.

Les lignes sont suivies à travers les gabarits. Une diapositive générée par un `{{ range }}`
suit la ligne de son gabarit, et ce qu'apporte un `{{ include }}`
suit la ligne de l'inclusion. Une présentation découpée par `headingDivider` se synchronise sur ses
titres.

## Présentation {#presenting}

Depuis la v0.4.0 de l'extension (go-slides v0.9.0), une présentation peut être projetée
sans quitter VS Code. La présentation ouverte dans l'éditeur, depuis la palette de commandes
ou le menu `…` de l'éditeur :

| Commande | Ouvre |
|---|---|
| *go-slides: Open Overview* | un panneau avec toutes les diapositives dans une grille. Une diapositive choisie affiche sa source dans l'éditeur ; `Esc` ferme le panneau |
| *go-slides: Start Presenter View* | les diapositives et, à côté, le mode présentateur : l'étape suivante, les notes de l'orateur et un chronomètre |

Les panneaux se placent à côté de l'éditeur et se suivent : changez de diapositive dans l'un
et les autres changent avec lui. Ils suivent aussi vos modifications de la présentation. La touche `p`
dans les diapositives ouvre le mode présentateur.

Pour projeter sur un autre écran, déplacez-y le panneau des diapositives avec *View: Move
Editor into New Window*.

Chaque panneau exécute la présentation dans un cadre isolé (sans `allow-same-origin`) sous une
politique `default-src 'none'` : le cadre charge les images, les médias, les styles et les polices,
rien d'autre, et n'exécute que les scripts que go-slides a intégrés à la page,
désignés par leur empreinte. Un script écrit par la présentation ne s'exécute pas.

## Problèmes, complétion, survol {#problems-completion-hover}

- **Problèmes** : les erreurs de gabarit, de YAML et de [HCL]({{< relref "hcl" >}}), sur la
  ligne où elles se trouvent, dans la présentation ou dans le fichier de *kinds* dont elles proviennent.
- **Complétion** : les clés de directive dans le front matter et dans les commentaires `<!-- -->`
  (avec `_` pour une directive ponctuelle), les valeurs de `kind` et de `_kind` (avec ce que
  définit chaque *kind*), `theme`, `size`, `paginate`, `math`, `headingDivider` ; dans
  `{{ }}`, les fonctions de gabarit et les champs de données.
- **Survol** : ce que fait une directive, ce que définit un *kind*, les arguments d'une
  fonction de gabarit.

## Export {#export}

*go-slides: Export Slide Deck...* (également dans la barre de titre de l'éditeur) écrit du HTML,
du PDF vectoriel, du PPTX et de l'ODP modifiables, du PNG ou du JPEG avec la commande `slides`, comme
le ferait `slides -o deck.pdf deck.md` ; voir [Ligne de commande]({{< relref "cli" >}}).
Il nécessite un espace de travail approuvé. Si `slides` est introuvable, l'extension propose de
lancer le `go install` ci-dessus dans un terminal.

## Paramètres {#settings}

| Paramètre | Défaut | |
|---|---|---|
| `goSlides.cliPath` | `""` | la commande `slides` ; vide : `slides` dans le `PATH`. Paramètre de machine : un espace de travail ne peut pas le définir |
| `goSlides.exportType` | `pdf` | le format que la boîte de dialogue d'export propose en premier |
| `goSlides.exportAutoOpen` | `true` | ouvrir le fichier exporté |
| `goSlides.html` | `false` | autoriser tout le HTML brut au lieu de la liste d'autorisation de Marp ; pris en compte uniquement dans un espace de travail approuvé |
| `goSlides.preview.marpDocuments` | `auto` | prévisualiser les présentations `marp: true` : `auto`, `always`, `never` |

## Sécurité {#security}

La présentation et l'espace de travail ne sont pas dignes de confiance. Le HTML brut suit la liste d'autorisation de Marp,
sauf si `goSlides.html` est activé dans un espace de travail approuvé (un `slides.yaml` avec
`html: true` ne l'active pas dans l'aperçu), et l'aperçu de VS Code n'exécute aucun
script de la présentation. L'aperçu ne lit que les fichiers du dossier de l'espace de travail de la présentation,
ne suit aucun lien symbolique, et les gabarits ne voient aucune variable d'environnement.
L'export exécute un programme : il attend donc que l'espace de travail soit approuvé, et le programme n'est
jamais désigné par l'espace de travail.

## Performances {#performance}

Sur la présentation de référence de 100 diapositives de go-slides, dans VS Code 1.142 sur un Apple M4 Max,
avec Marp for VS Code 3.6.1 mesuré dans la même session :

| | go-slides | Marp for VS Code |
|---|---:|---:|
| activation (chargement du module WebAssembly) | 189 ms | — |
| premier aperçu | 52 ms | 78 ms |
| aperçu après une modification (médiane) | **19 ms** | 23–24 ms |
| `.vsix` | 7,5 Mo | 36,3 Mo |

Les chiffres concernent l'extension v0.3.0 (go-slides v0.8.0) sur une présentation de 100 diapositives,
en tapant un caractère dans la diapositive 50, sur des tours entrelacés avec une charge moyenne sur 1 minute
inférieure à 8. Le premier aperçu est plus rapide que celui de Marp, et l'aperçu après une
modification l'est désormais aussi.

### Dans la v0.8.0 et l'extension v0.3.0 {#in-v080-and-extension-v030}

go-slides v0.8.0 conserve les diapositives rendues de chaque présentation ouverte et, après une modification,
ne refait le rendu que de celles dont la source ou l'état hérité a changé (directives, numéro de
page, total, ids des titres, références de liens, lignes sources) ; la sortie est
identique à celle d'un rendu complet, octet pour octet. Taper dans la diapositive 50 sur 100 refait le rendu d'une seule
diapositive : 2,5 ms en natif, 11,5 ms en WebAssembly. L'aperçu de VS Code ne transforme alors
que le DOM de la diapositive modifiée, et le test de fumée de l'extension tape un
caractère dans la diapositive 50 entre deux rendus chronométrés, pour les deux extensions. Le
mode présentateur de l'extension se synchronise via la fenêtre parente (`postMessage`).

### Dans la v0.7.1 et l'extension v0.2.2 {#in-v071-and-extension-v022}

go-slides v0.7.1 met en cache, d'un rendu à l'autre, les thèmes analysés, les feuilles de style compactées et le code coloré,
marque les lignes de synchronisation du défilement sans parcourir les octets un à un, et réutilise
ses writers : un rendu de la présentation de 100 diapositives alloue 5,7 Mo, contre 52,7 Mo
pour l'extension v0.2.0. Grâce à cela, l'extension v0.2.2 abandonne le forçage `GOGC=400`.
À titre de comparaison, l'extension v0.2.1 prenait 72 ms pour le premier aperçu et 44 ms
après une modification avec le `GOGC` par défaut (36 ms avec `GOGC=400`), et la v0.2.0 prenait
47 à 51 ms.

L'extension v0.2.2 a mesuré 52 ms pour le premier aperçu et 24 ms après une modification,
à égalité avec les 23 ms de Marp.
