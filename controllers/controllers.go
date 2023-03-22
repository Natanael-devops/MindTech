package controllers

import (
	"net/http"

	"github.com/Natanael-devops/MindTech/database"
	"github.com/Natanael-devops/MindTech/models"
	"github.com/gin-gonic/gin"
)

func PaginaIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func PaginaContato(c *gin.Context) {
	c.HTML(http.StatusOK, "contato.html", nil)
}

func PaginaQuemSomos(c *gin.Context) {
	c.HTML(http.StatusOK, "quemsomos.html", nil)
}

func RotaNaoEncontrada(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}

func SalvarContato(c *gin.Context) {
	contato := models.Contato{}

	nome := c.PostForm("nome")
	email := c.PostForm("email")
	telefone := c.PostForm("telefone")

	database.DB.Where(&models.Contato{Telefone: contato.Telefone}).First(&contato.Telefone)

	if contato.ID == 0 {

		contato.Nome = nome
		contato.Email = email
		contato.Telefone = telefone

		if contato.Nome == "" || contato.Telefone == "" {
			c.HTML(http.StatusOK, "resultado.html", gin.H{
				"Nome":     contato.Nome,
				"Telefone": contato.Telefone,
				"Email":    contato.Email,
			})
			return
		}

		database.DB.Create(&contato)
		c.HTML(http.StatusOK, "index.html", nil)

	}
}
