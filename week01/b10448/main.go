package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// myData 구조체에 표준입력을 통해 받은 테스트 케이스 수와
// 테스트 값을 inputSlice 배열에 저장
type myData struct {
	R *bufio.Reader
	W *bufio.Writer

	num        int
	inputSlice []int
	trioArray  []int
}

// 표준입력을 통해 값을 받아옴
func inputAction() *myData {
	d := &myData{}
	d.R = bufio.NewReader(os.Stdin)
	d.W = bufio.NewWriter(os.Stdout)

	fmt.Fscanln(d.R, &d.num)
	for i := 0; i < d.num; i++ {
		line, _, _ := d.R.ReadLine()
		num, _ := strconv.Atoi(string(line))
		d.inputSlice = append(d.inputSlice, num)
	}
	return d
}

// val(val+1)/2 == input : n이 무한대로 발산할 경우
// val 만큼의 삼각수 생성하여 trioArray에 저장
func makeTrioArray(input int, d *myData) {
	val := 0
	for val*val < 2*input {
		val++
	}
	for i := 1; i < val; i++ {
		d.trioArray = append(d.trioArray, i*(i+1)/2)
	}
}

// 3개의 삼각수로 표현 가능한지 검사
// Complexity - Time: O(n^3) Cubic | Space: O(n^2)
func checkAvail(num int, trioArray []int) {
	check := 0
	for i := 0; i < len(trioArray); i++ {
		val1 := trioArray[i]
		if val1 > num {
			break
		}
		for j := 0; j < len(trioArray); j++ {
			val2 := trioArray[j]
			if val1+val2 > num {
				break
			}
			for k := 0; k < len(trioArray); k++ {
				val3 := trioArray[k]
				if val1+val2+val3 == num {
					check = 1
					goto END
				}
				if val1+val2+val3 > num {
					break
				}

			}
		}
	}

END:
	fmt.Println(check)
}

func getReady() {
	d := inputAction()
	defer d.W.Flush()
	for i := 0; i < d.num; i++ {
		makeTrioArray(d.inputSlice[i], d)
		checkAvail(d.inputSlice[i], d.trioArray)
	}
}

func main() {
	getReady()
}
