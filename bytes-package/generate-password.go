package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().UnixNano())

const _charsetLowercase = "abcdefghijklmnpqrstuvwxyz"
const _charsetUppercase = "ABCDEFGHIJKLMNPQRSTUVWXYZ"
const _charsetNumeric = "123456789"
const _specialChars = "!@#$%&*_+;',.?"

type Option struct {
	length    int
	upperCase bool
	lowerCase bool
	numeric   bool
	specials  bool
}

func main() {
	// yeni bir password oluşturmak için generatePassword fonksiyonu çağırılıyor.
	passord, err := generatePassword(Option{
		length:    10,
		upperCase: true,
		lowerCase: true,
		numeric:   true,
		specials:  true,
	})
	if err != nil {
		panic(err)
	}
	// oluşturulan password ekrana yazdırılıyor.
	fmt.Println(passord)
}

func generatePassword(opt Option) (string, error) {
	x := make([]byte, opt.length)
	charset := "."
	if opt.lowerCase {
		charset += _charsetLowercase
	}
	if opt.upperCase {
		charset += _charsetUppercase
	}
	if opt.numeric {
		charset += _charsetNumeric
	}
	if opt.specials {
		charset += _specialChars
	}
	for i := range x {
		x[i] = charset[source.Int63()%int64(len(charset))]
	}
	if charset == "." {
		return "Non generate password.", errors.New("You have to choose at least one option.")
	}
	return string(x), nil
}
