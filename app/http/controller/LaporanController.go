package controller

import (
	"fmt"
	"laporin_go/database"
	"log"
	"net/http"
	"time"

	"github.com/foolin/goview"
	"github.com/gorilla/mux"

	"laporin_go/app/helper"
	"laporin_go/app/model"
)

// var ctx = context.Background()

func PostLaporan(w http.ResponseWriter, r *http.Request) {

	laporan := r.FormValue("laporan")
	kategori := r.FormValue("kategori")
	title := r.FormValue("title")
	session, err := store.Get(r, "login")
	date := time.Now()
	user_id := session.Values["user_id"]
	username := session.Values["username"]
	user_foto := session.Values["user_foto"]
	full_name := session.Values["fullname"]
	fmt.Println(full_name)

	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO `laporan`(`Id`,`Title`, `Laporan`, `User_id`,`Username`,`User_Foto`,`Foto`,`FullName`, `Kategori`, `Time`, `Status`) VALUES (?,?,?,?,?,?,?,?,?,?,?)", nil, title, laporan, user_id, username,user_foto,nil,full_name, kategori, date.Format("02-Jan-2006"), "Sedang Di Verifikasi")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if err != nil {
		log.Fatal(err.Error())
	}

	http.Redirect(w, r, "/user", 303)

}

func GetDetailsLaporan(w http.ResponseWriter, r *http.Request) {
	session, errsession := store.Get(r, "login")
	if errsession != nil {
		fmt.Println("error")
	}
	vars := mux.Vars(r)

	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM `laporan` WHERE id = ?", vars["id"])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []model.Laporan

	for rows.Next() {
		var each = model.Laporan{}
		var err = rows.Scan(&each.Id, &each.Title, &each.Laporan, &each.User_id,&each.Username,&each.User_Foto,&each.Foto, &each.FullName, &each.Kategori, &each.Time, &each.Status)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	komentar, err := db.Query("SELECT * FROM `komentarLaporan` WHERE Id_laporan = ?", vars["id"])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer komentar.Close()

	var resultkomentar []model.KomentarLaporan

	for komentar.Next() {
		var each = model.KomentarLaporan{}
		var err = komentar.Scan(&each.Id, &each.Id_laporan, &each.Isi, &each.Full_name, &each.Foto_profile, &each.Id_user, &each.Time)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		resultkomentar = append(resultkomentar, each)
	}

	pending := false
	process := false
	success := false

	if result[0].Status == "Sedang Di Verifikasi" {
		pending = true
	} else if result[0].Status == "Sedang Process" {
		process = true
	} else {
		success = true
	}

	CountKomentar := helper.GetCountWhere("komentarLaporan", "id_laporan", vars["id"])
	CountLike := helper.GetCountWhere("Like_laporan", "id_laporan", vars["id"])
	user_id := fmt.Sprintf("%v", session.Values["user_id"])

	likeornot := helper.GetCountWhereLike(user_id, vars["id"])
	fmt.Println(likeornot, user_id, vars["id"])


	goview.Render(w, http.StatusOK, "details.html", goview.M{
		"title":         result[0].Title,
		"laporan":       result[0].Laporan,
		"full_name":     result[0].FullName,
		"kategori":      result[0].Kategori,
		"time":          result[0].Time,
		"status":        result[0].Status,
		"komentar":      resultkomentar,
		"laporan_id":    vars["id"],
		"pending":       pending,
		"CountKomentar": CountKomentar,
		"CountLike":     CountLike,
		"user_id":       user_id,
		"likeornot":     likeornot,
		"process":       process,
		"success":       success,
	})

}
func PostGantiStatus(w http.ResponseWriter, r *http.Request) {

	id_laporan := r.FormValue("id_laporan")
	status := r.FormValue("status")
	// title := r.FormValue("title")
	fmt.Println(id_laporan, status)
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE `laporan` SET `Status`= ? WHERE Id = ?", status, id_laporan)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if err != nil {
		log.Fatal(err.Error())
	}

	// requestBody, err := json.Marshal(map[string]string{
	// 	"phone": "6287883594486",
	// 	"body":  `Laporan Anda Dengan id ` + id_laporan + ` sekarang dalam tahap ` + status + ``,
	// })
	// timeout := time.Duration(5 * time.Second)
	// client := http.Client{Timeout: timeout}
	// request, err := http.NewRequest("POST", "https://eu4.chat-api.com/instance107376/sendMessage?token=tbeomej7t1jxjkbb", bytes.NewBuffer(requestBody))
	// request.Header.Set("Content-type", "application/json")
	// resp, err := client.Do(request)
	// defer resp.Body.Close()
	http.Redirect(w, r, r.Header.Get("Referer"), 302)

}

func GetDeleteLaporan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	delete := helper.PostDelete("laporan", id)
	fmt.Println(delete)
	http.Redirect(w, r, r.Header.Get("Referer"), 302)

}

func PostTambahKomentar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	isi := r.FormValue("isi")
	// title := r.FormValue("title")
	session, err := store.Get(r, "login")
	date := time.Now()
	userid := session.Values["user_id"]
	fullname := session.Values["fullname"]

	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO `komentarLaporan`(`id`, `Id_laporan`, `Isi`, `Full_name`, `Foto_profile`, `Id_user`, `Time`) VALUES (?,?,?,?,?,?,?)", nil, id, isi, fullname, "poto.png", userid, date.Format("02-Jan-2006"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if err != nil {
		log.Fatal(err.Error())
	}
	http.Redirect(w, r, r.Header.Get("Referer"), 302)

}
