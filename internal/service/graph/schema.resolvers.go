package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/jaksonkallio/radiate/internal/service/graph/generated"
	"github.com/jaksonkallio/radiate/internal/service/graph/model"
)

// Scan is the resolver for the scan field.
func (r *queryResolver) Scan(ctx context.Context, originIds []string, query string, onlyStarred *bool) ([]*model.Media, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
