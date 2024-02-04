# Threads in Go

## Karşılıklı Hariç Tutma (Mutual Exclusion)

Önceki konularda kanallardan bahsettik. Gorutinler veri alışverişi için kanallar aracılığıyla iletişim kuruyorlar. Ancak, gorutinler arasında değişken paylaşımı sorunlara yol açabilir. Bu durumu size bir örnek göstereceğim. İki gorutin, aynı paylaşılan değişkene yazmaya çalışırlarsa, biri bir sayıyı yazmaya çalışacak, diğeri diğer sayıyı yazmaya çalışacak ve birbirlerini etkileyecekler.

Bir program, diğer gorutinlere müdahale etmeden başka gorutinlerle aynı anda çağrılabiliyorsa, bu program eş zamanlı (concurrency) olarak güvenli olarak adlandırılır. Yani, bir gorutinin diğer gorutinlerin değişkenlerini karıştırmaması ve bu değişkenleri bozmamasını sağlamalısınız.

Fonksiyonlar ise bu güvensiz yollarda birbirini etkilememesi durumunda eş zamanlı olarak güvenlidir. Örneğin, bir gorutin bir değişkeni kullanırken diğer gorutin o değişkeni yazmamalıdır. Bu şekilde, bir gorutin diğer gorutinlere müdahale edemez.

Örneğin, iki gorutin arasında paylaşılan bir değişkenin kötü paylaşımına dair bir örnek veriyorum ve bunun nasıl sorunlara yol açabileceğini anlatıyorum. İki gorutin, bir değişkeni artırmakla ilgili bir örneği inceliyoruz. Ancak, bu durumda sorunlar ortaya çıkabilir çünkü eş zamanlı bir şekilde paylaşılan değişkeni yazmaya çalışıyorlar.

Daha somut bir örnek için şu kod yapısını kullanıyorum:

```
var i int
var wg sync.WaitGroup

func inc() {
    i = i + 1
    wg.Done()
}

func main() {
    wg.Add(2)
    go inc()
    go inc()
    wg.Wait()
    fmt.Println(i)
}
```

Bu örnekte, iki gorutin de "i" değişkenini artırmaya çalışıyor. Ancak, bu durumda beklenen sonuç olan "2"yi elde etmek garanti değil. İki gorutin arasındaki sıralama ve etkileşim sorunlara neden olabilir.

## Mutex

Bu problemleri çözmek için, veriyi doğru bir şekilde paylaşmak gereklidir. İki gorutinin aynı anda bir paylaşılan değişkene yazmasına izin vermemek önemlidir. Bu durum, karşılıklı hariç tutma (mutex) ile sağlanır.

Mutex, iki gorutinin bir paylaşılan değişkeni aynı anda yazmalarını önlemek için kullanılır. Bir gorutin bir paylaşılan değişkeni kullanmaya başlamadan önce, ilgili kodu "kilit" ile kilitler (lock). Bu, diğer gorutinlerin aynı değişkeni kullanmalarını engeller. Kullanım bittikten sonra, kilit açılır (unlock) ve diğer gorutinler paylaşılan değişkeni kullanabilir hale gelir.

Bu işlemi sağlamak için sync paketindeki Mutex yapısı kullanılır. Bu yapı, iki metod içerir: Lock ve Unlock. Lock metodu, bir gorutinin kilit almasını sağlar, eğer başka bir gorutin zaten kilidi almışsa, bekler. Unlock metodu ise kilidi serbest bırakır ve başka bir gorutin kilit alabilir hale gelir.

Önceki örneği düzeltmek için, her iki gorutin içinde de Lock ve Unlock kullanılarak mutex mekanizması uygulanabilir:

```
var i int
var wg sync.WaitGroup
var mu sync.Mutex

func inc() {
    mu.Lock()
    i = i + 1
    mu.Unlock()
    wg.Done()
}

func main() {
    wg.Add(2)
    go inc()
    go inc()
    wg.Wait()
    fmt.Println(i)
}
```

Bu şekilde, her iki gorutin de i değişkenini güvenli bir şekilde artırabilir ve birbirlerini etkilemez.