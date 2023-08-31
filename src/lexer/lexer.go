package main

import "fmt"

var keywords = [...]string{
	"function",
	"if",
	"endif",
	"else",
	"ret",
	"put",
	"get",
}
var operators = [...]string{
	"+",
	"-",
	"*",
	"/",
}
var separators = [...]string{
	"(",
	")",
	"{",
	"}",
	",",
	";",
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

func main() {
	fmt.Println("Welcome to the Peanut Lexer!")
	fmt.Println("Here are the keywords, operators, and separators:")
	fmt.Println(keywords)
	fmt.Println(operators)
	fmt.Println(separators)
	fmt.Println("Is 'function' a keyword?", isKeyword("function"))
	fmt.Println("Is 'pineapple' a keyword?", isKeyword("pineapple"))
	fmt.Println("Is '+' an operator?", isOperator("+"))
	fmt.Println("Is '!' an operator?", isOperator("!"))
	fmt.Println("Is '(' a separator?", isSeparator("("))
	fmt.Println("Is '!' a separator?", isSeparator("!"))
}
