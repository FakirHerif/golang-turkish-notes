# Encapsulation / Kapsülleme

Go kapsülleme ve özel verilerin tutulması için birçok farklı destek sağlar. Ancak verilere kontrollü bir şekilde erişim olsun istiyorsunuz. Yani tüm verileri gizlemeden sadece bazı verileri gizleyerek kontrollü bir erişim olmasını isteyebilirsiniz. Bu durumda kapsülleme kullanabilirsiniz.

Kapsülleme, nesne yönelimli programlamada bir programın içsel mantığını dışarıdan gelen etkilerden koruma ve sadece belirli yöntemlerle erişilebilir olmasını sağlama sürecidir. Bu, bir nesnenin (yapının veya sınıfın) içindeki verilerin doğrudan dışarıdan değiştirilmesini önleyerek verilerin güvenliğini ve bütünlüğünü korur. Verilerin ve işlevlerin (metodların) bir araya getirilerek bir bütün oluşturulması ve dışarıdan sadece belirli işlevlerle bu bileşenlere erişilmesi esas alınır.

Kapsülleme, aynı zamanda bilgi gizleme prensibini de içerir. Bu, bir nesnenin iç yapısının dışarıya kapalı olması ve dışarıdaki kullanıcıların sadece belirli fonksiyonlar aracılığıyla bu yapının işlevselliğine erişmesini sağlar. Bu sayede nesnenin iç yapısı değişse bile dışarıdaki kullanımı etkilemez, yalnızca erişim metodlarına müdahale edilmesi gerekir.

Örnek olarak, bir banka hesabı düşünelim. Hesap sahibinin adı, hesap numarası, bakiyesi gibi bilgiler bir yapı içinde tutulabilir ve bu bilgilere sadece belirli metodlar (örneğin, para çekme, para yatırma) aracılığıyla erişilebilir. Doğrudan hesap bakiyesini değiştirmek yerine, bu işlemleri gerçekleştiren metodlar aracılığıyla bakiye güvenli bir şekilde yönetilebilir.

Örneğin;

```
// Person yapı tipi
type Person struct {
	name   string
	age    int
	gender string
}

// SetName, Person yapısının adını değiştiren metod
func (p *Person) SetName(newName string) {
	p.name = newName
}

// GetName, Person yapısının adını döndüren metod
func (p Person) GetName() string {
	return p.name
}

func main() {
	// Person yapısı üzerinden örnek oluşturma
	var person Person
	person.name = "Alice"
	person.age = 30
	person.gender = "Female"

	// GetName metodu ile adı alma
	name := person.GetName()
	fmt.Println("Kişinin adı:", name)

	// SetName metodu ile adı değiştirme
	person.SetName("Bob")

	// Yeni adı alma
	newName := person.GetName()
	fmt.Println("Yeni ad:", newName)
}
```

Bu örnekte, Person yapısında name, age ve gender alanları bulunuyor. Ancak SetName ve GetName adında iki yöntem tanımlanmış. SetName metodunu kullanarak name alanını değiştirebiliriz ve GetName metodunu kullanarak bu değeri alabiliriz.

Bu sayede Person yapısının name alanına doğrudan erişilmiyor, ancak belirli metodlar aracılığıyla kontrol ediliyor. Bu da kapsüllemenin bir örneğidir çünkü verilerin dışarıdan korunması ve sadece belirli yöntemlerle erişilebilir olması sağlanmış olur.