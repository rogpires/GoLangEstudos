package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Endereco representa a estrutura de endereço
type Endereco struct {
	Rua    string `json:"rua"`
	Numero string `json:"numero"`
	Cidade string `json:"cidade"`
	Estado string `json:"estado"`
	CEP    string `json:"cep"`
}

// Contato representa informações de contato
type Contato struct {
	Telefone string `json:"telefone"`
	Celular  string `json:"celular"`
	Email    string `json:"email"`
}

// Pessoa representa a estrutura de dados completa
type Pessoa struct {
	Nome     string   `json:"nome"`
	Idade    int      `json:"idade"`
	CPF      string   `json:"cpf"`
	Endereco Endereco `json:"endereco"`
	Contato  Contato  `json:"contato"`
}

func getPessoas(w http.ResponseWriter, r *http.Request) {
	// Criando um array de pessoas com dados de exemplo
	pessoas := []Pessoa{
		{
			Nome:  "João Silva",
			Idade: 25,
			CPF:   "123.456.789-00",
			Endereco: Endereco{
				Rua:    "Rua das Flores",
				Numero: "123",
				Cidade: "São Paulo",
				Estado: "SP",
				CEP:    "01234-567",
			},
			Contato: Contato{
				Telefone: "(11) 3456-7890",
				Celular:  "(11) 98765-4321",
				Email:    "joao.silva@email.com",
			},
		},
		{
			Nome:  "Maria Santos",
			Idade: 30,
			CPF:   "987.654.321-00",
			Endereco: Endereco{
				Rua:    "Avenida Principal",
				Numero: "456",
				Cidade: "Rio de Janeiro",
				Estado: "RJ",
				CEP:    "20000-000",
			},
			Contato: Contato{
				Telefone: "(21) 2345-6789",
				Celular:  "(21) 98765-4321",
				Email:    "maria.santos@email.com",
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pessoas)
}

func main() {
	// Atualizando a rota para retornar múltiplas pessoas
	http.HandleFunc("/pessoas", getPessoas)

	fmt.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", nil)
}
