# Point Receivers, Referencing, Dereferencing

**Point Receivers (Nokta Alıcıları):**

- Go dilinde metodların alıcıları vardır. Bir metodun alıcısı, o metodun bağlı olduğu türü belirtir.
- Bir tür üzerinde tanımlanan bir metodun alıcısı, o türden bir değişken üzerinde çalışırken kullanılabilir.
- Nokta alıcıları, bir işlevin veya metodu bir tür üzerine bağlamak ve bu türdeki verileri manipüle etmek için kullanılır.
- Nokta alıcıları, bir metodun alıcısı olarak işaretçi türleri kullanarak o türün içindeki verileri doğrudan değiştirmek için kullanılabilir.

**Referans Alma (Referencing):**

- Referans alma, bir değişkenin bellekteki adresini elde etme işlemidir.
- Go dilinde, bir değişkenin adresini elde etmek için & operatörü kullanılır.

Örnek:

```
x := 10
ptr := &x   // x'in adresini ptr değişkenine atar
```

**Değer Çözme (Dereferencing):**

- Değer çözme, bir işaretçinin (pointer) işaret ettiği bellek adresindeki değere erişme işlemidir.
- Go dilinde, bir işaretçinin gösterdiği değere erişmek için * operatörü kullanılır.

Örnek:

```
x := 10
ptr := &x           // x'in adresini ptr değişkenine atar
fmt.Println(*ptr)   // ptr'nin işaret ettiği değeri (x'in değeri) yazdırır
```

**İşaretçi Alıcıları Kullanımı:**

İşaretçi alıcıları, bir metodun belirli bir türün işaretçisiyle ilişkilendirilmesini sağlar. Örneğin, bir metodun alıcısı *Point (Point işaretçisi) ise, bu metodun o türün işaretçileriyle kullanılması gerektiği anlamına gelir.

**Dereferansın Gerekli Olmaması:**

İşaretçi alıcıları kullanıldığında, yöntemler içinde işaretçilerin açıkça çıkarılmasına gerek kalmaz. Örneğin, p.x = p.x + v yerine p.x += v gibi bir ifade kullanılabilir. Bu durumda, işaretçi alıcısı kullanıldığı için Go dilinin otomatik olarak referansı kaldırması (dereferans) sağlanır.

**Referans Verme ve Dereferansın Otomatik Olarak Yapılması:**

İşaretçi alıcıları kullanıldığında, kod içinde referans verme (&) veya dereferans (*) operatörlerini kullanma ihtiyacı azalır. Bu durumda, örneğin bir metodun alıcısı *Point olarak tanımlandığında, metodun içinde işaretçinin kendisi ile çalışırken referans verme veya dereferans işlemleri yapılmadan işaretçi üzerinden işlem yapılabilir.

**İyi Programlama Uygulamaları:**

İşaretçi alıcıları kullanırken, tüm metodların işaretçi alıcılarını kullanması veya hiçbirinin kullanmaması iyi bir uygulama olarak belirtilir. Bu, kodun okunabilirliğini artırır ve kafa karışıklığını önler. Belirli bir tür veya işaretçi gerektirmeyen tüm referanslar için işaretçi alıcılarını kullanmak ise iyi bir uygulama olarak önerilir.

Bu kavramları bir örnekle birleştirelim:

```
type Point struct {
    x int
    y int
}

func (p *Point) OffsetX(v int) {
    p.x += v // Nokta alıcısı kullanılarak p.x değeri doğrudan değiştirilebilir.
}

func main() {
    // Bir Point türünden bir değişken oluşturuldu.
    p := Point{3, 4}

    fmt.Println("Başlangıçta p.x:", p.x) // Başlangıçta p.x değeri: 3

    // Nokta alıcısı kullanarak p.x değeri 5 arttırıldı.
    p.OffsetX(5)

    fmt.Println("Sonuçta p.x:", p.x) // Sonuçta p.x değeri: 8
}
```

Bu örnekte, Point türünden bir p değişkeni oluşturuldu. OffsetX adında bir metodun alıcısı olarak *Point belirtildi, bu sayede bu metodun içinde p.x değeri doğrudan değiştirilebildi. Bu nokta alıcıları sayesinde, bir türün bir değişkeni üzerinde işlemler yapılabilir ve bu işlemler türün içindeki verileri değiştirebilir. Bu işaretçi türleri ve nokta alıcıları, referans alma ve değer çözme işlemlerini doğrudan yürütmeyi sağlar.


