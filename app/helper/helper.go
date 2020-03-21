package helper

import (
	"fmt"
	"laporin_go/database"
)

func GetCountWhere(tablename string, column string, value string) int {
	var cnt int
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	_ = db.QueryRow(`select count(*) from `+tablename+` where `+column+`=?`, value).Scan(&cnt)
	return cnt
}

func GetCountWhereLike(user_id string, id_laporan string) int {
	var cnt int
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	_ = db.QueryRow("select count(*) from `Like_laporan`  WHERE user_id = ? AND id_laporan = ? LIMIT 1", user_id, id_laporan).Scan(&cnt)
	return cnt
}

func GetCountWhereFollow(user_id string, target_id int) int {
	var cnt int
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	_ = db.QueryRow("select count(*) from `follow`  WHERE user_id = ? AND target_id = ? LIMIT 1", user_id, target_id).Scan(&cnt)
	return cnt
}

func GetCount(tablename string) int {
	var cnt int
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	_ = db.QueryRow(`select count(*) from ` + tablename + ``).Scan(&cnt)
	return cnt
}

func PostDelete(tablename string, where string) string {
	var cnt string
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	_, err = db.Exec(`DELETE FROM `+tablename+` WHERE Id = ?`, where)
	// _ = db.QueryRow(`select count(*) from `+tablename+` where id =?`, where).Scan(&cnt)
	cnt = "success"
	return cnt
}

func PostDeleteLike(tablename string, user_id string, id_laporan string) string {
	var cnt string
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	_, err = db.Exec(`DELETE FROM `+tablename+` WHERE user_id = `+user_id+` AND id_laporan = ?`, id_laporan)
	// _ = db.QueryRow(`select count(*) from `+tablename+` where id =?`, where).Scan(&cnt)
	cnt = "success"
	return cnt
}
