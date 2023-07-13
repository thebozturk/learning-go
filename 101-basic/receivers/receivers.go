package main
import (
	"fmt"
	"errors"
)

// receiverlar ile structlara fonksiyon tanımlayabiliriz !
type Ogrenci struct {
	Adi string
	Soyadi string
	Yasi int
}

// receiver tanımlamak için func keywordünden sonra fonksiyon isminden önce parantez içerisinde receiver'ın ismi ve tipi yazılır
func (ogrenci Ogrenci) AdiSoyadi() string {
	return ogrenci.Adi + " " + ogrenci.Soyadi
}

// eğer receiver'ın tipini pointer olarak tanımlarsak struct içerisindeki değişkenleri değiştirebiliriz
func (ogrenci *Ogrenci) YasiArttir() {
	if ogrenci.Yasi < 18 {
		panic(errors.New("ogrenci 18 yaşından küçük"))
	}
}

func main() {
	ogrenci := Ogrenci{
		Adi: "bugrahan",
		Soyadi: "ozturk",
		Yasi: 22,
	}

	fmt.Println(ogrenci.AdiSoyadi())

	ogrenci.YasiArttir()

	fmt.Println(ogrenci.Yasi)
}
