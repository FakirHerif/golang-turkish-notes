# First-Class Values / Fonksiyonların birinci sınıf değerler olarak ele alınması 

Bir fonksiyonu birinci sınıf bir değer olarak
ele almak, genellikle bir fonksiyonu; int, string veya diğer herhangi bir tür gibi ele alabilmek anlamına gelir. Diğer türler de yapabildiğimiz herşeyi bir fonksiyona da yapabiliriz. Örnek olarak değişkenler fonksiyon türü olarak bildirilebilir. Böylece bir değişkeni string, int gibi bildirebildiğimiz gibi fonksiyon olarak da bildirebiliriz.

**Fonksiyonların Değer Olarak Kullanılması:** Go'da fonksiyonlar diğer türler gibi kullanılabilir. Bir değişken, bir fonksiyon türüne atanabilir veya fonksiyonlar bir fonksiyonun parametresi olarak geçirilebilir, dönüş değeri olarak kullanılabilir veya veri yapıları içine yerleştirilebilir.

**Değişken Olarak Fonksiyonlar:** Bir programlama dilinde fonksiyonların diğer veri türleri gibi değişkenlere atanabilmesini ve fonksiyonları değişken olarak kullanabilme özelliğini ifade eder.

Örneğin, Go programlama dilinde bir fonksiyonu bir değişkene atayabilir ve sonra bu değişkeni kullanarak ilgili fonksiyonu çağırabilirsiniz. Fonksiyonlar bu şekilde değişkenlere atanabildiğinden, daha sonra bu değişkenler fonksiyon çağrıları için kullanılabilir.

```
var funcVar func(int) int       // Bir değişken tanımladık, funcVar isminde. Bu değişken, int tipinden bir parametre alan ve int tipinde bir değer döndüren fonksiyonları tutmak için kullanılacak.
    func incFn(x int) int {     // incFn isminde bir fonksiyon tanımladık. Bu fonksiyon, int tipinden bir x parametresi alır ve int tipinde bir değer döndürür.
        return x + 1
    }
func main() {
    funcVar = incFn             // incFn fonksiyonunu funcVar değişkenine atadık
    fmt.Print(funcVar(1))       // funcVar değişkeni üzerinden incFn fonksiyonunu çağırdık
}
```

**Fonksiyonları Argüman Olarak Kullanma:** Bir program içinde bir fonksiyonu başka bir fonksiyonun parametresi olarak iletebilme yeteneğini ifade eder. Bu, fonksiyonları diğer fonksiyonlara argüman olarak iletmek ve fonksiyonları çağıran fonksiyonlarda kullanmak anlamına gelir. Bu özellik, genellikle işlevsel programlama dillerinde kullanılan ve programın esnekliğini artıran bir yapıdır.

```
func ApplyIt(afunc func(int) int, val int) int {
    return afunc(val)
}

func Increment(x int) int {
    return x + 1
}

func main() {
    result := ApplyIt(Increment, 5)
    fmt.Println(result) // Çıktı: 6
}
```

Yukarıdaki örnekte, "ApplyIt" adlı bir fonksiyon tanımladık. Bu fonksiyon, bir tamsayı fonksiyonu (int alan ve int döndüren) ve bir tamsayı değeri alır. Ardından, "Increment" adlı başka bir fonksiyon tanımladık, bu fonksiyon bir tamsayıyı 1 arttırır. main fonksiyonunda "ApplyIt" fonksiyonunu kullanarak "Increment" fonksiyonunu çağırdık ve 5 değerini artırdık. Bu örnekte, "Increment" fonksiyonunu "ApplyIt" fonksiyonuna bir argüman olarak ilettik ve "ApplyIt" fonksiyonu, verilen fonksiyonu ve değeri çağırdı ve sonucu ekrana yazdırdık.

**Anonim Fonksiyonlar:** Adları olmayan fonksiyonlardır. Yani, bir isme ihtiyaç duymadan doğrudan bir değişkene atanabilir veya başka bir fonksiyonun argümanı olarak kullanılabilirler. Anonim fonksiyonlar, genellikle tek seferlik kullanımlar için tanımlanır ve genellikle başka bir fonksiyon veya algoritma içinde kullanılmak üzere oluşturulurlar.

Go dilinde anonim fonksiyonlar **lambda ifadeleri** olarak da adlandırılabilir. Fonksiyonları doğrudan ifade etmek için kullanılırlar ve bu fonksiyonlar genellikle bir değişkene atanmadan kullanılır.

```
func main() {
    // Anonim bir işlevi funcVar değişkenine atama
    funcVar := func(x int) int {
        return x + 1
    }

    result := funcVar(5)    // funcVar değişkeni ile anonim işlevi çağırma
    fmt.Println(result)     // Çıktı: 6
}
```

Yukarıdaki örnekte, funcVar adında bir değişken tanımlıyoruz ve bu değişken anonim bir işlevi temsil ediyor. funcVar değişkenine atanmış olan anonim işlev, bir tamsayı alıp bu sayıya bir ekleyerek sonucu döndürüyor.

Anonim işlevler genellikle, sadece bir kez kullanılmak üzere bir değişkene atanır veya bir işlevin içinde kullanılır. Kod içinde daha az yer kaplarlar ve ihtiyaç duyulduğunda kullanılabilirler. Bu sayede, kodun okunabilirliğini artırabilir ve programcıya esneklik sağlayabilirler.