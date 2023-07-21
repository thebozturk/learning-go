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
	result, err := evenNum(0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func checkError(err error) {
	if err != nil {
		errors.New("Bu bir hata mesajıdır.")
		os.Exit(1)
	}
}

func evenNum(num int) (int, error) {
	fmt.Println("num->", num)
	if num%2 != 0 {
		return num, errors.New("Bu bir hata mesajıdır.")
	}
	return num, nil
}
