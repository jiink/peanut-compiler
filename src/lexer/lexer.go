package main

import (
	"fmt"
	"regexp"
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

// Split source code based on whitespace, separators, and operators.
// Return the pieces of strings.
func splitSourceCode(sourceCode string, separators []string, operators []string) []string {
	stuffThatNeedsToBeEscaped := []string{
		`]`,
		`-`,
	}
	splitterPattern := ``
	splitters := append(separators, operators...)
	// Construct the separators pattern given the list of separators and operators
	for _, splitter := range splitters {
		if slices.Contains(stuffThatNeedsToBeEscaped, splitter) {
			splitterPattern += `\`
		}
		splitterPattern += splitter
	}
	finalRegEx := `\w+|[` + splitterPattern + `]`
	re := regexp.MustCompile(finalRegEx)
	return re.FindAllString(sourceCode, -1)
}

func main() {
	const sourceCode = "while (fahr < upper) a = 23.00;"
	fmt.Println("Welcome to the Peanut Lexer!")
	fmt.Println("Here are the keywords, operators, and separators:")
	fmt.Println(keywords)
	fmt.Println(operators)
	fmt.Println(separators)
	fmt.Println("Let's separate the source code into lexemes...")
	fmt.Println("Source code: " + sourceCode)
	sourceCodeSplit := splitSourceCode(sourceCode, separators, operators)
	fmt.Println(sourceCodeSplit)
	fmt.Println("Let the main lexing loop begin...")
	for _, tok := range sourceCodeSplit {
		// TODO: for each lexeme, create a record and identify its token type
		fmt.Println("Lexeme: '" + tok + "'")
	}
}
