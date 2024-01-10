/*
1. 1 ile 10 arasındaki sayıları if yapısıyla tek ve çift olarak yazdırınız.
2. for yapısını kullanarak go'da olmayan while döngüsüne örnek veriniz.
3. switch fallthrough ifadesini açıklayınız.
4. aşağıdaki if döngüsünü daha idiomatic hale getiriniz:
	if b := 20; b % 2 == 0 {
		fmt.Println(b, "çifttir")
	} else {
		fmt.Println(b, "tektir")
	}
5. 1 ile 50 arasındaki asal sayıları yazdırınız.
*/

package main

import "fmt"

func main() {

	// 1.
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "çifttir")
		} else {
			fmt.Println(i, "tektir")
		}
	}

	fmt.Println("**************")

	// 2. (while koşul sağlanana kadar çalışır, bu yüzden aşağıdaki yapımız while yapısıyla hemen hemen aynı mantıkla çalışır)
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	fmt.Println("**************")

	// 3. "fallthrough" diğer case'lerinde işleme alımasını istediğimiz durumlarda kullanılır. Koşulun devam etmesi durumunda kullanıyoruz. Koşulu sağlamayan durumlarda fallthrough kullansak bile koşulu sağlamadığı için ekrana yazdırılmaz.
	switch a := 25; {
	case a < 20:
		fmt.Printf("%d küçüktür 20\n", a)
		fallthrough

	case a < 50:
		fmt.Printf("%d küçüktür 50\n", a)
		fallthrough

	case a < 100:
		fmt.Printf("%d küçüktür 100\n", a)
		fallthrough

	case a < 200:
		fmt.Printf("%d küçüktür 200\n", a)
	}

	fmt.Println("**************")

	// 4. idiomatic kod yazımı demek; parçaların daha bağımsız ve daha sade, okunabilir biçimde yazılması demektir.

	b := 20
	if b%2 == 0 {
		fmt.Println(b, "çifttir")
		// (burada return kullanıyoruz ama kodları durdurmaması için şimdilik kullanmadım, kullanmasamda kodlar sorunsuz çalışıyor ancak ilerleyen konularda neden return kullanmamız gerektiğini anlatacağım.)
	}
	fmt.Println(b, "tektir")

	fmt.Println("**************")

	// 5.
	var c, d int
	for c = 2; c < 50; c++ {
		for d = 2; d <= (c / d); d++ {
			if c%d == 0 {
				break
			}
		}

		if d > (c / d) {
			fmt.Printf("%v bir asal sayıdır\n", c)
		}
	}
}
