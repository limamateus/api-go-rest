package controllers

import (
	"api/rest/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page") // Função da rota principal irá retorna apenas um texto home na tela
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(models.Personalidades) // Aqui irá retornar todas as personalinalidades que deixe mocadas.
}
