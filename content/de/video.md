---
title: "Video"
weight: 39
description: "Ein Video in einer Präsentation: vom Browser in der HTML-Seite abgespielt, im PDF ein Posterbild mit Link, in PowerPoint und OpenDocument eingebettet."
tags: [video]
---

{{< callout type="warning" >}}
Video ist noch in keinem Release enthalten.
{{< /callout >}}

Ein Video ist ein Bild, dessen Datei auf `.mp4`, `.m4v` oder `.webm` endet:

```markdown
![video](clip.mp4)
![The demo](clip.mp4#t=12.5)
![demo loop muted](clip.mp4)
```

## Syntax {#syntax}

Die Datei wird aus dem Ordner der Präsentation gelesen oder von einer `http`- oder `https`-URL. Ein absoluter Pfad, ein Pfad, der den Ordner verlässt, oder jedes andere Schema wird abgelehnt. `#t=12.5`, oder `#t=1:23` für Minuten und Sekunden, wählt das Posterbild: das Bild, das vor der Wiedergabe und von den Ausgaben gezeigt wird, die das Video nicht abspielen können. Die Media-Fragment-Form `#t=10,20` wird ebenfalls akzeptiert, und ihr Anfang wird verwendet. Ohne sie ist das Poster das erste Bild.

Schlüsselwörter im Alternativtext, geschrieben wie Marp-Bildschlüsselwörter, aber eigene von go-slides, steuern die Wiedergabe: `autoplay`, `loop` und `muted`. `autoplay` gilt nur zusammen mit `muted`; sonst wird es mit einer Warnung ignoriert. Die übrigen Wörter des Alternativtexts sind die barrierefreie Beschriftung.

## Was jede Ausgabe zeigt {#outputs}

- **HTML** (die Seite und die Vorschau des Playgrounds): ein `<video controls playsinline preload="metadata" poster=...>`-Element. Das Poster ist das dekodierte Bild, als JPEG-Data-URI eingebettet. Ein Video spielt nie von selbst mit Ton. Wenn der Browser weniger Bewegung wünscht (`prefers-reduced-motion`), werden `autoplay` und `loop` entfernt und das Video angehalten. Ein Video, das (stumm) automatisch startet, spielt nur, solange mindestens die Hälfte sichtbar ist, und pausiert, wenn seine Folie nicht angezeigt wird.
- **PDF** (`--pdf`): das Posterbild als Bild, mit einem Link zur Videodatei: der relative Pfad, wie er in der Präsentation steht, oder die URL bei einer `http`- oder `https`-Quelle. Abgespielt wird nur über das Öffnen des Links. Das PDF enthält keine `RichMedia`-Annotation: Sie würde nur in Acrobat abspielen und wird bewusst nicht verwendet.
- **PowerPoint** (`--pptx`): Das Video wird als Medienteil (`ppt/media/slidesvideoN.mp4|webm`) mit seinem Poster eingebettet, als `a:videoFile` plus die Erweiterung `p14:media`. Es spielt in PowerPoint; die Wiedergabe in PowerPoint und Keynote wurde nicht geprüft. Die Schlüsselwörter `autoplay`, `loop` und `muted` werden nicht in PowerPoint oder OpenDocument übernommen.
- **OpenDocument** (`--odp`): Das Video wird unter `Video/` als `draw:plugin` eingebettet (ODF 1.3 Teil 3, 10.4.8), vor dem `draw:image` des Posters, und im Manifest aufgeführt. Es spielt in Impress, wo dessen Medien-Backend den Codec unterstützt (nicht geprüft).
- **URL-Quellen**: Die HTML-Seite streamt das Video von der URL. Die Exporte zeigen eine neutrale Karte mit Link zur URL. Beim Bauen wird nichts geladen, und nichts wird in PowerPoint oder OpenDocument eingebettet.

## Container und Poster {#codecs}

MP4- (ISO BMFF) und WebM-Dateien (Matroska) werden in Go gelesen; ffmpeg wird nie gestartet. Das Poster wird nur für H.264 dekodiert (`avc1` oder `avc3`), den einzigen Videodekoder von go-avkit. Andere Codecs (H.265, VP8, VP9, AV1) spielen in HTML trotzdem und werden in PowerPoint und OpenDocument eingebettet, ihr Poster ist aber eine neutrale dunkle Karte mit Wiedergabedreieck, und die Warnung `video.codec` wird ausgegeben.

## Grenzen {#limits}

Ein Video ist nicht vertrauenswürdige Eingabe und daher begrenzt: die Datei (64 MiB), die Boxen oder Elemente (262.144) und ihre Verschachtelung (16 Ebenen), das Bild (4096 mal 4096) und die Poster-Dekodierung (höchstens 600 Samples ab dem vorangehenden Keyframe, 8 MiB pro Sample). Ein Poster, das zu weit von einem Keyframe entfernt ist, erhält die Karte und die Warnung `video.limit`.

Ein abgelehntes Video wird auf der Folie durch einen Hinweis ersetzt, `[video not shown: ...]`; der Build läuft weiter, und das Problem ist eine Warnung mit einem Code:

- `video.syntax`: keine MP4- oder WebM-Datei, keine Videospur oder eine ungültige Zeit;
- `video.limit`: eine der obigen Grenzen wird überschritten;
- `video.codec`: das Poster lässt sich für diesen Codec nicht dekodieren;
- `video.path`: ein absoluter Pfad, ein Pfad, der den Ordner der Präsentation verlässt, oder ein anderes Schema als `http` und `https`;
- `file.missing`: die Datei existiert nicht.

## Im Playground {#playground}

Nach der Veröffentlichung lässt sich ein Video wie die anderen Dateien zum Projekt hinzufügen. Seine Bytes bleiben im Browser, als Blob-URLs, und werden nie hochgeladen. Probieren Sie es im [Playground](https://go-slides.github.io/playground/).
