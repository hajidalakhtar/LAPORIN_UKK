package query

import (
	"laporin_go/app/model"
	"laporin_go/database"

	"github.com/graphql-go/graphql"
)

func LaporanResolve(param graphql.ResolveParams) (interface{}, error) {
	id, success := param.Args["Laporan_id"].(int)
	if success {
		var a model.Laporan
		var b []model.Laporan
		db, err := database.Connect()
		if err != nil {
			panic(err.Error())
		}
		b = b[:0]
		result, err := db.Query("SELECT `Id`, `Title`, `Laporan`, `User_id`, `Username`, `User_Foto`, `Foto`, `FullName`, `Kategori`, `Time`, `Status` FROM `laporan` WHERE `Id` = ? ", id)
		if err != nil {
			panic(err.Error())
		}

		for result.Next() {
			err = result.Scan(&a.Id, &a.Title, &a.Laporan, &a.User_id, &a.Username, &a.User_Foto, &a.Foto, &a.FullName, &a.Kategori, &a.Time, &a.Status)
			if err != nil {
				panic(err.Error())
			}
			b = append(b, a)

		}

		return b, nil
	}
	return nil, nil
}
