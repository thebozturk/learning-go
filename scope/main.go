package main

import "fmt"

var packVar = "Pack Scope"

// package scope'a ait olan değişkenler tüm paket içerisinde kullanılabilir.
// func scope'a ait olan değişkenler sadece o func içerisinde kullanılabilir.
// package scope olan değişkenler sürekli olarak hafızada tutulur.
// func scope olan değişkenler ise fonksiyon çalıştırıldığında hafızaya alınır ve fonksiyon çalışması bittiğinde hafızadan silinir.

func main() {
	var funcVar = "Func Scope"

	if true {
		// bu blok içerisinde tanımlanan değişken sadece bu blok içerisinde kullanılabilir.
		var blockVar = "Block Scope"
		fmt.Println(blockVar)
	}

	fmt.Println(funcVar)
	fmt.Println(packVar)
}
