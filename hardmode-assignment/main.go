package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := os.Args[1]
	data, err := os.Open(filename)
	if err != nil {
		fmt.Println("An error occurred while trying to read in a file. Error:", err)
	}
	io.Copy(os.Stdout, data)
}
