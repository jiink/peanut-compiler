// TODO: Check for type match when rat23f identifier is being used.
package main

import (
	"fmt"
	"slices"
)

//---- Definitions ----------------------------------------------------------------------
/////////////////////////////////////////////////////////////////////////////////////////

type identifierType int
type operationType int

const (
	TypeInteger identifierType = iota
	TypeReal
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

// So "0" doesn't appear in the instruction table for instructions that don't take args.
var instructionsWithOperands = []operationType{PUSHI, PUSHM, POPM, JUMPZ, JUMP}

type symbolTableEntry struct {
	identifier     string
	memoryLocation int
	symbolType     identifierType
}

type instruction struct {
	operation operationType
	operand   int
}

const instructionStartAddress = 1
const symbolTableStartAddress = 7000

//---- Variables ------------------------------------------------------------------------
/////////////////////////////////////////////////////////////////////////////////////////

var symbolTable []symbolTableEntry
var instructionTable []instruction
var currentInstructionAddress = instructionStartAddress
var currentSymbolTableAddress = symbolTableStartAddress
var jumpStack stack

//---- Functions ------------------------------------------------------------------------
/////////////////////////////////////////////////////////////////////////////////////////

// Returns a string representation of the given operationType.
func (e operationType) String() string {
	switch e {
	case PUSHI:
		return "PUSHI"
	case PUSHM:
		return "PUSHM"
	case POPM:
		return "POPM"
	case STDOUT:
		return "STDOUT"
	case STDIN:
		return "STDIN"
	case ADD:
		return "ADD"
	case SUB:
		return "SUB"
	case MUL:
		return "MUL"
	case DIV:
		return "DIV"
	case GRT:
		return "GRT"
	case LES:
		return "LES"
	case EQU:
		return "EQU"
	case NEQ:
		return "NEQ"
	case GEQ:
		return "GEQ"
	case LEQ:
		return "LEQ"
	case JUMPZ:
		return "JUMPZ"
	case JUMP:
		return "JUMP"
	case LABEL:
		return "LABEL"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

// Returns a string representation of the given identifierType.
func (e identifierType) String() string {
	switch e {
	case TypeInteger:
		return "integer"
	case TypeBool:
		return "bool"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

func addSymbol(identifier string, symbolType identifierType) {
	// First see if symbol is already in table
	if isIdentifierInSymbolTable(identifier) {
		fmt.Printf("[ERROR] Symbol %s already in symbol table\n", identifier)
		return
	}
	symbolTable = append(symbolTable, symbolTableEntry{identifier, currentSymbolTableAddress, symbolType})
	currentSymbolTableAddress++
	printSymbolTable()
}

func getSymbol(identifier string) (symbolTableEntry, bool) {
	for _, symbol := range symbolTable {
		if symbol.identifier == identifier {
			return symbol, true
		}
	}
	return symbolTableEntry{}, false
}

func printSymbolTable() {
	fmt.Println("Symbol Table:")
	for _, symbol := range symbolTable {
		fmt.Printf("%s\t%d\t%s\n", symbol.identifier, symbol.memoryLocation, symbol.symbolType.String())
	}
}

func isIdentifierInSymbolTable(identifier string) bool {
	_, found := getSymbol(identifier)
	return found
}

func generateInstruction(op operationType, operand int) {
	instructionTable = append(instructionTable, instruction{op, operand})
	logDebug("\t op ::= %s, operand ::= %d, intr_address ::= %d\n", op.String(), operand, currentInstructionAddress)
	currentInstructionAddress++
}

func printInstructionTable() {
	fmt.Println("Instruction Table:")
	for i, instruction := range instructionTable {
		fmt.Printf("%d\t%s", i+instructionStartAddress, instruction.operation.String())
		if slices.Contains(instructionsWithOperands, instruction.operation) {
			fmt.Printf("\t%d", instruction.operand)
		}
		fmt.Printf("\n")

	}
}

func backPatch(jumpAddress int) {
	addr := jumpStack.Pop()
	instructionTable[addr].operand = jumpAddress
}
