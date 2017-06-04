// Author: Vincent Composieux <vincent.composieux@gmail.com>

package mutations

import (
	"github.com/graphql-go/graphql"
)

// GetRootFields returns all the available mutations.
func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"createUser": GetCreateUserMutation(),
	}
}
