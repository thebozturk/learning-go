package main

import (
	"fmt"
	"strings"
)

// go dilinde stringler char array olarak tutulur ve stringler immutable'dır (değiştirilemez)
var str_1 string = "string 1"

func main() {
	str_2 := "string 2"

	// stringler + operatörü ile birleştirilebilir
	fmt.Println(str_1 + " " + str_2)

	// stringler'de eğer belli bir indexe kadar almak istersek : operatörü kullanılır
	fmt.Println(str_1[:3])

	// eğer stringin toplam uzunluğunu almak istersek len fonksiyonu kullanılır
	fmt.Println(len(str_1))

	// stringin toplam uzunlugundan belli bir sayı kadar geri gelmek istersek - operatörü kullanılır
	fmt.Println(str_1[len(str_1)-3:])

	// go string formatlama için fmt paketindeki Sprintf fonksiyonu kullanılır
	str_3 := fmt.Sprintf("%s %s", str_1, str_2)

	fmt.Println(str_3)

	// strings paketindeki fonksiyonlar ile stringler üzerinde işlemler yapılabilir

	// strings paketindeki Contains fonksiyonu ile stringin içerisinde belli bir stringin olup olmadığını kontrol edebiliriz
	fmt.Println(strings.Contains(str_3, "string"))

	// strings paketindeki Count fonksiyonu ile stringin içerisinde belli bir stringin kaç defa geçtiğini bulabiliriz
	fmt.Println(strings.Count(str_3, "string"))

	// strings paketindeki HasPrefix fonksiyonu ile stringin belli bir string ile başlayıp başlamadığını kontrol edebiliriz
	fmt.Println(strings.HasPrefix(str_3, "string"))

	// strings paketindeki HasSuffix fonksiyonu ile stringin belli bir string ile bitip bitmediğini kontrol edebiliriz
	fmt.Println(strings.HasSuffix(str_3, "string"))

	// strings paketindeki Index fonksiyonu ile stringin belli bir stringin indexini bulabiliriz
	fmt.Println(strings.Index(str_3, "string"))

	// strings paketindeki Join fonksiyonu ile stringleri birleştirebiliriz
	fmt.Println(strings.Join([]string{"bugrahan", "ozturk"}, " "))
}

