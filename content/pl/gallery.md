---
title: "Galeria"
weight: 12
description: "Dwanaście małych prezentacji, z których każda dobrze pokazuje jedną rzecz: motywy, kinds, szablony, HCL, matematykę, obrazy, kod, emoji, przejścia i kompletne wystąpienie. Otwórz dowolną w Playground."
tags: [gallery, examples]
---

Dwanaście małych prezentacji z katalogu
[`examples/`](https://github.com/go-slides/slides/tree/main/examples) projektu
go-slides, z których każda dobrze pokazuje jedną rzecz. Miniatury to pierwszy
slajd każdej prezentacji, wyrenderowany przez sam go-slides podczas budowania tej
witryny; CI buduje każdą prezentację do HTML i PDF, więc żadna się nie psuje.

**Open in the playground** wczytuje prezentację w przeglądarce, przez odnośnik,
który nie wysyła niczego na żaden serwer. Prezentacji, która ma pliki poza
`deck.md` (plik kinds, dołączane części, obrazy), nie da się przenieść w
odnośniku: pobierz jej **.zip** i zaimportuj go panelem plików w Playground albo
rozpakuj i zbuduj poleceniem `slides build --allow-local-files deck.md`.

Obrazy to oryginalne pliki SVG autorów go-slides, udostępnione na licencji
[CC0 1.0](https://creativecommons.org/publicdomain/zero/1.0/), a prezentacje,
które ich używają, o tym mówią.

{{< gallery >}}
