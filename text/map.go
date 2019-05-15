package text

import "fmt"

func main() {
	m := map[string]string{
		"name": "shan_Yi",
		"age": "1997",
		"sex": "男",
	}
	fmt.Println(m)

	m2:=make(map[string]int) // empty map
	var m3 map[string]int // nil
	fmt.Println(m2,m3)
	//map[] map[]

	for k,v := range m {
		fmt.Println(k,v)
	}
	// map hash map 无序

	// 如果没有存在key 则为 zero value
	if causeName, ok := m["colors"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("delete values")
	name, ok := m["name"]
	fmt.Println(name,ok)
	delete(m,"name")
	name, ok = m["name"]
	fmt.Println(name,ok)
}
