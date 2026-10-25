---
title: "갤러리"
weight: 12
description: "각각 한 가지를 잘 보여 주는 작은 프레젠테이션 12개: 테마, kind, 템플릿, HCL, 수식, 이미지, 코드, 이모지, 전환 효과, 그리고 완성된 발표 하나. 어느 것이든 Playground에서 열 수 있습니다."
tags: [gallery, examples]
---

go-slides의
[`examples/`](https://github.com/go-slides/slides/tree/main/examples) 폴더에 있는
작은 프레젠테이션 12개로, 각각 한 가지를 잘 보여 줍니다. 미리 보기 이미지는 각
프레젠테이션의 첫 슬라이드이며, 이 사이트를 빌드할 때 go-slides가 직접
렌더링합니다. CI가 모든 프레젠테이션을 HTML과 PDF로 빌드하므로 어느 것도 깨지지
않습니다.

**Playground에서 열기**를 누르면 브라우저에서 프레젠테이션을 불러오며, 이 링크는
어떤 서버에도 데이터를 보내지 않습니다. `deck.md` 외에 다른 파일(kind 파일,
포함된 부분, 이미지)이 있는 프레젠테이션은 링크에 담을 수 없습니다. 이 경우
**.zip**을 내려받아 Playground의 파일 패널로 가져오거나, 압축을 풀고
`slides build --allow-local-files deck.md`로 빌드하십시오.

그림은 go-slides 작성자들이 직접 만든 SVG로
[CC0 1.0](https://creativecommons.org/publicdomain/zero/1.0/)으로 공개되어 있으며,
이를 사용하는 프레젠테이션에는 그 사실이 적혀 있습니다.

{{< gallery >}}
