package main

func getNextIndex(curIdx int, arr []int) int {
	jumpDis := arr[curIdx]
	nxtIdx := (curIdx + jumpDis) % len(arr)
	if nxtIdx >= 0 {
		return nxtIdx
	}
	return nxtIdx + len(arr)
}

func HasSingleCycle(array []int) bool {
	count := 0
	curIdx := 0
	for count < len(array) {
		if curIdx == 0 && count > 0 {
			return false
		}
		count++
		curIdx = getNextIndex(curIdx, array)
	}
	return curIdx == 0
}
