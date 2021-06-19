package main

import (
	"os"
	"fmt"
	"io"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}