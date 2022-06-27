package database

import (
	"log"

	"github.com/PabloCassio/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexão := "host=localhost user=postgres password=fla12198130 dbname=alunos port=5432  sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexão))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
}
