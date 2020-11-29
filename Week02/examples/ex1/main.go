package main

import (
	"errors"
	"fmt"
	"regexp"
)

type errorString string

func (e errorString) Error() string {
	return string(e)
}

// New 创建一个错误
func New(text string) error {
	return errorString(text)
}

var ErrNamedType = New("EOF")
var ErrStructType = errors.New("EOF")

func main() {
	if ErrNamedType == New("EOF") {
		fmt.Println("ex1: Named Type Error")
	}
	if ErrStructType == errors.New("EOF") {
		fmt.Println("ex2: Struct Type Error")
	}
	regexp.Compile("ddd")
}
