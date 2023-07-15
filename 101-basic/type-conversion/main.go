package main

import "strconv"

func main() {
	var myString = "1"
	var x = 10

	println(myString, x)

	// string to int(strconv.Atoi)
	// _ orada hata var ama ben hatayı kullanmak istemiyorum demek
	newNumber, _ := strconv.Atoi(myString)
	println(newNumber)

	// int to string(strconv.Itoa)
	newString := strconv.Itoa(x)
	println(newString)
}
