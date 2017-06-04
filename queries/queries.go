// Author: Vincent Composieux <vincent.composieux@gmail.com>

package queries

import (
	"github.com/graphql-go/graphql"
)

// GetRootFields returns all the available queries.
func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"user": GetUserQuery(),
	}
}
