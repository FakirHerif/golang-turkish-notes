# Synchronized Communication

Kanallar, farklı gorutineler arasında veri iletişimi ve senkronizasyon sağlamak için kullanılan bir mekanizmadır.

## Kanallarda Bloklama (Blocking on Channels):

Gorutineler arasında veri alışverişi için kanalların yaygın bir kullanımı, kanal üzerinden iteratif bir şekilde veri okuma işlemidir. Bu durum genellikle bir üretici ve tüketici senaryosunda gerçekleşir. Tüketici, kanaldan sürekli olarak veri almak ve ardından bu veriyi işlemek ister. Bu durumu gerçekleştirmek için for döngüsü ve range keyword'ü kullanılır. for i := range c ifadesi, kanaldan sürekli olarak veri alır ve bu veriyi işler.

Örneğin:

```
for i := range c {
    // Veriyi işle (örneğin, yazdır)
    fmt.Println(i)
}
```

Bu döngü, kanaldan veri geldikçe devam eder ve kanal kapatıldığında sonlanır.

## Kanal Kapatma:

Kanalın kapatılması, kanalın gönderim tarafından kapatılması anlamına gelir. Bu, kanal üzerinden veri alacak olan gorutinenin, kanalın artık veri gönderilmeyeceğini ve döngünün sonlanması gerektiğini bilmesini sağlar. 

Kapatma işlemi şu şekilde yapılır:

```
close(c)
```

## Kanal Kullanımında Kapatmanın Zorunluluğu:

Kanalın kapatılması, özellikle range yapısıyla sürekli kanaldan veri okuyan senaryolarda önemlidir. Eğer bir gorutine, artık veri gönderilmeyeceğini biliyorsa, kanalı kapatması gerekmektedir. Bu durum, range ile sürekli bir döngüde beklemekte olan gorutinenin kanal kapatılmadan çıkmasını sağlar.

## Birden Fazla Kanaldan Okuma (Select):

Birden fazla kanaldan veri okuma durumu da ele alınmıştır. Örneğin, bir tüketici gorutini, birden fazla kanaldan veri almaya çalışabilir. Bu durumda, eğer her iki kanaldan da veri gerekiyorsa, sırayla her iki kanaldan da okuma yapılır. Eğer sadece bir kanaldan veri gerekiyorsa ve hangi kanaldan geleceği önceden bilinmiyorsa, select ifadesi kullanılır.

## Select İfadesi:

select, birden fazla iletişim operasyonu arasında beklemek için kullanılır. Örneğin:

```
select {
case a := <-c1:
    // c1'den veri geldi, a'ya atandı
case b := <-c2:
    // c2'den veri geldi, b'ye atandı
}
```

Bu ifade, ilk veri gelen kanalın işlenmesini sağlar. Eğer hem c1 hem de c2'ye veri gönderilmişse, hangisi önce gelirse o çalışır.

## Select İle Gönderme ve Alma:

select ifadesi, bir veya birden fazla kanaldan veri almayı (receive) veya veri göndermeyi (send) beklerken kullanılabilir.

```
select {
case a := <-inChan:
    // inChan kanalından veri alındı, a değişkenine atanacak
    fmt.Println("Received:", a)
case outChan <- b:
    // b değeri outChan kanalına gönderildi
    fmt.Println("Sent:", b)
}
```

Bu örnekte, select ifadesi içinde iki farklı case bulunmaktadır. İlk case, inChan kanalından veri almayı temsil eder ve bu veri a değişkenine atanır. İkinci case ise outChan kanalına veri göndermeyi temsil eder ve b değeri bu kanala gönderilir.

select ifadesi, her iki durumu da bekler ve hangi case'in hazır olduğuna bağlı olarak uygun olanı çalıştırır. Eğer inChan kanalından veri gelirse, ilk case çalışır ve alınan veriyi işler. Eğer outChan kanalına veri gönderilebiliyorsa, ikinci case çalışır ve b değeri outChan kanalına gönderilir.

Bu yapı, farklı gorutineler arasında veri iletimini dinamik bir şekilde yönetmek için kullanılabilir.

## Select ile Abort Kanalı Kullanımı:

## Default Durumu:

select ifadesi, ayrı bir iptal (abort) kanalını kullanmak için de yaygın olarak kullanılır. Örneğin, bir görevi sürekli olarak gerçekleştiren bir tüketici gorutini düşünelim. Bu tüketici, bir kanaldan gelen veriyi alır ve işler. Ancak, belirli bir noktada görevin iptal edilmesi gerekebilir.

```
for {
    select {
    case data := <-c:
        // Kanaldan veri al ve işle
        fmt.Println(data)
    case <-abort:
        // Abort kanalına veri gelirse, işlemi iptal et
        return
    }
}
```

Bu örnekte, select ifadesi içinde iki farklı case bulunmaktadır. Birincisi, normal veri alma ve işleme işlemidir. İkincisi ise bir abort (iptal etme) kanalıdır. Eğer abort kanalına bir değer gönderilirse, select içindeki ikinci case çalışacak ve tüketici gorutini işlemi iptal edip döngüden çıkacaktır.

Bu tasarım, genellikle görevin normal akışının yanı sıra, görevi iptal etmek için başka bir mekanizmanın kullanılmasını sağlar. Örneğin, kullanıcıdan gelen bir iptal talebi, abort kanalına bir değer göndererek tüketici gorutinin çalışmasını güvenli bir şekilde sonlandırabilir.

## Default Select

Select ifadesinde bir default durumu da kullanılabilir. Bu durumda, hiçbir kanaldan veri gelmediğinde çalışacak olan bir bloktur. 

Örneğin:

```
select {
case a := <-c1:
    // c1'den veri geldi, a'ya atandı
case b := <-c2:
    // c2'den veri geldi, b'ye atandı
default:
    // Hiçbir kanaldan veri gelmedi
}
```

Bu şekilde, birden fazla kanaldan veri almak veya veri göndermek için select ifadesi kullanılabilir. Bu yapılar, gorutineler arasında senkronizasyon ve veri iletişimi sağlayan güçlü bir mekanizmadır.