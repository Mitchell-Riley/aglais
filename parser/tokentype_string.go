// Code generated by "stringer -type=TokenType"; DO NOT EDIT

package parser

import "fmt"

const _TokenType_name = "expressionmessageargumentsargumentsymbol"

var _TokenType_index = [...]uint8{0, 10, 17, 26, 34, 40}

func (i TokenType) String() string {
	if i < 0 || i >= TokenType(len(_TokenType_index)-1) {
		return fmt.Sprintf("TokenType(%d)", i)
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}
