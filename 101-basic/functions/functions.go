package main
import "fmt"

// go dilinde fonksiyonlar birden fazla değer döndürebilir ve fonksiyonlar bir değişken gibi tanımlanabilir
func main() {
	// defer keywordü ile fonksiyonun sonunda çalıştırılacak kod bloğu tanımlanabilir
	defer fmt.Println("defer çalıştı")
	// recover fonksiyonu ile panic fonksiyonu ile durdurulan programı tekrar çalıştırabiliriz
	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	sayi_1 := 10
	sayi_2 := 20

	toplam := topla(sayi_1, sayi_2)

	fmt.Println(toplam)

	// panic fonksiyonu ile programı istediğimiz yerde durdurabiliriz
	if toplam > 10 {
		panic("toplam 10'dan büyük")
	}
}

func topla(a int, b int) int {
	return a + b
}

