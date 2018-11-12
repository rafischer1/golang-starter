package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays!")
	var a [3]int
	fmt.Println("emp:", a)
	fmt.Println("length a:", len(a))

	var b [5]int
	b[4] = 100
	fmt.Println("set:", b[4])

	c := [5]int{1, 2, 3, 4, 5}
	fmt.Println("set c:", c)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
