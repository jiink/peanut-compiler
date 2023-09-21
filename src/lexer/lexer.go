package main

import (
	"fmt"
	"io"
	"os"
	"slices"
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

type symbolType int

const (
	Letter symbolType = iota
	Digit
	Special
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

var sourceCode = ""

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

func charToSymbolType(r rune) symbolType {
	if isLetter(r) {
		return Letter
	}
	if isDigit(r) {
		return Digit
	}
	return Special
}

func readCharSourceCode(index *int) rune {
	char := []rune(sourceCode)[*index]
	*index = *index + 1
	return char
}

func backUp(index *int) bool {
	if *index < 1 {
		return false
	}
	*index = *index - 1
	return true
}

func isAcceptingState(currentState int, acceptingStates []int) bool {
	return slices.Contains(acceptingStates, currentState)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func dfsmIdentifier(sourceCodePointer *int, currentChar rune) bool {
	inputSymbolSet := []symbolType{Letter, Digit}
	// Made from making a graph from a regular expression and
	// converting it to a DFSM by hand, then assigning a number
	// to each unique state.
	transitionTable := [][]int{
		// l  d
		{0, 0}, // 0
		{2, 0}, // 1
		{3, 4}, // 2
		{3, 4}, // 3
		{3, 4}, // 4
	}
	acceptingStates := []int{2, 3, 4}
	currentState := 1

	// Back up char pointer
	backUp(sourceCodePointer)

	maxSteps := 100
	for i := 0; i < maxSteps; i++ {
		newChar := readCharSourceCode(sourceCodePointer)
		fmt.Printf("new char: %c\n", newChar)
		symbol := charToSymbolType(newChar)
		columnIndex := slices.Index(inputSymbolSet, symbol)
		if columnIndex == -1 {
			fmt.Printf("Invalid symbol: %d\n", symbol)
			break
		}
		fmt.Printf("Current state: %d, Symbol: %d\n", currentState, symbol)
		currentState = transitionTable[currentState][columnIndex]
		fmt.Printf("New state: %d\n\n", currentState)
		// 0 is the unrecoverable state.
		if currentState == 0 {
			fmt.Printf("Fell into unrecoverable state.\n")
			return false
		}
	}
	fmt.Printf("Final state: %d\n", currentState)
	return isAcceptingState(currentState, acceptingStates)
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
	sourceCode = string(content)
	fmt.Println("Source code: " + sourceCode)

	fmt.Println("Let the main lexing loop begin...")
	sourceCodePointer := 0 // Points to the current character in the source code
	for sourceCodePointer < len(sourceCode) {
		accepted := false
		lexemeStartIndex := sourceCodePointer
		currentChar := readCharSourceCode(&sourceCodePointer)
		//fmt.Printf("Current character: %c is type %d\n", currentChar, charToSymbolType(currentChar))
		if isLetter(currentChar) {
			//fmt.Println("It's a letter")
			// Call relevant DFSM
			accepted = dfsmIdentifier(&sourceCodePointer, currentChar)
		} else if isDigit(currentChar) {
			//fmt.Println("It's a digit")
			// Call relevant DFSM
			//dfsmInteger(&sourceCodePointer, currentChar)
		} else {
			fmt.Printf("Unrecognized character %c\n", currentChar)
		}
		if accepted {
			backUp(&sourceCodePointer)
			lexemeEndIndex := sourceCodePointer
			fmt.Printf("\"%s\" Accepted! (from %d to %d)\n", sourceCode[lexemeStartIndex:lexemeEndIndex], lexemeStartIndex, lexemeEndIndex)
		} else {
			fmt.Printf("Rejected.\n")
		}
	}
}
