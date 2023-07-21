package main

import "fmt"

// pointerlar ile birlikte go dilinde call by reference yapısı kullanılabilir
// mesela normal bir string asla nil olamaz ama pointer ile tanımlanmış bir string nil olabilir

func main() {
	// herhangi bir değişkeni pointer olarak tanımlamak için * operatörü kullanılır
	type String *string

	// eğer değer vermezsek pointer'ın değeri nil olur
	var str_1 String

	fmt.Println(str_1)

	// eğer değer verirsek pointer'ın değeri verdiğimiz değer olur
	str_2 := "string 2"

	fmt.Println(str_2)

	// pointeri almak için & operatörü kullanılır
	ref_1 := &str_2

	fmt.Println(ref_1)

	// pointerin değerini almak için * operatörü kullanılır
	fmt.Println(*ref_1)
}
