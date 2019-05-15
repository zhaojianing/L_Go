package text

import "fmt"

// 值传递 没有引用传递 -> 指针修改
//func swap(a,b *int) {
//	*b, *a = *a, *b
//}
//
//func main() {
//	a,b := 3,4
//	swap(&a, &b)
//	fmt.Println(a,b)
//}

// 换一种方式交换两个的值
func swap(a,b int) (int, int) {
	return b, a
}

func main() {
	a,b := 3,4
	a,b = swap(a, b)
	fmt.Println(a,b)
}