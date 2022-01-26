package main

import "fmt"

func action(x int) int {
	if x == 1 {
		return 1
	} else if x == 2 {
		return 2
	} else {
		if x%2 == 0 {
			return action(x/2) * 2
		} else {
			return (action(x/2+1) - 1) * 2
		}
	}
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	fmt.Printf("%d\n", action(n))
}
