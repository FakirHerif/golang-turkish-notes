# String Packages

Bir önceki notlarda bahsettiğimiz üzere; Dizeler ortak bir yapıdır, bu yapıları tüm programlama dillerinde görürüz ve dizeler Unicode Rune'lerden oluşur.

Go'da bulunan string paketi; dize işlemleri için bir dizi işlev sağlar. Dize birleştirme, bölme, değiştirme, arama gibi işlemleri gerçekleştirmek için kullanılır.

**strings.Compare() Fonksiyonu:** Bu fonksiyon, strings paketinde bulunur ve iki dizeyi karşılaştırır. Karşılaştırma sonucunda, eğer ilk dize ikinci dizeden küçükse, sonuç -1 olur; eğer eşitlerse sonuç 0 olur; ve eğer ilk dize ikinci dizeden büyükse, sonuç 1 olur. Örneğin:

```
result := strings.Compare("abc", "def")
fmt.Println(result)                     // Çıktı: -1
```

**strings.Contains() Fonksiyonu:** Go programlama dilindeki strings paketinde bulunan Contains fonksiyonu, bir dizenin başka bir dizeyi içerip içermediğini kontrol eder. Örneğin:

```
str := "Merhaba, bu bir örnek cümle."

contains := strings.Contains(str, "örnek")
fmt.Println(contains)                       // Çıktı: true
```

**strings.HasPrefix() Fonksiyonu:** Go programlama dilindeki strings paketinde bulunan HasPrefix fonksiyonu, bir dizenin belirli bir önek ile başlayıp başlamadığını kontrol eder. Örneğin:

```
str := "Merhaba, nasılsın?"

hasPrefix := strings.HasPrefix(str, "Merhaba")
fmt.Println(hasPrefix)                          // Çıktı: true
```

