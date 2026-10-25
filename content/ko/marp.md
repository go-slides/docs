---
title: "Marp 호환성"
linkTitle: "Marp"
weight: 20
description: "slides가 읽는 Markdown은 Marpit과 Marp Core의 Markdown 그 자체이며, 별도의 방언이 아닙니다."
tags: [marp]
---

[Marp](https://marp.app)용으로 작성한 프레젠테이션은 그대로 go-slides
프레젠테이션입니다. 아래 규칙은 [Marpit](https://marpit.marp.app)과
[Marp Core](https://github.com/marp-team/marp-core)의 규칙이며, go-slides가
다른 부분은 이 페이지에 따로 밝힙니다.

## 슬라이드 {#slides}

수평선, 즉 Marpit이 나누는 방식대로 한 줄에 단독으로 쓴 `---`는 새 슬라이드를
시작합니다. `headingDivider` 디렉티브가 지정한 수준의 제목도 마찬가지입니다.
처음 두 `---` 사이의 프런트 매터에는 전역 디렉티브와 프레젠테이션 데이터가
들어갑니다.

## 디렉티브 {#directives}

| 종류 | 위치 | 적용 범위 |
| --- | --- | --- |
| 전역 | 프런트 매터, 또는 어디서나 `<!-- theme: gaia -->` | 프레젠테이션 전체: `theme`, `style`, `headingDivider`, `lang`… |
| 로컬 | `<!-- paginate: true -->` | 이 슬라이드와 그 뒤의 슬라이드 |
| 스폿 | `<!-- _class: lead -->` (밑줄) | 이 슬라이드만 |

로컬 디렉티브는 Marpit의 것입니다: `paginate`, `header`, `footer`, `class`,
`backgroundColor`, `backgroundImage`, `backgroundPosition`,
`backgroundRepeat`, `backgroundSize`, `color`. 디렉티브는 Marp Core처럼
느슨한 YAML로 파싱하며, `--strict-yaml`을 쓰면 순수 Marpit처럼 파싱합니다.

## 이미지 {#images}

`![bg](image.jpg)`는 이미지를 슬라이드 배경으로 만들며, Marp의 키워드를 그대로
씁니다. `bg contain`, `bg cover`, `bg 50%`, 화면을 나누는 배경을 위한
`bg left`/`bg right`, 나란히 놓는 여러 장의 `bg` 이미지, 그리고 이미지 필터
(`blur`, `sepia`, `grayscale`…)를 지원합니다. 인라인 이미지 크기(`w:200px`,
`h:`)도 Marp와 똑같이 읽습니다.

## 슬라이드 주변 요소 {#around-the-slide}

- **헤더와 푸터**: Markdown으로 쓰는 `header:`와 `footer:` 디렉티브.
- **페이지 번호**: `paginate: true`.
- **발표자 노트**: 디렉티브가 아닌 HTML 주석은 노트가 됩니다.
  `slides notes`로 출력할 수 있습니다.
- **단계별 목록**: `*`(또는 `1)`)로 쓴 목록은 Marp처럼 항목이 하나씩
  나타납니다.
- **`<style scoped>`**: 한 슬라이드에만 적용되는 CSS.
- **원시 HTML**: 기본값은 Marp Core의 허용 목록이며, `--html`을 쓰면 모두
  허용합니다.
- **이모지**: `:shortcodes:`와 Unicode 이모지는 Marp Core처럼 Twemoji 이미지가
  되지만, CDN에서 불러오지 않고 내장됩니다. [이모지]({{< relref "emoji" >}})를
  참고하십시오.

## 테마 {#themes}

`default`, `gaia`, `uncover`가 기본 제공되며 `theme:`로 선택합니다. 직접 만든
테마는 Marp의 `/* @theme name */` 주석이 들어 있는 CSS 파일이며,
`--theme-set`으로 지정하거나 `slides.yaml`의 `themes:` 아래에 나열합니다.

## go-slides가 다른 점 {#where-go-slides-differs}

- **템플릿.** 본문의 `{{ }}`는 Go 템플릿 액션입니다
  ([템플릿]({{< relref "templates" >}}) 참고). 코드 블록 밖에서 `{{`를
  텍스트로 보여 주려면 `--no-template`을 쓰거나 Go 템플릿 방식대로 이스케이프해야
  합니다. 코드 블록과 인라인 코드는 건드리지 않습니다.
- **Kind.** `kind:`와 `_kind:`는 Marp에 없는 디렉티브입니다
  ([Kind]({{< relref "kinds" >}}) 참고). `slides expand`가 이를 해당하는
  디렉티브로 바꾸므로, 펼친 프레젠테이션은 일반 Marp입니다.
- **출력.** HTML, PNG, JPEG, PDF, PPTX, ODP를 헤드리스 브라우저가 아닌 Go
  라이브러리로 만듭니다(알려진 제한 사항은
  [명령줄]({{< relref "cli#images-pdf-pptx-and-odp" >}}) 참고). v0.4.0에서는
  marp-cli처럼 PPTX에 그림이 들어갔지만, v0.5.0부터는 슬라이드가 편집 가능한
  텍스트 상자와 도형이며 `--pptx-raster`를 쓰면 그림으로 유지됩니다.
- **제목 맞춤.** `# <!-- fit --> Title`을 지원합니다. 제목이 페이지와 모든
  내보내기 결과에서 한 줄로 슬라이드 너비를 채웁니다.
- **수식과 코드.** `$…$`와 `$$…$$`는 Marp Core의 구분자 규칙을 따르며
  go-tex/math로 조판됩니다. 펜스 코드는 go-rouge가 highlight.js 클래스 이름으로
  강조 표시합니다([수식과 코드]({{< relref "math-and-code" >}}) 참고).
