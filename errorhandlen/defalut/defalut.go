package main

import (
	"bufio"
	"fmt"
	"google_go/errorhandlen/fib"
	"os"
)

func tryDefer()  {
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//fmt.Println(3)
	for i:=0;i<100;i++{
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writerFile(filename string) {
	// deger 确保在函数调用结束时发生 return 之前
	// 参数在 deger 语句是计算
	// deger 列表为 后进先出/先进后出 一样啦
	file,err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()

	// 写到内存
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()

	for i:=0;i<20;i++ {
		fmt.Fprintln(writer,f())
	}

}

func main() {
	tryDefer()
	writerFile("fib.txt")
}
