package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

//---- Definitions ----------------------------------------------------------------------
/////////////////////////////////////////////////////////////////////////////////////////

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
	Period
	Special
)

type record struct {
	tokenType  tokenType
	lexeme     string
	lineNumber int
}

var keywords = []string{
	"function",
	"if",
	"endif",
	"else",
	"ret",
	"put",
	"get",
	"integer",
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
	"=>",
}
var separators = []string{
	"(",
	")",
	"{",
	"}",
	",",
	";",
	"#",
}

type FSM struct {
	inputSymbolSet  []symbolType
	transitionTable [][]int
	acceptingStates []int
	initialState    int
	currentState    int
}

//---- Functions ------------------------------------------------------------------------
/////////////////////////////////////////////////////////////////////////////////////////

// Returns a string representation of the given tokenType.
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

// Returns a string representation of the given symbolType.
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

// Returns true if the given string is found in the
// list of operators Rat23F recognizes.
func isOperator(str string) bool {
	for _, operator := range operators {
		if str == operator {
			return true
		}
	}
	return false
}

// Returns true if the given string is found in the
// list of separators Rat23F recognizes.
func isSeparator(str string) bool {
	for _, separator := range separators {
		if str == separator {
			return true
		}
	}
	return false
}

// Returns true if the given rune is in the English alphabet.
func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

// Returns true if the given rune is a 0 through 9 digit.
func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

// Returns true if the given rune is a period.
func isPeriod(r rune) bool {
	return r == '.'
}

// Returns the symbolType associated with the given rune.
// This symbol type can be given to a finite state machine.
func charToSymbolType(r rune) symbolType {
	if isLetter(r) {
		return Letter
	}
	if isDigit(r) {
		return Digit
	}
	if isPeriod(r) {
		return Period
	}
	return Special
}

// Returns the character of the source code at the given index,
// and increments the index for the subsequent calls.
func readCharSourceCode(index *int) rune {
	char := ' '
	if *index < len([]rune(sourceCode)) {
		char = []rune(sourceCode)[*index]
	}
	*index = *index + 1
	return char
}

// Decrements the given index.
func backUp(index *int) bool {
	if *index < 1 {
		return false
	}
	*index = *index - 1
	return true
}

// Panics if the given error is not nil.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Converts a given list of records into a 2-column table string.
func recordsToString(records []record) string {
	columnWidth := 12
	s := "----------------------\n"
	s += "[Token]" + strings.Repeat(" ", columnWidth-7) + ": [Lexeme]\n"
	for _, r := range records {
		s += fmt.Sprintf("%s%s: %s\n", r.tokenType, strings.Repeat(" ", columnWidth-len(r.tokenType.String())), r.lexeme)
	}
	s += "----------------------\n"
	return s
}

// Logs the given lexer records to the console and to an output file
// whos path depends on the Rat23F input file.
func logRecords(records []record, printToConsole bool, printToFile bool) {
	recordsReport := recordsToString(records)
	if printToConsole {
		fmt.Println(recordsReport)
	}

	// Create output file
	if printToFile {
		outputPath := inputFilePath + ".lexr"
		f, err := os.Create(outputPath)
		check(err)
		defer f.Close()
		f.WriteString(recordsReport)
		fmt.Printf("Created file: %s\n", outputPath)
	}
}

// Strips all CRs from CR+LF line endings on Windows.
// Makes counting newlines easier.
func trimCarriageReturns(s string) string {
	return strings.ReplaceAll(s, "\r", "")
}

// Removes all Rat23F comments from the given string.
func removeComments(s string) string {

	re := regexp.MustCompile(`(?s)\[\*.*?\*\]`)

	return re.ReplaceAllString(s, "")
}

/* ---- The main attractions -------------------------- */

// Returns false when the FSM must end, and true when it's ready
// to accept a new symbol.
func (f *FSM) transition(symbol symbolType) bool {
	columnIndex := slices.Index(f.inputSymbolSet, symbol)
	if columnIndex == -1 {
		logDebug("Invalid symbol: %s. Ending FSM...\n", symbol)
		return false
	}
	logDebug("Current state: %d, Symbol: %s\n", f.currentState, symbol)
	f.currentState = f.transitionTable[f.currentState][columnIndex]
	logDebug("New state: %d\n\n", f.currentState)
	// 0 is the unrecoverable state.
	if f.currentState == 0 {
		logDebug("Fell into unrecoverable state. Ending FSM.\n")
		return false
	}
	return true
}

// Resets the FSM to its initial state.
func (f *FSM) reset() {
	f.currentState = f.initialState
}

// Returns true if the FSM is in an accepting state.
func (f *FSM) isInAcceptingState() bool {
	return slices.Contains(f.acceptingStates, f.currentState)
}

