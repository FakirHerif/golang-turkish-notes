// 1- studentName --> John Doe, grade --> 77, isPassed --> true
// 2- farklı yöntem ile oluşturup çıktısını yazdırın
// 3- bu değişkenleri tek satır içinde tanımlayınız
// 4- Declaration, assign, initialization, inital value kavramlarını açıklayın (terminoloji)
// 5- statically typed vs dynamically typed ifadelerini açıklayın
// 6- := ve = arasındaki farkı gösteriniz

package main

func main() {

	/* 1-
	    var studentName string = "John Doe"
	   	var grade int = 77
	   	var isPassed bool = true
	*/

	/* 2-
		var studentName = "John Doe"
	   	var grade = 77
	   	var isPassed = true

		studentName := "John Doe"
		grade := 77
		isPassed := true
	*/

	/* 3-
	   var studentName, grade, isPassed = "John Doe", 77, true

		studentName, grade, isPassed := "John Doe", 77, true
	*/

	/* 4-
	   - Declaration = Beyan etmek. Yani bu değişkeni ileride kullanıcam:
	   		var studentName string

			Not: Bir değişken aynı scope içinde ikinci defa declare edilemez!

	   - Assign = Yukarıda oluşturduğumuz değişken hafıza bir yer kapladı, daha sonra değeri git o hafıza yer kaplayan yere yerleştir diyoruz
	   		var studentName string = "John Doe"
		"John Doe" değerimizi studentName değişkenimize assign etmiş olduk

		- Initialization = Başlatma, bir değeri oluşturup başlatma anlamına geliyor. Yani değişkeni oluşturduk ve değer atadık bu duruma initialization deniyor.

		- Initial Value = İlk değer anlamına geliyor. Değişkene atadığımız ilk değeri ifade eder.
	*/

	/* 5-
	Örneğin Python, js dinamik yazıma sahiptir, yani değişken oluşturduğumuz zaman değişkenin tipini belirtmeden yazıyoruz. Örneğin:
	var meyve = "elma" --> gördüğünüz gibi herhangi bir string tip olduğunu belirtmedik ama js bunu algılayıp bize string olduğu sonucunu veriyor. Ardından;
	meyve = 12 ---> yazıp değiştirebiliriz ve bu sefer veri tipini sorguladığımızda bize int olduğunu söyleyecek.
		ANCAK, bunu go gibi statik olan dillerde yapmamız mümkün değil. Veri tiplerini değişkene bildirmeliyiz. Örneğin;
	var x int = 100 --> dedikten sonra,
	var = "Ali" --> yazmayı denersek hata alırız bu yanlış bir kullanımdır.
	*/

	/* 	6-
	:= Kısa gösterimdir, bununla hem declare edip hem assign ediyoruz
	studentName := "Ali"
	studentName = "Veli"
	olarak değişken değerini değiştirebiliriz ancak aynı değişken isimlerine sahip oldukları için ikisinde de := kullanamayız
	*/

	/* 	fmt.Println(studentName)
	   	fmt.Println(grade)
	   	fmt.Println(isPassed) */
}
