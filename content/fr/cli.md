---
title: "Ligne de commande"
weight: 50
description: "Toutes les commandes et options de slides, et le fichier de configuration slides.yaml."
tags: [cli]
---

{{< callout type="warning" >}}
Cette page décrit la v0.9.0.
{{< /callout >}}

```text
slides [deck.md] [flags]
slides [command] [flags] deck.md
```

## Commandes {#commands}

| Commande | Rôle |
| --- | --- |
| `build` | écrit la présentation en HTML, en PDF ou en images (voir plus bas). **Commande par défaut** : `slides deck.md` équivaut à `slides build deck.md` |
| `expand` | affiche la présentation en Markdown Marp ordinaire : gabarits développés, types de diapositives (kinds) remplacés par les directives qu'ils représentent |
| `notes` | affiche les notes de l'orateur |
| `kinds` | liste les *kinds* que la présentation peut utiliser et ce qu'ils définissent ; `--json` (depuis la v0.6.0) les écrit sous la forme `[{"name", "extends", "directives": [{"key", "value"}]}]`, `extends` étant résolu |
| `themes` | liste les thèmes |
| `version` | affiche la version (tout comme `slides --version`) |
| `completion` | affiche un script de complétion pour le shell : `bash`, `zsh`, `fish` ou `powershell` |

`slides <command> --help` affiche les options d'une commande. Les options peuvent précéder
ou suivre la présentation : `slides deck.md -o out.html` et `slides -o out.html
deck.md` reviennent au même. Les options longues prennent deux tirets (`--html`) ; `-o` est la seule
option courte. `build` écrit `deck.html` à côté de `deck.md`, sauf si `-o` en décide
autrement.

## Options {#flags}

