---
title: "명령줄"
weight: 50
description: "slides의 모든 명령과 플래그, 그리고 slides.yaml 설정 파일."
tags: [cli]
---

{{< callout type="warning" >}}
이 페이지는 v0.14.0을 기준으로 설명합니다.
{{< /callout >}}

```text
slides [deck.md] [flags]
slides [command] [flags] deck.md
```

## 명령 {#commands}

| 명령 | 하는 일 |
| --- | --- |
| `build` | 프레젠테이션을 HTML, PDF, 이미지로 씁니다(아래 참고). **기본 명령**이므로 `slides deck.md`는 `slides build deck.md`와 같습니다 |
| `expand` | 프레젠테이션을 일반 Marp Markdown으로 출력합니다. 템플릿은 펼쳐지고, kind(슬라이드 유형)는 그에 해당하는 디렉티브로 바뀝니다 |
| `notes` | 발표자 노트를 출력합니다 |
| `kinds` | 프레젠테이션에서 쓸 수 있는 kind와 각각이 설정하는 내용을 나열합니다. `--json`(v0.6.0부터)을 쓰면 `extends`를 해석한 상태로 `[{"name", "extends", "directives": [{"key", "value"}]}]` 형식으로 씁니다 |
| `themes` | 테마를 나열합니다 |
| `version` | 버전을 출력합니다(`slides --version`도 같음) |
| `completion` | 셸 자동 완성 스크립트를 출력합니다: `bash`, `zsh`, `fish`, `powershell` |

`slides <command> --help`는 명령의 플래그를 출력합니다. 플래그는 프레젠테이션
앞에 와도, 뒤에 와도 됩니다. `slides deck.md -o out.html`과 `slides -o out.html
deck.md`는 같습니다. 긴 플래그는 대시 두 개(`--html`)로 쓰며, 짧은 플래그는
`-o` 하나뿐입니다. `build`는 `-o`로 따로 지정하지 않으면 `deck.md` 옆에
`deck.html`을 씁니다.

## 플래그 {#flags}

