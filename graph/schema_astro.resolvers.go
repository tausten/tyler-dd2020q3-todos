package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/tausten/tyler-dd2020q3-todos/graph/generated"
	"github.com/tausten/tyler-dd2020q3-todos/graph/model"
)

func (r *subscriptionResolver) Constellations(ctx context.Context) (<-chan []*model.Constellation, error) {
	panic(fmt.Errorf("not implemented"))
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
