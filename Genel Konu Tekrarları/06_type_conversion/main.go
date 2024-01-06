/*

--- BASIC TYPES (BUILT-IN TYPES)
NUMERIC TYPES
uint8       the set of all unsigned  8-bit integers (0 to 255) = 256
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127) = 256
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32


STRING TYPES
A string type represents the set of string values. A string value is a (possibly empty) sequence of bytes.
The number of bytes is called the length of the string and is never negative.

BOOLEAN TYPES
A boolean type represents the set of Boolean truth values denoted by the predeclared constants true and false.


--- COMPOSITE TYPES
Slice - Struct - Pointer - Function - Interface - Map - Channel - Array
*/

package main

import "fmt"

func main() {

	x := 10
	y := 10.0

	fmt.Printf("%v %T\n", x, x) // "%v" x değişkeninin veri tipini belirtiyor, "%T" ise x değişkeninin aldığı değeri belirtiyor. "\n" ise terminalde yeni bir satırda gösterim sağlıyor

	fmt.Printf("%v %T\n", y, y)

	/* fmt.Println(x + y) */ // type mismatch hatası alırız. çünkü int ile float64 toplamaya çalışıyoruz bunu yapamayız. Birbiriyle farklı veri tipine sahip değişkenleri birlikte herhangi bir işleme koyamayız

	// TYPE CONVERSION - type(value)

	fmt.Println(x + int(y))     // y değişkenimizi int'e dönüştürdük
	fmt.Println(float64(x) + y) // x değişkenini float64'e dönüştürdük

	// NOT: Type conversion'da asıl veri tipi değişmez. sadece işlem yaptığımız yere etki eder.

	// NOT2: Daha küçük kapsama sahip veri tipini daha geniş kapsama sahip veri tipine dönüştürmek daha sağlıklıdır.

	var z int16 = 128
	var w int8

	w = int8(z)

	fmt.Println(w) // z yi int8e dönüştürük ama int8'in alabileceği en yüksek değer 127 olduğu için sınırı geçti ve sonuç olarak bize -128 olan en küçük değeri döndürdü bu yüzden daha geniş kapsama sahip veri tipine dönüştürmek daha mantıklıdır. yani w değerini int16'ya dönüştürmeliyiz.

	var a = 10
	var b = "10"

	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)

	/* fmt.Println(a + int(b)) */ // verilerin de veri tiplerinin de birbirine dönüşüyor olması gerekiyor. type conversion yöntemiyle string'i int'e dönüştüremeyiz.

	num1 := 106
	str1 := string(num1)

	fmt.Printf("%v %T\n", num1, num1)
	fmt.Printf("%v %T\n", str1, str1) // az önce dönüştüremiyoruz dedik ama burada durum farklı 106 ya denk gelen ASCII karakterine dönüştürmüş oluyoruz ve sonucu "j string" olarak alıyoruz

}
