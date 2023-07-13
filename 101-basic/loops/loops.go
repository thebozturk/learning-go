package main
import "fmt"

// go dilinde fonksiyonlar birden fazla değer döndürebilir ve fonksiyonlar bir değişken gibi tanımlanabilir
func main() {
	slc_1 := []int{1, 2, 3, 4, 5}
	slc_2 := []int{}

	// range ile slice'ın elemanlarına erişebiliriz. forEach gibi.
	for i, v := range slc_1 {
		fmt.Println(i, v)
	}

	for i := 0; i < len(slc_1); i++ {
		fmt.Println(slc_1[i])
	}

	for i := 0; i < 10; i++ {
		slc_2 = append(slc_2, i)
	}

	// iki dizideki ortak elemanları bulmak için iç içe for döngüsü kullanabiliriz
	for _, v1 := range slc_1 {
		for _, v2 := range slc_2 {
			if v1 == v2 {
				fmt.Println(v1, " ", v2)
			}
		}
	}

	// go dilinde while döngüsü yoktur. for döngüsünü while gibi kullanabiliriz
	i := 0
	for i < 10 {
		fmt.Println("i->",i)
		i++
	}
}