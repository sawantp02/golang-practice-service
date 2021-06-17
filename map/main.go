package main

import "fmt"

func main() {
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}
	print(colors)

	// var sample map[string]string
	sample := make(map[string]string)
	sample["white"] = "#ffffff"
	delete(sample, "white")
	fmt.Println(sample)
}

func print(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}