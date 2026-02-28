package main

import (
	"bytes"
	"testing"
)

func TestSaudar(t *testing.T) {
	buffer := bytes.Buffer{}
	Saudar(&buffer, WithName("Maria"))

	resultado := buffer.String()
	esperado := "Olá, Maria"

	if resultado != esperado {
		t.Errorf("resultado %q esperado %q", resultado, esperado)
	}
}
