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

// Função de visibilidade privada quando começa com letra caixa baixa
func getTimeGreeting(hour int) string {
	if hour >= 5 && hour < 12 {
		return "Bom dia"
	} else if hour >= 12 && hour < 18 {
		return "Boa Tarde"
	}

	return "Boa Noite"
}

// Função de visibilidade pública quando começa com letra caixa alta
func HelloWithTime(name string, hour int) string {
	if name == "" {
		name = "Mundo"
	}

	return getTimeGreeting(hour) + ", " + name
}

// efeito colateral
func main() {
	fmt.Println(Hello("mundo"))
}