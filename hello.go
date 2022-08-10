package main

import "fmt"

const helloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {

	prefex := helloPrefix
	if language == "Spanish" {
		prefex = spanishHelloPrefix
	}

	if name == "" {
		name = "World"
	}
	return prefex + name
}

func main() {
	fmt.Println(Hello("Steve", ""))
}
