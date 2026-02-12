package main

import "testing"

func TestHello(t * testing.T) {
	got := Hello("Web II") // é o que foi recebido de fato pelo retorno da função
	want := "Olá, Web II" // aqui é o que eu gostaria que fosse recebido

	if got != want {
		t.Errorf("got %q ant %q", got, want)
	}
}