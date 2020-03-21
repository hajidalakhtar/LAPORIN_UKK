package controller

import (
	"fmt"
	"laporin_go/database"
	"log"
	"net/http"
	"time"

	"github.com/foolin/goview"
	"github.com/gorilla/mux"

	"laporin_go/app/model"
)

// var ctx = context.Background()

func ViewBlogDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM `blog` WHERE id = ?", vars["id"])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []model.Blog

	for rows.Next() {
		var each = model.Blog{}
		var err = rows.Scan(&each.Id, &each.Title, &each.Isi, &each.Kategori, &each.Time)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	goview.Render(w, http.StatusOK, "details_blog.html", goview.M{
		"blog": result,
	})

}

func PostBlog(w http.ResponseWriter, r *http.Request) {

	title := r.FormValue("title")
	isi := r.FormValue("isi")
	kategori := r.FormValue("kategori")
	date := time.Now()
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO `blog`(`Id`, `Title`, `Isi`, `Kategori`, `Time`) VALUES (?,?,?,?,?)", nil, title, isi, kategori, date.Format("02-Jan-2006"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if err != nil {
		log.Fatal(err.Error())
	}

	http.Redirect(w, r, "/create/blog", 303)

}
func ViewCreateBlog(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BlogRow, err := db.Query("SELECT * FROM `blog` WHERE 1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer BlogRow.Close()

	var blog []model.Blog

	for BlogRow.Next() {
		var each = model.Blog{}
		var err = BlogRow.Scan(&each.Id, &each.Title, &each.Isi, &each.Kategori, &each.Time)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		blog = append(blog, each)
	}

	goview.Render(w, http.StatusOK, "create_blog.html", goview.M{
		"Blog": blog,
	})

}
