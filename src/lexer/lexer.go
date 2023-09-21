package main

import (
	"fmt"
	"io"
	"os"
	"slices"
)

////////////////////////////////////////////////////////////////////
//---- Definitions -------------------------------------------------

type tokenType int

const (
	Identifier tokenType = iota
	Keyword
	Integer
	Real
	Operator
	Separator
	Unrecognized
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

////////////////////////////////////////////////////////////////////
//---- Variables ---------------------------------------------------

var sourceCode = ""

////////////////////////////////////////////////////////////////////
//---- Functions ---------------------------------------------------

/* ---- Helpers --------------------------------------- */

func (e tokenType) String() string {
	switch e {
	case Identifier:
		return "Identifier"
	case Keyword:
		return "Keyword"
	case Integer:
		return "Integer"
	case Real:
		return "Real"
	case Operator:
		return "Operator"
	case Separator:
		return "Separator"
	case Unrecognized:
		return "Unrecognized"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

func (e symbolType) String() string {
	switch e {
	case Letter:
		return "Letter"
	case Digit:
		return "Digit"
	case Special:
		return "Special"
	default:
		return fmt.Sprintf("%d", int(e))
	}
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

func printRecord(r record) {
	fmt.Printf("%s\t:\t%s\n", r.tokenType, r.lexeme)
}

func printRecords(records []record) {
	fmt.Println("----------------------")
	fmt.Println("[Token]\t:\t[Lexeme]")
	for _, r := range records {
		printRecord(r)
	}
	fmt.Println("----------------------")
}

/* ---- The main attractions -------------------------- */

func dfsmIdentifier(sourceCodePointer *int) bool {
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

	backUp(sourceCodePointer)
	maxLength := 500
	for i := 0; i < maxLength; i++ {
		newChar := readCharSourceCode(sourceCodePointer)
		fmt.Printf("New char: '%c'\n", newChar)
		symbol := charToSymbolType(newChar)
		columnIndex := slices.Index(inputSymbolSet, symbol)
		if columnIndex == -1 {
			fmt.Printf("Invalid symbol: %s. Ending FSM...\n", symbol)
			break
		}
		fmt.Printf("Current state: %d, Symbol: %s\n", currentState, symbol)
		currentState = transitionTable[currentState][columnIndex]
		fmt.Printf("New state: %d\n\n", currentState)
		// 0 is the unrecoverable state.
		if currentState == 0 {
			fmt.Printf("Fell into unrecoverable state. Ending FSM.\n")
			return false
		}
	}
	fmt.Printf("Final state: %d\n", currentState)
	return isAcceptingState(currentState, acceptingStates)
}

func readInSourceCode() {
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
}

func main() {
	fmt.Println("Welcome to the Peanut Lexer for Rat23F!")
	fmt.Println("Here are the keywords, operators, and separators:")
	fmt.Println(keywords)
	fmt.Println(operators)
	fmt.Println(separators)

	readInSourceCode()

	fmt.Println("Let the main lexing loop begin...")
	var records []record
	sourceCodePointer := 0 // Points to the current character in the source code
	for sourceCodePointer < len(sourceCode) {
		tokenType := Unrecognized
		lexemeStartIndex := sourceCodePointer
		currentChar := readCharSourceCode(&sourceCodePointer)

		if isLetter(currentChar) {
			// Call relevant DFSM. Identifiers start with a letter.
			if dfsmIdentifier(&sourceCodePointer) {
				tokenType = Identifier
			}
		} else if isDigit(currentChar) {
			// Call relevant DFSM. Integers and reals start with a digit.
		} else {
			fmt.Printf("Unhandled character '%c'\n", currentChar)
			tokenType = Unrecognized
		}

		if tokenType != Unrecognized {
			backUp(&sourceCodePointer)
			lexemeEndIndex := sourceCodePointer
			lexeme := sourceCode[lexemeStartIndex:lexemeEndIndex]
			fmt.Printf("\"%s\" Accepted!\n", lexeme)
			// If it's an identifier, it might be a keyword
			if tokenType == Identifier && isKeyword(lexeme) {
				tokenType = Keyword
			}
			records = append(records, record{tokenType: tokenType, lexeme: lexeme})
		} else {
			fmt.Printf("Rejected.\n")
		}
	}
	printRecords(records)
}
