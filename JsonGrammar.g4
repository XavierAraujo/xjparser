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
    : INT          #IntValue
    | FLOAT        #FloatValue
    | STRING       #StringValue
    | BOOL         #BoolValue
    | NULL         #NullValue
    | json_object  #ObjectValue
    | json_array   #ArrayValue
    ;


/* Lexer */

INT: [1-9]+[0-9]* ;
FLOAT: [1-9]+[0-9]* '.' [0-9]+ ;
STRING: ('"' (~["\t\r\n] | '\\' . )* '"' );
BOOL: 'true' | 'false' ;
NULL: 'null' ;
WS  : [ \t\r\n]+ -> skip ;
