# Slices

Go dilinde slices, belirli bir veri türünden oluşan dinamik boyutlu bir veri yapısıdır. Slices (dilimler), dizilerin bir görüntüsü olarak düşünülebilir; ancak, dizilere göre daha esnek ve dinamik bir yapıya sahiptirler.

Slice'lar, dizilere benzer ancak sabit bir boyuta sahip değillerdir. Dinamik boyutları olduğu için, elemanları eklemek, çıkarmak veya değiştirmek mümkündür. Slice'lar, altta yatan bir dizinin bir kesitini veya parçasını temsil ederler.

Go'da bir slice tanımlamak için şu sözdizimi kullanılır:

```
var slice []T
```

Burada T, slice'ın içinde bulunacak elemanın türünü temsil eder. []T ifadesi, bir slice olduğunu belirtir ve T veri türünde elemanlar içerir.

Ayrıca, bir slice oluşturmak için make() fonksiyonu kullanılabilir:

```
slice := make([]T, length, capacity)
```

Bu kullanım, belirli bir uzunluk (length) ve kapasite (capacity) ile bir slice oluşturur. Uzunluk, şu anda slice'ta bulunan eleman sayısını ifade ederken, kapasite slice'ın içindeki elemanların maksimum sayısını temsil eder.

Slice'lar append(), copy(), len(), cap() gibi işlevlerle yönetilebilir:

- **append():** Yeni elemanları slice'a eklemek için kullanılır.
- **copy():** Bir slice'ın içeriğini başka bir slice'a kopyalamak için kullanılır.
- **len():** Bir slice'ın uzunluğunu (içindeki mevcut eleman sayısını) verir.
- **cap():** Bir slice'ın kapasitesini (bellekteki toplam eleman sayısını) verir.

Örnek

```
// Elemanları tamsayı olan bir slice oluşturma
slice := []int{1, 2, 3, 4, 5}

// make() fonksiyonu ile boş bir slice oluşturma
emptySlice := make([]string, 0)

// append() fonksiyonu ile eleman ekleme
slice = append(slice, 6, 7, 8)
```

**Diziden dilimler (slice'lar) oluşturulması**

Örnek:

```
arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}
s1 := arr[1:3]
s2 := arr[2:5]
```

Bu örnekte ilk satır, 7 elemanlı ve string türünde bir dizi oluşturuyor. arr adı altında dizinin her bir elemanı, sırasıyla "a", "b", "c", "d", "e", "f", "g" değerlerine sahip olacak şekilde tanımlanmıştır.

İkinci satır, arr dizisinin 1. indisinden (dahil) başlayarak 3. indise (hariç) kadar olan dilimi oluşturur. Yani, s1 dizisi, arr dizisinin "b" (1. indis) ve "c" (2. indis) elemanlarını içerir. Not olarak, dilimlerde başlangıç indis dahilken, bitiş indis hariç tutulur.

Üçüncü satır, arr dizisinin 2. indisinden (dahil) başlayarak 5. indise (hariç) kadar olan dilimi oluşturur. Yani, s2 dizisi, arr dizisinin "c" (2. indis), "d" (3. indis) ve "e" (4. indis) elemanlarını içerir.

Özetle:

- **arr** adlı dizi, 7 elemana sahip ve "a"dan "g"ye kadar olan string değerlerini içerir.
- **s1, arr** dizisinin 1. indisinden (dahil) başlayarak 3. indise (hariç) kadar olan elemanları içeren bir dilimi temsil eder ("b" ve "c").
- **s2, arr** dizisinin 2. indisinden (dahil) başlayarak 5. indise (hariç) kadar olan elemanları içeren bir dilimi temsil eder ("c", "d" ve "e").

Bu dilimler (slices), verilerin belirli aralıklarını seçmek için kullanılır ve orijinal dizinin bir "görüntüsünü" temsil eder. Dilimler, veri üzerinde işlem yapmayı ve alt kümelerdeki verileri yönetmeyi kolaylaştırır.

**slice literals (dilim ifadeleri) // Ayrıca Slice Oluşturmak İçin [] kullanılır**

Go dilinde slice literals (dilim ifadeleri), yeni bir slice oluşturmak için kullanılan kısa bir ifade yöntemidir. Slice literals, var olan bir diziden veya başka bir slice'tan yeni bir slice oluşturmak için kullanılır.

Genel olarak slice literals, []T formundadır. Burada T, slice içindeki elemanların türünü temsil eder. Örneğin:

```
slice := []int{1, 2, 3, 4, 5}
```

Bu örnekte, {1, 2, 3, 4, 5} içindeki değerler, yeni bir int türündeki slice'ı temsil eder. Süslü parantez içindeki değerler, oluşturulan slice'ın elemanlarını ifade eder. Köşeli parantezin içinin boş olması slice olduğunu gösterir.