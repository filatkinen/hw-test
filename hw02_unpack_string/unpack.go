package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

// check if char == 0.
func isCharZero(char rune) bool {
	return char == '0'
}

// check if char from 1 to 9.
func isCharDigit(char rune) bool {
	return char >= '1' && char <= '9'
}

func Unpack(inputString string) (string, error) {
	var resultString = new(strings.Builder) //nolint:gofumpt
	var prevChar rune
	var skipLoop bool
	var wasSlash bool

	// check if string is empty
	if len(inputString) == 0 {
		return "", nil
	}

	for idx, char := range []rune(inputString + "_") {
		// add to the end of string inpustString  "_" for purpose not to repeat cheking condition outside the loop

		// first char has not to be digit is 0
		if idx == 0 && (isCharZero(char) || isCharDigit(char)) {
			return "", ErrInvalidString
		}

		// find  if sequence ==digit+0
		if isCharZero(char) && isCharDigit(prevChar) {
			return "", ErrInvalidString
		}

		// go to the next loop: we are adding to result checking  previous char due current char.
		if idx == 0 || skipLoop {
			prevChar = char
			skipLoop = false
			continue
		}

		switch {
		case prevChar == 'n' && char == '\\': // find if "\n" is in the string
			return "", ErrInvalidString
		case prevChar == '\\' && !wasSlash: // find if "\" is in the string and was not slash before
			prevChar = char
			wasSlash = true
		case isCharZero(char): // find if char is 0, so we are not  writing prevChar to result
			prevChar = char
			skipLoop = true
		case isCharDigit(char): // find if this char is digit-> add to result "previous char digit-1 copies"
			for i := 0; i < int(char)-'0'; i++ {
				resultString.WriteRune(prevChar)
			}
			prevChar = char
			skipLoop = true
		default:
			resultString.WriteRune(prevChar)
			if wasSlash {
				wasSlash = false
			}
			prevChar = char
		}
	}
	return resultString.String(), nil
}
