package exercise

import (
	"fmt"
	"os"
)

type MyError struct {
	Msg string
	File string
	Line int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s:%d: %s", e.File, e.Line, e.Msg)
}

func test() error {
	return &MyError{"Someting happended","server.go",42}
}

func main() {
	err := test()
	// 判断error的类型
	switch  err := err.(type) {
	case nil:
		// call succeeded, nothing to do
	case *MyError:
			fmt.Println("error occurred on line:", err.Line)
	default:
		// unknown error
		fmt.Println("unknown error")
	}

	const name, age = "Kim", 22
	n, err := fmt.Fprintf(os.Stdout, "%s is %d years old.\n", name, age)

	// The n and err return values from Fprintf are
	// those returned by the underlying io.Writer.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
	}
	fmt.Printf("%d bytes written.\n", n)


}
