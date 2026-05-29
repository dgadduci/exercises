grammar Test;

program : stat* EOF;
stat    : assignstat
        ;
assignstat  : ID '=' operation;

operation   : operation (MUL|DIV) operation
            | operation (ADD | MINUS) operation
            | value
            ;

value   : MINUS ? (number | ID);

number  : FLOAT
        | INT;

ID  : [a-zA-Z_]+ [0-9]* ;
FLOAT: [0-9]+ '.' [0-9]+;
INT : [0-9]+ ;

ADD     : '+';
MINUS   : '-';
MUL    : '*';
DIV     : '/';

WS  : [ \t\r\n]+ -> skip ;