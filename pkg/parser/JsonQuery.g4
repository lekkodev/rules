grammar JsonQuery;

query
   : NOT? SP? '(' query ')'                                                                         #parenExp
   | query SP AND_OPERATOR SP query                                                                 #andLogicalExp
   | query SP OR_OPERATOR SP query                                                                  #orLogicalExp
   | attrPath SP 'pr'                                                                               #presentExp
   | attrPath SP op=( EQ | NE | GT | LT | GE | LE | CO | SW | EW | IN ) SP value                    #compareExp
   | attrPath '(' (functionArg COMMA)* functionArg ')'                                              #callExp
   ;

NOT
   : 'not' | 'NOT' | '!'
   ;

AND_OPERATOR
   : 'and' | 'AND' | '&&'
   ;

OR_OPERATOR
   : 'or' | 'OR' | '||'
   ;

BOOLEAN
   : 'true' | 'false'
   ;

NULL
   : 'null'
   ;

IN:  'IN' | 'in';
EQ : 'eq' | 'EQ' | '==';
NE : 'ne' | 'NE' | '!=';
GT : 'gt' | 'GT' | '>';
LT : 'lt' | 'LT' | '<';
GE : 'ge' | 'GE' | '>=';
LE : 'le' | 'LE' | '<=';
CO : 'co' | 'CO';
SW : 'sw' | 'SW';
EW : 'ew' | 'EW';

attrPath
   : ATTRNAME subAttr?
   ;

subAttr
   : '.' attrPath
   ;

ATTRNAME
   : ALPHA ATTR_NAME_CHAR* ;

fragment ATTR_NAME_CHAR
   : '-' | '_' | ':' | DIGIT | ALPHA
   ;

fragment DIGIT
   : ('0'..'9')
   ;

fragment ALPHA
   : ( 'A'..'Z' | 'a'..'z' )
   ;

value
   : BOOLEAN           #boolean
   | NULL              #null
   | VERSION           #version
   | STRING            #string
   | DOUBLE            #double
   | '-'? INT EXP?     #long
   | listInts          #listOfInts
   | listDoubles       #listOfDoubles
   | listStrings       #listOfStrings
   | listBooleans      #listOfBooleans
   ;

VERSION
   : INT '.' INT '.' INT 
   ;

STRING
   : '"' (ESC | ~ ["\\])* '"'
   | '\'' (ESC | ~ ['\\])* '\''
   ;

listStrings
   : '[' subListOfStrings
   ;

subListOfStrings
   : STRING COMMA subListOfStrings
   | STRING ']';

fragment ESC
   : '\\' (["\\/bfnrt] | UNICODE)
   ;

fragment UNICODE
   : 'u' HEX HEX HEX HEX
   ;

fragment HEX
   : [0-9a-fA-F]
   ;

DOUBLE
   : '-'? INT '.' [0-9] + EXP?
   ;

listDoubles
   : '[' subListOfDoubles
   ;

subListOfDoubles
   : DOUBLE COMMA subListOfDoubles
   | DOUBLE ']';

listInts
   : '[' subListOfInts
   ;

subListOfInts
   : INT COMMA subListOfInts
   | INT ']';

listBooleans
   : '[' subListOfBooleans
   ;

subListOfBooleans
   : BOOLEAN COMMA subListOfBooleans
   | BOOLEAN ']';

functionArg
   : query
   | attrPath
   | value
   ;

// INT no leading zeros.
INT
   : '0' | [1-9] [0-9]*
   ;

// EXP we use "\-" since "-" means "range" inside [...]
EXP
   : [Ee] [+\-]? INT
   ;

NEWLINE
   : '\n' ;

COMMA
   : ',' ' '*;
SP
   : ' ' NEWLINE*
   ;
