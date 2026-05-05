package main

import (
	"embed"
	"io/fs"
	"net/http"
	"os"

	"mi-api-go/db"
	"mi-api-go/handler"
	"mi-api-go/service"

	"github.com/gin-gonic/gin"

	repoPostgres "mi-api-go/repository/postgres"
	repoSurreal  "mi-api-go/repository/surrealdb"
)

//go:embed frontend
var frontendFS embed.FS

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

	// Serve embedded frontend
	subFS, _ := fs.Sub(frontendFS, "frontend")

	serve := func(name, mime string) gin.HandlerFunc {
		return func(c *gin.Context) {
			data, err := fs.ReadFile(subFS, name)
			if err != nil {
				c.Status(http.StatusNotFound)
				return
			}
			c.Data(http.StatusOK, mime, data)
		}
	}

	r.GET("/", serve("index.html", "text/html; charset=utf-8"))
	r.GET("/style.css", serve("style.css", "text/css; charset=utf-8"))
	r.GET("/app.js", serve("app.js", "application/javascript; charset=utf-8"))

	// API routes
	api := r.Group("/api")
	api.GET("/productos", productoHandler.GetAll)
	api.GET("/productos/:id", productoHandler.GetById)
	api.POST("/productos", productoHandler.Create)
	api.PUT("/productos/:id", productoHandler.Update)
	api.DELETE("/productos/:id", productoHandler.Delete)

	r.Run(":8080")
}
