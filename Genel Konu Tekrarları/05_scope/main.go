package main

import "fmt"

var packVar = "Package Scope"
var funcVar = "Func(Package) Scope"

func main() {

	if true {
		var blockVar = "Block Scope"
		fmt.Println(blockVar)
	}

	var funcVar = "Func Scope"

	fmt.Println(funcVar)

	fmt.Println(packVar)

	myFunc()

	myFuncTwo()

}

func myFunc() {
	fmt.Println(packVar)
}

func myFuncTwo() {
	fmt.Println(funcVar)
}

// Block kısaca kapsadığı süslü parantezlerin içidir.

// Paket değişkenine tüm pakette, Block değişkenine ise sadece var olduğu block içinde ulaşabiliyoruz

// Örneğin if true içindeki blockVar değişkenini scope dışından ulaşamayız hata alırız.

// NOT: Kısa gösterimi := paket değişkenlerinde kullanamayız, Block scope'da kısa gösterim := kullanabiliriz.

// ÇOK ÖNEMLİ NOT: TÜM DEĞİŞKENLERİ paket kapsamında en üste yazarsak paket düzeyindeki değişkenler programın tüm çalışma süresi boyunca hafızada yer kaplar. Bu yüzden hepsini en üstte tanımlamıyoruz. Block içinde tanımlıyoruz ki sadece o block'un çalışma süresi boyunca hafızada yer kaplasın, block'un çalışması bitince boşuna hafıza yer kaplamamış olur.

/* --- ÖNEMLİ ÖRNEK ---
var name = "Ali"
name := "Veli" --->>> KULLANILAMAZ AMA:
name, surname := "Veli", "Can" --->>> ŞEKLİNDE KULLANILABİLİR.
*/
