package main

import (
	"strconv"
	"strings"
)

// func make_num(ret int, n int) int {
// 	if ret == 0 {
// 		return n
// 	} else {
// 		return ret*10 + n
// 	}
// }

func solution(s string) int {
	s = strings.ReplaceAll(s, "zero", "0")
	s = strings.ReplaceAll(s, "one", "1")
	s = strings.ReplaceAll(s, "two", "2")
	s = strings.ReplaceAll(s, "three", "3")
	s = strings.ReplaceAll(s, "four", "4")
	s = strings.ReplaceAll(s, "five", "5")
	s = strings.ReplaceAll(s, "six", "6")
	s = strings.ReplaceAll(s, "seven", "7")
	s = strings.ReplaceAll(s, "eight", "8")
	s = strings.ReplaceAll(s, "nine", "9")

	ret, _ := strconv.Atoi(s)

	return ret
}
