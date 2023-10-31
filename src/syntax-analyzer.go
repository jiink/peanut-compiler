package main

import "fmt"

//---- Variables ------------------------------------------------------------------------
/////////////////////////////////////////////////////////////////////////////////////////

var records []record
var currentRecord record

//---- Functions ------------------------------------------------------------------------
/////////////////////////////////////////////////////////////////////////////////////////

/* ---- Helpers --------------------------------------- */

// Reports a syntax error. To be used when a syntax error is encountered.
// Prof. said this can just exit after finding the first syntax error.
func syntaxError(format string, args ...interface{}) {
	errorInfo := fmt.Sprintf("Line %d: Unexpected token \"%s\" | ", currentRecord.lineNumber, currentRecord.lexeme)
	errorMessagePrefix := fmt.Sprintf("[ERROR] %s ", errorInfo)
	log(errorMessagePrefix+format+"\n", args...)
	promptExit()
}

// Returns the next record in each subsequent call.
// Also updates global variable `currentRecord.lexeme` with the lexeme of the new record.
func nextRecord() record {
	var record = record{tokenType: Unrecognized, lexeme: "<EOF>"}
	if len(records) > 0 {
		record = records[0]
		records = records[1:]
		logDebug("Token: %s\tLexeme: %s\n", record.tokenType.String(), record.lexeme)
	} else {
		logDebug("Reached end of file.\n\n")
	}
	currentRecord = record
	return record
}

/* ---- Productions ("prods") ------------------------- */

func prodRat23F() {
	logDebug("\t<Rat23F> ::= <Opt Function Definitions> # <Opt Declaration List> <Statement List> #\n")
	prodOptFunctionDefinitions()
	if currentRecord.lexeme == "#" {
		nextRecord()
		prodOptDeclarationList()
		prodStatementList()
		if currentRecord.lexeme == "#" {
			nextRecord()
			if currentRecord.lexeme != "<EOF>" { // Check for the end
				syntaxError("<EOF> expected")
			}
		} else {
			syntaxError("'#' expected")
		}
	} else {
		syntaxError("'#' expected")
	}
}

func prodOptFunctionDefinitions() {
	logDebug("\t<Opt Function Definitions> ::= <Function Definitions> | <Empty>\n")
	if currentRecord.lexeme == "function" {
		prodFunctionDefinitions()
	}
}

func prodFunctionDefinitions() {
	logDebug("\t<Function Definitions> ::= <Function> <Function Definitions Continued>\n")
	prodFunction()
	prodFunctionDefinitionsContinued()
}

func prodFunctionDefinitionsContinued() {
	logDebug("\t<Function Definitions Continued> ::= <Empty> | <Function Definitions>\n")
	if currentRecord.lexeme == "function" {
		prodFunctionDefinitions()
	}
}

func prodFunction() {
	logDebug("\t<Function> ::= function <Identifier> ( <Opt Parameter List> ) <Opt Declaration List> <Body>\n")
	if currentRecord.lexeme == "function" {
		nextRecord()
		if currentRecord.tokenType == Identifier {
			nextRecord()
			if currentRecord.lexeme == "(" {
				nextRecord()
				prodOptParameterList()
				if currentRecord.lexeme == ")" {
					nextRecord()
					prodOptDeclarationList()
					prodBody()
				} else {
					syntaxError("')' expected")
				}
			} else {
				syntaxError("'(' expected")
			}
		} else {
			syntaxError("Identifier expected")
		}
	} else {
		syntaxError("'function' expected")
	}
}

func prodOptParameterList() {
	logDebug("\t<Opt Parameter List> ::= <Parameter List> | <Empty>\n")
	if currentRecord.tokenType == Identifier {
		prodParameterList()
	}
}

func prodParameterList() {
	logDebug("\t<Parameter List> ::= <Parameter> <Parameter List Continued>\n")
	prodParameter()
	prodParameterListContinued()
}

