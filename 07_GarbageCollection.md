# Garbage Collection

Go içinde çöp toplama bulunan derlenmiş bir dildir.

Garbage Collection (Çöp Toplama), bir programın çalışması sırasında artık kullanılmayan bellek alanlarını otomatik olarak tespit edip serbest bırakan bir süreçtir. Garbage Collection, dilin çalışma zamanı (runtime) tarafından gerçekleştirilir ve programcının müdahalesi gerektirmez. Go dilinde, otomatik bir Garbage Collector bulunmaktadır ve programcılar genellikle bellek yönetimiyle doğrudan ilgilenmek zorunda kalmazlar.

**Pointerlar ve deallocation ile Garbage Collection arasındaki ilişki:**

Pointerlar, programın bellek kullanımını kontrol etmek ve işaret etmek için kullanılırken, deallocation bellek alanlarını elle serbest bırakmaya yönelik bir işlemdir. Garbage Collection ise kullanılmayan bellek alanlarını otomatik olarak tespit edip serbest bırakarak programcıların bellek yönetimiyle ilgilenme ihtiyacını azaltır. Pointerlar, doğru kullanıldığında ve gerektiğinde bellek tahsisi ve serbest bırakma işlemlerini Garbage Collector'e bildirerek bellek yönetimini optimize etmeye yardımcı olabilir.

Örnek:

```
func foo() *int {
    x := 1
    return &x
}

func main () {
    var y *int
    y = foo()
    fmt.Printf("%d", *y)
}
```

Bu go örneği go'da legal ama bazı dillerde legal değildir.

Bu örnekte, foo() adında bir fonksiyon ve main() adında başka bir fonksiyon bulunmaktadır. foo() fonksiyonu bir int tipinde bir pointer döndürmektedir. main() fonksiyonu ise bu foo() fonksiyonunu çağırarak bir değer alır ve ekrana bu değeri yazdırır.

- **foo() fonksiyonu:**
    - foo() fonksiyonu içinde, x adında bir int değişkeni tanımlanır ve değeri 1 olarak atanır.
    - &x kullanılarak x'in bellek adresi (pointer) döndürülür. Ancak bu adres, fonksiyonun dışında olduğu için fonksiyonun sona ermesiyle birlikte bellekten silinir.

- **main() fonksiyonu:**
    - var y *int ifadesi ile y adında bir int tipinde pointer oluşturulur.
    - y = foo() ile foo() fonksiyonu çağrılır ve döndürülen adres y'ye atanır.
    - fmt.Printf("%d", *y) ile y pointer'ının gösterdiği adresteki değer ekrana yazdırılır.

Ancak bu kodun önemli bir problemi var. foo() fonksiyonunda oluşturulan x değişkeninin adresi, fonksiyonun sona ermesiyle birlikte geçersiz hale gelir. Yani, main() fonksiyonundaki y pointer'ı geçersiz bellek adresine işaret eder. Bu, **"dangling pointer"** hatasına yol açabilir ve beklenmeyen sonuçlar doğurabilir.

**Ancak Go dilinde,** garbage collection mekanizması devreye girerek bu tür geçersiz bellek adreslerini izler ve kullanılmayan bellek alanlarını belirler. Kullanılmayan bu bellek alanlarını tanımlar ve geri dönüşüme alır, böylece programcının bellek sızıntısı gibi sorunlarla uğraşması gerekmez. Bu tür durumlar otomatik olarak tespit edilir ve bellek sızıntılarından kaçınılır.

Yani, **Go dilinde garbage collection**, bellek yönetimiyle ilgili hataların önlenmesine yardımcı olur ve programcının kodunu daha güvenli ve verimli hale getirir. Ancak bu durum, verimlilik açısından bazı performans maliyetleriyle de gelir. Bellek toplama işlemi, programın çalışma süresini etkileyebilir, ancak genellikle Go'nun bu maliyeti verimlilik ve güvenlik avantajlarıyla telafi ettiğini düşünebiliriz.