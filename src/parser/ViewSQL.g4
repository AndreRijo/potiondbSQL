grammar ViewSQL;

//reserved words, symbols

ADD: '+';
SUB: '-';
MULT: '*';
DIV: '/';

SUM: 'SUM';
AVG: 'AVG';
MAX: 'MAX';
MIN: 'MIN';
COUNT: 'COUNT';

AND: 'AND';

ASC: 'ASC';
DESC: 'DESC';

EQUAL: '=';
HIGHER: '>';
HIGHER_EQUAL: '>=';
LOWER: '<';
LOWER_EQUAL: '<=';
NOT_EQUAL: '!=';

DOT: '.';
SEPARATOR: ',';
RANGE_SEP: ':';
LEFT_P: '(';
RIGHT_P: ')';
LEFT_SP: '[';
RIGHT_SP: ']';
INV_COMMA: '"';
LEFT_CURLY: '{';
RIGHT_CURLY: '}';

AS: 'AS';

CREATE: 'CREATE';
VIEW: 'VIEW';
SELECT: 'SELECT';
FROM: 'FROM';
WHERE: 'WHERE';
GROUP: 'GROUP';
BY: 'BY';
ORDER: 'ORDER';
LIMIT: 'LIMIT';
WITH: 'WITH';
IN: 'IN';
PRIMARY: 'PRIMARY';
KEY: 'KEY';

TABLE: 'TABLE';
INDEX: 'INDEX';
UPDATE: 'UPDATE';
DELETE: 'DELETE';
INSERT: 'INSERT';
INTO: 'INTO';
DROP: 'DROP';
SET: 'SET';
FOREIGN: 'FOREIGN';
REFERENCES: 'REFERENCES';
VALUES: 'VALUES';
CHECK: 'CHECK';
ON: 'ON';

AW: 'AW';
RW: 'RW';
LWW: 'LWW';
DEFAULT: 'DEFAULT';
INTEGER: 'integer';
COUNTER: 'counter';
BOOLEAN: 'boolean';
VARCHAR: 'varchar';
DATE_TYPE: 'date';


//Types

DATE: YEAR '-' MONTH '-' DAY;

fragment DAY: '0'[1-9] | [1-2][0-9] | '3'[0-1];
fragment MONTH: '0'[1-9] | '1'[1-2];
fragment YEAR: [0-9][0-9][0-9][0-9];

BOOL: 'true' | 'false';

//STRING: [a-zA-Z_] (' ' | [a-zA-Z_0-9])*;
STRING: [a-zA-Z_] [a-zA-Z_0-9]*;

//NUMBER: INT | FLOAT;

//POSITIVE_INT: [1-9] [0-9]*;

INT: '0' | '-'? [1-9] [0-9]*;

FLOAT: '-'? [0-9] '.' [0-9]*[1-9]+;

WHITESPACE: [ \r\n\t]+ -> skip;





name: STRING ;

constant: DATE | BOOL | INT | FLOAT | '"' STRING (WHITESPACE | STRING)* '"';

aggrFunc: 'SUM' | 'AVG' | 'MAX' | 'MIN';

field: name '.' name;       //This is only used for Dates.

parameter: '[' name ']';

nameable: name | field | parameter;

key: 'PRIMARY' 'KEY' '(' nameable ')';

//op: '*' | '/' | '+' | '-';

//value: nameable | constant;

//math: value (op value)* | ;

//math: value (op math)? | '(' value (op math) ')';

//Later if needed can add power too. Like this: math ^ math
math                                                
    : (nameable | constant)     # Value
    | SUB math                  # Minus
    | '(' math ')'              # Parentheses
    | math opType=('*' | '/') math     # MultOrDiv
    | math opType=('+' | '-') math     # AddOrSub
    ;

asClause: math 'AS' name;                       //DONE!

aggregation: aggrFunc '(' math ')' 'AS' name;   //DONE!

count: 'COUNT' '(' (nameable | '*') ')' 'AS' name;  //DONE!

calc: key | nameable | asClause | aggregation | count | all='*';    //Last one is only supported by queries, not views

comp: opType=('=' | '>' | '<' | '>=' | '<=' | '!=') | name;       //name: some function name (e.g.: contains). Yes the grammar is very generic on this.

condition: nameable comp math;

sortOrder: 'DESC' | 'ASC';

continuousRange: left=('['|']') (constant ':' constant) right=(']'|'[');

sparseRange: '{' constant (',' constant)* '}';

range: continuousRange | sparseRange;



create: 'CREATE' 'VIEW' '('  STRING ',' STRING ')';

//This STRING here is actually a parameter but without '[' and ']'
with: 'WITH' STRING 'IN' range (',' STRING 'IN' range)*;

select: 'SELECT' calc (',' calc)*;              //Done! Hopefully.

from: 'FROM' name (',' name)*;                  //Done!

where: 'WHERE' condition ('AND' condition)*;

groupby: 'GROUP' 'BY' nameable (',' nameable)*;

orderby: 'ORDER' 'BY' nameable sortOrder (',' nameable sortOrder)*;

limit: 'LIMIT' INT;     //Done!

view: create with? 'AS' select from where? firstGroupBy=groupby? secondGroupBy=groupby orderby? limit?;

//view: CREATE;

//start: view EOF;



check: 'CHECK' type=('>' | '<' | '>=' | '<=') constant;

foreignkey: 'FOREIGN' 'KEY' 'REFERENCES' tableName=name '('columnName=name')';

primarykey: 'PRIMARY' 'KEY';

constraint: primarykey | foreignkey | check;

columns: STRING type=(COUNTER|INTEGER|BOOLEAN|VARCHAR|DATE_TYPE) ('DEFAULT' constant)? constraint?;

createtable: 'CREATE' policy=('AW'|'RW'|'LWW') 'TABLE' name '(' (columns ',')* columns ')';


createindex: 'CREATE' 'INDEX' indexName=name 'ON' tableName=name '(' columnName=name ')';

drop: 'DROP' 'TABLE' name;


delete: 'DELETE' 'FROM' name where;


set: 'SET' (name '=' constant ',')* name '=' constant where;

update: 'UPDATE' name set;


values: 'VALUES' '(' (constant ',')* constant ')';

columnNames: '(' (name ',')* name ')';

insert: 'INSERT' 'INTO' name columnNames? values;


query: select from where?;  //TODO: Think later if we want to support group by, order by and limit.

statement: view | createtable | createindex | insert | update | delete | drop | query;

start: statement EOF;

/*
UPDATE Customers
SET ContactName = 'Alfred Schmidt', City= 'Frankfurt'
WHERE CustomerID = 1;

CREATE AW/RW/LWW TABLE Persons (
    PersonID int DEFAULT value constraint,
    LastName varchar(255) DEFAULT value constraint,
    FirstName varchar(255) DEFAULT value constraint,
    Address varchar(255) DEFAULT value constraint,
    City varchar(255) DEFAULT value constraint 
);

INSERT INTO table_name (column1, column2, column3, ...)
VALUES (value1, value2, value3, ...);
(or)
INSERT INTO table_name
VALUES (value1, value2, value3, ...);

DELETE FROM table_name WHERE condition;

DROP TABLE table_name
*/

/*
CREATE VIEW (ID, bucket) AS
SELECT calc (, calc)*
FROM name (, name)*
WHERE condition (and|or condition)*
GROUP BY nameable (, nameable)*
GROUP BY nameable (, nameable)*
ORDER BY nameable desc|asc (, nameable desc|asc)*
LIMIT number
*/