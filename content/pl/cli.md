---
title: "Wiersz poleceń"
weight: 50
description: "Każde polecenie i flaga slides oraz plik konfiguracyjny slides.yaml."
tags: [cli]
---

{{< callout type="warning" >}}
Ta strona opisuje v0.14.0.
{{< /callout >}}

```text
slides [deck.md] [flags]
slides [command] [flags] deck.md
```

## Polecenia {#commands}

| Polecenie | Co robi |
| --- | --- |
| `build` | zapisuje prezentację jako HTML, PDF lub obrazy (zobacz niżej). **Domyślne**: `slides deck.md` to `slides build deck.md` |
| `expand` | wypisuje prezentację jako zwykły Markdown Marp: szablony rozwinięte, kinds zastąpione dyrektywami, które reprezentują |
| `notes` | wypisuje notatki prelegenta |
| `kinds` | wylicza kinds, z których prezentacja może korzystać, i to, co ustawiają; `--json` (w v0.6.0) zapisuje je jako `[{"name", "extends", "directives": [{"key", "value"}]}]`, z rozwiązanym `extends` |
| `themes` | wylicza motywy |
| `version` | wypisuje wersję (tak samo `slides --version`) |
| `completion` | wypisuje skrypt uzupełniania powłoki: `bash`, `zsh`, `fish` lub `powershell` |

`slides <command> --help` wypisuje flagi polecenia. Flagi mogą stać przed
prezentacją lub po niej: `slides deck.md -o out.html` i `slides -o out.html
deck.md` to to samo. Długie flagi mają dwa myślniki (`--html`); `-o` jest jedyną
krótką. `build` zapisuje `deck.html` obok `deck.md`, chyba że `-o` mówi
inaczej.

## Flagi {#flags}

