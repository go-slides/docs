---
title: "Animowany SVG"
weight: 40
description: "Obraz SVG z animacjami SMIL lub CSS: odtwarzany przez przeglądarkę w HTML, animowane kształty w PowerPoint i OpenDocument, animowany GIF tam, gdzie one nie potrafią tego wyrazić, stan początkowy w PDF i obrazach."
tags: [svg, animation]
---

Obraz SVG, który animuje się sam, zapisuje się jak każdy obraz:

```markdown
![](spinner.svg)
![w:400](chart.svg)
```

## Co jest czytane {#read}

Animacje zapisane w samym pliku SVG: SMIL (`animate`, `animateTransform`, `animateMotion`, `animateColor`, `set`, z `begin`, `dur`, `repeatCount`, `fill`, `calcMode`, `keyTimes` i `keySplines`) oraz CSS (`@keyframes` z właściwościami `animation`, `transform-origin` i `transform-box`). To, czego nie rozumie, na przykład zdarzenie takie jak `begin="click"`, jest pomijane i zgłaszane ostrzeżeniem: nic nie jest odrzucane po cichu.

## Co pokazuje każde wyjście {#outputs}

- **HTML**: przeglądarka sama odtwarza animacje.
- **PowerPoint** (`--pptx`) i **OpenDocument** (`--odp`): każdy animowany element jest osobnym kształtem, a jego animacje startują razem ze slajdem: ruch, obrót i skala, przezroczystość, zmiana między dwoma kolorami. Ruch elementu jest próbkowany w chwilach właściwych animacji, więc przyspieszenia, wiele klatek kluczowych i obrót wokół innego punktu są dokładne. Powtórzenia (także `indefinite`), opóźnienia i odwrócenia są przenoszone.
- **Klatki**: gdy animacji nie da się tak wyrazić (ścieżka, która się zniekształca, kreska, która się rysuje, pochylenie), cały obraz jest renderowany w chwilach właściwych animacji i osadzany jako animowany GIF, który odtwarzają PowerPoint i Impress. Jest to ograniczone: 12 klatek na sekundę, najwyżej 240 klatek, 16 milionów pikseli we wszystkich klatkach razem.
- **PDF** (`--pdf`) i obrazy (`--images`): nie potrafią się animować. Rysunek jest pokazany w stanie początkowym, bez JavaScriptu i bez `RichMedia`, a jedno ostrzeżenie wymienia animowane obrazy.

Każdy animowany obraz dostaje jedną linię ostrzeżenia z kodem `svg-animation`, która mówi, jak został przeniesiony: natywnie albo jako klatki i dlaczego.

## Obsługa według wyjścia {#support}

| Cecha | HTML | PowerPoint | OpenDocument | PDF i obrazy |
|---|---|---|---|---|
| Przesunięcie, ruch wzdłuż ścieżki | odtwarzane | natywnie | natywnie | stan początkowy |
| Obrót i skala wokół dowolnego punktu | odtwarzane | natywnie | natywnie | stan początkowy |
| Promień koła, rozmiar prostokąta (bez obrysu) | odtwarzane | natywnie | natywnie | stan początkowy |
| Przezroczystość elementu z jednego kształtu | odtwarzane | natywnie | natywnie | stan początkowy |
| Kolor wypełnienia lub obrysu, od jednego koloru do drugiego | odtwarzane | natywnie | natywnie | stan początkowy |
| Powtórzenia, opóźnienia, odwrócenia, zatrzymanie na końcu | odtwarzane | natywnie | natywnie | stan początkowy |
| Przyspieszenia, wiele klatek kluczowych | odtwarzane | natywnie (próbkowane) | natywnie (próbkowane) | stan początkowy |
| Stopnie, zmiany skokowe | odtwarzane | klatki | klatki | stan początkowy |
| Zniekształcanie ścieżki, przesunięcie kreskowania, atrybuty położenia i rozmiaru, gradienty | odtwarzane | klatki | klatki | stan początkowy |
| Pochylenie, odbicie, kilka zegarów w jednym elemencie, zagnieżdżone animowane elementy | odtwarzane | klatki | klatki | stan początkowy |
| Zdarzenia | odtwarzane | nie przenoszone (ostrzeżenie) | nie przenoszone (ostrzeżenie) | stan początkowy |

## Pliki lokalne {#local-files}

Obraz lokalny jest czytany tylko z `--allow-local-files`, z folderu prezentacji, tak jak w marp-cli. Bez tej opcji obraz jest pomijany w eksportach, a kompilacja mówi o tym przez `local-file-blocked`, podając plik, jego slajdy i opcję. Plik, który nie istnieje, jest zgłaszany jako `local-file-missing`.

## Ograniczenia {#limits}

SVG to dane niezaufane, więc jest ograniczony: 8 MiB SVG, 50 000 elementów, 2 000 animacji, 256 klatek kluczowych jednej animacji, 60 sekund osi czasu i 240 klatek. Rysunek przekraczający granicę jest odrzucany z ostrzeżeniem, a kompilacja trwa dalej.

## Co sprawdzono {#checked}

LibreOffice Impress (macOS) otwiera pliki PowerPoint i OpenDocument i zachowuje każdą animację przy ponownym zapisie, w obu formatach. Keynote (macOS) importuje kształty i animowany GIF; jego animacji nie sprawdzono. PowerPoint: jeszcze niezweryfikowany.
