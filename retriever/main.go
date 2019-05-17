package main

import (
	"fmt"
	"google_go/retriever/mock"
	"google_go/retriever/reals"
	"time"
)

type Reteriever interface {
	Get(url string) string
}

func download(r Reteriever) string {
	return r.Get("https://ladies.ren/")
}

func main() {
	var r Reteriever
	// copy
	r = mock.Reteriever{"this is 易山博客"}
	// -> r
	r = &reals.Reteriever{
		"Mozilla/5.0",
		time.Minute,
	}

	inspect(r)
	// Type assertion
	realsReteriever := r.(*reals.Reteriever)
	fmt.Println("Type assertion: ",realsReteriever)
	
	//fmt.Println(download(r))
}

func inspect(r Reteriever)  {
	fmt.Println("%T %v\n",r,r)
	switch v:= r.(type) {
	case mock.Reteriever:
		fmt.Println("mock ",v.Contents)
	case *reals.Reteriever:
		fmt.Println("reals ",v.UserAgent)
	}
}
