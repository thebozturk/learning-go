package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// bytes.Buffer ile string birleştirme işlemi yapılırken performans açısından daha iyi sonuçlar alınır.
	var x bytes.Buffer
	// 10 defa randomString fonksiyonu çağırılıp string birleştirme işlemi yapılıyor.
	for i := 0; i < 10; i++ {
		x.WriteString(randomString())
	}
	fmt.Println(x.String())

	// strings.Builder ile string birleştirme işlemi en performanslı şekilde yapılır.
	var y strings.Builder
	// 10 defa randomString fonksiyonu çağırılıp string birleştirme işlemi yapılıyor.
	for i := 0; i < 10; i++ {
		y.WriteString("Data")
	}
	fmt.Println(y.String())
}

func randomString() string {
	return "Hello World!"
}
