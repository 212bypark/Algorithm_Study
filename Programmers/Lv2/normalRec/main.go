package main

func solution(w int, h int) int64 {
	var answer int64 = 0
	for i := 1; i < w; i++ {
		answer += int64(i * h / w)
	}
	answer *= 2
	return answer
}
