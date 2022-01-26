package main

import (
	"bufio"
	"container/list"
	"os"
	"strconv"
)

var (
	scanner *bufio.Scanner
	queue   = list.New()
	wr      = bufio.NewWriter(os.Stdout)
)

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
}

func scanInt() int {
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	return x
}

func Size() {
	wr.WriteString(strconv.Itoa(queue.Len()))
	wr.WriteByte('\n')
}

func Empty() {
	if queue.Len() == 0 {
		wr.WriteString("1")
	} else {
		wr.WriteString("0")
	}
	wr.WriteByte('\n')
}

func PushFront(val int) {
	queue.PushFront(val)
}

func PushBack(val int) {
	queue.PushBack(val)
}

func PopFront() {
	if queue.Len() == 0 {
		wr.WriteString("-1")
		wr.WriteByte('\n')
		return
	}
	front := queue.Front()
	wr.WriteString(strconv.Itoa(front.Value.(int)))
	wr.WriteByte('\n')
	queue.Remove(front)
}

func PopBack() {
	if queue.Len() == 0 {
		wr.WriteString("-1")
		wr.WriteByte('\n')
		return
	}
	back := queue.Back()
	wr.WriteString(strconv.Itoa(back.Value.(int)))
	wr.WriteByte('\n')
	queue.Remove(back)
}

func Front() {
	front := queue.Front()
	if front == nil {
		wr.WriteString("-1")
	} else {
		wr.WriteString(strconv.Itoa(front.Value.(int)))
	}
	wr.WriteByte('\n')
}

func Back() {
	back := queue.Back()
	if back == nil {
		wr.WriteString("-1")
	} else {
		wr.WriteString(strconv.Itoa(back.Value.(int)))
	}
	wr.WriteByte('\n')
}

func main() {
	defer wr.Flush()

	n := scanInt()
	for i := 0; i < n; i++ {
		scanner.Scan()
		switch scanner.Text() {
		case "push_front":
			PushFront(scanInt())
		case "push_back":
			PushBack(scanInt())
		case "pop_front":
			PopFront()
		case "pop_back":
			PopBack()
		case "front":
			Front()
		case "back":
			Back()
		case "size":
			Size()
		case "empty":
			Empty()
		}
	}
}
