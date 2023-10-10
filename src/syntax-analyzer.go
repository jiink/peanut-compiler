package main

//---- Variables ------------------------------------------------------------------------
/////////////////////////////////////////////////////////////////////////////////////////

var records []record

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

/* ---- The main attraction --------------------------- */

func syntaxAnalyzer(recordsIn []record) error {
	records = recordsIn // Store a copy for getNextRecord() usage
	logDebug("Hello from syntaxAnalyzer()!\n")
	return nil
}
