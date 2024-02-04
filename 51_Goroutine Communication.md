# Goroutine Communication

Goroutines, genellikle daha büyük bir görevi yerine getirmek üzere bir araya gelir ve genellikle tamamen bağımsız değildir. Örneğin, bir web sunucusu oluşturuyorsanız ve çok sayıda bağlantıyı aynı anda ele almanız gerekiyorsa, her bağlantı için bir gönderim birimi oluşturabilir ve bunları çoklu gönderim birimleriyle yönetebilirsiniz.

Bu gönderim birimleri genellikle belirli bir işlevi yerine getirir ve diğer gönderim birimleriyle işbirliği yaparak daha büyük bir görevi tamamlarlar. Örneğin, web sunucusu senaryosunda, her bir gönderim birimi bir tarayıcı bağlantısını ele alabilir ve bu bağlantı ile iletişim kurabilir. Ancak, bu gönderim birimleri tamamen bağımsız değildir; çünkü aynı web sayfasını sunmaktadırlar ve sayfa verisini paylaşırlar.

İletişim, genellikle veri alışverişi ile sağlanır ve bu noktada "kanallar" devreye girer. Kanallar, gönderim birimleri arasında veri iletimini sağlayan ve tip bazlı veri aktarımına olanak tanıyan yapısal öğelerdir. Kanalların varsayılan kapasitesi sıfırdır ve bu nedenle gönderme ve alma işlemleri bloke edicidir. Yani, bir gönderim birimi veri gönderdiğinde, alıcı bir gönderim birimi tarafından alınana kadar bekler ve tersi durumda da aynıdır.

Kanallar, sadece veri iletişimini değil, aynı zamanda senkronizasyonu da sağlar. Örneğin, bir gönderim birimi diğerine veri gönderdiğinde, gönderen bekleyebilir ve alıcı veriyi aldığında devam edebilir. Bu durum, programın düzenli ve kontrol edilebilir olmasına yardımcı olur.

Kapasiteli kanallar, belirli bir miktar veriyi depolamak için bir arabellek sağlar. Bu, gönderim birimlerinin eşzamanlı çalışmasını kolaylaştırır, çünkü bir gönderim birimi bir miktar veri gönderebilir ve diğer gönderim birimi bunları okuyana kadar devam edebilir.

**Go dilinde kanalların nasıl kullanıldığını anlamak için örnek verelim:**

Örnek 1:

```
package main

import "fmt"

// prod fonksiyonu, iki sayıyı çarpar ve sonucu bir kanala gönderir.
func prod(v1, v2 int, c chan int) {
	c <- v1 * v2
}

func main() {
	// Bir tampon olmayan bir kanal oluşturun
	c := make(chan int)

	// İki goroutine oluşturun ve her biri prod fonksiyonunu çalıştırsın
	go prod(1, 2, c)
	go prod(3, 4, c)

	// İki sonucu kanaldan alın
	a := <-c
	b := <-c

	// Sonuçları çarpın ve ekrana yazdırın
	result := a * b
	fmt.Println(result)
}
```

Bu örnekte, prod fonksiyonu iki sayıyı çarpar ve sonucu c adlı bir kanala gönderir. main fonksiyonunda ise bir kanal oluşturulur (make(chan int)), ardından iki ayrı goroutine başlatılır. Her bir goroutine, prod fonksiyonunu çalıştırarak sonuçları kanala gönderir.

Ana goroutine, kanaldan iki sonucu alır (a ve b), bu sonuçları çarpar ve ekrana yazdırır. Kanal operasyonları, gönderme (c <- v1 * v2) ve alma (a := <-c) işlemleri arasında senkronizasyon sağlar.

Örneğin çalışma akışı şu şekildedir:

- prod fonksiyonları iki ayrı goroutine'de çalıştırılır ve çarpım sonuçları kanala gönderilir.
- main fonksiyonu, bu iki sonucu kanaldan alır (a ve b).
- a ve b değerleri çarpılır ve sonuç (result) hesaplanır.
- Sonuç ekrana yazdırılır.

Kanallar, iki goroutine arasında senkronizasyon ve veri iletişimi sağlamak için kullanılır. Bu örnekte, kanal kullanımıyla farklı goroutine'ler arasında veri iletimi ve sonuçların toplanması gerçekleştirilmiştir.

Örnek 2:

```
package main

import (
	"fmt"
	"time"
)

func main() {
	// Kanal oluştur
	messageChannel := make(chan string)

	// Goroutine ile mesaj gönderme
	go func() {
		time.Sleep(2 * time.Second) // 2 saniye bekleyelim
		messageChannel <- "Merhaba, bu bir kanal mesajıdır!"
	}()

	// Ana goroutine'de kanaldan mesajı al
	message := <-messageChannel

	// Alınan mesajı yazdır
	fmt.Println(message)
}
```

