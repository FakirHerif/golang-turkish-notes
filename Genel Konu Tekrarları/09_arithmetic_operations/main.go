/*
addition 		+  ==> 15, 10 -> operand / + -> operator
substruction 	-
product 		*
division 		/
remainder 		%
increment 		++
decrement		--
*/

package main

import "fmt"

func main() {

	x, y := 15, 10

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)

	fmt.Printf("%T, %v\n", (x + y), (x + y)) // 25
	fmt.Printf("%T, %v\n", (x - y), (x - y)) // 5
	fmt.Printf("%T, %v\n", (x * y), (x * y)) // 150

	fmt.Printf("%T, %v\n", (x / y), (x / y)) // 1 sonucunu veriyor ama bizim beklediğimiz sonuç 1,5 ama burada veri tipi int olduğu için bize sadece 1 sonucu dönüyor. İki int sonucu bize sonuçta int olarak dönüyor bu yüzden ,5 kısmı atılıyor ve 1 sonucu dönüyor

	z := (5.0 / 2)

	fmt.Printf("%T, %v\n", z, z) // veri tipimizi float64'e döndürüyor ve 2.5 sonucunu alıyoruz ancak 5/2 float64 yazsak bile sonuç 2 verecekti, 5.0 ondalıklı belirtmemiz önemli

	fmt.Printf("%T, %v\n", (x % y), (x % y)) // kalan 5 olarak sonucumuz döner.

	a := 10
	fmt.Println(a) // 10
	a = a + 1
	fmt.Println(a) // 11
	a++
	fmt.Println(a) // 12

	b := 20
	fmt.Println(b) // 20
	b = b - 1
	fmt.Println(b) // 19
	b--
	fmt.Println(b) // 18

	// ÖNEMLİ NOT: DECREMENT VE INCREMENT GO Dilinde bir statement'dır ve bir satırda yalnızca bir statement bulunur. bu yüzden println içinde a++ yapmamız bize hata döndürecektir. Örneğin:

	// fmt.Println(a++)  ==> bu şekilde kullanılamaz ve hata verir bu yüzden yukarıdaki örnekteki gibi ayrı satırda a++ yazıp yeni değeri println içinde döndürüyoruz.

	/*
		STATEMENT VE EXPRESSION ARASINDAKİ FARK NEDİR?

		w = 5 + 3 // Bu bir statement'dır çünkü x değişkenine bir değer atar.

		5 + 3 // Bu bir expression'dur çünkü 5 ile 3'ü toplar ve sonucu döndürür.
	*/

}
