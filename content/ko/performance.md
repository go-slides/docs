---
title: "성능"
weight: 60
description: "go-slides가 프레젠테이션을 얼마나 빠르게, 얼마만큼의 메모리로 빌드하는지 같은 컴퓨터에서 marp-cli와 비교해 측정했습니다."
tags: [performance]
---

{{< callout type="info" >}}
아래 PPTX 수치는 v0.5.0, 벡터 PDF 수치는 v0.4.0 기준입니다.
첫 번째 표는 이전 릴리스인 v0.3.0 기준이며, 이 버전부터 슬라이드를 동시에
렌더링하고 글꼴을 한 번만 파싱합니다
([go-slides/slides#11](https://github.com/go-slides/slides/pull/11)).
{{< /callout >}}

go-slides와 [marp-cli](https://github.com/marp-team/marp-cli) 4.5.1을 같은
프레젠테이션에 대해 하위 프로세스로 실행했습니다. Apple M4 Max(CPU 16개)에서
1분 평균 부하가 8 미만인 상태로, 예열 후 5회씩 번갈아 측정했습니다. 시간은
경과 시간(wall time)의 중앙값이고, 메모리는 전체 프로세스 트리의 최대값입니다
(marp-cli는 Chrome을 구동하므로 그 프로세스도 포함됩니다).

| 프레젠테이션 | 출력 | go-slides | marp-cli | go-slides 메모리 | marp-cli 메모리 |
| --- | --- | ---: | ---: | ---: | ---: |
| 슬라이드 100장 | HTML | 36 ms | 315 ms | 47 MiB | 140 MiB |
| 슬라이드 100장 | PDF | 1.14 s | 1.88 s | 242 MiB | 1,509 MiB |
| 슬라이드 100장 | PNG, 전체 슬라이드 | 0.97 s | 5.44 s | 225 MiB | 1,354 MiB |
| 슬라이드 100장 | PPTX | 3.39 s | 8.03 s | 530 MiB | 1,373 MiB |
| gaia, 슬라이드 20장 | PDF | 283 ms | 1.57 s | 154 MiB | 1,315 MiB |
| 수식과 코드, 슬라이드 20장 | PNG, 전체 슬라이드 | 250 ms | 2.32 s | 168 MiB | 1,286 MiB |

## 벡터 PDF {#vector-pdf}

v0.4.0부터 PDF는 벡터로 그려집니다
([명령줄]({{< relref "cli#vector-pdf" >}}) 참고). 같은 컴퓨터와 같은 방법으로
2026-10-19에 측정했습니다.

| 프레젠테이션 | go-slides | marp-cli | v0.3.0 | go-slides 크기 | marp-cli 크기 | v0.3.0 크기 |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| 슬라이드 100장 | 306 ms | 1,763 ms | 1,147 ms | 758 KiB | 837 KiB | 10.0 MiB |
| gaia, 슬라이드 20장 | 93 ms | 1,546 ms | 285 ms | 32 KiB | 66 KiB | 582 KiB |
| 수식과 코드, 슬라이드 20장 | 140 ms | 1,643 ms | 294 ms | 124 KiB | 335 KiB | 638 KiB |

슬라이드 100장 PDF의 메모리는 158 MiB로, marp-cli의 1,512 MiB와 비교됩니다.
CPU 시간은 0.9 s 대 2.7 s입니다.

슬라이드 100장 PDF의 v0.4.0 수치는 2026-10-19 21:29에 1분 평균 부하 7.7에서
3회 측정한 것입니다. go-slides는 306 ms, 758 KiB, 158 MiB였고 marp-cli는
1,838 ms, 837 KiB, 1,512 MiB였습니다.

## 편집 가능한 PPTX {#editable-pptx}

v0.5.0부터 PPTX는 텍스트 상자, 도형, 그림으로 이루어집니다
([명령줄]({{< relref "cli#editable-pptx-and-odp" >}}) 참고). 2026-10-19 22:51에
1분 평균 부하 7.1에서 3회 측정했습니다.

| 프레젠테이션 | go-slides | marp-cli | go-slides 크기 | marp-cli 크기 | go-slides 메모리 | marp-cli 메모리 |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| 슬라이드 100장 | 348 ms | 7,795 ms | 760 KiB | 7.2 MiB | 204 MiB | 1,370 MiB |
| gaia, 슬라이드 20장 | 112 ms | 4,608 ms | 893 KiB | 4.8 MiB | | |
| 수식과 코드, 슬라이드 20장 | 207 ms | 2,780 ms | 842 KiB | 1.5 MiB | | |

슬라이드 100장 프레젠테이션에서 v0.4.0은 4.7 s, 4.8 MiB, 540 MiB였습니다.
marp-cli의 PPTX에는 그림이 들어 있습니다. 편집 가능한 프레젠테이션 크기의
대부분은 내장 글꼴입니다.

## 빠른 이유 {#why-it-is-fast}

- **시작할 브라우저가 없습니다.** marp-cli는 PDF, 이미지, PPTX를 내보낼 때마다
  Chrome을 실행하지만, go-slides는 자체 프로세스 안에서
  [go-webengine](https://github.com/go-webengine/engine)으로 렌더링합니다.
- **슬라이드를 병렬로 렌더링합니다.** 기본값은 한 번에 네 장이며, 각 이미지는
  차례가 오는 즉시 기록되므로 긴 프레젠테이션도 이미지 형태로 메모리에 쌓아 두지
  않습니다.
- **글꼴은 한 번만 파싱합니다.** 슬라이드마다가 아니라 프로세스당 한 번입니다.

## 앞서지 못하는 부분 {#what-it-does-not-win}

- **v0.3.0의 PDF 크기.** PDF 안의 각 슬라이드가 이미지였기 때문에 슬라이드
  100장 프레젠테이션에서 10 MiB로, marp-cli의 837 KiB보다 컸습니다. 위의 벡터
  PDF는 marp-cli보다 작습니다.
- **큰 프레젠테이션의 CPU 시간.** 워커 네 개는 경과 시간이 보여 주는 것보다 더
  많은 CPU를 씁니다. v0.3.0에서 슬라이드 100장 PDF는 3.2 s로, marp-cli의
  2.7 s보다 많았습니다(벡터 PDF는 0.9 s).
- **marp-cli의 메모리는 상한값입니다.** Chrome 프로세스들은 페이지를 공유하는데,
  합계에서는 이를 프로세스마다 중복해 셉니다. 가장 큰 단일 프로세스는
  220–320 MiB입니다.

측정 방법, 모든 프레젠테이션과 모드, 수치를 재현하는 명령은
[BENCHMARKS.md](https://github.com/go-slides/slides/blob/main/BENCHMARKS.md)에
있습니다.
