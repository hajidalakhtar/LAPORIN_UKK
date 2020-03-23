package routes

import (
	"encoding/json"
	"fmt"
	"laporin_go/app/http/controller"
	"laporin_go/app/schema/query"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func SetRoutes() *mux.Router {
	var Schame, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: query.RootQuery,
		// Mutation: mutation.Mutation,
	})
	if err != nil {
		panic(err.Error())
	}

	r := mux.NewRouter()

	r.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("Query"), Schame)
		json.NewEncoder(w).Encode(result)
	})
	// graphql api webservice http://localhost:9000/api/graphql?Query={laporan(Laporan_id:28){Id,Title,Laporan}}

	r.HandleFunc("/api/users", controller.GetUsers).Methods("GET")
	r.HandleFunc("/api/login", controller.PostLogin).Methods("POST")
	r.HandleFunc("/api/post/laporan", controller.PostLaporan).Methods("POST")
	r.HandleFunc("/api/post/ganti/status", controller.PostGantiStatus).Methods("POST")
	r.HandleFunc("/api/delete/{id}", controller.GetDeleteLaporan).Methods("GET")
	r.HandleFunc("/api/delete/user/{id}", controller.GetDeleteUser).Methods("GET")

	r.HandleFunc("/api/register", controller.PostRegister).Methods("POST")
	r.HandleFunc("/api/post/edit/user", controller.PostEditUser).Methods("POST")
	r.HandleFunc("/api/post/blog", controller.PostBlog).Methods("POST")
	r.HandleFunc("/api/post/like/{id}/{user_id}", controller.PostLike).Methods("GET")
	r.HandleFunc("/api/post/follow/{id}", controller.PostFollow).Methods("GET")
	r.HandleFunc("/api/post/unfollow/{user_id}/{target_id}", controller.PostUnfollow).Methods("GET")

	r.HandleFunc("/api/post/unlike/{id}/{user_id}", controller.PostUnlike).Methods("GET")

	r.HandleFunc("/api/post/tambah/komentar/laporan/{id}", controller.PostTambahKomentar).Methods("POST")

	r.HandleFunc("/profile/{username}", controller.ViewProfile).Methods("GET")

	r.HandleFunc("/home", controller.ViewLandingPage).Methods("GET")
	r.HandleFunc("/generate", controller.ViewGenetateExcel).Methods("GET")

	r.HandleFunc("/admin", controller.ViewAdmin).Methods("GET")
	r.HandleFunc("/staff", controller.ViewStaff).Methods("GET")
	r.HandleFunc("/admin/user", controller.ViewAdminUser).Methods("GET")
	r.HandleFunc("/login", controller.ViewLogin).Methods("GET")
	r.HandleFunc("/register", controller.ViewRegister).Methods("GET")
	r.HandleFunc("/search", controller.ViewCari).Methods("GET")

	r.HandleFunc("/post/logout", controller.PostLogout).Methods("GET")
	r.HandleFunc("/d/{id}", controller.GetDetailsLaporan).Methods("GET")
	r.HandleFunc("/blog/{id}", controller.ViewBlogDetails).Methods("GET")
	r.HandleFunc("/create/blog", controller.ViewCreateBlog).Methods("GET")

	r.HandleFunc("/user", controller.ViewHome).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	return r
}