| Option | Commandes | Rôle |
| --- | --- | --- |
| `-o`, `--output file` | `build`, `expand`, `notes` | fichier de sortie ; `-` pour la sortie standard |
| `-I`, `--input-dir dir` | `build` | convertit récursivement les fichiers Markdown sous `dir` (depuis la v0.6.0) ; voir plus bas |
| `-P`, `--parallel n` | `build` | nombre de présentations générées simultanément (depuis la v0.6.0) : 5 par défaut, au plus le nombre de cœurs ; 1 les génère l'une après l'autre |
| `-w`, `--watch` | `build` | régénère quand une présentation, un thème, un fichier de *kinds* ou de données, ou la configuration change, et recharge les pages ouvertes (depuis la v0.6.0) ; voir plus bas |
| `-s`, `--server dir` | `build` | sert un répertoire en convertissant ses présentations à la demande (depuis la v0.6.0) ; voir plus bas |
| `-p`, `--preview` | `build` | ouvre le résultat dans le navigateur par défaut (depuis la v0.6.0) : les pages HTML avec `--watch` (implicite), le serveur avec `--server` |
| `--host addr` | `build` | l'adresse sur laquelle écoute `--server`, `127.0.0.1` par défaut (depuis la v0.6.0) |
| `--template bespoke\|bare` | `build` | le gabarit de la page HTML (depuis la v0.6.0) : `bespoke`, par défaut, avec le script de présentation, le contrôleur à l'écran et les transitions, ou `bare`, les diapositives empilées (`--bare`) |
| `--bespoke.osc`, `--bespoke.progress`, `--bespoke.transition` | `build` | les options du gabarit bespoke, avec les valeurs par défaut de marp-cli (depuis la v0.6.0) : contrôleur activé, **barre de progression désactivée** (elle était auparavant toujours affichée), transitions activées ; `--bespoke.osc=false` en désactive une |
| `-c`, `--config-file file`, `--no-config-file` | `build` | le fichier de configuration de marp-cli, en JSON ou en YAML (depuis la v0.6.0) ; voir plus bas. `--config` désigne le fichier propre à `slides` |
| `--browser`, `--browser-path`, `--browser-protocol`, `--browser-timeout` | `build` | acceptées et ignorées avec un avertissement (depuis la v0.6.0) : il n'y a pas de navigateur |
| `--pptx-editable`, `--no-parallel` | `build` | les équivalents marp-cli de `--pptx` et de `-P 1` (depuis la v0.6.0) |
| `--config file` | toutes | fichier de configuration ; par défaut, `slides.yaml` à côté de la présentation, s'il existe |
| `--no-config` | toutes | ne lit aucun fichier de configuration |
| `--theme-set path` | toutes | un fichier CSS de thème, ou un répertoire qui en contient (répétable) |
| `--kinds file` | toutes | un fichier YAML de *kinds* (répétable) |
| `--data file` | toutes | un fichier YAML ou JSON de données de gabarit (répétable) |
| `--set key=value` | toutes | une valeur de gabarit (répétable) |
| `--no-template` | toutes | ne développe pas les gabarits Go |
| `--strict-yaml` | toutes | analyse les directives en YAML strict, comme Marpit seul (Marp Core les analyse avec souplesse) |
| `--no-emoji` | toutes | laisse les `:shortcodes:` et les emoji Unicode sous forme de texte au lieu de les dessiner en images (depuis la v0.4.0 : [Emoji]({{< relref "emoji" >}})) |
| `--html` | `build` | autorise tout le HTML brut ; par défaut, la liste d'autorisation de Marp Core |
| `--lang code` | `build` | le `lang` de la page quand la présentation n'en définit pas |
| `--bare` | `build` | empile les diapositives, sans le script de présentation |
| `--pdf` | `build` | écrit un PDF, une page par diapositive, avec un plan — texte et formes en vectoriel (également choisi par `-o file.pdf`) |
| `--pdf-raster` | `build` | écrit le PDF avec, pour chaque page, une seule image de la diapositive, comme le faisait la v0.3.0 |
| `--notes` | `build` | écrit les notes de l'orateur dans un fichier texte, les diapositives séparées par `---`, comme marp-cli (depuis la v0.6.0 ; également choisi par `-o file.txt`). La commande `notes` les affiche dans une autre mise en forme |
| `--pdf-notes` | `build` | ajoute les notes de l'orateur de chaque diapositive à sa page PDF sous forme d'annotation de type note, comme marp-cli (depuis la v0.6.0). Demande un PDF quand rien d'autre n'indique quoi écrire |
| `--pdf-outlines`, `--pdf-outlines.pages`, `--pdf-outlines.headings` | `build` | les signets du PDF (depuis la v0.6.0) : « Page N » pour chaque diapositive et les titres imbriqués dessous. Activés par défaut ici, contrairement à marp-cli ; `=false` en désactive un. Voir plus bas |
| `--images png\|jpeg` | `build` | écrit une image par diapositive : `deck.001.png`, … |
| `--image png\|jpeg` | `build` | écrit la première diapositive sous forme d'image (également choisi par `-o file.png` ou `-o file.jpg`) |
| `--pptx` | `build` | écrit un fichier PowerPoint avec ses notes de l'orateur (également choisi par `-o file.pptx`) : modifiable depuis la v0.5.0, une image par diapositive en v0.4.0 |
| `--pptx-raster` | `build` | écrit un fichier PowerPoint composé d'une image pleine diapositive par diapositive, comme marp-cli (depuis la v0.5.0) ; `--image-scale` vaut 2 par défaut |
| `--odp` | `build` | écrit une présentation OpenDocument avec ses notes de l'orateur (également choisi par `-o file.odp`) : modifiable depuis la v0.5.0 |
| `--odp-raster` | `build` | écrit une présentation OpenDocument composée d'une image par diapositive (depuis la v0.5.0) |
| `--jpeg-quality n` | `build` | qualité JPEG, de 1 à 100 (85 par défaut) |
| `--image-scale n` | `build` | pixels physiques par pixel CSS pour les images et le PDF (1 par défaut) ; 2 donne 2560×1440 en 16:9, texte et vecteurs étant dessinés à cette densité, et non agrandis |
| `--allow-local-files` | `build` | permet aux images et au PDF de lire des fichiers locaux, uniquement sous le répertoire de la présentation ; refusé par défaut, comme avec marp-cli |
| `--theme name\|file.css` | `build` | remplace le thème de la présentation (depuis la v0.6.0) : un nom de thème, ou un fichier CSS, qui devient le thème quel que soit le `@theme` qu'il déclare ; un répertoire est refusé, utilisez `--theme-set` |
| `--title`, `--description`, `--author`, `--keywords`, `--url`, `--og-image` | `build` | les métadonnées de la page (depuis la v0.6.0), prioritaires sur les directives `title`, `description`, `author`, `keywords`, `url` et `image` de la présentation ; voir plus bas |

