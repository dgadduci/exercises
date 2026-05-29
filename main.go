package main

import (
	"fmt"

	// Runtime de ANTLR para Go
	"github.com/antlr4-go/antlr/v4"

	// Parser generado automáticamente desde Test.g4
	"github.com/dgadduci/exercises/parser"
)

func main() {

	// Crea un stream de entrada en memoria.
	// Más adelante esto vendrá de un archivo.
	input := antlr.NewInputStream(`
a=1
b=2
c=a+b*a
`)

	// El lexer transforma caracteres en tokens.
	//
	// Ejemplo:
	// a=1+2
	//
	// se convierte en:
	//
	// ID("a")
	// '='
	// INT("1")
	// '+'
	// INT("2")
	lexer := parser.NewTestLexer(input)

	// Buffer de tokens consumido por el parser.
	tokens := antlr.NewCommonTokenStream(
		lexer,
		antlr.TokenDefaultChannel,
	)

	// Crea una instancia del parser.
	p := parser.NewTestParser(tokens)

	// Indica la regla inicial de la gramática.
	//
	// Como tu gramática tiene:
	//
	// program : stat* EOF ;
	//
	// esta llamada devuelve el nodo raíz del árbol.
	tree := p.Program()

	// Imprime el árbol sintáctico generado.
	fmt.Println(tree.ToStringTree(nil, p))

	fmt.Println("Tipo:", fmt.Sprintf("%T", tree))
	fmt.Println("Texto:", tree.GetText())
}
