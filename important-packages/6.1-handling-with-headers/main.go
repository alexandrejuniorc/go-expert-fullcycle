package main

import "net/http"

func main() {

	// Define a rota "/" e associa à função BuscaCEP
	http.HandleFunc("/", BuscaCEPHandler)

	// Inicia um servidor HTTP na porta 8080
	http.ListenAndServe(":8080", nil)
}

// BuscaCEP responde às requisições na rota "/"
// Funcionalidade separada conhecida também como controller ou handler
func BuscaCEPHandler(w http.ResponseWriter, r *http.Request) {
	// Verifica se a rota é exatamente "/"
	// Se não for, retorna 404 Not Found
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Exemplo simples: lê o parâmetro "cep" da query string
	// Ex: /?cep=01310-100
	// Note que este exemplo não faz validação ou busca real de CEP
	// Apenas demonstra como ler parâmetros e responder
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Responde com um JSON simples
	w.Header().Set("Content-Type", "application/json")
	// Retorna status 200 OK
	w.WriteHeader(http.StatusOK)
	// Corpo da resposta
	w.Write([]byte(`{"message": "Hello, World!"}`))
}
