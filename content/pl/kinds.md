---
title: "Kinds"
weight: 40
description: "Slajd mówi, czym jest: kind nazywa zestaw dyrektyw lokalnych."
tags: [kinds]
---

**Kind** (rodzaj slajdu) to nazwany zestaw dyrektyw lokalnych: nagłówek, stopka,
numeracja stron, klasa, tło. Slajd mówi, czym jest — `title`,
`section` — zamiast powtarzać, jak wygląda, a wygląd prezentacji zmienia się
w jednym miejscu.

## Użycie kind {#using-a-kind}

```markdown
<!-- kind: section -->
# Part two

---

<!-- _kind: blank -->
![bg](photo.jpg)
```

- `kind:` działa jak dyrektywa lokalna: na ten slajd i następne.
- `_kind:` działa jak dyrektywa punktowa: tylko na ten slajd.
- Przełączenie kind najpierw **czyści** to, co ustawił poprzedni, więc klasa
  czy stopka nigdy nie przecieka z jednego kind do następnego.
- Dyrektywy zapisane na samym slajdzie **mają pierwszeństwo** przed kind.

## Wbudowane {#built-in}

| Kind | Ustawia |
| --- | --- |
| `title` | `class: lead`, bez numeru strony, bez nagłówka, bez stopki |
| `section` | `class: lead` |
| `content` | nic: własne wartości domyślne prezentacji |
| `blank` | bez numeru strony, bez nagłówka, bez stopki |

Każdy z nich można przedefiniować.

## Definiowanie kinds {#defining-kinds}

We front matter:

```yaml
---
marp: true
kinds:
  chapter:
    class: invert
    backgroundColor: "#1e1b4b"
    paginate: false
  quote:
    extends: blank
    class: quote
---
```

`extends` zaczyna kind od dyrektyw innego. Kind przyjmuje dyrektywy
lokalne i `extends`.

## Współdzielenie kinds {#sharing-kinds}

Kinds zespołu należą do jednego pliku, który wskazuje każda prezentacja:

```yaml
# kinds.yaml
chapter:
  class: invert
  paginate: false
```

```yaml
---
marp: true
kinds: kinds.yaml        # or a list of files
---
```

albo `--kinds kinds.yaml` w wierszu poleceń, albo `kinds:` w `slides.yaml`.
Kolejność jest taka: wbudowane kinds, potem wiersz poleceń i plik
konfiguracyjny, potem własne `kinds:` prezentacji.

`slides kinds talk.md` wypisuje kinds, z których prezentacja może korzystać, i
to, co każdy z nich ustawia.

## Opuszczenie go-slides {#leaving-go-slides}

`slides expand` zapisuje każdy kind jako dyrektywy punktowe, które
reprezentuje, slajd po slajdzie, więc wynik to zwykły Markdown Marp, który
odczyta każde narzędzie Marp.
