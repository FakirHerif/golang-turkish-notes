# JSON

SON (JavaScript Object Notation), veri alışverişi için hafif, metin tabanlı bir veri formatıdır. JSON, insanlar tarafından okunabilir ve yazılabilir bir yapıya sahip olduğu için, veri taşıma ve değiş tokuşu için yaygın olarak kullanılır.

JSON, anahtar-değer çiftlerinden oluşan bir veri formatıdır. Bir JSON verisi genellikle şu öğeleri içerir:

Nesneler (Objects): Anahtar-değer çiftlerinden oluşan bir nesne. Anahtarlar bir string, değerler ise herhangi bir JSON veri türü olabilir (string, sayı, boolean, dizi, başka bir nesne vb.).
Diziler (Arrays): Değerlerin bir listesi. Bu değerler de yine herhangi bir JSON veri türü olabilir.
Değerler: Sayılar, metinler, boolean (true/false), null veya diğer JSON nesneleri olabilir.

Örnek bir JSON verisi şu şekilde gösterilebilir:

```
{
  "ürün": "bilgisayar",
  "fiyat": 1500,
  "stoktaMı": true,
  "özellikler": {
    "işlemci": "Intel Core i7",
    "bellek": "16GB RAM",
    "depolama": ["SSD", "HDD"]
  },
  "garantiSüresi": null
}
```

JSON, verilerin tarafsız ve hafif bir şekilde taşınmasını ve farklı platformlar arasında kolayca paylaşılmasını sağlar. Web uygulamalarında, API'ler aracılığıyla veri alışverişi yapmak, bu verileri depolamak veya yapılandırmak için JSON sıkça tercih edilen bir formattır. JSON, JavaScript diline dayanan bir format olmasına rağmen, birçok programlama dilinde desteklenir ve kullanılabilir.

**Marshalling**

Go programlama dilinde "Marshalling", bir veri yapısını (örneğin, bir struct) JSON formatına dönüştürme işlemidir. Go'da JSON veri yapısına dönüştürme işlemi "encoding/json" paketi kullanılarak gerçekleştirilir.

Bir veri yapısını JSON formatına dönüştürmek için, Go'da genellikle json.Marshal() veya json.MarshalIndent() fonksiyonları kullanılır.

İşte basit bir örnek:

```
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string
	Age     int
	City    string
	Hobbies []string
}

func main() {
	person := Person{
		Name:    "Alice",
		Age:     30,
		City:    "New York",
		Hobbies: []string{"Reading", "Hiking", "Cooking"},
	}

	// JSON formatına dönüştürme (Marshalling)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("JSON dönüştürme hatası:", err)
		return
	}

	fmt.Println(string(jsonData))
}
```

Bu örnekte, Person adında bir struct tanımlanmış ve bu struct JSON formatına dönüştürülmüştür. json.Marshal() fonksiyonu, verilen Person nesnesini JSON formatına dönüştürerek jsonData değişkenine atar. fmt.Println(string(jsonData)) ifadesiyle JSON verisi ekrana yazdırılır.

Go'da JSON Marshalling işlemi, veri yapılarını JSON formatına dönüştürmek için kullanılır ve bu dönüşüm, genellikle web servislerinde veya veri depolama işlemlerinde veri taşıma veya paylaşma amacıyla kullanılır.

**Unmarshalling**

JSON Unmarshalling, JSON formatındaki verileri Go programlama dilinde kullanılabilir veri yapılarına dönüştürme işlemidir. Yani, JSON formatındaki verilerin bir Go veri yapısına (örneğin, struct) dönüştürülmesidir. Go'da bu işlem "encoding/json" paketi içerisinde yer alan json.Unmarshal() fonksiyonu kullanılarak gerçekleştirilir.

İşte basit bir örnek:

```
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	City    string   `json:"city"`
	Hobbies []string `json:"hobbies"`
}

func main() {
	jsonData := []byte(`{"name":"Alice","age":30,"city":"New York","hobbies":["Reading","Hiking","Cooking"]}`)

	// JSON verisini Go veri yapısına dönüştürme (Unmarshalling)
	var person Person
	err := json.Unmarshal(jsonData, &person)
	if err != nil {
		fmt.Println("JSON dönüştürme hatası:", err)
		return
	}

	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("City:", person.City)
	fmt.Println("Hobbies:", person.Hobbies)
}
```

Bu örnekte, JSON formatında bir veri olan jsonData değişkeni tanımlanmıştır. json.Unmarshal() fonksiyonu, jsonData içeriğini Person tipindeki bir değişkene dönüştürür. &person ifadesi, dönüştürülen JSON verisinin Person türündeki değişkene atanmasını sağlar.

JSON Unmarshalling, genellikle dış kaynaklardan (örneğin, bir API'den gelen veriler) alınan JSON formatındaki verilerin, Go programında kullanılabilir veri yapılarına dönüştürülmesi için kullanılır. Bu şekilde, Go programları JSON verilerini kolayca işleyebilir, depolayabilir veya kullanabilir.