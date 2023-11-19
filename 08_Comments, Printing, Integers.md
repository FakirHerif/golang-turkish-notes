# Comments, Printing, Integers

**Single-line Comments:** Yorumlar " // " çift eğik çizgi koyulduktan sonra kullanılır.

Örnek

```
var x int   // YORUM SATIRI

// BAŞKA BİR YORUM SATIRI
```

**Block Comments:** Yorumlar " /* " ve " */ " arasında da yazılabilir.

Örnek

```
var x int

/*  YORUM
BAŞKA BİR YORUM
BİR BAŞKA YORUM DAHA */
```

**Printing:** Go'da yazdırma işlemi **FMT** paketi kullanılarak yapılır. Go'da Printf ve Println kullanılır.

**Printf()** : Biçimlendirilmiş çıktı almak için kullanılır. Özel formatlama belirteçleri (örneğin %s, %d, %f) kullanarak değişken değerlerini belirli bir formatta ekrana yazdırmak için kullanılır. Örneğin, Printf ile bir metin içinde belirli bir yerde değişkenlerin değerlerini yerleştirebilirsiniz.

Örnek

```
name := "Alice"
age := 30
fmt.Printf("İsim: %s, Yaş: %d\n", name, age)
```

**Println()** Değerleri doğrudan yazdırmak için kullanılır. Parametre olarak verilen değerleri ekrana basar ve ardından bir satır atlar. Biçimlendirme veya özel formatlama yapmaz, sadece verilen değerleri yazdırır.

Örnek

```
fmt.Println("Merhaba, Dünya!")
```

Kısacası, **Printf**, biçimlendirilmiş çıktı almak için kullanılırken **Println**, değerleri doğrudan yazdırmak için kullanılır ve yeni bir satır ekler.

**Integers:** Go dilinde, tamsayıları ifade etmek için int veri türü kullanılır. Tamsayılar, negatif ve pozitif tam sayıları içerebilir. int türü, platformunun sözdizimine göre boyutlandırılır. Örneğin, 32-bit sistemlerde int, 32-bit (4 byte) uzunluğunda iken, 64-bit sistemlerde 64-bit (8 byte) uzunluğunda olabilir. Go dilinde ayrıca int8, int16, int32, int64 gibi özel boyutlarda tamsayı türleri de bulunur. Her türün bir alt ve üst sınırları vardır. Örneğin, int8 türü -128 ile 127 arasında değer alabilir. Bu türlerin boyutları ve sınırları platforma bağlı olabilir.

Ayrıca Go dilinde uint adında bir veri türü bulunmaktadır. uint, işaretli olmayan (yani sıfır veya pozitif değerleri temsil eden) tam sayıları ifade eder. Ayrıca uint8, uint16, uint32, uint64 gibi özel boyutlarda işaretli olmayan tamsayı türleri de bulunur. Bu türlerin boyutları ve sınırları platforma bağlı olabilir. Örneğin, uint8 türü 0 ile 255 arasında değer alabilir. Bu türler, özellikle sayılar sıfır veya daha büyük olduğunda kullanılır.

**Binary Operators:** Go dilindeki binary operatörler, iki operandı kullanarak matematiksel ve mantıksal işlemleri gerçekleştiren operatörlerdir. Bazı yaygın binary operatörler şunlardır:

- **Arithmetic:** + - * / % << >>
- **Comparison:** == != > < >= <=
- **Boolean:** && ||