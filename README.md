# golang-starter package

package main
import {}
func main() {
  fmt.Println("Hi")
}
  go run __filename
  
  "FMT" ➡ based on C's printf()
  ```
    %.2f ➡ two decimal place precision
    %-8.1f ➡ "100.6   " "left alignment"
    % x ➡ "01 02 03 04" "spacing"
  ```
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

```
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
```

--------------------------------------
