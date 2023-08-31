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

func main() {
	fmt.Println("Welcome to the Peanut Lexer!")
	fmt.Println("Here are the keywords, operators, and separators:")
	fmt.Println(keywords)
	fmt.Println(operators)
	fmt.Println(separators)
}
