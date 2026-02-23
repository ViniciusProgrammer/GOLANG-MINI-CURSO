package main

import "testing"

func assertMessage(t testing.TB, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
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
		want := "Bom dia, Rodolfo"
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