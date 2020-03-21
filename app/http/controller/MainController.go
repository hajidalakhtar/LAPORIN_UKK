package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"laporin_go/database"
	"net/http"
	"laporin_go/app/helper"
	"sort"

	"github.com/foolin/goview"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"

	"laporin_go/app/model"
)

var ctx = context.Background()

var store = sessions.NewCookieStore([]byte("haloo"))

func PostRegister(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	gender := r.FormValue("gender")
	alamat := r.FormValue("alamat")
	email := r.FormValue("email")
	fullname := r.FormValue("fullname")
	notelp := r.FormValue("notelp")
	role := "user"
	bidang := "user"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO `users`(`id`, `Username`, `Password`,`Foto_profile`, `Gender`, `Role`, `Bidang`,`Alamat`, `Email`, `FullName`, `NoTelp`) VALUES (?,?,?,?,?,?,?,?,?,?,?)", nil, username, hash,`default.png`, gender, role, bidang, alamat, email, fullname, notelp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// fmt.Println("insert success!")
	http.Redirect(w, r, "/login", 301)


}

func PostEditUser(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	username := r.FormValue("username")
	password := r.FormValue("password")
	gender := r.FormValue("gender")
	bidang := r.FormValue("bidang")
	alamat := r.FormValue("alamat")
	email := r.FormValue("email")
	fullname := r.FormValue("fullname")
	notelp := r.FormValue("notelp")
	role := "user"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	_, err = db.Exec("UPDATE `users` SET `id`=?,`Username`=?,`Password`=?,`Foto_profile`=?,`Gender`=?,`Role`=?,`Bidang`=?,`Alamat`=?,`Email`=?`FullName`=?,`NoTelp`=? WHERE id = ?",nil,username,hash,"default.png",gender,role,bidang,alamat,email,fullname,notelp,id)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	http.Redirect(w, r, r.Header.Get("Referer"), 302)

}

func ViewProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
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

// get user
	user, err := db.Query("SELECT * FROM `users` WHERE Username = ?", vars["username"])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer user.Close()

	var users []model.User

	for user.Next() {
		var each = model.User{}
		var err = user.Scan(&each.Id, &each.Username, &each.Password, &each.Foto_profile, &each.Gender, &each.Role, &each.Bidang, &each.Alamat, &each.Email, &each.FullName, &each.NoTelp)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		users = append(users, each)
	}

