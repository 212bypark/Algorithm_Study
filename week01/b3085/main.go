package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type myData struct {
	R *bufio.Reader
	W *bufio.Writer

	num        int
	candySlice [][]string
}

func inputAction() *myData {
	d := &myData{}
	d.R = bufio.NewReader(os.Stdin)
	d.W = bufio.NewWriter(os.Stdout)

	fmt.Fscanln(d.R, &d.num)
	for i := 0; i < d.num; i++ {
		input, _ := d.R.ReadString('\n')
		inputArray := strings.Split(strings.ReplaceAll(input, "\n", ""), "")
		d.candySlice = append(d.candySlice, inputArray)
	}
	return d
}

func swap(i, j, a, b int, candySlice [][]string) [][]string {
	candySlice[i][j], candySlice[a][b] = candySlice[a][b], candySlice[i][j]
	return candySlice
}

func checkMax(maxCount int, candySlice [][]string) int {
	max := 0
	previous := ""
	// compare each row value
	for i := 0; i < len(candySlice); i++ {
		count := 0
		for j := 0; j < len(candySlice); j++ {
			current := candySlice[i][j]
			if previous == current {
				count++
				if max < count {
					max = count
				}
			} else {
				count = 1
			}
			previous = current
		}
	}
	// compare each column value
	previous = ""
	for i := 0; i < len(candySlice); i++ {
		count := 0
		for j := 0; j < len(candySlice); j++ {
			current := candySlice[j][i]
			if previous == current {
				count++
				if max < count {
					max = count
				}
			} else {
				count = 1
			}
			previous = current
		}
	}
	if max < maxCount {
		max = maxCount
	}
	return max
}

func countMax(d *myData) int {
	maxCount := 0
	for i := 0; i < d.num; i++ {
		for j := 0; j < d.num; j++ {
			if j != d.num-1 {
				d.candySlice = swap(i, j, i, j+1, d.candySlice)
				maxCount = checkMax(maxCount, d.candySlice)
				d.candySlice = swap(i, j, i, j+1, d.candySlice)
			}
			if i != d.num-1 {
				d.candySlice = swap(i, j, i+1, j, d.candySlice)
				maxCount = checkMax(maxCount, d.candySlice)
				d.candySlice = swap(i, j, i+1, j, d.candySlice)
			}
		}
	}
	return maxCount
}

func main() {
	d := inputAction()
	defer d.W.Flush()
	ans := countMax(d)
	fmt.Println(ans)
}
