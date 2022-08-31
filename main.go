package main

import (
	"Alura-API_go/database"
	"Alura-API_go/models"
	"Alura-API_go/routes"
	"log"
)

func main() {
	log.Print("Iniciando o servidor Rest com Go")
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	database.ConectaComBancoDeDados()
	routes.HandleResquest()
}
