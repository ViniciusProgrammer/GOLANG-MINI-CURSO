package main

import "fmt"

const prefixoPortugues = "Olá, "

const(
	RegionRN = "rn"
	RegionSP = "sp"
	RegionMG = "mg"
	RegionRS = "rs"
)

func getRegionalVocative(region string) string {
	switch region {
	case RegionRN:
			return "boy"
	case RegionSP:
		return "mano"
	case RegionMG:
		return "sô"
	case RegionRS:
		return "tchê"
	default:
		return "" // sem vocativo regional
	} 	
}

func HelloWithRegion(hour int, region string) string {
	greeting := getTimeGreeting(hour)
	vocative := getRegionalVocative(region)

	if vocative != "" {
		return greeting + ", " + vocative
	}

	return greeting + ", Mundo"
}

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
		return "Bom Dia"
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