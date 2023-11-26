package routes

import (
	"api/rest/controllers"
	"log"
	"net/http"
)

func HandleResquest() { // Função que irá redirecionar para as controllers
	http.HandleFunc("/", controllers.Home)                                  // Rota principal que irá redirencionar para funçao home da homeController.go
	http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades) // Rota que irá retorna todas as personalidades
	log.Fatal(http.ListenAndServe(":8000", nil))                            // Subir um servidor.
}
