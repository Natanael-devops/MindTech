package main

import (
	"github.com/Natanael-devops/MindTech/controllers"
	"github.com/Natanael-devops/MindTech/database"
	"github.com/gin-gonic/gin"
)

// //go:embed templates/*
//var tmplEmbed embed.FS

// //go:embed assets/*
//var staticEmbedFS embed.FS

//type staticFS struct {
//	fs fs.FS
//}

//func (sfs *staticFS) Open(name string) (fs.File, error) {
//	return sfs.fs.Open(filepath.Join("assets", name))
//}

//var staticEmbed = &staticFS{staticEmbedFS}

func CarregarRotas() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.Static("/js", "./js")
	r.LoadHTMLGlob("templates/*")

	//templ := template.Must(template.New("").ParseFS(tmplEmbed, "templates/*.html"))
	//r.SetHTMLTemplate(templ)

	//r.StaticFS("assets/", http.FS(staticEmbed))

	r.GET("/", controllers.PaginaIndex)
	r.GET("/contato", controllers.PaginaContato)
	r.GET("/quemsomos", controllers.PaginaQuemSomos)
	r.POST("/contato", controllers.SalvarContato)

	r.NoRoute(controllers.RotaNaoEncontrada)

	r.Run(":7000")
}

func main() {
	CarregarRotas()
	database.ConectaComBancoDeDados()
}
