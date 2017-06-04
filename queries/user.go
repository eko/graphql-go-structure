// Author: Vincent Composieux <vincent.composieux@gmail.com>

package queries

import (
	"../types"

	"github.com/graphql-go/graphql"
	"log"
)

// GetUserQuery returns the queries available against user type.
func GetUserQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.UserType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			log.Printf("[query] user\n")
			var users []types.User

			return users, nil
		},
	}
}
