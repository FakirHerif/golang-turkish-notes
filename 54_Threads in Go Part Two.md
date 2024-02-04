# 54_53_Threads in Go Part Two

## Once Senkronizasyon (sync.Once)

Go dilinde birden çok goroutine arasındaki senkronizasyonu sağlamak için sync paketi kullanılır. Bu paket, sync.Once gibi yöntemlerle bir işin yalnızca bir kez gerçekleşmesini ve bu işlemin diğer işlemlerden önce tamamlanmasını sağlar.

Örnek:

```
package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func setup() {
	fmt.Println("Init")
}

func dostuff() {
	once.Do(setup)
	fmt.Println("Hello")
	// additional logic
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		dostuff()
	}()

	go func() {
		defer wg.Done()
		dostuff()
	}()

	wg.Wait()
}
```

Bu örnekte, sync.Once kullanılarak setup fonksiyonunun yalnızca bir kez çalıştırılması sağlanır. İki farklı goroutine'de dostuff fonksiyonu çağrılır, ancak once.Do sayesinde setup fonksiyonu yalnızca bir kez çalıştırılır.

## Deadlock

Deadlock, bir goroutine'in diğer bir goroutine'in tamamlanmasını beklediği, ancak diğer goroutine'in de beklemeye geçtiği durumdur.

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

	go func() {
		defer wg.Done()
		x := <-ch
		fmt.Println(x)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 42
	}()

	wg.Wait()
}
```

Bu örnekte, iki goroutine arasında bir channel (ch) üzerinden iletişim var. Ancak deadlock durumu oluşuyor çünkü bir goroutine kanaldan değer almayı beklerken, diğer goroutine kanala değer göndermeyi bekliyor. Bu durumda her ikisi de birbirini bekleyeceği için deadlock meydana gelir.

## Yemek Yiyen Filozoflar Problemi // Dining Philosophers Problem

Bu klasik eşzamanlılık problemi, beş filozofun bir masada oturduğu ve birbirleri arasındaki çatal (mutex) kullanımıyla yemek yemeye çalıştığı bir senaryoyu simgeler.

Örnek:

```
package main

import (
	"fmt"
	"sync"
)

type chopstick struct {
	sync.Mutex
}

type philosopher struct {
	leftCS, rightCS *chopstick
}

func (p philosopher) eat() {
	for {
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Println("Eating")
		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup

	CSticks := make([]*chopstick, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(chopstick)
	}

	philos := make([]philosopher, 5)
	for i := 0; i < 5; i++ {
		philos[i] = philosopher{CSticks[i], CSticks[(i+1)%5]}
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			philos[i].eat()
		}(i)
	}

	wg.Wait()
}
```

Bu örnekte, beş filozofun eat metodu sürekli olarak çalışır. Ancak, deadlock olasılığı vardır çünkü tüm filozoflar aynı anda sol çatalı kavrayabilirler, ardından sağ çatalı kavrayarak birbirlerini bekleyebilirler. Bu durumda, deadlock meydana gelir.

Deadlock durumlarından kaçınmak için çeşitli yöntemler kullanılabilir. İşte bazı çözüm yaklaşımları:

- **Chopstick Sıralamasını Değiştirme:**

Önceki örnekte, Dykstra'nın önerisine uyarak her filozofun çatalını kilitlediği sırayı değiştirmek bir çözüm olabilir. Her filozof, en düşük numaralı çatalı ilk kilitlerse, deadlock olasılığı azalır:

```
// Önceki kodlar...

type philosopher struct {
    leftCS, rightCS *chopstick
}

func (p philosopher) eat() {
    for {
        // Çatal sırasını değiştirme
        if p.leftCS.num < p.rightCS.num {
            p.leftCS.Lock()
            p.rightCS.Lock()
        } else {
            p.rightCS.Lock()
            p.leftCS.Lock()
        }

        fmt.Println("Eating")

        p.rightCS.Unlock()
        p.leftCS.Unlock()
    }
}

// Önceki kodların devamı...
```

- **Farklı Bir Senkronizasyon Yaklaşımı:**

Mutex ve deadlock durumlarını önlemek için başka senkronizasyon mekanizmalarını (kanallar, bekleyen gruplar vb.) düşünmek, kodunuzu deadlock'a yol açmayacak şekilde tasarlamak önemlidir.

- **Zamanlayıcı Kullanma:**

Bir filozof bir çatalı kilitledikten belirli bir süre sonra kilidi bırakabilir. Eğer diğer çatalı alamazsa, kilidi bırakarak sırayla diğer filozoflara geçebilir. Ancak bu, açlık durumuna (starvation) yol açabilir:

```
func (p philosopher) eatWithTimeout(timeout time.Duration) {
    for {
        if p.leftCS.TryLock(timeout) {
            if p.rightCS.TryLock(timeout) {
                fmt.Println("Eating")
                time.Sleep(100 * time.Millisecond) // Yemek yeme simülasyonu
                p.rightCS.Unlock()
            }
            p.leftCS.Unlock()
        }
    }
}
```

Bu çözümler, deadlock durumlarından kaçınmak için kullanılabilir. Ancak her durum benzersizdir ve tasarım, senaryo ve performans gereksinimlerine bağlı olarak farklı çözümler gerekebilir.