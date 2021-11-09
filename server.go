package main

import (
	"github.com/go-chi/chi"
	_ "github.com/golang-migrate/migrate/database"
	"github.com/gula16/hackernews/graph"
	_ "github.com/gula16/hackernews/graph"
	"github.com/gula16/hackernews/graph/generated"
	_ "github.com/gula16/hackernews/graph/generated"
	"github.com/gula16/hackernews/internal/auth"
	database "github.com/gula16/hackernews/internal/pkg/db/mysql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/gula16/hackernews/graph"
	_ "github.com/gula16/hackernews/graph/generated"
)


const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(auth.Middleware())

	database.InitDB()
	database.Migrate()
	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}