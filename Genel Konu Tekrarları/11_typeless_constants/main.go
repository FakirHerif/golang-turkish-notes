// typeless constant = veri tipine sahip olmayan sabit değişken

package main

import "fmt"

func main() {

	const x = 5 // varsayılan bir veri tipi atanmamıştır. bu örnek typeless constant'a örnektir

	fmt.Printf("%T, %v\n", x, x) // varsayılan veri tipi atıyor yani int

	var y int16 = 15
	fmt.Printf("%T, %v\n", y, y)

	fmt.Printf("%T, %v\n", x+y, x+y) // type conversion gerçekleşiyor ve şu anlama geliyor ==> int16(x) + y ==> ve int16, 20 sonucu döndürür.

	const a = 3
	const b = 5.6

	fmt.Printf("%T, %v\n", a, a) // int, 3
	fmt.Printf("%T, %v\n", b, b) // float64, 5.6

	fmt.Printf("%T, %v\n", a+b, a+b) // float64, 8.6

	// typeless değişkenlerde go veri tiplerini varsayılan olarak otomatik atıyor. a+b gibi durumda ise uygun olan veri tipini atıyor.
}
