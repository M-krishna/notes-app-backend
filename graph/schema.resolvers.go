package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/M-krishna/notes-app-backend/graph/generated"
	"github.com/M-krishna/notes-app-backend/graph/model"
)

func (r *mutationResolver) CreateNote(ctx context.Context, input *model.NewNote) (*model.Note, error) {
	user := &model.User{
		ID:   fmt.Sprintf("T%d", rand.Int()),
		Name: input.UserName,
	}

	note := &model.Note{
		ID:      fmt.Sprintf("T%d", rand.Int()),
		Title:   input.Title,
		Content: input.Content,
		User:    user,
	}

	r.notes = append(r.notes, note)
	return note, nil
}

func (r *mutationResolver) DeleteNote(ctx context.Context, noteID string) (string, error) {
	var notes []*model.Note
	for _, v := range r.notes {
		if v.ID != noteID {
			notes = append(notes, v)
		}
	}
	r.notes = notes
	return noteID, nil
}

func (r *queryResolver) Notes(ctx context.Context) ([]*model.Note, error) {
	return r.notes, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
