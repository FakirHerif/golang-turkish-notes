# Variables (Değişkenler)

Bildiğimiz üzere değişkenler, değerleri depolamak için kullanılan isimlendirilmiş depolama alanlarıdır.

# Go dilinde değişken adı belirleme kuralları

- Değişken adları harf (a-z, A-Z) veya alt çizgi (_) ile başlamalıdır.
- İsimler, harfler, rakamlar ve alt çizgi içerebilir, ancak rakamla başlayamazlar.
- Değişken adları büyük-küçük harf duyarlıdır (case-sensitive).
- Gömülü anahtar kelimeler (keywords) olarak belirlenmiş olan ifadeler değişken adları olarak kullanılamaz. Örneğin: break, case, chan, const, continue, default, defer, else, fallthrough, for, func, go, goto, if, import, interface, map, package, range, return, select, struct, switch, type, var gibi.
- İsimlendirme genellikle açıklayıcı ve anlamlı olmalıdır. Değişken adları, o değişkenin amacını ve içeriğini anlatmalıdır.

# Go'da basit bir değişken örneği

```
var x int
```

**var** = değişken beyanı için anahtar kelime
**x** = isim
**int** = tür

**Not:** Go dilinde type belirtmemiz gerekir. (int, string, float64 gibi..)

Go'da tek satırda birden fazla değişken tanımlanabilir(virgül ile ayırarak istediğimiz kadar yazabiliriz):

```
var x,y int
```

# Bazı Değişken Türleri

- **Integer**: Tam sayıları ifade eder. (-1, 5, -72, 3, 10)
- **Fractioal(decimal)**: Tam olmayan sayıları ifade eder. (0.5, 1.25, -3.75) **Not:** Go dilinde sabir bir decimal türü yoktur. Ondalık sayılar için "float32" ve "float64" gibi veri tipleri bulunur.
- **Strings**: Metin dizilerini (karakter dizilerini) ifade eder.
- **bool**: Boolean değerleri temsil etmek için kullanılır (true, false)

**Not:** Go'da değeri bildirilmeyen değişkenler 0 veya boş dize değerini alır. Yani **"var x int" değeri x = 0 olarak**, **"var x string" değeri x = ""** olarak değer alır.

**Go'da var yerine := kullanılabilir**

```
x := 100
```

Değişken tür bildirimi yapmazsak otomatik olarak yapılacaktır. Ama bu şekilde değişken türü belirtmeden kullanmak çok uygun değildir. 