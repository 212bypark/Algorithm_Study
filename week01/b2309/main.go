package main

import (
	"fmt"
	"sort"
	"strconv"
)

type DwarfData struct {
	dwarfSlice []int
	answer     []int
}

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
