package controller

import (
	"fmt"
	"laporin_go/database"
	"net/http"

	"github.com/foolin/goview"

	"laporin_go/app/helper"
	"laporin_go/app/model"
)

// var ctx = context.Background()

func ViewAdmin(w http.ResponseWriter, r *http.Request) {
	session, errsession := store.Get(r, "login")
	if errsession != nil {
		fmt.Println("error")
	}
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()
	// get laporan
	LaporanRow, err := db.Query("SELECT * FROM `laporan` WHERE 1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer LaporanRow.Close()

	var laporan []model.Laporan

	for LaporanRow.Next() {
		var each = model.Laporan{}
		// var err = rows.Scan()
		var err = LaporanRow.Scan(&each.Id, &each.Title, &each.Laporan, &each.User_id, &each.Username, &each.User_Foto, &each.Foto, &each.FullName, &each.Kategori, &each.Time, &each.Status)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		laporan = append(laporan, each)
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
		var err = BlogRow.Scan(&each.Id, &each.User_id, &each.Title, &each.Isi, &each.Kategori, &each.Time)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		blog = append(blog, each)
	}

	// laporanCount := helper.GetCount("laporan")
	CountVerifikasi := helper.GetCountWhere("laporan", "Status", "Sedang Di Verifikasi")
	CountProcess := helper.GetCountWhere("laporan", "Status", "Sedang Process")
	CountFinish := helper.GetCountWhere("laporan", "Status", "Finish")
	CountReject := helper.GetCountWhere("laporan", "Status", "Reject")

	// laporanPersenFinish := CountFinish / laporanCount * 100

	goview.Render(w, http.StatusOK, "home_admin.html", goview.M{
		"laporan":         laporan,
		"role":            session.Values["role"],
		"username":        session.Values["username"],
		"alamat":          session.Values["alamat"],
		"email":           session.Values["email"],
		"fullname":        session.Values["fullname"],
		"CountVerifikasi": CountVerifikasi,
		"CountProcess":    CountProcess,
		"CountFinish":     CountFinish,
		"CountReject":     CountReject,
		"Blog":            blog,
		"notelp":          session.Values["notelp"]})

}

// func GetDetailsLaporan(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	db, err := database.Connect()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT * FROM `laporan` WHERE id = ?", vars["id"])
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	defer rows.Close()

// 	var result []model.Laporan

// 	for rows.Next() {
// 		var each = model.Laporan{}
// 		var err = rows.Scan(&each.Id, &each.Title, &each.Laporan, &each.User_id, &each.FullName, &each.Kategori, &each.Time, &each.Status)

// 		if err != nil {
// 			fmt.Println(err.Error())
// 			return
// 		}

// 		result = append(result, each)
// 	}

// 	pending := false
// 	process := false
// 	success := false

// 	if result[0].Status == "Sedang Di Verifikasi" {
// 		pending = true
// 	} else if result[0].Status == "Sedang Kami Lakukan Perbaikan" {
// 		process = true
// 	} else {
// 		success = true
// 	}
// 	goview.Render(w, http.StatusOK, "details.html", goview.M{
// 		"title":     result[0].Title,
// 		"laporan":   result[0].Laporan,
// 		"full_name": result[0].FullName,
// 		"kategori":  result[0].Kategori,
// 		"time":      result[0].Time,
// 		"status":    result[0].Status,
// 		"pending":   pending,
// 		"process":   process,
// 		"success":   success,
// 	})
// }
