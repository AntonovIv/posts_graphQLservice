package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/AntonovIv/post_graphQlservice/graph"
	"github.com/AntonovIv/post_graphQlservice/internal/service/posts"
	graphTr "github.com/AntonovIv/post_graphQlservice/internal/transport/graph"
)

func main() {

	port := "8080"

	service := posts.New()

	resolver := graphTr.New(service)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
