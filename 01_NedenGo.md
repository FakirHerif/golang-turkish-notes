# KISACA NEDEN GO?

Go, nesne yönelimlidir. Nesne yönelimi diğer dillere göre basitleştirilmiştir. Ancak diğer dillere göre örneğin python, java veya c++ gibi nesne yönelimli dillerde gördüğümüz daha az özelliğe sahiptir.

Go'da class terimi yerine structs kullanıyoruz.

Geleneksel class uygulamalarına kıyasla basit bir şekildedir. Yani inheritance, constructors, generics yoktur. Bunun kod yazmayı kolaylaştırdığını ve çalıştırmayı verimli hale getirdiğini söyleyebiliriz. Yani, genellikle daha hızlı çalışır.

Go'nun en büyük avantajlarından biri eşzamanlılık uygulamasıdır. Eşzamanlılık kısaca aynı anda birden fazla görevin yönetimidir. Büyük sistemler için anahtardır. Yani birçok parça çalışıyor, devam ediyor ama hepsi sırayla yürütülmüyor. Bir nevi paralellik diyebiliriz. Ama eşzamanlılık her zaman performans iyileştirmesini garanti etmez.  Birden fazla görevin aynı anda yürütülmesine izin vererek performansı potansiyel olarak artırabilse de, etkinlik büyük ölçüde görevlerin doğası, donanım sınırlamaları, eşzamanlı işlemleri yönetme yükü, senkronizasyon gereksinimleri ve daha fazlası gibi çeşitli faktörlere bağlıdır.  Bazı durumlarda eşzamanlı görevleri yönetmenin getirdiği ek yük, düzgün şekilde uygulanmadığı takdirde performansı düşürebilir. Bu nedenle eşzamanlılık güçlü bir kavram olsa da tüm senaryolarda performansın artmasını garanti etmez.

**Özetlemek Gerekirse:**

- Kodlar hızlı çalışır
- Garbage collection (Artık kullanılmayan nesnelerin dağıtımının kaldırılması)
- Basit nesneler
- Yerleşik eşzamanlılık
- Machine language (Doğrudan işlemci üzerinde çalışır)
- Assembly language 
- High-level language
    - Go bu üç kategoride de üst düzey bir dildir
- Compile yani derlenmiş bir dildir, yorumlayıcı diller bilindiği üzere bizi yavaşlatır. Compile olması hız açısından bir avantajdır. Go yorumlanmış dillerin bazı iyi özelliklerine de sahiptir. Üstte yazdığımız gibi "Garbage Collection" (yani otomatik bellek yönetimi)

# GO Tools

- **go build** komutu Go projelerinizi derlemek ve çalıştırılabilir dosyalar oluşturmak için kullanılır.
- **go doc** komutu Go programlama dilinde belgelendirme oluşturmanın ve kullanmanın temel bir yoludur.
- **go fmt** komutu belirtilen klasördeki veya dosyadaki tüm Go kaynak kodlarını standart bir biçimde düzenleyecektir. Bu sayede, projenizdeki tüm kodlar Go'nun önerdiği standartlara uygun hale gelir.
- **go get** Go dilinde harici paketleri indirmek ve yüklemek için kullanılan bir komuttur. Genellikle Go dilindeki projelerde dışarıdan alınan paketleri veya kütüphaneleri kullanmak istediğinizde go get komutunu kullanırsınız.
- **go list** Go dilinde paketlerin, modüllerin veya çalışma zamanındaki öğelerin listesini oluşturmak ve göstermek için kullanılan bir komuttur.
- **go run** Go programlama dilinde kaynak dosyalarını derlemek ve ardından elde edilen çalıştırılabilir dosyayı doğrudan çalıştırmak için kullanılan bir komuttur.
- **go test** Go dilinde test dosyalarını çalıştırmak ve test sonuçlarını görüntülemek için kullanılan bir komuttur.