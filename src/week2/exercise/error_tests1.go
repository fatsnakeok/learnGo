package exercise

import (
	"errors"
	"fmt"
)

type errorString string

func (e errorString) Error() string {
	return string(e)
}

func New(text string) error {
	return errorString(text)
}

var ErrNameType = New("EOF")
var ErrStructType = errors.New("EOF")

func main2() {
	if ErrNameType == New("EOF") {
		fmt.Println("Named Type Error")
	}
	if ErrStructType == errors.New("EOF") {
		fmt.Println("struct Type Error")
	}

}
