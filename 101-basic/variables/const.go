package main

import "fmt"

const sabit_1 = "deger 1"

// eğer birden fazla sabit tanımlanacaksa () içerisinde tanımlanır
const (
	sabit_2 = "deger 2"
	sabit_3 = "deger 3"
)

// eğer değerler sıralı artan değerlerse iota adında bir enum var
// ilk değer 0'dan başlar ve her sabit için 1 artar
const (
	sabit_4 = iota // 0
	sabit_5
	sabit_6
)

func main() {
	name := "bugrahan"
	fmt.Println(sabit_1, sabit_2, sabit_3, sabit_4, sabit_5, sabit_6, name)
}