| 플래그 | 명령 | 하는 일 |
| --- | --- | --- |
| `-o`, `--output file` | `build`, `expand`, `notes` | 출력 파일. 표준 출력은 `-` |
| `-I`, `--input-dir dir` | `build` | `dir` 아래의 Markdown 파일을 재귀적으로 변환합니다(v0.6.0부터). 아래 참고 |
| `-P`, `--parallel n` | `build` | 동시에 빌드할 프레젠테이션 수(v0.6.0부터). 기본값 5, 최대 코어 수이며, 1이면 하나씩 차례로 빌드합니다 |
| `-w`, `--watch` | `build` | 프레젠테이션, 테마, kind 파일이나 데이터 파일, 설정이 바뀌면 다시 빌드하고 열린 페이지를 새로 고칩니다(v0.6.0부터). 아래 참고 |
| `-s`, `--server dir` | `build` | 디렉터리를 제공하며, 요청이 오면 그 안의 프레젠테이션을 변환합니다(v0.6.0부터). 아래 참고 |
| `-p`, `--preview` | `build` | 결과를 기본 브라우저로 엽니다(v0.6.0부터). `--watch`(자동 적용)와 함께면 HTML 페이지를, `--server`와 함께면 서버를 엽니다 |
| `--host addr` | `build` | `--server`가 수신할 주소. 기본값은 `127.0.0.1`입니다(v0.6.0부터) |
| `--allowed-host name` | `build` | `--server`가 응답할 리버스 프록시의 공개 호스트 이름. 여러 번 지정할 수 있습니다(v0.14.0부터). [셀프 호스팅]({{< relref "self-hosting" >}}) 참고 |
| `--template bespoke\|bare` | `build` | HTML 페이지 템플릿(v0.6.0부터). 기본값인 `bespoke`는 발표 스크립트, 화면 컨트롤러, 전환 효과를 포함하고, `bare`는 슬라이드를 쌓아 놓습니다(`--bare`) |
| `--bespoke.osc`, `--bespoke.progress`, `--bespoke.transition` | `build` | bespoke 템플릿의 옵션으로, 기본값은 marp-cli와 같습니다(v0.6.0부터). 컨트롤러 켜짐, **진행률 막대 꺼짐**(이전에는 항상 그려졌음), 전환 효과 켜짐. `--bespoke.osc=false`처럼 하나씩 끌 수 있습니다 |
| `-c`, `--config-file file`, `--no-config-file` | `build` | marp-cli의 설정 파일(JSON 또는 YAML, v0.6.0부터). 아래 참고. `--config`는 `slides` 자체의 설정 파일입니다 |
| `--browser`, `--browser-path`, `--browser-protocol`, `--browser-timeout` | `build` | 받아들이되 경고와 함께 무시합니다(v0.6.0부터). 브라우저가 없기 때문입니다 |
| `--pptx-editable`, `--no-parallel` | `build` | `--pptx`와 `-P 1`의 marp-cli식 표기(v0.6.0부터) |
| `--config file` | 모두 | 설정 파일. 기본값은 프레젠테이션 옆에 있는 `slides.yaml`(있을 경우) |
| `--no-config` | 모두 | 설정 파일을 읽지 않습니다 |
| `--theme-set path` | 모두 | 테마 CSS 파일 또는 그런 파일이 든 디렉터리(여러 번 지정 가능) |
| `--kinds file` | 모두 | kind를 정의한 YAML 파일(여러 번 지정 가능) |
| `--data file` | 모두 | 템플릿 데이터를 담은 YAML 또는 JSON 파일(여러 번 지정 가능) |
| `--set key=value` | 모두 | 템플릿 값 하나(여러 번 지정 가능) |
| `--no-template` | 모두 | Go 템플릿을 펼치지 않습니다 |
| `--strict-yaml` | 모두 | 순수 Marpit처럼 디렉티브를 엄격한 YAML로 파싱합니다(Marp Core는 느슨하게 파싱합니다) |
| `--no-emoji` | 모두 | `:shortcodes:`와 Unicode 이모지를 이미지로 그리지 않고 텍스트로 둡니다(v0.4.0부터: [이모지]({{< relref "emoji" >}})) |
| `--html` | `build` | 모든 원시 HTML을 허용합니다. 기본값은 Marp Core의 허용 목록입니다 |
| `--lang code` | `build` | 프레젠테이션이 지정하지 않았을 때 쓸 페이지의 `lang` |
| `--bare` | `build` | 발표 스크립트 없이 슬라이드를 쌓아 놓습니다 |
| `--pdf` | `build` | 슬라이드당 한 페이지, 개요가 포함된 PDF를 씁니다. 텍스트와 도형은 벡터입니다(`-o file.pdf`로도 선택됨) |
| `--pdf-raster` | `build` | v0.3.0처럼 각 페이지가 슬라이드 이미지 한 장인 PDF를 씁니다 |
| `--notes` | `build` | marp-cli처럼 슬라이드를 `---`로 구분한 텍스트 파일로 발표자 노트를 씁니다(v0.6.0부터, `-o file.txt`로도 선택됨). `notes` 명령은 다른 형식으로 출력합니다 |
| `--pdf-notes` | `build` | marp-cli처럼 각 슬라이드의 발표자 노트를 PDF 페이지에 노트 주석으로 추가합니다(v0.6.0부터). 출력 형식을 따로 지정하지 않으면 PDF를 만듭니다 |
| `--pdf-outlines`, `--pdf-outlines.pages`, `--pdf-outlines.headings` | `build` | PDF 책갈피(v0.6.0부터): 슬라이드마다 "Page N"과 그 아래에 중첩된 제목. marp-cli와 달리 여기서는 기본적으로 켜져 있으며, `=false`로 하나씩 끌 수 있습니다. 아래 참고 |
| `--images png\|jpeg` | `build` | 슬라이드마다 이미지 한 장을 씁니다: `deck.001.png`, … |
| `--image png\|jpeg` | `build` | 첫 슬라이드를 이미지로 씁니다(`-o file.png`나 `-o file.jpg`로도 선택됨) |
| `--pptx` | `build` | 발표자 노트를 포함한 PowerPoint 파일을 씁니다(`-o file.pptx`로도 선택됨). v0.5.0부터 편집 가능하며, v0.4.0에서는 슬라이드당 그림 한 장이었습니다 |
| `--pptx-raster` | `build` | marp-cli처럼 슬라이드마다 슬라이드 전체 크기 그림 한 장으로 이루어진 PowerPoint 파일을 씁니다(v0.5.0부터). `--image-scale` 기본값은 2입니다 |
| `--odp` | `build` | 발표자 노트를 포함한 OpenDocument 프레젠테이션을 씁니다(`-o file.odp`로도 선택됨). v0.5.0부터 편집 가능합니다 |
| `--odp-raster` | `build` | 슬라이드당 그림 한 장으로 이루어진 OpenDocument 프레젠테이션을 씁니다(v0.5.0부터) |
| `--jpeg-quality n` | `build` | JPEG 품질, 1~100(기본값 85) |
| `--image-scale n` | `build` | 이미지와 PDF에서 CSS 픽셀당 장치 픽셀 수(기본값 1). 2이면 16:9에서 2560×1440이 되며, 텍스트와 벡터는 확대되지 않고 그 밀도로 그려집니다 |
| `--allow-local-files` | `build` | 이미지와 PDF가 로컬 파일을 읽도록 허용하되, 프레젠테이션 디렉터리 아래의 파일만 허용합니다. marp-cli와 마찬가지로 기본값은 거부입니다 |
| `--theme name\|file.css` | `build` | 프레젠테이션의 테마를 재정의합니다(v0.6.0부터). 테마 이름이나 CSS 파일을 지정하며, CSS 파일은 어떤 `@theme`을 선언하든 그 테마가 됩니다. 디렉터리는 거부되므로 `--theme-set`을 쓰십시오 |
| `--title`, `--description`, `--author`, `--keywords`, `--url`, `--og-image` | `build` | 페이지 메타데이터(v0.6.0부터). 프레젠테이션의 `title`, `description`, `author`, `keywords`, `url`, `image` 디렉티브보다 우선합니다. 아래 참고 |

