package main

import (
	"bufio"
	"fmt"
	"os"
)

// myData 구조체에 표준입력을 통해 받은 숫자 저장
type Data struct {
	R   *bufio.Reader
	num int
}

// 표준입력을 통해 값을 받아옴
func InputAction() *Data {
	d := &Data{}
	d.R = bufio.NewReader(os.Stdin)
	fmt.Fscan(d.R, &d.num)
	return d
}

// 분해합
func SplitSum(num int) int {
	sum := num
	for num != 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

// Complexity - Time: O(n) Quardratic | Space: O(1)
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
