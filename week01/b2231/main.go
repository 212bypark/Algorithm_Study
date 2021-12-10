package main

import (
	"bufio"
	"os"
	"fmt"
)

type Data struct {
	R *bufio.Reader
	num int
}

func InputAction() *Data{
	d := &Data{}
	d.R = bufio.NewReader(os.Stdin)
	fmt.Fscan(d.R, &d.num)
	return d
}

func SplitSum(num int) int {
	sum := num
	for num != 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

func Solution(d *Data) {
	ans := 0
	for i := 0; i < d.num; i++ {
		if SplitSum(i) == d.num {
			ans = i
			break
		}
	}
	fmt.Println(ans)
}

func main() {
	d := InputAction()
	Solution(d)
}