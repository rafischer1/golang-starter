# golang-starter package

package main
import {}
func main() {
  fmt.Println("Hi")
}
  go run __filename
________________________________

variables:
name string
age  int

// variable declaration
var d = true
e := false
fmt.Println("d, e:", d, e)  // d, e: true false
var b, c int = 1, 2


-----------------------------------
values:

-----------------------------------
libraries: 	
  "errors"
	"fmt"
	"math"
  
-----------------------------------
pointers * => &
dereference the incrementer and point to the original variable
func inc(x *int) {
	*x++
}

________________________________
for loops:

for j := 7; j <= 9; j++ {
  fmt.Println(j)
}

i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }
    

-------------------------------------

fizzBuzz:

"func fizzBuzz(x int) { \n
	for i := 1; i <= x; i++ { \n
		if i%3 == 0 && i%5 == 0 { \n
			fmt.Println("fizzBuzz", i) \n
		} else if i%5 == 0 { \n
			fmt.Println("buzz", i) \n
		} else if i%3 == 0 { \n
			fmt.Println("fizz", i) \n
		}
	}
}"

--------------------------------------
