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

assignment: x := 20
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
