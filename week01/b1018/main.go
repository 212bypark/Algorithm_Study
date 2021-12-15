package main

import (
	"bufio"
	"fmt"
	"os"
)

// myData 구조체에 표준입력을 통해 받은 n,m 값과 보드를 수열로 저장
type myData struct {
	R *bufio.Reader
	W *bufio.Writer

	n     int
	m     int
	chess []string
}

// 표준입력을 통해 값을 받아옴
func inputAction() *myData {
	d := &myData{}
	d.R = bufio.NewReader(os.Stdin)
	d.W = bufio.NewWriter(os.Stdout)

	fmt.Fscanln(d.R, &d.n, &d.m)
	var nat = make([]string, d.n)
	for i := 0; i < d.n; i++ {
		fmt.Fscan(d.R, &nat[i])
	}
	d.chess = nat
	return d
}

// 두 매개변수를 비교하여 더 작은 값 반환
// math 패키지의 Min(x, y float64) float64 함수로 대체 가능
func myMin(num1, num2 int) int {
	if num1 > num2 {
		return num2
	} else {
		return num1
	}
}

// 8X8 보드를 만들어야 함 i, j 크기가 8만큼 순회
// a, b 변수는 보드에서 각각 x, y 좌표 의미
// Complexity - Time: O(n^2) Quardratic | Space: O(1)
func count(d *myData) int {
	min := d.n * d.m
	for i := 0; i < d.n-7; i++ {
		for j := 0; j < d.m-7; j++ {
			count1, count2 := 0, 0
			for a := i; a < i+8; a++ {
				for b := j; b < j+8; b++ {
					if (a+b)%2 == 0 {
						if string(d.chess[a][b]) == "B" {
							count1++
						} else {
							count2++
						}
					} else {
						if string(d.chess[a][b]) == "B" {
							count2++
						} else {
							count1++
						}
					}
				}
			}
			if min > myMin(count1, count2) {
				min = myMin(count1, count2)
			}
		}
	}
	return min
}

func main() {
	d := inputAction()
	fmt.Println(count(d))
}
