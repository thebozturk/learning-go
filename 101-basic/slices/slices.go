package main

import "fmt"

// go dilinde arraylar boyutludur. eğer boyutu belirtmek istemiyorsak slice kullanırız
// slice büyüyebilir
var slc_1 [10]int // int tipinde slice

func main() {
	fmt.Println(slc_1)

	// go da bir slice'ın kapasitesini öğrenmek için cap fonksiyonu kullanılır
	fmt.Println(cap(slc_1))

	// fonksiyonda slice tanımlamak için make fonksiyonu kullanılır
	slc_2 := make([]int, 5) // 5 elemanlı int tipinde bir slice

	fmt.Println(slc_2)

	// make ile olustururken 2 parametre verilebilir. 2. parametre slice'ın kapasitesini belirtir
	slc_3 := make([]int, 5, 10) // 5 elemanlı 10 kapasiteli int tipinde bir slice

	fmt.Println(slc_3)

	// append fonksiyonu ile slice'a eleman ekleyebiliriz
	slc_3 = append(slc_3, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	fmt.Println(slc_3)

	// her append işleminde kapasiteyi aşarsak kapasiteyi 2 katına çıkarır
}