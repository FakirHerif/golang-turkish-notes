# Function Parameters and Return Values

Go'da fonksiyon parametreleri, fonksiyonun alacağı girdi değerlerini belirtmek için kullanılır. Parametreler fonksiyon adından sonra parantez içinde tanımlanır.

```
// foo fonksiyonu, iki tamsayı parametresi alır ve çarpımlarını yazdırır.
func foo(x int, y int) {
    fmt.Println(x * y)
}

// Fonksiyonu çağırma
func main() {
    foo(2, 3) // Çıktı: 6
}
```

**Parametrelerin Geçirilmesi:**

Fonksiyon çağrıldığında, o fonksiyonun parametrelerine karşılık gelen argümanlar verilir. Parametreler tanımlandıkları sırada sıralı olarak argümanlarla eşleştirilir.

```
// add fonksiyonu, iki tamsayı parametresi alır ve toplamlarını döndürür
func add(x, y int) int {
    return x + y
}

func main() {
    result := add(5, 3) // add fonksiyonunu çağırma ve sonucu result değişkenine atama
    fmt.Println("Sonuç:", result) // Çıktı: Sonuç: 8
}
```

**Parametresiz Fonksiyonlar:**

Bazı durumlarda fonksiyonun parametresi olmayabilir. Bu durumda dahi parantez içi boş bırakılmalıdır.

```
// Parametresiz fonksiyon
func bar() {
    fmt.Println("Hello")
}

// Fonksiyonu çağırma
func main() {
    bar() // Çıktı: Hello
}
```

**Birden Fazla Dönüş Değeri:**

Go dilinde bir fonksiyon birden fazla değer döndürebilir. Bu değerlerin türleri fonksiyon tanımından sonra belirtilmelidir. Fonksiyon çağrısının sağ tarafında kullanılarak dönüş değerleri alınır.

```
// foo2 fonksiyonu, bir tamsayıyı alır ve o tamsayıyı ve bir artırımını döndürür.
func foo2(x int) (int, int) {
    return x, x + 1
}

// Fonksiyonu çağırma ve dönüş değerlerini kullanma
func main() {
    a, b := foo2(3)
    fmt.Println(a, b) // Çıktı: 3 4
}
```

Bu örnekler, Go dilinde fonksiyon parametreleri ve dönüş değerlerinin kullanımını göstermektedir. Fonksiyonlar, belirli görevleri yerine getirirken parametreleri alır ve belirli sonuçları döndürebilirler, bu da programın gerekli işlevleri gerçekleştirmesini sağlar.