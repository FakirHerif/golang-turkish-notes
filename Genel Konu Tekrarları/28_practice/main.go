/*
1. 5 string elemandan oluşan bir array ve 5 int elemandan oluşan slice oluşturup index numaralarıyla birlikte gösteriniz.
2. [0,1,2,3,4,5,6,7,8] slice dan [0,1,2,3], [4,5,6], [6,7,8], [2,3,6,7] slice'larını oluşturunuz.
3. slice'lar için copy methodunu ve assign ("=") farkını açıklayınız.
4. map gösterimi ile yazar ve yazarlara ait kitapların isimlerini ofr range ile gösterin.
*/

package main

import "fmt"

func main() {

	// 1.
	myArr1 := [5]string{"tahran", "belgrad", "roma", "tiflis", "moskova"}
	fmt.Println(myArr1) // [tahran belgrad roma tiflis moskova]
	for index, value := range myArr1 {
		fmt.Println(index, value) // index numaralarıyla birlikte sonuç döner
	}

	fmt.Println("**********")

	mySlc1 := []int{1, 2, 3, 4, 5}
	fmt.Println(mySlc1) // [1 2 3 4 5]
	for index, value := range mySlc1 {
		fmt.Println(index, value) // index numaralarıyla birlikte sonuç döner
	}

	fmt.Println("**********")

	// 2.
	mySlc2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	mySlc2PartOne := mySlc2[:4]
	fmt.Println(mySlc2PartOne) // [0 1 2 3]

	mySlc2PartTwo := mySlc2[4:7]
	fmt.Println(mySlc2PartTwo) // [4 5 6]

	mySlc2PartThree := mySlc2[6:]
	fmt.Println(mySlc2PartThree) // [6 7 8]

	mySlc2PartFour := append(mySlc2[2:4], mySlc2[6:8]...)
	fmt.Println(mySlc2PartFour) // [2 3 6 7]

	fmt.Println("**********")

	// 3.

	mySlc3 := []int{1, 2, 3}
	mySlc3PartTwo := make([]int, 2)

	fmt.Println(mySlc3)        // [1 2 3]
	fmt.Println(mySlc3PartTwo) // [0 0]

	copy(mySlc3PartTwo, mySlc3)
	fmt.Println(mySlc3)        // [1 2 3]
	fmt.Println(mySlc3PartTwo) // [1 2] --->> sadece iki elemanlık yer oluşturduğumuz için iki elemanı kopyaladı.

	fmt.Println("**********")

	mySlc3[0] = 100
	fmt.Println(mySlc3) // [100 2 3]
	mySlc3PartTwo = mySlc3
	fmt.Println(mySlc3)        // [100 2 3]
	fmt.Println(mySlc3PartTwo) // [100 2 3]	--> iki elemanlık yer olmasına rağmen mySlc3PartTwo = mySlc3 eşitledik bu yüzden birebir olarak aynı değerleri almasını sağladık.

	fmt.Println("**********")

	// 4.
	myMap4 := map[string][]string{
		"Yaşar Kemal":     {"ince memed", "yer demir gök bakır"},
		"Sabahattin Ali":  {"kuyucaklı yusuf", "kürk mantolu madonna", "değirmen"},
		"Haruki Murakami": {"1Q84", "dans dans dans", "kumandanı öldürmek"},
	}
	fmt.Println(myMap4)                       // tüm elemanlar
	fmt.Println(myMap4["Sabahattin Ali"])     // [kuyucaklı yusuf kürk mantolu madonna değirmen]
	fmt.Println(myMap4["Haruki Murakami"][0]) // 1Q84

	for key, value := range myMap4 {
		fmt.Println("Yazar: ", key)
		fmt.Println("Bazı Kitapları: ")
		for i, v := range value {
			fmt.Println("\t", i+1, v) // liste halinde tüm elemanları gösterecek.
		}
	}
}
