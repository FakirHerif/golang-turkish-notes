# Deallocating Memory

Bellek yönetimi, bir program çalışırken kullanılan bellek kaynaklarını tahsis etme ve serbest bırakma işlemini içerir. Bellek tahsisi, programın çalıştığı süre boyunca gereken bellek miktarını tahsis etmek ve işletim sistemine iade etmek anlamına gelir. Serbest bırakma ise artık ihtiyaç duyulmayan bellek alanlarını geri verme işlemidir.

**"Deallocating memory"** veya bellek serbest bırakma, bir programın çalışması sırasında kullanılan ancak artık gerekli olmayan bellek alanlarını serbest bırakma işlemidir. Bu, programın bellek kullanımını etkili bir şekilde yönetmesine yardımcı olur. Bellek yönetiminin iyi bir şekilde yapılması, sistem kaynaklarının verimli kullanımını sağlar ve gereksiz bellek tüketimini önler.

**Yanlış bellek yönetimi**, bellek sızıntılarına **(memory leaks)** veya kullanılmayan bellek bloklarına yönelik **gereksiz tahsislere** neden olabilir. Bu durumlar, programın performansını olumsuz etkileyebilir ve hatta programın çökmesine sebep olabilir. Bu yüzden, bellek yönetimi programlama sürecinde önemli bir rol oynar.

# Stack vs. Heap

Go dilinde de diğer birçok programlama dilinde olduğu gibi bellek yönetimi için iki önemli kavram olan "stack" ve "heap" bulunmaktadır.

**Stack (Yığın):** Stack, genellikle yerel değişkenler ve fonksiyon çağrıları gibi sıralı ve yapısal olarak organize edilmiş verilerin saklandığı bellek bölgesidir. Stack belleği, son giren ilk çıkar (Last In First Out - LIFO) prensibine göre çalışır. Her bir fonksiyon çağrısı, çağrıldığı sırada bir çağrı çerçevesi (call frame) oluşturur ve bu çerçeve, fonksiyonun parametreleri, yerel değişkenleri ve diğer çalışma bilgilerini içerir. Fonksiyon çağrısı tamamlandığında, bu çerçeve stack'ten kaldırılır ve bellek otomatik olarak serbest bırakılır.

Örnek:

```
func f() {
    var x = 4               // STACK
    fmt.Printf("%d", x)
}
func g() {
    fmt.Printf("%d", x)
}
```

**Heap (Öbek):** Heap, dinamik olarak tahsis edilen bellek ve genellikle daha büyük ve yapısı daha karmaşık verilerin depolandığı bellek bölgesidir. Programcı tarafından doğrudan kontrol edilmez ve referanslar aracılığıyla erişilir. Heap belleği, bellek tahsisi yapıldığında ve tahsis edilen bellek alanı kullanılmadığında serbest bırakılana kadar kalır. Go dilinde, make() ve new() gibi özel fonksiyonlar kullanılarak heap belleğinde bellek tahsis edilebilir.

Örnek:

```
var x = 4                   // HEAP

func f() {
    fmt.Printf("%d", x)
}
func g() {
    fmt.Printf("%d", x)
}
```