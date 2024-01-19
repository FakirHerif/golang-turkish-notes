// bir önceki slices konusunda bahsetmiştik slicelar array veri tiplerinin üzerine geliştirilen farklı veri tipleridir

package main

import "fmt"

func main() {
	underArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(underArray)

	// SLICE'LARI ARRAY'DEN OLUŞTURMA YÖNTEMİ:

	mySlc := underArray[2:5] // 2 numaralı indeksten başla ve 5'e kadar ama 5. indeksi dahil etmeden al diyoruz (başlangıç indeksi dahil bitiş indeksi dahil değil)
	fmt.Println(mySlc)       // [2 3 4]

	myslc2 := underArray[:6] // başlangıç değeri vermedik ve bitiş indeksini 6 olarak belirledik. unutmuyoruz; bitiş değeri dahil değil.
	fmt.Println(myslc2)      // [0 1 2 3 4 5]

	myslc3 := underArray[3:] // başlangıcı 3.indeks olarak belirledik ve bitiş indeksi belirlemedik. Yani 3. indeksten başlayacak ve geri kalan tüm elemanları dahil edecek
	fmt.Println(myslc3)      // [3 4 5 6 7 8 9]

	myslc4 := underArray[:] // başlangıç ve bitişe herhangi bir sınırlama getirmiyoruz. bu bize tüm elemanları yazdıracak.
	fmt.Println(myslc4)     // [0 1 2 3 4 5 6 7 8 9]

	// NOT: EĞER BİR SLICE'DA DEĞİŞİKLİK YAPARSAK BU AYNI UNDERLY ARRAYDEN TÜRETİLMİŞ OLAN DİĞER TÜM SLICE'LARIN TAMAMINI ETKİLER.

	fmt.Println(mySlc) // [2 3 4]
	mySlc[0] = 100     // [2 3 4] olan slice'mızın 0. indeksini 100 yaptık ===> [100 3 4]
	fmt.Println(mySlc) // [100 3 4]
	// ve tüm diğer slicelar da bu durumdan etkilenecek "2" olan değerler "100" olacak:
	fmt.Println(myslc2) // [0 1 100 3 4 5]
	fmt.Println(myslc4) // [0 1 100 3 4 5 6 7 8 9]

	fmt.Println("**************")

	// SLICE'A YENİ ELEMAN EKLEME

	mySlice := []int{1, 2, 3}
	fmt.Println(mySlice)            // [1 2 3]
	mySlice = append(mySlice, 4, 5) // append methoduyla 4 ve 5 elemanlarını ekliyoruz
	fmt.Println(mySlice)            // [1 2 3 4 5]

	mySlice2 := append(mySlice, 4, 5) // mySlice2 oluşturulurken kendine ait bir underarray oluşturuyor.
	fmt.Println(mySlice2)             // [1 2 3 4 5 4 5]

	mySlice[0] = 100
	fmt.Println(mySlice)  // [100 2 3 4 5]
	fmt.Println(mySlice2) // [1 2 3 4 5 4 5]

	fmt.Println("**************")

	mySlice3 := []int{1, 2, 3}
	mySlice4 := []int{4, 5}

	/* mySlice3 = append(mySlice3, mySlice4) */ // mySlice3 int veri tipine sahip, mySlice4'de int veri tipine sahip aslında ikisi de aynı veri tipine sahip gözüküyor bir sorun olmaması lazım append ederken ama maalesef bunda mantık bu şekilde ilerlemiyor. mySlice4 ===> "slice int" veri tipine ait bir slice oluyor. Yani go ikisini farklı veri tipi olarak görüyor. Bu hatayı çözmek için bu şekilde kullanıyoruz:
	mySlice3 = append(mySlice3, mySlice4...)    // [1 2 3 4 5]
	fmt.Println(mySlice3)

	fmt.Println("**************")

	// SLICE'DAN ELEMAN SİLME

	mySlice5 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(mySlice5)                 // [0 1 2 3 4 5 6 7 8 9]
	mySlice5 = mySlice5[2:]               // dilimleme işlemi yapıyoruz. İlk 2 elemanı silmiş olduk
	fmt.Println(mySlice5)                 // [2 3 4 5 6 7 8 9]
	mySlice5 = mySlice5[:len(mySlice5)-3] // dilimleme işlemi yapıyoruz ve son 3 elemanı silmiş oluyoruz. (başlangıcı boş bıraktık ve slice uzunluğundan çıkarmak istediğimiz eleman sayısını çıkardık)
	fmt.Println(mySlice5)                 // [2 3 4 5 6]

	fmt.Println("**************")

	// SLICE'LARA BOŞ DEĞER VERMEK

	var myArr10 [10]int  // arraylere sorunsuz bir şekilde boş değer atayabiliyoruz hata almadan ama bunu bu şekilde slice için yapamayız, bu yüzden slicelar da make methodu kullanıyoruz
	fmt.Println(myArr10) // [0 0 0 0 0 0 0 0 0 0]

	var mySlice10 []int
	mySlice10 = make([]int, 10) // make metoduyla Zero değerleri verdik
	fmt.Println(mySlice10)      // [0 0 0 0 0 0 0 0 0 0]

	var mySlice11 []bool
	mySlice11 = make([]bool, 2) // zero değerler - slice elemanlarına ait olan değerler
	fmt.Println(mySlice11)      // [false false] - slice elemanlarına ait olan değerler

	var mySlice12 []int
	fmt.Printf("%#v", mySlice12) // []int(nil) ---> HİÇBİR DEĞER YOK ! bu slice'ı declare ettik ama oluşturmadık. (OLUŞTURULMAMIŞ SLICE)	---- YOK SLICE

	println()

	mySlice13 := make([]int, 0)
	fmt.Printf("%#v", mySlice13) // []int{} ---> (OLUŞTURULMUŞ SLICE) yani bu slice var elemanı yok, boş bir slice ---- BOŞ SLICE
}
