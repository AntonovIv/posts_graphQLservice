package main

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/AntonovIv/post_graphQlservice/graph"
)

func main() {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema())

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
}
