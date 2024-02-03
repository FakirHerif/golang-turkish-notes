// örneğin bizim açımızdan önemli olan bir method belirleyelim ve bu method yazmak olsun. Bu yazmak methodunu kurşun kalemle veya tükenmez kalemle yazmak bizim için farketmez, bizim için önemli olan yazmak methodunu uygulamaktır. İşte bu duruma interfaces kavramı diyebiliriz

package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	a, b float64
}

func (r rectangle) area() float64 {
	return r.a * r.b
}

func (r rectangle) circumference() float64 {
	return 2 * (r.a + r.b)
}

type circle struct {
	r float64
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c circle) circumference() float64 {
	return 2 * math.Pi * c.r
}

// kendi interface'imizi oluşturalım ve yukarıdan ilgili methodlarımızı alalım.
type shape interface { // interface methodlarla ilgilenir.
	area() float64          // burda bu methodların imzasına sahip oluyoruz - signature
	circumference() float64 // burda bu methodların imzasına sahip oluyoruz - signature
}

func interfaceFunc(i shape) {
	fmt.Println(i)                 // rectangle için = {3 8} --- circle için ={5}
	fmt.Println(i.area())          // rectangle için = 24 --- circle için = 78.53981633974483
	fmt.Println(i.circumference()) // rectangle için = 22 --- circle için = 31.41592653589793
	fmt.Printf("%T\n", i)          // rectangle için = main.rectangle --- circle için = main.circle
	fmt.Println("**********")
}

func main() {

	/* 	r1 := rectangle{3, 8}
	   	fmt.Println("Area: ", r1.area())                   // Area:  24
	   	fmt.Println("Circumference: ", r1.circumference()) // Circumference:  22

	   	fmt.Println("**********") */

	r1 := rectangle{3, 8}
	interfaceFunc(r1)

	c1 := circle{5}
	interfaceFunc(c1)
}

// interface bize kullanılacak methodları belirtiyor. Methodların her biri kendine ait yani interface sadece kullanılacak methodu belirtiyor, sıfırdan oluşturmuyor.

// UYGULANAN METHOD İSİMLERİ AYNI ANCAK UYGUNALIŞ BİÇİMLERİ FARKLI... AREA HEM DAİRE HEM DİKTÖRGEN İÇİN AYNI İSİMDE ANCAK UYGULANIŞ BİÇİMLERİ FARKLI
