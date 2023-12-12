# Maps

Go programlama dilinde, hash tabloları veya dilimlerle eşdeğer olan map veri yapısı bulunmaktadır. Yani maps, go dilinin bir hash tablosu uygulamasıdır. Kısaca map; hash tablodur.

Map oluşturmak için make() fonksiyonu kullanabiliriz. Örneğin:

```
var idMap map[string]int        // map[string]int türünde bir değişken tanımladık
idMap = make(map[string]int)    // make() fonksiyonu ile map'i oluşturduk


[ string ]  =  key type
int         =  value type
```

var idMap map[ string]int ifadesiyle, idMap adında bir değişken tanımlanır. Bu değişken bir map veri yapısını temsil eder. string türünde anahtarlarla eşleştirilmiş int türünde değerler içerecek bir map olduğunu belirtir.

Ardından make() fonksiyonu kullanılarak idMap değişkenine boş bir map oluşturulur. make() fonksiyonu, map veri yapısını oluşturmak için kullanılır ve bu örnekte anahtarları string ve değerleri int türünde olan bir map oluşturur.

Bir başka örnek:

```
idMap := map[string]int{
    "joe": 123}
```

idMap adında yeni bir değişken oluşturuldu ve map[ string ]int türünde bir map olarak tanımlandı.

Harita oluşturulurken, süslü parantezler {} kullanılarak anahtar-değer çiftleri belirtilir. Burada anahtarlar string türünde ve değerler int türündedir.

"joe" anahtarı, 123 değeri ile eşleştirilmiştir.
Yani, bu kod, "joe" anahtarını 123 değeriyle eşleyen bir map oluşturur.

Bu yapı, map veri yapısını tanımlamak ve aynı zamanda başlangıç değerleriyle doldurmak için kısa ve pratik bir yol sunar. Bu tür yapılar, bir map veri yapısını hızlı bir şekilde başlatmak ve belirli anahtarlarla ilişkilendirilmiş değerler eklemek için oldukça kullanışlıdır.

**Accessing Maps / Map veri yapısına erişmek**

Go dilinde map veri yapısına erişmek, değer atamak ve bu yapılarda işlem yapmak oldukça kolaydır. map veri yapısına erişmek için, anahtarları kullanarak değerlere erişmek gerekmektedir. Örneğin:

```
idMap := map[string]int{
    "Alice": 25,
    "Bob":   30,
    "Eve":   28,
}
```

Bu örnekte, idMap adında bir map oluşturduk. string türündeki anahtarlarla int türündeki değerleri eşleştirdik. Şimdi bu map yapısına erişmek için kullanabileceğimiz yöntemlere bir göz atalım:

- **Değer Erişimi:**

```
fmt.Println(idMap["Alice"])     // "Alice" anahtarına karşılık gelen değeri yazdırır
fmt.Println(idMap["Bob"])       // "Bob" anahtarına karşılık gelen değeri yazdırır
```

Bu kod, idMap map yapısındaki belirli anahtarlara karşılık gelen değerleri yazdıracaktır.

- **Değer Atama:**

```
idMap["Eve"] = 29               // "Eve" anahtarına karşılık gelen değeri değiştirir
idMap["Charlie"] = 33           // Yeni bir anahtar ve değer ekler
```

Bu örnek, idMap map yapısındaki değerleri günceller veya yeni anahtar-değer çifti ekler.

- **Değer Silme:**

```
delete(idMap, "Eve")           // "Eve" anahtarını ve karşılık gelen değeri mapten siler
```

- **Anahtar Varlığını Kontrol Etme:**

```
value, exists := idMap["Bob"]
if exists {
    fmt.Println("Bob's age is", value)
} else {
    fmt.Println("Bob's age is not available")
}
```

Bu örnek, belirli bir anahtarın map yapısında var olup olmadığını kontrol eder. exists değişkeni, belirtilen anahtarın var olup olmadığını gösterir. Eğer anahtar varsa, ilgili değeri alırız; aksi takdirde, anahtarın olmadığını belirten bir mesaj yazdırırız.

map veri yapısı, anahtarlarla eşleştirilmiş değerleri hızlıca erişmek ve güncellemek için oldukça kullanışlıdır. Anahtar-değer çiftleriyle çalışırken, anahtarın map içinde olup olmadığını kontrol etmek gibi çeşitli işlemler yapmak mümkündür.

- **anahtar-değer çiftlerini döngü kullanarak gezmek**

```
idMap := map[string]int{
    "Alice": 25,
    "Bob":   30,
    "Eve":   28,
}
```

Bu kodda, idMap adında bir map oluşturduk. Şimdi bu map yapısını döngü kullanarak gezerek içerisindeki her bir anahtar-değer çiftini yazdıralım:

```
for key, val := range idMap {
    fmt.Println(key, val)
}
```

Bu döngü, idMap map yapısındaki her bir anahtar-değer çiftini key ve val adındaki değişkenlere atar ve bu çiftleri döngü boyunca kullanır. Her iterasyonda key değişkeni anahtar değerini (string tipinde) ve val değişkeni ise bu anahtara karşılık gelen değeri (int tipinde) temsil eder.

Döngü, idMap içindeki her bir anahtar-değer çiftini dolaşır ve her bir çift için key ve val değerlerini alır. Bu değerler döngü içinde kullanılarak her bir anahtarın ve ona karşılık gelen değerin ekrana yazdırılmasını sağlar. Bu sayede map yapısındaki tüm anahtar-değer çiftleri görüntülenir.