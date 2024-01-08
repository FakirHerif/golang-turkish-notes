package main

import (
	"fmt"
	"math"
)

func main() {

	r := 3.0

	const pi float64 = 3.14 // sabitler const anahtar kelimesiyle yazılır
	areaOfCircle := pi * (math.Pow(r, 2))
	fmt.Println(areaOfCircle)

	const x int = 2
	const y float32 = 3.4
	const z string = "test"
	const t bool = false

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", z, z)
	fmt.Printf("%T, %v\n", t, t)

	// Const'lara değer ataması sadece 1 defa olur. Yani const x'e bir değer atadığımız zaman tekrar x'e başka bir değer atayamayız. Örneğin:

	/* Yapamayacağımız işlemler:
	const x = 2
	x = 5
	x++
	x = x +1
	Şeklinde yeni atamalar yapamayız.
	*/

	// const ==> 	compile time -> yazdığımız kodun makine diline çevrilmiş durumu
	// var	 ==> 	run time 	 -> makine diline çevrilmiş kodun çalıştırılması

	var a = math.Pow(3, 4)
	fmt.Printf("%T, %v\n", a, a) // float64, 81 çıktısını alırız

	/* 	const b = math.Pow(3, 4)	// gördüğünüz gibi bu şekilde bir işlem yapamayız hata alırız. işte compile time ve run time farkı tam olarak burada ortaya çıkıyor
	 */

	// ÖNEMLİ NOT: const yani bir sabit oluşturduğumuz zaman ona değer atamalıyız, değer atamadan const kullanamayız. yani değer atamadan bu şekilde kullanım hata verecektir: "const b"

	// Ayrıca oluşturduğumuz bir değişkeni const değişkenine atayamayız. Örnek:

	/* 	w := 3
	   	const h = w 	// bu şekilde oluşturduğumuz w değişkenini const h değişkenine atama yapamayız.
	*/

	const o, u = 3, 5 // şeklinde atama yapabiliriz.

	const q = 1
	var k = 3
	fmt.Printf("%T, %v\n", q+k, q+k) // şeklinde bir kullanım yapabiliriz

}
