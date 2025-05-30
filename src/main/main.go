package main

import (
	"fmt"
	"strings"

	"sqlToKeyValue/src/parser"

	"sqlToKeyValue/src/sql"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

//TODO: Consider giving support for Q14 to be done with one single view

var q3 = `CREATE VIEW (Q3, views) WITH DATE IN [1995-03-01:1995-03-31] AS
			SELECT PRIMARY KEY(Orderkey), SUM(Extendedprice * (1 - Discount))
			AS Revenue, Orderdate, Shippriority
			FROM Lineitem
			WHERE Orderdate < [DATE] AND Shipdate > [DATE]
			GROUP BY Orderkey, Orderdate, Shippriority
			GROUP BY [DATE], Mktsegment
			ORDER BY Revenue DESC, Orderdate ASC
			LIMIT 10`

var q5 = `CREATE VIEW (Q5, views) WITH REGION IN {"AFRICA", "AMERICA", "ASIA", "EUROPE", "MIDDLE EAST"}, YEAR IN [1993:1997] AS
			SELECT N_Name, SUM(Extendedprice * (1 - Discount)) AS Revenue
			FROM Lineitem
			WHERE C_Nationkey = S_Nationkey AND R_Name = [REGION] AND Orderdate.year = [YEAR]
			GROUP BY N_Name
			GROUP BY [REGION], [YEAR]
			ORDER BY Revenue DESC`

var q11_v1 = `CREATE VIEW (Q11_p1, views) WITH NATION IN {"ALGERIA", "ARGENTINA", "BRAZIL", "CANADA", "EGYPT", 
				"ETHIOPIA", "FRANCE", "GERMANY", "INDIA", "INDONESIA", "IRAN", "IRAQ", "JAPAN", "JORDAN", "KENYA", 
				"MOROCCO", "MOZAMBIQUE", "PERU", "CHINA", "ROMANIA", "SAUDI ARABIA", "VIETNAM", "RUSSIA", "UNITED KINGDOM", 
				"UNITED STATES"} AS
			SELECT PRIMARY KEY(Partkey), SUM(Supplycost * Availqty) AS Value
			FROM Partsupp
			WHERE N_Name = [NATION]
			GROUP BY Partkey
			GROUP BY [NATION]
			ORDER BY Value DESC
			LIMIT 100`

var q11_v2 = `CREATE VIEW (Q11_p2, views) WITH NATION IN {"ALGERIA", "ARGENTINA", "BRAZIL", "CANADA", "EGYPT", 
					"ETHIOPIA", "FRANCE", "GERMANY", "INDIA", "INDONESIA", "IRAN", "IRAQ", "JAPAN", "JORDAN", "KENYA", 
					"MOROCCO", "MOZAMBIQUE", "PERU", "CHINA", "ROMANIA", "SAUDI ARABIA", "VIETNAM", "RUSSIA", "UNITED KINGDOM", 
					"UNITED STATES"} AS
				SELECT SUM(Supplycost * Availqty) AS Value
				FROM Partsupp
				WHERE N_Name = [NATION]
				GROUP BY [NATION]`

var q14_v1 = `CREATE VIEW (Q14_all, views) WITH YEAR IN [1993:1997], QUARTER IN [1:4] AS
			SELECT SUM(Extendedprice * (1 - Discount)) AS Revenue
			FROM Lineitem
			WHERE Shipdate.year = [YEAR] AND Shipdate.quarter = [QUARTER]
			GROUP BY [YEAR], [QUARTER]`

var q14_v2 = `CREATE VIEW (Q14_promo, views) WITH YEAR IN [1993:1997], QUARTER IN [1:4] AS
			SELECT SUM(Extendedprice * (1 - Discount)) AS Revenue
			FROM Lineitem
			WHERE P_Type startsWith "PROMO" AND Shipdate.year = [YEAR] AND Shipdate.quarter = [QUARTER]
			GROUP BY [YEAR], [QUARTER]`

var q15 = `CREATE VIEW (Q15, views) WITH YEAR IN [1993:1997], QUARTER IN [1:4] AS
			SELECT PRIMARY KEY(Suppkey), S_Name, S_Address, S_Phone, SUM(Extendedprice * (1 - Discount)) AS Revenue
			FROM Lineitem
			WHERE Shipdate.year = [YEAR] AND Shipdate.quarter = [QUARTER]
			GROUP BY Suppkey, S_Name, S_Address, S_Phone
			GROUP BY [YEAR], [QUARTER]
			ORDER BY Revenue DESC, Suppkey ASC
			LIMIT 1`

var q18 = `CREATE VIEW (Q18, views) WITH QUANTITY IN [312:315] AS
			SELECT C_Name, PRIMARY KEY(Custkey), Orderkey, Orderdate, Totalprice, Totalquantity
			FROM Orders
			WHERE Totalquantity > [QUANTITY]
			GROUP BY [QUANTITY]
			ORDER BY Totalprice DESC, Orderdate ASC
			LIMIT 100`

var base = `CREATE VIEW (Q5, views) AS
			SELECT
			FROM
			WHERE
			GROUP BY
			GROUP BY
			ORDER BY
			LIMIT`

var table = `CREATE AW TABLE Persons (
			PersonID integer PRIMARY KEY,
			LastName varchar,
			FirstName varchar,
			Address varchar DEFAULT "unknown",
			City varchar DEFAULT "unknown",
			NFriends counter DEFAULT 0 CHECK >= 0,
			HasPhone boolean DEFAULT true,
			WorkID integer FOREIGN KEY REFERENCES Work (WorkID)
			)`

var index = `CREATE INDEX CityPersons ON Persons (City)`

var drop = `DROP TABLE Persons`
var drop2 = `DROP TABLE SomeDifferentName`

var delete = `DELETE FROM Persons WHERE PersonID = 1`

var insert_1 = `INSERT INTO Persons (PersonID, FirstName, LastName, Address, WorkID)
				VALUES (2, "First", "Last", "Lost address", 3)`

var insert_2 = `INSERT INTO Persons
				VALUES (1, "LN", "FN", "Somewhere", "Somewhere City", 5, true, 2)`

var update = `UPDATE Persons
			  SET Address = "Beacon Hill", City = "Kansas City"
			  WHERE PersonID = 1`

var update2 = `UPDATE Persons
			   SET Address = "Woahhhhh", City = "Wowwwwww"
			   WHERE PersonID >= 995 + (1 * 5 + (-10 + 10)) AND PersonID > NFriends * 10 AND FirstName startsWith "A"`

var query = `SELECT FirstName, LastName, City, NFriends
			 FROM Persons
			 WHERE PersonID >= 0 AND PersonID <= 10`

var queryAll = `SELECT *
				FROM Persons
				WHERE City contains "City"`

//Other Select important clauses not supported atm:
//Group By, Order By, Limit

func testDrop(sqlCode string) {
	is := antlr.NewInputStream(sqlCode)
	lexer := parser.NewViewSQLLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parse := parser.NewViewSQLParser(stream)
	listener := sql.MakeDropTableListener(parse)
	antlr.ParseTreeWalkerDefault.Walk(listener, parse.Start())
	ignore(parse)
}

func testCreateTable(sqlCode string) {
	is := antlr.NewInputStream(sqlCode)
	lexer := parser.NewViewSQLLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parse := parser.NewViewSQLParser(stream)
	listener := sql.MakeCreateTableListener(parse)
	antlr.ParseTreeWalkerDefault.Walk(listener, parse.Start())
}

func testMultipleListeners(sqlCode string) {
	is := antlr.NewInputStream(sqlCode)
	lexer := parser.NewViewSQLLexer(is)
	firstToken, secondToken := lexer.NextToken(), lexer.NextToken()
	lexer.SetInputStream(antlr.NewInputStream(sqlCode)) //To reset the lexer
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parse := parser.NewViewSQLParser(stream)
	switch firstToken.GetText() {
	case "CREATE":
		switch secondToken.GetText() {
		case "VIEW":
			listener := sql.CreateViewSQLListener(parse)
			info := sql.ViewInfo{Listen: listener}
			antlr.ParseTreeWalkerDefault.Walk(info.Listen, parse.Start())
		case "INDEX":
			listener := sql.MakeCreateIndexListener(parse)
			antlr.ParseTreeWalkerDefault.Walk(listener, parse.Start())
		default: //AW/RW/LWW, so it is a CREATE TABLE
			listener := sql.MakeCreateTableListener(parse)
			antlr.ParseTreeWalkerDefault.Walk(listener, parse.Start())
		}
	case "DROP":
		listener := sql.MakeDropTableListener(parse)
		antlr.ParseTreeWalkerDefault.Walk(listener, parse.Start())
	case "DELETE":
		listener := sql.MakeDeleteListener(parse)
		antlr.ParseTreeWalkerDefault.Walk(listener, parse.Start())
	case "INSERT":
		listener := sql.MakeInsertListener(parse)
		antlr.ParseTreeWalkerDefault.Walk(listener, parse.Start())
	case "UPDATE":
		listener := sql.MakeUpdateListener(parse)
		antlr.ParseTreeWalkerDefault.Walk(listener, parse.Start())
	case "SELECT":
		listener := sql.MakeQueryListener(parse)
		antlr.ParseTreeWalkerDefault.Walk(listener, parse.Start())
	}
}

// Done: Drop, CreateTable, Insert, CreateIndex, Update, Delete
// TODO: View
func main() {
	//names := []string{"First insert", "Second insert", "delete", "drop", "drop2", "update", "update2",
	//"create table", "create index", "view", "query", "query with *"}
	//texts := []string{insert_1, insert_2, delete, drop, drop2, update, update2, table, index, q3, query, queryAll}
	//testDrop(drop)
	//testCreateTable(table)
	names := []string{"query", "query with *"}
	texts := []string{query, queryAll}
	for i, text := range texts {
		fmt.Println(names[i])
		testMultipleListeners(text)
		fmt.Println()
		fmt.Println()
	}
	//start()

	//doViews()
}

func testListenerType(list parser.ViewSQLListener) {

}

/*
func doViews() {
	queries := []string{q3, q5, q11_v1, q11_v2, q14_v1, q14_v2, q15, q18}
	queriesNames := []string{"Q3", "Q5", "Q11_P1", "Q11_P2", "Q14_P1", "Q14_P2", "Q15", "Q18"}

	infos := make([]*sql.ViewInfo, len(queries))
	queries, queriesNames = queries[:], queriesNames[:]
	for i, query := range queries {
		is := antlr.NewInputStream(query)

		lexer := parser.NewViewSQLLexer(is)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		parse := parser.NewViewSQLParser(stream)
		listener := sql.CreateViewSQLListener(parse)
		info := sql.ViewInfo{Listen: listener}
		infos[i] = &info

		//parse.Start: Makes the parser start on start. Could call any other rule of the grammar.
		//antlr.ParseTreeWalkerDefault.Walk(parser.BaseViewSQLListener{}, parse.Start())
		antlr.ParseTreeWalkerDefault.Walk(info.Listen, parse.Start())

		//could call parse.toStringTree() to show a parse tree

		fmt.Printf("***%s***\n", queriesNames[i])
		viewType := info.GetViewType()
		fmt.Printf("%d (%s)\n", viewType, sql.CrdtTypeToString[viewType])
		if viewType == sql.MAP {
			fmt.Println(mapViewTypeToString(info.GetViewMapFullType()))
		} else if viewType == sql.TOPK || viewType == sql.TOPSUM {
			fmt.Println(topViewTypeToString(info.GetViewTopFullType()))
		}
		info.PrepareInfo()
		fmt.Println()
		info.PrintPreparedQuery()
		return

	}
	queries = []string{q3, q5, q14_v1, q14_v2, q15}
	queriesNames = []string{"Q3", "Q5", "Q14_P1", "Q14_P2", "Q15"}
	items := []sql.Lineitem{sql.L1, sql.L2, sql.L3, sql.L4, sql.L5, sql.L6}
	infosItems := []*sql.ViewInfo{infos[0], infos[1], infos[4], infos[5], infos[6]}
	queries, queriesNames, infosItems = queries[:], queriesNames[:], infosItems[:]
	for i := range queries {
		fmt.Printf("***%s***\n", queriesNames[i])
		for _, item := range items {
			upd := item.ToMapUpdate()
			fmt.Printf("%+v\n", item)
			fmt.Printf("%v\n", item)
			fmt.Println(*infosItems[i])
			infosItems[i].GenerateViewUpdate(upd)
			fmt.Println()
		}
	}
	queries = []string{q11_v1, q11_v2}
	queriesNames = []string{"Q11_P1", "Q11_P2"}
	itemsPS := []sql.PartSupp{sql.P1, sql.P2, sql.P3, sql.P4, sql.P5, sql.P6}
	infosItems = []*sql.ViewInfo{infos[2], infos[3]}
	for i := range queries {
		fmt.Printf("***%s***\n", queriesNames[i])
		for _, item := range itemsPS {
			upd := item.ToMapUpdate()
			fmt.Printf("%+v\n", item)
			fmt.Printf("%v\n", item)
			fmt.Println(*infosItems[i])
			infosItems[i].GenerateViewUpdate(upd)
			fmt.Println()
		}
	}
	queries = []string{q18}
	queriesNames = []string{"Q18"}
	itemsO := []sql.Orders{sql.O1, sql.O2, sql.O3, sql.O4, sql.O5, sql.O6}
	infosItems = []*sql.ViewInfo{infos[7]}
	for i := range queries {
		fmt.Printf("***%s***\n", queriesNames[i])
		for _, item := range itemsO {
			upd := item.ToMapUpdate()
			fmt.Printf("%+v\n", item)
			fmt.Printf("%v\n", item)
			fmt.Println(*infosItems[i])
			infosItems[i].GenerateViewUpdate(upd)
			fmt.Println()
		}
	}
}
*/

func mapViewTypeToString(mapType map[string]int) string {
	var sb strings.Builder
	sb.WriteString("MAP: MAP[")
	for key, typeV := range mapType {
		sb.WriteString(fmt.Sprintf("%s -> %s, ", key, sql.CrdtTypeToString[typeV]))
	}
	sb.WriteString("]")
	return sb.String()
}

func topViewTypeToString(topType sql.TopKType) string {
	var sb strings.Builder
	sb.WriteString("TOP: TOP[")
	sb.WriteString("ID -> [")
	/*for _, key := range topType.Id {
		sb.WriteString(key + ", ")
	}
	*/
	sb.WriteString(topType.Id)
	sb.WriteString("], Value -> ")
	sb.WriteString(topType.Value)
	sb.WriteString("]")
	return sb.String()
}

func ignore(any interface{}) {

}

/*
var q3 = `CREATE VIEW (Q3, views) AS
			SELECT Orderkey, SUM(Extendedprice * (1 - Discount))
			AS Revenue, Orderdate, Shippriority
			FROM Lineitem
			WHERE Orderdate < [DATE] AND Shipdate > [DATE] AND [DATE] >= 1995-03-01 AND [DATE] <= 1995-03-31
			GROUP BY Orderkey, Orderdate, Shippriority
			GROUP BY [DATE], Mktsegment
			ORDER BY Revenue DESC, Orderdate ASC
			LIMIT 10`

var q5 = `CREATE VIEW (Q5, views) AS
			SELECT N_Name, SUM(Extendedprice * (1 - Discount)) AS Revenue
			FROM Lineitem
			WHERE C_Nationkey = S_Nationkey AND R_Name = [REGION] AND Orderdate.year = [YEAR] AND [YEAR] >= 1993 AND [YEAR] <= 1997
			GROUP BY N_Name
			GROUP BY [REGION], [YEAR]
			ORDER BY Revenue DESC`

var q11_v1 = `CREATE VIEW (Q11_p1, views) AS
			SELECT Partkey, SUM(Supplycost + Availqty) AS Value
			FROM Partsupp
			WHERE N_Name = [NATION]
			GROUP BY Partkey
			GROUP BY [NATION]
			ORDER BY Value DESC
			LIMIT 100`

var q11_v2 = `CREATE VIEW (Q11_p2, views) AS
				SELECT SUM(Supplycost + Availqty) AS Value
				FROM Partsupp
				WHERE N_Name = [NATION]
				GROUP BY [NATION]`

var q14_v1 = `CREATE VIEW (Q14_all, views) AS
			SELECT SUM(Extendedprice * (1 - Discount)) AS Revenue
			FROM Lineitem
			WHERE Shipdate.year = [YEAR] AND Shipdate.quarter = [QUARTER] AND Shipdate.year >= 1993 AND Shipdate.year <= 1997
			GROUP BY [YEAR], [QUARTER]`

var q14_v2 = `CREATE VIEW (Q14_promo, views) AS
			SELECT SUM(Extendedprice * (1 - Discount)) AS Revenue
			FROM Lineitem
			WHERE P_Type startsWith "PROMO" AND Shipdate.year = [YEAR] AND Shipdate.quarter = [QUARTER] AND Shipdate.year >= 1993 AND Shipdate.year <= 1997
			GROUP BY [YEAR], [QUARTER]`

var q15 = `CREATE VIEW (Q15, views) AS
			SELECT Suppkey, S_Name, S_Address, S_Phone, SUM(Extendedprice * (1 - Discount)) AS Revenue
			FROM Lineitem
			WHERE Shipdate.year = [YEAR] AND Shipdate.quarter = [QUARTER] AND Shipdate.year >= 1993 AND Shipdate.year <= 1997
			GROUP BY Suppkey, S_Name, S_Address, S_Phone
			GROUP BY [YEAR], [QUARTER]
			ORDER BY Revenue DESC, Suppkey ASC
			LIMIT 1`

var q18 = `CREATE VIEW (Q18, views) AS
			SELECT C_Name, Custkey, Orderkey, Orderdate, Totalprice, Totalquantity
			FROM Orders
			WHERE Totalquantity > [QUANTITY] AND [QUANTITY] >= 312 AND [QUANTITY] <= 315
			GROUP BY [QUANTITY]
			ORDER BY Totalprice DESC, Orderdate ASC
			LIMIT 100`

var base = `CREATE VIEW (Q5, views) AS
			SELECT
			FROM
			WHERE
			GROUP BY
			GROUP BY
			ORDER BY
			LIMIT`
*/
