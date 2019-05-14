package main

import "fmt"

func updateSlice(arr []int) {
	arr[0] = 100
}

func main()  {
	arr := [...]int{0,1,2,3,4,5,6,7}
	fmt.Println("dataArray:",arr[2:6])
	//dataArray: [2 3 4 5]
	fmt.Println("dataArray:",arr[2:])
	//dataArray: [2 3 4 5 6 7]
	fmt.Println("dataArray:",arr[:6])
	//dataArray: [0 1 2 3 4 5]
	fmt.Println("dataArray:",arr[:])
	//dataArray: [0 1 2 3 4 5 6 7]

	//Reslice
	a1 := arr[2:]
	updateSlice(a1)
	fmt.Println(a1)
	//[100 3 4 5 6 7]
	fmt.Println(arr)
	//[0 1 100 3 4 5 6 7]

	// slice 底层view 向后扩展
	//			    [3,6]
	//             [3,4]
	//arr := o,1,2,3,4,5,6,7
}
