# Hash Tables

Hash tabloları, key-value (anahtar-değer) çiftlerini depolayan ve bir anahtara karşılık gelen değeri hızlı bir şekilde almak için kullanılan veri yapılarıdır. Büyük veri kütlelerine hızlı erişmemizi sağlar.

Örneğin, GPS koordinatları ve adres bilgileri. Yabi her adresin kendisiyle ilişkili benzersiz bir GPS koordinatı vardır. Yani GPS koordinatları key(anahtar) olarak kullanılır ve adres ise bir value(değer) olur. Kısaca key benzersizdir.

Örneğin:

```
Key         Value
Joe           x
Jane          y
Pat           z
```

Bu örnekteki key'lerin hepsi benzersizdir. Yani Jone, Jane ve Pat. ve bazı değeleri vardır; x, y ve z. Yani Joe x ile ilişkili, Jane y ile ilişkili, Pat z ile ilişkili.