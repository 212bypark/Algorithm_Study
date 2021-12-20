package CP

import (
	"sort"
)

func ClassP(rSH []int, bSH []int) bool {
	result := false
	sort.Ints(rSH)
	sort.Ints(bSH)
	if rSH[0] > bSH[0] {
		for i := range rSH {
			if rSH[i] > bSH[i] {
				result = true
			} else {
				result = false
				return result
			}
		}
	} else {
		for i := range rSH {
			if rSH[i] < bSH[i] {
				result = true
			} else {
				result = false
				return result
			}
		}
	}
	return result
}
