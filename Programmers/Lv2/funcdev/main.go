package main

import (
	"fmt"
	// "sort"
)

func solution(progresses []int, speeds []int) []int {
	var arr, ret []int

	for i := 0; i < len(progresses); i++ {
		tmp := (100 - progresses[i]) / speeds[i]
		if (100-progresses[i])%speeds[i] > 0 {
			tmp++
		}
		arr = append(arr, tmp)
	}
	num := arr[0]
	count := 1
	for i := 1; i < len(arr); i++ {
		if num < arr[i] {
			ret = append(ret, count)
			num = arr[i]
			count = 1
		} else {
			count++
		}
	}
	ret = append(ret, count)
	return ret
}

func main() {
	p1 := []int{
		93, 30, 55,
	}
	s1 := []int{
		1, 30, 5,
	}
	fmt.Println(solution(p1, s1))
	p2 := []int{
		95, 90, 99, 99, 80, 99,
	}
	s2 := []int{
		1, 1, 1, 1, 1, 1,
	}
	fmt.Println(solution(p2, s2))
}
