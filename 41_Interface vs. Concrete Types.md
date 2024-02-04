# Interface vs. Concrete Types
(Go dilinde interface türleri ve somut (concrete) türler arasındaki temel farklar)

**Somut (Concrete) Türler:**

Somut türler, verinin ve metodların tam temsilini sağlayan düzenli türlerdir. Bir somut tür, verinin tam temsilini ve bu türün metodlarının tam uygulamalarını belirtir. Bu türün metodlarını kullanan herhangi bir metod tamamen belirlenmiştir. Somut türlerin önemli bir özelliği, verinin tam temsilini içermesidir.

**Arayüz (Interface) Türleri:**

Arayüz türleri ise sadece metod imzalarını belirtir. Veri belirtilmez, sadece metodlar belirtilir. Hatta metodların uygulamaları soyutlanmıştır, yani uygulamalar belirtilmez, yalnızca metodların imzaları bulunur. Arayüz türleri ile somut türler arasındaki büyük fark, verinin tam temsilinin olmamasıdır.

Ancak unutulmamalıdır ki, bir arayüz bir süre sonra bir somut türe eşlenir. 

**Arayüz Değerleri:**

Bir arayüz değeri oluşturduğunuzda, bu değeri diğer değerlerle (int, float vb.) aynı şekilde ele alabilirsiniz. Yani, bir değişken oluşturabilir ve bu değişkeni bir belirli bir arayüz türüne sahip bir değerle eşleyebilirsiniz. Örneğin, Shape2D arayüzüne sahip bir değişken oluşturabilirsiniz.

Arayüz değerinin iki bileşeni vardır: Dinamik tür (dynamic type) ve dinamik değer (dynamic value). Dinamik tür, atanan somut türdür ve dinamik değer, o türün değeridir. Yani, bir arayüz değeri aslında dinamik tür ve dinamik değerin bir çiftidir.

Örneğin, Shape2D arayüzüne sahip bir değişkeni, bu arayüzü karşılayan bir somut türe eşlerseniz, dinamik tür olarak bu somut tür atanır ve dinamik değer de bu somut türün değeri olur.

**Arayüz ve Nil Değer:**

Bir arayüzde dinamik tür olabilir ancak dinamik değer olmayabilir. Bu durumda, arayüzde bir dinamik tür olmasına rağmen değer atanmamışsa, bu durumu "nil değeri" olarak adlandırırız. Bu durumda arayüzün bir metodunu çağırmak mümkündür. Ancak, bu metodun içinde nil değeri kontrol etmek genellikle iyi bir uygulamadır, çünkü nil bir değere sahip olabilir ve bu durumda uygulamanın hata vermemesi önemlidir.

**Nil Arayüz Değeri:**

Eğer bir arayüzün hem dinamik türü hem de dinamik değeri nil ise, bu durum "nil arayüz değeri" olarak adlandırılır. Bu durumda, arayüz üzerinde herhangi bir metodun çağrılması hata verecektir, çünkü hiçbir somut tür atanmamıştır ve metodların uygulamaları bulunmamaktadır.

**ÖRNEK**

```
package main

import (
	"fmt"
)

// Speaker arayüzünü tanımla
type Speaker interface {
	Speak()
}

// Dog türünü tanımla
type Dog struct {
	Name string
}

// Dog türü, Speaker arayüzünü uygular
func (d Dog) Speak() {
	fmt.Println("Woof! My name is", d.Name)
}

func main() {
	// s1 adında bir Speaker arayüzü değişkeni tanımla
	var s1 Speaker

	// d1 adında bir Dog türü değişkeni tanımla ve ismini belirle
	d1 := Dog{Name: "Buddy"}

	// s1'e d1'i atayarak arayüzü d1 ile eşleştir
	s1 = d1

	// s1'in dinamik türü ve değerini yazdır
	fmt.Printf("Dynamic Type: %T, Dynamic Value: %v\n", s1, s1)

	// s1 üzerinden Speak metodunu çağır
	s1.Speak()

	// Nil değeri ile bir Speaker arayüzü tanımla
	var s2 Speaker

	// Nil arayüz değerini yazdır
	fmt.Printf("Dynamic Type: %T, Dynamic Value: %v\n", s2, s2)

	// Nil arayüz değeri ile bir metod çağrılmaz, hata verir
	// s2.Speak() // Commented out to prevent runtime error

	// Nil dinamik türü ile bir Speaker arayüzü tanımla
	var s3 Speaker = (*Dog)(nil)

	// Nil dinamik türü ile bir arayüz değeri yazdır
	fmt.Printf("Dynamic Type: %T, Dynamic Value: %v\n", s3, s3)

	// Nil dinamik türü ile bir metod çağrılmaz, hata verir
	// s3.Speak() // Commented out to prevent runtime error
}
```

Bu örnek, Speaker arayüzünü ve Dog somut türünü içermektedir. Dog türü, Speak metodunu uygulayan bir Speaker arayüzüdür. main fonksiyonunda bu türleri kullanarak arayüz değerlerini ve nil değerlerini açıklamaktadır.

Örneğin, s1 değişkeni bir arayüz değeri oluşturur ve dinamik türü Dog türüne eşlenir. s1.Speak() çağrısı yapılırken, Dog türündeki Speak metodunu çağıracaktır. s2 değişkeni ise nil bir arayüz değeridir ve metod çağrısı yapmak hata verecektir. s3 değişkeni ise nil bir dinamik türle oluşturulmuş bir arayüz değeridir ve metod çağrısı yine hata verecektir.