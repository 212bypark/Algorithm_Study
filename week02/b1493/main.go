package main

import (
	"bufio"
	"os"
)

type myData struct {
	R *bufio.Reader
}

func inputAction() *myData {
	d := &myData{}
	d.R = bufio.NewReader(os.Stdin)

	return d
}
