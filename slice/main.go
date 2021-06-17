package main

import "fmt"

func main() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	update(mySlice)
	fmt.Println(mySlice)
}

func update(s []string) {
	s[0] = "Bye"
}
