// fonksiyon en basit tanımıyla; belirli bir işlemi yapmak için kullanılan kod bloklarıdır.

// fonksiyonların ilk karakteri harf ile başlamalıdır. camelCase kullanılır. Paket içi kullanımda ilk harf küçük, paket dışı kullanımlarda ise ilk harf büyük olarak yazılır.

// func funcName(params) return type { codes }

package main

import "fmt"

func main() {

	/* 	var x, y, sum int
	   	x = 5
	   	y = 10
	   	sum = x + y
	   	fmt.Printf("%d ve %d toplamı %d\n", x, y, sum)

	   	x = 7
	   	y = 11
	   	sum = x + y
	   	fmt.Printf("%d ve %d toplamı %d\n", x, y, sum) */

	// Fonksiyonların yardımıyla kodumuzu daha modüler bir hale getiririz. Daha okunaklı olur. Ayrıca oluşturduğumuz bir fonksiyonu tekrar tekrar çalıştırabiliriz. Üstteki örneğimizi fonksiyon ile ifade edelim:

	fmt.Println(sum(5, 10))

	merhaba()

	sum2(6, 11)
}

func sum(x, y int) int {
	return x + y
}

// ÖNEMLİ NOT: Eğer bir fonksiyon herhangi bir değer döndürmeyecekse ve sadece belirli bir işlemi gerçekleştirecekse, o zaman return ifadesini kullanmak zorunda değilsiniz. Go dilinde, bir fonksiyonun dönüş değeri belirtilmediğinde veya return ifadesi kullanılmadığında, fonksiyon varsayılan olarak nil veya zero value değerini döndürür.

func merhaba() {
	fmt.Println("Benim adım Ali")
}

func sum2(x, y int) {
	fmt.Println(x + y)
}

// 1. Değer Döndüren Fonksiyon

/*
package main

import "fmt"

// Bu fonksiyon, iki sayıyı toplar ve sonucu döndürür
func toplama(x, y int) int {
    sonuc := x + y
    return sonuc
}

func main() {
    // toplama fonksiyonu çağrılır ve dönen değer alınarak ekrana yazdırılır
    result := toplama(3, 5)
    fmt.Println("Toplam:", result)
}
*/

// 2. Değer Döndürmeyen Fonksiyon

/*
package main

import "fmt"

// Bu fonksiyon, iki sayıyı toplar ve sonucu ekrana yazdırır, ancak bir değer döndürmez
func sadeceYazdir(x, y int) {
    sonuc := x + y
    fmt.Println("Toplam:", sonuc)
    // return ifadesi olmadığı için varsayılan olarak nil veya zero value dönecek
}

func main() {
    // sadeceYazdir fonksiyonu çağrılır, ancak dönen değer alınmaz
    sadeceYazdir(3, 5)
}
*/
