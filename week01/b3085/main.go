package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// myData 구조체에 표준입력을 통해 받은 사탕의 줄의 크기와
// 정렬되어 있는 사탕의 형태를 이차원배열에 저장
type myData struct {
	R *bufio.Reader
	W *bufio.Writer

	num        int
	candySlice [][]string
}

// 표준입력을 통해 값을 받아옴
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

// (i, j) 에 있는 사탕과 (a, b) 에 있는 사탕의 위치를 바꿈
func swap(i, j, a, b int, candySlice [][]string) [][]string {
	candySlice[i][j], candySlice[a][b] = candySlice[a][b], candySlice[i][j]
	return candySlice
}

// 2차원 배열에 정렬되어 있는 사탕을 보고 가장 길게 연속되는 크기를 max 변수에
// 저장하고 매개변수로 받은 maxCount와 비교하여 더 큰 값을 반환
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

// 모든 사탕을 오른쪽(j++)방향과 아래쪽(i++)방향으로 바꾸어 보고
// 각각의 상황에서 먹을 수 있는 최대 사탕 개수를 maxCount에 저장
// Complexity - Time: O(n^2) Quardratic | Space: O(1)
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
