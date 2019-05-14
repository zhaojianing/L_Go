package main

import "fmt"

func printArray(arr [5]int) {
	// 值类型数组
	for i, v:= range arr{
		fmt.Println(i,v)
	}
}

func main()  {
	// 数组值类型
	var arr1 [5]int
	arr2 := [3]int{1,2,3}
	// ... 编译器来决定个数
	arr3 := [...]int{2,4,6,8,10}
	fmt.Println(arr1,arr2,arr3)
	// 二维数组
	// 4个数组，5个内容
	var grid [4][5]int
	fmt.Println(grid)
	//[0 0 0 0 0] [1 2 3] [2 4 6 8 10]
	//[[0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0]]

	// 测试循环
	//for i:= 0; i < len(arr3);i++ {
	//	fmt.Println(arr3[i])
	//}
	//2
	//4
	//6
	//8
	//10

	// or

	// range 获取数组下标
	// 可以通过   _  省略变量
	// 任何地方都可以 使用 _ 省略变量
	//for i, v:= range arr3{
	//	fmt.Println(i,v)
	//}
	//0 2
	//1 4
	//2 6
	//3 8
	//4 10


	// range
	// 1. 美观，意义明确
	// 2. c++：没有类似功能
	// 3. java/python：只能for each value，不能同时获得 i，v

	printArray(arr1)
}