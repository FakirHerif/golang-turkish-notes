/* package main

import "fmt"

func main() { // GO ---> PASS BY VALUE

	x := 5
	fmt.Println(x) // 5
	double(x)
	fmt.Println(x) // 5 ---> ORİJİNAL DEĞER DEĞİŞMEDİ

}

func double(num int) {
	num *= 2
	fmt.Println(num) // 10
} */

// GO FONKSİYONLARDA ARGÜMANIN KOPYASINI ALIR. NORMALDE GO PASS BY VALUE OLDUĞU İÇİN ORİJİNALİ DEĞİŞTİRMEZ ANCAK BU DURUM SLICELARDA GEÇERLİ DEĞİLDİR. SLICEDA ORİJİNAL DEĞERDE DEĞİŞİR:

/* package main

import "fmt"

func main() { // SLICE
	mySlc := []int{1, 10, 100}
	fmt.Println(mySlc) // [1 10 100]
	double(mySlc)
	fmt.Println(mySlc) // [2 20 200] ---> ORİJİNAL DEĞER DEĞİŞTİ (GO KENDİSİ PASS BY VALUE OLARAK ÇALIŞIYOR BURDA DA AMA SLICE'IN KENDİ YAPISINDAN DOLAYI BU DEĞER DEĞİŞİYOR.)
}

func double(slc []int) {
	for i := 0; i < len(slc); i++ {
		slc[i] *= 2
	}
	fmt.Println(slc) // [2 20 200]
}
*/

/* package main

import "fmt"

func main() { // ARRAY
	myArr := [3]int{1, 10, 100}
	fmt.Println(myArr) // [1 10 100]
	double(myArr)
	fmt.Println(myArr) // [1 10 100] ---> ORİJİNAL DEĞER DEĞİŞMEDİ
}

func double(arr [3]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] *= 2
	}
	fmt.Println(arr) // [2 20 200]
} */

// PEKİ BU KONULARIN POINTERLARLA NE ALAKASI VAR? POINTERLARI NORMALDE BİR DEĞERİ OLDUĞU YERDE DEĞİŞTİREBİLMEK İÇİN KULLANIYORUZ, BU DURUMUN BİZE NE GİBİ BİR AVANTAJI OLABİLİR? VERİMİZİN ÇOK BÜYÜK BİR ARRAYDEN VEYA SLICEDAN OLUŞTUĞUNU DÜŞÜNELİM İŞLEM YAPARKEN BUNUN KOPYASINI ALMAK EFEKTİF BİR ÇÖZÜM OLMAZ. KOPYASINI ALMAK YERİNE ONU OLDUĞU YERDE DEĞİŞTİRMEK YANİ POINTERINI ALABİLMEK BİZİM AÇIMIZDAN DAHA AVANTAJLI OLUR. YANİ KISACA DEĞERİN KOPYASINI ALMADAN BULUNDUĞU YERDE DEĞİŞTİREBİLİRİZ.

package main

import "fmt"

func main() { // GO ---> PASS BY VALUE

	x := 5
	fmt.Println(x) // 5
	double(&x)
	fmt.Println(x) // 10 ---> pointer ile x değerinin değişmesini sağladık

}

func double(num *int) { // POINTER METHOD
	fmt.Println(num) // 0xc00000a0a8
	*num *= 2
	fmt.Println(*num) // 10
}
