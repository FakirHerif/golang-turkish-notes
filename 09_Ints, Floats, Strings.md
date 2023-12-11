# Ints, Floats, Strings

**Type Conversions**

```
var x int32 = 1
var y int16 = 2

şeklinde 2 değişkenimiz var, yapmak istdiğimiz şey ise
x'i y'ye eşitlemek.
```

Bu durumda her ikisi de tam sayı olmasına rağmen farklı türlerdir. Yani bunu yapabilmek için öncelikle int16 olan y'yi int32'ye dönüştürmemiz gerekir. Bunu da T işlemiyle gerçekleştirebiliriz "T() operation". "T" aslında bir argümandır, T yerine farklı argümanlar kullanabiliriz.

**Convert type with T() operation**

```
x = int32(y)    // Bu yöntemle y'yi int32 türüne dönüşmüştür oluyoruz.
```

**Floating Point**

```
float32 - ~6 digits of precision / 6 ondalıklı basamak
float64 - ~15 digits of precision / 15 ondalıklı basamak
```

Bu sayıları ondalık sayılarla veya bilimsel gösterilmle ifade edebiliriz:

```
var x float64 = 123.45      // ondalık gösterim
var y float64 = 1.2345e2    // bilimsel gösterim
```

Ayrıca karmaşık sayıları da (complex numbers) gösterebiliriz:

```
var z complex128 = complex(2,3)
```

# ASCII ve Unicode

Dizeler hakkında konuşmak için öncelikle ASCII ve Unicode hakkında konuşmamız gerekiyor.

**ASCII :** Temsil etmek istediğimiz her karakter 8 bitlik bir kodla temsil edilir. Bu sadece bir karakter kodlamasıdır. Yani örneğin:

```
'A' = 0X41      // A onaltılık sayı olarak 41'dir. Bu 8bitlik bir kod.
```

Kısaca, ASCII 8 bit uzunluğunda bir koddur, yani maksimum 256 olası karakteri temsil edebilir. 8 bitlik bir kod İngilizce için yeterlidir. Ama örneğin Çince gibi çok daha fazla karakteri olan bir alfabeyi işin içine dahil edersek 8 bitlik kod kullanamayız, 8 bitlik kodla Çince'yi temsil etmeyi bekleyemeyiz. Bu yüzden 8 bitten daha fazlasına ihtiyaç duyarız. İşte tam olarak 32 bit uzunluğunu destekleyen Unicode bu tür durumlar içindir.

**Unicode**

32 bit uzunluğundadır. UTF-8 ise Unicode'un bir alt kümesidir. 8bit UTF kodları ASCII değeriyle aynıdır. Yani 'A' değeri ASCII'de 41 iken UTF-8'de de 41'dir.

Go'da varsıyal UTF-8'dir. Ayrıca Go'da bir kod noktasına "rune" denir.