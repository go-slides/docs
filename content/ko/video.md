---
title: "비디오"
weight: 39
description: "프레젠테이션 속 비디오: HTML 페이지에서는 브라우저가 재생하고, PDF에서는 링크가 달린 포스터 프레임, PowerPoint와 OpenDocument에는 포함됩니다."
tags: [video]
---

비디오는 파일이 `.mp4`, `.m4v` 또는 `.webm`으로 끝나는 이미지입니다:

```markdown
![video](clip.mp4)
![The demo](clip.mp4#t=12.5)
![demo loop muted](clip.mp4)
```

## 문법 {#syntax}

파일은 프레젠테이션 폴더에서 읽거나 `http` 또는 `https` URL에서 읽습니다. 절대 경로, 폴더를 벗어나는 경로, 그 밖의 스킴은 거부됩니다. `#t=12.5`, 또는 분과 초로 쓴 `#t=1:23`은 포스터 프레임을 고릅니다. 재생 전에, 그리고 비디오를 재생할 수 없는 출력에서 보여 주는 그림입니다. 미디어 프래그먼트 형식 `#t=10,20`도 받아들이며 그 시작 값을 사용합니다. 없으면 포스터는 첫 프레임입니다.

대체 텍스트의 키워드는 Marp의 이미지 키워드처럼 쓰지만 go-slides 고유의 것이며, 재생 방식을 정합니다: `autoplay`, `loop`, `muted`. `autoplay`는 `muted`와 함께일 때만 적용되고, 그렇지 않으면 경고와 함께 무시됩니다. 대체 텍스트의 나머지 단어는 접근성 레이블입니다.

## 출력별로 보이는 것 {#outputs}

- **HTML** (페이지와 playground 미리보기): `<video controls playsinline preload="metadata" poster=...>` 요소. 포스터는 디코딩한 프레임이며 JPEG 데이터 URI로 포함됩니다. 비디오는 스스로 소리를 내며 재생되지 않습니다. 브라우저가 움직임 감소를 요청하면(`prefers-reduced-motion`) `autoplay`와 `loop`가 제거되고 비디오는 일시정지됩니다. 자동 재생되는(음소거) 비디오는 절반 이상이 보이는 동안에만 재생되고, 슬라이드가 표시되지 않으면 일시정지됩니다.
- **PDF** (`--pdf`): 포스터 프레임이 그림으로 들어가며 비디오 파일로의 링크가 붙습니다. 링크는 프레젠테이션에 쓴 상대 경로이거나, `http` 또는 `https` 소스라면 URL입니다. 링크를 열어야만 재생됩니다. PDF에는 `RichMedia` 주석이 없습니다. 그것은 Acrobat에서만 재생되므로 일부러 사용하지 않습니다.
- **PowerPoint** (`--pptx`): 비디오는 미디어 파트(`ppt/media/slidesvideoN.mp4|webm`)로 포스터와 함께 포함되며, `a:videoFile`에 `p14:media` 확장을 더한 형태입니다. Keynote: 동영상으로 가져와 재생합니다(확인함). LibreOffice Impress: PowerPoint와 OpenDocument 파일을 비디오로 가져와 디코딩합니다(확인함). PowerPoint: 아직 확인하지 않았습니다. `autoplay`, `loop`, `muted` 키워드는 PowerPoint와 OpenDocument로 옮겨지지 않습니다.
- **OpenDocument** (`--odp`): 비디오는 `Video/` 아래에 `draw:plugin`(ODF 1.3 파트 3, 10.4.8)으로, 포스터 `draw:image` 앞에 포함되고 매니페스트에 나열됩니다. LibreOffice Impress는 이를 비디오로 가져와 디코딩합니다(확인함). 슬라이드 쇼 재생은 미디어 백엔드와 코덱에 따라 달라집니다.
- **URL 소스**: HTML 페이지는 URL에서 비디오를 스트리밍합니다. 내보내기에서는 URL로 연결되는 중립 카드가 보입니다. 빌드할 때 아무것도 가져오지 않으며, PowerPoint나 OpenDocument에도 아무것도 포함되지 않습니다.

## 컨테이너와 포스터 {#codecs}

MP4(ISO BMFF)와 WebM(Matroska) 파일을 Go로 읽으며, ffmpeg는 실행하지 않습니다. 포스터는 H.264(`avc1` 또는 `avc3`)에 대해서만 디코딩되며, 이는 go-avkit의 유일한 비디오 디코더입니다. 다른 코덱(H.265, VP8, VP9, AV1)도 HTML에서 재생되고 PowerPoint와 OpenDocument에 포함되지만, 포스터는 재생 삼각형이 있는 중립 어두운 카드가 되고 경고 `video.codec`이 출력됩니다.

## 한도 {#limits}

비디오는 신뢰할 수 없는 입력이므로 한도가 있습니다: 파일(64 MiB), 박스 또는 요소(262,144개)와 그 중첩(16단계), 프레임(4096 x 4096), 포스터 디코딩(이전 키프레임부터 최대 600개 샘플, 샘플당 8 MiB). 키프레임에서 너무 먼 포스터는 카드와 경고 `video.limit`을 받습니다.

거부된 비디오는 슬라이드에서 `[video not shown: ...]` 노트로 대체됩니다. 빌드는 계속되고, 문제는 코드가 붙은 경고가 됩니다:

- `video.syntax`: MP4 또는 WebM 파일이 아니거나, 비디오 트랙이 없거나, 시간이 잘못됨;
- `video.limit`: 위의 한도를 초과함;
- `video.codec`: 이 코덱의 포스터를 디코딩할 수 없음;
- `video.path`: 절대 경로, 프레젠테이션 폴더를 벗어나는 경로, 또는 `http`와 `https` 이외의 스킴;
- `file.missing`: 파일이 없음.

## playground에서 {#playground}

릴리스되면 다른 파일처럼 프로젝트에 비디오를 추가할 수 있습니다. 비디오의 바이트는 blob URL로 브라우저 안에 머물며 업로드되지 않습니다. [playground](https://go-slides.github.io/playground/)에서 사용해 보세요.

비디오 지원은 필요할 때 로드됩니다. 포스터 프레임을 디코딩하는 리더(압축 시 약 3 MiB)는 별도의 모듈이며, 슬라이드가 처음 포스터를 필요로 할 때 내려받습니다. 비디오가 없는 슬라이드는 이를 전혀 내려받지 않으므로 playground의 첫 로딩이 그만큼 가벼워집니다. Web Worker에서 실행되어 입력이 막히지 않으며, 파일은 `slides`와 같은 제한(64 MiB, 박스와 프레임 크기)과 같은 `video.limit`, `video.syntax` 오류로 읽습니다.
