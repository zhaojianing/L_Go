package main

import "fmt"

// if else
func main() {
	aa := 1
	bb := 2
	// 第一种写法
	// err := true
	// if err {
	// 	fmt.Println(aa)
	// } else {
	// 	fmt.Println(bb)
	// }
	// 第二种写法(写法比较傻haha。。。)
	if err := true; err {
		fmt.Println(aa)
	} else {
		fmt.Println(bb)
	}
}
