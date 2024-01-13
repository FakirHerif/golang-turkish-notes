/* package main

import (
	"errors"
	"fmt"
)

func main() {

	result, err := evenNum(11)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Girdiğiniz sayı: ", result)
	}

}

func evenNum(num int) (int, error) {
	if num%2 != 0 {
		return 0, errors.New("Hata: Çift sayı girmediniz.")
	}
	return num, nil
} */

// nil: kısaca başlangıç değeri olmayan ifadelere verilen değerdir.

package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	result, err := sRoot(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}

func sRoot(num float64) (float64, error) { // error interface olan veri tipidir
	if num < 0 {
		return 0, errors.New("Negatif sayıların karekökü alınamaz")
	}
	return math.Sqrt(num), nil
}

// EK NOT: "os .Open" methodu 2 değer döner bu yüzden bu değerlerden ikisinide yakalamamız gerekir, sadece bir tanesini yakalamaya çalışırsak hata alırız.
