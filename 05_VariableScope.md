# Variable Scope

Bir değişkenin kapsamı, kodda bir değişkene erişilebilen yerlerdir. 

Örnek 1:

```
var x = 4

func f() {
    fmt.Printf("%d", x)
}
func g() {
    fmt.Printf("%d", x)
}
```

Birinci örnekte x fonksiyonların dışında, paket seviyesinde tanımlanmış bir değişkendir. Dolayısıyla f ve g fonksiyonları bu değişkene erişebilirler ve bu değişkenin değerini kullanabilirler. x'in kapsamı paket seviyesinde olduğu için f ve g içinde erişilebilir ve kullanılabilir.

Örnek 2:

```
func f() {
    var x = 4
    fmt.Printf("%d", x)
}
func g() {
    fmt.Printf("%d", x)
}
```

İkinci örnekte ise x, sadece f fonksiyonu içinde tanımlıdır. Bu durumda x'in kapsamı sadece f fonksiyonunun içindedir. g fonksiyonu, x değişkenine erişemez çünkü x'in kapsamı sadece f fonksiyonu ile sınırlıdır. Bu yüzden g fonksiyonunda x değişkenine erişilemez ve bir hata oluşur.

**NOT:** Go'da değişken kapsam belirleme bloklar kullanılarak yapılır. Yani Go dilinde kapsam, bir değişkenin tanımlandığı blok (fonksiyon, ifade veya paket) ile sınırlıdır.

# Blok hiyerarşisi

- **Paket Seviyesi Bloklar:** Bu bloklar, bir dosya içinde yer alan ve paket seviyesinde tanımlanan değişkenler için geçerlidir. Bu bloklar dosya içindeki tüm fonksiyonlar dışında kalan kısımları içerir.

- **Fonksiyon Blokları:** Her bir fonksiyonun kendi içinde bir bloğu vardır. Fonksiyon içinde tanımlanan değişkenler, sadece o fonksiyon içinde erişilebilir ve kullanılabilir.

- **Kontrol Yapısı Blokları:** Örneğin, if, for, switch gibi kontrol yapıları, içinde blok içerebilirler. Bu bloklar içinde tanımlanan değişkenler, yalnızca o kontrol yapısı içinde kullanılabilir ve o blok dışında erişilemez.

Bu hiyerarşiye göre, bir blok içinde tanımlanan değişkenler, sadece o blok içinde erişilebilirler. Dışarıdaki bloklar, içteki blokların içeriklerine erişebilir ancak içteki bloklar, dışarıdaki blokların içeriklerine erişemezler. Bu, değişken kapsamı (variable scope) ve blok hiyerarşisi ile ilgili önemli bir kavramdır.