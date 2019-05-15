package text

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	// byte 英文字符 阿斯克码
	// rune 中文转 1(改进）
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i,ch := range []rune(s) {

		if lastI,ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return  maxLength
}

func main() {
	// 寻找最长不含有重复字符的子串
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("b"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcdefg"))
	fmt.Println(lengthOfNonRepeatingSubStr("易山博客测试"))
	fmt.Println(lengthOfNonRepeatingSubStr("一一一二"))
	fmt.Println(lengthOfNonRepeatingSubStr("黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))
}
