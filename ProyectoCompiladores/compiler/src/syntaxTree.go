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
	productionName string
	token Token
	childs []*Node
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
		current.token = tokens[la]
		la++
	} else {
		//PANIC
		errorsFile.WriteString("ERROR en linea " + strconv.Itoa(tokens[la].Row) + ", columna " + strconv.Itoa(tokens[la].Column) + ". ")
		errorsFile.WriteString("Se esperaba " + tokenNames[Type] + ". Se encontro " + tokens[la].Lexeme + ".\n")
		for tokens[la].Type != Type && tokens[la].Type != TknEOF {
			la++
		}
		current.token = tokens[la]

	}

	return current

}

func tipo() *Node {
	var current *Node = new(Node)
	current.productionName = "tipo"
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
	current.productionName = "listaIdent"
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
	current.productionName = "declaracion"
	current.childs = append(current.childs, tipo())
	current.childs = append(current.childs, listaIdent())
	current.childs = append(current.childs, matchNode(TknSemi))
	return current
}

func listaDeclaraciones() *Node {
	var current *Node = new(Node)
	current.productionName = "listaDeclaraciones"
	for tokens[la].Type == TknFloat || tokens[la].Type == TknInt || tokens[la].Type == TknBool {
		current.childs = append(current.childs, declaracion())
		//current.childs = append(current.childs, listaDeclaraciones())
	}
	return current
}

func listaSentencias() *Node {
	var current *Node = new(Node)
	current.productionName = "listaSentencias"
	for tokens[la].Type == TknIf || tokens[la].Type == TknWhile || tokens[la].Type == TknDo || tokens[la].Type == TknRead || tokens[la].Type == TknWrite || tokens[la].Type == TknLeftBr || tokens[la].Type == TknIdent {
		current.childs = append(current.childs, sentencia())
		//listaSentencias()
	}
	return current
}

func sentencia() *Node {
	var current *Node = new(Node)
	current.productionName = "sentencia"
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
	current.productionName = "seleccion"
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
	current.productionName = "iteracion"
	current.childs = append(current.childs, matchNode(TknWhile))
	current.childs = append(current.childs, matchNode(TknLeftPar))
	current.childs = append(current.childs, bExpresion())
	current.childs = append(current.childs, matchNode(TknRightPar))
	current.childs = append(current.childs, bloque())
	return current
}

func repeticion() *Node {
	var current *Node = new(Node)
	current.productionName = "repeticion"
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
	current.productionName = "sentenciaRead"
	current.childs = append(current.childs, matchNode(TknRead))
	current.childs = append(current.childs, matchNode(TknIdent))
	current.childs = append(current.childs, matchNode(TknSemi))
	return current
}

func sentWrite() *Node {
	var current *Node = new(Node)
	current.productionName = "sentenciaWrite"
	current.childs = append(current.childs, matchNode(TknWrite))
	current.childs = append(current.childs, bExpresion())
	current.childs = append(current.childs, matchNode(TknSemi))
	return current
}

func bloque() *Node {
	var current *Node = new(Node)
	current.productionName = "bloque"
	current.childs = append(current.childs, matchNode(TknLeftBr))
	current.childs = append(current.childs, listaSentencias())
	current.childs = append(current.childs, matchNode(TknRightBr))
	return current
}

func asignacion() *Node {
	var current *Node = new(Node)
	current.productionName = "asignacion"
	current.childs = append(current.childs, matchNode(TknIdent))
	current.childs = append(current.childs, matchNode(TknAssign))
	current.childs = append(current.childs, bExpresion())
	current.childs = append(current.childs, matchNode(TknSemi))
	return current
}

func bExpresion() *Node {
	var current *Node = new(Node)
	current.productionName = "bExpresion"
	current.childs = append(current.childs, bTerm())
	for tokens[la].Type == TknOr {
		current.childs = append(current.childs, matchNode(TknOr))
		current.childs = append(current.childs, bTerm())
	}
	return current
}

func bTerm() *Node {
	var current *Node = new(Node)
	current.productionName = "bTerm"
	current.childs = append(current.childs, notFactor())
	for tokens[la].Type == TknAnd {
		current.childs = append(current.childs, matchNode(TknAnd))
		current.childs = append(current.childs, notFactor())
	}
	return current
}

func notFactor() *Node {
	var current *Node = new(Node)
	current.productionName = "notFactor"
	if tokens[la].Type == TknNot {
		current.childs = append(current.childs, matchNode(TknNot))
	}
	current.childs = append(current.childs, bFactor())
	return current
}

