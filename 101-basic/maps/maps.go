package main

import "fmt"

// map ile key-value şeklinde veriler tutulur. sıra garantisi olmadan adreste referans edilir

var mp_1 map[string]int // string tipinde key ve int tipinde value tutan bir map

func main() {
	// map tanımlamak için make fonksiyonu kullanılır
	mp_2 := make(map[string]int) // string tipinde key ve int tipinde value tutan bir map

	// map'e de boyut belirtilebilir
	mp_3 := make(map[string]int, 5) // string tipinde key ve int tipinde value tutan 5 boyutlu bir map

	fmt.Println(mp_2, mp_3)

	// map'e eleman eklemek için key-value şeklinde eklenir
	mp_3["key_1"] = 1
	mp_3["key_2"] = 2
	mp_3["key_3"] = 3

	fmt.Println(mp_3)
}