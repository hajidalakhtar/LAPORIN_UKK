package types

import (
	"github.com/graphql-go/graphql"
)

var LaporanTypes = graphql.NewObject(graphql.ObjectConfig{
	Name: "Products",
	Fields: graphql.Fields{
		"Id": &graphql.Field{
			Type: graphql.Int,
		},
		"Title": &graphql.Field{
			Type: graphql.String,
		},
		"User_id": &graphql.Field{
			Type: graphql.Int,
		},
		"Laporan": &graphql.Field{
			Type: graphql.String,
		},
		"Username": &graphql.Field{
			Type: graphql.String,
		},
		"User_Foto": &graphql.Field{
			Type: graphql.String,
		},
		"Foto": &graphql.Field{
			Type: graphql.String,
		},
		"FullName": &graphql.Field{
			Type: graphql.String,
		},
		"Kategori": &graphql.Field{
			Type: graphql.String,
		},
		"Time": &graphql.Field{
			Type: graphql.String,
		},
		"Status": &graphql.Field{
			Type: graphql.String,
		},
	},
})
