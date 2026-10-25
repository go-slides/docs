---
title: "Playground"
weight: 15
description: "Écrivez une présentation dans votre navigateur et regardez les diapositives suivre : go-slides compilé en WebAssembly, avec complétion, aide et exemples."
tags: [playground, browser]
---

**[Ouvrir le playground →](https://go-slides.github.io/playground/)**

Le playground, c'est go-slides lui-même, en v0.9.0, compilé en WebAssembly et
exécuté dans votre navigateur. Tapez une présentation à gauche ; les diapositives suivent à
droite. Rien n'est envoyé à un serveur : ni la présentation, ni un lien de partage.

## Ce qu'il propose {#what-it-has}

- **Aperçu en direct** de chaque diapositive, rendu hors du thread de la page (dans un Web
  Worker) : la saisie ne l'attend jamais. Cliquez sur une diapositive pour amener l'éditeur
  à sa ligne ; déplacez le curseur et l'aperçu défile jusqu'à sa diapositive.
- **Diaporama** : la présentation elle-même, avec le clavier, le
  contrôleur à l'écran et les transitions.
- **Mode présentateur** : **Present** (ou `p` dans le diaporama) l'ouvre dans une
  seconde fenêtre : la diapositive courante, l'étape suivante, les notes de l'orateur (les
  `<!-- comments -->` de la présentation), une horloge et un chronomètre. Les touches fléchées, Espace,
  Début et Fin fonctionnent dans l'une ou l'autre fenêtre, et chacune suit l'autre, y compris
  quand vous modifiez la présentation. Si la fenêtre ne s'ouvre pas, autorisez les fenêtres surgissantes pour
  le site. La vue d'ensemble (`o`) n'existe que dans la page téléchargée.
- **Complétion** pendant la saisie : clés de directive (les ponctuelles, avec `_`, dans les commentaires),
  les valeurs de `kind`, `theme`, `size`, `paginate`, ainsi que les fonctions de gabarit
  et les champs de données dans `{{ }}`.
- **Aide** sur ce qui se trouve sous le curseur : une directive, un type de diapositive (kind) (ce qu'il définit),
  une fonction de gabarit. Le volet d'aide renvoie aussi vers ces pages.
- **Problèmes** sur leur ligne : les erreurs de gabarit, de YAML et de HCL sont mises en évidence
  dans l'éditeur et listées en dessous, et l'aperçu conserve la dernière présentation
  rendue avec succès.
- **Exemples** : une présentation d'accueil, des gabarits qui génèrent des diapositives, des *kinds*,
  un front matter HCL, des thèmes et des directives ponctuelles, ainsi que des mathématiques, des images et des
  arrière-plans (`![bg]`, un `![bg left:40%]` partagé, des filtres).
- **Fichiers** : le panneau à gauche de l'éditeur est le projet de la présentation, organisé
  comme le lisent la [commande]({{< relref "cli" >}}) et
  l'[extension VS Code]({{< relref "vscode" >}}) : `deck.md`,
  `slides.yaml`, `themes/`, `kinds/`, `media/`.
  - Cliquez sur un fichier pour le modifier. L'éditeur le colore selon son langage
    (Markdown, YAML, HCL, JSON, CSS) et signale ses problèmes.
  - Clic droit (ou Maj+F10) pour New File, New Folder, Rename (F2),
    Delete (avec confirmation) et New theme / New kind. Faites glisser un fichier sur un dossier
    pour l'y déplacer. Le clavier suit le modèle d'arborescence WAI-ARIA, recherche par saisie
    comprise.
  - **Project ▸ New theme** crée `themes/my-gaia.css` (ou à partir de default,
    uncover) : un `@import` du thème intégré et les règles que vous ajoutez. Il
    rejoint la liste Theme. **New kind** et **New data file** partent eux aussi
    de modèles. Chacun est listé dans `slides.yaml`, si bien que `slides build deck.md`
    le trouve.
  - **Project ▸ Download project (.zip)** enregistre le tout. Une fois décompressé,
    `slides build deck.md` le génère. **Open project** reprend un `.zip` ou un
    dossier. Une archive qui sort de son dossier (`..`, chemins absolus),
    contient plus de 256 fichiers ou plus de 8 Mio dans un fichier, ou se décompresse à plus de 100
    fois sa taille est refusée.
  - Tout est conservé dans ce navigateur avec le brouillon. Un projet peut contenir jusqu'à
    64 Mio, et la ligne sous les fichiers indique l'espace utilisé.
- **Médias** : déposez ou collez une image, une vidéo ou un son dans l'éditeur : il
  va dans `media/`, et la présentation le désigne : `![](media/photo.png)`. Déposez
  des fichiers sur un dossier du panneau pour les y placer. La présentation reste légère et
  lisible, et rien n'est téléversé. Le bac à sable de l'aperçu les charge
  sans être assoupli (la page lui transmet les octets). **Save as ▸ HTML
  page** les intègre à la page. Les images `https:` fonctionnent aussi, mais le navigateur
  les récupère sur leur serveur, qui voit alors votre adresse. L'aperçu
  n'envoie pas de referrer.
- **Volets redimensionnables** : le panneau de fichiers, l'éditeur, les vignettes des diapositives,
  l'aperçu et l'aide sont séparés par des barres que vous pouvez faire glisser (le pointeur
  prend la forme d'un curseur de redimensionnement au-dessus). Le volet voisin de la barre prend la place,
  les autres gardent leur taille. Double-cliquez sur une barre, ou appuyez sur Entrée dessus, pour
  refermer le volet puis le rouvrir. Au clavier, F6 passe d'une barre à l'autre,
  les touches fléchées déplacent une barre de 16 px, Début et Fin amènent le volet
  qui la précède à sa taille minimale ou maximale, et Échap revient à l'éditeur.
  L'onglet View permet toujours de masquer des volets. Les tailles que vous définissez sont conservées dans ce
  navigateur, avec le brouillon ; les tailles par défaut (190, 168 et 280 px, l'éditeur
  et l'aperçu se partageant le reste à raison de 6 pour 5) ne sont jamais enregistrées.
- **Mise en forme** : l'onglet Format modifie le Markdown à votre place, et chaque clic
  correspond à une étape d'annulation.
  - **Texte** : Bold, Italic, Strikethrough et Inline code agissent sur la sélection
    ou sur le mot sous le curseur, et basculent (Ctrl+B, Ctrl+I, Ctrl+Maj+X et
    Ctrl+E ; Cmd sur Mac). Un bouton apparaît enfoncé quand le curseur est à l'intérieur
    de ses marques. Markdown n'a pas de soulignement : Underline (Ctrl+U) entoure le texte de
    `<u>…</u>`, une balise que conserve la liste d'autorisation HTML par défaut de Marp : cela fonctionne donc avec
    le HTML brut désactivé.
  - **Paragraphe** : Heading 1 à 3 (Ctrl+Alt+1..3), puces (Ctrl+Maj+8),
    numérotation (Ctrl+Maj+7) et citation agissent sur les lignes sélectionnées et basculent.
    La règle horizontale est la balise `<hr>` : `---`, `***` et `___` commenceraient chacun
    une nouvelle diapositive dans Marp.
  - **Tableau** : une grille permet de choisir colonnes et lignes. Avec le curseur dans un tableau, vous
    ajoutez ou supprimez une ligne ou une colonne et alignez la colonne (`:--`, `:-:`,
    `--:`) ; le tableau est réécrit aligné après chaque modification.
  - **Slide text** à gauche, au centre ou à droite écrit un bloc `<style scoped>` dans
    la diapositive courante. Marp n'a pas de syntaxe d'alignement, et un attribut `style`
    ne passe pas la liste d'autorisation HTML par défaut, mais le bloc scoped si : cela
    fonctionne donc avec le HTML brut désactivé, et se voit dans l'aperçu, la page HTML et le PDF.
  - **Link** (Ctrl+K) écrit `[text](url)` avec l'url sélectionnée.
- **Coloration** des clés du front matter à toutes les profondeurs, en YAML (y compris
  dans les mappings `{a: 1}`) et en HCL.
- **Sélecteur de thème** : default, gaia, uncover, les thèmes du projet
  (`themes/*.css`) et les thèmes que définit la présentation (`/* @theme name */` dans un
  `<style>`). Il écrit `theme:` dans le front
  matter, en YAML ou en HCL, si bien qu'un lien partagé le conserve, et il suit le front
  matter quand vous le modifiez.
- L'aperçu **garde sa position** pendant la saisie : seules les diapositives qui
  ont changé sont remplacées, et un changement de thème rouvre l'aperçu sur la
  même diapositive.
- **Lien de partage** : la présentation voyage dans le lien lui-même (dans sa partie `#`,
  compressée), si bien qu'ouvrir le lien ouvre la présentation. Les liens transportent jusqu'à
  64 Kio compressés, et la présentation seule : quand le projet contient davantage
  (thèmes, *kinds*, médias), Share le signale et propose le `.zip`.
- **Save as** :
  - la page HTML, identique à celle qu'écrit `slides deck.md` ;
  - un **PDF** (vectoriel, avec texte sélectionnable et plan) ;
  - une présentation **PowerPoint** ou **OpenDocument** modifiable.

  Ils sont produits dans votre navigateur, par le même code que `slides --pdf`,
  `--pptx` et `--odp` : le PDF de la présentation de référence de 100 diapositives est
  identique, octet pour octet, à celui de la commande. Le premier export télécharge le moteur
  d'export (environ 21 Mio, une seule fois ; votre navigateur le conserve jusqu'à la prochaine mise à jour),
  avec une ligne de progression. La présentation de 100 diapositives prend ensuite environ 3 secondes. Une
  image `https:` n'est incluse que si son serveur l'autorise (CORS) ; les images `data:`
  le sont toujours.
- **Votre brouillon est conservé** :
  - La présentation est enregistrée dans ce navigateur pendant la saisie (IndexedDB, pour ce site
    uniquement) et revient quand vous rechargez la page. Rien n'est envoyé nulle part. Les
    fichiers texte du projet (thèmes, *kinds*, données, `slides.yaml`) sont conservés de la
    même façon, même si vous fermez l'onglet juste après avoir tapé ; les médias sont
    enregistrés au moment où ils sont déposés.
  - Un lien partagé ouvre sa présentation et conserve votre brouillon : **Draft ▸ Restore my
    draft** le fait revenir.
  - Un exemple, **Draft ▸ Reset to the example** (qui demande confirmation) et
    un autre onglet ne l'effacent pas non plus : **Draft ▸ Restore replaced** le
    rétablit.
  - **Draft ▸ Forget my drafts** efface tout ce qui est enregistré.

Chaque export fait le rendu du projet qu'affiche l'aperçu : son `slides.yaml`,
ses thèmes, ses *kinds*, ses fichiers de données et ses médias, les médias étant intégrés au document.

Les images en tant que fichiers (`--images` de slides) ne sont pas produites dans le navigateur : utilisez la
[commande]({{< relref "cli" >}}) ou
l'[extension VS Code]({{< relref "vscode" >}}).

## Sécurité {#security}

La présentation s'affiche dans une iframe isolée, privée de l'origine de la page : elle ne peut
atteindre ni le playground ni son stockage, ne soumet aucun formulaire et ne navigue
nulle part. Sa page porte une Content Security Policy qui n'exécute que ses propres
scripts. Le HTML brut de la présentation suit la liste d'autorisation de Marp, comme `slides` le fait par
défaut. Un rendu qui dure plus de 5 secondes (un gabarit qui boucle)
est interrompu.

La fenêtre du présentateur conserve ce bac à sable. Elle est ouverte par le playground
lui-même et affiche la présentation dans une iframe isolée de la même façon. Les deux
iframes ne peuvent pas partager de stockage : le playground relaie donc le numéro de diapositive
de l'une à l'autre par messages. Aucun cadre qui exécute la présentation ne reçoit l'origine du
playground.

Les médias du projet parviennent aux cadres isolés sous forme d'octets dans un message, et
chaque cadre en fait ses propres URL `blob:`, dans sa propre origine. Une URL `blob:`
créée par le playground ne s'y charge pas, et la politique des cadres
n'ajoute que `blob:` pour les médias et les polices.

## Taille et vitesse {#size-and-speed}

Le module pèse 8,8 Mio tel que GitHub Pages le sert (gzip ; Pages ne
sert pas de brotli). Il est téléchargé une fois avec une barre de progression, puis mis en cache par le
navigateur jusqu'au prochain déploiement. Sur un Apple M4 Max avec Chrome 155, mesuré
sur la page en ligne, l'éditeur est prêt 2,1 s après l'ouverture de la page, dont 1,5 s
de téléchargement, et les premières diapositives apparaissent à 2,4 s. Une modification de
la présentation de référence de 100 diapositives atteint l'aperçu en 30 ms environ : go-slides
ne rend que les diapositives qui ont changé, et seules celles-ci sont remplacées dans
l'aperçu.
