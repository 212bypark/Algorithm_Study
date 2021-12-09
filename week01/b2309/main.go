// 표준입력으로 9 개의 숫자를 받아
// 그 중 전체 합이 100인 7개의 숫자를 뽑아 오름차순으로 출력
// 조건 : 중복되지 않고 100이 넘지 않는 9개의 자연수가 주어지며
// 가능한 답이 여러가지인 경우에는 아무거나 출력 (답이 없는 경우는 없다)
package main

import (
	"fmt"
	"sort"
	"strconv"
)

// 표준입력으로 받는 값을 dwarfSlice에 저장하고 출력할 값을 answer에 저장
type DwarfData struct {
	dwarfSlice []int
	answer     []int
}

// 표준입력 처리를 위한 함수이다. 개행을 기준으로 파싱
func inputAction() *DwarfData {
	d := &DwarfData{}

	for i := 0; i < 9; i++ {
		var str string
		fmt.Scanln(&str)
		num, _ := strconv.Atoi(str)
		d.dwarfSlice = append(d.dwarfSlice, num)
	}
	return d
}

// 조건에 만족하는 숫자를 찾아 DwarfData.answer에 저장하고 정렬
func chooseDwarf(d *DwarfData) {
	totalSum := 0
	for _, val := range d.dwarfSlice {
		totalSum += val
	}
	for i := 0; i < 8; i++ {
		for j := i + 1; j < 9; j++ {
			if i != j && totalSum-(d.dwarfSlice[i]+d.dwarfSlice[j]) == 100 {
				for _, num := range d.dwarfSlice {
					if num != d.dwarfSlice[i] && num != d.dwarfSlice[j] {
						d.answer = append(d.answer, num)
					}
				}
			}
		}
	}
	sort.Sort(sort.IntSlice(d.answer))
}

func main() {
	d := inputAction()
	chooseDwarf(d)
	// fmt.Println(d.answer)
	fmt.Println("=================")
	for _, realDwarf := range d.answer {
		fmt.Println(realDwarf)
	}
}
