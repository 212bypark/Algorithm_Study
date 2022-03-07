package main

import (
	"fmt"
	"strings"
)

func avail_seat(x int, y int) int {
	if x >= 0 && x < 5 && y >= 0 && y < 5 {
		return 1
	}
	return 0
}

func check(places [5][5]string, x int, y int) int {
	if avail_seat(x, y-2) == 1 {
		if places[x][y-2] == "P" && places[x][y-1] != "X" {
			return 0
		}
	}
	if avail_seat(x-1, y-1) == 1 {
		if places[x-1][y-1] == "P" && !(places[x][y-1] == "X" && places[x-1][y] == "X") {
			return 0
		}
	}
	if avail_seat(x, y-1) == 1 {
		if places[x][y-1] == "P" {
			return 0
		}
	}
	if avail_seat(x+1, y-1) == 1 {
		if places[x+1][y-1] == "P" && !(places[x][y-1] == "X" && places[x+1][y] == "X") {
			return 0
		}
	}
	if avail_seat(x-2, y) == 1 {
		if places[x-2][y] == "P" && places[x-1][y] != "X" {
			return 0
		}
	}
	if avail_seat(x-1, y) == 1 {
		if places[x-1][y] == "P" {
			return 0
		}
	}
	if avail_seat(x+1, y) == 1 {
		if places[x+1][y] == "P" {
			return 0
		}
	}
	if avail_seat(x+2, y) == 1 {
		if places[x+2][y] == "P" && places[x+1][y] != "X" {
			return 0
		}
	}
	if avail_seat(x+1, y+1) == 1 {
		if places[x+1][y+1] == "P" && !(places[x+1][y] == "X" && places[x][y+1] == "X") {
			return 0
		}
	}
	if avail_seat(x, y+2) == 1 {
		if places[x][y+2] == "P" && places[x][y+1] != "X" {
			return 0
		}
	}
	return 1
}

func solution(places [][]string) []int {
	num_place := len(places)
	var target [5][5]string
	ret := make([]int, num_place)
	for idx := 0; idx < num_place; idx++ {
		val := 1
		for i := 0; i < 5; i++ {
			tmp := strings.Split(places[idx][i], "")
			copy(target[i][:], tmp[:5])
		}
		for x := 0; x < 5; x++ {
			for y := 0; y < 5; y++ {
				if target[x][y] == "P" {
					if check(target, x, y) == 0 {
						val = 0
					}
				}
				if val == 0 {
					break
				}
			}
			if val == 0 {
				break
			}
		}

		ret[idx] = val
	}
	return ret
}

/*
						(x, y-2)
			(x-1, y-1)	(x, y-1)	(x+1, y-1)
(x-2, y)	(x-1, y)	(x, y)		(x+1, y)	(x+2, y)
			(x+1, y+1)	(x, y+1)	(x+1, y+1)
						(x, y+2)
*/

func main() {
	arr := [][]string{
		{"POOOP", "OXXOX", "OPXPX", "OOXOX", "POXXP"},
		{"POOPX", "OXPXP", "PXXXO", "OXXXO", "OOOPP"},
		{"PXOPX", "OXOXP", "OXPOX", "OXXOP", "PXPOX"},
		{"OOOXX", "XOOOX", "OOOXX", "OXOOX", "OOOOO"},
		{"PXPXP", "XPXPX", "PXPXP", "XPXPX", "PXPXP"},
	}
	arr2 := [][]string{
		{"PXOOO", "OOOOO", "PXOOO", "OOOOO", "OOOPO"},
		{"POOPX", "OXPXP", "PXXXO", "OXXXO", "OOOPP"},
		{"PXOPX", "OXOXP", "OXPOX", "OXXOP", "PXPOX"},
		{"OOOXX", "XOOOX", "OOOXX", "OXOOX", "OOOOO"},
		{"PXPXP", "XPXPX", "PXPXP", "XPXPX", "PXPXP"},
	}

	fmt.Println(solution(arr))
	fmt.Println(solution(arr2))
}