| Flaga | Polecenia | Co robi |
| --- | --- | --- |
| `-o`, `--output file` | `build`, `expand`, `notes` | plik wyjściowy; `-` dla standardowego wyjścia |
| `-I`, `--input-dir dir` | `build` | konwertuje pliki Markdown pod `dir`, rekurencyjnie (w v0.6.0); zobacz niżej |
| `-P`, `--parallel n` | `build` | liczba prezentacji budowanych naraz (w v0.6.0): domyślnie 5, co najwyżej liczba rdzeni; 1 buduje je jedna po drugiej |
| `-w`, `--watch` | `build` | buduje ponownie, gdy zmieni się prezentacja, motyw, plik kinds lub danych albo konfiguracja, i przeładowuje otwarte strony (w v0.6.0); zobacz niżej |
| `-s`, `--server dir` | `build` | serwuje katalog, konwertując jego prezentacje na żądanie (w v0.6.0); zobacz niżej |
| `-p`, `--preview` | `build` | otwiera wynik w domyślnej przeglądarce (w v0.6.0): strony HTML z `--watch` (domyślnie włączonym), serwer z `--server` |
| `--host addr` | `build` | adres, na którym nasłuchuje `--server`, domyślnie `127.0.0.1` (w v0.6.0) |
| `--allowed-host name` | `build` | publiczna nazwa hosta zwrotnego proxy, na którą odpowiada `--server`, można powtarzać (w v0.14.0); zobacz [Self-hosting]({{< relref "self-hosting" >}}) |
| `--template bespoke\|bare` | `build` | szablon strony HTML (w v0.6.0): `bespoke`, domyślny, ze skryptem prezentacji, kontrolerem na ekranie i przejściami, lub `bare`, slajdy ułożone jeden pod drugim (`--bare`) |
| `--bespoke.osc`, `--bespoke.progress`, `--bespoke.transition` | `build` | opcje szablonu bespoke, z wartościami domyślnymi marp-cli (w v0.6.0): kontroler włączony, **pasek postępu wyłączony** (wcześniej był zawsze rysowany), przejścia włączone; `--bespoke.osc=false` wyłącza jedną z nich |
| `-c`, `--config-file file`, `--no-config-file` | `build` | plik konfiguracyjny marp-cli, JSON lub YAML (w v0.6.0); zobacz niżej. `--config` to własny plik `slides` |
| `--browser`, `--browser-path`, `--browser-protocol`, `--browser-timeout` | `build` | akceptowane i ignorowane z ostrzeżeniem (w v0.6.0): nie ma przeglądarki |
| `--pptx-editable`, `--no-parallel` | `build` | zapisy marp-cli dla `--pptx` i `-P 1` (w v0.6.0) |
| `--config file` | wszystkie | plik konfiguracyjny; domyślnie `slides.yaml` obok prezentacji, jeśli istnieje |
| `--no-config` | wszystkie | nie czytaj pliku konfiguracyjnego |
| `--theme-set path` | wszystkie | plik CSS motywu lub katalog z nimi (można powtarzać) |
| `--kinds file` | wszystkie | plik YAML z kinds (można powtarzać) |
| `--data file` | wszystkie | plik YAML lub JSON z danymi szablonu (można powtarzać) |
| `--set key=value` | wszystkie | jedna wartość szablonu (można powtarzać) |
| `--no-template` | wszystkie | nie rozwijaj szablonów Go |
| `--strict-yaml` | wszystkie | parsuj dyrektywy jako ścisły YAML, jak zwykły Marpit (Marp Core parsuje je luźno) |
| `--no-emoji` | wszystkie | zostaw `:shortcodes:` i emoji Unicode jako tekst zamiast rysować je jako obrazy (od v0.4.0: [Emoji]({{< relref "emoji" >}})) |
| `--html` | `build` | dopuść cały surowy HTML; domyślnie lista dozwolonych z Marp Core |
| `--lang code` | `build` | `lang` strony, gdy prezentacja go nie ustawia |
| `--bare` | `build` | ułóż slajdy jeden pod drugim, bez skryptu prezentacji |
| `--pdf` | `build` | zapisz PDF, jedna strona na slajd, ze spisem treści — tekst i kształty jako wektory (wybierane także przez `-o file.pdf`) |
| `--pdf-raster` | `build` | zapisz PDF, w którym każda strona to jeden obraz slajdu, jak robiła v0.3.0 |
| `--notes` | `build` | zapisz notatki prelegenta jako plik tekstowy, slajdy rozdzielone `---` jak w marp-cli (w v0.6.0; wybierane także przez `-o file.txt`). Polecenie `notes` wypisuje je w innym układzie |
| `--pdf-notes` | `build` | dodaj notatki prelegenta każdego slajdu do jego strony PDF jako adnotację notatki, jak robi marp-cli (w v0.6.0). Prosi o PDF, gdy nic innego nie mówi, co zapisać |
| `--pdf-outlines`, `--pdf-outlines.pages`, `--pdf-outlines.headings` | `build` | zakładki PDF (w v0.6.0): „Page N" na slajd i zagnieżdżone pod nim nagłówki. Tu domyślnie włączone, inaczej niż w marp-cli; `=false` wyłącza jedną z nich. Zobacz niżej |
| `--images png\|jpeg` | `build` | zapisz jeden obraz na slajd: `deck.001.png`, … |
| `--image png\|jpeg` | `build` | zapisz pierwszy slajd jako obraz (wybierane także przez `-o file.png` lub `-o file.jpg`) |
| `--pptx` | `build` | zapisz plik PowerPoint z notatkami prelegenta (wybierane także przez `-o file.pptx`): edytowalny od v0.5.0, jeden obraz na slajd w v0.4.0 |
| `--pptx-raster` | `build` | zapisz plik PowerPoint z jednym obrazem całego slajdu na slajd, jak robi marp-cli (od v0.5.0); `--image-scale` domyślnie 2 |
| `--odp` | `build` | zapisz prezentację OpenDocument z notatkami prelegenta (wybierane także przez `-o file.odp`): edytowalna od v0.5.0 |
| `--odp-raster` | `build` | zapisz prezentację OpenDocument z jednym obrazem na slajd (od v0.5.0) |
| `--jpeg-quality n` | `build` | jakość JPEG, od 1 do 100 (domyślnie 85) |
| `--image-scale n` | `build` | piksele urządzenia na piksel CSS dla obrazów i PDF (domyślnie 1); 2 daje 2560×1440 dla 16:9, z tekstem i wektorami rysowanymi w tej gęstości, a nie powiększanymi |
| `--allow-local-files` | `build` | pozwól obrazom i PDF czytać pliki lokalne, tylko spod katalogu prezentacji; domyślnie odrzucane, jak w marp-cli |
| `--theme name\|file.css` | `build` | nadpisz motyw prezentacji (w v0.6.0): nazwa motywu lub plik CSS, który staje się motywem, jakikolwiek `@theme` deklaruje; katalog jest odrzucany, użyj `--theme-set` |
| `--title`, `--description`, `--author`, `--keywords`, `--url`, `--og-image` | `build` | metadane strony (w v0.6.0), nad dyrektywami prezentacji `title`, `description`, `author`, `keywords`, `url` i `image`; zobacz niżej |

