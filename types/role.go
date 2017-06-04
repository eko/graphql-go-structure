// Author: Vincent Composieux <vincent.composieux@gmail.com>

package types

import (
	"github.com/graphql-go/graphql"
)

// Role type definition.
type Role struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

// RoleType is the GraphQL schema for the user type.
var RoleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Role",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.Int},
		"name": &graphql.Field{Type: graphql.String},
	},
})
