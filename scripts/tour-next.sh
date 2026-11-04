#!/bin/bash
# tour/ 에서 마지막으로 푼 페이지를 찾아 그 다음 Tour of Go 페이지 폴더를 만들고 사파리 새 탭으로 연다.
# 폴더 이름이 곧 위치다: tour/01-basics/02-imports → go.dev/tour/basics/2
set -euo pipefail

cd "$(dirname "$0")/.."

sections=(basics flowcontrol moretypes methods generics concurrency)

# 섹션별 페이지 폴더 이름 (마지막 Congratulations 페이지는 뺐다)
# 원본 제목: github.com/golang/website/_content/tour/<섹션>.article 의 "* " 줄
slugs() {
	case "$1" in
		basics) echo packages imports exported-names functions functions-continued multiple-results named-return-values variables variables-with-initializers short-variable-declarations basic-types zero-values type-conversions type-inference constants numeric-constants ;;
		flowcontrol) echo for for-continued for-is-gos-while forever if if-with-a-short-statement if-and-else exercise-loops-and-functions switch switch-evaluation-order switch-with-no-condition defer stacking-defers ;;
		moretypes) echo pointers structs struct-fields pointers-to-structs struct-literals arrays slices slices-are-like-references-to-arrays slice-literals slice-defaults slice-length-and-capacity nil-slices creating-a-slice-with-make slices-of-slices appending-to-a-slice range range-continued exercise-slices maps map-literals map-literals-continued mutating-maps exercise-maps function-values function-closures exercise-fibonacci-closure ;;
		methods) echo methods methods-are-functions methods-continued pointer-receivers pointers-and-functions methods-and-pointer-indirection methods-and-pointer-indirection-2 choosing-a-value-or-pointer-receiver interfaces interfaces-are-implemented-implicitly interface-values interface-values-with-nil-underlying-values nil-interface-values the-empty-interface type-assertions type-switches stringers exercise-stringers errors exercise-errors readers exercise-readers exercise-rot13reader images exercise-images ;;
		generics) echo type-parameters generic-types ;;
		concurrency) echo goroutines channels buffered-channels range-and-close select default-selection exercise-equivalent-binary-trees exercise-equivalent-binary-trees sync-mutex exercise-web-crawler ;;
	esac
}

last=$(find tour -mindepth 2 -maxdepth 2 -type d -name '[0-9][0-9]-*' | sort | tail -1)

if [ -z "$last" ]; then
	section_idx=0
	page=1
else
	section_dir=$(basename "$(dirname "$last")") # 01-basics
	section_idx=$((10#${section_dir%%-*} - 1))
	page=$((10#$(basename "$last" | cut -d- -f1) + 1))

	read -ra names <<<"$(slugs "${sections[$section_idx]}")"
	if [ "$page" -gt "${#names[@]}" ]; then
		section_idx=$((section_idx + 1))
		if [ "$section_idx" -ge "${#sections[@]}" ]; then
			echo "투어 다 끝냈다. 마지막: $last"
			exit 0
		fi
		page=1
	fi
fi

section=${sections[$section_idx]}
read -ra names <<<"$(slugs "$section")"
dir=$(printf 'tour/%02d-%s/%02d-%s' $((section_idx + 1)) "$section" "$page" "${names[$((page - 1))]}")

mkdir -p "$dir"
if [ ! -e "$dir/main.go" ]; then
	printf 'package main\n\nfunc main() {\n}\n' >"$dir/main.go"
fi

url="https://go.dev/tour/$section/$page"
echo "마지막: ${last:-없음}"
echo "만든 폴더: $dir"
echo "다음: $url"

# VS Code 터미널이고 scripts/vscode-open-url 확장이 깔려 있으면 내장 브라우저로, 아니면 사파리로
if [ "${TERM_PROGRAM:-}" = vscode ] && code --list-extensions 2>/dev/null | grep -qx doldoldol21.open-url; then
	open "vscode://doldoldol21.open-url?url=$(printf %s "$url" | sed 's/:/%3A/g; s#/#%2F#g')"
else
	open -a Safari "$url"
fi
