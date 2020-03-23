package query

import (
	"laporin_go/app/schema/types"

	"github.com/graphql-go/graphql"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"laporan": &graphql.Field{
			Type: graphql.NewList(types.LaporanTypes),
			Args: graphql.FieldConfigArgument{
				"Laporan_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: LaporanResolve,
		},
	},
})
