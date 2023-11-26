package controllers

import (
	"api/rest/database"
	"api/rest/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page") // Função da rota principal irá retorna apenas um texto home na tela
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var personalidades []models.Personalidade
	database.DB.Find(&personalidades)

	json.NewEncoder(w).Encode(personalidades) // Aqui irá retornar todas as personalinalidades que estão no banco.
}

func PersonalidadePorId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Aqui estou declarando uma variavel que irá obter informações sobre request

	id := vars["id"] // Aqui estou pegando o Id que esta sendo passado no request

	//for _, personalidade := range models.Personalidades { // realizo um for para indetificar a personilidade mocada
	//	if strconv.Itoa(personalidade.Id) == id { // Realizo um conversão de string para int usando o recurso Itoa
	//		json.NewEncoder(w).Encode(personalidade) // retorno um json a personalidade encontrada
	//	}
	//}
	var personalidade = models.Personalidade{} // Intacio uma variavel
	database.DB.First(&personalidade, id)      // aqui eu irei pegar os dados do banco e passar para variavel pesonalidade

	json.NewEncoder(w).Encode(personalidade) // retorno objeto json.
}
