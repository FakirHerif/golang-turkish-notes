// BU KONU GO'NUN EN ÖNEMLİ KONULARINDAN BİR TANESİDİR.

package main

import (
	"fmt"
	"strings"
)

// struct --> underlying type, employee ---> defined type, named type

type mile float64 // mile, float64 üzerine kurulmuş bir veri tipidir. yani mile float64 veri tipi değildir. buna dikkat edelim.

type kilometer float64

type mystring string

func main() {

	var m1 mile
	m1 = 3.2
	fmt.Println(m1)              // 3.2
	fmt.Printf("%T, %v", m1, m1) // main.mile, 3.2

	fmt.Println()

	f1 := float64(4.4)
	fmt.Println(3.2 + 4.4) // 7.6 ---> ikiside aynı veri tipine sahip olduğu için sorunsuz şekilde işlemimiz gerçekleşti ve sonuç döndürüldü
	/* fmt.Println(m1 + f1) ---> ancak bu işlemi gerçekleştiremeyiz. Çünkü bu ikisi birbirlerinden farklı veri tipine sahip. Ancak bunları birbirine dönüştürebiliriz */
	fmt.Println("**********")

	fmt.Println(m1 + mile(f1))    // 7.6000000000000005 ---> f1'i mile veri tipine dönüştürdük ve bu sayede sorunsuz bir şekilde işlemimizi gerçekleştirdik. veya m1'i float64'e dönüştürebiliriz:
	fmt.Println(float64(m1) + f1) // 7.6000000000000005 ---> burda da m1'i float64 veri tipine dönüştürdük ve sorunsuz bir şekilde işlemimizi gerçekleştirdik.

	fmt.Println("**********")

	k1 := kilometer(7.8)
	fmt.Printf("%T, %v", k1, k1) // main.kilometer, 7.8

	// NOT: Oluşturduğumuz mile ve kilometer farklı veri tipleri buna dikkat edelim. m1 + k1 toplamaya çalışsırsak veri tipleri farklı olduğu için hata alırız ve işlemi gerçekleştiremeyiz.

	fmt.Println()

	m2 := mile(4.6)
	fmt.Println(m1 + m2) // 7.8 ---> m1 ve m2 aynı veri tipine sahip olduğu için sorunsuz bir şekilde işlemimizi gerçekleştirebiliriz.

	fmt.Println("**********")

	fmt.Println(m1 + m2 + 2.1) // 9.9 ---> float64 olduğu için yani sayısal bir değer, toplama, çıkarma, çarpma ve bölme gibi işlemleri sorunsuz gerçekleştirebiliriz.
	/* fmt.Println(m1 + m2 + "ali") ---> ancak burada kendimiz yeni bir method tanımladığımız için string bir ifade ile toplanamaz çünkü mile veri tipi float64 veri tipi üzerine kurulmuştur yani sayısal ifadeleri içerebilir */

	fmt.Println("**********")

	mystring := "ali"
	fmt.Println(strings.ToUpper(mystring)) // ALI

	// NOT: Peki yukarıda yaptığımız işlemleri kendi veri tipimizi oluşturmadan da basit şekilde yapamaz mıydık? Evet yapabilirdik. Örneğin;

	deneme := "deneme"
	fmt.Println(strings.ToUpper(deneme)) // DENEME ---> Kendi veri tipimizi oluşturmadan basit bir şekilde stringimizi tanımladık. Peki neden bu kadar basit bir şekilde yapabiliyorken kendi veri tipimizi oluşturuyoruz neden kendi veri tipimize ihtiyacımız var? Hemen açıklayalım:

	/*
		1. Var olan basit işlemleri yapmak için zaten kendi veri tipimizi oluşturmayız. (Burada yer alanlar sadece bir örnek olduğu için oluşturuldu)

		2. Kodun daha kolay okunmasını sağlar

		3. En temel sebebi o veri tipine ait olan yeni fonksiyonellikler tanımlamaktır. Yani o veri tipine yeni methodlar bağlamaktır. Bu tam olarak ne demek? Hemen bunu da açıklayalım, mile veri tipine ait main fonksiyonu dışında yeni bir fonksiyon tanımlayalım. (Önce main fonksiyonu dışında fonksiyonlarımızı tanımlıyoruz daha sonra main fonksiyonu içinde bu fonksiyonları kullanarak işlemlerimizi gerçekleştiriyoruz)
	*/

	mileOrnek1 := mile(10)
	kilometerOrnek1 := toKilometer(mileOrnek1)
	fmt.Println(kilometerOrnek1) // 16

	fmt.Println("**********")

	kilometreOrnek2 := kilometer(10)
	mileOrnek2 := toMile(kilometreOrnek2)
	fmt.Println(mileOrnek2) // 6.2

}

func toKilometer(m mile) kilometer {
	return kilometer(m * 1.6) // mile veri tipinin kendi ait bir fonksiyonu oluşturduk (mile'ı kilometreye çeviren fonskiyon)
}

func toMile(k kilometer) mile {
	return mile(k * 0.62) // // kilometer veri tipinin kendi ait bir fonksiyonu oluşturduk (kilometreyi mile çeviren fonskiyon)
}
