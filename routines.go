package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*With a goroutine we return immediately
to the next line and don't wait for the
function to complete. This is why the
call to the Scanln function has been included;
without it the program would exit before
being given the opportunity to print all
the numbers.*/

func f(n int) {
	for i := 0; i < 5; i++ {
		fmt.Println("n:", n, ":", "i:", i)
		amt := time.Duration(rand.Intn(500))
		fmt.Println("in f")
		time.Sleep(time.Millisecond * amt)
	}

}

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println("in main")
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}
