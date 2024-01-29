// pointer başka bir değişkenin adresini tutan değişkendir.

package main

import "fmt"

func main() {

	name := "Ali"
	fmt.Println(name)  // Ali
	fmt.Println(&name) // 0xc00002a070 ---> "Ali" değerine sahip name değişkenin bulunduğu yerin hafızadaki değeridir. & ---> address operator /// Not: bu değerler sizde farklı olabilir.

	fmt.Println("**************")

	x := 22
	fmt.Println(x)  // 22
	fmt.Println(&x) // 0xc00000a0c8

	fmt.Println("**************")

	fmt.Printf("%T, %v\n", x, x)   // int, 22
	fmt.Printf("%T, %v\n", &x, &x) // *int, 0xc000116088 ---> burda dikkat etmemiz gereken en önemli şey &x'in veri tipi int değildir, &x'in veri tipi ===> *int'tir. (* ===> asterisk(asteriks) veya yıldız olarak adlandırılır) Yani biz bu değeri int olan bir değişkene atayamayız:

	/* 	var y int = &x	// Burada int değerine sahip değişkene *int değerini atamaya çalışırsak hata alırız. Bu önemlidir.
	   	fmt.Println(y)

	Yapabileceğimiz şu şekildedir: */

	var y *int = &x
	// veya y := &x
	fmt.Println(y)               // 0xc000192088
	fmt.Printf("%T, %v\n", y, y) // *int, 0xc000192088

	fmt.Println("**************")

	z := &name
	fmt.Printf("%T, %v\n", z, z) // *string, 0xc00002a070

	// ÖNEMLİ NOT: string, int, bool tipleri pointerda dikkat edilmesi gereken önemli durumlardır. Yani *int olan değere *string değer atayamayız. Veya *string olan değere *int, int, string, bool vs atayamayız. *string değere sadece *string değer atayabiliriz. Ama unutmamak gerekir ki tüm veri tiplerine pointer atanabilir. Yani demek istediğim yukarıdaki örnekteki gibi string veri tipine sahip name değişkenine &name pointerı atayabildiğimiz gibi.

	fmt.Println("**************")

	fmt.Println(x)           // 22
	fmt.Println(&x)          // 0xc00000a0c8 ---> 22'nin adresi
	fmt.Println(*(&x))       // 22 ---> yıldız operatörü pointerın sahip olduğu gerçek değeri dönüyor yani pointer karşılığı olan adresteki değeri normal değerine döndürüyor. ---> dereferencing
	fmt.Println(&(*(&x)))    // 0xc00000a0c8 ---> * ---> ilgili adresteki değeri gösterir
	fmt.Println(*(&(*(&x)))) // 22

	// KISACA: & ---> adresi gösteriyor, * ---> ilgili adresteki değeri gösteriyor

	fmt.Println("**************")

	/* 	x1 := 10
	   	x2 := x1
	   	fmt.Println(x1, x2) // 10 10
	   	x1 = 5
	   	fmt.Println(x1, x2) // 5 10
	*/

	/* 	x1 := 10
	   	x2 := &x1
	   	fmt.Println(x1, x2)  // 10 0xc0001160e0
	   	fmt.Println(x1, *x2) // 10 10

	   	*x2 = 3              // x1'in adresinin bulunduğu değeri 3 yapmış olduk yani x1 değerinin kendisinide 3 yapmış olduk
	   	fmt.Println(x1, *x2) // 3 3 ---> *x2 ile x1'in bulunduğu yerdeki değeri alıyoruz ve x2'in bulunduğu yerdeki değeri yani x2'yi dereferencing yaparak 3'e eşitliyoruz. Yani varolan değer 10'dan 3'e dönüşüyor

	   	x3 := &x1
	   	*x3 = 5
	   	fmt.Println(x1, *x2, *x3) // 5 5 5
	*/

	fmt.Println("**************")

	/* 	x1 := [4]int{1, 10, 100, 1000} // array pass by value
	   	x2 := x1
	   	fmt.Println(x1, x2) // [1 10 100 1000] [1 10 100 1000]
	   	x2[0] = 3
	   	fmt.Println(x2) // [3 10 100 1000]
	   	fmt.Println(x1) // [1 10 100 1000]
	*/

	x1 := []int{1, 10, 100, 1000} // slice pass by reference
	x2 := x1
	fmt.Println(x1, x2) // [1 10 100 1000] [1 10 100 1000]
	x2[0] = 3
	fmt.Println(x2) // [3 10 100 1000]
	fmt.Println(x1) // [3 10 100 1000]

}
