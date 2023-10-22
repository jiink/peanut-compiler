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

/* ---- Productions ----------------------------------- */

func rulePrimary() {
	logDebug("<Primary> ::= <Identifier> | <Integer> | <Identifier> ( <IDs> ) | ( <Expression> ) | <Real> | true | false")
	// Too hard!
}

/* ---- The main attraction --------------------------- */

func syntaxAnalyzer(recordsIn []record) error {
	records = recordsIn // Store a copy for getNextRecord() usage
	logDebug("Hello from syntaxAnalyzer()!\n")
	return nil
}