## Plusieurs présentations, et un répertoire {#several-decks-and-a-directory}

Depuis la v0.6.0, `build` accepte plusieurs présentations, un répertoire (les présentations qu'il contient,
trouvées récursivement) ou un motif, et en génère jusqu'à `--parallel` simultanément.
`-` désigne l'entrée standard ; contrairement à marp-cli, un tube n'est lu que si `-` le demande.

Un motif est interprété comme le fait le glob de marp-cli (`**` et `{a,b}` depuis la v0.7.0 ;
mettez-le entre guillemets pour que le shell le laisse à `slides`) : `*`, `?`, `[a-z]` et
`{md,txt}` dans un nom, `**` pour un nombre quelconque de répertoires, si bien que
`'docs/**/*.md'` prend les présentations de toute l'arborescence. Un joker n'atteint pas un nom
qui commence par un point, rien n'est cherché dans `node_modules`, la casse compte
(`*.md` ne trouve pas `UPPER.MD`), et tout ce qu'un motif sélectionne est converti,
quelle que soit son extension. Un motif qui ne sélectionne rien est une erreur.

- **`-I dir`** trouve les fichiers Markdown (`.md`, `.mdown`, `.markdown`,
  `.markdn`, en minuscules) sous `dir`, comme marp-cli : `node_modules` et
  les répertoires commençant par un point sont ignorés. Avec `-o out`, chaque présentation est écrite à son
  chemin sous `dir`, reproduit sous `out` : `dir/sub/b.md` donne
  `out/sub/b.html` (`b.pdf`, `b.001.png`, `b.txt`, … selon les options).
  Sans `-o`, les sorties sont écrites à côté de leurs présentations. Nommer des présentations en plus de
  `-I` est une erreur, de même que `-o file` pour plusieurs présentations (`-o -` est autorisé).
- **`-P n`** génère `n` présentations simultanément. Le rendu d'une présentation exploite déjà les
  cœurs, donc `n` est plafonné à leur nombre ; avec `-o -`, les présentations sont générées l'une après
  l'autre. Les messages d'une présentation restent groupés, et une présentation en échec n'arrête
  pas les autres : le code de sortie vaut 1 et le message liste les échecs.

## Transitions et page de présentation {#transitions-and-the-presenting-page}

Depuis la v0.6.0, la directive `transition` de marp-cli fonctionne : `<!-- transition:
fade 1s -->`, un nom et une durée facultative, pour la diapositive et les suivantes,
ou `<!-- _transition: cover -->` pour une seule. Les noms intégrés sont ceux de marp-cli :
`clockwise`, `counterclockwise`, `cover`, `coverflow`, `cube`, `cylinder`,
`diamond`, `drop`, `explode`, `fade`, `fade-out`, `fall`, `flip`, `glow`,
`implode`, `in-out`, `iris-in`, `iris-out`, `melt`, `overlap`, `pivot`, `pull`,
`push`, `reveal`, `rotate`, `slide`, `star`, `swap`, `swipe`, `swoosh`, `wipe`,
`wiper` et `zoom` ; `none` la désactive. La page la joue avec la
**View Transition API** du navigateur quand le spectateur quitte vers l'avant la diapositive qui la porte,
et quand il y revient ; les navigateurs sans cette API changent de diapositive comme avant,
et un spectateur qui demande des mouvements réduits obtient un simple fondu. Un thème peut apporter
ses propres images clés, nommées `marp-transition-NAME` (ou `marp-outgoing-` et
`marp-incoming-`, et `-backward`), qui passent avant celles qui sont intégrées.

