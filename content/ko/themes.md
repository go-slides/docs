---
title: "테마 패키지"
weight: 42
description: "파일 하나로 된 테마: theme.hcl, CSS, 글꼴, 에셋, kind를 담은 .gsta.xz 아카이브이며, 신뢰할 수 없는 입력으로 검사합니다."
tags: [themes, hcl]
---

**테마 패키지**는 `<name>.gsta.xz`(Go-Slides Theme Archive)라는 파일 하나입니다.
루트에 `theme.hcl`이 있고 그 파일이 선언한 파일들을 담은, xz로 압축한 tar
아카이브(ustar 또는 pax)입니다. 테마의 CSS, 글꼴,
로고, 배경을 담으며, 원하면 슬라이드 kind도 담을 수 있고, 이름, 버전, 라이선스가
붙어 있어 테마를 주고받고 설치할 수 있습니다. 테마 패키지는 v0.18.0 이후 버전에서 제공됩니다.

## theme.hcl 파일 {#the-themehcl-file}

```hcl
# theme.hcl
theme {
  name        = "acme"
  version     = "1.0.0"                 # semantic version
  description = "Example theme"
  authors     = ["go-slides contributors"]
  license     = "MIT"                   # SPDX expression
  homepage    = "https://go-slides.github.io"
  extends     = "default"               # the base theme, built in or installed
  min_slides  = "0.18.0"                # oldest go-slides it works with
}

css "theme.css" {}                      # in order; sha256 = "…" is optional

font "Roboto Mono" {
  weight  = 400                         # or "100 900"
  style   = "normal"
  file    = "fonts/RobotoMono-Regular.ttf"
  license = "OFL-1.1"                   # required for a font
}

asset "logo" { file = "assets/logo.svg" }

kind "acme-title" {                     # optional slide kinds
  class    = "lead"
  paginate = "false"
}
```

`theme.hcl`은 리터럴 값만 쓰는 HCL의 기본 문법입니다. 함수도, 변수도, `file()`이나
`env()`도 없습니다. 이름은 소문자, 숫자, 대시로 이루어지며 `default`, `gaia`,
`uncover`는 쓸 수 없습니다. `extends`는 기반 테마를 지정하고, `css` 블록은 하나
추가될 때마다 순서대로 스타일시트를 더하며, `font`와 `asset` 블록은 각각 패키지
안의 파일을 가리키고, `kind` 블록은 하나하나가 [kind]({{< relref "kinds" >}})입니다.
`css`, `font`, `asset` 블록의 `sha256`은 패키지를 읽을 때 검사합니다.

CSS는 글꼴과 에셋을 패키지 안의 경로로, CSS 파일 기준 상대 경로로 참조합니다.
`background: url(assets/logo.svg)`처럼요. 패키지를 불러올 때 이런 참조는 모두
`data:` URL로 바뀌어 페이지가 그 자체로 완결되고, `font` 블록은 `@font-face`가
됩니다. 패키지에는 `LICENSE`, `NOTICE`, `README` 파일도 넣을 수 있으며,
`theme.hcl`이 선언하지 않은 다른 파일은 거부됩니다.

## 패키징, 설치, 사용 {#packing-installing-and-using}

```
slides theme pack dir       write dir/theme.hcl and its files as <name>.gsta.xz
slides theme verify file    run every check on a package
slides theme install file   check it, then install it for the current user
slides theme list           the installed packages
slides theme remove name
```

설치된 패키지는 사용자 설정 디렉터리 아래의 `go-slides/themes`에 있습니다
(`os.UserConfigDir()`: `~/Library/Application Support`, `$XDG_CONFIG_HOME` 또는
`~/.config`, `%AppData%`). 프레젠테이션에서 다음과 같이 쓰면:

```markdown
---
theme: acme
---
```

