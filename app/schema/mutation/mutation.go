// package mutation
//
// import (
// 	"GO-Mysql-grphql/schema/types"
//
// 	"github.com/graphql-go/graphql"
// )
//
// var Mutation = graphql.NewObject(graphql.ObjectConfig{
// 	Name: "Mutation",
// 	Fields: graphql.Fields{
// 		"CreateLaporan": &graphql.Field{
// 			Type: graphql.NewList(types.LaporanTypes),
// 			//config param argument
// 			Args: graphql.FieldConfigArgument{
//
// 				"Title": &graphql.ArgumentConfig{
// 					Type: graphql.NewNonNull(graphql.String),
// 				},
// 				"Laporan": &graphql.ArgumentConfig{
// 					Type: graphql.NewNonNull(graphql.String),
// 				},
// 				"User_id": &graphql.ArgumentConfig{
// 					Type: graphql.NewNonNull(graphql.Int),
// 				},
// 				"Username": &graphql.ArgumentConfig{
// 					Type: graphql.NewNonNull(graphql.String),
// 				},
// 				"User_Foto": &graphql.ArgumentConfig{
// 					Type: graphql.NewNonNull(graphql.String),
// 				},
// 				"Foto": &graphql.ArgumentConfig{
// 					Type: graphql.NewNonNull(graphql.String),
// 				},
// 				"FullName": &graphql.ArgumentConfig{
// 					Type: graphql.NewNonNull(graphql.String),
// 				},
// 				"Kategori": &graphql.ArgumentConfig{
// 					Type: graphql.NewNonNull(graphql.String),
// 				},
// 				"Time": &graphql.ArgumentConfig{
// 					Type: graphql.NewNonNull(graphql.String),
// 				},
// 				"Status": &graphql.ArgumentConfig{
// 					Type: graphql.NewNonNull(graphql.String),
// 				},
// 			},
// 			Resolve: CreateLaporanMutation,
// 		},
// 	},
// })
