// Package clause
package main

// Import statement
import (
	"fmt"
)

// My Code
func main() {

	var name string // var - degisken ismi - değişken tipi
	name = "Ali"

	var age int
	age = 30

	var isSmoking bool = true

	var firstName, lastName = "Ali", "Veli"

	car := "Honda"
	car = "Bmw"

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isSmoking)
	fmt.Println(firstName, lastName)
	fmt.Println(car)

	fmt.Println("------------------")

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isSmoking)

}
