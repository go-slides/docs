---
title: "Nagrania terminala"
weight: 38
description: "Nagranie asciinema w prezentacji: odtwarzane w małym odtwarzaczu na stronie HTML, rysowane jako tekst wektorowy w PDF, PowerPoint i OpenDocument."
tags: [asciinema, terminal]
---

{{< callout type="info" >}}
Nagrania terminala są w v0.18.0 i nowszych.
{{< /callout >}}

Nagranie [asciinema](https://docs.asciinema.org/) ([asciicast v2](https://docs.asciinema.org/manual/asciicast/v2/) lub [v3](https://docs.asciinema.org/manual/asciicast/v3/)) to obraz, którego plik kończy się na `.cast`:

```markdown
![cast](demo.cast)
![The failing build](demo.cast#t=12.5)
```

## Składnia {#syntax}

Plik jest czytany z folderu prezentacji: adres URL, ścieżka bezwzględna lub ścieżka wychodząca poza folder jest odrzucana. Tekst alternatywny nazywa nagranie dla czytników ekranu; samo `cast` oznacza własny tytuł nagrania. `#t=12.5`, albo `#t=1:23` w minutach i sekundach, wybiera chwilę, którą pokazują wyjścia nieruchome; bez tego jest to koniec nagrania.

## Co pokazuje każde wyjście {#outputs}

- **HTML** (strona i podgląd w playground): nagranie jest odtwarzane w małym odtwarzaczu go-slides osadzonym na stronie: odtwarzanie i pauza, przycisk prędkości i pasek położenia. Nie ładuje skryptu znikąd i nie korzysta z sieci.
- **PDF, PowerPoint i OpenDocument** (`--pdf`, `--pptx`, `--odp`): terminal w chwili wskazanej przez `#t=`, jako tekst wektorowy w czcionce o stałej szerokości z motywu, z kolorami 16, 256 i bezpośrednimi, pogrubieniem, kursywą, podkreśleniem i tekstem odwróconym. Tekst można zaznaczać i przeszukiwać. Nigdy nie jest to obraz. Pogrubienie i kursywa są rysowane tam, gdzie czcionka o stałej szerokości motywu je ma (Roboto Mono w motywie gaia).

## Odtwarzacz {#player}

Odtwarza raz, gdy nagranie pojawi się na ekranie po raz pierwszy. Gdy przeglądarka prosi o mniej ruchu (`prefers-reduced-motion`), startuje wstrzymany, pokazując chwilę wskazaną przez `#t=`, i odtwarza po naciśnięciu przycisku odtwarzania. Jego przyciski i klawisze nie zmieniają slajdu.

## Limity {#limits}

Nagranie jest niezaufanym wejściem, więc jest ograniczone: plik (16 MiB), zdarzenia (500 000), długość (3 godziny), terminal (400 na 200 komórek) i praca, jaką może wykonać emulator. Nagranie przekraczające limit lub niebędące poprawnym plikiem asciicast jest odrzucane: kompilacja trwa dalej, slajd pokazuje w jego miejscu notatkę, a problem jest ostrzeżeniem z kodem: `cast.syntax`, `cast.limit`, `cast.version` lub `cast.path`.

## W playground {#playground}

Dodaj plik `.cast` do projektu z panelu Pliki, przeciągając go lub z okna wyboru plików. Trafia do `media/` i jest wstawiany jako obraz. To tekst, który czyta prezentacja, więc liczy się razem z innymi plikami tekstowymi projektu, ograniczonymi do 2 MiB. Wypróbuj w [playground](https://go-slides.github.io/playground/).
