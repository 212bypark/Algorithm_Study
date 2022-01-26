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

func nextWord() string {
	scanner.Scan()
	return scanner.Text()
}

type Deque struct {
	list *list.List
}

func newDeque() *Deque {
	return &Deque{list: list.New()}
}

func (dq *Deque) push(s int) {
	dq.list.PushBack(s)
}

func (dq *Deque) pop() int {
	return dq.list.Remove(dq.list.Back()).(int)
}

func (dq *Deque) popLeft() int {
	return dq.list.Remove(dq.list.Front()).(int)
}

func (dq *Deque) front() int {
	return dq.list.Front().Value.(int)
}

func (dq *Deque) back() int {
	return dq.list.Back().Value.(int)
}

func (dq *Deque) len() int {
	return dq.list.Len()
}

func main() {
	defer wr.Flush()

	n, k := scanInt(), scanInt()
	queues := make([]*Deque, 21, 21)
	for i := 2; i < 21; i++ {
		queues[i] = newDeque()
	}
	ans := 0
	var name string
	for i := 0; i < n; i++ {
		name = nextWord()
		l := len(name)
		if queues[l].len() == 0 {
			queues[l].push(i)
		} else {
			for queues[l].len() != 0 && i-queues[l].front() > k {
				queues[l].popLeft()
			}
			ans += queues[l].len()
			queues[l].push(i)
		}
	}
	wr.WriteString(strconv.Itoa(ans))
	wr.WriteByte('\n')
}
