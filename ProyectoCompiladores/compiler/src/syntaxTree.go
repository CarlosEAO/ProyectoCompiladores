package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"io/ioutil"
	"encoding/json"
)

var fileName string
var outputFile *os.File
var errorsFile *os.File

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

//Token
type Token struct {
	Type      int `json:"Type"`
	Name      string `json:"Name"`
	Lexeme string `json:Lexeme`
	Row       int `json:"Row"`
	Column    int `json:"Column"`
}

//Node
type Node struct {
	ProductionName string
	Token Token
	Childs []*Node
}

var root *Node

func readTokens() {
	file, _ := ioutil.ReadFile(fileName)

	_ = json.Unmarshal([]byte(file), &tokens)

	for i:=0; i<len(tokens); i++{
		fmt.Println(string(tokens[i].Lexeme))
	}

}

func matchNode(Type int) *Node {
	var current *Node = new(Node)
	if tokens[la].Type == Type {
		current.Token = tokens[la]
		la++
	} else {
		//PANIC
		errorsFile.WriteString("ERROR en linea " + strconv.Itoa(tokens[la].Row) + ", columna " + strconv.Itoa(tokens[la].Column) + ". ")
		errorsFile.WriteString("Se esperaba " + tokenNames[Type] + ". Se encontro " + tokens[la].Lexeme + ".\n")
		for tokens[la].Type != Type && tokens[la].Type != TknEOF {
			la++
		}
		current.Token = tokens[la]

	}

	return current

}

func tipo() *Node {
	var current *Node = new(Node)
	current.ProductionName = "tipo"
	if tokens[la].Type == TknInt {
		current.Childs = append(current.Childs, matchNode(TknInt))
	} else if tokens[la].Type == TknBool {
		current.Childs = append(current.Childs, matchNode(TknBool))
	} else if tokens[la].Type == TknFloat {
		current.Childs = append(current.Childs, matchNode(TknFloat))
	}
	return current
}

func listaIdent() *Node {
	var current *Node = new(Node)
	current.ProductionName = "listaIdent"
	current.Childs = append(current.Childs, matchNode(TknIdent))
	for tokens[la].Type == TknComma {
		current.Childs = append(current.Childs, matchNode(TknComma))
		//current.Childs = append(current.Childs, listaIdent())
		current.Childs = append(current.Childs, matchNode(TknIdent))
	}
	return current
}

func declaracion() *Node {
	var current *Node = new(Node)
	current.ProductionName = "declaracion"
	current.Childs = append(current.Childs, tipo())
	current.Childs = append(current.Childs, listaIdent())
	current.Childs = append(current.Childs, matchNode(TknSemi))
	return current
}

func listaDeclaraciones() *Node {
	var current *Node = new(Node)
	current.ProductionName = "listaDeclaraciones"
	for tokens[la].Type == TknFloat || tokens[la].Type == TknInt || tokens[la].Type == TknBool {
		current.Childs = append(current.Childs, declaracion())
		//current.Childs = append(current.Childs, listaDeclaraciones())
	}
	return current
}

func listaSentencias() *Node {
	var current *Node = new(Node)
	current.ProductionName = "listaSentencias"
	for tokens[la].Type == TknIf || tokens[la].Type == TknWhile || tokens[la].Type == TknDo || tokens[la].Type == TknRead || tokens[la].Type == TknWrite || tokens[la].Type == TknLeftBr || tokens[la].Type == TknIdent {
		current.Childs = append(current.Childs, sentencia())
		//listaSentencias()
	}
	return current
}

func sentencia() *Node {
	var current *Node = new(Node)
	current.ProductionName = "sentencia"
	if tokens[la].Type == TknIf {
		current.Childs = append(current.Childs, seleccion())
	} else if tokens[la].Type == TknWhile {
		current.Childs = append(current.Childs, iteracion())
	} else if tokens[la].Type == TknDo {
		current.Childs = append(current.Childs, repeticion())
	} else if tokens[la].Type == TknRead {
		current.Childs = append(current.Childs, sentRead())
	} else if tokens[la].Type == TknWrite {
		current.Childs = append(current.Childs, sentWrite())
	} else if tokens[la].Type == TknLeftBr {
		current.Childs = append(current.Childs, bloque())
	} else {
		current.Childs = append(current.Childs, asignacion())
	}
	return current
}

