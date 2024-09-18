package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	delete(m, "two")
	fmt.Println(m["two"]) // 0

	for k, v := range m {
		println(k, v)
	}

	if v, ok := m["two"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("key not found")
	}

}
