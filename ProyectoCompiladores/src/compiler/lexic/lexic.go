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

var keywords = [...]string{"program", "if", "then", "else", "fi", "do", "until", "while", "read", "write", "float", "int", "bool", "true", "false", "not", "and", "or"}

//TOKEN TYPES
const (
	TknProgram  = 0
	TknIf       = 1
	TknThen     = 2
	TknElse     = 3
	TknFi       = 4
	TknDo       = 5
	TknUntil    = 6
	TknWhile    = 7
	TknRead     = 8
	TknWrite    = 9
	TknFloat    = 10
	TknInt      = 11
	TknBool     = 12
	TknTrue     = 13
	TknFalse    = 14
	TknNot      = 15
	TknAnd      = 16
	TknOr       = 17
	TknSum      = 18
	TknSub      = 19
	TknMul      = 20
	TknDiv      = 21
	TknExp      = 22
	TknLess     = 23
	TknLessEq   = 24
	TknGreat    = 25
	TknGreatEq  = 26
	TknEq       = 27
	TknNotEq    = 28
	TknAssign   = 29
	TknSemi     = 30
	TknComma    = 31
	TknLeftPar  = 32
	TknRightPar = 33
	TknLeftBr   = 34
	TknRightBr  = 35
	TknError    = 36
	TknIdent    = 37
	TknEOF      = 38
	TknComment  = 39
	TknConst    = 40
	TknWord     = 41
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
	Type      int
	Attribute string
	Row       int
	Column    int
}

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
			token.Attribute += string(nextChar)
			token.Row = currentLine
			token.Column = currentColumn
			if isAlpha(nextChar) { //START OF ANY WORD
				state = InWord
			} else if isNumeric(nextChar) { //START OF ANY NUMBER
				state = InNumber
			} else if nextChar == '+' {
				token.Type = TknSum
				state = DoneState
			} else if nextChar == '-' {
				token.Type = TknSub
				state = DoneState
			} else if nextChar == '*' {
				token.Type = TknMul
				state = DoneState
			} else if nextChar == '/' { //START OF COMMENT OR DIVISION
				state = InSlash
			} else if nextChar == '^' {
				token.Type = TknExp
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
				state = DoneState
			} else if nextChar == ',' {
				token.Type = TknComma
				state = DoneState
			} else if nextChar == '(' {
				token.Type = TknLeftPar
				state = DoneState
			} else if nextChar == ')' {
				token.Type = TknRightPar
				state = DoneState
			} else if nextChar == '{' {
				token.Type = TknLeftBr
				state = DoneState
			} else if nextChar == '}' {
				token.Type = TknRightBr
				state = DoneState
			} else if nextChar == 0 {
				token.Type = TknEOF
				state = DoneState
			} else {
				token.Type = TknError
				state = DoneState
			}
		case InWord:
			nextChar = getNextChar()
			if !isAlphaNumeric(nextChar) {
				token.Type = TknWord
				state = DoneState
				rollBackChar()
			} else {
				token.Attribute += string(nextChar)
			}
		case InNumber:
			nextChar = getNextChar()
			if isNumeric(nextChar) {
				token.Attribute += string(nextChar)
			} else if nextChar == '.' {
				token.Attribute += string(nextChar)
				state = InDecimal
			} else {
				token.Type = TknConst
				state = DoneState
				rollBackChar()
			}
		case InDecimal:
			nextChar = getNextChar()
			if isNumeric(nextChar) {
				token.Attribute += string(nextChar)
				state = InOptionalDecimal
			} else {
				token.Type = TknError
				state = DoneState
				rollBackChar()
			}
		case InOptionalDecimal:
			nextChar = getNextChar()
			if isNumeric(nextChar) {
				token.Attribute += string(nextChar)
			} else if nextChar == '.' {
				token.Type = TknError
				state = DoneState
				rollBackChar()
			} else {
				token.Type = TknConst
				state = DoneState
				rollBackChar()
			}
		case InSlash:
			nextChar = getNextChar()
			if nextChar == '/' {
				token.Attribute += string(nextChar)
				state = InLineComment
			} else if nextChar == '*' {
				token.Attribute += string(nextChar)
				state = InBlockComment
			} else {
				token.Type = TknDiv
				state = DoneState
				rollBackChar()
			}
		case InLineComment:
			nextChar = getNextChar()
			if nextChar == '\n' {
				token.Type = TknComment
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
				state = DoneState
			} else if nextChar != '*' {
				state = InBlockComment
			} else {
			}
		case InLessThan:
			nextChar = getNextChar()
			if nextChar == '=' {
				token.Attribute += string(nextChar)
				token.Type = TknLessEq
				state = DoneState
			} else {
				token.Type = TknLess
				state = DoneState
				rollBackChar()
			}
		case InGreaterThan:
			nextChar = getNextChar()
			if nextChar == '=' {
				token.Attribute += string(nextChar)
				token.Type = TknGreatEq
				state = DoneState
			} else {
				token.Type = TknGreat
				state = DoneState
				rollBackChar()
			}
		case InEqual:
			nextChar = getNextChar()
			if nextChar == '=' {
				token.Attribute += string(nextChar)
				token.Type = TknEq
				state = DoneState
			} else {
				token.Type = TknAssign
				state = DoneState
				rollBackChar()
			}
		case InNotEqual:
			nextChar = getNextChar()
			if nextChar == '=' {
				token.Attribute += string(nextChar)
				token.Type = TknNotEq
				state = DoneState
			} else {
				token.Type = TknError
				state = DoneState
				rollBackChar()
			}
		default:
			{
				token.Type = TknError
				state = DoneState
			}
		}

	}

	if token.Type == TknWord {

		token.Type = isKeyWord(token.Attribute)
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
				//fmt.Println("Valio madre en la linea ", currentLine, ", '", currentToken.Attribute, "' no corresponde a ningun token")
				exitCode = 2
				errorsFile.WriteString("Valio madre en la linea " + strconv.Itoa(currentLine) + ", columna " + strconv.Itoa(currentColumn) + ". '" + currentToken.Attribute + "' no corresponde a ningun token.\n")
			} else {
				//fmt.Println(currentToken)
				outputFile.WriteString(strconv.Itoa(currentToken.Type) + " " + string(currentToken.Attribute) + " " + strconv.Itoa(currentToken.Row) + " " + strconv.Itoa(currentToken.Column) + "\n")
			}
			currentToken = getToken()
		}
		outputFile.WriteString(strconv.Itoa(currentToken.Type) + " " + "EOF" + " " + strconv.Itoa(currentToken.Row) + " " + strconv.Itoa(currentToken.Column) + "\n")
	}

	os.Exit(exitCode)

}