## Wiele prezentacji i katalog {#several-decks-and-a-directory}

W v0.6.0 `build` przyjmuje kilka prezentacji, katalog (prezentacje pod nim,
znajdowane rekurencyjnie) lub wzorzec i buduje do `--parallel` z nich naraz.
`-` to standardowe wejście; inaczej niż w marp-cli, potok nie jest czytany, jeśli `-` tego nie mówi.

Wzorzec jest czytany tak, jak czyta glob marp-cli (`**` i `{a,b}` w v0.7.0;
ujmij go w cudzysłów, by powłoka zostawiła go `slides`): `*`, `?`, `[a-z]` i
`{md,txt}` w obrębie nazwy, `**` dla dowolnej liczby katalogów, tak że
`'docs/**/*.md'` bierze prezentacje całego drzewa. Wieloznacznik nie obejmuje nazwy
zaczynającej się od kropki, w `node_modules` niczego się nie szuka, wielkość liter ma znaczenie
(`*.md` nie znajdzie `UPPER.MD`), a wszystko, co wzorzec dopasuje, jest konwertowane,
jakiekolwiek ma rozszerzenie. Wzorzec, który niczego nie dopasuje, jest błędem.

- **`-I dir`** znajduje pliki Markdown (`.md`, `.mdown`, `.markdown`,
  `.markdn`, małymi literami) pod `dir`, jak robi marp-cli: w `node_modules` i
  katalogach z kropką się nie szuka. Z `-o out` każda prezentacja jest zapisywana w swojej
  ścieżce pod `dir`, odzwierciedlonej pod `out`: `dir/sub/b.md` daje
  `out/sub/b.html` (`b.pdf`, `b.001.png`, `b.txt`, … cokolwiek mówią flagi).
  Bez `-o` wyniki idą obok prezentacji. Wskazanie prezentacji oprócz
  `-I` jest błędem, tak samo `-o file` dla kilku prezentacji (`-o -` jest dozwolone).
- **`-P n`** buduje `n` prezentacji naraz. Renderowanie jednej prezentacji używa już
  rdzeni, więc `n` jest ograniczone ich liczbą; z `-o -` prezentacje są budowane jedna
  po drugiej. Komunikaty prezentacji pozostają razem, a prezentacja, która zawiedzie, nie
  zatrzymuje pozostałych: kod wyjścia to 1, a komunikat wymienia niepowodzenia.

## Przejścia i strona prezentacji {#transitions-and-the-presenting-page}

