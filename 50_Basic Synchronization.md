# Basic Synchronization

## Senkronizasyon Nedir?

Senkronizasyon, birden çok thread veya gorutinin, bir olayın zamanlaması konusunda anlaşması anlamına gelir.

Bu olaylar genellikle global olaylardır ve tüm thread'ler aynı anda bu olayları görür, böylece zamanlamada anlaşabilirler.

## Neden Senkronizasyon Gerekli?

Go dilinde, gorutinler genellikle birbirlerinin ne zaman çalıştığını umursamaz.
    
Ancak bazen belirli bir sıraya ihtiyaç duyabiliriz. Bu durumda senkronizasyon devreye girer.

## Senkronizasyonun Dezavantajları:

Senkronizasyon kullanmak, programın performansını azaltabilir çünkü senkronizasyon, thread'lerin bağımsızlığını kısıtlayarak çalışma hızını düşürebilir.

## Bekleme Grupları (Wait Groups):
    
Bekleme grupları, bir gorutinin diğer gorutinleri beklemesini sağlar.
    
Örneğin, bir ana gorutin, belirli bir sayıda diğer gorutini bekleyebilir ve bu diğer gorutinlerin tamamlanmasını bekleyerek devam edebilir.

**Senkronizasyon Amacı:**

Bekleme grupları, genellikle bir ana gorutinin diğer gorutinlerin tamamlanmasını beklemesini sağlamak için kullanılır.
    
Bu, belirli bir sıra veya koordinasyon gerektiren durumlar için kullanışlıdır.

**Nasıl Kullanılır:**

sync paketindeki WaitGroup türü kullanılarak oluşturulur.
    
Add metodu ile beklenen gorutin sayısı belirtilir.

Done metodu, bir gorutin tamamlandığında çağrılır ve beklenen gorutin sayısını azaltır.
    
Wait metodu, beklenen tüm gorutinler tamamlandığında devam etmesi için ana gorutini bekletir.

**Örnek Senaryo:**

Ana gorutin, bir veya birden fazla diğer gorutini başlatır.
    
Her başlatılan gorutin, işini tamamladığında Done metodunu çağırarak beklenen gorutin sayısını azaltır.
    
Ana gorutin, Wait metodunu çağırarak beklenen gorutinlerin tamamlanmasını bekler. Bu, ana gorutinin beklemesi için bir sinyaldir.

**Performans ve Senkronizasyon Dengesi:**

Wait Groups kullanmak, programın daha düzenli ve kontrol edilebilir olmasını sağlar.
    
Ancak, senkronizasyon kullanımı, performansı etkileyebilir çünkü gorutinlerin bağımsız çalışmasını kısıtlar.

**ÖRNEK**

```
var wg sync.WaitGroup

func main() {
    wg.Add(1)  // Beklenen gorutin sayısını artır
    go myGoroutine()
    wg.Wait()  // Tüm beklenen gorutinlerin tamamlanmasını bekle
}

func myGoroutine() {
    defer wg.Done()  // Gorutin tamamlandığında beklenen sayıyı azalt
    // Gorutinin işlemleri burada
}
```

Bu örnek kod parçasında, myGoroutine adlı bir gorutin başlatılır ve bu gorutin, işlemlerini tamamladığında Done metodu çağrılarak beklenen gorutin sayısı azaltılır. Ana gorutin ise Wait metodu ile bu azaltılan sayının sıfır olmasını bekler. Böylece, ana gorutin, diğer gorutinin işini tamamlamasını bekler.