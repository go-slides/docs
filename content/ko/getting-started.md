---
title: "시작하기"
weight: 10
description: "첫 프레젠테이션을 HTML로 빌드하고 다시 Marp Markdown으로 펼쳐 봅니다."
tags: [start]
---

## 설치 {#install}

Go 1.27.2 이상이 필요합니다.

```sh
go install github.com/go-slides/slides/cmd/slides@latest
```

`slides`는 cgo를 쓰지 않는 정적 Go 바이너리 하나입니다. Node도 Chrome도 필요
없고, 따로 설치할 것이 없습니다. 이 페이지는 v0.13.0을 기준으로 설명합니다.

## 첫 프레젠테이션 {#a-first-deck}

Marp 프레젠테이션이라면 무엇이든 그대로 쓸 수 있습니다. 다음 내용을 `talk.md`로
저장하십시오.

```markdown
---
marp: true
theme: default
paginate: true
footer: "{{ .talk }}"
data:
  talk: Pure Go slides
  topics: [Why, How, Next]
---

<!-- _kind: title -->
# {{ .talk }}

{{ range .topics }}
---

# {{ . }}
{{ end }}
```

프런트 매터는 Marp의 것에 `data:`가 더해진 것입니다. 본문은 Go 템플릿입니다.
`range`는 주제마다 `---`로 시작하는 슬라이드를 하나씩 만들고, 첫 슬라이드는
기본 제공 kind(슬라이드 유형)인 `title`(헤더, 푸터, 페이지 번호 없음)입니다.

## 빌드하기 {#build-it}

```sh
slides talk.md                 # writes talk.html beside the deck
slides -o slides.html talk.md  # or somewhere else; -o - for standard output
```

명령 없이 실행한 `slides`는 `slides build`와 같습니다. 플래그는 프레젠테이션
앞에 와도, 뒤에 와도 됩니다. HTML을 브라우저에서 열면 키보드로 슬라이드를
넘길 수 있습니다.

## 이미지, PDF, PPTX, ODP로 내보내기 {#export-images-pdf-pptx-and-odp}

```sh
slides talk.md --pdf                  # one page per slide, with an outline
slides talk.md --images png           # talk.001.png, talk.002.png, …
slides talk.md --image jpeg           # the first slide only
slides talk.md -o out/slides.pdf      # the extension picks the format; out/ is created
slides talk.md --images jpeg --jpeg-quality 90
slides talk.md --images png --image-scale 2   # 2560×1440 for 16:9
slides talk.md --pdf --allow-local-files      # images from under the deck's directory
slides talk.md --pptx                         # or -o talk.pptx, with notes
slides talk.md --odp                          # or -o talk.odp: the same, as OpenDocument
```

브라우저는 전혀 쓰이지 않습니다. 슬라이드는 go-webengine이 렌더링하고 PDF는
go-pdfkit이 작성합니다. 로컬 이미지는 `--allow-local-files`를 지정하지 않으면
거부됩니다. v0.5.0부터 PPTX와 ODP의 슬라이드는 텍스트 상자, 도형, 그림, 내장
글꼴로 이루어져 편집할 수 있습니다(marp-cli는 그림입니다). 그림으로 유지하려면
`--pptx-raster`를 쓰십시오([편집 가능한 PPTX와 ODP]({{< relref "cli#editable-pptx-and-odp" >}})). v0.4.0부터 PDF는 벡터이며
텍스트를 선택할 수 있습니다(v0.3.0은 페이지마다 이미지 한 장을 썼습니다). [명령줄]({{< relref "cli#vector-pdf" >}})을 참고하십시오.

## 펼치기 {#expand-it}

```sh
slides expand talk.md
```

이 명령은 같은 프레젠테이션을 **일반 Marp Markdown**으로 출력합니다
(`4ae5c49` 시점, 2026-10-18의 `main` 출력).

```markdown
---
marp: true
theme: default
paginate: true
footer: "Pure Go slides"
---
<!--
_class: "lead"
_footer: ""
_paginate: "false"
-->


# Pure Go slides


---

# Why

---

# How

---

# Next
```

`range`가 생성한 슬라이드 네 장, 값이 채워진 푸터, 그리고 `title` kind 대신
그에 해당하는 스폿 디렉티브가 보입니다. marp-cli나 VS Code 확장처럼 Marp를
읽는 도구라면 무엇이든 이어서 처리할 수 있으므로, go-slides의 추가 기능으로
작성한 내용이 특정 도구에 묶이지 않습니다.

## 다음 단계 {#next}

- [Marp 호환성]({{< relref "marp" >}}): Marp와 똑같이 읽는 부분.
- [템플릿]({{< relref "templates" >}})과 [Kind]({{< relref "kinds" >}}): 두 가지 추가 기능.
- [명령줄]({{< relref "cli" >}}): 모든 명령과 플래그.