func prodParameterListContinued() {
	logDebug("\t<Parameter List Continued> ::= <Empty> | ,<Parameter List>\n")
	if currentRecord.lexeme == "," {
		nextRecord()
		prodParameterList()
	}
}

func prodParameter() {
	logDebug("\t<Parameter> ::= <IDs> <Qualifier>\n")
	prodIDs()
	prodQualifier()
}

func prodQualifier() {
	logDebug("\t<Qualifier> ::= integer | bool | real\n")
	if currentRecord.lexeme == "integer" || currentRecord.lexeme == "bool" || currentRecord.lexeme == "real" {
		nextRecord()
	} else {
		syntaxError("'bool', 'real', or 'integer' expected")
	}
}

func prodBody() {
	logDebug("\t<Body> ::= { <Statement List> }\n")
	if currentRecord.lexeme == "{" {
		nextRecord()
		prodStatementList()
		if currentRecord.lexeme == "}" {
			nextRecord()
		} else {
			syntaxError("'}' expected")
		}
	} else {
		syntaxError("'{' expected")
	}
}

func prodOptDeclarationList() {
	logDebug("\t<Opt Declaration List> ::= <Declaration List> | <Empty>\n")
	if currentRecord.lexeme == "integer" || currentRecord.lexeme == "bool" || currentRecord.lexeme == "real" {
		prodDeclarationList()
	}
}

func prodDeclarationList() {
	logDebug("\t<Declaration List> ::= <Declaration> ; <Declaration List Continued>\n")
	prodDeclaration()
	if currentRecord.lexeme == ";" {
		nextRecord()
		prodDeclarationListContinued()
	} else {
		syntaxError("';' expected")
	}
}

func prodDeclarationListContinued() {
	logDebug("\t<Declaration List Continued> ::= <Empty> | <Declaration List>\n")
	if currentRecord.lexeme == "integer" || currentRecord.lexeme == "bool" || currentRecord.lexeme == "real" {
		prodDeclarationList()
	}
}

func prodDeclaration() {
	logDebug("\t<Declaration> ::= <Qualifier> <IDs>\n")
	prodQualifier()
	prodIDs()
}

func prodIDs() {
	logDebug("\t<IDs> ::= <Identifier> <IDs Continued>\n")
	if currentRecord.tokenType == Identifier {
		nextRecord()
		prodIDsContinued()
	} else {
		syntaxError("Identifier Expected")
	}
}

func prodIDsContinued() {
	logDebug("\t<IDs Continued> ::= <Empty> | , <IDs>\n")
	if currentRecord.lexeme == "," {
		nextRecord()
		prodIDs()
	}
}

func prodStatementList() {
	logDebug("\t<Statement List> ::= <Statement> <Statement List Continued>\n")
	prodStatement()
	prodStatementListContinued()
}

func prodStatementListContinued() {
	logDebug("\t<Statement List Continued> ::= <Empty> | <Statement List>\n")
	switch currentRecord.lexeme {
	case "{", "if", "ret", "put", "get", "while":
		prodStatementList()
	default:
		if currentRecord.tokenType == Identifier {
			prodStatementList()
		}
	}
}

func prodStatement() {
	logDebug("\t<Statement> ::= <Compound> | <Assign> | <If> | <Return> | <Print> | <Scan> | <While>\n")
	if currentRecord.lexeme == "{" {
		prodCompound()
	} else if currentRecord.tokenType == Identifier {
		prodAssign()
	} else if currentRecord.lexeme == "if" {
		prodIf()
	} else if currentRecord.lexeme == "ret" {
		prodReturn()
	} else if currentRecord.lexeme == "put" {
		prodPrint()
	} else if currentRecord.lexeme == "get" {
		prodScan()
	} else if currentRecord.lexeme == "while" {
		prodWhile()
	} else {
		syntaxError("Not valid")
	}
}

