---
title: "HCL"
weight: 45
description: "프레젠테이션이 YAML을 받는 모든 곳에서 HCL2를 쓸 수 있습니다: HCL 프런트 매터, slides.hcl, kind와 데이터 파일, 표현식과 함수."
tags: [hcl, configuration]
---

{{< callout type="warning" >}}
HCL은 v0.3.0에 추가되었습니다([go-slides/slides#13](https://github.com/go-slides/slides/pull/13)).
{{< /callout >}}

`slides`는 YAML을 읽는 모든 곳에서
[HCL2](https://github.com/hashicorp/hcl/blob/main/hclsyntax/spec.md)의 네이티브
문법도 읽습니다. 표현식과 함수를 쓸 수 있기 때문입니다.

| 위치 | YAML | HCL |
| --- | --- | --- |
| 프런트 매터 | `---` … `---` | `---hcl` … `---` |
| 설정 | `slides.yaml` | `slides.hcl`, 먼저 찾음 |
| kind 파일 | `--kinds kinds.yaml`, `kinds: kinds.yaml` | 이름이 `*.hcl`인 파일 |
| 데이터 파일 | `--data data.yaml` 또는 `.json` | 이름이 `*.hcl`인 파일 |

디렉티브 주석(`<!-- class: invert -->`)은 Marp와 마찬가지로 YAML로 남고,
프레젠테이션 본문은 그대로 [Go 템플릿]({{< relref "templates" >}})입니다.

## HCL로 쓴 프런트 매터 {#a-front-matter-in-hcl}

```hcl
---hcl
theme    = "gaia"
paginate = true
title    = "Quarterly review"
data = {
  company = "ACME"
  quarter = 3
}
header = format("%s — Q%d", data.company, data.quarter)
footer = "${upper(data.company)} · ${page()}/${total()}"
date   = formatdate("D MMMM YYYY", now())

kind "chapter" {
  extends = "section"
  class   = ["lead", "invert"]
}
---

# {{ .title }}
```

- **속성**은 같은 이름의 YAML 키에 해당합니다. Marp 디렉티브(`theme`,
  `paginate`, `header`, `footer`, `class`, `backgroundColor`…), 메타데이터
  (`title`, `author`…), `template`, `delims`, `kinds`(파일 경로), 그리고
  객체인 `data`가 있습니다.
- **`kind "<name>" { … }`** 블록은 로컬 디렉티브와 `extends`로
  [kind(슬라이드 유형)]({{< relref "kinds" >}})를 정의합니다.
- **값**은 YAML이 주는 것과 같은 문자열 형태로 디렉티브에 전달됩니다.
  `true`는 `"true"`가 되고 클래스 목록은 목록이 되므로, 어느 문법으로 설정했든
  디렉티브는 똑같이 동작합니다. `null`은 디렉티브를 설정하지 않은 상태로 둡니다.
- **`data`**는 타입을 유지합니다. 템플릿에서 정수는 `int`, 그 밖의 수는
  `float64`가 되고, 목록과 객체는 목록과 맵으로 남으므로
  `{{ add .quarter 1 }}`이 동작합니다.
- HCL 값은 최종 값입니다. Go 템플릿이 그 위에서 다시 실행되지 않습니다.

## 변수 {#variables}

`data`는 템플릿 데이터입니다. 프런트 매터 자체의 `data` 위에
`slides.hcl`의 데이터, `--data` 파일, `--set`이 차례로 덮어씁니다.
`{{ .company }}`를 쓸 수 있는 곳이면 어디서나 `"${data.company}"`를 쓸 수
있습니다.

프런트 매터 속성은 `data.*`를 참조할 수 있지만 **서로를 참조할 수는 없습니다**.
`footer = title`은 오류입니다. 스폿 디렉티브(`_class`…)는 HCL 식별자가
아니므로 주석으로 쓰십시오.

## 함수 {#functions}

| | |
| --- | --- |
| `page()`, `total()` | `{{ page }}`, `{{ total }}`과 같은 페이지 번호와 마지막 페이지 번호. 푸터에 페이지를 매길 수 있습니다 |
| `now()`, `timestamp()` | 현재 시각(RFC 3339, 로컬 시각과 UTC). `formatdate`로 서식을 지정합니다 |
| `env(name)` | 환경 변수. 설정되지 않았으면 `""` |
| `file(path)` | 프레젠테이션이나 설정 파일 기준 상대 경로의 UTF-8 텍스트 파일(최대 1 MiB) |
| 숫자 | `abs ceil floor log max min parseint pow signum` |
| 문자열 | `chomp format formatlist join lower regex regexall regexreplace replace split strlen strrev substr title trim trimprefix trimspace trimsuffix upper` |
| 컬렉션 | `chunklist coalesce coalescelist compact concat contains distinct element flatten keys length lookup merge range reverse slice sort values zipmap` |
| 인코딩 | `csvdecode jsondecode jsonencode` |
| 날짜 | `formatdate timeadd` |

처음 네 행을 제외한 나머지는 모두 [go-cty](https://github.com/zclconf/go-cty)의
표준 라이브러리입니다. 어느 함수도 네트워크에 접근하지 않습니다.

## slides.hcl {#slideshcl}

```hcl
kinds  = "kinds.hcl"          # or a list of files
themes = ["themes/acme.css"]
data   = { company = "ACME" } # or file paths
html   = false
lang   = "en"

kind "quote" {
  extends = "blank"
  footer  = "${data.company}"
}
```

알 수 없는 속성은 조용히 무시되지 않고 오류가 됩니다. kind 파일에는 `kind`
블록을 담으며, 그 밖의 속성은 무시되므로 `slides.hcl`을 kind 파일로도 쓸 수
있습니다. 데이터 파일의 속성은 데이터의 키가 됩니다.

```hcl
# data.hcl
company = "ACME"
year    = 2026
motto   = trimspace(file("motto.txt"))
```

## 오류 {#errors}

오류는 HCL의 진단 메시지이며, 위치는 프레젠테이션 파일 자체를 기준으로
표시됩니다. 프레젠테이션 넷째 줄의 `footer = data.compnay`는 다음 오류를
냅니다.

```
deck.md:4:14: Unsupported attribute: This object does not have an attribute named "compnay".
```

## Marp와 `slides expand` {#marp-and-slides-expand}

Marp도 `---hcl`을 프런트 매터의 시작으로 읽지만(Marp의 프런트 매터 규칙은
대시 세 개로 시작하는 첫 줄이면 무엇이든 받아들입니다), 이 블록은 YAML 매핑이
아니므로 아무것도 적용하지 않습니다. 따라서 프레젠테이션은 Marp에서도
렌더링되지만 디렉티브는 빠집니다. 이 동작은 Marp Core 자체의 출력과 비교하는
테스트로 고정되어 있습니다.

`slides expand`는 프런트 매터를 YAML로 다시 쓰므로, 펼친 프레젠테이션은 일반
Marp이며 `slides`가 원본을 렌더링한 것과 똑같이 렌더링됩니다.

```yaml
---
theme: "gaia"
paginate: true
title: "Quarterly review"
header: "ACME — Q3"
footer: "ACME · 1/2"
date: "19 October 2026"
---
```

## 라이선스 {#licences}

HCL 지원에는 [hashicorp/hcl](https://github.com/hashicorp/hcl)
(MPL-2.0, 수정 없이 라이브러리로 사용)과
[zclconf/go-cty](https://github.com/zclconf/go-cty)(MIT)가 링크됩니다.
