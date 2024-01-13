/*
1. iki rakam arasında toplama, çıkarma, çarpma işleminin yapıldığı bir fonksiyon yazınız.
2. kullanıcı tarafından girilen note göre geçtiniz veya kaldınız geri dönüşü yazdırınız.
3. 1 ile 100 arasındaki bir sayıyı tahmin etme uygulaması yazınız, toplam tahmin etme hakkı 10 olsun.
*/

/*
1-

package main

import "fmt"

func main() {

	x, y := 10, 4
	sum, dif, prod := math(x, y)
	fmt.Println("Toplam: ", sum)
	fmt.Println("Fark: ", dif)
	fmt.Println("Çarpım: ", prod)
}

func math(num1 int, num2 int) (sum int, dif int, prod int) {
	sum = num1 + num2
	dif = num1 - num2
	prod = num1 * num2

	return sum, dif, prod
}

veya aşağıdaki gibi de yazabiliriz:


func math(num1 int, num2 int) (int, int, int) {
	sum := num1 + num2
	dif := num1 - num2
	prod := num1 * num2

	return sum, dif, prod
}
*/

/*
2-

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Lütfen aldığınız notu giriniz: ")
	grade, _ := getGrade()

	var result string

	if grade >= 50 {
		result = "Geçtiniz"
	} else {
		result = "Kaldınız"
	}

	fmt.Println(result)
}

func getGrade() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	}

	return num, nil
}
*/

// 3-

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	target := numRand(1, 100)

	fmt.Println("1 ile 100 arasındaki sayıyı bulmaya çalışın")

	reader := bufio.NewReader(os.Stdin)

	for attempts := 0; attempts < 10; attempts++ {
		fmt.Println(10-attempts, "hakkınız kaldı")
		fmt.Print("Lütfen tahmininizi yapınız: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		}

		if num > target {
			fmt.Println("Tahmininiz daha büyük, daha küçük bir sayı giriniz")
		} else if num < target {
			fmt.Println("Tahmininiz daha küçük, daha büyük bir sayı giriniz")
		} else {
			fmt.Println("Tahmininiz Doğru! ", attempts, ". denemede buldunuz")
			break // tahmin doğru olduğu zaman if koşulundan çıkmak için break kullanıyoruz
		}
	}

}

func numRand(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
