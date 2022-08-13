package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/tokizuoh/contrail-server/graph/generated"
	"github.com/tokizuoh/contrail-server/graph/model"
)

// CreateWorkout is the resolver for the createWorkout field.
func (r *mutationResolver) CreateWorkout(ctx context.Context, input model.WorkoutInput) (*model.Workout, error) {
	panic(fmt.Errorf("not implemented"))
}

// Workouts is the resolver for the workouts field.
func (r *queryResolver) Workouts(ctx context.Context) ([]*model.Workout, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
