package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/tokizuoh/contrail-server/graph/generated"
	"github.com/tokizuoh/contrail-server/graph/model"
)

// CreateWorkout is the resolver for the createWorkout field.
func (r *mutationResolver) CreateWorkout(ctx context.Context, input model.WorkoutInput) (*model.Workout, error) {
	workout := &model.Workout{
		Distance:  input.Distance,
		Duration:  input.Duration,
		StartDate: input.StartDate,
		Type:      input.Type,
		ID:        "TODO",
	}
	r.workouts = append(r.workouts, workout)
	return workout, nil
}

// Workouts is the resolver for the workouts field.
func (r *queryResolver) Workouts(ctx context.Context) ([]*model.Workout, error) {
	return r.workouts, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
