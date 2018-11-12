package main

import (
	"errors"
	"fmt"
	"math"
)

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println("Hello World")

	x := 20
	fmt.Println(x)

	a := [5]int{2, 3, 4, 5, 6}
	a[2] = 3
	fmt.Println(a)

	//slices are an abstraction of arrays
	b := []int{2, 3, 4, 5, 6}
	b[2] = 3
	fmt.Println(b)

	vertices := make(map[string]int)
	vertices["one"] = 2
	vertices["two"] = 3
	vertices["three"] = 5

	delete(vertices, "three")

	fmt.Println(vertices)

	//loops
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	//range
	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index", index, "value", value)
	}

	//map[type]type
	m := make(map[string]string)
	m["A"] = "alpha"
	m["B"] = "beta"

	for key, value := range m {
		fmt.Println("key", key, "value", value)
	}

	//calling a function outside the main func scope
	result := sum(2, 3)
	fmt.Println(result)

	// result, err := sqrt(16)
	//
	// if err != nil {
	//   fmt.Println(err)
	// } else {
	//   fmt.Println(result)
	// }

	p := person{name: "Artie", age: 37}
	fmt.Println("Person:", p.name)

	//the pointer to the memory address of z
	z := 7
	inc(&z)
	fmt.Println(z)

}

func sum(x int, y int) int {
	return x + y

}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for neg numbers")
	}
	return math.Sqrt(x), nil
}

//pointers * => &
//dereference the incrementer and point to the original variable
func inc(x *int) {
	*x++
}
