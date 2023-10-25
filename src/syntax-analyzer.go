package main

//---- Variables ------------------------------------------------------------------------
/////////////////////////////////////////////////////////////////////////////////////////

var records []record
var currentToken string

//---- Functions ------------------------------------------------------------------------
/////////////////////////////////////////////////////////////////////////////////////////

/* ---- Helpers --------------------------------------- */

// Returns the next record in each subsequent call.
func getNextRecord() record {
	var record = records[0]
	records = records[1:]
	return record
}

/* ---- Productions ("prods") ------------------------- */

func prodRat23F() {
	logDebug("<Rat23F> ::= <Opt Function Definitions> # <Opt Declaration List> <Statement List> #")
}

func prodOptFunctionDefinitions() {
	logDebug("<Opt Function Definitions> ::= <Function Definitions> | <Empty>")
}

func prodFunctionDefinitions() {
	logDebug("<Function Definitions> ::= <Function> <Function Definitions Continued>")
}

func prodFunctionDefinitionsContinued() {
	logDebug("<Function Definitions Continued> ::= <Empty> | <Function Definitions>")
}

func prodFunction() {
	logDebug("<Function> ::= function <Identifier> ( <Opt Parameter List> ) <Opt Declaration List> <Body>")
}

func prodOptParameterList() {
	logDebug("<Opt Parameter List> ::= <Parameter List> | <Empty>")
}

func prodParameterList() {
	logDebug("<Parameter List> ::= <Parameter> <Parameter List Continued>")
}

func prodParameterListContinued() {
	logDebug("<Parameter List Continued> ::= <Empty> | ,<Parameter List>")
}

func prodParameter() {
	logDebug("<Parameter> ::= <IDs > <Qualifier>")
}

func prodQualifier() {
	logDebug("<Qualifier> ::= integer | bool | real")
}

func prodBody() {
	logDebug("<Body> ::= { < Statement List> }")
}

func prodOptDeclarationList() {
	logDebug("<Opt Declaration List> ::= <Declaration List> | <Empty>")
}

func prodDeclarationList() {
	logDebug("<Declaration List> ::= <Declaration> ; <Declaration List Continued>")
}

func prodDeclarationListContinued() {
	logDebug("<Declaration List Continued> ::= <Empty> | <Declaration List>")
}

func prodDeclaration() {
	logDebug("<Declaration> ::= <Qualifier > <IDs>")
}

func prodIDs() {
	logDebug("<IDs> ::= <Identifier> <IDs Continued>")
}

func prodIDsContinued() {
	logDebug("<IDs Continued> ::= <Empty> | , <IDs>")
}

func prodStatementList() {
	logDebug("<Statement List> ::= <Statement> <Statement List Continued>")
}

func prodStatementListContinued() {
	logDebug("<Statement List Continued> ::= <Empty> | <Statement List>")
}

func prodStatement() {
	logDebug("<Statement> ::= <Compound> | <Assign> | <If> | <Return> | <Print> | <Scan> | <While>")
}

func prodCompound() {
	logDebug("<Compound> ::= { <Statement List> }")
}

func prodAssign() {
	logDebug("<Assign> ::= <Identifier> = <Expression> ;")
}

func prodIf() {
	logDebug("<If> ::= if ( <Condition> ) <Statement> <If Continued>")
}

func prodIfContinued() {
	logDebug("<If Continued> ::= endif | else <Statement> endif")
}

func prodReturn() {
	logDebug("<Return> ::= ret <Return Continued> ")
}

func prodReturnContinued() {
	logDebug("<Return Continued> ::= ; | <Expression> ;")
}

func prodPrint() {
	logDebug("<Print> ::= put ( <Expression>);")
}

func prodScan() {
	logDebug("<Scan> ::= get ( <IDs> );")
}

func prodWhile() {
	logDebug("<While> ::= while ( <Condition> ) <Statement>")
}

func prodCondition() {
	logDebug("<Condition> ::= <Expression> <Relop> <Expression>")
}

func prodRelop() {
	logDebug("<Relop> ::= == | != | > | < | <= | =>")
}

func prodExpression() {
	logDebug("<Expression> ::= <Term> <Expression Prime>")
}

func prodExpressionPrime() {
	logDebug("<Expression Prime> ::= <Expression Continued> <Expression Prime> | <Empty>")
}

func prodExpressionContinued() {
	logDebug("<Expression Continued> ::= + <Term> | - <Term>")
}

func prodTerm() {
	logDebug("<Term> ::= <Factor> <Term Prime>")
}

func prodTermPrime() {
	logDebug("<Term Prime> ::= <Term Continued> <Term Prime> | <Empty>")
}

func prodTermContinued() {
	logDebug("<Term Continued> ::= * <Factor> | / <Factor>")
}

func prodFactor() {
	logDebug("<Factor> ::= - <Primary> | <Primary>")
}

func prodPrimary() {
	logDebug("<Primary> ::= <Identifier> <Primary Continued> | <Integer> | ( <Expression> ) | <Real> | true | false")
}

func prodPrimaryContinued() {
	logDebug("<Primary Continued> ::= <Empty> | ( <IDs> )")
}

func prodEmpty() {
	logDebug("<Empty> ::= ε")
}

/* ---- The main attraction --------------------------- */

func syntaxAnalyzer(recordsIn []record) error {
	records = recordsIn // Store a copy for getNextRecord() usage
	logDebug("Hello from syntaxAnalyzer()!\n")
	prodRat23F() // Call initial production
	return nil
}
