# Error Interface

Go dilinde hataları işlemek için sıkça kullanılan bir yöntem Error Interface'dir. Go'da birçok yerleşik paket, işlevlerinden bir şey döndüğünde ikinci bir değer olarak hata (error) döndürür. Hata kontrolü yapmak için bu hata değerini kontrol etmek iyi bir uygulama pratiğidir.

Öncelikle, Error Interface'ı kendimize bir bakalım:

```
type error interface {
    Error() string
}
```

Bu, Error Interface'ının temel bir tanımıdır. Bir tipin bu arayüzü uygulayabilmesi için Error adında bir metodunun olması gerekiyor ve bu metodun bir string döndürmesi bekleniyor.

Bir örnek üzerinden ilerleyelim:

```
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Println("Dosya açılırken bir hata oluştu:", err.Error())
        return
    }
    defer file.Close()

    // Dosya ile ilgili diğer işlemleri burada gerçekleştirin.
    fmt.Println("Dosya başarıyla açıldı.")
}
```

Bu örnekte, os.Open fonksiyonu bir dosyayı açarken iki değer döndürür: açılan dosya ve bir hata (error). Hemen ardından bu hatayı kontrol ediyoruz. Eğer hata değeri nil değilse, yani bir hata varsa, hatayı işleyip programa devam etmiyoruz. Hatayı işlemek için sıkça kullanılan bir yöntem, hatanın Error metodunu çağırarak hatanın açıklamasını elde etmek ve ekrana yazdırmaktır.

Bu, Go dilinde yaygın olarak kullanılan bir hata işleme pratiğidir. Hataların kontrol edilmesi ve uygun şekilde işlenmesi, güvenilir ve hata toleranslı yazılımların geliştirilmesinde önemlidir.