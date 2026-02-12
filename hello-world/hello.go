package main

import "fmt"

// dominio
func Hello() string {
	return "Olá, mundo"
}

// efeito colateral
func main() {
	fmt.Println(Hello())
}