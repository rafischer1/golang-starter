package main

import "fmt"

type Animal struct {
	name   string
	region string
	age    int
}

func (a Animal) animalFunc() string {
	return "The " + a.name + " lives in " + a.region + "."
}

func main() {
	fmt.Println("In structs")

	tiger := Animal{"Bengal Tiger", "India", 4}
	gorilla := Animal{"Gorilla", "Sub-Saharan Africa", 12}

	fmt.Println("gorilla:", gorilla)
	fmt.Println("gorilla's region:", gorilla.region)

	ta := &tiger
	fmt.Println(ta.animalFunc())

}
