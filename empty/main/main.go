package main

import (
	"fmt"
	"golang.org/x/tools/container/intsets"
	"google_go/empty"
)

func main() {
	s := empty.Shan{1}
	s.Push(2)
	fmt.Println(s)

	a := intsets.Sparse{}
	fmt.Println(a.Insert(1))
}
