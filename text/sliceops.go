package text

import "fmt"

func printSlice(s []int) {
	fmt.Println(len(s),cap(s))
}

func main() {
	var s []int // zero value for slice is nil
	for i:= 0;i<10;i++{
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
	//[1 3 5 7 9 11 13 15 17 19]
	printSlice(s)
	//10 16
	s1 := make([]int,10,16)
	printSlice(s1)
	//10 16
	copy(s1, s)
	fmt.Println(s1)
	//[1 3 5 7 9 11 13 15 17 19]

}
