# Basit bir hello, world yazımı

```
package main

import "fmt"

func main() {
	fmt.Println("Hello, World")
}
```

**1- source kodumuzu compile etmek için:**
go build main.go
yazıyoruz. Ardından "main.exe" adlı pcmizin okuyabileceği bir exe dosyası oluşturmuş olduk. terminalden "main" veya "main.exe" yazdığımız zaman "Hello, World" çıktımızı verecektir.

**2- Konsolda çıktıyı almak için:**
go run main.go
yazıyoruz. ve konsolda "Hello, World" çıktısına ulaşıyoruz

