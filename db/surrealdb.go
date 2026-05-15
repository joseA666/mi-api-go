package db

import (
	"context"
	"log"

	"github.com/surrealdb/surrealdb.go"
)

var SurrealDB *surrealdb.DB

func ConectarSurreal(url, user, pass, ns, dbName string) {
	ctx := context.Background()

	db, err := surrealdb.FromEndpointURLString(ctx, url)
	if err != nil {
		log.Fatalf("Error conectando a SurrealDB: %v", err)
	}

	if err = db.Use(ctx, ns, dbName); err != nil {
		log.Fatalf("Error seleccionando namespace/db: %v", err)
	}

	if _, err = db.SignIn(ctx, &surrealdb.Auth{
		Username: user,
		Password: pass,
	}); err != nil {
		log.Fatalf("Error autenticando SurrealDB: %v", err)
	}

	SurrealDB = db
	log.Println("SurrealDB conectado")
}

git add .

git commit -m "docs: minor update

Co-authored-by: josea666 <arieldomingues88@gmail.com>"