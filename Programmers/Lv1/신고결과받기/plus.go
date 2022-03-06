import (
	"string"
)

func solution(id_list []string, report []string, k int) []int {
	arr := map[string]map[string]bool{}
	ret := make([]int, len(id_list))

	for _, str := range report {
		cases := string.Split(report, " ")
		if _, ok := arr[cases[1]]; !ok {
			arr[cases[1]] = make(map[string]bool)
		}
		arr[cases[1]]cases[0] = true
	}

	for _, val := range arr {
		if len(val) >= k {
			for k, _ := range(val) {
				for idx, n := range(id_list) {
					if n == k {
						ret[idx]++
					}
				}
			}
		}
	}
	return ret
}