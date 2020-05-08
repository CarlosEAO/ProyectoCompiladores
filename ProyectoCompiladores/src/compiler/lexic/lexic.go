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

var keywords = [...]string{"program", "if", "else", "fi", "do", "until", "while", "read", "write", "float", "int", "bool", "not", "and", "or"}

//TOKEN TYPES
const (
	TknProgram  = 0
	TknIf       = 1
	TknElse     = 2
	TknFi       = 3
	TknDo       = 4
	TknUntil    = 5
	TknWhile    = 6
	TknRead     = 7
	TknWrite    = 8
	TknFloat    = 9
	TknInt      = 10
	TknBool     = 11
	TknNot      = 12
	TknAnd      = 13
	TknOr       = 14
	TknSum      = 15
	TknSub      = 16
	TknMul      = 17
	TknDiv      = 18
	TknExp      = 19
	TknLess     = 20
	TknLessEq   = 21
	TknGreat    = 22
	TknGreatEq  = 23
	TknEq       = 24
	TknNotEq    = 25
	TknAssign   = 26
	TknSemi     = 27
	TknComma    = 28
	TknLeftPar  = 29
	TknRightPar = 30
	TknLeftBr   = 31
	TknRightBr  = 32
	TknError    = 33
	TknIdent    = 34
	TknEOF      = 35
	TknComment  = 36
	TknConst    = 37
	TknWord     = 38
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
	return 38
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

		outputFile, err = os.Create("tokens.txt")
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

		currentToken = getToken()

		for currentToken.Type != TknEOF {
			if currentToken.Type == TknError {
				//fmt.Println("Valio madre en la linea ", currentLine, ", '", currentToken.Attribute, "' no corresponde a ningun token")
				exitCode = 2
				errorsFile.WriteString("Valio madre en la linea " + strconv.Itoa(currentLine) + ", '" + currentToken.Attribute + "' no corresponde a ningun token.\n")
			} else {
				//fmt.Println(currentToken)
				outputFile.WriteString(strconv.Itoa(currentToken.Type) + " " + string(currentToken.Attribute) + "\n")
			}
			currentToken = getToken()
		}
	}

	os.Exit(exitCode)

}
