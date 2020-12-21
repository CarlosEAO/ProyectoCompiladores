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
var noErrors bool

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
	Attributes map[string]string
}

var root *Node

func initializeAttributesMaps(current *Node,){
	current.Attributes = make(map[string]string)
	for i := 0; i < len(current.Childs); i++ {
		initializeAttributesMaps(current.Childs[i])
	}
}

func computeAttributes(current *Node, parent *Node){

	for i := 0; i < len(current.Childs); i++ {
		computeAttributes(current.Childs[i], current)
	}

	if current.ProductionName == "tipo"{
		current.Attributes["dtype"] = current.Childs[0].Token.Lexeme
	}
	if current.ProductionName == "declaracion"{
		current.Childs[1].Attributes["dtype"] = current.Childs[0].Attributes["dtype"]
	}
	if current.ProductionName == "listaIdent"{
		for i:=0; i<len(current.Childs); i++{
			if current.Childs[i].Token.Type == TknIdent{

				if symbols[current.Childs[i].Token.Lexeme]!=nil{
					errorsFile.WriteString("VALIO madres, redeclarando " + current.Childs[i].Token.Lexeme+"\n")
					noErrors = false
				}else{
					current.Childs[i].Attributes["dtype"] = parent.Childs[0].Attributes["dtype"];
					current.Childs[i].Attributes["value"] = "0";
					symbols[current.Childs[i].Token.Lexeme] = make(map[string]string)
					symbols[current.Childs[i].Token.Lexeme]["dtype"] = current.Childs[i].Attributes["dtype"]
					symbols[current.Childs[i].Token.Lexeme]["value"] = current.Childs[i].Attributes["value"]; 
				}


			}
		}
	}

	if(current.ProductionName  == "factor" ){
		if(len(current.Childs) == 1){
			if(current.Childs[0].Token.Type == TknConst){
				current.Attributes["value"] = current.Childs[0].Token.Lexeme;
			}else{
				if(symbols[current.Childs[0].Token.Lexeme]!=nil){
					current.Attributes["value"] = symbols[current.Childs[0].Token.Lexeme]["value"];
				}else{
					current.Attributes["value"] = "0";
					errorsFile.WriteString(current.Childs[0].Token.Lexeme + " no está declarado\n");
					noErrors = false
				}
			}
		}else{
			current.Attributes["value"] = current.Childs[1].Attributes["value"];
		}
	}

	if(current.ProductionName == "sumaOp" || current.ProductionName == "multOp" || current.ProductionName == "relOp"){
		current.Attributes["op"] = current.Childs[0].Token.Lexeme;
	}

	if(current.ProductionName == "signoFactor"){
		if(len(current.Childs) == 1){
			current.Attributes["value"] = current.Childs[0].Attributes["value"];
		}else{
			current.Attributes["value"] = current.Childs[0].Attributes["op"] + current.Childs[1].Attributes["value"];
		}
	}

	if(current.ProductionName == "termino"){
		if(len(current.Childs) == 1){
			current.Attributes["value"] = current.Childs[0].Attributes["value"];
		}else{
			term1,err := strconv.ParseFloat(current.Childs[0].Attributes["value"],64);
			if(err!=nil){
				os.Exit(1);
			}
			term2,err := strconv.ParseFloat(current.Childs[2].Attributes["value"],64);
			if(err!=nil){
				os.Exit(1);
			}
			ans:=0.0;
			if(current.Childs[1].Attributes["op"] == "*"){
				ans = term1*term2;
			}else{
				if(term2 == 0){
					errorsFile.WriteString("VALIO MADRE, ANDAS DIVIDIENDO ENTRE 0");
					noErrors = false
				}else{
					ans = term1/term2;
				}
			}
			current.Attributes["value"] = strconv.FormatFloat(ans, 'f', -1, 64);
		}
	}

	if(current.ProductionName == "expresion"){
		if(len(current.Childs) == 1){
			current.Attributes["value"] = current.Childs[0].Attributes["value"];
		}else{
			term1,err := strconv.ParseFloat(current.Childs[0].Attributes["value"],64);
			if(err!=nil){
				os.Exit(1);
			}
			term2,err := strconv.ParseFloat(current.Childs[2].Attributes["value"],64);
			if(err!=nil){
				os.Exit(1);
			}
			ans:=0.0;
			if(current.Childs[1].Attributes["op"] == "+"){
				ans = term1+term2;
			}else{
				ans = term1-term2;
				
			}
			current.Attributes["value"] = strconv.FormatFloat(ans, 'f', -1, 64);
		}
	}

	if(current.ProductionName == "relacion"){
		if(len(current.Childs) == 1){
			current.Attributes["value"] = current.Childs[0].Attributes["value"];
		}else{
			exp1,err := strconv.ParseFloat(current.Childs[0].Attributes["value"],64);
			if(err!=nil){
				os.Exit(1);
			}
			exp2,err := strconv.ParseFloat(current.Childs[2].Attributes["value"],64);
			if(err!=nil){
				os.Exit(1);
			}
			ans:=0;
			switch current.Childs[1].Attributes["op"]{
			case  "<=":
				if (exp1 <= exp2){
					ans = 1;
				}
			case "<":
				if(exp1<exp2){
					ans = 1;
				}
			case ">":
				if(exp1>exp2){
					ans = 1;
				}
			case ">=":
				if(exp1>=exp2){
					ans = 1;
				}
			case "==":
				if(exp1==exp2){
					ans = 1;
				}
			case "!=":
				if(exp1!=exp2){
					ans = 1;
				}
			default:
				ans = 1;
			}
			current.Attributes["value"] = strconv.Itoa(ans);
		}
	}

	if(current.ProductionName == "bFactor"){
		if(current.Childs[0].Token.Lexeme == "false"){
			current.Attributes["value"] = "0";
		}else if(current.Childs[0].Token.Lexeme == "true"){
			current.Attributes["value"] = "1";
		}else{
			current.Attributes["value"] = current.Childs[0].Attributes["value"];
		}
	}

	if(current.ProductionName == "notFactor"){
		if(len(current.Childs) == 1){
			current.Attributes["value"] = current.Childs[0].Attributes["value"];
		}else{
			if(current.Childs[1].Attributes["value"] == "0"){
				current.Attributes["value"]  = "1";
			}else{
				current.Attributes["value"] = "0";
			}
		}
	}

	if(current.ProductionName == "bTerm"){
		if(len(current.Childs) == 1){
			current.Attributes["value"] = current.Childs[0].Attributes["value"];
		}else{
			factor1:=true;
			factor2:=true;
			if(current.Childs[0].Attributes["value"] == "0"){
				factor1 = false;
			}
			if(current.Childs[2].Attributes["value"] == "0"){
				factor2 = false;
			}

			ans:= factor1 && factor2;
			if(ans == true){
				current.Attributes["value"] = "1";
			}else{
				current.Attributes["value"] = "0";
			}

		}
	}

	if(current.ProductionName == "bExpresion"){
		if(len(current.Childs) == 1){
			current.Attributes["value"] = current.Childs[0].Attributes["value"];
		}else{
			factor1:=true;
			factor2:=true;
			if(current.Childs[0].Attributes["value"] == "0"){
				factor1 = false;
			}
			if(current.Childs[2].Attributes["value"] == "0"){
				factor2 = false;
			}

			ans:= factor1 || factor2;
			if(ans == true){
				current.Attributes["value"] = "1";
			}else{
				current.Attributes["value"] = "0";
			}
		}
	}

	if(current.ProductionName == "asignacion"){

		if(symbols[current.Childs[0].Token.Lexeme]!=nil){
			current.Childs[0].Attributes["value"] = current.Childs[2].Attributes["value"];
			symbols[current.Childs[0].Token.Lexeme]["value"] = current.Childs[0].Attributes["value"];
		}else{
			errorsFile.WriteString(current.Childs[0].Token.Lexeme + " no está declarado\n");
			noErrors = false
		}
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


		outputFile, err = os.Create("output/attributedST.txt")
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

		initializeAttributesMaps(root)
		symbols = make(map[string]map[string]string)
		noErrors = true

		computeAttributes(root, nil)

		output, _ := json.Marshal(root)

		outputFile.WriteString(string(output))
		if noErrors == true{
			errorsFile.WriteString("No hubo errores semánticos\n")
		}


	}
	os.Exit(0)

}
