grammar Test;

program : stat* EOF;
stat    : assignstat
        ;
assignstat  : ID '=' operation;

operation   : operation (MUL|DIV) operation #MulDiv
            | operation (ADD | MINUS) operation #AddMinus
            | value     #ValueExpr
            ;

value   : MINUS ? (number | ID);

number  : FLOAT
        | INT;

ID  : [a-zA-Z_] [a-zA-Z0-9_]* ;
FLOAT: [0-9]+ '.' [0-9]+;
INT : [0-9]+ ;

ADD     : '+';
MINUS   : '-';
MUL    : '*';
DIV     : '/';

WS  : [ \t\r\n]+ -> skip ;