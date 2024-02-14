package main

import (
	"bytes"
	"fmt"
)

// DependencyInjection

func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
