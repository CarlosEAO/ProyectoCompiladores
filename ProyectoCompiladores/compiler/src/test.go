package main

import (
	"bufio"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
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
	token Token
	childs []*Node
}

var data []Token

func main(){
	file, _ := ioutil.ReadFile(string(os.Args[1]))

	_ = json.Unmarshal([]byte(file), &data)

	fmt.Println(len(data))

for i:=0; i<len(data); i++{
	fmt.Println((data[i].Type))
}

}
