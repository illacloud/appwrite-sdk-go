package appwrite

import (
)

// Graphql service
type Graphql struct {
	client Client
}

func NewGraphql(clt Client) *Graphql {
	return &Graphql{
		client: clt,
	}
}

// Query execute a GraphQL mutation.
func (srv *Graphql) Query(Query interface{}) (*ClientResponse, error) {
	path := "/graphql"
	
	params := map[string]interface{}{
		"query": Query,
	}

	headers := map[string]interface{}{
		"x-sdk-graphql": "true",
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// Mutation execute a GraphQL mutation.
func (srv *Graphql) Mutation(Query interface{}) (*ClientResponse, error) {
	path := "/graphql/mutation"
	
	params := map[string]interface{}{
		"query": Query,
	}

	headers := map[string]interface{}{
		"x-sdk-graphql": "true",
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}
