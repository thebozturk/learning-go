package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// os.Open() fonksiyonu dosya açmak için kullanılır.
	file, err := os.Open("test.txt")
	println(file)
	// eğer bir error nesnesi oluşturmak istiyorsak errors paketini kullanabiliriz.
	err = errors.New("Bu bir hata mesajıdır.")
	fmt.Println(err)

	checkError(err)
}

func checkError(err error) {
	if err != nil {
		errors.New("Bu bir hata mesajıdır.")
		os.Exit(1)
	}
}
