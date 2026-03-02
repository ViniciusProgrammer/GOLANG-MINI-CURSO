package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerPing(t *testing.T) {
	requisitado := httptest.NewRequest("GET", "/ping", nil) // cria uma requisição falsa
	recebido := httptest.NewRecorder()                      // grava o que vem de resposta

	handlerPing(recebido, requisitado) // chama a função

	resultado := recebido.Body.String()
	esperado := "pong"

	if resultado != esperado {
		t.Errorf("resultado %q esperado %q", resultado, esperado)
	}

	if recebido.Code != http.StatusOK {
		t.Errorf("status: resultado %d, esperado %d", recebido.Code, http.StatusOK)
	}
}

func verificarBody(t testing.TB, recebido *httptest.ResponseRecorder, esperado string) {
	t.Helper()
	resultado := recebido.Body.String()

	if resultado != esperado {
		t.Errorf("body: resultado %q, esperado %q", resultado, esperado)
	}
}

func TestHandleCelsius(t *testing.T) {
	t.Run("ponto de congelamento", func(t *testing.T) {
		requisitado := httptest.NewRequest("GET", "/celsius?valor=0", nil)
		recebido := httptest.NewRecorder()
		handlerCelsius(recebido, requisitado)
		verificarBody(t, recebido, "32ºF")
	})

	t.Run("ponto de ebulição", func(t *testing.T) {
		requisitado := httptest.NewRequest("GET", "/celsius?valor=100", nil)
		recebido := httptest.NewRecorder()
		handlerCelsius(recebido, requisitado)
		verificarBody(t, recebido, "212ºF")
	})

	t.Run("sem parâmetro usa zero", func(t *testing.T) {
		requisitado := httptest.NewRequest("GET", "/celsius", nil)
		recebido := httptest.NewRecorder()
		handlerCelsius(recebido, requisitado)
		verificarBody(t, recebido, "32ºF")
	})
}