Bu örnekte, make(chan string) ile bir kanal oluşturuyoruz. Daha sonra, go anahtar kelimesi ile bir goroutine başlatıyoruz. Bu goroutine, 2 saniye bekledikten sonra kanala bir mesaj gönderiyor. Ana goroutine ise kanaldan gelen mesajı alıp yazdırıyor.

Bu örnekte, kanalların senkronize bir şekilde veri iletimini sağladığını görebilirsiniz. Gönderim işlemi tamamlandıktan sonra alım işlemi gerçekleşir ve mesaj başarıyla alınarak ekrana yazdırılır. Kanallar, bu tür senkronizasyon ve iletişim senaryolarında kullanılarak farklı gönderim birimleri arasında güvenli bir şekilde veri alışverişi yapılmasını sağlar.

## Unbuffered Channel:

Unbuffered bir kanal, özel bir kapasiteye sahip olmayan bir kanaldır. Gönderilen bir veri, alıcı tarafından alınmadan önce bekler ve bu durum senkronizasyona neden olur. Gönderme ve alma işlemleri eş zamanlı olarak gerçekleşir.

Örnek:

```
package main

import (
	"fmt"
)

func main() {
	// Unbuffered bir kanal oluştur
	ch := make(chan int)

	// Gönderme işlemi (Bu, alıcı hazır olana kadar bloke olur)
	go func() {
		ch <- 42
	}()

	// Alım işlemi
	value := <-ch

	// Alınan değeri yazdır
	fmt.Println(value)
}
```

Bu örnekte, gönderme işlemi (ch <- 42) alıcı (<-ch) hazır olana kadar bekleyecektir.

## Blocking and Synchronization:

Kanallar, veri gönderme ve alma işlemleri sırasında senkronizasyonu sağlar. Gönderim veya alım gerçekleşene kadar bir işlem bekler (bloker). Bu, farklı gorutinler arasında eşgüdüm ve senkronizasyon sağlar.

Örnek:

```
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	// Gönderim işlemi
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 42
	}()

	// Alım işlemi
	wg.Add(1)
	go func() {
		defer wg.Done()
		value := <-ch
		fmt.Println(value)
	}()

	wg.Wait() // Ana goroutine'in diğerlerini beklemesi
}
```

Bu örnekte, sync.WaitGroup kullanarak gönderim ve alım işlemleri arasında eşzamanlılık sağlıyoruz.

## Channel Capacity:

Kanal kapasitesi, bir kanalın aynı anda kaç veriyi tutabileceğini belirler. Buffer kapasitesi arttıkça, gönderme ve alım işlemleri arasındaki senkronizasyon değişir.

Örnek:

```
package main

import (
	"fmt"
)

func main() {
	// Kapasitesi 2 olan bir kanal oluştur
	ch := make(chan int, 2)

	// Gönderim işlemleri (bloklamadan iki kez gönderilebilir)
	ch <- 42
	ch <- 99

	// Alım işlemleri (bloklamadan iki kez alınabilir)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
```

Bu örnekte, kanalın kapasitesi 2 olduğu için iki kez gönderim ve alım yapabiliyoruz.

## Channel Blocking:

Kanalın bloke olması, belirli bir işlem gerçekleşene kadar diğer işlemlerin beklemesi anlamına gelir. Kanallar, genellikle alıcı hazır olmadan gönderim yaparsa veya tam tersi durumda bloklanır.

## Send ve Receive İşlemleri:

Kanal üzerinden veri gönderme ve alma işlemleri sırasında kullanılan send ve receive işlemleri.

Örnek:

```
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	// Gönderim işlemi
	go func() {
		ch <- 42
	}()

	// Alım işlemi
	value := <-ch

	// Alınan değeri yazdır
	fmt.Println(value)
}
```

Bu örnekte, gönderim ve alım işlemleri (ch <- 42 ve value := <-ch) kullanılmıştır.

## Use of Buffering:

Tampon kullanımı, kanalın bir miktar veriyi bellekte tutma yeteneği anlamına gelir. Bu, gönderme ve alım işlemleri arasındaki hız farklılıklarını tolere etmeye yardımcı olabilir.

Örnek:

```
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)

	// Gönderim işlemleri (tampona 2 kez gönderilebilir)
	ch <- 42
	ch <- 99

	// Alım işlemleri (tampondan 2 kez alınabilir)
	fmt.Println(<-ch)
	time.Sleep(2 * time.Second) // 2 saniye bekleyelim
	fmt.Println(<-ch)
}
```

Bu örnekte, kanalın kapasitesi 2 olduğu için tampona iki kez gönderim ve alım yapabiliyoruz.