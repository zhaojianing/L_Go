package main

import "fmt"

func main() {
	a := 60
	call := ""
	switch  {
		case a < 60:
			call = "E"
		case a < 70:
			call = "D"
		case a < 80:
			call = "C"
		case a < 90:
			call = "B"
		case a < 100:
			call = "A"
		default:
			panic(fmt.Sprintf("Error: %d", a))
	}
	fmt.Println(call)
}
