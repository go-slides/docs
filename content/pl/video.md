---
title: "Wideo"
weight: 39
description: "Wideo w prezentacji: odtwarzane przez przeglądarkę na stronie HTML, kadr plakatu z odnośnikiem w PDF, osadzone w PowerPoint i OpenDocument."
tags: [video]
---

Wideo to obraz, którego plik kończy się na `.mp4`, `.m4v` lub `.webm`:

```markdown
![video](clip.mp4)
![The demo](clip.mp4#t=12.5)
![demo loop muted](clip.mp4)
```

## Składnia {#syntax}

Plik jest czytany z folderu prezentacji albo z adresu URL `http` lub `https`. Ścieżka bezwzględna, ścieżka wychodząca poza folder lub każdy inny schemat jest odrzucany. `#t=12.5`, albo `#t=1:23` w minutach i sekundach, wybiera kadr plakatu: obraz pokazywany przed odtwarzaniem i przez wyjścia, które nie potrafią odtworzyć wideo. Akceptowana jest też forma fragmentu multimediów `#t=10,20`, a używany jest jej początek. Bez tego plakatem jest pierwszy kadr.

Słowa kluczowe w tekście alternatywnym, zapisywane jak słowa kluczowe obrazów w Marp, ale własne dla go-slides, określają sposób odtwarzania: `autoplay`, `loop` i `muted`. `autoplay` działa tylko razem z `muted`; w przeciwnym razie jest ignorowane, z ostrzeżeniem. Pozostałe słowa tekstu alternatywnego są etykietą dostępności.

## Co pokazuje każde wyjście {#outputs}

- **HTML** (strona i podgląd w playground): element `<video controls playsinline preload="metadata" poster=...>`. Plakat to zdekodowany kadr, osadzony jako URI danych JPEG. Wideo nigdy nie odtwarza dźwięku samo z siebie. Gdy przeglądarka prosi o mniej ruchu (`prefers-reduced-motion`), `autoplay` i `loop` są usuwane, a wideo jest wstrzymane. Wideo uruchamiane automatycznie (bez dźwięku) gra tylko wtedy, gdy widoczna jest co najmniej połowa, i zatrzymuje się, gdy jego slajd nie jest pokazywany.
- **PDF** (`--pdf`): kadr plakatu jako obraz, z odnośnikiem do pliku wideo: ścieżką względną zapisaną w prezentacji albo adresem URL dla źródła `http` lub `https`. Odtwarzanie następuje tylko po otwarciu odnośnika. PDF nie zawiera adnotacji `RichMedia`: odtwarzałaby się tylko w Acrobacie i celowo nie jest używana.
- **PowerPoint** (`--pptx`): wideo jest osadzone jako część multimedialna (`ppt/media/slidesvideoN.mp4|webm`) wraz z plakatem, jako `a:videoFile` plus rozszerzenie `p14:media`. Keynote: importuje je jako film i odtwarza (sprawdzone). PowerPoint: jeszcze niesprawdzone. Słowa kluczowe `autoplay`, `loop` i `muted` nie są przenoszone do PowerPoint ani OpenDocument.
- **OpenDocument** (`--odp`): wideo jest osadzone w `Video/` jako `draw:plugin` (ODF 1.3 część 3, 10.4.8), przed `draw:image` plakatu, i wymienione w manifeście. Odtwarza się w Impress, gdy jego zaplecze multimedialne obsługuje kodek (niesprawdzone).
- **Źródła URL**: strona HTML strumieniuje wideo z adresu URL. Eksporty pokazują neutralną kartę z odnośnikiem do URL. Nic nie jest pobierane podczas budowania i nic nie jest osadzane w PowerPoint ani OpenDocument.

## Kontenery i plakaty {#codecs}

Pliki MP4 (ISO BMFF) i WebM (Matroska) są czytane w Go; ffmpeg nigdy nie jest uruchamiany. Plakat jest dekodowany tylko dla H.264 (`avc1` lub `avc3`), jedynego dekodera wideo go-avkit. Inne kodeki (H.265, VP8, VP9, AV1) nadal odtwarzają się w HTML i są osadzane w PowerPoint i OpenDocument, ale ich plakatem jest neutralna ciemna karta z trójkątem odtwarzania, a wypisywane jest ostrzeżenie `video.codec`.

## Limity {#limits}

Wideo jest niezaufanym wejściem, więc ma granice: plik (64 MiB), pudełka lub elementy (262 144) i ich zagnieżdżenie (16 poziomów), kadr (4096 na 4096) oraz dekodowanie plakatu (najwyżej 600 próbek od poprzedniej klatki kluczowej, 8 MiB na próbkę). Plakat zbyt daleko od klatki kluczowej dostaje kartę i ostrzeżenie `video.limit`.

Odrzucone wideo jest na slajdzie zastępowane notatką, `[video not shown: ...]`; budowanie trwa dalej, a problem jest ostrzeżeniem z kodem:

- `video.syntax`: to nie plik MP4 ani WebM, brak ścieżki wideo albo nieprawidłowy czas;
- `video.limit`: przekroczono jedną z powyższych granic;
- `video.codec`: plakatu nie można zdekodować dla tego kodeka;
- `video.path`: ścieżka bezwzględna, ścieżka wychodząca poza folder prezentacji albo schemat inny niż `http` i `https`;
- `file.missing`: plik nie istnieje.

## W playground {#playground}

Po wydaniu wideo będzie można dodać do projektu jak inne pliki. Jego bajty zostają w przeglądarce, jako adresy blob, i nigdy nie są wysyłane. Wypróbuj w [playground](https://go-slides.github.io/playground/).
