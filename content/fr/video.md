---
title: "Vidéo"
weight: 39
description: "Une vidéo dans une présentation : lue par le navigateur dans la page HTML, une image d'affiche avec un lien dans le PDF, intégrée dans PowerPoint et OpenDocument."
tags: [video]
---

{{< callout type="warning" >}}
La vidéo n'est dans aucune version pour l'instant.
{{< /callout >}}

Une vidéo est une image dont le fichier se termine par `.mp4`, `.m4v` ou `.webm` :

```markdown
![video](clip.mp4)
![The demo](clip.mp4#t=12.5)
![demo loop muted](clip.mp4)
```

## Syntaxe {#syntax}

Le fichier est lu dans le dossier de la présentation, ou depuis une URL `http` ou `https`. Un chemin absolu, un chemin qui sort du dossier ou tout autre schéma est refusé. `#t=12.5`, ou `#t=1:23` en minutes et secondes, choisit l'image d'affiche : l'image montrée avant la lecture et par les sorties qui ne peuvent pas lire la vidéo. La forme de fragment média `#t=10,20` est acceptée aussi, et son début est utilisé. Sans elle, l'affiche est la première image.

Des mots-clés dans le texte alternatif, écrits comme les mots-clés d'image de Marp mais propres à go-slides, règlent la lecture : `autoplay`, `loop` et `muted`. `autoplay` n'est pris en compte qu'avec `muted` ; sinon il est ignoré, avec un avertissement. Les autres mots du texte alternatif forment l'étiquette accessible.

## Ce que montre chaque sortie {#outputs}

- **HTML** (la page, et l'aperçu du bac à sable) : un élément `<video controls playsinline preload="metadata" poster=...>`. L'affiche est l'image décodée, intégrée sous forme d'URI de données JPEG. Une vidéo ne joue jamais le son toute seule. Quand le navigateur demande moins de mouvement (`prefers-reduced-motion`), `autoplay` et `loop` sont retirés et la vidéo est mise en pause. Une vidéo qui démarre seule (muette) ne joue que tant qu'au moins la moitié est visible, et se met en pause quand sa diapositive n'est pas affichée.
- **PDF** (`--pdf`) : l'image d'affiche comme une image, avec un lien vers le fichier vidéo : le chemin relatif tel qu'écrit dans la présentation, ou l'URL pour une source `http` ou `https`. La lecture n'a lieu qu'en ouvrant le lien. Le PDF ne porte pas d'annotation `RichMedia` : elle ne jouerait que dans Acrobat, et elle n'est délibérément pas utilisée.
- **PowerPoint** (`--pptx`) : la vidéo est intégrée comme une partie média (`ppt/media/slidesvideoN.mp4|webm`) avec son affiche, sous la forme `a:videoFile` plus l'extension `p14:media`. Elle se lit dans PowerPoint ; la lecture dans PowerPoint et Keynote n'a pas été vérifiée. Les mots-clés `autoplay`, `loop` et `muted` ne sont pas repris dans PowerPoint ni OpenDocument.
- **OpenDocument** (`--odp`) : la vidéo est intégrée sous `Video/` comme un `draw:plugin` (ODF 1.3 partie 3, 10.4.8), avant le `draw:image` de l'affiche, et listée dans le manifeste. Elle se lit dans Impress si son moteur multimédia gère le codec (non vérifié).
- **Sources URL** : la page HTML diffuse la vidéo depuis l'URL. Les exports montrent une carte neutre liée à l'URL. Rien n'est téléchargé à la construction, et rien n'est intégré dans PowerPoint ni OpenDocument.

## Conteneurs et affiches {#codecs}

Les fichiers MP4 (ISO BMFF) et WebM (Matroska) sont lus, en Go ; ffmpeg n'est jamais lancé. L'affiche n'est décodée que pour H.264 (`avc1` ou `avc3`), le seul décodeur vidéo de go-avkit. Les autres codecs (H.265, VP8, VP9, AV1) se lisent quand même en HTML et sont intégrés dans PowerPoint et OpenDocument, mais leur affiche est une carte sombre neutre avec un triangle de lecture, et l'avertissement `video.codec` est affiché.

## Limites {#limits}

Une vidéo est une entrée non fiable ; elle est donc bornée : le fichier (64 MiB), les boîtes ou éléments (262 144) et leur imbrication (16 niveaux), l'image (4096 par 4096) et le décodage de l'affiche (au plus 600 échantillons depuis l'image clé précédente, 8 MiB par échantillon). Une affiche trop loin d'une image clé reçoit la carte et l'avertissement `video.limit`.

Une vidéo refusée est remplacée dans la diapositive par une note, `[video not shown: ...]` ; la construction continue, et le problème est un avertissement portant un code :

- `video.syntax` : ce n'est pas un fichier MP4 ou WebM, pas de piste vidéo, ou un temps invalide ;
- `video.limit` : une borne ci-dessus est dépassée ;
- `video.codec` : l'affiche ne peut pas être décodée pour ce codec ;
- `video.path` : un chemin absolu, un chemin qui sort du dossier de la présentation, ou un schéma autre que `http` et `https` ;
- `file.missing` : le fichier n'existe pas.

## Dans le bac à sable {#playground}

Une fois publiée, une vidéo pourra être ajoutée au projet comme les autres fichiers. Ses octets restent dans le navigateur, sous forme d'URL blob, et ne sont jamais envoyés. Essayez-la dans le [bac à sable](https://go-slides.github.io/playground/).
