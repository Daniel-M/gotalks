package main

import (
	"fmt"
)

// Greeter greets
func Greeter(lang string) string {
	return fmt.Sprintf("Hello from %s\n", lang)
}

func main() {
	lang := "Go"
	fmt.Printf(Greeter(lang))
}
