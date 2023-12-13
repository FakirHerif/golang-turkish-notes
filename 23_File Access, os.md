# File Access, os

Dosya erişimi üzerinde biraz daha hassas bir kontrole sahip olmak istiyorsak genellikle OS paketini kullanmamız gerekiyor. Bu paket dosyalara erişmek için bir dizi işlev ve daha fazla kontrole sahip olmamızı sağlar.

Yani IOutil ile tüm dosyayı okuyabilir ve yazabiliriz ancak OS ile daha fazlasını yapabiliriz.

os.File, dosyaların okunması, yazılması, kapatılması ve konumlandırılması gibi işlemleri gerçekleştirmek için kullanılır. Bir dosyayı açmak ve üzerinde işlem yapmak için os.Open() veya os.Create() gibi işlevler kullanılır ve bu işlevler bir os.File türü döndürür.

İşte os.File ile dosya işlemlerinin temel örnekleri:

```
package main

import (
	"fmt"
	"os"
)

func main() {
	// Dosya oluşturma veya açma işlemi
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Dosya oluşturma hatası:", err)
		return
	}
	defer file.Close() // Dosyayı işlem sonunda kapat

	// Dosyaya yazma işlemi
	data := []byte("Bu dosyaya yazılacak metin.")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Dosyaya yazma hatası:", err)
		return
	}

	// Dosyadan okuma işlemi
	readData := make([]byte, 100) // 100 byte'lık bir buffer oluştur
	_, err = file.Read(readData)
	if err != nil {
		fmt.Println("Dosyadan okuma hatası:", err)
		return
	}
	fmt.Println("Dosyadan okunan veri:", string(readData))
}
```

Yukarıdaki örnek kodda, os.Create() işlevi kullanılarak "example.txt" adında bir dosya oluşturulur veya varsa açılır. os.Create() işlevi, bir dosyayı oluşturur veya üzerine yazar. os.Open() işlevi ise sadece bir dosyayı açar, eğer dosya yoksa hata döner.

Dosya oluşturma veya açma işlemi sonucunda dönen os.File türünden bir değer ile dosya üzerinde işlemler gerçekleştirilir. Örneğin, file.Write() metodu dosyaya veri yazmayı sağlar, file.Read() metodu dosyadan veri okur.

defer file.Close() ifadesi, dosya işlemi tamamlandıktan sonra dosyayı kapatmak için kullanılır. Dosya işlemi bitiminde dosyanın kapatılması, kaynakların etkin bir şekilde yönetilmesini sağlar.

os.File, dosyalarla çalışmak için bir dizi işlevi içerir ve dosyaların okunması, yazılması, konumlandırılması gibi temel işlemleri gerçekleştirmek için kullanılır.