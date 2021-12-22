// Memory: 976KB
// Time : 4ms
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type myData struct {
	R             *bufio.Reader
	mD            int
	assignmentArr [][2]int
}

func inputAction() *myData {
	data := &myData{}
	data.R = bufio.NewReader(os.Stdin)
	num, d, w, mD := 0, 0, 0, 0
	fmt.Fscanln(data.R, &num)
	for i := 0; i < num; i++ {
		fmt.Fscanln(data.R, &d, &w)
		if mD < d {
			mD = d
		}
		temp := [2]int{d, w}
		data.assignmentArr = append(data.assignmentArr, temp)
	}
	data.mD = mD
	sort.Slice(data.assignmentArr, func(i, j int) bool {
		if data.assignmentArr[i][0] == data.assignmentArr[j][0] {
			return data.assignmentArr[i][1] > data.assignmentArr[j][1]
		}
		return data.assignmentArr[i][0] > data.assignmentArr[j][0]
	})
	// fmt.Println(data.assignmentArr)
	return data
}

func popSum(d *myData, index int) int {
	val := 0
	val = d.assignmentArr[index][1]
	if index == 0 {
		d.assignmentArr = d.assignmentArr[index+1:]
	} else if index == len(d.assignmentArr)-1 {
		d.assignmentArr = d.assignmentArr[:index]
	} else {
		d.assignmentArr = append(d.assignmentArr[:index], d.assignmentArr[index+1:]...)
	}
	return val
}

func getMaxScore(d *myData) int {
	maxScore := 0
	for i := d.mD; i > 0; i-- {
		index := -1
		temp := 0
		for j := range d.assignmentArr {
			if d.assignmentArr[j][0] >= i {
				if d.assignmentArr[j][1] > temp {
					temp = d.assignmentArr[j][1]
					index = j
					// fmt.Println("CHECK: j=", j)
				}
			} else {
				break
			}
		}
		if index != -1 {
			maxScore += popSum(d, index)
		}
		fmt.Println(d.assignmentArr)
	}
	// fmt.Println("flag:", d.assignmentArr)
	return maxScore
}

func main() {
	d := inputAction()
	fmt.Println(getMaxScore(d))
}

/*
[3 30] [2 50] [4 40] [4 60] [6 5]

D6(6): [6 5]
D5(5~6): X
D4(4~6): [4 60]
D3(3~6): [4 40]
D2(2~6): [2 50]
D1(1~6): [3 30]
*/
