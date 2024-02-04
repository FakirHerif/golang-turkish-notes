# Using Interfaces

Arayüzler, Go programlama dilinde türler arasında benzerliği ifade etmek için kullanılır. İki tip arayüzü sağlıyorsa, uygulama açısından önemli bir benzerlikleri olmalıdır. Genellikle bir işlev, bir parametre olarak birden fazla tür almak istediğimizde kullanılır. Örneğin, bir işlev genellikle bir tamsayıyı alır, ancak bir tamsayı veya ondalık sayıyı da kabul etmek istiyorsak, buna bir arayüz kullanabiliriz.

u konsepti daha somut bir örnekle açıklayalım. Diyelim ki bir bahçemiz var ve bu bahçeye bir havuz yerleştirmek istiyoruz. Ancak havuzun belirli kriterleri karşılaması gerekiyor: bahçeye sığmalı ve çevresi bir çit ile çevrilmelidir. Havuz şekillerini kontrol eden bir işlev yazmak istiyoruz ve bu işlev farklı türlerle çalışabilmelidir. İşte burada arayüzler devreye girer.

Örneğin, Shape2D adında bir arayüz tanımlayabiliriz. Bu arayüz, herhangi bir 2 boyutlu şeklin uygulamasını gerektirir ve bu uygulama, şeklin alanını ve çevresini hesaplayabilmesini sağlar. Daha sonra, üçgen ve dikdörtgen gibi farklı şekil tipleri oluşturabiliriz. Bu tipler, Shape2D arayüzünü uygular ve kendi alan ve çevre hesaplamalarını içerir.

Son olarak, FitInYard adlı bir işlev oluşturabiliriz. Bu işlev, Shape2D arayüzünü parametre olarak alır. Yani, bu işlev herhangi bir şekil türünü kabul edebilir. İşte bu noktada, arayüzler sayesinde farklı şekil tiplerini aynı işlev içinde kullanabiliyoruz.

**ÖRNEK**

```
package main

import (
	"fmt"
	"math"
)

// Shape2D interface'ini tanımla
type Shape2D interface {
	Area() float64
	Perimeter() float64
}

// Triangle türünü tanımla
type Triangle struct {
	Base   float64
	Height float64
	Side1  float64
	Side2  float64
	Side3  float64
}

// Triangle, Shape2D interface'ini uygular
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float64 {
	return t.Side1 + t.Side2 + t.Side3
}

// Rectangle türünü tanımla
type Rectangle struct {
	Width  float64
	Height float64
}

// Rectangle, Shape2D interface'ini uygular
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// FitInYard fonksiyonunu tanımla
func FitInYard(s Shape2D) bool {
	// 100 birim karelik bir arka bahçe ve 100 birimlik çit uzunluğu gibi kabul edelim
	maxYardArea := 100.0
	maxFenceLength := 100.0

	return s.Area() < maxYardArea && s.Perimeter() < maxFenceLength
}

// PrintMe fonksiyonunu tanımla
func PrintMe(val interface{}) {
	fmt.Println(val)
}

func main() {
	// Triangle ve Rectangle türlerini kullanarak FitInYard fonksiyonunu test et
	triangle := Triangle{Base: 5, Height: 8, Side1: 3, Side2: 4, Side3: 5}
	rectangle := Rectangle{Width: 10, Height: 5}

	// FitInYard fonksiyonunu çağırarak bahçeye sığma kontrolü yap
	fmt.Println("Triangle fits in yard:", FitInYard(triangle))   // true
	fmt.Println("Rectangle fits in yard:", FitInYard(rectangle)) // false

	// PrintMe fonksiyonunu kullanarak farklı tipleri yazdır
	PrintMe(42)           // int
	PrintMe(3.14)         // float64
	PrintMe("Hello, Go!") // string
}
```