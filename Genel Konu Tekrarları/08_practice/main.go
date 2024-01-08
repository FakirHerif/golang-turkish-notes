/*
1- int x, float64 y type conversion sample
2- multiple assign sample x,y = y, x
3- non English variable names
4- shadowing ?
5- 40 as a string
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {

	/* 	1-
	   	x := 75
	   	var y float64
	   	y = float64(x) // type(value)
	   	fmt.Println(y)
	*/

	/* 	2-
	   	x := 5
	   	y := 10
	   	fmt.Println("x:", x, "y:", y)

	   	x, y = y, x
	   	fmt.Println("x:", x, "y:", y)
	*/

	/* 	3-
	   	yaş := 40
	   	fmt.Println(yaş) // türkçe karakterle oluşturabiliriz ama çok doğru bir kullanım şekli değildir. İngilizce yazım tarzı kullanmak daha doğrudur.
	*/

	/* 	4-
	   	x := 5
	   	if true {
	   		x := 10            // YENİ BİR x OLUŞTURMUŞ OLUYORUZ
	   		fmt.Println(x + 1) // 11
	   	}
	   	fmt.Println(x) // 5

	   	fmt.Println("*******")

	   	y := 10
	   	if true {
	   		y = 20              // var olan y'ye yeni bir değer vermiş olduk
	   		fmt.Println(y + 10) // 30
	   	}
	   	fmt.Println(y) // 20	YENİ OLUŞTURDUĞUMUZ DEĞERİ ALIR ÇÜNKÜ y'nin değeri 20 ile değiştirdik
	*/

	x := 40
	s := string(40) // Bu şekilde kullanımda 40 sayısının ASCII karşılığı olanı ekrana yazdırır yani ===> "("

	fmt.Printf("%v, %T\n", x, x) // 40 değerini ve veri tipi olan int çıktısını veriyor.
	fmt.Printf("%v, %T\n", s, s) // "(, string" çıktısını alırız.

	// Peki string ifade ile "40" şeklinde yazdırmak istiyorsak ne yapacağız:

	y := strconv.Itoa(x)         // strconv.Itoa kullanarak int'i direkt istediğimiz şekliyle string'e çevirebiliriz
	fmt.Printf("%v, %T\n", y, y) // "40, string" çıktısını alırız.
}
