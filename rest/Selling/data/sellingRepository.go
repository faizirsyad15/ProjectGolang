package data

import (
	"database/sql"
	"day15/Selling/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type SellingRepository struct {
	DB *sql.DB
}

// untuk nilai return get all butuh struktur dari pengarang
// 2.a buat model dari Pengarang

func GetAll(db *SellingRepository) []models.Selling {
	fmt.Println(db.DB)

	result, err := db.DB.Query("Select tblSellingID, Invoice From tblSelling")

	if err != nil {
		return nil
	}

	defer result.Close()
	fmt.Println(result)
	selling := []models.Selling{}
	for result.Next() {
		var p models.Selling
		if err := result.Scan(&p.TblSellingID, &p.Invoice); err != nil {
			return nil
		}
		selling = append(selling, p)
	}
	return selling
}
