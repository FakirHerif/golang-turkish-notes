# Arrays

**'array' (dizi)** bir programlama terimidir ve genellikle bir programlama dilinde birden fazla veriyi tek bir değişken altında depolamak için kullanılır.

Dizilerde index 0'dan başlar. Yani ilk elemanın index numarası 0'dır.

```
sayilar = [1, 2, 3, 4, 5]

print(sayilar[0])       // ÇIKTI: 1
sayilar[0] = 9         // İlk indexteki sayıyı 9 ile değiştirdik
print(sayilar[0])       // ÇIKTI: 9
```

Eğer örneğin 5 elemanlı int türünde bir dizi oluşturmak istersek bunu şu şekilde yapabiliriz:

```
var x [5]int        // 5 elemanlı int türünde bir dizi oluşturduk

x[0] = 2            // Dizinin ilk elemanına (index 0) 2 değerini atadık
fmt.Printf(x[1])    // Dizinin ikinci elemanını (index 1) ekrana yazdırdık. ÇIKTI: 0 --- Neden diye sorabileceğinizi tahmin edebiliyorum, bunun 0 olmasının sebebi ise, index 1'e herhangi bir atama yapmadığımız için otomatik olarak 0 değerini alıyor.
```

**Array Literal**

Go programlama dilinde, bir dizi oluşturmak için kullanılan bir ifade türüne "array literal" denir. Array literal, bir dizi içindeki elemanları belirlemek için kullanılır ve genellikle **{}** süslü parantezler içinde tanımlanır.

Array literal'lar, Go'da sabit uzunluktaki diziler oluşturmak için de kullanılır. Sabit uzunluktaki diziler, eleman sayısının başlangıçta belirli olduğu ve daha sonra değiştirilemeyen dizilerdir. Bu tür diziler genellikle array literal ile oluşturulur ve fonksiyonlar veya veri yapısı tanımlamaları için kullanılır.

Örnek olarak, bir sayılar dizisi oluşturmak için array literal kullanımı şu şekildedir:

```
sayilar := [4]int{1, 2, 3, 4}

veya

var sayilar2 [5]int = [5]{1,2,3,4,5}
```

Burada "[ 4 ]int" ifadesi, 4 elemanlı bir tamsayı dizisi tanımlar. Dizinin elemanları ise süslü parantezler içinde {} belirtilir ve virgülle ayrılır.

Go'da array literal kullanımı aynı zamanda elemanların türünü ve sayısını belirler. Bu nedenle, array literal oluştururken belirtilen eleman sayısı, oluşturulan dizinin boyutunu belirler. Bu boyutun aşılmaması önemlidir; aksi takdirde bir hata alabilirsiniz.

**[ ... ] / Dinamik Boyutlu Dizi**

Go programlama dilinde eleman sayısını belirtmeden (üç nokta ... kullanarak) bir dizi oluşturabiliriz.

Bu durum, eleman sayısını belirtmek yerine, dizinin kaç elemana sahip olduğunu derleyicinin otomatik olarak belirlemesine izin verir. Derleyici bu durumda, süslü parantez içinde belirtilen elemanların sayısına bakarak dizinin boyutunu otomatik olarak belirler. Örneğin:

```
x := [...]int{1, 2, 3, 4}
```

ifadesi, içindeki değerlere dayanarak otomatik olarak bir dizi boyutu belirleyen bir dizi oluşturur.

Bu örnek, 4 tamsayı elemanını içeren bir dizi oluşturur ve eleman sayısını belirtmeden (... kullanarak) derleyiciye boyutu belirlemesi için bırakır. Bu durumda, x dizisinin boyutu 4 olacaktır.

Ayrıca bu kullanım, bir dizi oluştururken eleman sayısını açıkça belirtmek yerine, derleyicinin dizinin boyutunu elemanların sayısına göre belirlemesini tercih ettiğiniz durumlarda faydalı olabilir. Bu şekilde, dizi elemanlarını girmekle yetinirsiniz ve derleyici dizinin boyutunu sizin için belirler.

**Diziler Üzerinde Döngü**

Go programlama dilinde diziler üzerinde döngü yaparak elemanlara erişmek ve işlem yapmak için çeşitli yöntemler bulunmaktadır. Dizi üzerinde döngü yapmak için genellikle **for** döngüsü kullanılır. Bunun yanında ayrıca **range** yöntemi de kullanılır.

- **Standart for Döngüsü ile İterasyon:**

```
sayilar := [4]int{10, 20, 30, 40}

for i := 0; i < len(sayilar); i++ {
    fmt.Println(sayilar[i])
}
```

Bu yöntemde, len() fonksiyonu ile dizinin uzunluğu alınır ve ardından for döngüsü içinde dizinin elemanlarına sırasıyla erişilir.

- **range Anahtar Kelimesi Kullanarak İterasyon:**

```
sayilar := [4]int{10, 20, 30, 40}

for index, value := range sayilar {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

range anahtar kelimesi, bir dizi veya diğer veri yapıları üzerinde döngü yapmayı sağlar. Bu yöntemde index ve value olmak üzere iki değişken kullanarak hem elemanın değeri hem de indis bilgisi elde edilir.

Dizi üzerinde döngü yaparken for döngüsü veya range anahtar kelimesi kullanılabilir. range genellikle daha temiz ve okunabilir bir yöntemdir.

Döngülerdeki for anahtar kelimesinin kullanımı, dizinin elemanlarını işlemek için tercih edilen yönteme bağlı olarak değişebilir.