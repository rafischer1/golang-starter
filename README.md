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
