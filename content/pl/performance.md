---
title: "Wydajność"
weight: 60
description: "Jak szybko go-slides buduje prezentację, przy jakim zużyciu pamięci, mierzone względem marp-cli na tej samej maszynie."
tags: [performance]
---

{{< callout type="info" >}}
Poniższe wyniki PPTX dotyczą v0.5.0, a wyniki PDF wektorowego v0.4.0.
Pierwsza tabela dotyczy v0.3.0,
poprzedniego wydania, które renderuje slajdy współbieżnie i parsuje swoje czcionki
raz ([go-slides/slides#11](https://github.com/go-slides/slides/pull/11)).
{{< /callout >}}

Wiersze PDF i PPTX ze 100 slajdami w trzech tabelach poniżej zmierzono przed v0.18.0. Zmierzono je ponownie w v0.26.1, po naprawie regresji śledzonej w [go-slides/slides#117](https://github.com/go-slides/slides/issues/117): zob. [Wiele rdzeni](#several-cores).

go-slides i [marp-cli](https://github.com/marp-team/marp-cli) 4.5.1 zostały
uruchomione jako podprocesy na tych samych prezentacjach, na Apple M4 Max (16 CPU) ze
średnim obciążeniem z minuty poniżej 8, w pięciu przeplatanych rundach po rozgrzewce.
Czas to mediana czasu ściennego; pamięć to szczyt całego drzewa procesów
(marp-cli steruje Chrome, którego procesy się liczą).

*Zmierzone w v0.3.0 na wszystkich rdzeniach (16 CPU).*

| Prezentacja | Format | go-slides | marp-cli | pamięć go-slides | pamięć marp-cli |
| --- | --- | ---: | ---: | ---: | ---: |
| 100 slajdów | HTML | 36 ms | 315 ms | 47 MiB | 140 MiB |
| 100 slajdów | PDF | 1,14 s | 1,88 s | 242 MiB | 1 509 MiB |
| 100 slajdów | PNG, każdy slajd | 0,97 s | 5,44 s | 225 MiB | 1 354 MiB |
| 100 slajdów | PPTX | 3,39 s | 8,03 s | 530 MiB | 1 373 MiB |
| gaia, 20 slajdów | PDF | 283 ms | 1,57 s | 154 MiB | 1 315 MiB |
| matematyka i kod, 20 slajdów | PNG, każdy slajd | 250 ms | 2,32 s | 168 MiB | 1 286 MiB |

## PDF wektorowy {#vector-pdf}

W v0.4.0 PDF jest rysowany jako wektory (zobacz
[Wiersz poleceń]({{< relref "cli#vector-pdf" >}})). Ta sama maszyna i metoda,
2026-10-19:

*Zmierzone w v0.4.0 na wszystkich rdzeniach (16 CPU), 2026-10-19.*

| Prezentacja | go-slides | marp-cli | v0.3.0 | rozmiar go-slides | rozmiar marp-cli | rozmiar v0.3.0 |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| 100 slajdów | 306 ms | 1 763 ms | 1 147 ms | 758 KiB | 837 KiB | 10,0 MiB |
| gaia, 20 slajdów | 93 ms | 1 546 ms | 285 ms | 32 KiB | 66 KiB | 582 KiB |
| matematyka i kod, 20 slajdów | 140 ms | 1 643 ms | 294 ms | 124 KiB | 335 KiB | 638 KiB |

Pamięć dla PDF ze 100 slajdów: 158 MiB wobec 1 512 MiB w marp-cli; czas CPU
0,9 s wobec 2,7 s.

Wyniki v0.4.0 dla PDF ze 100 slajdów zmierzono 2026-10-19 o 21:29,
przy średnim obciążeniu z minuty 7,7, w trzech rundach: go-slides 306 ms, 758
KiB i 158 MiB; marp-cli 1 838 ms, 837 KiB i 1 512 MiB.

## Edytowalny PPTX {#editable-pptx}

Od v0.5.0 PPTX składa się z pól tekstowych, kształtów i obrazów (zobacz
[Wiersz poleceń]({{< relref "cli#editable-pptx-and-odp" >}})). Zmierzono
2026-10-19 o 22:51, przy średnim obciążeniu z minuty 7,1, w trzech rundach:

*Zmierzone w v0.5.0 na wszystkich rdzeniach (16 CPU), 2026-10-19.*

| Prezentacja | go-slides | marp-cli | rozmiar go-slides | rozmiar marp-cli | pamięć go-slides | pamięć marp-cli |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| 100 slajdów | 348 ms | 7 795 ms | 760 KiB | 7,2 MiB | 204 MiB | 1 370 MiB |
| gaia, 20 slajdów | 112 ms | 4 608 ms | 893 KiB | 4,8 MiB | | |
| matematyka i kod, 20 slajdów | 207 ms | 2 780 ms | 842 KiB | 1,5 MiB | | |

Dla prezentacji ze 100 slajdów v0.4.0 potrzebowała 4,7 s, 4,8 MiB i 540 MiB. PPTX
z marp-cli zawiera obrazy. Większość rozmiaru edytowalnej prezentacji to osadzone czcionki.

## Wiele rdzeni {#several-cores}

Zmierzone w v0.26.1 (go-webengine/engine v0.48.0): czas zegarowy w ms przy 1, 2, 4 i 16 rdzeniach, wobec marp-cli. Średnie obciążenie z 1 minuty na początku wynosiło 5,7. Czas marp-cli nie zmienia się wraz z liczbą rdzeni, ponieważ jego `-P` tylko rozdziela osobne pliki między przeglądarki. `slidesbench -cores` tworzy tę tabelę. go-slides jest szybszy niż marp-cli w każdym wierszu, przy każdej liczbie rdzeni.

| Prezentacja, tryb | 1 rdzeń | 2 | 4 | 16 | marp-cli |
| --- | ---: | ---: | ---: | ---: | ---: |
| 100 slajdów, PNG | 2628 | 1397 | 899 | 766 | 5197 |
| gaia, 20 slajdów, PNG | 731 | 394 | 227 | 222 | 2331 |
| matematyka i kod, 20 slajdów, PNG | 600 | 333 | 199 | 192 | 2140 |
| gaia, 20 slajdów, PDF | 85 | 64 | 57 | 58 | 1411 |
| matematyka i kod, 20 slajdów, PDF | 278 | 183 | 139 | 135 | 1457 |
| gaia, 20 slajdów, PPTX | 99 | 77 | 69 | 70 | 3110 |
| matematyka i kod, 20 slajdów, PPTX | 395 | 268 | 189 | 179 | 2527 |
| 100 slajdów, PDF | 544 | 331 | 247 | 256 | 1534 |
| 100 slajdów, PPTX | 692 | 435 | 317 | 276 | 7248 |

**Naprawione w v0.26.1.** Od v0.18.0 do v0.26.0 PDF ze 100 slajdami był wolniejszy niż marp-cli, a PDF i PPTX ze 100 slajdami nie skalowały się z rdzeniami: silnik odbudowywał instancje czcionki zmiennej dla każdej strony. go-webengine/engine v0.48.0 współdzieli je między stronami ([go-webengine/engine#443](https://github.com/go-webengine/engine/issues/443)), a v0.26.1 go używa: PDF ze 100 slajdami zajmuje na jednym rdzeniu 544 ms zamiast 3166 ms, ze szczytem 169 MiB zamiast 2182 MiB, i ma 792,4 KiB zamiast 1005,5 KiB, bo czcionka jest osadzona 3 razy zamiast 102. Strony są identyczne co do piksela. Zgłoszenie to [go-slides/slides#117](https://github.com/go-slides/slides/issues/117) (zamknięte).

## Dlaczego jest szybki {#why-it-is-fast}

- **Brak przeglądarki do uruchomienia.** marp-cli uruchamia Chrome przy każdym eksporcie
  PDF, obrazów i PPTX; go-slides renderuje przez
  [go-webengine](https://github.com/go-webengine/engine) we własnym
  procesie.
- **Slajdy renderują się równolegle**, domyślnie po cztery naraz, a każdy obraz jest
  zapisywany, gdy tylko przyjdzie jego kolej: długa prezentacja nigdy nie jest trzymana w pamięci
  jako obrazy.
- **Czcionki są parsowane raz** na proces, a nie raz na slajd.

## Czego nie wygrywa {#what-it-does-not-win}

- **Rozmiar PDF, w v0.3.0.** Każdy slajd był w PDF obrazem: 10 MiB wobec
  837 KiB dla prezentacji ze 100 slajdów. PDF wektorowy powyżej jest mniejszy niż
  ten z marp-cli.
- **Czas CPU na dużych prezentacjach.** Cztery wątki robocze zużywają więcej CPU, niż
  pokazuje czas ścienny: w v0.3.0 3,2 s wobec 2,7 s w marp-cli dla PDF ze 100 slajdów
  (0,9 s dla PDF wektorowego).
- **Pamięć jest górnym ograniczeniem dla marp-cli**: procesy Chrome współdzielą strony,
  które suma liczy w każdym z nich. Jego największy pojedynczy proces ma 220–320 MiB.

Metoda, każda prezentacja i tryb oraz polecenia do powtórzenia wyników
są w
[BENCHMARKS.md](https://github.com/go-slides/slides/blob/main/BENCHMARKS.md).
