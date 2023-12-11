# Control Flow

Kontrol akışı (Control flow), bir programın çalışma sırasını belirleyen, hangi kod bloklarının çalışacağını, hangi koşullarda veya durumlarda hangi kod parçalarının yürütüleceğini kontrol eden yapıdır. Programlama dillerinde kontrol akışı, programın belirli şartlara veya durumlara göre davranışını belirler.

Kontrol akışı, genellikle şu temel yapılar üzerinde gerçekleştirilir:

**Sıralı İşlemler (Sequential Execution):** Kodun baştan sona doğru sırasıyla çalıştığı durumdur. Her satır veya ifade bir sonraki satır veya ifadenin ardından çalışır. Yani yukarıdan aşağıya.

**Koşul İfadeleri (Conditional Statements):** if, else if, else gibi yapılarla belirli koşullara bağlı olarak belirli kod bloklarının çalışmasını sağlar. Örneğin:

```
if x > 10 {
    // x değeri 10'dan büyükse burası çalışır
} else if x == 10 {
    // x değeri 10'a eşitse burası çalışır
} else {
    // x değeri 10'dan küçükse burası çalışır
}
```

**Döngüler (Loops):** for, while, do-while gibi döngü yapıları, belirli koşullar sağlandığı sürece belirli bir kod bloğunun tekrarlanmasını sağlar.

```
for i := 0; i < 5; i++ {
    // 0'dan 4'e kadar olan değerler için bu kod bloğu çalışır
}
```

**Deyimler (Statements / Switch-Case):** switch-case gibi kontrol yapıları, farklı durumlar için kodun hangi bloklarının çalışacağını belirler.

```
switch day {
case "Monday":
    // Pazartesi günü için bu kod bloğu çalışır
case "Tuesday":
    // Salı günü için bu kod bloğu çalışır
default:
    // Diğer durumlar için bu kod bloğu çalışır
}
```

Bu kontrol akışı yapıları, programların belirli koşullar altında nasıl davranacaklarını ve hangi kod parçalarının çalışacağını yönetir. Programların mantıklı bir şekilde işlemesini ve istenen sonuçları elde etmeyi sağlar.
