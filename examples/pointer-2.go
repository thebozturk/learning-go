package main

import "fmt"

// pointeri bir değeri oldugu yerde değiştirmek icin kullanırız. bu da bize performans kazandırır.

func main() {
	x := 5
	fmt.Println(x) // 5
	double(&x)
	fmt.Println(x) // 10
}

func double(number *int) {
	*number *= 2
	fmt.Println(*number)
}