La page comporte le **contrôleur à l'écran** — précédent, page, suivant, plein écran —,
qui apparaît quand le pointeur bouge ou qu'une touche est pressée et disparaît au bout d'un moment
(`--bespoke.osc=false` le supprime), ainsi que, avec `--bespoke.progress`, une barre de
progression. **La barre de progression est désormais désactivée par défaut**, comme dans marp-cli ; elle était auparavant
toujours affichée.

### Mode présentateur {#presenter-view}

Depuis la v0.7.0, la page dispose du mode présentateur de marp-cli. Appuyez sur `p`, ou sur le dernier
bouton du contrôleur : une fenêtre s'ouvre avec la diapositive courante, l'étape
suivante, vos notes de l'orateur, la page, l'heure et un chronomètre (cliquez sur le chronomètre
pour le remettre à zéro ; `+` et `-` redimensionnent les notes ; la barre entre les diapositives et les
notes peut être déplacée). Les deux fenêtres se suivent dans les deux sens : les
touches fléchées, un clic sur l'étape suivante ou sur Précédent/Suivant dans l'une font avancer
l'autre, un fragment à la fois.

Il fonctionne comme celui de marp-cli, d'après ses sources 4.5.1 : la fenêtre est
`?view=presenter&sync=KEY` du même fichier, l'étape suivante est la même page en
`?view=next` dans un cadre, et les fenêtres partagent leur position par
`localStorage` (`bespoke-marp-sync-KEY`) : il ne demande ni serveur ni
réseau, et ne fonctionne pas dans un navigateur dépourvu de `localStorage` (le
bouton le signale). Les notes sont du texte : `&`, `<` et `>` sont échappés, si bien qu'une note
ne peut jamais être du balisage. Elles figurent dans la page, masquées, pour que le mode présentateur
les reprenne, ce qui les tient à l'écart de ce qu'un lecteur d'écran lit au public.
Les pages `--bare` n'ont rien de tout cela. Ce mode ajoute environ 10 Ko à une page.

### Vue d'ensemble {#overview-view}

Depuis la v0.8.0, appuyez sur `o` (ou `Escape` : marp-cli bascule avec les deux), sur le
bouton en grille du contrôleur ou sur le compteur de pages du mode présentateur : toutes les diapositives
s'affichent en entier dans une grille par-dessus la page. Les touches fléchées choisissent une diapositive (haut et bas
par rangée), `Home` et `End` vont aux extrémités, et les autres fenêtres suivent ;
`Enter`, `Space` ou un clic choisit une diapositive et, par-dessus la page, ferme la
vue d'ensemble. Par-dessus le mode présentateur, un choix la laisse ouverte (le présentateur choisit
la diapositive à montrer et continue à regarder). Le bouton « Close overview », `o` et
`Escape` la ferment. C'est celle de marp-cli : la page, de nouveau dans un cadre, à
`?view=overview&sync=KEY`, qui partage la position par `localStorage`, sans
réseau. Les diapositives de la grille peuvent recevoir le focus et portent le libellé « Slide N », et la
vue d'ensemble fermée est inerte : les personnes qui naviguent au clavier ou au lecteur d'écran n'y
tombent jamais. Les deux vues ajoutent ensemble environ 15 Ko à une page.

### Fichiers de configuration {#configuration-files}

marp-cli lit ses options dans un fichier du répertoire de travail, et `slides` aussi,
quand le fichier contient des données : `.marprc` (YAML, le JSON étant aussi du YAML),
`.marprc.json`, `.marprc.yaml`, `.marprc.yml`, `.config/marprc*`, ou la
clé `marp` de `package.json` ; `--config-file` (`-c`) désigne n'importe quel fichier JSON ou YAML,
`marp.config.json` par exemple. Seul le répertoire de travail est examiné, pas
ses parents, comme dans la recherche de marp-cli. Un fichier JavaScript ou TypeScript
(`marp.config.js` ...) est un programme, que `slides` n'exécute pas : il le signale et
continue. Les options sont celles de la documentation de marp-cli qui ont un sens
ici : `title`, `description`, `author`, `keywords`, `url`, `ogImage`, `lang`,
`theme`, `themeSet`, `html`, `allowLocalFiles`, `template`, `bespoke`,
`inputDir`, `output`, `parallel`, `pdf`, `pptx`, `notes`, `pdfNotes`,
`pdfOutlines`, `images`, `image`, `imageScale`, `jpegQuality`, `watch`,
`server`, `preview` ; les chemins sont relatifs au fichier. La ligne de commande l'emporte.
`engine`, `options` et `browser*` nécessitent le moteur de Marp et un navigateur ; ils sont
ignorés avec un avertissement, tout comme les options `--browser*` : les scripts écrits pour
marp-cli continuent de fonctionner. (`slides.yaml` et `slides.hcl`, avec les *kinds* et les
données de gabarit, sont distincts et lus comme avant.)

