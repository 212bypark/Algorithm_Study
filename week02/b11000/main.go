// Memory: 14240KB
// Time : 336ms
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// 모든 강의 최소 강의실 수
// 사작시간과 종료시간 모두 오름차순 정렬 후
// 강의실 개수 카운트
// 시작시간 : 1 , 종료시간 : -1
// 주의 : 시작시간과 종료시간이 같다면 종료시간이 먼저!!
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
	// fmt.Println(arr)
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
