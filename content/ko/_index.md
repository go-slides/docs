---
title: "go-slides 문서"
linkTitle: "홈"
type: docs
cascade:
  type: docs
description: "순수 Go(CGO=0)로 만드는 Marp 호환 프레젠테이션. Go 템플릿 치환과 슬라이드 kind를 지원합니다."
---

**`slides`는 Marp 프레젠테이션을 순수 Go만으로 발표 자료로 만들어 줍니다.**
[Marp](https://marp.app)가 읽는 Markdown을 그대로 읽으며(Marpit 디렉티브,
슬라이드 사이의 `---`, `![bg]` 배경, 발표자 노트, `default`, `gaia`,
`uncover` 테마), Marp에 없는 두 가지 기능을 더합니다.

- 프레젠테이션 안의 **[Go 템플릿]({{< relref "templates" >}})**: Markdown을
  파싱하기 전에 평가되므로 프레젠테이션이 스스로 슬라이드를 생성할 수 있습니다.
- **[슬라이드 kind(슬라이드 유형)]({{< relref "kinds" >}})**: 슬라이드가 어떻게
  보이는지를 반복하는 대신 자신이 무엇인지(`title`, `section`…)를 밝힙니다.

{{< callout type="warning" >}}
**v0.14.0.** 이 문서는
[go-slides/slides v0.14.0](https://github.com/go-slides/slides/releases/tag/v0.14.0)을
설명하며 `/docs/0.14/` 아래에 게시됩니다.
{{< /callout >}}

## 현재 지원하는 기능 {#what-it-does-today}

| | |
| --- | --- |
| `slides build` (기본값) | 프레젠테이션 전체를 하나의 HTML 페이지로 만듭니다. 각 슬라이드는 창 크기에 맞춰 늘어나는 SVG 틀 안의 `<section>`이며, 테마 CSS와 키보드 뷰어가 함께 들어갑니다 |
| `slides build --pdf`, `--images png`, `--image jpeg` | 브라우저 없이 프레젠테이션을 PDF(슬라이드당 한 페이지, 개요 포함)나 PNG/JPEG 이미지로 만듭니다. 렌더링은 [go-webengine](https://github.com/go-webengine)이, PDF 작성은 [go-pdfkit](https://github.com/go-pdfkit)이 맡습니다. v0.4.0부터는 벡터로 작성되어 텍스트를 선택할 수 있습니다: [벡터 PDF]({{< relref "cli#vector-pdf" >}}) |
| 수식과 코드 | `$…$`와 `$$…$$`는 go-tex/math로 조판하고, 펜스 코드 블록은 go-rouge로 강조 표시합니다. HTML, 이미지, PDF 모두에 적용됩니다: [수식과 코드]({{< relref "math-and-code" >}}) |
| 이모지 | `:smile:`과 😄는 Twemoji 이미지로 바뀌어 내장되므로 네트워크 없이도 모든 출력에 그려집니다: [이모지]({{< relref "emoji" >}}) (v0.4.0부터) |
| `slides build --pptx` | 발표자 노트와 프레젠테이션 메타데이터를 담은 PowerPoint 파일로 만듭니다. v0.4.0에서는 marp-cli처럼 슬라이드당 그림 한 장이었고, v0.5.0부터는 텍스트 상자, 도형, 그림, 내장 글꼴로 이루어져 **편집할 수 있습니다** ([편집 가능한 PPTX와 ODP]({{< relref "cli#editable-pptx-and-odp" >}})) |
| `slides build --odp` | 같은 내용을 OpenDocument 프레젠테이션(`.odp`)으로 만듭니다 |
| 제목 맞춤 | `# <!-- fit --> Title`은 페이지와 모든 내보내기 결과에서 한 줄로 슬라이드 너비를 채웁니다 |
| `slides expand` | 프레젠테이션을 일반 Marp Markdown으로 펼칩니다. 템플릿은 평가되고, kind는 그에 해당하는 스폿 디렉티브로 바뀝니다 |
| `slides notes` | 발표자 노트를 출력합니다 |
| `slides kinds`, `slides themes` | 프레젠테이션에서 지정할 수 있는 이름을 보여 줍니다 |
| `slides.yaml` | 한 폴더의 프레젠테이션들이 공유하는 kind, 테마, 데이터, HTML 정책, 언어를 정합니다 |
| HCL | `---hcl` 프런트 매터, `slides.hcl`, HCL2로 쓴 kind와 데이터 파일을 표현식과 함수까지 지원합니다: [HCL]({{< relref "hcl" >}}) |

marp-cli와의 차이는 방식에
있습니다. marp-cli는 헤드리스 Chrome을 구동해 이미지와 PDF를 렌더링하지만,
여기서는 Go 라이브러리가 그 일을 합니다.

## 페이지 {#pages}

- [시작하기]({{< relref "getting-started" >}})
- [갤러리]({{< relref "gallery" >}}): Playground에서 열어 볼 수 있는 예제 프레젠테이션 12개
- [Marp 호환성]({{< relref "marp" >}})
- [템플릿]({{< relref "templates" >}})
- [Kind]({{< relref "kinds" >}})
- [HCL]({{< relref "hcl" >}})
- [수식과 코드]({{< relref "math-and-code" >}})
- [이모지]({{< relref "emoji" >}})
- [명령줄]({{< relref "cli" >}})
