package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	merhaba("ali", 30) // Argüman

	fmt.Println("**********")

	fmt.Println(result(40))

	fmt.Println("**********")

	// Farklı paketlerdeki fonksiyonlara metod denir. Metodu farklı veri tiplerine ait olan fonksiyonlar şeklinde de tanımlayabiliriz. Örneğin:

	var moment time.Time = time.Now() // Now ---> method
	fmt.Println(moment)

	fmt.Println("**********")

	fmt.Print("Lütfen sınav sonucunuzu giriniz: ") // girilen sonucun yeni satırda yazmaması için Println kullanmıyoruz, aynı satırda yazması için Print kullanıyoruz.
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n') // _ --> blank identifier (genellikle kullanılmayan değerleri veya değişkenleri ifade etmek için kullanılır)
	fmt.Println(value)

	fmt.Println("**********")

	bölüm, kalan := bölme(104, 5)
	fmt.Println(bölüm, kalan)
}

// bir fonksiyonda kullandığımız değişkenlerin scope'u/kapsamı o fonksiyon içerisindedir. bu örnekte ki name ve age sadece bu fonksiyon içinde geçerlidir.
func merhaba(name string, age int) { // Parametre
	fmt.Printf("Adım %s, yaşım %d\n", name, age)
}

func result(grade int) string {
	if grade >= 50 {
		return "geçtiniz"
	}
	return "kaldınız"
}

// return'dan sonra herhangi bir statement yazılamaz. return en sona gelir.

// go'da birden fazla değeri return edebiliriz. Örneğin:
func bölme(bölünen int, bölen int) (bölüm, kalan int) {
	bölüm = bölünen / bölen
	kalan = bölünen % bölen

	return bölüm, kalan
}