## 여러 프레젠테이션과 디렉터리 {#several-decks-and-a-directory}

v0.6.0부터 `build`는 여러 프레젠테이션, 디렉터리(그 아래에서 재귀적으로 찾은
프레젠테이션), 패턴을 받을 수 있으며, 최대 `--parallel`개를 동시에 빌드합니다.
`-`는 표준 입력입니다. marp-cli와 달리 `-`로 지정하지 않으면 파이프를 읽지
않습니다.

패턴은 marp-cli의 glob과 같은 방식으로 해석됩니다(`**`와 `{a,b}`는 v0.7.0부터.
셸이 펼치지 않고 `slides`에 넘기도록 따옴표로 감싸십시오). 이름 안에서는 `*`,
`?`, `[a-z]`, `{md,txt}`를, 디렉터리 여러 단계에는 `**`를 쓰므로
`'docs/**/*.md'`는 트리 안의 모든 프레젠테이션을 가리킵니다. 와일드카드는 점으로
시작하는 이름과 일치하지 않고, `node_modules` 안은 찾지 않으며, 대소문자를
구분하고(`*.md`는 `UPPER.MD`를 찾지 않음), 패턴과 일치하는 것은 확장자와
관계없이 변환됩니다. 아무것도 일치하지 않는 패턴은 오류입니다.

- **`-I dir`**는 marp-cli처럼 `dir` 아래의 Markdown 파일(소문자 `.md`, `.mdown`,
  `.markdown`, `.markdn`)을 찾으며, `node_modules`와 점으로 시작하는
  디렉터리는 찾지 않습니다. `-o out`을 지정하면 각 프레젠테이션은 `dir` 아래의
  경로를 `out` 아래에 그대로 옮긴 위치에 쓰입니다. `dir/sub/b.md`는
  `out/sub/b.html`(플래그에 따라 `b.pdf`, `b.001.png`, `b.txt`, …)이 됩니다.
  `-o`가 없으면 출력은 각 프레젠테이션 옆에 쓰입니다. `-I`와 함께
  프레젠테이션을 지정하면 오류이며, 여러 프레젠테이션에 `-o file`을 지정해도
  오류입니다(`-o -`는 허용).
- **`-P n`**은 프레젠테이션 `n`개를 동시에 빌드합니다. 프레젠테이션 하나의
  렌더링도 이미 여러 코어를 쓰므로 `n`은 코어 수로 제한되며, `-o -`이면 하나씩
  차례로 빌드합니다. 한 프레젠테이션의 메시지는 함께 묶여 출력되고, 실패한
  프레젠테이션이 있어도 나머지는 멈추지 않습니다. 종료 상태는 1이며 메시지에
  실패 목록이 나옵니다.

## 전환 효과와 발표 페이지 {#transitions-and-the-presenting-page}

v0.6.0부터 marp-cli의 `transition` 디렉티브가 동작합니다. `<!-- transition:
fade 1s -->`처럼 이름과 선택적 지속 시간을 쓰면 해당 슬라이드와 그 뒤의
슬라이드에 적용되고, `<!-- _transition: cover -->`는 한 슬라이드에만
적용됩니다. 기본 제공 이름은 marp-cli와 같습니다:
`clockwise`, `counterclockwise`, `cover`, `coverflow`, `cube`, `cylinder`,
`diamond`, `drop`, `explode`, `fade`, `fade-out`, `fall`, `flip`, `glow`,
`implode`, `in-out`, `iris-in`, `iris-out`, `melt`, `overlap`, `pivot`, `pull`,
`push`, `reveal`, `rotate`, `slide`, `star`, `swap`, `swipe`, `swoosh`, `wipe`,
`wiper`, `zoom`. `none`은 전환 효과를 끕니다. 페이지는 보는 사람이 전환 효과가
지정된 슬라이드에서 앞으로 넘어갈 때와 그 슬라이드로 돌아올 때 브라우저의
**View Transition API**로 효과를 재생합니다. 이 API가 없는 브라우저는 이전처럼
슬라이드를 바꾸고, 동작 줄이기를 요청한 사용자에게는 단순한 페이드가 적용됩니다.
테마는 `marp-transition-NAME`(또는 `marp-outgoing-`과 `marp-incoming-`, 그리고
`-backward`)이라는 이름으로 자체 키프레임을 제공할 수 있으며, 이 키프레임이 기본
제공 키프레임보다 먼저 쓰입니다.

