package main

import "fmt"

func main() {
	// mapler pass by reference'dır.
	var myMap map[string]string
	if myMap == nil {
		myMap = make(map[string]string)
	}
	myMap["dog"] = "Samson"
	fmt.Println(myMap)

	// delete a key

	delete(myMap, "dog")
	fmt.Println(myMap)
}
