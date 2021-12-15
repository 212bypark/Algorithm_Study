// 반복되는 수열을 제외한 나머지 수들의 개수를 반환
package main

import (
	"bufio"
	"fmt"
	"os"
)

// myData 구조체에 표준입력을 통해 받은 A, P 값과
// 반복되는 부분을 제외한 수열에 남게 되는 수들을 배열에 저장
type Data struct {
	R         *bufio.Reader
	W         *bufio.Writer
	num       int
	power     int
	iterSlice []int
}

// 표준입력을 통해 값을 받아옴
func inputAction() *Data {
	d := &Data{}

	d.R = bufio.NewReader(os.Stdin)
	d.W = bufio.NewWriter(os.Stdout)

	fmt.Fscan(d.R, &d.num, &d.power)
	return d
}

// power제곱
func power(value int, power int) int {
	num := value
	for i := power - 1; i != 0; i-- {
		num *= value
	}
	return num
}

// d.num에 값 대입 후 iterSlice 과 비교하여 반복 여부 체크
// Complexity - Time: O(n) Quardratic | Space: O(1)
func solution(d *Data) int {
	for {
		d.iterSlice = append(d.iterSlice, d.num)
		next_num := 0
		// 다음 수열에 나올 값 계산
		for d.num != 0 {
			next_num += power(d.num%10, d.power)
			d.num /= 10
		}
		d.num = next_num
		// 다음 수열 값이 기존 수열에 존재하는지 체크
		i := 0
		for _, value := range d.iterSlice {
			if value == d.num {
				d.iterSlice = d.iterSlice[:i]
				goto END
			}
			i += 1
		}
	}

END:
	answer := len(d.iterSlice)

	return answer
}

func main() {
	d := inputAction()
	defer d.W.Flush()
	ans := solution(d)
	fmt.Println(ans)
}
