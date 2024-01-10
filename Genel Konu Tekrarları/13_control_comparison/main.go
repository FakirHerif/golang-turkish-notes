/*
Comparison Operators
== 		Equal
!= 		Not Equal
< 		Less than
> 		Greater than
<= 		Less than or equal
<= 		Greater than or equal

Logical Operators
&& 	conditional 	AND 	p && q 	is 	"if p then q else false"
|| 	conditional 	OR 		p || 	is 	"if p then true else q"
! 					NOT 	!p 		is 	"not p"
*/

package main

import "fmt"

func main() {

	x, y := 3, 7
	fmt.Printf("%T, %v\n", x == y, x == y) // bool, false
	fmt.Printf("%T, %v\n", x != y, x != y) // bool, true
	fmt.Printf("%T, %v\n", x < y, x < y)   // bool, true
	fmt.Printf("%T, %v\n", x > y, x > y)   // bool, false
	fmt.Printf("%T, %v\n", x >= y, x >= y) // bool false
	fmt.Printf("%T, %v\n", x <= y, x <= y) // bool true

	fmt.Println("*************************")

	a, b := "a", "b"                       // decimal olarak a = 97 ve b = 98 olarak algılar
	fmt.Printf("%T, %v\n", a == b, a == b) // bool, false
	fmt.Printf("%T, %v\n", a != b, a != b) // bool, true
	fmt.Printf("%T, %v\n", a < b, a < b)   // bool, true
	fmt.Printf("%T, %v\n", a > b, a > b)   // bool, false
	fmt.Printf("%T, %v\n", a >= b, a >= b) // bool false
	fmt.Printf("%T, %v\n", a <= b, a <= b) // bool true

	// karşılaştırma operatörleri aynı veri tipinde olmalıdır. Örneğin int ile float64'ü veya string'i karşılaştıramayız.

	fmt.Println("*************************")

	q, w := 15, 20
	set1 := (q == w)                   //false
	set2 := (q < w)                    //true
	fmt.Printf("%v\n", (set1 && set2)) // false ===>> ve karşılaştırma operatöründe değerlerden herhangi bir false ise sonuç false olarak döner, eğer iki ifade de true ise sonuç true döner. Kısaca sadece 2 durum true ise sonuç true
	fmt.Printf("%v\n", (set1 || set2)) // true ===>> veya karşılaştırma operatöründe değerlerden herhangi biri true ise sonuç true olarak döner, eğer iki ifade de false ise sonuç false döner. Kısaca sadece 2 durum false ise sonuç false

	// mantıksal ifadelerde de aynı veri tipi olmalıdır, yani bool ile bool olmalıdır.

	fmt.Printf("%v\n", (!set1)) // true ===>> tam tersi olan sonuç döner. set1 normalde false ve değildir kullanarak sorgularsak tersi olan true sonucu döner
}
