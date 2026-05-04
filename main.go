package main

import (
	"os"

	"mi-api-go/db"
	"mi-api-go/handler"
	"mi-api-go/service"

	"github.com/gin-gonic/gin"

	repoPostgres "mi-api-go/repository/postgres"
	repoSurreal "mi-api-go/repository/surrealdb"
)

func main() {
	usarDB := os.Getenv("DB_BACKEND")

	var productoHandler *handler.ProductoHandler

	if usarDB == "surreal" {
		db.ConectarSurreal(
			os.Getenv("SURREAL_URL"),
			os.Getenv("SURREAL_USER"),
			os.Getenv("SURREAL_PASS"),
			os.Getenv("SURREAL_NS"),
			os.Getenv("SURREAL_DB"),
		)
		repo := repoSurreal.NewProductoRepository()
		svc := service.NewProductosService(repo)
		productoHandler = handler.NewProductoHandler(svc)
	} else {
		db.ConectarPostgres(os.Getenv("POSTGRES_URL"))
		repo := repoPostgres.NewProductoRepository()
		svc := service.NewProductosService(repo)
		productoHandler = handler.NewProductoHandler(svc)
	}

	r := gin.Default()
	r.GET("/productos", productoHandler.GetAll)
	r.GET("/productos/:id", productoHandler.GetById)
	r.POST("/productos", productoHandler.Create)
	r.Run(":8080")
}
