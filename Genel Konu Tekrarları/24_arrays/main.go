package main

import "fmt"

func main() {

	city1 := "istanbul"
	city2 := "izmir"
	city3 := "ankara"
	city4 := "antalya"

	fmt.Println(city1, city2, city3, city4) // bu şekilde tek tek yazmak ne kadar uğraştırıcı, düşünsenize 100 tane veye daha fazlasını tek tek yazmaya çalıştığımızı...

	cities := [4]string{"istanbul", "izmir", "ankara", "antalya"} // array oluştururken veri tipini belirtmemiz gerekiyor. 4 yazdığımız ise kaç tane verimizin olduğunu belirtiyor. Yani 4 dilim ayırdık ve elemanlarımızı yazdık
	fmt.Println(cities)                                           // tüm elemanları yazdırır
	fmt.Println(cities[0])                                        // birinci sıradaki elemanı yazdırır. "istanbul"

	// elemanın aldığı sayıya göre arrayimizi oluşturmak istersek [ ] şeklinde boş belirtebiliriz veya [...] yazarak belirtebiliriz. Örneğin:

	meyveler := [...]string{"elma", "armut", "kiraz"}
	fmt.Println(meyveler)
	fmt.Println(len(meyveler)) // array uzunluğumuzu yani kaç elemandan oluştuğunu bu şekilde sorgulayabiliriz

	// ÖNEMLİ NOT: Genellikle verileri örneğin database üzerinden çekerken arrayimize gelecek olan elemanları bilmiyor olabiliriz. Bu durumda başlangıç vermediğimiz array oluşturmamız gerekir:

	var myArr [5]int          // 5 elemanlı int veri tipine sahip bir array oluşturmuş oluyoruz
	fmt.Println(myArr)        // [0 0 0 0 0] sonucu döner. Yani içerdiği eleman sayısı kadar zero value alır
	myArr[0] = 1              // 0. indekse yani birinci sıraya 1 elemanımızı yerleştirdik
	fmt.Println(myArr)        // [1 0 0 0 0] sonucu döner.
	myArr[len(myArr)-1] = 100 // en son sıraya 100 elemanımızı ekledik
	fmt.Println(myArr)        // [1 0 0 0 100] sonucu döner.

	/* 	var myArr1 [5]int
	   	var myArr2 [4]int
	   	Bu iki array birbirinden farklı veri tipine sahip arraylerdir. Ama nasıl ikiside int dediğinizin farkındayım ancak bu arraylerden myArr1 olan 5 elemana sahip int veri tipine sahipken myArr2 ise 4 elemana sahip int veri tipine sahiptir.

		Hatta bunların birbirinden farklı olduğunu kanıtlamak için şöyle basit bir sorgulama yapabiliriz:

		if myArr1 == myArr2 {
			fmt.Println("Bu ikisi aynı")
		}

		direkt hata aldığımızı göreceksiniz. Eğer ikiside aynı eleman sayısına sahip int olsaydı bu durumda sorun olmayacaktı ve ikisi de aynı veri tipine sahip arraylerdir diyebilecektik.
	*/

	mevsim := [4]string{"ilkbahar", "yaz", "sonbahar", "kış"}
	// for döngüsü ile elemanlarımızı indeksleriyle birlikte tek tek yazdırabiliriz:
	for i := 0; i < len(mevsim); i++ {
		fmt.Println(i, mevsim[i])
	}

	// FOR --- RANGE KULLANIMI:

	cicekler := [4]string{"papatya", "menekşe", "gül", "lale"}

	for index, cicek := range cicekler {
		fmt.Println(index, cicek)
	}

	for _, cicek := range cicekler {
		fmt.Println(cicek)
	} // indeks olmadan yazdırmak istersek bu şekilde kullanabiliriz.

	myArr3 := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	myArr3 = mySquare(myArr3)
	fmt.Println(myArr3)
	/*
		for i := 0; i < len(myArr3); i++ {
			myArr3[i] = myArr3[i] * myArr3[i]
			fmt.Println(myArr3[i])
		} // şeklinde karelerini alabiliriz. peki karelerini fonksiyon kullanarak tek arrayde yazdıralım:
	*/
}

func mySquare(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
	}
	return arr
}
