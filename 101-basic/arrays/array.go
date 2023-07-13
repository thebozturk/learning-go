package main

import "fmt"

// eğer dizinin boyutunu belirtmek istersek
var arr_1 [5]int // 5 elemanlı int tipinde bir dizi
var arr_2 = [5]int{1, 2, 3, 4, 5} // 5 elemanlı int tipinde bir dizi ve ilk değerlerini belirttik

func main() {
	// eğer fonksiyonun için içerisinde tanımlanacaksa make fonksiyonu ile tanımlanır
	arr_3 := make([]int, 3) // 5 elemanlı int tipinde bir dizi

	// arrayın 1. indexine 1 değerini atadık
	arr_3[1] = 1

	fmt.Println(arr_1, arr_2, arr_3)

	// len ile arrayın boyutunu alabiliriz
	fmt.Println(len(arr_1), len(arr_2), len(arr_3))
}