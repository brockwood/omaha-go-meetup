package main

import (
	"github.com/containous/yaegi/interp"
	"github.com/containous/yaegi/stdlib"
)

func main() {
	i := interp.New(interp.Options{})

	i.Use(stdlib.Symbols)

	_, err := i.Eval(`import "fmt"`)
	if err != nil {
		panic(err)
	}

	_, err = i.Eval(`fmt.Println("Hello, World")`)
	if err != nil {
		panic(err)
	}
}
