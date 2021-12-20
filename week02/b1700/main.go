package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type inputData struct {
	plug  int
	order []int
	// amount [][2]int
	amount map[int]int
}

func inputAction() *inputData {
	R := bufio.NewReader(os.Stdin)
	d := &inputData{}
	var num int
	fmt.Fscanln(R, &d.plug, &num)
	for i := 0; i < num; i++ {
		val := 0
		fmt.Fscan(R, &val)
		d.order = append(d.order, val)
	}
	var temp []int
	temp = append(temp, d.order...)
	sort.Ints(temp)
	// e := [2]int{temp[0], 1}
	// d.amount = append(d.amount, e)
	// index := 0
	// for i := 1; i < len(temp); i++ {
	// 	if d.amount[index][0] == temp[i] {
	// 		d.amount[index][1]++
	// 	} else {
	// 		e = [2]int{temp[i], 1}
	// 		d.amount = append(d.amount, e)
	// 		index++
	// 	}
	// }
	amount := make(map[int]int)
	for _, i := range temp {
		if amount[i] == 0 {
			amount[i] = 1
		} else {
			amount[i] = amount[i] + 1
		}
	}
	d.amount = amount
	// fmt.Println(amount)
	// fmt.Println(d.amount)
	return d
}

func whoIsMin(d *inputData, arr []int) int {
	index := 0
	target := d.amount[arr[0]]
	for i := 1; i < len(arr); i++ {
		if target > d.amount[arr[i]] {
			target = d.amount[arr[i]]
			index = i
		}
	}
	return index
}

func getMinCount(d *inputData) {
	count := 0
	curr := 0
	arr := make([]int, d.plug)
	for _, a := range d.order {
		if curr != d.plug {
			arr[curr] = a
			d.amount[a] = d.amount[a] - 1
			curr++
		} else {
			for _, b := range arr {
				if a == b {
					d.amount[a] = d.amount[a] - 1
					// fmt.Println("a==b : a=", a, ", b=", b)
					break
				}
				if b == arr[len(arr)-1] {
					j := whoIsMin(d, arr)
					// fmt.Println("WHEN: a=", a, " , b=", b, " , j=", j, ", arr[j]=", arr[j])
					arr[j] = a
					d.amount[a] = d.amount[a] - 1
					count++
				}
			}
		}
	}
	fmt.Println(count)
}

func main() {
	d := inputAction()
	getMinCount(d)
}
