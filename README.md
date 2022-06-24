## go-simple-graphql

A simple GraphQL application written in Go with gqlgen.

## How to use

1. Run the application with this command.

```
go run server.go
```

2. Open the GraphQL playground. The URL for GraphQL playground can be checked in the console.

3. Try out some GraphQL queries.

- Query for creating a new product data

```graphql
mutation {
  createProduct(
    input: {
      name: "new product"
      price: 100
      quantity: 10
      description: "a new product"
    }
  ) {
    id
    name
    price
    quantity
    description
  }
}
```

- Query for fetching all product data. The required attributes are name, price and description.

```graphql
{
  products {
    name
    price
    description
  }
}
```
