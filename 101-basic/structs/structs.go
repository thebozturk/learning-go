package main
import (
	"fmt"
	"encoding/json"
)

// structs ile birlikte go dilinde class yapısı oluşturabiliriz !

// struct tanımlamak için type ve struct keywordleri kullanılır
// struct içerisindeki değişkenlerin isimleri büyük harfle başlarsa public olur
// struct içerisindeki değişkenlerin isimleri küçük harfle başlarsa private olur

type Ogrenci struct {
	Adi string
	Soyadi string
	Yasi int

	// struct içerisinde tag kullanılabilir ve tag'ler json formatında kullanılabilir
	// bu taglar ile json formatında veri alıp verebiliriz !
	username string `json:"username"`
}

func main() {
	ogrenci := Ogrenci{
		Adi: "bugrahan",
		Soyadi: "ozturk",
		Yasi: 22,
		username: "bugrahan61",
	}

	// struct içerisindeki değişkenlere nokta ile erişebiliriz
	fmt.Println(ogrenci)

	// eğer structın içerisinden json formatında veri almak istersek Marshal fonksiyonunu kullanırız
	ogrenci_json, _ := json.Marshal(ogrenci)

	fmt.Println(string(ogrenci_json))

	// eğer structın içerisine json formatında veri vermek istersek Unmarshal fonksiyonunu kullanırız
	ogrenci_2 := Ogrenci{}

	json.Unmarshal(ogrenci_json, &ogrenci_2)

	fmt.Println(ogrenci_2)
}