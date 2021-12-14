package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type myData struct {
	R *bufio.Reader
	W *bufio.Writer

	num        int
	inputSlice []int
	trioArray  []int
}

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

func makeTrioArray(input int, d *myData) {
	val := 0
	for val*val < 2*input {
		val++
	}
	// fmt.Println("input", input)
	// fmt.Println("value", val)
	// d.trioArray = make([]int, val)
	for i := 1; i < val; i++ {
		d.trioArray = append(d.trioArray, i*(i+1)/2)
		// fmt.Println("checking:", d.trioArray)
	}
	// fmt.Println("checking:", d.trioArray)
}

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
		// fmt.Println("inputslice", i, ":", d.inputSlice[i])
		makeTrioArray(d.inputSlice[i], d)
		// fmt.Println("trioArray:", d.trioArray)
		checkAvail(d.inputSlice[i], d.trioArray)
	}
}

func main() {
	getReady()
}
