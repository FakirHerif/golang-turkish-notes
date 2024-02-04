# Why Use Concurrency? / Neden Eşzamanlılık Kullanırız?

## Go Dilinde Concurrency

Go dilinin büyük bir özelliği, dilin içine concurrency (eş zamanlılık) özelliğinin dahil edilmiş olmasıdır. Diğer dillerde, örneğin C veya Python gibi dillerde, genellikle concurrent (eş zamanlı) programlama yapmak için dış kütüphaneleri içe aktarırsınız. Ancak Go dilinde, bu özellik dilin içine yerleştirilmiştir, yani dilin bir parçasıdır.

## Paralel Yürütme

Paralel yürütme, iki programın tam olarak aynı anda çalıştığı durumu ifade eder. Genellikle, işlemciler bir seferde bir komut çalıştırmak üzere tasarlanmıştır. Ancak paralel yürütme için gerçek bir hız artışı elde etmek için iki işlemci veya en azından iki işlemci çekirdeği gereklidir.

## Paralel Yürütmenin Faydaları

Paralel yürütme, görevlerin daha hızlı tamamlanmasına olanak tanır. Bu, bir görevin paralel olarak çalıştırıldığında, genel olarak tüm görevlerin daha hızlı tamamlandığı anlamına gelir. Ancak, bazı görevlerin sıralı bir şekilde gerçekleştirilmesi gerektiğini unutmamak önemlidir.

## Güç Duvarı

Güç duvarı, transistörlerin güç tüketiminin artması ve sıcaklık sorunlarının ortaya çıkması nedeniyle işlemcilerin frekanslarını artıramamalarını ifade eder. Bu durumu aşmak için, tasarımcılar çoklu çekirdekli sistemlere yönelmişlerdir. Ancak çoklu çekirdekli sistemlerden faydalanmak için paralel yürütme ve eş zamanlı programlama gerekir.

## Eş Zamanlı Programlamanın Önemi

Paralel sistemlerde performans artışı elde etmek için, programcıların kodlarını çoklu çekirdekli sistemlere uygun bir şekilde yazmaları gerekmektedir. Bu nedenle, eş zamanlı programlama önem kazanmıştır çünkü programcılar, görevleri çekirdekler arasında bölüştürmek ve paralel yürütme sağlamak zorundadır.

Sonuç olarak, Go dilinde eş zamanlılık (concurrency) özelliği bulunması ve paralel yürütmenin önemi, modern bilgisayar sistemlerinde performans artışı sağlamak için gereklidir. Bu, güç tüketimi ve sıcaklık sorunları nedeniyle frekans artışının sınırlı olduğu bir ortamda özellikle önemli hale gelmiştir.