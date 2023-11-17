# Pointers

Go dilinde, bir değişkenin bellek adresini işaret eden değerlere "pointers" denir. Pointerlar, değerlerin yerine bellek adresini saklar ve bu adresler üzerinden değerlere erişim sağlarlar. Go dilinde pointerlar, diğer birçok dilde olduğu gibi, daha etkili ve performanslı kod yazmak için kullanılır.

Örnek:

```
var x int = 1       // x adında bir integer (tam sayı) değişkeni oluşturuldu ve 1 değeri atandı
var y int           // y adında bir integer değişkeni tanımlandı (varsayılan olarak sıfır)
var ip *int         // ip, integer türünde bir pointer değişkeni oluşturuldu

ip = &x             // ip, x değişkeninin bellek adresini gösteriyor
y = *ip             // *ip ifadesi, ip'nin gösterdiği bellek adresindeki değeri y değişkenine atıyor. Yani y artık 1 değerine sahip.
```

**new()**

new() fonksiyonu, dinamik olarak bir bellek bloğu oluşturur ve bu bloğun adresini (bir pointer olarak) döndürür. Bu işlev genellikle değer döndüren ve yeni oluşturulan bellek bloğunun adresine işaret eden bir pointer sağlayarak bir nesnenin bellek içinde nasıl oluşturulacağını gösterir.

new() fonksiyonu, temel veri türleri için bellek tahsis etmek amacıyla kullanılır. Örneğin, bir integer, string veya başka bir temel veri türü için yeni bir bellek bloğu oluşturmak için kullanılabilir.

Örnek
```
var ptr *int

ptr = new(int)      // new() ile bir integer için bellek tahsis edildi ve ptr bu bellek adresini gösteriyor

*ptr = 10           // Bellek adresi üzerinden değer atama

fmt.Println(*ptr)   // Çıktı: 10
```