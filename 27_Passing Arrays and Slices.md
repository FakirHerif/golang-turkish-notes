# Passing Arrays and Slices

Go dilinde, işlevlere veri aktarmak için hem diziler (arrays) hem de dilin esnekliğini sağlayan dilimler (slices) kullanılabilir. İkisi arasındaki ana fark, bir dizinin sabit boyutta olması ve bir dilimin dinamik olarak boyutlandırılabilmesidir.

Go'da, bir fonksiyona bir dizi aktardığınızda, o dizi tümüyle kopyalanır. Eğer dizi büyükse bu kopyalama süreci zaman alabilir ve fazla bellek kullanabilir. Bu nedenle, büyük dizilerle çalışırken bu durum sorun olabilir.

Bu sorunu çözmek için, Go'da dizi işaretçilerini kullanabilirsiniz. Dizi işaretçileri, aslında diziye bir referans sağlar ve bu referans üzerinden diziye erişim sağlar. Ancak Go dilinde, daha yaygın ve pratik bir yöntem dilimlerin kullanılmasıdır.

Dilimler, dizilerin esnek versiyonlarıdır ve dinamik boyuta sahiptirler. Bir dilimin altında bir dizi bulunur ve dilim, bu dizinin bir kesitidir. Bir dilimi bir fonksiyona aktardığınızda, aslında bir işaretçi aktarmış olursunuz. Bu nedenle, fonksiyon, dilimin başlangıç adresini, uzunluğunu ve kapasitesini alır.

Dolayısıyla, Go'da büyük veri koleksiyonlarıyla çalışırken diziler yerine dilimleri kullanmak daha iyidir. Bu sayede, bellek ve zaman açısından daha verimli bir şekilde işlem yapabilirsiniz. Bir fonksiyona veri aktarırken, dilimleri kullanmak, kopyalama maliyetini azaltır ve bellek kullanımını optimize eder.

Özetle, Go'da diziler yerine dilimleri kullanmak, özellikle büyük veri koleksiyonlarıyla çalışırken daha etkili bir yöntemdir. Bu, daha az bellek kullanımı ve daha hızlı işlemler sağlar.

Örneklerle üç ana başlığı değerlendirelim:

**Normal bir dizi kullanımı:**

```
func foo(x [3]int) int {
    return x[0]                 // Dizinin ilk elemanını döndürür
}

func main() {
    a := [3]int{1, 2, 3}        // 3 elemanlı bir dizi tanımlar
    fmt.Print(foo(a))           // foo fonksiyonunu çağırır ve ilk elemanı yazdırır
}
```

Bu örnekte, foo fonksiyonu, [ 3]int türünden bir dizi alır ve bu dizinin ilk elemanını (x[ 0]) döndürür. main fonksiyonunda, a adında bir [ 3]int dizisi tanımlanır ve foo fonksiyonu a dizisini çağırır. Fonksiyon, dizinin ilk elemanını döndürür ve bu değer ekrana yazdırılır.

**Dizi İşaretçisi Kullanımı:**

```
func foo(x *[3]int) int {
    (*x) [0] = (*x) [0] + 1     // Dizinin ilk elemanını bir artırır
}

func main() {
    a := [3]int{1, 2, 3}        // 3 elemanlı bir dizi tanımlar
    foo(&a)                     // foo fonksiyonunu dizinin adresiyle çağırır
    fmt.Print(a)                // Dizinin tüm elemanlarını yazdırır
}
```

Bu örnekte, foo fonksiyonu, [ 3]int türünden bir dizi işaretçisi alır. main fonksiyonunda, a adında bir [ 3]int dizisi tanımlanır ve foo fonksiyonu a dizisinin adresini (&a) alarak çağırılır. Fonksiyon, işaretçi aracılığıyla dizinin ilk elemanını bir artırır ve bu değişiklik ana işlevdeki a dizisine yansır. Sonuç olarak, dizinin ilk elemanı arttırılır ve fmt.Print(a) ifadesi ile dizinin tüm elemanları ekrana yazdırılır.

**Dilim Kullanımı:**

```
func foo(sli int) int {
    sli[0] = sli[0] + 1         // Dizinin ilk elemanını bir artırır
}

func main() {
    a := []int{1,2,3}           // 3 elemanlı bir dizi tanımlar
    foo(a)                      // foo fonksiyonunu dilimle çağırır çağırır
    fmt.Print(a)                // Dilimin tüm elemanlarını yazdırır
}
```

Bu örnekte, foo fonksiyonu, []int türünden bir dilim alır. main fonksiyonunda, a adında bir []int dilimi tanımlanır ve foo fonksiyonu a dilimini çağırır. Fonksiyon, dilimin ilk elemanını bir artırır ve bu değişiklik ana işlevdeki a dilimine yansır. Sonuç olarak, dilimin ilk elemanı arttırılır ve fmt.Print(a) ifadesi ile dilimin tüm elemanları ekrana yazdırılır.

Bu örneklerde, farklı yaklaşımlarla dizi veya dilim elemanlarının değiştirilmesi gösteriliyor. Dizi işaretçileri doğrudan bellek adresleri üzerinden işlem yaparken, dilimler referanslarla çalışır ve bu nedenle dilimler, içerikleri değiştirilebilen esnek veri yapılarıdır.