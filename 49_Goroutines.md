# Goroutines

Önceki modüllerde, eşzamanlılık ve paralelizm konularını ele aldık, işlemlerin ve iş parçacıklarının nasıl çalıştığını, eşzamanlı bir şekilde çalıştıklarını ve belki de ek donanımınız varsa paralel olarak çalıştıklarını gördük. Ancak bu konuları nasıl uygulayacağınızı, özellikle de Go dilinde nasıl uygulayacağınızı henüz konuşmadık. İş parçacığı oluşturmak için özel olarak Go diline entegre edilmiş belirli programlama yapılarını kullanmanız gerekiyor ve bu, Go'nun avantajlarından biridir. Bu yapılar kolay kullanımlıdır.

Örneğin, bir Gorutine oluşturmak istediğinizi düşünelim. Gorutine, diğer Gorutinelerle eşzamanlı olarak çalışabilen bir iş parçacığıdır. Bir Gorutine her zaman main fonksiyonunu yürütmek için otomatik olarak oluşturulur. Diğer Gorutineler ise "go" anahtar kelimesi kullanılarak oluşturulur. Sol tarafta, yalnızca bir Gorutine olan main fonksiyonunun bir bölümü var. A = 1, foo çağrısı ve A = 2 tüm bu talimatları yürüten ana Gorutine. Sağ tarafta ise yeni bir Gorutine oluşturuyoruz sadece "foo" fonksiyonunu yürütmek için. Yani, solda main Gorutine'si tüm üç talimatı yürütüyor. Sağda ise "go foo" çağrısı, "foo" fonksiyonunu yürüten yeni bir Gorutine oluşturuyor ve sonra A = 2 talimatını ana Gorutine devam ediyor. "go" ifadesi ile bir Gorutine başlatmak oldukça basit. "go" ifadesi kullanılır ve çalıştırmak istediğiniz fonksiyonun adını veya fonksiyonu içerirsiniz, bu Gorutine o fonksiyonun içindeki kodu yürütür. Bir Gorutine genellikle kodu tamamlandığında sona erer. Main Gorutine tamamlandığında diğer tüm Gorutinelerin de sona ermesi zorunlu hale gelir.

Ancak, main Gorutine erken tamamlanırsa diğer Gorutineler de sona ermek zorunda kalır. Bu, Gorutinelerin kodlarını tamamlamadan önce sona erebileceği anlamına gelir. Bu durumu açıklamak için bir örnek kod parçası var. Bu örnekte, main Gorutine, "new routine" ifadesini yazdıran yeni bir Gorutine oluşturuyor ve ardından main Gorutine'nin farklı bir noktada "main routine" ifadesini yazdırıyoruz. Normalde, "new routine" ve "main routine" ifadelerini ekranda görmeyi bekleriz, ancak bu kodun çoğu zaman yalnızca "main routine" ifadesini yazdırdığını görürüz. Bu neden oluyor? Çünkü main Gorutine, yeni bir Gorutineyi başlatır başlatmaz tamamlanıyor. Main Gorutine'nin tamamlandığı bir noktada, diğer Gorutineler de sona ermek zorunda kalır. Bu durumu önlemek için bir çözüm olarak, main Gorutine'yi belirli bir süre bekletmek adına "time.Sleep(100 * time.Millisecond)" kullanılmıştır. Ancak bu, zamanlamaya dayalı bir çözüm olup, sistem ve dilin gelecekteki değişikliklerine karşı güvenli değildir ve kullanılmamalıdır. Bu tür zamanlamalara güvenmek, hatalı ve belirsiz sonuçlara yol açabilir. Güvenli ve deterministik senkronizasyon yapıları kullanmak her zaman daha iyidir.

**Basit temel bir örnek verelim:**

```
package main

import (
	"fmt"
)

func main() {
	go fmt.Println("New routine")
	fmt.Println("Main routine")
}
```

Bu örnekte, bir Gorutine başlatıyoruz (go fmt.Println("New routine")), ancak ana program bu Gorutine'nin tamamlanmasını beklemiyor. Ana program, Gorutine çalışmaya başladığında kendi yürütmesine devam eder ve program sonlandığında Gorutine çalışması tamamlanmamış olabilir.

