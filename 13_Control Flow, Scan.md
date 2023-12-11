# Control Flow, Scan

Go programlama dilinde **"scan tagless switch"** veya **"tagless switch"** olarak adlandırılan bir kontrol akışı yapısı bulunmaktadır. Bu yapının adı **"tagless"** olarak adlandırılmasının sebebi, switch ifadesinin her bir durumu değerlendirmek için spesifik bir etiket veya değer kullanmamasıdır.

Bu yapı, Go'daki switch ifadesinin daha esnek bir kullanımını sağlar. Burada switch ifadesinin kullanımı, case'lerde bir değer veya etiket yerine bir mantıksal ifade içerebilir. Bu sayede, case'lerin kontrolü switch ifadesi içindeki ifadelerin sonuçlarına göre gerçekleşir.

Örnek olarak, bir koşulun karşılanıp karşılanmadığını kontrol eden ve bunu switch ifadesi ile yapan bir örnek verelim:

```
    num := 10

	switch {
	case num < 0:
		fmt.Println("Negatif sayı")
	case num == 0:
		fmt.Println("Sıfır")
	case num > 0:
		fmt.Println("Pozitif sayı")
	}
```

Bu örnekte, switch ifadesinde herhangi bir etiket veya değer belirtilmemiş; bunun yerine case'lerde num değişkeninin hangi koşulu karşıladığı kontrol edilmiştir. switch ifadesi içindeki durumlar değerlendirilir ve karşılanan ilk durumun kod bloğu çalıştırılır.

Bu tarzda yapılan switch ifadeleri, belirli bir değişkenin değerine göre değil, koşulların mantıksal ifadelerine göre çalıştırılmasını sağlar. Bu da kodu daha esnek hale getirir ve belirli durumlar için daha genel ifadeler kullanma olanağı sağlar.

**Break ve Continue**

Go programlama dilinde break ve continue ifadeleri kontrol akışını değiştirmek için kullanılan yapısal ifadelerdir.

**Break:** break ifadesi, döngülerde veya switch ifadelerinde kullanılarak, döngüyü veya switch bloğunu anında sonlandırmak için kullanılır. Bir döngü içinde break ifadesi kullanıldığında, döngü anında sonlanır ve döngüden sonraki herhangi bir kod bloğu çalıştırılmaz. Örnek:

```
for i := 1; i <= 5; i++ {
    if i == 3 {
        break
    }
    fmt.Println(i)      // ÇIKTI: 1, 2
}
```

Bu örnekte, döngü i değeri 3 olduğunda break ifadesiyle sonlandırılır ve döngüden sonra gelen herhangi bir kod çalıştırılmaz.

**Continue:** continue ifadesi, döngülerde kullanıldığında, döngünün devam eden kısmını atlayarak döngünün bir sonraki adımına geçiş yapar. Yani, continue ifadesi kullanıldığında döngü içindeki geri kalan kodlar çalıştırılmaz ve döngü bir sonraki adımına geçer. Örnek:

```
for i := 1; i <= 5; i++ {
    if i == 3 {
        continue
    }
    fmt.Println(i)      // ÇIKTI: 1, 2, 4, 5 
}
```

Bu örnekte, i değeri 3 olduğunda continue ifadesiyle döngünün geri kalan kısmı atlanır ve döngü bir sonraki adıma geçer.

**break** ifadesi, bir döngüyü veya switch ifadesini tamamen sonlandırırken, **continue** ifadesi, bir döngü içinde sadece o adımdaki çalışmayı durdurur ve bir sonraki adıma geçer. Bu ifadeler, kontrol akışını değiştirerek belirli durumlarda kodu daha esnek hale getirmek için kullanılır.

**Scan**

Go programlama dilinde scan, genellikle konsoldan kullanıcıdan veri almak için kullanılan bir yöntemdir. fmt paketinin Scan ve Scanln fonksiyonları, kullanıcıdan girdi almak için kullanılır.

**fmt.Scan():** Kullanıcıdan girdi alırken, bir satırın tamamını okur. Boşluklarla ayrılmış verileri tek bir değişkene atar. Yani, boşlukla ayrılmış birden fazla giriş alır ve her bir girişi sırasıyla değişkenlere atar.

```
package main

import "fmt"

func main() {
    var name string
    var age int

    fmt.Print("Adınız: ")
    fmt.Scan(&name)

    fmt.Print("Yaşınız: ")
    fmt.Scan(&age)

    fmt.Println("Ad:", name)
    fmt.Println("Yaş:", age)
}
```

Bu örnekte, fmt.Scan() fonksiyonu kullanılarak sırasıyla kullanıcıdan Adınız ve Yaşınız bilgileri istenir. fmt.Scan() ile birlikte gelen değişkenlere kullanıcının girişi atılır.

Örneğin, kullanıcının girişi şu şekilde olabilir:

```
Adınız: Alice
Yaşınız: 25
Ad: Alice
Yaş: 25
```

**fmt.Scanln():** Kullanıcıdan girdi alırken, bir satırın tamamını okur. Boşluklarla ayrılmış verileri ayrı ayrı değişkenlere atar. Yani bir satırın tamamını okur. Enter tuşuna kadar olan tüm girişi bir satır olarak kabul eder. Bu nedenle, Scanln() bir satırın tamamını alır ve sonunda Enter tuşu girilene kadar bekler.

```
package main

import "fmt"

func main() {
    var name string
    var age int

    fmt.Print("Adınız: ")
    fmt.Scanln(&name)

    fmt.Print("Yaşınız: ")
    fmt.Scanln(&age)

    fmt.Println("Ad:", name)
    fmt.Println("Yaş:", age)
}
```

Bu örnekte, kullanıcıdan Adınız ve Yaşınız bilgilerini girmesi beklenir. fmt.Scanln() ifadesi ile kullanıcıdan girdiler okunur ve & operatörü kullanılarak bu girdiler ilgili değişkenlere atanır.

Örneğin, ekrana şu şekilde bir çıktı gelebilir:

```
Adınız: John Doe
Yaşınız: 30
Ad: John Doe
Yaş: 30
```

**Scan ve Scanln Farkı**

```
Adınız: John Doe
```

**fmt.Scan() fonksiyonu** "John" değerini name değişkenine atar ve "Doe" değeri ise bir sonraki Scan() çağrısında yaş değeri olarak işlenir.

**fmt.Scanln() fonksiyonu** ise, "John Doe" değerini name değişkenine atar ve Enter tuşuna basılana kadar age değişkenini bekler.

Bu farklar, girdilerin nasıl okunduğunu ve işlendiğini belirler. **Scan()** boşlukla ayrılmış birden fazla girişi alırken, **Scanln()** bir satırın tamamını alır.