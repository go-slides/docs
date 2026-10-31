---
title: "Matematyka i kod"
weight: 35
description: "Wzory między znakami dolara są składane, kod w blokach jest podświetlany, na stronie HTML, w obrazach i w PDF."
tags: [math, code]
---

{{< callout type="warning" >}}
Matematyka i kod są w v0.2.0.
{{< /callout >}}

## Matematyka {#math}

`$…$` to wzór w tekście, a `$$…$$` wzór wyświetlany osobno:

```markdown
The famous one is $E=mc^2$, and this is a display formula:

$$
\int_0^\infty e^{-x^2}\,dx = \frac{\sqrt{\pi}}{2}
$$
```

- Reguły ograniczników pochodzą z Marp Core, odczytane z marp-core 4.4.0: `$`
  po którym stoi spacja nie otwiera wzoru, a po którym stoi cyfra
  go nie zamyka, więc ceny takie jak `$5 and $10` pozostają tekstem.
- [go-tex/math](https://github.com/go-tex/math) składa każdy wzór w
  SVG w tekście o rozmiarze w `em`, więc podąża za rozmiarem tekstu i siedzi na jego
  linii bazowej. Ten sam SVG jest używany na stronie HTML, w obrazach i w PDF.
- Wzór, który się nie parsuje, pozostaje widoczny tak, jak go napisano; błąd jest zachowany w
  jego `title`.

## Kod {#code}

Bloki kodu są kolorowane przez [go-rouge](https://github.com/go-rouge/rouge),
który zapisuje nazwy klas highlight.js, te, którymi stylizują motywy Marp:

````markdown
```go
fmt.Println("hello")
```
````

Język, którego go-rouge nie zna, jest zapisywany jako zwykły tekst, jak robi to Marp Core.

### Numery linii {#line-numbers}

W v0.8.0 blok może numerować swoje linie. Marp Core nie ma numerów linii i
ignoruje to, co następuje po języku w ogrodzeniu, więc opcje idą tam w klamrach,
a prezentacja, która ich używa, renderuje się w Marp tak, jakby ich nie było:

````markdown
```go {lineNumbers}
func main() {
	fmt.Println("hello")
}
```

```go {lineNumbers=10}
// starts at 10
```
````

| Zapis | Znaczenie |
| --- | --- |
| `{lineNumbers}` | numeruj od 1 (czytane są też `{lines:true}` ze Slidev i `{.numberLines}` z Pandoc) |
| `{lineNumbers=10}` | numeruj od 10 (także `startLine=10`, `startFrom="10"` z Pandoc) |
| `{lineNumbers=false}` | bez numerów, cokolwiek mówi slajd |

Dla każdego bloku slajdu i następnych ustaw dyrektywę:
`<!-- lineNumbers: true -->` (`_lineNumbers` dla jednego slajdu). Rodzaj może ustawić ją
raz dla rodziny slajdów:

```yaml
kinds:
  listing: {lineNumbers: true}
```

Każdy numer zaczyna swoją linię wewnątrz ramki kodu motywu, oddzielony linią.
Zostaje na linii bazowej swojej linii, nawet gdy linia jest wyższa, na przykład
z objaśnieniem lub emoji. Numery są wyciszone kolorem zmieszanym z koloru
kodu i jego tła, z kontrastem 4,5:1 lub większym w motywach default,
gaia i uncover, jasnych i ciemnych. Na stronie **nie dają się zaznaczyć** i
nie są czytane przez czytniki ekranu. W PDF są narysowane, ale pominięte
w tekście, więc ekstraktor tekstu czyta sam kod. W pliku PowerPoint lub
OpenDocument są edytowalnym tekstem na początku każdej linii. Prezentacja,
która nie używa tej opcji, renderuje się jak dawniej, co do bajta. Podświetlanie
wybranych linii nie jest jeszcze obsługiwane.

### Objaśnienia kodu {#code-callouts}

W v0.8.0 linia kodu może kończyć się numerowanym krążkiem odsyłającym do notatki
pod blokiem. Składnia pochodzi z
[Quarto](https://quarto.org/docs/authoring/code-annotation.html), jest to też
forma komentarza z Asciidoctor. Znacznik jest komentarzem w języku samego kodu,
więc kod pozostaje poprawny, a notatki to lista numerowana
zaraz po bloku:

````markdown
```go
func main() { // <1>
	x := load() // <2>
	fmt.Println(x)
}
```

1. The entry point.
2. Reads the input.
````

| Zapis | Znaczenie |
| --- | --- |
| `// <1>`, `# <1>`, `-- <1>`, `; <1>`, `% <1>`, `' <1>` | komentarz liniowy zawierający tylko znacznik: oba są zastępowane krążkiem |
| `<!-- <1> -->`, `/* <1> */` | to samo, dla języków z komentarzami ujętymi w ograniczniki |
| `// <1> <2>` | dwa krążki w jednej linii |
| `x = 1 # count <1>` | tekst komentarza zostaje, a krążek idzie za nim |

Znaczniki stają się krążkami tylko wtedy, gdy lista zaraz po bloku zaczyna się od 1
i ma element dla każdego numeru. W przeciwnym razie blok zostaje tak, jak napisano. Krążki
to pierścienie w kolorze kodu, w motywach default, gaia i uncover,
jasnych lub ciemnych. Elementy listy są numerowane tymi samymi krążkami.

Krążków **nie da się zaznaczyć**. Skopiowanie kodu ze strony daje kod
bez znaczników. W PDF są wektorami pominiętymi w tekście, a w pliku
PowerPoint lub OpenDocument edytowalnymi kształtami. Objaśnienia łączą się
z numerami linii. Marp Core nie ma objaśnień (marp-core 4.4.0), więc w Marp
komentarze i lista numerowana wyświetlają się tak, jak napisano. Prezentacja, która ich nie używa,
renderuje się jak dawniej, co do bajta.

W v0.9.0 krążek i jego notatka znajdują się nawzajem na stronie HTML. Najechanie na
krążek lub dojście do niego klawiszem Tab podświetla jego notatkę, a notatka robi to samo dla
swojego krążka. Tak działa `code-annotations: hover` z Quarto, ale tutaj arkusz
stylów robi to sam, bez skryptu. Działa więc tak samo w
piaskownicowym podglądzie playgroundu i w podglądzie VS Code, które oba ograniczają
skrypty. Każdy krążek jest opisany swoją notatką, więc czytnik ekranu czyta notatkę,
gdy dojdzie do krążka. Podświetlenie pojawia się płynnie tylko wtedy, gdy czytelnik nie prosił
o ograniczenie ruchu. Obrazy, PDF, pliki PowerPoint i OpenDocument są
bez zmian, co do bajta.

## Czcionki {#fonts}

Lato Regular, Bold, Black i kursywy oraz Roboto Mono Regular, Bold i
kursywy są osadzone, więc nagłówki gaia dostają swoją grubość, a tekst
pochylony renderuje się jako kursywa.
