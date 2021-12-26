package main

import (
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func scan() string {
	sc.Scan()
	return sc.Text()
}
func scanInt() int {
	n, _ := strconv.Atoi(scan())
	return n
}

var p [][]byte

func main() {
	sc.Split(bufio.ScanWords)
	n := scanInt()
	p = make([][]byte, n)
	for i := 0; i < n; i++ {
		p[i] = []byte(scan())
	}
	all(0, 0, n)
	wr.Flush()
}

func all(a, b, n int) {
	c := p[a][b]
	for i := a; i < a+n; i++ {
		for j := b; j < b+n; j++ {
			if p[i][j] != c {
				wr.WriteByte('(')
				all(a, b, n/2)
				all(a, b+n/2, n/2)
				all(a+n/2, b, n/2)
				all(a+n/2, b+n/2, n/2)
				wr.WriteByte(')')
				return
			}
		}
	}
	wr.WriteByte(c)
}
