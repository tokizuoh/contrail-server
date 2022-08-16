package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strings"

	"github.com/tokizuoh/contrail-server/graph/generated"
	"github.com/tokizuoh/contrail-server/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	var workouts []*model.Workout
	for _, iw := range input.Workouts {
		nw := model.Workout{
			ID:       fmt.Sprintf("Workout:%d", r.wcnt),
			Distance: iw.Distance,
		}
		r.wcnt += 1
		r.workouts = append(r.workouts, &nw)
		workouts = append(workouts, &nw)
	}

	user := model.User{
		ID:       fmt.Sprintf("User:%d", r.ucnt),
		Name:     input.Name,
		Workouts: workouts,
	}
	r.ucnt += 1
	r.users = append(r.users, &user)
	return &user, nil
}

// Workouts is the resolver for the workouts field.
func (r *queryResolver) Workouts(ctx context.Context) ([]*model.Workout, error) {
	return r.workouts, nil
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	arr := strings.Split(id, ":")
	switch arr[0] {
	case "User":
		return r.User(ctx, id)
	case "Workout":
		return r.Workout(ctx, id)
	}
	panic(fmt.Errorf("不正なIDだよ〜"))
}

func (r *queryResolver) Workout(ctx context.Context, id string) (*model.Workout, error) {
	for _, workout := range r.workouts {
		if workout.ID == id {
			return workout, nil
		}
	}
	return nil, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
