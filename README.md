# my-go

Go 다시 공부하면서 쌓아두는 곳.

- `tour/` - [A Tour of Go](https://go.dev/tour) 연습 (한국어판으로 본다)
- `leetcode/` - 릿코드 풀이
- `projects/` - 이것저것 만든 것 (booking-app은 Udemy 강의)

## 릿코드

문제 하나에 폴더 하나, `main.go` 하나. 한 폴더에 `func main`이 여러 개면 IDE가 에러를 낸다.

```sh
go run ./leetcode/0013-roman-to-integer
```

VS Code에서는 `Cmd+Shift+B`로 열려 있는 파일의 폴더를 바로 실행한다 (`.vscode/tasks.json`).

## 투어 이어하기

```sh
./scripts/tour-next.sh
```

`tour/`의 마지막 폴더 다음 페이지를 한국어판으로 열고, 그 페이지 폴더(`tour/01-basics/03-exported-names/main.go`)를 만든다. 섹션이 끝났으면 다음 섹션 1페이지로 넘어간다.

VS Code 터미널에서 돌리면 내장 브라우저로, 아니면 사파리로 열린다. 내장 브라우저로 열려면 확장을 한 번 깔아 둔다.

```sh
cd scripts/vscode-open-url && npx @vscode/vsce package --skip-license -o /tmp/open-url.vsix && code --install-extension /tmp/open-url.vsix
```

두 번 돌리면 한 페이지를 건너뛴다. 그럴 땐 마지막 폴더를 지운다.

## 커밋

`폴더: 내용`으로 짧게. 폴더 밖 파일(README, go.mod 등)은 `chore:`.

```sh
git commit -m "tour: 01-basics 변수, 함수"
git commit -m "leetcode: 20 valid parentheses"
git commit -m "projects: booking-app 입력 검증 추가"
git commit -m "chore: README 수정"
```
