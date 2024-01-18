// Slice'lar arraylere benzer yapıdadır ancak slice'lar arraylerin biraz daha genişletilmiş versiyonudur.

package main

import "fmt"

func main() {

	myArr := [3]int{1, 2, 3}
	myArr2 := [...]int{1, 2, 3, 4}
	fmt.Println(myArr)
	fmt.Println(myArr2)
	fmt.Println(len(myArr))
	fmt.Println(len(myArr2))

	// arrayler oluşturulurken eleman sayısı bilirken ancak slicelar oluşturulurken eleman sayısını bilmiyor.

	mySlc := []int{1, 2, 3} // slicelar oluşturulurken eleman sayısı belirtilmiyor
	fmt.Println(mySlc)
	fmt.Println(len(mySlc))

	var myArr3 [4]int
	fmt.Println(myArr3) // [0 0 0 0]
	myArr3[0] = 5       // arrayda ilk elemanı 5 yaptık
	fmt.Println(myArr3) // [5 0 0 0]

	// şimdi bunu slice ile deneyelim:

	/* 	var mySlc2 []int
	   	fmt.Println(mySlc2) // []
	   	mySlc2[0] = 5
	   	fmt.Println(mySlc2) // panic: runtime error: index out of range [0] with length 0

		Gördüğünüz gibi sliceda eleman sayısını bilmediğimiz için herhangi bir eleman olmadığı için ilk elemana bir değer atamaya çalışırsak hata alırız. Peki slice'a eleman nasıl atarız, bunu yapalım:
	*/

	var mySlc3 []int
	mySlc3 = make([]int, 4) // make methodu yardımıyla slice oluşturduk, yani elemanları belirli bir değere (genellikle sıfır) ayarlamak için yaygın kullanılan bir yöntemdir
	fmt.Println(mySlc3)     // [0 0 0 0]
	mySlc3[0] = 10
	fmt.Println(mySlc3) // [10 0 0 0]

	// arrayde arrayi kopyaladığımız zaman arrayin kendisinin kopyasını alırız ancak slice kopyaladığımız zaman slice'ın referanslarının kopyasını alırız

	arr1 := [3]int{1, 2, 3}
	fmt.Println(arr1) // [1 2 3]
	arr2 := arr1      // arr2'ye arr1'i atadık
	fmt.Println(arr2) // [1 2 3]

	arr2[0] = 100     // arr2 nin ilk değerini 100 yaptık
	fmt.Println(arr2) // [100 2 3]
	// peki bu durumda arr1'e ne oldu hemen bakalım:
	fmt.Println(arr1) // [1 2 3] ---> arr1 hiçbir şekilde değişikliğe uğramadı ve ilk oluşturduğumuz şekilde duruyor. (PASS BY VALUE)

	println("************")
	// sliceda bu durumu inceleyelim, referans kopyasını almamızı tam olarak örnekle görelim: (NOT: Genelde diğer dillerde arrayler kopyalandığı zaman referanslarının kopyası alınmış olur)

	slc1 := []int{1, 2, 3}
	fmt.Println(slc1) // [1 2 3]
	slc2 := slc1
	fmt.Println(slc2) // [1 2 3]
	slc2[0] = 33
	fmt.Println(slc2) // [33 2 3]
	fmt.Println(slc1) // [33 2 3] ---> gördüğümüz üzere arrayde olduğu gibi ilk halinin çıktısını almadık. işte tam olarak referansının kopyalanması durumu bu oluyor. slc1'de durumdan etkilendi ve slc2'ye atanan değerleri almış oldu. (PASS BY REFERENCE)

	// Slice'lar array'lerin üzerine oluşturulmuş bir yapıdır. Yani bir slice oluşturduğumuz zaman arkaplanda underlying array denilen bir array oluşur. Veya underlying array'i kendimiz de oluşturup bu array'den slice oluşturabiliriz

}
