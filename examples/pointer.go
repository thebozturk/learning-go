package main

import "fmt"

// pointer: başka bir değişkenin adresini tutar. pass by reference
func main() {
	name := "Bugrahan"
	fmt.Println(name)  // Value
	fmt.Println(&name) // & -> Address: name değişkeninin memory'deki adresi

	// pointerin veri tipi: asteriks(*) ile belirtilir
	// addressin veri tipi: ampersand(&) ile belirtilir

	// deference: pointerin adresindeki değeri almak
	fmt.Println(*&name) // name değişkeninin değerini alır

	//x1 := 5
	//x2 := x1
	//fmt.Println(x1, x2) // 5 5
	//x2 = 10
	//fmt.Println(x1, x2) // 5 10

	x1 := 5
	x2 := &x1
	fmt.Println(x1, x2) // 5 address
	*x2 = 10
	fmt.Println(x1, *x2) // 10 10

	arr := [3]int{1, 2, 3}
	arr2 := arr
	fmt.Println(arr, arr2) // [1 2 3] [1 2 3]

	// !!!! arraylerde pass by value çalışır ama slice'larda pass by reference çalışır
}
