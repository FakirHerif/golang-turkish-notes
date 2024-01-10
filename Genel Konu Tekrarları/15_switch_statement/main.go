// switch yapısı if/else yapısı çok karışık görünmesin diye kullanılır. if/else yapısıyla aynı mantığa sahiptir. Yani çok fazla ifade bulunuyorsa switch yapısı if/else yapısına göre daha okunaklı oluyor.

// case'lerin altında otomatik olarak bir break yer alır.

package main

import "fmt"

func main() {

	grade := 3

	switch grade {
	case 5: // if grade == 5 { fmt.Println("Pekiyi") } anlamına geliyor.
		fmt.Println("Pekiyi")
	case 4:
		fmt.Println("İyi")
	case 3:
		fmt.Println("Orta")
	case 2:
		fmt.Println("Geçer")
	case 1:
		fmt.Println("Başarısız")
	default:
		fmt.Println("Geçersiz Not") // beklenenden daha farklı bir ifadeyi default ile yakalıyoruz. Örneğin grade değerinin -3 olması durumu gibi
	}

	// case yanına yazdığımız değerin veri tipi ile oluşturduğumuz değişkenin veri tipinin aynı olması gerekiyor. Yani grade := 3 bir int veri tipine sahip olduğu için case'in yanına örneğin bir string ifadesi yazamayız, int veri tipine sahip bir değer olmalıdır.

	switch {
	case false:
		fmt.Println("Bu yazdığımız konsolda görünmez.")
	case true:
		fmt.Println("Bu yazdığımız konsolda görünür.")
	}

}
