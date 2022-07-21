package main

import (
	"net/http"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/nadirbasalamah/go-simple-graphql/graph"
	"github.com/nadirbasalamah/go-simple-graphql/graph/generated"
	"github.com/steinfletcher/apitest"
)

func graphQLHandler() http.Handler {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	mux := http.NewServeMux()

	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	return mux
}

func TestGetProducts_Success(t *testing.T) {
	var query string = `{
		products {
			name
			price
			quantity
		}
	}`

	var result string = `{
		"data": {
		  "products": []
		}
	  }`

	apitest.New().
		Handler(graphQLHandler()).
		Post("/query").
		GraphQLQuery(query).
		Expect(t).
		Status(http.StatusOK).
		Body(result).
		End()
}

func TestCreateProduct_Success(t *testing.T) {
	var query string = `mutation {
		createProduct(input:{
		  name:"new product",
		  price:100,
		  quantity:100,
		  description:"a new product"
		}) {
		  name
		  price
		}
	  }`

	var result string = `{
		"data": {
		  "createProduct": {
			"name": "new product",
			"price": 100
		  }
		}
	  }`

	apitest.New().
		Handler(graphQLHandler()).
		Post("/query").
		GraphQLQuery(query).
		Expect(t).
		Status(http.StatusOK).
		Body(result).
		End()
}

func TestCreateProduct_Failed(t *testing.T) {
	var query string = `mutation {
		createProduct(input:{
		  name:"new product"
		}) {
		  name
		  price
		}
	  }`

	apitest.New().
		Handler(graphQLHandler()).
		Post("/query").
		GraphQLQuery(query).
		Expect(t).
		Status(http.StatusUnprocessableEntity).
		End()
}
