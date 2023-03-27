package handler

import (
	"fmt"
)

var ErrorHandler = func(err any) {
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
}
