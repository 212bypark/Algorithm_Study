package main

import (
	"bufio"
	"fmt"
	"os"
)

type myData struct {
	R *bufio.Reader
	W *bufio.Writer

	num      int
	eachTree []tree
}

type tree struct {
	x     int
	y     int
	fence int
}

func inputAction() *myData {
	d := &myData{}
	d.R = bufio.NewReader(os.Stdin)
	d.W = bufio.NewWriter(os.Stdout)

	fmt.Fscanln(d.R, &d.num)
	d.eachTree = make([]tree, d.num)
	for i := 0; i < d.num; i++ {
		fmt.Fscanln(d.R, &d.eachTree[i].x, &d.eachTree[i].y, &d.eachTree[i].fence)
	}
	return d
}

func getMin(num1, num2 int) int {
	if num1 > num2 {
		return num2
	} else {
		return num1
	}
}

func getMax(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func howManyFence(myTree []tree) int {
	countFence := 0
	xMin := 1000000
	xMax := 0
	for i := 0; i < len(myTree); i++ {
		xMin = getMin(myTree[i].x, xMin)
		xMax = getMax(myTree[i].x, xMax)
	}
	yMin := 1000000
	yMax := 0
	for i := 0; i < len(myTree); i++ {
		yMin = getMin(myTree[i].y, yMin)
		yMax = getMax(myTree[i].y, yMax)
	}

	countFence = 2*(xMax-xMin) + 2*(yMax-yMin)
	return countFence
}

func makeSubSlice(myTree []tree, num int) ([]tree, int) {
	tempTree := myTree
	selected := 0

	return tempTree, selected
}

func checkAvail(d *myData, num int) int {
	countTree := d.num
	for i := 0; i < d.num; i++ {
		restTree, selected := makeSubSlice(d.eachTree, i)
		if howManyFence(restTree) <= d.eachTree[selected].fence {

		}

	}
	return countTree
}

func main() {
	d := inputAction()
	fmt.Println(d.eachTree)
}
