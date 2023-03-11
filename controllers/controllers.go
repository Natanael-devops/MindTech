package controllers

import (
	"net/http"

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
