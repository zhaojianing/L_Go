package text

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s:= "hello易山"  // UTF-8 english1 china3
	for _,b := range []byte(s) {
		fmt.Printf("%X " , b)
	}
	fmt.Println()

	for i,ch := range s { // ch is a rune
		fmt.Printf("(%d %X)",i,ch)
	}
	fmt.Println()

	fmt.Println(utf8.RuneCountInString(s))
	//68 65 6C 6C 6F E6 98 93 E5 B1 B1
	//(0 68)(1 65)(2 6C)(3 6C)(4 6F)(5 6613)(8 5C71)
	//7
}
