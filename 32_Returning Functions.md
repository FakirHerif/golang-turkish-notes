# Returning Functions

Bir fonksiyon, başka bir fonksiyonu dönüş değeri olarak döndürebilir. Bu, özellikle belirli koşullara veya parametrelere göre özelleştirilmiş fonksiyonlar oluşturmak için kullanılabilir.

Örnek olarak, bir mesafenin orijine olan uzaklığını hesaplayan bir fonksiyon düşünelim. Ancak bu mesafenin, farklı başlangıç noktalarına göre değişebilmesini istiyoruz. Örneğin, farklı koordinatlarda farklı orijinlere sahip bir fonksiyon oluşturmak isteyebiliriz. Bu durumda, kapanışlar (closures) devreye girer.

**Kapanışlar (Closures)**

Kapanışlar, bir fonksiyonu ve o fonksiyonun tanımlandığı ortamı bir araya getirir. Yani, bir fonksiyon, kendi içinde tanımlanmış değişkenlere ve ortamına erişebilir.

Örneğin, orijin koordinatlarını (o_x ve o_y) içeren bir fonksiyon tanımlayalım ve bu fonksiyonun mesafe hesaplamak için kullanılacağını düşünelim. Bu fonksiyon döndürüldüğünde, o_x ve o_y değerleri fonksiyonla birlikte taşınır ve fonksiyon çağrıldığında kullanılabilir hale gelir.

```
func MakeDistOrigin(o_x, o_y float64) func(x, y float64) float64 {
    return func(x, y float64) float64 {
        return math.Sqrt(math.Pow(x-o_x, 2) + math.Pow(y-o_y, 2))
    }
}

func main() {
    // MakeDistOrigin kullanarak özelleştirilmiş fonksiyonları oluşturuyoruz
    Dist1 := MakeDistOrigin(0, 0) // Orijin için mesafe hesaplayan fonksiyon
    Dist2 := MakeDistOrigin(2, 2) // (2, 2) koordinatına göre mesafe hesaplayan fonksiyon

    // Oluşturulan fonksiyonları kullanarak mesafeleri hesaplayıp yazdırıyoruz
    fmt.Println(Dist1(2, 2)) // (2, 2) koordinatından orijine olan mesafe
    fmt.Println(Dist2(2, 2)) // (2, 2) koordinatından (2, 2) koordinatına olan mesafe
}
```

Yukarıdaki örnekte, MakeDistOrigin fonksiyonu, belirli bir orijine göre mesafe hesaplamak için kullanılabilen fonksiyonları oluşturur. Bu fonksiyonların, koordinatlar verildiğinde belirtilen orijine olan mesafeyi hesapladığı görülmektedir. Bu, fonksiyonların dönüş değeri olarak fonksiyonlar döndürerek, özelleştirilmiş fonksiyonlar oluşturmayı sağlar.

Kapanışlar (closures) sayesinde, Dist1 ve Dist2 fonksiyonları oluşturulurken tanımlanan orijin koordinatları, fonksiyonun çağrılması sırasında kullanılabilmektedir. Bu, fonksiyonun tanımlandığı ortamın, fonksiyon dışında çağrıldığında bile hala geçerli olmasını sağlar.