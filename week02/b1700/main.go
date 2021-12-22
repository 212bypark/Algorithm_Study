// Memory: 948KB
// Time : 4ms
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	R := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscanln(R, &n, &k)
	input := make([]int, k)
	for i := range input {
		fmt.Fscan(R, &input[i])
	}

	if n >= k {
		fmt.Println(0)
		return
	}

	plugOut := 0
	con := map[int]bool{}
	for i := 0; i < k; i++ {
		if _, check := con[input[i]]; !check {
			if len(con) < n {
				con[input[i]] = true
			} else {
				max, idx := 0, 0
				for plugedIn := range con {
					// fmt.Println("plugedIn: ", plugedIn)
					tempIdx := k
					for j := i + 1; j < k; j++ {
						if plugedIn == input[j] {
							tempIdx = j
							break
						}
					}
					if tempIdx > idx {
						idx = tempIdx
						max = plugedIn
					}
					// fmt.Println("-->max: ", max)
				}
				// fmt.Println("CHECK - max: ", max)
				delete(con, max)
				con[input[i]] = true
				plugOut++
			}
		}
	}
	fmt.Println(plugOut)
}
