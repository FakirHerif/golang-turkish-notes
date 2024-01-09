/*
1. x = x - 10 vs -=10
2. K = F - 32 / 1.8 + 273 => -40 F kaç K derecedir?
3.  age := 40
	const myAge = age
	fmt.Printf("%v, %T", myAge, myAge)
4. Sabitlerde shadowing kavramı çalışır mı?
5. const w = 4, const q = 5.4 , w + q?
6. const g float64 = 6.4 , h := 4 + g , bu durumda h ne olur?
*/

package main

import "fmt"

func main() {
	// 1
	x := 50
	x = x - 10                   // assignment statement -> sağdaki işlemi yani statement'ı sola x değerine assign ediyoruz yani x değerine atıyoruz.
	fmt.Printf("%T, %v\n", x, x) // int, 40
	// Kısa gösterimi ise:
	x -= 10                      // assignment operation ("+=", "*=", "/=" şeklinde diğer işlemler içinde kullanabiliriz)
	fmt.Printf("%T, %v\n", x, x) // int, 30

	// 2
	F := -40
	K := float64(F-32)/1.8 + 273
	fmt.Printf("%T, %v\n", K, K) // float64, 233

	// 3
	/*
		age := 40
		const myAge = age // hata alırız. Const compile time, variable runtime'a aittir.

		fmt.Printf("%v, %T", myAge, myAge)
	*/

	// 4
	// Shadowing: örneğin paket kapsamında yani global kapsamda var x = 14 tanımladık ve fonksiyon içinde ise var x = 24 tanımladık diyelim, bu durumda x'in değerini yazdırmak istediğimiz zaman bize 24 sonucu dönecektir. İşte bu durumda fonksiyonun içinde yer alan değişken, paket kapsamında yani global tanımladığımız değişkeni gölgelemiş oluyor. Bu durumda const tanımlasak bile sonuç aynı olur, shadowing kavramı constlarda da aynı şekilde çalışır.

	// 5
	const w = 4                      // typeless
	const q = 5.4                    // typeless
	fmt.Printf("%T, %v\n", w+q, w+q) // float64, 9.4  ---> typeless değişkenlerde go veri tiplerini varsayılan olarak otomatik atıyor. toplama işlemi gibi durumda ise uygun olan veri tipini atıyor.

	// 6
	const g float64 = 6.4
	h := 4 + g
	fmt.Printf("%T, %v\n", h, h) // float64, 10.4

}
