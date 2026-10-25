---
title: "이모지"
weight: 36
description: "이모지 단축 코드와 Unicode 이모지가 네트워크 없이 Twemoji 이미지로 바뀝니다. HTML 페이지, 이미지, PDF, PowerPoint, OpenDocument 모두에 적용됩니다."
tags: [emoji]
---

{{< callout type="warning" >}}
이모지는 v0.4.0에 추가되었습니다. v0.3.0에는 없었습니다.
{{< /callout >}}

이모지는 단축 코드로 쓰거나 문자 자체로 쓸 수 있습니다.

```markdown
# Launch :rocket: 🎉

Shortcodes :smile: :+1: :heart_eyes: and Unicode 😄 👍🏽 👩‍💻 🇫🇷 1️⃣.
```

각 이모지는 Marp Core가 그리는 [Twemoji](https://github.com/jdecked/twemoji)
이미지가 되며, 마크업도 Marp Core가 쓰는 것과 같습니다.

```html
<img class="emoji" draggable="false" alt="😄" src="data:image/svg+xml;base64,…" data-marp-twemoji=""/>
```

## 네트워크 불필요 {#no-network}

Marp Core의 `src`는 CDN의 URL이므로, 오프라인에서 보여 주거나 인쇄하거나
인터넷 연결이 없는 컴퓨터에서 PDF로 만든 프레젠테이션에는 깨진 이미지가
나타납니다. go-slides는 SVG 자체를 `data:` URI에 넣습니다. 그래픽 4,009개
(Twemoji 17.0.3, Emoji 17.0)는 순수 Go 모듈
[go-icons/twemoji](https://github.com/go-icons/twemoji)를 통해 프로그램에
내장되어 있으며, 압축 시 2 MB이고 프레젠테이션에 이모지가 있을 때만 읽습니다.

이미지는 HTML 페이지, `--images`/`--image`, `--pdf`, `--pptx`, `--odp` 모두에서
똑같이 그려지며, SVG 래스터화는 go-webengine이 맡습니다. 그래픽 4,009개 모두가
이 SVG 래스터라이저에서 파싱되고 그려지는지 테스트로 확인합니다. v0.5.0부터
래스터라이저는 [go-gfx/oksvg](https://github.com/go-gfx/oksvg) v0.2.0입니다.
이제 `clipPath`, `fill-rule: evenodd`, 원호 명령이 Chrome과 똑같이 그려지며,
이로 인해 Twemoji 하나와 그 밖의 211개가 달라졌고 SVG 로고도 더 이상 깨지지
않습니다.

이모지 크기를 텍스트에 맞추는 규칙(`img[data-marp-twemoji]`, 한 변이 `em`
하나인 정사각형, 기준선에 정렬)은 Marp Core와 같은 방식으로 테마에 추가되므로,
직접 만든 테마에도 적용됩니다.

## 변환 대상 {#what-is-converted}

규칙은 Marp Core의 것을 따르며, 모든 경우를 담은 프레젠테이션에 대해 Marp Core
4.4.0이 출력하는 HTML과 비교하는 테스트가 있습니다.

| | |
| --- | --- |
| `:smile:`, `:+1:`, `:heart_eyes:` | markdown-it-emoji의 이름 1,903개. 본문, 제목, 링크, 표, `header`, `footer`에서 변환되며 코드, URL, 이미지 대체 텍스트에서는 변환되지 않습니다 |
| `😄`, `👩‍💻`, `🇫🇷`, `1️⃣`, `👍🏽` | Unicode 이모지. 본문 **그리고 코드에서도** 변환됩니다. Marp Core가 렌더링된 코드 HTML을 다시 쓰기 때문입니다 |
| `✔`, `©`, `™`, `▶` | 텍스트로 남습니다. 기본 표시 형태가 텍스트인 문자는 U+FE0F가 붙을 때만(`✔️`) 이모지가 됩니다. `❤`는 예외입니다 |
| `☺︎` | 텍스트로 남습니다. 이모지 뒤의 U+FE0E는 일반 글리프를 요청합니다 |
| `12:30:45`, `:not_a_name:` | 쓴 그대로 남습니다 |

제목의 `id`와 개요에는 단축 코드 이름이 아니라 이모지가 쓰입니다.
`# Launch :rocket:`의 id는 Marp와 마찬가지로 `launch-%F0%9F%9A%80`입니다.

go-slides가 Marp Core와 다른 점은 더 많이 변환한다는 것입니다. Marp Core의
emoji-regex가 텍스트로 남겨 두는 `✌️ ☝️ ✍️ 🏋️` 등 기본 표시가 텍스트인 몇몇
문자와, Unicode 목록에는 없지만 Twemoji가 그리는 몇 가지 시퀀스도 변환합니다.

## 끄기 {#turning-it-off}

`--no-emoji`(코드에서는 `slides.Config{NoEmoji: true}`)를 쓰면 단축 코드와
Unicode 이모지가 쓴 그대로의 텍스트로 남습니다. Marp Core의
`emoji: false`에 해당합니다.

## 저작자 표시 {#attribution}

그래픽은 Twemoji이며 Copyright 2020 Twitter, Inc and other contributors,
[CC-BY 4.0](https://creativecommons.org/licenses/by/4.0/) 라이선스로 배포되고
[jdecked/twemoji](https://github.com/jdecked/twemoji) 프로젝트가 이어 가고
있습니다. 이를 표시하는 프레젠테이션은 푸터나 마지막 슬라이드 등에 출처를
밝혀야 합니다. 단축 코드 표는 markdown-it-emoji(MIT)의 것입니다.
