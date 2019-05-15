package text

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a , b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	default:
		panic("error:" + op)
	}
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("call function %s with args (%d, %d)",opName,a,b)
	return op(a, b)
}

func div(a,b int) (c,d int) {
	return a / b, a % b
}

func pow(a,b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	fmt.Println(eval(1,2,"+"))
	c, d := div(23,5)
	fmt.Println(c, d)
	fmt.Println(apply(pow,3,4))
	fmt.Println(sum(1,2,3,4,5))
}