package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
)

type Data struct {
	R *bufio.Reader
	num int
	candySlice [][]int
}

func InputAction() *Data {
	d := &Data{}
	d.R = bufio.NewReader(os.Stdin)
	str, _, _ := d.R.ReadLine()
	d.num, _ = strconv.Atoi(string(str))
	// fmt.Fscanln(d.R, &d.num)
	fmt.Println("num:", d.num)
	var candyArray [][] int
	for i := 0; i < d.num; i++ {
		// fmt.Fscan(d.R, candyArray[i][:])
		line, _, _ := d.R.ReadLine()
		for j := 0; j < d.num; j++ {
			val := 0
			fmt.Fscan(string(line), val)
			candyArray = append(candyArray, val)
		}
		fmt.Println("->", candyArray)
	}
	return d
}

func main() {
	d := InputAction()
	fmt.Println(d.candySlice)
}