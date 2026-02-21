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
		got := Hello("Web II")
		want := "Olá, Web II"

		assertMessage(t, got, want)
	})

	t.Run("usar 'Mundo' quando a string for vazia", func(t *testing.T) {
		got := Hello("")
		want := "Olá, Mundo"

		assertMessage(t, got, want)
	})
}