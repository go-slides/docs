---
title: "Rodzaje"
weight: 40
description: "Slajd mówi, czym jest: kind nazywa zestaw dyrektyw lokalnych."
tags: [kinds]
---

**Rodzaj (kind)** to nazwany zestaw dyrektyw lokalnych: nagłówek, stopka,
numeracja stron, klasa, tło. Slajd mówi, czym jest — `title`,
`section` — zamiast powtarzać, jak wygląda, a wygląd prezentacji zmienia się
w jednym miejscu.

## Użycie rodzaju {#using-a-kind}

```markdown
<!-- kind: section -->
# Part two

---

<!-- _kind: blank -->
![bg](photo.jpg)
```

- `kind:` działa jak dyrektywa lokalna: na ten slajd i następne.
- `_kind:` działa jak dyrektywa punktowa: tylko na ten slajd.
- Przełączenie rodzaju najpierw **czyści** to, co ustawił poprzedni, więc klasa
  czy stopka nigdy nie przecieka z jednego rodzaju do następnego.
- Dyrektywy zapisane na samym slajdzie **mają pierwszeństwo** przed rodzajem.

## Wbudowane {#built-in}

| Rodzaj | Ustawia |
| --- | --- |
| `title` | `class: lead`, bez numeru strony, bez nagłówka, bez stopki |
| `section` | `class: lead` |
| `content` | nic: własne wartości domyślne prezentacji |
| `blank` | bez numeru strony, bez nagłówka, bez stopki |

Każdy z nich można przedefiniować.

## Definiowanie rodzajów {#defining-kinds}

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

`extends` zaczyna rodzaj od dyrektyw innego. Rodzaj przyjmuje dyrektywy
lokalne i `extends`.

## Współdzielenie rodzajów {#sharing-kinds}

Rodzaje zespołu należą do jednego pliku, który wskazuje każda prezentacja:

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
Kolejność jest taka: rodzaje wbudowane, potem wiersz poleceń i plik
konfiguracyjny, potem własne `kinds:` prezentacji.

`slides kinds talk.md` wypisuje rodzaje, z których prezentacja może korzystać, i
to, co każdy z nich ustawia.

## Opuszczenie go-slides {#leaving-go-slides}

`slides expand` zapisuje każdy rodzaj jako dyrektywy punktowe, które
reprezentuje, slajd po slajdzie, więc wynik to zwykły Markdown Marp, który
odczyta każde narzędzie Marp.
