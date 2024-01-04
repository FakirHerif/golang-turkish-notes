package main

import "fmt"

func main() {

	/* 	var name string = "Ali"
	   	var name2 = "Veli"
	   	name3 := "Can" // Short Declaration

	   	var age float32 = -12.5
	   	var isSmoking bool = true

	   	fmt.Println(name, name2, name3)
	   	fmt.Printf("%T\n", age)
	   	fmt.Printf("%T\n", isSmoking)
	   	fmt.Println(age) */

	/* 	var (
		name      string = "Ali"
		age       int    = 30
		isSmoking bool   = true

		weight float32 = 85
		height int     = 183
	) */

	/* var name, age, isSmoking, weight, height = "Ali", 30, true, 85, 182 */

	/* name, age, isSmoking, weight, height := "Ali", 30, true, 85, 182 */

	var name string
	var age int
	var isSmoking bool
	fmt.Println(name)      // "" BOŞ STRING GELİR. Zero Values
	fmt.Println(age)       // 0 SIFIR INT GELİR.
	fmt.Println(isSmoking) // FALSE BOOL GELİR.

	/*
		fmt.Println(weight)
		fmt.Println(height) */

}
