---
title: "VS Code"
weight: 55
description: "Rozszerzenie go-slides dla VS Code: podgląd, prezentowanie, problemy, uzupełnianie i eksport, tak jak Marp for VS Code robi to dla Marp."
tags: [vscode, editor]
---

{{< callout type="info" >}}
Rozszerzenie to [go-slides/vscode](https://github.com/go-slides/vscode),
wersjonowane osobno (v0.4.0). Renderuje przez go-slides v0.9.0 (pakiet
`editor`, linie źródłowe synchronizacji przewijania oraz widok przeglądu i
prelegenta). Instaluje się je z `go-slides-0.4.0.vsix` dołączonego do
wydania na GitHubie: jego publikacja w VS Code Marketplace i Open VSX jest
przygotowana, ale jeszcze jej tam nie ma.
{{< /callout >}}

Rozszerzenie to odpowiednik
[Marp for VS Code](https://github.com/marp-team/marp-vscode) w go-slides: prezentacja pokazuje się w
podglądzie Markdown VS Code jako slajdy, jej błędy w panelu Problems, a
dyrektywy, rodzaje i funkcje szablonów uzupełniają się w trakcie pisania.

## Instalacja {#install}

Pobierz `go-slides-<version>.vsix` z
[wydań](https://github.com/go-slides/vscode/releases), potem uruchom
*Extensions: Install from VSIX...* albo:

```sh
code --install-extension go-slides-0.4.0.vsix
```

Podgląd nie potrzebuje niczego więcej. Eksport potrzebuje polecenia `slides`:

```sh
go install github.com/go-slides/slides/cmd/slides@latest
```

## Podgląd {#preview}

Otwórz prezentację i *Markdown: Open Preview to the Side*. Prezentacja to plik Markdown,
którego front matter mówi `slides: true`:

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

Pliki z `marp: true` też są podglądane, chyba że zainstalowano Marp for VS Code:
oba rozszerzenia renderowałyby ten sam dokument. Decyduje o tym ustawienie
`goSlides.preview.marpDocuments`: `auto` (domyślne), `always` lub
`never`.

Podgląd to sam go-slides, skompilowany do WebAssembly i ładowany przez
rozszerzenie: bez pliku binarnego, działa na [vscode.dev](https://vscode.dev). Czyta
to, czego używa prezentacja, z przestrzeni roboczej — `slides.yaml` lub `slides.hcl` obok
prezentacji, pliki [rodzajów]({{< relref "kinds" >}}), które wskazują, pliki dołączane przez
[szablon]({{< relref "templates" >}}), CSS motywów — i tylko z
folderu przestrzeni roboczej prezentacji.

## Synchronizacja przewijania {#scroll-sync}

Od v0.2.0 rozszerzenia (go-slides v0.7.0) edytor i
podgląd przewijają się razem, w obie strony, jak dla zwykłego Markdownu:
podgląd podąża za slajdem i blokiem pod kursorem edytora, a
przewijanie podglądu przesuwa edytor. Ustawienia to własne ustawienia VS Code,
`markdown.preview.scrollPreviewWithEditor` i
`markdown.preview.scrollEditorWithPreview`.

Linie są śledzone przez szablony. Slajd generowany przez `{{ range }}`
podąża za linią swojego szablonu, a to, co wnosi `{{ include }}`,
podąża za linią include. Prezentacja podzielona przez `headingDivider` synchronizuje się po
nagłówkach.

## Prezentowanie {#presenting}

Od v0.4.0 rozszerzenia (go-slides v0.9.0) prezentację można prezentować
bez opuszczania VS Code. Mając prezentację w edytorze, z palety poleceń
lub z menu `…` edytora:

| Polecenie | Otwiera |
|---|---|
| *go-slides: Open Overview* | panel ze wszystkimi slajdami w siatce. Wybrany slajd pokazuje swoje źródło w edytorze; `Esc` zamyka panel |
| *go-slides: Start Presenter View* | slajdy, a obok nich widok prelegenta: następny krok, notatki prelegenta i stoper |

Panele stoją obok edytora i podążają jeden za drugim: zmień slajd w jednym,
a pozostałe zmienią się razem z nim. Podążają też za twoimi zmianami w prezentacji. Klawisz
`p` na slajdach otwiera widok prelegenta.

Aby prezentować na innym ekranie, przenieś tam panel slajdów poleceniem *View: Move
Editor into New Window*.

Każdy panel uruchamia prezentację w ramce z piaskownicą (bez `allow-same-origin`) pod
polityką `default-src 'none'`: ramka ładuje obrazy, media, style i czcionki
i nic więcej, a uruchamia tylko skrypty, które go-slides wbudował w stronę,
wskazane przez skrót. Skrypt napisany w prezentacji się nie uruchamia.

## Problemy, uzupełnianie, podpowiedzi {#problems-completion-hover}

- **Problems**: błędy szablonów, YAML i [HCL]({{< relref "hcl" >}}), w
  linii, w której są, w prezentacji lub w pliku rodzajów, z którego pochodzą.
- **Uzupełnianie**: klucze dyrektyw we front matter i w komentarzach `<!-- -->`
  (z `_` dla dyrektywy punktowej), wartości `kind` i `_kind` (z tym, co
  ustawia każdy rodzaj), `theme`, `size`, `paginate`, `math`, `headingDivider`; wewnątrz
  `{{ }}` funkcje szablonów i pola danych.
- **Hover**: co robi dyrektywa, co ustawia rodzaj, argumenty funkcji szablonu.

## Eksport {#export}

*go-slides: Export Slide Deck...* (także na pasku tytułu edytora) zapisuje HTML,
wektorowy PDF, edytowalne PPTX i ODP, PNG lub JPEG poleceniem `slides`, tak jak
zrobiłoby `slides -o deck.pdf deck.md`; zobacz [Wiersz poleceń]({{< relref "cli" >}}).
Wymaga zaufanej przestrzeni roboczej. Jeśli `slides` nie zostanie znalezione, rozszerzenie proponuje
uruchomienie powyższego `go install` w terminalu.

## Ustawienia {#settings}

| Ustawienie | Domyślnie | |
|---|---|---|
| `goSlides.cliPath` | `""` | polecenie `slides`; puste: `slides` w `PATH`. Ustawienie maszyny: przestrzeń robocza nie może go ustawić |
| `goSlides.exportType` | `pdf` | format, który okno eksportu proponuje jako pierwszy |
| `goSlides.exportAutoOpen` | `true` | otwórz wyeksportowany plik |
| `goSlides.html` | `false` | dopuść cały surowy HTML zamiast listy dozwolonych Marp; uwzględniane tylko w zaufanej przestrzeni roboczej |
| `goSlides.preview.marpDocuments` | `auto` | podglądaj prezentacje `marp: true`: `auto`, `always`, `never` |

## Bezpieczeństwo {#security}

Prezentacja i przestrzeń robocza są niezaufane. Surowy HTML podlega liście dozwolonych Marp,
chyba że `goSlides.html` jest włączone w zaufanej przestrzeni roboczej (`slides.yaml` z
`html: true` nie włącza go w podglądzie), a podgląd VS Code nie uruchamia żadnego
skryptu z prezentacji. Podgląd czyta tylko pliki z folderu przestrzeni roboczej prezentacji,
nie podąża za żadnym dowiązaniem symbolicznym, a szablony nie widzą żadnej zmiennej środowiskowej.
Eksport uruchamia program, więc czeka na zaufanie do przestrzeni roboczej, a program
nigdy nie jest nazywany przez przestrzeń roboczą.

## Wydajność {#performance}

Na 100-slajdowej prezentacji testowej go-slides, w VS Code 1.142 na Apple M4 Max,
z Marp for VS Code 3.6.1 mierzonym w tej samej sesji:

| | go-slides | Marp for VS Code |
|---|---:|---:|
| aktywacja (ładowanie modułu WebAssembly) | 189 ms | — |
| pierwszy podgląd | 52 ms | 78 ms |
| podgląd po edycji (mediana) | **19 ms** | 23–24 ms |
| `.vsix` | 7,5 MB | 36,3 MB |

Wyniki dotyczą rozszerzenia v0.3.0 (go-slides v0.8.0) na prezentacji ze 100 slajdów,
przy wpisaniu jednego znaku w slajdzie 50, z przeplatanych rund przy średnim obciążeniu z minuty
poniżej 8. Pierwszy podgląd jest szybszy niż w Marp, a podgląd po
edycji jest teraz szybszy niż w Marp.

### W v0.8.0 i rozszerzeniu v0.3.0 {#in-v080-and-extension-v030}

go-slides v0.8.0 zachowuje wyrenderowane slajdy każdej otwartej prezentacji i po edycji
renderuje ponownie tylko te, których źródło lub odziedziczony stan się zmieniły (dyrektywy, numer
strony, suma, id nagłówków, odwołania do odnośników, linie źródłowe); wynik jest
taki sam jak przy pełnym renderowaniu, bajt w bajt. Pisanie w slajdzie 50 ze 100 renderuje ponownie jeden
slajd: 2,5 ms natywnie, 11,5 ms w WebAssembly. Podgląd VS Code przekształca wtedy
tylko DOM zmienionego slajdu, a test dymny rozszerzenia wpisuje jeden
znak w slajdzie 50 między mierzonymi renderowaniami, dla obu rozszerzeń. Widok
prelegenta rozszerzenia synchronizuje się przez okno nadrzędne (`postMessage`).

### W v0.7.1 i rozszerzeniu v0.2.2 {#in-v071-and-extension-v022}

go-slides v0.7.1 buforuje sparsowane motywy, spakowane arkusze stylów i podświetlony kod
między renderowaniami, oznacza linie synchronizacji przewijania bez przechodzenia bajt po bajcie i ponownie używa
swoich zapisywaczy: renderowanie prezentacji ze 100 slajdów alokuje 5,7 MB wobec 52,7 MB
dla rozszerzenia v0.2.0. Dzięki temu rozszerzenie v0.2.2 porzuca nadpisanie `GOGC=400`.
Dla porównania rozszerzenie v0.2.1 potrzebowało 72 ms na pierwszy podgląd i 44 ms
po edycji przy domyślnym `GOGC` (36 ms z `GOGC=400`), a v0.2.0 potrzebowało
47–51 ms.

Rozszerzenie v0.2.2 zmierzyło 52 ms dla pierwszego podglądu i 24 ms po edycji,
na równi z 23 ms w Marp.
