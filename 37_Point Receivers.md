# Point Receivers

Go programlama dilinde, bir metodu çağırmak için kullanılan receiver (alıcı), o metodu bir türe bağlar. Bu, bir fonksiyonu bir türe özgü hale getirir ve o türün örnekleri üzerinde çağrılabilir hale gelmesini sağlar. Receiver, Go'da metodun bir türle ilişkilendirilmesini sağlar.

Point receiver (nokta alıcısı) ise, Go'da belirli bir veri türünün (struct gibi) bir metodu ile ilişkilendirilmesi için kullanılan bir yapıdır. Genellikle bu, bir fonksiyonun belirli bir türe ait olması ve bu türün örnekleri üzerinde çalışabilmesi için kullanılır.

Örnek olarak, bir Non-Point Receiver ve Point Receiver yapısını birlikte inceleyelim:

**Non-Point Receiver Örneği:**

```
type Point struct {
	x float64
	y float64
}

func Distance(p1, p2 Point) float64 {
	return math.Sqrt((p2.x-p1.x)*(p2.x-p1.x) + (p2.y-p1.y)*(p2.y-p1.y))
}

func main() {
	p1 := Point{x: 0, y: 0}
	p2 := Point{x: 3, y: 4}

	distance := Distance(p1, p2)
	fmt.Println("İki nokta arasındaki mesafe:", distance)
}
```

**Point Receiver Örneği:**

```
type Point struct {
	x float64
	y float64
}

func (p Point) DistanceToPoint(q Point) float64 {
	return math.Sqrt((q.x-p.x)*(q.x-p.x) + (q.y-p.y)*(q.y-p.y))
}

func main() {
	p1 := Point{x: 0, y: 0}
	p2 := Point{x: 3, y: 4}

	distance := p1.DistanceToPoint(p2)
	fmt.Println("İki nokta arasındaki mesafe:", distance)
}
```

**Farklar:**

- **Metod İsmi:** Non-Point Receiver örneğinde Distance adında bir fonksiyon tanımlanmışken, Point Receiver örneğinde DistanceToPoint adında bir metod tanımlanmıştır.

- **Alıcı (Receiver):** Non-Point Receiver örneğinde Distance fonksiyonu, Point yapısına ait değildir ve dışarıdan çağrılırken parametre olarak Point yapısı kullanılır. Oysa Point Receiver örneğinde DistanceToPoint metodu Point yapısına ait bir metot olarak tanımlanmış ve Point türünden bir alıcı üzerinden çağrılır.

Point Receiver örneğinde metot, belirli bir türe (burada Point yapısına) özgüdür ve bu türden bir alıcı üzerinden çağrılır. Bu durum, ilgili veri yapısının metoduyla daha doğrudan bir ilişki sağlar ve kodun daha anlaşılır olmasını sağlar. Ayrıca, bu yapı, kodu daha modüler hale getirir ve daha kolay bakım yapılmasını sağlar.

Diğer taraftan Non-Point Receiver örneğinde ise, metot bir türe ait değil ve dışarıdan çağrıldığında türlerle daha az bağlantılıdır. Bu durum, kodu daha dağılmış hale getirebilir ve bakımı daha zor hale getirebilir.

Point Receiver kullanımı, Go'da daha yaygın olarak tercih edilen bir yapıdır çünkü metotlar daha spesifik bir türe bağlıdır ve bu türe daha güçlü bir şekilde entegre olur. Bu da yazılım geliştirmede daha iyi bir yapılanma ve daha düzenli kod anlamına gelir.