func prodCompound() {
	logDebug("\t<Compound> ::= { <Statement List> }\n")
	if currentRecord.lexeme == "{" {
		nextRecord()
		prodStatementList()
		if currentRecord.lexeme == "}" {
			nextRecord()
		} else {
			syntaxError("'}' Expected")
		}
	} else {
		syntaxError("'{' Expected")
	}
}

func prodAssign() {
	logDebug("\t<Assign> ::= <Identifier> = <Expression> ;\n")
	if currentRecord.tokenType == Identifier {
		nextRecord()
		if currentRecord.lexeme == "=" {
			nextRecord()
			prodExpression()
			if currentRecord.lexeme == ";" {
				nextRecord()
			} else {
				syntaxError("';' Expected")
			}
		} else {
			syntaxError("'=' Expected")
		}
	} else {
		syntaxError("Identifier Expected")
	}
}

func prodIf() {
	logDebug("\t<If> ::= if ( <Condition> ) <Statement> <If Continued>\n")
	if currentRecord.lexeme == "if" {
		nextRecord()
		if currentRecord.lexeme == "(" {
			nextRecord()
			prodCondition()
			if currentRecord.lexeme == ")" {
				nextRecord()
				prodStatement()
				prodIfContinued()
			} else {
				syntaxError("')' Expected")
			}
		} else {
			syntaxError("'(' Expected")
		}
	} else {
		syntaxError("'if' Expected")
	}
}

func prodIfContinued() {
	logDebug("\t<If Continued> ::= endif | else <Statement> endif\n")
	if currentRecord.lexeme == "endif" {
		nextRecord()
	} else if currentRecord.lexeme == "else" {
		nextRecord()
		prodStatement()
		if currentRecord.lexeme == "endif" {
			nextRecord()
		} else {
			syntaxError("'endif' Expected")
		}
	} else {
		syntaxError("'else' or 'endif' Expected")
	}
}

func prodReturn() {
	logDebug("\t<Return> ::= ret <Return Continued> \n")
	if currentRecord.lexeme == "ret" {
		nextRecord()
		prodReturnContinued()
	} else {
		syntaxError("'ret' Expected")
	}
}

func prodReturnContinued() {
	logDebug("\t<Return Continued> ::= ; | <Expression> ;\n")
	if currentRecord.lexeme == ";" {
		nextRecord()
		return
	}
	prodExpression()
	if currentRecord.lexeme == ";" {
		nextRecord()
	} else {
		syntaxError("';' Expected")
	}
}

func prodPrint() {
	logDebug("\t<Print> ::= put ( <Expression> );\n")
	if currentRecord.lexeme == "put" {
		nextRecord()
		if currentRecord.lexeme == "(" {
			nextRecord()
			prodExpression()
			if currentRecord.lexeme == ")" {
				nextRecord()
				if currentRecord.lexeme == ";" {
					nextRecord()
				} else {
					syntaxError("';' expected")
				}
			} else {
				syntaxError("')' expected")
			}
		} else {
			syntaxError("'(' expected")
		}
	} else {
		syntaxError("'put' expected")
	}
}

func prodScan() {
	logDebug("\t<Scan> ::= get ( <IDs> );\n")
	if currentRecord.lexeme == "get" {
		nextRecord()
		if currentRecord.lexeme == "(" {
			nextRecord()
			prodIDs()
			if currentRecord.lexeme == ")" {
				nextRecord()
				if currentRecord.lexeme == ";" {
					nextRecord()
				} else {
					syntaxError("';' expected")
				}
			} else {
				syntaxError("')' expected")
			}
		} else {
			syntaxError("'(' expected")
		}
	} else {
		syntaxError("'get' expected")
	}
}

func prodWhile() {
	logDebug("\t<While> ::= while ( <Condition> ) <Statement>\n")
	if currentRecord.lexeme == "while" {
		nextRecord()
		if currentRecord.lexeme == "(" {
			nextRecord()
			prodCondition()
			if currentRecord.lexeme == ")" {
				nextRecord()
				prodStatement()
			} else {
				syntaxError("')' expected")
			}
		} else {
			syntaxError("'(' expected")
		}
	} else {
		syntaxError("'while' expected")
	}
}

