# Support for Classes

Go'da, bir "class" anahtar kelimesi yoktur; bunun yerine, tipler üzerinde metotların tanımlanması sağlanır. Veri tipine özgü metodlar oluşturmak için, bir tipte "receiver" olarak adlandırılan özel bir parametre tanımlanır. Örneğin:

```
type MyInt int
```

Sonra, bu MyInt türü üzerinde çalışacak bir metod tanımlanır.

```
func (mi MyInt) Double() int {
    return int(mi * 2)
}
```

main() fonksiyonunda ise, MyInt tipinden bir değişken oluşturulur ve bu değişken üzerinde Double() metodunu çağırarak işlem yapılır.

```
func main() {
    v := MyInt(3)
    fmt.Println(v.Double())
}
```

Bu kod, MyInt tipine ait bir değeri ikiyle çarparak sonucunu döndüren Double() metodunu kullanarak, v değişkeninin değerini işlemleyip çıktı olarak görüntüler.

Daha detaylı anlatmak gerekirse; 

**mi:** mi, bu örnekte bir alıcı ("receiver") olarak kullanılan isimdir. Alıcı, bir metodun hangi tür üzerinde tanımlandığını belirtir. func (mi MyInt) Double() int ifadesinde mi, MyInt türü üzerinde çalışacak olan Double() metodunun alıcısıdır. 

Peki, **Alıcı (Receiver)** Nedir?: Go dilinde bir fonksiyonun başında belirtilen ve metodun hangi tür üzerinde çalışacağını belirten bir parametredir. Bu, diğer dillerdeki "this" veya "self" kavramlarına benzerdir. Örneğin, mi MyInt ifadesi, Double() metodunun MyInt türü üzerinde kullanılacağını gösterir.

**Metod (Method)** Nedir?: Bir metod, bir nesnenin özelliklerini veya davranışlarını tanımlayan fonksiyonlardır. Go dilinde bir metodun tanımı, bir alıcı (receiver) ile başlar. Bu alıcı, belirli bir türe (struct, int, string vb.) ait fonksiyonların çalışacağı türü ifade eder.


Diyelim ki bir oyun karakteri için konumları temsil eden bir yapı oluşturmak istiyoruz. Her karakterin x ve y koordinatları var ve bu koordinatlar üzerinde bazı işlemler yapmak istiyoruz. Bunun için bir yapı (struct) ve bununla ilişkili yöntemler tanımlayabiliriz.

```
// Point struct tanımı
type Point struct {
	x float64
	y float64
}

// Point tipi için yeni bir yöntem tanımlama
func (p Point) DistanceToOrigin() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func main() {
	// Yeni bir Point örneği oluşturma
	p1 := Point{x: 3, y: 4}

	// DistanceToOrigin() metodunu çağırarak karakterin orijine uzaklığını hesaplama
	distance := p1.DistanceToOrigin()

	fmt.Printf("Karakterin orijine uzaklığı: %f\n", distance)
}
```

Yukarıdaki örnekte, Point adında bir yapı (struct) tanımladık. Bu yapı, karakterin x ve y koordinatlarını içerir.

Sonra, DistanceToOrigin() adında bir metod tanımladık. Bu metod, Point türüne ait bir alıcı (receiver) olan p üzerinde çalışır. Bu metod, p noktasının orijine olan uzaklığını hesaplamak için Pisagor teoremini kullanır.

main() fonksiyonunda, p1 adında bir Point örneği oluşturduk ve x ve y koordinatlarını (3 ve 4) belirttik. Sonra p1.DistanceToOrigin() metodunu çağırarak, karakterin orijine olan uzaklığını hesaplayıp ekrana yazdırdık.

Bu örnek, bir yapı içinde verileri (x ve y koordinatları) tutar ve bu yapıya ait bir metod ile bu veriler üzerinde bir işlem yapılmasını sağlar. Karakterin orijine olan uzaklığını hesaplamak için bu metod kullanılmıştır.