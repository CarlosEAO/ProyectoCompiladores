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

//ejemplo de cagadero: En esta estructura una cosa se llama attribute y en la que está declarada en el archivo de lexico se llama lexema  
type Token struct {
	Type      int
	Name      string
	Attribute string
	Row       int
	Column    int
}

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

func match(Type int) bool {
	if tokens[la].Type == Type {
		la++
		return true
	}
	fmt.Println("ERROR en linea " + strconv.Itoa(tokens[la].Row))
	return false
}

func tipo() {

	if tokens[la].Type == TknInt {
		match(TknInt)
	} else if tokens[la].Type == TknBool {
		match(TknBool)
	} else if tokens[la].Type == TknFloat {
		match(TknFloat)
	}
}

func listaIdent() {
	match(TknIdent)
	if tokens[la].Type == TknComma {
		match(TknComma)
		listaIdent()
	} else {

	}
}

func declaracion() {
	tipo()
	listaIdent()
	match(TknSemi)
}

func listaDeclaraciones() {
	if tokens[la].Type == TknFloat || tokens[la].Type == TknInt || tokens[la].Type == TknBool {
		declaracion()
		listaDeclaraciones()
	}
}

func listaSentencias() {
	if tokens[la].Type == TknIf || tokens[la].Type == TknWhile || tokens[la].Type == TknDo || tokens[la].Type == TknRead || tokens[la].Type == TknWrite || tokens[la].Type == TknLeftBr || tokens[la].Type == TknIdent {
		sentencia()
		listaSentencias()
	}
}

func sentencia() {
	if tokens[la].Type == TknIf {
		seleccion()
	} else if tokens[la].Type == TknWhile {
		iteracion()
	} else if tokens[la].Type == TknDo {
		repeticion()
	} else if tokens[la].Type == TknRead {
		sentRead()
	} else if tokens[la].Type == TknWrite {
		sentWrite()
	} else if tokens[la].Type == TknLeftBr {
		bloque()
	} else {
		asignacion()
	}
}

func seleccion() {
	match(TknIf)
	match(TknLeftPar)
	bExpresion()
	match(TknRightPar)
	match(TknThen)
	bloque()
	if tokens[la].Type == TknElse {
		match(TknElse)
		bloque()
	}
	match(TknFi)
}

func iteracion() {
	match(TknWhile)
	match(TknLeftPar)
	bExpresion()
	match(TknRightPar)
	bloque()
}

func repeticion() {
	match(TknDo)
	bloque()
	match(TknUntil)
	match(TknLeftPar)
	bExpresion()
	match(TknRightPar)
	match(TknSemi)
}

func sentRead() {
	match(TknRead)
	match(TknIdent)
	match(TknSemi)
}

func sentWrite() {
	match(TknWrite)
	bExpresion()
	match(TknSemi)
}

func bloque() {
	match(TknLeftBr)
	listaSentencias()
	match(TknRightBr)
}

func asignacion() {
	match(TknIdent)
	match(TknAssign)
	bExpresion()
	match(TknSemi)
}

func bExpresion() {
	bTerm()
	for tokens[la].Type == TknOr {
		match(TknOr)
		bTerm()
	}
}

func bTerm() {
	notFactor()
	for tokens[la].Type == TknAnd {
		match(TknAnd)
		notFactor()
	}
}

func notFactor() {
	if tokens[la].Type == TknNot {
		match(TknNot)
	}
	bFactor()
}

func bFactor() {
	if tokens[la].Type == TknTrue || tokens[la].Type == TknFalse {
		match(tokens[la].Type)
	} else {
		relacion()
	}
}

func relacion() {
	expresion()
	if tokens[la].Type == TknLessEq || tokens[la].Type == TknLess || tokens[la].Type == TknGreat || tokens[la].Type == TknGreatEq || tokens[la].Type == TknEq || tokens[la].Type == TknNotEq {
		relOp()
		expresion()
	}
}

func relOp() {
	if tokens[la].Type == TknLessEq {
		match(TknLessEq)
	} else if tokens[la].Type == TknLess {
		match(TknLess)
	} else if tokens[la].Type == TknGreat {
		match(TknGreat)
	} else if tokens[la].Type == TknGreatEq {
		match(TknGreatEq)
	} else if tokens[la].Type == TknEq {
		match(TknEq)
	} else {
		match(TknNotEq)
	}
}

func expresion() {
	termino()
	if tokens[la].Type == TknSum || tokens[la].Type == TknSub {
		sumaOp()
		termino()
	} else {

	}
}

func termino() {
	signoFactor()
	for tokens[la].Type == TknDiv || tokens[la].Type == TknMul {
		multOp()
		signoFactor()
	}

}

func multOp() {
	if tokens[la].Type == TknMul {
		match(TknMul)
	} else {
		match(TknDiv)
	}
}

func sumaOp() {
	if tokens[la].Type == TknSub {
		match(TknSub)
	} else {
		match(TknSum)
	}
}

func signoFactor() {
	if tokens[la].Type == TknSum || tokens[la].Type == TknSub {
		sumaOp()
		factor()
	} else {
		factor()
	}
}

func factor() {
	if tokens[la].Type == TknLeftPar {
		match(TknLeftPar)
		bExpresion()
		match(TknRightPar)
	} else if tokens[la].Type == TknConst {
		match(TknConst)
	} else {
		match(TknIdent)
	}
}

func programa() {
	match(TknProgram)
	match(TknLeftBr)
	listaDeclaraciones()
	listaSentencias()
	match(TknRightBr)
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

		outputFile, err = os.Create("output.txt")
		if err != nil {
			os.Exit(1)
			log.Fatal(err)
		}
		errorsFile, err = os.Create("errors.txt")
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
		programa()

		if tokens[la].Type == TknEOF {
			fmt.Println("Parseo Terminado")
		}

	}

	os.Exit(0)

}
