package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var fileName string
var buffer string
var sourceFile *os.File
var outputFile *os.File
var errorsFile *os.File

var reader *bufio.Reader

var currentColumn int = 0
var currentLine int = 0

var err error

var keywords = [...]string{"", "program", "if", "then", "else", "fi", "do", "until", "while", "read", "write", "float", "int", "bool", "true", "false", "not", "and", "or"}

//Los índices del arreglo tokenNames coinciden con las constantes numericas de la "enumeracion" token types. Quizas hubiera sido mejor idea utilizar un mapa en lugar de esta cosa bien cerda?

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

//STATES
const (
	StartState        = 0
	InWord            = 1
	InNumber          = 2
	InDecimal         = 3
	InOptionalDecimal = 4
	InSlash           = 5
	InLineComment     = 6
	InBlockComment    = 7
	InBlockCommentEnd = 8
	InLessThan        = 9
	InGreaterThan     = 10
	InEqual           = 11
	InNotEqual        = 12
	DoneState         = 30
)

//Token blaablabla
type Token struct {
	//Recordando que: type viene de tokn types, name viene de token names, lexeme es exactamente lo que está escrito en el código (para palabras reservadas, lexeme = name, o bueno algo así)
	Type   int
	Name   string
	Lexeme string
	Row    int
	Column int
}

//Quizas la estructura del token (los arreglos y la madre de arriba) deberían estar en un solo archivo?) Hubiera sido más conveniente y menos propenso a este cagadero
//Hacer ese cambio a estas alturas implica cambiar un buen (creo que un buen la neta no sé) de cosas abajo, entooonces pues alv


func getNextChar() rune {
	if currentColumn == 0 || currentColumn >= len(buffer) {
		currentColumn = 0

		buffer, err = reader.ReadString('\n')
		currentLine++
		if err != nil && len(buffer) == 0 {
			return 0
		}

	}
	//13 por el retorno de carro
	if buffer[currentColumn] == 13 {
		currentColumn++
	}
	var nextChar = buffer[currentColumn]
	currentColumn++
	return rune(nextChar)
}

func rollBackChar() {
	if currentColumn > 0 {
		currentColumn--
	}
}

func isDelim(c rune) bool {
	if c == ' ' || c == '\t' || c == '\n' {
		return true
	}
	return false

}

func isAlpha(c rune) bool {
	if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
		return true
	}
	return false
}

