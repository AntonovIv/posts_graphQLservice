package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/AntonovIv/post_graphQlservice/graph"

	"github.com/AntonovIv/post_graphQlservice/internal/config"
	"github.com/AntonovIv/post_graphQlservice/internal/repository/mem"
	"github.com/AntonovIv/post_graphQlservice/internal/service/posts"
	"github.com/AntonovIv/post_graphQlservice/internal/service/posts/subscription"
	graphTr "github.com/AntonovIv/post_graphQlservice/internal/transport/graph"

	manager "github.com/AntonovIv/post_graphQlservice/internal/db/postgre"
	"github.com/AntonovIv/post_graphQlservice/internal/repository/postgre"
)

const (
	memoryDB         = "memory"
	postgresDB       = "postgres"
	memoryDBCapacity = 100
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	ctx := context.Background()
	cfg := config.MustLoadCfg()
	var repository posts.PostRepo

	switch cfg.DbConfig.DbType {
	case memoryDB:
		repository = mem.NewInMemDb(memoryDBCapacity)
		logger.Info("Server uses in-memory database")
	case postgresDB:
		postgresManager, err := manager.New(ctx, cfg.DbConfig)
		if err != nil {
			log.Fatal("failed to create postgres manager: ", err)
		}
		repository = postgre.New(postgresManager)
		logger.Info("Server uses postgres database")
	}

	obsPool := subscription.NewObserverPool()
	service := posts.New(repository, obsPool)

	resolver := graphTr.New(service, logger)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))
	srv.Use(extension.FixedComplexityLimit(25))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	logger.Debug("server start",
		"port", cfg.Port,
		"dbType", cfg.DbConfig.DbType,
		"dbPort", cfg.DbConfig.Port)
	log.Printf("connect to http://localhost:%v/ for GraphQL playground", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(cfg.Port), nil))
}
