#!/bin/bash
# 바뀐 파일을 보고 README 커밋 규칙대로 메시지를 만들어 커밋하고 푸시한다.
#   tour/01-basics/03-exported-names ~ 09-...  → tour: 01-basics exported-names ~ variables-with-initializers
#   leetcode/0013-..., leetcode/0014-...       → leetcode: 13, 14
#   projects/booking-app/...                   → projects: booking-app
# 폴더 밖 파일(README, scripts 등)이 섞여 있으면 메시지를 직접 넘긴다: ./scripts/commit.sh "chore: README 수정"
set -euo pipefail

cd "$(dirname "$0")/.."

paths=$(git status --porcelain -uall | cut -c4- | sed 's/.* -> //')
if [ -z "$paths" ]; then
	echo "바뀐 게 없다."
	exit 0
fi

if [ $# -gt 0 ]; then
	msg="$*"
	git add -A
else
	others=$(printf '%s\n' "$paths" | grep -vE '^(tour|leetcode|projects)/' || true)
	if [ -n "$others" ]; then
		echo "tour/leetcode/projects 밖 파일이 바뀌어서 메시지를 정할 수 없다. 직접 넘겨줘:" >&2
		printf '  %s\n' $others >&2
		echo '  ./scripts/commit.sh "chore: ..."' >&2
		exit 1
	fi

	parts=()

	# tour: 섹션별로 페이지 이름을 모은다 (한 페이지면 그 이름, 여러 페이지면 첫 ~ 끝)
	for section in $(printf '%s\n' "$paths" | awk -F/ '$1=="tour" && NF>2 {print $2}' | sort -u); do
		pages=$(printf '%s\n' "$paths" | awk -F/ -v s="$section" '$1=="tour" && $2==s && $3 ~ /^[0-9][0-9]-/ && NF>3 {print $3}' | sort -u)
		if [ -z "$pages" ]; then
			parts+=("tour: $section")
			continue
		fi
		first=$(printf '%s\n' "$pages" | head -1)
		last=$(printf '%s\n' "$pages" | tail -1)
		if [ "$first" = "$last" ]; then
			parts+=("tour: $section ${first#*-}")
		else
			parts+=("tour: $section ${first#*-} ~ ${last#*-}")
		fi
	done

	# leetcode: 문제 번호
	nums=$(printf '%s\n' "$paths" | awk -F/ '$1=="leetcode" && NF>2 {print $2}' | sort -u | while read -r d; do echo $((10#${d%%-*})); done | paste -sd, - | sed 's/,/, /g')
	[ -n "$nums" ] && parts+=("leetcode: $nums")

	# projects: 프로젝트 이름
	projs=$(printf '%s\n' "$paths" | awk -F/ '$1=="projects" && NF>2 {print $2}' | sort -u | paste -sd, - | sed 's/,/, /g')
	[ -n "$projs" ] && parts+=("projects: $projs")

	# 여러 종류가 섞이면 " / " 로 한 줄에 잇는다
	msg=$(printf '%s\n' "${parts[@]}" | paste -sd/ - | sed 's#/# / #g')
	git add tour leetcode projects
fi

git commit -q -m "$msg"
git log --oneline -1
git push -q
echo "푸시 완료"
