/*
General:

%v	the value in a default format
	when printing structs, the plus flag (%+v) adds field names
%#v	a Go-syntax representation of the value
%T	a Go-syntax representation of the type of the value
%%	a literal percent sign; consumes no value

Boolean:

%t	the word true or false

Integer:

%b	base 2
%c	the character represented by the corresponding Unicode code point
%d	base 10
%o	base 8
%O	base 8 with 0o prefix
%q	a single-quoted character literal safely escaped with Go syntax.
%x	base 16, with lower-case letters for a-f
%X	base 16, with upper-case letters for A-F
%U	Unicode format: U+1234; same as "U+%04X"

*/

package main

import "fmt"

func main() {

	// (Print - Println) - Printf Farkı

	// Print ve Println aldıkları parametreyi yani neyi yazdıracaklarsa raw string şeklinde yazdırırlar yani metin olarak işlerler. Print fonksiyonu parametreyi olduğu gibi yani formatlanmamış olarak yazdırırken Println ise parametreyi yazdırdıktan sonra bir satır atlar. Printf ise formatlanmış şekilde yani biçimlendirilmiş bir şekilde yazdırır.

	fmt.Println("Merhaba") // Üstte yazdığımız println yeni bir satır başlangıcı yapar. bu yüzden bu ikisi yazdırılırken alt alta 2 adet "Merhaba" yazdırılır.
	fmt.Print("Merhaba")

	fmt.Print("Selam") // Üstte print yazarsak eğer alt alta yazdırmayacak ve 2 adet "Selam" yan yana yazdırılacak. Hatta üstteki iki ifade de bulunan "Merhaba" ' da bu yazdırılanların solunda yer alacak ve çıktımız MerhabaSelamSelam şeklinde olacak
	fmt.Println("Selam")
	fmt.Println("") // sadece yeni bir satır başlangıcı yapıyor

	name := "Ali"
	/* fmt.Print(name)
	fmt.Println(name)
	fmt.Printf(name) */

	fmt.Print("Benim Adım", name)   // "Benim AdımAli" araya herhangi bir boşluk bırakmadı
	fmt.Println("")                 // yeni bir satır başlangıcı
	fmt.Println("Benim Adım", name) // "Benim Adım Ali" şeklinde boşluk bırakarak yazdırdı
	fmt.Printf("Benim Adım", name)  // "Benim Adım%!(EXTRA string=Ali)" şeklinde yazdırıyor ancak eklememiz gereken ifade/ifadeler var;
	fmt.Println("")
	fmt.Printf("Benim Adım %v", name) // "Benim Adım Ali"
	fmt.Println("")
	fmt.Printf("Benim Adım %v %T", name, name) // "Benim Adım Ali string" Ayrıca %T ile değişkenin tipini de yazdırdık iki ifade eklediğimiz için iki defa name belirtiyoruz yani her biri için bir tane. Örneğin ek olarak %X ekleseydik üç adet olacaktı ve üç adet name eklememiz gerekecekti:
	fmt.Println("")
	fmt.Printf("Benim Adım %v %T %X", name, name, name) // "Benim Adım Ali string 416C69"

	fmt.Println("")

	x := 100
	fmt.Printf("%b", x) // %b bize 100'ün ikilik sayı sistemindeki ifadesini gösteriyor. ==> "1100100"

	fmt.Println("")

	y := 20
	z := 30
	fmt.Printf("%d %o", y, z) // %d 10luk tabanda, %o ise 8lik tabanda gösteriyor. ==> "20 36"

	fmt.Println("")

	nameTwo, age := "Ali", 30
	fmt.Print("Benim Adım ", nameTwo, " ve ben ", age, " yaşındayım") // "Benim Adım Ali ve ben 30 yaşındayım" şeklinde yazdırıyoruz. Burada gördüğünüz gibi boşlukları koymamız gerekiyor. Ancak bunu Println ile yazdırırsak bu boşlukları koymamıza gerek kalmayacak:
	fmt.Println("")
	fmt.Println("Benim Adım", nameTwo, "ve ben", age, "yaşındayım") // "Benim Adım Ali ve ben 30 yaşındayım" şeklinde çıktımızı aldık, gördüğümüz gibi print'te yazdırırken koyduğumuz boşlukları burada koymamıza gerek kalmadı ve kendisi otomatik olarak oluşturdu
	fmt.Println("")
	fmt.Printf("Benim Adım %v ve ben %v yaşındayım", nameTwo, age) // "Benim Adım Ali ve ben 30 yaşındayım" şeklinde çıktımızı aldık.

	fmt.Println("")

	// -----------

	// Go'da camel case isimlendirme kullanılır:
	/* 	var coinType string
	   	var custName string */

	// Kısaltmalar büyük harflerle yazılır:
	/* 	var URL	// Url değil
	   	var HTTP // http değil. yanında başka ifadeler olsa dahi kısaltmalar büyük harfle yazılır; "xyzHTTP" gibi */

	// index yapısı üzerinde çalışırken tek harften oluşan ifadeler kullanılır.
	// Örneğin 1. ifade için: "i", 2. ifade için "j", 3. ifade için "k"

	// Paket kapsamında tanımladıklarımızı diğer paketlerden çağırabilmek için ilk harf büyük yazılır: "var Deneme" gibi

	// Paket kapsamında tanımladığımız değişkenlerin ismi kesinlikle ve kesinlikle anlaşılabilir olmalıdır. Daha sonra kodumuzu incelediğimiz zaman ne işe yaradığını/ne anlama geldiğini anlayabilmeliyiz.
}