## Surveillance, serveur et aperçu {#watch-serve-and-preview}

Depuis la v0.6.0, avec les routes et le protocole de rechargement à chaud de marp-cli (lus dans ses
sources 4.5.1 et comparés à un `marp -s` en cours d'exécution).

**`slides --server dir`** (`-s`) sert le répertoire sur `http://127.0.0.1:8080/`
(la variable d'environnement `PORT` change le port ; un port déjà utilisé est une
erreur, signalée comme telle) :

| Requête | Réponse |
| --- | --- |
| `/deck.md` | la présentation en HTML, convertie à la demande |
| `/deck.md?pdf`, `?pptx`, `?png`, `?jpg` (ou `?jpeg`), `?notes` (ou `?txt`) | la présentation convertie selon la requête ; le PPTX est servi en pièce jointe |
| un répertoire | son `index.md` ou son `PITCHME.md` en tant que page, sinon un index avec un lien PDF et un lien PPTX pour chaque présentation |
| tout autre fichier | le fichier |
| une présentation dont la conversion échoue | 503, avec l'erreur |

Une page ouvre un WebSocket vers `/.__marp-cli-watch-notifier__/<id>`, où l'id est
le SHA-256 du chemin de la présentation ; le serveur envoie `ready`, puis `reload` quand la
présentation, ou quoi que ce soit sous le répertoire, change. Un répertoire demandé sans
sa barre oblique finale est redirigé vers celle-ci (marp-cli y sert son index, ce qui
casse les liens relatifs de la page).

**`--watch`** (`-w`) génère les présentations, puis régénère celle qui change
et, quand un fichier `--theme-set`, un fichier de *kinds* ou de données, ou la configuration
change, toutes. Les pages qu'il écrit ouvrent une socket sur l'interface de bouclage (port
37717, ou le suivant qui est libre) et se rechargent quand leur présentation est régénérée. Une présentation
dont la génération échoue est signalée et reste surveillée. **`--preview`** (`-p`) ouvre
le navigateur par défaut sur le serveur, ou sur les pages HTML tout juste générées ; il n'a pas de
fenêtre propre, puisque `slides` n'a aucun navigateur à piloter.

Le serveur est une porte d'entrée sur la machine : il est donc plus strict que celui de marp-cli, qui
écoute sur toutes les interfaces et ne vérifie rien :

- il écoute sur **127.0.0.1** ; `--host` change cela, avec un avertissement, et une
  adresse générique ne permet plus de vérifier l'en-tête `Host` ;
- l'en-tête **`Host`** doit être un nom de bouclage (`localhost`, `127.0.0.1`,
  `[::1]`, `*.localhost`) ou l'hôte sur lequel il écoute, ce qui bloque le DNS rebinding ;
- l'**`Origin`** d'un WebSocket doit être celle du serveur ; les sockets du mode surveillance
  acceptent aussi l'origine `null` d'une page ouverte depuis un fichier ; une requête venant
  d'un autre site qui n'est pas une navigation (un script, une image, un `fetch`) est
  refusée ;
- un chemin est normalisé, ses liens symboliques sont résolus, et le résultat doit se trouver
  dans le répertoire : `..` et un lien qui en sort sont refusés, de même que
  `?pdf` sur une présentation extérieure ; les noms commençant par un point (`.git`, `.env`) ne sont pas
  servis ;
- seulement `GET` et `HEAD` ; quelques conversions à la fois ; au plus 4096 répertoires
  surveillés, et aucun dans `node_modules` ni sous un répertoire commençant par un point.

