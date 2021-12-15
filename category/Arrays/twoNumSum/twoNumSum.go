package twoNumSum

import "sort"

// Brute-force Search
// O(n^2) time | O(1) space
func QuardraticTwoNumberSum(array []int, target int) []int {
	for _, num1 := range array {
		for _, num2 := range array {
			if num1 != num2 && target == num1+num2 {
				return []int{num1, num2}
			}
		}
	}
	return []int{}
}

// O(n) time | O(n) space
func LinearTwoNumberSum(array []int, target int) []int {
	check := map[int]bool{}
	for _, value := range array {
		potentialNum := target - value
		// check[value] = true -> because it shouldn't be added itself
		if _, found := check[potentialNum]; found {
			return []int{potentialNum, value}
		}
		check[value] = true
	}
	return []int{}
}

// O(nlog(n)) time | O(1) space
func LogLinearTwoNumberSum(array []int, target int) []int {
	sort.Ints(array)
	left, right := 0, len(array)-1
	for left < right {
		sum := array[left] + array[right]
		if sum == target {
			return []int{array[left], array[right]}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}
