## Paket Oluşturma

Go dilinde bir paket oluşturmak için, aynı dizindeki dosyaları bir araya getirip bir paket adı ile başlamalısınız. Örneğin, matematik adında bir paket oluşturmak istiyorsanız, şu şekilde bir dosya yapısı kullanabiliriz:

```
matematik/
|-- toplama.go
|-- carpma.go
|-- main.go
```

**toplama.go dosyasında:**

```
package matematik

func Topla(a, b int) int {
    return a + b
}
```

**carpma.go dosyasında:**

```
package matematik

func Carp(a, b int) int {
    return a * b
}
```

## Paketi Kullanma

**main.go dosyasında bu paketi kullanmak için:**

```
package main

import (
    "fmt"
    "matematik"
)

func main() {
    sonucToplama := matematik.Topla(3, 4)
    sonucCarpma := matematik.Carp(2, 5)

    fmt.Println("Toplama Sonucu:", sonucToplama)
    fmt.Println("Çarpma Sonucu:", sonucCarpma)
}
```

## Çağırma ve Kullanma Avantajları

**Modülerlik:** Paketler, kodunuzu modüler hale getirir ve farklı parçalardaki işlevselliği düzenlemenize olanak tanır.

**Tekrar Kullanılabilirlik:** Bir paket, farklı projelerde veya aynı projenin farklı bölümlerinde kolayca tekrar kullanılabilir.

**Bakım Kolaylığı:** Kodunuzu küçük, bağımsız parçalara bölmek, bakımı daha kolay hale getirir. Her bir paket kendi sorumluluk alanına odaklanır.

**Okunabilirlik:** İyi adlandırılmış ve mantıklı bir paket yapısı, kodunuzun okunabilirliğini artırır.

## Dikkat Edilmesi Gerekenler

Paket isimleri genellikle küçük harfle başlamalıdır ve mümkünse kısa ve anlamlı olmalıdır.

Fonksiyon ve değişken adları da genellikle büyük harfle başlamalıdır, çünkü paket dışındaki kodlara ihraç edilen (exported) öğeler büyük harfle başlamalıdır.

Paketler genellikle bir dizin içinde gruplanır, ancak bu bir zorunluluk değildir.