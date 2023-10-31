package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//---- Definitions ----------------------------------------------------------------------
/////////////////////////////////////////////////////////////////////////////////////////

var noDebugLogArgument = "silent"
var noDebugLogArgumentShort = "s"

//---- Variables ------------------------------------------------------------------------
/////////////////////////////////////////////////////////////////////////////////////////

var inputFilePath = ""
var sourceCode = ""
var debugEnabled = true          // Enables/disables debug log messages in the console
var debugLogToFileEnabled = true // Enables/disables debug log messages being written to a file
var logFile *os.File

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
	fmt.Printf("\t-%s, --%s\t\tSilence debug messages such as lexer state machine statuses and syntax analysis production trees.\n", noDebugLogArgumentShort, noDebugLogArgument)
}

func createLogFile() {
	outputPath := inputFilePath
	// Remove everything after the dot (including the dot itself)
	dotIndex := strings.Index(outputPath, ".")
	if dotIndex != -1 {
		result := outputPath[:dotIndex]
		fmt.Println("Result:", result)
	}
	outputPath += "-out.txt"
	f, err := os.Create(outputPath)
	check(err)
	logFile = f
}

// Prints a debug message to the console and/or a file
func logDebug(format string, args ...interface{}) {
	if debugEnabled {
		fmt.Printf("[DEBUG] "+format, args...)
	}
	if debugLogToFileEnabled && (logFile != nil) {
		logFile.WriteString(fmt.Sprintf(format, args...))
	}
}

// Prints a message to the console and/or a file
func log(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	if debugLogToFileEnabled && (logFile != nil) {
		logFile.WriteString(fmt.Sprintf(format, args...))
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
	sourceCode = removeComments(sourceCode)
	sourceCode = trimCarriageReturns(sourceCode)
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
		if os.Args[2] == "-"+noDebugLogArgumentShort || os.Args[2] == "--"+noDebugLogArgument {
			debugEnabled = false
		}
	}
	fmt.Println("Welcome to the Peanut Compiler for Rat23F!")
	readInSourceCode(inputFilePath)

	if debugLogToFileEnabled {
		createLogFile()
	}

	enableDebugSoon := debugEnabled
	debugEnabled = false // Silence debug messages just for the lexer
	enableDebugLogToFileSoon := debugLogToFileEnabled
	debugLogToFileEnabled = false
	//fmt.Println("Beginning lexical analysis...")
	var records, err = lexer(sourceCode)
	if err != nil {
		log("The lexer encountered an error.")
	} else {
		logRecords(records, false, false)
	}
	//fmt.Println("Lexical analysis complete.")

	//fmt.Println("Beginning syntax analysis...")
	debugEnabled = enableDebugSoon
	debugLogToFileEnabled = enableDebugLogToFileSoon
	err = syntaxAnalyzer(records)
	if err != nil {
		log("The syntax analyzer encountered an error.\n")
	} else {
		log("The code is syntactically correct!\n")
	}
	fmt.Println("This is the end of the compiler so far.")
	if logFile != nil {
		logFile.Close()
	}
}
