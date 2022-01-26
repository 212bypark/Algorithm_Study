package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
}

var (
	scanner *bufio.Scanner
	wr      = bufio.NewWriter(os.Stdout)
)

func nextInt() int {
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	return x
}

func nextWord() string {
	scanner.Scan()
	return scanner.Text()
}

func nextNum() *[]int {
	var numArr []int
	var strArr []string
	str := nextWord()
	str = strings.ReplaceAll(str, "[", "")
	str = strings.ReplaceAll(str, "]", "")
	strArr = strings.Split(str, ",")
	for _, val := range strArr {
		num, _ := strconv.Atoi(val)
		numArr = append(numArr, num)
	}
	return &numArr
}

func actionLoop(comb string, arr []int) {
	reverse := false
	for _, command := range comb {
		switch command {
		case 'R':
			reverse = !reverse
		case 'D':
			if len(arr) == 0 {
				wr.WriteString("error\n")
				return
			}
			if reverse {
				arr = arr[:len(arr)-1]
			} else {
				arr = arr[1:]
			}
		}
	}

	wr.WriteByte('[')
	if reverse {
		for i := len(arr) - 1; i >= 0; i-- {
			temp := strconv.Itoa(arr[i])
			wr.WriteString(temp)
			if i > 0 {
				wr.WriteByte(',')
			}
		}
	} else {
		for i := 0; i <= len(arr)-1; i++ {
			temp := strconv.Itoa(arr[i])
			wr.WriteString(temp)
			if i < len(arr)-1 {
				wr.WriteByte(',')
			}
		}
	}
	wr.WriteString("]\n")

}

func main() {
	defer wr.Flush()

	testCase := nextInt()
	for i := 0; i < testCase; i++ {
		comb := nextWord()
		n := nextInt()
		arr := make([]int, n)
		inputArr := nextWord()
		if n != 0 {
			temp, index := 0, 0
			for _, c := range inputArr {
				if c >= '0' && c <= '9' {
					temp = 10*temp + int(c-'0')
				} else if c == ',' || c == ']' {
					arr[index] = temp
					index++
					temp = 0
				}
			}
		}
		actionLoop(comb, arr)
	}
}
