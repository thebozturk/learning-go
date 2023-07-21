package main

import "fmt"

func main() {
	//cities := []string{"New York", "Paris", "Berlin", "Madrid"}

	//for i := 0; i < len(cities); i++ {
	//	fmt.Println(cities[i])
	//}

	// for -range
	//for _, c := range cities {
	//	fmt.Println(c)
	//}

	// arrayler slice'lara göre daha hızlıdır. arraylerin kopyasını alabiliriz ama slice'ların kopyasını alamayız. referans tip olduğu için.
	// arrayler pass by value, slice'lar pass by reference
	// underlying array, slice'ın arkasındaki array demek. slice'ın kapasitesi array'in kapasitesiyle aynıdır.

	//underArr := [...]int{1, 2, 3, 4, 5}
	//fmt.Println(underArr)
	//
	//mySlc := underArr[0:3]
	//fmt.Println(mySlc)
	//
	//mySlc2 := underArr[2:5]
	//fmt.Println(mySlc2)
	//
	//mySlc3 := underArr[:2]
	//fmt.Println(mySlc3)
	//
	//mySlc4 := underArr[3:]
	//fmt.Println(mySlc4)

	// slice'lar array'lerin üzerinde çalışır. slice'lar array'lerin bir parçasıdır.
	//mySlc5 := underArr[:]
	//fmt.Println(mySlc5)

	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice)

	mySlice = append(mySlice, 6, 7, 8)
	fmt.Println(mySlice)

	mySlc2 := []int{9, 10, 11}
	mySlice = append(mySlice, mySlc2...)
	fmt.Println(mySlice)

	// ilk 2 elemanı siler
	mySlice = mySlice[2:]
	fmt.Println(mySlice)

	// son 2 elemanı siler
	mySlice = mySlice[:len(mySlice)-2]
	fmt.Println(mySlice)
}
