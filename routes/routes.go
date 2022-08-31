package routes

import (
	"Alura-API_go/controllers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func HandleResquest() {
	r := mux.NewRouter()
	Port := os.Getenv("PORT")
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	log.Fatal(http.ListenAndServe(":"+Port, r))
}
