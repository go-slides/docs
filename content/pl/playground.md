---
title: "Playground"
weight: 15
description: "Napisz prezentację w przeglądarce i patrz, jak slajdy za tobą podążają: go-slides skompilowany do WebAssembly, z uzupełnianiem, pomocą i przykładami."
tags: [playground, browser]
---

**[Otwórz Playground →](https://go-slides.github.io/pl/playground/)**

Playground to sam go-slides, v0.15.0, skompilowany do WebAssembly i
działający w twojej przeglądarce. Pisz prezentację po lewej; slajdy podążają
po prawej. Nic nie jest wysyłane na serwer: ani prezentacja, ani odnośnik do udostępnienia.

Interfejs w Playground (Wstążka, pomoc, komunikaty) mówi językiem strony:
otwórz `/pl/playground/`, aby uzyskać polski, i tak samo dla dwunastu języków
tej witryny. Selektor języka na stronie go zmienia, a Twój wybór jest
zapamiętywany w tej przeglądarce; link wskazujący język zawsze ma pierwszeństwo.
Przykładowe prezentacje pozostają po angielsku. Ukraiński, rosyjski, chiński,
japoński i koreański są rysowane przy użyciu podzbioru Noto Sans (od 0,06 do
2,3 MiB), który strona pobiera tylko dla tych języków.

## Co ma {#what-it-has}

- **Podgląd na żywo** każdego slajdu, renderowany poza wątkiem strony (Web
  Worker), więc pisanie nigdy na niego nie czeka. Kliknij slajd, by przenieść
  edytor do jego linii; przesuń kursor, a podgląd przewinie się do jego slajdu.
- **Pokaz slajdów**: sama prezentacja, z klawiaturą,
  kontrolerem na ekranie i przejściami.
- **Widok prelegenta**: **Present** (lub `p` w pokazie slajdów) otwiera go w
  drugim oknie: bieżący slajd, następny krok, notatki prelegenta (komentarze
  `<!-- comments -->` prezentacji), zegar i stoper. Strzałki, Spacja,
  Home i End działają w obu oknach, a każde podąża za drugim, także
  gdy edytujesz prezentację. Jeśli okno się nie otwiera, zezwól na wyskakujące okna
  dla witryny. Widok przeglądu (`o`) jest tylko na pobranej stronie.
- **Uzupełnianie** w trakcie pisania: klucze dyrektyw (punktowe z `_` w komentarzach),
  wartości `kind`, `theme`, `size`, `paginate` oraz funkcje szablonów
  i pola danych wewnątrz `{{ }}`.
- **Pomoc** dla tego, co jest pod kursorem: dyrektywa, kind (co ustawia),
  funkcja szablonu. Panel pomocy odsyła też do tych stron.
- **Problemy** w ich linii: błędy szablonów, YAML i HCL są podświetlane
  w edytorze i wypisane pod nim, a podgląd zachowuje ostatnią prezentację,
  która się wyrenderowała.
- **Przykłady**: prezentacja powitalna, szablony generujące slajdy, kinds,
  front matter w HCL, motywy i dyrektywy punktowe oraz matematyka, obrazy i
  tła (`![bg]`, dzielone `![bg left:40%]`, filtry).
- **Pliki**: panel po lewej od edytora to projekt prezentacji, ułożony
  tak, jak czytają go [polecenie]({{< relref "cli" >}}) i
  [rozszerzenie VS Code]({{< relref "vscode" >}}): `deck.md`,
  `slides.yaml`, `themes/`, `kinds/`, `media/`.
  - Kliknij plik, by go edytować. Edytor koloruje go według języka
    (Markdown, YAML, HCL, JSON, CSS) i zaznacza jego problemy.
  - Kliknij prawym przyciskiem (lub Shift+F10), by uzyskać New File, New Folder, Rename (F2),
    Delete (pyta wcześniej) i New theme / New kind. Przeciągnij plik na folder,
    by go przenieść. Klawiatura podąża za wzorcem drzewa WAI-ARIA, łącznie
    z wyszukiwaniem po wpisywaniu.
  - **Project ▸ New theme** tworzy `themes/my-gaia.css` (lub z default,
    uncover): `@import` wbudowanego motywu i reguły, które dodasz. Trafia
    na listę motywów. **New kind** i **New data file** także startują
    z szablonów. Każdy jest wpisany do `slides.yaml`, więc `slides build deck.md`
    go znajduje.
  - **Project ▸ Download project (.zip)** zapisuje wszystko. Po rozpakowaniu
    `slides build deck.md` go zbuduje. **Open project** przyjmuje `.zip` lub
    folder z powrotem. Archiwum, które wychodzi poza swój folder (`..`, ścieżki bezwzględne),
    zawiera więcej niż 256 plików albo 8 MiB w pliku, lub rozpakowuje się ponad 100
    razy, jest odrzucane.
  - Wszystko jest przechowywane w tej przeglądarce razem ze szkicem. Projekt mieści
    do 64 MiB, a linia pod plikami mówi, ile jest zajęte.
- **Media**: upuść lub wklej obraz, wideo albo dźwięk na edytor, a trafi do
  `media/`, a prezentacja go nazwie: `![](media/photo.png)`. Upuść
  pliki na folder panelu, by umieścić je tam. Prezentacja pozostaje mała i
  czytelna, a nic nie jest wysyłane. Piaskownica podglądu ładuje je
  bez otwierania się (strona przekazuje jej bajty). **Save as ▸ HTML
  page** umieszcza je w stronie. Obrazy `https:` też działają, ale przeglądarka
  pobiera je z ich serwera, który widzi wtedy twój adres. Podgląd
  nie wysyła referera.
- **Panele ze zmiennym rozmiarem**: panel plików, edytor, miniatury slajdów,
  podgląd i pomoc są rozdzielone paskami, które można przeciągać (wskaźnik
  zmienia się na kursor zmiany rozmiaru). Panel obok paska zajmuje miejsce,
  pozostałe zachowują rozmiar. Kliknij pasek dwukrotnie lub naciśnij na nim Enter,
  by zwinąć panel i przywrócić go. Z klawiatury F6 przechodzi między
  paskami, strzałki przesuwają jeden o 16 px, Home i End doprowadzają panel
  przed nim do najmniejszego i największego rozmiaru, a Escape wraca do edytora.
  Karta View nadal wyłącza panele. Ustawione rozmiary są zapamiętywane w tej
  przeglądarce razem ze szkicem; rozmiary domyślne (190, 168 i 280 px, edytor
  i podgląd dzielą resztę w stosunku 6 do 5) nigdy nie są zapisywane.
- **Format**: karta Format edytuje Markdown za ciebie, a każde naciśnięcie to
  jeden krok cofania.
  - **Tekst**: Bold, Italic, Strikethrough i Inline code działają na zaznaczeniu
    lub słowie pod kursorem i się przełączają (Ctrl+B, Ctrl+I, Ctrl+Shift+X i
    Ctrl+E; Cmd na Macu). Przycisk jest wciśnięty, gdy kursor jest wewnątrz
    jego znaczników. Markdown nie ma podkreślenia: Underline (Ctrl+U) otacza tekst
    znacznikiem `<u>…</u>`, który domyślna lista dozwolonego HTML z Marp zachowuje, więc działa
    przy wyłączonym surowym HTML.
  - **Akapit**: Heading 1 do 3 (Ctrl+Alt+1..3), punktory (Ctrl+Shift+8),
    numerowanie (Ctrl+Shift+7) i cytat działają na zaznaczonych liniach i się przełączają.
    Linia pozioma to znacznik `<hr>`: `---`, `***` i `___` zaczęłyby w Marp
    każde nowy slajd.
  - **Tabela**: siatka wybiera kolumny i wiersze. Przy kursorze w tabeli możesz
    dodać lub usunąć wiersz albo kolumnę i wyrównać kolumnę (`:--`, `:-:`,
    `--:`); tabela jest po każdej edycji przepisywana wyrównana.
  - **Slide text** do lewej, do środka lub do prawej zapisuje blok `<style scoped>` w
    bieżącym slajdzie. Marp nie ma składni wyrównania, a atrybut `style`
    nie przechodzi przez domyślną listę dozwolonego HTML, ale blok scoped przechodzi, więc
    działa przy wyłączonym surowym HTML i widać go w podglądzie, na stronie HTML i w PDF.
  - **Link** (Ctrl+K) zapisuje `[text](url)` z zaznaczonym url.
- **Podświetlanie** kluczy front matter na dowolnej głębokości, w YAML (także
  w mapach `{a: 1}`) i w HCL.
- Wybór **motywu**: default, gaia, uncover, motywy projektu
  (`themes/*.css`) i motywy zdefiniowane w prezentacji (`/* @theme name */` w
  `<style>`). Zapisuje `theme:` w front
  matter, YAML lub HCL, więc udostępniony odnośnik go zachowuje, i podąża za front
  matter, gdy go edytujesz.
- Podgląd **zachowuje swoje miejsce** w trakcie pisania: wymieniane są tylko slajdy,
  które się zmieniły, a zmiana motywu otwiera podgląd ponownie na
  tym samym slajdzie.
- **Odnośnik do udostępnienia**: prezentacja podróżuje w samym odnośniku (w jego części `#`,
  skompresowana), więc otwarcie odnośnika otwiera prezentację. Odnośniki niosą do
  64 KiB po kompresji, i tylko prezentację: gdy projekt zawiera więcej
  (motywy, kinds, media), Share o tym mówi i oferuje `.zip`.
- **Save as**:
  - strona HTML, taka sama jak ta, którą zapisuje `slides deck.md`;
  - **PDF** (wektorowy, z zaznaczalnym tekstem, ze spisem treści);
  - edytowalna prezentacja **PowerPoint** lub **OpenDocument**.

  Powstają w twojej przeglądarce, przez ten sam kod co `slides --pdf`,
  `--pptx` i `--odp`: PDF z 100-slajdowej prezentacji testowej jest
  bajt w bajt identyczny z wynikiem polecenia. Pierwszy eksport pobiera
  silnik eksportu (około 5 MiB, jednorazowo; przeglądarka trzyma go do następnej
  aktualizacji), z linią postępu, i kompiluje go w miarę napływu. Potem pobiera tylko
  czcionki, które rysują twoje slajdy. Na łączu 3 MB/s pierwszy PDF z prezentacji
  powitalnej jest gotowy w mniej niż 3 sekundy. Prezentacja ze 100 slajdów zajmuje wtedy
  około 3 sekund. Obraz
  `https:` jest dołączany tylko wtedy, gdy jego serwer na to pozwala (CORS); obrazy `data:`
  zawsze.
- **Twój szkic jest zachowywany**:
  - Prezentacja jest zapisywana w tej przeglądarce w trakcie pisania (IndexedDB, tylko dla tej
    witryny) i wraca po przeładowaniu. Nic nie jest nigdzie wysyłane.
    Pliki tekstowe projektu (motywy, kinds, dane, `slides.yaml`) są zachowywane
    tak samo, nawet gdy zamkniesz kartę zaraz po pisaniu; media są
    zapisywane w chwili upuszczenia.
  - Udostępniony odnośnik otwiera swoją prezentację i zachowuje twój szkic: **Draft ▸ Restore my
    draft** go przywraca.
  - Przykład, **Draft ▸ Reset to the example** (który pyta wcześniej) i
    inna karta też go nigdy nie kasują: **Draft ▸ Restore replaced** zamienia go
    z powrotem.
  - **Draft ▸ Forget my drafts** kasuje wszystko, co zapisano.

Każdy eksport renderuje projekt, który pokazuje podgląd: jego `slides.yaml`,
motywy, kinds, pliki danych i media, z mediami wewnątrz dokumentu.

Obrazy jako pliki (slides `--images`) nie powstają w przeglądarce: użyj
[polecenia]({{< relref "cli" >}}) lub
[rozszerzenia VS Code]({{< relref "vscode" >}}).

## Bezpieczeństwo {#security}

Prezentacja jest pokazywana w iframe z piaskownicą bez pochodzenia strony: nie
sięgnie ani do Playground, ani do jego pamięci, nie wyśle żadnego formularza i nie
przejdzie nigdzie. Jej strona niesie Content Security Policy, która uruchamia tylko jej własne
skrypty. Surowy HTML w prezentacji podlega liście dozwolonych Marp, tak jak `slides` domyślnie.
Renderowanie trwające dłużej niż 5 sekund (szablon, który się zapętla)
jest przerywane.

Okno prelegenta zachowuje tę piaskownicę. Jest otwierane przez sam Playground
i pokazuje prezentację w iframe z taką samą piaskownicą. Dwa
iframe nie mogą współdzielić pamięci, więc Playground przekazuje numer slajdu
między nimi w komunikatach. Żadna ramka uruchamiająca prezentację nie dostaje
pochodzenia Playground.

Media projektu docierają do ramek w piaskownicy jako bajty w komunikacie, a
każda ramka tworzy z nich własne adresy `blob:`, w swoim pochodzeniu. Adres `blob:`
utworzony przez Playground tam się nie ładuje, a polityka ramek
zyskuje tylko `blob:` dla mediów i czcionek.

Playground można serwować z własnej maszyny, za logowaniem:
zobacz [Self-hosting]({{< relref "self-hosting" >}}).

## Rozmiar i szybkość {#size-and-speed}

Moduł ma 8,8 MiB w postaci, w jakiej serwuje go GitHub Pages (gzip; Pages nie
serwuje brotli). Jest pobierany raz z paskiem postępu, potem trzymany w pamięci podręcznej
przeglądarki do następnego wdrożenia. Na Apple M4 Max z Chrome 155, mierzone
na żywej stronie, edytor jest gotowy 2,1 s po otwarciu strony, z czego 1,5 s
to pobieranie, a pierwsze slajdy pojawiają się po 2,4 s. Zmiana w
100-slajdowej prezentacji testowej dociera do podglądu w około 30 ms: go-slides
renderuje tylko slajdy, które się zmieniły, i tylko te są wymieniane w
podglądzie.
