// kısaca paket nedir? Benzer işlemleri kod parçaların toplamına paket denir

package main

import (
	"fmt" // fmt bir pakettir
	"strings"
)

func main() {

	fmt.Println(strings.Contains("seafood", "foo")) // strings paketine ait Contains methodunu kullanıyoruz ve burada "foo" kelimesinin "seafood" kelimesi içinde olup olmadığı kontrolü gerçekleştiriyoruz, varsa true yoksa false dönmesini sağlıyor. Bu örnekte gördüğümüz üzere foo yer aldığı için true döner.
	fmt.Println(strings.Contains("seafood", "bar")) // "bar", "seafood" içinde yer almadığı için false döner
	fmt.Println(strings.Contains("seafood", ""))    // boş string yer alır bu yüzden true döner
	fmt.Println(strings.Contains("", ""))           // boş string boş string içinde vardır, true döner

	fmt.Println(strings.Count("animatrix", "a")) // "a" 'nın "animatrix" içinde kaç tane olduğu sonucunu döndürür. Yani 2

	fmt.Println(strings.ToUpper("Gopher")) // tüm harfleri büyük harfe çevirir. "GOPHER" döner
}

// go'da tüm paketleri ezberlemeye gerek yok, zamanla hangi paketin nerede kullanılacağını öğrenmek en iyi yoldur.
