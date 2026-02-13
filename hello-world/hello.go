package main

import "fmt"

const prefixoPortugues = "Olá, "

// dominio
func Hello(name string) string {
	if name == "" {
		name = "Mundo"
	}
	
	return prefixoPortugues + name
}

// efeito colateral
func main() {
	fmt.Println(Hello("mundo"))
}