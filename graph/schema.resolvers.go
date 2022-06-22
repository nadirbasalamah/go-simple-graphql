package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/nadirbasalamah/go-simple-graphql/graph/generated"
	"github.com/nadirbasalamah/go-simple-graphql/graph/model"
)

func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*model.Product, error) {
	var newProduct *model.Product = &model.Product{
		ID:          fmt.Sprintf("T%d", rand.Int()),
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Quantity:    input.Quantity,
	}

	r.products = append(r.products, newProduct)

	return newProduct, nil
}

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	return r.products, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
