package model

import "html/template"

type User struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Foto_profile string `json:"foto_profile"`

	Gender   string `json:"gender"`
	Role     string `json:"role"`
	Bidang   string `json:"bidang"`
	Alamat   string `json:"alamat"`
	Email    string `json:"email"`
	FullName string `json:"fullname"`
	NoTelp   string `json:"notelep"`
}

type Response struct {
	Status int `json:"status"`
	Data   []User
}
type Laporan struct {
	Id       int           `json:"id"`
	Title    string        `json:"Title"`
	Laporan  template.HTML `json:"laporan"`
	User_id  string        `json:"user_id"`
	Username  string        `json:"username"`
	User_Foto  string        `json:"user_foto"`
	Foto  string        `json:"Foto"`
	FullName string        `json:"full_name"`
	Kategori string        `json:"kategori"`
	Time     string        `json:"time"`
	Status   string        `json:"status"`
}

type Blog struct {
	Id       int           `json:"id"`
	Title    string        `json:"title"`
	Isi      template.HTML `json:"isi"`
	Kategori string        `json:"kategori"`
	Time     string        `json:"time"`
}

type Follow struct {
	Id       int           `json:"id"`
	User_id    string        `json:"user_id"`
	Target_id     string `json:"target_id"`
}

type KomentarLaporan struct {
	Id           int           `json:"id"`
	Id_laporan   string        `json:"id_laporan"`
	Isi          template.HTML `json:"isi"`
	Full_name    string        `json:"full_name"`
	Foto_profile string        `json:"foto_profile"`
	Id_user      string        `json:"Id_user"`
	Time         string        `json:"time"`
}
