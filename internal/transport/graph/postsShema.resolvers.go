package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"github.com/AntonovIv/post_graphQlservice/graph"
)

// Comment returns graph.CommentResolver implementation.
func (r *Resolver) Comment() graph.CommentResolver { return &commentResolver{r} }

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Post returns graph.PostResolver implementation.
func (r *Resolver) Post() graph.PostResolver { return &postResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

// Subscription returns graph.SubscriptionResolver implementation.
func (r *Resolver) Subscription() graph.SubscriptionResolver { return &subscriptionResolver{r} }

type commentResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
