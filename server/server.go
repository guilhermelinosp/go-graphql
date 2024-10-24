package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var data = map[string]string{
	"1": "Hello, World!",
	"2": "GraphQL in Go!",
}

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"message": &graphql.Field{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if msg, ok := data[p.Args["id"].(string)]; ok {
						return msg, nil
					}
					return nil, nil
				},
			},
		},
	}),
})

func main() {
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request received from %s\n", r.RemoteAddr)
		h.ServeHTTP(w, r)
	})

	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