W v0.6.0 działa dyrektywa `transition` z marp-cli: `<!-- transition:
fade 1s -->`, nazwa i opcjonalny czas trwania, dla slajdu i następnych,
lub `<!-- _transition: cover -->` dla jednego. Wbudowane nazwy to nazwy marp-cli:
`clockwise`, `counterclockwise`, `cover`, `coverflow`, `cube`, `cylinder`,
`diamond`, `drop`, `explode`, `fade`, `fade-out`, `fall`, `flip`, `glow`,
`implode`, `in-out`, `iris-in`, `iris-out`, `melt`, `overlap`, `pivot`, `pull`,
`push`, `reveal`, `rotate`, `slide`, `star`, `swap`, `swipe`, `swoosh`, `wipe`,
`wiper` i `zoom`; `none` je wyłącza. Strona odtwarza je za pomocą **View Transition API**
przeglądarki, gdy widz opuszcza slajd, który je ma, do przodu,
i gdy do niego wraca; przeglądarki bez tego API zmieniają slajdy jak dawniej,
a widz, który prosi o ograniczenie ruchu, dostaje zwykłe płynne przejście. Motyw może wnieść
własne klatki kluczowe, nazwane `marp-transition-NAME` (albo `marp-outgoing-` i
`marp-incoming-`, oraz `-backward`), które są używane przed wbudowanymi.

Strona ma **kontroler na ekranie** — poprzedni, strona, następny, pełny ekran —
który pokazuje się, gdy porusza się wskaźnik lub klawisz, i znika po chwili
(`--bespoke.osc=false` go usuwa), oraz, z `--bespoke.progress`, pasek
postępu. **Pasek postępu jest teraz domyślnie wyłączony**, jak w marp-cli; wcześniej był zawsze
rysowany.

### Widok prelegenta {#presenter-view}

W v0.7.0 strona ma widok prelegenta z marp-cli. Naciśnij `p` lub ostatni
przycisk kontrolera, a otworzy się okno z bieżącym slajdem, krokiem
po nim, twoimi notatkami, stroną, godziną i stoperem (kliknij stoper,
by go wyzerować; `+` i `-` zmieniają rozmiar notatek; pasek między slajdami a
notatkami można przeciągać). Dwa okna podążają jedno za drugim w obie strony: strzałki,
kliknięcie następnego kroku albo Previous/Next w którymkolwiek z nich przesuwa
drugie, po jednym fragmencie.

Działa tak jak w marp-cli, odczytane z jego źródeł 4.5.1: okno to
`?view=presenter&sync=KEY` tego samego pliku, następny krok to ta sama strona co
`?view=next` w ramce, a okna współdzielą swoją pozycję przez
`localStorage` (`bespoke-marp-sync-KEY`), więc nie potrzeba serwera ani
sieci, i nie działa w przeglądarce bez `localStorage` (przycisk
o tym mówi). Notatki są tekstem: `&`, `<` i `>` są zamieniane na encje, więc notatka
nigdy nie może być znacznikiem. Są na stronie, ukryte, by widok prelegenta mógł
je pobrać, co trzyma je poza tym, co czytnik ekranu czyta odbiorcom.
Strony `--bare` nie mają tego wszystkiego. Widok dodaje do strony około 10 KB.

### Widok przeglądu {#overview-view}

W v0.8.0 naciśnij `o` (lub `Escape`: marp-cli przełącza obydwoma),
przycisk siatki kontrolera lub licznik stron widoku prelegenta, a każdy slajd
pokaże się w całości w siatce nad stroną. Strzałki wybierają slajd (góra i dół
o wiersz), `Home` i `End` przechodzą na końce, a pozostałe okna podążają;
`Enter`, `Space` lub kliknięcie wybiera slajd i, nad stroną, zamyka
przegląd. Nad widokiem prelegenta wybór zostawia go otwartym (prelegent wybiera
slajd do pokazania i dalej patrzy). Przycisk „Close overview", `o` i
`Escape` go zamykają. To rozwiązanie z marp-cli: strona ponownie w ramce pod
`?view=overview&sync=KEY`, współdzieląca pozycję przez `localStorage`, bez
sieci. Slajdy w siatce można fokusować i mają etykietę „Slide N”, a zamknięty
przegląd jest nieaktywny, więc użytkownicy klawiatury i czytników ekranu nigdy
w nim nie wylądują. Oba widoki razem dodają do strony około 15 KB.

### Pliki konfiguracyjne {#configuration-files}