Bu durum, Gorutinelerin asenkron olarak çalıştıkları gerçeğini yansıtır. Ana programın yürütme akışı ve Gorutinelerin yürütme akışı bağımsızdır. Ana program biterse, Gorutineler de sonlandırılabilir.

Eğer main fonksiyonu içinde bir bekleme yapmazsak, program, Gorutinelerin tamamlanmasını beklemeden sona erecektir. Bu nedenle, Gorutineleri başlatırken bir şekilde bekleme yapmak, sync.WaitGroup veya başka senkronizasyon mekanizmaları kullanmak genellikle daha iyi bir uygulama tasarımıdır. Aksi takdirde, Gorutinelerin işi bitmeden program sona erdiğinde beklemediğiniz beklenmeyen davranışlarla karşılaşabilirsiniz.

**Zamanlama (timing) ile senkronizasyon yapmaya çalışan bir örnek verelim:**

```
package main

import (
	"fmt"
	"time"
)

func main() {
	// Yeni bir Gorutine başlatıyoruz.
	go func() {
		time.Sleep(100 * time.Millisecond) // Belirli bir süre bekleme
		fmt.Println("New routine")
	}()

	// Ana Gorutine, belirli bir süre bekledikten sonra devam eden kodları çalıştırıyor.
	time.Sleep(200 * time.Millisecond) // Yeterince uzun bir süre bekliyoruz
	fmt.Println("Main routine")
}
```

Bu örnekte, ana Gorutine, yeni bir Gorutine başlatıldıktan sonra belirli bir süre boyunca bekler ve ardından devam eden kodları çalıştırır. Ancak, bu yöntemle elde edilen süreç, tamamen zamanlama bağımlıdır. Yani, yeni Gorutine'nin işi ne kadar sürede tamamlayacağından emin olamayız. Örneğin, eğer yeni Gorutine'nin işi bitene kadar beklenenden daha uzun süre alırsa, ana Gorutine devam etmiş olabilir ve beklenmeyen sonuçlar ortaya çıkabilir.

Bu nedenle, sync.WaitGroup gibi senkronizasyon araçları kullanmak daha güvenli ve deterministik bir yaklaşımdır. sync.WaitGroup, Gorutinelerin tamamlanmasını belirli ve güvenli bir şekilde beklememizi sağlar ve böylece programın çalışma sırasını belirleyebiliriz. Zamanlama ile senkronizasyon, programın farklı koşullarda beklenmeyen sonuçlar üretme olasılığını artırır ve bu tür hataları tespit etmek ve düzeltmek zordur.

**Gorutineler arasında güvenli ve deterministik bir senkronizasyon sağlamak amacıyla kullanılan "sync" paketini içeren bir örnek verelim:**

```
package main

import (
	"fmt"
	"sync"
)

func main() {
	// WaitGroup oluşturuyoruz, Gorutinelerin bitişini beklemek için kullanılacak.
	var wg sync.WaitGroup

	// WaitGroup'e bir Gorutine ekliyoruz.
	wg.Add(1)

	// Yeni bir Gorutine başlatıyoruz.
	go func() {
		// Gorutine tamamlandığında WaitGroup'e haber veriyoruz.
		defer wg.Done()
		fmt.Println("New routine")
	}()

	// Main Gorutine, diğer Gorutinelerin tamamlanmasını bekler.
	wg.Wait()

	// Diğer Gorutineler tamamlandıktan sonra devam eden kodlar.
	fmt.Println("Main routine")
}
```

Bu örnekte sync.WaitGroup kullanılarak main Gorutine, diğer Gorutinelerin tamamlanmasını bekliyor. sync.WaitGroup bir sayaçtır ve bu sayaç, beklenen Gorutine sayısını temsil eder. Her yeni Gorutine başlatıldığında Add fonksiyonu ile sayaç bir arttırılır ve Gorutine tamamlandığında Done fonksiyonu ile sayaç bir azaltılır. Wait fonksiyonu ise sayaç sıfırlanana kadar bekler. Bu sayede main Gorutine, diğer Gorutinelerin tamamlanmasını bekleyebilir.

Bu yöntem, belirli bir süre bekletmek yerine, Gorutinelerin gerçekten tamamlandığına dair güvenli bir senkronizasyon sağlar.