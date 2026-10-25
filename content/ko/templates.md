---
title: "템플릿"
weight: 30
description: "프레젠테이션 안의 Go text/template 치환: 값, 생성된 슬라이드, 페이지 번호, 포함."
tags: [templates]
---

프레젠테이션은 Go [`text/template`](https://pkg.go.dev/text/template)이며,
Markdown을 파싱하기 **전에** 평가됩니다. 액션, 파이프라인, `if`, `range`,
`with`, 변수 등 Go 템플릿 언어의 모든 기능을 프레젠테이션에서 쓸 수 있습니다.

## 데이터 {#data}

템플릿의 점(dot)에는 프런트 매터의 키, 그다음 `data:` 매핑의 키, 그다음
명령줄이나 `slides.yaml`에서 넘긴 데이터가 차례로 담기며, 뒤의 것이 앞의 것을
덮어씁니다.

```markdown
---
marp: true
title: Quarterly review
data:
  team: Platform
---

# {{ .title }}

Presented by the {{ .team }} team.
```

명령줄에서는 다음과 같이 넘깁니다.

| 플래그 | 넘기는 값 |
| --- | --- |
| `--data file.yaml` | 데이터를 담은 YAML 또는 JSON 파일(여러 번 지정 가능) |
| `--set key=value` | 값 하나(여러 번 지정 가능) |

## 슬라이드 생성하기 {#generating-slides}

본문 전체가 하나의 템플릿이므로, `---`를 출력하는 `range`는 슬라이드를 만들어
냅니다.

```markdown
{{ range .releases }}
---

# {{ .version }}

{{ range .changes }}- {{ . }}
{{ end }}
{{ end }}
```

## 페이지 번호와 전체 수 {#page-and-total}

`{{ page }}`와 `{{ total }}`은 슬라이드의 번호와 전체 슬라이드 수입니다. 이 값은
프레젠테이션을 슬라이드로 나눈 뒤에야 알 수 있으므로 그때 채워집니다. 따라서
생성된 프레젠테이션도 자신이 생성한 슬라이드를 정확히 셉니다.

```markdown
<!-- footer: "{{ page }} / {{ total }}" -->
```

## 포함 {#include}

`{{ include "parts/agenda.md" }}`는 프레젠테이션 기준 상대 경로의 다른 파일을
삽입합니다.

## 코드는 건드리지 않음 {#code-is-left-alone}

펜스 코드 블록과 인라인 코드는 템플릿을 실행하기 전에 따로 빼 두었다가 그대로
되돌려 놓습니다. 그래서 Go 템플릿을 다루는 슬라이드에서도 이스케이프 없이
`{{ .Name }}`을 보여 줄 수 있습니다.

## 끄기 {#turning-it-off}

`--no-template`을 쓰면 프레젠테이션을 일반 Marp로 읽으며, `{{`는 텍스트로
취급됩니다.
