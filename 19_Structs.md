# Structs

Go programlama dilinde, struct bir veri yapısıdır ve farklı tipte verileri bir arada tutmak için kullanılır. struct, farklı veri türlerinin (int, string, float vb.) birleştirilmesini sağlar. Temel olarak, struct'lar, farklı özelliklere sahip alanları (fields) bir araya getirerek bir veri yapısı oluşturur. Örneğin:

```
type Person struct {
    Name    string
    Age     int
    Address string
}
```

Bu örnekte, Person adında bir struct oluşturduk. Name, Age ve Address olmak üzere üç farklı alan (field) içerir. Bu alanlar, Person yapısının özelliklerini tanımlar. Her bir alanın kendine özgü bir veri türü bulunabilir.

**Struct Oluşturulması**

```
// Yeni bir Person struct örneği oluşturma
var person1 Person
person1.Name = "Alice"
person1.Age = 30
person1.Address = "123 Main St"

// veya kısa bir şekilde
person2 := Person{Name: "Bob", Age: 25, Address: "456 Elm St"}

// Oluşturulan struct örneklerini kullanma
fmt.Println("Person 1:", person1.Name, person1.Age, person1.Address)
fmt.Println("Person 2:", person2.Name, person2.Age, person2.Address)
```

Bu örneklerde, Person struct'ını kullanarak person1 ve person2 adında iki farklı struct örneği oluşturuldu. Her bir örnek, Person yapısının alanlarını kullanarak bir kişinin adını, yaşını ve adresini temsil eder. Bu alanlar, struct içinde farklı veri türleri ile tanımlanabilir ve bu şekilde özelleştirilebilir.

struct'lar, benzer özelliklere sahip verileri gruplamak ve bir veri yapısı olarak kullanmak için oldukça kullanışlıdır. Bu sayede daha karmaşık veri yapıları oluşturabilir ve verileri düzenli bir şekilde yönetebilirsiniz.

**Struct başlatmanın diğer yolları**

Go'da struct'ları başlatmanın birkaç yolu vardır. new() fonksiyonu ve struct literal kullanımı, struct'ları başlatmak için kullanılan iki yaygın yöntemdir.

- **Struct Literal Kullanımı:**

```
type Person struct {
    Name    string
    Age     int
    Address string
}

// Struct literal ile başlatma
person := Person{
    Name:    "Alice",
    Age:     30,
    Address: "123 Main St",
}
```

Burada, Person adında bir struct tanımlandı ve person adlı bir örnek oluşturuldu. Bu örnek, struct alanlarını başlatmak için Person yapısının alan isimlerini ve değerlerini kullandı.

- **new() Fonksiyonu ile Struct Başlatma**

new() fonksiyonu, bir struct'ın bellekte alan ayırmak ve bu alanı sıfır değerleriyle doldurmak için kullanılır. new() fonksiyonu bir adres (pointer) döndürür ve bu adres, oluşturulan yeni struct örneğinin yerini temsil eder.

```
personPtr := new(Person)
```

Bu örnekte, new() fonksiyonu Person tipinde bir struct oluşturur ve bu struct'ın bir pointer'ını (*Person) döndürür. personPtr, oluşturulan struct'ın adresini içerir.

new() fonksiyonu ile oluşturulan struct örneği, struct literal ile oluşturulan örnekle aynı değerlere sahip olur, fakat new() kullanıldığında bu değerler varsayılan (sıfır) değerlerle başlatılır.

```
personPtr := Person(Name: "Joe", Age: "25", Address: "456 Main St")
```

şeklinde boş olarak gelen değerleri doldurabiliriz.