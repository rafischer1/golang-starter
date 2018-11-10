package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("go" + "lang")

	//like console.log with the print out
	fmt.Println("1+1 =", 1+1)

	//true/false
	fmt.Println("true && false:", true && false)
	fmt.Println("true || false:", true || false)
	fmt.Println("!true:", !true)

	// variable declaration
	var d = true
	e := false
	fmt.Println("d, e:", d, e)

	var b, c int = 1, 2
	fmt.Println(b, c)

	const s string = "constant"
	fmt.Println(s)

	const n = 500000000
	const m = 3e20 / n
	fmt.Println(m)
	fmt.Println(int64(m))
	fmt.Println(math.Sin(n))

}
