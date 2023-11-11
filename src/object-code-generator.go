// TODO: Check for type match when rat23f identifier is being used.
package main

import "fmt"

//---- Definitions ----------------------------------------------------------------------
/////////////////////////////////////////////////////////////////////////////////////////

type identifierType int
type operationType int

const (
	TypeInteger identifierType = iota
	TypeBool
)

const (
	PUSHI operationType = iota
	PUSHM
	POPM
	STDOUT
	STDIN
	ADD
	SUB
	MUL
	DIV
	GRT
	LES
	EQU
	NEQ
	GEQ
	LEQ
	JUMPZ
	JUMP
	LABEL
)

type symbolTableEntry struct {
	identifier     string
	memoryLocation int
	symbolType     identifierType
}

type instruction struct {
	operation operationType
	operand   int
}

const symbolTableStartAddress = 7000

//---- Variables ------------------------------------------------------------------------
/////////////////////////////////////////////////////////////////////////////////////////

var symbolTable []symbolTableEntry
var instructionTable []instruction
var currentInstructionAddress = 1
var currentSymbolTableAddress = symbolTableStartAddress
var jumpStack = make(stack, 0)

//---- Functions ------------------------------------------------------------------------
/////////////////////////////////////////////////////////////////////////////////////////

func addSymbol(identifier string, symbolType identifierType) {
	// First see if symbol is already in table
	if isIdentifierInSymbolTable(identifier) {
		fmt.Printf("[ERROR] Symbol %s already in symbol table\n", identifier)
		return
	}
	symbolTable = append(symbolTable, symbolTableEntry{identifier, len(symbolTable), symbolType})
	currentSymbolTableAddress++
}

func getSymbol(identifier string) (symbolTableEntry, bool) {
	for _, symbol := range symbolTable {
		if symbol.identifier == identifier {
			return symbol, true
		}
	}
	fmt.Printf("[ERROR] Symbol %s used before declaration\n", identifier)
	return symbolTableEntry{}, false
}

func printSymbolTable() {
	fmt.Println("Symbol Table:")
	for _, symbol := range symbolTable {
		fmt.Printf("%s\t%d\t%d\n", symbol.identifier, symbol.memoryLocation, symbol.symbolType)
	}
}

func isIdentifierInSymbolTable(identifier string) bool {
	_, found := getSymbol(identifier)
	return found
}

func generateInstruction(op operationType, operand int) {
	instructionTable = append(instructionTable, instruction{op, operand})
	currentInstructionAddress++
}

func printInstructionTable() {
	fmt.Println("Instruction Table:")
	for _, instruction := range instructionTable {
		fmt.Printf("%d\t%d\t%d\n", instruction.operation, instruction.operand, currentInstructionAddress)
	}
}

func backPatch(jumpAddress int) {
	jumpStack, addr := jumpStack.Pop()
	instructionTable[addr].operand = jumpAddress
}
