package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var fileName string
var buffer string
var sourceFile *os.File
var outputFile *os.File
var errorsFile *os.File

var scanner *bufio.Scanner

var err error

var la int

var tokenNames = [...]string{"", "programa", "if", "then", "else", "fi", "do", "until", "while", "read", "write", "float", "int", "bool", "true", "false", "not", "and", "or",
	"+", "-", "*", "/", "^", "<", "<=", ">", ">=", "==", "!=", "=", ";", ",", "(", ")", "{", "}", "error", "identificador", "eof", "comentario", "constante numerica", "palabra",
}

//TOKEN TYPES
const (
	TknProgram  = 1
	TknIf       = 2
	TknThen     = 3
	TknElse     = 4
	TknFi       = 5
	TknDo       = 6
	TknUntil    = 7
	TknWhile    = 8
	TknRead     = 9
	TknWrite    = 10
	TknFloat    = 11
	TknInt      = 12
	TknBool     = 13
	TknTrue     = 14
	TknFalse    = 15
	TknNot      = 16
	TknAnd      = 17
	TknOr       = 18
	TknSum      = 19
	TknSub      = 20
	TknMul      = 21
	TknDiv      = 22
	TknExp      = 23
	TknLess     = 24
	TknLessEq   = 25
	TknGreat    = 26
	TknGreatEq  = 27
	TknEq       = 28
	TknNotEq    = 29
	TknAssign   = 30
	TknSemi     = 31
	TknComma    = 32
	TknLeftPar  = 33
	TknRightPar = 34
	TknLeftBr   = 35
	TknRightBr  = 36
	TknError    = 37
	TknIdent    = 38
	TknEOF      = 39
	TknComment  = 40
	TknConst    = 41
	TknWord     = 42
)

var tokens []Token

//Token blaablabla
type Token struct {
	Type      int
	Name      string
	Attribute string
	Row       int
	Column    int
}

//Node para el arbol
type Node struct {
	name   string
	lexeme string
	Type   int
	childs []*Node
}

var root *Node

func readTokens() {
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		var auxToken Token
		auxToken.Type, err = strconv.Atoi(s[0])
		auxToken.Name = s[1]
		auxToken.Attribute = s[2]
		auxToken.Row, err = strconv.Atoi(s[3])
		auxToken.Column, err = strconv.Atoi(s[4])
		tokens = append(tokens, auxToken)

	}
}

func matchNode(Type int) *Node {
	var current *Node = new(Node)
	if tokens[la].Type == Type {

		current.name = tokens[la].Attribute
		current.Type = Type
		la++
	} else {
		errorsFile.WriteString("ERROR en linea " + strconv.Itoa(tokens[la].Row) + ", columna " + strconv.Itoa(tokens[la].Column) + ". ")
		errorsFile.WriteString("Se esperaba " + tokenNames[Type] + ". Se encontro " + tokens[la].Attribute + ".\n")
	}

	return current

}

func tipo() *Node {
	var current *Node = new(Node)
	current.name = "tipo"
	if tokens[la].Type == TknInt {
		current.childs = append(current.childs, matchNode(TknInt))
	} else if tokens[la].Type == TknBool {
		current.childs = append(current.childs, matchNode(TknBool))
	} else if tokens[la].Type == TknFloat {
		current.childs = append(current.childs, matchNode(TknFloat))
	}
	return current
}

func listaIdent() *Node {
	var current *Node = new(Node)
	current.name = "listaIdent"
	current.childs = append(current.childs, matchNode(TknIdent))
	for tokens[la].Type == TknComma {
		current.childs = append(current.childs, matchNode(TknComma))
		//current.childs = append(current.childs, listaIdent())
		current.childs = append(current.childs, matchNode(TknIdent))
	}
	return current
}

func declaracion() *Node {
	var current *Node = new(Node)
	current.name = "declaracion"
	current.childs = append(current.childs, tipo())
	current.childs = append(current.childs, listaIdent())
	current.childs = append(current.childs, matchNode(TknSemi))
	return current
}

func listaDeclaraciones() *Node {
	var current *Node = new(Node)
	current.name = "listaDeclaraciones"
	for tokens[la].Type == TknFloat || tokens[la].Type == TknInt || tokens[la].Type == TknBool {
		current.childs = append(current.childs, declaracion())
		//current.childs = append(current.childs, listaDeclaraciones())
	}
	return current
}

func listaSentencias() *Node {
	var current *Node = new(Node)
	current.name = "listaSentencias"
	for tokens[la].Type == TknIf || tokens[la].Type == TknWhile || tokens[la].Type == TknDo || tokens[la].Type == TknRead || tokens[la].Type == TknWrite || tokens[la].Type == TknLeftBr || tokens[la].Type == TknIdent {
		current.childs = append(current.childs, sentencia())
		//listaSentencias()
	}
	return current
}