// get laporan
	var laporan []model.Laporan
	lapor, err := db.Query("SELECT * FROM `laporan` WHERE User_id = ?", users[0].Id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer lapor.Close()


	for lapor.Next() {
		var each = model.Laporan{}
		var err = lapor.Scan(&each.Id, &each.Title, &each.Laporan, &each.User_id,&each.Username,&each.User_Foto,&each.Foto, &each.FullName, &each.Kategori, &each.Time, &each.Status)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		laporan = append(laporan, each)
	}
	fmt.Println(laporan)
	session_user_id := fmt.Sprintf("%v", session.Values["user_id"])
	// json.NewEncoder(w).Encode(users)
	isFollow := helper.GetCountWhereFollow(session_user_id,users[0].Id)
	fmt.Println(isFollow)


	// goview.Render(w, http.StatusOK, "user_profile.html", goview.M{"user": users})
	goview.Render(w, http.StatusOK, "user_profile.html", goview.M{
		"user_id":  users[0].Id,
		"session_user_id":	session_user_id,
		"username": users[0].Username,
		"isFollow":isFollow,
		"password": users[0].Password,
		"gender":   users[0].Gender,
		"role":     users[0].Role,
		"alamat":   users[0].Alamat,
		"email":    users[0].Email,
		"fullname": users[0].FullName,
		"notelp":   users[0].NoTelp,
		"laporan":  laporan,
	})

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from users")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []model.User

	for rows.Next() {
		var each = model.User{}
		var err = rows.Scan(&each.Id, &each.Username, &each.Password, &each.Foto_profile, &each.Gender, &each.Role, &each.Bidang, &each.Alamat, &each.Email, &each.FullName, &each.NoTelp)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}
	json.NewEncoder(w).Encode(result)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	username := r.FormValue("username")
	password := r.FormValue("password")
	session, err := store.Get(r, "login")
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM `users` WHERE Username = ?", username)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	var result []model.User
	for rows.Next() {
		var each = model.User{}
		var err = rows.Scan(&each.Id, &each.Username, &each.Password, &each.Foto_profile, &each.Gender, &each.Role, &each.Bidang, &each.Alamat, &each.Email, &each.FullName, &each.NoTelp)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		result = append(result, each)
	}
	match := CheckPasswordHash(password, result[0].Password)
	session.Values["user_id"] = result[0].Id
	session.Values["username"] = result[0].Username
	session.Values["password"] = result[0].Password
	session.Values["gender"] = result[0].Gender
	session.Values["user_foto"] = result[0].Foto_profile
	session.Values["role"] = result[0].Role
	session.Values["bidang"] = result[0].Bidang
	session.Values["alamat"] = result[0].Alamat
	session.Values["email"] = result[0].Email
	session.Values["fullname"] = result[0].FullName
	session.Values["notelp"] = result[0].NoTelp
	err = session.Save(r, w)
	// fmt.Println(session.Values["bidang"])
	if result[0].Role == "user" {
		if match {
			http.Redirect(w, r, "/user", 301)

		} else {
			http.Redirect(w, r, "/login", 301)

		}
	} else if result[0].Role == "admin" {
		http.Redirect(w, r, "/admin", 301)

	} else {
		http.Redirect(w, r, "/staff", 301)

	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// json.NewEncoder(w).Encode(response)
}

func PostLogout(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")

	session, _ := store.Get(r, "login")
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/home", http.StatusSeeOther)

	// json.NewEncoder(w).Encode(response)
}

func ViewLandingPage(w http.ResponseWriter, r *http.Request) {
	err := goview.Render(w, http.StatusOK, "landingpage.html", goview.M{
		"title": "Page file title!!"})

	if err != nil {
		fmt.Fprintf(w, "Render page.html error: %v!", err)
	}
}

func ViewHome(w http.ResponseWriter, r *http.Request) {
	session, errsession := store.Get(r, "login")

	user_id := session.Values["user_id"]
	if user_id == nil {
		http.Redirect(w, r, "/home", 301)

	}
	if errsession != nil {
		fmt.Println("error")
	}
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

// get Follow
follow, err := db.Query("SELECT * FROM `follow` WHERE user_id = ?", user_id)
if err != nil {
	fmt.Println(err.Error())
	return
}
defer follow.Close()

var follows []model.Follow

for follow.Next() {
	var each = model.Follow{}
	var err = follow.Scan(&each.Id, &each.User_id, &each.Target_id)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	follows = append(follows, each)
}
	var result []model.Laporan

	rows, err := db.Query("SELECT * FROM `laporan` WHERE User_id = ?", user_id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()


	for rows.Next() {
		var each = model.Laporan{}
		var err = rows.Scan(&each.Id, &each.Title, &each.Laporan, &each.User_id,&each.Username,&each.User_Foto,&each.Foto, &each.FullName, &each.Kategori, &each.Time, &each.Status)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}


	for i := 0;i < len(follows);i++ {
		rows, err := db.Query("SELECT * FROM `laporan` WHERE User_id = ?", follows[i].Target_id)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer rows.Close()


		for rows.Next() {
			var each = model.Laporan{}
			var err = rows.Scan(&each.Id, &each.Title, &each.Laporan, &each.User_id,&each.Username,&each.User_Foto,&each.Foto, &each.FullName, &each.Kategori, &each.Time, &each.Status)

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			result = append(result, each)
		}
	}
	// laporan
	defer db.Close()

	sort.SliceStable(result, func(i, j int) bool {
	    return result[j].Id < result[i].Id
	})

	goview.Render(w, http.StatusOK, "home.html", goview.M{"username": session.Values["username"],
		"laporan":  result,
		"role":     session.Values["role"],
		"alamat":   session.Values["alamat"],
		"email":    session.Values["email"],
		"fullname": session.Values["fullname"],
		"notelp":   session.Values["notelp"]})

}

func ViewLogin(w http.ResponseWriter, r *http.Request) {
	err := goview.Render(w, http.StatusOK, "login.html", goview.M{"title": "Page file title!!"})
	if err != nil {
		fmt.Fprintf(w, "Render page.html error: %v!", err)
	}
}

func ViewCari(w http.ResponseWriter, r *http.Request) {
	err := goview.Render(w, http.StatusOK, "cari.html", goview.M{"title": "Page file title!!"})
	if err != nil {
		fmt.Fprintf(w, "Render page.html error: %v!", err)
	}
}
func ViewAdminUser(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM `users`")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	var result []model.User
	for rows.Next() {
		var each = model.User{}
		var err = rows.Scan(&each.Id, &each.Username, &each.Password, &each.Foto_profile, &each.Gender, &each.Role, &each.Bidang, &each.Alamat, &each.Email, &each.FullName, &each.NoTelp)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		result = append(result, each)
	}

	goview.Render(w, http.StatusOK, "admin_user.html", goview.M{
		"user": result,
	})
	if err != nil {
		fmt.Fprintf(w, "Render page.html error: %v!", err)
	}
}
func ViewRegister(w http.ResponseWriter, r *http.Request) {
	err := goview.Render(w, http.StatusOK, "register.html", goview.M{"title": "Page file title!!"})
	if err != nil {
		fmt.Fprintf(w, "Render page.html error: %v!", err)
	}
}

func PostLike(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)


	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO `Like_laporan`(`id`, `user_id`, `id_laporan`) VALUES (?,?,?)", nil, vars["user_id"], vars["id"])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// fmt.Println(vars["user_id"], vars["id"])

	http.Redirect(w, r, r.Header.Get("Referer"), 302)

}

func PostUnlike(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = db.Exec("DELETE FROM `Like_laporan` WHERE user_id = ? AND id_laporan = ? LIMIT 1", vars["user_id"], vars["id"])
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	http.Redirect(w, r, r.Header.Get("Referer"), 302)

}

func PostFollow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

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

	_, err = db.Exec("INSERT INTO `follow`(`id`, `user_id`, `target_id`) VALUES (?,?,?)", nil, session.Values["user_id"], vars["id"])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("success")
	http.Redirect(w, r, r.Header.Get("Referer"), 302)

}


func PostUnfollow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = db.Exec("DELETE FROM `follow` WHERE target_id = ? AND user_id = ? LIMIT 1",  vars["target_id"], vars["user_id"])
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
