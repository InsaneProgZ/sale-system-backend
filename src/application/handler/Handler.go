package handler

import (
	"fmt"
	"go/printer"
)

var ErrorHandler = func(err any) {
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
}
