package main

import "testing"

func assertMessage(t testing.TB, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloWithRegion(t * testing.T) {
	t.Run("Rio Grande do Norte morning", func(t *testing.T) {
		got := HelloWithRegion(9, RegionRN)
		want := "Bom Dia, boy"
		assertMessage(t, got, want)
	})

	t.Run("São Paulo afternoon", func(t *testing.T) {
		got := HelloWithRegion(14, RegionSP)
		want := "Boa Tarde, mano"
		assertMessage(t, got, want)
	})

	t.Run("Minas Gerais morning", func(t *testing.T) {
		got := HelloWithRegion(10, RegionMG)
		want := "Bom Dia, sô"
		assertMessage(t, got, want)
	})

	t.Run("Rio Grande do Sul evening", func(t *testing.T) {
		got := HelloWithRegion(22, RegionRS)
		want := "Boa Noite, tchê"
		assertMessage(t, got, want)
	})
}

// Função de teste
func TestHello(t *testing.T) {
	// sub-teste para retorno da função que recebe uma string não vazia
	t.Run("cumprimentar pessoas", func(t *testing.T) {
		got := Hello("Maria")
		want := "Olá, Maria"

		assertMessage(t, got, want)
	})

	t.Run("usar 'Mundo' quando a string for vazia", func(t *testing.T) {
		got := Hello("")
		want := "Olá, Mundo"

		assertMessage(t, got, want)
	})

	t.Run("saudação de manhã", func(t *testing.T) {
		hour := 9
		got := HelloWithTime("Rodolfo", hour)
		want := "Bom Dia, Rodolfo"
		assertMessage(t, got, want)
	})

	t.Run("saudação de tarde", func(t *testing.T) {
		hour := 16
		got := HelloWithTime("Maria", hour)
		want := "Boa Tarde, Maria"
		assertMessage(t, got, want)
	})

	t.Run("saudação da noite", func(t *testing.T) {
		hour := 21
		got := HelloWithTime("Pedro", hour)
		want := "Boa Noite, Pedro"
		assertMessage(t, got, want)
	})
}