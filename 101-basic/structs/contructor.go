package main

type Person struct {
	firstName string
	lastName  string
	age       int
}

// Pointer ile struct'ın adresini döndürür. Çünkü struct'ın kopyası oluşturulmaz ve bellekten tasarruf edilir.
func NewPerson(firstName string, lastName string, age int) *Person {
	return &Person{firstName, lastName, age}
}

func main() {
	/*
			x := Person{firstName: "John", lastName: "Doe", age: 25}
				println(x.firstName, x.lastName, x.age)

		x := new(Person)
			x.firstName = "John"
			println(x.firstName)
	*/

	// Golang'da constructor yoktur. Yapıcı fonksiyonlar kullanılır.
	// Yapıcı fonksiyonlar, struct'ın ilk harfi büyük olacak şekilde isimlendirilir.

	// Yapıcı fonksiyon kullanımı
	x := NewPerson("John", "Doe", 25)
	println(x.firstName, x.lastName, x.age)
}
