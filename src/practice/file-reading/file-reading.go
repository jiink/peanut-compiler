package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Specify the file path you want to read
	filePath := "file.txt"

	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Read the contents of the file
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Convert the file content to a []rune
	runes := []rune(string(content))

	// Print the content as runes
	fmt.Println("Content as runes:")
	fmt.Println(string(runes))

	// Get one rune
	fmt.Printf("First rune: %c\n", runes[0])
}
