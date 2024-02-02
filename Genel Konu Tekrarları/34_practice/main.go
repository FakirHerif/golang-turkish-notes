/*
1. Underlying type 'int' olacak şekilde kendi veri tipinizi oluşturunuz. Veri tipi ve değerini '10' yazdırınız.
2. Başlangıç değeri 10 olan bir x değişkeni oluşturun, bulunduğu adres üzerinden y değişkenini tanımlayıp x değerini 20 yapınız.
3. Underlying type 'struct' olan Rectangle type oluşturun. Display, area, circumference methodlarını yazınız.
4. family aile bireyleri şeklinde veri tipi oluşturun (Underlying struct). Aile bireylerinin hepsini farklı şekilde tanımlayın. ve tüm family değerlerini for döngüsünde gösteriniz.
*/

/* package main

import "fmt"

// 1.
type myType int

func main() {

	var z myType
	z = 10
	fmt.Printf("%T, %v\n", z, z) // main.myType, 10

	fmt.Println("***********")

	// 2.
	x := 10
	fmt.Println(x) // 10
	y := &x
	fmt.Println(y) // 0xc00000a0a8
	*y = 20
	fmt.Println(x) // 20

} */

/* package main

import "fmt"

// 3.
type rectangle struct {
	a, b int
}

func (r rectangle) display() {
	fmt.Println("Kenar uzunlukları:", r.a, " ve ", r.b, " olan dikdörtgen ")
}

func (r rectangle) area() int {
	return r.a * r.b
}

func (r rectangle) circumference() int {
	return 2 * (r.a + r.b)
}

func main() {

	r1 := rectangle{3, 8}
	r1.display()
	fmt.Println("Alanı: ", r1.area())            // Alanı:  24
	fmt.Println("Çevresi: ", r1.circumference()) // Çevresi:  22

} */

package main

import "fmt"

// 4.
type family struct {
	name      string
	age       int
	isMarried bool
}

func showFamily() []family {

	family1 := family{
		name:      "Ali",
		age:       30,
		isMarried: false,
	}

	family2 := family{
		name: "Can",
		age:  35,
	}

	family3 := family{
		"Veli",
		40,
		true,
	}

	var family4 family
	family4.name = "Ayşe"
	family4.age = 40
	family4.isMarried = true

	return []family{family1, family2, family3, family4}
}

func main() {

	families := showFamily()
	for i := 0; i < len(families); i++ {
		fmt.Printf("%v, %T\n", families[i], families[i])
		/*
			{Ali 30 false}, main.family
			{Can 35 false}, main.family
			{Veli 40 true}, main.family
			{Ayşe 40 true}, main.family
		*/
	}

}