func bFactor() *Node {
	var current *Node = new(Node)
	current.productionName = "bFactor"
	if tokens[la].Type == TknTrue || tokens[la].Type == TknFalse {
		current.childs = append(current.childs, matchNode(tokens[la].Type))
	} else {
		current.childs = append(current.childs, relacion())
	}
	return current
}

func relacion() *Node {
	var current *Node = new(Node)
	current.productionName = "relacion"
	current.childs = append(current.childs, expresion())
	if tokens[la].Type == TknLessEq || tokens[la].Type == TknLess || tokens[la].Type == TknGreat || tokens[la].Type == TknGreatEq || tokens[la].Type == TknEq || tokens[la].Type == TknNotEq {
		current.childs = append(current.childs, relOp())
		current.childs = append(current.childs, expresion())
	}
	return current
}

func relOp() *Node {
	var current *Node = new(Node)
	current.productionName = "relOp"
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
	current.productionName = "expresion"
	current.childs = append(current.childs, termino())
	if tokens[la].Type == TknSum || tokens[la].Type == TknSub {
		current.childs = append(current.childs, sumaOp())
		current.childs = append(current.childs, termino())
	}
	return current
}

func termino() *Node {
	var current *Node = new(Node)
	current.productionName = "termino"
	current.childs = append(current.childs, signoFactor())
	for tokens[la].Type == TknDiv || tokens[la].Type == TknMul {
		current.childs = append(current.childs, multOp())
		current.childs = append(current.childs, signoFactor())
	}
	return current
}

func multOp() *Node {
	var current *Node = new(Node)
	current.productionName = "multOp"
	if tokens[la].Type == TknMul {
		current.childs = append(current.childs, matchNode(TknMul))
	} else {
		current.childs = append(current.childs, matchNode(TknDiv))
	}
	return current
}

func sumaOp() *Node {
	var current *Node = new(Node)
	current.productionName = "sumaOP"
	if tokens[la].Type == TknSub {
		current.childs = append(current.childs, matchNode(TknSub))
	} else {
		current.childs = append(current.childs, matchNode(TknSum))
	}
	return current
}

func signoFactor() *Node {
	var current *Node = new(Node)
	current.productionName = "signoFactor"
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
	current.productionName = "factor"
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

	current.productionName = "programa"
	current.childs = append(current.childs, matchNode(TknProgram))
	current.childs = append(current.childs, matchNode(TknLeftBr))
	current.childs = append(current.childs, listaDeclaraciones())
	current.childs = append(current.childs, listaSentencias())
	current.childs = append(current.childs, matchNode(TknRightBr))

	return current
}

func traverse(current *Node, tab int) {


	/*outputFile.WriteString("{\n")
	outputFile.WriteString(current.productionName + " , " + strconv.Itoa(current.Type) + "\n")

	for i := 0; i < len(current.childs); i++ {
		traverse(current.childs[i], tab+1)
	}
	outputFile.WriteString("}\n")*/

}

//Attribute ST
//Dont forget to initializee Attributes after calling the new operator
type NodeWAttributes struct {
	productionName string
	token Token
	Attributes map[string]string
	childs []*NodeWAttributes
}

var rootAAST *NodeWAttributes 

//COPY ATTRIBUTE ST from ST
/*
func copyAttribute(currentST *Node, currentAAST *NodeWAttributes) {
	currentAAST.Attributes = make(map[string]string)
	currentAAST.name = currentST.name
	currentAAST.lexeme = currentST.lexeme
	currentAAST.lexeme = currentST.lexeme

	for i := 0; i < len(currentST.childs); i++ {
		var newNodeWAttributes * NodeWAttributes= new(NodeWAttributes) 
		currentAAST.childs = append(currentAAST.childs, newNodeWAttributes)
		copyAttribute(currentST.childs[i],currentAAST.childs[i] )
	}
}*/


//compute attributes
func computeAttributes(currentAAST *NodeWAttributes){
	
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

		fmt.Println(len(root.childs))

		output, _ := json.Marshal(root)

		outputFile.WriteString(string(output))

		if tokens[la].Type == TknEOF {
			errorsFile.WriteString("Parseo Terminado")
		}

		rootAAST = new(NodeWAttributes)

		//copyAttribute(root, rootAAST)



	}

	os.Exit(0)

}
