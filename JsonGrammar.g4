grammar JsonGrammar;

prog: expr;

expr: json_expr EOF;

json_expr: json_object #JsonObjectExpr
    | json_array  #JsonArrayExpr
    ;


/* Parser */

json_key_value: (key=STRING ':' value=json_value);
json_object: '{' (json_key_value (',' json_key_value)*)? '}';
json_array: '[' (json_value (',' json_value)*)? ']';
json_value
    : INT         #IntValue
    | FLOAT       #FloatValue
    | STRING      #StringValue
    | BOOL        #BoolValue
    | NULL        #NullValue
    | json_object #ObjectValue
    | json_array  #ArrayValue
    ;


/* Lexer */

INT: INT_NUMBER ;
FLOAT: FLOAT_NUMBER | (FLOAT_NUMBER [Ee] [\-+]? [0-9]+) | (INT_NUMBER [Ee] [\-+]? [0-9]+) ;
STRING: '"' ( ~["\\\t\r\n] | '\\' [btnrf"\\/] | HEX_VALUE )* '"';
BOOL: 'true' | 'false' ;
NULL: 'null' ;
WS  : [ \t\r\n]+ -> skip ;

fragment HEX_VALUE: '\\u' HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR;
fragment HEX_CHAR: [0-9-A-Fa-f];
fragment INT_NUMBER:  '-'* ([0] | ([1-9]+[0-9]*)) ; 
fragment FLOAT_NUMBER: '-'* [0-9]+ '.' [0-9]+ ;
