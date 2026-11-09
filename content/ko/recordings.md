---
title: "터미널 녹화"
weight: 38
description: "프레젠테이션 속 asciinema 녹화: HTML 페이지의 작은 플레이어에서 재생하고, PDF, PowerPoint, OpenDocument에서는 벡터 텍스트로 그립니다."
tags: [asciinema, terminal]
---

{{< callout type="info" >}}
터미널 녹화는 v0.18.0 이후 버전에서 제공됩니다.
{{< /callout >}}

[asciinema](https://docs.asciinema.org/) 녹화([asciicast v2](https://docs.asciinema.org/manual/asciicast/v2/) 또는 [v3](https://docs.asciinema.org/manual/asciicast/v3/))는 파일이 `.cast`로 끝나는 이미지입니다.

```markdown
![cast](demo.cast)
![The failing build](demo.cast#t=12.5)
```

## 구문 {#syntax}

파일은 프레젠테이션 폴더에서 읽습니다. URL, 절대 경로, 폴더 밖으로 나가는 경로는 거부됩니다. 대체 텍스트는 화면 낭독기를 위해 녹화의 이름을 붙입니다. `cast`만 쓰면 녹화 자체의 제목을 뜻합니다. `#t=12.5`(분과 초로는 `#t=1:23`)는 움직일 수 없는 출력이 보여 줄 순간을 고릅니다. 생략하면 녹화의 끝입니다.

## 출력별로 보이는 것 {#outputs}

- **HTML**(페이지와 playground 미리 보기): 녹화는 페이지에 내장된 go-slides의 작은 플레이어에서 재생됩니다. 재생과 일시 정지, 속도 버튼, 위치 막대가 있습니다. 어디서도 스크립트를 불러오지 않으며 네트워크를 쓰지 않습니다.
- **PDF, PowerPoint, OpenDocument**(`--pdf`, `--pptx`, `--odp`): `#t=`가 가리키는 순간의 터미널을 테마의 고정폭 글꼴로 된 벡터 텍스트로 그립니다. 16색, 256색, 직접 색, 굵게, 기울임, 밑줄, 반전 텍스트를 지원합니다. 텍스트는 선택하고 검색할 수 있으며 결코 그림이 아닙니다. 굵게와 기울임은 테마의 고정폭 글꼴에 있을 때 그려집니다(gaia 테마의 Roboto Mono).

## 플레이어 {#player}

녹화가 처음 화면에 나타날 때 한 번 재생합니다. 브라우저가 움직임 줄이기를 요청하면(`prefers-reduced-motion`) 일시 정지 상태로 시작해 `#t=`가 가리키는 순간을 보여 주고, 재생을 누르면 재생합니다. 플레이어의 버튼과 키는 슬라이드를 넘기지 않습니다.

## 제한 {#limits}

녹화는 신뢰할 수 없는 입력이므로 제한이 있습니다. 파일(16 MiB), 이벤트(500,000개), 길이(3시간), 터미널(400×200 셀), 에뮬레이터가 할 수 있는 작업량입니다. 제한을 넘거나 올바른 asciicast 파일이 아닌 녹화는 거부됩니다. 빌드는 계속되고, 슬라이드에는 그 자리에 안내가 표시되며, 문제는 코드가 붙은 경고가 됩니다: `cast.syntax`, `cast.limit`, `cast.version`, `cast.path`.

## playground에서 {#playground}

파일 패널에서 끌어 놓거나 파일 선택기로 프로젝트에 `.cast` 파일을 추가합니다. 파일은 `media/`에 들어가고 이미지로 삽입됩니다. 프레젠테이션이 읽는 텍스트이므로 프로젝트의 다른 텍스트 파일과 함께 계산되며, 그 한도는 2 MiB입니다. [playground](https://go-slides.github.io/playground/)에서 사용해 보세요.
