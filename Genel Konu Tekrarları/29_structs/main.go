// farklı veri tiplerinden faydalanma durumu go da yeni bir veri tipi olarak düşünülür ve bu veri tipine structs denir. Yani structs, programlamada veri gruplarını bir araya getiren ve farklı veri türlerini içeren bir yapıdır.

package main

import "fmt"

func main() {

	var employee struct { // struct aşağıda yer alan string, int, bool veri tiplerimizin üzerine inşa edilecek.
		name      string
		age       int
		isMarried bool
	}

	fmt.Println(employee)         // { 0 false}
	fmt.Printf("%#v\n", employee) // struct { name string; age int; isMarried bool }{name:"", age:0, isMarried:false}
	fmt.Println(employee.age)     // 0

	fmt.Println("**********")

	employee.name = "Veli"
	employee.age = 30
	employee.isMarried = false
	fmt.Printf("%#v\n", employee)   // struct { name string; age int; isMarried bool }{name:"Veli", age:30, isMarried:false}
	fmt.Println(employee.name)      // Veli
	fmt.Println(employee.age)       // 30
	fmt.Println(employee.isMarried) // false

	fmt.Println("**********")

	// struct'ı kendi tanımladığımız bir veri tipi olarak düşünebiliriz. Birden fazla değer atamak istediğimizi düşünelim yukarıdaki yapıda örneğin employee2, employee3 şeklinde tanımlamak yapabilirdik. Evet bu durum yanlış değil ancak yüzlerce, binlerce struct oluşturmamız gerekirdi. Bu durumda tek tek struct oluşturmak doğru bir yapı olmazdı. Bu yüzden "var" anahtar kelimesine yerine "type" anahtar kelimesi ile yeni bir type oluşturmamız daha mantıklı olacaktır.

	type employees struct { // bir değişken değil veri tipi oluşturuyoruz. yukarıda değişken oluşturduk burada ise oluşturduğumuz şey bir veri tipi. ---BU ÇOK ÖNEMLİ--- Yani employess veri tipimizi oluşturduk ve bunu oluştururken struct'tan yararlandık ---> UNDERLYING TYPE
		name      string
		age       int
		isMarried bool
	}

	var e1 employees // e1 değişkeni oluşturduk ve e1'e diyoruz ki senin veri tipin "employees"
	e1.name = "Gurcan"
	e1.age = 40
	e1.isMarried = true

	var e2 employees
	e2.name = "Arin"
	e2.age = 5
	e2.isMarried = false

	// bu şekilde de tanımlama yapabiliriz:
	e3 := employees{
		name:      "Can",
		age:       20,
		isMarried: false,
	}

	fmt.Printf("%#v\n", e1) // main.employees{name:"Gurcan", age:40, isMarried:true}
	fmt.Printf("%#v\n", e2) // main.employees{name:"Arin", age:5, isMarried:false}
	fmt.Printf("%#v\n", e3) // main.employees{name:"Can", age:20, isMarried:false}

	// Tekrarlıyorum; e1 ve e2'nin veri tipi "struct" değil. e1 ve e2'nin veri tipi "employees"	("employees" kendi oluşturduğumuz bir veri tipidir. "employees" veri tipini oluştururken çekirdek struct veri tipinden faydalandık)

	// --- NOT --- Kendi oluşturduğumuz veri tipleri tüm paket içerisinde ulaşılabilir olması için main fonksiyonunun dışarısına yazılır. Ve dışarıdan ulaşılmasını istiyorsak büyük harfle başlatılır (burada örnek olması için içine yazdık)

	fmt.Println("**********")

	// ayrıca yapımızın içine slice'da tanımlayabiliriz:

	type persons struct {
		name      string
		age       int
		isMarried bool
		kids      []string // slice ekledik
	}

	p1 := persons{
		name:      "Veysel",
		age:       50,
		isMarried: true,
		kids: []string{
			"Ayşe",
			"Kemal",
		},
	}

	fmt.Printf("%#v\n", p1)      // main.persons{name:"Veysel", age:50, isMarried:true, kids:[]string{"Ayşe", "Kemal"}}
	fmt.Printf("%#v\n", p1.kids) // []string{"Ayşe", "Kemal"}
	fmt.Println(p1.kids)         // [Ayşe Kemal]
	fmt.Println(p1.kids[0])      // Ayşe
	fmt.Println(p1.kids[1])      // Kemal

}
