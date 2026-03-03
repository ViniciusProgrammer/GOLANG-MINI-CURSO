package main

import (
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
)

func configurarRotas() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/ajuda", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Endpoints disponíveis: /ping, /celsius, /calcular"))
	})

	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		dados := map[string]interface{} {
			"online": true,
			"versao": "1.0",
		}

		json.NewEncoder(w).Encode(dados)
	})
	
	return mux
}

func handlerPing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func celsiusParaFahrenheit(c float64) float64 {
	return c * 9 / 5 + 32
}

func handlerCelsius(w http.ResponseWriter, r *http.Request) {
	valorStr := r.URL.Query().Get("valor")

	celsius := 0.0

	if valorStr != "" {
		celsius, _ = strconv.ParseFloat(valorStr, 64)
	}

	fmt.Fprintf(w, "%.0fºF", celsiusParaFahrenheit(celsius))
}
