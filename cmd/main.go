package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/AntonovIv/post_graphQlservice/graph"
	manager "github.com/AntonovIv/post_graphQlservice/internal/db/postgre"

	repo "github.com/AntonovIv/post_graphQlservice/internal/repository/postgre"
	"github.com/AntonovIv/post_graphQlservice/internal/service/posts"
	"github.com/AntonovIv/post_graphQlservice/internal/service/posts/subscription"
	graphTr "github.com/AntonovIv/post_graphQlservice/internal/transport/graph"
)

func main() {

	ctx := context.Background()
	port := "8080"

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	testcfg := manager.Config{
		User:     "user",
		Password: "password",
		Hostname: "localhost",
		Port:     5432,
		Name:     "record_db",
	}

	postgresManager, err := manager.New(ctx, testcfg)
	if err != nil {
		log.Fatal("failed to create postgres manager: ", err)
	}

	postgresRepo := repo.New(postgresManager)
	obsPool := subscription.NewObserverPool()
	service := posts.New(postgresRepo, obsPool)

	resolver := graphTr.New(service, logger)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))
	srv.Use(extension.FixedComplexityLimit(25))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
