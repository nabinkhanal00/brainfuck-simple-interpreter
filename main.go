package main

import (
	"fmt"
	"io"
	"os"
)

type Token byte

var (
	INC_POINTER  Token = '>'
	DEC_POINTER  Token = '<'
	INC_BYTE     Token = '+'
	DEC_BYTE     Token = '-'
	OUTPUT       Token = '.'
	INPUT        Token = ','
	JUMP_FORWARD Token = '['
	JUMP_BACK    Token = ']'
	SPACE        Token = ' '
	TAB          Token = '\t'
	NEWLINE      Token = '\n'
)

func main() {
	const bufferSize int = 30000
	var buffer [bufferSize]byte
	if len(os.Args) != 2 {
		err := fmt.Errorf("invalid number of arguments: Expected 1, Got %d", len(os.Args)-1)
		fmt.Println(err)
		return
	}
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	code, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	dataPointer := 0
	instructionPointer := 0
	for instructionPointer < len(code) {
		token := Token(code[instructionPointer])
		switch Token(token) {
		case INC_POINTER:
			dataPointer++
			if dataPointer >= bufferSize {
				fmt.Println(fmt.Errorf("datapointer out of range: expected between 0 and %d, got: %d", bufferSize-1, dataPointer))
				return
			}
		case DEC_POINTER:
			dataPointer--
			if dataPointer < 0 {
				fmt.Println(fmt.Errorf("datapointer out of range: expected between 0 and %d , got: %d", bufferSize-1, dataPointer))
				return
			}
		case INC_BYTE:
			buffer[dataPointer]++
		case DEC_BYTE:
			buffer[dataPointer]--
		case OUTPUT:
			fmt.Print(string(buffer[dataPointer]))
		case INPUT:
			var value byte
			fmt.Scanf("%c", &value)
			buffer[dataPointer] = value
		case JUMP_FORWARD:
			if buffer[dataPointer] == 0 {
				instructionPointer += 1
				startPosition := instructionPointer
				count := 0
				for instructionPointer < len(code) {
					if code[instructionPointer] == byte(JUMP_FORWARD) {
						count++
					}
					if code[instructionPointer] == byte(JUMP_BACK) {
						if count == 0 {
							break
						} else {
							count--
						}
					}
					instructionPointer++
				}
				if instructionPointer == len(code) || code[instructionPointer] != byte(JUMP_BACK) {
					err = fmt.Errorf("syntax error: no matching ']' provided for %d", startPosition)
					fmt.Println(err)
					return
				}
			}
		case JUMP_BACK:
			if buffer[dataPointer] != 0 {
				instructionPointer -= 1
				startPosition := instructionPointer
				count := 0
				for instructionPointer > -1 {
					if code[instructionPointer] == byte(JUMP_BACK) {
						count++
					}
					if code[instructionPointer] == byte(JUMP_FORWARD) {
						if count == 0 {
							break
						} else {
							count--
						}
					}
					instructionPointer--
				}
				if instructionPointer == -1 || code[instructionPointer] != byte(JUMP_FORWARD) {
					err = fmt.Errorf("syntax error: no matching '[' provided for %d", startPosition)
					fmt.Println(err)
					return
				}
			}
		default:
			Nothing()
		}
		instructionPointer++
	}
}

func TODO() {
	fmt.Println("TODO")
	os.Exit(1)
}

func Nothing() {
}
