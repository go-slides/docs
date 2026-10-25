---
title: "수식과 코드"
weight: 35
description: "달러 기호 사이의 수식은 조판되고 펜스 코드는 강조 표시됩니다. HTML 페이지, 이미지, PDF 모두에 적용됩니다."
tags: [math, code]
---

{{< callout type="warning" >}}
수식과 코드는 v0.2.0에 추가되었습니다.
{{< /callout >}}

## 수식 {#math}

`$…$`는 인라인 수식이고 `$$…$$`는 디스플레이 수식입니다.

```markdown
The famous one is $E=mc^2$, and this is a display formula:

$$
\int_0^\infty e^{-x^2}\,dx = \frac{\sqrt{\pi}}{2}
$$
```

- 구분자 규칙은 marp-core 4.4.0에서 확인한 Marp Core의 규칙입니다. 뒤에 공백이
  오는 `$`는 수식을 열지 않고, 뒤에 숫자가 오는 달러 기호는 수식을 닫지 않으므로
  `$5 and $10` 같은 가격은 텍스트로 남습니다.
- [go-tex/math](https://github.com/go-tex/math)는 각 수식을 `em` 단위로 크기가
  정해진 인라인 SVG로 조판하므로, 수식이 텍스트 크기를 따르고 기준선 위에
  놓입니다. HTML 페이지, 이미지, PDF에서 같은 SVG를 씁니다.
- 파싱할 수 없는 수식은 쓴 그대로 보이며, 오류는 그 `title`에 담깁니다.

## 코드 {#code}

펜스 코드 블록은 [go-rouge](https://github.com/go-rouge/rouge)가 색을 입힙니다.
go-rouge는 Marp 테마가 스타일을 지정하는 highlight.js 클래스 이름을 씁니다.

````markdown
```go
fmt.Println("hello")
```
````

go-rouge가 모르는 언어는 Marp Core와 마찬가지로 일반 텍스트로 출력합니다.

### 줄 번호 {#line-numbers}

v0.8.0부터 블록에 줄 번호를 붙일 수 있습니다. Marp Core에는 줄 번호가 없고
펜스의 언어 뒤에 오는 내용을 무시하므로, 옵션은 그 자리에 중괄호로 씁니다.
이 옵션을 쓴 프레젠테이션도 Marp에서는 옵션이 없을 때와 똑같이 렌더링됩니다.

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

| 표기 | 의미 |
| --- | --- |
| `{lineNumbers}` | 1부터 번호를 매깁니다(Slidev의 `{lines:true}`와 Pandoc의 `{.numberLines}`도 인식합니다) |
| `{lineNumbers=10}` | 10부터 번호를 매깁니다(`startLine=10`, Pandoc의 `startFrom="10"`도 가능) |
| `{lineNumbers=false}` | 슬라이드 설정과 관계없이 번호를 매기지 않습니다 |

한 슬라이드와 그 뒤 슬라이드의 모든 블록에 적용하려면 디렉티브
`<!-- lineNumbers: true -->`를 설정합니다(한 슬라이드만이면 `_lineNumbers`).
kind(슬라이드 유형)로 같은 계열의 슬라이드 전체에 한 번에 설정할 수도 있습니다.

```yaml
kinds:
  listing: {lineNumbers: true}
```

각 번호는 테마의 코드 상자 안에서 줄 앞에 놓이며 구분선으로 분리됩니다.
콜아웃이나 이모지 때문에 줄이 더 높아져도 번호는 그 줄의 기준선에 머뭅니다.
번호는 코드 색과 배경색을 섞은 흐린 색으로 표시되며, default, gaia, uncover
테마의 밝은 모드와 어두운 모드 모두에서 명암비가 4.5:1 이상입니다. 번호는
페이지에서 **선택할 수 없고** 화면 낭독기도 읽지 않습니다. PDF에서는 그려지되
텍스트에서는 제외되므로, 텍스트 추출기는 코드만 읽습니다. PowerPoint나
OpenDocument 파일에서는 각 줄 앞의 편집 가능한 텍스트가 됩니다. 이 옵션을
쓰지 않는 프레젠테이션은 바이트 단위까지 이전과 똑같이 렌더링됩니다. 특정
줄을 강조하는 기능은 아직 지원하지 않습니다.

### 코드 콜아웃 {#code-callouts}

v0.8.0부터 코드 줄 끝에 블록 아래의 설명을 가리키는 번호 원을 붙일 수 있습니다.
문법은 [Quarto의 문법](https://quarto.org/docs/authoring/code-annotation.html)으로,
Asciidoctor의 주석 형식이기도 합니다. 표식은 해당 코드 언어의 주석이므로 코드는
유효한 상태로 유지되며, 설명은 블록 바로 뒤에 오는 번호 목록입니다.

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

| 표기 | 의미 |
| --- | --- |
| `// <1>`, `# <1>`, `-- <1>`, `; <1>`, `% <1>`, `' <1>` | 표식만 담은 한 줄 주석: 주석과 표식이 모두 번호 원으로 바뀝니다 |
| `<!-- <1> -->`, `/* <1> */` | 감싸는 형식의 주석을 쓰는 언어에서도 마찬가지입니다 |
| `// <1> <2>` | 한 줄에 번호 원 두 개 |
| `x = 1 # count <1>` | 주석 텍스트는 남고 그 뒤에 번호 원이 붙습니다 |

표식은 블록 바로 뒤의 목록이 1부터 시작하고 모든 번호에 해당하는 항목이 있을
때만 번호 원이 됩니다. 그렇지 않으면 블록은 쓴 그대로 남습니다. 번호 원은
default, gaia, uncover 테마의 밝은 모드와 어두운 모드 모두에서 코드 색의 고리로
그려집니다. 목록 항목에도 같은 번호 원이 붙습니다.

번호 원은 **선택할 수 없습니다**. 페이지에서 코드를 복사하면 표식이 없는 코드만
복사됩니다. PDF에서는 텍스트에서 제외된 벡터이고, PowerPoint나 OpenDocument
파일에서는 편집 가능한 도형입니다. 콜아웃은 줄 번호와 함께 쓸 수 있습니다.
Marp Core에는 콜아웃이 없으므로(marp-core 4.4.0) Marp에서는 주석과 번호 목록이
쓴 그대로 보입니다. 콜아웃을 쓰지 않는 프레젠테이션은 바이트 단위까지 이전과
똑같이 렌더링됩니다.

v0.9.0부터는 HTML 페이지에서 번호 원과 그 설명이 서로 연결됩니다. 번호 원에
마우스를 올리거나 Tab 키로 이동하면 해당 설명에 색이 입혀지고, 설명에서도
마찬가지로 번호 원이 강조됩니다. Quarto의 `code-annotations: hover`와 같은
동작이지만, 여기서는 스크립트 없이 스타일시트만으로 처리합니다. 따라서 둘 다
스크립트를 제한하는 Playground의 샌드박스 미리 보기와 VS Code 미리 보기에서도
똑같이 동작합니다. 각 번호 원은 설명으로 기술되므로, 화면 낭독기는 번호 원에
이르면 설명을 읽습니다. 색 변화는 사용자가 동작 줄이기를 요청하지 않은 경우에만
서서히 나타납니다. 이미지, PDF, PowerPoint, OpenDocument 파일은 바이트 단위까지
바뀌지 않습니다.

## 글꼴 {#fonts}

Lato Regular, Bold, Black과 이탤릭, 그리고 Roboto Mono Regular, Bold와
이탤릭이 내장되어 있으므로 gaia의 제목은 제 굵기로 표시되고 이탤릭 텍스트는
이탤릭으로 렌더링됩니다.