내장 테마와 `--theme-set` 다음 순서로 설치된 패키지를 찾습니다.
`--theme acme.gsta.xz`와 `--theme-set acme.gsta.xz`는 설치하지 않고 파일을
바로 사용합니다. [명령줄]({{< relref "cli" >}})을 참고하십시오.

패키지의 kind는 `--theme file.gsta.xz`로 적용되며, 설치된 패키지라면
`--theme name`으로, 또는 프레젠테이션의 프런트 매터에 `theme: name`을 적어서
적용됩니다. 프레젠테이션 자신의 kind가 이들보다 우선합니다.

패키지의 `extends`는 기본 테마를 지정하며, 기본 테마는 다음 순서로 찾습니다. 먼저 내장된 것(또는 `--theme-set`으로 지정한 것), 다음으로 설치된 것, 그다음으로 패키지 파일과 같은 디렉터리에 있는 형제 `<base>.gsta.xz`입니다. 이렇게 하면 테마 군(기본 테마 하나와 작은 자식 테마 여럿)을 파일만으로 바로 시험해 볼 수 있습니다. 오류 메시지에는 검색한 세 곳이 표시됩니다. 패키지 체인의 길이는 최대 8이며, 순환은 거부됩니다(`theme.meta`).

## 패키지는 신뢰할 수 없는 입력입니다 {#a-package-is-untrusted-input}

테마 패키지는 어디에서든 올 수 있으므로, 패키지를 읽는 모든 경로(`pack`, `verify`,
`install`, `--theme`, 빌드)가 똑같이 검사하고, 코드가 붙은 오류로 거부합니다
(`theme.archive`, `theme.limit`, `theme.meta`, `theme.css`, `theme.svg`, `theme.font`,
`theme.checksum`, `theme.version`).

- **아카이브.** xz 스트림은 하나뿐이며, 사전(dictionary)에는 한도가 있고 윈도는
  데이터에 따라 커집니다. tar 안에서는 이름에 `..`, 절대 경로, 드라이브 문자,
  백슬래시가 없어야 하고, 일반 파일과 디렉터리만 허용되며(심볼릭 링크, 하드 링크,
  장치, FIFO는 안 됨), 중복된 이름이 없어야 합니다(대소문자를 구분하지 않고 비교).
  ustar와 pax 헤더만 허용되며(GNU 확장도, 스파스 파일도 안 됨), pax 헤더는 최대
  8 KiB입니다. 크기는 헤더가 선언한 값이 아니라 실제로 읽은 값입니다.
- **한도.** 아카이브는 최대 16 MiB, 항목은 256개, 항목당 8 MiB, 압축을 푼 뒤 24 MiB,
  xz 사전은 16 MiB, 압축을 푼 스트림은 아카이브 크기의 100배 이하(64 KiB 이상일 때),
  `theme.hcl`은 최대 64 KiB에 괄호 중첩 32단계,
  CSS는 파일당 1 MiB, 글꼴이나 이미지는 4 MiB까지입니다.
- **CSS.** `@import`도, `@import-theme`도(`extends`를 쓰십시오), `image-set()`,
  `src()` 등 리소스를 불러오는 함수도, 문자열 밖의 백슬래시 이스케이프도,
  `</style`도 안 됩니다. `url()`은 `#fragment`, `data:` 이미지나 글꼴, 또는 패키지에
  선언된 파일이어야 하며, 패키지 밖으로 나가는 것은 안 됩니다.
- **SVG.** `script`, `foreignObject`, 프레임, embed, `set`, DOCTYPE, 이벤트 핸들러는
  안 됩니다. `href`는 `#fragment`여야 하고, 파일 밖으로 나가는 `url()`도 안 됩니다.
- **글꼴과 이미지**는 파싱합니다. 글꼴은 글꼴이어야 하고, 이미지는 확장자와
  일치해야 하며, 각각 크기 한도가 있습니다.

저장소의 `testdata/themes/acme`가 작은 예제입니다.
