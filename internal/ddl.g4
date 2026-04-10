grammar ddl;
/

// Parser rules
parse
    : createTable* EOF
    ;

// CREATE TABLE statement
createTable
    : CREATE TABLE tableName '(' columnDefList ')' SEMICOLON
    ;

// Table name
tableName
    : IDENTIFIER
    ;

// Column definitions — supports optional leading commas for formatting
columnDefList
    : columnDef (COMMA columnDef)*   // first column mandatory, commas between columns
    ;

// Column definition
columnDef
    : IDENTIFIER dataType
    ;

// Data types
dataType
    : NUMBER ('(' INT (',' INT)? ')')?   // e.g., NUMBER or NUMBER(10,2)
    | VARCHAR2 '(' INT ')'               // e.g., VARCHAR2(50)
    ;

// Lexer rules
CREATE   : [cC][rR][eE][aA][tT][eE];
TABLE    : [tT][aA][bB][lL][eE];
NUMBER   : [nN][uU][mM][bB][eE][rR];
VARCHAR2 : [vV][aA][rR][cC][hH][aA][rR] '2';
DATE     : [Dd][Aa][Tt][Ee];
SEMICOLON: ';';
COMMA    : ',';

// Identifiers and literals
IDENTIFIER
    : [a-zA-Z_][a-zA-Z_0-9]*
    ;

INT
    : [0-9]+
    ;

// Skip whitespace and newlines
WS
    : [ \t\r\n]+ -> skip
    ;