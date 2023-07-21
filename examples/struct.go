package main

import "fmt"

type employee struct { // underlying type ve pass by value
	firstName, lastName string
	age, salary         int
}

type manager struct {
	employee  // employee stuct'ının tüm özelliklerini manager struct'ına dahil ettik.
	hasDegree bool
}

func main() {
	employee1 := employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}

	employee2 := employee{"Thomas", "Paul", 29, 800}

	fmt.Println("Employee 1", employee1)
	fmt.Println("Employee 2", employee2)

	m1 := manager{
		employee: employee{
			firstName: "John",
			lastName:  "Doe",
			age:       32,
			salary:    600,
		},
		hasDegree: true,
	}

	fmt.Println("Manager 1", m1)

	// is relation nedir: manager is an employee
	// has relation nedir: manager has an employee

	// anonymous struct: bir struct'ı sadece bir yerde kullanacaksak. örneğin bir fonksiyonun içinde.
	emp3 := struct {
		firstName, lastName string
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
	}

	fmt.Println("Employee 3", emp3)
}
