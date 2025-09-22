package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	// Loop para cada URL passada como argumento na linha de comando
	// O range itera sobre os argumentos passados na linha de comando, ignorando o primeiro (nome do programa)
	for _, cep := range os.Args[1:] {
		// Faz a requisição HTTP para a API do ViaCEP
		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição %s: %v\n", cep, err)
		}
		defer req.Body.Close() // Fecha o corpo da resposta ao final da função

		// Lê o corpo da resposta
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta %s: %v\n", cep, err)
		}

		// Faz o unmarshal do JSON para a struct ViaCEP
		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer unmarshal %s: %v\n", cep, err)
		}

		// Cria (ou sobrescreve) o arquivo cidade.txt e escreve os dados formatados
		file, err := os.Create("cidade.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo %s: %v\n", cep, err)
		}
		defer file.Close()

		// Escreve os dados formatados no arquivo
		_, err = file.WriteString(fmt.Sprintf("CEP: %s, localidade: %s, UF: %s\n", data.Cep, data.Localidade, data.Uf))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao escrever no arquivo %s: %v\n", cep, err)
		}

		fmt.Println("Arquivo criado com sucesso!")
		fmt.Println("Cidade:", data.Localidade)
	}
}
