package main

import "fmt"

func main() {
	m := map[string]int{"apple": 100, "banana": 198}

	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["no"]
	fmt.Println(v2, ok2)

	m["grape"] = 500
	fmt.Println(m)
}
