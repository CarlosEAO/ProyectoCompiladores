package main

import (
	"fmt"
	"log"
	"os"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

var fileName string
var outputFile *os.File
var errorsFile *os.File

var err error

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

var symbols map[string]map[string]string

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
	ProductionName string `json:"ProductionName"`
	Token Token `json:"Token"`
	Childs []*Node `json:"Childs"`
	Attributes map[string]string `json:"Attributes"`
}

var root *Node

var labelTotal int
var temporaryTotal int

func generateCode(current *Node){

	for i := 0; i < len(current.Childs); i++ {
		generateCode(current.Childs[i])
	}

	if current.ProductionName == "factor"{
		temporaryTotal = temporaryTotal + 1;
		current.Attributes["idTemp"] = "temp"+strconv.Itoa(temporaryTotal);
		if(len(current.Childs) == 1){
			outputFile.WriteString("( "+current.Attributes["idTemp"] +" , "+ current.Childs[0].Token.Lexeme +" )\n");
		}else{
			outputFile.WriteString("( "+current.Attributes["idTemp"] +" , "+ current.Childs[1].Attributes["idTemp"] +" )\n");
		}
	}

	if current.ProductionName == "signoFactor" {
		if(len(current.Childs) == 1){
			current.Attributes["idTemp"] = current.Childs[0].Attributes["idTemp"];
		}else{
			temporaryTotal = temporaryTotal + 1;
			current.Attributes["idTemp"] = "temp"+strconv.Itoa(temporaryTotal);
			outputFile.WriteString("( "+current.Attributes["idTemp"] +" , * , " + current.Childs[1].Attributes["idTemp"] +" -1)\n");
		}
	}


	if(current.ProductionName == "termino" || current.ProductionName == "expresion" || current.ProductionName == "relacion"){
		if(len(current.Childs) == 1){
			current.Attributes["idTemp"] = current.Childs[0].Attributes["idTemp"];
		}else{
			temporaryTotal = temporaryTotal + 1;
			current.Attributes["idTemp"] = "temp"+strconv.Itoa(temporaryTotal);
			outputFile.WriteString("( "+current.Attributes["idTemp"] +" , " + current.Childs[0].Attributes["idTemp"] +" , " +current.Childs[1].Attributes["op"] + " , "+ current.Childs[2].Attributes["idTemp"] + " )\n");
		}
	}

	if(current.ProductionName == "bFactor"){
		if(current.Childs[0].Token.Lexeme == "true" || current.Childs[0].Token.Lexeme == "false" ){
			temporaryTotal = temporaryTotal + 1;
			current.Attributes["idTemp"] = "temp"+strconv.Itoa(temporaryTotal);
			if(current.Childs[0].Token.Lexeme == "true"){
				outputFile.WriteString("( "+current.Attributes["idTemp"] +" , 1 )\n");
			}else{
				outputFile.WriteString("( "+current.Attributes["idTemp"] +" , 0 )\n");
			}
		}else{
			current.Attributes["idTemp"] = current.Childs[0].Attributes["idTemp"];
		}
	}

	if current.ProductionName == "notFactor" {
		if(len(current.Childs) == 1){
			current.Attributes["idTemp"] = current.Childs[0].Attributes["idTemp"];
		}else{
			temporaryTotal = temporaryTotal + 1;
			current.Attributes["idTemp"] = "temp"+strconv.Itoa(temporaryTotal);
			outputFile.WriteString("( "+current.Attributes["idTemp"] +" , not , " + current.Childs[1].Attributes["idTemp"] +" )\n");
		}
	}

	if(current.ProductionName == "bTerm" || current.ProductionName == "bExpresion"){
		if(len(current.Childs) == 1){
			current.Attributes["idTemp"] = current.Childs[0].Attributes["idTemp"];
		}else{
			temporaryTotal = temporaryTotal + 1;
			current.Attributes["idTemp"] = "temp"+strconv.Itoa(temporaryTotal);
			outputFile.WriteString("( "+current.Attributes["idTemp"] +" , " + current.Childs[0].Attributes["idTemp"] +" , " +current.Childs[1].Token.Lexeme + " , "+ current.Childs[2].Attributes["idTemp"] + " )\n");
		}
	}

	if(current.ProductionName == "asignacion"){
		outputFile.WriteString("( "+current.Childs[0].Token.Lexeme +" , " + current.Childs[2].Attributes["idTemp"]+ " )\n");
	}


}


func main() {

	if len(os.Args) != 2 {
		fmt.Println("SE DEBE PROPORCIONAR EXACTAMENTE UN ARCHIVO")
	} else {
		fileName = string(os.Args[1])
		file, _ := ioutil.ReadFile(fileName)
		root = new(Node)
		_ = json.Unmarshal([]byte(file), root)


		outputFile, err = os.Create("output/intermediateCode.txt")
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

		generateCode(root);

		output, _ := json.Marshal(root)
		fmt.Println(string(output))


	}
	os.Exit(0)

}
