package main
import "fmt"

// interface'ler go dilinde birbirinden farklı tipteki verileri tutmak için kullanılır

type Ogrenci interface {
	AdiSoyadi() string
	YasiArttir()
}

type ogrenci struct {
	Adi string
	Soyadi string
	Yasi int
}

// interface'lerin içerisinde fonksiyon tanımlarken fonksiyonun ismi, parametreleri ve dönüş tipi yazılır
func (ogrenci ogrenci) AdiSoyadi() string {
	return ogrenci.Adi + " " + ogrenci.Soyadi
}

func (ogrenci *ogrenci) YasiArttir() {
	if ogrenci.Yasi < 18 {
		panic("ogrenci 18 yaşından küçük")
	}
	ogrenci.Yasi++
}
