# File Access, ioutil

Files, programlarda çok yaygın olarak kullanılır ve programlar arasında veri ticareti yapmak için kullanılır. Yani Golang'ın dosyalara erişmemize nasıl izin verdiğinden bahsedeceğiz.

Dosya erişimi, genellikle doğrusal (linear) bir süreçtir. Bu, dosyanın başından sonuna kadar sıralı olarak okunmasını veya yazılmasını gerektirir. Bu, dosyaların fiziksel olarak depolama ortamında, örneğin bir kasede, doğrusal olarak saklandığı anlamına gelir. Dosyaların bu yapıya sahip olması, başından okumak istediğinizde, okuma işleminin başlangıç noktasından başlaması ve sona ulaşmak için doğrusal olarak devam etmesi gerektiği anlamına gelir.

Dosya erişim işlemleri temelde birkaç aşamadan oluşur: dosyayı açma, okuma veya yazma işlemleri ve ardından dosyayı kapatma. Dosya işlemleri, ReadFile veya WriteFile gibi işlevler aracılığıyla gerçekleştirilebilir. Örneğin, ReadFile, bir dosyayı okuyup içeriğini bir bayt dizisine dönüştürürken, WriteFile belirli bir dosyaya veri yazmayı sağlar.

Büyük dosyaların işlenmesi durumunda, bellek sınırlamaları göz önünde bulundurulmalıdır. Özellikle ReadFile gibi işlevler, tüm dosyayı belleğe yükleyerek işlem yapar ve bu, çok büyük dosyaların bellek sınırlarını aşmasıyla sonuçlanabilir. Bu durumda, alternatif yöntemler veya farklı dosya okuma ve yazma işlevleri kullanılabilir.

Bu nedenle, dosya erişimi ve işlemleri, bir programın veri alışverişi yapmasını sağlayan önemli bir unsurdur. Doğrusal erişim yapısına sahip olmaları, dosya işlemlerinin doğru şekilde yönetilmesi ve büyük dosyaların işlenmesi durumunda dikkat edilmesi gereken önemli bir noktadır.

**ioutil.ReadFile()**

ioutil.ReadFile() fonksiyonu, Go programlama dilinde dosya okuma işlemlerini gerçekleştirmek için kullanılan bir işlevdir. Bu işlev, ioutil paketi içinde yer alır ve dosya içeriğini bir byte dizisine dönüştürür.

İşte ioutil.ReadFile() fonksiyonunun kullanımı:

```
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filePath := "example.txt"

	// Dosya okuma işlemi
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Dosya okuma hatası:", err)
		return
	}

	// Dosya içeriğini ekrana yazdırma
	fmt.Println("Dosya İçeriği:")
	fmt.Println(string(content))
}
```

Yukarıdaki örnek kodda, ioutil.ReadFile() fonksiyonu kullanılarak bir dosyanın içeriği okunur. filePath değişkeni, okunacak dosyanın yolunu belirtir. ioutil.ReadFile() fonksiyonu bu dosyayı okur ve dosya içeriğini bir byte dizisine dönüştürerek content değişkenine atar.

Eğer bir hata oluşursa, err değişkeni nil olmayacak ve hata mesajı fmt.Println("Dosya okuma hatası:", err) ifadesiyle ekrana yazdırılacaktır.

Dosya içeriğini ekrana yazdırmak için fmt.Println(string(content)) ifadesi kullanılır. content bir byte dizisi olduğu için string() dönüşümü ile bu byte dizisi stringe dönüştürülerek ekrana yazdırılır.

ioutil.ReadFile() fonksiyonu, dosyanın tüm içeriğini bir defada okur ve bellekteki bayt dizisine yükler. Bu nedenle, büyük dosyaların okunması durumunda bellek sınırlamalarını göz önünde bulundurmak önemlidir. Büyük dosyaların işlenmesi durumunda, dosya okuma işlemi için farklı yöntemler veya os paketindeki diğer fonksiyonlar tercih edilebilir.

**ioutil.WriteFile()**

ioutil.WriteFile() fonksiyonu, Go programlama dilinde dosya yazma işlemlerini gerçekleştirmek için kullanılan bir işlevdir. Bu işlev, ioutil paketi içinde yer alır ve belirtilen dosyaya veri yazmayı sağlar.

İşte ioutil.WriteFile() fonksiyonunun kullanımı:

```
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filePath := "example.txt"
	content := []byte("Bu dosyaya yazılacak metin.")

	// Dosyaya yazma işlemi
	err := ioutil.WriteFile(filePath, content, 0644)
	if err != nil {
		fmt.Println("Dosya yazma hatası:", err)
		return
	}

	fmt.Println("Dosya başarıyla yazıldı:", filePath)
}
```

Yukarıdaki örnek kodda, ioutil.WriteFile() fonksiyonu kullanılarak bir dosyaya veri yazılır. filePath değişkeni, yazılacak dosyanın yolunu belirtir. content değişkeni ise dosyaya yazılacak veriyi içerir. Bu örnekte, string bir metni byte dizisine dönüştürüp dosyaya yazıyoruz.

ioutil.WriteFile() fonksiyonu, belirtilen dosya yolundaki dosyayı oluşturur veya var olan dosyayı silip yeniden oluşturarak veriyi yazar. Eğer dosya yoksa, bu işlev otomatik olarak dosyayı oluşturur. Eğer dosya varsa ve üzerine yazma izni varsa, varolan içeriği siler ve yeni veriyi yazar. Dosyayı sadece yazılabilir durumda açabilir.

Eğer bir hata oluşursa, err değişkeni nil olmayacak ve hata mesajı fmt.Println("Dosya yazma hatası:", err) ifadesiyle ekrana yazdırılacaktır.

Bu örnekte, 0644 izinlerini belirten üçüncü bir parametre olarak verilmiştir. Bu izinler, Unix benzeri sistemlerde dosya erişim izinlerini temsil eder. 0644, dosyanın okunabilir ve yazılabilir olduğunu ve diğer kullanıcıların sadece okuyabileceğini belirtir.

Örneğin 0777 ise dosyanın herkes için izni var demektir, yani herkes herşeyi yapabilir.

ioutil.WriteFile() fonksiyonu, dosyaya belirtilen veriyi yazmak için kullanılır. Farklı izinlerle veya farklı dosya yazma işlemleriyle kullanılabilir, ancak büyük dosyaların yazılması durumunda veya işlemin performansı gerektiğinde alternatif yöntemler düşünülebilir.

ioutil.WriteFile() işlevi çağrıldığında, belirtilen dosya üzerinde bir değişiklik yapılmak istenirse, bu işlev dosyanın içeriğini tamamen değiştirir. Eğer dosya zaten varsa, WriteFile() mevcut içeriği siler ve belirtilen veriyi dosyaya yazar. Dolayısıyla, dosya işleminin sonunda dosya üzerine eklenen veriyle sınırlı kalır; yani, **dosya baştan sona yeniden yazılır.**

Eğer dosyaya **ek veri eklemek veya mevcut içeriği koruyarak yeni veri eklemek gerekiyorsa**, ioutil.WriteFile() işlevi **kullanılmaz.** Bunun yerine, mevcut dosyanın içeriğini koruyarak ek veri yazmak için farklı bir yaklaşım veya farklı işlevler (örneğin, dosyanın sonuna veri ekleyebilen os.File gibi) kullanılması gerekir. Bu tür özelliği destekleyen işlevler, mevcut dosyanın içeriğini değiştirmeden dosyaya ek veri yazmaya olanak tanır.