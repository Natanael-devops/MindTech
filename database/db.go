package database

import (
	"log"

	"github.com/Natanael-devops/MindTech/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=dpg-cesp9h82i3mh51vbsk00-a user=nate password=ujEUxrtXbBSH3kNZeyVLLY0I46hUPd3h dbname=freshair port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Contato{})

}