페이지에는 이전, 페이지, 다음, 전체 화면 버튼으로 이루어진 **화면 컨트롤러**가
있습니다. 포인터를 움직이거나 키를 누르면 나타났다가 잠시 뒤 사라지며
(`--bespoke.osc=false`로 제거), `--bespoke.progress`를 쓰면 진행률 막대도
표시됩니다. **이제 진행률 막대는 marp-cli처럼 기본적으로 꺼져 있습니다.**
이전에는 항상 그려졌습니다.

### 발표자 보기 {#presenter-view}

v0.7.0부터 페이지에 marp-cli의 발표자 보기가 있습니다. `p`를 누르거나 컨트롤러의
마지막 버튼을 누르면 현재 슬라이드, 그다음 단계, 발표자 노트, 페이지, 시각,
타이머가 있는 창이 열립니다(타이머를 클릭하면 초기화되고, `+`와 `-`로 노트 크기를
조절하며, 슬라이드와 노트 사이의 막대는 끌어서 움직일 수 있습니다). 두 창은
양방향으로 연동됩니다. 어느 창에서든 화살표 키, 다음 단계 클릭, Previous/Next
클릭으로 다른 창도 한 단계(프래그먼트)씩 움직입니다.

marp-cli 4.5.1 소스에서 확인한 것과 같은 방식으로 동작합니다. 창은 같은 파일의
`?view=presenter&sync=KEY`이고, 다음 단계는 같은 페이지를 `?view=next`로
프레임에 띄운 것이며, 두 창은 `localStorage`(`bespoke-marp-sync-KEY`)로 위치를
공유합니다. 따라서 서버도 네트워크도 필요 없지만, `localStorage`가 없는
브라우저에서는 동작하지 않습니다(버튼이 이를 알려 줍니다). 노트는 텍스트이며
`&`, `<`, `>`가 이스케이프되므로 노트가 마크업이 될 수 없습니다. 노트는 발표자
보기가 가져갈 수 있도록 페이지 안에 숨겨져 있어, 화면 낭독기가 청중에게 읽어
주는 내용에는 포함되지 않습니다. `--bare` 페이지에는 이 기능이 없습니다. 발표자
보기는 페이지 크기를 약 10 KB 늘립니다.

### 개요 보기 {#overview-view}

v0.8.0부터 `o`(또는 `Escape`. marp-cli는 두 키 모두로 전환합니다), 컨트롤러의
격자 버튼, 발표자 보기의 페이지 카운터를 누르면 모든 슬라이드가 페이지 위의
격자에 온전히 표시됩니다. 화살표 키로 슬라이드를 고르고(위아래는 행 단위),
`Home`과 `End`로 처음과 끝으로 가며, 다른 창도 따라 움직입니다. `Enter`,
`Space`, 클릭으로 슬라이드를 선택하면 페이지 위에서는 개요가 닫힙니다. 발표자
보기 위에서는 선택해도 개요가 열린 채로 남습니다(발표자가 보여 줄 슬라이드를
고른 뒤에도 계속 살펴볼 수 있도록). "Close overview" 버튼, `o`, `Escape`로
닫습니다. 이 기능도 marp-cli와 같습니다. 같은 페이지를 다시
`?view=overview&sync=KEY` 프레임에 띄우고 `localStorage`로 위치를 공유하며,
네트워크를 쓰지 않습니다. 격자 안의 슬라이드는 포커스를 받을 수 있고 "Slide N"
레이블이 붙으며, 닫힌 개요는 비활성(inert) 상태이므로 키보드나 화면 낭독기
사용자가 그 안으로 들어가는 일은 없습니다. 두 보기를 합치면 페이지 크기가 약
15 KB 늘어납니다.

### 설정 파일 {#configuration-files}

