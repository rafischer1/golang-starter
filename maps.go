package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["one"] = 7
	m["two"] = 13

	fmt.Println("map:", m)

	v1 := m["one"]
	fmt.Println("v1: ", v1)

	fmt.Println("length:", len(m))

	delete(m, "two")
	fmt.Println("map:", m)

	_, prs := m["two"]
	fmt.Println("prs:", prs)

	n := map[string]int{"one": 1, "two": 2}
	fmt.Println("map:", n)

}
