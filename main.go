package main

import (
	"api/rest/database"
	"api/rest/routes"
	"fmt"
)

func main() {
	//models.Personalidades = []models.Personalidade{ // mocando os dados
	//		{Id: 1, Nome: "Nome 1", Historio: "Historia 1"},
	//	{Id: 2, Nome: "Nome 2", Historio: "Historia 2"},
	//		{Id: 3, Nome: "Nome 3", Historio: "Historia 3"},
	//	}

	fmt.Println("Inicializando o servidor Rest Com Go")

	database.ConexaoComBanco() // Inicialização com banco

	routes.HandleResquest()
}
