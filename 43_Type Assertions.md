# Type Assertions / Tür Belirtimi

Go dilinde "type assertion" yani "tür belirtimi", bir interface değerinin altındaki (underlying) değeri almak için kullanılan bir mekanizmadır. İnterface türü içinde saklanan değerin asıl türünü belirlemek için kullanılır.

İlk olarak, bir örnekle başlayalım:

```
var val interface{}  // boş bir interface tanımlıyoruz

val = 42  // val içine bir int değeri atıyoruz

// Şimdi, bu interface içindeki değeri almak için type assertion kullanalım:
if num, ok := val.(int); ok {
    fmt.Println("Bu bir int:", num)
} else {
    fmt.Println("Bu bir int değil")
}
```

Bu örnekte, val adlı boş bir interface oluşturduk ve içine bir int değeri atadık. Sonra val içindeki değeri almak için val.(int) ifadesini kullandık. Bu ifade, val'in altındaki değerin int türünde olup olmadığını kontrol eder. Eğer değer int ise, num adlı bir değişken içine alınır ve ok adlı ikinci bir değerle birlikte döner. ok, tür belirtiminin başarılı olup olmadığını belirtir.

Buna karşılık, eğer val içindeki değer bir int değilse, ok değeri false olur ve num herhangi bir anlam ifade etmez.

Bunu bir başka örnekle pekiştirelim:

```
var val interface{}  // boş bir interface

val = "Merhaba"  // val içine bir string değeri atıyoruz

// Şimdi, bu interface içindeki değeri almak için type assertion kullanalım:
if str, ok := val.(string); ok {
    fmt.Println("Bu bir string:", str)
} else {
    fmt.Println("Bu bir string değil")
}
```

Bu sefer val içine bir string değeri atadık ve aynı type assertion yapısını kullanarak, bu değerin bir string olup olmadığını kontrol ettik.

Tür belirtimi, interface içindeki değeri çeşitli türlerde kullanmak istediğinizde oldukça kullanışlıdır. Ancak, dikkatli olunmalıdır çünkü eğer belirtilen tür ile gerçekteki tür uyuşmazsa, program hata verebilir. Bu nedenle, genellikle ok değerini kontrol etmek önemlidir.

```
var val interface{}  // boş bir interface tanımlıyoruz

val = "Merhaba"  // val içine bir string değeri atıyoruz

// Şimdi, bu interface içindeki değeri almak için type assertion kullanalım:
if num, ok := val.(int); ok {
    fmt.Println("Bu bir int:", num)
} else {
    fmt.Println("Bu bir int değil")
}
```

Burada, val aslında bir string değerini içeriyor ve biz onu bir int olarak almaya çalışıyoruz. Bu durumda ok değeri false olacak ve "Bu bir int değil" mesajını verecek.

Bu kontrolü eklemek, programın hata vermesini önleyerek daha güvenli bir şekilde çalışmasını sağlar. Özellikle runtime sırasında bir hataya neden olabilecek bu tür durumları kontrol etmek iyi bir uygulama pratiğidir.