func sentencia() *Node {
	var current *Node = new(Node)
	current.name = "sentencia"
	if tokens[la].Type == TknIf {
		current.childs = append(current.childs, seleccion())
	} else if tokens[la].Type == TknWhile {
		current.childs = append(current.childs, iteracion())
	} else if tokens[la].Type == TknDo {
		current.childs = append(current.childs, repeticion())
	} else if tokens[la].Type == TknRead {
		current.childs = append(current.childs, sentRead())
	} else if tokens[la].Type == TknWrite {
		current.childs = append(current.childs, sentWrite())
	} else if tokens[la].Type == TknLeftBr {
		current.childs = append(current.childs, bloque())
	} else {
		current.childs = append(current.childs, asignacion())
	}
	return current
}

func seleccion() *Node {
	var current *Node = new(Node)
	current.name = "seleccion"
	current.childs = append(current.childs, matchNode(TknIf))
	current.childs = append(current.childs, matchNode(TknLeftPar))
	current.childs = append(current.childs, bExpresion())
	current.childs = append(current.childs, matchNode(TknRightPar))
	current.childs = append(current.childs, matchNode(TknThen))
	current.childs = append(current.childs, bloque())
	if tokens[la].Type == TknElse {
		current.childs = append(current.childs, matchNode(TknElse))
		current.childs = append(current.childs, bloque())
	}
	current.childs = append(current.childs, matchNode(TknFi))
	return current
}

func iteracion() *Node {
	var current *Node = new(Node)
	current.name = "iteracion"
	current.childs = append(current.childs, matchNode(TknWhile))
	current.childs = append(current.childs, matchNode(TknLeftPar))
	current.childs = append(current.childs, bExpresion())
	current.childs = append(current.childs, matchNode(TknRightPar))
	current.childs = append(current.childs, bloque())
	return current
}

func repeticion() *Node {
	var current *Node = new(Node)
	current.name = "repeticion"
	current.childs = append(current.childs, matchNode(TknDo))
	current.childs = append(current.childs, bloque())
	current.childs = append(current.childs, matchNode(TknUntil))
	current.childs = append(current.childs, matchNode(TknLeftPar))
	current.childs = append(current.childs, bExpresion())
	current.childs = append(current.childs, matchNode(TknRightPar))
	current.childs = append(current.childs, matchNode(TknSemi))
	return current
}

func sentRead() *Node {
	var current *Node = new(Node)
	current.name = "sentenciaRead"
	current.childs = append(current.childs, matchNode(TknRead))
	current.childs = append(current.childs, matchNode(TknIdent))
	current.childs = append(current.childs, matchNode(TknSemi))
	return current
}

func sentWrite() *Node {
	var current *Node = new(Node)
	current.name = "sentenciaWrite"
	current.childs = append(current.childs, matchNode(TknWrite))
	current.childs = append(current.childs, bExpresion())
	current.childs = append(current.childs, matchNode(TknSemi))
	return current
}

func bloque() *Node {
	var current *Node = new(Node)
	current.name = "bloque"
	current.childs = append(current.childs, matchNode(TknLeftBr))
	current.childs = append(current.childs, listaSentencias())
	current.childs = append(current.childs, matchNode(TknRightBr))
	return current
}

func asignacion() *Node {
	var current *Node = new(Node)
	current.name = "asignacion"
	current.childs = append(current.childs, matchNode(TknIdent))
	current.childs = append(current.childs, matchNode(TknAssign))
	current.childs = append(current.childs, bExpresion())
	current.childs = append(current.childs, matchNode(TknSemi))
	return current
}

func bExpresion() *Node {
	var current *Node = new(Node)
	current.name = "bExpresion"
	current.childs = append(current.childs, bTerm())
	for tokens[la].Type == TknOr {
		current.childs = append(current.childs, matchNode(TknOr))
		current.childs = append(current.childs, bTerm())
	}
	return current
}

func bTerm() *Node {
	var current *Node = new(Node)
	current.name = "bTerm"
	current.childs = append(current.childs, notFactor())
	for tokens[la].Type == TknAnd {
		current.childs = append(current.childs, matchNode(TknAnd))
		current.childs = append(current.childs, notFactor())
	}
	return current
}

func notFactor() *Node {
	var current *Node = new(Node)
	current.name = "notFactor"
	if tokens[la].Type == TknNot {
		current.childs = append(current.childs, matchNode(TknNot))
	}
	current.childs = append(current.childs, bFactor())
	return current
}

func bFactor() *Node {
	var current *Node = new(Node)
	current.name = "bFactor"
	if tokens[la].Type == TknTrue || tokens[la].Type == TknFalse {
		current.childs = append(current.childs, matchNode(tokens[la].Type))
	} else {
		current.childs = append(current.childs, relacion())
	}
	return current
}

func relacion() *Node {
	var current *Node = new(Node)
	current.name = "relacion"
	current.childs = append(current.childs, expresion())
	if tokens[la].Type == TknLessEq || tokens[la].Type == TknLess || tokens[la].Type == TknGreat || tokens[la].Type == TknGreatEq || tokens[la].Type == TknEq || tokens[la].Type == TknNotEq {
		current.childs = append(current.childs, relOp())
		current.childs = append(current.childs, expresion())
	}
	return current
}