marp-cli czyta swoje opcje z pliku w bieżącym katalogu, tak samo
`slides`, gdy plik jest danymi: `.marprc` (YAML, JSON to też YAML),
`.marprc.json`, `.marprc.yaml`, `.marprc.yml`, `.config/marprc*` lub klucz
`marp` w `package.json`; `--config-file` (`-c`) wskazuje dowolny plik JSON lub YAML,
`marp.config.json` jeden z nich. Przeszukiwany jest tylko bieżący katalog, nie jego
nadrzędne, jak w wyszukiwaniu marp-cli. Plik JavaScript lub TypeScript
(`marp.config.js` ...) jest programem, którego `slides` nie uruchamia: mówi o tym i
idzie dalej. Opcje to te z dokumentacji marp-cli, które coś tu znaczą:
`title`, `description`, `author`, `keywords`, `url`, `ogImage`, `lang`,
`theme`, `themeSet`, `html`, `allowLocalFiles`, `template`, `bespoke`,
`inputDir`, `output`, `parallel`, `pdf`, `pptx`, `notes`, `pdfNotes`,
`pdfOutlines`, `images`, `image`, `imageScale`, `jpegQuality`, `watch`,
`server`, `preview`; ścieżki są względem pliku. Wiersz poleceń wygrywa.
`engine`, `options` i `browser*` wymagają silnika Marp i przeglądarki i są
ignorowane z ostrzeżeniem, tak jak flagi `--browser*`: skrypty napisane dla
marp-cli nadal działają. (`slides.yaml` i `slides.hcl`, z kinds i
danymi szablonu, są osobne i czytane jak dawniej.)

## Watch, serve i podgląd {#watch-serve-and-preview}

W v0.6.0, z trasami i protokołem przeładowania na żywo z marp-cli (odczytanymi z jego
źródeł 4.5.1 i porównanymi z działającym `marp -s`).

**`slides --server dir`** (`-s`) serwuje katalog pod `http://127.0.0.1:8080/`
(zmienna środowiskowa `PORT` zmienia port; zajęty port jest
błędem, który o tym mówi):

| Żądanie | Odpowiedź |
| --- | --- |
| `/deck.md` | prezentacja jako HTML, konwertowana na żądanie |
| `/deck.md?pdf`, `?pptx`, `?png`, `?jpg` (lub `?jpeg`), `?notes` (lub `?txt`) | prezentacja skonwertowana tak, jak prosi zapytanie; PPTX jest załącznikiem |
| katalog | jego `index.md` lub `PITCHME.md` jako strona, w przeciwnym razie indeks z odnośnikiem do PDF i PPTX dla każdej prezentacji |
| każdy inny plik | ten plik |
| prezentacja, która się nie konwertuje | 503 z błędem |

Strona otwiera WebSocket do `/.__marp-cli-watch-notifier__/<id>`, gdzie id to
SHA-256 ścieżki prezentacji; serwer wysyła `ready`, potem `reload`, gdy
prezentacja lub cokolwiek pod katalogiem się zmieni. Katalog zażądany bez
końcowego ukośnika jest do niego przekierowywany (marp-cli serwuje tam swój indeks, co
psuje względne odnośniki strony).

**`--watch`** (`-w`) buduje prezentacje, potem buduje ponownie tę, która się zmienia, a
gdy zmieni się plik `--theme-set`, plik kinds lub danych albo konfiguracja,
wszystkie. Zapisywane strony otwierają gniazdo na pętli zwrotnej (port
37717 lub następny wolny) i przeładowują się, gdy ich prezentacja jest zbudowana ponownie. Prezentacja,
której budowa się nie powiedzie, jest zgłaszana i dalej obserwowana. **`--preview`** (`-p`) otwiera
domyślną przeglądarkę na serwerze lub na właśnie zbudowanych stronach HTML; nie ma
własnego okna, ponieważ `slides` nie ma przeglądarki do sterowania.

Serwer jest drogą do maszyny, więc jest surowszy niż ten z marp-cli, który
nasłuchuje na każdym interfejsie i niczego nie sprawdza:

