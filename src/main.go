package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//---- Definitions ----------------------------------------------------------------------
/////////////////////////////////////////////////////////////////////////////////////////

var debugLogArgument = "debugLog"
var debugLogArgumentShort = "d"

//---- Variables ------------------------------------------------------------------------
/////////////////////////////////////////////////////////////////////////////////////////

var inputFilePath = ""
var sourceCode = ""

//---- Functions ------------------------------------------------------------------------
/////////////////////////////////////////////////////////////////////////////////////////

/* ---- Helpers --------------------------------------- */

// Terminates the program.
func exit() {
	fmt.Println("Exiting...")
	os.Exit(1)
}

func promptExit() {
	fmt.Print("Press 'Enter' to quit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	exit()
}

// Tells the user how to use this program.
// Show this when the user isn't using the program correctly.
func showUsage() {
	fmt.Println(">>>> Please provide the path to a Rat23F source code file as an argument.")
	fmt.Println("Usage: peanut-compiler.exe FILE [-d]")
	fmt.Println("Compiles Rat23F source code in the given FILE.")
	fmt.Println("Lexical analysis step produces output of \"FILE.lexer\".")
	fmt.Println()
	fmt.Printf("\t-%s, --%s\t\tDisplay debug messages such as lexer state machine statuses and syntax analysis production trees.\n", debugLogArgumentShort, debugLogArgument)
}

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
		showUsage()
		promptExit()
	}
	inputFilePath = os.Args[1]
	if len(inputFilePath) < 1 {
		showUsage()
		promptExit()
	}

	if len(os.Args) > 2 {
		if os.Args[2] == "-"+debugLogArgumentShort || os.Args[2] == "--"+debugLogArgument {
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
