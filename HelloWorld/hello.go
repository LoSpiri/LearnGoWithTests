package main

import (
	"fmt"
)

const helloPrefix = "Hello "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := helloPrefix
	switch language {
	case "Spanish":
		prefix = "Hola "
	case "Italian":
		prefix = "Ciao "
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("LoSpiri", "English"))
}
