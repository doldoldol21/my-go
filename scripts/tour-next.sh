#!/bin/bash
# tour/ 에서 마지막으로 푼 페이지를 찾아 그 다음 Tour of Go 페이지를 사파리 새 탭으로 연다.
# 폴더 이름이 곧 위치다: tour/01-basics/02-imports → go.dev/tour/basics/2
set -euo pipefail

cd "$(dirname "$0")/.."

sections=(basics flowcontrol moretypes methods generics concurrency)

# 섹션별 페이지 수 (마지막 페이지는 Congratulations 라 건너뛴다)
pages() {
	case "$1" in
		basics) echo 17 ;;
		flowcontrol) echo 14 ;;
		moretypes) echo 27 ;;
		methods) echo 26 ;;
		generics) echo 3 ;;
		concurrency) echo 11 ;;
	esac
}

last=$(find tour -mindepth 2 -maxdepth 2 -type d -name '[0-9][0-9]-*' | sort | tail -1)

if [ -z "$last" ]; then
	section=basics
	page=1
else
	section_dir=$(basename "$(dirname "$last")") # 01-basics
	section_idx=$((10#${section_dir%%-*} - 1))
	section=${section_dir#*-}
	page=$((10#$(basename "$last" | cut -d- -f1) + 1))

	if [ "$page" -ge "$(pages "$section")" ]; then
		section_idx=$((section_idx + 1))
		if [ "$section_idx" -ge "${#sections[@]}" ]; then
			echo "투어 다 끝냈다. 마지막: $last"
			exit 0
		fi
		section=${sections[$section_idx]}
		page=1
	fi
fi

url="https://go.dev/tour/$section/$page"
echo "마지막: ${last:-없음}"
echo "다음: $url"
open -a Safari "$url"
