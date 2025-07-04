package sql

// TODO: Consider giving support for Q14 to be done with one single view
var Q3 = `CREATE VIEW (Q3, views) WITH DATE IN [1995-03-01:1995-03-31] AS
			SELECT PRIMARY KEY(Orderkey), SUM(Extendedprice * (1 - Discount))
			AS Revenue, Orderdate, Shippriority
			FROM Lineitem
			WHERE Orderdate < [DATE] AND Shipdate > [DATE]
			GROUP BY Orderkey, Orderdate, Shippriority
			GROUP BY [DATE], Mktsegment
			ORDER BY Revenue DESC, Orderdate ASC
			LIMIT 10`

var Q5 = `CREATE VIEW (Q5, views) WITH REGION IN {"AFRICA", "AMERICA", "ASIA", "EUROPE", "MIDDLE EAST"}, YEAR IN [1993:1997] AS
			SELECT N_Name, SUM(Extendedprice * (1 - Discount)) AS Revenue
			FROM Lineitem
			WHERE C_Nationkey = S_Nationkey AND R_Name = [REGION] AND Orderdate.year = [YEAR]
			GROUP BY N_Name
			GROUP BY [REGION], [YEAR]
			ORDER BY Revenue DESC`

var Q11_v1 = `CREATE VIEW (Q11_p1, views) WITH NATION IN {"ALGERIA", "ARGENTINA", "BRAZIL", "CANADA", "EGYPT", 
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

var Q11_v2 = `CREATE VIEW (Q11_p2, views) WITH NATION IN {"ALGERIA", "ARGENTINA", "BRAZIL", "CANADA", "EGYPT", 
					"ETHIOPIA", "FRANCE", "GERMANY", "INDIA", "INDONESIA", "IRAN", "IRAQ", "JAPAN", "JORDAN", "KENYA", 
					"MOROCCO", "MOZAMBIQUE", "PERU", "CHINA", "ROMANIA", "SAUDI ARABIA", "VIETNAM", "RUSSIA", "UNITED KINGDOM", 
					"UNITED STATES"} AS
				SELECT SUM(Supplycost * Availqty) AS Value
				FROM Partsupp
				WHERE N_Name = [NATION]
				GROUP BY [NATION]`

var Q14_v1 = `CREATE VIEW (Q14_all, views) WITH YEAR IN [1993:1997], QUARTER IN [1:4] AS
			SELECT SUM(Extendedprice * (1 - Discount)) AS Revenue
			FROM Lineitem
			WHERE Shipdate.year = [YEAR] AND Shipdate.quarter = [QUARTER]
			GROUP BY [YEAR], [QUARTER]`

var Q14_v2 = `CREATE VIEW (Q14_promo, views) WITH YEAR IN [1993:1997], QUARTER IN [1:4] AS
			SELECT SUM(Extendedprice * (1 - Discount)) AS Revenue
			FROM Lineitem
			WHERE P_Type startsWith "PROMO" AND Shipdate.year = [YEAR] AND Shipdate.quarter = [QUARTER]
			GROUP BY [YEAR], [QUARTER]`

var Q15 = `CREATE VIEW (Q15, views) WITH YEAR IN [1993:1997], QUARTER IN [1:4] AS
			SELECT PRIMARY KEY(Suppkey), S_Name, S_Address, S_Phone, SUM(Extendedprice * (1 - Discount)) AS Revenue
			FROM Lineitem
			WHERE Shipdate.year = [YEAR] AND Shipdate.quarter = [QUARTER]
			GROUP BY Suppkey, S_Name, S_Address, S_Phone
			GROUP BY [YEAR], [QUARTER]
			ORDER BY Revenue DESC, Suppkey ASC
			LIMIT 1`

var Q18 = `CREATE VIEW (Q18, views) WITH QUANTITY IN [312:315] AS
			SELECT C_Name, PRIMARY KEY(Custkey), Orderkey, Orderdate, Totalprice, Totalquantity
			FROM Orders
			WHERE Totalquantity > [QUANTITY]
			GROUP BY [QUANTITY]
			ORDER BY Totalprice DESC, Orderdate ASC
			LIMIT 100`

var Base = `CREATE VIEW (Q5, views) AS
			SELECT
			FROM
			WHERE
			GROUP BY
			GROUP BY
			ORDER BY
			LIMIT`

var Table = `CREATE AW TABLE Persons (
			PersonID integer PRIMARY KEY,
			LastName varchar MW,
			FirstName varchar LWW,
			Address varchar DEFAULT "unknown" LWW,
			City varchar DEFAULT "unknown" MW,
			NFriends counter DEFAULT 0 CHECK >= 0,
			HasPhone boolean DEFAULT true EW,
			WorkID integer FOREIGN KEY REFERENCES Work (WorkID)
			)`

var Index = `CREATE INDEX CityPersons ON Persons (City)`

var Drop = `DROP TABLE Persons`
var Drop2 = `DROP TABLE SomeDifferentName`

var Delete = `DELETE FROM Persons WHERE PersonID = 1`

var Insert_1 = `INSERT INTO Persons (PersonID, FirstName, LastName, Address, WorkID)
				VALUES (2, "First", "Last", "Lost address", 3)`

var Insert_2 = `INSERT INTO Persons
				VALUES (1, "LN", "FN", "Somewhere", "Somewhere City", 5, true, 2)`

var Update = `UPDATE Persons
			  SET Address = "Beacon Hill", City = "Kansas City"
			  WHERE PersonID = 1`

var Update2 = `UPDATE Persons
			   SET Address = "Woahhhhh", City = "Wowwwwww"
			   WHERE PersonID >= 995 + (1 * 5 + (-10 + 10)) AND PersonID > NFriends * 10 AND FirstName startsWith "A"`

var Query = `SELECT FirstName, LastName, City, NFriends
			 FROM Persons
			 WHERE PersonID >= 0 AND PersonID <= 10`

var QueryAll = `SELECT *
				FROM Persons
				WHERE City contains "City"`
