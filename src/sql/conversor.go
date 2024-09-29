package sql

import (
	"strconv"
	"strings"

	"sqlToKeyValue/src/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

//Contains auxiliary methods to convert from SQL to CRDTs

//Note: This method returns CRDTType_COUNTER for "counter",
//even if it should be a bounded counter (CRDTType_FATCOUNTER)
/*
func DataTypeSQLToCRDTType(dataTypeString string) proto.CRDTType {
	switch dataTypeString {
	case "counter":
		return proto.CRDTType_COUNTER
	case "integer":
		return proto.CRDTType_LWWREG
	case "boolean":
		return proto.CRDTType_FLAG_EW
	case "varchar":
		return proto.CRDTType_LWWREG
	case "date":
		return proto.CRDTType_LWWREG
	}
	return proto.CRDTType_LWWREG
}

func SQLUpdateToCRDTUpdate(dataTypeString string, sqlValue parser.IConstantContext) crdt.UpdateArguments {
	switch dataTypeString {
	case "counter":
		value, _ := strconv.Atoi(sqlValue.INT().GetText())
		return crdt.Increment{Change: int32(value)}
	case "integer":
		return crdt.SetValue{NewValue: sqlValue.INT().GetText()}
	case "boolean":
		defaultBool := strings.ToLower(sqlValue.BOOL().GetText())
		if defaultBool == "true" {
			return crdt.EnableFlag{}
		}
		return crdt.DisableFlag{}
	case "varchar":
		allStrings := sqlValue.AllSTRING()
		fullString := allStrings[0].GetText()
		for _, str := range allStrings[1:] {
			fullString += " " + str.GetText()
		}
		return crdt.SetValue{NewValue: fullString}
	case "date":
		return crdt.SetValue{NewValue: sqlValue.DATE().GetText()}
	}
	return nil
}*/

func ParsedTypeToSQLDatatype(dataTypeString string) SQLDatatype {
	switch dataTypeString {
	case "counter":
		return COUNTER
	case "integer":
		return INTEGER
	case "boolean":
		return BOOLEAN
	case "varchar":
		return VARCHAR
	case "date":
		return DATE
	}
	return UNSUPPORTED
}

func SQLUpdateToString(sqlValue parser.IConstantContext) string {
	if date := sqlValue.DATE(); date != nil {
		return date.GetText()
	}
	if boolean := sqlValue.BOOL(); boolean != nil {
		return boolean.GetText()
	}
	if integer := sqlValue.INT(); integer != nil {
		return integer.GetText()
	}
	if strings := sqlValue.AllSTRING(); strings != nil {
		fullString := strings[0].GetText()
		for _, str := range strings[1:] {
			fullString += " " + str.GetText()
		}
		return fullString
	}
	//Float, not currently supported
	return sqlValue.FLOAT().GetText()
}

func StringOperatorToCondOp(opType antlr.Token) CondType {
	switch opType.GetText() {
	case "=":
		return EQUAL
	case ">":
		return HIGHER
	case "<":
		return LOWER
	case ">=":
		return HIGHER_EQ
	case "<=":
		return LOWER_EQ
	default: //!=
		return NOT_EQ
	}
}

func MakeConstantFromContext(ctx parser.IConstantContext) Constant {
	if allStrings := ctx.AllSTRING(); len(allStrings) > 0 {
		fullString := allStrings[0].GetText()
		for _, str := range allStrings[1:] {
			fullString += " " + str.GetText()
		}
		return &StringConstant{Value: fullString}
	} else if number := ctx.INT(); number != nil {
		value, _ := strconv.ParseInt(number.GetText(), 10, 64)
		return &IntConstant{Value: int(value)}
	} else if number := ctx.FLOAT(); number != nil {
		text := number.GetText()
		value, _ := strconv.ParseFloat(text, 64)
		dotPos := strings.Index(text, ".")
		return &FloatConstant{Value: value, Precision: len(text) - dotPos - 1}
	} else { //Date
		dateText := ctx.DATE().GetText()
		date := &DateConstant{}
		date.FromString(dateText)
		return date
	}
}

func MakeNameableFromContext(ctx parser.INameableContext) Nameable {
	if name := ctx.Name(); name != nil {
		return &Name{Name: name.STRING().GetText()}
	}
	field := ctx.Field()
	return &Field{Name: field.Name(0).STRING().GetText(), FieldType: GetDateFieldType(field.Name(1).STRING().GetText())}
}

func SQLStringToListener(sqlCode string) (listener parser.ViewSQLListener) {
	lexer := parser.NewViewSQLLexer(antlr.NewInputStream(sqlCode))
	firstToken, secondToken := lexer.NextToken(), lexer.NextToken()
	lexer.SetInputStream(antlr.NewInputStream(sqlCode)) //Reset the lexer
	parse := parser.NewViewSQLParser(antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel))
	switch firstToken.GetText() {
	case "CREATE":
		switch secondToken.GetText() {
		case "VIEW":
			listener = CreateViewSQLListener(parse)
			info := ViewInfo{Listen: listener.(*MyViewSQLListener)} //TODO: Handle this one
			antlr.ParseTreeWalkerDefault.Walk(info.Listen, parse.Start())
		case "INDEX":
			listener = MakeCreateIndexListener(parse)
			antlr.ParseTreeWalkerDefault.Walk(listener, parse.Start())
		default: //AW/RW/LWW, so it is a CREATE TABLE
			listener = MakeCreateTableListener(parse)
			antlr.ParseTreeWalkerDefault.Walk(listener, parse.Start())
		}
	case "DROP":
		listener = MakeDropTableListener(parse)
		antlr.ParseTreeWalkerDefault.Walk(listener, parse.Start())
	case "DELETE":
		listener = MakeDeleteListener(parse)
		antlr.ParseTreeWalkerDefault.Walk(listener, parse.Start())
	case "INSERT":
		listener = MakeInsertListener(parse)
		antlr.ParseTreeWalkerDefault.Walk(listener, parse.Start())
	case "UPDATE":
		listener = MakeUpdateListener(parse)
		antlr.ParseTreeWalkerDefault.Walk(listener, parse.Start())
	case "SELECT":
		listener = MakeQueryListener(parse)
		antlr.ParseTreeWalkerDefault.Walk(listener, parse.Start())
	}
	return nil
}
