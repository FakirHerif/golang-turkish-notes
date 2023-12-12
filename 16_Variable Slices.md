# Variable Slices

make fonksiyonu, özellikle slice, map veya channel gibi belirli veri yapılarını oluşturmak için kullanılan bir fonksiyondur. Genellikle dinamik boyutlara sahip veri yapılarını oluştururken kullanılır. Bu fonksiyon, belirli bir veri tipi için bir yapı oluşturur ve başlangıç değerleriyle birlikte kullanılabilir.

İşte make() fonksiyonunun slice için kullanımı ve örnekleri:

```
sli := make([]int, 10)
```

Bu kullanım, 10 elemanlı bir integer slice oluşturur. Bu slice'ın başlangıç uzunluğu ve kapasitesi 10 olur. Ancak, başlangıçta bu slice'ın elemanları sıfır değerine eşit olur.

```
sli := make([]int, 10, 15)
```

Bu kullanım, 10 elemanlı bir integer slice oluşturur ve bu slice'ın kapasitesini 15 olarak belirler. Başlangıç uzunluğu 10 olacaktır, ancak kapasitesi 15 olur. Kapasite, bir slice'ın mevcut elemanları depolayabileceği maksimum eleman sayısını ifade eder. Genellikle slice'ın daha büyük bir kapasiteye sahip olmasını ve daha sonra daha fazla eleman eklenirken performans artışı sağlanmasını sağlar. Kapasite, slice'ın dinamik olarak büyütülebileceği bellek alanını belirtir.

Özetle, make() fonksiyonu, slice, map ve channel gibi dinamik veri yapılarını oluşturmak için kullanılır. Bu yapılar önceden belirlenmiş başlangıç uzunluğuna ve kapasiteye sahip olabilirler ve ihtiyaç halinde bu kapasite dinamik olarak genişletilebilir.

Bir slice'a yeni bir veya birden fazla eleman eklemek için append() fonksiyonunu kullandığımızı hatırlayalım.

```
sli := make([]int, 0, 3)    // Uzunluğu 0, kapasitesi 3 olan bir integer slice oluşturduk
sli = append(sli, 100)      // 100 değerini slice'a ekledik
```

Bu işlem, sli adındaki slice'a 100 değerini ekler. Şu anki durumda sli slice'ının uzunluğu 1 ve kapasitesi 3 olacaktır. Çünkü başlangıçta sıfır elemanlı bir slice'a 100 değeri eklendi.

Eğer slice'ın kapasitesi aşılırsa, append() fonksiyonu otomatik olarak kapasiteyi genişletecek yeni bir slice oluşturur ve mevcut elemanları yeni slice'a kopyalar. Bu, kapasitenin aşıldığı durumda otomatik olarak büyüyen dinamik bir yapı sağlar. Yani, eğer slice'ın uzunluğu kapasiteyi aşarsa, append işlemi yeni bir slice oluşturarak daha fazla eleman eklemeyi mümkün kılar. Bu özellik, dinamik boyutlara sahip veri yapılarını yönetirken oldukça faydalıdır.