- nasłuchuje na **127.0.0.1**; `--host` to zmienia, z ostrzeżeniem, a adres
  wieloznaczny nie może już sprawdzać nagłówka `Host`;
- nagłówek **`Host`** musi być nazwą pętli zwrotnej (`localhost`, `127.0.0.1`,
  `[::1]`, `*.localhost`) albo hostem, na którym nasłuchuje, co zatrzymuje DNS rebinding;
- **`Origin`** WebSocketu musi być własnym pochodzeniem serwera; gniazda trybu watch
  przyjmują też pochodzenie `null` strony otwartej z pliku; żądanie z
  innej witryny, które nie jest nawigacją (skrypt, obraz, `fetch`), jest
  odrzucane;
- ścieżka jest oczyszczana, jej dowiązania symboliczne rozwiązywane, a wynik musi leżeć
  wewnątrz katalogu: `..` i dowiązanie na zewnątrz są odrzucane, tak samo
  `?pdf` prezentacji spoza katalogu; nazwy zaczynające się od kropki (`.git`, `.env`) nie są
  serwowane;
- tylko `GET` i `HEAD`; kilka konwersji naraz; co najwyżej 4096 katalogów
  obserwowanych, żadnego w `node_modules` ani pod kropką.

Za zwrotnym proxy `--allowed-host name` (w v0.14.0, można powtarzać)
dodaje publiczną nazwę proxy do akceptowanych `Host` i `Origin`; kontrole
pozostają włączone, a serwer pozostaje na pętli zwrotnej.
[Self-hosting]({{< relref "self-hosting" >}}) stawia go za Caddy, z
HTTPS i logowaniem.

## Metadane strony {#page-metadata}

Wiersz poleceń wygrywa z prezentacją, jak w marp-cli, gdzie flagi są
dołączane do prezentacji jako ostatni komentarz-dyrektywa. Reguły odczytane z marp-cli
4.5.1 i przypięte testami względem jego wyniku (w v0.6.0):

- Prezentacja, która nie ustawia `title`, bierze tekst swojego pierwszego nagłówka o najniższym
  poziomie, tak jak jest zapisany (`# **Bold** title` daje `**Bold** title`). Bez
  nagłówka nie ma `<title>`.
- `keywords` są dzielone po przecinkach, przycinane i pozbawiane duplikatów; lista we front matter
  jest brana taka, jaka jest. Jak marp-cli, strona niesie znacznik `keywords` od
  drugiego słowa kluczowego, nie dla pojedynczego.
- Nieprawidłowy `url` (niebezwzględny) jest ignorowany z ostrzeżeniem, a poprzednia
  wartość zachowana; prawidłowy zapisuje też `<link rel="canonical">` i `og:url`.
- `author` i `description` idą ze swoimi bliźniakami Open Graph; `og:image:alt`
  i duża karta Twittera wymagają i tytułu, i obrazu.

## Obrazy, PDF, PPTX i ODP {#images-pdf-pptx-and-odp}

`-o` z rozszerzeniem `.pdf`, `.pptx`, `.odp`, `.png` lub `.jpg` wybiera format, jak robi
marp-cli. Katalogi wyjściowe są tworzone. PDF ma jedną stronę na slajd, w
rozmiarze slajdu, oraz spis slajdów i ich nagłówków.

### Notatki i spis treści PDF {#pdf-notes-and-outline}

W v0.6.0 PDF przyjmuje dwie opcje marp-cli, porównane z
PDF-ami, które marp-cli 4.5.1 zapisuje dla tej samej prezentacji:

- `--pdf-notes` umieszcza notatki prelegenta każdego slajdu na jego stronie jako adnotację
  notatki: żółta ikona Note w lewym dolnym rogu, której zawartością są
  notatki (kilka komentarzy połączonych pustą linią), a tytułem `author`
  prezentacji. Samo `--pdf-notes` prosi o PDF, jak w marp-cli.