marp-cli는 작업 디렉터리의 파일에서 옵션을 읽으며, `slides`도 그 파일이 데이터일
때 똑같이 읽습니다: `.marprc`(YAML. JSON도 YAML입니다), `.marprc.json`,
`.marprc.yaml`, `.marprc.yml`, `.config/marprc*`, 또는 `package.json`의 `marp`
키. `--config-file`(`-c`)로 임의의 JSON 또는 YAML 파일(예: `marp.config.json`)을
지정할 수 있습니다. marp-cli의 검색과 마찬가지로 작업 디렉터리만 찾고 상위
디렉터리는 찾지 않습니다. JavaScript나 TypeScript 파일(`marp.config.js` ...)은
프로그램이므로 `slides`는 실행하지 않으며, 그 사실을 알리고 계속 진행합니다.
옵션은 marp-cli 문서에 있는 것 중 여기서 의미가 있는 것들입니다:
`title`, `description`, `author`, `keywords`, `url`, `ogImage`, `lang`,
`theme`, `themeSet`, `html`, `allowLocalFiles`, `template`, `bespoke`,
`inputDir`, `output`, `parallel`, `pdf`, `pptx`, `notes`, `pdfNotes`,
`pdfOutlines`, `images`, `image`, `imageScale`, `jpegQuality`, `watch`,
`server`, `preview`. 경로는 파일 기준 상대 경로이며, 명령줄이 우선합니다.
`engine`, `options`, `browser*`는 Marp 엔진과 브라우저가 필요하므로
`--browser*` 플래그와 마찬가지로 경고와 함께 무시됩니다. 따라서 marp-cli용으로
작성한 스크립트도 계속 동작합니다. (`slides.yaml`과 `slides.hcl`은 kind,
템플릿 데이터와 함께 별개이며 이전과 똑같이 읽습니다.)

## 감시, 서버, 미리 보기 {#watch-serve-and-preview}

v0.6.0부터 marp-cli의 경로와 실시간 새로 고침 프로토콜을 지원합니다(marp-cli
4.5.1 소스에서 확인하고 실행 중인 `marp -s`와 비교했습니다).

**`slides --server dir`**(`-s`)는 디렉터리를 `http://127.0.0.1:8080/`에서
제공합니다(`PORT` 환경 변수로 포트를 바꿀 수 있으며, 이미 사용 중인 포트는 그
사실을 알리는 오류가 됩니다).

| 요청 | 응답 |
| --- | --- |
| `/deck.md` | 요청 시 변환한 HTML 프레젠테이션 |
| `/deck.md?pdf`, `?pptx`, `?png`, `?jpg` (또는 `?jpeg`), `?notes` (또는 `?txt`) | 쿼리가 요청한 형식으로 변환한 프레젠테이션. PPTX는 첨부 파일로 제공됩니다 |
| 디렉터리 | 그 안의 `index.md`나 `PITCHME.md`를 페이지로, 없으면 프레젠테이션마다 PDF와 PPTX 링크가 있는 색인 |
| 그 밖의 파일 | 파일 그대로 |
| 변환되지 않는 프레젠테이션 | 오류와 함께 503 |

페이지는 `/.__marp-cli-watch-notifier__/<id>`로 WebSocket을 엽니다. id는
프레젠테이션 경로의 SHA-256입니다. 서버는 `ready`를 보낸 다음, 프레젠테이션이나
디렉터리 아래의 무언가가 바뀌면 `reload`를 보냅니다. 끝의 슬래시 없이 요청한
디렉터리는 슬래시가 붙은 경로로 리디렉션됩니다(marp-cli는 그 경로에서 색인을
제공하므로 페이지의 상대 링크가 깨집니다).

**`--watch`**(`-w`)는 프레젠테이션을 빌드한 뒤 바뀐 프레젠테이션을 다시
빌드하고, `--theme-set` 파일, kind 파일이나 데이터 파일, 설정이 바뀌면 전체를 다시
빌드합니다. 이렇게 쓴 페이지는 루프백에 소켓(포트 37717 또는 그다음 빈 포트)을
열고, 해당 프레젠테이션이 다시 빌드되면 새로 고침합니다. 빌드에 실패한
프레젠테이션은 보고된 뒤에도 계속 감시됩니다. **`--preview`**(`-p`)는 기본
브라우저로 서버나 방금 빌드한 HTML 페이지를 엽니다. `slides`에는 구동할
브라우저가 없으므로 자체 창은 없습니다.

서버는 컴퓨터로 들어오는 통로이므로, 모든 인터페이스에서 수신하고 아무것도
검사하지 않는 marp-cli보다 엄격합니다.

- **127.0.0.1**에서 수신합니다. `--host`로 바꿀 수 있지만 경고가 나오며,
  와일드카드 주소에서는 더 이상 `Host` 헤더를 검사할 수 없습니다.