func prodCondition() {
	logDebug("\t<Condition> ::= <Expression> <Relop> <Expression>\n")
	prodExpression()
	prodRelop()
	prodExpression()
}

func prodRelop() {
	logDebug("\t<Relop> ::= == | != | > | < | <= | =>\n")
	switch currentRecord.lexeme {
	case "==", "!=", ">", "<", "<=", "=>":
		nextRecord()
	default:
		syntaxError("'==', '!=', '>', '<', '<=', or '=>' expected")
	}
}

func prodExpression() {
	logDebug("\t<Expression> ::= <Term> <Expression Prime>\n")
	prodTerm()
	prodExpressionPrime()
}

func prodExpressionPrime() {
	logDebug("\t<Expression Prime> ::= <Expression Continued> <Expression Prime> | <Empty>\n")
	if currentRecord.lexeme == "+" || currentRecord.lexeme == "-" {
		prodExpressionContinued()
		prodExpressionPrime()
	}
}

func prodExpressionContinued() {
	logDebug("\t<Expression Continued> ::= + <Term> | - <Term>\n")
	if currentRecord.lexeme == "+" || currentRecord.lexeme == "-" {
		nextRecord()
		prodTerm()
	} else {
		syntaxError("'+' or '-' expected")
	}
}

func prodTerm() {
	logDebug("\t<Term> ::= <Factor> <Term Prime>\n")
	prodFactor()
	prodTermPrime()
}

func prodTermPrime() {
	logDebug("\t<Term Prime> ::= <Term Continued> <Term Prime> | <Empty>\n")
	if currentRecord.lexeme == "*" || currentRecord.lexeme == "/" {
		prodTermContinued()
		prodTermPrime()
	}
}

func prodTermContinued() {
	logDebug("\t<Term Continued> ::= * <Factor> | / <Factor>\n")
	if currentRecord.lexeme == "*" || currentRecord.lexeme == "/" {
		nextRecord()
		prodFactor()
	} else {
		syntaxError("'*' or '/' expected")
	}
}

func prodFactor() {
	logDebug("\t<Factor> ::= - <Primary> | <Primary>\n")
	if currentRecord.lexeme == "-" {
		nextRecord()
		prodPrimary()
		return
	}
	prodPrimary()
}

func prodPrimary() {
	logDebug("\t<Primary> ::= <Identifier> <Primary Continued> | <Integer> | ( <Expression> ) | <Real> | true | false\n")
	if currentRecord.tokenType == Identifier {
		nextRecord()
		prodPrimaryContinued()
		return
	}
	if currentRecord.tokenType == Integer {
		nextRecord()
		return
	}
	if currentRecord.lexeme == "(" {
		nextRecord()
		prodExpression()
		if currentRecord.lexeme == ")" {
			nextRecord()
		} else {
			syntaxError("')' expected")
		}
		return
	}
	if currentRecord.tokenType == Real {
		nextRecord()
		return
	}
	if currentRecord.lexeme == "true" {
		nextRecord()
		return
	}
	if currentRecord.lexeme == "false" {
		nextRecord()
		return
	}
	syntaxError("Identifier, integer, '(', Real, 'true', or 'false' expected")
}

func prodPrimaryContinued() {
	logDebug("\t<Primary Continued> ::= <Empty> | ( <IDs> )\n")
	if currentRecord.lexeme == "(" {
		nextRecord()
		prodIDs()
		if currentRecord.lexeme == ")" {
			nextRecord()
		} else {
			syntaxError("Expected ')'")
		}
	}
}

func prodEmpty() {
	logDebug("\t<Empty> ::= ε\n")
}

/* ---- The main attraction --------------------------- */

func syntaxAnalyzer(recordsIn []record) error {
	records = recordsIn // Store a copy for getNextRecord() usage
	nextRecord()        // Set up global variables with first token
	prodRat23F()        // Call initial production
	return nil
}
