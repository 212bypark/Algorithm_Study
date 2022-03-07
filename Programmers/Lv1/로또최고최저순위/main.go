package main

func solution(lottos []int, win_nums []int) []int {
	z_count, e_count := 0, 0
	for _, val := range lottos {
		if val == 0 {
			z_count++
		} else {
			for _, w_num := range win_nums {
				if val == w_num {
					e_count++
				}
			}
		}
	}
	if e_count < 2 {
		if z_count == 0 && e_count == 0 {
			return []int{6, 6}
		}
		return []int{7 - (e_count + z_count), 6}
	} else {
		return []int{7 - (e_count + z_count), 7 - e_count}
	}
}
