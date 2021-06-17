package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	var michael person
	michael.firstName = "Michael"
	michael.lastName = "Jackson"
	fmt.Println(michael)
	fmt.Printf("%+v", michael)
}
