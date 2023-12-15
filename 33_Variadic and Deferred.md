# Variadic and Deferred

**Değişken Sayıda Argüman Alan Fonksiyonlar**

Go dilinde değişken sayıda argüman alan fonksiyonlar, **... (üç nokta)** operatörü kullanılarak tanımlanır. Bu fonksiyonlar, farklı tiplerde birden fazla argüman alabilirler. Örneğin:

```
// Topla; değişken sayıda tamsayı argüman alarak toplamını döndüren bir fonksiyondur.
func Topla(sayilar ...int) int {
	toplam := 0
	for _, sayi := range sayilar {
		toplam += sayi
	}
	return toplam
}

func main() {
	// Değişken sayıda argümanları farklı yöntemlerle iletme

	// 1. Yöntem: Birden fazla argümanı doğrudan ileterek
	sonuc1 := Topla(1, 2, 3, 4)
	fmt.Println("Sonuc 1:", sonuc1) // Çıktı: Sonuc 1: 10

	// 2. Yöntem: Bir dilim kullanarak
	dizi := []int{5, 6, 7, 8}
	sonuc2 := Topla(dizi...)
	fmt.Println("Sonuc 2:", sonuc2) // Çıktı: Sonuc 2: 26
}
```

Yukarıdaki örnekte, Topla adında bir fonksiyon tanımlıyoruz. Bu fonksiyon, ...int kullanarak değişken sayıda tamsayı argüman alabiliyor. Topla fonksiyonu, aldığı tüm tamsayıları toplar ve toplamı döndürür. main fonksiyonunda bu fonksiyonu kullanırken iki farklı yöntemi görebilirsiniz.

Doğrudan birden fazla argümanı fonksiyona iletme (Topla(1, 2, 3, 4))
Bir dilim kullanarak argümanları iletme (Topla(dizi...))

**Ertelenmiş Fonksiyon Çağrıları**

Go'da defer ifadesi, bir fonksiyon çağrısının çevreleyen fonksiyonun sonunda çalışmasını sağlar. Bu, kaynakların temizlenmesi veya belirli bir zaman diliminde bir işlemin gerçekleştirilmesi gibi durumlarda kullanışlıdır.

```
func main() {
	i := 1

	// Ertelenmiş fonksiyon çağrıları
	defer fmt.Println("Güle güle") // Fonksiyon çağrısı ertelendi

	fmt.Println("Merhaba") // Çıktı: Merhaba

	// Değişken değerlerini değerlendirme
	defer fmt.Println(i + 1) // Çıktı: 2

    i++

	// Zamanlama
	defer fmt.Println("Ertelenmiş fonksiyonlar daha sonra değerlendirilir")

	fmt.Println("Ertelenmiş fonksiyonlar daha sonra çalıştırılır")
}
```

Yukarıdaki örnekte, defer ifadesiyle işaretlenmiş fonksiyonlar, çevreleyen fonksiyon tamamlanana kadar çağrılmaz ve bu fonksiyonlar, çevreleyen fonksiyon tamamlandıktan sonra çalıştırılır. Bu özellik, kaynakların serbest bırakılması veya belirli bir zaman diliminde bir işlemin gerçekleştirilmesi için kullanılabilir.

**DEFER KULLANIMIYLA İLGİLİ BASİT VE ÖNEMLİ ÖRNEKLER:**

```
func main() {
	i := 1

	defer fmt.Println(i + 1) // Çıktı: 2

	i++
}
```

Bu kod örneğinde, i isimli bir değişken oluşturulur ve başlangıç değeri 1 olarak atanır. Ardından defer ifadesiyle bir fonksiyon çağrısı ertelenir. defer, fonksiyon çağrısını bulunduğu fonksiyonun sonunda çalıştırmak üzere bekletir. fmt.Println(i + 1) ifadesi defer ile ertelendiği için, fonksiyon main() tamamlandıktan sonra çalışır. Bu durumda i değişkeni bir artırıldığından (i++), fonksiyon çağrısı sırasında i değeri 2 olur ve 2 + 1 işlemi sonucunda ekrana 2 yazdırılır.


```
func main() {
	i := 1
	i++
	defer fmt.Println(i + 1) // Çıktı: 3
}
```

Bu kod örneğinde, i isimli bir değişken oluşturulur ve başlangıç değeri 1 olarak atanır. Daha sonra i bir artırılır (i++). Ardından defer ifadesiyle bir fonksiyon çağrısı ertelenir. defer ifadesi, fonksiyon çağrısını bulunduğu fonksiyonun sonunda çalıştırmak üzere bekletir. fmt.Println(i + 1) ifadesi defer ile ertelendiği için, fonksiyon main() tamamlandıktan sonra çalışır. i değişkeni zaten bir artırıldığından (i önceden 1 idi ve i++ ile 2 oldu), fonksiyon çağrısı sırasında i değeri 2 olur ve 2 + 1 işlemi sonucunda ekrana 3 yazdırılır.

Özetle, defer ifadesi, fonksiyon çağrısını o ifadeyi içeren fonksiyonun sonunda çalıştırmak üzere erteleyerek, o anki değerleri alır ve sonuçları çalıştırma anında değil, fonksiyonun tamamlanmasından sonra kullanır. Bu örneklerde defer ifadesi, i değişkeninin değerini çalıştırma anında değil, çağrıldığı zaman belirlenen değer üzerinden işlem yapar.