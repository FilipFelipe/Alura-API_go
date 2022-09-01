package main

import (
	"Alura-API_go/database"
	"Alura-API_go/routes"
	"log"
)

func main() {
	log.Print("UTC [INFO] LOG: Iniciando o servidor Rest com Go")
	database.ConectaComBancoDeDados()
	routes.HandleResquest()
}
