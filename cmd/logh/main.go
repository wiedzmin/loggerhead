package main

import (
	"fmt"

	"github.com/wiedzmin/loggerhead/impl"
)

func main() {
	e, err := impl.Capture()
	if err != nil {
		fmt.Println("error capturing stats")
	} else {
		e.Print()
	}
}