- Spis treści zawiera pogrubioną zakładkę „Page N" na slajd, a pod nią
  nagłówki slajdu zagnieżdżone według poziomu (`--pdf-outlines.pages` i
  `--pdf-outlines.headings` wybierają te dwie części). **Spis jest tu domyślnie włączony,
  a w marp-cli wyłączony.** `slides` zawsze go zapisywał; kosztuje kilkaset
  bajtów i czyni długą prezentację nawigowalną, więc domyślna wartość zostaje.
  `--pdf-outlines=false` lub `--no-pdf-outlines` z marp-cli go wyłącza, a
  `=false` na `.pages` lub `.headings` odrzuca tę część. Zakładka nagłówka przeskakuje
  do swojego slajdu; marp-cli przewija też do nagłówka.

Renderowanie wykonuje [go-webengine](https://github.com/go-webengine/engine),
a PDF zapisuje [go-pdfkit](https://github.com/go-pdfkit/pdfkit);
czcionki, o które proszą wbudowane motywy (Lato, Roboto Mono), pochodzą z
[go-opentype/fonts](https://github.com/go-opentype/fonts) i są osadzone, więc
żadne renderowanie nie dotyka sieci. Nie ma Chrome.

`--image-scale 2` renderuje przy dwóch pikselach urządzenia na piksel CSS. Obrazy lokalne są
odrzucane, chyba że podano `--allow-local-files`, a nawet wtedy czytane są tylko pliki spod
katalogu prezentacji: prezentacja z innego miejsca nie sięgnie do `~/.ssh`.
Tła wielokrotne i dzielone (`![bg left:40%]`) renderują się tak, jak warstwy Marp.

`--pptx` zapisuje prezentację przez `WritePresentation` z
[go-ooxml](https://github.com/go-ooxml), z notatkami prelegenta i metadanymi
prezentacji. Od v0.5.0 slajdy są edytowalne (niżej); `--pptx-raster` zapisuje jeden
obraz całego slajdu na slajd, jak robi marp-cli, w skali domyślnie 2, którą
nadpisuje `--image-scale`. LibreOffice otwiera wynik. Testowano tylko LibreOffice;
PowerPoint i Keynote nie były testowane.

`--odp` zapisuje prezentację OpenDocument, przez `WritePresentation` z go-odf,
z jej notatkami prelegenta; `--odp-raster` zapisuje jeden obraz
na slajd.

### Edytowalne PPTX i ODP {#editable-pptx-and-odp}

Od v0.5.0 slajdy PPTX lub ODP są **edytowalne** — PPTX z marp-cli
to jeden obraz na slajd, a jego eksperymentalne `--pptx-editable` konwertuje
PDF przez LibreOffice. Tutaj każdy slajd jest rysowany ze swojego układu, tego samego,
który rysują eksporty PNG i PDF:

- każda linia tekstu to pole tekstowe w swoim wyliczonym położeniu, w czcionkach,
  w których slajd został złożony — **osadzonych w pliku**, więc wygląda tak
  samo na maszynie bez Lato czy Inter (sprawdzone w LibreOffice 25.8) —
  z pogrubieniem, kursywą, kolorami, odstępami między literami i klikalnymi odnośnikami;
- tła, obramowania i zaokrąglone ramki to kształty; gradienty to wypełnienia
  gradientowe w PPTX i obrazy w ODP (gradient ODF ma dwa
  kolory);
- obrazy są przechowywane raz, niezależnie od tego, ile slajdów je pokazuje;
- to, czego slajd nie umie wyrazić — element pod `opacity`, `filter`,
  rozmyty cień, wzór — jest obrazem tego elementu, z tekstem zachowanym
  nad nim.

Odrenderowany z powrotem przez LibreOffice, slajd mieści się w 0,04–2 % pikseli
eksportu PNG. `--pptx-raster` i `--odp-raster` zachowują jeden obraz na slajd.

### PDF wektorowy {#vector-pdf}

Od v0.4.0 PDF jest rysowany jako wektory: go-webengine maluje złożony
slajd do PDF zamiast do pikseli. Tekst jest tekstem, w czcionkach, w których
slajd został złożony, osadzonych jako podzbiór z mapą `/ToUnicode` — można go
zaznaczać, przeszukiwać i kopiować. Tła, obramowania i gradienty to ścieżki
i cieniowania, obraz pokazany na wielu slajdach jest osadzony raz, a odnośniki są
klikalne. 100-slajdowa prezentacja testowa ma 758 KiB wobec 837 KiB w
marp-cli i 10 MiB w PDF-ie z obrazów z v0.3.0, a zapisuje się w 0,3 s.

Element bez formy wektorowej tutaj — pod `opacity` < 1, z `filter`
lub obrotem, rozmyty `box-shadow`, gradient, którego przezroczystość zanika — jest
rysowany jako bitmapa samego tego elementu, z tekstem nałożonym na nią niewidocznie, by
pozostał zaznaczalny. Wzory są nadal bitmapami. `--pdf-raster` zapisuje jeden obraz
na stronę, jak robiła v0.3.0.

### Dostępny, otagowany PDF {#accessible-tagged-pdf}

Od v0.12.0 wektorowy PDF jest otagowany. Niesie swoją strukturę logiczną,
którą czytnik ekranu czyta po kolei, tak jak czyta stronę internetową:

| W prezentacji | W PDF |
| --- | --- |
| slajd | sekcja (`Sect`) |
| `#` … `######` | `H1` … `H6` |
| akapit, lista, tabela | `P`; `L` z etykietą i treścią każdego elementu; `Table`, komórki nagłówka z zakresem |
| odnośnik | `Link`, razem z klikalnym obszarem, jego tekst jako opis |
| obraz | `Figure` z tekstem alternatywnym; `![](…)` bez tekstu alternatywnego jest dekoracją |
| `$…$`, `$$…$$` | `Formula`, jej źródło TeX jako opis |
| kod | `Code`, bez numerów linii i krążków objaśnień |
| `<span lang="fr">` | span w tym języku |
| tła, obramowania, numery stron | artefakty, które czytnik pomija |
| notatki prelegenta (`--pdf-notes`) | adnotacje w strukturze |

Plik deklaruje PDF/UA-1, standard dostępnego PDF. Niesie
`title` prezentacji, który pokazuje przeglądarka, i jej `lang` (`--lang`, gdy prezentacja
go nie ustawia, w przeciwnym razie `en`).

Prezentacja sprawdzona w veraPDF przechodzi PDF/UA-1, chyba że jej własne nagłówki łamią
kolejność wymaganą przez standard. PDF/UA wymaga, by pierwszy nagłówek był `H1` i by
żaden poziom nie został pominięty (`#`, a potem `###`). Tylko autor prezentacji może to naprawić.
PDF z Chrome tej samej prezentacji, który zapisuje marp-cli, jest także otagowany, ale łamie
siedem reguł PDF/UA, nie ma sekcji na slajd ani wzorów.

Otagowanie 100-slajdowej prezentacji testowej kosztuje 4% rozmiaru (788 wobec 757 KiB)
i nie daje mierzalnej różnicy czasu. `--pdf-tags=false` zapisuje PDF bez tagów.

## `slides.yaml` {#slidesyaml}

Plik konfiguracyjny obok prezentacji dotyczy jej, więc katalog prezentacji
współdzieli jeden:

```yaml
kinds: kinds.yaml        # a mapping of kinds, or file(s) holding one
themes:                  # theme CSS files or directories
  - themes/
data:                    # template data
  company: Example Inc.
html: false              # allow all raw HTML
lang: en
```

Kinds: własne `kinds:` prezentacji nadpisują te z pliku i z wiersza poleceń.
Dane: `data` z pliku i z wiersza poleceń nadpisują front matter prezentacji,
więc jedną prezentację można zbudować z różnymi danymi.

## Planowane {#planned}

Jeszcze niezbudowane: wzory jako wektory w PDF. Gdy się pojawią, będą tu
poleceniem lub flagą.
