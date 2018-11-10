package main

import (
	"fmt"
)

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	//if else
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	//fizz fizzBuzz
	fizzBuzz(30)

}

func fizzBuzz(x int) {
	for i := 1; i <= x; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzBuzz", i)
		} else if i%5 == 0 {
			fmt.Println("buzz", i)
		} else if i%3 == 0 {
			fmt.Println("fizz", i)
		}
	}
}
