package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//---- Variables ------------------------------------------------------------------------
/////////////////////////////////////////////////////////////////////////////////////////

var inputFilePath = ""
var sourceCode = ""

//---- Functions ------------------------------------------------------------------------
/////////////////////////////////////////////////////////////////////////////////////////

/* ---- Helpers --------------------------------------- */

// Prints a debug message to the console if debugEnabled is true.
func logDebug(format string, args ...interface{}) {
	if debugEnabled {
		fmt.Printf("[DEBUG] "+format, args...)
	}
}

// Reads in the Rat23F source code from the given path and
// stores it in the global variable `sourceCode`.
func readInSourceCode(path string) {
	sourceCodePath := path
	file, err := os.Open(sourceCodePath)
	check(err)
	defer file.Close()

	// To make things simple we'll just put it all in a string.
	content, err := io.ReadAll(file)
	check(err)
	sourceCode = string(content)
	logDebug("Source code: %s\n", sourceCode)
	sourceCode = removeComments(sourceCode)
	sourceCode = trimWhiteSpace(sourceCode)
	fmt.Printf("Opened file: %s\n", sourceCodePath)
}

/* ---- Main ------------------------------------------ */

func main() {
	if len(os.Args) < 2 {
		fmt.Println(">>>> Please provide the path to a Rat23F source code file as an argument.")
		fmt.Print("Press 'Enter' to quit...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		return
	}
	inputFilePath = os.Args[1]
	if len(inputFilePath) < 1 {
		fmt.Println(">>>> Please provide the path to a Rat23F source code file as an argument.")
		fmt.Print("Press 'Enter' to quit...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		return
	}

	if len(os.Args) > 2 {
		if os.Args[2] == "-d" {
			debugEnabled = true
		}
	}
	fmt.Println("Welcome to the Peanut Compiler for Rat23F!")
	readInSourceCode(inputFilePath)

	fmt.Println("Beginning lexical analysis...")
	var records, err = lexer(sourceCode)
	if err != nil {
		fmt.Println("The lexer encountered an error.")
	} else {
		logRecords(records)
	}
	fmt.Println("Lexical analysis complete.")

	fmt.Println("Beginning syntax analysis...")
	err = syntaxAnalyzer(records)
	if err != nil {
		fmt.Println("The syntax analyzer encountered an error.")
	}
	fmt.Println("Syntax analysis complete.")

	fmt.Println("This is the end of the compiler so far.")
}