## Métadonnées de la page {#page-metadata}

La ligne de commande l'emporte sur la présentation, comme dans marp-cli, où les options sont
ajoutées à la présentation sous la forme d'un dernier commentaire de directive. Règles lues dans marp-cli
4.5.1 et vérifiées par des tests sur sa sortie (depuis la v0.6.0) :

- Une présentation qui ne définit pas de `title` prend le texte de son premier titre du niveau le plus
  élevé, tel qu'il est écrit (`# **Bold** title` donne `**Bold** title`). Sans
  titre, il n'y a pas de `<title>`.
- Les `keywords` sont découpés aux virgules, débarrassés des espaces et dédoublonnés ; une liste dans le front matter
  est reprise telle quelle. Comme marp-cli, la page porte la balise `keywords` à partir
  du deuxième mot-clé, et non pour un seul.
- Une `url` invalide (non absolue) est ignorée avec un avertissement, et la valeur
  précédente conservée ; une URL valide écrit aussi `<link rel="canonical">` et `og:url`.
- `author` et `description` sont accompagnés de leurs équivalents Open Graph ; `og:image:alt`
  et la grande carte Twitter nécessitent à la fois un titre et une image.

## Images, PDF, PPTX et ODP {#images-pdf-pptx-and-odp}

`-o` avec une extension `.pdf`, `.pptx`, `.odp`, `.png` ou `.jpg` choisit le format, comme
marp-cli. Les répertoires de sortie sont créés. Un PDF comporte une page par diapositive, à la
taille de la diapositive, et un plan des diapositives et de leurs titres.

### Notes et plan du PDF {#pdf-notes-and-outline}

Depuis la v0.6.0, le PDF prend en charge les deux options de marp-cli, comparées avec
les PDF qu'écrit marp-cli 4.5.1 pour la même présentation :

- `--pdf-notes` place les notes de l'orateur de chaque diapositive sur sa page sous forme d'annotation
  de type note : une icône Note jaune en bas à gauche, dont le contenu est constitué des
  notes (plusieurs commentaires séparés par une ligne vide) et dont le titre est l'`author`
  de la présentation. `--pdf-notes` seul demande un PDF, comme dans marp-cli.
- Le plan contient un signet en gras « Page N » pour chaque diapositive et, en dessous, les
  titres de la diapositive imbriqués par niveau (`--pdf-outlines.pages` et
  `--pdf-outlines.headings` choisissent ces deux parties). **Le plan est activé par défaut
  ici et désactivé dans marp-cli.** `slides` en a toujours écrit un ; il coûte quelques
  centaines d'octets et rend une longue présentation navigable : ce choix par défaut est donc conservé.
  `--pdf-outlines=false`, ou le `--no-pdf-outlines` de marp-cli, le désactive, et
  `=false` sur `.pages` ou `.headings` supprime la partie correspondante. Un signet de titre renvoie
  à sa diapositive ; marp-cli fait en plus défiler jusqu'au titre.

