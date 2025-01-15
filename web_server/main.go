package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	// Define o handler para a rota principal "/"
	http.HandleFunc("/", handleHome)

	// Inicia o servidor na porta 8080
	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	// HTML da página de boas-vindas
	html := `
<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <title>Bem-vindo!</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
            background-color: #f0f2f5;
        }
        .container {
            text-align: center;
            padding: 20px;
            background-color: white;
            border-radius: 8px;
            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
        }
        h1 {
            color: #1a73e8;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Bem-vindo ao Servidor Web Go!</h1>
        <p>Esta é uma página simples criada com Go.</p>
    </div>
</body>
</html>
`
	// Parse e executa o template
	tmpl, err := template.New("home").Parse(html)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
