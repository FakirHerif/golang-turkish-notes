// Polymorphism = çok biçimlilik. Bir işlemin farklı şekillerde yapılması olarak düşünebiliriz.

package main

import "fmt"

type shape interface { // interface'imiz
	area() float64
}

func printArea(shapes ...shape) {
	for _, shape := range shapes {
		fmt.Println("Alan: ", shape.area())
	}
}

type triangle struct { // üçgen için
	a float64
	h float64
}

func (t triangle) area() float64 {
	return (t.a * t.h) / 2 // üçgenin alanı
}

type square struct { // kare için
	a float64
}

func (s square) area() float64 {
	return (s.a * s.a) // karenin alanı
}

type rectangle struct { // dikdörtgen için
	a, b float64
}

func (r rectangle) area() float64 {
	return (r.a * r.b) // dikdörtgenin alanı
}

func main() {

	t := triangle{3, 4}
	s := square{5}
	r := rectangle{4, 5}

	printArea(t, s, r)
	/*
		Alan:  6
		Alan:  25
		Alan:  20
	*/
}
