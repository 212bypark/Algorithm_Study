package main

import (
	"bufio"
	"fmt"
	"os"
)

type myData struct {
	R *bufio.Reader
	W *bufio.Writer

	num   int
	games []game
}

type game struct {
	num  int
	stk  int
	ball int
}

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

func isVal(a, b, c int) bool {
	if a == 0 || b == 0 || c == 0 {
		return false
	}
	if a == b || b == c || c == a {
		return false
	}
	return true
}

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
