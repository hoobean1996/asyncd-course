package main

import (
	"context"
	"log"
	"net/http"

	"entgo.io/ent/dialect"
	"fs.io/asyncd"
	"fs.io/asyncd/ent"
	"fs.io/asyncd/ent/migrate"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/rs/cors"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Create ent.Client and run the schema migration.
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	{
		client.EntTask.Create().SetHandler("echo1").SetPriority(0).SetParameter("hello").Save(context.Background())
	}
	{
		client.EntTask.Create().SetHandler("echo2").SetPriority(1).SetParameter("world").Save(context.Background())
	}
	{
		client.EntTask.Create().SetHandler("echo3").SetPriority(1).SetParameter("xxx").Save(context.Background())
	}
	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)
	// Configure the server and start listening on :8081.
	srv := handler.NewDefaultServer(asyncd.NewSchema(client))
	router.Handle("/", playground.Handler("Asynd", "/graphql"))
	router.Handle("/graphql", srv)
	if err := http.ListenAndServe(":8081", router); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
