package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Environ() ile sistemde tanımlı tüm değişkenleri alabiliriz.
	env := os.Environ()
	for _, v := range env {
		fmt.Println(v)
	}

	// os.Getenv() ile belirli bir değişkenin değerini alabiliriz.
	fmt.Println(os.Getenv("PATH"))
	fmt.Println(os.Getenv("HOME"))
}