- **`Host`** 헤더는 루프백 이름(`localhost`, `127.0.0.1`, `[::1]`,
  `*.localhost`)이거나 서버가 수신하는 호스트여야 하며, 이것이 DNS 리바인딩을
  막습니다.
- WebSocket의 **`Origin`**은 서버 자신의 출처여야 합니다. 감시 모드의 소켓은
  파일에서 연 페이지의 `null` 출처도 받습니다. 다른 사이트에서 온 요청 중
  탐색이 아닌 것(스크립트, 이미지, `fetch`)은 거부됩니다.
- 경로는 정규화하고 심볼릭 링크를 해석하며, 그 결과가 디렉터리 안에 있어야
  합니다. `..`와 디렉터리 밖을 가리키는 링크는 거부되며, 밖에 있는
  프레젠테이션의 `?pdf`도 거부됩니다. 점으로 시작하는 이름(`.git`, `.env`)은
  제공하지 않습니다.
- `GET`과 `HEAD`만 받습니다. 동시 변환 수는 몇 개로 제한되고, 감시하는
  디렉터리는 최대 4096개이며 `node_modules` 안이나 점으로 시작하는 디렉터리
  아래는 감시하지 않습니다.

리버스 프록시 뒤에서는 `--allowed-host name`(v0.14.0부터, 여러 번 지정 가능)이
프록시의 공개 이름을 허용되는 `Host`와 `Origin`에 추가합니다. 검사는 켜진 채로
유지되고 서버는 루프백에 그대로 머뭅니다.
[셀프 호스팅]({{< relref "self-hosting" >}})은 이를 Caddy 뒤에 두고
HTTPS와 로그인을 함께 구성합니다.

## 페이지 메타데이터 {#page-metadata}

marp-cli와 마찬가지로 명령줄이 프레젠테이션보다 우선합니다. marp-cli에서는
플래그가 마지막 디렉티브 주석 하나로 프레젠테이션 끝에 덧붙여집니다. 규칙은
marp-cli 4.5.1에서 확인했으며 그 출력과 비교하는 테스트로 고정되어 있습니다
(v0.6.0부터).

- `title`을 지정하지 않은 프레젠테이션은 가장 낮은 수준의 첫 제목 텍스트를 쓴
  그대로 가져옵니다(`# **Bold** title`이면 `**Bold** title`). 제목이 없으면
  `<title>`도 없습니다.
- `keywords`는 쉼표로 나누고 앞뒤 공백을 없앤 뒤 중복을 제거합니다. 프런트 매터의
  목록은 그대로 받아들입니다. marp-cli와 마찬가지로 키워드가 두 개 이상일 때만
  페이지에 `keywords` 태그가 들어가며, 하나뿐이면 들어가지 않습니다.
- 유효하지 않은(절대 경로가 아닌) `url`은 경고와 함께 무시되고 이전 값이
  유지됩니다. 유효한 값은 `<link rel="canonical">`과 `og:url`도 씁니다.
- `author`와 `description`에는 대응하는 Open Graph 태그가 함께 붙습니다.
  `og:image:alt`와 큰 Twitter 카드에는 제목과 이미지가 모두 필요합니다.

## 이미지, PDF, PPTX, ODP {#images-pdf-pptx-and-odp}

marp-cli와 마찬가지로 `-o`에 `.pdf`, `.pptx`, `.odp`, `.png`, `.jpg` 확장자를
쓰면 그에 맞는 형식이 선택됩니다. 출력 디렉터리는 자동으로 만들어집니다. PDF는
슬라이드 크기의 페이지가 슬라이드당 하나씩 있고, 슬라이드와 그 제목으로 이루어진
개요가 들어갑니다.

### PDF 노트와 개요 {#pdf-notes-and-outline}

v0.6.0부터 PDF는 marp-cli의 두 옵션을 지원하며, 같은 프레젠테이션에 대해
marp-cli 4.5.1이 쓴 PDF와 비교했습니다.

- `--pdf-notes`는 각 슬라이드의 발표자 노트를 해당 페이지에 노트 주석으로
  넣습니다. 왼쪽 아래에 노란 노트 아이콘이 표시되며, 내용은 노트(주석이 여러
  개면 빈 줄로 연결)이고 제목은 프레젠테이션의 `author`입니다. marp-cli와
  마찬가지로 `--pdf-notes`만 지정해도 PDF를 만듭니다.
