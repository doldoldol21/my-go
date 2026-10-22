# my-go

Go 다시 공부하면서 쌓아두는 곳.

- `tour/` - [A Tour of Go](https://go.dev/tour) 따라가면서 연습한 것
- `leetcode/` - 릿코드 풀이
- `projects/` - 이것저것 만들어본 것 (booking-app은 Udemy 강의 따라한 거)

릿코드는 문제마다 폴더 하나에 `main.go` 하나. 한 폴더에 `func main`이 여러 개 있으면 IDE에서 에러가 나서 이렇게 나눴다.

```sh
go run ./leetcode/0013-roman-to-integer
```

투어 이어서 할 때는 아래를 돌리면 `tour/`에서 마지막으로 만든 폴더 다음 페이지가 사파리 새 탭으로 열린다. 섹션 마지막 페이지를 했으면 다음 섹션 1페이지로 넘어간다.

```sh
./scripts/tour-next.sh
```

## 커밋

`폴더: 내용` 으로 짧게. 폴더 밖 파일(README, go.mod 등)은 `chore:`.

```sh
git commit -m "tour: 01-basics 변수, 함수"
git commit -m "tour: 06-concurrency 채널 연습 문제"
git commit -m "leetcode: 20 valid parentheses"
git commit -m "leetcode: 13 디버그 출력 제거"
git commit -m "projects: booking-app 입력 검증 추가"
git commit -m "chore: README 수정"
```
