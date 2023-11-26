package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ( // Variavel de inicialização do banco
	DB  *gorm.DB
	err error
)

func ConexaoComBanco() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable" // String de conexão com banco

	DB, err = gorm.Open(postgres.Open(stringDeConexao)) // abrindo a conexão com banco

	if err != nil {
		log.Panic("Erro de conexão com banco!")
	}

}