func seleccion() *Node {
	var current *Node = new(Node)
	current.ProductionName = "seleccion"
	current.Childs = append(current.Childs, matchNode(TknIf))
	current.Childs = append(current.Childs, matchNode(TknLeftPar))
	current.Childs = append(current.Childs, bExpresion())
	current.Childs = append(current.Childs, matchNode(TknRightPar))
	current.Childs = append(current.Childs, matchNode(TknThen))
	current.Childs = append(current.Childs, bloque())
	if tokens[la].Type == TknElse {
		current.Childs = append(current.Childs, matchNode(TknElse))
		current.Childs = append(current.Childs, bloque())
	}
	current.Childs = append(current.Childs, matchNode(TknFi))
	return current
}

func iteracion() *Node {
	var current *Node = new(Node)
	current.ProductionName = "iteracion"
	current.Childs = append(current.Childs, matchNode(TknWhile))
	current.Childs = append(current.Childs, matchNode(TknLeftPar))
	current.Childs = append(current.Childs, bExpresion())
	current.Childs = append(current.Childs, matchNode(TknRightPar))
	current.Childs = append(current.Childs, bloque())
	return current
}

func repeticion() *Node {
	var current *Node = new(Node)
	current.ProductionName = "repeticion"
	current.Childs = append(current.Childs, matchNode(TknDo))
	current.Childs = append(current.Childs, bloque())
	current.Childs = append(current.Childs, matchNode(TknUntil))
	current.Childs = append(current.Childs, matchNode(TknLeftPar))
	current.Childs = append(current.Childs, bExpresion())
	current.Childs = append(current.Childs, matchNode(TknRightPar))
	current.Childs = append(current.Childs, matchNode(TknSemi))
	return current
}

func sentRead() *Node {
	var current *Node = new(Node)
	current.ProductionName = "sentenciaRead"
	current.Childs = append(current.Childs, matchNode(TknRead))
	current.Childs = append(current.Childs, matchNode(TknIdent))
	current.Childs = append(current.Childs, matchNode(TknSemi))
	return current
}

func sentWrite() *Node {
	var current *Node = new(Node)
	current.ProductionName = "sentenciaWrite"
	current.Childs = append(current.Childs, matchNode(TknWrite))
	current.Childs = append(current.Childs, bExpresion())
	current.Childs = append(current.Childs, matchNode(TknSemi))
	return current
}

func bloque() *Node {
	var current *Node = new(Node)
	current.ProductionName = "bloque"
	current.Childs = append(current.Childs, matchNode(TknLeftBr))
	current.Childs = append(current.Childs, listaSentencias())
	current.Childs = append(current.Childs, matchNode(TknRightBr))
	return current
}

func asignacion() *Node {
	var current *Node = new(Node)
	current.ProductionName = "asignacion"
	current.Childs = append(current.Childs, matchNode(TknIdent))
	current.Childs = append(current.Childs, matchNode(TknAssign))
	current.Childs = append(current.Childs, bExpresion())
	current.Childs = append(current.Childs, matchNode(TknSemi))
	return current
}

func bExpresion() *Node {
	var current *Node = new(Node)
	current.ProductionName = "bExpresion"
	current.Childs = append(current.Childs, bTerm())
	for tokens[la].Type == TknOr {
		current.Childs = append(current.Childs, matchNode(TknOr))
		current.Childs = append(current.Childs, bTerm())
	}
	return current
}

func bTerm() *Node {
	var current *Node = new(Node)
	current.ProductionName = "bTerm"
	current.Childs = append(current.Childs, notFactor())
	for tokens[la].Type == TknAnd {
		current.Childs = append(current.Childs, matchNode(TknAnd))
		current.Childs = append(current.Childs, notFactor())
	}
	return current
}

func notFactor() *Node {
	var current *Node = new(Node)
	current.ProductionName = "notFactor"
	if tokens[la].Type == TknNot {
		current.Childs = append(current.Childs, matchNode(TknNot))
	}
	current.Childs = append(current.Childs, bFactor())
	return current
}

func bFactor() *Node {
	var current *Node = new(Node)
	current.ProductionName = "bFactor"
	if tokens[la].Type == TknTrue || tokens[la].Type == TknFalse {
		current.Childs = append(current.Childs, matchNode(tokens[la].Type))
	} else {
		current.Childs = append(current.Childs, relacion())
	}
	return current
}

