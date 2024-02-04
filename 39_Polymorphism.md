# Polymorphism

Polymorphism, nesne yönelimli programlama ile çok yaygın olarak ilişkilendirilen bir özelliktir. Bir nesnenin içeriğe bağlı olarak farklı biçimlere sahip olma yeteneğidir. Yani polimorfizm, aynı işlevin veya yöntemin, farklı nesne türleri için farklı şekillerde davranabilme yeteneğini ifade eder.

Örneğin, "alan" hesaplama işlevi alalım. Bir dikdörtgenin alanını hesaplamak için uzunluğu ve genişliği çarparız, ancak bir üçgenin alanını hesaplamak için tabanı yükseklik ile çarpıp yarıya böleriz. Bu durumda, aynı "alan" işlevi, içeriğe bağlı olarak farklı biçimlerde davranır. Bu, bir tür polimorfizmdir.

Kısaca bu örnekte alan hesaplama işlevi aynı ancak hesaplama yöntemleri farklı.

```
package main

import (
	"fmt"
)

// Şekil arayüzü
type Shape interface {
	Area() float64
}

// Dikdörtgen struct'ı, Shape arayüzünü uygular
type Rectangle struct {
	Width  float64
	Height float64
}

// Dikdörtgen'in alanını hesaplayan metot
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Üçgen struct'ı, Shape arayüzünü uygular
type Triangle struct {
	Base   float64
	Height float64
}

// Üçgen'in alanını hesaplayan metot
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func main() {
	// Dikdörtgen ve Üçgen nesneleri oluşturuluyor
	rect := Rectangle{Width: 4, Height: 5}
	tri := Triangle{Base: 3, Height: 6}

	// Her ikisi de Shape arayüzünü uygular
	shapes := []Shape{rect, tri}

	// Polimorfizm sayesinde aynı arayüzü uygulayan farklı nesneleri tek bir dilim içinde kullanabiliriz
	for _, shape := range shapes {
		fmt.Printf("Alan: %f\n", shape.Area())
	}
}
```

Genellikle nesne yönelimli dillerde
polimofilmi desteklemek için kullanılan şey kalıtımdır ve **Golang'ın kalıtımı yoktur** Bu yüzden Go dilinde bu kavramlar farklı şekilde ele alınır.

Polimorfizm, bir işlemin veya yöntemin farklı nesne türleri için farklı şekillerde çalışabilme yeteneğidir. Geleneksel nesneye yönelik dillerde, bu genellikle kalıtım (miras) kullanılarak sağlanır. Örneğin, bir üst sınıf olan "Konuşmacı" sınıfının alt sınıfları olan "Kedi" ve "Köpek" sınıfları, üst sınıfın metodlarını miras alabilir ve gerektiğinde bu metodları değiştirebilir (geçersiz kılabilir). Bu durumda, "Konuşmacı" sınıfının "Konuşma" metodunu, "Kedi" sınıfı "Miyav" ve "Köpek" sınıfı "Hav" olarak geçersiz kılabilir.

Ancak, Go dilinde kalıtım (miras) olmadığı için, polimorfizmi başka bir şekilde sağlamamız gerekiyor. Go, arayüzler (interfaces) kullanarak polimorfizmi destekler. Yani, bir arayüzü uygulayan herhangi bir tip (struct) otomatik olarak o arayüzü uygular. Bu da bize farklı tiplerin aynı arayüzü kullanarak polimorfik davranış sergilemesini sağlar.

Örneğin, yine "Konuşmacı" arayüzü oluşturabilir ve "Kedi" ve "Köpek" tiplerini bu arayüzü uygulayacak şekilde tasarlayabiliriz. Ardından, bu tipleri kullanarak polimorfik davranış elde edebiliriz.

```
package main

import "fmt"

// Konuşmacı arayüzü
type Konuşmacı interface {
    Konuş() string
}

// Kedi tipi, Konuşmacı arayüzünü uygular
type Kedi struct{}

// Kedi'nin Konuş metodunu geçersiz kılma
func (k Kedi) Konuş() string {
    return "Miyav"
}

// Köpek tipi, Konuşmacı arayüzünü uygular
type Köpek struct{}

// Köpek'in Konuş metodunu geçersiz kılma
func (d Köpek) Konuş() string {
    return "Hav"
}

func main() {
    kedi := Kedi{}
    köpek := Köpek{}

    konuşmacılar := []Konuşmacı{kedi, köpek}

    for _, konuşmacı := range konuşmacılar {
        fmt.Println(konuşmacı.Konuş())
    }
}
```

Bu örnekte, "Kedi" ve "Köpek" tipleri, "Konuşmacı" arayüzünü uygular. Ve ardından, bu tipleri içeren bir dilim oluşturarak polimorfik davranışı elde edebiliriz. Bu, Go dilinde kalıtım olmadan polimorfizmin nasıl uygulanabileceğini gösterir.