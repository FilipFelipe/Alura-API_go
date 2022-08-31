package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func carregaEnv() string {
	host := os.Getenv("HOST_DB")
	user := os.Getenv("USER_DB")
	password := os.Getenv("PASSWORD_DB")
	dbname := os.Getenv("NAME_DB")
	port := os.Getenv("PORT_DB")
	stringDeConexao := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable"
	return stringDeConexao
}

func ConectaComBancoDeDados() {
	stringDeConexao := carregaEnv()
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	} else {
		log.Print("Conectado com sucesso ao banco de dados")
	}

}