func relacion() *Node {
	var current *Node = new(Node)
	current.ProductionName = "relacion"
	current.Childs = append(current.Childs, expresion())
	if tokens[la].Type == TknLessEq || tokens[la].Type == TknLess || tokens[la].Type == TknGreat || tokens[la].Type == TknGreatEq || tokens[la].Type == TknEq || tokens[la].Type == TknNotEq {
		current.Childs = append(current.Childs, relOp())
		current.Childs = append(current.Childs, expresion())
	}
	return current
}

func relOp() *Node {
	var current *Node = new(Node)
	current.ProductionName = "relOp"
	if tokens[la].Type == TknLessEq {
		current.Childs = append(current.Childs, matchNode(TknLessEq))
	} else if tokens[la].Type == TknLess {
		current.Childs = append(current.Childs, matchNode(TknLess))
	} else if tokens[la].Type == TknGreat {
		current.Childs = append(current.Childs, matchNode(TknGreat))
	} else if tokens[la].Type == TknGreatEq {
		current.Childs = append(current.Childs, matchNode(TknGreatEq))
	} else if tokens[la].Type == TknEq {
		current.Childs = append(current.Childs, matchNode(TknEq))
	} else {
		current.Childs = append(current.Childs, matchNode(TknNotEq))
	}
	return current
}

func expresion() *Node {
	var current *Node = new(Node)
	current.ProductionName = "expresion"
	current.Childs = append(current.Childs, termino())
	if tokens[la].Type == TknSum || tokens[la].Type == TknSub {
		current.Childs = append(current.Childs, sumaOp())
		current.Childs = append(current.Childs, termino())
	}
	return current
}

func termino() *Node {
	var current *Node = new(Node)
	current.ProductionName = "termino"
	current.Childs = append(current.Childs, signoFactor())
	for tokens[la].Type == TknDiv || tokens[la].Type == TknMul {
		current.Childs = append(current.Childs, multOp())
		current.Childs = append(current.Childs, signoFactor())
	}
	return current
}

func multOp() *Node {
	var current *Node = new(Node)
	current.ProductionName = "multOp"
	if tokens[la].Type == TknMul {
		current.Childs = append(current.Childs, matchNode(TknMul))
	} else {
		current.Childs = append(current.Childs, matchNode(TknDiv))
	}
	return current
}

func sumaOp() *Node {
	var current *Node = new(Node)
	current.ProductionName = "sumaOP"
	if tokens[la].Type == TknSub {
		current.Childs = append(current.Childs, matchNode(TknSub))
	} else {
		current.Childs = append(current.Childs, matchNode(TknSum))
	}
	return current
}

func signoFactor() *Node {
	var current *Node = new(Node)
	current.ProductionName = "signoFactor"
	if tokens[la].Type == TknSum || tokens[la].Type == TknSub {
		current.Childs = append(current.Childs, sumaOp())
		current.Childs = append(current.Childs, factor())
	} else {
		current.Childs = append(current.Childs, factor())
	}
	return current
}

func factor() *Node {
	var current *Node = new(Node)
	current.ProductionName = "factor"
	if tokens[la].Type == TknLeftPar {
		current.Childs = append(current.Childs, matchNode(TknLeftPar))
		current.Childs = append(current.Childs, bExpresion())
		current.Childs = append(current.Childs, matchNode(TknRightPar))
	} else if tokens[la].Type == TknConst {
		current.Childs = append(current.Childs, matchNode(TknConst))
	} else {
		current.Childs = append(current.Childs, matchNode(TknIdent))
	}
	return current
}

func programa() *Node {
	var current = new(Node)

	current.ProductionName = "programa"
	current.Childs = append(current.Childs, matchNode(TknProgram))
	current.Childs = append(current.Childs, matchNode(TknLeftBr))
	current.Childs = append(current.Childs, listaDeclaraciones())
	current.Childs = append(current.Childs, listaSentencias())
	current.Childs = append(current.Childs, matchNode(TknRightBr))

	return current
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("SE DEBE PROPORCIONAR EXACTAMENTE UN ARCHIVO")
	} else {
		fileName = string(os.Args[1])

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

		fmt.Println(len(root.Childs))

		output, _ := json.Marshal(root)

		outputFile.WriteString(string(output))

		if tokens[la].Type == TknEOF {
			errorsFile.WriteString("Parseo Terminado")
		}
	}
	os.Exit(0)

}
