# Constants (Sabit)

Derleme zamanında değeri bilinen bir ifadedir. Yani değişmeyen değerdir. Program çalıştığı sürece bu değeri her zaman sabit bir değer olarak tutar ve değiştirilemez.

```
const x = 1.3
const (
    y = 4
    z = "Hi"
)
```

**IOTA**

Go programlama dilinde kullanılan özel bir sabit değerdir ve özellikle const bloklarında kullanılır. iota, genellikle artan sayısal sabitler oluşturmak için kullanılır ve bir tamsayı değeri olarak başlar, ardından her yeni sabit tanımlamasında bir artış gösterir.

iota özelliği, sabit değerlerde kullanıldığında, her yeni const tanımlamasında artan bir sayı değeri olarak çalışır. Özellikle gruplandırılmış sabit değerlerde kullanıldığında oldukça kullanışlıdır. 

Örneğin, haftanın yedi günü vardır ve bu bir değişmezdir. Haftanın her günü için bir sabit tanımlamak istediğimizde ancak farklı olan sabitler olmasını istiyorsak ve özellikle sabitin değerinin ne olduğunu umarsamıyorsak bu durum için iota özelliğini kullanırız.

Kısaca, sabitlerin farklı olması gerekir ama sabitlerin gerçek değeri önemli değildir. Yani örneğin pazartesi ile salı farklı sabitlerse ve değerlerinin 500, 5000 veya herhangi birşey olması önemsizse iota'yı kullanacağımız durum tam olarak budur.

Farklı sabitler olduğu sürece iota kullanabiliriz.

```
const (
	Pazartesi = iota // 0
	Sali             // 1
	Carsamba         // 2
	Persembe         // 3
	Cuma             // 4
	Cumartesi        // 5
	Pazar            // 6
)

func main() {
	fmt.Println("Pazartesi:", Pazartesi)
	fmt.Println("Sali:", Sali)
	fmt.Println("Carsamba:", Carsamba)
	fmt.Println("Persembe:", Persembe)
	fmt.Println("Cuma:", Cuma)
	fmt.Println("Cumartesi:", Cumartesi)
	fmt.Println("Pazar:", Pazar)
}
```

Bu örnekte, iota ile başlayan const bloğu içindeki sabitler, haftanın günlerini temsil eder. iota her yeni sabitte bir artış gösterdiği için her sabit, bir öncekinin bir fazlası olarak artan değerler alır.

Çıktı, her bir sabitin haftanın hangi gününü temsil ettiğini gösterir. Bu şekilde iota, ardışık olarak artan sabit değerlerini otomatik olarak belirler.