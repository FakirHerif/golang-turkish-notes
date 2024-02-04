# Parallel Programming

## Sıralı Kod İçinde Gezme / Interleavings

Paralel kod yazmak zordur çünkü programın herhangi bir anındaki durumu zihinde canlı bir şekilde tutmak zordur. Sıralı bir kodda, program belirli bir satırda çökerse, önceki satırların hangilerinin çalıştığını biliyorsunuz. Ancak paralel kodda, her bir görevde (task) çökme meydana gelebilir, bu durumda diğer görevlerin hangi aşamada olduğunu kestirmek zordur.

Paralel kodun karmaşık olmasının nedenlerinden biri, programın genel durumunun belirli olmamasıdır. Bu durum non-deterministik olduğu anlamına gelir. Aynı satırda çökerse bile, diğer görevlerin genel sistem durumu her seferinde farklı olabilir. Bu durum, programın doğru çalışıp çalışmadığını anlamak için zihinsel çaba gerektirir.

Interleaving (gezdirme), farklı görevlerdeki (tasks) komutların sırasının belirsiz olduğu bir durumu ifade eder. Örneğin, task 1'in komutları sırayla 1, 2, 3 ise, task 2'nin komutları da sırayla 1, 2, 3 olabilir. Fakat farklı görevler arasındaki komut sıralaması belirsiz ve her seferinde farklı olabilir.

İnterleaving problemi, komutların sadece Go kodu veya C kodu düzeyinde değil, aynı zamanda makine kodu düzeyinde de gerçekleşebilir. Yani, bir komutun tamamlanması, diğer görevdeki bir komutun başlamasından önce tamamlanmış olmayabilir.

## Yarış Koşulları / Race Conditions

Yarış koşulları, paralel programlamanın zorluklarından biridir. Yarış koşulları, programın sonucunun gezdirme sırasına bağlı olduğu bir durumu ifade eder. Bu durumda, gezdirme sıralaması non-deterministik olduğu için programın sonucu da non-deterministik olabilir.

İletişim, görevler arasında veya Go dilindeki tabiriyle Goroutine'ler arasında gerçekleşir ve yarış koşullarına neden olabilir. Eğer iki görev arasında iletişim varsa, yani bir değişken paylaşılıyorsa, bu durum yarış koşullarına neden olabilir. Bu iletişim, programın sonucunun gezdirme sırasına bağlı olmasına yol açabilir.

Örnek olarak, iki görev arasında paylaşılan bir değişken olan 'x' üzerinde işlemler gerçekleştiriliyor. Bu durumda, gezdirme sırasına bağlı olarak 'x' değeri değişebilir ve program non-deterministik hale gelebilir. İletişimi kontrol etmek için belirli teknikler kullanmak önemlidir.

Genel olarak, paralel programlamada görevler genellikle bağımsızdır, ancak arada iletişim olabilir. İletişim, yarış koşullarına neden olabilir ve bu da programın doğru çalışmasını engelleyebilir. İletişimi kontrol etmek ve yarış koşullarını önlemek için belirli teknikler kullanmak önemlidir.