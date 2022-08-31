package routes

import (
	"Alura-API_go/controllers"
	"log"
	"net/http"
	"os"
)

func HandleResquest() {
	Port := os.Getenv("PORT")
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)

	log.Fatal(http.ListenAndServe(":"+Port, nil))
}