// Runs the FSM on the Rat23F source code, starting at the given index.
func (f *FSM) run(sourceCodePointer *int) bool {
	f.reset()
	backUp(sourceCodePointer)
	maxLength := 500
	for i := 0; i < maxLength; i++ {
		newChar := readCharSourceCode(sourceCodePointer)
		logDebug("New char: '%c'\n", newChar)
		symbol := charToSymbolType(newChar)
		if !f.transition(symbol) {
			break
		}
	}
	logDebug("Final state: %d\n", f.currentState)
	return f.isInAcceptingState()
}

// Runs the lexer on the given Rat23F source code.
// Returns a slice of lexer records, showing
// the token type and lexeme for each token identified.
func lexer(sourceCode string) ([]record, error) {

	records := []record{}
	sourceCodePointer := 0
	currentLineNumber := 1

	// The transition tables are made by hand. After drawing a
	// non-deterministic FSM diagram, it is converted
	// into a DFSM table and then each state is given
	// a number and entered into tables in these structs.

	fsmIdentifier := FSM{
		inputSymbolSet: []symbolType{Letter, Digit},
		transitionTable: [][]int{
			// l  d
			{0, 0}, // 0
			{2, 0}, // 1
			{3, 4}, // 2
			{3, 4}, // 3
			{3, 4}, // 4
		},
		acceptingStates: []int{2, 3},
		initialState:    1,
	} // FSM for identifiers

	fsmReal := FSM{
		inputSymbolSet: []symbolType{Digit, Period, Letter},
		transitionTable: [][]int{
			// d, p, l
			{0, 0, 0}, // 0
			{2, 0, 0}, // 1
			{2, 3, 5}, // 2
			{4, 0, 5}, // 3
			{4, 0, 5}, // 4
			{5, 5, 5}, // 5, ensures that incorrect tokens are fully read instead of split
		},
		acceptingStates: []int{4},
		initialState:    1,
	} // FSM for reals

	fsmInteger := FSM{
		inputSymbolSet: []symbolType{Digit, Period, Letter},
		transitionTable: [][]int{
			// d, p, l
			{0, 0, 0}, // 0
			{2, 0, 0}, // 1
			{2, 3, 3}, // 2
			{3, 3, 3}, // 3, ensures that incorrect tokens are fully read instead of split
		},
		acceptingStates: []int{2},
		initialState:    1,
	} // FSM for integers

	for sourceCodePointer < len(sourceCode) {
		tokenType := Unrecognized
		lexemeStartIndex := sourceCodePointer
		currentChar := readCharSourceCode(&sourceCodePointer)
		logDebug("new loop with %q\n", currentChar)

		if isLetter(currentChar) {
			// Identifiers start with a letter.
			if fsmIdentifier.run(&sourceCodePointer) {
				tokenType = Identifier
			}
		} else if isDigit(currentChar) {
			// Reals and integers start with a digit.
			// First check if a real is here. If not, back up and try again as an integer
			sourceCodePointerBookmark := sourceCodePointer
			if fsmReal.run(&sourceCodePointer) {
				tokenType = Real
			} else {
				sourceCodePointer = sourceCodePointerBookmark
				if fsmInteger.run(&sourceCodePointer) {
					tokenType = Integer
				}
			}
		} else {
			// Check for a 2-char operator. If not, back up and check for a 1-char operator, then separator.
			nextChar := readCharSourceCode(&sourceCodePointer) // Peek at next character for 2-char operators (e.g. ==)
			logDebug("next char: %q\n", nextChar)
			if isOperator(string(currentChar) + string(nextChar)) {
				tokenType = Operator
				_ = readCharSourceCode(&sourceCodePointer) // Backs up later
			} else if isOperator(string(currentChar)) {
				tokenType = Operator
			} else if isSeparator(string(currentChar)) {
				tokenType = Separator
			} else if isPeriod(currentChar) {
				fsmInteger.run(&sourceCodePointer) // checks for tokens like ".123"
				tokenType = Unrecognized
			} else {
				tokenType = Unrecognized
				logDebug("Unhandled character %q\n", currentChar)
				if currentChar == '\n' {
					currentLineNumber++
				}
			}
		}
		backUp(&sourceCodePointer)
		if tokenType != Unrecognized {
			lexemeEndIndex := sourceCodePointer
			lexeme := sourceCode[lexemeStartIndex:lexemeEndIndex]
			logDebug("\"%s\" Accepted!\n", lexeme)
			// If it's an identifier, it might be a keyword
			if tokenType == Identifier && isKeyword(lexeme) {
				tokenType = Keyword
			}
			records = append(records, record{tokenType: tokenType, lexeme: lexeme, lineNumber: currentLineNumber})
		} else {
			lexeme := sourceCode[lexemeStartIndex:sourceCodePointer]
			if strings.TrimSpace(lexeme) != "" {
				fmt.Printf("[ERROR] Unrecognized token \"%q\".\n", lexeme)
				records = append(records, record{tokenType: tokenType, lexeme: lexeme, lineNumber: currentLineNumber})
			}
		}
	}
	return records, nil
}
