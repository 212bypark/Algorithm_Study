package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	R := bufio.NewReader(os.Stdin)
	var N int
	fmt.Fscanln(R, &N)
	arr := make([][2]int64, N<<1)
	for i := 0; i < N; i++ {
		var s, t int64
		fmt.Fscanln(R, &s, &t)
		arr[i<<1] = [2]int64{1, s}
		arr[i<<1+1] = [2]int64{-1, t}
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i][1] == arr[j][1] {
			return arr[i][0] < arr[j][0]
		}
		return arr[i][1] < arr[j][1]
	})

	concurr := int64(0)
	max := int64(0)
	for _, v := range arr {
		concurr += v[0]
		if concurr > max {
			max = concurr
		}
	}
	fmt.Println(max)
}

/*
test case
6
1 3
2 5
7 8
9 10
7 11
4 12
*/
