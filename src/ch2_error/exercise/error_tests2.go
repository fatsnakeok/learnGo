package exercise

import (
	"fmt"
)


type errorString1 struct {
	s string
}

func (e errorString1) Error() string {
	return e.s
}

func NewError(text string) error {
	return errorString1{text}
}

var ErrType = NewError("EOF")

func main1() {
	if ErrType == NewError("EOF") {
		fmt.Println("Error:", ErrType)
	}

}