- 개요에는 슬라이드마다 굵은 글씨의 "Page N" 책갈피가 있고, 그 아래에 슬라이드의
  제목이 수준별로 중첩됩니다(`--pdf-outlines.pages`와
  `--pdf-outlines.headings`로 각각 선택). **개요는 여기서는 기본적으로 켜져 있고
  marp-cli에서는 꺼져 있습니다.** `slides`는 항상 개요를 써 왔습니다. 몇백
  바이트밖에 들지 않고 긴 프레젠테이션을 쉽게 탐색할 수 있게 해 주므로 기본값을
  유지합니다. `--pdf-outlines=false`나 marp-cli의 `--no-pdf-outlines`로 끄며,
  `.pages`나 `.headings`에 `=false`를 쓰면 해당 부분만 뺍니다. 제목 책갈피는
  해당 슬라이드로 이동합니다. marp-cli는 제목 위치까지 스크롤합니다.

렌더링은 [go-webengine](https://github.com/go-webengine/engine)이 하고 PDF는
[go-pdfkit](https://github.com/go-pdfkit/pdfkit)이 씁니다. 기본 제공 테마가
요구하는 글꼴(Lato, Roboto Mono)은
[go-opentype/fonts](https://github.com/go-opentype/fonts)에서 가져와 내장하므로,
렌더링 중 네트워크에 접근하는 일이 없습니다. Chrome은 쓰지 않습니다.

`--image-scale 2`는 CSS 픽셀당 장치 픽셀 2개로 렌더링합니다. 로컬 이미지는
`--allow-local-files`를 지정하지 않으면 거부되며, 지정하더라도 프레젠테이션
디렉터리 아래의 파일만 읽습니다. 따라서 다른 곳에서 받은 프레젠테이션이
`~/.ssh`에 접근할 수 없습니다. 여러 장의 배경과 화면을 나누는 배경
(`![bg left:40%]`)은 Marp의 레이어와 똑같이 렌더링됩니다.

`--pptx`는 [go-ooxml](https://github.com/go-ooxml)의 `WritePresentation`으로
발표자 노트와 프레젠테이션 메타데이터를 포함해 씁니다. v0.5.0부터 슬라이드는
편집할 수 있습니다(아래 참고). `--pptx-raster`는 marp-cli처럼 슬라이드마다
슬라이드 전체 크기 그림 한 장을 쓰며, 배율 기본값은 2이고 `--image-scale`로
바꿀 수 있습니다. 결과는 LibreOffice에서 열립니다. LibreOffice에서만
테스트했으며 PowerPoint와 Keynote에서는 테스트하지 않았습니다.

`--odp`는 go-odf의 `WritePresentation`으로 발표자 노트를 포함한 OpenDocument
프레젠테이션을 씁니다. `--odp-raster`는 슬라이드당 그림 한 장을 씁니다.

### 편집 가능한 PPTX와 ODP {#editable-pptx-and-odp}

v0.5.0부터 PPTX와 ODP의 슬라이드는 **편집할 수 있습니다**. marp-cli의 PPTX는
슬라이드당 그림 한 장이며, 실험적 기능인 `--pptx-editable`은 PDF를 LibreOffice로
변환합니다. 여기서는 각 슬라이드를 PNG와 PDF 내보내기가 그리는 것과 같은
레이아웃에서 그립니다.

- 텍스트의 각 줄은 레이아웃된 위치에 놓인 텍스트 상자이며, 슬라이드 레이아웃에
  쓰인 글꼴을 씁니다. 글꼴은 **파일에 내장되므로** Lato나 Inter가 없는
  컴퓨터에서도 똑같이 보이며(LibreOffice 25.8로 확인), 굵게, 기울임, 색, 자간,
  클릭 가능한 링크도 유지됩니다.
- 배경, 테두리, 둥근 상자는 도형입니다. 그레이디언트는 PPTX에서는 그레이디언트
  채우기, ODP에서는 그림입니다(ODF 그레이디언트는 색이 두 개뿐입니다).
- 그림은 여러 슬라이드에 나오더라도 한 번만 저장됩니다.
- 슬라이드로 표현할 수 없는 것(`opacity`가 적용된 요소, `filter`, 흐린 그림자,
  수식)은 해당 요소의 그림이 되며, 그 위에 텍스트가 유지됩니다.

LibreOffice로 다시 렌더링하면 슬라이드는 PNG 내보내기와 픽셀 기준 0.04–2 %
이내로 일치합니다. `--pptx-raster`와 `--odp-raster`는 슬라이드당 그림 한 장을
유지합니다.

### 벡터 PDF {#vector-pdf}

v0.4.0부터 PDF는 벡터로 그려집니다. go-webengine이 레이아웃된 슬라이드를 픽셀
대신 PDF에 직접 그립니다. 텍스트는 슬라이드 레이아웃에 쓰인 글꼴의 텍스트로,
`/ToUnicode` 맵과 함께 서브셋으로 내장되므로 선택, 검색, 복사할 수 있습니다.
배경, 테두리, 그레이디언트는 패스와 셰이딩이며, 여러 슬라이드에 나오는 이미지는
한 번만 내장되고 링크는 클릭할 수 있습니다. 슬라이드 100장 벤치마크
프레젠테이션은 758 KiB로, marp-cli의 837 KiB와 v0.3.0 이미지 PDF의 10 MiB보다
작으며 0.3 s 만에 작성됩니다.

여기서 벡터로 표현할 수 없는 요소(`opacity` < 1이 적용된 요소, `filter`나
회전이 있는 요소, 흐린 `box-shadow`, 불투명도가 점점 옅어지는 그레이디언트)는
해당 요소만 비트맵으로 그리고, 그 위에 텍스트를 보이지 않게 겹쳐 선택할 수 있게
합니다. 수식은 아직 비트맵입니다. `--pdf-raster`는 v0.3.0처럼 페이지마다 이미지
한 장을 씁니다.

### 접근성 있는 태그 PDF {#accessible-tagged-pdf}

v0.12.0부터 벡터 PDF에는 태그가 붙습니다. PDF에 논리 구조가 담기므로, 화면
낭독기는 웹 페이지를 읽을 때처럼 이를 순서대로 읽습니다.

| 프레젠테이션에서 | PDF에서 |
| --- | --- |
| 슬라이드 | 섹션(`Sect`) |
| `#` … `######` | `H1` … `H6` |
| 문단, 목록, 표 | `P`, 항목마다 레이블과 본문이 있는 `L`, 머리글 셀에 범위가 지정된 `Table` |
| 링크 | 클릭 영역을 포함하고 텍스트를 설명으로 갖는 `Link` |
| 이미지 | 대체 텍스트를 갖는 `Figure`. 대체 텍스트가 없는 `![](…)`는 장식으로 처리 |
| `$…$`, `$$…$$` | TeX 소스를 설명으로 갖는 `Formula` |
| 코드 | 줄 번호와 콜아웃 원을 뺀 `Code` |
| `<span lang="fr">` | 해당 언어의 스팬 |
| 배경, 테두리, 페이지 번호 | 낭독기가 건너뛰는 아티팩트 |
| 발표자 노트(`--pdf-notes`) | 구조 안의 주석 |

파일은 접근성 있는 PDF의 표준인 PDF/UA-1을 선언합니다. 뷰어가 표시하는
프레젠테이션의 `title`과, `lang`(프레젠테이션에 지정이 없으면 `--lang`, 그것도
없으면 `en`)을 담습니다.

veraPDF로 검사한 프레젠테이션은 PDF/UA-1을 통과합니다. 단, 프레젠테이션 자체의
제목이 표준의 순서를 어기는 경우는 예외입니다. PDF/UA는 첫 제목이 `H1`이어야 하고
수준을 건너뛰지 않아야 한다고(`#` 다음에 `###`) 요구합니다. 이는 프레젠테이션
작성자만 고칠 수 있습니다. marp-cli가 만드는 같은 프레젠테이션의 Chrome PDF에도
태그가 있지만, PDF/UA 규칙 일곱 가지를 어기고, 슬라이드별 섹션도 수식도 없습니다.

슬라이드 100장 벤치마크 프레젠테이션에 태그를 붙이면 크기가 4% 늘어나며(757 KiB
대비 788 KiB) 측정할 만한 시간은 들지 않습니다. `--pdf-tags=false`는 태그 없는
PDF를 씁니다.

## `slides.yaml` {#slidesyaml}

프레젠테이션 옆의 설정 파일이 그 프레젠테이션에 적용되므로, 한 폴더의
프레젠테이션들이 설정 파일 하나를 공유합니다.

```yaml
kinds: kinds.yaml        # a mapping of kinds, or file(s) holding one
themes:                  # theme CSS files or directories
  - themes/
data:                    # template data
  company: Example Inc.
html: false              # allow all raw HTML
lang: en
```

kind: 프레젠테이션 자체의 `kinds:`가 설정 파일과 명령줄의 kind보다 우선합니다.
데이터: 설정 파일과 명령줄의 `data`가 프레젠테이션의 프런트 매터보다 우선하므로,
같은 프레젠테이션을 서로 다른 데이터로 빌드할 수 있습니다.

## 계획된 기능 {#planned}

아직 구현되지 않은 기능은 PDF 안의 수식을 벡터로 출력하는 것입니다. 구현되면
이 페이지에 명령이나 플래그로 추가됩니다.
