// Author: Vincent Composieux <vincent.composieux@gmail.com>

package types

import (
	"github.com/graphql-go/graphql"
)

// User type definition.
type User struct {
	ID        int    `db:"id" json:"id"`
	Firstname string `db:"firstname" json:"firstname"`
	Lastname  string `db:"lastname" json:"lastname"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"password"`
	CreatedAt string `db:"created_at" json:"created_at"`
}

// UserType is the GraphQL schema for the user type.
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: graphql.Int},
		"firstname":  &graphql.Field{Type: graphql.String},
		"lastname":   &graphql.Field{Type: graphql.String},
		"email":      &graphql.Field{Type: graphql.String},
		"password":   &graphql.Field{Type: graphql.String},
		"created_at": &graphql.Field{Type: graphql.String},
		"roles": &graphql.Field{
			Type: graphql.NewList(RoleType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var roles []Role

				// userID := params.Source.(User).ID
				// Implement logic to retrieve user associated roles from user id here.

				return roles, nil
			},
		},
	},
})
