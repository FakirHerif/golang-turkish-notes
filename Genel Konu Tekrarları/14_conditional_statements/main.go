// Koşullu ifadeler
// if <boolean expression> { code }

// boolean expression kısmına yazdığımız ifadelerin sonucu bool veya true/false olmalıdır. js'de ki geniş bir kapsam yoktur. Nedeni ise go'da veri tipleri çok önemlidir.

// if <boolean expression> { code } else if <boolean expression> { code } ... else if <boolean expression> { code } else { code }

package main

import "fmt"

func main() {

	x := 27

	if x%2 == 0 {
		fmt.Println(x, "çift sayıdır")
	} else {
		fmt.Println(x, "tek sayıdır")
	}

	if !true {
		fmt.Println("Mesaj gösterilmeyecek :(") // değer false olduğu için herhangi bir println mesajını göstermez.
	}

	if true {
		fmt.Println("Mesaj gösterilecek :)")
	}

	a := 6

	if a < 0 {
		fmt.Println(a, "negatif sayıdır")
	} else if a%2 == 0 {
		fmt.Println(a, "çift sayıdır")
	} else {
		fmt.Println(a, "tek sayıdır")
	}

	// genellikle tanımladığımız değerleri dışarıda değil, if yapısında yani if bloğunun içinde tanımlamak isteriz bu durumda:

	if b := -5; x < 0 { // NOT: İki statement aynı satırda olduğu zaman ";" kullanmayı unutmuyoruz.
		fmt.Println(b, "negatif sayıdır")
	} else if b%2 == 0 {
		fmt.Println(b, "çift sayıdır")
	} else {
		fmt.Println(b, "tek sayıdır")
	}

}
