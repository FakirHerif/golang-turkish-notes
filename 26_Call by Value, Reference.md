# Call by Value, Reference

**Call by Value" (Değer ile Çağrı)**, bir fonksiyona parametre olarak değişken gönderildiğinde, o değişkenin bir kopyasının fonksiyona iletilmesi anlamına gelir. Yani, fonksiyon içinde bu parametrenin değerinde yapılan değişiklikler, orijinal değişkeni etkilemez. Gönderilen değişkenin kendisi değil, bir kopyası fonksiyona aktarılır.

Örnek olarak, Go dilinde "Call by Value" davranışını gösteren bir örnek verelim:

```
// increase fonksiyonu, bir sayıyı artırır
func increase(num int) {
    num = num + 1           // Parametre olarak alınan num değerini 1 artırır
}

func main() {
    x := 5

    // increase fonksiyonu çağrıldığında, x değişkeninin bir kopyası oluşturulur ve bu kopya fonksiyona geçirilir
    increase(x)

    fmt.Println("x:", x)    // Çıktı: x: 5 (Fonksiyon içindeki değişiklik orijinal x'i etkilemez)
}
```

Yukarıdaki örnekte, increase fonksiyonu num adında bir parametre alır. Bu fonksiyon, parametre olarak aldığı sayıyı bir artırır. main fonksiyonunda x adında bir değişken tanımlanır ve bu değişken increase fonksiyonuna gönderilir.

Ancak, fonksiyon içinde yapılan değişiklikler, fonksiyona gönderilen x değişkenini etkilemez. Çünkü fonksiyon num adında bir kopya aldı ve bu kopya üzerinde işlem yapıldı. Orijinal x değişkeni aynı kaldı ve fonksiyon içindeki değişikliklerden etkilenmedi.

Bu örnek, Go dilinde "Call by Value" prensibini göstermektedir: fonksiyona parametre olarak gönderilen değişkenlerin bir kopyası oluşturulur ve bu kopyalar üzerinde işlem yapılır, bu işlemler orijinal değişkenleri etkilemez.

"Call by Value" (Değer ile Çağrı) yönteminin birkaç avantajı vardır:

Güvenlik ve Beklentilerin Korunması: Değişkenlerin fonksiyonlara kopyalanarak iletilmesi, beklenmeyen değişikliklerin önlenmesine yardımcı olur. Bu sayede, bir fonksiyonun yan etki yaratma olasılığı azalır. Fonksiyonun, dışarıda bulunan değişkenleri doğrudan değiştirme riski olmadığından, kodun beklenmeyen sonuçlar üretme olasılığı düşer.

Daha İyi Anlaşılabilirlik: Değer ile çağrı, fonksiyonun davranışının daha tahmin edilebilir olmasını sağlar. Fonksiyonun içinde yapılan değişikliklerin, dışarıda bulunan değişkenleri etkilemeyeceği bilindiğinden, kodun okunabilirliği artar.

Paralel Programlama ve Eşzamanlılık: Değer ile çağrı, paralel programlama ve eşzamanlılık durumlarında daha güvenli bir yaklaşım sunabilir. Çünkü farklı iş parçacıkları veya eş zamanlı çalışan farklı süreçler arasında değişkenlerin paylaşılmasının kontrol edilmesi daha kolay olabilir.

Bu avantajlar, programların beklenmedik davranışlardan korunmasına, hata ayıklamanın kolaylaştırılmasına ve kodun daha güvenli hale getirilmesine yardımcı olur. Ancak bazı durumlarda, belirli bir değişkenin değerinin orijinal olarak değişmesi gereken durumlar olabilir. Bu tür durumlar için farklı bir yaklaşım gerekebilir. Bu nedenle, programcılar ihtiyaçlarına göre farklı yöntemleri kullanabilirler. Örneğin, Go dilinde pointerlar kullanarak fonksiyonlara referansları geçirmek, fonksiyon içinde değişikliklerin orijinal değişkeni etkilemesine olanak sağlar. Bu da programcılara istedikleri durumlarda değişkenlerin orijinal değerlerini değiştirme imkanı sunar. Bu esneklik, programcılara farklı senaryolara uygun çözümler üretme imkanı sağlar.

**Call by Reference (Referans ile Çağrı)**, bir fonksiyona parametre olarak değişkenin bellek adresini geçirerek çalışır. Yani, orijinal değişkenin kendisi değil, onun bellek adresi (referansı) fonksiyona geçirilir. Fonksiyon bu referans aracılığıyla doğrudan orijinal değişkenin bellek adresine erişir ve değişiklikler orada yapılır. Böylelikle, fonksiyon içinde yapılan değişiklikler doğrudan orijinal değişkeni etkiler.

Bu yöntem, bir değer kopyası yerine değişkenin kendisinin fonksiyona geçirilmesini sağlar. Call by Reference, genellikle dillerin bazılarında (örneğin C++, Python, vb.) kullanılabilir; fakat Go dilinde doğrudan call by reference mekanizması bulunmaz.

Go dilinde benzer bir davranış elde etmek için pointerlar kullanılır. Pointerlar, bir değişkenin bellek adresini tutan özel veri türleridir ve bu adresler üzerinden değişkenlere erişim sağlarlar. Bir fonksiyona bir değişkenin adresini geçirerek, bu değişkenin referansı üzerinden değişiklikler yapılabilir, benzer şekilde call by reference benzeri bir işlevsellik elde edilebilir.

Go dilinde call by reference benzeri işlevsellik elde etmek için bir örnek:

```
// increase fonksiyonu, bir pointer ile çağrılır ve parametre olarak bir adres alır
func increase(num *int) {
    *num = *num + 1         // Pointer aracılığıyla değişkenin değeri bir artırılır
}

func main() {
    x := 5

    // increase fonksiyonuna x değişkeninin adresi (pointer'ı) geçirilir
    increase(&x)

    fmt.Println("x:", x)    // Çıktı: x: 6 (Fonksiyon içindeki değişiklik orijinal x'i etkiler)
}
```

Yukarıdaki örnekte, increase fonksiyonu *int türünde bir parametre alır, yani bir pointer alır. increase fonksiyonu çağrılırken &x kullanılarak x değişkeninin adresi (pointer'ı) fonksiyona geçirilir. Bu sayede, fonksiyon içindeki *num üzerinden değişkenin değeri bir artırılır ve orijinal x değişkeni fonksiyon çağrıldıktan sonra artırılmış değere sahip olur. Bu durum, Go dilinde call by reference benzeri bir işlevsellik sağlar.

**"&"** işareti, bir değişkenin bellek adresini almak için kullanılır. Örneğin, &x ifadesi x değişkeninin bellek adresini verir. Bu, bir değişkenin adresini elde etmek için kullanılan bir işlemdir.

**"*"** işareti ise bir pointer'ın değerini dereference etmek veya işaret ettiği değere erişmek için kullanılır. Örneğin, *ptr ifadesi ptr adındaki pointer'ın işaret ettiği değeri döndürür. Bu, bir pointer'ın işaret ettiği bellek adresindeki değere erişmek için kullanılır.

Pointerlar, değişkenlerin bellek adreslerini tutarlar ve bu adresler üzerinden doğrudan değişkenlere erişim sağlarlar. Bu nedenle, pointerlar Go dilinde referans benzeri bir işlevsellik sağlarlar ve fonksiyonlara değişkenlerin adresleri (pointer'ları) geçirilerek call by reference benzeri bir davranış elde edilebilir.