package MWT

import (
	"sort"
)

// O(nlog(n)) time | O(1) space
func MinWaitTime(queries []int) int {
	// sort.Sort(sort.IntSlice(queries))
	sort.Ints(queries)
	total, sum := 0, 0
	for _, wt := range queries {
		total += sum
		sum += wt
	}
	return total
}
