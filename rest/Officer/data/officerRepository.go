package data

import (
	"database/sql"
	"day15/Officer/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type OfficerRepository struct {
	DB *sql.DB
}

// untuk nilai return get all butuh struktur dari pengarang
// 2.a buat model dari Pengarang

func GetAll(db *OfficerRepository) []models.Officer {
	fmt.Println(db.DB)

	result, err := db.DB.Query("Select * From tblOfficer")

	if err != nil {
		return nil
	}

	defer result.Close()
	fmt.Println(result)
	officer := []models.Officer{}
	for result.Next() {
		var p models.Officer
		if err := result.Scan(&p.TblOfficerID, &p.OfficerCode, &p.OfficerName, &p.OfficerPassword, &p.OfficerStatus); err != nil {
			return nil
		}
		officer = append(officer, p)
	}
	return officer
}