**strings.Index() Fonksiyonu:** Go programlama dilindeki strings paketinde bulunan Index fonksiyonu, bir dizenin içinde belirli bir alt dizenin ilk kez göründüğü konumu (index'ini) döndürür. Örneğin:

```
str := "Merhaba, nasılsın?"

index := strings.Index(str, "nasıl")
fmt.Println(index) // Çıktı: 9
```

Yukarıdaki örnekte, strings.Index() fonksiyonu "Merhaba, nasılsın?" dizesindeki "nasıl" alt dizesinin ilk kez hangi indekste başladığını bulur. Bu durumda, "nasıl" alt dizesi 9. indexten başlamaktadır ve bu değer döndürülür. Eğer belirtilen alt dize dizede bulunmuyorsa, strings.Index() fonksiyonu -1 değerini döndürür.

**Unicode Packages**

Ayrıca dizelerin içindeki farklı rune'lerin özelliklerini gerçekten değerlendiren ve bir dizi işlev sağlayan Unicode Paketi vardır.

Bu paket, Unicode karakterlerinin işlenmesi için kullanılır. Bu paket, Unicode karakterlerinin özelliklerini değerlendirmek ve işlemek için çeşitli işlevler sunar. 

Özellikle, runelerin özelliklerini değerlendirmek için sağladığı işlevler vardır. Örneğin, **"IsDigit"** fonksiyonu, bir rune karakterinin bir sayısal basamak olup olmadığını belirler. Yani, verilen rune karakterin bir sayı olup olmadığını kontrol eder. Örneğin:

```
var r rune = '7'
    if unicode.IsDigit(r) {
		fmt.Println("Bu bir sayısal basamaktır.")
	} else {
		fmt.Println("Bu bir sayısal basamak değildir.")
	}
```

Ayrıca;

**IsSpace:** Bu fonksiyon, bir rune karakterinin boşluk karakteri olup olmadığını belirler. Boşluk karakterleri, bir metinde görsel olarak boşluk bırakan karakterlerdir. Boşluk, sekme, yeni satır gibi karakterleri içerir. Örneğin:

```
var r rune = ' '

	if unicode.IsSpace(r) {
		fmt.Println("Bu bir boşluk karakteridir.")
	} else {
		fmt.Println("Bu bir boşluk karakteri değildir.")
	}
```

Bu örnek, ' ' (boşluk karakteri) için "Bu bir boşluk karakteridir." çıktısını verecektir.

**IsLetter:** Bu fonksiyon, bir rune karakterinin bir harf karakteri olup olmadığını belirler. Yani, alfabedeki harfleri temsil eden karakterlerdir. Örneğin:

```
var r rune = 'A'

	if unicode.IsLetter(r) {
		fmt.Println("Bu bir harf karakteridir.")
	} else {
		fmt.Println("Bu bir harf karakteri değildir.")
	}
```

Bu örnek, 'A' için "Bu bir harf karakteridir." çıktısını verecektir.

**IsLower:** Bu fonksiyon, bir rune karakterinin küçük harf karakteri olup olmadığını belirler. Yani, alfabedeki küçük harf karakterleri için doğrudur. Örneğin:

```
var r rune = 'b'

	if unicode.IsLower(r) {
		fmt.Println("Bu bir küçük harf karakteridir.")
	} else {
		fmt.Println("Bu bir küçük harf karakteri değildir.")
	}
```

Bu örnek, 'b' için "Bu bir küçük harf karakteridir." çıktısını verecektir.

**IsPunct:** Bu fonksiyon, bir rune karakterinin noktalama işareti olup olmadığını belirler. Noktalama işaretleri, noktalar, virgüller, ünlem işaretleri gibi metnin anlamını vurgulayan karakterlerdir. Örneğin:

```
var r rune = '!'

	if unicode.IsPunct(r) {
		fmt.Println("Bu bir noktalama işareti karakteridir.")
	} else {
		fmt.Println("Bu bir noktalama işareti karakteri değildir.")
	}
```

Bu örnek, '!' için "Bu bir noktalama işareti karakteridir." çıktısını verecektir.


**String Manipulation**

Go'da string manipülasyonu, metin veya dize verileri üzerinde çeşitli işlemler yapmayı ifade eder. Dize manipülasyonu genellikle metinlerin birleştirilmesi, bölünmesi, değiştirilmesi, aranması, dönüştürülmesi ve işlenmesi gibi çeşitli işlemleri içerir. 

**strings.Replace() fonksiyonu:** Bu fonksiyon, belirli bir dizede belirli bir alt dizeyi başka bir alt dizeyle değiştirir. Örneğin:

```
str := "Bu bir örnek metindir. Örnek metin."

// "Örnek" alt dizelerini "Deneme" ile değiştir
replaced := strings.Replace(str, "Örnek", "Deneme", -1)
fmt.Println(replaced)                   // Çıktı: Bu bir deneme metindir. Deneme metin.
```

Eğer belirli bir alt dize sadece belirli sayıda değiştirilsin isteniyorsa, üçüncü bir parametre olarak değiştirme sayısını belirtebilirsiniz. -1 kullanarak tüm alt dizelerin değiştirilmesini sağlarsınız, diğer sayılar ise belirli bir sayıda değiştirme yapılmasını belirtir.

**ToUpper:** Bu fonksiyon, bir dizenin içindeki tüm harf karakterlerini büyük harfe dönüştürmek için kullanılır. Örneğin:

```
str := "Merhaba Dünya"

upper := strings.ToUpper(str)
fmt.Println(upper)              // Çıktı: MERHABA DÜNYA
```

**ToLower:** Bu fonksiyon, bir dizenin içindeki tüm harf karakterlerini küçük harfe dönüştürmek için kullanılır. Örneğin:

```
str := "MERHABA DÜNYA"

lower := strings.ToLower(str)
fmt.Println(lower)              // Çıktı: merhaba dünya
```

**strings.TrimSpace() Fonksiyonu:** bir dizenin başındaki ve sonundaki boşlukları (boşluk, tab, yeni satır gibi karakterler) temizler ve dizeyi temizlenmiş bir hale getirir. Örneğin:

```
str := "   Bu bir örnek cümledir.   "

trimmed := strings.TrimSpace(str)
fmt.Println(trimmed)                    // Çıktı: "Bu bir örnek cümledir."
```

**Strconv Package**

Strconv paketi, (string conversion) Go programlama dilinde string dönüşümlerini ve sayısal türlerin stringlere dönüştürülmesini sağlayan bir pakettir. Bu paket, string ve sayısal türler arasında dönüşüm yaparken kullanılır.

- **Sayısal Veri Tiplerini String'e Dönüştürme:**
    - **FormatInt:** Bir tamsayıyı belirli bir sayı tabanında string olarak biçimlendirir.
    - **FormatUint:** Bir işaretli olmayan tamsayıyı belirli bir sayı tabanında string olarak biçimlendirir.
    - **FormatFloat:** Bir float sayıyı string olarak biçimlendirir.
    - **Itoa:** Bir tamsayıyı onluk sistemde string bir formata dönüştüren bir fonksiyondur. "Itoa", "Integer to ASCII"nin kısaltmasıdır 

- **String'i Sayısal Veri Tiplerine Dönüştürme:**
    - **Atoi:** Bir string'i integer türüne dönüştürür.
    - **ParseInt:** Bir string'den tamsayı türüne dönüşüm yapar.
    - **ParseUint:** Bir string'den işaretli olmayan tamsayı türüne dönüşüm yapar.
    - **ParseFloat:** Bir string'den float türüne dönüşüm yapar.

Örnek bir kullanım:

```
// Sayısal veriyi string'e dönüştürme
integerToString := strconv.Itoa(42)
fmt.Println("Integer to String:", integerToString)

// String'i sayısal veri türüne dönüştürme
str := "123"
strToInt, err := strconv.Atoi(str)
if err != nil {
	fmt.Println("Hata:", err)
	return
}
fmt.Println("String to Integer:", strToInt)
```

Bu örnek, strconv paketindeki Itoa fonksiyonu ile bir tamsayıyı string'e ve Atoi fonksiyonu ile bir string'i tamsayıya dönüştürmeyi göstermektedir.