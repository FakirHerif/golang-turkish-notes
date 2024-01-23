// GO NESNE YÖNELİMLİ BİR PROGRAMLAMA DİLİ MİDİR? BUNUN CEVABI HEM EVET HEM HAYIRDIR. Go bazı nesne yönelimli yaklaşımları kullanıyor ancak kendi bakış açısıyla kullanıyor. Bu durumu detaylı olarak ilerleyen konularda inceleyeceğiz.

package main

import "fmt"

// bir önceki konuda bahsettiğimiz gibi, tüm pakette ulaşılabilir olması için main fonksiyonu dışında yazıyoruz.
type employee struct {
	name      string
	age       int
	isMarried bool
}

func main() {

	e1 := employee{
		name:      "Gurcan",
		age:       40,
		isMarried: true,
	}

	fmt.Println(e1) // {Gurcan 40 true}
	e2 := e1
	fmt.Println(e2) // {Gurcan 40 true}
	e2.name = "Arin"
	fmt.Println(e2) // {Arin 40 true}
	fmt.Println(e1) // {Gurcan 40 true}

	// STRUCT'lar birbirleri arasında değerleri paylaşırlar referansı değil. Hafızada iki tane aynı yeri tutmazlar, hafızada birbirinden farklı yer tutarlar.

	fmt.Println("**********")

	// Biraz da Klasik nesne yönelimli programlama mantığından bahsedelim; "inheritance" şeklinde bir kavram vardır. Şu şekilde örnek verebiliriz; bir class oluşturulur animal şeklinde, burada hayvanların genel özellikleri yazar, hayvanlarda ortak olan özellikler isim gibi yani tüm hayvanlarda ortak bulunan özellikler yazılır. Ancak bunun yanında aslanın kendi has özellikleri vardır, mesela yeleye sahip olması gibi. Yani bu aslan classı animal class'ını kapsayacak şekilde tasarlanır. Yani "animal" bu noktada super class'tır, "aslan" ise subclass'tır. Aslan animal'ın tüm özelliklerini içerir ve ekstradan kendi özelliklerini taşır. Bu tanımlama klasik OOP'tur.

	// Go'da benzer yaklaşım olarak struct yardımıyla örneğin bir "manager" tipi oluşturuyoruz ve bu "manager" veri tipi "employee"'un sahip olduğu tüm özellikleri taşıyor. Yani "manager" bir sub oluyor. "employee" ise super. Peki neden go'da böyle bir yaklaşımda bulunuyoruz? Çünkü go'da class yoktur.

	type manager struct {
		employee       // gördüğünüz gibi employee'un taşıdığı tüm özellikleri taşıdığı için "name, age ve isMarried" özelliklerini tek tek yazmak yerine direkt employee yazarak tüm bu özellikleri burada tanımlayabiliyoruz.
		hasDegree bool // employee'dan bağımsız olarak ekstra özelliği de bu şekilde tanımlayabiliyoruz.
	}

	m1 := manager{
		employee: employee{
			name:      "Ayşe",
			age:       28,
			isMarried: false,
		},
		hasDegree: true,
	}

	fmt.Println(m1) // {{Ayşe 28 false} true}

	// animal/aslan örneğimizden gidelim. klasik oop mantığında animal bir animal'dır, aslan'da bir animal'dır. ---> (IS A RELATION)

	// ÖNEMLİ NOT: YUKARIDA Kİ ÖRNEKTE; "employee" ve "manager" tamamen birbirinden bağımsız ve farklı olan struct üzerine kurulmuş olan veri tipleridir. Örnekte "m1" içinde yaptığımız işlem kısa bir yazım şeklinden faydalanmaktır. (aslında anlatmak istediğim şey şu; "m1" bir employee değil, "m1" bir manager'dır) ---> (HAS A RELATION)

	/* m1 Örneğinin diğer yazım şekli:

	m1 := manager{}
	m1.name:      "Ayşe"
	m1.age:       28
	m1.isMarried: false
	m1.hasDegree: true

	fmt.Println(m1) // {{Ayşe 28 false} true}
	*/

	fmt.Println("**********")

	// Anonim struct oluşturmak: (bazı durumlarda tek kullanımlık yapılar oluşturmak isteyebiliriz) Örneğin genellikle patron bir tane olur diyebiliriz. Bu durumda gidip yeni bir type "boss" şeklinde yazdırmaya gerek yok. Peki neden? Çünkü patronun bir tane olduğunu varsaydık yani bir tane patron var dedik. İşte bu tarz durumlarda bazen hemen hızlıca bir struct yapısı kullanıp geçmek isteriz. İşte bu tarz durumlarda "anonim struct" yapısını kullanırız.

	theBoss := struct {
		name  string
		money bool
	}{name: "THE BOSS", money: true}

	fmt.Println(theBoss) // {THE BOSS true}

}
