package main

import (
	"Alura-API_go/models"
	"Alura-API_go/routes"
	"fmt"
)

func main() {
	fmt.Println("Iniciando o servidor REST com GO")
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome 2", Historia: "Historia 2"},
	}
	routes.HandleResquest()
}
