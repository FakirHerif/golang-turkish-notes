/* package main

import (
	"fmt"
	"math"
)

type circle struct {
	r float64
}

func (c circle) display() {
	fmt.Println("A Circle")
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c1 := circle{5} // yarı çapı 5 olan çember. yarıçapına 5 tanımladık

	//	area1 := go c1.area()	// Go'da routine oluşturduğumuz fonksiyon herhangi bir return değer alamaz. Alırsa routine oluştururken hata alırız
	area1 := c1.area()
	fmt.Printf("%.2f\n", area1)

	go c1.display()
}
*/

/* package main

func merhaba() string {
	return "Merhaba"	// return değeri var bu yüzden goroutine tanımlamaya çalışırsak hata alırız
}

func main() {
	fmt.Println (go merhaba())	// return değeri aldığı için burada goroutine tanımladığımız zaman hata alırız. return yerine değerimizi channel'a göndeririz
} */

/* package main

import "fmt"

func merhaba(chan1 chan string) { // parametresi channel olan bir fonksiyon oluşturduk
	chan1 <- "Merhaba" // burada string değerimizi channel'ın içine göndermiş oluyoruz

}

func main() {

	// CHANNEL OLUŞTURMAK (UZUN YÖNTEM):

	// var myChannel chan string // channel tanımlamak için "chan" anahtar kelimesini kullanırız
	// myChannel = make(chan string)


	// CHANNEL OLUŞTURMAK (KISA YÖNTEM):

	myChannel := make(chan string)

	go merhaba(myChannel)

	fmt.Println(myChannel)   // 0xc0000200c0 -- değerimizin adresini görüyoruz, değerimizin kendisini görmek istiyoruz:
	fmt.Println(<-myChannel) // Merhaba
}
*/

package main

import (
	"fmt"
	"math"
)

type circle struct {
	r float64
}

func (c circle) display() {
	fmt.Println("A Circle")
}

func area(c circle, myChannel chan float64) {
	result := math.Pi * c.r * c.r
	myChannel <- result
}

func main() {
	c1 := circle{5}
	chan1 := make(chan float64)

	go area(c1, chan1)
	fmt.Printf("%.2f\n", <-chan1) // 78.54

	go c1.display()
}