func relOp() *Node {
	var current *Node = new(Node)
	current.name = "relOp"
	if tokens[la].Type == TknLessEq {
		current.childs = append(current.childs, matchNode(TknLessEq))
	} else if tokens[la].Type == TknLess {
		current.childs = append(current.childs, matchNode(TknLess))
	} else if tokens[la].Type == TknGreat {
		current.childs = append(current.childs, matchNode(TknGreat))
	} else if tokens[la].Type == TknGreatEq {
		current.childs = append(current.childs, matchNode(TknGreatEq))
	} else if tokens[la].Type == TknEq {
		current.childs = append(current.childs, matchNode(TknEq))
	} else {
		current.childs = append(current.childs, matchNode(TknNotEq))
	}
	return current
}

func expresion() *Node {
	var current *Node = new(Node)
	current.name = "expresion"
	current.childs = append(current.childs, termino())
	if tokens[la].Type == TknSum || tokens[la].Type == TknSub {
		current.childs = append(current.childs, sumaOp())
		current.childs = append(current.childs, termino())
	}
	return current
}

func termino() *Node {
	var current *Node = new(Node)
	current.name = "termino"
	current.childs = append(current.childs, signoFactor())
	for tokens[la].Type == TknDiv || tokens[la].Type == TknMul {
		current.childs = append(current.childs, multOp())
		current.childs = append(current.childs, signoFactor())
	}
	return current
}

func multOp() *Node {
	var current *Node = new(Node)
	current.name = "multOp"
	if tokens[la].Type == TknMul {
		current.childs = append(current.childs, matchNode(TknMul))
	} else {
		current.childs = append(current.childs, matchNode(TknDiv))
	}
	return current
}

func sumaOp() *Node {
	var current *Node = new(Node)
	current.name = "sumaOP"
	if tokens[la].Type == TknSub {
		current.childs = append(current.childs, matchNode(TknSub))
	} else {
		current.childs = append(current.childs, matchNode(TknSum))
	}
	return current
}

func signoFactor() *Node {
	var current *Node = new(Node)
	current.name = "signoFactor"
	if tokens[la].Type == TknSum || tokens[la].Type == TknSub {
		current.childs = append(current.childs, sumaOp())
		current.childs = append(current.childs, factor())
	} else {
		current.childs = append(current.childs, factor())
	}
	return current
}

func factor() *Node {
	var current *Node = new(Node)
	current.name = "factor"
	if tokens[la].Type == TknLeftPar {
		current.childs = append(current.childs, matchNode(TknLeftPar))
		current.childs = append(current.childs, bExpresion())
		current.childs = append(current.childs, matchNode(TknRightPar))
	} else if tokens[la].Type == TknConst {
		current.childs = append(current.childs, matchNode(TknConst))
	} else {
		current.childs = append(current.childs, matchNode(TknIdent))
	}
	return current
}

func programa() *Node {
	var current = new(Node)

	current.name = "programa"
	current.childs = append(current.childs, matchNode(TknProgram))
	current.childs = append(current.childs, matchNode(TknLeftBr))
	current.childs = append(current.childs, listaDeclaraciones())
	current.childs = append(current.childs, listaSentencias())
	current.childs = append(current.childs, matchNode(TknRightBr))

	return current
}

func traverse(current *Node, tab int) {
	outputFile.WriteString("{\n")
	outputFile.WriteString(current.name + " , " + strconv.Itoa(current.Type) + "\n")

	for i := 0; i < len(current.childs); i++ {
		traverse(current.childs[i], tab+1)
	}
	outputFile.WriteString("}\n")

}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("SE DEBE PROPORCIONAR EXACTAMENTE UN ARCHIVO")
	} else {
		fileName = string(os.Args[1])

		sourceFile, err = os.Open(fileName)
		if err != nil {
			os.Exit(1)
			log.Fatal(err)
		}

		scanner = bufio.NewScanner(sourceFile)

		outputFile, err = os.Create("output/parseTree.txt")
		if err != nil {
			os.Exit(1)
			log.Fatal(err)
		}
		errorsFile, err = os.OpenFile("output/errors.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			os.Exit(1)
			log.Fatal(err)
		}
		defer func() {
			if err = sourceFile.Close(); err != nil {
				os.Exit(1)
				log.Fatal(err)
			}
			if err = outputFile.Close(); err != nil {
				os.Exit(1)
				log.Fatal(err)
			}
			if err = errorsFile.Close(); err != nil {
				os.Exit(1)
				log.Fatal(err)
			}
		}()

		readTokens()

		la = 0
		root = programa()

		traverse(root, 0)

		if tokens[la].Type == TknEOF {
			errorsFile.WriteString("Parseo Terminado")
		}

	}

	os.Exit(0)

}
