package main

import (
	"bufio"
	"os"
	"strconv"
)

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
}

var (
	scanner *bufio.Scanner
	wr      = bufio.NewWriter(os.Stdout)
)

func scanInt() int {
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	return x
}

type node struct {
	value int
	link  *node
}

func newNode(value int, link *node) node {
	return node{value: value, link: link}
}

type queue struct {
	front, back *node
	size        int
}

func newQueue() queue {
	return queue{nil, nil, 0}
}

func (q queue) isEmpty() bool {
	return q.front == nil
}

func (q queue) getSize() int {
	return q.size
}

// queue 맨 뒤에 값 삽입
func (q *queue) Enqueue(value int) {
	myNode := newNode(value, nil)
	if q.isEmpty() {
		q.front = &myNode
		q.back = &myNode
	} else {
		q.back.link = &myNode
		q.back = &myNode
	}
	q.size += 1
}

// queue 맨 앞 삭제 - 그 값을 반환
func (q *queue) Dequeue() int {
	if q.isEmpty() {
		return 0
	}
	value := q.front.value
	q.front = q.front.link
	if q.front == nil {
		q.back = nil
	}
	q.size -= 1
	return value
}

func (q queue) Peek() int {
	if q.isEmpty() {
		return 0
	}
	return q.front.value
}

func (q queue) getMaxVal() int {
	if q.isEmpty() {
		return 0
	}
	max := -1
	for i := q.front; i != nil; i = i.link {
		if i.value > max {
			max = i.value
		}
	}
	return max
}

func myWriter(n int) {
	wr.WriteString(strconv.Itoa(n))
	wr.WriteByte('\n')
}

func main() {
	defer wr.Flush()

	testCase := scanInt()
	for i := 0; i < testCase; i++ {
		n, m := scanInt(), scanInt()
		queue := newQueue()
		for j := 0; j < n; j++ {
			queue.Enqueue(scanInt())
		}
		ans, max := 1, queue.getMaxVal()
		for {
			if queue.Peek() != max {
				queue.Enqueue(queue.Dequeue())
			} else {
				if m == 0 {
					myWriter(ans)
					break
				} else {
					queue.Dequeue()
					max = queue.getMaxVal()
					ans += 1
				}
			}
			m -= 1
			if m < 0 {
				m = queue.size - 1
			}
		}
	}
}
