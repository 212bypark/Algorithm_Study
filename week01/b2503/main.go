package main

import (
	"bufio"
	"fmt"
	"os"
)

// myData 구조체에 표준입력을 통해 받은 num 값과 각 게임 결과를 수열로 저장
// num : 물어본 횟수
type myData struct {
	R *bufio.Reader
	W *bufio.Writer

	num   int
	games []game
}

// 각각의 게임에는 질문했던 숫자(num), 스트라이크(stk), 볼(ball)이 있다
type game struct {
	num  int
	stk  int
	ball int
}

// 표준입력을 통해 값을 받아옴
func inputAction() *myData {
	d := &myData{}
	d.R = bufio.NewReader(os.Stdin)
	d.W = bufio.NewWriter(os.Stdout)

	fmt.Fscanln(d.R, &d.num)
	for i := 0; i < d.num; i++ {
		game := game{}
		fmt.Fscanln(d.R, &game.num, &game.stk, &game.ball)
		d.games = append(d.games, game)
	}
	return d
}

// 물어본 수의 각자리는 중복되지 않는다 (서로 다른 세자리 숫자)
// -> 유효성 검사
func isVal(a, b, c int) bool {
	if a == 0 || b == 0 || c == 0 {
		return false
	}
	if a == b || b == c || c == a {
		return false
	}
	return true
}

// 123부터 987까지 모든 숫자들에 대해 결과값과 비교
// Complexity - Time: O(n) | Space: O(1)
func calAvail(d *myData) {
	count := 0
	for i := 123; i <= 987; i++ {
		isAns := true
		a, b, c := i/100, (i/10)%10, i%10
		if !isVal(a, b, c) {
			continue
		}
		for j := 0; j < d.num; j++ {
			stk, ball, n := 0, 0, d.games[j].num
			a2, b2, c2 := n/100, (n/10)%10, n%10
			if a == a2 {
				stk++
			}
			if a == b2 || a == c2 {
				ball++
			}
			if b == b2 {
				stk++
			}
			if b == a2 || b == c2 {
				ball++
			}
			if c == c2 {
				stk++
			}
			if c == a2 || c == b2 {
				ball++
			}
			if !(stk == d.games[j].stk && ball == d.games[j].ball) {
				isAns = false
				break
			}
		}
		if isAns {
			count++
		}
	}
	fmt.Println(count)
}

func solution() {
	d := inputAction()
	defer d.W.Flush()
	calAvail(d)
}

func main() {
	solution()
}
