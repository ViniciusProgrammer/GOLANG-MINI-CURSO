package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRotas(t *testing.T) {
	mux := configurarRotas()

	t.Run("raiz", func(t *testing.T) {
		requisitado := httptest.NewRequest("GET", "/", nil)
		recebido := httptest.NewRecorder()
		mux.ServeHTTP(recebido, requisitado)

		verificarBody(r, recebido, "Calculadora API")
		verificarStatus(t, recebido, http.StatusOK)
	})

	t.Run("ajuda", func(t *testing.T) {
		requisitado := httptest.NewRequest("GET", "/ajuda", nil) 
		recebido := httptest.NewRecorder()
		mux.ServeHTTP(recebido, requisitado)

		verificarStatus(t, recebido, "disponiveis: /ping, /celsius, /calcular")
	})

	t.Run("status JSON", func(t *testing.T) {
		requisitado := httptest.NewRequest("GET", "/status", nil)
		recebido := httptest.NewRecorder()
		mux.ServeHTTP(recebido, requisitado)

		verificarStatus(t, recebido, http.StatusOK)

		contentType := recebido.Header().Get("Content-Type")

		if contentType != "application/json" {
			t.Errorf("Content-Type: resultado %q, esperado: %q", contentType, "application")
		}
	})

	t.Run("não encontrado", func(t *testing.T) {
		requisitado := httptest.NewRequest("GET", "/xyz", nil)
		recebido := httptest.NewRecorder()
		mux.ServeHTTP(recebido, requisitado)

		verificarStatus(t, recebido, http.StatusNotFound)
	})
}

func verificarStatus(t testing.TB, recebido *httptest.ResponseRecorder, esperado int) {
	t.Helper()

	if recebido.Code != esperado {
		t.Errorf("status: resultado %d, esperado %d", recebido.Code, esperado)
	}
}

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
