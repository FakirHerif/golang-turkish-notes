// bizim için özel bir anlam ifade eden bir değeri bulmak için map kullanabiliriz

package main

import (
	"fmt"
)

func main() {

	myMap := map[string]int{ //string key değerleri oluyor. Yani ahmet, ayşe, selim, mustafa. Yanlarında yer alan int değerler.
		"Ahmet":   40,
		"Ayşe":    17,
		"Selim":   14,
		"Mustafa": 70,
	}

	fmt.Println(myMap)                          // map[Ahmet:40 Ayşe:17 Mustafa:70 Selim:14]
	fmt.Println(myMap["Ahmet"])                 // 40 ---> key değerini yazarak değere ulaştık
	fmt.Println(myMap["Ahmet"], myMap["Selim"]) // 40 14
	fmt.Println(myMap["abcd"])                  // 0 ---> olmayan key değeri yazarsak zero value döner.	(Değer int olduğu için 0 dönüyor)

	fmt.Println("*********")

	isMarried := map[string]bool{
		"Ahmet": true,
		"Ayşe":  false,
		"Selim": false,
	}

	fmt.Println(isMarried)

	fmt.Println("*********")

	// Slice'larda olduğu gibi map'de de make methodunu kullanabiliriz

	myMap2 := make(map[string]int)
	fmt.Println(myMap2) // map[]

	studentGrades := map[string]int{
		"Arin":  80,
		"Ahmet": 29,
		"Selim": 72,
		"Ayşe":  77,
		"Çınar": 0,
	}
	fmt.Println(studentGrades)          // map[Ahmet:29 Arin:80 Ayşe:77 Selim:72 Çınar:0]
	fmt.Println(studentGrades["Arin"])  // 80
	fmt.Println(studentGrades["Çınar"]) // 0

	// Peki çınar 0 aldı ve 0 olarak yazdırıldı, listede olmayan Elis adlı birini yazdırırsak o da 0/zero value dönecek. Bu durumda key'in listede olup olmadığını nasıl anlayabiliriz:
	value, ok := studentGrades["Elis"]
	fmt.Println(value, ok) // 0 false ---> Yer almadığını bu şekilde anlayabiliriz.

	// Yeni key ve value eklemek:
	studentGrades["Mahmut"] = 55
	fmt.Println(studentGrades) // map[Ahmet:29 Arin:80 Ayşe:77 Mahmut:55 Selim:72 Çınar:0] ---> Mehmut eklendi. Ayrıca keyler otomatik alfabetik sıraya koyuluyor

	// Mapteki elemanları silmek:
	delete(studentGrades, "Selim")
	fmt.Println(studentGrades) // map[Ahmet:29 Arin:80 Ayşe:77 Mahmut:55 Çınar:0] ---> Selim silindi.

	fmt.Println(len(studentGrades)) // 5 ---> len yöntemiyle eleman sayımızıda yazdırabiliriz

	fmt.Println("*********")

	// Map'ler de slice'lar daki gibi array'lerden farklı olarak DEĞERİ DEĞİL, REFERANSI PAYLAŞIR.

	studentGrades2 := map[string]int{
		"Arin":  80,
		"Ahmet": 29,
		"Selim": 72,
		"Ayşe":  77,
		"Çınar": 0,
	}

	sg := studentGrades2
	fmt.Println(studentGrades2) // map[Ahmet:29 Arin:80 Ayşe:77 Selim:72 Çınar:0]
	fmt.Println(sg)             // map[Ahmet:29 Arin:80 Ayşe:77 Selim:72 Çınar:0]

	delete(sg, "Arin")
	fmt.Println(sg)             // map[Ahmet:29 Ayşe:77 Selim:72 Çınar:0]
	fmt.Println(studentGrades2) // map[Ahmet:29 Ayşe:77 Selim:72 Çınar:0] ---> mapte bir değişiklik yaptığımız zaman tüm kopyalarına etki ediyor. sg'de "Arin" sildik ve studentGrades2'de de silindi. (tutuculardan birinde değişiklik yapıldığı anda tüm hepsi etkilenir) ---> PASS BY REFERENCE

	fmt.Println("*********")

	// Map'lerde for döngüsü

	for k, v := range studentGrades2 { // for'da key için "k", value için "v" kısaltması kullanılır
		fmt.Println(k, v) // elemanlarımızın çıktısını alırız
	}
}
