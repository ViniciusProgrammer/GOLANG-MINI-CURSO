package main

import "testing"

func assertMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
			t.Errorf("got %q want %q", got, want)
		}
}

func TestHello(t *testing.T) {
	t.Run("cumprimentar pessoas", func(t *testing.T) {
		got := Hello("Web II") // é o que foi recebido de fato pelo retorno da função
		want := "Olá, Web II"  // aqui é o que eu gostaria que fosse recebido
		
		assertMessage(t, got, want)
	})

	t.Run("usar 'Mundo' quando string for vazia", func(t *testing.T) {
		got := Hello("")
		want := "Olá, Mundo"

		assertMessage(t, got, want)
	})
}