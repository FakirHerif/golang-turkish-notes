// go'da diğer dillerde olduğu gibi while gibi döngü yoktur. sadece for döngüsü vardır

// for <init statement>; <condition>; <post statement> { code }

package main

import "fmt"

func main() {

	for i := 0; i <= 10; i += 5 {
		fmt.Println(i)
	}

	fmt.Println("*******************")

	// i'yi for döngüsü dışında da tanımlayabiliriz:
	i := 0
	for ; i <= 10; i += 5 {
		fmt.Println(i)
	}

	fmt.Println("*******************")

	/* for ifadesi kendi başına sonsuzdur. Infinite Loop:
	for {
		fmt.Println("Bu yazı durdurulana kadar sonsuz olarak tekrarlanır")
	}

	- koşul ifademiz true olduğu zamanda döngü sonsuza kadar devam eder:
	for i := 0; true; i += 5 {
		fmt.Println(i)
	}

	- false durumda ise döngüye girmez.
	*/

	// post statement'ı { code } içinde de yazabiliriz:

	a := 10
	for a >= 0 {
		fmt.Println(a)
		a--
	}

	fmt.Println("*******************")

	for b := 0; b <= 10; b++ {
		if b%3 == 0 {
			continue // continue --> döngünün başına git (yani bu örnekte b'nin 3 ile bölümünden kalan 0 olduğu durumlarda döngünün başına git ve kalanın 0 olmadığı durumları yazdır demiş oluyoruz)
		}
		fmt.Println(b)
	}

	fmt.Println("*******************")

	for c := 0; c <= 10; c++ {
		if c == 3 {
			break // break --> döngüden çık (yani bu örnekte c 3'e eşit olduğu anda döngüden çıkıyor ve c 3'e eşit olana kadar ki tüm adımları ekrana yazdırmış oluyoruz)
		}
		fmt.Println(c)
	}
}
