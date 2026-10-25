---
title: "Kind"
weight: 40
description: "슬라이드가 자신이 무엇인지 밝힙니다. kind는 로컬 디렉티브 프리셋에 붙인 이름입니다."
tags: [kinds]
---

**kind**(슬라이드 유형)는 헤더, 푸터, 페이지 번호, 클래스, 배경 같은 로컬
디렉티브를 묶어 이름을 붙인 프리셋입니다. 슬라이드는 어떻게 보이는지를
반복하는 대신 자신이 무엇인지(`title`, `section`)를 밝히고, 프레젠테이션의
모양은 한곳에서 바꿉니다.

## kind 사용하기 {#using-a-kind}

```markdown
<!-- kind: section -->
# Part two

---

<!-- _kind: blank -->
![bg](photo.jpg)
```

- `kind:`는 로컬 디렉티브처럼 적용됩니다. 즉 이 슬라이드와 그 뒤의 슬라이드에 적용됩니다.
- `_kind:`는 스폿 디렉티브처럼 이 슬라이드에만 적용됩니다.
- kind를 바꾸면 먼저 이전 kind가 설정한 내용을 **지우므로**, 클래스나 푸터가
  한 kind에서 다음 kind로 새어 나가지 않습니다.
- 슬라이드에 직접 쓴 디렉티브가 kind보다 **우선합니다**.

## 기본 제공 kind {#built-in}

| Kind | 설정 내용 |
| --- | --- |
| `title` | `class: lead`, 페이지 번호·헤더·푸터 없음 |
| `section` | `class: lead` |
| `content` | 없음: 프레젠테이션 자체의 기본값 |
| `blank` | 페이지 번호·헤더·푸터 없음 |

모두 다시 정의할 수 있습니다.

## kind 정의하기 {#defining-kinds}

프런트 매터에 정의합니다.

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

`extends`는 다른 kind의 디렉티브에서 출발해 kind를 만듭니다. kind에는 로컬
디렉티브와 `extends`를 쓸 수 있습니다.

## kind 공유하기 {#sharing-kinds}

팀에서 쓰는 kind는 모든 프레젠테이션이 참조하는 파일 하나에 두는 것이 좋습니다.

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

명령줄의 `--kinds kinds.yaml`이나 `slides.yaml`의 `kinds:`로도 지정할 수
있습니다. 적용 순서는 기본 제공 kind, 명령줄과 설정 파일, 그리고 프레젠테이션
자체의 `kinds:` 순입니다.

`slides kinds talk.md`는 프레젠테이션에서 쓸 수 있는 kind와 각 kind가 설정하는
내용을 나열합니다.

## go-slides 없이 쓰기 {#leaving-go-slides}

`slides expand`는 슬라이드마다 모든 kind를 그에 해당하는 스폿 디렉티브로 풀어
쓰므로, 결과는 어떤 Marp 도구로도 읽을 수 있는 일반 Marp Markdown이 됩니다.
