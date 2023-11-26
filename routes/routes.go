package routes

import (
	"api/rest/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleResquest() { // Função que irá redirecionar para as controllers
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Home).Methods("Get")                                  // Rota principal que irá redirencionar para funçao home da homeController.go
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get") // Rota que irá retorna todas as personalidades
	r.HandleFunc("/api/personalidades/{id}", controllers.PersonalidadePorId).Methods("Get")
	r.HandleFunc("/api/personalidades/nova", controllers.NovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/deletar/{id}", controllers.DeletarPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/atualizar/{id}", controllers.AtualizarPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r)) // Subir um servidor.
}