func isNumeric(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func isAlphaNumeric(c rune) bool {
	if isAlpha(c) || isNumeric(c) {
		return true
	}
	return false
}

func isKeyWord(word string) int {
	for i := 0; i < len(keywords); i++ {
		if word == keywords[i] {
			return i
		}
	}
	return TknIdent
}

//getToken dsff...
func getToken() Token {

	var token Token
	var nextChar rune
	var state = StartState

	for state != DoneState {
		switch state {
		case StartState:
			nextChar = getNextChar()
			for isDelim(nextChar) {
				nextChar = getNextChar()
			}
			token.Lexeme += string(nextChar)
			token.Row = currentLine
			token.Column = currentColumn
			if isAlpha(nextChar) { //START OF ANY WORD
				state = InWord
			} else if isNumeric(nextChar) { //START OF ANY NUMBER   
				state = InNumber
			} else if nextChar == '+' {
				token.Type = TknSum
				token.Name = "Suma"
				state = DoneState
			} else if nextChar == '-' {
				token.Type = TknSub
				token.Name = "Resta"
				state = DoneState
			} else if nextChar == '*' {
				token.Type = TknMul
				token.Name = "Multiplicacion"
				state = DoneState
			} else if nextChar == '/' { //START OF COMMENT OR DIVISION
				state = InSlash
			} else if nextChar == '^' {
				token.Type = TknExp
				token.Name = "Exponente"
				state = DoneState
			} else if nextChar == '<' { //LESS THAN OR LTE
				state = InLessThan
			} else if nextChar == '>' { //GREATER THAN OR GTE
				state = InGreaterThan
			} else if nextChar == '=' { //ASSIGNMENT OR EQUALITY
				state = InEqual
			} else if nextChar == '!' { //INEQUALITY
				state = InNotEqual
			} else if nextChar == ';' {
				token.Type = TknSemi
				token.Name = "PuntoYComa"
				state = DoneState
			} else if nextChar == ',' {
				token.Type = TknComma
				token.Name = "Coma"
				state = DoneState
			} else if nextChar == '(' {
				token.Type = TknLeftPar
				token.Name = "ParentesisIzq"
				state = DoneState
			} else if nextChar == ')' {
				token.Type = TknRightPar
				token.Name = "ParentesisDer"
				state = DoneState
			} else if nextChar == '{' {
				token.Type = TknLeftBr
				token.Name = "LlaveIzq"
				state = DoneState
			} else if nextChar == '}' {
				token.Type = TknRightBr
				token.Name = "LlaveDer"
				state = DoneState
			} else if nextChar == 0 {
				token.Type = TknEOF
				token.Name = "FindeArchivo"
				state = DoneState
			} else {
				token.Type = TknError
				token.Name = "Error"
				state = DoneState
			}
		case InWord:
			nextChar = getNextChar()
			if !isAlphaNumeric(nextChar) {
				token.Type = TknWord
				state = DoneState
				rollBackChar()
			} else {
				token.Lexeme += string(nextChar)
			}
		case InNumber:
			nextChar = getNextChar()
			if isNumeric(nextChar) {
				token.Lexeme += string(nextChar)
			} else if nextChar == '.' {
				token.Lexeme += string(nextChar)
				state = InDecimal
			} else {
				token.Type = TknConst
				token.Name = "ConstanteNumerica"
				state = DoneState
				rollBackChar()
			}
		case InDecimal:
			nextChar = getNextChar()
			if isNumeric(nextChar) {
				token.Lexeme += string(nextChar)
				state = InOptionalDecimal
			} else {
				token.Type = TknError
				token.Name = "Error"
				state = DoneState
				rollBackChar()
			}
		case InOptionalDecimal:
			nextChar = getNextChar()
			if isNumeric(nextChar) {
				token.Lexeme += string(nextChar)
			} else if nextChar == '.' {
				token.Type = TknError
				token.Name = "Error"
				state = DoneState
				rollBackChar()
			} else {
				token.Type = TknConst
				token.Name = "ConstanteNumerica"
				state = DoneState
				rollBackChar()
			}
		case InSlash:
			nextChar = getNextChar()
			if nextChar == '/' {
				token.Lexeme += string(nextChar)
				state = InLineComment
			} else if nextChar == '*' {
				token.Lexeme += string(nextChar)
				state = InBlockComment
			} else {
				token.Type = TknDiv
				token.Name = "Division"
				state = DoneState
				rollBackChar()
			}
		case InLineComment:
			nextChar = getNextChar()
			if nextChar == '\n' {
				token.Type = TknComment
				token.Name = "Comentario"
				state = DoneState
				rollBackChar()
			} else {
			}
		case InBlockComment:
			nextChar = getNextChar()
			if nextChar == '*' {
				state = InBlockCommentEnd
			} else {
			}
		case InBlockCommentEnd:
			nextChar = getNextChar()
			if nextChar == '/' {
				token.Type = TknComment
				token.Name = "Comentario"
				state = DoneState
			} else if nextChar != '*' {
				state = InBlockComment
			} else {
			}
		case InLessThan:
			nextChar = getNextChar()
			if nextChar == '=' {
				token.Lexeme += string(nextChar)
				token.Type = TknLessEq
				token.Name = "MenoroIgual"
				state = DoneState
			} else {
				token.Type = TknLess
				token.Name = "Menorestricto"
				state = DoneState
				rollBackChar()
			}
		case InGreaterThan:
			nextChar = getNextChar()
			if nextChar == '=' {
				token.Lexeme += string(nextChar)
				token.Type = TknGreatEq
				token.Name = "MayorOIgual"
				state = DoneState
			} else {
				token.Type = TknGreat
				token.Name = "MayorEstricto"
				state = DoneState
				rollBackChar()
			}
		case InEqual:
			nextChar = getNextChar()
			if nextChar == '=' {
				token.Lexeme += string(nextChar)
				token.Type = TknEq
				token.Name = "Igualdad"
				state = DoneState
			} else {
				token.Type = TknAssign
				token.Name = "Asignacion"
				state = DoneState
				rollBackChar()
			}
		case InNotEqual:
			nextChar = getNextChar()
			if nextChar == '=' {
				token.Lexeme += string(nextChar)
				token.Type = TknNotEq
				token.Name = "Diferente"
				state = DoneState
			} else {
				token.Type = TknError
				token.Name = "Error"
				state = DoneState
				rollBackChar()
			}
		default:
			{
				token.Type = TknError
				token.Name = "Error"
				state = DoneState
			}
		}

	}

	if token.Type == TknWord {

		token.Type = isKeyWord(token.Lexeme)
		if token.Type == TknIdent {
			token.Name = "Identificador"
		} else {
			token.Name = keywords[token.Type]
		}
	}

	return token
}

func main() {

	var currentToken Token
	var exitCode = 0

	if len(os.Args) != 2 {
		fmt.Println("SE DEBE PROPORCIONAR EXACTAMENTE UN ARCHIVO")
	} else {
		fileName = string(os.Args[1])

		sourceFile, err = os.Open(fileName)
		if err != nil {
			os.Exit(1)
			log.Fatal(err)
		}

		reader = bufio.NewReader(sourceFile)

		outputFile, err = os.Create("output/tokens.txt")
		if err != nil {
			os.Exit(1)
			log.Fatal(err)
		}
		errorsFile, err = os.Create("output/errors.txt")
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

		currentToken = getToken()

		for currentToken.Type != TknEOF {
			if currentToken.Type == TknError {
				//fmt.Println("Valio madre en la linea ", currentLine, ", '", currentToken.Lexeme, "' no corresponde a ningun token")
				exitCode = 2
				errorsFile.WriteString("Valio madre en la linea " + strconv.Itoa(currentLine) + ", columna " + strconv.Itoa(currentColumn) + ". '" + currentToken.Lexeme + "' no corresponde a ningun token.\n")
			} else {
				//fmt.Println(currentToken)
				if currentToken.Type != TknComment {
					outputFile.WriteString(strconv.Itoa(currentToken.Type) + " " + currentToken.Name + " " + string(currentToken.Lexeme) + " " + strconv.Itoa(currentToken.Row) + " " + strconv.Itoa(currentToken.Column) + "\n")
				}
			}
			currentToken = getToken()
		}
		outputFile.WriteString(strconv.Itoa(currentToken.Type) + " " + "FinDeArchivo" + " " + "EOF" + " " + strconv.Itoa(currentToken.Row) + " " + strconv.Itoa(currentToken.Column) + "\n")
	}

	os.Exit(exitCode)

}
