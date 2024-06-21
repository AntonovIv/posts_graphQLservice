package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/AntonovIv/post_graphQlservice/graph"

	graphTr "github.com/AntonovIv/post_graphQlservice/internal/transport/graph"
)

func main() {

	port := "8080"
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graphTr.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