Le rendu est assuré par [go-webengine](https://github.com/go-webengine/engine)
et le PDF est écrit par [go-pdfkit](https://github.com/go-pdfkit/pdfkit) ;
les polices demandées par les thèmes intégrés (Lato, Roboto Mono) proviennent de
[go-opentype/fonts](https://github.com/go-opentype/fonts) et sont intégrées : aucun
rendu n'accède au réseau. Il n'y a pas de Chrome.

`--image-scale 2` fait le rendu à deux pixels physiques par pixel CSS. Les images locales sont
refusées sauf si `--allow-local-files` est indiqué, et même alors seuls les fichiers situés sous
le répertoire de la présentation sont lus : une présentation venue d'ailleurs ne peut pas atteindre `~/.ssh`.
Les arrière-plans multiples et partagés (`![bg left:40%]`) sont rendus comme les calques de Marp.

`--pptx` écrit la présentation avec la fonction `WritePresentation` de [go-ooxml](https://github.com/go-ooxml),
avec les notes de l'orateur et les métadonnées de la présentation. Depuis
la v0.5.0, les diapositives sont modifiables (voir ci-dessous) ; `--pptx-raster` écrit une image
pleine diapositive par diapositive, comme marp-cli, à une échelle de 2 par défaut que
`--image-scale` remplace. LibreOffice ouvre le résultat. Seul LibreOffice a été
testé ; PowerPoint et Keynote ne l'ont pas été.

`--odp` écrit une présentation OpenDocument, au moyen de la fonction
`WritePresentation` de go-odf, avec ses notes de l'orateur ; `--odp-raster` écrit une image
par diapositive.

### PPTX et ODP modifiables {#editable-pptx-and-odp}

Depuis la v0.5.0, les diapositives d'un PPTX ou d'un ODP sont **modifiables** — le PPTX de marp-cli
est une image par diapositive, et son option expérimentale `--pptx-editable` convertit
un PDF au moyen de LibreOffice. Ici, chaque diapositive est dessinée à partir de sa mise en page, celle-là même
que dessinent les exports PNG et PDF :

- chaque ligne de texte est une zone de texte placée à sa position dans la mise en page, dans les polices
  avec lesquelles la diapositive a été mise en page — **intégrées au fichier**, si bien qu'elle a le même
  aspect sur une machine dépourvue de Lato ou d'Inter (vérifié avec LibreOffice 25.8) —
  avec le gras, l'italique, les couleurs, l'espacement des lettres et les liens cliquables ;
- les arrière-plans, les bordures et les cadres arrondis sont des formes ; les dégradés sont des remplissages
  en dégradé dans le PPTX et des images dans l'ODP (un dégradé ODF n'a que deux
  couleurs) ;
- les images ne sont stockées qu'une fois, quel que soit le nombre de diapositives qui les affichent ;
- ce qu'une diapositive ne peut pas exprimer — un élément sous `opacity`, un `filter`, une
  ombre floue, une formule — devient une image de cet élément, son texte étant conservé
  par-dessus.

Rendue à nouveau par LibreOffice, une diapositive ne s'écarte que de 0,04 à 2 % des pixels de l'export
PNG. `--pptx-raster` et `--odp-raster` conservent une image par diapositive.

### PDF vectoriel {#vector-pdf}

Depuis la v0.4.0, le PDF est dessiné en vectoriel : go-webengine peint la diapositive mise en page
dans le PDF au lieu de la peindre en pixels. Le texte reste du texte, dans les polices avec lesquelles la
diapositive a été mise en page, intégrées sous forme de sous-ensembles avec une table `/ToUnicode` : on peut le
sélectionner, y rechercher et le copier. Les arrière-plans, les bordures et les dégradés sont des tracés
et des ombrages, une image affichée sur plusieurs diapositives n'est intégrée qu'une fois, et les liens sont
cliquables. La présentation de référence de 100 diapositives pèse 758 Kio, contre 837 Kio pour
marp-cli et 10 Mio pour le PDF en images de la v0.3.0, et elle est écrite en 0,3 s.

Un élément sans forme vectorielle ici — sous `opacity` < 1, avec un `filter`
ou une rotation, une `box-shadow` floue, un dégradé dont l'opacité s'estompe — est
dessiné comme un bitmap de ce seul élément, son texte étant posé dessus de façon invisible pour
rester sélectionnable. Les formules sont encore des bitmaps. `--pdf-raster` écrit une image
par page, comme le faisait la v0.3.0.

## `slides.yaml` {#slidesyaml}

Un fichier de configuration placé à côté de la présentation s'y applique : un dossier de présentations
peut ainsi en partager un seul :

```yaml
kinds: kinds.yaml        # a mapping of kinds, or file(s) holding one
themes:                  # theme CSS files or directories
  - themes/
data:                    # template data
  company: Example Inc.
html: false              # allow all raw HTML
lang: en
```

*Kinds* : les `kinds:` propres à la présentation remplacent ceux du fichier et de la ligne de commande.
Données : les `data` du fichier et de la ligne de commande remplacent celles du front
matter de la présentation, si bien qu'une même présentation peut être générée avec des données différentes.

## Prévu {#planned}

Pas encore réalisé : les formules en vectoriel dans le PDF. Quand elles arriveront, elles prendront la forme d'une
commande ou d'une option décrite ici.
