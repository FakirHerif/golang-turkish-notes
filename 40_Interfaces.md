# Interfaces / Arayüzler

Interface, Go'da kullanılan bir kavramdır ve polimorfizm elde etmemize yardımcı olur.
Yani miras almıyoruz, mirasa ihtiyacımız yok, geçersiz kılmaya ihtiyacımız yok. 

Go dilinde arayüzler, tip sistemine esneklik kazandıran ve polimorfizmi sağlayan önemli bir özelliktir. Bu bölümde özellikle, bir arayüzün nasıl tanımlanacağını, bir tipin bu arayüzü nasıl tatmin edeceğini ve arayüzlerin nasıl kullanıldığını ele alıyoruz.

**Interface Tanımlama:**

İlk olarak, bir arayüzün nasıl tanımlandığını gözden geçirelim. Arayüz, bir türün uygulamasını beklediğimiz metodların imzalarını içerir. Bu metod imzaları, adı, parametreleri, ve dönüş değerleri içerir, ancak metodun uygulamasını değil sadece imzasını belirtir.

Örneğin, Shape2D adında bir arayüz tanımlayalım:

```
type Shape2D interface {
    Area() float64
    Perimeter() float64
}
```

Bu arayüz, iki metodun (Area ve Perimeter) imzasını içerir. Her iki metod da hiçbir argüman almaz ve float64 türünde bir değer döndürür.

**Türlerin Arayüzü Karşılaması:**

Bir türün bu arayüzü karşılaması için, arayüzde belirtilen tüm metodların aynı imzalara sahip olması gerekir. Ancak bu metodların uygulanması veya geçersiz kılınması gerekmez, sadece imzalarına uygun bir şekilde tanımlanmış olmaları yeterlidir.

Örneğin, Triangle adında bir tür tanımlayalım:

```
type Triangle struct {
    // Triangle'a ait veriler
}

func (t Triangle) Area() float64 {
    // Area metodunun uygulanması
    // Üçgenin alanını hesapla
    return // hesaplanan alan değeri
}

func (t Triangle) Perimeter() float64 {
    // Perimeter metodunun uygulanması
    // Üçgenin çevresini hesapla
    return // hesaplanan çevre değeri
}
```

Burada, Triangle türü, Shape2D arayüzündeki Area ve Perimeter metodlarının imzalarına uygun şekilde uygulanmıştır. İçerisinde hangi veriler olduğu önemli değildir, yeter ki ilgili metodların imzalarını taşıyorsunuz.

**Arayüzü Karşılayan Tiplerin Otomatik Algılanması:**

Go dilinde, bir türün bir arayüzü karşıladığını ifade etmek için ek bir bildirimde bulunmaya gerek yoktur. Derleyici, türün arayüzü karşıladığını otomatik olarak algılar. Yani, yukarıdaki Triangle örneğinde olduğu gibi, açıkça "Triangle, Shape2D'yi karşılar" gibi bir ifade kullanmamıza gerek yoktur.

Bu durum, kodun daha temiz ve esnek olmasını sağlar. Çünkü bir türün hangi arayüzleri karşıladığını görmek için ilgili türün tanımına bakmamız yeterlidir.