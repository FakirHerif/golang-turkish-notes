// Goroutines sayesinde birden fazla işlemin aynı anda yapılması sağlanır. Eşzamanlı işlemler, yani eşzamanlı olarak yapılan görevlerin herbirine bir goroutine denir.

// Programlamada kodun en hızlı çalışma düzeni her zaman sıralı işlemler değildir. Yani bu ne demek örneğin bir x ve y fonksiyonu oluşturduk ve önce y'yi çağırdık ve ardından x'i çağırdık. Bu durumda önce y çalışacak daha sonra x çalışacak. Evet ama bir websitemiz için y fonksiyonunun dışarıdan bir veri aldığını düşünelim, bir backendden olabilir. Peki bu işlemin uzun sürdüğünü düşünelim bu durumda y işlemi tamamlanmadan x işlemine geçmeyecek ve bu durumda websitemize dışarıdan gelen kullanıcı bu uzun süren işlemin tamamlanmasını muhtemelen beklemeyecek. Bu durum bizim açımızdan çok can sıkıcı olabilir. İşte bu durumda ne yapmak isteriz? Birden fazla işlemin aynı anda yapılmasını isteriz.

// VEYA... kendi isteğimizle işlemlerin farklı sırada çalışmasını isteyebiliriz. Bunun için tabi ki en basit olan yöntem olarak kod bloğumuza gidip çağırdığımız fonksiyonların yerlerini değiştirebiliriz ancak bu çok efektif bir yöntem olmaz. Binlerce satırlık kod bloğunda veya sürekli yapılacak durumlarda bu değişikliği yapmak dediğim gibi hiç efektif olmaz.

// İşte bu durumda eşzamanlı çalışmalar için goroutines kullanırız. Go'da kullandığımız methodların başına "go" anahtar kelimesini koyarız.

package main

import (
	"fmt"
	"time"
)

func main() {

	/* 	 	go printY()
		   	// fmt.Println()
		   	go printX()
	--- Burda go' ları başına koyduk ancak herhangi bir çıktı almadık. Neden? Bunun sebebi şudur; aslında main fonksiyonunu çağırdığımız anda yani main fonksiyonu çalışmaya başlar başlamaz main go routine oluşturmuş oluyoruz. Yani çalışan tüm go programlarında bizim bir tane ana goroutine'ımız var yani main go routine. Yani bunun içinde oluşturduğumuz diğer routine'ları beklemeden çalışmanın sonuna geliyor, main go routine tamamlanıyor ve diğer routine'ların başlayıp başlamadığını bile dikkate almadan programı tamamlıyor bu yüzden burda herhangi bir çıktı alamıyoruz.

	Peki bu durumda ne yapabiliriz? Main goroutine'ın çalışmasını belirli bir süre durdurabiliriz time paketinden sleep methodunu çağırabiliriz. */

	go printX()
	go printY()
	time.Sleep(time.Second) // burda main fonskiyonunun tamamlanmasını 1 saniye durdur, yani 1saniye bekle diğer işlemler yapılsın diyoruz ve konsolda yazdırdığımız zaman sonuç olarak bazen "XXXXXXXXXXXXXXXXXXXXYYYYYYYYYYYYYYYYYYYY" çıktısı bazen de "YYYYYYYYYYYYYYYYYYYYXXXXXXXXXXXXXXXXXXXX" çıktısı alıyoruz. Bazen daha farklı çıktılarda alabiliriz (yani 2bin tane x yazdırıyoruz diyelim y'ler her zaman en başta veya en sonda gelmeyebilir x çalışmaya devam ederken x'in bulunduğu satır içinde de yazdırılabilir). (Bu durumların sırasını kontrol edebileceğimiz yöntemler var, channel gibi vs ama daha o konuya gelmedik.) Ayrıca bu durumda da şöyle bir sorunla karşılaşıyoruz. Örneğin 20 tane x değilde 50bin tane x yazdırmak istediğimizi düşünelim. Bu durumda 1 saniye içinde 50bin tane x yazdıramadan 1 saniyelik süre dolacak ve main go çalışıp programı tamamlamış olacak.

	// Kısaca time.Sleep'de temelde 2 sorunumuz var; 1- sadece goroutine'lerle methodların hangi sırayla çalışacağını tam olarak bilemiyoruz. 2- belirlediğimiz süre içerisinde methodumuz tamamlanmayabilir ve eksik/yarım kalabilir.

	// Tek çekirdekte goroutines çalışma şekli şudur; sadece tek çekirdekte aynı anda sadece 1 tane goroutine çalışır. Ancak goroutine'ın çalışması bitmeden diğerleri devreye girebilir. En efektif olacak şekilde çalışma durumu belirlenir.

	// ÖNEMLİ NOT: Eşzamanlılığı paralellik ile karıştırmamak gerek. Eşzamanlılıkta tüm işlemler aynı zamanda aynı anda başlamıyor. Eşzamanlı(concurrency) işlemlerde örneğin 1. goroutine çalışmaya başladı veri gelmeye devam ediyor uzun sürüyor ya da bir sorun var hemen diğerine geçiyor. bunda da sorun var hemen bir diğerine geçiyor. Ama paralellikte(parallelism) tüm işlemler aynı anda başlatılıyor.

}

func printX() {
	for i := 0; i < 20; i++ {
		fmt.Print("X")
	}
}

func printY() {
	for i := 0; i < 20; i++ {
		fmt.Print("Y")
	}

}

// x'lerin her zaman daha önce yazılmasını istediğimiz bir senaryomuz olsun. "sync.WaitGroup" kullanabiliriz:

/*
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // yöntemimizi oluşturduk

func main() {

	wg.Add(1) // Burada main go routine'a sen 1 tane goroutine bekleyeceksin diyoruz

	go printX()
	wg.Wait() // Burada da alttaki fonksiyonda "wg.Done()" sinyalini veriyoruz ve buradan önceki goroutine tamamlanana kadar bekle ve tamamlandıktan sonra alttaki methoddan yani "printY()" methodundan devam et diyoruz.
	printY()
}

func printX() {
	for i := 0; i < 20; i++ {
		fmt.Print("X")
	}
	wg.Done() // işlemin tamamlandığı sinyalini gönderiyoruz
}

func printY() {
	for i := 0; i < 20; i++ {
		fmt.Print("Y")
	}
}

// bu senaryoda her zaman x önce çalışacak daha sonra y çalışacak
*/
