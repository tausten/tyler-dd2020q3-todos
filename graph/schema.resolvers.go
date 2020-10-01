package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/tausten/tyler-dd2020q3-todos/graph/generated"
	"github.com/tausten/tyler-dd2020q3-todos/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) CreateConstellation(ctx context.Context, input model.CreateConstellationInput) (*model.CreateConstellationPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateConstellations(ctx context.Context, input []*model.CreateConstellationInput) (*model.CreateConstellationsPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) Galaxies(ctx context.Context, after *string, before *string, first *int, last *int) (*model.GalaxyConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Constellations(ctx context.Context, after *string, before *string, first *int, last *int) (*model.ConstellationConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Planets(ctx context.Context, after *string, before *string, first *int, last *int, planetType *model.PlanetTypeEnum) (*model.PlanetConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Stars(ctx context.Context, after *string, before *string, first *int, last *int) (*model.StarConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
