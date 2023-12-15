# Classes and Encapsulation

**Sınıflar ve Nesne Yönelimli Programlama (OOP)**

Nesne Yönelimli Programlama (OOP), veri ve işlevleri birleştirerek kodu daha modüler ve düzenli hale getiren bir paradigmadır. Golang, sınıf tabanlı OOP yapısına sahip olmasa da, benzer özellikleri sağlayan bir dil olarak kullanılabilir.

**Sınıflar**, veri alanlarını (alanlar, özellikler veya değişkenler olarak adlandırılabilir) ve bu alanlarda çalışan fonksiyonları (yöntemler olarak adlandırılır) içeren yapısal bir yapıdır.
Örnek olarak, 2D bir uzayda bir noktayı temsil eden bir "Point" sınıfını ele alalım. Bu sınıf, x ve y koordinatlarını içeren veri alanlarına ve bu koordinatlar üzerinde çalışan işlevlere sahip olabilir.

```
// Car struct (sınıf) tanımlaması
type Car struct {
	brand string
	year  int
}

func main() {
	// Araba sınıfından bir örnek oluşturuyoruz
	myCar := Car{
		brand: "Toyota",
		year:  2022,
	}

	// Araba örneğinin özelliklerini yazdırma
	fmt.Println("Marka:", myCar.brand)
	fmt.Println("Yıl:", myCar.year)
}
```

Bu örnek, Go dilinde bir sınıf yapısını temsil eder. Car adlı bir struct oluşturulur ve bu struct içerisinde arabaların marka ve üretim yılı gibi özellikleri yer alır.

**Nesneler**, bir sınıfın örneğidir. Bir sınıftan oluşturulan bir nesne, o sınıfın veri alanlarını içeren ve sınıfın yöntemlerini kullanabilen bir yapıdır.

```
type Car struct {
	brand string
	year  int
}

// Car struct'ına ait metodlar
func (c *Car) drive() {
	fmt.Println("Araba sürülüyor:", c.brand)
}

func main() {
	// Araba nesnesi oluşturma
	myCar := Car{
		brand: "Toyota",
		year:  2022,
	}

	// Arabayı sürme (metodu çağırma)
	myCar.drive()
}
```

Bu örnek, Go dilinde bir nesne yönelimli programlama örneğidir. Car adlı bir struct oluşturulur ve bu struct içerisinde arabaların marka ve üretim yılı gibi özellikleri yer alır. drive() adında bir metod tanımlanır ve bu metod, arabayı sürmek için kullanılır.

**Kapsülleme**, bir sınıfın iç verilerini ve iç işlevlerini dış dünyadan gizleyerek, sadece belirlenmiş yöntemlerle erişilebilmesini sağlar. Bu, verilerin tutarlılığını ve güvenliğini korumak için kullanılır.

```
// Car struct'ı oluşturuluyor
type Car struct {
	brand     string
	year      int
}

// Car struct'ının markasını değiştiren bir fonksiyon
func (c *Car) SetBrand(newBrand string) {
	c.brand = newBrand
}

// Car struct'ının yılını döndüren bir fonksiyon
func (c *Car) GetYear() int {
	return c.year
}

func main() {
	myCar := Car{
		brand: "Toyota",
		year:  2022,
	}

	// Car struct'ının içindeki değerlere erişim
	fmt.Println("Marka:", myCar.brand) // Direkt erişim
	fmt.Println("Yıl:", myCar.GetYear())

	// SetBrand fonksiyonu ile markayı değiştirme (kapsülleme)
	myCar.SetBrand("Honda") // Özel fonksiyon aracılığıyla değiştirme

	// Değiştirilmiş marka değeri
	fmt.Println("Yeni Marka:", myCar.brand) // Direkt erişim
}
```

Bu örnek, Car adında bir struct oluşturur. brand ve year adında iki özel alan içerir. SetBrand adında bir fonksiyon, arabaların markasını değiştirmek için kullanılır. GetYear adında bir fonksiyon ise arabaların yılını almak için kullanılır.

myCar adında bir arabaya sahibiz. Öncelikle, brand ve year alanlarına doğrudan erişim sağlıyoruz. Sonra SetBrand fonksiyonu aracılığıyla markayı değiştiriyoruz.

Burada kapsülleme, brand adlı özgün bir alanın doğrudan değiştirilmesini engelleyerek bir fonksiyon aracılığıyla bu değişimi sağlar. Bu, verilere sınırlı erişim sağlamak ve veri bütünlüğünü korumak için önemli bir tekniktir.

**Ders Sonu Notu: ("c *Car" ve "c Car" Farkı)**

**"c Car" :** Bu ifade, metodun alıcısının Car türünden bir değer olduğunu belirtir. Bu durumda metodun işlem yaptığı nesne değeri kopyalanır, yani aslında metodun içinde yapılan değişiklikler, orijinal Car nesnesine etki etmez.

Örnek:

```
type Car struct {
	brand string
}

func (c Car) GetBrand() string {
	return c.brand
}

```

GetBrand metodu Car türünden bir değer üzerinde çalışır. Bu durumda, metodun içindeki değişiklikler aslında orijinal Car nesnesine yansımaz. Yani, bu metodu çağırdığınızda, orijinal Car nesnesinin üzerinde bir **değişiklik yapılmaz.**

**"c *Car" :** Bu ifade, metodun alıcısının *Car türünden bir işaretçi olduğunu belirtir. Bu durumda, metodun işlem yaptığı nesne üzerinde yapılan değişiklikler, orijinal *Car işaretçisinin gösterdiği Car nesnesine etki eder.

Örnek:

```
type Car struct {
	brand string
}

func (c *Car) GetBrand() string {
	return c.brand
}
```

GetBrand metodu *Car türünden bir işaretçi üzerinde çalışır. Bu durumda, metodu çağırdığınızda, orijinal *Car işaretçisi üzerinde yapılan değişiklikler, Car nesnesine **etki eder.**