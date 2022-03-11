package main

import "fmt"

type buildingStack struct {
	heights   map[int]int
	direction int
	ret       []int
}

func newBuildingStack() *buildingStack {
	stack := &buildingStack{}
	stack.heights = make(map[int]int)
	return stack
}

func solution(stack *buildingStack) {
	i := 0
	if stack.direction == 1 {

	} else if stack.direction == -1 {

	}
}

func pop(stack *buildingStack, idx int) {
	delete(stack.heights, idx)
}

func SunsetViews(buildings []int, direction string) []int {
	stack := newBuildingStack()
	if direction == "EAST" {
		stack.direction = 1
	} else if direction == "WEST" {
		stack.direction = -1
	}
	for i := 0; i < len(buildings); i++ {
		stack.heights[i] = buildings[i]
	}

	return []int{}
}

func main() {
	testCase1 := []int{3, 5, 4, 4, 3, 1, 3, 2}
	fmt.Println(SunsetViews(testCase1, "EAST"))
}
