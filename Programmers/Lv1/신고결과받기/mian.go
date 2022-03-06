package main

import (
	"strings"
)

type id_report struct {
	id         string
	report_ids []string
	count      int
}

func is_exist(report_ids []string, id string) bool {
	for _, check_id := range report_ids {
		if check_id == id {
			return true
		}
	}
	return false
}

func solution(id_list []string, report []string, k int) []int {
	data := make([]id_report, len(id_list))
	ret := make([]int, len(id_list))
	idx := 0
	for _, id := range id_list {
		data[idx].id = id
		data[idx].count = 0
		idx++
	}
	for _, str := range report {
		val := strings.Split(str, " ")
		for i := 0; i < len(data); i++ {
			if data[i].id == val[0] && !is_exist(data[i].report_ids, val[1]) {
				data[i].report_ids = append(data[i].report_ids, val[1])
				for j := 0; j < len(data); j++ {
					if data[j].id == val[1] {
						data[j].count++
						break
					}
				}
			}
		}
	}
	for i := 0; i < len(id_list); i++ {
		temp := 0
		for _, name := range data[i].report_ids {
			for j := 0; j < len(id_list); j++ {
				if data[j].id == name && data[j].count >= k {
					temp++
				}
			}
		}
		ret[i] = temp
	}
	return ret
}
