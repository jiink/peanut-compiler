package main

import (
	"fmt"
	"io"
	"os"
)

type tokenType int

const (
	Identifier tokenType = iota
	Keyword
	Integer
	Real
	Operator
	Separator
)

type record struct {
	tokenType tokenType
	lexeme    string
}

type stateMachine struct {
	currentState int
}

var keywords = []string{
	"function",
	"if",
	"endif",
	"else",
	"ret",
	"put",
	"get",
	"bool",
	"real",
	"while",
	"true",
	"false",
}
var operators = []string{
	"+",
	"-",
	"*",
	"/",
	"=",
	">",
	"<",
	"==",
	"!=",
	"<=",
	">=",
}
var separators = []string{
	"(",
	")",
	"{",
	"}",
	".",
	",",
	";",
	"#",
}

// Returns true if the given string is found in the
// list of keywords Rat23F recognizes.
func isKeyword(str string) bool {
	for _, keyword := range keywords {
		if str == keyword {
			return true
		}
	}
	return false
}

func isOperator(str string) bool {
	for _, operator := range operators {
		if str == operator {
			return true
		}
	}
	return false
}

func isSeparator(str string) bool {
	for _, separator := range separators {
		if str == separator {
			return true
		}
	}
	return false
}

func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func readChar(str string, index *int) rune {
	char := []rune(str)[*index]
	*index = *index + 1
	return char
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Welcome to the Peanut Lexer for Rat23F!")
	fmt.Println("Here are the keywords, operators, and separators:")
	fmt.Println(keywords)
	fmt.Println(operators)
	fmt.Println(separators)

	fmt.Println("Let's read in the source code file.")
	const sourceCodePath = "test.rat"
	file, err := os.Open(sourceCodePath)
	check(err)
	defer file.Close()
	// To make things simple we'll just put it all in a string.
	content, err := io.ReadAll(file)
	check(err)
	sourceCode := string(content)
	fmt.Println("Source code: " + sourceCode)

	fmt.Println("Let the main lexing loop begin...")
	sourceCodePointer := 0 // Points to the current character in the source code
	for sourceCodePointer < len(sourceCode) {
		currentChar := readChar(sourceCode, &sourceCodePointer)
		fmt.Printf("Current character: %c\n", currentChar)
		if isLetter(currentChar) {
			//fmt.Println("It's a letter")
			// Call relevant DFSM
		} else if isDigit(currentChar) {
			//fmt.Println("It's a digit")
			// Call relevant DFSM
		} else {
			fmt.Printf("Unrecognized character %c\n", currentChar)
		}
	}
}
