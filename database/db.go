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
	dsn := "postgres://nate:ujEUxrtXbBSH3kNZeyVLLY0I46hUPd3h@dpg-cesp9h82i3mh51vbsk00-a/freshair"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	// Verifica se a conexão com o banco de dados está funcionando

	DB.AutoMigrate(&models.Contato{})

}
