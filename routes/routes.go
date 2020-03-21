package routes

import (
	"laporin_go/app/http/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func SetRoutes() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/api/users", controller.GetUsers).Methods("GET")
	r.HandleFunc("/api/login", controller.PostLogin).Methods("POST")
	r.HandleFunc("/api/post/laporan", controller.PostLaporan).Methods("POST")
	r.HandleFunc("/api/post/ganti/status", controller.PostGantiStatus).Methods("POST")
	r.HandleFunc("/api/delete/{id}", controller.GetDeleteLaporan).Methods("GET")
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
