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

func ViewStaff(w http.ResponseWriter, r *http.Request) {
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
	LaporanRow, err := db.Query("SELECT * FROM `laporan` WHERE Kategori = ?", session.Values["bidang"])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer LaporanRow.Close()

	var laporan []model.Laporan

	for LaporanRow.Next() {
		var each = model.Laporan{}
		var err = LaporanRow.Scan(&each.Id, &each.Title, &each.Laporan, &each.User_id, &each.FullName, &each.Kategori, &each.Time, &each.Status)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		laporan = append(laporan, each)
	}
	fmt.Println(laporan, session.Values["bidang"])
	// laporanCount := helper.GetCount("laporan")
	CountVerifikasi := helper.GetCountWhere("laporan", "Status", "Sedang Di Verifikasi")
	CountProcess := helper.GetCountWhere("laporan", "Status", "Sedang Process")
	CountFinish := helper.GetCountWhere("laporan", "Status", "Finish")
	CountReject := helper.GetCountWhere("laporan", "Status", "Reject")

	// laporanPersenFinish := CountFinish / laporanCount * 100
	goview.Render(w, http.StatusOK, "home_staff.html", goview.M{
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
		"notelp":          session.Values["notelp"]})

}
