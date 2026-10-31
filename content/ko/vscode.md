---
title: "VS Code"
weight: 55
description: "VS Code용 go-slides 확장: Marp for VS Code가 Marp에 제공하는 것처럼 미리 보기, 발표, 문제 표시, 자동 완성, 내보내기를 제공합니다."
tags: [vscode, editor]
---

{{< callout type="info" >}}
이 확장은 [go-slides/vscode](https://github.com/go-slides/vscode)이며, 버전은
따로 관리됩니다(v0.5.0). 렌더링에는 go-slides v0.14.0을 사용합니다(`editor`
패키지, 스크롤 동기화용 소스 줄, 개요와 발표자 보기). GitHub 릴리스에 첨부된
`go-slides-0.5.0.vsix`로 설치합니다. VS Code Marketplace와 Open VSX 게시는
준비되어 있지만 아직 올라가 있지 않습니다.
{{< /callout >}}

이 확장은 [Marp for VS Code](https://github.com/marp-team/marp-vscode)에
대응하는 go-slides용 확장입니다. 프레젠테이션은 VS Code의 Markdown 미리 보기에
슬라이드로 표시되고, 오류는 문제 패널에 나타나며, 디렉티브, kind(슬라이드
유형), 템플릿 함수는 입력하는 동안 자동 완성됩니다.

## 설치 {#install}

[릴리스](https://github.com/go-slides/vscode/releases)에서
`go-slides-<version>.vsix`를 내려받은 다음 *Extensions: Install from VSIX...*를
실행하거나, 다음 명령을 실행하십시오.

```sh
code --install-extension go-slides-0.5.0.vsix
```

미리 보기에는 그 밖에 필요한 것이 없습니다. 내보내기에는 `slides` 명령이
필요합니다.

```sh
go install github.com/go-slides/slides/cmd/slides@latest
```

## 미리 보기 {#preview}

프레젠테이션을 열고 *Markdown: Open Preview to the Side*를 실행합니다.
프레젠테이션은 프런트 매터에 `slides: true`가 있는 Markdown 파일입니다.

```markdown
---
slides: true
paginate: true
data:
  who: World
---

# Hello {{ .who }}

---

<!-- _kind: title -->
# Page {{ page }} of {{ total }}
```

`marp: true`가 있는 파일도 미리 보기를 제공합니다. 단, Marp for VS Code가
설치되어 있으면 두 확장이 같은 문서를 렌더링하게 되므로 예외입니다. 설정
`goSlides.preview.marpDocuments`로 정하며, 값은 `auto`(기본값), `always`,
`never`입니다.

미리 보기는 WebAssembly로 컴파일되어 확장이 불러오는 go-slides 자체입니다.
바이너리가 필요 없으며 [vscode.dev](https://vscode.dev)에서도 동작합니다.
프레젠테이션이 사용하는 파일, 즉 프레젠테이션 옆의 `slides.yaml`이나
`slides.hcl`, 거기서 지정한 [kind]({{< relref "kinds" >}}) 파일,
[템플릿]({{< relref "templates" >}})이 포함하는 파일, 테마 CSS를 작업 영역에서
읽으며, 프레젠테이션이 속한 작업 영역 폴더에서만 읽습니다.

## 스크롤 동기화 {#scroll-sync}

확장 v0.2.0(go-slides v0.7.0)부터 편집기와 미리 보기가 일반 Markdown에서처럼
양방향으로 함께 스크롤됩니다. 미리 보기는 편집기 커서가 있는 슬라이드와 블록을
따라가고, 미리 보기를 스크롤하면 편집기가 움직입니다. 설정은 VS Code 자체의
`markdown.preview.scrollPreviewWithEditor`와
`markdown.preview.scrollEditorWithPreview`입니다.

줄 위치는 템플릿을 거쳐서도 추적됩니다. `{{ range }}`가 생성한 슬라이드는
템플릿의 해당 줄을 따라가고, `{{ include }}`로 가져온 내용은 include가 있는
줄을 따라갑니다. `headingDivider`로 나뉜 프레젠테이션은 제목을 기준으로
동기화됩니다.

## 발표하기 {#presenting}

확장 v0.4.0(go-slides v0.9.0)부터 VS Code를 떠나지 않고 프레젠테이션을 발표할
수 있습니다. 편집기에 프레젠테이션을 연 상태에서 명령 팔레트나 편집기의 `…`
메뉴에서 다음 명령을 실행합니다.

| 명령 | 여는 화면 |
|---|---|
| *go-slides: Open Overview* | 모든 슬라이드를 격자로 보여 주는 패널. 슬라이드를 고르면 편집기에 그 소스가 표시되며, `Esc`로 패널을 닫습니다 |
| *go-slides: Start Presenter View* | 슬라이드와 그 옆의 발표자 보기: 다음 단계, 발표자 노트, 타이머 |

패널은 편집기 옆에 놓이며 서로 연동됩니다. 한 패널에서 슬라이드를 바꾸면 다른
패널도 함께 바뀝니다. 프레젠테이션을 편집한 내용도 따라갑니다. 슬라이드에서 `p`
키를 누르면 발표자 보기가 열립니다.

다른 화면에서 발표하려면 *View: Move Editor into New Window*로 슬라이드 패널을
그쪽으로 옮기십시오.

각 패널은 `default-src 'none'` 정책 아래 샌드박스 프레임(`allow-same-origin`
없음)에서 프레젠테이션을 실행합니다. 프레임은 이미지, 미디어, 스타일, 글꼴만
불러오고, go-slides가 페이지에 넣은 해시로 지정된 스크립트만 실행합니다.
프레젠테이션에 작성된 스크립트는 실행되지 않습니다.

## 문제 표시, 자동 완성, 호버 {#problems-completion-hover}

- **문제 표시**: 템플릿, YAML, [HCL]({{< relref "hcl" >}}) 오류를 해당 줄에
  표시합니다. 오류가 프레젠테이션에 있든, 그 출처인 kind 파일에 있든 마찬가지입니다.
- **자동 완성**: 프런트 매터와 `<!-- -->` 주석의 디렉티브 키(스폿 디렉티브는
  `_` 포함), `kind`와 `_kind`의 값(각 kind가 설정하는 내용과 함께), `theme`,
  `size`, `paginate`, `math`, `headingDivider`, 그리고 `{{ }}` 안에서는 템플릿
  함수와 데이터 필드를 완성합니다.
- **호버**: 디렉티브의 역할, kind가 설정하는 내용, 템플릿 함수의 인수를
  보여 줍니다.

## 내보내기 {#export}

*go-slides: Export Slide Deck...*(편집기 제목 표시줄에도 있음)은
`slides -o deck.pdf deck.md`처럼 `slides` 명령으로 HTML, 벡터 PDF, 편집 가능한
PPTX와 ODP, PNG, JPEG를 만듭니다. [명령줄]({{< relref "cli" >}})을
참고하십시오. 신뢰할 수 있는 작업 영역이 필요합니다. `slides`를 찾지 못하면
확장이 위의 `go install`을 터미널에서 실행하도록 제안합니다.

확장 v0.5.0(go-slides v0.14.0)부터 내보낸 PDF는 **접근성**을 갖춥니다.
`slides` v0.12.0 이상에서는 벡터 PDF에 화면 낭독기용 태그가 붙고 PDF/UA-1을
선언하며, 프레젠테이션의 `title`과 `lang`, 슬라이드마다 하나의 섹션을 담습니다.
`slides --pdf-tags=false`는 태그 없는 파일을 만듭니다. 위의 `go install`로
명령을 업데이트하십시오.

## 설정 {#settings}

| 설정 | 기본값 | |
|---|---|---|
| `goSlides.cliPath` | `""` | `slides` 명령. 비어 있으면 `PATH`의 `slides`를 씁니다. 컴퓨터 설정이므로 작업 영역에서는 지정할 수 없습니다 |
| `goSlides.exportType` | `pdf` | 내보내기 대화 상자가 처음 제안하는 형식 |
| `goSlides.exportAutoOpen` | `true` | 내보낸 파일 열기 |
| `goSlides.html` | `false` | Marp 허용 목록 대신 모든 원시 HTML 허용. 신뢰할 수 있는 작업 영역에서만 적용됩니다 |
| `goSlides.preview.marpDocuments` | `auto` | `marp: true` 프레젠테이션 미리 보기: `auto`, `always`, `never` |

## 보안 {#security}

프레젠테이션과 작업 영역은 신뢰하지 않는 것으로 취급합니다. 원시 HTML은
신뢰할 수 있는 작업 영역에서 `goSlides.html`을 켜지 않는 한 Marp 허용 목록을
따르며(`html: true`가 있는 `slides.yaml`로는 미리 보기에서 켜지지 않습니다),
VS Code 미리 보기는 프레젠테이션의 스크립트를 실행하지 않습니다. 미리 보기는
프레젠테이션이 속한 작업 영역 폴더의 파일만 읽고, 심볼릭 링크를 따라가지 않으며,
템플릿은 환경 변수를 볼 수 없습니다. 내보내기는 프로그램을 실행하므로 작업 영역
신뢰를 기다리며, 실행할 프로그램을 작업 영역이 지정하는 일은 없습니다.

## 성능 {#performance}

go-slides의 슬라이드 100장짜리 벤치마크 프레젠테이션으로, Apple M4 Max의
VS Code 1.142에서 같은 세션에 Marp for VS Code 3.6.1도 함께 측정했습니다.

| | go-slides | Marp for VS Code |
|---|---:|---:|
| 활성화(WebAssembly 모듈 로드) | 189 ms | — |
| 첫 미리 보기 | 52 ms | 78 ms |
| 편집 후 미리 보기(중앙값) | **19 ms** | 23–24 ms |
| `.vsix` | 7.5 MB | 36.3 MB |

수치는 확장 v0.3.0(go-slides v0.8.0) 기준이며, 슬라이드 100장 프레젠테이션의
50번째 슬라이드에 문자 하나를 입력하는 작업을 1분 부하 8 미만에서 번갈아
측정한 결과입니다. 첫 미리 보기는 Marp보다 빠르고, 이제 편집 후 미리 보기도
Marp보다 빠릅니다.

### v0.8.0과 확장 v0.3.0 {#in-v080-and-extension-v030}

go-slides v0.8.0은 열려 있는 각 프레젠테이션의 렌더링된 슬라이드를 보관하고,
편집 후에는 소스나 상속된 상태(디렉티브, 페이지 번호, 전체 수, 제목 id, 링크
참조, 소스 줄)가 바뀐 슬라이드만 다시 렌더링합니다. 결과는 전체 렌더링과
바이트 단위까지 같습니다. 100장 중 50번째 슬라이드에 입력하면 슬라이드 한 장만
다시 렌더링하며, 네이티브로 2.5 ms, WebAssembly로 11.5 ms가 걸립니다. 그러면
VS Code 미리 보기는 바뀐 슬라이드의 DOM만 변형하며, 확장의 스모크 테스트는 두
확장 모두에 대해 시간을 재는 렌더링 사이에 50번째 슬라이드에 문자 하나를
입력합니다. 확장의 발표자 보기는 부모 창을 통해(`postMessage`) 동기화합니다.

### v0.7.1과 확장 v0.2.2 {#in-v071-and-extension-v022}

go-slides v0.7.1은 파싱한 테마, 압축한 스타일시트, 강조 표시한 코드를 렌더링
사이에 캐시하고, 스크롤 동기화용 줄을 바이트 단위로 훑지 않고 표시하며, 작성기를
재사용합니다. 슬라이드 100장 프레젠테이션 한 번 렌더링에 할당하는 메모리는
5.7 MB로, 확장 v0.2.0의 52.7 MB보다 적습니다. 덕분에 확장 v0.2.2는
`GOGC=400` 재정의를 없앴습니다. 비교하자면 확장 v0.2.1은 기본 `GOGC`에서 첫 미리
보기 72 ms, 편집 후 44 ms(`GOGC=400`에서는 36 ms)였고, v0.2.0은 47–51 ms였습니다.

확장 v0.2.2는 첫 미리 보기 52 ms, 편집 후 24 ms로 측정되어 Marp의 23 ms와 대등한
수준이었습니다.
